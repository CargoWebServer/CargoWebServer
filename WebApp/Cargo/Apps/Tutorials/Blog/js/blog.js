
// global variable.
var databaseName = "Blog."
var schemaId = "" //"dbo."

var userTypeName = databaseName + schemaId + "blog_user"
var authorTypeName = databaseName + schemaId + "blog_author"
var blogPostTypeName = databaseName + schemaId + "blog_post"
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
            // Login text
            "login-lnk": "login",
            "login-header": "Log In",
            "login-form-username-lbl": "Username",
            "login-form-password-lbl": "Password",
            "forgot-password-id": "Forgot Password?",
            "login-submit": "Log In",
            "logout-lnk": "logout",
            "remember-login-form-lbl": " Remember Me",
            // Account text (user info)
            "user-info-header": "Personnal Informations",
            "user-info-lnk": "myUserName",
            "user-info-first-name-label": "First Name",
            "user-info-middle-label": "Middle.",
            "user-info-last-name-label": "Last Name",
            "user-info-email-label": "E-Mail",
            "user-info-phone-label": "Phone",
            "user-info-submit": "Save",
            // Register text.
            "register-lnk": "register",
            "register-header": "Register",
            "register-submit": "Register Now",
            // error text...
            "register-form-confirm-password-input": "Whoops, these don't match",
            "register-form-emailaddress-input": "That email address is invalid",
            "register-username-input-error": "",
            // Other parts of header.
            "blog-search-title": "Blog Search",
            "blog-categories": "Blog Categories",
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

    // The active blog view.
    this.activePostView = null

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
        .appendElement({ "tag": "li", "id": "new-blog-lnk", "style": "display: none;" }).down()
        .appendElement({ "tag": "a", "id": "create-lnk", "href": "#" }).up()
        .appendElement({ "tag": "li" }).down()
        .appendElement({ "tag": "a", "id": "about-lnk", "href": "#" }).up().up()
        .appendElement({ "tag": "ul", "class": "nav navbar-nav navbar-right" }).down()

        // The register link.
        .appendElement({ "tag": "li", "class": "dropdown" }).down()
        .appendElement({ "tag": "a", "id": "register-lnk", "href": "#" })
        .appendElement({ "tag": "ul", "id": "register-dropdown", "class": "dropdown-menu dropdown-lr animated slideInRight", "role": "menu", "style": "min-width: 300px;" }).down()
        .appendElement({ "tag": "div", "class": "col-lg-12" }).down()
        .appendElement({ "tag": "div", "class": "text-center" }).down()
        .appendElement({ "tag": "h3", "id": "register-header" }).up()
        .appendElement({ "tag": "div", "role": "form", "id": "register-form" }).down()
        // The user name
        .appendElement({ "tag": "div", "class": "form-group" }).down()
        .appendElement({ "tag": "input", "name": "username", "type": "text", "pattern": "^[_A-z0-9]{1,}$", "maxlength": "15", "id": "register-form-username-input", "tabindex": "1", "class": "form-control", "placeholder": "Username", "autocomplete": "off", "required": "" })
        .appendElement({ "tag": "div", "class": "help-block with-errors", "id": "register-username-input-error" })
        .up()
        // The email
        .appendElement({ "tag": "div", "class": "form-group" }).down()
        .appendElement({ "tag": "input", "name": "emailAddress", "type": "email", "id": "register-form-emailaddress-input", "tabindex": "2", "class": "form-control", "placeholder": "Email Address", "autocomplete": "off", "data-error": "That email address is invalid" })
        .appendElement({ "tag": "div", "class": "help-block with-errors" })
        .up()
        // The password
        .appendElement({ "tag": "div", "class": "form-group" }).down()
        .appendElement({ "tag": "input", "name": "password", "type": "password", "id": "register-form-password-input", "tabindex": "3", "class": "form-control", "placeholder": "Password", "autocomplete": "off" }).up()
        // The confirm password
        .appendElement({ "tag": "div", "class": "form-group" }).down()
        .appendElement({ "tag": "input", "name": "confirmPassword", "type": "password", "id": "register-form-confirm-password-input", "tabindex": "4", "class": "form-control", "placeholder": "Confirm Password", "autocomplete": "off", "data-match": "#register-form-password-input", "data-match-error": "Whoops, these don't match" })
        .appendElement({ "tag": "div", "class": "help-block with-errors" })
        .up()
        // The register button.
        .appendElement({ "tag": "div", "class": "form-group" }).down()
        .appendElement({ "tag": "div", "class": "row" }).down()
        .appendElement({ "tag": "div", "class": "col-lg-12" }).down()
        .appendElement({ "tag": "div", "class": "col-xs-6 col-xs-offset-3" }).down()
        .appendElement({ "tag": "button", "name": "register-submit", "id": "register-submit", "class": "form-control btn btn-info", "tabindex": "5" })
        // Move up
        .up().up().up().up().up().up().up().up()
        // If the user is logged here I will display it information.
        .appendElement({ "tag": "li", "id": "user-info-dropdown-lnk", "class": "dropdown", "style": "display: none;" }).down()
        .appendElement({ "tag": "a", "id": "user-info-lnk", "href": "#" })
        .appendElement({ "tag": "ul", "id": "user-info-dropdown", "class": "dropdown-menu dropdown-lr animated slideInRight", "role": "menu", "style": "min-width: 500px;" }).down()
        .appendElement({ "tag": "div", "class": "col-lg-12" }).down()
        .appendElement({ "tag": "div", "class": "text-center" }).down()
        .appendElement({ "tag": "h3", "id": "user-info-header" }).up()
        // The form.
        .appendElement({ "tag": "div", "id": "user-info-form" }).down()
        // The First name.
        .appendElement({ "tag": "div", "class": "form-group has-feedback" }).down()
        .appendElement({ "tag": "div", "class": "row" }).down()
        .appendElement({ "tag": "label", "for": "first_name", "id": "user-info-first-name-label", "class": "col-md-4 control-label" })
        .appendElement({ "tag": "div", "class": "col-md-8 inputGroupContainer" }).down()
        .appendElement({ "tag": "div", "class": "input-group" }).down()
        .appendElement({ "tag": "span", "class": "input-group-addon" }).down()
        .appendElement({ "tag": "i", "class": "fa fa-user" }).up()
        .appendElement({ "tag": "input", "id": "first_name", "placeholder": "First Name", "pattern": "^[a-zA-ZàáâäãåąčćęèéêëėįìíîïłńòóôöõøùúûüųūÿýżźñçčšžÀÁÂÄÃÅĄĆČĖĘÈÉÊËÌÍÎÏĮŁŃÒÓÔÖÕØÙÚÛÜŲŪŸÝŻŹÑßÇŒÆČŠŽ∂ð ,.'-]+$", "class": "form-control", "type": "text", "tabindex": "1", "required": "" })
        .up().up().up()
        .appendElement({ "tag": "div", "class": "help-block with-errors", "style": "text-align: center;" })
        .up()
        // The middle name.
        .appendElement({ "tag": "div", "class": "form-group has-feedback" }).down()
        .appendElement({ "tag": "div", "class": "row" }).down()
        .appendElement({ "tag": "label", "for": "middle_name", "id": "user-info-middle-label", "class": "col-md-4 control-label" })
        .appendElement({ "tag": "div", "class": "col-md-8 inputGroupContainer" }).down()
        .appendElement({ "tag": "div", "class": "input-group" }).down()
        .appendElement({ "tag": "span", "class": "input-group-addon" }).down()
        .appendElement({ "tag": "i", "class": "fa fa-user" }).up()
        .appendElement({ "tag": "input", "id": "middle_name", "placeholder": "Middle Name", "class": "form-control", "type": "text", "data-bv-field": "middle_name", "tabindex": "2" })
        .up().up().up()
        .appendElement({ "tag": "div", "class": "help-block with-errors", "style": "text-align: center;" })
        .up()
        // The last name.
        .appendElement({ "tag": "div", "class": "form-group has-feedback" }).down()
        .appendElement({ "tag": "div", "class": "row" }).down()
        .appendElement({ "tag": "label", "for": "last_name", "id": "user-info-last-name-label", "class": "col-md-4 control-label" })
        .appendElement({ "tag": "div", "class": "col-md-8 inputGroupContainer" }).down()
        .appendElement({ "tag": "div", "class": "input-group" }).down()
        .appendElement({ "tag": "span", "class": "input-group-addon" }).down()
        .appendElement({ "tag": "i", "class": "fa fa-user" }).up()
        .appendElement({ "tag": "input", "id": "last_name", "placeholder": "Last Name", "class": "form-control", "type": "text", "tabindex": "3", "pattern": "^[a-zA-ZàáâäãåąčćęèéêëėįìíîïłńòóôöõøùúûüųūÿýżźñçčšžÀÁÂÄÃÅĄĆČĖĘÈÉÊËÌÍÎÏĮŁŃÒÓÔÖÕØÙÚÛÜŲŪŸÝŻŹÑßÇŒÆČŠŽ∂ð ,.'-]+$", "required": "" })
        .up().up().up()
        .appendElement({ "tag": "div", "class": "help-block with-errors", "style": "text-align: center;" })
        .up()
        // The email address.
        .appendElement({ "tag": "div", "class": "form-group has-feedback" }).down()
        .appendElement({ "tag": "div", "class": "row" }).down()
        .appendElement({ "tag": "label", "for": "email", "id": "user-info-email-label", "class": "col-md-4 control-label" })
        .appendElement({ "tag": "div", "class": "col-md-8 inputGroupContainer" }).down()
        .appendElement({ "tag": "div", "class": "input-group" }).down()
        .appendElement({ "tag": "span", "class": "input-group-addon" }).down()
        .appendElement({ "tag": "i", "class": "fa fa-envelope" }).up()
        .appendElement({ "tag": "input", "id": "email", "placeholder": "E-Mail Address", "class": "form-control", "type": "email", "tabindex": "4", "data-error": "That email address is invalid" })
        .up().up().up()
        .appendElement({ "tag": "div", "class": "help-block with-errors", "style": "text-align: center;" })
        .up()
        // The phone address.
        .appendElement({ "tag": "div", "class": "form-group has-feedback" }).down()
        .appendElement({ "tag": "div", "class": "row" }).down()
        .appendElement({ "tag": "label", "for": "phone", "id": "user-info-phone-label", "class": "col-md-4 control-label" })
        .appendElement({ "tag": "div", "class": "col-md-8 inputGroupContainer" }).down()
        .appendElement({ "tag": "div", "class": "input-group" }).down()
        .appendElement({ "tag": "span", "class": "input-group-addon" }).down()
        .appendElement({ "tag": "i", "class": "fa fa-phone" }).up()
        .appendElement({ "tag": "input", "id": "phone", "type": "tel", "pattern": "^[\+]?[(]?[0-9]{3}[)]?[-\s\.]?[0-9]{3}[-\s\.]?[0-9]{4,6}$", "title": "Phone Number (Format: +99(99)9999-9999)", "placeholder": "(845)555-1212", "class": "form-control", "data-bv-field": "phone", "tabindex": "5" })
        .up().up().up()
        .appendElement({ "tag": "div", "class": "help-block with-errors", "style": "text-align: center;" })
        .up()
        // The save button.
        .appendElement({ "tag": "div", "class": "form-group" }).down()
        .appendElement({ "tag": "div", "class": "row" }).down()
        .appendElement({ "tag": "div", "class": "col-lg-12" }).down()
        .appendElement({ "tag": "div", "class": "col-xs-6 col-xs-offset-3" }).down()
        .appendElement({ "tag": "button", "id": "user-info-submit", "class": "form-control btn btn-info", "tabindex": "6" }).up().up().up().up()
        .up().up().up().up()

        // Now the login dialog...
        .appendElement({ "tag": "li", "id": "login-dropdown-lnk", "class": "dropdown" }).down()
        .appendElement({ "tag": "a", "id": "login-lnk", "href": "#" })
        .appendElement({ "tag": "ul", "id": "login-dropdown", "class": "dropdown-menu dropdown-lr animated slideInRight", "role": "menu", "style": "min-width: 300px;" }).down()
        .appendElement({ "tag": "div", "class": "col-lg-12" }).down()
        .appendElement({ "tag": "div", "class": "text-center" }).down()
        .appendElement({ "tag": "h3", "id": "login-header" }).up()
        .appendElement({ "tag": "div" }).down()
        // The user name
        .appendElement({ "tag": "div", "class": "form-group" }).down()
        .appendElement({ "tag": "label", "for": "username", "id": "login-form-username-lbl" })
        .appendElement({ "tag": "input", "name": "username", "id": "login-form-username-input", "tabindex": "1", "class": "form-control", "placeholder": "Username", "autocomplete": "off" })
        .appendElement({ "tag": "div", "id": "login-form-username-error", "class": "help-block with-errors", "style": "text-align: center;" })
        .up()
        // The password
        .appendElement({ "tag": "div", "class": "form-group" }).down()
        .appendElement({ "tag": "label", "for": "password", "id": "login-form-password-lbl" })
        .appendElement({ "tag": "input", "name": "password", "id": "login-form-password-input", "tabindex": "2", "type": "password", "class": "form-control", "placeholder": "Password", "autocomplete": "off" })
        .appendElement({ "tag": "div", "id": "login-form-password-error", "class": "help-block with-errors", "style": "text-align: center;" })
        .up()
        // The remember me and login button.
        .appendElement({ "tag": "div", "class": "form-group" }).down()
        .appendElement({ "tag": "div", "class": "row" }).down()
        .appendElement({ "tag": "div", "class": "col-xs-7" }).down()
        .appendElement({ "tag": "input", "type": "checkbox", "name": "remember", "id": "remember", "tabindex": "3" })
        .appendElement({ "tag": "label", "for": "remember", "id": "remember-login-form-lbl", "style": "padding-left: 5px;" }).up()
        .appendElement({ "tag": "div", "class": "col-xs-5 pull-right" }).down()
        .appendElement({ "tag": "button", "name": "login-submit", "id": "login-submit", "class": "form-control btn btn-success", "tabindex": "4" }).up().up().up()
        .appendElement({ "tag": "div", "class": "form-group" }).down()
        .appendElement({ "tag": "div", "class": "row" }).down()
        .appendElement({ "tag": "div", "class": "col-lg-12" }).down()
        .appendElement({ "tag": "div", "class": "text-center" }).down()
        .appendElement({ "tag": "a", "class": "forgot-password", "id": "forgot-password-id", "tabindex": "5" }).up().up().up().up()
        .up().up().up().up()
        // The logout link.
        .appendElement({ "tag": "li", "id": "logout-dropdown-lnk", "style": "display: none;" }).down()
        .appendElement({ "tag": "a", "id": "logout-lnk", "href": "#" })

    // Set boostrap validation.
    $('#register-form').validator()
    $('#user-info-form').validator()

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

        // The side well widget.
        .appendElement({ "tag": "div", "class": "well" }).down()
        .appendElement({ "tag": "h4" })
        .appendElement({ "tag": "p", "id": "side-well-widget" }).up().appendElement({ "tag": "hr" }).up()

        // The footer.
        .appendElement({ "tag": "footer" }).down()
        .appendElement({ "tag": "div", "class": "row" }).down()
        .appendElement({ "tag": "div", "class": "col-lg-12" }).down()
        .appendElement({ "tag": "p", "id": "footer-text" })

    // The side well widget is use to display additionall information like the list of blog for a user.
    this.sideWellWidget = this.container.getChildById("side-well-widget")

    //////////////////////////////////////////////////////////////////////
    // Blog manager action.
    //////////////////////////////////////////////////////////////////////
    // login
    this.loginLnk = this.navBar.getChildById("login-lnk")
    this.registerLnk = this.navBar.getChildById("register-lnk")
    this.userInfoLnk = this.navBar.getChildById("user-info-lnk")

    this.loginLnk.element.onclick = function (registerLnk, userInfoLnk) {
        return function () {
            registerLnk.element.parentNode.className = "dropdown"
            userInfoLnk.element.parentNode.className = "dropdown"
            if (this.parentNode.className.indexOf("open") > 0) {
                this.parentNode.className = "dropdown"
            } else {
                this.parentNode.className = "dropdown open"
            }
            document.getElementById("login-form-username-input").focus()
        }
    } (this.registerLnk, this.userInfoLnk)

    // register
    this.registerLnk.element.onclick = function (loginLnk, userInfoLnk) {
        return function () {
            loginLnk.element.parentNode.className = "dropdown"
            userInfoLnk.element.parentNode.className = "dropdown"
            if (this.parentNode.className.indexOf("open") > 0) {
                this.parentNode.className = "dropdown"
            } else {
                this.parentNode.className = "dropdown open"
            }
            document.getElementById("register-form-username-input").focus()
        }
    } (this.loginLnk, this.userInfoLnk)

    // The user informations.
    this.userInfoLnk.element.onclick = function (loginLnk, registerLnk) {
        return function () {
            loginLnk.element.parentNode.className = "dropdown"
            registerLnk.element.parentNode.className = "dropdown"
            if (this.parentNode.className.indexOf("open") > 0) {
                this.parentNode.className = "dropdown"
            } else {
                this.parentNode.className = "dropdown open"
            }
        }
    } (this.loginLnk, this.registerLnk)

    // Now the mouse out event.
    this.registerDropDown = this.navBar.getChildById("register-dropdown")

    // The login dropdown.
    this.loginDropDown = this.navBar.getChildById("login-dropdown")

    // The user info dropdown
    this.userInfoDropDown = this.navBar.getChildById("user-info-dropdown")

    // Hide the window if the mouse leave it...
    this.userInfoDropDown.element.onmouseleave = this.registerDropDown.element.onmouseleave = this.loginDropDown.element.onmouseleave = function () {
        this.parentNode.className = "dropdown"
    }

    // Create a new blog.
    this.createBlog = this.navBar.getChildById("create-lnk")

    this.createBlog.element.onclick = function (blogManager) {
        return function () {
            // Create a new blog.
            blogManager.createNewPost()
        }
    } (this)

    // Register.
    this.registerBtn = this.navBar.getChildById("register-submit")

    // First i will get the values...
    var userNameInput = this.registerDropDown.getChildById("register-form-username-input")
    var emailInput = this.registerDropDown.getChildById("register-form-emailaddress-input")
    var passwordInput = this.registerDropDown.getChildById("register-form-password-input")
    var confirmPasswordInput = this.registerDropDown.getChildById("register-form-confirm-password-input")

    // Register a new account.
    this.registerBtn.element.onclick = function (userNameInput, emailInput, passwordInput, confirmPasswordInput) {
        return function () {
            // TODO make validation with a predefined bootstrap validator lib...
            // Get the variable content.
            var userName = userNameInput.element.value
            userNameInput.element.value = ""

            var email = emailInput.element.value
            emailInput.element.value = ""

            var password = passwordInput.element.value
            passwordInput.element.value = ""

            var confirmPassword = confirmPasswordInput.element.value
            confirmPasswordInput.element.value = ""

            // Now i will create a new account.
            server.accountManager.register(userName, password, email,
                function (result, caller) {
                    console.log("register succed!", result)
                },
                function (errMsg, caller) {
                    console.log("register fail!", errMsg)
                }, {})

        }
    } (userNameInput, emailInput, passwordInput, confirmPasswordInput)

    // Now the login button.
    this.loginBtn = this.navBar.getChildById("login-submit")
    var loginNameInput = this.navBar.getChildById("login-form-username-input")
    var loginPasswordInput = this.navBar.getChildById("login-form-password-input")

    this.userInfoDropdownLnk = this.navBar.getChildById("user-info-dropdown-lnk")
    this.loginDropdownLnk = this.navBar.getChildById("login-dropdown-lnk")
    this.logoutDropdownLnk = this.navBar.getChildById("logout-dropdown-lnk")

    // The logged user.
    this.account = null
    this.authorPostDiv = null

    // Only registered user can create a blog.
    this.newBlogLnk = this.navBar.getChildById("new-blog-lnk")

    // Login the user.
    this.loginBtn.element.onclick = function (blogManager, loginNameInput, loginPasswordInput) {
        return function () {
            // Get the value login information from interface.
            var userName = loginNameInput.element.value
            var password = loginPasswordInput.element.value

            // Here I will try to login in on the server.
            server.sessionManager.login(userName, password, "",
                function (result, caller) {
                    var userInfoDropdownLnk = caller.blogManager.userInfoDropdownLnk
                    var loginDropdownLnk = caller.blogManager.loginDropdownLnk
                    var logoutDropdownLnk = caller.blogManager.logoutDropdownLnk
                    var registerLnk = caller.blogManager.registerLnk
                    var newBlogLnk = caller.blogManager.newBlogLnk

                    // Set visibility
                    loginDropdownLnk.element.style.display = "none"
                    registerLnk.element.style.display = "none"
                    logoutDropdownLnk.element.style.display = ""
                    userInfoDropdownLnk.element.style.display = ""
                    newBlogLnk.element.style.display = ""

                    userInfoDropdownLnk.element.firstChild.innerHTML = result.M_accountPtr.M_id

                    // Keep track of the account ptr.
                    blogManager.account = result.M_accountPtr

                    // If the account dosent have user info I will show it.
                    if (result.M_accountPtr.M_userRef == "") {
                        userInfoDropdownLnk.element.className = "dropdown open"
                        // Set the account email here.
                        document.getElementById("email").value = result.M_accountPtr.M_email
                        document.getElementById("first_name").focus()
                    } else {
                        // In that case i will set the user values.
                        caller.blogManager.userInfoDropDown.getChildById("first_name").element.value = result.M_accountPtr.M_userRef.M_firstName
                        caller.blogManager.userInfoDropDown.getChildById("last_name").element.value = result.M_accountPtr.M_userRef.M_lastName
                        caller.blogManager.userInfoDropDown.getChildById("middle_name").element.value = result.M_accountPtr.M_userRef.M_middle
                        caller.blogManager.userInfoDropDown.getChildById("email").element.value = result.M_accountPtr.M_userRef.M_email
                        caller.blogManager.userInfoDropDown.getChildById("phone").element.value = result.M_accountPtr.M_userRef.M_phone
                    }

                    // Display the author post inside the the side well widget.
                    blogManager.displayAuthorPost()
                },
                function (errObj, caller) {
                    var err = errObj.dataMap.errorObj
                    var loginNameInput = caller.loginNameInput.element
                    var loginPasswordInput = caller.loginPasswordInput.element
                    var loginFormPasswordError = document.getElementById("login-form-password-error")
                    var loginFormUsernameError = document.getElementById("login-form-username-error")
                    loginPasswordInput.value = ""
                    if (err["M_id"] == "PASSWORD_MISMATCH_ERROR") {
                        loginFormPasswordError.innerHTML = err["M_body"]
                        loginPasswordInput.focus()
                    } else if (err["M_id"] == "ACCOUNT_DOESNT_EXIST_ERROR") {
                        loginFormUsernameError.innerHTML = err["M_body"]
                        loginNameInput.setSelectionRange(0, loginNameInput.value.length)
                        loginNameInput.focus()
                    }
                    // Clear the message after 2 seconds.
                    setTimeout(function (loginFormUsernameError, loginFormPasswordError) {
                        return function () {
                            loginFormUsernameError.innerHTML = ""
                            loginFormPasswordError.innerHTML = ""
                        }
                    } (loginFormUsernameError, loginFormPasswordError), 2000)

                }, { "blogManager": blogManager, "loginNameInput": loginNameInput, "loginPasswordInput": loginPasswordInput })
        }
    } (this, loginNameInput, loginPasswordInput)

    // Logout the user.
    this.logoutDropdownLnk.element.onclick = function (blogManager) {
        return function () {
            // So here I will call logout on the server.
            server.sessionManager.logout(
                // The session to close.
                server.sessionId,
                // Success callback
                function () {
                    var userInfoDropdownLnk = caller.blogManager.userInfoDropdownLnk
                    var loginDropdownLnk = caller.blogManager.loginDropdownLnk
                    var logoutDropdownLnk = caller.blogManager.logoutDropdownLnk
                    var newBlogLnk = caller.blogManager.newBlogLnk
                    // Reset the current account.
                    blogManager.account = null

                    // Set user interface element.
                    loginDropdownLnk.element.style.display = ""
                    logoutDropdownLnk.element.style.display = "none"
                    userInfoDropdownLnk.element.style.display = ""
                    newBlogLnk.element.style.display = "none"

                },
                // Error callback
                function () { },
                { "loginDropdownLnk": blogManager.loginDropdownLnk, "userInfoDropdownLnk": blogManager.userInfoDropdownLnk, "logoutDropdownLnk": blogManager.logoutDropdownLnk })
        }
    } (this)

    // Now the user infos.
    var firstNameInput = this.userInfoDropDown.getChildById("first_name")
    var lastNameInput = this.userInfoDropDown.getChildById("last_name")
    var middleNameInput = this.userInfoDropDown.getChildById("middle_name")
    var email = this.userInfoDropDown.getChildById("email")
    var phone = this.userInfoDropDown.getChildById("phone")
    this.saveUserInfoBtn = this.userInfoDropDown.getChildById("user-info-submit")

    // Now I will save the user information.
    this.saveUserInfoBtn.element.onclick = function (blogManager, firstNameInput, lastNameInput, middleNameInput, emailInput, phoneInput) {
        return function () {
            // So here I will create a new user if none exist.
            var user = null
            var exist = false
            if (blogManager.account.m_userRef != null) {
                user = blogManager.account.M_userRef
                exist = true
            } else {
                user = new CargoEntities.User()
                user.M_accountPtr = [blogManager.account.UUID] // Set the array 
            }

            // Set the user fields
            user.M_id = blogManager.account.M_id
            user.M_firstName = firstNameInput.element.value
            user.M_lastName = lastNameInput.element.value
            user.M_middle = middleNameInput.element.value
            user.M_email = emailInput.element.value
            user.M_phone = phone.element.value
            user.M_entitiesPtr = blogManager.account.M_entitiesPtr
            user.NeedSave = true

            if (!exist) {
                server.entityManager.createEntity(blogManager.account.M_entitiesPtr, "M_entities", "CargoEntities.User", "", user,
                    // Success callback
                    function (result, caller) {
                        // I that case I will set the account ref.
                        blogManager.account.M_userRef = result.UUID
                        blogManager.account.NeedSave = true
                        // Now I will save the account with it newly created user.
                        server.entityManager.saveEntity(blogManager.account)
                        caller.blogManager.userInfoLnk.className = "dropdown"
                        // Create the author...
                        caller.blogManager.saveAuthor(result)
                    },
                    // Error callback
                    function (errObj, caller) {
                    }, { "blogManager": blogManager })
            } else {
                // I that case I will save the user instead of create it.
                server.entityManager.saveEntity(user,
                    // The success callback
                    function (result, caller) {
                        caller.blogManager.userInfoLnk.className = "dropdown"
                        caller.blogManager.saveAuthor(result)
                    },
                    // The error callback
                    function (errObj, caller) {
                    },
                    { "blogManager": blogManager })
            }

        }
    } (this, firstNameInput, lastNameInput, middleNameInput, email, phone)

    /////////////////////////////////////////////////////////////////////////////////////////////////
    //  Event listener's
    /////////////////////////////////////////////////////////////////////////////////////////////////
    // The delete entity event.
    server.entityManager.attach(this, DeleteEntityEvent, function (evt, blogManager) {
        if (evt.dataMap["entity"].TYPENAME == blogPostTypeName) {
            if (evt.dataMap["entity"] != null && blogManager.activePostView.post != null) {
                if (blogManager.activePostView.post.UUID == evt.dataMap["entity"].UUID) {
                    console.log("delete post!")
                }
            }
        }
    })

    // The new entity event.
    server.entityManager.attach(this, NewEntityEvent, function (evt, blogManager) {
        // I will reinit the panel here...
        if (evt.dataMap["entity"].TYPENAME == blogPostTypeName) {
            if (evt.dataMap["entity"] != null && blogManager.activePostView.post != null) {
                if (blogManager.activePostView.post.UUID == evt.dataMap["entity"].UUID) {
                    console.log("new post!")
                }
            }
        }
    })

    // The update entity event.
    server.entityManager.attach(this, UpdateEntityEvent, function (evt, blogManager) {
        if (evt.dataMap["entity"].TYPENAME == blogPostTypeName) {
            if (evt.dataMap["entity"] != null && blogManager.activePostView.post != null) {
                // I will reinit the panel here...
                if (blogManager.activePostView.post.UUID == evt.dataMap["entity"].UUID) {
                    // Set the entity.
                    server.entityManager.entities[evt.dataMap["entity"].UUID] = evt.dataMap["entity"]

                    // Udate the author post.
                    blogManager.activePostView = new BlogPostView(blogManager.blogContainer, evt.dataMap["entity"])
                    
                    // Set the blog view editable.
                    blogManager.setEditable(blogManager.activePostView)
                    blogManager.displayAuthorPost()
                }
            }
        }
    })

    return this
}

