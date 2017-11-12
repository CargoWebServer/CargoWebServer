package JS

import (
	"bytes"
	"errors"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"code.myceliUs.com/Utility"
	"github.com/robertkrimen/otto"
	_ "github.com/robertkrimen/otto/underscore"
)

/////////////////////////////////////////////////////////////////////////////
// Structures used by channel.
/////////////////////////////////////////////////////////////////////////////

/**
 * That structure is use to manipulate Javacsript function call infromations.
 */
type JsFunctionInfos struct {
	// Parameters...
	m_messageId      string
	m_sessionId      string
	m_functionStr    string
	m_functionParams []interface{}

	// Result part...
	m_results []interface{}
	m_err     error
}

/**
 * That structure is use to put variable information inside channels.
 */
type JsVarInfos struct {
	m_name string
	m_val  interface{}
}

/**
 * Wrap operation information...
 */
type OperationInfos struct {
	m_params  map[string]interface{}
	m_returns chan ([]interface{})
}

/**
 * Operation chan
 */
type OperationChannel chan (OperationInfos)

/**
 * Channel to use to access session.
 */
type SessionInfos struct {
	m_sessionId string
	m_return    chan (*otto.Otto)
}

/**
 * Channel use to set funtion.
 */
type FunctionInfos struct {
	m_functionId string
	m_function   interface{}
}

var (
	jsRuntimeManager *JsRuntimeManager
)

/**
 * The javascript runtime is use to extent the functionnality of the system
 * via javascript...
 */
type JsRuntimeManager struct {
	/** The directory where the script is store (recursive) **/
	m_searchDir string

	/** Contain all script. **/
	m_scripts []string

	/** Each connection has it own VM. */
	m_sessions map[string]*otto.Otto

	/** Go function interfaced in JS. **/
	m_functions map[string]interface{}

	/**Channel use for communication with vm... **/

	// Script channels
	m_appendScript OperationChannel
	m_initScripts  OperationChannel

	// Open/Close channel.
	m_createVm OperationChannel
	m_closeVm  OperationChannel

	// JS channel
	m_setFunction       chan (FunctionInfos)
	m_getSession        chan (SessionInfos)
	m_setVariable       map[string]OperationChannel
	m_getVariable       map[string]OperationChannel
	m_executeJsFunction map[string]OperationChannel
	m_runScript         map[string]OperationChannel
	m_stopVm            map[string]chan (bool)
}

