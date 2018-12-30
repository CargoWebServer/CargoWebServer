var AdminLocalisationPage = function(panel){
    /** The tab content panel **/
    this.panel = panel
    
    /** The panel sections. **/
    this.panel.appendElement({"tag" : "div", "class" : "row", "style" : "height:85vh;margin:0;"}).down()
        .appendElement({"tag" : "div", "class" : "col-md-3 bg-dark", "id" : "localisationsAdminNavigation"}).down()
        .appendElement({"tag" : "div", "class" : "input-group mb-3", "style":"padding:15px;"}).down()
        .appendElement({"tag" : "span", "class" : "input-group-prepend"}).down()
        .appendElement({"tag":"span", "class":"input-group-text"}).down()
        .appendElement({"tag" : "i", "class" : "fa fa-search text-dark"}).up().up()
        .appendElement({"tag" : "input", "type" : "text", "class" : "form-control", "placeholder" : "Filtrer", "id" : "localisationFilterKeywords"}).up()
        .appendElement({"tag" : "div", "class" : "list-group","role" :"tablist", "id" : "localisationsFilteredList"})
        .appendElement({"tag" : "div", "style" : "position:absolute;bottom:0; right:0;margin:30px;"}).down()
        .appendElement({"tag" : "button", "class" : "btn btn-success", "id" : "addLocalisationButton"}).down()
        .appendElement({"tag" : "i", "class" : "fa fa-plus"}).up().up().up()
        .appendElement({"tag" : "div", "class" : "col-md-9 bg-light"}).down()
         .appendElement({"tag" : "div", "class" : "tab-content", "id" : "localisationsAdminControl"}).up().up().up()
        

    this.modifiedItems = {}
  
    this.removedItems = {}
    
    this.panel.getChildById("localisationFilterKeywords").element.addEventListener('change', function(){
        results = mainPage.adminPage.adminLocalisationPage.getLocalisationsFromKeyword(this.value.toUpperCase())
    })
    
    
    this.panel.getChildById("addLocalisationButton").element.onclick = function(){
        mainPage.adminPage.adminLocalisationPage.panel.getChildById("localisationsFilteredList").removeAllChilds()
        mainPage.adminPage.adminLocalisationPage.panel.getChildById("localisationsAdminControl").removeAllChilds()
        mainPage.adminPage.adminLocalisationPage.panel.getChildById("localisationsFilteredList").appendElement({"tag":"a", "class" : "list-group-item list-group-item-action active show","id" : "newLocalisation-selector","data-toggle" : "tab","href" : "#"+ "newLocalisation-control","role":"tab", "innerHtml": "Nouvelle localisation","aria-controls": "newLocalisation-control"})
         mainPage.adminPage.adminLocalisationPage.panel.getChildById("localisationsAdminControl").appendElement({"tag" : "div", "class" : "tab-pane active show", "id" : "newLocalisation-control", "role" : "tabpanel", "aria-labelledby" : "newLocalisation-selector", "style" : "padding:15px;"}).down()
        .appendElement({"tag" : "div", "class" : "input-group mb-3"}).down()
        .appendElement({"tag" : "div", "class" : "input-group-prepend"}).down()
        .appendElement({"tag" : "span", "class" : "input-group-text", "innerHtml" : "ID"}).up()
        .appendElement({"tag" : "input", "class" : "form-control",  "type" : "text", "id" : "newLocalisation-localisationID"}).up()
        .appendElement({"tag" : "div", "class" : "input-group mb-3"}).down()
        .appendElement({"tag" : "div", "class" : "input-group-prepend"}).down()
        .appendElement({"tag" : "span", "class" : "input-group-text", "innerHtml" : "Nom"}).up()
        .appendElement({"tag" : "input", "class" : "form-control", "type" : "text", "id" : "newLocalisation-name"}).up()
        .appendElement({"tag" : "div", "class" : "input-group mb-3"}).down()
        .appendElement({"tag" : "div", "class" : "input-group-prepend"}).down()
        .appendElement({"tag" : "span", "class" : "input-group-text", "innerHtml" : "Parent"}).up()
        .appendElement({"tag" : "input", "class" : "form-control", "type" : "text", "id" : "newLocalisation-parent"}).up()
        .appendElement({"tag" : "button", "class" : "btn btn-primary mr-3", "innerHtml" : "Enregistrer", "id" : "newLocalisation-saveBtn"})
        
        
        autocomplete("CatalogSchema.LocalisationType",mainPage.adminPage.adminLocalisationPage.panel.getChildById("newLocalisation-parent"))
        
        mainPage.adminPage.adminLocalisationPage.panel.getChildById("newLocalisation-saveBtn").element.onclick = function(){
            document.getElementById("waitingDiv").style.display = ""
            var id = mainPage.adminPage.adminLocalisationPage.panel.getChildById("newLocalisation-localisationID").element.value
            var name = mainPage.adminPage.adminLocalisationPage.panel.getChildById("newLocalisation-name").element.value
            var parent = mainPage.adminPage.adminLocalisationPage.panel.getChildById("newLocalisation-parent").element.value
            var newLocalisation = new CatalogSchema.LocalisationType
            newLocalisation.M_id = id
            newLocalisation.M_name = name
            if(parent != ""){
                  var q = new EntityQuery()
                    q.TypeName = "CatalogSchema.LocalisationType"
                    q.Fields = ["M_id"]
                    q.Query = 'CatalogSchema.LocalisationType.M_id=="'+ parent +'"'
                    
                    server.entityManager.getEntities("CatalogSchema.LocalisationType", "CatalogSchema", q, 0, -1, [], true, true,
                        function(index,total,caller){},
                        function(localisations,caller){
                            caller.newLocalisation.M_parent = localisations[0]
                            server.entityManager.createEntity(catalog.UUID, "M_localisations", caller.newLocalisation,
                            function(success,caller){
                                caller.parentLocalisation.M_subLocalisations.push(success.UUID)
                                server.entityManager.saveEntity(caller.parentLocalisation,
                                    function(success,caller){
                                        document.getElementById("waitingDiv").style.display = "none"
                                        mainPage.adminPage.adminLocalisationPage.panel.getChildById("localisationsFilteredList").removeAllChilds()
                                        mainPage.adminPage.adminLocalisationPage.panel.getChildById("localisationsAdminControl").removeAllChilds()
                                        mainPage.adminPage.adminLocalisationPage.panel.getChildById("localisationsFilteredList").appendElement({"tag":"a", "class" : "list-group-item list-group-item-action","id" : caller.newLocalisation.M_id + "-selector","data-toggle" : "tab","href" : "#"+ caller.newLocalisation.M_id + "-control","role":"tab", "innerHtml": caller.newLocalisation.M_id,"aria-controls":  caller.newLocalisation.M_id + "-control"})
                                        mainPage.adminPage.adminLocalisationPage.loadAdminControl(caller.newLocalisation)
                                    }, function(){},{"newLocalisation" : success})
                                
                            },function(){},{"parentLocalisation" : localisations[0]})
                        },
                        function(){}, {"newLocalisation" : newLocalisation})
            
            }else{
                newLocalisation.M_parent = ""
                server.entityManager.createEntity(catalog.UUID, "M_localisations", newLocalisation,
                    function(success,caller){
                        
              
                                document.getElementById("waitingDiv").style.display = "none"
                                mainPage.adminPage.adminLocalisationPage.panel.getChildById("localisationsFilteredList").removeAllChilds()
                                mainPage.adminPage.adminLocalisationPage.panel.getChildById("localisationsAdminControl").removeAllChilds()
                                mainPage.adminPage.adminLocalisationPage.panel.getChildById("localisationsFilteredList").appendElement({"tag":"a", "class" : "list-group-item list-group-item-action","id" : caller.newLocalisation.M_id + "-selector","data-toggle" : "tab","href" : "#"+ caller.newLocalisation.M_id + "-control","role":"tab", "innerHtml": caller.newLocalisation.M_id,"aria-controls":  caller.newLocalisation.M_id + "-control"})
                                mainPage.adminPage.adminLocalisationPage.loadAdminControl(success)
                            
                    },function(){},{"newLocalisation" : newLocalisation})
            }
           
            
       
            
            
        }
    }
    
    return this
}

