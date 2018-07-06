/*
 * The list of instances.
 */
var InstanceListView = function (parent, svgDiagram) {
	// The parent div.
	this.parent = parent

	/* This is the link to the related diagram view **/
	this.svgDiagram = svgDiagram

	/** So here I will get the list of instance for the process */
	//this.svgDiagram.bpmnDiagram.M_BPMNPlane.M_bpmnElement
	var definitions = this.svgDiagram.bpmnDiagram.getParentDefinitions()
	var bpmnProcess = null
	for(var i=0; i < definitions.M_rootElement.length; i++){
		if(definitions.M_rootElement[i].UUID == this.svgDiagram.bpmnDiagram.M_BPMNPlane.M_bpmnElement){
			bpmnProcess = definitions.M_rootElement[i]
			break
		}
	}

	if(bpmnProcess != null){
		server.workflowProcessor.getActiveProcessInstances(bpmnProcess.UUID, 
		// Success callback...
		function(instances, instanceListView){
		    instanceListView.parent.element.style.width = "100%"
		    instanceListView.parent.element.style.margin = "8px"
		    instanceListView.svgDiagram.parent.element.classList.add("workflow-viewer")
		    console.log(instances)
			instanceListView.parent.appendElement({"tag" : "table", "id" : "instancesList", "class" : "instanceListTable"}).down()
			.appendElement({"tag" : "tr"}).down()
			.appendElement({"tag" : "th", "innerHtml" : "Number","class" : "instanceListHeader", "style" : "width:20%;"})
			.appendElement({"tag" : "th", "innerHtml" : "Color", "class" : "instanceListHeader","style" : "width:80%;"})
		
			for(var i = 0; i < instances.length; i++){
			    instanceListView.parent.getChildById("instancesList").appendElement({"tag" : "tr", "class" : "instanceListRow", "title" : "Load this process instance", "id" : instances[i].M_id + "-processInstance"}).down()
			    .appendElement({"tag" : "th" ,"class" : "instanceListHeader", "innerHtml" : instances[i].M_number.toString(),})
			    .appendElement({"tag" : "th" , "class" : "instanceListHeader","innerHtml" : instances[i].M_colorName, "style" : "color:" + instances[i].M_colorNumber})
			    
			    instanceListView.parent.getChildById(instances[i].M_id + "-processInstance").element.onclick = function(instance){
			        return function(){
			            instanceListView.loadProcessInstance(instance, instanceListView)
			        }
			    }(instances[i])
			}
		}, 
		function(){

		}, this)
	}
	
	svgDiagram.startEventClick = function(parent, diagramElement,group,svgDiagram,event,bpmnElement){
	    event.preventDefault();
        event.stopPropagation();
    	// Get the menu position
    	var x = event.offsetX - 5;     // Get the horizontal coordinate
    	var y = event.offsetY - 5;     // 
        var playEventItem = new MenuItem("play_startevent", "Play", {}, 0, function (diagramElement,parent) {
            return function () {
                var startEvent = entities[diagramElement.M_bpmnElement]
    			// Remove the menu.
    
    			// Create a new process instance wizard...
    			new ProcessWizard(parent, startEvent)
            }
        }(diagramElement,parent), "fa fa-play")
    	var contextMenu = new PopUpMenu(group, [playEventItem], event)
	}
	
	server.entityManager.attach(this, NewEntityEvent, function (evt, instanceListView) {
	    var entity = evt.dataMap.entity 
        if (entity !== undefined) {
            if(entity.TYPENAME.startsWith("BPMS.ProcessInstance")){
                console.log("---> entity create: ", entity)
                instanceListView.parent.getChildById("instancesList").appendElement({"tag" : "tr", "class" : "instanceListRow", "title" : "Load this process instance", "id" : entity.M_id + "-processInstance"}).down()
			    .appendElement({"tag" : "th" ,"class" : "instanceListHeader", "innerHtml" : entity.M_number.toString(),})
			    .appendElement({"tag" : "th" , "class" : "instanceListHeader","innerHtml" : entity.M_colorName, "style" : "color:" + entity.M_colorNumber})
			    
			    instanceListView.parent.getChildById(entity.M_id + "-processInstance").element.onclick = function(instance){
			        return function(){
			            instanceListView.loadProcessInstance(instance, instanceListView)
			        }
			    }(entity)
            }
        }
	})
	
    server.entityManager.attach(this, DeleteEntityEvent, function (evt, instanceListView) {
        var entity = evt.dataMap.entity 
        if (entity  !== undefined) {
            if(entity.TYPENAME.startsWith("BPMS.ProcessInstance")){
                console.log("---> entity deleted: ", entity)
            }
        }
	})
	
	/*server.entityManager.attach(this, UpdateEntityEvent, function (evt, instanceListView) {
        var entity = evt.dataMap.entity 
        if (entity  !== undefined) {
            if(entity.TYPENAME.startsWith("BPMS.ProcessInstance")){
                console.log("---> entity update: ", entity)
            }
        }
 	 
	  })*/
	return this
}

