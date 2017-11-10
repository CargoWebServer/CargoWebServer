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
	"net/http/pprof"

	"golang.org/x/net/websocket"
)

func main() {

	// Handle application path...
	root := Server.GetServer().GetConfigurationManager().GetApplicationDirectoryPath()
	port := Server.GetServer().GetConfigurationManager().GetServerPort()

	r := http.NewServeMux()

	// Register pprof handlers
	r.HandleFunc("/debug/pprof/", pprof.Index)
	r.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	r.HandleFunc("/debug/pprof/profile", pprof.Profile)
	r.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	r.HandleFunc("/debug/pprof/trace", pprof.Trace)

	// Start the web socket handler
	r.Handle("/ws", websocket.Handler(Server.HttpHandler))

	// The http handler
	r.Handle("/", http.FileServer(http.Dir(root)))

	// The file upload handler.
	r.HandleFunc("/uploads", Server.FileUploadHandler)

	// The http query handler use by external http client or OAuth2
	r.HandleFunc("/api/", Server.HttpQueryHandler)

	// Test values...
	//	ClientId:     "1234",
	// 	ClientSecret: "aabbccdd",
	//	AuthorizeUrl: "http://localhost:9393/authorize",
	//	TokenUrl:     "http://localhost:9393/token",
	//	RedirectUrl:  "http://localhost:9393/oauth2callback"
	// Now the server

	// OAuth2 http handler's
	r.HandleFunc("/authorize", Server.AuthorizeHandler)
	r.HandleFunc("/token", Server.TokenHandler)
	r.HandleFunc("/info", Server.InfoHandler)

	// Client redirect address.
	r.HandleFunc("/oauth2callback", Server.AppAuthCodeHandler)

	// OpenId service.
	r.HandleFunc("/.well-known/openid-configuration", Server.DiscoveryHandler)
	r.HandleFunc("/publickeys", Server.PublicKeysHandler)

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

	err := http.ListenAndServe(":"+strconv.Itoa(port), r)

	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}

}
