/**
 * The panel where to edit the comment.
 */
var BlogPostCommentEditor = function (parent, parentDiv) {

    // Can be the a PostView or a CommentView depending if the comment is a 
    // new comment or a reply to existing one.
    this.parent = parent

    // The blog comments input.
    this.div = parentDiv.appendElement({ "tag": "div", "class": "well" }).down()
    this.div.appendElement({ "tag": "h4", id: "comment-header" })
        .appendElement({ "tag": "div", "class": "form-group" }).down()
        .appendElement({ "tag": "textarea", "id": "comment-text-area", "class": "form-control", "rows": "3" }).up()
        .appendElement({ "tag": "div", "class": "row" }).down()
        .appendElement({ "tag": "div", "class": "col-md-2" }).down()
        .appendElement({ "tag": "button", "type": "submit", "class": "btn btn-primary", "style": "display: none;", "id": "submit-comment" }).up()
        // social bar.
        .appendElement({ "tag": "div", "class": "col-md-10" }).down()
        .appendElement({ "tag": "div", "class": "social" }).down()
        // push the list to right.
        .appendElement({ "tag": "div", "style": "display: table-cell; width: 100%" })
        // Cargo
        .appendElement({ "tag": "div", "id": "cargo-oauth2-btn", "class": "socialBtn cargo", "title": "Cargo" }).down()
        .appendElement({ "tag": "i", "class": "icon-cargo" }).up()
        // facebook
        .appendElement({ "tag": "div", "id": "facebook-oauth2-btn", "class": "socialBtn face", "title": "Facebook" }).down()
        .appendElement({ "tag": "i", "class": "fa fa-facebook" }).up()
        // google+
        .appendElement({ "tag": "div", "id": "google-plus-oauth2-btn", "class": "socialBtn google", "title": "Google+" }).down()
        .appendElement({ "tag": "i", "class": "fa fa-google-plus" })

    // Here I will keep the block user that want to make a comment.
    this.activeUser = null
    this.activeUserSpan = this.div.getChildById("comment-header")

    // The summit button.
    this.newCommentTextArea = this.div.getChildById("comment-text-area")
    this.submitCommentBtn = this.div.getChildById("submit-comment")

    // The OAuth2 button.
    this.connectCargo = this.div.getChildById("cargo-oauth2-btn")
    this.connectFacebook = this.div.getChildById("facebook-oauth2-btn")
    this.connectGooglePlus = this.div.getChildById("google-plus-oauth2-btn")

    // The Cargo OAuth provider.
    this.connectCargo.element.onclick = function (commentEditor) {
        return function () {
            var sessionId = server.sessionId.replace("%", "%25")

            // TODO create function 
            var query = "http://" + server.hostName + ":" + server.port + "/api/Server/AccountManager/Me?p0=" + sessionId
            server.securityManager.getResource("1234", "openid profile email", query,
                function (results, caller) {
                    if (results.TYPENAME != undefined) {
                        // In that case I have an Account object.
                        if (results.TYPENAME == "CargoEntities.Account") {
                            // I will made use of the entity manager to initilyse the account properly
                            server.entityManager.getEntityByUuid(results.UUID,
                                // success callback
                                function (results, caller) {
                                    var userUuid
                                    if (isString(results.M_userRef)) {
                                        userUuid = results.M_userRef
                                    } else {
                                        userUuid = results.M_userRef.UUID
                                    }

                                    // Now init the user reference.
                                    results["set_M_userRef_" + userUuid + "_ref"](
                                        function (caller, results) {
                                            return function (ref) {
                                                caller.activeUser = eval("new " + userTypeName + "()")
                                                caller.activeUser.M_id = results.M_id
                                                caller.activeUser.M_first_name = results.M_userRef.M_firstName
                                                caller.activeUser.M_last_name = results.M_userRef.M_lastName
                                                caller.activeUser.M_email = results.M_userRef.M_email
                                                caller.activeUser.M_website = "www.cargoWebserver.com" // Powered by cargo.
                                                caller.connectFacebook.element.className = caller.connectFacebook.element.className.replace(" active", "")
                                                caller.connectGooglePlus.element.className = caller.connectGooglePlus.element.className.replace(" active", "")
                                                caller.connectCargo.element.className += " active"

                                                caller.activeUserSpan.element.innerHTML = "Leave a Comment as " + caller.activeUser.M_first_name + " " + caller.activeUser.M_last_name

                                                // here I will display the button post.
                                                caller.submitCommentBtn.element.style.display = ""
                                                // Set the text focus.
                                                caller.newCommentTextArea.element.focus()
                                            }
                                        }(caller, results)
                                    )

                                },
                                // error callback
                                function () {

                                }, caller)
                        }
                    }

                },
                function (errMsg, caller) {
                }, commentEditor)
        }
    }(this)

    // The summit comment action.
    this.submitCommentBtn.element.onclick = function (commentEditor) {
        return function () {
            // Create the comment.
            commentEditor.parent.appendComment(commentEditor.parent.post, commentEditor.activeUser, commentEditor.newCommentTextArea.element.value)
            commentEditor.newCommentTextArea.element.value = ""
        }
    }(this)

    return this
}

/**
 * The of a blog comment view.
 */
