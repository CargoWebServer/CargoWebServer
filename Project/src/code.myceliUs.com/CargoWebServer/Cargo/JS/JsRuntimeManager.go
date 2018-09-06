package JS

import (
	"bytes"
	b64 "encoding/base64"
	"encoding/json"
	"errors"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"code.myceliUs.com/GoJavaScript"
	"code.myceliUs.com/GoJavaScript/GoJavaScriptClient"
	"code.myceliUs.com/Utility"
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
	m_name    string
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
	m_return    chan (*GoJavaScriptClient.Client)
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
	m_sessions map[string]*GoJavaScriptClient.Client

	/** Go function interfaced in JS. **/
	m_functions map[string]interface{}

	/** Exported values for each file for each session **/
	m_exports map[string]map[string]GoJavaScript.Object

	/**Channel use for communication with vm... **/

	// Contain the list of active interval JS function.
	m_intervals     map[string]*IntervalInfo
	m_setInterval   chan *IntervalInfo
	m_clearInterval chan string

	// Script channels
	m_appendScript OperationChannel

	// Open/Close channel.
	m_createVm OperationChannel
	m_closeVm  OperationChannel

	// Use that channel to execute various vm operation instead of
	// accessing sessions channels maps (m_setVariable, m_getVariable...)
	// directly.
	m_execVmOperation OperationChannel

	// JS channel
	m_setFunction chan (FunctionInfos)
	m_getSession  chan (SessionInfos)

	// Map of channel by session.
	m_setVariable       map[string]OperationChannel
	m_getVariable       map[string]OperationChannel
	m_executeJsFunction map[string]OperationChannel
	m_runScript         map[string]OperationChannel
	m_stopVm            map[string]chan (bool)

	// The starting port number for session VM.
	m_startPort int
}

