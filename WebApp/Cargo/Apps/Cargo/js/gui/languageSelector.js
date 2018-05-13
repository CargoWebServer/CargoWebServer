/**
 * Append language selector i parent div. The languages are JSON object containing information: 
 * ex [{"flag":"us", "name":"English"}, {"flag":"fr", "name":"Frainçais"}]
 */
var LanguageSelector = function(parent, languages, callback){

    this.languages = languages;

    this.panel = parent.appendElement({"tag":"div", "class":"language-selector dropdown"}).down()
    
    this.currentLanguageDiv = this.panel.appendElement({"tag":"button", "id":"language_select_btn", 
    "class":"current-language btn btn-outline-secondary dropdown-toggle", "type":"button", "data-toggle":"dropdown", 
    "aria-haspopup":"true", "aria-expanded" : "false", "style":"background-color: transparent; border: none;"}).down()
    
    // This is the button.
    this.flagDiv = this.currentLanguageDiv.appendElement({"tag":"div"}).down()
    this.currentLanguageSpan = this.currentLanguageDiv.appendElement({"tag":"span", "style":"margin-left: 5px;"}).down()
    
    this.languageListDiv = this.panel.appendElement({"tag":"ul", "id":"language-list-div", "class":"dropdown-menu dropdown-menu-right bg-dark", "role":"menu", "aria-labelledby":"language_select_btn"}).down()
    
    var panelLeft = getCoords(this.panel.element).left
    if(panelLeft == 0){
        this.languageListDiv.element.style.left = "0px"
    }

    for(var i=0; i < this.languages.length; i++){
        var row = this.languageListDiv.appendElement({"tag":"li", "class":"language-list-div-row", "role":"presentation"}).down()
            var item = row.appendElement({"tag":"div", "class":"menuitem"}).down();
            item.appendElement({"tag":"div", "class":"flag-icon flag-icon-" + this.languages[i].flag})
            var languageSpan = item.appendElement({"tag":"span", "innerHtml":this.languages[i].name}).down()
            row.element.onclick = function(selector, language){
                return function(){
                    selector.setLanguage(language)
                }
            }(this, this.languages[i])
    }
    
    if(languages.length > 0){
        this.setLanguage(languages[0]);
    }

    // Call when the language change.
    this.callback = callback

    return this;
}

/** Set the current language. */
LanguageSelector.prototype.setLanguage = function(language){
    this.flagDiv.element.className = "flag-div flag-icon flag-icon-" + language.flag;
    this.currentLanguageSpan.element.innerText = language.name;
    
    // keep track of the current language...
    this.currentLanguage = language;

    // call the callback function.
    if(this.callback != null){
        this.callback(language)
    }
}