AdminLocalisationPage.prototype.getLocalisationsFromKeyword = function(keyword){
    var q = new EntityQuery()
    q.TypeName = "CatalogSchema.LocalisationType"
    q.Fields = ["M_name","M_id"]
    q.Query = 'CatalogSchema.LocalisationType.M_id=="'+ keyword +'"'



 server.entityManager.getEntities("CatalogSchema.LocalisationType", "CatalogSchema", q, 0,-1,[],true,true, 
        function(index,total, caller){
            
        },
        function(localisations,caller){
        console.log(localisations)
            mainPage.adminPage.adminLocalisationPage.panel.getChildById("localisationsFilteredList").removeAllChilds()
            mainPage.adminPage.adminLocalisationPage.panel.getChildById("localisationsAdminControl").removeAllChilds()
           
            
           
            for(var i = 0; i < localisations.length; i++){
                mainPage.adminPage.adminLocalisationPage.panel.getChildById("localisationsFilteredList").appendElement({"tag":"a", "class" : "list-group-item list-group-item-action","id" : localisations[i].M_id + "-selector","data-toggle" : "tab","href" : "#"+ localisations[i].M_id + "-control","role":"tab", "innerHtml": localisations[i].M_id,"aria-controls":  localisations[i].M_id + "-control"})
                mainPage.adminPage.adminLocalisationPage.loadAdminControl(localisations[i])
               
                
            }
            
        },
        function(){
        },{})
}

