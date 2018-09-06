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

// Golang struct...
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
	p.TYPENAME = "GoJavaScript.Person"
	p.Age = 42
	p.FirstName = "Dave"
	p.LastName = "Courtois"
	p.NickNames = make([]string, 0)
	p.NickNames = append(p.NickNames, "Natural")
	p.Gender = Male

	p.Contacts = make([]*Person, 0)

	// Append a contact.
	c0 := new(Person)
	c0.TYPENAME = "GoJavaScript.Person"
	c0.Age = 21
	c0.FirstName = "Emmanuel"
	c0.LastName = "Proulx"
	c0.Gender = Male
	p.Contacts = append(p.Contacts, c0)

	c1 := new(Person)
	c1.TYPENAME = "GoJavaScript.Person"
	c1.Age = 42
	c1.FirstName = "Eric"
	c1.LastName = "Boucher"
	c1.Gender = Male
	p.Contacts = append(p.Contacts, c1)

	log.Println("---> get Person was call!", p)
	return p
}

// A simple function.
func (self *Person) Name() string {
	log.Println("---> name was called ", self.FirstName+" "+self.LastName)
	return self.FirstName + " " + self.LastName
}

// A method that use a Go type as pointer.
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

// Simple function handler.
func PrintValue(value interface{}) {
	log.Println(value)
}

var engine = GoJavaScriptClient.NewClient("127.0.0.1", 8081, "jerryscript" /*, "chakracore"*/)

/**
 * Simple Hello world test.
 * Start JavaScript interpreter append simple script and run it.
 */
func TestHelloJava(t *testing.T) {

	// Register the function SayHelloTo. The function take one parameter.
	engine.RegisterJsFunction("SayHelloTo", "function SayHelloTo(greething, to){return greething + ' ' + to + '!';}")
	str, _ := engine.CallFunction("SayHelloTo", "Hello", "Java Script")
	str_, _ := str.ToString()

	if str_ != "Hello Java Script!" {
		t.Error("Expected 'Hello Java Script!', got ", str)
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
	a := GoJavaScript.NewVariable("a", 1)
	b := GoJavaScript.NewVariable("b", 2.25)
	number, _ := engine.EvalScript("Add(a, b);", []interface{}{a, b})
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
	var toto = "This is Java"
	engine.SetGlobalVariable("toto", toto)

	toto_, _ := engine.GetGlobalVariable("toto")
	toto__, _ := toto_.Export()

	if toto__ != toto {
		t.Log("Set/Get global variables fail! ")
	}

}

func TestCreateJsObjectFromGo(t *testing.T) {

	engine.RegisterGoFunction("print", PrintValue)

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
	obj.SetJsMethode("helloTo", `function helloTo(to){print("Hello " + to + "!"); return "Hello " + to + "!";}`)

	helloToResult, _ := obj.Call("helloTo", "Java")
	helloToResult_, _ := helloToResult.ToString()
	if helloToResult_ != "Hello Java!" {
		t.Error("---> fail to set js property!")
	}

}

func TestRegisterGoObject(t *testing.T) {

	engine.RegisterGoFunction("print", PrintValue)

	// Create the object to register.
	engine.RegisterGoType((*Person)(nil))

	p := GetPerson()

	// Here I will register a go Object in JavaScript and set
	// it as global variable named Dave.
	engine.SetGlobalVariable("Dave", p)

	// Now I will eval sricpt on it...
	engine.RegisterJsFunction("Test1", `function Test1(){print('Hello ' + Dave.Name() + ' your first contacts is ' + Dave.GetContacts()[0].Name())}`)

	// Eval script that contain Go object in it.
	engine.EvalScript("Test1();", []interface{}{})

	// Now I will try to register a function that return a Go type and use it
	// result to access it contact.
	engine.RegisterGoFunction("GetPerson", GetPerson)

	// Eval single return type (not array)
	engine.EvalScript("print('Hello: ' + GetPerson().Name() + ' Your age are ' + GetPerson().Age + ' ' + GetPerson().SayHelloTo(Dave))", []interface{}{})

	// Eval array...
	engine.EvalScript("print(Dave.SayHelloToAll(GetPerson().GetContacts()))", []interface{}{})
}

/**
 * Test calling a go function from JS.
 */
func TestCreateGoObjectFromJs(t *testing.T) {

	// Test with structure
	// The type must be register before being usable by the vm.
	engine.RegisterGoType((*Person)(nil))

	// Register the dynamic type.
	engine.RegisterJsFunction("TestJsToGoStruct", `function TestJsToGoStruct(){var jerry = {TYPENAME:"GoJavaScriptTest.Person", FirstName:"Java", LastName:"Script", Age:20, NickNames:["toto", "titi", "tata"]}; return jerry; }`)
	p, err := engine.EvalScript("TestJsToGoStruct();", []interface{}{})

	if err != nil {
		t.Error("fail to create Go from Js: ", err)
	}

	p_, _ := p.Export()

	if err == nil {
		t.Log(p_)
	} else {
		t.Error("test fail!")
	}
}

func TestStopJava(t *testing.T) {
	engine.Stop()
}
