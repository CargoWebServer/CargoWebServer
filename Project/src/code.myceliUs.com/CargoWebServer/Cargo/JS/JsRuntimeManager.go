package JS

import (
	"bytes"
	"errors"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"code.myceliUs.com/Utility"
	"github.com/robertkrimen/otto"
	_ "github.com/robertkrimen/otto/underscore"
)

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

var (
	jsRuntimeManager *JsRuntimeManager
)

/**
 * The javascript runtime is use to extent the functionnality of the system
 * via javascript...
 */
type JsRuntimeManager struct {
	/** Where the script are **/
	m_searchDir string

	/** Contain the scripts source **/
	m_scripts []string

	/** One js interpreter by session */
	m_sessions map[string]*otto.Otto

	/** Each session has it own channel **/
	m_channels map[string]chan (JsFunctionInfos)

	/** a list of function define in cargo to inject in the VM. **/
	m_functions map[string]interface{}

	/**
	 * Use to protected the entitiesMap access...
	 */
	sync.RWMutex
}

func NewJsRuntimeManager(searchDir string) *JsRuntimeManager {
	// assign memory...
	jsRuntimeManager = new(JsRuntimeManager)
	jsRuntimeManager.m_searchDir = searchDir
	jsRuntimeManager.m_sessions = make(map[string]*otto.Otto)
	jsRuntimeManager.m_functions = make(map[string]interface{})
	jsRuntimeManager.m_channels = make(map[string]chan (JsFunctionInfos))

	// Load the script from the script repository...
	jsRuntimeManager.AppendScriptFiles()

	return jsRuntimeManager
}

func GetJsRuntimeManager() *JsRuntimeManager {
	return jsRuntimeManager
}

/**
 * Return the current vm for a given session.
 */
func (this *JsRuntimeManager) CreateVm(sessionId string) *otto.Otto {
	// Protectect the map access...
	this.Lock()
	defer this.Unlock()

	// Create a new js interpreter for the given session.
	this.m_sessions[sessionId] = otto.New()
	this.m_sessions[sessionId].Set("sessionId", sessionId)

	this.m_channels[sessionId] = make(chan (JsFunctionInfos)) // open it channel.

	// I will start a process loop here. It will run until the vm is open.
	go func(sessionId string, channel chan (JsFunctionInfos)) {
		// I will get the vm...
		vm := GetJsRuntimeManager().GetVm(sessionId)
		for vm != nil {
			select {
			case jsFunctionInfos := <-channel:
				log.Println("----------> received message on channel", jsFunctionInfos.m_messageId)
				vm := this.GetVm(sessionId)

				// Set the current session id.
				vm.Set("messageId", jsFunctionInfos.m_messageId)
				// TODO set the answer here...

				channel <- jsFunctionInfos
			}
		}

		// remove the channel and release it ressources.

	}(sessionId, this.m_channels[sessionId])

	return this.m_sessions[sessionId]
}

/**
 * Init script.
 */
func (this *JsRuntimeManager) InitScripts(sessionId string) {
	// Compile the list of script...
	for i := 0; i < len(this.m_scripts); i++ {
		script, err := this.m_sessions[sessionId].Compile("", this.m_scripts[i])
		if err == nil {
			if this.m_sessions[sessionId] != nil {
				this.m_sessions[sessionId].Run(script)
				if err != nil {
					log.Println("runtime script compilation error:", script, err)
				} else {
					//log.Println(sessionId, " Load: ", this.m_scripts[i])
				}
			}
		} else {
			log.Println("-------> error in script: ", this.m_scripts[i])
			log.Println(err)
		}
	}

	// Set the list of binded function.
	for name, function := range this.m_functions {
		this.m_sessions[sessionId].Set(name, function)
	}
}

/**
 * Return the current vm for a given session.
 */
func (this *JsRuntimeManager) GetVm(sessionId string) *otto.Otto {
	// Protectect the map access...
	this.Lock()
	defer this.Unlock()

	// If the vm exist, i will simply return the vm...
	if vm, ok := this.m_sessions[sessionId]; ok {
		return vm
	}

	return nil
}

func (this *JsRuntimeManager) CloseSession(sessionId string) {
	// Protectect the map access...
	this.Lock()
	defer this.Unlock()

	// simply remove the js interpreter...
	delete(this.m_sessions, sessionId)
}

/** Append all scripts **/
func (this *JsRuntimeManager) AppendScriptFiles() error {
	fileList := []string{}
	err := filepath.Walk(this.m_searchDir, func(path string, f os.FileInfo, err error) error {
		fileList = append(fileList, path)
		return nil
	})

	for _, file := range fileList {
		// Only .js are valid extension here...
		if strings.HasSuffix(file, ".js") {
			err = this.AppendScriptFile(file)
			if err != nil {
				return err
			}
		}
	}
	return err
}

/**
 * Append function in the JS.
 */
func (this *JsRuntimeManager) AppendFunction(name string, function interface{}) {
	// I will compile the script and set it in each session...
	this.m_functions[name] = function
}

/**
 * Compile and run a given script...
 */
func (this *JsRuntimeManager) AppendScriptFile(filePath string) error {
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
	this.AppendScript(src)

	return err
}

func (this *JsRuntimeManager) AppendScript(src string) {
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
			functionName = "fct_" + strings.Replace(Utility.RandomUUID(), "-", "_", -1)
			functionStr = "function " + functionName + strings.TrimSpace(functionStr)[8:]
		}
		_, err = vm.Run(functionStr)
		if err != nil {
			log.Println("Error in code of ", functionName)
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
		return nil, err
	}

	// Append val to results...
	results = append(results, val)

	return results, err
}

/**
 * Append and excute a javacript function on the JS...
 */
func (this *JsRuntimeManager) ExecuteJsFunction(messageId string, sessionId string, functionStr string, functionParams []interface{}) (results []interface{}, err error) {
	// Set the function on the JS runtime...

	// Put function call information into a struct.
	var jsFunctionInfos JsFunctionInfos
	jsFunctionInfos.m_messageId = messageId
	jsFunctionInfos.m_sessionId = sessionId
	jsFunctionInfos.m_functionStr = functionStr
	jsFunctionInfos.m_functionParams = functionParams

	log.Println("----------> 307")
	this.m_channels[sessionId] <- jsFunctionInfos

	log.Println("----------> 309")
	jsFunctionInfos = <-this.m_channels[sessionId]

	// with the answer...
	log.Println("----------> 310", jsFunctionInfos.m_messageId)

	return
}
