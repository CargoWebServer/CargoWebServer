package JS

import (
	"bytes"
	b64 "encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
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
 * Interval information, used by JS function setInterval to repeatedly execute
 * a function until it ends by clearInterval
 */
type IntervalInfo struct {
	sessionId string
	uuid      string
	callback  string
	ticker    *time.Ticker
	timer     *time.Timer
}

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

	/** Contain script the time of loading. **/
	m_scripts map[string]string

	/** If a script is runing I keep it path here */
	m_script string

	/** Each connection has it own VM. */
	m_sessions map[string]*otto.Otto

	/** Go function interfaced in JS. **/
	m_functions map[string]interface{}

	/** Exported values for each file **/

	// A module is a directory in src...
	m_modules map[string]*otto.Object

	// each file contain it own export variable.
	m_exports map[string]*otto.Object

	/**Channel use for communication with vm... **/

	// Contain the list of active interval JS function.
	m_intervals     map[string]*IntervalInfo
	m_setInterval   chan *IntervalInfo
	m_clearInterval chan string

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

	// The map of script with their given path.
	jsRuntimeManager.m_scripts = make(map[string]string)

	// List of functions.
	jsRuntimeManager.m_functions = make(map[string]interface{})

	// Map of exported values.
	// each directory in the /Scrip/src ar consider module.
	jsRuntimeManager.m_modules = make(map[string]*otto.Object)
	// each file contain it exports.
	jsRuntimeManager.m_exports = make(map[string]*otto.Object)

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

	// Interval functions stuff.
	jsRuntimeManager.m_intervals = make(map[string]*IntervalInfo, 0)
	jsRuntimeManager.m_setInterval = make(chan *IntervalInfo, 0)
	jsRuntimeManager.m_clearInterval = make(chan string, 0)

	////////////////////////////////////////////////////////////////////////////
	// Javascript std function.
	////////////////////////////////////////////////////////////////////////////

	// Base 64 encoding and decoding.
	// Convert a utf8 string to a base 64 string
	jsRuntimeManager.appendFunction("utf8_to_b64", func(data string) string {
		sEnc := b64.StdEncoding.EncodeToString([]byte(data))
		return sEnc
	})

	jsRuntimeManager.appendFunction("atob", func(data []byte) string {
		str, err := b64.StdEncoding.DecodeString(string(data))
		if err == nil {
			return string(str)
		} else {
			return string(data)
		}
	})

	jsRuntimeManager.appendFunction("btoa", func(str string) string {
		return b64.StdEncoding.EncodeToString([]byte(str))
	})

	////////////////////////////////////////////////////////////////////////////
	// JSON function...
	////////////////////////////////////////////////////////////////////////////
	jsRuntimeManager.appendFunction("stringify", func(object interface{}) string {
		data, _ := json.Marshal(object)
		str := string(data)
		return str
	})

	////////////////////////////////////////////////////////////////////////////
	// String functions...
	////////////////////////////////////////////////////////////////////////////
	jsRuntimeManager.appendFunction("startsWith", func(str string, val string) bool {
		return strings.HasPrefix(str, val)
	})

	jsRuntimeManager.appendFunction("endsWith", func(str string, val string) bool {
		return strings.HasSuffix(str, val)
	})

	jsRuntimeManager.appendFunction("replaceAll", func(str string, val string, by string) string {
		return strings.Replace(str, val, by, -1)
	})

	jsRuntimeManager.appendFunction("capitalizeFirstLetter", func(str string) string {
		return strings.ToUpper(str[0:1]) + str[1:]
	})

	////////////////////////////////////////////////////////////////////////////
	// UUID
	////////////////////////////////////////////////////////////////////////////
	jsRuntimeManager.appendFunction("randomUUID", func() string {
		return Utility.RandomUUID()
	})

	jsRuntimeManager.appendFunction("generateUUID", func(val string) string {
		return Utility.GenerateUUID(val)
	})

	////////////////////////////////////////////////////////////////////////////
	// Timeout/Interval
	////////////////////////////////////////////////////////////////////////////
	jsRuntimeManager.appendFunction("setInterval_", func(callback string, interval int64, sessionId string) string {
		// The intetifier of the function.
		intervalInfo := new(IntervalInfo)
		intervalInfo.sessionId = sessionId
		intervalInfo.uuid = Utility.RandomUUID()
		intervalInfo.callback = callback
		intervalInfo.ticker = time.NewTicker(time.Duration(interval) * time.Millisecond)

		// Set the interval info.
		GetJsRuntimeManager().m_setInterval <- intervalInfo

		return intervalInfo.uuid
	})

	jsRuntimeManager.appendFunction("clearInterval", func(uuid string) {
		GetJsRuntimeManager().m_clearInterval <- uuid
	})

	jsRuntimeManager.appendFunction("setTimeout_", func(callback string, timeout int64, sessionId string) string {
		// The intetifier of the function.
		intervalInfo := new(IntervalInfo)
		intervalInfo.sessionId = sessionId
		intervalInfo.uuid = Utility.RandomUUID()
		intervalInfo.callback = callback
		intervalInfo.timer = time.NewTimer(time.Duration(timeout) * time.Millisecond)

		// Set the interval info.
		GetJsRuntimeManager().m_setInterval <- intervalInfo

		return intervalInfo.uuid
	})

	jsRuntimeManager.appendFunction("clearTimeout", func(uuid string) {
		GetJsRuntimeManager().m_clearInterval <- uuid
	})

	////////////////////////////////////////////////////////////////////////////
	// Node.js compatibility
	////////////////////////////////////////////////////////////////////////////

	// Init NodeJs api.
	jsRuntimeManager.initNodeJs()

	// Load the script from the script repository...
	jsRuntimeManager.appendScriptFiles()

	// That function will process api call's
	go func(jsRuntimeManager *JsRuntimeManager) {
		for {
			select {
			case intervalInfo := <-jsRuntimeManager.m_setInterval:
				// Keep the intervals info in the map.
				jsRuntimeManager.m_intervals[intervalInfo.uuid] = intervalInfo
				// Wait util the timer ends...
				go func(intervalInfo *IntervalInfo) {
					// Set the variable as function.
					functionName := "callback_" + strings.Replace(intervalInfo.uuid, "-", "_", -1)
					_, err := jsRuntimeManager.m_sessions[intervalInfo.sessionId].Run("var " + functionName + "=" + intervalInfo.callback)
					// I must run the script one and at interval after it...
					if err == nil {
						if intervalInfo.ticker != nil {
							// setInterval function.
							for t := range intervalInfo.ticker.C {
								// So here I will call the callback.
								// The callback contain unamed function...
								_, err := GetJsRuntimeManager().GetSession(intervalInfo.sessionId).Run(functionName + "()")
								if err != nil {
									log.Println("---> Run interval callback error: ", err, t)
								}
							}
						} else if intervalInfo.timer != nil {
							// setTimeout function
							<-intervalInfo.timer.C
							_, err := GetJsRuntimeManager().RunScript(intervalInfo.sessionId, functionName+"()")
							if err != nil {
								log.Println("---> Run timeout callback error: ", err)
							}
						}
					} else {
						log.Println("---> Run interval callback error: ", "var "+functionName+"="+intervalInfo.callback, err)
					}

				}(intervalInfo)
			case uuid := <-jsRuntimeManager.m_clearInterval:
				intervalInfo := jsRuntimeManager.m_intervals[uuid]
				if intervalInfo != nil {
					if intervalInfo.ticker != nil {
						intervalInfo.ticker.Stop()
					} else if intervalInfo.timer != nil {
						intervalInfo.timer.Stop()
					}
					// Remove the interval/timeout information.
					delete(jsRuntimeManager.m_intervals, uuid)
				} else {
					// In that case the uuid is a session id and not an interval id.
					for _, intervalInfo_ := range jsRuntimeManager.m_intervals {
						if intervalInfo_.sessionId == uuid {
							if intervalInfo_.ticker != nil {
								intervalInfo_.ticker.Stop()
							} else if intervalInfo_.timer != nil {
								intervalInfo_.timer.Stop()
							}
							// Remove the interval/timeout information.
							delete(jsRuntimeManager.m_intervals, intervalInfo_.uuid)
						}
					}
				}
			case functionInfo := <-jsRuntimeManager.m_setFunction:
				jsRuntimeManager.m_functions[functionInfo.m_functionId] = functionInfo.m_function
			case sessionInfo := <-jsRuntimeManager.m_getSession:
				sessionInfo.m_return <- jsRuntimeManager.m_sessions[sessionInfo.m_sessionId]
			case operationInfos := <-jsRuntimeManager.m_appendScript:
				callback := operationInfos.m_returns
				path := operationInfos.m_params["path"].(string)
				script := operationInfos.m_params["script"].(string)
				jsRuntimeManager.appendScript(path, script)
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
	this.appendScript(filePath, src)

	return err
}

/**
 * Append a function in the function map.
 */
func (this *JsRuntimeManager) appendFunction(functionId string, function interface{}) {
	this.m_functions[functionId] = function
}

/**
 * Append a script and initialyse all vm with it.
 */
func (this *JsRuntimeManager) appendScript(path string, src string) {
	if val, ok := this.m_scripts[path]; ok {
		// append sricpt.
		this.m_scripts[path] = src + "\n" + val
	} else {
		this.m_scripts[path] = src
	}
}

/**
 * Part of the module/exports
 */
func (this *JsRuntimeManager) getExports(path string) (*otto.Object, error) {

	// Here from the path information and the current running script i will
	// return infromation conatain in the modules/export.
	var currentPath = this.m_script
	// First of all I will retreive the module.

	// If the path begin by ./ that means the current module must be use...
	var moduleId string
	var dir string

	if strings.HasPrefix(path, "./") {
		// In that case I will use the current file path to locate the module.
		strartIndex := strings.Index(filepath.ToSlash(this.m_script), "/WebApp/Cargo/Script/src/")
		if strartIndex > 0 {
			strartIndex += len("/WebApp/Cargo/Script/src/")
			endIndex := strings.Index(filepath.ToSlash(this.m_script[strartIndex:]), "/")
			if endIndex != -1 {
				moduleId = this.m_script[strartIndex:][0:endIndex]
			}
			dir = this.m_script[0:strings.LastIndex(filepath.ToSlash(this.m_script), "/")]
		}
	} else {
		// Here the first part of the path define the module.
		index := strings.Index(path, "/")
		if index == -1 {
			moduleId = path[0:]
		} else {
			moduleId = path[0:index]
		}
		dir = this.m_searchDir + "/src"
	}

	// Now I will count the number of up
	upCount := strings.Count(path, "../")
	if upCount > 0 {
		for i := 0; i < upCount; i++ {
			dir = dir[0:strings.LastIndex(filepath.ToSlash(dir), "/")]
		}
	}

	// Now I will recreate the export path from the dir ant the path.
	exportPath := dir
	// reconstruct the file path.
	if dir != path {
		if string(filepath.Separator) == "/" {
			exportPath += "/" + strings.Replace(strings.Replace(path, "../", "", -1), "./", "", -1)
			exportPath = strings.Replace(exportPath, "\\", "/", -1)
		} else {
			exportPath += "\\" + strings.Replace(strings.Replace(strings.Replace(path, "../", "", -1), "./", "", -1), "/", "\\", -1)
			exportPath = strings.Replace(exportPath, "/", "\\", -1)
		}
	}

	// Now I will get the export from the error...
	exports := this.m_exports[exportPath]
	if exports == nil {
		// Here I need to import a file...
		exports = this.initScript("", exportPath+".js")
		this.m_script = currentPath // Set back the path to the file before the call.bytes
	}

	// set the global variable exports...
	currentExportPath := currentPath
	if strings.HasSuffix(currentExportPath, ".js") {
		currentExportPath = currentExportPath[0 : len(currentExportPath)-3]
	}

	// If the exports is not in the path.
	if this.m_exports[currentExportPath] == nil {
		this.m_exports[currentExportPath], _ = this.m_sessions[""].Object("exports = {}")
		this.m_exports[currentExportPath].Set("__path__", currentExportPath)
		this.m_modules[moduleId].Set("exports", this.m_exports[currentExportPath]) // Set it export values.
		exports = this.m_exports[currentExportPath]
	}

	this.m_sessions[""].Set("module", this.m_modules[moduleId])
	this.m_sessions[""].Set("exports", this.m_exports[currentExportPath])

	return exports, nil

}

/**
 * Initialisation of script for a newly created session.
 */
func (this *JsRuntimeManager) initScripts(sessionId string) {
	// Get the vm.
	vm := this.m_sessions[sessionId]

	// Exported golang JS function.
	for name, function := range this.m_functions {
		if strings.Index(name, ".") != -1 {
			// Thats means the function is part of a module so I will
			// create the module if it not exist and append the funtion
			// in it exports variable.
			values := strings.Split(name, ".")
			name = values[1]
			moduleId := values[0]
			var exportPath string
			if string(filepath.Separator) == "/" {
				exportPath = filepath.ToSlash(this.m_searchDir + "/src/" + moduleId)
			} else {
				exportPath = filepath.FromSlash(this.m_searchDir + "/src/" + moduleId)
			}

			if this.m_modules[moduleId] == nil {
				this.m_modules[moduleId], _ = vm.Object("module = {}")
			}

			if this.m_exports[exportPath] == nil {
				this.m_exports[exportPath], _ = vm.Object("exports = {}")
				this.m_exports[exportPath].Set("__path__", exportPath)
				this.m_modules[moduleId].Set("exports", this.m_exports[exportPath]) // Set it export values.
			}

			// Set the function in the module object.
			this.m_exports[exportPath].Set(name, function)
		}
		vm.Set(name, function)
	}

	// Init srcipts.
	// Native Cargo script first.
	for path, _ := range this.m_scripts {
		// Start initalyse the scripts.
		if strings.HasPrefix(path, "CargoWebServer") || strings.Index(path, "/WebApp/Cargo/Script/src/Cargo/") != -1 {
			this.initScript(sessionId, path)
		}
	}

	// Other application srcipt.
	for path, _ := range this.m_scripts {
		// Start initalyse the scripts.
		this.initScript(sessionId, path)
	}
}

/**
 * Set init a script.
 */
func (this *JsRuntimeManager) initScript(sessionId string, path string) *otto.Object {

	var exports *otto.Object
	this.m_script = path

	var moduleId string
	strartIndex := strings.Index(filepath.ToSlash(this.m_script), "/WebApp/Cargo/Script/src/")
	if strartIndex > 0 {
		strartIndex += len("/WebApp/Cargo/Script/src/")
		endIndex := strings.Index(filepath.ToSlash(this.m_script[strartIndex:]), "/")
		if endIndex != -1 {
			moduleId = this.m_script[strartIndex:][0:endIndex]
		}
	} else {
		// Here the first part of the path define the module.
		index := strings.Index(filepath.ToSlash(path), "/")
		if index == -1 {
			moduleId = path[0:]
		} else {
			moduleId = path[0:index]
		}
	}

	// The vm to run the script.
	var vm *otto.Otto
	// All function of Cargo and CargoWebServer are public, no exports needed.
	// require will be use to synchronize the order of initialysation in their case.
	if moduleId == "Cargo" || moduleId == "CargoWebServer" {
		// initialyse cargo anonymous session here, that session is the one
		// that contain remote access code.
		vm = this.m_sessions[sessionId]
	} else {
		// Local code session.
		// Create new vm and run code inside it.
		vm = otto.New()
		for name, function := range this.m_functions {
			// Append general scope function only. ex require, atoa, setInterval...
			if strings.Index(name, ".") == -1 {
				vm.Set(name, function)
			}
		}
	}

	if this.m_modules[moduleId] == nil {
		this.m_modules[moduleId], _ = vm.Object("module = {}")
	}

	exportPath := path
	if strings.HasSuffix(exportPath, ".js") {
		exportPath = exportPath[0 : len(exportPath)-3]
	}

	if this.m_exports[exportPath] == nil {
		this.m_exports[exportPath], _ = vm.Object("exports = {}")
		this.m_exports[exportPath].Set("__path__", exportPath)
		this.m_modules[moduleId].Set("exports", this.m_exports[exportPath]) // Set it export values.
		exports = this.m_exports[exportPath]
	} else {
		// Here I will return the exports
		// set the global variable exports...
		exports = this.m_exports[exportPath]
	}

	if src, ok := this.m_scripts[path]; ok {
		script, err := vm.Compile("", src)
		if err == nil {
			// set the global variable exports...
			vm.Set("exports", this.m_exports[exportPath])
			// set the current module.
			vm.Set("module", this.m_modules[moduleId])
			// set the export as return value
			exports = this.m_exports[exportPath]

			_, err := vm.Run(script)

			if err != nil {
				log.Println("---> script running error:  ", path, err)
			}
		} else {
			log.Println("---> script compilation error:  ", path, err)
		}

		// remove it from the map.
		delete(this.m_scripts, path)
	}

	return exports
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

	// I will also clear intervals for the
	// the session.
	this.m_clearInterval <- sessionId
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
func (this *JsRuntimeManager) GetSession(sessionId string) *otto.Otto {
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
func (this *JsRuntimeManager) AppendScript(path string, script string) {

	var op OperationInfos
	op.m_params = make(map[string]interface{})
	op.m_params["path"] = path
	op.m_params["script"] = script
	op.m_returns = make(chan ([]interface{}))
	defer close(op.m_returns)
	this.m_appendScript <- op

	// wait for completion
	<-op.m_returns
}

/**
 * Must be called once after all script are imported.
 */
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

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()

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
