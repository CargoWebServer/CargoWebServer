/**
 * Contain the documentation of the framework
 */
var DocumentationPage = function (parent) {
    // Text of the page.
    var languageInfo = {
        // english
        "en": {
            "documentation-intro-lnk": "Introduction",
            "documentation-introduction-title": "Introduction",
            "documentation-programming-background": "Programming background",
            "documentation-introduction-p1": "In order to take advantage of a framework it's important to know what it philosophy is. There a lot of choice out there and "
            + "find the good tool that meet your needs can be confusing. Chosing a framework simply because it's the flavor of the "
            + "day whiout knowing to much about it is a risky buisness. Extensibility, simplicity, efficiency, flexibility are overused "
            + "word's these day's. So let's put all of this in perspective and find out if Cargo is the good tool for you.",
            "documentation-introduction-p2": "First of all simplicity. My background is OOP the first programming language I learn was C++, and later "
            + "Java. Far from being perfect, those languages have the advantage to impose structure at there core. There goal is to "
            + "protect the integrity of the system. In a language like JavaScript structures are more a matter of desing, but trust "
            + "me without structures, simplicity give place to complexity. OOP it's a good starting point for using Cargo, but it's "
            + "not the destination...",
            "documentation-introduction-p3": "If is easy to switch from C++ to Java, thing get more complicated when you switch to JavaScript. Because JavaScript is "
            + "not a pure OOP, thinking purely in OOP result in ugly code. I like to thing OOP language as lego block and functional "
            + "language as playdough. My goal here is not to teach you JavaScript, but to encourage you to take advantage of the flexibility "
            + "of this language without give up all good habit learn from OOP.",
            "documentation-introduction-p4": "If you came from C++, Java, Objective-C or any other desktop application programming language you will enjoy working with Cargo.",
            "philosophy-intro-lnk": "Philosophy",
            "documentation-philosophy-title": "Philosophy",
            "documentation-philosophy-desc": "In tow words <span>non-intrusive</span>. Cargo dosen't impose you any rules. The role of Cargo is give you access to "
            + "server functionalies via the server object. Simplicity here is achieve by making abstraction of the network. We also "
            + "create some helper class to simplify the use of HTML and SVG element. In order to keep thing simple, we keep external "
            + "dependencies at mimimum. You can use tool like jQuery, Bootstrap or Angular but those libraries are not part of the "
            + "framework. We will see later how to achieve extensibity whit Cargo.",
            "documentation-architecture-lnk": "Architecture",
            "documentation-architecture-title": "Architecture",
            "documentation-from-desk-to-web": "From desktop application to web application",
            "documentation-architecture-p1": "Application programming in Java, C++ or objective-C fowllow the same principles. Framework in these languages offer basicaly "
            + "the same kind of functionalities like database access, interface design tools, openGL api... They are composed of various libraries "
            + "that exetend functionalities of the language and facilitate the development of application. Swing in Java, Cacao in Objective-c,"
            + "MFC and Qt in C++ are good example of desktop application framework. If we think HTML5 as a JavaScript framework we can see some similarities "
            + "with these frameworks. Each new functionnality added in HTML5 address a need of application programming. With all these "
            + "new functionalities we can now considere JavaScript/HTML5 as a desktop application framework.",
            "documentation-architecture-p2": "There's no formal definition of what is a <span>web</span> application. For the sake of clarity, let's define web application as "
            + "distibuted application. The major difference between a desktop application and a web application is the model of execution. Desktop "
            + "application run on single computer, and distributed application on a network. In those applications, functionalities are access with use "
            + "of libraries. The communication with the framework and the application are made via object instances or function call. Instead of libraries, "
            + "distributed application made use of services, and call to function are made via RPC or distributed object reference.",
            "documentation-architecture-p3": "Services offer by Cargo are, objects persitence, sessions management, events management (channels, listener), ressource management (files, user, group, role, action...), "
            + "data access (SQL, Key/Value), authentication (LDAP). Because Cargo take care of all the burden of distributed aspects of your "
            + "application, developping web application with it isn't more difficult than developping desktop application.",
            "documentation-big-picture": "The big picture",
            "documentation-big-picture-lnk": "The big picture",
            "documentation-big-picture-p1": "Cargo simplify web application programming by exposing distant services as local objects. Just like you use <span>document</span> object to intereact with "
            + "HTML elements, you use <span>server</span> object to interact with the backend. The only difference is the communication model. Local object methods call are "
            + "synchronous and distant object method call are asynchronous. To work around it and simplify the code, we made use of callback's functions. The "
            + "three callbacks functions are, success, error and progress. They represent the state of the action call on the backend. In backround, when a method on "
            + "server object is invoke, a request to the backend is send. The request is process by the backend and a response message is sent back to client. Depending of "
            + "the response type, the appropriate callback function is call.",
            "documentation-big-picture-p2": "Distributed system based from OOP made use of interface definition language IDL to describe component's application programming interface. With Cargo, in place of "
            + "IDL, we made use of JavaScript. Because Cargo backend can interpret JavaScript and because all of it's functionalities are interface with JavaScript, you can ask Cargo to do "
            + "all you want. Following the the open closed principle, you can exetend the server functionalities from the client side without modifying it's code.",
            "documentation-event-handling-title": "Event handling",
            "documentation-event-handling-lnk": "Event handling",
            "documentation-event-handling-subtitle": "don't call us, we'll call you",
            "documentation-event-handling-p": "Cargo event handling are pretty simple and pretty usefull. The design of the event system follow the publish-subscribe pattern principles. Event can be see as piece "
            + "of inforamtion travelling through channels. Each channel can carry one type of event and each event can travel throw one channel. Channel are dynamic, when a listener "
            + "register itself on an event channel, the channel is created, and inversely the channel is close when no more listeners use it. One of the key concept in Cargo event "
            + "handling is the <span>EventHub</span>. From the point of view of server, who play the role of publisher, the <span>EventHub</span> play the role of subscriber. But from the point "
            + "of view of local object, who play the role of subscriber, it play the role of publisher. Just like event channels and events, <span>EventHub</span> are specialize in on type of event. "
            + "Class derived from <span>EventHub</span> play the role of event controller in your application. Here we made use of a variation of the observer pattern. If an object want's "
            + "to be inform of a given event, it define an update function for that particular event and attach itself to the <span>EventHub</span> reponsable of that event."
            + "Latter when the object has no more interest to receive event message, he simply detach itself from the <span>EventHub</span>. Because there can be a lot of events travelling the "
            + "network, we create event filters. Event filters are <span>regex</span> regular expression. By default channel are silent, in order to receive event you must register a regular expression (filter) "
            + "for a given event type. To get all event of a given type the regex * can be use. By using the event model in your application you decoupling classes and making your "
            + "code more robust and manageable.",
            "documentation-entity-title": "Entity",
            "documentation-entity-lnk": "Entity",
            "documentation-entity-subtitle": "something that exists as itself",
            "documentation-entity-p1": "JavaScript Object Notation JSON is the facto information model those days. Shinning by it's simplicity and it's versatility every structures can be serialysed with it. "
            + "But when you serialyse an instance of a class and instantiate it back you end up with a json object and not the initial instance. That represent a problem when it's time "
            + "to make object persistent. Data store need meta data in order to be able to manipulated his contained data. The concept of schema is use to define data structure inside the "
            + " data sotre and the relations between those structures. Query language also need meta data to be able work properly. To fullfill our need we create the entity prototype concept."
            + "The entity prototype contain the additionals informations needed by application to persist class instance. The entity prototype was inspire by XML schema <span>XSD</span>, "
            + "he's fully compatible with it, and also with SQL schema.",
            "documentation-entity-p2": "Being the model part of your application, the entities are instances of JavaScript classes. By simply creating a data store connection in your configuration, or importing XSD schema, "
            + "Cargo automatically do all the work for you. The mapping between the data source and the classes instances is done transparently. From the entity prototype, Cargo generate the "
            + "classes definitions and use it to instantiate entities. Class definition are injected on the JavaScript runtime when you need it, and became accessible to your application. Because "
            + "entity query on it. Cargo has it own query language named entity query language EQL. Not only you can get access to exsiting data but you can define new one directly in JavaScript,"
            + "simple by creating a new entity prototype instance. ",
            "documentation-services-title":"Services",
            "documentation-services-lnk":"Services",
            "documentation-services-p":"The <span>sever</span> object contain few functionalities by itself, it role is to give access to other services objects. IF CSS3/HTML5/JavaScript, in the context of the browser,"
            +"are suitable to create desktop application, they lack most of the required features needed by distributed applications. Remember, with Cargo the application code run on the client side. "
            +"This way to do have the advantage of minimizing ressource usage on the server side. Aslo the connection model of Cargo is different of most web application framework, instead of using "
            +"http, it make use of it own protocol runing over the websocket. That connection model is best suited for service oriented architecture. In that model, the client and the server side "
            +"are well definied and independant of each other. Most of Cargo services are in fact interface to other external services like SQL DBMS, LDAP or SMTP server. Cargo give acces to those "
            +"underlying services in a uniform manner. Because services configurations informations are keep in the server side their access are protected. I will tell more about each services "
            +"in the tutorials.",
            "documentation-security-title":"Security",
            "documentation-security-subtitle":"With Great Power Comes Great Responsibility",
            "documentation-security-lnk":"Security",
            "documentation-security-p":"Internally Cargo work with Entities. As exemple, File, Group, User, Role, Session and Action are entities. Draft...",
        },
        //francais
        "fr": {
            // TODO translate it...
            "documentation-intro-lnk": "Introduction",
            "documentation-introduction-title": "Introduction",
            "documentation-programming-background": "Programming background",
            "documentation-introduction-p1": "In order to take advantage of a framework it's important to know what it philosophy is. There a lot of choice out there and "
            + "find the good tool that meet your needs can be confusing. Chosing a framework simply because it's the flavor of the "
            + "day whiout knowing to much about it is a risky buisness. Extensibility, simplicity, efficiency, flexibility are overused "
            + "word's these day's. So let's put all of this in perspective and find out if Cargo is the good tool for you.",
            "documentation-introduction-p2": "First of all simplicity. My background is OOP the first programming language I learn was C++, and later"
            + "Java. Far from being perfect, those languages have the advantage to impose structure at there core. There goal is to "
            + "protect the integrity of the system. In a language like JavaScript structures are more a matter of desing, but trust "
            + "me without structures, simplicity give place to complexity. OOP it's a good starting point for using Cargo, but it's"
            + "not the destination...",
            "documentation-introduction-p3": "If is easy to switch from C++ to Java, thing get more complicated when you switch to JavaScript. Because JavaScript is"
            + "not a pure OOP, thinking purely in OOP result in ugly code. I like to thing OOP language as lego block and functional "
            + "language as playdough. My goal here is not to teach you JavaScript, but to encourage you to take advantage of the flexibility "
            + "of this language without give up all good habit learn from OOP.",
            "documentation-introduction-p4": "If you came from C++, Java, Objective-C or any other desktop application programming language you will enjoy working with Cargo.",
            "philosophy-intro-lnk": "Philosophy",
            "documentation-philosophy-title": "Philosophy",
            "documentation-philosophy-desc": "In tow words <span>non-intrusive</span>. Cargo dosen't impose you any rules. The role of Cargo is give you access to"
            + "server functionalies via the server object. Simplicity here is achieve by making abstraction of the network. We also "
            + "create some helper class to simplify the use of HTML and SVG element. In order to keep thing simple, we keep external "
            + "dependencies at mimimum. You can use tool like jQuery, Bootstrap or Angular but those libraries are not part of the "
            + "framework. We will see later how to achieve extensibity whit Cargo.",
            "documentation-architecture-lnk": "Architecture",
            "documentation-architecture-title": "Architecture",
            "documentation-from-desk-to-web": "From desktop application to web application",
            "documentation-architecture-p1": "Application programming in Java, C++ or objective-C fowllow the same principles. Framework in these languages offer basicaly"
            + "the same kind of functionalities like database access, interface design tools, openGL api... They are composed of various libraries"
            + "that exetend functionalities of the language and facilitate the development of application. Swing in Java, Cacao in Objective-c,"
            + "MFC and Qt in C++ are good example of desktop application framework. If we think HTML5 as a JavaScript framework we can see some similarities "
            + "with these frameworks. Each new functionnality added in HTML5 address a need of application programming. With all these "
            + "new functionalities we can now considere JavaScript/HTML5 as a desktop application framework.",
            "documentation-architecture-p2": "There's no formal definition of what is a <span>web</span> application. For the sake of clarity, let's define web application as "
            + "distibuted application. The major difference between a desktop application and a web application is the model of execution. Desktop "
            + "application run on single computer, and distributed application on a network. In those applications, functionalities are access with use "
            + "of libraries. The communication with the framework and the application are made via object instances or function call. Instead of libraries, "
            + "distributed application made use of services, and call to function are made via RPC or distributed object reference.",
            "documentation-architecture-p3": "Services offer by Cargo are, objects persitence, sessions management, events management (channels, listener), ressource management (files, user, group, role, action...),"
            + "data access (SQL, Key/Value), authentication (LDAP). Because Cargo take care of all the burden of distributed aspects of your "
            + "application, developping web application with it isn't more difficult than developping desktop application.",
            "documentation-big-picture": "The big picture",
            "documentation-big-picture-lnk": "The big picture",
            "documentation-big-picture-p1": "Cargo simplify web application programming by exposing distant services as local objects. Just like you use <span>document</span> object to intereact with "
            + "HTML elements, you use <span>server</span> object to interact with the backend. The only difference is the communication model. Local object methods call are "
            + "synchronous and distant object method call are asynchronous. To work around it and simplify the code, we made use of callback's functions. The "
            + "three callbacks functions are, success, error and progress. They represent the state of the action call on the backend. In backround, when a method on "
            + "server object is invoke, a request to the backend is send. The request is process by the backend and a response message is sent back to client. Depending of "
            + "the response type, the appropriate callback function is call.",
            "documentation-big-picture-p2": "Distributed system based from OOP made use of interface definition language IDL to describe component's application programming interface. With Cargo, in place of "
            + "IDL, we made use of JavaScript. Because Cargo backend can interpret JavaScript and because all of it's functionalities are interface with JavaScript, you can ask Cargo to do "
            + "all you want. Following the the open closed principle, you can exetend the server functionalities from the client side without modifying it's code.",
            "documentation-event-handling-title": "Event handling",
            "documentation-event-handling-lnk": "Event handling",
            "documentation-event-handling-subtitle": "don't call us, we'll call you",
            "documentation-event-handling-p": "Cargo event handling are pretty simple and pretty usefull. The design of the event system follow the publish-subscribe pattern principles. Event can be see as piece "
            + "of inforamtion travelling through channels. Each channel can carry one type of event and each event can travel throw one channel. Channel are dynamic, when a listener "
            + "register itself on an event channel, the channel is created, and inversely the channel is close when no more listeners use it. One of the key concept in Cargo event "
            + "handling is the <span>EventHub</span>. From the point of view of server, who play the role of publisher, the <span>EventHub</span> play the role of subscriber. But from the point "
            + "of view of local object, who play the role of subscriber, it play the role of publisher. Just like event channels and events, <span>EventHub</span> are specialize in on type of event. "
            + "Class derived from <span>EventHub</span> play the role of event controller in your application. Here we made use of a variation of the observer pattern. If an object want's "
            + "to be inform of a given event, it define an update function for that particular event and attach itself to the <span>EventHub</span> reponsable of that event."
            + "Latter when the object has no more interest to receive event message, he simply detach itself from the <span>EventHub</span>. Because there can be a lot of events travelling the "
            + "network, we create event filters. Event filters are <span>regex</span> regular expression. By default channel are silent, in order to receive event you must register a regular expression (filter) "
            + "for a given event type. To get all event of a given type the regex * can be use. By using the event model in your application you decoupling classes and making your "
            + "code more robust and manageable.",
            "documentation-entity-title": "Entity",
            "documentation-entity-lnk": "Entity",
            "documentation-entity-subtitle": "something that exists as itself",
            "documentation-entity-p1": "JavaScript Object Notation JSON is the facto information model those days. Shinning by it's simplicity and it's versatility every structures can be serialysed with it."
            + "But when you serialyse an instance of a class and instantiate it back you end up with a json object and not the initial instance. That represent a problem when it's time "
            + "to make object persistent. Data store need meta data in order to be able to manipulated his contained data. The concept of schema is use to define data structure inside the "
            + " data sotre and the relations between those structures. Query language also need meta data to be able work properly. To fullfill our need we create the entity prototype concept."
            + "The entity prototype contain the additionals informations needed by application to persist class instance. The entity prototype was inspire by XML schema <span>XSD</span>, "
            + "he's fully compatible with it, and also with SQL schema.",
            "documentation-entity-p2": "Being the model part of your application, the entities are instances of JavaScript classes. By simply creating a data store connection in your configuration, or importing XSD schema, "
            + "Cargo automatically do all the work for you. The mapping between the data source and the classes instances is done transparently. From the entity prototype, Cargo generate the "
            + "classes definitions and use it to instantiate entities. Class definition are injected on the JavaScript runtime when you need it, and became accessible to your application. Because "
            + "entity query on it. Cargo has it own query language named entity query language EQL. Not only you can get access to exsiting data but you can define new one directly in JavaScript,"
            + "simple by creating a new entity prototype instance. ",
            "documentation-services-title":"Services",
            "documentation-services-lnk":"Services",
            "documentation-services-p":"The <span>sever</span> object contain few functionalities by itself, it role is to give access to other services objects. IF CSS3/HTML5/JavaScript, in the context of the browser,"
            +"are suitable to create desktop application, they lack most of the required features needed by distributed applications. Remember, with Cargo the application code run on the client side. "
            +"This way to do have the advantage of minimizing ressource usage on the server side. Aslo the connection model of Cargo is different of most web application framework, instead of using "
            +"http, it make use of it own protocol runing over the websocket. That connection model is best suited for service oriented architecture. In that model, the client and the server side "
            +"are well definied and independant of each other. Most of Cargo services are in fact interface to other external services like SQL DBMS, LDAP or SMTP server. Cargo give acces to those "
            +"underlying services in a uniform manner. Because services configurations informations are keep in the server side their access are protected. I will tell more about each services "
            +"in the tutorials.",
            "documentation-security-title":"Security",
            "documentation-security-subtitle":"With Great Power Comes Great Responsibility",
            "documentation-security-lnk":"Security",
            "documentation-security-p1":"Internally Cargo work with Entities. As exemple, File, Group, User, Role, Session and Action are entities. Draft...",
        }
    }

    // Depending of the language the correct text will be set.
    server.languageManager.appendLanguageInfo(languageInfo)

    // The getting started panel.
    this.panel = new Element(parent, { "tag": "div", "class": "row hidden pageRow", "id": "page-cargoDocumentation" })
        .appendElement({ "tag": "div", "class": "col-lg-12" }).down()

    this.links = this.panel.appendElement({ "tag": "div", "class": "col-sm-2", "role": "complementary" }).down()
        .appendElement({ "tag": "nav", "class": "bs-docs-sidebar affix-top" }).down()
        .appendElement({ "tag": "ul", "class": "nav bs-docs-sidenav" }).down()

    // Links
    this.links.appendElement({ "tag": "li" }).down()
        .appendElement({ "tag": "a", "href": "#documentationIntroductionDiv", "id": "documentation-intro-lnk" }).up()
        .appendElement({ "tag": "li" }).down()
        .appendElement({ "tag": "a", "href": "#documentation-philosophy-title", "id": "philosophy-intro-lnk" }).up()
        .appendElement({ "tag": "li" }).down()
        .appendElement({ "tag": "a", "href": "#documentationArchitectureDiv", "id": "documentation-architecture-lnk" }).up()
        .appendElement({ "tag": "li" }).down()
        .appendElement({ "tag": "a", "href": "#documentation-big-picture", "id": "documentation-big-picture-lnk" }).up()
        .appendElement({ "tag": "li" }).down()
        .appendElement({ "tag": "a", "href": "#documentation-event-handling-title", "id": "documentation-event-handling-lnk" }).up()
        .appendElement({ "tag": "li" }).down()
        .appendElement({ "tag": "a", "href": "#documentation-entity-title", "id": "documentation-entity-lnk" }).up()
        .appendElement({ "tag": "li" }).down()
        .appendElement({ "tag": "a", "href": "#documentation-services-title", "id": "documentation-services-lnk" }).up()
        .appendElement({ "tag": "li" }).down()
        .appendElement({ "tag": "a", "href": "#documentation-security-title", "id": "documentation-security-lnk" }).up()

    // Now the introduction panel.
    this.main = this.panel.appendElement({ "tag": "div", "class": "col-sm-10", "role": "main" }).down()
        .appendElement({ "tag": "div", "class": "row" }).down()
        .appendElement({ "tag": "div", "class": "col-xs-12" }).down()

    var intro = this.main.appendElement({ "tag": "div", "class": "col-xs-12", "id": "documentationIntroductionDiv" })
        // Introduction
        .appendElement({ "tag": "h2", "id": "documentation-introduction-title" })
        .appendElement({ "tag": "h4", "id": "documentation-programming-background" })
        .appendElement({ "tag": "p", "id": "documentation-introduction-p1" })
        .appendElement({ "tag": "p", "id": "documentation-introduction-p2" })
        .appendElement({ "tag": "p", "id": "documentation-introduction-p3" })
        .appendElement({ "tag": "p", "id": "documentation-introduction-p4" })
        // Philosophy
        .appendElement({ "tag": "h4", "id": "documentation-philosophy-title" })
        .appendElement({ "tag": "p", "id": "documentation-philosophy-desc" })

    // Architecture
    var architecture = this.main.appendElement({ "tag": "div", "class": "col-xs-12", "id": "documentationArchitectureDiv" })
        // Architecture intro
        .appendElement({ "tag": "h2", "id": "documentation-architecture-title" })
        .appendElement({ "tag": "h4", "id": "documentation-from-desk-to-web" })
        .appendElement({ "tag": "p", "id": "documentation-architecture-p1" })
        .appendElement({ "tag": "p", "id": "documentation-architecture-p2" })
        .appendElement({ "tag": "p", "id": "documentation-architecture-p3" })
        // The big picture
        .appendElement({ "tag": "h4", "id": "documentation-big-picture" })
        .appendElement({ "tag": "p", "id": "documentation-big-picture-p1" })
        .appendElement({ "tag": "p", "id": "documentation-big-picture-p2" })
        // Event handling
        .appendElement({ "tag": "h4", "id": "documentation-event-handling-title" })
        .appendElement({ "tag": "h5", "id": "documentation-event-handling-subtitle" })
        .appendElement({ "tag": "p", "id": "documentation-event-handling-p" })
        // Entity
        .appendElement({ "tag": "h4", "id": "documentation-entity-title" })
        .appendElement({ "tag": "h5", "id": "documentation-entity-subtitle" })
        .appendElement({ "tag": "p", "id": "documentation-entity-p1" })
        .appendElement({ "tag": "p", "id": "documentation-entity-p2" })
        // Services
        .appendElement({ "tag": "h4", "id": "documentation-services-title" })
        .appendElement({ "tag": "p", "id": "documentation-services-p" })
        // Security
        .appendElement({ "tag": "h4", "id": "documentation-security-title" })
        .appendElement({ "tag": "h5", "id": "documentation-security-subtitle" })
        .appendElement({ "tag": "p", "id": "documentation-security-p" })
}