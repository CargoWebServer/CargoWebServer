/**
 * That will contain page of the website.
 * @param {*} parent 
 */
var HomePage = function (parent) {
    this.div = new Element(parent, { "tag": "div", "class": "page" })

    // The webpages.
    this.webPage = {}
    
    return this
}