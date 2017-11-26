/**
 * Display website page.
 * 
 * @param {*} parent 
 */
var HomePage = function (parent) {
    this.div = new Element(parent, { "tag": "div", "class": "page" })

    
    // The webpages.
    this.webPage = {}

    return this
}

/**
 * Display the page.
 */
HomePage.prototype.display = function(parent){
    parent.removeAllChilds()
    parent.appendElement(this.div)
    
}