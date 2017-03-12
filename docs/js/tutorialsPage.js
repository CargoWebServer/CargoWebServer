
var TutorialsPage = function (parent) {
    // Text of the page.
    var languageInfo = {
        // english
        "en": {
            "tutorial-title": "Cargo Tutorial",
            "tutorial-presentation": "Cargo Tutorial",
            "tutorial-presentation-description": "",
            "website-tutorila-title": "",
            "website-tutorial-content": "",
            "website-tutorial-lnk": "Website tutorial",
            "blog-tutorila-title": "",
            "blog-tutorial-content": "",
            "blog-tutorial-lnk": "Blog tutorial",
        },
        //francais
        "fr": {
            "tutorial-title": "Cargo Tutorial",
            "tutorial-presentation": "Cargo Dactilogiciel",
            "tutorial-presentation-description": "",
            "website-tutorila-title": "",
            "website-tutorial-content": "",
            "website-tutorial-lnk": "",
            "blog-tutorila-title": "",
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
    this.links.appendElement({ "tag": "li" }).down()
        .appendElement({ "tag": "a", "href": "#websiteTutorialDiv", "id": "website-tutorial-lnk" }).up()
        // getting cargo lnk
        .appendElement({ "tag": "li" }).down()
        .appendElement({ "tag": "a", "href": "#blogTutorialDiv", "id": "blog-tutorial-lnk" }).up()


    // Now the introduction panel.
    this.main = this.panel.appendElement({ "tag": "div", "class": "col-sm-10", "role": "main" }).down()
        .appendElement({ "tag": "div", "class": "row" }).down()
        .appendElement({ "tag": "div", "class": "col-xs-12" }).down()

    // Introduction text.
    this.intro = this.main
        .appendElement({ "tag": "div", "class": "jumbotron" }).down()
        .appendElement({ "tag": "div", "class": "container" }).down()
        .appendElement({ "tag": "h2", "id": "tutorial-title" })
        .appendElement({ "tag": "p" }).down()
        .appendElement({ "tag": "span", "id": "tutorial-presentation" }).up()

        // Description of the tutorials.
        .appendElement({ "tag": "p", "id": "tutorial-presentation-description" })

    this.websiteTutorialSection = this.main
        .appendElement({ "tag": "div", "class": "row", "id": "websiteTutorialDiv" }).down()
        .appendElement({ "tag": "div", "class": "col-xs-12" }).down()
        .appendElement({ "tag": "h2", "id": "website-tutorila-title" })
        .appendElement({ "tag": "p", "id": "website-tutorial-content" })

    this.blogTutorialSection = this.main
        .appendElement({ "tag": "div", "class": "row", "id": "blogTutorialDiv" }).down()
        .appendElement({ "tag": "div", "class": "col-xs-12" }).down()
        .appendElement({ "tag": "h2", "id": "blog-tutorila-title" })
        .appendElement({ "tag": "p", "id": "blog-tutorial-content" })
}