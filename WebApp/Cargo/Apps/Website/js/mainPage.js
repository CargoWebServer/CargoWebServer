/**
 * Utility class use to manage data element of the blog.
 */
var MainPage = function (parent) {
    // The language Info JSON contain a mapping of id and text.
    // The id's here are the same as the html element's in the entity panel,
    // so we can control the panel title and field's labal text.
    var languageInfo = {
        "en": {
            // The navigation bar.
            "toggle-navigation": "Toggle navigation",
            "start-lnk": "Start Boostrap",
            "home-lnk": "home",
            "tutorial-lnk": "tutorials",
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
            "blog-search-title": "Search",
            "blog-categories-title": "Categories",
            "footer-text": "CargoWebServer&copy; 2017",
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
        .appendElement({"tag":"img", "src":"img/wheel.svg", "id":"cargo-home-logo", "style":"width:30px; height: 30px; margin-top: 10px;"})
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
        .appendElement({ "tag": "a", "id": "home-lnk", "href": "#" }).up()
        .appendElement({ "tag": "li" }).down()
        .appendElement({ "tag": "a", "id": "tutorial-lnk", "href": "#" }).up()
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
    this.pageContent = this.container.appendElement({ "tag": "div", "class": "col-lg-8" }).down()

    // Blog Sidebar Widgets Column
    // The search panel.
    this.container.appendElement({ "tag": "div", "class": "col-md-4" }).down()
        .appendElement({ "tag": "div", "class": "well" }).down()
        .appendElement({ "tag": "span", "class": "input-group" }).down()
        .appendElement({ "tag": "input", "type": "text", "class": "form-control" })
        .appendElement({ "tag": "span", "class": "input-group-btn" }).down()
        .appendElement({ "tag": "button", class: "btn btn-default" }).down()
        .appendElement({ "tag": "span", "class": "fa fa-search" }).up().up().up().up()

        // Blog Categories Well
        .appendElement({ "tag": "div", "class": "well", "id": "blog-categories" }).down()
        .appendElement({ "tag": "h4", "id": "blog-categories-title" }).down()
        .appendElement({ "tag": "div", "id": "new-blog-categories-btn" }).down()
        .appendElement({ "tag": "i", "class": "fa fa-plus" }).up().up()
        .appendElement({ "tag": "div", "class": "row", "id": "blog-categories-content" }).down() // The list of all categories
        .appendElement({ "tag": "div", "class": "col-lg-6" }).down() // col 1
        .appendElement({ "tag": "ul", "class": "list-unstyled" }).up()
        .appendElement({ "tag": "div", "class": "col-lg-6" }).down() // col 2
        .appendElement({ "tag": "ul", "class": "list-unstyled" }).up().up().up()

        // The side well widget.
        .appendElement({ "tag": "div", "class": "well", "style": "display: none;" }).down()
        .appendElement({ "tag": "h4", "id": "blog-post-by-author" })
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
    this.homeLnk = this.navBar.getChildById("home-lnk")
    this.tutorialLnk = this.navBar.getChildById("tutorial-lnk")
    this.homeLogoLnk = this.navBar.getChildById("cargo-home-logo")
    this.registerLnk = this.navBar.getChildById("register-lnk")
    this.userInfoLnk = this.navBar.getChildById("user-info-lnk")
    this.authorPostTitle = this.container.getChildById("blog-post-by-author")
    
    // The new category button...
    this.newCategoryBtn = this.container.getChildById("new-blog-categories-btn")
    this.categoryContentDiv = this.container.getChildById("blog-categories-content")

    // Set the tutorial page.
    this.tutorialPage = new TutorialPage()
    this.tutorialLnk.element.onclick = function(tutorialPage){
        return function(){
            tutorialPage.display(mainPage.pageContent)
        }
    }(this.tutorialPage)

    // Set the home page.
    this.homePage = new HomePage(this.pageContent)
    this.homeLnk.element.onclick = this.homeLogoLnk.element.onclick =  function (homePage) {
        return function () {
            // Display the home page in the page content area.
            homePage.display(mainPage.pageContent)
        }
    }(this.homePage)


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
    }(this.registerLnk, this.userInfoLnk)

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
    }(this.loginLnk, this.userInfoLnk)

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
    }(this.loginLnk, this.registerLnk)

    // Append a new category...
    this.newCategoryBtn.element.onclick = function (mainPage) {
        return function () {
            // Here I will create a div to display to edit the new category name...
            if (document.getElementById("save_category_btn") == undefined) {

                var newCategoryPanel = mainPage.newCategoryBtn.appendElement({ "tag": "div", "class": "well", "style": "display: table; background-color: white; position: absolute; width: 200px;" }).down()
                newCategoryPanel.appendElement({ "tag": "input", "id": "new_category_input", "class": "form-control", "placeholder": "Category Name" })
                    .appendElement({ "tag": "button", "id": "save_category_btn", "class": "form-control btn btn-info", "style": "margin-top: 5px;" }).down()
                    .appendElement({ "tag": "span", "innerHtml": " Save", "style": "padding-right: 10px;" })
                    .appendElement({ "tag": "i", "class": "fa fa fa-floppy-o" })

                // Here the save 
                var saveCategoryBtn = mainPage.newCategoryBtn.getChildById("save_category_btn")
                var newCategoryInput = mainPage.newCategoryBtn.getChildById("new_category_input")
                newCategoryInput.element.focus() // set the focus...

                saveCategoryBtn.element.onclick = function (mainPage, newCategoryPanel, newCategoryInput) {
                    return function (evt) {
                        evt.stopPropagation();

                        var category = eval("new " + categoryTypeName + "()")
                        category.M_name = newCategoryInput.element.value
                        category.M_enable = true
                        category.M_date_created = new Date()

                        mainPage.appendCategory(category)

                        // Remove the panel
                        mainPage.newCategoryBtn.removeElement(newCategoryPanel)
                    }
                }(mainPage, newCategoryPanel, newCategoryInput)

            } else {
                var newCategoryPanel = mainPage.newCategoryBtn.childs[Object.keys(mainPage.newCategoryBtn.childs)[1]]
                mainPage.newCategoryBtn.removeElement(newCategoryPanel)
            }
        }
    }(this)

    // Get the list of existing categories.
    server.entityManager.getEntities(categoryTypeName, "Blog", null, 0, -1, [], true,
        // progress callback
        function (index, total, caller) {
            // Nothing to do here.
        },
        // success callback
        function (results, caller) {
            for (var i = 0; i < results.length; i++) {
                // Append the category to the category panel.
                caller.mainPage.appendCategory(results[i])
            }
        },
        // error callback
        function (errObj, caller) {

        },
        { "mainPage": this })

    // Now the mouse out event.
    this.registerDropDown = this.navBar.getChildById("register-dropdown")

    // The login dropdown.
    this.loginDropDown = this.navBar.getChildById("login-dropdown")

    // The user info dropdown
    this.userInfoDropDown = this.navBar.getChildById("user-info-dropdown")

    // Hide the window if the mouse leave it...
    this.userInfoDropDown.element.onmouseleave = this.registerDropDown.element.onmouseleave = this.loginDropDown.element.onmouseleave = function (evt) {
        // Here I will test if the mouse is inside it parent area, a little workaround!
        var coord = getCoords(this)
        if (evt.pageX < coord.left || evt.pageX > coord.left + this.offsetWidth || evt.pageY < coord.top || evt.pageY > coord.top + this.offsetWidth) {
            this.parentNode.className = "dropdown"
        }
    }

    // Create a new blog.
    this.createBlog = this.navBar.getChildById("create-lnk")

    this.createBlog.element.onclick = function (mainPage) {
        return function () {
            // Create a new blog.
            mainPage.createNewPost()
        }
    }(this)

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
    }(userNameInput, emailInput, passwordInput, confirmPasswordInput)

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
    this.loginBtn.element.onclick = function (mainPage, loginNameInput, loginPasswordInput) {
        return function () {
            // Get the value login information from interface.
            var userName = loginNameInput.element.value
            var password = loginPasswordInput.element.value

            // Here I will try to login in on the server.
            server.sessionManager.login(userName, password, "",
                function (session, caller) {
                    caller.successCallback = function (session) {

                        // short variable name.
                        var userInfoDropdownLnk = this.mainPage.userInfoDropdownLnk
                        var loginDropdownLnk = this.mainPage.loginDropdownLnk
                        var logoutDropdownLnk = this.mainPage.logoutDropdownLnk
                        var registerLnk = this.mainPage.registerLnk
                        var newBlogLnk = this.mainPage.newBlogLnk

                        // Set visibility
                        loginDropdownLnk.element.style.display = "none"
                        registerLnk.element.style.display = "none"
                        logoutDropdownLnk.element.style.display = ""
                        userInfoDropdownLnk.element.style.display = ""
                        newBlogLnk.element.style.display = ""
                        mainPage.newCategoryBtn.element.style.display = "block"
                        mainPage.authorPostTitle.element.innerHTML = session.M_accountPtr.M_userRef.M_firstName + " " + session.M_accountPtr.M_userRef.M_lastName + " post's"

                        userInfoDropdownLnk.element.firstChild.innerHTML = session.M_accountPtr.M_id

                        // Keep track of the account ptr.
                        mainPage.account = session.M_accountPtr

                        // If the account dosent have user info I will show it.
                        if (session.M_accountPtr.M_userRef == "") {
                            userInfoDropdownLnk.element.className = "dropdown open"
                            // Set the account email here.
                            document.getElementById("email").value = session.M_accountPtr.M_email
                            document.getElementById("first_name").focus()
                        } else {
                            // In that case i will set the user values.
                            this.mainPage.userInfoDropDown.getChildById("first_name").element.value = session.M_accountPtr.M_userRef.M_firstName
                            this.mainPage.userInfoDropDown.getChildById("last_name").element.value = session.M_accountPtr.M_userRef.M_lastName
                            this.mainPage.userInfoDropDown.getChildById("middle_name").element.value = session.M_accountPtr.M_userRef.M_middle
                            this.mainPage.userInfoDropDown.getChildById("email").element.value = session.M_accountPtr.M_userRef.M_email
                            this.mainPage.userInfoDropDown.getChildById("phone").element.value = session.M_accountPtr.M_userRef.M_phone
                            // Display the author post inside the the side well widget.
                            this.mainPage.displayAuthorPost()

                        }

                        // TODO display personal author page here instead of nothing...
                        this.mainPage.pageContent.removeAllChilds()
                        if (this.mainPage.activePostView != null) {
                            // redisplay the active post to unlock category edit button's.
                            this.mainPage.displayPost(this.mainPage.activePostView.post)
                        }

                    }


                    // Display the category delete button...
                    var deleteCategoryButtons = document.getElementsByClassName("category_delete_btn")
                    for (var i = 0; i < deleteCategoryButtons.length; i++) {
                        deleteCategoryButtons[i].style.display = ""
                    }

                    // Save the active session id on the session manager.
                    activeSessionAccountId = session.M_accountPtr

                    // Initialisation of the account pointer and call the login callback on success.
                    session["set_M_accountPtr_" + session.M_accountPtr + "_ref"](function (caller, session) {
                        return function (accountPtr) {
                            // call the callback after the session is intialysed.
                            // Keep track of the accountUuid for future access.
                            localStorage.setItem("accountUuid", accountPtr.UUID)
                            if (accountPtr["set_M_userRef_" + accountPtr.M_userRef + "_ref"] != undefined) {
                                accountPtr["set_M_userRef_" + accountPtr.M_userRef + "_ref"](function (session, caller) {
                                    return function () {
                                        caller.successCallback(session)
                                    }
                                }(session, caller))
                            } else {
                                caller.successCallback(session)
                            }
                        }
                    }(caller, session))

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
                    }(loginFormUsernameError, loginFormPasswordError), 2000)

                }, { "mainPage": mainPage, "loginNameInput": loginNameInput, "loginPasswordInput": loginPasswordInput })
        }
    }(this, loginNameInput, loginPasswordInput)

    // Logout the user.
    this.logoutDropdownLnk.element.onclick = function (mainPage) {
        return function () {
            // So here I will call logout on the server.
            server.sessionManager.logout(
                // The session to close.
                server.sessionId,
                // Success callback
                function () {
                    var userInfoDropdownLnk = caller.mainPage.userInfoDropdownLnk
                    var loginDropdownLnk = caller.mainPage.loginDropdownLnk
                    var logoutDropdownLnk = caller.mainPage.logoutDropdownLnk
                    var newBlogLnk = caller.mainPage.newBlogLnk
                    // Reset the current account.
                    mainPage.account = null

                    // Set user interface element.
                    loginDropdownLnk.element.style.display = ""
                    logoutDropdownLnk.element.style.display = "none"
                    userInfoDropdownLnk.element.style.display = ""
                    newBlogLnk.element.style.display = "none"

                },
                // Error callback
                function () { },
                { "loginDropdownLnk": mainPage.loginDropdownLnk, "userInfoDropdownLnk": mainPage.userInfoDropdownLnk, "logoutDropdownLnk": mainPage.logoutDropdownLnk })
        }
    }(this)

    // Now the user infos.
    var firstNameInput = this.userInfoDropDown.getChildById("first_name")
    var lastNameInput = this.userInfoDropDown.getChildById("last_name")
    var middleNameInput = this.userInfoDropDown.getChildById("middle_name")
    var email = this.userInfoDropDown.getChildById("email")
    var phone = this.userInfoDropDown.getChildById("phone")
    this.saveUserInfoBtn = this.userInfoDropDown.getChildById("user-info-submit")

    // Now I will save the user information.
    this.saveUserInfoBtn.element.onclick = function (mainPage, firstNameInput, lastNameInput, middleNameInput, emailInput, phoneInput) {
        return function () {
            // So here I will create a new user if none exist.
            var user = null
            var exist = false
            if (isObject(mainPage.account.M_userRef)) {
                user = mainPage.account.M_userRef
                user.setAccounts(mainPage.account)
                exist = true
            } else {
                user = new CargoEntities.User()
                // user.M_accountPtr = [mainPage.account.UUID] // Set the array 
                user.setAccounts(mainPage.account)
            }

            // Set the user fields
            user.M_id = mainPage.account.M_id
            user.M_firstName = firstNameInput.element.value
            user.M_lastName = lastNameInput.element.value
            user.M_middle = middleNameInput.element.value
            user.M_email = emailInput.element.value
            user.M_phone = phone.element.value
            user.M_entitiesPtr = mainPage.account.M_entitiesPtr
            user.NeedSave = true

            if (!exist) {
                server.entityManager.createEntity(mainPage.account.M_entitiesPtr, "M_entities", "CargoEntities.User", "", user,
                    // Success callback
                    function (result, caller) {
                        // I that case I will set the account ref.
                        mainPage.account.setUserRef(result)
                        mainPage.account.NeedSave = true

                        // Now I will save the account with it newly created user.
                        server.entityManager.saveEntity(mainPage.account)
                        caller.mainPage.userInfoLnk.className = "dropdown"
                        // Create the author...
                        caller.mainPage.saveAuthor(result)
                    },
                    // Error callback
                    function (errObj, caller) {
                    }, { "mainPage": mainPage })
            } else {
                // I that case I will save the user instead of create it.
                server.entityManager.saveEntity(user,
                    // The success callback
                    function (result, caller) {
                        caller.mainPage.userInfoLnk.className = "dropdown"
                        caller.mainPage.saveAuthor(result)
                    },
                    // The error callback
                    function (errObj, caller) {
                    },
                    { "mainPage": mainPage })
            }

        }
    }(this, firstNameInput, lastNameInput, middleNameInput, email, phone)

    /////////////////////////////////////////////////////////////////////////////////////////////////
    //  Event listener's
    /////////////////////////////////////////////////////////////////////////////////////////////////

    // The delete entity event.
    server.entityManager.attach(this, DeleteEntityEvent, function (evt, mainPage) {
        if (evt.dataMap["entity"].TYPENAME == blogPostTypeName) {
            // remove the post lnk for the lst.
            if (mainPage.activePostView != null) {
                if (evt.dataMap["entity"] != null && mainPage.activePostView.post != null) {
                    if (mainPage.activePostView.post.UUID == evt.dataMap["entity"].UUID) {
                        // Remove the content of the page.
                        mainPage.pageContent.removeAllChilds()
                    }
                }
            }
        } else if (evt.dataMap["entity"].TYPENAME == categoryTypeName) {
            var category = evt.dataMap["entity"]
            // I will use DOM to remove the category li
            var li = document.getElementById(category.UUID + "_li")
            if (li != undefined) {
                li.parentNode.removeChild(li)
            }
        }
    })

    // The update entity event.
    server.entityManager.attach(this, UpdateEntityEvent, function (evt, mainPage) {
        if (evt.dataMap["entity"].TYPENAME == blogPostTypeName) {
            // In case of post change.
            mainPage.displayAuthorPost()
            if (mainPage.activePostView != null) {
                if (evt.dataMap["entity"] != null && mainPage.activePostView.post != null) {
                    // I will reinit the panel here...
                    if (mainPage.activePostView.post.UUID == evt.dataMap["entity"].UUID) {
                        // Udate the author post.
                        mainPage.displayPost(entities[evt.dataMap["entity"].UUID])

                        // Set the blog view editable.
                        mainPage.setEditable(mainPage.activePostView)
                    }
                }
            }
        } else if (evt.dataMap["entity"].TYPENAME == authorTypeName) {
            if (mainPage.account != null) {
                // In case of author change...
                mainPage.displayAuthorPost()
            }
        }
    })

    // New entity event.
    server.entityManager.attach(this, NewEntityEvent, function (evt, blogView) {
        if (evt.dataMap["entity"].TYPENAME == categoryTypeName) {
            var category = evt.dataMap["entity"]
            blogView.appendCategory(category)
        }
    })

    return this
}