var BlogPostCommentView = function (parent, user, comment, post, id) {

    // So here i will create the comment.
    function setView(comment, blogPostCommentView) {
        // The formated date.
        var m = moment(comment.M_date);
        blogPostCommentView.commentDiv = parent.prependElement({ "tag": "div", class: "media" }).down()
        blogPostCommentView.commentDiv.appendElement({ "tag": "a", "class": "pull-left", "href": "#" }).down()
            .appendElement({ "tag": "img", "class": "media-object", "src": "http://placehold.it/64x64" }).up()
            .appendElement({ "tag": "div", "class": "media-body" }).down()
            .appendElement({ "tag": "h4", "innerHtml": comment.M_user_id }).down().appendElement({ "tag": "br" }).appendElement({ "tag": "small", "innerHtml": m.format('LLL') }).up().up()
            .appendElement({ "tag": "div", "innerHtml": comment.M_comment })
    }

    // Set the id.
    if (id == undefined) {
        var query = "SELECT MAX(id) FROM " + blogCommentTypeName
        server.dataManager.read("Blog", query, ["int"], [],
            // Success callback.
            function (result, caller) {
                var user = caller.user
                var post = caller.post
                var commentStr = caller.commentStr

                var lastId = 0
                if (result[0][0][0] != null) {
                    lastId = result[0][0][0]
                }

                var comment = eval("new " + blogCommentTypeName + "()")
                comment.M_id = lastId + 1
                comment.M_post_id = post.M_id
                comment.M_user_id = user.M_id
                comment.M_comment = commentStr
                comment.M_mark_read = false
                comment.M_enabled = false
                comment.M_date = new Date()

                // So here I will save the comment send by the user.
                // * Note that the parent of a comment is always a post, even comment reply. So 
                //   if a post is remove comments are also removed.
                server.entityManager.createEntity(post.UUID, "M_FK_blog_comment_blog_post", blogCommentTypeName, "", comment,
                    // Success callback.
                    function (comment, caller) {
                        // Create the comment view.
                        setView(caller.comment, caller.blogPostCommentView)
                    },
                    // Error callback
                    function (errObj, caller) {
                        // Error here.
                    }, { "blogPostCommentView": caller.blogPostCommentView, "comment": comment, "post": post })
            },
            // The progress callback
            function () { },
            // error callback.
            function (errObj, caller) {

            },
            // Caller.
            { "blogPostCommentView": this, "post": post, "commentStr": comment, "user": user })
    } else {
        // Get the comment by id.
        setView(comment, this)
    }

    return this
}

/**
 * Append a comment editor at the end of a comment to be able to create a reply.
 */
BlogPostCommentView.prototype.appendCommentReply = function () {

}

/**
 * Append a reply (comment) to a main comment.
 */
BlogPostCommentView.prototype.appendComment = function (parent, user, comment, post, id) {

    // So here i will create the comment.
    function setView(comment, blogPostCommentView) {
        // The formated date.
        var m = moment(comment.M_date);
        blogPostCommentView.commentDiv = parent.prependElement({ "tag": "div", class: "media" }).down()
        blogPostCommentView.commentDiv.appendElement({ "tag": "a", "class": "pull-left", "href": "#" }).down()
            .appendElement({ "tag": "img", "class": "media-object", "src": "http://placehold.it/64x64" }).up()
            .appendElement({ "tag": "div", "class": "media-body" }).down()
            .appendElement({ "tag": "h4", "innerHtml": comment.M_user_id }).down().appendElement({ "tag": "br" }).appendElement({ "tag": "small", "innerHtml": m.format('LLL') }).up().up()
            .appendElement({ "tag": "div", "innerHtml": comment.M_comment })
    }

    // Set the id.
    if (id == undefined) {
        var query = "SELECT MAX(id) FROM " + blogCommentTypeName
        server.dataManager.read("Blog", query, ["int"], [],
            // Success callback.
            function (result, caller) {
                var user = caller.user
                var post = caller.post
                var commentStr = caller.commentStr

                var lastId = 0
                if (result[0][0][0] != null) {
                    lastId = result[0][0][0]
                }

                var comment = eval("new " + blogCommentTypeName + "()")
                comment.M_id = lastId + 1
                comment.M_post_id = post.M_id
                comment.M_user_id = user.M_id
                comment.M_comment = commentStr
                comment.M_mark_read = false
                comment.M_enabled = false
                comment.M_date = new Date()

                // So here I will save the comment send by the user.
                // * Note that the parent of a comment is always a post, even comment reply. So 
                //   if a post is remove comments are also removed.
                server.entityManager.createEntity(post.UUID, "M_FK_blog_comment_blog_post", blogCommentTypeName, "", comment,
                    // Success callback.
                    function (comment, caller) {
                        // Create the comment view.
                        setView(caller.comment, caller.blogPostCommentView)
                    },
                    // Error callback
                    function (errObj, caller) {
                        // Error here.
                    }, { "blogPostCommentView": caller.blogPostCommentView, "comment": comment, "post": post })
            },
            // The progress callback
            function () { },
            // error callback.
            function (errObj, caller) {

            },
            // Caller.
            { "blogPostCommentView": this, "post": post, "commentStr": comment, "user": user })
    } else {
        // Get the comment by id.
        setView(comment, this)
    }

    return this
}

/**
 * Append a comment editor at the end of a comment to be able to create a reply.
 */
BlogPostCommentView.prototype.appendReply = function () {

}

/**
 * Append a reply (comment) to a main comment.
 */
BlogPostCommentView.prototype.appendComment = function (post, user, comment) {

}