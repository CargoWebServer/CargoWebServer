/**
 * That function is use to navigage from page to page.
 */
var PageSelector = function (parent, search, max) {
    this.max = max;
    if (this.max == undefined) {
        this.max = 10;
    }

    // The seach function whit prototype function(offest,  pageSize)
    this.search = search;

    // The parent where the panel will be display.
    this.parent = parent;

    this.panel = parent.appendElement({ "tag": "div", "class": "page-selector-panel" }).down();

    return this;
}

PageSelector.prototype.setResults = function (index, total, pagesize) {

    // Here I will simply erease the buttons and recreate it...
    this.panel.removeAllChilds();

    // The current page number.
    var pageNumber = parseInt(index / pagesize);

    // The total number of page.
    var numberPages = parseInt(total / pagesize + .5);

    // The start index.
    var startIndex = parseInt(pageNumber / this.max);

    for (var i = 0; i < this.max; i++) {
        if (i + (startIndex * this.max) < numberPages) {
            // if all page are already display.
            var selector = this.panel.appendElement({ "tag": "div", "innerHtml": ((startIndex * this.max) + i + 1).toString() }).down()

            // If the page is active...
            if ((startIndex * this.max) + i == pageNumber) {
                selector.element.className += " active";
            }

            // Now I will set the action...
            selector.element.onclick = function (search, index, pagesize) {
                return function () {
                    search(index, pagesize)
                }
            }(this.search, (startIndex * this.max * pagesize) + (i * pagesize), pagesize)
        }
    }

    // Now I will append the button to navigate from batch of page...

    // Move Foward...
    if (parseInt(numberPages / this.max + .5) > 1 && startIndex < parseInt(numberPages / this.max + .5) - 1) {
        var moveNext = this.panel.appendElement({ "tag": "div" }).down()
        moveNext.appendElement({ "tag": "i", "class": "fa fa-angle-right" }).down()
        var nextIndex = ((startIndex + 1) * this.max * pagesize)
        moveNext.element.onclick = function (search, index, pagesize) {
            return function () {
                search(index, pagesize)
            }
        }(this.search, nextIndex, pagesize)

        if (parseInt(numberPages / this.max + .5) > 2) {
            var moveLast = this.panel.appendElement({ "tag": "div" }).down()
            moveLast.appendElement({ "tag": "i", "class": "fa fa-angle-double-right" }).down()
            moveLast.element.onclick = function (search, index, pagesize) {
                return function () {
                    search(index, pagesize)
                }
            }(this.search, (numberPages - 1) * pagesize, pagesize)
        }
    }

    // Move backward...
    if (startIndex > 0) {
        var movePrevious = this.panel.prependElement({ "tag": "div" }).down()
        movePrevious.appendElement({ "tag": "i", "class": "fa fa-angle-left" }).down()
        var previousIndex = ((startIndex - 1) * this.max * pagesize)
        movePrevious.element.onclick = function (search, index, pagesize) {
            return function () {
                search(index, pagesize)
            }
        }(this.search, previousIndex, pagesize)
    }

    if (startIndex > 1) {
        var movePrevious = this.panel.prependElement({ "tag": "div" }).down()
        movePrevious.appendElement({ "tag": "i", "class": "fa fa-angle-double-left" }).down()
        movePrevious.element.onclick = function (search, index, pagesize) {
            return function () {
                search(index, pagesize)
            }
        }(this.search, 0, pagesize)
    }
}