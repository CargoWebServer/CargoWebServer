// GoJavaScriptServer project main.go
package main

import "os"
import "strconv"

//import "log"

// Process command.
func main() {

	// Uncomment to dump log to file.
	/*f, err := os.OpenFile("./GoJavaScript_"+os.Args[1]+"_log.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}

	//defer to close when you're done with it, not because you think it's idiomatic!
	defer f.Close()

	//set output of logs to f
	log.SetOutput(f)*/

	port, _ := strconv.Atoi(os.Args[1])
	server := NewServer("127.0.0.1", port, os.Args[2])
	server.processActions()
}
