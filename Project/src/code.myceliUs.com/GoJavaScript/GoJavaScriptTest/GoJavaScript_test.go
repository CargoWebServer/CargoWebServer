package GoJavaScript

import "testing"
import "log"

import "code.myceliUs.com/GoJavaScript"
import "code.myceliUs.com/GoJavaScript/GoJavaScriptClient"

type GenderType int

const (
	Male GenderType = 1 + iota
	Female
)

// A simple Go struct that contain array, properties and constant values.
type Person struct {
	// Must be GoJavaScript.Person
	TYPENAME  string
	FirstName string
	LastName  string
	Age       int
	NickNames []string
	// List of contact.
	Contacts []*Person

	// Gender
	Gender GenderType
}

// Create a person and return it as a reasult.
func GetPerson() *Person {

	var p *Person
	p = new(Person)
	p.TYPENAME = "GoJavaScriptTest.Person"
	p.Age = 42
	p.FirstName = "Dave"
	p.LastName = "Courtois"
	p.NickNames = make([]string, 0)
	p.NickNames = append(p.NickNames, "Natural")
	p.Gender = Male

	p.Contacts = make([]*Person, 0)

	// Append a contact.
	c0 := new(Person)
	c0.TYPENAME = "GoJavaScriptTest.Person"
	c0.Age = 21
	c0.FirstName = "Emmanuel"
	c0.LastName = "Proulx"
	c0.Gender = Male
	p.Contacts = append(p.Contacts, c0)

	c1 := new(Person)
	c1.TYPENAME = "GoJavaScriptTest.Person"
	c1.Age = 42
	c1.FirstName = "Eric"
	c1.LastName = "Boucher"
	c1.Gender = Male
	p.Contacts = append(p.Contacts, c1)

	log.Println("---> get Person was call!", p)
	return p
}

// Function that return a string
func (self *Person) Name() string {
	log.Println("---> name was called ", self.FirstName+" "+self.LastName)
	return self.FirstName + " " + self.LastName
}

// Return a object, also recursive.
func (self *Person) Myself() *Person {
	return self
}

// A method that take an object as parameter.
func (self *Person) SayHelloTo(to *Person) string {
	return self.FirstName + " say hello to " + to.Name() + "!"
}

// A function with an array of object as parameter.
func (self *Person) SayHelloToAll(all []*Person) string {
	var greathing string

	for i := 0; i < len(all); i++ {
		greathing += self.FirstName + " say hello to " + all[i].Name() + "!\n"
	}

	return greathing
}

// A method that return an array.
func (self *Person) GetContacts() []*Person {
	log.Println("---> get contact was call!", self.Contacts)
	return self.Contacts
}

// Simple function to test adding tow number in Go
// that function will be call inside JS via the handler.
func AddNumber(a float64, b float64) float64 {
	return a + b
}

// only ducktape is avalaible a this time.
var engine = GoJavaScriptClient.NewClient("127.0.0.1", 8081, "duktape")

/**
 * Simple Hello world test.
 */
func TestHelloJava(t *testing.T) {

	// Register the function SayHelloTo. The function take one parameter.
	engine.RegisterJsFunction("SayHelloTo", "function SayHelloTo(greething, to){return greething + ' ' + to + '!';}")
	str, _ := engine.CallFunction("SayHelloTo", "Hello", "Java Script")

	if str.(string) != "Hello Java Script!" {
		t.Error("Expected 'Hello Java Script!', got ", str)
	} else {
		// display hello jerry!
		log.Println(str)
	}
}

/*
 * Test numeric function and use EvalScript instead of CallFunction...
 */
func TestNumericValue(t *testing.T) {

	// Create a remote javascript server.
	engine.RegisterJsFunction("Add", "function Add(a, b){return a + b;}")
	a := GoJavaScript.NewVariable("a", 1)
	b := GoJavaScript.NewVariable("b", 2.25)
	number, _ := engine.EvalScript("Add(a, b);", []interface{}{a, b})

	if number.(float64) != 3.25 {
		t.Error("Expected 3.25, got ", number)
	}
}

/**
 * Test function with boolean values.
 */
func TestBooleanValue(t *testing.T) {
	engine.RegisterJsFunction("TestBool", "function TestBool(val){return val>0;}")
	boolean, _ := engine.CallFunction("TestBool", 1)

	if boolean.(bool) == false {
		t.Error("Expected true, got ", boolean)
	}
}

/**
 * Test playing with array value.
 */
func TestArray(t *testing.T) {
	engine.RegisterJsFunction("TestArray", "function TestArray(arr, val){arr.push(val); return arr;}")

	// Array of anything is supported as you can see.
	arr, err0 := engine.CallFunction("TestArray", []interface{}{1.0, 3.0, 4.0}, 2.25)
	if err0 == nil {
		t.Log(arr)
	}
}

/**
 * Test calling a go function from JS.
 */
func TestGoFunction(t *testing.T) {
	// First of all I will register the tow go function in the Engine.
	engine.RegisterGoFunction("AddNumber", AddNumber)
	engine.RegisterJsFunction("TestAddNumber", `function TestAddNumber(){var result = AddNumber(3, 8); console.log("The result is:" + result); return result;}`)
	addNumberResult, err := engine.CallFunction("TestAddNumber")

	if err == nil {
		t.Log("Add number result: ", addNumberResult)
	}
}

// Test global variable, Set and Get
func TestGlobalVariable(t *testing.T) {

	var toto = "This is Java"
	engine.SetGlobalVariable("toto", toto)

	toto_, _ := engine.GetGlobalVariable("toto")

	if toto_ != toto {
		t.Log("Set/Get global variables fail! ")
	}
}

/**
 * Test creating JavaScript object with go function and Js method.
 */
