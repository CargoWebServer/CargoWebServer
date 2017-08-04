/**
 * That page contain information about how to install a begin working with cargo.
 */
var GettingStartedPage = function (parent) {
    // Text of the page.
    var languageInfo = {
        // english
        "en": {
            "prerequisites-lnk": "Prerequisites",
            "prerequisites-title": "Prerequisites",
            "getting-cargo-lnk": "Getting Cargo",
            "getting-cargo-title": "Getting Cargo",
            "installing-cargo-lnk": "Installation",
            "installing-cargo-title": "Installation",
            "files-struct-cargo-lnk": "Files Structure",
            "files-struct-cargo-title": "Files Structure",
            "hello-cargo-lnk": "Hello Cargo!",
            "hello-cargo-title": "Hello Cargo!",
            "getting-started-title": "Getting Started with Cargo",
            "getting-started-p1": "After reading this guide you will know:",
            "how-to-intall-li": "How to install and configure your Cargo application server.",
            "general-structure-li": "The general structure of a Cargo application server.",
            "create-simple-app-li": "How to create a simple application that uses basic parts of the framework.",
            "prerequisites-content": "To start developing web applications with Cargo, you will need some knowledge of HTML, CSS and JavaScript. Because the back-end "
            + "of Cargo was witten with Go, it and has no external dependencies so there nothing more than Cargo itself to install."
            + "However, you will not need to know or use Go to build a web application with Cargo.",
            "cargo-web-server-git-lnk": "Visit our Github page",
            "cargo-linux-64-download-lnk": "Cargo v1.0 linux 64bit",
            "cargo-windows-64-download-lnk": "Cargo v1.0 windows 64bit",
            "installing-cargo-content": "Unzip the file where you want your sever to be and voîlà! You have a server ready to run. "
            + "</br>To start the server simply call the command <code>./CargoWebServer</code> in linux, or <code>CargoWebServer.exe</code> "
            + "in windows. Note that the server might take several minutes to load the first time.</br>"
            + "When the indexation is done, Bridge will open. To login, in the user name input enter <kbd>admin</kbd> and in the "
            + "password input enter <kbd>adminadmin</kbd>.",
            "files-struct-intro": "Let's now explore the files we got in <span>CargoWebServer</span> directory.",
            "files-structure-description": "In order to work properly, the <span>WebApp</span> folder and the sever executable file should be at the same directory level. If you need to separate "
            + "these tow files, you must define a global variable named CARGOROOT and make it point to <span>WebApp</span> directory."
            + "The application directory <span>Apps</span> contain applications files. From the browser point of view it's the root "
            + "of your server file system. With a running server, clicking the link <a href=\"http://127.0.0.1:9393\">Apps</span></a>"
            + " display the <span>Apps</span> folder content.</br>Here's a list of server directories description: </br>",
            "lib-folder-desc-li": "The <span>lib</span> folder is the place to put external libraries used by all your applications."
            + "<span> E.g. underscore, jQuery, ...</span>",
            "cargo-folder-desc-li": "The <span>Cargo</span> folder contain the code used by all applications to interact with the server. We will see "
            + "later which files must be included by application.",
            "data-folder-desc-li": "The <span>Data</span> folder contain applications data. The <span>MimeTypeIcon</span> directory contain icon's "
            + "named (mime type extension).png </br>** Never delete files <span>colorName.csv</span> and <span>mimeType.csv</span>",
            "schema-folder-desc-li": "The <span>Schemas</span> folder contain XML schemas used by applications. "
            + "</br>** Never delete files <span>xs.xsd</span> and <span>SqlTypes.xsd</span> ",
            "script-folder-desc-li": "The <span>Script</span> folder contain sever side JavaScript files. ",
            "hello-cargo-intro-p1": "With Cargo up and runing on your computer, it's time to create our first application. The application will contain one input "
            + "box and one button. The user will enter a value in the input box and when it click on the button the server will "
            + "respond with a greeting composed from that value. You can get the code <a href=\"http://cargowebserver.com/example/HelloCargo.zip\">here</a> "
            + " or try it <a href=\"http://54.218.110.52:9393/HelloCargo/\">here</a>. ",
            "hello-cargo-intro-p2": "First of all, we will create the directory inside the <span>Apps</span> folder. Let's name it <span>HelloCargo</span>. "
            + "Inside it we will create tow folders, one named <span>css</span> the other named <span>js</span>. I like to regroup "
            + "Javascript files in one folder and style sheets in other folder, but there's no rules about it. In the <span>HelloCargo</span> "
            + "directory create a new file named <span>index.html</span>. Each application must have an <span>index</span> file "
            + "at it top level directory. The <span>index</span> file has tow part, the header and the body.</br> Let's get a look of it: ",
            "index-html-desc-p": "The first tow <span>meta</span> elements defines how the content of the application must be interpreted by "
            + "the browser. The title element must contain your application name. Other lines of the header contains various stylesheets "
            + "and script files to be imported. Those files are needed by all applications in order to work. In the body element "
            + "I import specif application files, in your case only the main file are import.",
            "main-js-desc-p": "Every C++ or Java programmer know the main function. In those language it represent the entry point of all program. JavaScript dosen't have a main function by itself. Cargo impose the use of a main "
            + "dosen't have a main function by itself. Cargo impose the use of a main function because before starting your application "
            + "a lot of work must be done, like opening a connection with the server, tranfering files, initializing data structures "
            + "to name a few. So let's get a look of the <scan>main.js</scan> file.",
            "hello-cargo-conclusion-p" : "Dont worry if you can't understand all at once, the point here is to get a idea of what is Cargo. As you see "
            +"there is no trace of http, in fact there no trace of protocol at all, the only thing we know is the <span>server</span> "
            +"run it and return the result. The rest of the code was to create the user interface."
        },
        //francais
        "fr": {
            "prerequisites-lnk": "Prérequis",
            "prerequisites-title": "Prérequis",
            "getting-cargo-lnk": "Obtenir Cargo",
            "getting-cargo-title": "Obtenir Cargo",
            "installing-cargo-lnk": "Installation",
            "installing-cargo-title": "Installation",
            "files-struct-cargo-lnk": "Structure des fichiers",
            "files-struct-cargo-title": "Structure des fichiers",
            "hello-cargo-lnk": "Hello Cargo!",
            "hello-cargo-title": "Hello Cargo!",
            "getting-started-title": "Demarrer avec Cargo!",
            "getting-started-p1": "Vous apprenderez dans ce guide:",
            "how-to-intall-li": "À installer et configurer votre serveur Cargo",
            "general-structure-li": "La sturcture général d'un serveur d'application Cargo",
            "create-simple-app-li": "Les bases de la création d'une application web avec Cargo.",
            "prerequisites-content": "Pour dévellopper des application web avec Cargo il est nécessaire de connaitre HTML, CSS et JavaScript."
            + "Comme le côté seveur de Cargo à été codé dans en language Go, il n'y a rien d'autre a installer que Cargo. De plus aucune connaissance"
            + "du langage Go est requise pour dévelloper des application web avec Cargo.",
            "cargo-web-server-git-lnk": "Visitez notre page sur GitHub",
            "cargo-linux-download-lnk": "Cargo v1.0 linux 64bit",
            "cargo-windows-64-download-lnk": "Cargo v1.0 windows 64bit",
            "installing-cargo-content": "Décompresser le dossier à l'endroit ou vous voulez installer votre serveur et voîlà! Vous avez un seveur prêt a être démarrer"
            + "</br>Pour démarrer votre serveur vous devez exécuter la commande <code>./CargoWebServer</code> sous linux, où <code>CargoWebServer.exe</code>"
            + "sous windows. Veuillez prendre note que le premier démarrage du serveur est plus long.</br>"
            + "Lorsque l'indexation des fichier sera complété, la console du serveur Bridge apparaitera. Pour vous identifier, entrer le nom d'usager <kbd>admin</kbd> dans la boîte user et "
            + " le mot de passe <kbd>adminadmin</kbd> dans la boîte d'entré du mot de passe.",
            "files-struct-intro": "Explorons maintenant le contenu du répertoire <span>CargoWebServer</span>",
            // A traduire en francais...
            "files-structure-description": "In order to work properly, the <span>WebApp</span> folder and the sever executable file should be at the same directory level. If you need to separate"
            + "these tow files, you must define a global variable named CARGOROOT and make it point to <span>WebApp</span> directory."
            + "The application directory <span>Apps</span> contain applications files. From the browser point of view it's the root "
            + "of your server file system. With a running server, clicking the link <a href=\"http://127.0.0.1:9393\">Apps</span></a>"
            + " display the <span>Apps</span> folder content.</br>Here's a list of server directories description: </br>",
            "lib-folder-desc-li": "The <span>lib</span> folder is the place to put external libraries used by all your applications."
            + "<span> E.g. underscore, jQuery, ...</span>",
            "cargo-folder-desc-li": "The <span>Cargo</span> folder contain the code used by all applications to interact with the server. We will see"
            + "later which files must be included by application.",
            "data-folder-desc-li": "The <span>Data</span> folder contain applications data. The <span>MimeTypeIcon</span> directory contain icon's"
            + "named (mime type extension).png </br>** Never delete files <span>colorName.csv</span> and <span>mimeType.csv</span>",
            "schema-folder-desc-li": "The <span>Schemas</span> folder contain XML schemas used by applications."
            + "</br>** Never delete files <span>xs.xsd</span> and <span>SqlTypes.xsd</span>",
            "script-folder-desc-li": "The <span>Script</span> folder contain sever side JavaScript files.",
            "hello-cargo-intro-p1": "With Cargo up and runing on your computer, it's time to create our first application. The application will contain one input"
            + "box and one button. The user will enter a value in the input box and when it click on the button the server will"
            + "respond with a greeting composed from that value. You can try it <a href=\"http://54.218.110.52:9393/HelloCargo/\">here</a>.",
            "hello-cargo-intro-p2": "First of all, we will create the directory inside the <span>Apps</span> folder. Let's name it <span>HelloCargo</span>."
            + "Inside it we will create tow folders, one named <span>css</span> the other named <span>js</span>. I like to regroup"
            + "Javascript files in one folder and style sheets in other folder, but there's no rules about it. In the <span>HelloCargo</span> "
            + "directory create a new file named <span>index.html</span>. Each application must have an <span>index</span> file"
            + "at it top level directory. The <span>index</span> file has tow part, the header and the body.</br> Let's get a look of it:",
            "index-html-desc-p": "The first tow <span>meta</span> elements defines how the content of the application must be interpreted by "
            + "the browser. The title element must contain your application name. Other lines of the header contains various stylesheets "
            + "and script files to be imported. Those files are needed by all applications in order to work. In the body element "
            + "I import specif application files, in your case only the main file are import.",
            "main-js-desc-p": "Every C++ or Java programmer know the main function. In those language it represent the entry point of all program. JavaScript dosen't have a main function by itself. Cargo impose the use of a main "
            + "dosen't have a main function by itself. Cargo impose the use of a main function because before starting your application"
            + "a lot of work must be done, like opening a connection with the server, tranfering files, initializing data structures"
            + "to name a few. So let's get a look of the <scan>main.js</scan> file.",
            "hello-cargo-conclusion-p" : "Dont worry if you can't understand all at once, the point here is to get a idea of what is Cargo. As you see "
            +"there is no trace of http, in fact there no trace of protocol at all, the only thing we know is the <span>server</span> "
            +"run it and return the result. The rest of the code was to create the user interface."
        }
    }

    // Depending of the language the correct text will be set.
    server.languageManager.appendLanguageInfo(languageInfo)

    // The getting started panel.
    this.panel = new Element(parent, { "tag": "div", "class": "row hidden pageRow", "id": "page-gettingStarted" })
        .appendElement({ "tag": "div", "class": "col-lg-12" }).down()

    this.links = this.panel.appendElement({ "tag": "div", "class": "col-sm-2", "role": "complementary" }).down()
        .appendElement({ "tag": "nav", "class": "bs-docs-sidebar affix-top" }).down()
        .appendElement({ "tag": "ul", "class": "nav bs-docs-sidenav" }).down()

    // Links
    this.links.appendElement({ "tag": "li" }).down()
        .appendElement({ "tag": "a", "href": "#prerequisitesDiv", "id": "prerequisites-lnk" }).up()
        // getting cargo lnk
        .appendElement({ "tag": "li" }).down()
        .appendElement({ "tag": "a", "href": "#gettingCargoDiv", "id": "getting-cargo-lnk" }).up()
        // installing cargo
        .appendElement({ "tag": "li" }).down()
        .appendElement({ "tag": "a", "href": "#installingCargoDiv", "id": "installing-cargo-lnk" }).up()
        // file structure
        .appendElement({ "tag": "li" }).down()
        .appendElement({ "tag": "a", "href": "#filesStructCargoDiv", "id": "files-struct-cargo-lnk" }).up()
        // hello cargo lnk
        .appendElement({ "tag": "li" }).down()
        .appendElement({ "tag": "a", "href": "#helloCargoDiv", "id": "hello-cargo-lnk" })

    // Now the introduction panel.
    this.main = this.panel.appendElement({ "tag": "div", "class": "col-sm-10", "role": "main" }).down()
        .appendElement({ "tag": "div", "class": "row" }).down()
        .appendElement({ "tag": "div", "class": "col-xs-12" }).down()

    this.intro = this.main
        .appendElement({ "tag": "div", "class": "jumbotron" }).down()
        .appendElement({ "tag": "div", "class": "container" }).down()
        .appendElement({ "tag": "h2", "id": "getting-started-title" })
        .appendElement({ "tag": "p" }).down()
        .appendElement({ "tag": "span", "id": "getting-started-p1" }).up()

        // List of think you will be able to do...
        .appendElement({ "tag": "ul" }).down()
        .appendElement({ "tag": "li", "id": "how-to-intall-li" })
        .appendElement({ "tag": "li", "id": "general-structure-li" })
        .appendElement({ "tag": "li", "id": "create-simple-app-li" })

    ////////////////////////////////////////////////////////////////////////
    // Section of the getting started page.
    ////////////////////////////////////////////////////////////////////////

    // Declaration.
    // prerequisites section. 
    this.prerequisitesSection = this.main
        .appendElement({ "tag": "div", "class": "row", "id": "prerequisitesDiv" }).down()
        .appendElement({ "tag": "div", "class": "col-xs-12" }).down()
        .appendElement({ "tag": "h2", "id": "prerequisites-title" })
        .appendElement({ "tag": "p", "id": "prerequisites-content" })

    // Getting cargo section.
    this.gettingCargoSection = this.main
        .appendElement({ "tag": "div", "class": "row", "id": "gettingCargoDiv" }).down()
        .appendElement({ "tag": "div", "class": "col-xs-12" }).down()
        .appendElement({ "tag": "h2", "id": "getting-cargo-title" })
        .appendElement({ "tag": "div" }).down()

    // Installing cargo section
    this.installingCargoSection = this.main
        .appendElement({ "tag": "div", "class": "row", "id": "installingCargoDiv" }).down()
        .appendElement({ "tag": "div", "class": "col-xs-12" }).down()
        .appendElement({ "tag": "h2", "id": "installing-cargo-title" })
        .appendElement({ "tag": "p", "id": "installing-cargo-content" }).down()

    // Files structure section
    this.FilesStructureSection = this.main
        .appendElement({ "tag": "div", "class": "row", "id": "filesStructCargoDiv" }).down()
        .appendElement({ "tag": "div", "class": "col-xs-12" }).down()
        .appendElement({ "tag": "h2", "id": "files-struct-cargo-title" })
        .appendElement({ "tag": "div" }).down()

    // Hello cargo section.
    this.HelloCargoSection = this.main
        .appendElement({ "tag": "div", "class": "row", "id": "helloCargoDiv" }).down()
        .appendElement({ "tag": "div", "class": "col-xs-12" }).down()
        .appendElement({ "tag": "h2", "id": "hello-cargo-title" })
        .appendElement({ "tag": "div" }).down()


    // content...

    // Getting cargo stuff...
    this.gettingCargoSection.appendElement({ "tag": "p" }).down()
        // git hub link
        .appendElement({ "tag": "i", "class": "fa fa-github" })
        .appendElement({ "tag": "a", "id": "cargo-web-server-git-lnk", "style": "padding-left: 5px;", "href": "https://github.com/CargoWebServer/CargoWebServer" })
        .appendElement({ "tag": "br" }).appendElement({ "tag": "br" })

        // Downlaod link's
        .appendElement({ "tag": "i", "class": "fa fa-linux" })
        .appendElement({ "tag": "a", "id": "cargo-linux-64-download-lnk", "style": "padding-left: 5px;", "href": "http://cargowebserver.com/distro/linux/64/CargoWebServer.tar.gz" })
        .appendElement({ "tag": "br" }).appendElement({ "tag": "br" })

        .appendElement({ "tag": "i", "class": "fa fa-windows" })
        .appendElement({ "tag": "a", "id": "cargo-windows-64-download-lnk", "style": "padding-left: 5px;", "href": "http://cargowebserver.com/distro/windows/64/CargoWebServer.zip" })

    // Files structure...
    this.FilesStructureSection.appendElement({ "tag": "p", "id": "files-struct-intro" }).down()
        // The top level
        .appendElement({ "tag": "div", "class": "fileStructure" }).down()

        // The exec file.
        .appendElement({ "tag": "div" }).down().appendElement({ "tag": "i", "class": "fa fa-cog", "aria-hidden": "true" })
        .appendElement({ "tag": "a", "href": "#", "data-toggle": "tooltip", "title": "Cargo executable file.", "innerHtml": "CargoWebServer(.exe)" }).up()

        // The webApp directory.
        .appendElement({ "tag": "div", "class": "fileStructure" })
        .appendElement({ "tag": "div", "style": "padding-top: 5px;" }).down()
        .appendElement({ "tag": "i", "class": "fa fa-folder-open-o", "aria-hidden": "true" }).down()
        .appendElement({ "tag": "span", "innerHtml": "WebApp" })

        // WebApp/Cargo
        .appendElement({ "tag": "div", "class": "fileStructure" }).down()
        .appendElement({ "tag": "div", "style": "padding-top: 5px;" }).down()
        .appendElement({ "tag": "i", "class": "fa fa-folder-open-o", "aria-hidden": "true" }).down()
        .appendElement({ "tag": "span", "innerHtml": "Cargo" })

        // WebApp/Cargo/Apps
        .appendElement({ "tag": "div", "class": "fileStructure" }).down()
        .appendElement({ "tag": "div", "style": "padding-top: 5px;" }).down()
        .appendElement({ "tag": "i", "class": "fa fa-folder-open-o", "aria-hidden": "true" }).down()
        .appendElement({ "tag": "span", "innerHtml": "Apps" })

        // WebApp/Cargo/Apps/Bridge   
        .appendElement({ "tag": "div", "class": "fileStructure" }).down()
        .appendElement({ "tag": "div", "style": "padding-top: 5px;" }).down()
        .appendElement({ "tag": "i", "class": "fa fa-folder-o", "aria-hidden": "true" }).down()
        .appendElement({ "tag": "span", "innerHtml": "Bridge" })
        .appendElement({ "tag": "br" })

        // WebApp/Cargo/Apps/lib
        .appendElement({ "tag": "div", "style": "padding-top: 5px;" }).down()
        .appendElement({ "tag": "i", "class": "fa fa-folder-o", "aria-hidden": "true" }).down()
        .appendElement({ "tag": "span", "innerHtml": "lib" })
        .appendElement({ "tag": "br" })

        // WebApp/Cargo/Apps/Test
        .appendElement({ "tag": "div", "style": "padding-top: 5px;" }).down()
        .appendElement({ "tag": "i", "class": "fa fa-folder-o", "aria-hidden": "true" }).down()
        .appendElement({ "tag": "span", "innerHtml": "Test" })
        .appendElement({ "tag": "br" })

        // WebApp/Cargo/Apps/Cargo project.
        .appendElement({ "tag": "div", "class": "fileStructure" }).down()
        .appendElement({ "tag": "div", "style": "padding-top: 5px;" }).down()
        .appendElement({ "tag": "i", "class": "fa fa-folder-open-o", "aria-hidden": "true" }).down()
        .appendElement({ "tag": "span", "innerHtml": "Cargo" })
        .appendElement({ "tag": "br" })

        // WebApp/Cargo/Apps/Cargo/css 
        .appendElement({ "tag": "div", "class": "fileStructure" }).down()
        .appendElement({ "tag": "div", "style": "padding-top: 5px;" }).down()
        .appendElement({ "tag": "i", "class": "fa fa-folder-o", "aria-hidden": "true" }).down()
        .appendElement({ "tag": "span", "innerHtml": "css" })
        .appendElement({ "tag": "br" })

        // WebApp/Cargo/Apps/Cargo/doc
        .appendElement({ "tag": "div", "style": "padding-top: 5px;" }).down()
        .appendElement({ "tag": "i", "class": "fa fa-folder-o", "aria-hidden": "true" }).down()
        .appendElement({ "tag": "span", "innerHtml": "doc" })
        .appendElement({ "tag": "br" })

        // WebApp/Cargo/Apps/Cargo/svg
        .appendElement({ "tag": "div", "style": "padding-top: 5px;" }).down()
        .appendElement({ "tag": "i", "class": "fa fa-folder-o", "aria-hidden": "true" }).down()
        .appendElement({ "tag": "span", "innerHtml": "svg" })
        .appendElement({ "tag": "br" })


        // WebApp/Cargo/Apps/Cargo/proto
        .up().up().up().up().up().up().up().up().up()
        .appendElement({ "tag": "div", "class": "fileStructure" }).down()
        .appendElement({ "tag": "div", "style": "padding-top: 5px;" }).down()
        .appendElement({ "tag": "i", "class": "fa fa-folder-open-o", "aria-hidden": "true" }).down()
        .appendElement({ "tag": "span", "innerHtml": "proto" })
        .appendElement({ "tag": "br" })

        // WebApp/Cargo/Apps/Cargo/proto/rpc.proto
        .appendElement({ "tag": "div", "class": "fileStructure" }).down()
        .appendElement({ "tag": "div", "style": "padding-top: 5px;" }).down()
        .appendElement({ "tag": "i", "class": "fa fa-file-code-o", "aria-hidden": "true" }).down()
        .appendElement({ "tag": "span" }).down()
        .appendElement({
            "tag": "a", "href": "https://github.com/CargoWebServer/CargoWebServer/blob/master/WebApp/Cargo/Apps/Cargo/proto/rpc.proto",
            "title": "Protocol definition", "innerHtml": "rpc.proto"
        }).up()
        .appendElement({ "tag": "br" })

        // WebApp/Cargo/Apps/Cargo/js
        .up().up().up().appendElement({ "tag": "div", "style": "padding-top: 5px;" }).down()
        .appendElement({ "tag": "i", "class": "fa fa-folder-o", "aria-hidden": "true" }).down()
        .appendElement({ "tag": "span", "innerHtml": "js" })
        .appendElement({ "tag": "br" })

        // WebApp/Cargo/Apps/Cargo/js/gui
        .appendElement({ "tag": "div", "class": "fileStructure" }).down()
        .appendElement({ "tag": "div", "style": "padding-top: 5px;" }).down()
        .appendElement({ "tag": "i", "class": "fa fa-folder-o", "aria-hidden": "true" }).down()
        .appendElement({ "tag": "span", "innerHtml": "gui" })
        .appendElement({ "tag": "br" })

        // WebApp/Cargo/Apps/Cargo/js/*.js
        .appendElement({ "tag": "div", "style": "padding-top: 5px;" }).down()
        .appendElement({ "tag": "i", "class": "fa fa-file-code-o", "aria-hidden": "true" }).down()
        .appendElement({
            "tag": "a", "href": "https://github.com/CargoWebServer/CargoWebServer/tree/master/WebApp/Cargo/Apps/Cargo/js",
            "title": "Javascript script files.", "innerHtml": "*.js"
        }).up()
        .appendElement({ "tag": "br" })

        // WebApp/Cargo/Data
        .up().up().up().up().up().up().up().up().up().up().up().up().up().up().up().up().up()
        .appendElement({ "tag": "div", "style": "padding-top: 5px;" }).down()
        .appendElement({ "tag": "i", "class": "fa fa-folder-open-o", "aria-hidden": "true" }).down()
        .appendElement({ "tag": "span", "innerHtml": "Data" })
        .appendElement({ "tag": "br" })

        // WebApp/Cargo/Data/MimeTypeIcon
        .appendElement({ "tag": "div", "class": "fileStructure" }).down()
        .appendElement({ "tag": "div", "style": "padding-top: 5px;" }).down()
        .appendElement({ "tag": "i", "class": "fa fa-folder-o", "aria-hidden": "true" }).down()
        .appendElement({ "tag": "span", "innerHtml": "MimeTypeIcon" })
        .appendElement({ "tag": "br" })

        // WebApp/Cargo/Data/MimeTypeIcon/colorName.csv 
        .appendElement({ "tag": "div", "style": "padding-top: 5px;" }).down()
        .appendElement({ "tag": "i", "class": "fa fa-file-code-o", "aria-hidden": "true" }).down()
        .appendElement({ "tag": "span" }).down()
        .appendElement({
            "tag": "a", "href": "https://github.com/CargoWebServer/CargoWebServer/blob/master/WebApp/Cargo/Data/colorName.csv",
            "title": "Protocol definition", "innerHtml": "colorName.csv"
        }).up()
        .appendElement({ "tag": "br" })

        // WebApp/Cargo/Data/MimeTypeIcon/mimeType.csv
        .appendElement({ "tag": "div", "style": "padding-top: 5px;" }).down()
        .appendElement({ "tag": "i", "class": "fa fa-file-code-o", "aria-hidden": "true" }).down()
        .appendElement({ "tag": "span" }).down()
        .appendElement({
            "tag": "a", "href": "https://github.com/CargoWebServer/CargoWebServer/blob/master/WebApp/Cargo/Data/mimeType.csv",
            "title": "Protocol definition", "innerHtml": "mimeType.csv"
        }).up()
        .appendElement({ "tag": "br" })

        // WebApp/Cargo/Data/Schemas
        .up().up().up().up().up().up().up().up().up().up()
        .appendElement({ "tag": "div", "style": "padding-top: 5px;" }).down()
        .appendElement({ "tag": "i", "class": "fa fa-folder-open-o", "aria-hidden": "true" }).down()
        .appendElement({ "tag": "span", "innerHtml": "Schemas" })
        .appendElement({ "tag": "br" })

        // WebApp/Cargo/Data/Schemas/xs.xsd
        .appendElement({ "tag": "div", "class": "fileStructure" }).down()
        .appendElement({ "tag": "div", "style": "padding-top: 5px;" }).down()
        .appendElement({ "tag": "i", "class": "fa fa-file-code-o", "aria-hidden": "true" }).down()
        .appendElement({ "tag": "span" }).down()
        .appendElement({
            "tag": "a", "href": "https://github.com/CargoWebServer/CargoWebServer/blob/master/WebApp/Cargo/Schemas/xs.xsd",
            "title": "Protocol definition", "innerHtml": "xs.xsd"
        }).up()
        .appendElement({ "tag": "br" })

        // WebApp/Cargo/Data/Schemas/SqlTypes.xsd
        .appendElement({ "tag": "div", "style": "padding-top: 5px;" }).down()
        .appendElement({ "tag": "i", "class": "fa fa-file-code-o", "aria-hidden": "true" }).down()
        .appendElement({ "tag": "span" }).down()
        .appendElement({
            "tag": "a", "href": "https://github.com/CargoWebServer/CargoWebServer/blob/master/WebApp/Cargo/Schemas/SqlTypes.xsd",
            "title": "Protocol definition", "innerHtml": "SqlTypes.xsd"
        }).up()
        .appendElement({ "tag": "br" })

        .up().up().up().up().up().up()
        .appendElement({ "tag": "div", "style": "padding-top: 5px;" }).down()
        .appendElement({ "tag": "i", "class": "fa fa-folder-o", "aria-hidden": "true" }).down()
        .appendElement({ "tag": "span", "innerHtml": "Script" })
        .appendElement({ "tag": "br" })

    // The file structure description.
    this.FilesStructureSection.appendElement({ "tag": "p", "id": "files-structure-description" })
        .appendElement({ "tag": "ul" }).down()
        .appendElement({ "tag": "li", "id": "lib-folder-desc-li" })
        .appendElement({ "tag": "li", "id": "cargo-folder-desc-li" })
        .appendElement({ "tag": "li", "id": "data-folder-desc-li" })
        .appendElement({ "tag": "li", "id": "schema-folder-desc-li" })
        .appendElement({ "tag": "li", "id": "script-folder-desc-li" })

    //////////////////////////////////////////////////////////////////
    // Hello cargo.
    //////////////////////////////////////////////////////////////////
    this.HelloCargoSection
        .appendElement({ "tag": "p", "id": "hello-cargo-intro-p1" })
        .appendElement({ "tag": "p", "id": "hello-cargo-intro-p2" })

    //////////////////////////// Html index //////////////////////////
    this.indexHtmlCode = this.HelloCargoSection.appendElement({ "tag": "pre" }).down()
        .appendElement({ "tag": "code", "class": "xml hljs" }).down()

    // Append the meta element   
    appendHtmlMeta(this.indexHtmlCode, "&lt;!DOCTYPE html&gt;")

    // Open brace...
    appendHtmlTag(this.indexHtmlCode, 0, "html")
    appendHtmlTag(this.indexHtmlCode, 1, "head", [{ "name": "lang", "value": "en" }])

    appendHtmlTag(this.indexHtmlCode, 2, "meta", [{ "name": "http-equiv", "value": "x-ua-compatible" }, { "name": "content", "value": "IE=edge" }])
    appendHtmlTag(this.indexHtmlCode, 2, "meta", [{ "name": "charset", "value": "UTF-8" }])
    appendHtmlTag(this.indexHtmlCode, 2, "title", [], true, "HelloCargo")

    // css
    appendHtmlComment(this.indexHtmlCode, "css", 2)

    // Closing...
    appendHtmlTag(this.indexHtmlCode, 2, "link", [{ "name": "rel", "value": "stylesheet" }, { "name": "type", "value": "text/css" }, { "name": "href", "value": "/Cargo/css/main.css" }])

    // js
    appendHtmlComment(this.indexHtmlCode, "js", 2)

    appendHtmlTag(this.indexHtmlCode, 2, "script", [{ "name": "src", "value": "/lib/protobuf/bytebuffer.js" }], true, "")
    appendHtmlTag(this.indexHtmlCode, 2, "script", [{ "name": "src", "value": "/lib/protobuf/long.js" }], true, "")
    appendHtmlTag(this.indexHtmlCode, 2, "script", [{ "name": "src", "value": "/lib/protobuf/protobuf.js" }], true, "")
    appendHtmlTag(this.indexHtmlCode, 2, "script", [{ "name": "src", "value": "/lib/rollups/sha256.js" }], true, "")
    appendHtmlTag(this.indexHtmlCode, 2, "script", [{ "name": "src", "value": "/lib/stringview.js" }], true, "")

    appendHtmlComment(this.indexHtmlCode, "Base Server files", 2)
    appendHtmlTag(this.indexHtmlCode, 2, "script", [{ "name": "src", "value": "/Cargo/js/utility.js" }], true, "")
    appendHtmlTag(this.indexHtmlCode, 2, "script", [{ "name": "src", "value": "/Cargo/js/handlers.js" }], true, "")
    appendHtmlTag(this.indexHtmlCode, 2, "script", [{ "name": "src", "value": "/Cargo/js/error.js" }], true, "")
    appendHtmlTag(this.indexHtmlCode, 2, "script", [{ "name": "src", "value": "/Cargo/js/event.js" }], true, "")
    appendHtmlTag(this.indexHtmlCode, 2, "script", [{ "name": "src", "value": "/Cargo/js/entity.js" }], true, "")
    appendHtmlTag(this.indexHtmlCode, 2, "script", [{ "name": "src", "value": "/Cargo/js/languageManager.js" }], true, "")
    appendHtmlTag(this.indexHtmlCode, 2, "script", [{ "name": "src", "value": "/Cargo/js/server.js" }], true, "")
    appendHtmlTag(this.indexHtmlCode, 2, "script", [{ "name": "src", "value": "/Cargo/js/main.js" }], true, "")

    appendHtmlComment(this.indexHtmlCode, "GUI", 2)
    appendHtmlTag(this.indexHtmlCode, 2, "script", [{ "name": "src", "value": "/Cargo/js/element.js" }], true, "")

    appendHtmlComment(this.indexHtmlCode, "HelloCargo files", 2)
    appendHtmlTag(this.indexHtmlCode, 2, "link", [{ "name": "rel", "value": "stylesheet" }, { "name": "type", "value": "text/css" }, { "name": "href", "value": "css/main.css" }])
    appendHtmlTag(this.indexHtmlCode, 1, "head", [], false)

    appendHtmlTag(this.indexHtmlCode, 1, "body")
    appendHtmlComment(this.indexHtmlCode, "HelloCargo files ", 2)
    appendHtmlTag(this.indexHtmlCode, 2, "script", [{ "name": "src", "value": "js/main.js" }], true, "")
    appendHtmlTag(this.indexHtmlCode, 1, "body", [], false)

    appendHtmlTag(this.indexHtmlCode, 0, "html", [], false)
    //////////////////////////// end Html index //////////////////////////

    // The index file description.
    this.HelloCargoSection.appendElement({ "tag": "p", "id": "index-html-desc-p" })

    // The main.js description.
    this.HelloCargoSection.appendElement({ "tag": "p", "id": "main-js-desc-p" })

    // Now the content of the main.js file.
    this.mainJsCode = this.HelloCargoSection.appendElement({ "tag": "pre" }).down().appendElement({ "tag": "code", "class": "javascript hljs" }).down()

    // So here I will get the code from the file itself...
    server.fileManager.getFileByPath("/HelloCargo/js/main.js",
        // success callback
        function (result, caller) {
            /** here I will set the innerHtml with the file content. */
            caller.element.innerHTML = decode64(result.M_data)
            hljs.highlightBlock(caller.element);
        },
        // error callback
        function (errMsg, caller) {
            /** Nothing to do here. */
        },
        this.mainJsCode)

    // Conclusion code.
    this.HelloCargoSection.appendElement({ "tag": "p", "id": "hello-cargo-conclusion-p" })

    return this
}
