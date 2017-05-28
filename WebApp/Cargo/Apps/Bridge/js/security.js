/**
 * Role and permission are use to manage access to ressources.
 */
var RolePermissionManager = function (parent) {

    // Keep reference to the parent.
    this.parent = parent

    // The interface panel.
    this.panel = parent.appendElement({ "tag": "div", "class": "role_permission_manager" }).down()

    // So here I will create a tab panel...
    this.content = this.panel.appendElement({ "tag": "div", "style": "display:table;" }).down()
        .appendElement({ "tag": "div", "id": "role_tab", "class": "file_tab active", "style": "display: table-cell;", "innerHtml": "Roles" })
        .appendElement({ "tag": "div", "id": "permission_tab", "class": "file_tab", "style": "display: table-cell;", "innerHtml": "Permissions" })
        .appendElement({ "tag": "div", "style": "display: table-cell; width: 100%; background-color: #f0f0f0;" }).up()
        .appendElement({"tag":"div", "style":"display:table; width:100%; height: 100%; border-top: 1px splid grey;"}).down()

    this.roleTab = this.panel.getChildById("role_tab")
    this.permissionTab = this.panel.getChildById("permission_tab")

    // The roles...
    this.roleManager = new RoleManager(this.content)

    // The permissions...
    this.permissionManager = new PermissionManager(this.content)

    // Connect the actions
    this.roleTab.element.onclick = function (permissionTab, roleManager, permissionManager) {
        return function () {
            permissionTab.element.className = "file_tab"
            this.className = "file_tab active"
            roleManager.panel.element.style.display = ""
            permissionManager.panel.element.style.display = "none"
        }
    }(this.permissionTab, this.roleManager, this.permissionManager)

    this.permissionTab.element.onclick = function (roleTab, roleManager, permissionManager) {
        return function () {
            roleTab.element.className = "file_tab"
            this.className = "file_tab active"
            roleManager.panel.element.style.display = "none"
            permissionManager.panel.element.style.display = ""
        }
    }(this.roleTab, this.roleManager, this.permissionManager)



    return this
}

/**
 * The role manager is use to create, modify or delete roles.
 */
var RoleManager = function (parent) {
    // Keep reference to the parent.
    this.parent = parent

    // The interface panel.
    this.panel = parent.appendElement({ "tag": "div", "class": "role_manager",  }).down()

    // Return the role manager.
    return this
}

/**
 * The permission manager is use to change permission of an entity.
 */
var PermissionManager = function (parent) {
    // Keep reference to the parent.
    this.parent = parent

    // The interface panel.
    this.panel = parent.appendElement({ "tag": "div", "class": "permission_manager", "style":"display:none;"}).down()

    // Return the permission manager.
    return this
}