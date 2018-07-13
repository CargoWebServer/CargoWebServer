/*
 * The list of instances.
 */
var InstanceListView = function (parent, svgDiagram) {
    
	// The parent div.
	this.parent = parent

	/* This is the link to the related diagram view **/
	this.svgDiagram = svgDiagram
	
	this.svgDiagram.canvas.parentElement.appendElement({"tag" : "div", "style" : "position:absolute;top:0;left:0;margin:5px;", "id" : "processInstanceTitle"})

	/** So here I will get the list of instance for the process */
	//this.svgDiagram.bpmnDiagram.M_BPMNPlane.M_bpmnElement
	this.definitions = this.svgDiagram.bpmnDiagram.getParentDefinitions()
	this.bpmnProcess = null
	for(var i=0; i < this.definitions.M_rootElement.length; i++){
		if(this.definitions.M_rootElement[i].UUID == this.svgDiagram.bpmnDiagram.M_BPMNPlane.M_bpmnElement){
			this.bpmnProcess = this.definitions.M_rootElement[i]
			break
		}
	}

	if(this.bpmnProcess != null){
		server.workflowProcessor.getActiveProcessInstances(this.bpmnProcess.UUID, 
		// Success callback...
		function(instances, instanceListView){
		    instanceListView.svgDiagram.parent.element.classList.add("workflow-viewer")
		  
			instanceListView.parent.appendElement({"tag" : "table", "id" : "instancesList", "class" : "scrolltable"}).down()
			.appendElement({"tag" : "thead", "class":"table_header"}).down()
			.appendElement({"tag" : "th", "innerHtml" : "Number", "class":"header_cell", "style" : "width:20%;"})
			.appendElement({"tag" : "th", "innerHtml" : "Color", "class":"header_cell", "style" : "width:80%;"})
			
		    var content = instanceListView.parent.getChildById("instancesList").appendElement({"tag":"tbody", "class":"table_body", "style":"display: table-row-group;"}).down()
		    
		    for(var i = 0; i < instances.length; i++){
			    content.appendElement({"tag" : "tr", "title" : "Load this process instance", "id" : instances[i].M_id + "-processInstance", "class":"table_row"}).down()
			    .appendElement({"tag" : "td" , "innerHtml" : instances[i].M_number.toString(), "class":"body_cell"})
			    .appendElement({"tag" : "td" ,"innerHtml" : instances[i].M_colorName, "style" : "color:" + instances[i].M_colorNumber, "class":"body_cell instanceListHeader"})
			    
			    content.getChildById(instances[i].M_id + "-processInstance").element.onclick = function(instance){
			        return function(){
			            instanceListView.loadProcessInstance(instance, instanceListView)
			        }
			    }(instances[i])
			    
			  
			}
		}, 
		function(){

		}, this)
	}
	
	this.reloadDiagram = function(svgDiagram,instanceListView){
	    return function(){
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
        	
        	var classNameList = ["completed", "ready", "active"]
        	for(var i = 0; i < classNameList.length; i++){
        	    var elements = document.getElementsByClassName(classNameList[i])
        	    for(var j = elements.length-1; j >= 0 ; j--){
        	        elements[j].classList.remove(classNameList[i])
        	    }
        	}
        	svgDiagram.endEventClick = function(instance){}
        	
        	svgDiagram.taskClick = function(instance){}
        	
    		svgDiagram.scriptTaskClick = function(){}
        	
        	svgDiagram.dataObjectReferenceClick = function(){}
	    }
	}(svgDiagram,this)
	
	//Initial definitions of click functions, before an instance is loaded
	
	svgDiagram.scriptTaskClick = function(){}
	
	svgDiagram.userTaskClick = function(){}
	
	svgDiagram.endEventClick = function(){}
	
	svgDiagram.taskClick = function(){}
	
	svgDiagram.dataObjectReferenceClick = function(){}
	
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
                if(entity.M_bpmnElementId == instanceListView.bpmnProcess.UUID){
                    instanceListView.parent.getChildById("instancesList").appendElement({"tag" : "tr", "class" : "table_row", "title" : "Load this process instance", "id" : entity.M_id + "-processInstance"}).down()
    			    .appendElement({"tag" : "td" ,"class" : "body_cell", "innerHtml" : entity.M_number.toString(),})
    			    .appendElement({"tag" : "td" , "class" : "body_cell instanceListHeader","innerHtml" : entity.M_colorName, "style" : "color:" + entity.M_colorNumber})
    			    
    			    instanceListView.parent.getChildById(entity.M_id + "-processInstance").element.onclick = function(instance){
    			        return function(){
    			            server.entityManager.getEntityByUuid(instance.UUID,false,
    			                function(instance,caller){
    			                    caller.loadProcessInstance(instance, caller)
    			                },function(){}, instanceListView)
    			            
    			        }
    			    }(entity)
                }
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

/////////////////////////////////////////////////////////////////////////////////////////
// Load all of the required data and functions for the given instance
/////////////////////////////////////////////////////////////////////////////////////////
InstanceListView.prototype.loadProcessInstance = function(instance,instanceListView){
    console.log(instance)
     instanceListView.svgDiagram.canvas.parentElement.getChildById("processInstanceTitle").removeAllChilds()
    instanceListView.svgDiagram.canvas.parentElement.getChildById("processInstanceTitle").appendElement({"tag" : "span", "innerHtml"  : instance.M_colorName})
    .appendElement({"tag" : "a","style" : "cursor:pointer;", "id" : "clearProcessInstanceBtn"}).down()
    .appendElement({"tag" : "i", "class" : "fa fa-close"})
    instanceListView.svgDiagram.canvas.parentElement.getChildById("processInstanceTitle").element.style.color = instance.M_colorNumber
   
   instanceListView.svgDiagram.canvas.parentElement.getChildById("clearProcessInstanceBtn").element.onclick = function(instance,instanceListView){
       return function(){
            
            instanceListView.svgDiagram.canvas.parentElement.getChildById("processInstanceTitle").element.innerHTML = null;
            instanceListView.reloadDiagram()
            
       }
   }(instance,instanceListView)
   
    for(var i = 0; i < instance.M_flowNodeInstances.length; i++){
        var element = document.getElementsByName(instance.M_flowNodeInstances[i].M_bpmnElementId)[0]
        element.flowNodeInstance = instance.M_flowNodeInstances[i]
        element.instanceListView = instanceListView
        //Attach a listener for updated entities in the process instance
        server.entityManager.attach(element, UpdateEntityEvent, function (evt, element) {
            var entity = evt.dataMap.entity 
             console.log("---> entity update: ", entity)
            if (entity  !== undefined) {
                if(entity.UUID == element.flowNodeInstance.UUID){
                   
                    //Lifecycle state handler, sets the CSS class to the equivalent state
                    switch(entity.M_lifecycleState){
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
	    })
	    
        //Lifecycle state handler, sets the CSS class to the equivalent state
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
    
    instanceListView.svgDiagram.dataObjectReferenceClick = function(instance){
        return function(parent, diagramElement,group,svgDiagram,event,bpmnElement){
    	    event.preventDefault();
            event.stopPropagation();
            
        	// Get the menu position
        	var x = event.offsetX - 5;     // Get the horizontal coordinate
        	var y = event.offsetY - 5;     // 
            
            //Menu item for the log info
            var getDataInfoItem = new MenuItem("data_objectreference", "Data", {}, 0, function (diagramElement,parent,instance,diagramElement) {
                return function () {
                    //Check that the entity actually exists first
       
                    if(instance.M_data.find(x => x.M_bpmnElementId == bpmnElement.M_dataObjectRef.UUID) != undefined){
                        data = instance.M_data.find(x => x.M_bpmnElementId == bpmnElement.M_dataObjectRef.UUID)
                        server.entityManager.getEntityByUuid(data.M_bpmnElementId,false,
                            function(entity,caller){
                                data.M_bpmnElement = entity
                                loadDataDialog(caller.data,caller.parent,caller.diagramElement)
                            },function(){}, {"data" : data, "parent" : parent, "diagramElement" : diagramElement})
                        
                        
                    }
                    
                    
                }
            }(diagramElement,parent,instance,diagramElement), "fa fa-database")

        	var popup = parent.appendElement({ "tag": "div", "style": "top:" + y + "px; left:" + x + "px;" }).down()
        	var contextMenu = new PopUpMenu(group, [ getDataInfoItem], event)
        }
	  
	}(instance)
    
    //Override the click function on the start event
    instanceListView.svgDiagram.startEventClick = function(instance){
        return function(parent, diagramElement,group,svgDiagram,event,bpmnElement){
    	    event.preventDefault();
            event.stopPropagation();
            
        	// Get the menu position
        	var x = event.offsetX - 5;     // Get the horizontal coordinate
        	var y = event.offsetY - 5;     // 
            
            //Menu item for the log info
            var getLogInfoItem = new MenuItem("loginfo_startevent", "Info", {}, 0, function (diagramElement,parent,instance,diagramElement) {
                return function () {
                    //Check that the entity actually exists first
                    if(instance.M_flowNodeInstances.find(x => x.M_bpmnElementId == bpmnElement.UUID) != undefined){
                        logInfos = instance.M_flowNodeInstances.find(x => x.M_bpmnElementId == bpmnElement.UUID).M_logInfoRef
                        //Check if the log infos were fully loaded
                        if(logInfos[0].M_action != undefined){
                            loadLogInfosDialog(logInfos,parent,diagramElement)
                        }else{
                            //If not, loop through each log info and load their attributes
                            function drawLogInfo(logInfos, index,parent){
                        	    var logInfo = logInfos[index]
                        		server.entityManager.getEntityByUuid(logInfo, false,
                        			function (logInfo, caller) { 
                          				caller.logInfos[caller.index] = logInfo
                        				if(caller.index < caller.logInfos.length - 1){
                        				    drawLogInfo(caller.logInfos, caller.index + 1)
                        				}else{
                        				    loadLogInfosDialog(caller.logInfos,caller.parent,caller.diagramElement)
                        				}
                        			},
                        			function (errObj, caller) {},
                        			{"index":index, "logInfos":logInfos, "parent" : parent, "diagramElement" : diagramElement})
                        	}
                        	drawLogInfo(logInfos, 0,parent)
                        }
                    }
                    
                    
                }
            }(diagramElement,parent,instance,diagramElement), "fa fa-book")

        	var popup = parent.appendElement({ "tag": "div", "style": "top:" + y + "px; left:" + x + "px;" }).down()
        	var contextMenu = new PopUpMenu(group, [ getLogInfoItem], event)
        }
	  
	}(instance)
	
	//Override the click function on the end event
	instanceListView.svgDiagram.endEventClick = function(instance){
        return function(parent, diagramElement,group,svgDiagram,event,bpmnElement){
    	    event.preventDefault();
            event.stopPropagation();
            
        	// Get the menu position
        	var x = event.offsetX - 5;     // Get the horizontal coordinate
        	var y = event.offsetY - 5;     // 
            
            //Menu item for the log info
            var getLogInfoItem = new MenuItem("loginfo_startevent", "Info", {}, 0, function (diagramElement,parent,instance) {
                return function () {
                    //Check that the entity actually exists first
                    if(instance.M_flowNodeInstances.find(x => x.M_bpmnElementId == bpmnElement.UUID) != undefined){
                        logInfos = instance.M_flowNodeInstances.find(x => x.M_bpmnElementId == bpmnElement.UUID).M_logInfoRef
                        //Check if the log infos were fully loaded
                        if(logInfos[0].M_action != undefined){
                            loadLogInfosDialog(logInfos,parent,diagramElement)
                        }else{
                            //If not, loop through each log info and load their attributes
                            function drawLogInfo(logInfos, index,parent){
                        	    var logInfo = logInfos[index]
                        		server.entityManager.getEntityByUuid(logInfo, false,
                        			function (logInfo, caller) {
                        				caller.logInfos[caller.index] = logInfo
                        				if(caller.index < caller.logInfos.length - 1){
                        				    drawLogInfo(caller.logInfos, caller.index + 1)
                        				}else{
                        				    loadLogInfosDialog(caller.logInfos,caller.parent,caller.diagramElement)
                        				}
                        			},
                        			function (errObj, caller) {},
                        			{ "index":index, "logInfos":logInfos, "parent" : parent, "diagramElement" : diagramElement})
                        	}
                        	drawLogInfo(logInfos, 0,parent)
                        }
                    }
                }
            }(diagramElement,parent,instance), "fa fa-book")

        	var popup = parent.appendElement({ "tag": "div", "style": "top:" + y + "px; left:" + x + "px;" }).down()
        	var contextMenu = new PopUpMenu(group, [ getLogInfoItem], event)
        }
	  
	}(instance)
	
	//Override the click function of a generic task
	instanceListView.svgDiagram.taskClick = function(instance){
        return function(parent, diagramElement,group,svgDiagram,event,bpmnElement){
    	    event.preventDefault();
            event.stopPropagation();
            
        	// Get the menu position
        	var x = event.offsetX - 5;     // Get the horizontal coordinate
        	var y = event.offsetY - 5;     // 
            
            //Menu item for the log info
            var getLogInfoItem = new MenuItem("loginfo_startevent", "Info", {}, 0, function (diagramElement,parent,instance) {
                return function () {
                    //Check that the entity actually exists first
                    if(instance.M_flowNodeInstances.find(x => x.M_bpmnElementId == bpmnElement.UUID) != undefined){
                        logInfos = instance.M_flowNodeInstances.find(x => x.M_bpmnElementId == bpmnElement.UUID).M_logInfoRef
                        //Check if the log infos were fully loaded
                        if(logInfos[0].M_action != undefined){
                            loadLogInfosDialog(logInfos,parent,diagramElement)
                        }else{
                            //If not, loop through each log info and load their attributes
                            function drawLogInfo(logInfos, index,parent){
                        	    var logInfo = logInfos[index]
                        		server.entityManager.getEntityByUuid(logInfo, false,
                        			function (logInfo, caller) { 
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
                }
            }(diagramElement,parent,instance), "fa fa-book")

        	var popup = parent.appendElement({ "tag": "div", "style": "top:" + y + "px; left:" + x + "px;" }).down()
        	var contextMenu = new PopUpMenu(group, [ getLogInfoItem], event)
        }
	  
	}(instance)
	
	//Override the click function of an user task
	instanceListView.svgDiagram.userTaskClick = function(instance){
        return function(parent, diagramElement,group,svgDiagram,event,bpmnElement){
    	    event.preventDefault();
            event.stopPropagation();
            
        	// Get the menu position
        	var x = event.offsetX - 5;     // Get the horizontal coordinate
        	var y = event.offsetY - 5;     // 
            
            //Menu item for the log info
            var getLogInfoItem = new MenuItem("loginfo_startevent", "Info", {}, 0, function (diagramElement,parent,instance) {
                return function () {
                    //Check that the entity actually exists first
                    if(instance.M_flowNodeInstances.find(x => x.M_bpmnElementId == bpmnElement.UUID) != undefined){
                        logInfos = instance.M_flowNodeInstances.find(x => x.M_bpmnElementId == bpmnElement.UUID).M_logInfoRef
                        //Check if the log infos were fully loaded
                        if(logInfos[0].M_action != undefined){
                            loadLogInfosDialog(logInfos,parent,diagramElement)
                        }else{
                            //If not, loop through each log info and load their attributes
                            function drawLogInfo(logInfos, index,parent){
                        	    var logInfo = logInfos[index]
                        		server.entityManager.getEntityByUuid(logInfo, false,
                        			function (logInfo, caller) { 
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
                }
            }(diagramElement,parent,instance), "fa fa-book")
            
            //Menu item to play the user task
            var playItem = new MenuItem("play_activity", "Play", {}, 0, function(){}, "fa fa-play")
            
            //Menu item to cancel the user task
            var cancelItem = new MenuItem("cancel_activity", "Cancel", {}, 0, function(){}, "fa fa-close")

        	var popup = parent.appendElement({ "tag": "div", "style": "top:" + y + "px; left:" + x + "px;" }).down()
        	var contextMenu = new PopUpMenu(group, [ getLogInfoItem, playItem, cancelItem], event)
        }
	  
	}(instance)
   
}


//////////////////////////////////////////////////////////////////////////////////////////////////////
// Load and display a data instance
//////////////////////////////////////////////////////////////////////////////////////////////////////
function loadDataDialog(data, parent, diagramElement){
    var dialog = new Dialog(randomUUID(),parent,false, data.M_id)
   
    var x = diagramElement.M_Bounds.M_x
	var y = diagramElement.M_Bounds.M_y + diagramElement.M_Bounds.M_height + 5
	
	console.log(data)
	
	dialog.content.appendElement({"tag" : "table", "id" : "logInfosTable", "class" : "instanceListTable"}).down()
	.appendElement({"tag" : "tr"}).down()
	.appendElement({"tag" : "th","class" : "", "innerHtml" : data.M_bpmnElement.M_name})
	dialog.setPosition(x, y)
	
}


//////////////////////////////////////////////////////////////////////////////////////////////////////
// Load the history of logs into the given parent, including the actor, the action and the date
//////////////////////////////////////////////////////////////////////////////////////////////////////
 function loadLogInfosDialog(logInfos, parent, diagramElement){
    var dialog = new Dialog(randomUUID(),parent,false, "Infos")
    
    var x = diagramElement.M_Bounds.M_x
	var y = diagramElement.M_Bounds.M_y + diagramElement.M_Bounds.M_height + 5
	
	dialog.content.appendElement({"tag" : "table", "id" : "logInfosTable", "class" : "instanceListTable"}).down()
	.appendElement({"tag" : "tr"}).down()
	.appendElement({"tag" : "th","class" : "", "innerHtml" : "Date"})
	.appendElement({"tag" : "th","class" : "", "innerHtml" : "Actor"})
	.appendElement({"tag" : "th","class" : "","innerHtml" : "Action"})
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
                if(account.M_name != ""){
                    dialog.content.getChildById("logInfosTable").appendElement({"tag" : "tr"}).down()
                    .appendElement({"tag" : "th", "class" : "", "innerHtml" : caller.date})
                    .appendElement({"tag" : "th", "class" : "", "innerHtml" : account.M_name})
                    .appendElement({"tag" : "th", "class" : "", "innerHtml" : caller.action}) 
                }else{
                     dialog.content.getChildById("logInfosTable").appendElement({"tag" : "tr"}).down()
                    .appendElement({"tag" : "th", "class" : "", "innerHtml" : caller.date})
                    .appendElement({"tag" : "th", "class" : "", "innerHtml" : account.M_id})
                    .appendElement({"tag" : "th", "class" : "", "innerHtml" : caller.action})
                }
               
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