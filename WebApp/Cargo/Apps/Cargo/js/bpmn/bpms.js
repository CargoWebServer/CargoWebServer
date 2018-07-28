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
                
                var getInfoItem = new MenuItem("bpmninfo", "BPMN Element", {}, 0, function (bpmnElement,parent,diagramElement) {
                        return function () {
                            showBpmnInfo(bpmnElement,parent,diagramElement)
                        }
                }(bpmnElement,parent,diagramElement), "fa fa-cube")
            	var contextMenu = new PopUpMenu(group, [playEventItem,getInfoItem], event)
            	
            	function showBpmnInfo(bpmnElement,parent,diagramElement){
                     //Initialize the dialog element and set its position
                    var dialog = new Dialog(randomUUID(),parent,false, diagramElement.M_id)
                    var x = diagramElement.M_Bounds.M_x
                	var y = diagramElement.M_Bounds.M_y + diagramElement.M_Bounds.M_height + 5
                	dialog.content.appendElement({"tag" : "table", "id" : "dataInfosTable", "class" : "instanceListTable"}).down()
                    dialog.setPosition(x, y)
                    
                     var entityPanelParent = dialog.content.getChildById("dataInfosTable").appendElement({"tag" : "tr"}).down()
                	.appendElement({"tag" : "th","class" : "borderHeader", "innerHtml" : bpmnElement.M_name})
                	.appendElement({"tag" : "th", "class" : "borderHeader", "style" : "overflow-y: auto;max-height: 350px;max-width:600px;display: block;" }).down()
                	//Create the entity panel from the entity we loaded
                    new EntityPanel(entityPanelParent,bpmnElement.TYPENAME,
                        function(bpmnElement){
                           return function(panel){
                             panel.setEntity(bpmnElement)
                            } 
                        }(bpmnElement))
                    
                }
            	
            
            	
        	}
        	
        	var classNameList = ["completed", "ready", "active", "passedConnector", "inactive"]
        	for(var i = 0; i < classNameList.length; i++){
        	    var elements = document.getElementsByClassName(classNameList[i])
        	    for(var j = elements.length-1; j >= 0 ; j--){
        	        elements[j].classList.remove(classNameList[i])
        	    }
        	}
        	svgDiagram.endEventClick = function(parent, diagramElement,group,svgDiagram,event,bpmnElement){
    		    loadBpmnElementInfo(parent, diagramElement,group,svgDiagram,event,bpmnElement)
    		}
        	
        	svgDiagram.taskClick = function(parent, diagramElement,group,svgDiagram,event,bpmnElement){
    		    loadBpmnElementInfo(parent, diagramElement,group,svgDiagram,event,bpmnElement)
    		}
        	
    		svgDiagram.scriptTaskClick = function(parent, diagramElement,group,svgDiagram,event,bpmnElement){
    		    loadBpmnElementInfo(parent, diagramElement,group,svgDiagram,event,bpmnElement)
    		}
    		
    		svgDiagram.userTaskClick = function(parent, diagramElement,group,svgDiagram,event,bpmnElement){
    		    loadBpmnElementInfo(parent, diagramElement,group,svgDiagram,event,bpmnElement)
    		}
        	
        	svgDiagram.dataObjectReferenceClick = function(parent, diagramElement,group,svgDiagram,event,bpmnElement){
    		    loadBpmnElementInfo(parent, diagramElement,group,svgDiagram,event,bpmnElement)
    		}
        	
        	svgDiagram.dataInputClick = function(parent, diagramElement,group,svgDiagram,event,bpmnElement){
    		    loadBpmnElementInfo(parent, diagramElement,group,svgDiagram,event,bpmnElement)
    		}
	
	        svgDiagram.dataOutputClick = function(parent, diagramElement,group,svgDiagram,event,bpmnElement){
    		    loadBpmnElementInfo(parent, diagramElement,group,svgDiagram,event,bpmnElement)
    		}
    		
    		svgDiagram.gatewayClick = function(parent, diagramElement,group,svgDiagram,event,bpmnElement){
    		    loadBpmnElementInfo(parent, diagramElement,group,svgDiagram,event,bpmnElement)
    		}
    		
	    }
	}(svgDiagram,this)
	
	this.reloadDiagram()

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
    
    
    instanceListView.reloadDiagram()
    
    var elements = document.getElementsByClassName("activity")
    for(var i = 0; i < elements.length;i++){
        elements[i].firstElementChild.classList.add("inactive")
    }
    
    var elements = document.getElementsByClassName("gateway")
    for(var i = 0; i < elements.length;i++){
        elements[i].firstElementChild.classList.add("inactive")
    }
    
    
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
   
   
   
    
    instanceListView.svgDiagram.dataObjectReferenceClick = function(instance){
        return function(parent, diagramElement,group,svgDiagram,event,bpmnElement){
            //Check that the entity actually exists first
            if(instance.M_data.find(x => x.M_bpmnElementId == bpmnElement.M_dataObjectRef.UUID) != undefined){
                event.preventDefault();
                event.stopPropagation();
                
            	// Get the menu position
            	var x = event.offsetX - 5;     // Get the horizontal coordinate
            	var y = event.offsetY - 5;     // 
                
                //Menu item for the log info
                var getDataInfoItem = new MenuItem("data_objectreference", "Data", {}, 0, function (diagramElement,parent,instance) {
                    return function () {
                        
                        data = instance.M_data.find(x => x.M_bpmnElementId == bpmnElement.M_dataObjectRef.UUID)
                        server.entityManager.getEntityByUuid(data.M_bpmnElementId,false,
                            function(entity,caller){
                                data.M_bpmnElement = entity
                                loadDataDialog(caller.data,caller.parent,caller.diagramElement,false)
                            },function(){}, {"data" : data, "parent" : parent, "diagramElement" : diagramElement})
                        
                        
                    }
                }(diagramElement,parent,instance), "fa fa-database")
    
            	var popup = parent.appendElement({ "tag": "div", "style": "top:" + y + "px; left:" + x + "px;" }).down()
            	var contextMenu = new PopUpMenu(group, [ getDataInfoItem], event)
            }
    	    
        }
	  
	}(instance)
	
	instanceListView.svgDiagram.dataInputClick = function(instance){
        return function(parent, diagramElement,group,svgDiagram,event,bpmnElement){
            //Check that the entity actually exists first
            if(instance.M_data.find(x => x.M_bpmnElementId == bpmnElement.UUID) != undefined){
                event.preventDefault();
                event.stopPropagation();
                
            	// Get the menu position
            	var x = event.offsetX - 5;     // Get the horizontal coordinate
            	var y = event.offsetY - 5;     // 
                
                //Menu item for the log info
                var getDataInfoItem = new MenuItem("data_input", "Data", {}, 0, function (diagramElement,parent,instance,diagramElement) {
                    return function () {
                        
                        data = instance.M_data.find(x => x.M_bpmnElementId == bpmnElement.UUID)
                        server.entityManager.getEntityByUuid(data.M_bpmnElementId,false,
                            function(entity,caller){
                                data.M_bpmnElement = entity
                                loadDataDialog(caller.data,caller.parent,caller.diagramElement,false)
                            },function(){}, {"data" : data, "parent" : parent, "diagramElement" : diagramElement})
                    }
                }(diagramElement,parent,instance,diagramElement), "fa fa-database")
    
            	var popup = parent.appendElement({ "tag": "div", "style": "top:" + y + "px; left:" + x + "px;" }).down()
            	var contextMenu = new PopUpMenu(group, [ getDataInfoItem], event)
            }
    	    
        }
	  
	}(instance)
	
	instanceListView.svgDiagram.dataOutputClick = function(instance){
        return function(parent, diagramElement,group,svgDiagram,event,bpmnElement){
    	     //Check that the entity actually exists first
            if(instance.M_data.find(x => x.M_bpmnElementId == bpmnElement.UUID) != undefined){
                event.preventDefault();
                event.stopPropagation();
                
            	// Get the menu position
            	var x = event.offsetX - 5;     // Get the horizontal coordinate
            	var y = event.offsetY - 5;     // 
                
                //Menu item for the log info
                var getDataInfoItem = new MenuItem("data_output", "Data", {}, 0, function (diagramElement,parent,instance,diagramElement) {
                    return function () {
                        
                        data = instance.M_data.find(x => x.M_bpmnElementId == bpmnElement.UUID)
                        server.entityManager.getEntityByUuid(data.M_bpmnElementId,false,
                            function(entity,caller){
                                data.M_bpmnElement = entity
                                loadDataDialog(caller.data,caller.parent,caller.diagramElement,false)
                            },function(){}, {"data" : data, "parent" : parent, "diagramElement" : diagramElement})
                    }
                }(diagramElement,parent,instance,diagramElement), "fa fa-database")
    
            	var popup = parent.appendElement({ "tag": "div", "style": "top:" + y + "px; left:" + x + "px;" }).down()
            	var contextMenu = new PopUpMenu(group, [ getDataInfoItem], event)
            }
        }
	  
	}(instance)
    
    //Override the click function on the start event
    instanceListView.svgDiagram.startEventClick = function(instance){
        return function(parent, diagramElement,group,svgDiagram,event,bpmnElement){
            //Check that the entity actually exists first
            if(instance.M_flowNodeInstances.find(x => x.M_bpmnElementId == bpmnElement.UUID) != undefined){
                event.preventDefault();
                event.stopPropagation();
                
            	// Get the menu position
            	var x = event.offsetX - 5;     // Get the horizontal coordinate
            	var y = event.offsetY - 5;     // 
                
                //Menu item for the log info
                var getLogInfoItem = new MenuItem("loginfo_startevent", "Info", {}, 0, function (diagramElement,parent,instance,diagramElement) {
                    return function () {
                        
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
                        				    drawLogInfo(caller.logInfos, caller.index + 1,caller.parent)
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
                }(diagramElement,parent,instance,diagramElement), "fa fa-book")
    
            	var popup = parent.appendElement({ "tag": "div", "style": "top:" + y + "px; left:" + x + "px;" }).down()
            	var contextMenu = new PopUpMenu(group, [ getLogInfoItem], event)
            }
    	    
        }
	  
	}(instance)
	
	//Override the click function on the end event
	instanceListView.svgDiagram.endEventClick = function(instance){
        return function(parent, diagramElement,group,svgDiagram,event,bpmnElement){
    	    //Check that the entity actually exists first
            if(instance.M_flowNodeInstances.find(x => x.M_bpmnElementId == bpmnElement.UUID) != undefined){
                event.preventDefault();
                event.stopPropagation();
                
            	// Get the menu position
            	var x = event.offsetX - 5;     // Get the horizontal coordinate
            	var y = event.offsetY - 5;     // 
                
                //Menu item for the log info
                var getLogInfoItem = new MenuItem("loginfo_stopevent", "Info", {}, 0, function (diagramElement,parent,instance,diagramElement) {
                    return function () {
                        
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
                        				    drawLogInfo(caller.logInfos, caller.index + 1,caller.parent)
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
                }(diagramElement,parent,instance,diagramElement), "fa fa-book")
    
            	var popup = parent.appendElement({ "tag": "div", "style": "top:" + y + "px; left:" + x + "px;" }).down()
            	var contextMenu = new PopUpMenu(group, [ getLogInfoItem], event)
            }
        }
	  
	}(instance)
	
	//Override the click function of a generic task
	instanceListView.svgDiagram.taskClick = function(instance){
        return function(parent, diagramElement,group,svgDiagram,event,bpmnElement){
            //Check that the entity actually exists first
            if(instance.M_flowNodeInstances.find(x => x.M_bpmnElementId == bpmnElement.UUID) != undefined){
                event.preventDefault();
                event.stopPropagation();
                
            	// Get the menu position
            	var x = event.offsetX - 5;     // Get the horizontal coordinate
            	var y = event.offsetY - 5;     // 
                
                //Menu item for the log info
                var getLogInfoItem = new MenuItem("loginfo_task", "Info", {}, 0, function (diagramElement,parent,instance) {
                    return function () {
                        
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
                            				    drawLogInfo(caller.logInfos, caller.index + 1,caller.parent)
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
                }(diagramElement,parent,instance), "fa fa-book")
    
            	var popup = parent.appendElement({ "tag": "div", "style": "top:" + y + "px; left:" + x + "px;" }).down()
            	var contextMenu = new PopUpMenu(group, [ getLogInfoItem], event)
            }
    	    
        }
	  
	}(instance)
	
	//Override the click function of a user task
	instanceListView.svgDiagram.userTaskClick = function(instance){
        return function(parent, diagramElement,group,svgDiagram,event,bpmnElement){
            //Check that the entity actually exists first
            if(instance.M_flowNodeInstances.find(x => x.M_bpmnElementId == bpmnElement.UUID) != undefined){
                event.preventDefault();
                event.stopPropagation();
                
            	// Get the menu position
            	var x = event.offsetX - 5;     // Get the horizontal coordinate
            	var y = event.offsetY - 5;     // 
                
                //Menu item for the log info
                var getLogInfoItem = new MenuItem("loginfo_usertask", "Info", {}, 0, function (diagramElement,parent,instance) {
                    return function () {
                        
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
                        				    drawLogInfo(caller.logInfos, caller.index + 1,caller.parent)
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
                }(diagramElement,parent,instance), "fa fa-book")
                
                //Menu item to play the user task
                var playItem = new MenuItem("play_activity", "Play", {}, 0, function(diagramElement,parent,instance,bpmnElement){
                    return function() {
                        var activityInstance;
                        if(instance.M_flowNodeInstances.find(x => x.M_bpmnElementId == bpmnElement.UUID) != undefined){
                            activityInstance = instance.M_flowNodeInstances.find(x => x.M_bpmnElementId == bpmnElement.UUID)
                            
                        }
                        new UserTaskWizard(parent, bpmnElement,activityInstance)
                    }
                }(diagramElement,parent,instance,bpmnElement), "fa fa-play")
                
                //Menu item to cancel the user task
                var cancelItem = new MenuItem("cancel_activity", "Cancel", {}, 0, function(){}, "fa fa-close")
    
            	var popup = parent.appendElement({ "tag": "div", "style": "top:" + y + "px; left:" + x + "px;" }).down()
            	var contextMenu = new PopUpMenu(group, [playItem, cancelItem,getLogInfoItem], event)
            }
    	    
        }
	  
	}(instance)
	
	//Override the click function of a script task
	instanceListView.svgDiagram.scriptTaskClick = function(instance){
        return function(parent, diagramElement,group,svgDiagram,event,bpmnElement){
            //Check that the entity actually exists first
            if(instance.M_flowNodeInstances.find(x => x.M_bpmnElementId == bpmnElement.UUID) != undefined){
                event.preventDefault();
                event.stopPropagation();
                
            	// Get the menu position
            	var x = event.offsetX - 5;     // Get the horizontal coordinate
            	var y = event.offsetY - 5;     // 
                
                //Menu item for the log info
                var getLogInfoItem = new MenuItem("loginfo_scriptask", "Info", {}, 0, function (diagramElement,parent,instance) {
                    return function () {
                        
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
                        				    drawLogInfo(caller.logInfos, caller.index + 1,caller.parent)
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
                }(diagramElement,parent,instance), "fa fa-book")
                
                //Menu item for the log info
                var getDataInputInfoItem = new MenuItem("data_objectreference", "Data", {}, 0, function (diagramElement,parent,instance) {
                    return function () {
                       
                        if(instance.M_flowNodeInstances.find(x => x.M_bpmnElementId == bpmnElement.UUID) != undefined){
                            data = instance.M_flowNodeInstances.find(x => x.M_bpmnElementId == bpmnElement.UUID)
                            
                            loadDataDialog(data.M_data,parent,diagramElement,false)
                            
                        }
                        
                        
                    }
                }(diagramElement,parent,instance), "fa fa-database")
    
            	var popup = parent.appendElement({ "tag": "div", "style": "top:" + y + "px; left:" + x + "px;" }).down()
            	var contextMenu = new PopUpMenu(group, [ getLogInfoItem,getDataInputInfoItem], event)
            }
    	    
        }
	  
	}(instance)
	
	//Override the click function on the end event
	instanceListView.svgDiagram.gatewayClick = function(instance){
        return function(parent, diagramElement,group,svgDiagram,event,bpmnElement){
            //Check that the entity actually exists first
            if(instance.M_flowNodeInstances.find(x => x.M_bpmnElementId == bpmnElement.UUID) != undefined){
                event.preventDefault();
                event.stopPropagation();
                
            	// Get the menu position
            	var x = event.offsetX - 5;     // Get the horizontal coordinate
            	var y = event.offsetY - 5;     // 
                
                //Menu item for the log info
                var getLogInfoItem = new MenuItem("loginfo_gateway", "Info", {}, 0, function (diagramElement,parent,instance) {
                    return function () {
                        
                        
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
                        				    drawLogInfo(caller.logInfos, caller.index + 1,caller.parent)
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
                }(diagramElement,parent,instance), "fa fa-book")
    
            	var popup = parent.appendElement({ "tag": "div", "style": "top:" + y + "px; left:" + x + "px;" }).down()
            	var contextMenu = new PopUpMenu(group, [ getLogInfoItem], event)
            }
    	    
        }
	  
	}(instance)
	
	//Attach a listener for updated entities in the process instance
    server.entityManager.attach(instance, NewEntityEvent, function (svgDiagram,instance) {
        return function(evt, element){
             var entity = evt.dataMap.entity 
            if (entity  !== undefined) {
                if(entity.ParentUuid == instance.UUID && entity.TYPENAME == "BPMS.ConnectingObject"){
                    console.log(entity)
                    var element = document.getElementsByName(entity.M_bpmnElementId)[0]
                    element.firstElementChild.classList.add("passedConnector")
                    document.getElementById("sequenceflow-end").firstElementChild.classList.add("passedConnector")
                    
                }
            }
        }
    }(instanceListView.svgDiagram,instance))
	
	//Attach a listener for updated entities in the process instance
    server.entityManager.attach(instance, UpdateEntityEvent, function (svgDiagram,instance) {
        return function(evt, element){
             var entity = evt.dataMap.entity 
            if (entity  !== undefined) {
                if(entity.ParentUuid == instance.UUID){
                   var element = document.getElementsByName(entity.M_bpmnElementId)[0]
                    //Lifecycle state handler, sets the CSS class to the equivalent state
                    switch(entity.M_lifecycleState){
                        case 1:
                            element.parentElement.classList.add("completed")
                            element.parentElement.classList.remove("ready")
                            element.parentElement.classList.remove("active")
                            
                            if(entity.M_bpmnElementId.startsWith("BPMN20.UserTask")){
                                //Override the click function of a user task when it has been completed
                            	instanceListView.svgDiagram.userTaskClick = function(instance){
                                    return function(parent, diagramElement,group,svgDiagram,event,bpmnElement){
                                        //Check that the entity actually exists first
                                        if(instance.M_flowNodeInstances.find(x => x.M_bpmnElementId == bpmnElement.UUID) != undefined){
                                            event.preventDefault();
                                            event.stopPropagation();
                                            
                                        	// Get the menu position
                                        	var x = event.offsetX - 5;     // Get the horizontal coordinate
                                        	var y = event.offsetY - 5;     // 
                                            
                                            //Menu item for the log info
                                            var getLogInfoItem = new MenuItem("loginfo_startevent", "Info", {}, 0, function (diagramElement,parent,instance) {
                                                return function () {
                                                    
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
                                                    				    drawLogInfo(caller.logInfos, caller.index + 1,caller.parent)
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
                                            }(diagramElement,parent,instance), "fa fa-book")
                                        
                                            //Menu item for the log info
                                            var getDataInputInfoItem = new MenuItem("data_objectreference", "Data", {}, 0, function (diagramElement,parent,instance) {
                                                return function () {
                                                   
                                                    if(instance.M_flowNodeInstances.find(x => x.M_bpmnElementId == bpmnElement.UUID) != undefined){
                                                        data = instance.M_flowNodeInstances.find(x => x.M_bpmnElementId == bpmnElement.UUID)
                                                        
                                                        loadDataDialog(data.M_data,parent,diagramElement,false)
                                                        
                                                    }
                                                    
                                                    
                                                }
                                            }(diagramElement,parent,instance), "fa fa-database")
                                
                                        	var popup = parent.appendElement({ "tag": "div", "style": "top:" + y + "px; left:" + x + "px;" }).down()
                                        	var contextMenu = new PopUpMenu(group, [ getLogInfoItem,getDataInputInfoItem], event)
                                        }
                                	    
                                    }
                        	  
                        	    }(instance)
                            }
                        
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
        }
    }(instanceListView.svgDiagram,instance))
   
    for(var i = 0; i < instance.M_flowNodeInstances.length; i++){
        var element = document.getElementsByName(instance.M_flowNodeInstances[i].M_bpmnElementId)[0]
        element.flowNodeInstance = instance.M_flowNodeInstances[i]
        element.instanceListView = instanceListView
	    
        //Lifecycle state handler, sets the CSS class to the equivalent state
        switch(instance.M_flowNodeInstances[i].M_lifecycleState){
            
            case 1:
                
                element.parentElement.classList.add("completed")
                element.parentElement.classList.remove("ready")
                element.parentElement.classList.remove("active")
                
                if(instance.M_flowNodeInstances[i].M_bpmnElementId.startsWith("BPMN20.UserTask") ){
                    //Override the click function of a user task when it has been completed
                	instanceListView.svgDiagram.userTaskClick = function(instance){
                        return function(parent, diagramElement,group,svgDiagram,event,bpmnElement){
                            //Check that the entity actually exists first
                            if(instance.M_flowNodeInstances.find(x => x.M_bpmnElementId == bpmnElement.UUID) != undefined){
                                event.preventDefault();
                                event.stopPropagation();
                                
                            	// Get the menu position
                            	var x = event.offsetX - 5;     // Get the horizontal coordinate
                            	var y = event.offsetY - 5;     // 
                                
                                //Menu item for the log info
                                var getLogInfoItem = new MenuItem("loginfo_startevent", "Info", {}, 0, function (diagramElement,parent,instance) {
                                    return function () {
                                        
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
                                        				    drawLogInfo(caller.logInfos, caller.index + 1,caller.parent)
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
                                }(diagramElement,parent,instance), "fa fa-book")
                            
                                //Menu item for the log info
                                var getDataInputInfoItem = new MenuItem("data_objectreference", "Data", {}, 0, function (diagramElement,parent,instance) {
                                    return function () {
                                       
                                        if(instance.M_flowNodeInstances.find(x => x.M_bpmnElementId == bpmnElement.UUID) != undefined){
                                            data = instance.M_flowNodeInstances.find(x => x.M_bpmnElementId == bpmnElement.UUID)
                                            
                                            loadDataDialog(data.M_data,parent,diagramElement,false)
                                            
                                        }
                                        
                                        
                                    }
                                }(diagramElement,parent,instance), "fa fa-database")
                    
                            	var popup = parent.appendElement({ "tag": "div", "style": "top:" + y + "px; left:" + x + "px;" }).down()
                            	var contextMenu = new PopUpMenu(group, [ getLogInfoItem,getDataInputInfoItem], event)
                            }
                    	    
                        }
            	  
            	    }(instance)
                }
                
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
    
    for(var i = 0; i < instance.M_connectingObjects.length; i++){
        var element = document.getElementsByName(instance.M_connectingObjects[i].M_bpmnElementId)[0]
        element.firstElementChild.classList.add("passedConnector")
        document.getElementById("sequenceflow-end").firstElementChild.classList.add("passedConnector")
    }
   
}


//////////////////////////////////////////////////////////////////////////////////////////////////////
// Load and display a data instance by creating a dialog
//////////////////////////////////////////////////////////////////////////////////////////////////////
function loadDataDialog(data, parent, diagramElement,decoded){
    
    //Initialize the dialog element and set its position
    var dialog = new Dialog(randomUUID(),parent,false, diagramElement.M_id)
    var x = diagramElement.M_Bounds.M_x
	var y = diagramElement.M_Bounds.M_y + diagramElement.M_Bounds.M_height + 5
	dialog.content.appendElement({"tag" : "table", "id" : "dataInfosTable", "class" : "instanceListTable"}).down()
    dialog.setPosition(x, y)
    
    
    //If the data we received is not part of an array, just force it into one so we can iterate it
	if(data.length == undefined){
	    data = [data]
	}
	
	
	//Iterate through the data array
    for(var i = 0; i < data.length; i++){
        //Only decode from 64 bit if specified
        if(decoded == false){
            if(decode64(data[i].M_data).startsWith("[") && decode64(data[i].M_data).endsWith("]")){
                var objects = JSON.parse(decode64(data[i].M_data))
                if(objects.length != 0){
                    for(var j = 0; j < objects.length; j++){
                        var entity = eval("new " + objects[j].TYPENAME + "()")
                        entity.init(objects[j],false)
                        objects[j] = entity
                    }
                    var div = dialog.content.getChildById("dataInfosTable").appendElement({"tag" : "tr"}).down()
                        	.appendElement({"tag" : "th","class" : "borderHeader", "innerHtml" : data[i].M_name})
                        	.appendElement({"tag" : "th", "class" : "borderHeader", "style" : "overflow-y: auto;max-height: 350px;display: block;" }).down()
        			var table = new Table(randomUUID(), div)
        
        			var model = new EntityTableModel(getEntityPrototype(objects[0].TYPENAME))
        			
        			model.entities = objects
        			
        			var callback = function(){}
        
        			table.setModel(model,
        				function (table, callback) {
        					return function () {
        						table.init()
        						table.refresh()
        						table.header.maximizeBtn.element.click()
        					}
        				}(table, callback))
                }
                
            }else{
                if(typeof decode64(data[i].M_data) != "string"){
                    data[i].M_data= JSON.stringify(decode64(data[i].M_data))
                }
                //Verify if the data is contained inside a map
                if(decode64(data[i].M_data).startsWith("{") && decode64(data[i].M_data).endsWith("}")){
                    //Try to parse it into a JSON object
                    if(JSON.parse(decode64(data[i].M_data)).UUID != undefined){
                        //Create an entity using the JSON we parsed
                        var entityData = JSON.parse(decode64(data[i].M_data))
                        var entity = eval("new " + entityData.TYPENAME + "()")
                        entity.init(entityData,false)
                        var entityPanelParent = dialog.content.getChildById("dataInfosTable").appendElement({"tag" : "tr"}).down()
                    	.appendElement({"tag" : "th","class" : "borderHeader", "innerHtml" : data[i].M_name})
                    	.appendElement({"tag" : "th", "class" : "borderHeader", "style" : "overflow-y: auto;max-height: 350px;display: block;" }).down()
                        //Create an entity panel with the entity we loaded
                        new EntityPanel(entityPanelParent,entity.TYPENAME,
                            function(entity){
                               return function(panel){
                                 panel.setEntity(entity)
                                } 
                            }(entity))
                            
                    }
                }else{
                    //Verify if the data is an UUID by splitting the string
                    var prototype = ""
                    if(decode64(data[i].M_data).split(".").length > 1){
                        if(decode64(data[i].M_data).split(".")[1].split("%").length > 1){
                            prototype = "CargoEntities." + decode64(data[i].M_data).split(".")[1].split("%")[0]
                        }
                    }
                    
                    //Try to get the entity prototype from the string we splitted
                    server.entityManager.getEntityPrototype(prototype ,false,
                        //If we enter this function, the string was an UUID, therefore we can load an entity panel
                        function(result,caller){
                            //Get the entity from the server
                            server.entityManager.getEntityByUuid(caller.entityUUID, false,
                                function(entity,caller){
                                     var entityPanelParent = caller.dialog.content.getChildById("dataInfosTable").appendElement({"tag" : "tr"}).down()
                                	.appendElement({"tag" : "th","class" : "borderHeader", "innerHtml" : caller.data.M_name})
                                	.appendElement({"tag" : "th", "class" : "borderHeader", "style" : "overflow-y: auto;max-height: 350px;max-width:600px;display: block;" }).down()
                                	//Create the entity panel from the entity we loaded
                                    new EntityPanel(entityPanelParent,entity.TYPENAME,
                                        function(entity){
                                            return function(panel){
                                                panel.setEntity(entity)
                                            } 
                                        }(entity))
                                },function(){}, {"dialog" : caller.dialog, "data" : caller.data})
                          
                        },
                        //If we enter this function, the server cannot get the entity prototype, so the string is not an UUID
                        function(error,caller){
                   
                            //Simply inject data into the table
                            dialog.content.getChildById("dataInfosTable").appendElement({"tag" : "tr"}).down()
                        	.appendElement({"tag" : "th","class" : "borderHeader", "innerHtml" : caller.data.M_name})
                        	.appendElement({"tag" : "th", "class" : "borderHeader", "innerHtml" : decode64(caller.data.M_data)})
                        },{"data" : data[i], "dialog" : dialog, "entityUUID" : decode64(data[i].M_data)})
                } 
            }
             
            
        }else{
            if(typeof data[i].M_data != "string"){
                data[i].M_data = JSON.stringify(data[i].M_data)
            }
            //Verify if the data is contained inside a map
            if(data[i].M_data.startsWith("{") && data[i].M_data.endsWith("}")){
                //Try to parse it into a JSON object
                if(JSON.parse(data[i].M_data).UUID != undefined){
                    //Create an entity using the JSON we parsed
                    var entityData = JSON.parse(data[i].M_data)
                    var entity = eval("new " + entityData.TYPENAME + "()")
                    entity.init(entityData,false)
                    var entityPanelParent = dialog.content.getChildById("dataInfosTable").appendElement({"tag" : "tr"}).down()
                	.appendElement({"tag" : "th","class" : "borderHeader", "innerHtml" : data[i].M_name})
                	.appendElement({"tag" : "th", "class" : "borderHeader", "style" : "overflow-y: auto;max-height: 350px;display: block;" }).down()
                    //Create an entity panel with the entity we loaded
                    new EntityPanel(entityPanelParent,entity.TYPENAME,
                        function(entity){
                           return function(panel){
                             panel.setEntity(entity)
                            } 
                        }(entity))
                        
                }
            }else{
                //Verify if the data is an UUID by splitting the string
                var prototype = ""
                if(data[i].M_data.split(".").length > 1){
                    if(data[i].M_data.split(".")[1].split("%").length > 1){
                        prototype = "CargoEntities." + data[i].M_data.split(".")[1].split("%")[0]
                    }
                }
                
                //Try to get the entity prototype from the string we splitted
                server.entityManager.getEntityPrototype(prototype ,false,
                    //If we enter this function, the string was an UUID, therefore we can load an entity panel
                    function(result,caller){
                        //Get the entity from the server
                        server.entityManager.getEntityByUuid(caller.entityUUID, false,
                            function(entity,caller){
                                 var entityPanelParent = caller.dialog.content.getChildById("dataInfosTable").appendElement({"tag" : "tr"}).down()
                            	.appendElement({"tag" : "th","class" : "borderHeader", "innerHtml" : caller.data.M_name})
                            	.appendElement({"tag" : "th", "class" : "borderHeader", "style" : "overflow-y: auto;max-height: 350px;max-width:600px;display: block;" }).down()
                            	//Create the entity panel from the entity we loaded
                                new EntityPanel(entityPanelParent,entity.TYPENAME,
                                    function(entity){
                                       return function(panel){
                                         panel.setEntity(entity)
                                        } 
                                    }(entity))
                            },function(){}, {"dialog" : caller.dialog, "data" : caller.data})
                      
                    },
                    //If we enter this function, the server cannot get the entity prototype, so the string is not an UUID
                    function(error,caller){
               
                        //Simply inject data into the table
                        dialog.content.getChildById("dataInfosTable").appendElement({"tag" : "tr"}).down()
                    	.appendElement({"tag" : "th","class" : "borderHeader", "innerHtml" : caller.data.M_name})
                    	.appendElement({"tag" : "th", "class" : "borderHeader", "innerHtml" : caller.data.M_data})
                    },{"data" : data[i], "dialog" : dialog, "entityUUID" : data[i].M_data})
            }
        }
        
	}
}


//////////////////////////////////////////////////////////////////////////////////////////////////////
// Load the history of logs into the given parent, including the actor, the action and the date
//////////////////////////////////////////////////////////////////////////////////////////////////////
 function loadLogInfosDialog(logInfos, parent, diagramElement){
     
    logInfos.sort(function(a,b){
        return a.M_date - b.M_date
    })
     
    var dialog = new Dialog(randomUUID(),parent,false, "Infos")
    
    var x = diagramElement.M_Bounds.M_x
	var y = diagramElement.M_Bounds.M_y + diagramElement.M_Bounds.M_height + 5
	
	dialog.content.appendElement({"tag" : "table", "id" : "logInfosTable", "class" : "instanceListTable"}).down()
	.appendElement({"tag" : "tr"}).down()
	.appendElement({"tag" : "th","class" : "borderHeader", "innerHtml" : "Date"})
	.appendElement({"tag" : "th","class" : "borderHeader", "innerHtml" : "Actor"})
	.appendElement({"tag" : "th","class" : "borderHeader","innerHtml" : "Action"})
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
        var dialogContent = dialog.content.getChildById("logInfosTable").appendElement({"tag" : "tr"}).down()
        server.entityManager.getEntityByUuid(logInfos[i].M_actor,false,
            function(account,caller){
                if(account.M_name != ""){
                    
                    caller.dialogContent.appendElement({"tag" : "th", "class" : "borderHeader", "innerHtml" : caller.date})
                    .appendElement({"tag" : "th", "class" : "borderHeader", "innerHtml" : account.M_name})
                    .appendElement({"tag" : "th", "class" : "borderHeader", "innerHtml" : caller.action}) 
                }else{
                    caller.dialogContent.appendElement({"tag" : "th", "class" : "borderHeader", "innerHtml" : caller.date})
                    .appendElement({"tag" : "th", "class" : "borderHeader", "innerHtml" : account.M_id})
                    .appendElement({"tag" : "th", "class" : "borderHeader", "innerHtml" : caller.action})
                }
               
            },function(){},{"dialogContent" : dialogContent, "date" : date, "action" : logInfos[i].M_action})
    }
    
}

//////////////////////////////////////////////////////////////////////////////////////////////////////
// Load the generic bpmn element info
//////////////////////////////////////////////////////////////////////////////////////////////////////
function loadBpmnElementInfo(parent, diagramElement,group,svgDiagram,event,bpmnElement){
     event.preventDefault();
    event.stopPropagation();
    
	// Get the menu position
	var x = event.offsetX - 5;     // Get the horizontal coordinate
	var y = event.offsetY - 5;     // 
    
    var getInfoItem = new MenuItem("bpmninfo", "BPMN Element", {}, 0, function (bpmnElement,parent,diagramElement) {
            return function () {
                showBpmnInfo(bpmnElement,parent,diagramElement)
            }
    }(bpmnElement,parent,diagramElement), "fa fa-cube")
    
    var popup = parent.appendElement({ "tag": "div", "style": "top:" + y + "px; left:" + x + "px;" }).down()
	var contextMenu = new PopUpMenu(group, [getInfoItem], event)
    
    function showBpmnInfo(bpmnElement,parent,diagramElement){
         //Initialize the dialog element and set its position
        var dialog = new Dialog(randomUUID(),parent,false, diagramElement.M_id)
        var x = diagramElement.M_Bounds.M_x
    	var y = diagramElement.M_Bounds.M_y + diagramElement.M_Bounds.M_height + 5
    	dialog.content.appendElement({"tag" : "table", "id" : "dataInfosTable", "class" : "instanceListTable"}).down()
        dialog.setPosition(x, y)
        
         var entityPanelParent = dialog.content.getChildById("dataInfosTable").appendElement({"tag" : "tr"}).down()
    	.appendElement({"tag" : "th","class" : "borderHeader", "innerHtml" : bpmnElement.M_name})
    	.appendElement({"tag" : "th", "class" : "borderHeader", "style" : "overflow-y: auto;max-height: 350px;max-width:600px;display: block;" }).down()
    	//Create the entity panel from the entity we loaded
        new EntityPanel(entityPanelParent,bpmnElement.TYPENAME,
            function(bpmnElement){
               return function(panel){
                 panel.setEntity(bpmnElement)
                } 
            }(bpmnElement))
        
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
	this.dataView = new BpmnDataView(this.content, [], startEvent.M_dataOutput)
	this.dialog.ok.element.onclick = function (dataView, process, dialog, startEvent) {
		return function () {
			// I will close the dialogue first...
			dialog.close()

			// I will save the data view...
			dataView.save(function (process, startEvent) {
				return function (itemAwareInstances) {
                    // Create a process trigger event...
                    var trigger = new BPMS.Trigger()
                    trigger.M_id = startEvent.UUID
                    trigger.M_eventTriggerType = 14//BPMS.EventTriggerType_Start
                    trigger.M_sessionId = server.sessionId
                    trigger.M_processUUID = process.UUID
                    trigger.UUID = "BPMS.Trigger%" + randomUUID()
                    trigger.M_data = itemAwareInstances
    
                    var evtData = {"TYPENAME":"Server.MessageData", "Name":"trigger", "Value":trigger}
                    server.eventManager.broadcastNetworkEvent(StartProcessEvent, WorkflowEvent, [evtData], function(){}, function(){}, undefined) 
				}
			} (process, startEvent))
			
		}
	} (this.dataView, startEvent.getParent(), this.dialog, startEvent)

	return this
}

var UserTaskWizard = function (parent, userTask,instance) {
    this.parent = parent
	this.id = randomUUID()

	// The wizard dialog...
	this.dialog = new Dialog(this.id, this.parent, false, "User Task")

	// Set the dialog position...
	var diagramElement = userTask.getDiagramElement()
	var x = diagramElement.M_Bounds.M_x
	var y = diagramElement.M_Bounds.M_y + diagramElement.M_Bounds.M_height + 5
	this.dialog.setPosition(x, y)

	this.content = this.dialog.content.appendElement({ "tag": "div", "class": "process_wizard_content" }).down()

	// That will contain the values ask by the user...
	this.dataView = new BpmnDataView(this.content, userTask.M_ioSpecification.M_dataInput, userTask.M_ioSpecification.M_dataOutput, userTask.M_property, instance)
	
	this.dialog.ok.element.onclick = function (dataView, dialog) {
		return function () {
			// I will close the dialogue first...
			dialog.close()

			// I will save the data view...
			dataView.save(function (itemAwareElements, activity) {
    			server.workflowProcessor.runActivity(activity,
				    function(success,caller){
				        console.log(success)
				    },function(){},{})
			})
			
		}
	} (this.dataView, this.dialog)
	
	
	
	return this

}


