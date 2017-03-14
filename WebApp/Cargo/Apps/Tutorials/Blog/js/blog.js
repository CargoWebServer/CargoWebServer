
/**
 * Utility class use to manage data element of the blog.
 */
var BlogManager = function (parent) {
    var databaseName = "Blog"
    var schemaId = "dbo"

    // The language Info JSON contain a mapping of id and text.
    // The id's here are the same as the html element's in the entity panel,
    // so we can control the panel title and field's labal text.
    var languageInfo = {
        "en": {
            // The user input panel
            "Blog.dbo.blog_user": "User's",
            "Blog.blog_user": "User's",
            "FK_blog_comment_blog_user": "Comments",
            "fk_comment_2":"Comments",
            "Name": "Name",
            "email": "Email",
            "website": "Website"
        },
        "fr": {
            // The user input panel
            "Blog.dbo.blog_user": "Utilisateurs",
            "Blog.blog_user": "Utilisateurs",
            "FK_blog_comment_blog_user": "Commentaires",
            "fk_comment_2": "Commentaires",
            "Name": "Nom",
            "email": "Email",
            "website": "Site web"
        }
    }

    // Depending of the language the correct text will be set.
    server.languageManager.appendLanguageInfo(languageInfo)

    // The main panel of the blog manager.
    this.panel = parent.appendElement({ "tag": "div", "class": "blog-manager" }).down()

    // At the configuration time server create all needed entities from SQL schema,
    // here Blog was the database we made use of and dbo are the sql schema and finaly
    // the blog_user was the table. Cargo offer you the EntityPanel class to 
    // manipulate entity of every type, just tell Cargo the type name you want and boom!
    var userTypeName = databaseName
    //userTypeName += "." + schemaId
    userTypeName += ".blog_user"

    var userPanel = new EntityPanel(this.panel, userTypeName, function (panel) {
        panel.maximizeBtn.element.click()
    }, null, false, null, null)

    return this
}
