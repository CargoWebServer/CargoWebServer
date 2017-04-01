package main

import (
	"log"
	"net"
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

	// Test values...
	//	ClientId:     "1234",
	// 	ClientSecret: "aabbccdd",
	//	AuthorizeUrl: "http://localhost:9393/authorize",
	//	TokenUrl:     "http://localhost:9393/token",
	//	RedirectUrl:  "http://localhost:9393/oauth2callback"

	// OAuth2 http handler's
	http.HandleFunc("/authorize", Server.AuthorizeHandler)
	http.HandleFunc("/token", Server.TokenHandler)
	http.HandleFunc("/info", Server.InfoHandler)

	// OpenId
	http.HandleFunc("/.well-known/openid-configuration", Server.DiscoveryHandler)
	http.HandleFunc("/publickeys", Server.PublicKeysHandler)

	// Client redirect address.
	http.HandleFunc("/oauth2callback", Server.AppAuthCodeHandler)

	// stop the server...
	defer Server.GetServer().Stop()

	// Start the server...
	Server.GetServer().Start()

	open.Run("http://127.0.0.1:9393/Bridge")

	listener, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}

	log.Println("Server listen on Port:", port)
	http.Serve(listener, nil)

}
