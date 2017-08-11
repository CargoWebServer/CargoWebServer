/**
 * General page is the 
 */
var HomePage = function (parent) {
    this.div = parent.appendElement({ "tag": "div", "class": "home_page" }).down()
    this.postPreviews = {}

    server.entityManager.getObjectsByType("Blog.blog_post", "Blog", "",
        function (index, total, caller) {

        },
        function (results, caller) {
            for (var i = 0; i < results.length; i++) {
                caller.homePage.appendPostPreview(results[i])
            }
        },
        function (errObj, caller) {

        },
        { "homePage": this })
}

/**
 * Append a post previews in the home page.
 */
HomePage.prototype.appendPostPreview = function (post) {
    var postPreview = new PostPreview(this.div, post)
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
            //img.element.style.height = div.element.clientHeight + "px";
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