func NewJsRuntimeManager(searchDir string) *JsRuntimeManager {
	// The singleton.
	jsRuntimeManager = new(JsRuntimeManager)

	// The dir where the js file are store on the server side.
	jsRuntimeManager.m_searchDir = searchDir

	// List of vm one per connection.
	jsRuntimeManager.m_sessions = make(map[string]*otto.Otto)

	// List of functions.
	jsRuntimeManager.m_functions = make(map[string]interface{})

	// Initialisation of channel's.
	jsRuntimeManager.m_appendScript = make(OperationChannel)
	jsRuntimeManager.m_initScripts = make(OperationChannel)
	jsRuntimeManager.m_closeVm = make(OperationChannel)
	jsRuntimeManager.m_createVm = make(OperationChannel)

	// Create one channel by vm to set variable.
	jsRuntimeManager.m_setFunction = make(chan (FunctionInfos))
	jsRuntimeManager.m_getSession = make(chan (SessionInfos))
	jsRuntimeManager.m_setVariable = make(map[string]OperationChannel)
	jsRuntimeManager.m_getVariable = make(map[string]OperationChannel)
	jsRuntimeManager.m_executeJsFunction = make(map[string]OperationChannel)
	jsRuntimeManager.m_runScript = make(map[string]OperationChannel)
	jsRuntimeManager.m_stopVm = make(map[string]chan (bool))

	// Load the script from the script repository...
	jsRuntimeManager.appendScriptFiles()

	// That function will process api call's
	go func(jsRuntimeManager *JsRuntimeManager) {
		for {
			select {
			case functionInfo := <-jsRuntimeManager.m_setFunction:
				jsRuntimeManager.m_functions[functionInfo.m_functionId] = functionInfo.m_function
			case sessionInfo := <-jsRuntimeManager.m_getSession:
				sessionInfo.m_return <- jsRuntimeManager.m_sessions[sessionInfo.m_sessionId]
			case operationInfos := <-jsRuntimeManager.m_appendScript:
				callback := operationInfos.m_returns
				script := operationInfos.m_params["script"].(string)
				jsRuntimeManager.appendScript(script)
				callback <- []interface{}{true} // unblock the channel...
			case operationInfos := <-jsRuntimeManager.m_initScripts:
				callback := operationInfos.m_returns
				sessionId := operationInfos.m_params["sessionId"].(string)
				jsRuntimeManager.initScripts(sessionId)
				callback <- []interface{}{true} // unblock the channel...
			case operationInfos := <-jsRuntimeManager.m_createVm:
				callback := operationInfos.m_returns
				sessionId := operationInfos.m_params["sessionId"].(string)
				jsRuntimeManager.createVm(sessionId)
				jsRuntimeManager.m_setVariable[sessionId] = make(OperationChannel)
				jsRuntimeManager.m_getVariable[sessionId] = make(OperationChannel)
				jsRuntimeManager.m_executeJsFunction[sessionId] = make(OperationChannel)
				jsRuntimeManager.m_runScript[sessionId] = make(OperationChannel)
				jsRuntimeManager.m_stopVm[sessionId] = make(chan (bool))
				// The vm processing loop...
				go func(vm *otto.Otto, setVariable OperationChannel, getVariable OperationChannel, executeJsFunction OperationChannel, runScript OperationChannel, stopVm chan (bool), sessionId string) {
					done := false
					for !done {
						select {
						case operationInfos := <-setVariable:
							callback := operationInfos.m_returns
							varInfos := operationInfos.m_params["varInfos"].(JsVarInfos)
							vm.Set(varInfos.m_name, varInfos.m_val)
							callback <- []interface{}{true} // unblock the channel...
						case operationInfos := <-getVariable:
							callback := operationInfos.m_returns
							varInfos := operationInfos.m_params["varInfos"].(JsVarInfos)
							value, err := vm.Get(varInfos.m_name)
							if err == nil {
								varInfos.m_val = value
							}
							callback <- []interface{}{varInfos} // unblock the channel...
						case operationInfos := <-executeJsFunction:
							callback := operationInfos.m_returns
							jsFunctionInfos := operationInfos.m_params["jsFunctionInfos"].(JsFunctionInfos)
							vm.Set("messageId", jsFunctionInfos.m_messageId)
							vm.Set("sessionId", jsFunctionInfos.m_sessionId)
							jsFunctionInfos.m_results, jsFunctionInfos.m_err = GetJsRuntimeManager().executeJsFunction(vm, jsFunctionInfos.m_functionStr, jsFunctionInfos.m_functionParams)
							callback <- []interface{}{jsFunctionInfos} // unblock the channel...
						case operationInfos := <-runScript:
							callback := operationInfos.m_returns
							script := operationInfos.m_params["script"].(string)
							results, err := vm.Run(script)
							callback <- []interface{}{results, err} // unblock the channel...
						case stop := <-stopVm:
							if stop {
								// Wait until the vm is stop
								wait := make(chan string, 1)
								timer := time.NewTimer(5 * time.Second)

								// Call the interupt function on the VM.
								vm.Interrupt <- func(sessionId string, wait chan string, timer *time.Timer) func() {
									return func() {
										timer.Stop()
										// Continue the processing.
										wait <- "--> Interrupt execution of VM with id " + sessionId
										// The panic error will actually kill the vm
										panic(errors.New("Stahp"))
									}
								}(sessionId, wait, timer)

								// If nothing append for 5 second I will return.
								go func(wait chan string, timer *time.Timer) {
									<-timer.C
									wait <- "--> Stop execution of VM with id " + sessionId
								}(wait, timer)

								// Synchronyse exec of interuption here.
								log.Println(<-wait)
								done = true
								break // exit the loop.
							}
						}
					}
					jsRuntimeManager.removeVm(sessionId)
				}(jsRuntimeManager.m_sessions[sessionId], jsRuntimeManager.m_setVariable[sessionId], jsRuntimeManager.m_getVariable[sessionId], jsRuntimeManager.m_executeJsFunction[sessionId], jsRuntimeManager.m_runScript[sessionId], jsRuntimeManager.m_stopVm[sessionId], sessionId)
				callback <- []interface{}{true} // unblock the channel...

			case operationInfos := <-jsRuntimeManager.m_closeVm:
				callback := operationInfos.m_returns
				sessionId := operationInfos.m_params["sessionId"].(string)
				jsRuntimeManager.m_stopVm[sessionId] <- true // send kill
				callback <- []interface{}{true}              // unblock the channel...
			}
		}
	}(jsRuntimeManager)

	return jsRuntimeManager
}

