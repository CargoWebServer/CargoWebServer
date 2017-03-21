package main

import (
	"log"
	"net/http"
	"strconv"

	"code.myceliUs.com/CargoWebServer/Cargo/Server"
	"github.com/skratchdot/open-golang/open"
	"golang.org/x/net/websocket"
)

func main() {

	// Handle application path...
	root := Server.GetServer().GetConfigurationManager().GetApplicationDirectoryPath()
	port := Server.GetServer().GetConfigurationManager().GetServerPort()

	log.Println("Start serve files from ", root)

	// Start the web socket handler
	http.Handle("/ws", websocket.Handler(Server.HttpHandler))

	// The http handler
	http.Handle("/", http.FileServer(http.Dir(root)))

	// The file upload handler.
	http.HandleFunc("/uploads", Server.FileUploadHandler)

	// OAuth2 http handler's
	http.HandleFunc("/authorize", Server.AuthorizeHandler)
	http.HandleFunc("/token", Server.TokenHandler)
	http.HandleFunc("/info", Server.InfoHandler)
	http.HandleFunc("/app", Server.AppHandler)
	http.HandleFunc("/appauth/code", Server.AppAuthCodeHandler)

	// stop the server...
	defer Server.GetServer().Stop()

	// Start the server...
	Server.GetServer().Start()
	log.Println("Port:", port)
	open.Run("http://127.0.0.1:9393/Bridge")

	err := http.ListenAndServe(":"+strconv.Itoa(port), nil)

	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}

}
