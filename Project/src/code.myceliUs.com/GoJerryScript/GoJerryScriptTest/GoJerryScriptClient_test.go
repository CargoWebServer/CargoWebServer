package GoJerryScript

import "testing"
import "log"

//import "code.myceliUs.com/GoJerryScript"
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

// Simple function to test adding tow number in Go
// that function will be call inside JS via the handler.
func AddNumber(a float64, b float64) float64 {
	return a + b
}

// Simple function handler.
func PrintValue(value interface{}) {
	log.Println(value)
}

///**
// * Simple Hello world test.
// * Start JerryScript interpreter append simple script and run it.
// */
//func TestHelloJerry(t *testing.T) {

//	// Create a remote javascript server.
//	engine := GoJerryScriptClient.NewClient("127.0.0.1", 8081)

//	// Register the function SayHelloTo. The function take one parameter.
//	engine.RegisterJsFunction("SayHelloTo", "function SayHelloTo(to){return 'Hello ' + to + '!';}")

//	str, _ := engine.EvalScript("SayHelloTo(jerry);", GoJerryScript.Variables{{Name: "jerry", Value: "Jerry Script"}})
//	str_, _ := str.ToString()

//	if str_ != "Hello Jerry Script!" {
//		t.Error("Expected 'Hello Jerry Script!', got ", str)
//	} else {
//		// display hello jerry!
//		log.Println(str_)
//	}

//	// Stop the engine
//	engine.Stop()
//}

///*
// * Test function with numeric values.
// */
//func TestNumericValue(t *testing.T) {

//	// Create a remote javascript server.
//	engine := GoJerryScriptClient.NewClient("127.0.0.1", 8081)
//	engine.RegisterJsFunction("Add", "function Add(a, b){return a + b;}")

//	number, _ := engine.EvalScript("Add(a, b);", GoJerryScript.Variables{{Name: "a", Value: 1}, {Name: "b", Value: 2.25}})
//	number_, _ := number.Export()

//	if number_ != 3.25 {
//		t.Error("Expected 3.25, got ", number_)
//	}

//	// Stop the engine
//	engine.Stop()
//}

///**
// * Test function with boolean values.
// */
//func TestBooleanValue(t *testing.T) {

//	// Create a remote javascript server.
//	engine := GoJerryScriptClient.NewClient("127.0.0.1", 8081)

//	engine.RegisterJsFunction("TestBool", "function TestBool(val){return val>0;}")
//	boolean, _ := engine.EvalScript("TestBool(val)", GoJerryScript.Variables{{Name: "val", Value: 1}})
//	boolean_, _ := boolean.ToBoolean()
//	if boolean_ == false {
//		t.Error("Expected true, got ", boolean)
//	}

//	// Stop the engine
//	engine.Stop()
//}

///**
// * Test playing with array value.
// */
//func TestArray(t *testing.T) {

//	// Create a remote javascript server.
//	engine := GoJerryScriptClient.NewClient("127.0.0.1", 8081)

//	engine.RegisterJsFunction("TestArray", "function TestArray(arr, val){arr.push(val); return arr;}")

//	arr, err0 := engine.EvalScript("TestArray(arr, val);", GoJerryScript.Variables{{Name: "arr", Value: []interface{}{1.0, 3.0, 4.0}}, {Name: "val", Value: 2.25}})
//	if err0 == nil {
//		t.Log(arr)
//	}

//	// Stop the engine
//	engine.Stop()
//}

///**
// * Test calling a go function from JS.
// */
//func TestGoFunction(t *testing.T) {

//	// Create a remote javascript server.
//	engine := GoJerryScriptClient.NewClient("127.0.0.1", 8081)

//	// First of all I will register the tow go function in the Engine.
//	engine.RegisterGoFunction("AddNumber", AddNumber)
//	engine.RegisterGoFunction("Print", PrintValue)
//	engine.RegisterJsFunction("TestAddNumber", `function TestAddNumber(){var result = AddNumber(3, 8); Print("The result is:" + result); return result;}`)

//	addNumberResult, err := engine.EvalScript("TestAddNumber();", GoJerryScript.Variables{})

//	if err == nil {
//		t.Log("Add number result: ", addNumberResult)
//	}

