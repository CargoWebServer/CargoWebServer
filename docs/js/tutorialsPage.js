
var TutorialsPage = function (parent) {
    // Text of the page.
    var languageInfo = {
        // english
        "en": {
            "tutorial-title": "Tutorials",
            "tutorial-presentation": "'Any fool can know. The point is to understand.'</br> <span style='font-size: 10pt;'>Albert Einstein</span>",
            "tutorial-presentation-description": "Learning a new tool can be challenging not to say frustrating. I believe true knowledge came from experience, "
            + " tutorials are there to guide you through the process of learning. They put you in context of web application programming and give you the experience you "
            + "need to take advantage of Cargo. Each tutorial are workable application. Making something usefull from start?... That's rewording!!!",
            "introduction-tutorial-lnk": "Introduction",
            "introduction-tutorial-title": "Introduction",
            "introduction-tutorial-content": "First of all, welcome on board! Cargo is a work in progress, we try to make your experience of it the best we can. "
            + "Programming and documenting a framework is a huge task. This website was created and power by Cargo, each day, you will see more and more usefull informations on it. "
            + "So don't Hesitate to communicate with us to share comment or idea at <a href=\"mailto:cargowebserver@gmail.com\">CargoWebServer@gmail.com</a>, your opinion is important to us.",
            "convention-tutorial-title": "Conventions",
            "convention-tutorial-content": "Let's me now explain to you the tutorials conventions. Because of the name Cargo, we made use of nautic symbols.",
            "tutorial-lifesaver-desc": "The life saver symbol is use when we want to teach you good practices or usefull techniques.",
            "tutorial-buoy-desc": "The buoy symbol is use to warning you about common mistake or bad programming practices.",
            "tutorial-anchor-desc": "The anchor symbol are use for complementary explanation or discution that can be skip by hurry reader's.",
            "SOM-tutorial-title": "The Service Object Model (SOM)",
            "SOM-tutorial-lnk": "SOM tutorial",
            "SOM-tutorial-content": "Making a distant service appear like local object, that's the essence of the SOM. Object are intuitive to work with, they regroup related actions and informations."
            + "Let's now create a simple application to get a better view of how to work with SOM."
        },
        //francais
        "fr": {
            "tutorial-title": "Cargo Tutorial",
            "tutorial-presentation": "Cargo Dactilogiciel",
            "tutorial-presentation-description": "",
            "SOM-tutorial-title": "",
            "SOM-tutorial-content": "",
            "SOM-tutorial-lnk": "",
        }
    }

    // Depending of the language the correct text will be set.
    server.languageManager.appendLanguageInfo(languageInfo)

    // The getting started panel.
    this.panel = new Element(parent, { "tag": "div", "class": "row hidden pageRow", "id": "page-cargoTutorial" })
        .appendElement({ "tag": "div", "class": "col-lg-12" }).down()

    this.links = this.panel.appendElement({ "tag": "div", "class": "col-sm-2", "role": "complementary" }).down()
        .appendElement({ "tag": "nav", "class": "bs-docs-sidebar affix-top" }).down()
        .appendElement({ "tag": "ul", "class": "nav bs-docs-sidenav" }).down()

    // Links
    this.links
        // Tutorial introduction.
        .appendElement({ "tag": "li" }).down()
        .appendElement({ "tag": "a", "href": "#introductionTutorialDiv", "id": "introduction-tutorial-lnk" }).up()
        // The blog tutorial presentation.
        .appendElement({ "tag": "li" }).down()
        .appendElement({ "tag": "a", "href": "#somTutorialDiv", "id": "SOM-tutorial-lnk" }).up()


    // Now the introduction panel.
    this.main = this.panel.appendElement({ "tag": "div", "class": "col-sm-10", "role": "main" }).down()
        .appendElement({ "tag": "div", "class": "row" }).down()
        .appendElement({ "tag": "div", "class": "col-xs-12" }).down()

    // Introduction text.
    this.intro = this.main
        .appendElement({ "tag": "div", "class": "jumbotron" }).down()
        .appendElement({ "tag": "div", "class": "container" }).down()
        .appendElement({ "tag": "div", "class": "col-xs-2" }).down()
        .appendElement({ "tag": "img", "src": "img/Lighthouse.svg", "style": "max-width: 100px;" }).up()
        .appendElement({ "tag": "div", "class": "col-xs-10" }).down()
        .appendElement({ "tag": "h2", "id": "tutorial-title" })
        .appendElement({ "tag": "p" }).down()
        .appendElement({ "tag": "span", "id": "tutorial-presentation" }).up()

        // Description of the tutorials.
        .appendElement({ "tag": "p", "id": "tutorial-presentation-description" })

    // The  tutorial.
    this.main
        .appendElement({ "tag": "div", "class": "row", "id": "introductionTutorialDiv" }).down()
        .appendElement({ "tag": "div", "class": "col-xs-12" }).down()
        .appendElement({ "tag": "h2", "id": "introduction-tutorial-title" })
        .appendElement({ "tag": "p", "id": "introduction-tutorial-content" })
        .appendElement({ "tag": "h4", "id": "convention-tutorial-title" })
        .appendElement({ "tag": "p", "id": "convention-tutorial-content" })
        .appendElement({ "tag": "ul", "style": "list-style: none;" }).down()
        // Life saver symbol
        .appendElement({ "tag": "li" }).down()
        .appendElement({ "tag": "div", "class": "tutorial-symbol-div" }).down()
        .appendElement({ "tag": "img", "src": "img/lifesaver.svg", "style": "width: 35px; height:35px;" })
        .appendElement({ "tag": "span", "id": "tutorial-lifesaver-desc" }).up().up()
        // buoy
        .appendElement({ "tag": "li" }).down()
        .appendElement({ "tag": "div", "class": "tutorial-symbol-div" }).down()
        .appendElement({ "tag": "img", "src": "img/buoy.svg", "style": "width: 35px; height:35px;" })
        .appendElement({ "tag": "span", "id": "tutorial-buoy-desc" }).up().up()
        // buoy
        .appendElement({ "tag": "li" }).down()
        .appendElement({ "tag": "div", "class": "tutorial-symbol-div" }).down()
        .appendElement({ "tag": "img", "src": "img/anchor.svg", "style": "width: 35px; height:35px;" })
        .appendElement({ "tag": "span", "id": "tutorial-anchor-desc" }).up().up()

    // The blog tutorial.
    var somTutorialDiv = this.main.appendElement({ "tag": "div", "class": "row", "id": "SOM_TutorialDiv" }).down()
    somTutorialDiv.appendElement({ "tag": "div", "class": "col-xs-12" }).down()
        .appendElement({ "tag": "h2", "id": "SOM-tutorial-title" })
        .appendElement({ "tag": "p", "id": "SOM-tutorial-content" })
        // The tutorial content will be created here in english, translation will come latter.
        .appendElement({ "tag": "h2", "innerHtml": "Create a new project" })
        .appendElement({
            "tag": "p", "innerHtml": "In order to create a new project in Cargo simply create a new folder inside "
            + " CargoWebServer/WebApp/Cargo/Apps folder and put file named index.hml inside it. The index.html file is the "
            + "entry point of your web application, it contain the title of your application and all needed ressources file path."
            + "You can create html elements in index.htlm, but personnaly, I prefer to create elements dynamicaly,"
            + "I will show you how to do so latter.</br>"
            + "In order to keep the content of the application folder clean, I create folders, one "
            + "for each kind of ressource use by the application, css, js, svg, img... "
        }).appendElement({ "tag": "div", "style": "width: 100%; text-align: center; padding: 10px;" }).down()
        .appendElement({ "tag": "img", "src": "img/screenShot1.png", "style": "" }).up()
        .appendElement({
            "tag": "p", "innerHtml": "Here is the index.html page for our exemple let me explain to you in more detail what's in it."
            + "Index.html it's a standard html page, it is compose of a 'head' and a 'body' section. The diffence is not the file itself but how we use it."
            + "Because Index.html is the entry point of your application, it must contain all necessary information needed by your application. If we take a "
            + "closer look we can see that most of index.html code is compose of 'script' and 'link' element's."
        }).down().appendElement({ "tag": "div", "class": "tutorial-symbol-div" }).down()
        .appendElement({ "tag": "img", "src": "img/buoy.svg", "style": "width: 35px; height:35px;" })
        .appendElement({
            "tag": "span", "innerHtml": "Ressource on the server are store at the CARGOROOT path, so /Cargo/js/utility.js mean's CARGOROOT/Cargo/js/utility.js. "
            + "Without the '/' at the begining of the path it became relative to the index.html file itself."
        })


    //////////////////////////// Html index //////////////////////////
    this.indexHtmlCode = somTutorialDiv.appendElement({ "tag": "pre" }).down()
        .appendElement({ "tag": "code", "class": "xml hljs" }).down()

    // Append the meta element   
    appendHtmlMeta(this.indexHtmlCode, "&lt;!DOCTYPE html&gt;")

    // Open brace...
    appendHtmlTag(this.indexHtmlCode, 0, "html")
    appendHtmlTag(this.indexHtmlCode, 1, "head", [{ "name": "lang", "value": "en" }])

    appendHtmlTag(this.indexHtmlCode, 2, "meta", [{ "name": "http-equiv", "value": "x-ua-compatible" }, { "name": "content", "value": "IE=edge" }])
    appendHtmlTag(this.indexHtmlCode, 2, "meta", [{ "name": "charset", "value": "UTF-8" }])
    appendHtmlTag(this.indexHtmlCode, 2, "title", [], true, "Tutorial: SOM (Server Object Model)")

    // css
    appendHtmlComment(this.indexHtmlCode, "css", 2)

    // Closing...
    appendHtmlTag(this.indexHtmlCode, 2, "link", [{ "name": "rel", "value": "stylesheet" }, { "name": "type", "value": "text/css" }, { "name": "href", "value": "/Cargo/css/main.css" }])


    appendHtmlComment(this.indexHtmlCode, "External dependencies files", 2)
    appendHtmlTag(this.indexHtmlCode, 2, "script", [{ "name": "src", "value": "/lib/protobuf/bytebuffer.js" }], true, "")
    appendHtmlTag(this.indexHtmlCode, 2, "script", [{ "name": "src", "value": "/lib/protobuf/long.js" }], true, "")
    appendHtmlTag(this.indexHtmlCode, 2, "script", [{ "name": "src", "value": "/lib/protobuf/protobuf.js" }], true, "")
    appendHtmlTag(this.indexHtmlCode, 2, "script", [{ "name": "src", "value": "/lib/rollups/sha256.js" }], true, "")
    appendHtmlTag(this.indexHtmlCode, 2, "script", [{ "name": "src", "value": "/lib/stringview.js" }], true, "")

    appendHtmlComment(this.indexHtmlCode, "Cargo files", 2)
    appendHtmlTag(this.indexHtmlCode, 2, "script", [{ "name": "src", "value": "/Cargo/js/utility.js" }], true, "")
    appendHtmlTag(this.indexHtmlCode, 2, "script", [{ "name": "src", "value": "/Cargo/js/handlers.js" }], true, "")
    appendHtmlTag(this.indexHtmlCode, 2, "script", [{ "name": "src", "value": "/Cargo/js/error.js" }], true, "")
    appendHtmlTag(this.indexHtmlCode, 2, "script", [{ "name": "src", "value": "/Cargo/js/event.js" }], true, "")
    appendHtmlTag(this.indexHtmlCode, 2, "script", [{ "name": "src", "value": "/Cargo/js/entity.js" }], true, "")
    appendHtmlTag(this.indexHtmlCode, 2, "script", [{ "name": "src", "value": "/Cargo/js/languageManager.js" }], true, "")
    appendHtmlTag(this.indexHtmlCode, 2, "script", [{ "name": "src", "value": "/Cargo/js/server.js" }], true, "")
    appendHtmlTag(this.indexHtmlCode, 2, "script", [{ "name": "src", "value": "/Cargo/js/main.js" }], true, "")
    appendHtmlTag(this.indexHtmlCode, 2, "script", [{ "name": "src", "value": "/Cargo/js/element.js" }], true, "")

    appendHtmlComment(this.indexHtmlCode, "Application CSS files", 2)
    appendHtmlTag(this.indexHtmlCode, 2, "link", [{ "name": "rel", "value": "stylesheet" }, { "name": "type", "value": "text/css" }, { "name": "href", "value": "css/main.css" }])
    appendHtmlTag(this.indexHtmlCode, 1, "head", [], false)

    appendHtmlTag(this.indexHtmlCode, 1, "body")
    appendHtmlComment(this.indexHtmlCode, "Application JS files ", 2)
    appendHtmlTag(this.indexHtmlCode, 2, "script", [{ "name": "src", "value": "js/main.js" }], true, "")
    appendHtmlTag(this.indexHtmlCode, 1, "body", [], false)

    appendHtmlTag(this.indexHtmlCode, 0, "html", [], false)
    //////////////////////////// end Html index //////////////////////////

    somTutorialDiv.appendElement({
        "tag": "p", "innerHtml": "In order to uniformize the initialisation sequence we define the main entry point. "
        + "When the user reach the index.html page, Cargo do a lot of stuff in bacground, and when he's ready to run your application,"
        + " he call the main function. So each of your application must define it own main function. I put that function in the main.js file, "
        + "but you can be put it anywhere else."
    })
        .appendElement({ "tag": "h2", "innerHtml": "The server object" })
        .appendElement({
            "tag": "p", "innerHtml": "Most of Web framework's made use of http as communication protocol, but with the avenement of Html5 and the webSocket thing's are about to change... "
            + "Like TCP socket, the webSocket use a connected communication model. When you open a new communication channel, "
            + "it connection stay active until one of tow side end's it. In my opinion, that way to operate are more suitable for Web application. It facilate connections management on both side."
            + "The webSocket has it own protocol, it's a low level protocol similar to TCP/IP. It goal are to establish connection between client and server in order to transfert paquets of information. "
            + "The signification of tranfered informations "
            + "are out of the scope of the webSocket. I made use of Google protocol buffer to define messages semantics. Here is the application protocol definition file "
            + "<a href=\"https://github.com/CargoWebServer/CargoWebServer/blob/master/WebApp/Cargo/Apps/Cargo/proto/rpc.proto\" title=\"Protocol definition\">rpc.proto</a>. "
        })
        .appendElement({
            "tag": "p", "innerHtml": "Each application has a 'server' object define " +
                "in file <a href=\"https://github.com/CargoWebServer/CargoWebServer/blob/master/WebApp/Cargo/Apps/Cargo/js/main.js\" title=\"Protocol definition\">main.js</a>, at line 27. "
                + "The default 'server' object is initialyse with the local address 127.0.0.1. It's correct to keep the localhost address at developement time, but don't forgot to change it "
                + "at the time of deployement. "
        })
}