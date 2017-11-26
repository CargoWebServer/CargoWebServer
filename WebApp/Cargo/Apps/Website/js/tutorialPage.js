/**
 * That page display list of tutorial.
 */
var TutorialPage = function (parent) {
    this.div = new Element(parent, { "tag": "div", "class": "page" })
    this.postPreviews = {}
    this.newestPostDiv = null

    // Events... 
    // Delete post.
    server.entityManager.attach(this, DeleteEntityEvent, function (evt, tutorialPage) {
        if (evt.dataMap["entity"] !== undefined) {
            if (tutorialPage.postPreviews[evt.dataMap["entity"].UUID] != undefined) {
                var postPreview = tutorialPage.postPreviews[evt.dataMap["entity"].UUID]
                delete tutorialPage.postPreviews[postPreview.UUID]
                // Remove it from display.
                tutorialPage.newestPostDiv.removeElement(postPreview.div)
            }
        }
    })

    // New post
    server.entityManager.attach(this, NewEntityEvent, function (evt, tutorialPage) {
        if (evt.dataMap["entity"] !== undefined) {
            if (evt.dataMap["entity"].TYPENAME == blogPostTypeName) {
                if (tutorialPage.postPreviews[evt.dataMap["entity"].UUID] == undefined) {
                    tutorialPage.appendPostPreview(evt.dataMap["entity"], tutorialPage.newestPostDiv)
                }
            }
        }
    })

    // Update post
    server.entityManager.attach(this, UpdateEntityEvent, function (evt, tutorialPage) {
        if (evt.dataMap["entity"] !== undefined) {
            if (tutorialPage.postPreviews[evt.dataMap["entity"].UUID] != undefined) {
                var postPreview = tutorialPage.postPreviews[evt.dataMap["entity"].UUID]
                tutorialPage.appendPostPreview(evt.dataMap["entity"], tutorialPage.newestPostDiv)
            }
        }
    })

    return this
}

TutorialPage.prototype.display = function (parent) {
    parent.removeAllChilds()
    if (this.newestPostDiv == null) {
        parent.element.style.textAlign = "center"
        parent.appendElement({ "tag": "col-xs-12" }).down()
            .appendElement({ "tag": "img", "src": "img/wheel_.svg", "style": "color: black;", "class": "cargo-turning-wheel" })

        // Display the post by newest order.
        this.displayNewestPost(
            function (parent) {
                return function (tutorialPage) {
                    parent.removeAllChilds()
                    parent.element.style.textAlign = ""
                    parent.appendElement(tutorialPage.div)
                }
            }(parent)
        )

    } else {
        parent.appendElement(this.div)
    }
}
/**
 * Display the newest post.
 */
TutorialPage.prototype.displayNewestPost = function (callback) {

    this.newestPostDiv = this.div.appendElement({ "tag": "div", "style": "display: table;" }).down()
        .appendElement({ "tag": "div", "style": "display: table-row;" }).down()
        .appendElement({ "tag": "h3", "style": "display: table-cell;", "innerHtml": "Newest" }).up()
        .appendElement({ "tag": "div", "style": "display: table-row;" }).down()
        .appendElement({ "tag": "div", "style": "display: table-cell;" }).down()
        .appendElement({ "tag": "div", "class": "post_preview_container" }).down()

    // Display the 8 last articles...
    server.entityManager.getEntities("Blog.blog_post", "Blog", null, 0, 8, ["M_date_published"], false,
        function (index, total, caller) {

        },
        function (results, caller) {
            for (var i = 0; i < results.length; i++) {
                caller.tutorialPage.appendPostPreview(results[i], caller.div)
            }
            caller.callback(caller.tutorialPage)
        },
        function (errObj, caller) {

        },
        { "tutorialPage": this, "div": this.newestPostDiv, "callback": callback })
}

TutorialPage.prototype.displayMostViewedPost = function () {

    this.newestPostDiv = this.div.appendElement({ "tag": "div", "style": "display: table;" }).down()
        .appendElement({ "tag": "div", "style": "display: table-row;" }).down()
        .appendElement({ "tag": "h3", "style": "display: table-cell;", "innerHtml": "Most view" }).up()
        .appendElement({ "tag": "div", "style": "display: table-row;" }).down()
        .appendElement({ "tag": "div", "style": "display: table-cell;" }).down()
        .appendElement({ "tag": "div", "class": "post_preview_container" }).down()

    // Display the 8 last articles...
    server.entityManager.getEntities("Blog.blog_post", "Blog", null, 0, 8, ["M_date_published"], false,
        function (index, total, caller) {

        },
        function (results, caller) {
            for (var i = 0; i < results.length; i++) {
                caller.tutorialPage.appendPostPreview(results[i], caller.div)
            }
        },
        function (errObj, caller) {

        },
        { "tutorialPage": this, "div": this.newestPostDiv })
}

/**
 * Append a post previews in the home page.
 */
TutorialPage.prototype.appendPostPreview = function (post, div) {
    var postPreview = new PostPreview(div, post)
    this.postPreviews[post.UUID] = postPreview
}

/**
 * Post previews
 */
var PostPreview = function (parent, post) {
    this.post = post
 
    this.div = parent.getChildById(post.UUID + "_post_preview")
    if (this.div == undefined) {
        this.div = parent.appendElement({ "tag": "div", "class": "post_preview", "id": post.UUID + "_post_preview" }).down()
        this.title = this.div.appendElement({ "tag": "span", "id": post.UUID + "_post_preview_title", "class": "post_title", "innerHtml": post.M_title }).down()
        this.img = this.div.appendElement({ "tag": "img", "id": post.UUID + "_post_preview_img", "src": post.M_thumbnail }).down()

    } else {
        this.title = parent.getChildById(post.UUID + "_post_preview_title")
        this.img = parent.getChildById(post.UUID + "_post_preview_img")
        this.title.element.innerHTML = post.M_title
        this.img.element.src = post.M_thumbnail
    }

    // Here I will generate the post thumbnail from the content of the post
    this.img.element.onload = function (img, div) {
        return function () {
            // I will set the with of the image...
            if (div.element.clientHeight > 0 && div.element.clientWidth > 0) {
                img.element.style.height = div.element.clientHeight + "px";
                img.element.style.width = div.element.clientWidth + "px";
            }
        }
    }(this.img, this.div)


    this.div.element.onclick = function () {
        return function () {
            mainPage.displayPost(post)
        }
    }(post)

    return this
}