//	// Stop the engine
//	engine.Stop()
//}

/**
 * Test calling a go function from JS.
 */
/*func TestCreateGoObjectFromJs(t *testing.T) {

	// Create a remote javascript server.
	engine := GoJerryScriptClient.NewClient("127.0.0.1", 8081)

	// Test with structure
	// The type must be register before being usable by the vm.
	engine.RegisterGoType((*Person)(nil))

	// Register the dynamic type.
	engine.RegisterJsFunction("TestJsToGoStruct", `function TestJsToGoStruct(){var jerry = {TYPENAME:"GoJerryScript.Person", FirstName:"Jerry", LastName:"Script", Age:20, NickNames:["toto", "titi", "tata"]}; return jerry; }`)
	p, err1 := engine.EvalScript("TestJsToGoStruct();", []GoJerryScript.Variable{})

	p_, _ := p.Export()
	if err1 == nil {
		t.Log(p_)
	}

	// Stop the engine
	engine.Stop()
}*/

func TestCreateJsObjectFromGo(t *testing.T) {

	engine := GoJerryScriptClient.NewClient("127.0.0.1", 8081)

	// First of all I will create the object.
	obj := engine.CreatObject("test")

	// Set a property on test.
	obj.Set("number", 1.01)

	number, err := obj.Get("number")
	if err == nil {
		number_, _ := number.ToFloat()
		if number_ != 1.01 {
			t.Error("---> fail to get object property!")
		}
	}

	// Now I will try to set a go function.
	obj.Set("add", AddNumber)

	// and call the go function on the object.
	addReuslt, _ := obj.Call("add", 2, 3)
	addReuslt_, _ := addReuslt.ToFloat()
	if addReuslt_ != 5.0 {
		t.Error("---> fail to get object property!")
	}

	obj.SetJsMethode("helloTo", []string{"to"}, `function helloTo(to){return "Hello " + to + "!";}`)

	helloToResult, _ := obj.Call("helloTo", "Jerry")
	helloToResult_, _ := helloToResult.ToString()
	if helloToResult_ != "Hello Jerry!" {
		t.Error("---> fail to set js property!")
	}

	// Stop the engine
	engine.Stop()
}

/*func TestCallJsFunction(t *testing.T) {
	// Test Call function...
	engine := GoJerryScriptClient.NewClient("127.0.0.1", 8081)
	engine.RegisterGoFunction("Add", AddNumber)

	r, err5 := engine.CallFunction("Add", []interface{}{215, 5})
	if err5 == nil {
		result, _ := r.Export()
		log.Println("151 ---> ", result)
	}
	// Stop the engine
	engine.Stop()
}*/

/*
func TestEvalScript(t *testing.T) {

	engine := NewEngine(9696, JERRY_INIT_EMPTY)
	engine.RegisterGoFunction("Print", PrintValue)



	// Now I will try to call go function from Js.


	// Now I will share a go object with JS engine.
	engine.SetGlobalVariable("person", p_)

	//* Note that person properties are accessible via Getter/Setter only, that's because
	//  the object
	engine.AppendJsFunction("TestGoObject", []interface{}{}, `function TestGoObject(){person.Age=42; Print("Hello " + person.FirstName + " your age is " + person.Age + " full name: " + person.Name());person.NickNames = person.NickNames.concat(["test1", "test2"]); Print(person.NickNames);  Print(person.SayHelloTo("World!")); return person;}`)
	p__, err3 := engine.EvalScript("TestGoObject();", []Variable{})
	if err3 == nil {
		t.Log("Go object: ", p__)
	}

	// Test Call function...
	r, err5 := engine.CallFunction("Add", []interface{}{215, 5})
	if err5 == nil {
		result, _ := r.Export()
		log.Println("151 ---> ", result)
	}

	// I will test a more structured resutls...
	engine.RegisterGoFunction("GetPerson", GetPerson)

	// Now I will run the scipt...
	engine.EvalScript("Print('Hello ' + GetPerson().Contacts[0].Name()); Print('Object person has uuid: ' + GetPerson().uuid_ + ' and his contact has: ' + GetPerson().Contacts[0].uuid_ )", []Variable{})

	engine.Clear()
}
*/
