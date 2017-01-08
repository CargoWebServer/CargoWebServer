/*
 * The menu item contain the information to be use
 * by the menu interface...
 */
var MenuItem = function (id, name, subItems, level, action, icon) {

    // Must be unique in it context...
    this.id = id

    // The name to display in the menu
    this.name = name

    // The submenu item...
    this.subItems = subItems

    // This is the panel associated whit the button...
    this.panel = null

    // The action assciated with the panel click...
    this.action = action

    // The level of the element...
    this.level = level

    // The parent of the item...
    this.parent = null

    // The font-awsone icon name
    this.icon = ""
    if (icon != undefined) {
        this.icon = icon
    }

    for (var key in subItems) {
        subItems[key].parent = this
    }

    return this
}

/*
 * That class contain the code for vertical menu.
 */
var VerticalMenu = function (parent, items) {

    this.parent = parent
    this.panel = this.parent.appendElement({ "tag": "div", "style":"relative;"}).down()
    this.subItemPanel = null
    this.currentItem = null

    this.items = items
    if (this.items == undefined) {
        this.items = []
    }

    // I will now initialyse the items...
    for (var i = 0; i < this.items.length; i++) {
        this.appendItem(this.items[i])
    }

    return this
}

/*
 * Append a new Item in the menu panel...
 */
VerticalMenu.prototype.appendItem = function (item) {

    var currentPanel = null
    if (item.level == 0) {
        // Here the menu is show in row...
        currentPanel = this.panel.appendElement({ "tag": "div", "id": item.id, "class": "vertical_menu", "innerHtml": item.name }).down()
    } else {
        // Here the menu is 
        currentPanel = this.panel.appendElement({ "tag": "div", "class": "menu_row", "style": "display: table-row; width: 100%" }).down()

        var iconPanel = currentPanel.appendElement({ "tag": "i", "class": item.icon + " menu_icon", "style": "display: table-cell;" }).down()
        var subMenuPanel = currentPanel.appendElement({ "tag": "div", "id": item.id, "class": "vertical_submenu", "innerHtml": item.name }).down()

        iconPanel.element.onmouseenter = function (subMenuPanel) {
            return function () {
                subMenuPanel.element.style.border = "1px solid lightgrey"
            }
        } (subMenuPanel)

        iconPanel.element.onmouseleave = function (subMenuPanel) {
            return function () {
                subMenuPanel.element.style.border = ""
            }
        } (subMenuPanel)

        subMenuPanel.element.onmouseenter = function (iconPanel) {
            return function () {
                iconPanel.element.style.color = "white"
                iconPanel.element.style.backgroundColor = "#657383"
            }
        } (iconPanel)

        subMenuPanel.element.onmouseleave = function (iconPanel) {
            return function () {
                iconPanel.element.style.color = ""
                iconPanel.element.style.backgroundColor = ""
            }
        } (iconPanel)
    }

    // Append the subitem panel.
    this.subItemPanel = currentPanel.appendElement({ "tag": "div", "id": item.id, "class": "vertical_submenu_items" }).down()
    item.panel = this.subItemPanel

    // Display the menu automatically...
    currentPanel.element.onmouseenter = function (menuPanel, subItemPanel, item, menu) {
        return function () {
            var subItemPanels = document.getElementsByClassName("vertical_submenu_items")
            for (var i = 0; i < subItemPanels.length; i++) {
                subItemPanels[i].style.display = "none"
            }

            if (item.level > 0) {
                // Now I will offset the menu...
                menu.currentItem = item
                function setVisible(item) {

                    item.panel.element.style.display = "block"
                    if (item.parent != undefined) {
                        setVisible(item.parent)
                    }
                    if (item.level > 0) {
                        item.panel.element.style.left = item.panel.element.parentNode.offsetWidth + "px"
                        item.panel.element.style.top = "0px"
                    }
                    if (Object.keys(item.panel.childs).length == 0) {
                        item.panel.element.style.display = "none"
                    }
                }
                // Set the item menu visibility...
                setVisible(item)
            }
        }
    } (currentPanel, this.subItemPanel, item, this)

    // Create sub items...
    var subItems = []
    if (item.subItems != undefined) {
        for (var key in item.subItems) {
            subItems.push(item.subItems[key])
        }
        if (subItems.length > 0) {
            new VerticalMenu(this.subItemPanel, subItems)
        }
    }

    // Now the actions...
    currentPanel.element.onclick = function (menuPanel, subItemPanel) {
        return function (evt) {
            evt.stopPropagation()

            if (subItemPanel.element.style.display == "block") {
                subItemPanel.element.style.display = "none"
            } else {
                subItemPanel.element.style.display = "block"
                if (item.level >= 1) {
                    // Now I will offset the menu...
                    subItemPanel.element.style.left = menuPanel.element.offsetWidth + "px"
                    subItemPanel.element.style.top = "0px"
                }
            }
            if (Object.keys(subItemPanel.childs).length == 0) {
                subItemPanel.element.style.display = "none"
            }
            if (item.action != undefined) {
                item.action()
                function setInvisible(item) {
                    item.panel.element.style.display = "none"
                    if (item.parent != undefined) {
                        setInvisible(item.parent)
                    }
                }
                setInvisible(item)
            }
        }
    } (currentPanel, this.subItemPanel, item)

}