InstanceListView.prototype.loadProcessInstance = function(instance,instanceListView){

   
    for(var i = 0; i < instance.M_flowNodeInstances.length; i++){
        if(instance.M_flowNodeInstances[i].M_bpmnElementId == undefined){
            server.entityManager.getEntityByUuid(instance.M_flowNodeInstances[i],false,
                function(flowNodeInstance,caller){
                    var element = document.getElementsByName(flowNodeInstance.M_bpmnElementId)[0]
                     element.flowNodeInstance = flowNodeInstance
                     element.instanceListView = caller.instanceListView
                    server.entityManager.attach(element, UpdateEntityEvent, function (evt, element) {
                        var entity = evt.dataMap.entity 
                        if (entity  !== undefined) {
                            if(entity.UUID == element.flowNodeInstance.UUID){
                                console.log("---> entity update: ", entity)
                                console.log(element)
                            }
                        }
             	 
            	  })
                    
                    switch(flowNodeInstance.M_lifecycleState){
                        case 1:
                            element.parentElement.classList.add("completed")
                            element.parentElement.classList.remove("ready")
                            element.parentElement.classList.remove("active")
                            break;
                        case 9:
                            element.parentElement.classList.add("ready")
                            element.parentElement.classList.remove("completed")
                            element.parentElement.classList.remove("active")
                            break;
                        case 10:
                            element.parentElement.classList.add("active")
                            element.parentElement.classList.remove("ready")
                            element.parentElement.classList.remove("completed")
                            break;
                        default:
                            break;
                    }
                },function(){},{"instanceListView"  : instanceListView})
        }else{
            var element = document.getElementsByName(instance.M_flowNodeInstances[i].M_bpmnElementId)[0]
             element.flowNodeInstance = instance.M_flowNodeInstances[i]
             element.instanceListView = instanceListView
            server.entityManager.attach(element, UpdateEntityEvent, function (evt, element) {
                var entity = evt.dataMap.entity 
                if (entity  !== undefined) {
                    if(entity.UUID == element.flowNodeInstance.UUID){
                        console.log("---> entity update: ", entity)
                        console.log(element)
                    }
                }
     	 
    	  })
            
            switch(instance.M_flowNodeInstances[i].M_lifecycleState){
                case 1:
                    element.parentElement.classList.add("completed")
                    element.parentElement.classList.remove("ready")
                    element.parentElement.classList.remove("active")
                    break;
                case 9:
                    element.parentElement.classList.add("ready")
                    element.parentElement.classList.remove("completed")
                    element.parentElement.classList.remove("active")
                    break;
                case 10:
                    element.parentElement.classList.add("active")
                    element.parentElement.classList.remove("ready")
                    element.parentElement.classList.remove("completed")
                    break;
                default:
                    break;
            }
        }
         
    }
    
    instanceListView.svgDiagram.startEventClick = function(instance){
        return function(parent, diagramElement,group,svgDiagram,event,bpmnElement){
    	    event.preventDefault();
            event.stopPropagation();
            
        	// Get the menu position
        	var x = event.offsetX - 5;     // Get the horizontal coordinate
        	var y = event.offsetY - 5;     // 
            
            // Delete a file from a project.
            var getLogInfoItem = new MenuItem("loginfo_startevent", "Info", {}, 0, function (diagramElement,parent,instance,diagramElement) {
                return function () {
                    logInfos = instance.M_flowNodeInstances.find(x => x.M_bpmnElementId == bpmnElement.UUID).M_logInfoRef
                    if(logInfos[0].M_action != undefined){
                        loadLogInfosDialog(logInfos,parent,diagramElement)
                    }else{
                        
                        function drawLogInfo(logInfos, index,parent){
                    	    var logInfo = logInfos[index]
                    		server.entityManager.getEntityByUuid(logInfo, false,
                    			function (logInfo, caller) { 
                    			    console.log(logInfo)
                    				
                    				caller.logInfos[caller.index] = logInfo
                    				if(caller.index < caller.logInfos.length - 1){
                    				    drawLogInfo(caller.logInfos, caller.index + 1)
                    				}else{
                    				    loadLogInfosDialog(caller.logInfos,caller.parent,caller.diagramElement)
                    				}
                    			},
                    			function (errObj, caller) { 
                    
                    			},
                    			{ "index":index, "logInfos":logInfos, "parent" : parent, "diagramElement" : diagramElement})
                    	}
                    	drawLogInfo(logInfos, 0,parent)
                    }
                    
                }
            }(diagramElement,parent,instance,diagramElement), "fa fa-book")

        	var popup = parent.appendElement({ "tag": "div", "style": "top:" + y + "px; left:" + x + "px;" }).down()
        	var contextMenu = new PopUpMenu(group, [ getLogInfoItem], event)
        }
	  
	}(instance)
   
}

 function loadLogInfosDialog(logInfos, parent, diagramElement){
    var dialog = new Dialog(randomUUID(),parent,false, "Infos")
    
    var x = diagramElement.M_Bounds.M_x
	var y = diagramElement.M_Bounds.M_y + diagramElement.M_Bounds.M_height + 5
	
	dialog.content.appendElement({"tag" : "table", "id" : "logInfosTable", "class" : "instanceListTable"}).down()
	.appendElement({"tag" : "tr"}).down()
	.appendElement({"tag" : "th","class" : "instanceListHeader", "innerHtml" : "Date"})
	.appendElement({"tag" : "th","class" : "instanceListHeader", "innerHtml" : "Actor"})
	.appendElement({"tag" : "th","class" : "instanceListHeader","innerHtml" : "Action"})
	dialog.setPosition(x, y)
	
	dialog.ok.element.onclick = function (dialog) {
		return function () {
			// I will close the dialogue first...
			dialog.close()

			// I will save the data view...
		
			
		}
	} (dialog)
	
    for(var i = 0; i < logInfos.length; i++){
        var date = moment.unix(logInfos[i].M_date).format("dddd, MMMM Do YYYY, h:mm:ss a")
       
        server.entityManager.getEntityByUuid(logInfos[i].M_actor,false,
            function(account,caller){
                dialog.content.getChildById("logInfosTable").appendElement({"tag" : "tr"}).down()
                .appendElement({"tag" : "th", "class" : "instanceListHeader", "innerHtml" : caller.date})
                .appendElement({"tag" : "th", "class" : "instanceListHeader", "innerHtml" : account.M_name})
                .appendElement({"tag" : "th", "class" : "instanceListHeader", "innerHtml" : caller.action})
            },function(){},{"dialog" : dialog, "date" : date, "action" : logInfos[i].M_action})
    }
    
}

