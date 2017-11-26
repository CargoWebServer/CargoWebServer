/**
 * That page display list of tutorial.
 */
var TutorialPage = function (parent) {
    this.div = new Element(parent, { "tag": "div", "class": "page" })
    this.postPreviews = {}
    this.newestPostDiv = null

    return this
}

TutorialPage.prototype.display = function (parent) {
    mainPage.pageContent.removeAllChilds()
    if (this.newestPostDiv == null) {
        parent.element.style.textAlign = "center"
        parent.appendElement({"tag":"col-xs-12"}).down()
            .appendElement({"tag":"img", "src":"img/wheel_.svg", "style":"color: black;", "class":"cargo-turning-wheel"})
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
    this.div = parent.appendElement({ "tag": "div", "class": "post_preview" }).down()
    this.div.appendElement({ "tag": "span", "class": "post_title", "innerHtml": post.M_title })

    // Here I will generate the post thumbnail from the content of the post
    var img = this.div.appendElement({ "tag": "img", "src": post.M_thumbnail }).down()

    img.element.onload = function (img, div) {
        return function () {
            // I will set the with of the image...
            img.element.style.height = div.element.clientHeight + "px";
            img.element.style.width = div.element.clientWidth + "px";
        }
    }(img, this.div)


    this.div.element.onclick = function () {
        return function () {
            mainPage.displayPost(post)
        }
    }(post)

    return this
}