
/**
 * Funtion that create the home page.
 */
var HomePage = function (parent) {
    // Language of the starting page.
    var languageInfo = {
        "en": {
            "first-msg": "Cargo is an open-source, complete web application framework. It's fast, easy and lightweight.",
            "getStarted-btn": "Get started!",
            "feature-complete-title": "Complete",
            "feature-complete-desc": "Front-end, server, database. All in one.",
            "feature-easy-title": "Easy setup",
            "feature-easy-desc": "One download.",
            "feature-smooth-title": "Smooth learning curve",
            "feature-smooth-desc": "Code your web application in only one language: javascript.",
            "feature-lightweight-title": "Lightweight, fast, non-blocking",
            "feature-lightweight-desc": "Backend written in Go.",
            "getting-started-lnk": "Getting Started",
            "tutorials-lnk": "Tutorials",
            "docs-lnk": "Documentation",
            "about-lnk": "About",
        },
        "fr": {
            "first-msg": "Cargo est une platforme de création d'application web au code source ouvert. Cargo est rapide, simple et léger.",
            "getStarted-btn": "Commencer!",
            "feature-complete-title": "Complet",
            "feature-complete-desc": "Interface, serveur, base de donnée. Tout en un.",
            "feature-easy-title": "Facile a installer",
            "feature-easy-desc": "Un simple téléchargement",
            "feature-smooth-title": "Douce courbe d'apprentisage",
            "feature-smooth-desc": "Coder vos applications web dans un langage: javascript.",
            "feature-lightweight-title": "Léger, rapide, disponible",
            "feature-lightweight-desc": "Côté serveur entitèrement écrit en Go.",
            "getting-started-lnk": "démarrer",
            "tutorials-lnk": "Didacticiels",
            "docs-lnk": "Documentation",
            "about-lnk": "À propos",
        }
    }
    // Depending of the language the correct text will be set.
    server.languageManager.appendLanguageInfo(languageInfo)

    // Create the main div
    this.panel = new Element(parent, { "tag": "div", "class": "container-fluid" })

    // The navigation bar.
    this.pages = this.panel.appendElement({ "tag": "div", "class": "row" }).down()

    this.navBar = this.pages.appendElement({ "tag": "div", "class": "navbar navbar-default" }).down()
        .appendElement({ "tag": "nav", "class": "container-fluid" }).down()


    this.navBar.appendElement({ "tag": "div", "class": "navbar-header" }).down()
        .appendElement({ "tag": "button", "class": "navbar-toggle collapsed", "data-toggle": "collapse", "data-target": "#bs-example-navbar-collapse-1", "aria-expanded": "false" }).down()
        .appendElement({ "tag": "div", "class": "sr-only", "innerHtml": "Toggle navigation" })
        .appendElement({ "tag": "span", "class": "icon-bar" })
        .appendElement({ "tag": "span", "class": "icon-bar" })
        .appendElement({ "tag": "span", "class": "icon-bar" }).up()
        .appendElement({ "tag": "a", "class": "navbar-brand ", "href": "#", "id": "a-cargo", "innerHtml": "Cargo" }).up()
        .appendElement({ "tag": "div", "class": "collapse navbar-collapse", "id": "bs-example-navbar-collapse-1" }).down()
        .appendElement({ "tag": "ul", "class": "nav navbar-nav" }).down()
        .appendElement({ "tag": "li", "id": "li-gettingStarted", "class":"nav_button" }).down()
        .appendElement({ "tag": "a", "href": "#", "id": "getting-started-lnk" }).up()
        .appendElement({ "tag": "li", "id": "li-tutorials", "class":"nav_button" }).down()
        .appendElement({ "tag": "a", "href": "#", "id": "tutorials-lnk" }).up()
        .appendElement({ "tag": "li", "id": "li-docs", "class":"nav_button" }).down()
        .appendElement({ "tag": "a", "href": "#", "id": "docs-lnk" }).up().up()
        .appendElement({ "tag": "ul", "class": "nav navbar-nav navbar-right" }).down()
        .appendElement({ "tag": "li", "id": "li-about", "class":"nav_button" }).down()
        .appendElement({ "tag": "a", "href": "#", "id": "about-lnk" })

    // The main page content.
    this.pages.appendElement({ "tag": "div", "class": "row pageRow", "id": "page-cargo" }).down()
        .appendElement({ "tag": "div", "class": "row" }).down()
        .appendElement({ "tag": "div", "class": "col-xs-12 text-center", "id": "cargo-intro-container" }).down()
        .appendElement({ "tag": "div", "class": "col-xs-12" }).down()
        .appendElement({ "tag": "img", "src": "img/wheel.svg", "id": "cargo-intro-wheel" }).up()
        .appendElement({ "tag": "div", "class": "col-sm-6 col-sm-offset-3 text-center" }).down()
        .appendElement({ "tag": "span" }).appendElement({ "tag": "h2", "id": "first-msg" }).appendElement({ "tag": "br" })
        .appendElement({ "tag": "p" }).appendElement({ "tag": "a", "id": "getStarted-btn", "class": "btn btn-primary btn-lg", "href": "#", "role": "button" }).up().up()
        .appendElement({ "tag": "div", "class": "row feature" }).down()
        // Complet
        .appendElement({ "tag": "div", "class": "col-md-3" }).down()
        .appendElement({ "tag": "span", "class": "glyphicon glyphicon glyphicon-ok" })
        .appendElement({ "tag": "h2", "id": "feature-complete-title" })
        .appendElement({ "tag": "p", "id": "feature-complete-desc" }).up()
        // easy
        .appendElement({ "tag": "div", "class": "col-md-3" }).down()
        .appendElement({ "tag": "span", "class": "glyphicon glyphicon glyphicon-play" })
        .appendElement({ "tag": "h2", "id": "feature-easy-title" })
        .appendElement({ "tag": "p", "id": "feature-easy-desc" }).up()
        // Smooth
        .appendElement({ "tag": "div", "class": "col-md-3" }).down()
        .appendElement({ "tag": "span", "class": "glyphicon glyphicon glyphicon-road" })
        .appendElement({ "tag": "h2", "id": "feature-smooth-title" })
        .appendElement({ "tag": "p", "id": "feature-smooth-desc" }).up()
        // lightweight
        .appendElement({ "tag": "div", "class": "col-md-3" }).down()
        .appendElement({ "tag": "span", "class": "glyphicon glyphicon glyphicon-flash" })
        .appendElement({ "tag": "h2", "id": "feature-lightweight-title" })
        .appendElement({ "tag": "p", "id": "feature-lightweight-desc" }).up()

    function HidePages(notHide) {
        var pages = document.getElementsByClassName("pageRow")
        for (var i = 0; i < pages.length; i++) {
            if (notHide == pages[i].id) {
                pages[i].className = "row pageRow"
            } else {
                pages[i].className = "row hidden pageRow"
            }
        }
    }

    // Now the actions.

    // Go home
    this.navBar.getChildById("a-cargo").element.onclick = function () {
        HidePages("page-cargo")
        var lnks = document.getElementsByClassName("active")
        for (var i = 0; i < lnks.length; i++) {
            lnks[i].className = lnks[i].className.replace(" active", "")
        }
    }

    // Getting started.
    this.navBar.getChildById("li-gettingStarted").element.onclick = this.pages.getChildById("getStarted-btn").element.onclick = function () {
        HidePages("page-gettingStarted")
        var lnks = document.getElementsByClassName("active")
        for (var i = 0; i < lnks.length; i++) {
            lnks[i].className = lnks[i].className.replace(" active", "")
        }
        this.className += " active"
    }

    // Tutorials.
    this.navBar.getChildById("li-tutorials").element.onclick = function () {
        HidePages("page-cargoTutorial")
        var lnks = document.getElementsByClassName("active")
        for (var i = 0; i < lnks.length; i++) {
            lnks[i].className = lnks[i].className.replace(" active", "")
        }
        this.className += " active"
    }

    // Documentation
    this.navBar.getChildById("li-docs").element.onclick = function () {
        HidePages("page-cargoDocumentation")
        var lnks = document.getElementsByClassName("active")
        for (var i = 0; i < lnks.length; i++) {
            lnks[i].className = lnks[i].className.replace(" active", "")
        }
        this.className += " active"
    }

    // About.

    return this
}