package GoJerryScriptTest

import "testing"
import "log"

import "code.myceliUs.com/GoJerryScript"
import "code.myceliUs.com/GoJerryScript/GoJerryScriptClient"

// Golang struct...
type Person struct {
	// Must be GoJerryScript.Person
	TYPENAME  string
	FirstName string
	LastName  string
	Age       int
	NickNames []string

	Contacts []*Person
}

// Create a person and return it as a reasult.
func GetPerson() *Person {

	var p *Person
	p = new(Person)
	p.TYPENAME = "GoJerryScript.Person"
	p.Age = 40
	p.FirstName = "Dave"
	p.LastName = "Courtois"
	p.NickNames = make([]string, 0)
	p.NickNames = append(p.NickNames, "Natural")

	p.Contacts = make([]*Person, 0)

	// Append a contact.
	c := new(Person)
	c.TYPENAME = "GoJerryScript.Person"
	c.Age = 21
	c.FirstName = "Emmanuel"
	c.LastName = "Proulx"

	p.Contacts = append(p.Contacts, c)

	return p
}

// A simple function.
func (self *Person) Name() string {
	return self.FirstName + " " + self.LastName
}

func (self *Person) SayHelloTo(to string) string {
	return self.FirstName + " say hello to " + to + "!"
}

// A method that return an array.
func (self *Person) GetContacts() []*Person {
	return self.Contacts
}

// Simple function to test adding tow number in Go
// that function will be call inside JS via the handler.
func AddNumber(a float64, b float64) float64 {
	return a + b
}

// Simple function handler.
func PrintValue(value interface{}) {
	log.Println(value)
}

var engine = GoJerryScriptClient.NewClient("127.0.0.1", 8081)

/**
 * Simple Hello world test.
 * Start JerryScript interpreter append simple script and run it.
 */
func TestHelloJerry(t *testing.T) {

	// Register the function SayHelloTo. The function take one parameter.
	engine.RegisterJsFunction("SayHelloTo", "function SayHelloTo(to){return 'Hello ' + to + '!';}")
	str, _ := engine.CallFunction("SayHelloTo", "Jerry Script")

	str_, _ := str.ToString()

	if str_ != "Hello Jerry Script!" {
		t.Error("Expected 'Hello Jerry Script!', got ", str)
	} else {
		// display hello jerry!
		log.Println(str_)
	}

}

/*
 * Test numeric function and use EvalScript instead of CallFunction...
 */
func TestNumericValue(t *testing.T) {

	// Create a remote javascript server.
	engine.RegisterJsFunction("Add", "function Add(a, b){return a + b;}")
	number, _ := engine.EvalScript("Add(a, b);", GoJerryScript.Variables{{Name: "a", Value: 1}, {Name: "b", Value: 2.25}})
	number_, _ := number.Export()

	if number_ != 3.25 {
		t.Error("Expected 3.25, got ", number_)
	}

}

/**
 * Test function with boolean values.
 */
func TestBooleanValue(t *testing.T) {

	engine.RegisterJsFunction("TestBool", "function TestBool(val){return val>0;}")
	boolean, _ := engine.CallFunction("TestBool", 1)
	boolean_, _ := boolean.ToBoolean()
	if boolean_ == false {
		t.Error("Expected true, got ", boolean)
	}

}

/**
 * Test playing with array value.
 */
func TestArray(t *testing.T) {

	engine.RegisterJsFunction("TestArray", "function TestArray(arr, val){arr.push(val); return arr;}")
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
	engine.RegisterGoFunction("Print", PrintValue)
	engine.RegisterJsFunction("TestAddNumber", `function TestAddNumber(){var result = AddNumber(3, 8); Print("The result is:" + result); return result;}`)
	addNumberResult, err := engine.CallFunction("TestAddNumber")

	if err == nil {
		t.Log("Add number result: ", addNumberResult)
	}

}

func TestGlobalVariable(t *testing.T) {

	// First of all I will register the tow go function in the Engine.
	var toto = "This is Jerry"
	engine.SetGlobalVariable("toto", toto)

	toto_, _ := engine.GetGlobalVariable("toto")
	toto__, _ := toto_.Export()

	if toto__ != toto {
		t.Log("Set/Get global variables fail! ")
	}

}

func TestCreateJsObjectFromGo(t *testing.T) {

	// First of all I will create the object.
	obj := engine.CreateObject("test")

	// Set a property on test.
	obj.Set("number", 1.01)

	number, err := obj.Get("number")
	if err == nil {
		number_, _ := number.ToFloat()
		if number_ != 1.01 {
			t.Error("---> fail to get object property!")
		}
	}

	// Now set a go function.
	obj.Set("add", AddNumber)

	// and call the go function on the object.
	addReuslt, _ := obj.Call("add", 2, 3)
	addReuslt_, _ := addReuslt.ToFloat()
	if addReuslt_ != 5.0 {
		t.Error("---> fail to get object property!")
	}

	// set a Js function on the object.
	obj.SetJsMethode("helloTo", `function helloTo(to){return "Hello " + to + "!";}`)

	helloToResult, _ := obj.Call("helloTo", "Jerry")
	helloToResult_, _ := helloToResult.ToString()
	if helloToResult_ != "Hello Jerry!" {
		t.Error("---> fail to set js property!")
	}

}

func TestRegisterGoObject(t *testing.T) {

	engine.RegisterGoFunction("print", PrintValue)

	// Create the object to register.
	p := GetPerson()

	// Here I will register a go Object in JavaScript and set
	// it as global variable named Dave.
	engine.RegisterGoObject(p, "Dave")

	// Now I will eval sricpt on it...
	engine.RegisterJsFunction("Test1", `function Test1(){print('--> Hello ' + Dave.Name() + ' your first contacts is ' + Dave.GetContacts()[0].Name())}`)

	// Eval script that contain Go object in it.
	engine.EvalScript("Test1();", GoJerryScript.Variables{})

	// Now I will try to register a function that return a Go type and use it
	// result to access it contact.
	engine.RegisterGoFunction("GetPerson", GetPerson)

	// Eval single return type (not array)
	engine.EvalScript("print('Hello: ' + GetPerson().Name() + 'Your age are ' + GetPerson().Age)", []GoJerryScript.Variable{})

	// Eval array...
	engine.EvalScript("print('Hello ' + GetPerson().Contacts[0].Name())", []GoJerryScript.Variable{})

	//engine.Stop()
}

/**
 * Test calling a go function from JS.
 */
func TestCreateGoObjectFromJs(t *testing.T) {

	// Test with structure
	// The type must be register before being usable by the vm.
	engine.RegisterGoType((*Person)(nil))

	// Register the dynamic type.
	engine.RegisterJsFunction("TestJsToGoStruct", `function TestJsToGoStruct(){var jerry = {TYPENAME:"GoJerryScriptTest.Person", FirstName:"Jerry", LastName:"Script", Age:20, NickNames:["toto", "titi", "tata"]}; return jerry; }`)
	p, err := engine.EvalScript("TestJsToGoStruct();", []GoJerryScript.Variable{})

	if err != nil {
		t.Error("fail to create Go from Js: ", err)
	}

	p_, _ := p.Export()

	log.Println("----> p_", p_)
	if err == nil {
		t.Log(p_)
	}
}
