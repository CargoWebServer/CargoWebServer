/*
 * The menu item contain the information to be use
 * by the menu interface...
 */
var MenuItem = function (id, name, subItems, level, action, icon) {
    // The parent menu.
    this.menu = null

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

/**
 * Append a sub menu item to an existing menu item.
 */
MenuItem.prototype.appendItem = function (item) {
    item.parent = this
    this.subItems[item.id] = item
    this.menu.appendItem(item)
}

/**
 * Append a sub menu item to an existing menu item.
 */
MenuItem.prototype.deleteItem = function () {
    // remove it from it parent,
    delete this.parent.subItems[this.id]
    var menu = document.getElementById(this.id)
    if (menu != null) {
        menu.parentNode.parentNode.removeChild(menu.parentNode)
    }
}

/**
 * Append a sub menu item to an existing menu item.
 */
MenuItem.prototype.renameItem = function (name, id) {
    delete this.parent.subItems[this.id]
    this.parent.subItems[this.id] = this
    var menu = document.getElementById(this.id)
    if (menu != null) {
        // set the new name.
        menu.firstChild.innerHTML = name
    }
}

/**
 * 
 * @param {*} parent The div where the menu will be append.
 * @param {*} items  The list of menu Item to display in the menu.
 */
var Menu = function (parent, items) {
    if (parent == null) {
        return
    }

    this.parent = parent
    this.panel = this.parent.appendElement({ "tag": "div", "style": "position: relative; display: none;" }).down()
    this.subItemPanel = null
    this.currentItem = null
    this.subMenus = {}

    this.items = items
    if (this.items == undefined) {
        this.items = []
    }
}

/**
 * Append separator between menu items.
 */
Menu.prototype.appendSeparator = function (item) {
    this.panel.appendElement({ "tag": "div", "class": "menu_separator" }).down()
}

/*
 * That class contain the code for vertical menu.
 */
var VerticalMenu = function (parent, items) {

    Menu.call(this, parent, items)

    // I will now initialyse the items...
    for (var i = 0; i < this.items.length; i++) {
        this.appendItem(this.items[i])
    }

    return this
}

VerticalMenu.prototype = new Menu(null);
VerticalMenu.prototype.constructor = Menu;

/*
 * Append a new Item in the menu panel...
 */
VerticalMenu.prototype.appendItem = function (item) {
    if (isString(item)) {
        // In that case I will append a separator.
        this.appendSeparator()
        return
    }

    this.panel.element.style.display = ""

    var currentPanel = null
    if (item.level == 0) {
        // Here the menu is show in row...
        currentPanel = this.panel.appendElement({ "tag": "div", "id": item.id, "class": "vertical_menu" }).down()
        currentPanel.appendElement({"tag":"span", "innerHtml": item.name })
    } else {
        // Here the menu is
        currentPanel = this.panel.appendElement({ "tag": "div", "class": "menu_row", "id": item.id + "_vertical_submenu_row", "style": "display: table-row; width: 100%" }).down()
        currentPanel.appendElement({ "tag": "i", "class": item.icon + " menu_icon", "style": "display: table-cell;" }).down()
        currentPanel.appendElement({ "tag": "div", "id": item.id, "class": "vertical_submenu"}).down()
            .appendElement({"tag":"span", "innerHtml": item.name })

    }

    // Append the subitem panel.
    this.subItemPanel = currentPanel.appendElement({ "tag": "div", "id": item.id + "_vertical_submenu_items", "class": "vertical_submenu_items", "style": "display: none;" }).down()
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
                        item.panel.element.style.top = item.panel.element.parentNode.offsetTop + "px"
                    }
                    if (Object.keys(item.subItems).length == 0) {
                        item.panel.element.style.display = "none"
                    }
                }
                // Set the item menu visibility...
                setVisible(item)
            }
        }
    }(currentPanel, this.subItemPanel, item, this)

    // Create sub items...
    var subItems = []
    if (item.subItems != undefined) {
        for (var key in item.subItems) {
            subItems.push(item.subItems[key])
        }
    }

    item.menu = new VerticalMenu(this.subItemPanel, subItems)

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
                item.action(item.menu)
                function setInvisible(item) {
                    item.panel.element.style.display = "none"
                    if (item.parent != undefined) {
                        setInvisible(item.parent)
                    }
                }
                setInvisible(item)
            }
        }
    }(currentPanel, this.subItemPanel, item)
}

/*
 * That class create a popup menu.
 */
var PopUpMenu = function (parent, items, e) {

    Menu.call(this, parent.getTopParent(), items)

    // Only one menu must be display at any time.
    var popups = document.getElementsByClassName("popup_menu")
    for (var i = 0; i < popups.length; i++) {
        popups[i].parentNode.removeChild(popups[i])
    }

    this.panel.element.className = "popup_menu"
    this.panel.element.style.position = "absolute"
    this.panel.element.style.zIndex = "10"
    var coord = getCoords(parent.element)
    this.panel.element.style.top = coord.top + 16 + "px"
    this.panel.element.style.left = coord.left + 24 + "px"

    this.displayed = false

    // Body events, used to remove the context menu.
    var listener0 = function (popup, listener) {
        return function () {
            if (popup.displayed == true) {
                try {
                    popup.panel.element.parentNode.removeChild(popup.panel.element)
                    document.getElementsByTagName("body")[0].removeEventListener("contextmenu", listener)
                } catch (err) {
                    /** nothing to do here. */
                }
            } else {
                popup.displayed = true
            }
            return false
        }
    }(this, listener0)

    document.getElementsByTagName("body")[0].addEventListener("contextmenu", listener0);

    var listener1 = function (listener) {
        return function (evt) {
            var popups = document.getElementsByClassName("popup_menu")
            for (var i = 0; i < popups.length; i++) {
                popups[i].parentNode.removeChild(popups[i])
            }
            document.getElementsByTagName("body")[0].removeEventListener("click", listener)
        }
    }(listener1)
    document.getElementsByTagName("body")[0].addEventListener("click", listener1);

    var listener2 = function (popup, listener) {
        return function (evt) {
            if (evt.keyCode == 27) {
                try {
                    popup.panel.element.parentNode.removeChild(popup.panel.element)
                    document.getElementsByTagName("body")[0].removeEventListener("keyup", listener)
                } catch (err) {
                    /** nothing to do here. */
                }
            }
        }
    }(this, listener2)
    document.getElementsByTagName("body")[0].addEventListener("keyup", listener2);

    // I will now initialyse the items...
    for (var i = 0; i < this.items.length; i++) {
        this.appendItem(this.items[i])
    }

    return this
}

