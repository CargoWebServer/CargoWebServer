#Welcome to Cargo 
Cargo is a complete web application framework. It's fast, easy and lightweight. 

Cargo was created with a Service Oriented Achitecture (SOA) in mind. The basic services offered by the framework are;

- Entities service: an easy way to make json object persistent (CRUD)
- Data service: give access to SQL or Key/value store from the browser.
- Session service: login functionalities
- Account service: users & groups management
- Event service: channel and events functionalities
- Security service: management by roles and permissions of actions and access respectively.
- Other service: LDAP, SMTP.

Extensibility and modularity are key concepts in the design of Cargo. You can create your own service, all you have to do is implement the Serivce interface in Go. You can also use the service container and create a plugin in C++.

Cargo uses the websocket/tcp scoket to communicate with client's. To do so it has it own protocol written with google protobuffer, and similiar to [JSON/RPC](https://github.com/CargoWebServer/CargoWebServer/blob/master/WebApp/Cargo/Apps/Cargo/proto/rpc.proto). The Server Object Model (SOM) was created to simplify the interaction with the sever. Similar to the document object model (DOM) which gives access to the browser, the server object model SOM gives you access to a server. All you have to do is create the server object and invoke actions on it. With the help of callbacks and events, communication with the SOM is easy and intuive.

www.cargowebserver.com