// The main processing loop...
func run(jsRuntimeManager *JsRuntimeManager) {
	// That function will process api call's
	for {
		select {
		case intervalInfo := <-jsRuntimeManager.m_setInterval:
			// Keep the intervals info in the map.
			jsRuntimeManager.m_intervals[intervalInfo.uuid] = intervalInfo
			// Wait util the timer ends...
			go func(intervalInfo *IntervalInfo) {
				// Set the variable as function.
				functionName := "callback_" + strings.Replace(intervalInfo.uuid, "-", "_", -1)
				_, err := jsRuntimeManager.m_sessions[intervalInfo.sessionId].EvalScript("var "+functionName+"="+intervalInfo.callback, []interface{}{})

				// I must run the script one and at interval after it...
				if err == nil {
					if intervalInfo.ticker != nil {
						// setInterval function.
						for t := range intervalInfo.ticker.C {
							// So here I will call the callback.
							// The callback contain unamed function...
							_, err := jsRuntimeManager.m_sessions[intervalInfo.sessionId].EvalScript(functionName+"()", []interface{}{})
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
			jsRuntimeManager.appendFunction(functionInfo.m_functionId, functionInfo.m_function)
		case sessionInfo := <-jsRuntimeManager.m_getSession:
			sessionInfo.m_return <- jsRuntimeManager.m_sessions[sessionInfo.m_sessionId]
		case operationInfos := <-jsRuntimeManager.m_appendScript:
			callback := operationInfos.m_returns
			path := operationInfos.m_params["path"].(string)
			script := operationInfos.m_params["script"].(string)
			jsRuntimeManager.appendScript(path, script)
			// In case the script must be run...
			if operationInfos.m_params["run"].(bool) {
				for _, session := range jsRuntimeManager.m_sessions {
					session.EvalScript(script, []interface{}{})
				}
			}
			callback <- []interface{}{true} // unblock the channel...
		case operationInfos := <-jsRuntimeManager.m_execVmOperation:
			// Set the function on the JS runtime...
			var sessionId = operationInfos.m_params["sessionId"].(string)

			if jsRuntimeManager.m_sessions[sessionId] != nil {
				// Here I will execute various session action.
				if operationInfos.m_name == "GetVar" {
					jsRuntimeManager.m_getVariable[operationInfos.m_params["sessionId"].(string)] <- operationInfos
				} else if operationInfos.m_name == "SetVar" {
					jsRuntimeManager.m_setVariable[operationInfos.m_params["sessionId"].(string)] <- operationInfos
				} else if operationInfos.m_name == "ExecuteJsFunction" {
					jsRuntimeManager.m_executeJsFunction[operationInfos.m_params["sessionId"].(string)] <- operationInfos
				} else if operationInfos.m_name == "RunScript" {
					jsRuntimeManager.m_runScript[operationInfos.m_params["sessionId"].(string)] <- operationInfos
				}
			} else {
				log.Println("---> try to execute function on close channel whit id: ", sessionId)
				// return nil, errors.New("Session " + sessionId + " is closed!")
			}

		case operationInfos := <-jsRuntimeManager.m_createVm:
			callback := operationInfos.m_returns
			sessionId := operationInfos.m_params["sessionId"].(string)
			if jsRuntimeManager.m_sessions[sessionId] == nil {
				// Create one JS engine by session, or task...
				jsRuntimeManager.createVm(sessionId)

				// Create the list of channel to communicate with the
				// JS engine.
				jsRuntimeManager.m_setVariable[sessionId] = make(OperationChannel)
				jsRuntimeManager.m_getVariable[sessionId] = make(OperationChannel)
				jsRuntimeManager.m_executeJsFunction[sessionId] = make(OperationChannel)
				jsRuntimeManager.m_runScript[sessionId] = make(OperationChannel)
				jsRuntimeManager.m_stopVm[sessionId] = make(chan (bool))

				// The session processing loop...
				go func(vm *GoJavaScriptClient.Client, setVariable OperationChannel, getVariable OperationChannel, executeJsFunction OperationChannel, runScript OperationChannel, stopVm chan (bool), sessionId string) {
					// The session was interrupt!
					defer func() {
						// Stahp mean the VM was kill by the admin.
						if caught := recover(); caught != nil {
							if caught.(error).Error() == "Stahp" {
								// Here the task was cancel.
								log.Println("---> session: ", sessionId, " is now closed")
								return
							} else {
								panic(caught) // Something else happened, repanic!
							}
						}
					}()

				process:
					for {
						select {
						case operationInfos := <-setVariable:
							callback := operationInfos.m_returns
							if operationInfos.m_params["varInfos"] != nil {
								varInfos := operationInfos.m_params["varInfos"].(JsVarInfos)
								vm.SetGlobalVariable(varInfos.m_name, varInfos.m_val)
							}
							callback <- []interface{}{true} // unblock the channel...
						case operationInfos := <-getVariable:
							callback := operationInfos.m_returns
							var varInfos JsVarInfos
							if operationInfos.m_params["varInfos"] != nil {
								varInfos = operationInfos.m_params["varInfos"].(JsVarInfos)
								value, err := vm.GetGlobalVariable(varInfos.m_name)
								if err == nil {
									varInfos.m_val = value
								}
							}
							callback <- []interface{}{varInfos} // unblock the channel...
						case operationInfos := <-executeJsFunction:
							callback := operationInfos.m_returns
							var jsFunctionInfos JsFunctionInfos
							if operationInfos.m_params["jsFunctionInfos"] != nil {
								jsFunctionInfos = operationInfos.m_params["jsFunctionInfos"].(JsFunctionInfos)
								vm.SetGlobalVariable("messageId", jsFunctionInfos.m_messageId)
								vm.SetGlobalVariable("sessionId", jsFunctionInfos.m_sessionId)
								//log.Println("---> execute jsFunction: ", jsFunctionInfos.m_functionStr)
								jsFunctionInfos.m_results, jsFunctionInfos.m_err = GetJsRuntimeManager().executeJsFunction(vm, jsFunctionInfos.m_functionStr, jsFunctionInfos.m_functionParams)
							}
							callback <- []interface{}{jsFunctionInfos} // unblock the channel...
						case operationInfos := <-runScript:
							callback := operationInfos.m_returns
							var results interface{}
							var err error
							if operationInfos.m_params["script"] != nil {
								results, err = vm.EvalScript(operationInfos.m_params["script"].(string), []interface{}{})
							}
							callback <- []interface{}{results, err} // unblock the channel...
						case stop := <-stopVm:
							if stop {
								vm.Stop()
								break process // exit the processing loop.
							}
						}
					}
				}(jsRuntimeManager.m_sessions[sessionId], jsRuntimeManager.m_setVariable[sessionId], jsRuntimeManager.m_getVariable[sessionId], jsRuntimeManager.m_executeJsFunction[sessionId], jsRuntimeManager.m_runScript[sessionId], jsRuntimeManager.m_stopVm[sessionId], sessionId)
			}
			// Close vm callback...
			callback <- []interface{}{true} // unblock the channel...

		case operationInfos := <-jsRuntimeManager.m_closeVm:
			sessionId := operationInfos.m_params["sessionId"].(string)
			callback := operationInfos.m_returns
			if jsRuntimeManager.m_sessions[sessionId] != nil {
				callback <- []interface{}{true} // unblock the channel...
				// here I will not wait for the session to clean before retrun.
				go func() {
					// Wait until the vm is stop
					jsRuntimeManager.m_stopVm[sessionId] <- true // stop processing loop
					jsRuntimeManager.removeVm(sessionId)
				}()
			} else {
				// simply call the callback if the sesion is already close.
				callback <- []interface{}{true} // unblock the channel...
			}
		}
	}
}

func NewJsRuntimeManager(searchDir string) *JsRuntimeManager {
	// The singleton.
	jsRuntimeManager = new(JsRuntimeManager)

	jsRuntimeManager.m_startPort = 9696

	// The dir where the js file are store on the server side.
	jsRuntimeManager.m_searchDir = filepath.ToSlash(searchDir)

	// List of vm one per connection.
	jsRuntimeManager.m_sessions = make(map[string]*GoJavaScriptClient.Client)

	// The map of script with their given path.
	jsRuntimeManager.m_scripts = make(map[string]string)

	// List of functions.
	jsRuntimeManager.m_functions = make(map[string]interface{})

	// each file contain it exports.
	jsRuntimeManager.m_exports = make(map[string]map[string]GoJavaScript.Object)

	// Initialisation of channel's.
	jsRuntimeManager.m_appendScript = make(OperationChannel)
	jsRuntimeManager.m_closeVm = make(OperationChannel)
	jsRuntimeManager.m_createVm = make(OperationChannel)
	jsRuntimeManager.m_execVmOperation = make(OperationChannel)

	// Create one channel by vm to set variable.
	jsRuntimeManager.m_setFunction = make(chan (FunctionInfos))
	jsRuntimeManager.m_getSession = make(chan (SessionInfos))

	// Create sessions channel container.
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
		b, _ := json.Marshal(object)
		b, _ = Utility.PrettyPrint(b)
		s := string(b)
		return s
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

	// Bypass error here and try to loop again...
	defer func(jsRuntimeManager *JsRuntimeManager) {
		if caught := recover(); caught != nil {
			log.Println("----> error caught at line 539 of JsRuntimeManager. ", caught.(error).Error())
			go run(jsRuntimeManager)
		}
	}(jsRuntimeManager)

	// Run the main processing loop.
	go run(jsRuntimeManager)

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
			err = this.appendScriptFile(filepath.ToSlash(file))
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
func (this *JsRuntimeManager) getExports(path string, sessionId string) (GoJavaScript.Object, error) {

	// If the path begin by ./ that means the current module must be use...
	var dir string

	if strings.HasPrefix(path, "./") {
		// In that case I will use the current file path to locate the module.
		strartIndex := strings.Index(this.m_script, "/WebApp/Cargo/Script/src/")
		if strartIndex > 0 {
			dir = this.m_script[0:strings.LastIndex(this.m_script, "/")]
		}
	} else {
		// Here the first part of the path define the module.
		dir = this.m_searchDir + "/src"
	}

	// Now I will count the number of up
	upCount := strings.Count(path, "../")
	if upCount > 0 {
		for i := 0; i < upCount; i++ {
			dir = dir[0:strings.LastIndex(dir, "/")]
		}
	}

	// Now I will recreate the export path from the dir ant the path.
	exportPath := dir

	// reconstruct the file path.
	if dir != path {
		exportPath += "/" + strings.Replace(strings.Replace(path, "../", "", -1), "./", "", -1)
	}

	// Before I call init script i will keep current script in local variable...
	// and set it back after the function return.
	var currentPath = this.m_script

	// Here I need to import a file...
	exports := this.initScript(exportPath+".js", sessionId)

	// Set it back...
	this.m_script = currentPath // Set back the path to the file before the call.bytes

	if strings.HasSuffix(currentPath, ".js") {
		currentPath = currentPath[0 : len(currentPath)-3]
	}

	// TODO be sure
	/*if export, ok := this.m_exports[sessionId][currentPath]; ok {
		log.Println("---> 670: set exports!")
		this.m_sessions[sessionId].SetGlobalVariable("exports", export)
	} else {
		log.Panicln("---> no exports found for path ", currentPath)
	}*/

	return exports, nil

}

/**
 * Return the module id from a given path.
 */
func (this *JsRuntimeManager) getModuleId(path string) string {
	var moduleId string
	var startIndex = strings.Index(path, this.m_searchDir+"/src/")
	if startIndex != -1 {
		startIndex += len(this.m_searchDir + "/src/")
		endIndex := strings.Index(path[startIndex:], "/")
		if endIndex != -1 {
			moduleId = path[startIndex : startIndex+endIndex]
		} else {
			moduleId = path[startIndex:]
		}
	} else if strings.Index(path, "/") != -1 {
		moduleId = path[0:strings.Index(path, "/")]
	} else {
		moduleId = path
	}
	return moduleId
}

/**
 * Initialisation of script for a newly created session.
 */
func (this *JsRuntimeManager) initScripts(sessionId string) {

	// Get the vm.
	vm := this.m_sessions[sessionId]

	// Need to be set as a global variable.
	vm.SetGlobalVariable("sessionId", sessionId)

	// I will register the require function first.
	vm.RegisterJsFunction("require", "function require(moduleId){return require_(moduleId, sessionId)}")

	// Create the map of exports.
	jsRuntimeManager.m_exports[sessionId] = make(map[string]GoJavaScript.Object)

	// Exported golang JS function.
	for name, function := range this.m_functions {
		// CargoWebServer function are set in the global scope.
		name = strings.Replace(name, "CargoWebServer.", "", -1)
		if strings.Index(name, ".") != -1 {
			// Thats means the function is part of a module so I will
			// create the module if it not exist and append the funtion
			// in it exports variable.
			values := strings.Split(name, ".")
			name = values[1]
			moduleId := values[0]
			var exportPath string
			exportPath = filepath.ToSlash(this.m_searchDir + "/src/" + moduleId)

			if export, ok := this.m_exports[sessionId][exportPath]; ok {
				// Set the function as part of the exports object.
				export.Set(name, function)
			} else {
				export = vm.CreateObject("exports")

				// Keep it path
				export.Set("path__", exportPath)
				export.Set("module_id__", moduleId)

				this.m_exports[sessionId][exportPath] = export

				// Set the function as part of the exports object.
				export.Set(name, function)
			}

		} else {
			// Register go function.
			vm.RegisterGoFunction(name, function)
		}
	}

	// Init srcipts.
	// Native Cargo script first.
	for path, _ := range this.m_scripts {
		// Start initalyse the scripts.
		if this.getModuleId(path) == "CargoWebServer" {
			this.initScript(path, sessionId)
		}
	}

	// Load only the Cargo directory by default,
	// Other module will be load with use of require in server side script execution.
	for path, _ := range this.m_scripts {
		// Start initalyse the scripts.
		if this.getModuleId(path) == "Cargo" {
			this.initScript(path, sessionId)
		}
	}
}

/**
 * Set init a script.
 */
func (this *JsRuntimeManager) initScript(path string, sessionId string) GoJavaScript.Object {
	vm := this.m_sessions[sessionId]
	moduleId := this.getModuleId(path)
	this.m_script = path

	exportPath := path
	if strings.HasSuffix(exportPath, ".js") {
		exportPath = exportPath[0 : len(exportPath)-3]
	}

	if jsRuntimeManager.m_exports[sessionId] == nil {
		jsRuntimeManager.m_exports[sessionId] = make(map[string]GoJavaScript.Object)
	}

	if export, ok := this.m_exports[sessionId][exportPath]; ok {
		return export
	} else {
		// create a new path
		log.Println("---> init script: ", path)
		export := vm.CreateObject("exports")
		export.Set("path__", exportPath)
		export.Set("module_id__", moduleId)
		this.m_exports[sessionId][exportPath] = export

		if src, ok := this.m_scripts[path]; ok {
			_, err := vm.EvalScript(src, []interface{}{})
			if err != nil {
				log.Panicln("---> script running error:  ", path, err)
			}
		}
		return export
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

	log.Println("---> create a new vm for session: ", sessionId)

	// Create a new js interpreter for the given session.
	this.m_sessions[sessionId] = GoJavaScriptClient.NewClient("127.0.0.1", this.m_startPort, "jerryscript" /* "chakracore"*/)

	this.m_startPort += 1

	if sessionId != "" {
		this.initScripts(sessionId)
	}
}

/**
 *  Close and remove VM for a given session id.
 */
func (this *JsRuntimeManager) removeVm(sessionId string) {
	if sessionId == "" {
		return // never delete the general session
	}

	log.Println("---> remove ressource for session: ", sessionId)

	// I will also clear intervals for the
	// the session.
	this.m_clearInterval <- sessionId

	// Execute channel id
	if this.m_executeJsFunction[sessionId] != nil {
		close(this.m_executeJsFunction[sessionId])
		delete(this.m_executeJsFunction, sessionId)
	}

	if this.m_runScript[sessionId] != nil {
		close(this.m_runScript[sessionId])
		delete(this.m_runScript, sessionId)
	}

	if this.m_setVariable[sessionId] != nil {
		close(this.m_setVariable[sessionId])
		delete(this.m_setVariable, sessionId)
	}

	if this.m_stopVm[sessionId] != nil {
		close(this.m_stopVm[sessionId])
		delete(this.m_stopVm, sessionId)
	}

	if this.m_setVariable[sessionId] != nil {
		close(this.m_setVariable[sessionId])
		delete(this.m_setVariable, sessionId)
	}

	if this.m_getVariable[sessionId] != nil {
		close(this.m_getVariable[sessionId])
		delete(this.m_getVariable, sessionId)
	}

	if this.m_getVariable[sessionId] != nil {
		close(this.m_getVariable[sessionId])
		delete(this.m_getVariable, sessionId)
	}

	if this.m_exports[sessionId] != nil {
		delete(this.m_exports, sessionId)
	}

	if this.m_sessions[sessionId] != nil {
		delete(this.m_sessions, sessionId)
	}

}

/**
 * Execute javascript function.
 */
func (this *JsRuntimeManager) executeJsFunction(vm *GoJavaScriptClient.Client, functionStr string, functionParams []interface{}) (results []interface{}, err error) {

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

		_, err = vm.EvalScript(functionStr, []interface{}{})
		if err != nil {
			log.Println("fail to run ", functionStr)
			return nil, err
		}

	} else {
		functionName = functionStr
	}
	result, err_ := vm.CallFunction(functionName, functionParams...)

	if err_ != nil {
		log.Println("Error found with function ", functionName, err_.(error), "params: ", functionParams)
		log.Println("Src ", functionStr)
		return nil, err_.(error)
	}

	// Return the result if there is one...
	val, err := result.Export()

	if err != nil {
		log.Panicln("---> error ", err)
		return nil, err
	}

	// Append val to results...
	results = append(results, val)

	return
}

/**
 * Run given script for a given session.
 */
func (this *JsRuntimeManager) GetSession(sessionId string) *GoJavaScriptClient.Client {
	// Protectect the map access...
	var sessionInfo SessionInfos
	sessionInfo.m_return = make(chan (*GoJavaScriptClient.Client))
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
func (this *JsRuntimeManager) RunScript(sessionId string, script string) (GoJavaScript.Value, error) {
	// Protectect the map access...
	var op OperationInfos
	op.m_name = "RunScript"
	op.m_params = make(map[string]interface{})
	op.m_params["script"] = script
	op.m_params["sessionId"] = sessionId
	op.m_returns = make(chan ([]interface{}))
	defer close(op.m_returns)

	this.m_execVmOperation <- op

	// wait for completion
	results := <-op.m_returns
	var value GoJavaScript.Value
	var err error

	if results[0] != nil {
		value = results[0].(GoJavaScript.Value)
	}

	if results[1] != nil {
		err = results[1].(error)
		log.Println("--> 914 err: ", err)
		log.Println("--> 915 script: ", script)
	}

	return value, err
}

/**
 * Append a script to all VM's
 */
func (this *JsRuntimeManager) AppendScript(path string, script string, run bool) {
	var op OperationInfos
	op.m_params = make(map[string]interface{})
	op.m_params["path"] = path
	op.m_params["script"] = script
	op.m_params["run"] = run
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
	this.initScripts(sessionId)
}

/**
 * Append and excute a javacript function on the JS...
 */
func (this *JsRuntimeManager) ExecuteJsFunction(messageId string, sessionId string, functionStr string, functionParams []interface{}) ([]interface{}, error) {

	// Put function call information into a struct.
	var jsFunctionInfos JsFunctionInfos
	jsFunctionInfos.m_messageId = messageId
	jsFunctionInfos.m_sessionId = sessionId
	jsFunctionInfos.m_functionStr = functionStr
	jsFunctionInfos.m_functionParams = functionParams

	var op OperationInfos
	op.m_name = "ExecuteJsFunction"
	op.m_params = make(map[string]interface{})
	op.m_params["jsFunctionInfos"] = jsFunctionInfos
	op.m_params["sessionId"] = sessionId
	op.m_returns = make(chan ([]interface{}))
	defer close(op.m_returns)
	this.m_execVmOperation <- op

	// wait for completion
	results := <-op.m_returns
	return results[0].(JsFunctionInfos).m_results, results[0].(JsFunctionInfos).m_err

}

/**
 * Set variable value for a given session
 */
func (this *JsRuntimeManager) SetVar(sessionId string, name string, val interface{}) {

	// Protectect the map access...
	var info JsVarInfos
	info.m_name = name
	info.m_val = val

	var op OperationInfos
	op.m_name = "SetVar"
	op.m_params = make(map[string]interface{})
	op.m_params["varInfos"] = info
	op.m_params["sessionId"] = sessionId
	op.m_returns = make(chan ([]interface{}))
	defer close(op.m_returns)
	this.m_execVmOperation <- op
	// wait for completion
	<-op.m_returns
}

/**
 * Run given script for a given session.
 */
func (this *JsRuntimeManager) GetVar(sessionId string, name string) interface{} {

	// Protectect the map access...
	var info JsVarInfos
	info.m_name = name

	var op OperationInfos
	op.m_name = "GetVar"
	op.m_params = make(map[string]interface{})
	op.m_params["varInfos"] = info
	op.m_params["sessionId"] = sessionId
	op.m_returns = make(chan ([]interface{}))
	defer close(op.m_returns)
	this.m_execVmOperation <- op

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
	// Error handler.

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
