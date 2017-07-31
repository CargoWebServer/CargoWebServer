
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
            "tutorial-buoy-desc" :"The buoy symbol is use to warning you about common mistake or bad programming practices.",
            "tutorial-anchor-desc":"The anchor symbol are use for complementary explanation or discution that can be skip by hurry reader's.",
            "SOM-tutorial-title": "The Service Object Model (SOM)",
            "SOM-tutorial-lnk": "SOM tutorial",
            "SOM-tutorial-content": ""
        },
        //francais
        "fr": {
            "tutorial-title": "Cargo Tutorial",
            "tutorial-presentation": "Cargo Dactilogiciel",
            "tutorial-presentation-description": "",
            "blog-tutorial-title": "",
            "blog-tutorial-content": "",
            "blog-tutorial-lnk": "",
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
        .appendElement({ "tag": "a", "href": "#somTutorialDiv", "id": "blog-tutorial-lnk" }).up()


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
    this.main
        .appendElement({ "tag": "div", "class": "row", "id": "SOM_TutorialDiv" }).down()
        .appendElement({ "tag": "div", "class": "col-xs-12" }).down()
        .appendElement({ "tag": "h2", "id": "SOM-tutorial-title" })
        .appendElement({ "tag": "p", "id": "SOM-tutorial-content" })
}