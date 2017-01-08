package main

import (
	"log"
	"net/http"
	"strconv"

	"code.google.com/p/go.net/websocket"
	"code.myceliUs.com/CargoWebServer/Cargo/Server"
	//	"github.com/pkg/profile"
)

func main() {

	// Handle application path...
	root := Server.GetServer().GetConfigurationManager().GetApplicationDirectoryPath()
	port := Server.GetServer().GetConfigurationManager().GetPort()
	log.Println("Start serve files from ", root)

	// Start the web socket handler
	http.Handle("/ws", websocket.Handler(Server.HttpHandler))

	// The http handler
	http.Handle("/", http.FileServer(http.Dir(root)))

	// The file upload handler.
	http.HandleFunc("/uploads", Server.FileUploadHandler)

	// Start the server...
	Server.GetServer().Start()
	log.Println("Port:", port)
	err := http.ListenAndServe(":"+strconv.Itoa(port), nil)

	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