PopUpMenu.prototype = new Menu(null);
PopUpMenu.prototype.constructor = Menu;

/*
 * Append a new Item in the menu panel...
 */
PopUpMenu.prototype.appendItem = function (item) {
    if (isString(item)) {
        // In that case I will append a separator.
        this.appendSeparator()
        return
    }

    this.panel.element.style.display = ""

    var currentPanel = null

    // Here the menu is 
    currentPanel = this.panel.appendElement({ "tag": "div", "class": "menu_row", "style": "display: table-row; width: 100%" }).down()

    var iconPanel = currentPanel.appendElement({ "tag": "i", "class": item.icon + " menu_icon", "style": "display: table-cell;" }).down()
    var subMenuPanel = currentPanel.appendElement({ "tag": "div", "id": item.id, "class": "vertical_submenu" }).down()
    subMenuPanel.appendElement({"tag":"span", "innerHtml": item.name})

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
    }(currentPanel, this.subItemPanel, item, this)

    // Create sub items...
    var subItems = []
    if (item.subItems != undefined) {
        for (var key in item.subItems) {
            subItems.push(item.subItems[key])
        }
    }

    item.menu = new VerticalMenu(this.subItemPanel, subItems)


    // Now the actions...
    currentPanel.element.onclick = function (menu, menuPanel, subItemPanel) {
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
                item.action(menu)
                menu.panel.element.parentNode.removeChild(menu.panel.element)
                function setInvisible(item) {
                    item.panel.element.style.display = "none"
                    if (item.parent != undefined) {
                        setInvisible(item.parent)
                    }
                }
                setInvisible(item)
            }
        }
    }(this, currentPanel, this.subItemPanel, item)

}

/*
 * That class contain de the code for horizontal menu
 */
var HorizontalMenu = function (parent, items) {

    Menu.call(this, parent, items)

    // Set specific panel properties.
    this.panel.element.className = "horizontal_menu"
    this.panel.element.style.display = "table"
    this.panel.element.style.width = "100%"

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

HorizontalMenu.prototype = new Menu(null);
HorizontalMenu.prototype.constructor = Menu;

/*
 * Append a new Item in the menu panel...
 */
HorizontalMenu.prototype.appendItem = function (item) {

    item.menu = this

    // The top element...
    if (item.level == 0) {
        if (item.panel == null) {
            item.panel = this.menuRow.appendElement({ "tag": "div", "class": "menu_item", "id": item.id, "style": "display:inline-block; padding-left:10px; padding-top: 20px; vertical-align:middle; height:40px;" }).down()
            itmep.panel.appendElement({"tag":"span", "innerHtml": item.name})
        }
    } else {
        if (item.parent.panel.childs[item.id] == undefined) {
            item.panel = item.parent.panel.appendElement({ "tag": "div", "id": item.id, "style": "display:inline-block;border-left:1px solid grey; margin-left:4px; padding-left:4px; height:auto;" }).down()
            itmep.panel.appendElement({"tag":"span", "innerHtml": item.name})
        }
    }

    item.panel.level = item.level

    // Mouse action...
    item.panel.element.onmouseover = function (menu) {
        return function () {
            this.style.cursor = "pointer"
        }
    }(this)

    item.panel.element.onmouseleave = function (menu) {
        return function () {
            this.style.cursor = "default"
        }
    }(this)

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
                        menu.items[i].panel.element.firstChild.innerHTML = menu.items[i].name
                    }
                }
            }

            this.style.color = "#428bca";
            if (Object.keys(menu.subMenuRow.childs).length > 0 && item_.level == 0 && Object.keys(item_.panel.childs).length == 0) {
                menu.subMenuRow.removeAllChilds()
                item.panel.element.firstChild.innerHTML = item_.name
                this.style.color = "";
                menu.appendItem(item)
            } else {
                menu.setItem(item_)
            }
            if (item_.action != null) {
                // Call the action...
                item_.action(menu)
            }
        }
    }(this, item)
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
            subItem.panel = this.subMenuRow.appendElement({ "tag": "div", "class": "submenu_item", "style": "display:inline-block; padding-left:10px;padding-top: 4px; vertical-align:middle;height:20px; ", "innerHtml": subItem.name }).down()
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
                        item.action(menu)
                    }
                }
            }(this, subItem)

            // Mouse action...
            subItem.panel.element.onmouseover = function (menu) {
                return function () {
                    this.style.cursor = "pointer"
                }
            }(this)

            subItem.panel.element.onmouseleave = function (menu) {
                return function () {
                    this.style.cursor = "default"
                }
            }(this)

        }

        // Set the menu size...
        this.menuRow.element.style.height = "40px"
        this.subMenuRow.element.style.height = "20px"
    }
}