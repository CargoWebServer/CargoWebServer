#Welcome to Cargo 
Cargo is a complete web application framework. It's fast, easy and lightweight. 

Cargo was created with Service Oriented Achitecture (SOA) in mind. The basics services offer by the framework are;

- Entities service: an easy way to make json object persistent (CRUD)
- Data service: give access to SQL or Key/value store from the browser.
- Session service: login functionalities
- Account service: user, group functionalities
- Event service: channel and events management functionalities
- Security service: management by roles and permissions of actions and access respectively.
- Other service: LDAP, SMTP.

Extensibility and modularity are key concepts in the design of Cargo. You can create your own service, all you have to do is implementing the Serive interface in Go. You can also use the service container and create a plugin in C++.

Cargo made use of the websocket/tcp scoket to communicate with client's. To do so it has it own protocol written with google protobuffer, and similiar to [JSON/RPC](https://github.com/CargoWebServer/CargoWebServer/blob/master/WebApp/Cargo/Apps/Cargo/proto/rpc.proto). The Server Object Model (SOM) was created to simplify the interaction with the sever. Like DOM give access to the browser, the SOM give you access to a server, all you have to do is create the server object and invoke action on it. With help of callback's and event's, the communication with the SOM are made very easy and intuive.

www.cargowebserver.com