func GetJsRuntimeManager() *JsRuntimeManager {
	return jsRuntimeManager
}

/////////////////////////////////////////////////////////////////////////////
// Initialisation relatead functions
/////////////////////////////////////////////////////////////////////////////

/** Append all scripts **/
func (this *JsRuntimeManager) appendScriptFiles() error {

	fileList := []string{}
	err := filepath.Walk(this.m_searchDir, func(path string, f os.FileInfo, err error) error {
		fileList = append(fileList, path)
		return nil
	})

	for _, file := range fileList {
		// Only .js are valid extension here...
		if strings.HasSuffix(file, ".js") {
			err = this.appendScriptFile(file)
			if err != nil {
				return err
			}
		}
	}
	return err
}

/**
 * Compile and run a given script...
 */
func (this *JsRuntimeManager) appendScriptFile(filePath string) error {
	log.Println("--> Load script file ", filePath)
	srcFile, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Read JS src file:", err)
	}

	// Close the file when is no more needed...
	defer srcFile.Close()
	buf := bytes.NewBuffer(nil)
	io.Copy(buf, srcFile)
	src := string(buf.Bytes())
	this.appendScript(src)

	return err
}

/**
 * Append a script and initialyse all vm with it.
 */
func (this *JsRuntimeManager) appendScript(src string) {

	if Utility.Contains(this.m_scripts, src) == false {
		this.m_scripts = append(this.m_scripts, src)
	}

	// I will compile the script and set it in each session...
	for scriptName, vm := range this.m_sessions {
		script, err := vm.Compile("", src)
		if err != nil {
			log.Println("runtime script compilation error:", scriptName, err)
		} else {
			vm.Run(script)
		}
	}
}

/**
 * Initialisation of script for a newly created session.
 */
func (this *JsRuntimeManager) initScripts(sessionId string) {
	// Get the vm.
	vm := this.m_sessions[sessionId]

	// Compile the list of script...
	for i := 0; i < len(this.m_scripts); i++ {
		script, err := vm.Compile("", this.m_scripts[i])
		if err == nil {
			vm.Run(script)
			if err != nil {
				log.Println("runtime script compilation error:", script, err)
			} else {
				//log.Println(sessionId, " Load: ", this.m_scripts[i])
			}

		} else {
			log.Println("-------> error in script: ", this.m_scripts[i])
			log.Println(err)
		}
	}

	// Set the list of binded function.
	for name, function := range this.m_functions {
		vm.Set(name, function)
	}
}

//////////////////////////////////////////////////////////////////////////////
// VM
//////////////////////////////////////////////////////////////////////////////

/**
 *  Create a new VM for a given session id.
 */
func (this *JsRuntimeManager) createVm(sessionId string) {
	if this.m_sessions[sessionId] != nil {
		return // Nothing to do if the session already exist.
	}
	// Create a new js interpreter for the given session.
	if sessionId == "" {
		this.m_sessions[sessionId] = otto.New()
	} else {
		// The runtime will be the base cargo runtime for each session.
		this.m_sessions[sessionId] = this.m_sessions[""].Copy()
	}

	// That channel is use to interrupt vm machine, it must be created before
	// the vm start.
	this.m_sessions[sessionId].Interrupt = make(chan func(), 1) // The buffer prevents blocking

	// Put the sessionId variable on the global scope.
	this.m_sessions[sessionId].Set("sessionId", sessionId)
}

/**
 *  Close and remove VM for a given session id.
 */
