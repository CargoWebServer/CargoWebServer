package JS

import (
	"bytes"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"code.myceliUs.com/CargoWebServer/Cargo/Utility"
	"github.com/robertkrimen/otto"
	_ "github.com/robertkrimen/otto/underscore"
)

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

	/** Contain the js/go function mapping **/
	m_functors map[string]interface{}

	/** One js interpreter by session */
	m_session map[string]*otto.Otto

	/**
	 * Use to protected the entitiesMap access...
	 */
	sync.RWMutex
}

func NewJsRuntimeManager(searchDir string) *JsRuntimeManager {
	// assign memory...
	jsRuntimeManager = new(JsRuntimeManager)
	jsRuntimeManager.m_searchDir = searchDir
	jsRuntimeManager.m_session = make(map[string]*otto.Otto)
	jsRuntimeManager.m_functors = make(map[string]interface{})

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
func (this *JsRuntimeManager) GetVm(sessionId string) *otto.Otto {
	// Protectect the map access...
	this.Lock()
	defer this.Unlock()

	// If the vm exist, i will simply return the vm...
	if vm, ok := this.m_session[sessionId]; ok {
		return vm
	}

	// Create a new js interpreter for the given session.
	this.m_session[sessionId] = otto.New()

	// Compile the list of script...
	for i := 0; i < len(this.m_scripts); i++ {
		script, err := this.m_session[sessionId].Compile("", this.m_scripts[i])
		this.m_session[sessionId].Run(script)
		if err != nil {
			log.Fatal("runtime script compilation error:", err)
		}
	}

	// Create the list of functors...
	for prototype, functor := range this.m_functors {
		this.m_session[sessionId].Set(prototype, functor)
	}

	return this.m_session[sessionId]
}

func (this *JsRuntimeManager) CloseSession(sessionId string) {
	// Protectect the map access...
	this.Lock()
	defer this.Unlock()

	// simply remove the js interpreter...
	delete(this.m_session, sessionId)
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
}

/**
 * Append and excute a javacript function on the JS...
 */
func (this *JsRuntimeManager) ExecuteJsFunction(messageId string, sessionId string, functionStr string, functionParams []interface{}) (results []interface{}, err error) {

	// Here i wil find the name of the function...
	startIndex := strings.Index(functionStr, " ")
	endIndex := strings.Index(functionStr, "(")

	functionName := strings.Trim(functionStr[startIndex:endIndex], " ")

	// Set the function on the JS runtime...
	vm := this.GetVm(sessionId).Copy()

	// Set the current session id.
	vm.Set("sessionId", sessionId)
	vm.Set("messageId", messageId)
	
	_, err = vm.Run(functionStr)
	if err != nil {
		log.Println("Error in code of ", functionName)
		return nil, err
	}

	var params []interface{}
	params = append(params, functionName)
	params = append(params, nil)

	// Now I will make otto digest the parameters...
	for i := 0; i < len(functionParams); i++ {
		//log.Println("parameter: ", functionParams[i])
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
