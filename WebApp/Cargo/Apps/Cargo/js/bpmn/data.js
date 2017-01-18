// Contain various object constructor function for bpmn data
BpmnsData = {}

/**
 * That class is use to create bpmn data view.
 */
var BpmnDataView = function (parent, bpmnElement, initCallback) {

    // The parent bpmn element that contain the data.
    this.bpmnElement = bpmnElement

    this.parent = parent

    // I will use an entity panel to display
    this.panel = this.parent.appendElement({ "tag": "div" }).down()

    // Bpmn data
    this.dataOutputsPanel = null
    this.dataOutputs = this.bpmnElement.M_dataOutput
    this.dataInputsPanel = null
    this.dataInputs = this.bpmnElement.M_dataInput

    // The function to call when the initialisation is done.
    this.initCallback = initCallback


    this.init()

    return this
}

BpmnDataView.prototype.init = function () {
    // first of all I will generate an entity prototype
    // from the data array.
    if (this.dataInputs != undefined) {
        var typeName = this.bpmnElement.UUID + "_dataInput"
    }

    if (this.dataOutputs != undefined) {

        var typeName = "BpmnsData.DataOutput_" + this.bpmnElement.UUID.split("%")[1].replaceAll("-", "_")

        this.generateEntityPrototype(this.dataOutputs, typeName,
            // The callback...
            function (dataView, dataOutputsPanel) {
                return function (prototype) {
                    // Generate the entity constructor.
                    prototype.generateConstructor()
                    prototype.generateInit()
                    server.entityManager.entityPrototypes[prototype.TypeName] = prototype

                    // So here with the entity prototype i will generate the entity panel.
                    dataOutputsPanel = new EntityPanel(dataView.panel, prototype.TypeName, function () { }, null, true, null, "")
                    //dataOutputsPanel.header.element.style.display = "none"
                    dataOutputsPanel.maximizeBtn.element.click()

                    // Set an empty entity.
                    var entity = eval("new " + prototype.TypeName + "()")
                    server.entityManager.entities[entity.UUID] = entity

                    entity.onChange = function (dataOutputsPanel) {
                        return function(entity){
                            dataOutputsPanel.setEntity(entity)
                        }
                    }(dataOutputsPanel)

                    dataOutputsPanel.setEntity(entity)

                }
            } (this, this.dataOutputsPanel))
    }
}

/**
 * Generate an entity prototype for a given set of data.
 */
BpmnDataView.prototype.generateEntityPrototype = function (data, typeName, callback) {
    var prototype = new EntityPrototype()
    prototype.TypeName = typeName
    prototype.PackageName = typeName.split(".")[0]
    prototype.ClassName = typeName.split(".")[1]

    for (var i = 0; i < data.length; i++) {
        // I will retreive it item definition...
        console.log(data[i])

        if (isObjectReference(data[i].M_itemSubjectRef)) {
            // Here I will get the item definition...
            server.entityManager.getEntityByUuid(data[i].M_itemSubjectRef,
                function (result, caller) {
                    var done = caller.done
                    var callback = caller.callback
                    var prototype = caller.prototype

                    var fieldType = result.M_structureRef + ":Ref"
                    if (result.M_isCollection) {
                        fieldType = "[]" + fieldType
                    }

                    prototype.FieldsOrder.push(prototype.Fields.length)
                    prototype.Fields.push("M_" +data[i].M_name)
                    prototype.FieldsType.push(fieldType)
                    prototype.FieldsVisibility.push(true)
                    prototype.FieldsNillable.push(true)

                    if (done) {
                        callback(prototype)
                    }
                },
                function (errObj, caller) {

                },
                { "prototype": prototype, "callback": callback, "done": i == data.length - 1 })
        } else {

            // Here I will append the base type...
            var fieldType = data[i].M_itemSubjectRef.replace("xsd:", "xs.")
            if (data[i].M_isCollection) {
                fieldType = "[]" + fieldType
            }
            prototype.FieldsOrder.push(prototype.Fields.length)
            prototype.Fields.push("M_" + data[i].M_name)
            prototype.FieldsType.push(fieldType)
            prototype.FieldsVisibility.push(true)
            prototype.FieldsNillable.push(true)

            if (i == data.length - 1) {
                callback(prototype)
            }
        }
    }
}