package JS

import (
	"bytes"
	"errors"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

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
	m_setVariable       map[string]OperationChannel
	m_executeJsFunction map[string]OperationChannel
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
	jsRuntimeManager.m_setVariable = make(map[string]OperationChannel)
	jsRuntimeManager.m_executeJsFunction = make(map[string]OperationChannel)
	jsRuntimeManager.m_stopVm = make(map[string]chan (bool))

	// Load the script from the script repository...
	jsRuntimeManager.appendScriptFiles()

	// That function will process api call's
	go func(jsRuntimeManager *JsRuntimeManager) {
		for {
			select {
			case operationInfos := <-jsRuntimeManager.m_appendScript:
				callback := operationInfos.m_returns
				script := operationInfos.m_params["script"].(string)
				jsRuntimeManager.appendScript(script)
				callback <- []interface{}{true} // unblock the channel...
			case operationInfos := <-jsRuntimeManager.m_initScripts:
				callback := operationInfos.m_returns
				sessionId := operationInfos.m_params["sessionId"].(string)
				log.Println("-----> init script: ", sessionId)
				jsRuntimeManager.initScripts(sessionId)
				callback <- []interface{}{true} // unblock the channel...
			case operationInfos := <-jsRuntimeManager.m_createVm:
				callback := operationInfos.m_returns
				sessionId := operationInfos.m_params["sessionId"].(string)
				jsRuntimeManager.createVm(sessionId)
				jsRuntimeManager.m_setVariable[sessionId] = make(OperationChannel)
				jsRuntimeManager.m_executeJsFunction[sessionId] = make(OperationChannel)
				jsRuntimeManager.m_stopVm[sessionId] = make(chan (bool))
				// The vm processing loop...
				go func(vm *otto.Otto, setVariable OperationChannel, executeJsFunction OperationChannel, stopVm chan (bool)) {
					for {
						select {
						case operationInfos := <-setVariable:
							callback := operationInfos.m_returns
							varInfos := operationInfos.m_params["varInfos"].(JsVarInfos)
							vm.Set(varInfos.m_name, varInfos.m_val)
							callback <- []interface{}{true} // unblock the channel...
						case operationInfos := <-executeJsFunction:
							callback := operationInfos.m_returns
							jsFunctionInfos := operationInfos.m_params["jsFunctionInfos"].(JsFunctionInfos)
							vm.Set("messageId", jsFunctionInfos.m_messageId)
							jsFunctionInfos.m_results, jsFunctionInfos.m_err = GetJsRuntimeManager().executeJsFunction(vm, jsFunctionInfos.m_functionStr, jsFunctionInfos.m_functionParams)
							callback <- []interface{}{jsFunctionInfos} // unblock the channel...
						case stop := <-stopVm:
							// clear ressource here...
							if stop {
								log.Println("-----> vm: ", sessionId, " is now removed!")
								return // exit the loop.
							}
						}
					}
				}(jsRuntimeManager.m_sessions[sessionId], jsRuntimeManager.m_setVariable[sessionId], jsRuntimeManager.m_executeJsFunction[sessionId], jsRuntimeManager.m_stopVm[sessionId])
				callback <- []interface{}{true} // unblock the channel...

			case operationInfos := <-jsRuntimeManager.m_closeVm:
				callback := operationInfos.m_returns
				sessionId := operationInfos.m_params["sessionId"].(string)
				log.Println("-----> remove vm: ", sessionId)
				jsRuntimeManager.m_stopVm[sessionId] <- true // send kill
				jsRuntimeManager.removeVm(sessionId)
				callback <- []interface{}{true} // unblock the channel...
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
	log.Println("Append script file ", filePath)
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
	for _, vm := range this.m_sessions {
		script, err := vm.Compile("", src)
		vm.Run(script)

		if err != nil {
			log.Println("runtime script compilation error:", script, err)
		} else {
			//log.Println(sessionId, " Load: ", this.m_scripts[i])
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
	// Create a new js interpreter for the given session.
	this.m_sessions[sessionId] = otto.New()

	// Put the sessionId variable on the global scope.
	this.m_sessions[sessionId].Set("sessionId", sessionId)
}

/**
 *  Create a new VM for a given session id.
 */
func (this *JsRuntimeManager) removeVm(sessionId string) {
	// Remove vm ressources.
	delete(this.m_sessions, sessionId)
	delete(this.m_executeJsFunction, sessionId)
	delete(this.m_setVariable, sessionId)
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

//////////////////////////////////////////////////////////////////////////////
// Api
//////////////////////////////////////////////////////////////////////////////

/**
 * Append function in the JS, must be call before the js is started.
 */
func (this *JsRuntimeManager) AppendFunction(name string, function interface{}) {
	// I will compile the script and set it in each session...
	this.m_functions[name] = function
}

/**
 * Run given script for a given session. Must be call from inside the JS routine.
 */
func (this *JsRuntimeManager) RunScript(sessionId string, script string) (otto.Value, error) {
	results, err := this.m_sessions[sessionId].Run(script)
	return results, err
}

/**
 * Append a script to all VM's
 */
func (this *JsRuntimeManager) AppendScript(script string) {

	var op OperationInfos
	op.m_params = make(map[string]interface{})
	op.m_params["script"] = script
	op.m_returns = make(chan ([]interface{}))

	this.m_appendScript <- op

	// wait for completion
	<-op.m_returns
}

func (this *JsRuntimeManager) InitScripts(sessionId string) {
	var op OperationInfos
	op.m_params = make(map[string]interface{})
	op.m_params["sessionId"] = sessionId
	op.m_returns = make(chan ([]interface{}))

	this.m_initScripts <- op

	// wait for completion
	<-op.m_returns
}

/**
 * Append and excute a javacript function on the JS...
 */
func (this *JsRuntimeManager) ExecuteJsFunction(messageId string, sessionId string, functionStr string, functionParams []interface{}) ([]interface{}, error) {
	// Set the function on the JS runtime...

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

	this.m_executeJsFunction[sessionId] <- op

	// wait for completion
	results := <-op.m_returns

	return results[0].(JsFunctionInfos).m_results, results[0].(JsFunctionInfos).m_err
}

/**
 * Run given script for a given session.
 */
func (this *JsRuntimeManager) SetVar(sessionId string, name string, val interface{}) {
	// Protectect the map access...
	var info JsVarInfos
	info.m_name = name
	info.m_val = val

	var op OperationInfos
	op.m_params = make(map[string]interface{})
	op.m_params["varInfos"] = info
	op.m_returns = make(chan ([]interface{}))

	this.m_setVariable[sessionId] <- op

	// wait for completion
	<-op.m_returns
}

/**
 * Open a new session on the server.
 */
func (this *JsRuntimeManager) OpendSession(sessionId string) {
	var op OperationInfos
	op.m_params = make(map[string]interface{})
	op.m_params["sessionId"] = sessionId
	op.m_returns = make(chan ([]interface{}))

	this.m_createVm <- op

	// wait for completion
	<-op.m_returns
}

/**
 * Close a given session.
 */
func (this *JsRuntimeManager) CloseSession(sessionId string) {
	// Send message to stop vm...
	var op OperationInfos
	op.m_params = make(map[string]interface{})
	op.m_params["sessionId"] = sessionId
	op.m_returns = make(chan ([]interface{}))
	this.m_closeVm <- op

	// wait for completion
	<-op.m_returns
}