/**
 * The post creator.
 */
MainPage.prototype.saveAuthor = function (user) {
    // Here I will use the display name to keep the user uuid inside the 
    // field display name.
    var author = eval("new " + authorTypeName + "()")
    author.M_id = user.UUID // The related user uuid.
    // save the author information.
    server.entityManager.saveEntity(author)
}

/**
 * Display a given post entity.
 */
MainPage.prototype.displayPost = function (post) {
    this.pageContent.removeAllChilds()
    this.activePostView = new BlogPostView(this.pageContent, post, this.categoryContentDiv)
}

/**
 * Set the content of author post.
 */
MainPage.prototype.displayAuthorPost = function () {
    // The list of post by an author
    if (this.authorPostDiv == null) {
        this.authorPostDiv = this.sideWellWidget.appendElement({ "tag": "div" }).down()
    } else {
        this.authorPostDiv.removeAllChilds()
    }

    this.sideWellWidget.element.parentNode.style.display = "block"

    // Now I will get the all post from a given author.
    if (isObject(this.account.M_userRef)) {
        server.entityManager.getEntityById(authorTypeName, "sql_info", [this.account.M_userRef.UUID],
            // The success callback.
            function (author, caller) {

                var authorPostDiv = caller.mainPage.authorPostDiv
                for (var i = 0; i < author.M_FK_blog_post_blog_author.length; i++) {
                    var post = author.M_FK_blog_post_blog_author[i]
                    if (entities[post.UUID] != undefined) {
                        post = entities[post.UUID]
                        author.M_FK_blog_post_blog_author[i] = post
                    }

                    // Here I will create the link with the title if it not already exist.
                    if (document.getElementById(post.UUID + "_lnk") == undefined) {
                        authorPostDiv.appendElement({ "tag": "div", "class": "row" }).down()
                            .appendElement({ "tag": "div", "class": "col-md-1" }).down()
                            .appendElement({ "tag": "i", "id": post.UUID + "_delete_btn", "class": "fa fa-trash-o delete-button", "style": "vertical-align: middle;" }).up()
                            .appendElement({ "tag": "a", "id": post.UUID + "_lnk", "class": "col-md-10 control-label", "innerHtml": post.M_title, "href": "#", "style": "padding-left: 5px" })

                        var postLnk = authorPostDiv.getChildById(post.UUID + "_lnk")
                        postLnk.element.onclick = function (post, mainPage) {
                            return function () {
                                mainPage.displayPost(post)

                                // Set the blog view editable.
                                mainPage.setEditable(caller.mainPage.activePostView)

                                // Here I will also display the new category button.
                                mainPage.newCategoryBtn.element.style.display = "block"

                            }
                        }(post, caller.mainPage)

                        var postDeleteBtn = authorPostDiv.getChildById(post.UUID + "_delete_btn")
                        postDeleteBtn.element.onclick = function (post) {
                            return function () {
                                console.log("-------> delete post: ", post)
                                if (post != undefined) {
                                    server.entityManager.removeEntity(post.UUID,
                                        // Success callback 
                                        function (results, caller) {
                                            console.log("Entity was remove sucessfully")
                                        },
                                        // Error callback
                                        function (errObj, caller) {
                                            // Nothing to do here...
                                        }, {/* no caller. */ })
                                }
                            }
                        }(post)
                    }
                }
            },
            // Error callback.
            function () { },
            { "mainPage": this })
    }
}