AdminLocalisationPage.prototype.loadAdminControl = function(localisation){
    mainPage.adminPage.panel.getChildById("adminLocalisationSaveState").element.innerHTML = ""
    mainPage.adminPage.adminLocalisationPage.modifiedItems[localisation.M_id] = new Map()
    mainPage.adminPage.adminLocalisationPage.removedItems[localisation.M_id] = new Map()
     
    if(mainPage.adminPage.adminLocalisationPage.panel.getChildById(localisation.M_id + "-control") != undefined){
        mainPage.adminPage.adminLocalisationPage.panel.getChildById(localisation.M_id + "-control").delete()
    }
    mainPage.adminPage.adminLocalisationPage.panel.getChildById("localisationsAdminControl").appendElement({"tag" : "div", "class" : "tab-pane", "id" : localisation.M_id + "-control", "role" : "tabpanel", "aria-labelledby" : localisation.M_id + "-selector", "style" : "padding:15px;"}).down()
    .appendElement({"tag" : "div", "class" : "input-group mb-3"}).down()
    .appendElement({"tag" : "div", "class" : "input-group-prepend"}).down()
    .appendElement({"tag" : "span", "class" : "input-group-text", "innerHtml" : "ID"}).up()
    .appendElement({"tag" : "span", "class" : "form-control",  "innerHtml" : localisation.M_id}).up()
    .appendElement({"tag" : "div", "class" : "input-group mb-3"}).down()
    .appendElement({"tag" : "div", "class" : "input-group-prepend"}).down()
    .appendElement({"tag" : "span", "class" : "input-group-text", "innerHtml" : "Nom"}).up()
    .appendElement({"tag" : "input", "class" : "form-control", "type" : "text", "value" : localisation.M_name, "id" : localisation.M_id + "-name"}).up()
    .appendElement({"tag" : "div", "class" : "input-group mb-3"}).down()
    .appendElement({"tag" : "div", "class" : "input-group-prepend"}).down()
    .appendElement({"tag" : "span", "class" : "input-group-text", "innerHtml" : "Parent"}).up()
    .appendElement({"tag" : "input", "class" : "form-control", "id" : localisation.M_id + "-parent", "type" : "text"}).up()
    .appendElement({"tag" : "div", "class" : "row"}).down()
    .appendElement({"tag":"label", "for" : localisation.M_id + "-itemsList","class" : "col-3 d-flex align-items-center", "innerHtml" : "Sous-localisations"})
    .appendElement({"tag": "input", "class" : "form-control col-5 m-1", "id" : localisation.M_id + "-filterSubLocalisations","type" :"text", "placeholder" : "Filtrer.."}).up()
    .appendElement({"tag" : "div"}).down()
    .appendElement({"tag":"table", "class":"table table-hover table-condensed", "style" : "display: block;height: 30em;overflow: auto;"}).down()
    .appendElement({"tag":"thead"}).down()
    .appendElement({"tag" : "tr"}).down()
    .appendElement({"tag" : "th", "innerHtml" : "ID", "style" : "width: 30em;"})
    .appendElement({"tag" : "th", "innerHtml" : "Nom", "style" : "width:30em;"})
    .appendElement({"tag" : "th", "innerHtml" : ""}).down()
    .appendElement({"tag" : "button", "class" : "btn btn-success btn-sm", "id" : localisation.M_id +"-addSubLocalisationButton"}).down()
    .appendElement({"tag" : "i", "class" : "fa fa-plus"}).up().up()
    .up().up()
    .appendElement({"tag" : "tbody", "id" : localisation.M_id + "-itemsList"})
    .up().up()
    
    .appendElement({"tag" : "div", "class" : "d-flex justify-content-center"}).down()
    .appendElement({"tag" : "button", "class" : "btn btn-primary disabled mr-3", "innerHtml" : "Enregistrer", "id" : localisation.M_id + "-saveBtn"})
    .appendElement({"tag" : "button", "class" : "btn disabled mr-3", "innerHtml" : "Annuler les modifications", "id" : localisation.M_id + "-cancelBtn"})
    .appendElement({"tag" : "button", "class" : "btn btn-danger", "innerHtml" : "Supprimer l'entité", "id" : localisation.M_id + "-deleteBtn", "data-toggle" : "modal", "data-target" : "#" + localisation.M_id + "-modal"})
    .appendElement({"tag" : "div", "class" : "modal fade", "id"  : localisation.M_id + "-modal", "tabindex" : "-1", "role"  :"dialog", "aria-labelledby" : localisation.M_id + "-modal", "aria-hidden" : "true"}).down()
    .appendElement({"tag" : "div", "class" : "modal-dialog", "role" : "document"}).down()
    .appendElement({"tag" : "div", "class" : "modal-content"}).down()
    .appendElement({"tag" : "div", "class" : "modal-header"}).down()
    .appendElement({"tag" : "h5", "class" : "modal-title", "innerHtml" : "Supprimer la localisation"})
    .appendElement({"tag" : "button", "class" : "close", "data-dismiss" : "modal", "aria-label" : "Close"}).down()
    .appendElement({"tag" : "span", "aria-hidden" : "true", "innerHtml" : "&times;"}).up().up()
    .appendElement({"tag" : "div", "class" : "modal-body", "innerHtml" : "Êtes-vous certain de vouloir supprimer cette localisation? Cela supprimera aussi toutes les sous-localisations associées à celles-ci."})
    .appendElement({"tag" : "div", "class" : "modal-footer"}).down()
    .appendElement({"tag" : "button", "class" : "btn btn-secondary", "data-dismiss" : "modal", "innerHtml" : "Fermer", "id" : localisation.M_id + "-closeDelete"})
    .appendElement({"tag" : "button", "type" : "button", "class" : "btn btn-danger", "id" : localisation.M_id + "-deleteConfirmBtn", "innerHtml" : "Confirmer la suppression"})
    
    autocomplete("CatalogSchema.LocalisationType", mainPage.adminPage.adminLocalisationPage.panel.getChildById(localisation.M_id + "-parent"))
    
    
    
    mainPage.adminPage.adminLocalisationPage.panel.getChildById(localisation.M_id + "-deleteConfirmBtn").element.onclick = function(localisation){
        return function(){
            document.getElementById("waitingDiv").style.display = ""
            mainPage.adminPage.adminLocalisationPage.panel.getChildById(localisation.M_id + "-closeDelete").element.click()
            mainPage.adminPage.adminLocalisationPage.deleteSubLocalisations(localisation.M_subLocalisations)
            
                
            server.entityManager.getEntityByUuid(localisation.M_parent,false,
                function(parent,caller){
                    parent.M_subLocalisations.pop(caller.localisation)
                    server.entityManager.saveEntity(parent,
                        function(success,caller){
                             mainPage.adminPage.adminLocalisationPage.panel.getChildById(caller.localisation.M_id + "-control").delete()
                             mainPage.adminPage.adminLocalisationPage.panel.getChildById(caller.localisation.M_id + "-selector").delete()
                        },function(){},{"localisation" : caller.localisation})
                },function(){}, {"localisation" : localisation})
                
            server.entityManager.removeEntity(localisation.UUID,
            function(success,caller){
                document.getElementById("waitingDiv").style.display = "none"
            }, function(){},{})
        }
    }(localisation)
    
    server.entityManager.getEntityByUuid(localisation.M_parent, false, 
    function(parentLocalisation, caller){
        console.log(parentLocalisation.M_id)
        mainPage.adminPage.adminLocalisationPage.panel.getChildById(caller.localisationID + "-parent").element.value = parentLocalisation.M_id
    }, function(){}, {"localisationID" : localisation.M_id})
    
    
     mainPage.adminPage.adminLocalisationPage.panel.getChildById(localisation.M_id + "-filterSubLocalisations").element.onkeyup = function(localisationID){
        return function(){
            var input, filter, table, tr, td, i;
            input = document.getElementById(localisationID+"-filterSubLocalisations");
            filter = input.value.toUpperCase();
            table = document.getElementById(localisationID + "-itemsList");
            tr = table.getElementsByTagName("tr");
        
            // Loop through all table rows, and hide those who don't match the search query
            for (i = 0; i < tr.length; i++) {
                td = tr[i].getElementsByTagName("td")[1];
                if (td) {
                    if (td.innerHTML.toUpperCase().indexOf(filter) > -1) {
                        tr[i].style.display = "";
                    } else {
                        tr[i].style.display = "none";
                    }
                } 
            }
       }
    }(localisation.M_id)
    
    for(var i = 0; i < localisation.M_subLocalisations.length; i++){
        server.entityManager.getEntityByUuid(localisation.M_subLocalisations[i], false,
            function(subLocalisation, caller){
                mainPage.adminPage.adminLocalisationPage.appendSubLocalisation(subLocalisation, caller.localisationID)
            }, function() {}, {"localisationID" : localisation.M_id})
    }
    
    mainPage.adminPage.adminLocalisationPage.panel.getChildById(localisation.M_id + "-addSubLocalisationButton").element.onclick = function(localisation){
        return function(){
            var date = new Date()
            mainPage.adminPage.adminLocalisationPage.panel.getChildById(localisation.M_id + "-itemsList").prependElement({"tag" : "tr", "id" : date + "-row"}).down()
            .appendElement({"tag" : "th", "scope" : "row"}).down()
            .appendElement({"tag" : "input", "class" : "form-control", "type" : "text", "placeholder" : "ID", "id" : date  + "-id"}).up()
            .appendElement({"tag" : "td", "scope" : "row"}).down()
            .appendElement({"tag" : "input", "class" : "form-control", "type" : "text", "placeholder" : "ID", "id" : date  + "-name"}).up()
                     .appendElement({"tag":"td"}).down()
                    .appendElement({"tag" : "button", "class" : "btn btn-success btn-sm", "id" : date + "-confirmBtn"}).down()
                    .appendElement({"tag" : "i", "class" : "fa fa-check"}).up()
               
                    
            mainPage.adminPage.adminLocalisationPage.panel.getChildById(date+ "-confirmBtn").element.onclick = function (date, localisation) {
               return function(){
                    document.getElementById("waitingDiv").style.display = ""
                    var localisationID = mainPage.adminPage.adminLocalisationPage.panel.getChildById(date+"-id").element.value
                    var localisationName = mainPage.adminPage.adminLocalisationPage.panel.getChildById(date+"-name").element.value
                    var newItem = new CatalogSchema.LocalisationType()
            
                    newItem.M_id = localisationID
                    newItem.M_name = localisationName
          
                    newItem.M_parent = localisation.UUID
                    
                 
          
                    server.entityManager.createEntity(localisation.UUID, "M_subLocalisations", newItem, 
                        function(success, caller){
         
                             caller.localisation.M_subLocalisations.push(success.UUID)
                            console.log(success)
                             server.entityManager.saveEntity(caller.localisation, 
                                function(result,caller){
                                    mainPage.adminPage.adminLocalisationPage.panel.getChildById(date + "-row").delete()
                                    document.getElementById("waitingDiv").style.display = "none"
                                    mainPage.adminPage.adminLocalisationPage.appendSubLocalisation(caller.newLine, result.M_id)
                                    console.log(result)
                                },function(){},{"localisation" : caller.localisation,  "newLine" : success})
                            
                        },function(){},{"localisation" : localisation})
                                    
                               
                                
                                
                                
                                
                            
                    
               }
            }(date, localisation)
        }
    }(localisation)
    
    
    
    
    mainPage.adminPage.adminLocalisationPage.panel.getChildById(localisation.M_id + "-saveBtn").element.onclick = function(localisation){
        return function(){
            
            if(mainPage.adminPage.adminLocalisationPage.modifiedItems[localisation.M_id].size > 0){
                 var i = 0
                 document.getElementById("waitingDiv").style.display = ""

                for(var object of mainPage.adminPage.adminLocalisationPage.modifiedItems[localisation.M_id].values()){
                
                    server.entityManager.saveEntity(object,
                        function(result, caller){
                           mainPage.adminPage.panel.getChildById("adminLocalisationSaveState").element.innerHTML = ""
                            caller.i++
                             document.getElementById("waitingDiv").style.display = "none"
                            if(caller.i == caller.length){
                                if(mainPage.adminPage.adminLocalisationPage.removedItems[caller.localisation.M_id].size > 0){
                                    var j = 0
                                    for(var removeObject of mainPage.adminPage.adminLocalisationPage.removedItems[caller.localisation.M_id].values()){
                                     server.entityManager.removeEntity(removeObject.UUID,
                                        function(result,caller){
                                            caller.j++
                                           
                                            if(caller.j >= caller.length){
                                                
                                                server.entityManager.getEntityByUuid(caller.localisation.UUID, true,
                                                    function(localisation,caller){
                                                        
                                                        mainPage.adminPage.adminLocalisationPage.panel.getChildById(localisation.M_id + "-selector").element.innerHTML = localisation.M_name
                                                        mainPage.adminPage.adminLocalisationPage.loadAdminControl(localisation)
                                                    mainPage.adminPage.adminLocalisationPage.panel.getChildById(localisation.M_id + "-saveBtn").element.classList.add("disabled")
                                                    mainPage.adminPage.adminLocalisationPage.panel.getChildById(localisation.M_id + "-cancelBtn").element.classList.add("disabled")
                                                    
                                                    mainPage.adminPage.adminLocalisationPage.panel.getChildById(localisation.M_id + "-control").element.classList.add("active")
                                                     mainPage.adminPage.adminLocalisationPage.panel.getChildById(localisation.M_id + "-control").element.classList.add("show")
                                                    },function(){},{})
                                                
                                            }
                                         
                                        },function(){},{"length" : mainPage.adminPage.adminLocalisationPage.removedItems[localisation.M_id].size, "j" : j, "localisation" : caller.localisation})
                       
                                     }
                                }else{
                                   server.entityManager.getEntityByUuid(caller.localisation.UUID, true,
                                    function(localisation,caller){
                                  
                                        mainPage.adminPage.adminLocalisationPage.panel.getChildById(localisation.M_id + "-selector").element.innerHTML = localisation.M_name
                                        mainPage.adminPage.adminLocalisationPage.loadAdminControl(localisation)
                                    mainPage.adminPage.adminLocalisationPage.panel.getChildById(localisation.M_id + "-saveBtn").element.classList.add("disabled")
                                    mainPage.adminPage.adminLocalisationPage.panel.getChildById(localisation.M_id + "-cancelBtn").element.classList.add("disabled")
                                    
                                    mainPage.adminPage.adminLocalisationPage.panel.getChildById(localisation.M_id + "-control").element.classList.add("active")
                                     mainPage.adminPage.adminLocalisationPage.panel.getChildById(localisation.M_id + "-control").element.classList.add("show")
                                    },function(){},{})
                                }
                               
                                
                            }
                        }, function () {}, {"i" : i, "length" : mainPage.adminPage.adminLocalisationPage.modifiedItems[localisation.M_id].size, "localisation" : localisation})
                }
            }else{
                if(mainPage.adminPage.adminLocalisationPage.removedItems[localisation.M_id].size > 0){
                     document.getElementById("waitingDiv").style.display = ""
                    var j = 0
                    for(var removeObject of mainPage.adminPage.adminLocalisationPage.removedItems[localisation.M_id].values()){
                     server.entityManager.removeEntity(removeObject.UUID,
                        function(result,caller){
                            caller.j++
                            mainPage.adminPage.panel.getChildById("adminLocalisationSaveState").element.innerHTML = ""
                            document.getElementById("waitingDiv").style.display = "none"
                            if(caller.j >= caller.length){
                                server.entityManager.getEntityByUuid(caller.localisation.UUID, true,
                                function(localisation,caller){
                                    
                                    mainPage.adminPage.adminLocalisationPage.panel.getChildById(localisation.M_id + "-selector").element.innerHTML = localisation.M_name
                                    mainPage.adminPage.adminLocalisationPage.loadAdminControl(localisation)
                                mainPage.adminPage.adminLocalisationPage.panel.getChildById(localisation.M_id + "-saveBtn").element.classList.add("disabled")
                                mainPage.adminPage.adminLocalisationPage.panel.getChildById(localisation.M_id + "-cancelBtn").element.classList.add("disabled")
                                
                                mainPage.adminPage.adminLocalisationPage.panel.getChildById(localisation.M_id + "-control").element.classList.add("active")
                                 mainPage.adminPage.adminLocalisationPage.panel.getChildById(localisation.M_id + "-control").element.classList.add("show")
                                },function(){},{})
                            }
                         
                        },function(){},{"length" : mainPage.adminPage.adminLocalisationPage.removedItems[localisation.M_id].size, "j" : j, "localisation" : localisation})
                    }
                }
                
            }
            
            
            
           
        
        }
    }(localisation)
    
    mainPage.adminPage.adminLocalisationPage.panel.getChildById(localisation.M_id + "-cancelBtn").element.onclick = function(localisation){
        return function(){
           mainPage.adminPage.panel.getChildById("adminLocalisationSaveState").element.innerHTML = ""
            mainPage.adminPage.adminLocalisationPage.panel.getChildById(localisation.M_id + "-saveBtn").element.classList.add("disabled")
            mainPage.adminPage.adminLocalisationPage.panel.getChildById(localisation.M_id + "-cancelBtn").element.classList.add("disabled")
            mainPage.adminPage.adminLocalisationPage.loadAdminControl(localisation)
            mainPage.adminPage.adminLocalisationPage.panel.getChildById(localisation.M_id + "-control").element.classList.add("active")
             mainPage.adminPage.adminLocalisationPage.panel.getChildById(localisation.M_id + "-control").element.classList.add("show")
        }
    }(localisation)
    
    mainPage.adminPage.adminLocalisationPage.panel.getChildById(localisation.M_id + "-name").element.onkeyup = function(localisation){
       return function(){
           mainPage.adminPage.panel.getChildById("adminLocalisationSaveState").element.innerHTML = "*"
            mainPage.adminPage.adminLocalisationPage.panel.getChildById(localisation.M_id + "-saveBtn").element.classList.remove("disabled")
            mainPage.adminPage.panel.getChildById(localisation.M_id + "-cancelBtn").element.classList.remove("disabled")
           var newName = mainPage.adminPage.adminLocalisationPage.panel.getChildById(localisation.M_id + "-name").element.value
           var newLocalisation = jQuery.extend({}, localisation)
           newLocalisation.M_name = newName
           
           mainPage.adminPage.adminLocalisationPage.modifiedItems[localisation.M_id].set(localisation.M_id, newLocalisation)
       }
    }(localisation)
    
    
    mainPage.adminPage.adminLocalisationPage.panel.getChildById(localisation.M_id + "-parent").element.oncomplete = function(localisation){
        return function(uuid){
            mainPage.adminPage.adminLocalisationPage.panel[localisation.M_id + "-parent"] = uuid
            mainPage.adminPage.panel.getChildById("adminLocalisationSaveState").element.innerHTML = "*"
            mainPage.adminPage.adminLocalisationPage.panel.getChildById(localisation.M_id + "-saveBtn").element.classList.remove("disabled")
            mainPage.adminPage.panel.getChildById(localisation.M_id + "-cancelBtn").element.classList.remove("disabled")
           var newParent = mainPage.adminPage.adminLocalisationPage.panel[localisation.M_id + "-parent"]
           if(newParent != ""){
               server.entityManager.getEntityByUuid(newParent,false,
                function(newParent,caller){
                    if(newParent != undefined){
                        
                         var newLocalisation = caller.localisation
                         
                         if(!newParent.M_subLocalisations.includes(newLocalisation.UUID)){
                            newParent.M_subLocalisations.push(newLocalisation.UUID) 
                            mainPage.adminPage.adminLocalisationPage.modifiedItems[localisation.M_id].set(newParent.M_id, newParent)
                            
                         }
                         if(newLocalisation.M_parent != null){
                                server.entityManager.getEntityByUuid(newLocalisation.M_parent,false,
                                    function(result,caller){
                                        result.M_subLocalisations.pop(caller.localisation.UUID)
                                        mainPage.adminPage.adminLocalisationPage.modifiedItems[caller.localisation.M_id].set(result.M_id, result)
                                    },function(){},{"localisation" : newLocalisation})
                            }
                        newLocalisation.M_parent = newParent.UUID
                         mainPage.adminPage.adminLocalisationPage.modifiedItems[caller.localisation.M_id].set(newLocalisation.M_id, newLocalisation)
                 
                    }
                   
                },function(){},{"localisation": localisation})
           }
        }
    }(localisation)

}

