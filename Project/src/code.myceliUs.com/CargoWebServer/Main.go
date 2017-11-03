package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"code.myceliUs.com/CargoWebServer/Cargo/Server"
	//	"github.com/skratchdot/open-golang/open"
	"golang.org/x/net/websocket"
)

func main() {

	// Handle application path...
	root := Server.GetServer().GetConfigurationManager().GetApplicationDirectoryPath()
	port := Server.GetServer().GetConfigurationManager().GetServerPort()

	// Start the web socket handler
	http.Handle("/ws", websocket.Handler(Server.HttpHandler))

	// The http handler
	http.Handle("/", http.FileServer(http.Dir(root)))

	// The file upload handler.
	http.HandleFunc("/uploads", Server.FileUploadHandler)

	// The http query handler use by external http client or OAuth2
	http.HandleFunc("/api/", Server.HttpQueryHandler)

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

	// Client redirect address.
	http.HandleFunc("/oauth2callback", Server.AppAuthCodeHandler)

	// OpenId service.
	http.HandleFunc("/.well-known/openid-configuration", Server.DiscoveryHandler)
	http.HandleFunc("/publickeys", Server.PublicKeysHandler)

	// In case of interuption of the program i will
	// stop service before return.
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		Server.GetServer().Stop()
		os.Exit(1)
	}()

	// Print the error message in the console in case of panic.
	defer func() {
		log.Println(recover()) // 1
	}()

	// Start the server...
	Server.GetServer().Start()

	//open.Run("http://127.0.0.1:9393/Bridge")
	log.Println("--> server is ready and listen at port ", port)
	err := http.ListenAndServe(":"+strconv.Itoa(port), nil)

	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}

}
