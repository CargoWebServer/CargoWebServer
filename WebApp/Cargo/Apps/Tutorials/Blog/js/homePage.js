/**
 * General page is the 
 */
var HomePage = function (parent) {
    this.div = parent.appendElement({ "tag": "div", "class": "home_page" }).down()

    server.entityManager.getObjectsByType("Blog.blog_post", "Blog", "",
        function(index, total, caller){

        },
        function (results, caller) {
            console.log("--------> hey! ", results)
        },
        function (errObj, caller) {

        },
        { "homePage": this })
}