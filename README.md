#Welcome to Cargo 
Cargo is a complete web application framework. It's fast, easy and lightweight. 

Cargo was created with a Service Oriented Achitecture (SOA) in mind. The basic services offered by the framework are:

- Entities service: an easy way to make json object persistent (CRUD)
- Data service: gives access to SQL or Key/value store from the browser.
- Session service: login functionalities
- Account service: users & groups management
- Event service: channel and events functionalities
- Security service: management by roles and permissions of actions and access respectively.
- Other service: LDAP, SMTP.
- Made use of JerryScript (done) / Chakra (in progress) / V8 (futur work)runing in her own process, one per session and extended with Go funtions and Objects.

Extensibility and modularity are key concepts in the design of Cargo. You can create your own service, all you have to do is implement the Serivce interface in Go. You can also use the service container and create a plugin in C++.

Cargo uses the websocket/tcp scoket to communicate with clients. To do so it has it's own protocol written with google protobuffer, and similiar to [JSON/RPC](https://github.com/CargoWebServer/CargoWebServer/blob/master/WebApp/Cargo/Apps/Cargo/proto/rpc.proto). The Server Object Model (SOM) was created to simplify the interaction with the sever. Similar to the document object model (DOM) which gives access to the browser, the server object model SOM gives you access to a server. All you have to do is create the server object and invoke actions on it. With the help of callbacks and events, communication with the SOM is easy and intuive.

To compile the project on your computer you must have a Go environement properly configured.

###Compiling
The Go source code of Cargo are in the directories:
CargoWebServer(master depending if you clone the project or not)
  * Project (That must be part of your GOPATH)
  * Project/src (Go source code)
  * Project/src/code.myceliUs.com (Cargo source code)
  
  You must get the following dependencies:
- go get github.com/pborman/uuid
- go get github.com/alexbrainman/odbc
- go get github.com/denisenkom/go-mssqldb
- go get github.com/go-sql-driver/mysql
- go get github.com/golang/protobuf/proto
- go get github.com/kokardy/saxlike
- go get github.com/mavricknz/ldap
- go get github.com/nfnt/resize
- go get github.com/polds/imgbase64
- go get github.com/syndtr/goleveldb/leveldb
- go get github.com/xrash/smetrics
- go get github.com/bytbox/go-pop3
- go get github.com/skratchdot/open-golang/open
- go get gopkg.in/gomail.v1
- go get golang.org/x/net/websocket
- go get golang.org/x/text/runes
- go get github.com/RangelReale/osin
- go get golang.org/x/oauth2

###Build
To build Cago from the top level directory:

`cd CargoWebServer/Project/src/code.myceliUs.com/CargoWebServer`

You should see the file 'Main.go' it this directory, now call: 

`go build -i -tags "Config CargoEntities"`

The output file should be 'CargoWebServer' in linux and 'CargoWebServer.exe' in windows.

###Run
Now to run Cargo you must move CargoWebServer(.exe) at the top level directory, or create a new directory and put the following files inside:

- CargoWebServer/WebApp
- CargoWebServer/CargoWebServer(.exe)

Now you can start the command CargoWebServer(.exe) et voîlà!

type 127.0.0.1:9393 in your browser's address bar to go to the root of your newly installed server.

For more info go to our website.
www.cargowebserver.com