/**
 * Create a new post.
 */
MainPage.prototype.createNewPost = function (author) {
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
            if (isString(caller.mainPage.account.M_userRef)) {
                userUuid = caller.mainPage.account.M_userRef
            } else {
                userUuid = caller.mainPage.account.M_userRef.UUID
            }

            // Now I will save the post...
            // The post is own by author, so if we delete an author all it's post will be deleted.
            server.entityManager.getEntityById(authorTypeName, "sql_info", [userUuid],
                function (author, caller) {
                    if (author.M_id.length == 0) {
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
                    post.M_FK_blog_post_blog_author = author.UUID


                    server.entityManager.createEntity(author.UUID, "M_FK_blog_post_blog_author", blogPostTypeName, "", post,
                        // Success callback.
                        function (post, caller) {
                            // Create a new Blog.
                            caller.mainPage.displayPost(post)
                            // Set the blog view editable.
                            caller.mainPage.setEditable(caller.mainPage.activePostView)
                        },
                        // Error callback
                        function (errObj, caller) {
                            // Error here.
                        }, { "mainPage": caller.mainPage })

                },
                function (errObj, caller) {

                }, { "mainPage": caller.mainPage, "lastId": lastId })
        },
        // progress callback
        function (index, total, caller) {

        },
        // error callback
        function (errObj, caller) {

        },
        { "mainPage": this })
}