/////////////////////////////////////////////////////////////////////////////////////////
// The data input wizard...
/////////////////////////////////////////////////////////////////////////////////////////
var ProcessWizard = function (parent, startEvent) {
	this.parent = parent
	this.id = randomUUID()

	// The wizard dialog...
	this.dialog = new Dialog(this.id, this.parent, false, "New Process")

	// Set the dialog position...
	var diagramElement = startEvent.getDiagramElement()
	var x = diagramElement.M_Bounds.M_x
	var y = diagramElement.M_Bounds.M_y + diagramElement.M_Bounds.M_height + 5
	this.dialog.setPosition(x, y)

	this.content = this.dialog.content.appendElement({ "tag": "div", "class": "process_wizard_content" }).down()

	// That will contain the values ask by the user...
	this.dataView = new BpmnDataView(this.content, startEvent)

	this.dialog.ok.element.onclick = function (dataView, process, dialog) {
		return function () {
			// I will close the dialogue first...
			dialog.close()

			// I will save the data view...
			dataView.save(function (process) {
				return function (itemAwareInstances) {
					// Here I will create the new process...
					server.workflowManager.startProcess(process.UUID, itemAwareInstances, [],
						// Success Callback
						function (result, caller) {
							/* Nothing here */
						},
						// Error Callback
						function () {/* Nothing here */ },
						{})
				}
			} (process))
			
		}
	} (this.dataView, startEvent.getParent(), this.dialog)

	return this
}