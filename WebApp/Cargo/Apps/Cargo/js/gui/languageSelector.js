/**
 * Append language selector i parent div. The languages are JSON object containing information: 
 * ex [{"flag":"us", "name":"English"}, {"flag":"fr", "name":"Frainçais"}]
 */
var LanguageSelector = function(parent, languages, callback){

    this.languages = languages;

    this.panel = parent.appendElement({"tag":"div", "class":"language-selector"}).down()
    
    this.currentLanguageDiv = this.panel.appendElement({"tag":"div", "style":"display: table; border-spacing:2px 2px; position: relative;"}).down()
    this.flagDiv = this.currentLanguageDiv.appendElement({"tag":"div", "style":"display: table-cell;"}).down()
  
    this.currentLanguage = null;
    this.languageListDiv = null;

    if(languages.length > 0){
        this.setLanguage(languages[0]);
    }

    this.flagDiv.element.onclick = function(selector){
        return function(){
            selector.displayLanguages();
        }
    }(this)

    // Call when the language change.
    this.callback = callback

    return this;
}

/** Set the current language. */
LanguageSelector.prototype.setLanguage = function(language){
    this.flagDiv.element.className = "flag-div flag-icon flag-icon-" + language.flag;
    this.flagDiv.element.title = language.name;
    // keep track of the current language...
    this.currentLanguage = language;
    if(this.languageListDiv != null){
        this.languageListDiv.parentElement.removeElement(this.languageListDiv)
        this.languageListDiv = null
    }
    // call the callback function.
    if(this.callback != null){
        this.callback(language)
    }
}

LanguageSelector.prototype.displayLanguages = function(){
    // remove the list if one already exist.
    if(this.languageListDiv != null){
        this.languageListDiv.parentElement.removeElement(this.languageListDiv)
        this.languageListDiv = null
        return;
    }

    this.languageListDiv = this.panel.appendElement({"tag":"div", "id":"language-list-div"}).down()
    this.languageListDiv.element.style.top = this.flagDiv.element.offsetHeight + 5 + "px"

    var panelLeft = getCoords(this.panel.element).left
    if(panelLeft == 0){
        this.languageListDiv.element.style.left = "0px"
    }

    for(var i=0; i < this.languages.length; i++){
        if(this.currentLanguage.name != this.languages[i].name){
        var row = this.languageListDiv.appendElement({"tag":"div", "class":"language-list-div-row", "style":"display: table-row;"}).down()
            var languageSpan = row.appendElement({"tag":"span","style":"display: table-cell;", "innerHtml":this.languages[i].name}).down()
            row.appendElement({"tag":"div", "style":"display: table-cell;", "class":"flag-icon flag-icon-" + this.languages[i].flag})
            
            languageSpan.element.onclick = function(selector, language){
                return function(){
                    selector.setLanguage(language)
                }
            }(this, this.languages[i])
        }
    }
}