AdminLocalisationPage.prototype.appendSubLocalisation = function(subLocalisation, localisationID){
    mainPage.adminPage.adminLocalisationPage.panel.getChildById(localisationID+"-itemsList").appendElement({"tag" : "tr", "id" : subLocalisation.M_id + "-adminSubLocalisationRow"}).down()
    .appendElement({"tag" : "td", "data-th" : "ID", "innerHtml" : subLocalisation.M_id, "id" : subLocalisation.M_id + "-id"})
    .appendElement({"tag" : "td", "innerHtml" : subLocalisation.M_name, "data-th" : "Nom"})   
    .appendElement({"tag":"td", "data-th" : ""}).down()
    .appendElement({"tag" : "button", "class" : "btn btn-danger btn-sm", "id" : subLocalisation.M_id + "-deleteRowAdminBtn"}).down()
    .appendElement({"tag" : "i", "class" : "fa fa-trash-o"}).up()
    
     mainPage.adminPage.adminLocalisationPage.panel.getChildById(subLocalisation.M_id + "-deleteRowAdminBtn").element.onclick = function(subLocalisation, localisationID){
        return function(){
            mainPage.adminPage.panel.getChildById("adminLocalisationSaveState").element.innerHTML = "*"
            mainPage.adminPage.adminLocalisationPage.removedItems[localisationID].set(subLocalisation.M_id, subLocalisation)
            var row = mainPage.adminPage.adminLocalisationPage.panel.getChildById(subLocalisation.M_id + "-adminSubLocalisationRow")
            row.delete()
            mainPage.adminPage.adminLocalisationPage.panel.getChildById(localisationID + "-saveBtn").element.classList.remove("disabled")
            mainPage.adminPage.panel.getChildById(localisationID + "-cancelBtn").element.classList.remove("disabled")
            if(subLocalisation.M_subLocalisations.length > 0){
                mainPage.adminPage.adminLocalisationPage.removeSubLocalisations(subLocalisation.M_subLocalisations, localisationID)
            }
        }
    }(subLocalisation, localisationID)
}