func (this *JsRuntimeManager) removeVm(sessionId string) {
	// Remove vm ressources.
	delete(this.m_sessions, sessionId)
	close(this.m_executeJsFunction[sessionId])
	delete(this.m_executeJsFunction, sessionId)
	close(this.m_runScript[sessionId])
	delete(this.m_runScript, sessionId)
	close(this.m_setVariable[sessionId])
	delete(this.m_setVariable, sessionId)
	close(this.m_stopVm[sessionId])
	delete(this.m_stopVm, sessionId)

}

/**
 * Execute javascript function.
 */
func (this *JsRuntimeManager) executeJsFunction(vm *otto.Otto, functionStr string, functionParams []interface{}) (results []interface{}, err error) {

	if len(functionStr) == 0 {
		return nil, errors.New("No function string.")
	}

	// Here i wil find the name of the function...
	startIndex := strings.Index(functionStr, "function")
	if startIndex != -1 {
		startIndex += len("function")
	} else {
		startIndex = 0
	}

	endIndex := strings.Index(functionStr, "(")

	var functionName string

	// Remove withe space.
	if endIndex != -1 {
		functionName = strings.TrimSpace(functionStr[startIndex:endIndex])
		if len(functionName) == 0 {
			functionName = "fct_" + strings.Replace(Utility.GenerateUUID(functionStr), "-", "_", -1)
			functionStr = "function " + functionName + strings.TrimSpace(functionStr)[8:]
		}
		script, err := vm.Compile("", functionStr)
		if err == nil {
			_, err = vm.Run(script)
			if err != nil {
				log.Println("fail to run ", functionStr)
				return nil, err
			}
		} else {
			log.Println("fail to compile  ", functionStr)
			return nil, err
		}
	} else {
		functionName = functionStr
	}

	var params []interface{}
	params = append(params, functionName)
	params = append(params, nil)

	// Now I will make otto digest the parameters...
	for i := 0; i < len(functionParams); i++ {
		p, err := vm.ToValue(functionParams[i])
		if err != nil {
			log.Println("Error binding parameter", err)
			params = append(params, nil)
		} else {
			params = append(params, p)
		}
	}

	// Call the call...
	result, err_ := Utility.CallMethod(vm, "Call", params)

	if err_ != nil {
		log.Println("Error found with function ", functionName, err_.(error), "params: ", params)
		log.Println("Src ", functionStr)
		return nil, err_.(error)
	}

	// Return the result if there is one...
	val, err := result.(otto.Value).Export()
	if err != nil {
		log.Panicln("---------> error ", err)
		return nil, err
	}

	// Append val to results...
	results = append(results, val)

	return
}

/**
 * Run given script for a given session.
 */
func (this *JsRuntimeManager) getSession(sessionId string) *otto.Otto {
	// Protectect the map access...
	var sessionInfo SessionInfos
	sessionInfo.m_return = make(chan (*otto.Otto))
	defer close(sessionInfo.m_return)
	sessionInfo.m_sessionId = sessionId
	this.m_getSession <- sessionInfo
	session := <-sessionInfo.m_return
	return session
}

//////////////////////////////////////////////////////////////////////////////
// Api
//////////////////////////////////////////////////////////////////////////////

/**
 * Append function in the JS, must be call before the js is started.
 */
func (this *JsRuntimeManager) AppendFunction(name string, function interface{}) {
	// I will compile the script and set it in each session...
	var functionInfo FunctionInfos
	functionInfo.m_functionId = name
	functionInfo.m_function = function
	// Set the function via the channel
	this.m_setFunction <- functionInfo
}

/**
 * Run given script for a given session.
 */
func (this *JsRuntimeManager) RunScript(sessionId string, script string) (otto.Value, error) {
	// Protectect the map access...
	var op OperationInfos
	op.m_params = make(map[string]interface{})
	op.m_params["script"] = script
	op.m_returns = make(chan ([]interface{}))
	defer close(op.m_returns)
	this.m_runScript[sessionId] <- op
	// wait for completion
	results := <-op.m_returns
	var value otto.Value
	var err error
	if results[0] != nil {
		value = results[0].(otto.Value)
	}

	if results[1] != nil {
		err = results[0].(error)
	}

	return value, err
}