/*
 * That class contain de the code for horizontal menu
 */
var HorizontalMenu = function (parent, items) {
    this.items = items
    this.currentItem = null

    // The menu panel.
    this.panel = new Element(parent, { "tag": "div", "class": "horizontal_menu", "style": "display: table;width:100%" })

    // The menu content...
    this.menuRow = this.panel.appendElement({ "tag": "div", "style": "display:table-row;height:60px;" }).down()

    // That panel will contain the sub menu...
    this.subMenuRow = this.panel.appendElement({ "tag": "div", "style": "display:table-row;height:0px; background-color:#657383; color:lightgrey;" }).down()

    // I will now initialyse the items...
    for (var i = 0; i < this.items.length; i++) {
        this.appendItem(this.items[i])
    }

    return this
}

/*
 * Append a new Item in the menu panel...
 */
HorizontalMenu.prototype.appendItem = function (item) {

    // The top element...
    if (item.level == 0) {
        if (item.panel == null) {
            item.panel = this.menuRow.appendElement({ "tag": "div", "class": "menu_item", "id": item.id, "style": "display:inline-block; padding-left:10px; padding-top: 20px; vertical-align:middle; height:40px;", "innerHtml": item.name }).down()
        }
    } else {
        if (item.parent.panel.childs[item.id] == undefined) {
            item.panel = item.parent.panel.appendElement({ "tag": "div", "id": item.id, "style": "display:inline-block;border-left:1px solid grey; margin-left:4px; padding-left:4px; height:auto;", "innerHtml": item.name }).down()
        }
    }

    item.panel.level = item.level

    // Mouse action...
    item.panel.element.onmouseover = function (menu) {
        return function () {
            this.style.cursor = "pointer"
        }
    } (this)

    item.panel.element.onmouseleave = function (menu) {
        return function () {
            this.style.cursor = "default"
        }
    } (this)

    item.panel.element.onclick = function (menu, item_) {
        return function (evt) {
            evt.stopPropagation()
            if (item_.level == 0) {
                var items = document.getElementsByClassName("menu_item")
                for (var i = 0; i < items.length; i++) {
                    items[i].style.color = ""
                }
                for (var i = 0; i < menu.items.length; i++) {
                    if (menu.items[i].id != item_.id) {
                        // Clear it childs...
                        menu.items[i].panel.element.innerHTML = menu.items[i].name
                    }
                }
            }

            this.style.color = "#428bca";
            if (Object.keys(menu.subMenuRow.childs).length > 0 && item_.level == 0 && Object.keys(item_.panel.childs).length == 0) {
                menu.subMenuRow.removeAllChilds()
                item.panel.element.innerHTML = item_.name
                this.style.color = "";
                menu.appendItem(item)
            } else {
                menu.setItem(item_)
            }
            if (item_.action != null) {
                // Call the action...
                item_.action()
            }
        }
    } (this, item)
}

HorizontalMenu.prototype.setItem = function (item) {

    if (Object.keys(item.subItems).length > 0) {
        item.panel.removeAllChilds()
        if (item.level > 0) {
            function setParentPanel(parent, menu) {
                if (parent.parent != null) {
                    setParentPanel(parent.parent, menu)
                }
                menu.appendItem(parent)
            }
            setParentPanel(item, this)
        } else {
            this.appendItem(item)
        }

        this.subMenuRow.removeAllChilds()
        this.currentItem = item
        for (var itemId in this.currentItem.subItems) {
            var subItem = this.currentItem.subItems[itemId]
            subItem.parent = item
            subItem.panel = this.subMenuRow.appendElement({ "tag": "div", "class": "submenu_item", "style": "display:inline-block; padding-left:10px;padding-top: 1px; vertical-align:middle;height:20px; ", "innerHtml": subItem.name }).down()
            subItem.panel.element.onclick = function (menu, item) {
                return function () {
                    menu.setItem(item)
                    var subitems = document.getElementsByClassName("submenu_item")
                    for (var i = 0; i < subitems.length; i++) {
                        subitems[i].style.color = ""
                    }
                    if (this.style.color == "") {
                        this.style.color = "white";
                    } else {
                        this.style.color == ""
                    }
                    if (item.action != null) {
                        // Call the action...
                        item.action()
                    }
                }
            } (this, subItem)

            // Mouse action...
            subItem.panel.element.onmouseover = function (menu) {
                return function () {
                    this.style.cursor = "pointer"
                }
            } (this)

            subItem.panel.element.onmouseleave = function (menu) {
                return function () {
                    this.style.cursor = "default"
                }
            } (this)

        }

        // Set the menu size...
        this.menuRow.element.style.height = "40px"
        this.subMenuRow.element.style.height = "20px"
    }
}