/**
 * The post creator.
 */
BlogManager.prototype.saveAuthor = function (user) {
    // Here I will use the display name to keep the user uuid inside the 
    // field display name.
    var author = eval("new " + authorTypeName + "()")
    author.M_id = user.UUID // The related user uuid.
    // save the author information.
    server.entityManager.saveEntity(author)
}

/**
 * Set the content of author post.
 */
BlogManager.prototype.displayAuthorPost = function () {
    // The list of post by an author.
    if (this.authorPostDiv == null) {
        this.authorPostDiv = this.sideWellWidget.appendElement({ "tag": "div" }).down()
    } else {
        // reset it content.
        this.authorPostDiv.removeAllChilds()
        this.authorPostDiv.element.innerHTML = ""
    }

    // Now I will get the all post from a given author.
    server.entityManager.getEntityById("sql_info", authorTypeName, this.account.M_userRef.UUID,
        // The success callback.
        function (author, caller) {

            var authorPostDiv = caller.blogManager.authorPostDiv
            for (var i = 0; i < author.M_FK_blog_post_blog_author.length; i++) {
                var post = author.M_FK_blog_post_blog_author[i]
                if(server.entityManager.entities[post.UUID]!=undefined){
                    post = server.entityManager.entities[post.UUID]
                    author.M_FK_blog_post_blog_author[i] = post
                }
                // Here I will create the link with the title.
                authorPostDiv.appendElement({ "tag": "div", "class": "row" }).down()
                    .appendElement({ "tag": "div", "class": "col-md-1" }).down()
                    .appendElement({ "tag": "i", "class": "fa fa-trash-o delete-button", "style": "vertical-align: center;" }).up()
                    .appendElement({ "tag": "a", "id": post.UUID + "_lnk", "class": "col-md-10 control-label", "innerHtml": post.M_title, "href": "#" })

                var postLnk = authorPostDiv.getChildById(post.UUID + "_lnk")
                postLnk.element.onclick = function (post, blogManager) {
                    return function () {
                        blogManager.activePostView = new BlogPostView(blogManager.blogContainer, post)
                        // Set the blog view editable.
                        blogManager.setEditable(caller.blogManager.activePostView)
                    }
                } (post, caller.blogManager)
            }
        },
        // Error callback.
        function () { },
        { "blogManager": this })
}