func TestCreateJsObjectFromGo(t *testing.T) {

	// First of all I will create the object.
	obj := engine.CreateObject("test")

	// Set a property on test.
	obj.Set("number", 1.01)

	number, err := obj.Get("number")
	if err == nil {
		if number.(float64) != 1.01 {
			t.Error("---> fail to get object property!")
		}
	}

	// Now set a go function.
	obj.Set("add", AddNumber)

	// set a Js function on the object.
	obj.SetJsMethode("helloTo", `function helloTo(to){return "Hello " + to + "!";}`)

	// and call the go function on the object.
	addReuslt, _ := obj.Call("add", 2, 3)
	if addReuslt.(float64) != 5.0 {
		t.Error("---> fail to get object property!")
	}

	helloToResult, _ := obj.Call("helloTo", "Java")
	if helloToResult.(string) != "Hello Java!" {
		t.Error("---> fail to set js property!")
	}
}

// Test using go object created from go function and from global variable.
func TestRegisterGoObject(t *testing.T) {

	// Create the object to register.
	p := GetPerson()

	// Here I will register a go Object in JavaScript and set
	// it as global variable named Dave.
	engine.SetGlobalVariable("Dave", p)

	// Now I will eval sricpt on it...
	engine.RegisterJsFunction("Test1", `function Test1(){console.log('Hello ' + Dave.Name() + ' your first contacts is ' + Dave.GetContacts()[0].Name())}`)

	// Eval script that contain Go object in it.
	engine.EvalScript("Test1();", []interface{}{})

	// Eval single return type (not array)
	engine.RegisterGoFunction("GetPerson", GetPerson)

	engine.EvalScript("console.log('Hello: ' + GetPerson().Name() + ' Your age are ' + GetPerson().Age + ' ' + GetPerson().SayHelloTo(Dave))", []interface{}{})

	// Eval array...
	engine.EvalScript("console.log(Dave.SayHelloToAll(GetPerson().GetContacts()))", []interface{}{})

	// Eval object chain call...
	engine.EvalScript("console.log('I am ' + GetPerson().Myself().Myself().Myself().Name() + '!')", []interface{}{})

	// I will now register a function that call GetPerson in it.
	engine.EvalScript(`function SayHelloToContact(index){console.log('---> '+ GetPerson().SayHelloTo(GetPerson().GetContacts()[index].Myself()))}`, []interface{}{})
	engine.CallFunction("SayHelloToContact", []interface{}{1})
}

/**
 * Test creating a Go object from Javascript.
 * ** The Class, here Person, must contain a field TYPENAME and be register
 *    with RegisterGoType
 */
func TestCreateGoObjectFromJs(t *testing.T) {

	// Test with structure
	engine.RegisterJsFunction("TestJsToGoStruct", `function TestJsToGoStruct(){var jerry = {TYPENAME:"GoJavaScriptTest.Person", FirstName:"Java", LastName:"Script", Age:20, NickNames:["toto", "titi", "tata"]}; console.log('---> TestJsToGoStruct ' + jerry ); return jerry; }`)
	p, err := engine.EvalScript("TestJsToGoStruct();", []interface{}{})

	if err != nil {
		t.Error("fail to create Go from Js: ", err)
	}

	if err == nil {
		t.Log(p)
	} else {
		t.Error("test fail!")
	}
}

/**
 * Test using javascript object in go function.
 */
func TestJsObjInGo(t *testing.T) {

	// The test function will be use inside a JsClass method.
	engine.RegisterGoFunction("test_", func(obj GoJavaScript.Object) {
		// so here obj is call from toto.test();
		// I will try to call sayHello on it.
		obj.Call("sayHello", "Go!")
	})

	// a simple javacript function with a simple function.
	src := "var JsClass = function(name){ this.name = name; this.sayMyName = function(){console.log('my name is ', name);}; return this; };"
	src += "JsClass.prototype.sayHello = function(from){console.log('hello ', this.name, ' from ', from);};"
	src += "JsClass.prototype.test = function(){test_(this);};"

	// so here i will eval the script.
	engine.EvalScript(src, []interface{}{})

	// Try call toto.test()
	engine.EvalScript("var toto = new JsClass('toto'); toto.sayMyName(); toto.test();", []interface{}{})
}

/**
 * Test js callback function as Go function parameter.
 */
func TestCallback(t *testing.T) {

	// test a go function that take a Js function as parameter.
	engine.RegisterGoFunction("go_callback", func(engine *GoJavaScriptClient.Client) func(string) {
		// Callback is a javascript bytecode containing the function pass ass go function
		// parameter.
		return func(callback string) {
			engine.CallFunction(callback, []interface{}{"GoJavaScript"})
		}
	}(engine))

	// a simple javacript function with a simple function.
	src := `go_callback(function(name){console.log("I'm " + name + "!")});`

	// so here i will eval the script.
	engine.EvalScript(src, []interface{}{})

	// Now test a Js function that take a go function as callback.
	src = `function fct0(fct1, fct2){fct1(fct2);}; fct0(go_callback, function(name){console.log("I'm " + name + "!")})`
	engine.EvalScript(src, []interface{}{})

	// Now I will test a go function that take another go funciton as parameter.
	engine.RegisterGoFunction("fct0", func(fct interface{}) {
		// fct is the function itself pass as interface so I will cast it
		// and call it with value "GoJavaScript".
		fct.(func(toto string))("GoJavaScript!")
	})

	engine.RegisterGoFunction("fct1", func(toto string) {
		log.Println("---> call go function ftc0 with value ", toto)
	})

	// Here the function fct0 is call with fct1 as parameter...
	src = `fct0(fct1);`
	engine.EvalScript(src, []interface{}{})
}

func TestStopJava(t *testing.T) {
	engine.Stop()
}
