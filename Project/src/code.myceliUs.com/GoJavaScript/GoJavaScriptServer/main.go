// GoJavaScriptServer project main.go
package main

import "os"
import "strconv"

// Process command.
func main() {
	port, _ := strconv.Atoi(os.Args[1])
	server := NewServer("127.0.0.1", port, os.Args[2])
	server.processActions()
}
