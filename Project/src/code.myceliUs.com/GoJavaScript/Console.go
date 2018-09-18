package GoJavaScript

import "log"
import "fmt"

/**
 * The Console object provides access to the server debugging console.
 */
type Console struct {
	// use by group...
	level    int
	TYPENAME string
}

// Output a message to the console
func (self *Console) Log(values ...interface{}) {
	indent := ""
	for i := 0; i < self.level; i++ {
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
func (self *Console) Group() {
	self.level++
}

// Exits the current inline group in the console
func (self *Console) GroupEnd() {
	if self.level > 0 {
		self.level--
	}
}

// Clear the console
func (self *Console) Clear() {
	self.level = 0
	fmt.Println("\033[H\033[2J")
}
