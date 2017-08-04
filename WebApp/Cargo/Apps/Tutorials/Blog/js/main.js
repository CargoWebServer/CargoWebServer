var applicationName = document.getElementsByTagName("title")[0].text

// global variable.
var databaseName = "Blog."
var schemaId = "" //"dbo."

var userTypeName = databaseName + schemaId + "blog_user"
var authorTypeName = databaseName + schemaId + "blog_author"
var blogPostTypeName = databaseName + schemaId + "blog_post"
var blogCommentTypeName = databaseName + schemaId + "blog_comment"
var categoryTypeName = databaseName + schemaId + "blog_category"

// TODO create la language bar.
server.languageManager.setLanguage("en")

// This is the body that will be use by all other panel.
var bodyElement = new Element(document.getElementsByTagName("body")[0], {"tag":"div", "style":"width: 100%; height: 100%;"})

var mainPage = null

/**
 * This function is the entry point of the application...
 */
function main() {

    // get the prototypes of the blog schema.
    server.entityManager.getEntityPrototypes("sql_info",
    // success callback
    function(results, caller){
        mainPage = new MainPage(bodyElement)
    },
    // error callback.
    function(){

    }, 
    // caller.
    {} )
}