/**
 * Inject functionality to edit parts of the blog.
 */
MainPage.prototype.setEditable = function (blogView) {

    function setEditable(div, onClickCallback) {
        div.element.style.position = "relative"
        div.appendElement({ "tag": "div", "class": "edit-button", "title": "edit" }).down()
            .appendElement({ "tag": "i", "class": "fa fa-pencil-square" })

        // The action to do on click.
        div.element.onclick = function (onClickCallback, div) {
            return function () {
                onClickCallback(div, this)
            }
        }(onClickCallback, div)
    }

    // Set the title div.
    var setTitleCallback = function (mainPage) {
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
                saveBtn.element.onclick = function (inputTitle, div, mainPage) {
                    return function (evt) {
                        // Stop event propagation so we will not return direclty here...
                        evt.stopPropagation()

                        // Here I will set back the text inside the h1 element.
                        div.element.innerText = inputTitle.element.value

                        mainPage.activePostView.post.M_title = inputTitle.element.value

                        // Here I will remove edit button.
                        var editButtons = document.getElementsByClassName("edit-button")
                        for (var i = 0; i < editButtons.length; i++) {
                            editButtons[i].parentNode.removeChild(editButtons[i])
                        }

                        // I will also save the post content...
                        mainPage.activePostView.post.M_article = document.getElementById("article-div").innerHTML

                        setEditable(div, setTitleCallback)

                        mainPage.saveActivePost()
                    }
                }(inputTitle, div, mainPage)
            }
        }
    }(this)

    // Set the title.
    setEditable(blogView.titleDiv, setTitleCallback)

    // The author will be the logged user.

    // Set the content div
    setEditable(blogView.pageContentDiv,
        function (mainPage) {
            return function (div) {
                // I made use of http://summernote.org/ Great work folk's
                $('#article-div').summernote();

                // I will append the save button to existing toolbar, it's so easy...
                var saveBtn = new Element(document.getElementsByClassName("note-view")[0], { "tag": "button", "tabindex": -1, "type": "button", "class": "note-btn btn btn-default btn-sm btn btn-primary btn-save", "title": "save", "data-original-title": "save" })
                saveBtn.appendElement({ "tag": "i", "class": "fa fa-floppy-o" })

                // Now the save action.
                saveBtn.element.onclick = function (mainPage) {
                    return function () {
                        // Set the inner html value to the post.
                        // Here I will remove edit button.
                        var editButtons = document.getElementsByClassName("edit-button")
                        for (var i = 0; i < editButtons.length; i++) {
                            editButtons[i].parentNode.removeChild(editButtons[i])
                        }

                        $('#article-div').summernote('destroy');
                        mainPage.activePostView.post.M_article = document.getElementById("article-div").innerHTML
                        mainPage.saveActivePost()
                    }
                }(mainPage)
            }
        }(this))

    // Here I will use use the clock to set the time.
}

