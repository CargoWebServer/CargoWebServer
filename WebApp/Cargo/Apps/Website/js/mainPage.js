
var languageInfo = {
    "en": {
        "IntroductionText": "Create by applications programmers for applications programmers, Cargo is a complete web application framework. Cargo make web programming intuitive and fun. Written in Go, Cago is fast, easy to configure and lightweight. All Cargo functionnalities are exposed to applications in JavaScript, the official browser language.",
        "AboutText": "<span style=\"font-size: 12pt;\">If you want people to understand you, speak their language. –African Proverb </span></br></br>Programming a desktop application is not easy, programming a web application is even more complicated. For desktop application it make sense to use programming language like Java or C++ but when we talk about web application there is only one contender and it's JavaScript. So if you want to be a good web application programmer you must, at least, know Javascript. With JavaScript and the document object model, you will be able to interact with the browser and the end user. But what about the server? Why server's use languages other than JavaScript like PHP, Java or C#... Wouldn't it be simpler to have a server object model SOM and interact with it the same way DOM interact with the browser?</br>This is Cargo! ",
        "DownloadText": "Download",
        "DocsText": "Documentation here...",
        "NewsText": "News here...",
    },
    "fr": {
        "IntroductionText": "Create by applications programmers for applications programmers, Cargo is a complete web application framework. Cargo make web programming intuitive and fun. Written in Go, Cago is fast, easy to configure and lightweight. All Cargo functionnalities are exposed to applications in JavaScript, the official browser language.",
        "AboutText": "<span style=\"font-size: 12pt;\">If you want people to understand you, speak their language. –African Proverb </span></br></br>Programming a desktop application is not easy, programming a web application is even more complicated. For desktop application it make sense to use programming language like Java or C++ but when we talk about web application there is only one contender and it's JavaScript. So if you want to be a good web application programmer you must, at least, know Javascript. With JavaScript and the document object model, you will be able to interact with the browser and the end user. But what about the server? Why server's use languages other than JavaScript like PHP, Java or C#... Wouldn't it be simpler to have a server object model SOM and interact with it the same way DOM interact with the browser?</br>This is Cargo! ",
        "DownloadText": "Download",
        "DocsText": "Documentation here...",
        "NewsText": "News here...",
    }
}

// Set the text...
server.languageManager.appendLanguageInfo(languageInfo)


/**
 * This contain the main website page.
 */
