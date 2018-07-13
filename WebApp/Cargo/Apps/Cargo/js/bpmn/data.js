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

    // Input data
    this.dataInputsPanel = null
    this.dataInputs = this.bpmnElement.M_dataInput

    // Output data
    this.dataOutputsPanel = null
    this.dataOutputs = this.bpmnElement.M_dataOutput

    // The function to call when the initialisation is done.
    this.initCallback = initCallback

    this.init()

    return this
}

BpmnDataView.prototype.init = function () {

    // first of all I will generate an entity prototype
    // from the data array.
    if (this.dataInputs != undefined) {
        var typeName = "BpmnsData.DataInput_" + this.bpmnElement.UUID.split("%")[1].replaceAll("-", "_")
        this.generateEntityPrototype(this.dataInputs, typeName,
            // The callback...
            function (dataView) {
                return function (prototype) {
                    // Generate the entity constructor.
                    prototype.generateConstructor()
                    setEntityPrototype(prototype)

                    // So here with the entity prototype i will generate the entity panel.
                    dataView.dataInputsPanel = new EntityPanel(dataView.panel, prototype.TypeName, function (prototype) { 
                        return function(panel){
                            // Set an empty entity.
                            var entity = eval("new " + prototype.TypeName + "()")
                            entity.UUID = typeName  + "%" + randomUUID()
                            panel.setEntity(entity)
                        }
                    }(prototype), null, true, null, "")
                }
            } (this))
    }

    if (this.dataOutputs != undefined) {
        var typeName = "BpmnsData.DataOutput_" + this.bpmnElement.UUID.split("%")[1].replaceAll("-", "_")
        this.generateEntityPrototype(this.dataOutputs, typeName,
            // The callback...
            function (dataView) {
                return function (prototype) {
                    // Generate the entity constructor.
                    prototype.generateConstructor()
                    setEntityPrototype(prototype)
                    
                    // So here with the entity prototype i will generate the entity panel.
                    dataView.dataOutputsPanel = new EntityPanel(dataView.panel, prototype.TypeName, function (prototype) { 
                        return function(panel){
                            // Set an empty entity.
                            var entity = eval("new " + prototype.TypeName + "()")
                            entity.UUID = typeName  + "%" + randomUUID()
                            panel.setEntity(entity)
                        }
                    }(prototype), null, true, null, "")
                }
            } (this))
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
    // use to keep track of the associated data.
    prototype.FieldsId = []
    
    prototype.Fields.unshift("ParentLnk")
    prototype.FieldsType.unshift("xs.string")
    prototype.FieldsVisibility.unshift(false)
    prototype.FieldsNillable.unshift(false)
    prototype.FieldsDocumentation.unshift("Relation with it parent.")
    prototype.FieldsOrder.push(prototype.FieldsOrder.length)
    prototype.FieldsDefaultValue.unshift("")
        
    // Append parent uuid if none is define.
    prototype.Fields.unshift("ParentUuid")
    prototype.FieldsType.unshift("xs.string")
    prototype.FieldsVisibility.unshift(false)
    prototype.FieldsNillable.unshift(false)
    prototype.FieldsDocumentation.unshift("The parent object UUID")
    prototype.FieldsOrder.push(prototype.FieldsOrder.length)
    prototype.FieldsDefaultValue.unshift("")

    // append the uuid...
    prototype.Fields.unshift("UUID")
    prototype.FieldsType.unshift("xs.string")
    prototype.FieldsVisibility.unshift(false)
    prototype.FieldsOrder.push(prototype.FieldsOrder.length)
    prototype.FieldsNillable.unshift(false)
    prototype.FieldsDocumentation.unshift("The object UUID")
    prototype.FieldsDefaultValue.unshift("")
    prototype.Ids.unshift("UUID")
        
    for (var i = 0; i < data.length; i++) {
        // I will retreive it item definition...
        if (isObjectReference(data[i].M_itemSubjectRef)) {
            // Here I will get the item definition...
            server.entityManager.getEntityByUuid(data[i].M_itemSubjectRef, false,
                function (result, caller) {
                    var done = caller.done
                    var callback = caller.callback
                    var prototype = caller.prototype

                    var fieldType = result.M_structureRef + ":Ref"
                    if (result.M_isCollection) {
                        fieldType = "[]" + fieldType
                    }

                    prototype.FieldsOrder.push(prototype.Fields.length)
                    prototype.FieldsId.push(data[i].UUID)
                    prototype.Fields.push("M_" + data[i].M_name.replaceAll(" ", "_"))
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
            prototype.FieldsId.push(data[i].UUID) // Set the uuid...
            prototype.Fields.push("M_" + data[i].M_name.replaceAll(" ", "_"))
            prototype.FieldsType.push(fieldType)
            prototype.FieldsVisibility.push(true)
            prototype.FieldsNillable.push(true)

            if (i == data.length - 1) {
                callback(prototype)
            }
        }
    }
}

/**
 * Save the data display in the view.
 */
BpmnDataView.prototype.save = function (callback) {

    // Save the entity data.
    function saveData(entity, callback) {
        var prototype = entity.getPrototype()
        var itemAwareInstances = []
        // The first tree field are not needed..
        for (var i = 3; i < prototype.Fields.length; i++) {
            var data = entity[prototype.Fields[i]]
            var fieldType = prototype.FieldsType[i]
            var isArray = fieldType.startsWith("[]")
            var dataStr = ""
            // Here I will stringify the value content.
            if (isArray) {
                data_ = []
                for (var j = 0; j < data.length; j++) {
                    // serialyse the object...
                    if (data[j].stringify != undefined) {
                        data_.push(data[j].stringify())
                    } else {
                        data_.push(data[j])
                    }
                }
                dataStr = JSON.stringify(data_)
            } else {
                if (data.stringify != undefined) {
                    dataStr = data.stringify()
                } else {
                    dataStr = data
                }
            }

            // Now I will create the itemaware element...
            server.workflowProcessor.newItemAwareElementInstance(prototype.FieldsId[i-3], dataStr,
                function (result, caller) {
                    caller.itemAwareInstances.push(result)
                    if (caller.itemAwareInstances.length == caller.count) {
                        caller.callback(caller.itemAwareInstances)
                    }
                },
                function () { /* Nothing here */ },
                { "itemAwareInstances": itemAwareInstances, "count": prototype.Fields.length - 3, "callback":callback })

        }
    }

    // The data input...
    if (this.dataInputsPanel != undefined) {
        saveData(this.dataOutputsPanel.entity, callback)
    }

    if (this.dataOutputsPanel != undefined) {
        saveData(this.dataOutputsPanel.entity, callback)
    }

    if(this.dataOutputsPanel == undefined && this.dataInputsPanel == undefined){
        // call the callback function...
        callback([])
    }
}