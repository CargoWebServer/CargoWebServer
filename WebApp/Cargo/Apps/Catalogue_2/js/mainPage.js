/** Text contain in the main page **/
var mainPageText = {
    "en": {
        "quickSearchText": " Quick search",
        "welcomeText" : "Welcome to the catalog!",
         "welcomeTipText1" : "To find an item, type a keyword in the search bar, or use the shortcuts ",
        "welcomeTipText2" : " for a category search and ",
        "welcomeTipText3" : " for a filtered search.",
        "advancedSearchText": " Advanced search",
        "navbarSearchText" : "Search in the catalog...",
        "recentSearchText" : "Recent searches",
        "supplierSearchLnk" : "by supplier",
        "localisationSearchLnk" : "by localisation",
        "generalSearchLnk" : "générale"
    },
    "fr": {
        "quickSearchText": " Recherche rapide",
        "welcomeText" : "Bienvenue dans le catalogue!",
        "welcomeTipText1" : "Pour trouver un item, tapez un mot-clé dans la barre de recherche ou utilisez les raccourcis ",
        "welcomeTipText2" : " pour une recherche par catégorie et ",
        "welcomeTipText3" : " pour une recherche filtrée.",
        "advancedSearchText" : " Recherche avancée",
        "navbarSearchText" : "Recherche dans le catalogue...",
        "recentSearchText" : "Recherches récentes",
        "supplierSearchLnk" : "par fournisseur",
        "localisationSearchLnk" : "par localisation",
        "generalSearchLnk" : "générale"
    },
    
    "es" : {
        "quickSearchText" : " Búsqueda rápida",
        "welcomeText" : "¡Bienvenido al catálogo!",
        "advancedSearchText" : " Búsqueda avanzada",
        "welcomeTipText1" : "Para encontrar un artículo, escriben una palabra clave en la barra de búsqueda o usen las atajos ",
        "welcomeTipText2" : " para una búsqueda de categoría y ",
        "welcomeTipText3" : " para una búsqueda filtrada.",
        "navbarSearchText" : "Buscar en el catálogo ...",
        "recentSearchText" : "Búsquedas recientes",
        "supplierSearchLnk" : "",
        "localisationSearchLnk" : "",
        "generalSearchLnk" : ""
    }
};

// Set the text info.
server.languageManager.appendLanguageInfo(mainPageText);


