/**
 * The properitie manager is a simple listener to manage event...
 */


/**
 * The propertie view contain the code to display propertie
 * of an entity to the end user.
 */
var PropertiesView = function (parent) {

    /** The panel view */
    this.panel = new Element(parent, { "tag": "div", "class": "propertieView" })

    this.propertiesPanel = {}

    server.entityManager.attach(this, OpenEntityEvent, function (evt, propertieView) {
            if (evt.dataMap["entityInfo"] != undefined) {
                propertieView.displayProperties(evt.dataMap["entityInfo"])
            }
        })

    // Attach the file close event.
    server.entityManager.attach(this, CloseEntityEvent,function (evt, propertieView) {
            if (evt.dataMap["entityId"] != undefined) {

            }
        })

    return this
}

PropertiesView.prototype.displayProperties = function (entity) {
    if (this.propertiesPanel[entity.UUID] == undefined) {
        // Here I will display the propertie panel...
        var entityView = new EntityPanel(this.panel, entity.TYPENAME, function (entity, propertiesView) {
            return function (panel) {
                
                // First I will hide all other panel
                for (var panelId in propertiesView.propertiesPanel) {
                    propertiesView.propertiesPanel[panelId].hide()
                }

                // Here I will set the entity
                panel.setTitle(entity.TYPENAME)
                panel.setEntity(entity)
                panel.show()
            }
        } (entity, this))

        // Keep the reference to the propertie view.
        this.propertiesPanel[entity.UUID] = entityView
    } else {
        for (var panelId in this.propertiesPanel) {
            this.propertiesPanel[panelId].hide()
        }
        // Here I will set the entity
        this.propertiesPanel[entity.UUID].setEntity(entity)
        this.propertiesPanel[entity.UUID].show()
    }
}