/**
 * Create a new blog.
 */
BlogManager.prototype.createNewPost = function (author) {
    // Here I will use the data manager to get the number of post.
    var query = "SELECT MAX(id) FROM " + blogPostTypeName
    server.dataManager.read("Blog", query, ["int"], [],
        // success callback
        function (result, caller) {
            var lastId = 0
            if (result[0][0][0] != null) {
                lastId = result[0][0][0]
            }

            var userUuid
            if(isString(caller.blogManager.account.M_userRef)){
                userUuid = caller.blogManager.account.M_userRef
            }else{
                userUuid = caller.blogManager.account.M_userRef.UUID
            }

            // Now I will save the post...
            // The post is own by author, so if we delete an author all it's post will be deleted.
            server.entityManager.getEntityById("sql_info", authorTypeName, userUuid,
                function (author, caller) {
                    if(author.M_id.length == 0){
                        return // Do nothing if the author id is not set properly.
                    }

                    // So here I will create a new blog post and from it I will set it view.
                    var post = eval("new " + blogPostTypeName + "()")
                    post.M_id = caller.lastId + 1
                    post.M_title = "Blog Post Title"
                    post.M_article = "<p class=\"lead\">Lorem ipsum dolor sit amet, consectetur adipisicing elit. Ducimus, vero, obcaecati, aut, error quam sapiente nemo saepe quibusdam sit excepturi nam quia corporis eligendi eos magni recusandae laborum minus inventore?</p>"
                        + "<p>Lorem ipsum dolor sit amet, consectetur adipisicing elit. Ut, tenetur natus doloremque laborum quos iste ipsum rerum obcaecati impedit odit illo dolorum ab tempora nihil dicta earum fugiat. Temporibus, voluptatibus.</p>"
                        + " <p>Lorem ipsum dolor sit amet, consectetur adipisicing elit. Eos, doloribus, dolorem iusto blanditiis unde eius illum consequuntur neque dicta incidunt ullam ea hic porro optio ratione repellat perspiciatis. Enim, iure!</p>"
                        + "<p>Lorem ipsum dolor sit amet, consectetur adipisicing elit. Error, nostrum, aliquid, animi, ut quas placeat totam sunt tempora commodi nihil ullam alias modi dicta saepe minima ab quo voluptatem obcaecati?</p>"
                        + "<p>Lorem ipsum dolor sit amet, consectetur adipisicing elit. Harum, dolor quis. Sunt, ut, explicabo, aliquam tenetur ratione tempore quidem voluptates cupiditate voluptas illo saepe quaerat numquam recusandae? Qui, necessitatibus, est!</p>"
                    post.M_blog_author_id = author.M_id
                    post.M_date_published = new Date()
                    post.M_featuread = 1
                    post.M_enabled = 1
                    post.M_comments_enabled = 1
                    post.M_views = 0
                    
                    // Link to the blog author object.

                    //post.M_FK_blog_post_blog_author = author.UUID


                    server.entityManager.createEntity(author.UUID, "M_FK_blog_post_blog_author", blogPostTypeName, "", post,
                        // Success callback.
                        function (post, caller) {
                            // Create a new Blog.
                            caller.blogManager.activePostView = new BlogPostView(caller.blogManager.blogContainer, post)

                            // Set the blog view editable.
                            caller.blogManager.setEditable(caller.blogManager.activePostView)
                        },
                        // Error callback
                        function (errObj, caller) {
                            // Error here.
                        }, { "blogManager": caller.blogManager })

                },
                function (errObj, caller) {

                }, { "blogManager": caller.blogManager, "lastId": lastId })
        },
        // progress callback
        function (index, total, caller) {

        },
        // error callback
        function (errObj, caller) {

        },
        { "blogManager": this })
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
    var setTitleCallback = function (blogManager) {
        return function (div, btn) {
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
                saveBtn.element.onclick = function (inputTitle, div, blogManager) {
                    return function (evt) {
                        // Stop event propagation so we will not return direclty here...
                        evt.stopPropagation()
                        // Here I will set back the text inside the h1 element.
                        div.element.innerText = inputTitle.element.value
                        blogManager.activePostView.post.M_title = inputTitle.element.value
                        setEditable(div, setTitleCallback)
                        blogManager.saveActivePost()
                    }
                } (inputTitle, div, blogManager)
            }
        }
    } (this)

    // Set the title.
    setEditable(blogView.titleDiv, setTitleCallback)

    // The author will be the logged user.

    // Set the content div
    setEditable(blogView.pageContentDiv,
        function (blogManager) {
            return function (div) {
                // I made use of http://summernote.org/ Great work folk's
                $('#page-content').summernote();

                // I will append the save button to existing toolbar, it's so easy...
                var saveBtn = new Element(document.getElementsByClassName("note-view")[0], { "tag": "button", "tabindex": -1, "type": "button", "class": "note-btn btn btn-default btn-sm btn btn-primary btn-save", "title": "save", "data-original-title": "save" })
                saveBtn.appendElement({ "tag": "i", "class": "fa fa-floppy-o" })

                // Now the save action.
                saveBtn.element.onclick = function (blogManager) {
                    return function () {

                        // remove the editor.
                        $('#page-content').summernote('destroy');

                        // Set the inner html value to the post.
                        blogManager.activePostView.post.M_article = document.getElementById("page-content").innerHTML

                        blogManager.saveActivePost()
                    }
                } (blogManager)
            }
        } (this))

    // Here I will use use the clock to set the time.
}