/**
 * Append a script to all VM's
 */
func (this *JsRuntimeManager) AppendScript(script string) {

	var op OperationInfos
	op.m_params = make(map[string]interface{})
	op.m_params["script"] = script
	op.m_returns = make(chan ([]interface{}))
	defer close(op.m_returns)
	this.m_appendScript <- op

	// wait for completion
	<-op.m_returns
}

func (this *JsRuntimeManager) InitScripts(sessionId string) {

	var op OperationInfos
	op.m_params = make(map[string]interface{})
	op.m_params["sessionId"] = sessionId
	op.m_returns = make(chan ([]interface{}))
	defer close(op.m_returns)
	this.m_initScripts <- op

	// wait for completion
	<-op.m_returns
}

/**
 * Append and excute a javacript function on the JS...
 */
func (this *JsRuntimeManager) ExecuteJsFunction(messageId string, sessionId string, functionStr string, functionParams []interface{}) ([]interface{}, error) {
	// Set the function on the JS runtime...
	if this.m_executeJsFunction[sessionId] == nil {
		return nil, errors.New("Session " + sessionId + " is closed!")
	}

	// Put function call information into a struct.
	var jsFunctionInfos JsFunctionInfos
	jsFunctionInfos.m_messageId = messageId
	jsFunctionInfos.m_sessionId = sessionId
	jsFunctionInfos.m_functionStr = functionStr
	jsFunctionInfos.m_functionParams = functionParams

	var op OperationInfos
	op.m_params = make(map[string]interface{})
	op.m_params["jsFunctionInfos"] = jsFunctionInfos
	op.m_returns = make(chan ([]interface{}))
	defer close(op.m_returns)
	this.m_executeJsFunction[sessionId] <- op
	// wait for completion
	results := <-op.m_returns
	return results[0].(JsFunctionInfos).m_results, results[0].(JsFunctionInfos).m_err

}

/**
 * Set variable value for a given session
 */
func (this *JsRuntimeManager) SetVar(sessionId string, name string, val interface{}) {
	if this.m_setVariable[sessionId] == nil {
		return
	}
	// Protectect the map access...
	var info JsVarInfos
	info.m_name = name
	info.m_val = val

	var op OperationInfos
	op.m_params = make(map[string]interface{})
	op.m_params["varInfos"] = info
	op.m_returns = make(chan ([]interface{}))
	defer close(op.m_returns)
	this.m_setVariable[sessionId] <- op
	// wait for completion
	<-op.m_returns

}

/**
 * Run given script for a given session.
 */
func (this *JsRuntimeManager) GetVar(sessionId string, name string) interface{} {
	// Nothing to do with a close channel
	if this.m_setVariable[sessionId] == nil {
		return nil
	}

	// Protectect the map access...
	var info JsVarInfos
	info.m_name = name

	var op OperationInfos
	op.m_params = make(map[string]interface{})
	op.m_params["varInfos"] = info
	op.m_returns = make(chan ([]interface{}))
	defer close(op.m_returns)
	this.m_setVariable[sessionId] <- op

	// wait for completion
	results := <-op.m_returns
	return results[0].(JsVarInfos).m_val

}

/**
 * Open a new session on the server.
 */
func (this *JsRuntimeManager) OpenSession(sessionId string) {

	var op OperationInfos
	op.m_params = make(map[string]interface{})
	op.m_params["sessionId"] = sessionId
	op.m_returns = make(chan ([]interface{}))
	this.m_createVm <- op
	defer close(op.m_returns)
	// wait for completion
	<-op.m_returns
	log.Println("--> vm with id", sessionId, "is now open!")
}

/**
 * Close a given session.
 */
func (this *JsRuntimeManager) CloseSession(sessionId string, callback func()) {
	// Send message to stop vm...
	var op OperationInfos
	op.m_params = make(map[string]interface{})
	op.m_params["sessionId"] = sessionId
	op.m_returns = make(chan ([]interface{}))
	this.m_closeVm <- op
	defer close(op.m_returns)
	// wait for completion
	<-op.m_returns

	// Call the close callback
	callback()
}