AdminLocalisationPage.prototype.removeSubLocalisations = function(subLocalisations, localisationID){
    for(var i = 0; i< subLocalisations.length; i++){
        server.entityManager.getEntityByUuid(subLocalisations[i],false,
            function(subLocalisation, caller){
                mainPage.adminPage.adminLocalisationPage.removedItems[caller.localisationID].set(subLocalisation.M_id, subLocalisation)
                if(subLocalisation.M_subLocalisations.length > 0){
                    mainPage.adminPage.adminLocalisationPage.removeSubLocalisations(subLocalisation.M_subLocalisations, localisationID)
                }
            },function(){},{"localisationID" : localisationID})
        
    }
}

AdminLocalisationPage.prototype.deleteSubLocalisations = function(subLocalisations){
    for(var i = 0; i< subLocalisations.length; i++){
        server.entityManager.getEntityByUuid(subLocalisations[i],false,
            function(subLocalisation, caller){
                
                if(subLocalisation.M_subLocalisations.length > 0){
                    mainPage.adminPage.adminLocalisationPage.deleteSubLocalisations(subLocalisation.M_subLocalisations)
                }
                server.entityManager.removeEntity(subLocalisation.UUID,
                    function(success,caller){
                        console.log(success)
                    },function(){},{})
            },function(){},{})
        
    }
}