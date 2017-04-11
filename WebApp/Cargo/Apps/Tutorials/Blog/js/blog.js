
/**
 * Utility class use to manage data element of the blog.
 */
var BlogManager = function (parent) {

    // The language Info JSON contain a mapping of id and text.
    // The id's here are the same as the html element's in the entity panel,
    // so we can control the panel title and field's labal text.
    var languageInfo = {
        "en": {
            // The navigation bar.
            "toggle-navigation": "Toggle navigation",
            "start-lnk": "Start Boostrap",
            "about-lnk": "about",
            "create-lnk": "New blog",
            "contact-lnk": "Contact",
            "login-lnk": "login",
            "login-header": "Log In",
            "login-form-username-lbl":"Username",
            "login-form-password-lbl":"Password",
            "remember-login-form-lbl":" Remember Me",
            "forgot-password-id" :"Forgot Password?",
            "login-submit":"Log In",
            "register-lnk": "register",
            "blog-search-title": "Blog Search",
            "blog-categories": "Blog Categories",
            "side-well-widget": "Lorem ipsum dolor sit amet, consectetur adipisicing elit. Inventore, perspiciatis adipisci accusamus laudantium odit aliquam repellat tempore quos aspernatur vero.",
            "footer-text": "Copyright &copy; Your Website 2017",

            // The user input panel
            "Blog.dbo.blog_user": "User's",
            "Blog.blog_user": "User's",
            "FK_blog_comment_blog_user": "Comments",
            "fk_comment_2": "Comments",
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

    // Local variable.
    var databaseName = "Blog"
    var schemaId = "" //"dbo"

    // At the configuration time server create all needed entities from SQL schema,
    // here Blog was the database we made use of and dbo are the sql schema and finaly
    // the blog_user was the table. Cargo offer you the EntityPanel class to 
    // manipulate entity of every type, just tell Cargo the type name you want and boom!
    var userTypeName = databaseName
    if (schemaId.length > 0) {
        userTypeName += "." + schemaId
    }
    userTypeName += ".blog_user"

    /*
        var userPanel = new EntityPanel(this.panel, userTypeName, function (panel) {
            panel.maximizeBtn.element.click()
        }, null, false, null, null)
    */

    // The active blog view.
    this.activeBlogView = null

    /////////////////////////////////////////////////////////////////////////////////
    // Interface.
    /////////////////////////////////////////////////////////////////////////////////

    // Navigation bar
    this.navBar = parent.appendElement({ "tag": "nav", "class": "navbar navbar-inverse navbar-fixed-top", "role": "navigation" }).down()

    // Brand and toggle get grouped for better mobile display
    this.navBar.appendElement({ "tag": "div", "class": "container", "id": "nav-container" }).down()
        .appendElement({ "tag": "div", class: "navbar-header" }).down()
        .appendElement({ "tag": "button", "type": "button", "class": "navbar-toggle", "data-toggle": "collapse", "data-target": "#bs-example-navbar-collapse-1" }).down()
        .appendElement({ "tag": "span", "class": "sr-only", "id": "toggle-navigation" }).down()
        .appendElement({ "tag": "span", "class": "icon-bar" })
        .appendElement({ "tag": "span", "class": "icon-bar" })
        .appendElement({ "tag": "span", "class": "icon-bar" }).up()
        .appendElement({ "tag": "a", "class": "navbar-brand", "id": "start-lnk", "href": "#" })

    // Collect the nav links, forms, and other content for toggling
    this.navBar.getChildById("nav-container").appendElement({ "tag": "div", "class": "collapse navbar-collapse", "id": "bs-example-navbar-collapse-1" }).down()
        .appendElement({ "tag": "ul", "class": "nav navbar-nav" }).down()
        .appendElement({ "tag": "li" }).down()
        .appendElement({ "tag": "a", "id": "create-lnk", "href": "#" }).up()
        .appendElement({ "tag": "li" }).down()
        .appendElement({ "tag": "a", "id": "about-lnk", "href": "#" }).up().up()
        .appendElement({ "tag": "ul", "class": "nav navbar-nav navbar-right" }).down()
        .appendElement({ "tag": "li" }).down()
        // The register link.
        .appendElement({ "tag": "a", "id": "register-lnk", "href": "#" }).up()
        .appendElement({ "tag": "li", "class": "dropdown open" }).down()
        // Now the login dialog...
        .appendElement({ "tag": "a", "id": "login-lnk", "href": "#" })
        .appendElement({"tag":"ul", "class":"dropdown-menu dropdown-lr animated slideInRight", "role":"menu", "style":"min-width: 300px;"}).down()
        .appendElement({ "tag": "div", "class": "col-lg-12" }).down()
        .appendElement({ "tag": "div", "class": "text-center" }).down()
        .appendElement({ "tag": "h3", "id": "login-header" }).up()
        .appendElement({ "tag": "form"}).down()
        // The user name
        .appendElement({"tag":"div", "class":"form-group"}).down()
        .appendElement({"tag":"label", "for":"username", "id":"login-form-username-lbl"})
        .appendElement({"tag":"input", "name":"username", "id":"login-form-username-input", "tabindex":"1", "class":"form-control", "placeholder":"Username", "autocomplete":"off"}).up()
        // The password
        .appendElement({"tag":"div", "class":"form-group"}).down()
        .appendElement({"tag":"label", "for":"password", "id":"login-form-password-lbl"})
        .appendElement({"tag":"input", "name":"password", "id":"login-form-password-input", "tabindex":"2", "type":"password", "class":"form-control", "placeholder":"Password", "autocomplete":"off"}).up()
        
        // The remember me and login button.
        .appendElement({"tag":"div", "class":"form-group"}).down()
        .appendElement({"tag":"div", "class":"row"}).down()
        .appendElement({"tag":"div", "class":"col-xs-7"}).down()
        .appendElement({"tag":"input", "type":"checkbox", "name":"remember", "id":"remember", "tabindex":"3" })
        .appendElement({"tag":"label", "for":"remember", "id":"remember-login-form-lbl", "style":"padding-left: 5px;"}).up()
        .appendElement({"tag":"div", "class":"col-xs-5 pull-right"}).down()
        .appendElement({"tag":"button", "name":"login-submit", "id":"login-submit", "class":"form-control btn btn-success", "tabindex":"4"}).up().up().up()
        .appendElement({"tag":"div", "class":"form-group"}).down()
        .appendElement({"tag":"div", "class":"row"}).down()
        .appendElement({"tag":"div", "class":"col-lg-12"}).down()
        .appendElement({ "tag": "div", "class": "text-center" }).down()
        .appendElement({"tag":"a", "class":"forgot-password", "id":"forgot-password-id", "tabindex":"5"})


    // The content
    this.container = parent.appendElement({ "tag": "div", "class": "container" }).down()
        .appendElement({ "tag": "div", "class": "row" }).down()

    // Blog Post Content Column
    this.blogContainer = this.container.appendElement({ "tag": "div", "class": "col-lg-8" }).down()

    // Blog Sidebar Widgets Column
    // The search panel.
    this.container.appendElement({ "tag": "div", "class": "col-md-4" }).down()
        .appendElement({ "tag": "div", "class": "well" }).down()
        .appendElement({ "tag": "span", "class": "input-group" }).down()
        .appendElement({ "tag": "input", "type": "text", "class": "form-control" })
        .appendElement({ "tag": "span", "class": "input-group-btn" }).down()
        .appendElement({ "tag": "button", class: "btn btn-default" }).down()
        .appendElement({ "tag": "span", "class": "fa fa-search" }).up().up().up().up()

        // Blog Categories Well // TODO populate it with the db content.
        .appendElement({ "tag": "div", "class": "well" }).down()
        .appendElement({ "tag": "h4", "id": "blog-categories" })
        .appendElement({ "tag": "div", "class": "row" }).down()
        .appendElement({ "tag": "div", "class": "col-lg-6" }).down()
        .appendElement({ "tag": "ul", "class": "list-unstyled" }).down()
        .appendElement({ "tag": "li" }).down()
        .appendElement({ "tag": "a", innerHtml: "Category Name" }).up()
        .appendElement({ "tag": "li" }).down()
        .appendElement({ "tag": "a", innerHtml: "Category Name" }).up()
        .appendElement({ "tag": "li" }).down()
        .appendElement({ "tag": "a", innerHtml: "Category Name" }).up()
        .appendElement({ "tag": "li" }).down()
        .appendElement({ "tag": "a", innerHtml: "Category Name" }).up().up().up()
        .appendElement({ "tag": "div", "class": "col-lg-6" }).down()
        .appendElement({ "tag": "ul", "class": "list-unstyled" }).down()
        .appendElement({ "tag": "li" }).down()
        .appendElement({ "tag": "a", innerHtml: "Category Name" }).up()
        .appendElement({ "tag": "li" }).down()
        .appendElement({ "tag": "a", innerHtml: "Category Name" }).up()
        .appendElement({ "tag": "li" }).down()
        .appendElement({ "tag": "a", innerHtml: "Category Name" }).up()
        .appendElement({ "tag": "li" }).down()
        .appendElement({ "tag": "a", innerHtml: "Category Name" }).up().up().up().up().up()

        // The side Widget Well
        .appendElement({ "tag": "div", "class": "well" }).down()
        .appendElement({ "tag": "h4" })
        .appendElement({ "tag": "p", "id": "side-well-widget" }).up().appendElement({ "tag": "hr" }).up()

        // The footer.
        .appendElement({ "tag": "footer" }).down()
        .appendElement({ "tag": "div", "class": "row" }).down()
        .appendElement({ "tag": "div", "class": "col-lg-12" }).down()
        .appendElement({ "tag": "p", "id": "footer-text" })

    //////////////////////////////////////////////////////////////////////
    // Blog manager action.
    //////////////////////////////////////////////////////////////////////
    // login
    this.loginLnk = this.navBar.getChildById("login-lnk")
    this.loginLnk.element.onclick = function () {
        alert("Aille!!!!")

    }

    // register
    this.registerLnk = this.navBar.getChildById("register-lnk")
    this.registerLnk.element.onclick = function () {
        alert("Hooo!!!!")
    }

    // Create a new blog.
    this.createBlog = this.navBar.getChildById("create-lnk")

    this.createBlog.element.onclick = function (blogManager) {
        return function () {
            // Create a new blog.
            blogManager.createNewBlog()
        }
    } (this)

    return this
}


/**
 * Create a new blog.
 */
BlogManager.prototype.createNewBlog = function () {
    // So here I will create a new blog post and from it I will set it view.
    // TODO Create Author's from oauth2
    // Create a new Blog.
    this.activeBlogView = new BlogPostView(this.blogContainer, 1)

    // Set the blog view editable.
    this.setEditable(this.activeBlogView)
}

/**
 * Inject functionality to edit parts of the blog.
 */
BlogManager.prototype.setEditable = function (blogView) {

    function setEditable(div, onClickCallback) {
        div.element.style.position = "relative"
        div.appendElement({ "tag": "div", "class": "edit-button", "title": "edit" }).down()
            .appendElement({ "tag": "i", "class": "fa fa-pencil-square" })

        // The action to do on click.
        div.element.onclick = function (onClickCallback, div) {
            return function () {
                onClickCallback(div, this)
            }
        } (onClickCallback, div)
    }

    // Set the title div.
    var setTitleCallback = function (div, btn) {
        // Here I will use a simple input...
        if (div.element.firstChild.id != "blog-post-title") {
            // I will get the actual title value.
            var title = div.element.innerText
            div.element.innerText = ""
            var inputTitle = new Element(div.element, { "tag": "input", "id": "blog-post-title", "value": title })
            inputTitle.element.select()

            // Now the save button.
            var saveBtn = div.appendElement({ "tag": "div", "class": "edit-button" }).down()
                .appendElement({ "tag": "i", "class": "fa fa fa-floppy-o" })

            // Now the save action.
            saveBtn.element.onclick = function (inputTitle, div) {
                return function (evt) {
                    // Stop event propagation so we will not return direclty here...
                    evt.stopPropagation()
                    // Here I will set back the text inside the h1 element.
                    div.element.innerText = inputTitle.element.value
                    setEditable(div, setTitleCallback)
                }
            } (inputTitle, div)
        }
    }
    // Set the title.
    setEditable(blogView.titleDiv, setTitleCallback)

    // The author will be the logged user.

    // Set the content div
    setEditable(blogView.pageContentDiv,
        function (div) {
            // I made use of http://summernote.org/ Great work folk's
            $('#page-content').summernote();

            // I will append the save button to existing toolbar, it's so easy...
            var saveBtn = new Element(document.getElementsByClassName("note-view")[0], { "tag": "button", "tabindex": -1, "type": "button", "class": "note-btn btn btn-default btn-sm btn btn-primary btn-save", "title": "save", "data-original-title": "save" })
            saveBtn.appendElement({ "tag": "i", "class": "fa fa-floppy-o" })

            // Now the save action.
            saveBtn.element.onclick = function () {
                // remove the editor.
                $('#page-content').summernote('destroy');
                blogManager.saveActiveBlog()
            }
        })

    // Here I will use use the clock to set the time.
}

/**
 * Save the active blog.
 */
BlogManager.prototype.saveActiveBlog = function () {
    console.log("save blog!!!!")
}

/**
 * There is the blog the blog view.
 */
var BlogPostView = function (parent, id) {

    // The blog text interface elements
    var languageInfo = {
        "en": {
            "blog-post-title": "Blog Post Title",
            "written-by": "by ",
            "author-name": "Start Bootstrap",
            "created-date": " Posted on August 24, 2013 at 9:00 PM",
            "page-content": "<p class=\"lead\">Lorem ipsum dolor sit amet, consectetur adipisicing elit. Ducimus, vero, obcaecati, aut, error quam sapiente nemo saepe quibusdam sit excepturi nam quia corporis eligendi eos magni recusandae laborum minus inventore?</p>"
            + "<p>Lorem ipsum dolor sit amet, consectetur adipisicing elit. Ut, tenetur natus doloremque laborum quos iste ipsum rerum obcaecati impedit odit illo dolorum ab tempora nihil dicta earum fugiat. Temporibus, voluptatibus.</p>"
            + " <p>Lorem ipsum dolor sit amet, consectetur adipisicing elit. Eos, doloribus, dolorem iusto blanditiis unde eius illum consequuntur neque dicta incidunt ullam ea hic porro optio ratione repellat perspiciatis. Enim, iure!</p>"
            + "<p>Lorem ipsum dolor sit amet, consectetur adipisicing elit. Error, nostrum, aliquid, animi, ut quas placeat totam sunt tempora commodi nihil ullam alias modi dicta saepe minima ab quo voluptatem obcaecati?</p>"
            + "<p>Lorem ipsum dolor sit amet, consectetur adipisicing elit. Harum, dolor quis. Sunt, ut, explicabo, aliquam tenetur ratione tempore quidem voluptates cupiditate voluptas illo saepe quaerat numquam recusandae? Qui, necessitatibus, est!</p>",
            "comment-header": "Leave a Comment",
            "submit-comment": "Submit"

        },
        "fr": {

        }
    }

    // Depending of the language the correct text will be set.
    server.languageManager.appendLanguageInfo(languageInfo)

    // Page Content.
    this.pageContainer = parent.appendElement({ "tag": "div", "class": "container" }).down()
        .appendElement({ "tag": "div", "class": "row" }).down()

    // Blog Post Content Column
    this.pageContainer.appendElement({ "tag": "div", "class": "col-lg-8" }).down()
        .appendElement({ "tag": "h1", "id": "blog-post-title" })
        // The author
        .appendElement({ "tag": "p", "class": "lead" }).down().appendElement({ "tag": "span", "id": "written-by" })
        .appendElement({ "tag": "a", "id": "author-name", "href": "#" }).up()
        .appendElement({ "tag": "hr" })
        // The date and time 
        .appendElement({ "tag": "p" }).down().appendElement({ "tag": "span", "class": "fa fa-clock-o" }).appendElement({ "tag": "span", "id": "created-date" }).up()
        .appendElement({ "tag": "hr" })
        // The page content.
        .appendElement({ "tag": "div", "id": "page-content" })
        .appendElement({ "tag": "hr" })

        // The blog comments input.
        .appendElement({ "tag": "div", "class": "well" }).down()
        .appendElement({ "tag": "h4", id: "comment-header" })
        .appendElement({ "tag": "div", "class": "form-group" }).down()
        .appendElement({ "tag": "textarea", "id": "comment-text-area", "class": "form-control", "rows": "3" }).up()
        .appendElement({ "tag": "button", "type": "submit", "class": "btn btn-primary", "id": "submit-comment" }).up().appendElement({ "tag": "hr" })
        .appendElement({ "tag": "div", "id": "comments-container" })

    // The comment container.
    this.pageContainer.commentContainer = this.pageContainer.getChildById("comments-container")
    this.pageContainer.commentContainer.up().appendElement({ "tag": "hr" })

    ////////////////////////////////////////////////////////////////////////
    // Blog sections.
    ////////////////////////////////////////////////////////////////////////

    // The title div
    this.titleDiv = this.pageContainer.getChildById("blog-post-title")

    // The author div
    this.authorDiv = this.pageContainer.getChildById("author-name")

    // The created div.
    this.createdDiv = this.pageContainer.getChildById("created-date")

    // The page content.
    this.pageContentDiv = this.pageContainer.getChildById("page-content")

    // The summit button.
    this.submitCommentBtn = this.pageContainer.getChildById("submit-comment")

    ////////////////////////////////////////////////////////////////////////
    // Blog actions.
    ////////////////////////////////////////////////////////////////////////

    this.submitCommentBtn.element.onclick = function (pageContainer) {
        return function () {
            new BlogPostCommentView(pageContainer.commentContainer, "toto", pageContainer.getChildById("comment-text-area").element.value, "121212")
            pageContainer.getChildById("comment-text-area").element.value = ""
        }
    } (this.pageContainer)

    return this
}

/**
 * Append a blog comment.
 */
BlogPostView.prototype.appendComment = function (user, comment, date) {
    // Create the comment
    new BlogPostCommentView(this.pageContainer, user, comment, date)
}

/**
 * The interface of a blog comment.
 */
BlogPostCommentView = function (parent, user, comment, date) {

    // Here I will append the new comment.
    this.commentContainer = parent.prependElement({ "tag": "div", class: "media" }).down()
    this.commentContainer.appendElement({ "tag": "a", "class": "pull-left", "href": "#" }).down()
        .appendElement({ "tag": "img", "class": "media-object", "src": "http://placehold.it/64x64" }).up()
        .appendElement({ "tag": "div", "class": "media-body" }).down()
        .appendElement({ "tag": "h4", "innerHtml": user }).down().appendElement({ "tag": "small", innerHtml: date }).up().up()
        .appendElement({ "tag": "div", "innerHtml": comment })


    return this
}