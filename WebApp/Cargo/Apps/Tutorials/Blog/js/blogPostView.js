/**
 * There is the blog the blog view.
 */
var BlogPostView = function (parent, post, categoryContentDiv) {

    // Reset the parent content here.
    parent.removeAllChilds()
    parent.element.innerHTML = ""

    // A reference to the post.
    this.post = post

    // The list of category for that post.
    this.categoryContentDiv = categoryContentDiv
    this.categoryContentDiv.element.parentNode.style.display = ""
    
    // I will uncheck all categories
    //this.categoryContentDiv.removeAllChilds()
    // the category div contain tow row...
    var categoryLnks = this.categoryContentDiv.getChildsByClassName("category_lnk")
    var inputs = {}
    for(i=0; i < categoryLnks.length; i++){
        var input = categoryLnks[i].element.childNodes[2]
        input.style.display=""
        input.checked = false

        // Now i will set the action if the input box state change...
        input.onchange = function(post){
            return function(){
                // So here I will append the category of remove it from the post...
                var categoryUuid = this.id.replace("_checkbox", "")
                if(this.checked){
                    // append a new category to the post
                    //console.log("append category ", categoryUuid, "to post ", post)
                    post.M_fk_posts_to_categories_1.push(categoryUuid)
                    server.entityManager.saveEntity(post, function(){}, function(){}, {})
                }else{
                    // remove existing category from the post
                }
            }
        }(this.post)

        // Keep track of input box...
        inputs[input.id] = input
    }

    // Now I will set the categories related to this post...
    for (var i = 0; i < post.M_fk_posts_to_categories_1.length; i++) {
        this.appendCategory(post.M_fk_posts_to_categories_1[i])
        inputs[post.M_fk_posts_to_categories_1[i] + "_checkbox"].checked = true
    }

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

    // The comment editor.
    this.commentEditor = new BlogPostCommentEditor(this, this.pageContainer.up())

    // The comment section.
    this.pageContainer.up().appendElement({ "tag": "hr" })
        .appendElement({ "tag": "div", "id": "comments-container" })

    // The comment container.
    this.commentContainer = this.pageContainer.up().getChildById("comments-container")
    this.commentContainer.up().appendElement({ "tag": "hr" })

    ////////////////////////////////////////////////////////////////////////
    // Blog sections.
    ////////////////////////////////////////////////////////////////////////

    // The title div
    this.titleDiv = this.pageContainer.getChildById("blog-post-title")
    this.titleDiv.element.innerHTML = this.post.M_title

    // The author div
    this.authorDiv = this.pageContainer.getChildById("author-name")

    // The entity uuid.
    server.entityManager.getEntityByUuid(this.post.M_blog_author_id,
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

    // Format the date time to local time.
    var m = moment(this.post.M_date_published);
    this.createdDiv.element.innerHTML = " " + m.format('LLL') // Returns "February 8 2013 8:30 AM" on en-us

    // The page content.
    this.pageContentDiv = this.pageContainer.getChildById("page-content")
    this.pageContentDiv.element.innerHTML = this.post.M_article

    // I will display the comment.
    for (var i = 0; i < post.M_FK_blog_comment_blog_post.length; i++) {
        var comment = post.M_FK_blog_comment_blog_post[i]
        if (isString(comment.M_FK_blog_comment_blog_user)) {
            comment["set_M_FK_blog_comment_blog_user_" + comment.M_FK_blog_comment_blog_user + "_ref"](
                function (blogPostView, comment, post) {
                    return function (ref) {
                        new BlogPostCommentView(blogPostView.commentContainer, ref, comment, post, comment.M_id)
                    }
                }(this, comment, post)
            )
        } else {
            new BlogPostCommentView(this.commentContainer, comment.M_FK_blog_comment_blog_user, comment, post, comment.M_id)
        }
    }

    ////////////////////////////////////////////////////////////////////////
    // Post view actions.
    ////////////////////////////////////////////////////////////////////////

    ////////////////////////////////////////////////////////////////////////
    // Post view listener.
    ////////////////////////////////////////////////////////////////////////


    return this
}

/**
 * Append a new category in the category div.
 */
BlogPostView.prototype.appendCategory = function (category) {

}

/**
 * Append a blog comment.
 */
BlogPostView.prototype.appendComment = function (post, user, comment) {
    // First of all I will test if the user exist in the blog_user table.
    server.entityManager.getEntityById(userTypeName, "sql_info", [user.M_id],
        function (result, caller) {
            // update entity here.
            server.entityManager.saveEntity(user)
            // Here the user exist.
            new BlogPostCommentView(caller.blogPostView.commentContainer, result, caller.comment, caller.post)
        },
        function (errObj, caller) {
            var user = caller.user
            // The user dosent exist yet...
            server.entityManager.saveEntity(user)
            new BlogPostCommentView(caller.blogPostView.commentContainer, caller.user, caller.comment, caller.post)
        }, { "user": user, "comment": comment, "blogPostView": this, "post": post })
}