var MainPage = function (parent) {

    // keep the ref to the main layout.
    this.parent = parent

    this.panel = this.parent.appendElement({ "tag": "div", "style": "display: table; width: 100%; height:100%;" }).down()

    // The header section...
    this.header = this.panel.appendElement({ "tag": "div", "class": "header" }).down()

    // The body section
    this.body = this.panel.appendElement({ "tag": "div", "class": "body" }).down()

    // The footer section
    this.footer = this.panel.appendElement({ "tag": "div", "class": "footer" }).down()

    // The header...
    var titleDiv = this.header.appendElement({ "tag": "div", "class": "title" }).down()
    //var title = titleDiv.appendElement({ "tag": "div", "innerHtml": "Cargo" }).down()

    // The logo...
    //title.appendElement({ "tag": "img", "src": "/Website/image/wheel.svg", "style": "width:60px; height:60px;vertical-align:middle;" })

    // Now I will append the menu...
    var mainMenuDiv = this.header.appendElement({ "tag": "div", "class": "mainMenu" }).down()

    // Create the menu item's
    var homeMenuItem = mainMenuDiv.appendElement({ "tag": "div", "class": "mainMenuItem", "innerHtml": "Home", "style": "border-right: 1px solid #ccc;" }).down()
    var downloadsMenuItem = mainMenuDiv.appendElement({ "tag": "div", "class": "mainMenuItem", "innerHtml": "Downloads", "style": "border-right: 1px solid #ccc;" }).down()
    var newsMenuItem = mainMenuDiv.appendElement({ "tag": "div", "class": "mainMenuItem", "innerHtml": "News", "style": "border-right: 1px solid #ccc;" }).down()
    var tutorialsMenuItem = mainMenuDiv.appendElement({ "tag": "div", "class": "mainMenuItem", "innerHtml": "Tutorials", "style": "border-right: 1px solid #ccc;" }).down()
    var documentationsMenuItem = mainMenuDiv.appendElement({ "tag": "div", "class": "mainMenuItem", "innerHtml": "Docs", "style": "border-right: 1px solid #ccc;" }).down()
    var aboutMenuItem = mainMenuDiv.appendElement({ "tag": "div", "class": "mainMenuItem", "innerHtml": "About", "style": "" }).down()


    // Now The page content.
    var homeIntroItem = new Element(null, { "tag": "div", "id": "home-intro" })

    // Now I will write the text inside the panel...
    function appendSection(menuItem, textId, grid, callback) {
        var sectionItem = new Element(null, { "tag": "div", "id": "home-intro" })

        var p = sectionItem.appendElement({ "tag": "p" }).down()
        server.languageManager.setElementText(p, textId)
        menuItem.element.onclick = function (sectionItem, grid, callback) {
            return function () {
                // Set the content!
                grid.removeAllChilds()
                grid.moveChildElements([sectionItem])
                if (callback != undefined) {
                    callback(sectionItem)
                }
            }
        } (sectionItem, grid, callback)

        return sectionItem
    }

    // The geneneral introduciton...
    var homeSectionElement = appendSection(homeMenuItem, "IntroductionText", this.body,
        function (footer) {
            return function (homeSectionElement) {
                // Reset the initial position...
                document.getElementsByTagName("body")[0].style.overflowY = "auto"
                footer.element.style.display = ""
                var homeIntro = document.getElementById("home-intro")
                homeIntro.style.maxWidth = ""
            }
        } (this.footer))

    // The download button...
    homeSectionElement.appendElement({ "tag": "div", "style": "text-align: center;" }).down()
        .appendElement({ "tag": "div", "style": "font-size: 18pt; font-weight: 600;padding:10px 30px 10px 30px;", "innerHtml": "</br>Download for " + navigator.platform })
        .appendElement({ "tag": "i", "id": "osLogo", "style": "font-size: 28pt;" })
        .appendElement({ "tag": "div", "class": "home-downloadbutton", "id": "downloadButton", "innerHtml": "v1.0.1" }).down()
        .appendElement({ "tag": "small", "innerHtml": "the first version!" })
        .up().up()
        .appendElement({ "tag": "span", "style": "font-size: 18pt;", "innerHtml": "Welcome aboard we expecting you!" })

    // Here I will append the icon...
    if (navigator.platform == "Linux x86_64" || navigator.platform == "Linux x86_32") {
        homeSectionElement.getChildById("osLogo").element.className = "fa fa-linux"
    } else if (navigator.platform == "Windows x86_64" || navigator.platform == "Windows x86_32") {
        homeSectionElement.getChildById("osLogo").element.className = "fa fa-windows"
    }

    homeSectionElement.getChildById("downloadButton").element.onclick = function () {
        // Here is the list of available distro...

        // So here I will open the link...
        if (navigator.platform == "Linux x86_64") {
            // So here I will get the file at link,
            var lnk = new Element(null, { "tag": "a", "href": "http://" + server.ipv4 + ":" + server.port + "/Dist/Linux/x86_64/CargoWebserver" })
            lnk.element.click()
        } else if (navigator.platform == "Linux x86_32") {
            // So here I will get the file at link,
            var lnk = new Element(null, { "tag": "a", "href": "http://" + server.ipv4 + ":" + server.port + "/Dist/Linux/x86_32/CargoWebserver" })
            lnk.element.click()
        } else if (navigator.platform == "Windows x86_64") {
            // So here I will get the file at link,
            var lnk = new Element(null, { "tag": "a", "href": "http://" + server.ipv4 + ":" + server.port + "/Dist/Windows/x86_64/CargoWebserver" })
            lnk.element.click()
        } else if (navigator.platform == "Windows x86_32") {
            // So here I will get the file at link,
            var lnk = new Element(null, { "tag": "a", "href": "http://" + server.ipv4 + ":" + server.port + "/Dist/Windows/x86_32/CargoWebserver" })
            lnk.element.click()
        }
    }

    // Set the text...
    homeMenuItem.element.click()


    // The about section...
    var aboutSectionElement = appendSection(aboutMenuItem, "AboutText", this.body,
        function (footer) {
            return function (aboutSectionElement) {
                // Here I will set the editor...
                var editor = ace.edit("editor");
                editor.setTheme("ace/theme/monokai");
                editor.getSession().setMode("ace/mode/javascript");
                aboutSectionElement.element.parentNode.style.overflowY = "auto"
                footer.element.style.display = "none"
                document.getElementsByTagName("body")[0].style.overflowY = ""
                var homeIntro = document.getElementById("home-intro")
                homeIntro.style.maxWidth = ""
            }
        } (this.footer)
    )

    // Here I will create the example...
    var editor = aboutSectionElement.appendElement({ "tag": "div", "id": "editor", "style": "height: 410px; text-align: left;" }).down()
    editor.element.innerHTML = "/**\n * Calculate the sum of tow numbers, it will be execute by the server.\n */\n"
    editor.element.innerHTML += "function add(a, b){\n  return a + b; \n}\n"
    editor.element.innerHTML += "\n/**\n *The main entry point called when the server is ready \n */\n"
    editor.element.innerHTML += "function main(){\n"
    editor.element.innerHTML += "	// I will ask the server to calculate 2 + 4...\n"
    editor.element.innerHTML += "	var a = 2;\n"
    editor.element.innerHTML += "	var b = 4;\n"
    editor.element.innerHTML += "	var params = [];\n"
    editor.element.innerHTML += "	// wrapp variable in RpcData to give cargo a hint about data types.\n"
    editor.element.innerHTML += "	params.push(new RpcData({ \"name\": \"a\", \"type\": 1, \"dataBytes\": utf8_to_b64(a) }));\n"
    editor.element.innerHTML += "	params.push(new RpcData({ \"name\": \"b\", \"type\": 1, \"dataBytes\": utf8_to_b64(b) }));\n"
    editor.element.innerHTML += "	// Execute the function by the server\n"
    editor.element.innerHTML += "	server.executeJsFunction( add.toString(), params,\n"
    editor.element.innerHTML += "	    // The progress callback.\n"
    editor.element.innerHTML += "	    function(index, total, caller){\n"
    editor.element.innerHTML += "	        console.log(\"progress:\" + (index/progress) + \"%\");\n"
    editor.element.innerHTML += "	    },\n"
    editor.element.innerHTML += "	    // The success callback.\n"
    editor.element.innerHTML += "	    function(result, caller){\n"
    editor.element.innerHTML += "	        console.log(\"sum is:\" + result[0]);\n"
    editor.element.innerHTML += "	    },\n"
    editor.element.innerHTML += "	    // The error callback.\n"
    editor.element.innerHTML += "	    function(errorMsg, caller){\n"
    editor.element.innerHTML += "	        console.log(\"error:\" + errorMsg);\n"
    editor.element.innerHTML += "	    },\n"
    editor.element.innerHTML += "       {/* put caller variable here to get back their references in the response context. */});\n"
    editor.element.innerHTML += "}"

    // The download section
    var downlaodSection = appendSection(downloadsMenuItem, "DownloadText", this.body,
        function (footer) {
            return function (downloadSectionElement) {
                document.getElementsByTagName("body")[0].style.overflowY = "auto"
                var homeIntro = document.getElementById("home-intro")
                homeIntro.style.maxWidth = ""
                footer.element.style.display = ""
            }
        } (this.footer))

    // The documentation section
    var documentationSectionElement = appendSection(documentationsMenuItem, "DocsText", this.body,
        function (footer) {
            return function (documentationSectionElement) {
                var page = "http://" + server.ipv4 + ":" + server.port + "/Cargo/doc/index.html"
                documentationSectionElement.element.style = "max-width: 100%; position: absolute; top: 129px; bottom: 0px; left: 0px; right: 0px;"
                documentationSectionElement.element.innerHTML = '<object style="width: 100%" data="' + page + '" type="text/html"><embed src="' + page + '" type="text/html" /></object>';
                documentationSectionElement.element.firstChild.style.height = documentationSectionElement.element.offsetHeight + "px"
                documentationSectionElement.element.parentNode.style.overflow = "hidden"
                footer.element.style.display = "none"
                document.getElementsByTagName("body")[0].style.overflowY = "hidden"
                var homeIntro = document.getElementById("home-intro")
                homeIntro.style.maxWidth = "100%"
                // Now the windows resize event...
                window.addEventListener('resize', function (documentationSectionElement) {
                    return function (event) {
                        documentationSectionElement.element.firstChild.style.height = documentationSectionElement.element.offsetHeight + "px"
                    }
                } (documentationSectionElement));

            }
        } (this.footer))


    // The news section
    appendSection(newsMenuItem, "NewsText", this.body)

    // The tutorials section
    var tutorialsSectionElement = appendSection(
        tutorialsMenuItem,
        "Tutorials",
        this.body,
        function (footer) {
            return function (tutorialsSectionElement) {
                tutorialsSectionElement.element.style = "max-width: 100%; position: absolute; top: 129px; bottom: 0px; left: 0px; right: 0px;"
                tutorialsSectionElement.element.parentNode.style.overflow = "hidden"
                footer.element.style.display = "none"
                tutorialsSectionElement.removeAllChilds()
                var homeIntro = document.getElementById("home-intro")
                homeIntro.style.maxWidth = "100%"
                homeIntro.style.minHeight = "100%"

                var tutorialsBanner = tutorialsSectionElement.appendElement({ "tag": "div", "class": "tutorialsBanner", "innerHtml": "TUTORIALS" }).down()

                var tutorialsFlexConainer = tutorialsSectionElement.appendElement({ "tag": "div", "class": "flexboxRowContainer" }).down()

                var tutorialsSidebarContainer = tutorialsFlexConainer.appendElement({ "tag": "div", "class": "tutorialsSidebarContainer" }).down()
                var tutorialsSidebarMenu = tutorialsSidebarContainer.appendElement({ "tag": "div", "id": "tutorialsSidebarMenu", "class": "flexboxColumnContainer" }).down()

                tutorialsSidebarMenu.appendElement({ "tag": "div", "class": "sidebarMenuTitle", "innerHtml": "Getting Started" })
                    .appendElement({ "tag": "div", "class": "sidebarMenuItem", "innerHtml": "Welcome", "id": "WelcomeTutorialMenuDiv" })
                    .appendElement({ "tag": "div", "class": "sidebarMenuItem", "innerHtml": "Licence" })
                    .appendElement({ "tag": "div", "class": "sidebarMenuItem", "innerHtml": "Installation" })
                    .appendElement({ "tag": "div", "class": "sidebarMenuItem", "innerHtml": "Configuration" })
                    .appendElement({ "tag": "div", "class": "sidebarMenuTitle", "innerHtml": "Basics" })
                    .appendElement({ "tag": "div", "class": "sidebarMenuItem", "innerHtml": "'Hello World' Project" })
                    .appendElement({ "tag": "div", "class": "sidebarMenuItem", "innerHtml": "Front-End JS Framework" })
                    .appendElement({ "tag": "div", "class": "sidebarMenuItem", "innerHtml": "Back-End Golang Server" })
                    .appendElement({ "tag": "div", "class": "sidebarMenuItem", "innerHtml": "Element", "id": "ElementTutorialMenuDiv" })
                    .appendElement({ "tag": "div", "class": "sidebarMenuTitle", "innerHtml": "Core Modules" })
                    .appendElement({ "tag": "div", "class": "sidebarMenuItem", "innerHtml": "Events" })
                    .appendElement({ "tag": "div", "class": "sidebarMenuItem", "innerHtml": "Entities" })
                    .appendElement({ "tag": "div", "class": "sidebarMenuItem", "innerHtml": "Data" })
                    .appendElement({ "tag": "div", "class": "sidebarMenuTitle", "innerHtml": "Advanced Tutorials" })
                    .appendElement({ "tag": "div", "class": "sidebarMenuItem", "innerHtml": "MVC Applications" })
                    .appendElement({ "tag": "div", "class": "sidebarMenuItem", "innerHtml": "Permissions" })
                    .appendElement({ "tag": "div", "class": "sidebarMenuItem", "innerHtml": "..." })

                var tutorialsContentContainer = tutorialsFlexConainer.appendElement({ "tag": "div", "class": "tutorialsContentContainer" }).down()

                tutorialsSectionElement.getChildById("ElementTutorialMenuDiv").element.onclick = function (tutorialsContentContainer) {
                    return function () {
                        tutorialsContentContainer.removeAllChilds()
                        new elementTutorial(tutorialsContentContainer)
                    }
                } (tutorialsContentContainer)

            }
        } (this.footer)
    )
    return this

}