/** Contain the main page layout **/
var MainPage = function () {
    // Set special search context to general (item's).
    this.searchContext = "generalSearchLnk";
    
    // The panel.
    this.panel = new Element(document.body, { "tag": "div", "id": "main-page" });
    
    // Building the navbar with all the content inside
    this.panel.appendElement({"tag" : "nav", "class" : "navbar navbar-expand-md navbar-fixed-top navbar-dark bg-dark", "style" : "z-index : 10;", "id" : "top-navigation"}).down()
    .appendElement({"tag" : "div", "class" : "container-fluid"}).down()
    .appendElement({"tag" : "div", "class" : "navbar-header"}).down()
 
    .appendElement({"tag" : "a", "class" : "navbar-brand"}).down()
    .appendElement({"tag" : "img", "src" : "image/safran.png","style" : "width : 24px;", "class" : "nav-item"}).up()

    //.appendElement({"tag" : "a", "class" : "navbar-brand navbar-text", "href" : "#","style" : "font-family: 'Questrial', sans-serif;", "innerHtml" : "Catalogue"})
    .appendElement({"tag" : "button" ,"class" : "btn btn-default text-light", "style" :"font-size : 0.8rem;cursor:pointer;", "onclick" : "openNav()"}).down()
    .appendElement({"tag" : "i", "class" : "fa fa-navicon"})
    .appendElement({"tag" : "span", "id" : "quickSearchText"}).up().up()
    .appendElement({"tag" : "button", "class" : "navbar-toggler", "type" : "button", "data-toggle": "collapse", "data-target" : "#collapsibleNavbar"}).down()
    .appendElement({"tag" : "span", "class" : "navbar-toggler-icon"}).up()
    .appendElement({"tag" : "div", "class" : "collapse navbar-collapse", "id" : "collapsibleNavbar"}).down()
    
    .appendElement({"tag" : "div", "class": "ml-auto", "id" : "searchbardiv"}).down()
    .appendElement({"tag" : "div", "class" : "navbar-form navbar-center", "id" : "searchbar"}).down()
    .appendElement({"tag" : "div", "class" : "input-group"}).down()
    .appendElement({"tag" : "input", "type" : "text", "class" : "form-control","style" : "font-family: 'Questrial', sans-serif;","id" : "navbarSearchText"})
    .appendElement({"tag" : "div", "class" : "input-group-btn"}).down()

    .appendElement({"tag" : "button", "class" : "btn btn-default text-light", "data-toggle" : "dropdown","type" : "button", "aria-expanded" : "false", "role": "menu", "id" : "advancedSearchButtonNavbar"}).down()
    .appendElement({"tag" : "i", "class" : "fa fa-caret-down", "style":"padding-right: 5px;"})
    .appendElement({"tag" : "span" , "id" : "advancedSearchText"}).up()
    .appendElement({"tag" : "div", "class" : "dropdown-menu bg-dark text-light border border-secondary", "id" : "advancedSearchCollapse"}).down()
    .appendElement({"tag":"a", "class":"dropdown-item bg-dark", "id":"generalSearchLnk"})
    .appendElement({"tag":"a", "class":"dropdown-item bg-dark", "id":"supplierSearchLnk"})
    .appendElement({"tag":"a", "class":"dropdown-item bg-dark", "id":"localisationSearchLnk"}).up()
    

    .appendElement({"tag" : "button", "type" : "button", "class" : "btn btn-default","id" :"searchButton"}).down()
    .appendElement({"tag" : "i", "class" :"fa fa-search"}).up().up().up().up().up()
    .appendElement({"tag" : "ul", "class" : "nav navbar-nav ml-auto"}).down()
    .appendElement({"tag" : "li", "id":"sessionPanelDiv", "class" : "navItem"})
    .appendElement({"tag" : "li", "class": "nav-item"}).down()
    .appendElement({"tag" : "div", "id" : "languageSelector", "class" : "navItem"}).up();
    
    new SessionPanel(this.panel.getChildById("sessionPanelDiv"))
     
    new LanguageSelector( this.panel.getChildById("languageSelector"),
        [{ "flag": "ca", "name": "français", "id": "fr" }, { "flag": "us", "name": "english", "id": "en" }, { "flag": "es", "name": "español", "id" : "es" }],
        function (language) {
            // set the language.
            server.languageManager.setLanguage(language.id)
        }
    );
    
    this.panel.appendElement({"tag" : "div", "class" : "row contentrow", "style" : "height : calc(100vh - 56px); width : 100%;"}).down()
    .appendElement({"tag" : "div", "class" : "sidenav", "id" : "mySideNav", "style" : "background-color : #343a40; position:fixed;top:50px;"}).down()
    .appendElement({"tag" : "a" , "href" : "#", "onclick" : "openNav()","class" : "closebtn", "innerHtml" : "&times", "style":"position:absolute;"}).up()
    .appendElement({"tag" : "span", "id" : "main", "class" : "col", "style" : "overflow-y : scroll;"}).down()
    .appendElement({"tag" : "div", "class" : "container", "style" : "height:100vh; padding-top : 20px;", "id" : "main-container"}).down()
    .appendElement({"tag" : "div","class" : "jumbotron bg-dark text-light border border-dark align-items-center", "style" : "margin: 20px; padding : 2rem 1rem;"}).down()
    .appendElement({"tag" : "h1", "class" : "display-5", "id" : "welcomeText", "style" : "text-align : center;"})
    .appendElement({"tag" : "hr", "class" : "my-4 dark"})
    .appendElement({"tag" : "p", "class" : "lead"}).down()
    .appendElement({"tag" : "span","id":"welcomeTipText1"})
    .appendElement({"tag" : "button" ,"class": "btn btn-outline-light disabled", "style":"opacity : 1; font-size : 0.8rem;cursor:default;"}).down()
    .appendElement({"tag" : "i", "class" :"fa fa-navicon"})
    .appendElement({"tag" : "span","id":"quickSearchText"}).up()
    .appendElement({"tag" : "span","id":"welcomeTipText2"})
    .appendElement({"tag":"button", "class": "btn btn-outline-light disabled", "style":"opacity :1;font-size : 0.8rem;cursor:default;"}).down()
    .appendElement({"tag":"i", "class":"fa fa-angle-double-down"})
    .appendElement({"tag" :"span","id":"advancedSearchText"}).up()
    .appendElement({"tag" : "span","id":"welcomeTipText3"}).up()
    .appendElement({"tag" : "p", "style" : "text-align:center;"}).down()
    .appendElement({"tag" : "hr", "class" : "my-4"}).up().up()
    .appendElement({"tag" : "div","class" : "row", "style" : "justify-content : space-between;"}).down()
    .appendElement({"tag" : "div", "class" : "container-fluid"}).down()
    .appendElement({"tag" : "div", "class" : "row contentrow ", "style" : "justify-content:space-between;"}).down()
    .appendElement({"tag" : "div","class" :"col-md-5 card bg-light mb-3 border border-dark align-items-center", "style" : "margin:20px; padding:0;"}).down()
    .appendElement({"tag" : "div", "class" : "card-header bg-dark text-light", "style" : "width:100%;"}).down()
    .appendElement({"tag" : "h1", "class" : "card-title smallertitle", "id" : "recentSearchText", "style" : "text-align:center;"}).up()
    .appendElement({"tag" : "div", "class" : "card-body", "style" : "width:100%;padding:0;"}).down()
    .appendElement({"tag" : "ul", "class" :"list-group list-group-flush"}).down()
    .appendElement({"tag" : "li", "class" : "list-group-item"})
    .appendElement({"tag" : "li", "class" : "list-group-item"})
    .appendElement({"tag" : "li", "class" : "list-group-item"})
    .appendElement({"tag" : "li", "class" : "list-group-item"}).up().up()
    .appendElement({"tag" : "hr", "class" : "my-4"}).up()
    .appendElement({"tag" : "div","class" : "col card border border-dark align-items-center", "style" : "margin:20px;"}).down()
    .appendElement({"tag" : "h1", "class" : "card-body display-5 smallertitle"})
    .appendElement({"tag" : "hr", "class" : "my-4"}).up().up().up().up()
    .appendElement({"tag" : "div", "id" : "notificationDiv" , "style" : "z-index : 100;position:fixed; width: 275px; right:2%;top:10%;"}).down()
    .up().up()
     openNav();
    this.contentPanel = this.panel.getChildById("main")
    this.welcomePage = this.panel.getChildById("main-container")
    catalogMessageHub.attach(this, welcomeEvent, function (evt, mainPage) {
        if(server.sessionId == evt.dataMap.welcomeEvent.sessionId){
            var user = evt.dataMap.welcomeEvent.user
            showNotification("primary", "Bienvenue "+ user.M_firstName + " !",4000)
           
        }else{
            var user = evt.dataMap.welcomeEvent.user
            showNotification("primary", user.M_firstName + " s'est connecté au catalogue!",4000)

        }
        
    })
    
    this.panel.getChildById("searchButton").element.onclick = function () {
        try{
            mainPage.welcomePage.element.parentNode.removeChild(mainPage.welcomePage.element)
        }
        catch (error){
            
        }
        spinner.panel.element.style.display = "";
        // Here I will display the seach results.
        mainPage.searchItems(mainPage.panel.getChildById("navbarSearchText").element.value)
    }
    
    this.panel.getChildById("navbarSearchText").element.onkeyup = function (evt) {
        if (evt.keyCode == 13) {
            try{
                mainPage.welcomePage.element.parentNode.removeChild(mainPage.welcomePage.element)
            }
            catch (error){
            
            }
            spinner.panel.element.style.display = "";
            // Here I will display the seach results.
            mainPage.searchItems(mainPage.panel.getChildById("navbarSearchText").element.value)
        }
    }
    
    
    this.searchResultPage = new SearchResultPage(this.contentPanel)
    this.itemDisplayPage = new ItemDisplayPage(this.contentPanel)

    // Advance search.
    var supplierLnk = this.panel.getChildById("supplierSearchLnk");
    var localisationLnk = this.panel.getChildById("localisationSearchLnk")
    var generalSearchLnk = this.panel.getChildById("generalSearchLnk")
    

    generalSearchLnk.element.onclick = supplierLnk.element.onclick = localisationLnk.element.onclick = function(contentPanel, advancedSearchText){
        return function(){
            mainPage.searchContext = this.id
            if(this.id == "generalSearchLnk"){
                server.languageManager.refresh()
            }else{
                advancedSearchText.element.innerHTML = this.innerHTML
            }
            document.getElementById("navbarSearchText").focus()
        }
    }(this.contentPanel, this.panel.getChildById("advancedSearchText"))
    

    this.itemDisplayPage.panel.element.style.display = "none";
    return this;
}

