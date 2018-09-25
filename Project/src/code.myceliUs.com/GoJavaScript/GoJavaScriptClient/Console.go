package GoJavaScriptClient

import "log"
import "code.myceliUs.com/GoJavaScript"
import "fmt"

var (
	// The indentation level for the group.
	level = 0

	// Set the console.
	console *GoJavaScript.Object
)

func getConsole() *GoJavaScript.Object {
	if console == nil {
		console = GoJavaScript.NewObject("console")

		// Set console functions.
		console.Set("log", ConsoleLog)
		console.Set("group", ConsoleGroup)
		console.Set("groupEnd", ConsoleGroupEnd)
		console.Set("clear", ConsoleClear)
	}

	return console
}

// Output a message to the console
func ConsoleLog(values ...interface{}) {
	indent := ""
	for i := 0; i < level; i++ {
		indent += "	"
	}

	// Create a new array and append the le indentation in it
	values_ := make([]interface{}, 0)
	values_ = append(values_, indent)
	values_ = append(values_, values)

	log.Println(values...)
}

// Create a new inline group in the console. This indents following console
// message by additional level, until console.groupEnd is called
func ConsoleGroup() {
	level++
}

// Exits the current inline group in the console
func ConsoleGroupEnd() {
	if level > 0 {
		level--
	}
}

// Clear the console
func ConsoleClear() {
	level = 0
	fmt.Println("\033[H\033[2J")
}
