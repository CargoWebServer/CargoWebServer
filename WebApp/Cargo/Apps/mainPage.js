/** Contain the main page layout **/
var MainPage = function () {

    // The panel.
    this.panel = new Element(document.body, { "tag": "div", "id": "main-page" });
    
    /*this.navBar = this.panel.appendElement({"tag" : "nav", "class" : "navbar navbar-expand-lg navbar-dark bg-dark navbar-fixed-top", "style" : ""}).down();
    var navbarContainer = this.navBar.appendElement({"tag" : "div", "class" : "container"}).down();
    
    var navbarHeader = navbarContainer.appendElement({"tag" : "div", "class" : "navbar-header"}).down();
   
    navbarHeader.appendElement({"tag" : "img", "src" : "image/safran.png", "style" : "width : 32px; height: 32px; margin-right : 10px;"});
    navbarHeader.appendElement({"tag" : "a", "class" : "navbar-brand", "href" : "#", "innerHtml" : "Catalogue"});
     var navbarCollapse = navbarContainer.appendElement({"tag" : "div" , "class" : "navbar-collapse collapse", "aria-expanded" : "false", "style:"}).down();
    var leftSection = navbarCollapse.appendElement({"tag" : "ul", "class" : "nav navbar-nav"}).down();
    var rightSection = navbarCollapse.appendElement({"tag" : "ul", "class" : "nav navbar-nav navbar-right"}).down();
    //this.navBar.appendElement({"tag" : "div", "class"})
    // Here I will set the language selector...
    var languageSelector = new LanguageSelector( rightSection.appendElement({"tag" : "li", "class" : "", "style" : ""}).down(),
        [{ "flag": "ca", "name": "français", "id": "fr" }, { "flag": "us", "name": "english", "id": "en" }, { "flag": "es", "name": "español" }],
        function (language) {
            // set the language.
            server.languageManager.setLanguage(language.id)
        }
    );
    rightSection.appendElement({"tag" : "li", "style" : "padding: 5px;"})
    rightSection.appendElement({"tag" : "li", "class" : "navbar-item", "innerHtml" : "Connection"})*/
    
    this.panel.appendElement({"tag" : "nav", "class" : "navbar navbar-expand-md navbar-dark bg-dark"}).down()
//    .appendElement({"tag": "div", "class" : "container", "style" : "flex-wrap : nowrap;"}).down()
    .appendElement({"tag" : "div", "class" : "navbar-header"}).down()
 
    .appendElement({"tag" : "a", "class" : "navbar-brand"}).down()
    .appendElement({"tag" : "img", "src" : "image/safran.png","style" : "width : 24px; align-items : center;"}).up()

    .appendElement({"tag" : "a", "class" : "navbar-brand", "href" : "#", "innerHtml" : "Catalogue"}).up()
    .appendElement({"tag" : "button", "class" : "navbar-toggler", "type" : "button", "data-toggle": "collapse", "data-target" : "#collapsibleNavbar"}).down()
    .appendElement({"tag" : "span", "class" : "navbar-toggler-icon"}).up()
    .appendElement({"tag" : "div", "class" : "collapse navbar-collapse justify-content-end", "id" : "collapsibleNavbar"}).down()
    .appendElement({"tag" : "ul", "class" : "navbar-nav"}).down()
    .appendElement({"tag" : "li", "class" : "nav-item"}).down()
    .appendElement({"tag" : "a", "class" : "nav-link", "href" : "#", "innerHtml" : "Connection"}).up()
    .appendElement({"tag" : "li", "class": "nav-item"}).down()
    .appendElement({"tag" : "div", "class" : "nav-link", "id" : "languageSelector"});
    
    
     new LanguageSelector( this.panel.getChildById("languageSelector"),
        [{ "flag": "ca", "name": "français", "id": "fr" }, { "flag": "us", "name": "english", "id": "en" }, { "flag": "es", "name": "español" }],
        function (language) {
            // set the language.
            server.languageManager.setLanguage(language.id)
        }
    );
    
    
    return this;
    
    
}