/**
 * The context can be general, supplier or localisation.
 */
MainPage.prototype.searchItems = function (keyword) {
    // Set specific fields here.
    var fields = []
    console.log(service.conn)
    // First of a ll I will clear the search panel.
    xapian.search(
        dbpaths,
        keyword.toUpperCase(),
        fields,
        "en",
        0,
        1000,
        // success callback
        function (results, caller) {
            mainPage.searchResultPage.displayResults(results, caller.keyword, mainPage.searchContext );
            //mainPage.searchInput.element.value = ""
            //mainPage.searchInput.element.focus()
        },
        // error callback
        function () {

        }, {"keyword":keyword })
}

function showNotification(alertType, message,delay){
    var id = randomUUID();
    var msg =  mainPage.panel.getChildById("notificationDiv").appendElement({"tag" : "div", "class" : "alert alert-"+alertType+" alert-dismissible", "role" : "alert"}).down()
    msg.appendElement({"tag" : "p","class" : "mb-0", "innerHtml" : message})
    .appendElement({"tag" : "button","id" : id, "class" : "close", "aria-label" : "Close"}).down()
    .appendElement({"tag" : "span", "aria-hidden" :"true", "innerHtml" : "&times;"}).up().up()
    msg.element.style.opacity = 0;
    //msg.element.style.transform = "translateX("+msg.element.offsetWidth+"px)"
    msg.element.style.left = msg.element.offsetWidth + "px";
    var keyframe = "100% {left:0px; opacity:1;}"
    msg.animate(keyframe, .75, 
        function(msg){
            return function(){
                 msg.element.style.opacity = 1;
                 msg.element.style.left = "0px";
               // msg.element.parentNode.removeChild(msg.element);
            }
    }(msg))
    msg.getChildById(id).element.onclick = function(msg){
        return function(){
            var keyframe = "100% {transform: translateX("+msg.element.offsetWidth+"px); opacity:0;}"
            msg.animate(keyframe, .75, 
            function(msg){
                return function(){
                    msg.element.parentNode.removeChild(msg.element);
                }
            }(msg))
            
        }
    }(msg)
    if(delay != undefined){
        setTimeout(function(closebtn) {
            return function(){
                closebtn.element.click();
            }
        }(msg.getChildById(id)), delay);
    }
}



function openNav() {
    if(!document.getElementById("mySideNav").classList.contains("closedNav")){
        document.getElementById("mySideNav").classList.remove("openedNav");
        document.getElementById("mySideNav").classList.add("closedNav");
    }else{
        document.getElementById("mySideNav").classList.remove("closedNav");
        document.getElementById("mySideNav").classList.add("openedNav");
    }
}
