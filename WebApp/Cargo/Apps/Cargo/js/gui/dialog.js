
/*
 * A dialog whit button, modal or not...
 */
var Dialog = function (id, parent, isModal, title) {

    this.parent = parent
    if (parent == undefined) {
        this.parent = document.getElementsByTagName("body")[0]
    }

    if (isModal) {
        var body = document.getElementsByTagName("body")[0]
        var html = document.getElementsByTagName("html")[0]
        var height = Math.max(body.scrollHeight, body.offsetHeight, html.clientHeight, html.scrollHeight, html.offsetHeight);
        this.modalDiv = new Element(document.getElementsByTagName("body")[0], { "tag": "div", "style": "position: absolute; top:0px; left:0px;width:100%; height:" + height + "px; background-color:rgba(0,0,0,.25); z-index:1000; display:none;" })
        this.parent = this.modalDiv
        this.isModal = isModal
        this.modalDiv.element.style.display = "block"
    }

    /* The dialogue id **/
    this.id = id
    if (this.id == "") {
        this.id = randomUUID()
    }

    /* The dialog div **/
    this.div = new Element(this.parent, { "tag": "div", "class": "dialog modal-content", "style": "z-index:1001;" })

    /* The header **/
    this.header = this.div.appendElement({ "tag": "div", "class": "dialog_header modal-header" }).down()

    /* The dialog title **/
    this.title = this.header.appendElement({ "tag": "div", "class": "dialog_title modal-title", "innerHtml": title }).down()


    this.deleteBtn = this.header.appendElement({ "tag": "div", "class": "dialog_delete_button close" }).down()
        .appendElement({ "tag": "i", "class": "fa fa-close" }).down()

    /* The content will contain text, input etc... */
    this.content = this.div.appendElement({ "tag": "div", "class": "dialog_content modal-body" }).down()

    /* The footer contain the button like ok, cancel etc...**/
    this.footer = this.div.appendElement({ "tag": "div", "class": "dialog_footer modal-footer" }).down()
    this.footerButtons = this.footer.appendElement({ "tag": "div", "class": "dialog_buttons", "style": "right: 2px;" }).down()

    /* The ok button **/
    this.ok = this.footerButtons.appendElement({ "tag": "div", "class": "diablog_button btn btn-default", "innerHtml": "ok" }).down()

    /* The cancel button **/
    this.cancel = this.footerButtons.appendElement({ "tag": "div", "class": "diablog_button btn btn-default", "innerHtml": "cancel" }).down()

    /* Property used to move the dialog **/
    this.isMoving = false
    this.offsetX = 0
    this.offsetY = 0

    this.header.element.onmousedown = function (dialog) {
        return function (evt) {
            if (evt.which == 1) {
                dialog.isMoving = true
                dialog.offsetX = evt.offsetX
                dialog.offsetY = evt.offsetY
                dialog.div.element.parentNode.className = "unselectable"
                dialog.div.element.style.cursor = "move"
            }
        }
    } (this)

    this.div.element.parentNode.onmouseup = this.header.element.onmouseup = function (dialog) {
        return function (evt) {
            if (evt.which == 1) {
                dialog.isMoving = false
                dialog.offsetX = 0
                dialog.offsetY = 0
                dialog.div.element.style.cursor = ""
                if (dialog.div.element.parentNode != null) {
                    dialog.div.element.parentNode.className = ""
                }
            }
        }
    } (this)

    this.mouseMoveListener = function (dialog) {
        return function (evt) {
            if (dialog.isMoving == true) {

                var rect = dialog.parent.element.getBoundingClientRect();
                var body = document.body;
                var docEl = document.documentElement;

                var scrollTop = window.pageYOffset || docEl.scrollTop || body.scrollTop;
                var scrollLeft = window.pageXOffset || docEl.scrollLeft || body.scrollLeft;

                // Here i will calculate the new position of the dialogue...
                var x = evt.pageX
                var y = evt.pageY

                // Now the new postion of the dialog...
                dialog.x = x - dialog.offsetX - rect.left - scrollLeft
                dialog.y = y - dialog.offsetY - rect.top - scrollTop
                dialog.div.element.style.left = dialog.x + "px"
                dialog.div.element.style.top = dialog.y + "px"
            }
        }
    } (this)

    this.div.element.parentNode.addEventListener("mousemove", this.mouseMoveListener)

    /* The button action **/
    this.cancel.element.onclick = this.deleteBtn.element.onclick = function (dialog) {
        return function (evt) {
            evt.stopPropagation()
            dialog.parent.element.removeChild(dialog.div.element)
            dialog.parent.element.removeEventListener("mousemove", dialog.mouseMoveListener)
            if (dialog.isModal) {
                dialog.modalDiv.element.parentNode.removeChild(dialog.modalDiv.element)
            }
        }
    } (this)

    return this
}

Dialog.prototype.close = function () {
    this.parent.element.removeChild(this.div.element)
    this.parent.element.removeEventListener("mousemove", this.mouseMoveListener)
    if (this.isModal) {
        this.modalDiv.element.parentNode.removeChild(this.modalDiv.element)
    }
}

Dialog.prototype.setCentered = function () {
    var docEl = document.documentElement;
    var body = document.body;
    var scrollTop = window.pageYOffset || docEl.scrollTop || body.scrollTop;
    var scrollLeft = window.pageXOffset || docEl.scrollLeft || body.scrollLeft;

    /* I will set the position of the dialog **/
    this.x = (this.parent.element.offsetWidth - scrollLeft - this.div.element.offsetWidth ) / 2 + scrollLeft
    this.div.element.style.left = this.x + "px"

    this.y = (this.parent.element.offsetHeight - scrollTop - this.div.element.offsetHeight ) / 2 + scrollTop
    this.div.element.style.top = this.y  + "px"
}

Dialog.prototype.setPosition = function (x, y) {
    /* I will set the position of the dialog **/
    this.x = x
    this.div.element.style.left = this.x + "px"

    this.y = y
    this.div.element.style.top = this.y + "px"
}

Dialog.prototype.appendButton = function (id, title, icon) {
    var button = this.footerButtons.prependElement({ "tag": "div", "class": "diablog_button", "id": id, "innerHtml": title }).down()
    if (icon != null) {
        button.prependElement({ "tag": "div", "class": icon })
    }
    return button
}