/**
 * Save the active blog.
 */
MainPage.prototype.saveActivePost = function () {
    // Here the post to save is in the active view.
    var post = this.activePostView.post

    // Here I will create the thumbnail image.
    var pageContent = document.getElementById("page-content")
    html2canvas(pageContent, {
        onrendered: function (post) {
            return function (canvas) {
                // Save the canvas png image by default as data url.
                post.M_thumbnail = canvas.toDataURL()
                server.entityManager.saveEntity(post)
            }
        }(post),
        width: pageContent.offsetWidth,
        height: pageContent.offsetHeigth
    });
}

/**
 * Append a new category in the category div.
 */
MainPage.prototype.appendCategory = function (category) {
    if (category.UUID != undefined) {
        // Here the category dosen't exist...
        var query = "SELECT MAX(id) FROM " + categoryTypeName
        server.dataManager.read("Blog", query, ["int"], [],
            // success callback
            function (results, category) {
                // The last id...
                var lastId = 1
                if (results[0][0][0] != null) {
                    lastId = parseInt(results[0][0]) + 1
                }

                // Now I will save the category...
                category.M_id = lastId
                server.entityManager.saveEntity(category,
                    // Success callback
                    function (result, caller) {
                        // Nothing to do here I will use new entity event instead.
                    },
                    // Error callback
                    function () {

                    },
                    {})
            },
            // progress callback
            function (index, total, caller) {

            },
            // error callback
            function (errObj) {

            },
            category)

    } else {
        // Here the category already exist.
        var categoryContentDiv
        var l0 = this.categoryContentDiv.childs[Object.keys(this.categoryContentDiv.childs)[0]].element.firstChild.childNodes.length
        var l1 = this.categoryContentDiv.childs[Object.keys(this.categoryContentDiv.childs)[1]].element.firstChild.childNodes.length
        if (l1 <= l0) {
            categoryContentDiv = this.categoryContentDiv.childs[Object.keys(this.categoryContentDiv.childs)[1]].lastChild
        } else {
            categoryContentDiv = this.categoryContentDiv.childs[Object.keys(this.categoryContentDiv.childs)[0]].lastChild
        }

        // I will append the category...
        var lnk = categoryContentDiv.appendElement({ "tag": "li", "id": category.UUID + "_li", "class": "category_lnk" }).down()
        var deleteCategoryBtn = lnk.appendElement({ "tag": "i", "id": category.UUID + "_delete_btn", "class": "fa fa-trash-o delete-button category_delete_btn", "style": "vertical-align: middle; display: none;" }).down()
        lnk.appendElement({ "tag": "a", innerHtml: category.M_name, "style": "padding-left: 7px; padding-right: 7px;" })

        // I will also append the checkbox to be use latter by post view...
        var categoryBtn = lnk.appendElement({ "tag": "input", "id": category.UUID + "_checkbox", "type": "checkbox", "style": "vertical-align: text-bottom; display: none;" }).down()

        // if the user is logged i will show the delete button.
        if (this.account != undefined) {
            deleteCategoryBtn.element.style.display = ""
            categoryBtn.element.style.display = ""
        }

        // Now the delete action.
        deleteCategoryBtn.element.onclick = function (category) {
            return function () {
                server.entityManager.removeEntity(category.UUID,
                    /** Success Callback */
                    function (result, caller) {

                    },
                    /** Error Callback  */
                    function (errObj, caller) {

                    },
                    {})
            }
        }(category)
    }
}
