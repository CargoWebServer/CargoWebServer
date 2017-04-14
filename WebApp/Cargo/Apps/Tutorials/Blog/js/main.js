var applicationName = document.getElementsByTagName("title")[0].text
var languageInfo = {
    "en": {
    },
    "fr": {
    }
}

// Depending of the language the correct text will be set.
server.languageManager.appendLanguageInfo(languageInfo)

// TODO create la language bar.
server.languageManager.setLanguage("en")

// This is the body that will be use by all other panel.
var bodyElement = new Element(document.getElementsByTagName("body")[0], {"tag":"div", "style":"width: 100%; height: 100%;"})

var blogManager = null

/**
 * This function is the entry point of the application...
 */
function main() {

    // get the prototypes of the blog schema.
    server.entityManager.getEntityPrototypes("sql_info",
    // success callback
    function(results, caller){
        blogManager = new BlogManager(bodyElement)
    },
    // error callback.
    function(){

    }, 
    // caller.
    {} )
}