/**
 * Save the active blog.
 */
BlogManager.prototype.saveActivePost = function () {
    // Here the post to save is in the active view.
    var post = this.activePostView.post
    var author = server.entityManager.
    server.entityManager.saveEntity(post)
}

/**
 * There is the blog the blog view.
 */
var BlogPostView = function (parent, post) {

    // Reset the parent content here.
    parent.removeAllChilds()
    parent.element.innerHTML = ""

    // A reference to the post.
    this.post = post

    // The blog text interface elements
    var languageInfo = {
        "en": {
            "blog-post-title": "Blog Post Title",
            "written-by": "by ",
            "comment-header": "Leave a Comment",
            "submit-comment": "Submit",
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
    this.titleDiv.element.innerHTML = this.post.M_title

    // The author div
    this.authorDiv = this.pageContainer.getChildById("author-name")

    // The entity uuid.
    server.entityManager.getEntityByUuid(this.post.M_author_id,
        // Success Callback
        function (result, caller) {
            caller.authorDiv.element.innerHTML = result.M_firstName + " " + result.M_lastName
        },
        // Error Callback
        function (errObj, caller) {
            // Print error here...
        },
        { "authorDiv": this.authorDiv })


    // The created div.
    this.createdDiv = this.pageContainer.getChildById("created-date")
    this.createdDiv.element.innerHTML = " " + this.post.M_date_published

    // The page content.
    this.pageContentDiv = this.pageContainer.getChildById("page-content")
    this.pageContentDiv.element.innerHTML = this.post.M_article

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