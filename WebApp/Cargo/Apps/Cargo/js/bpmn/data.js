// Contain various object constructor function for bpmn data
BpmsData = {}

/**
 * That class is use to create bpmn data view.
 */
var BpmnDataView = function (parent, dataInputs, dataOutputs, properties, instance) {

    // The parent bpmn element that contain the data.
    this.UUID = randomUUID()

    this.parent = parent

    // I will use an entity panel to display
    this.panel = this.parent.appendElement({ "tag": "div" }).down()

    // Bpmn data

    // Input data
    this.dataInputsPanel = null
    this.dataInputs = dataInputs

    // Output data
    this.dataOutputsPanel = null
    this.dataOutputs = dataOutputs
    
    // Properties.
    this.propertiesPanel = null
    this.properties = properties
    
    this.instance = instance

    /*// The function to call when the initialisation is done.
    this.initCallback = initCallback*/

    this.init()

    return this
}

BpmnDataView.prototype.init = function () {

    // first of all I will generate an entity prototype
    // from the data array.
    if (this.dataInputs != undefined) {
        var typeName = "BpmsData.DataInput_" + this.UUID.replaceAll("-", "_")
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
        var typeName = "BpmsData.DataOutput_" + this.UUID.replaceAll("-", "_")
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
    
    if (this.properties != undefined) {
        var typeName = "BpmsData.Properties_" + this.UUID.replaceAll("-", "_")
        this.generateEntityPrototype(this.properties, typeName,
            // The callback...
            function (dataView) {
                return function (prototype) {
                    // Generate the entity constructor.
                    prototype.generateConstructor()
                    setEntityPrototype(prototype)
                    
                    // So here with the entity prototype i will generate the entity panel.
                    dataView.propertiesPanel = new EntityPanel(dataView.panel, prototype.TypeName, function (prototype) { 
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
    function saveData(bpmnView,entity, callback) {
        
        function saveField(ids,fields,fieldType, index, callback, itemAwareInstances){
            var data = entity[fields[index+3]]
            var fieldType = fieldType[index+3]
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
            
            var itemAwareElementInstance
            if(itemAwareInstances.find(x => x.M_name == fields[index+3].substring(2)) == undefined){
                // Now I will create the itemaware element...
                server.workflowProcessor.newItemAwareElementInstance(ids[index], dataStr,
                    function (result, caller) {
                        caller.itemAwareInstances.push(result)
                        if (caller.index == caller.fields.length - 4) {
                            caller.callback(caller.itemAwareInstances,caller.instance)
                        }else{
                            saveField(caller.ids,caller.fields,caller.fieldType, caller.index + 1, caller.callback, caller.itemAwareInstances)
                        }
                    },
                    function () { },
                    { "itemAwareInstances": itemAwareInstances, "instance" : bpmnView.instance, "index" : index, "fields" : fields, "callback":callback,"ids" :ids, "fieldType" : fieldType })
            }else{
                itemAwareElementInstance = itemAwareInstances.find(x => x.M_name == fields[index+3].substring(2))
                itemAwareElementInstance.M_data = dataStr
                if (index == fields.length - 4) {
                    callback(itemAwareInstances,bpmnView.instance)
                }else{
                    saveField(ids,fields,fieldType, index + 1, callback, itemAwareInstances)
                }
            }
            
 
           
        }
        
        var prototype = entity.getPrototype()
        var itemAwareInstances = []
        if(bpmnView.instance != undefined){
            itemAwareInstances = itemAwareInstances.concat(bpmnView.instance.M_data)
            itemAwareInstances = itemAwareInstances.concat(bpmnView.instance.M_dataRef)
        }
       
        saveField(prototype.FieldsId, prototype.Fields,prototype.FieldsType, 0, callback, itemAwareInstances)
    }
    
    var entities = []
    

    // The data input...
    if (this.dataInputsPanel != undefined) {
        entities.push(this.dataInputsPanel.entity)
        //saveData(this,this.dataOutputsPanel.entity, callback)
    }

    if (this.dataOutputsPanel != undefined) {
        entities.push(this.dataOutputsPanel.entity)
        //saveData(this,this.dataOutputsPanel.entity, callback)
    }
    
    if (this.propertiesPanel!= undefined) {
        entities.push(this.propertiesPanel.entity)
        //saveData(this,this.propertiesPanel.entity, callback)
    }

    if(entities.length == 0){
        // call the callback function...
        callback([],this.instance)
    }else{
        function saveEntities(bpmnView,entities,index,callback){
            if(index < entities.length - 1){
                saveData(bpmnView, entities[index], function(bpmnView,entities,index,callback){
                    return function(){
                        console.log("---> im not done!")
                        saveEntities(bpmnView,entities, index + 1, callback)
                    }
                    
                }(bpmnView,entities,index,callback))
            }else{
                console.log("---> I'm done")
                saveData(bpmnView,entities[index],callback)
            }
        }
        
        saveEntities(this,entities, 0, callback)
    }

}