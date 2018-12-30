var AdminManufacturerPage = function(panel){
    /** The tab content panel **/
    this.panel = panel
    
    /** The panel sections. **/
    this.panel.appendElement({"tag" : "div", "class" : "row", "style" : "height:85vh;margin:0;"}).down()
        .appendElement({"tag" : "div", "class" : "col-md-3 bg-dark", "id" : "manufacturersAdminNavigation"}).down()
        .appendElement({"tag" : "div", "class" : "input-group mb-3", "style":"padding:15px;"}).down()
        .appendElement({"tag" : "span", "class" : "input-group-prepend"}).down()
        .appendElement({"tag":"span", "class":"input-group-text"}).down()
        .appendElement({"tag" : "i", "class" : "fa fa-search text-dark"}).up().up()
        .appendElement({"tag" : "input", "type" : "text", "class" : "form-control", "placeholder" : "Filtrer", "id" : "manufacturerFilterKeywords"}).up()
        .appendElement({"tag" : "div", "class" : "list-group","role" :"tablist", "id" : "manufacturersFilteredList"})
        .appendElement({"tag" : "div", "style" : "position:absolute;bottom:0; right:0;margin:30px;"}).down()
        .appendElement({"tag" : "button", "class" : "btn btn-success", "id" : "addManufacturerButton"}).down()
        .appendElement({"tag" : "i", "class" : "fa fa-plus"}).up().up().up()
        .appendElement({"tag" : "div", "class" : "col-md-9 bg-light","style":"overflow-y:auto;"}).down()
         .appendElement({"tag" : "div", "class" : "tab-content", "id" : "manufacturersAdminControl"}).up().up().up()
        
  
    this.removedItems = {}
    
    this.panel.getChildById("manufacturerFilterKeywords").element.addEventListener('change', function(){
        results = mainPage.adminPage.adminManufacturerPage.getManufacturersFromKeyword(this.value)
    })
    
    
    this.panel.getChildById("addManufacturerButton").element.onclick = function(){
      
        mainPage.adminPage.adminManufacturerPage.panel.getChildById("manufacturersFilteredList").appendElement({"tag":"a", "class" : "list-group-item list-group-item-action active show","id" : "newManufacturer-selector","data-toggle" : "tab","href" : "#"+ "newManufacturer-control","role":"tab", "innerHtml": "Nouveau manufacturier","aria-controls": "newManufacturer-control"})
         mainPage.adminPage.adminManufacturerPage.panel.getChildById("manufacturersAdminControl").appendElement({"tag" : "div", "class" : "tab-pane active show", "id" : "newManufacturer-control", "role" : "tabpanel", "aria-labelledby" : "newManufacturer-selector", "style" : "padding:15px;"}).down()
        .appendElement({"tag" : "div", "class" : "input-group mb-3"}).down()
        .appendElement({"tag" : "div", "class" : "input-group-prepend"}).down()
        .appendElement({"tag" : "span", "class" : "input-group-text", "innerHtml" : "ID"}).up()
        .appendElement({"tag" : "input", "class" : "form-control",  "type" : "text", "id" : "newManufacturer-manufacturerID"}).up()
        .appendElement({"tag" : "div", "class" : "input-group mb-3"}).down()
        .appendElement({"tag" : "div", "class" : "input-group-prepend"}).down()
        .appendElement({"tag" : "span", "class" : "input-group-text", "innerHtml" : "Nom"}).up()
        .appendElement({"tag" : "input", "class" : "form-control", "type" : "text", "id" : "newManufacturer-name"}).up()
 
        .appendElement({"tag" : "button", "class" : "btn btn-primary mr-3", "innerHtml" : "Enregistrer", "id" : "newManufacturer-saveBtn"})
        
        mainPage.adminPage.adminManufacturerPage.panel.getChildById("newManufacturer-saveBtn").element.onclick = function(){
            document.getElementById("waitingDiv").style.display = ""
            var id = mainPage.adminPage.adminManufacturerPage.panel.getChildById("newManufacturer-manufacturerID").element.value
            var name = mainPage.adminPage.adminManufacturerPage.panel.getChildById("newManufacturer-name").element.value
            var newManufacturer = new CatalogSchema.ManufacturerType
            newManufacturer.M_id = id
            newManufacturer.M_name = name

           
          
            server.entityManager.createEntity(catalog.UUID, "M_manufacturers", newManufacturer,
                function(success,caller){
                    
          
                            document.getElementById("waitingDiv").style.display = "none"
                            mainPage.adminPage.adminManufacturerPage.panel.getChildById("manufacturersFilteredList").removeAllChilds()
                            mainPage.adminPage.adminManufacturerPage.panel.getChildById("manufacturersAdminControl").removeAllChilds()
                            mainPage.adminPage.adminManufacturerPage.panel.getChildById("manufacturersFilteredList").appendElement({"tag":"a", "class" : "list-group-item list-group-item-action","id" : caller.newManufacturer.M_id + "-selector","data-toggle" : "tab","href" : "#"+ caller.newManufacturer.M_id + "-control","role":"tab", "innerHtml": caller.newManufacturer.M_name,"aria-controls":  caller.newManufacturer.M_id + "-control"})
                            mainPage.adminPage.adminManufacturerPage.loadAdminControl(caller.newManufacturer)
                            console.log(success)
                        
                },function(){},{"newManufacturer" : newManufacturer})
            
           
            
       
            
            
        }
    }
    
    return this
}

AdminManufacturerPage.prototype.getManufacturersFromKeyword = function(keyword){

    
    var q = new EntityQuery()
    q.TypeName = "CatalogSchema.ManufacturerType"
    q.Fields = ["M_name","M_id"]
    q.Query = 'CatalogSchema.ManufacturerType.M_name=="'+ keyword +'" '
    
    server.entityManager.getEntities("CatalogSchema.ManufacturerType", "CatalogSchema", q, 0,-1,[],true,false, 
        function(index,total, caller){
            
        },
        function(manufacturers,caller){
     
            mainPage.adminPage.adminManufacturerPage.panel.getChildById("manufacturersFilteredList").removeAllChilds()
            mainPage.adminPage.adminManufacturerPage.panel.getChildById("manufacturersAdminControl").removeAllChilds()
           
            
           
            for(var i = 0; i < manufacturers.length; i++){
                mainPage.adminPage.adminManufacturerPage.panel.getChildById("manufacturersFilteredList").appendElement({"tag":"a", "class" : "list-group-item list-group-item-action","id" : manufacturers[i].M_id + "-selector","data-toggle" : "tab","href" : "#"+ manufacturers[i].M_id + "-control","role":"tab", "innerHtml": manufacturers[i].M_name,"aria-controls":  manufacturers[i].M_id + "-control"})
                mainPage.adminPage.adminManufacturerPage.loadAdminControl(manufacturers[i])
               console.log(manufacturers[i])
                
            }
            
        },
        function(){
        },{})
}

AdminManufacturerPage.prototype.loadAdminControl = function(manufacturer){
    mainPage.adminPage.panel.getChildById("adminManufacturerSaveState").element.innerHTML = ""
    mainPage.adminPage.adminManufacturerPage.removedItems[manufacturer.M_id] = {}
    mainPage.adminPage.adminManufacturerPage.removedItems[manufacturer.M_id]["items"] = new Map()
    mainPage.adminPage.adminManufacturerPage.removedItems[manufacturer.M_id]["packages"] = new Map()
    if(mainPage.adminPage.adminManufacturerPage.panel.getChildById(manufacturer.M_id + "-control") != undefined){
        mainPage.adminPage.adminManufacturerPage.panel.getChildById(manufacturer.M_id + "-control").delete()
    }
    mainPage.adminPage.adminManufacturerPage.panel.getChildById("manufacturersAdminControl").appendElement({"tag" : "div", "class" : "tab-pane", "id" : manufacturer.M_id + "-control", "role" : "tabpanel", "aria-labelledby" : manufacturer.M_id + "-selector", "style" : "padding:15px;"}).down()
    .appendElement({"tag" : "div", "class" : "input-group mb-3"}).down()
    .appendElement({"tag" : "div", "class" : "input-group-prepend"}).down()
    .appendElement({"tag" : "span", "class" : "input-group-text", "innerHtml" : "ID"}).up()
    .appendElement({"tag" : "span", "class" : "form-control",  "innerHtml" : manufacturer.M_id}).up()
    .appendElement({"tag" : "div", "class" : "input-group mb-3"}).down()
    .appendElement({"tag" : "div", "class" : "input-group-prepend"}).down()
    .appendElement({"tag" : "span", "class" : "input-group-text", "innerHtml" : "Nom"}).up()
    .appendElement({"tag" : "input", "class" : "form-control", "type" : "text", "value" : manufacturer.M_name, "id" : manufacturer.M_id + "-name"}).up()
    .appendElement({"tag" : "div", "class" : "row"}).down()
    .appendElement({"tag":"label", "for" : manufacturer.M_id + "-packagesList", "innerHtml" : "Paquets","class" : "col-3 d-flex align-items-center"})
    .appendElement({"tag": "input", "class" : "form-control col-5 m-1", "id" : manufacturer.M_id + "-filterPackages","type" :"text", "placeholder" : "Filtrer.."}).up()
    .appendElement({"tag" : "div"}).down()
    .appendElement({"tag":"table", "class":"table table-hover table-condensed"}).down()
    .appendElement({"tag":"thead"}).down()
    .appendElement({"tag" : "tr"}).down()
    .appendElement({"tag" : "th", "innerHtml" : "ID Paquet", "style" : "width:30%;"})
    .appendElement({"tag" : "th", "innerHtml" : "Nom", "style":"width:50%;"})

    .appendElement({"tag" : "th", "style" : "width:20%;"}).down()
    .appendElement({"tag" : "button", "class" : "btn btn-success btn-sm", "id" : manufacturer.M_id +"-addPackageButton", "style":"display: inline-flex; height: 29px;"}).down()
    .appendElement({"tag" : "i", "class" : "fa fa-plus"}).up().up()
    .up().up()
    .appendElement({"tag" : "tbody", "id" : manufacturer.M_id + "-packagesList"})
    .up().up()
    .appendElement({"tag" : "div", "class" : "row"}).down()
    .appendElement({"tag":"label", "for" : manufacturer.M_id + "-itemsList", "innerHtml" : "Items","class" : "col-3 d-flex align-items-center"})
    .appendElement({"tag": "input", "class" : "form-control col-5 m-1", "id" : manufacturer.M_id + "-filterItems","type" :"text", "placeholder" : "Filtrer.."}).up()
    .appendElement({"tag" : "div"}).down()
    .appendElement({"tag":"table", "class":"table table-hover table-condensed"}).down()
    .appendElement({"tag":"thead"}).down()
    .appendElement({"tag" : "tr"}).down()
    .appendElement({"tag" : "th", "innerHtml" : "ID Item", "style" : "width:30%;"})
    .appendElement({"tag" : "th", "innerHtml" : "Nom", "style":"width:50%;"})

    .appendElement({"tag" : "th", "style" : "width:20%;"}).down()
    .appendElement({"tag" : "button", "class" : "btn btn-success btn-sm", "id" : manufacturer.M_id +"-addItemButton", "style":"display: inline-flex; height: 29px;"}).down()
    .appendElement({"tag" : "i", "class" : "fa fa-plus"}).up().up()
    .up().up()
    .appendElement({"tag" : "tbody", "id" : manufacturer.M_id + "-itemsList"})
    .up().up()
    .appendElement({"tag" : "div", "class" : "d-flex justify-content-center"}).down()
    .appendElement({"tag" : "button", "class" : "btn btn-primary disabled mr-3", "innerHtml" : "Enregistrer", "id" : manufacturer.M_id + "-saveBtn"})
    .appendElement({"tag" : "button", "class" : "btn disabled mr-3", "innerHtml" : "Annuler les modifications", "id" : manufacturer.M_id + "-cancelBtn"})
    .appendElement({"tag" : "button", "class" : "btn btn-danger", "innerHtml" : "Supprimer l'entité", "id" : manufacturer.M_id + "-deleteBtn", "data-toggle" : "modal", "data-target" : "#" + manufacturer.M_id + "-modal"})
    .appendElement({"tag" : "div", "class" : "modal fade", "id"  : manufacturer.M_id + "-modal", "tabindex" : "-1", "role"  :"dialog", "aria-labelledby" : manufacturer.M_id + "-modal", "aria-hidden" : "true"}).down()
    .appendElement({"tag" : "div", "class" : "modal-dialog", "role" : "document"}).down()
    .appendElement({"tag" : "div", "class" : "modal-content"}).down()
    .appendElement({"tag" : "div", "class" : "modal-header"}).down()
    .appendElement({"tag" : "h5", "class" : "modal-title", "innerHtml" : "Supprimer le manufacturier"})
    .appendElement({"tag" : "button", "class" : "close", "data-dismiss" : "modal", "aria-label" : "Close"}).down()
    .appendElement({"tag" : "span", "aria-hidden" : "true", "innerHtml" : "&times;"}).up().up()
    .appendElement({"tag" : "div", "class" : "modal-body", "innerHtml" : "Êtes-vous certain de vouloir supprimer ce manufacturier?"})
    .appendElement({"tag" : "div", "class" : "modal-footer"}).down()
    .appendElement({"tag" : "button", "class" : "btn btn-secondary", "data-dismiss" : "modal", "innerHtml" : "Fermer", "id" : manufacturer.M_id + "-closeDelete"})
    .appendElement({"tag" : "button", "type" : "button", "class" : "btn btn-danger", "id" : manufacturer.M_id + "-deleteConfirmBtn", "innerHtml" : "Confirmer la suppression"})
    
    mainPage.adminPage.adminManufacturerPage.panel.getChildById(manufacturer.M_id + "-deleteConfirmBtn").element.onclick = function(manufacturer){
        return function(){
            document.getElementById("waitingDiv").style.display = ""
            mainPage.adminPage.adminManufacturerPage.panel.getChildById(manufacturer.M_id + "-closeDelete").element.click()
      
                
            server.entityManager.removeEntity(manufacturer.UUID,
            function(success,caller){
                document.getElementById("waitingDiv").style.display = "none"
                 mainPage.adminPage.adminManufacturerPage.panel.getChildById(caller.manufacturer.M_id + "-control").delete()
                 mainPage.adminPage.adminManufacturerPage.panel.getChildById(caller.manufacturer.M_id + "-selector").delete()
                
            }, function(){},{"manufacturer" : manufacturer})
        }
    }(manufacturer)
    
    
    for(var i  = 0; i < manufacturer.M_packages.length; i++){
        mainPage.adminPage.adminManufacturerPage.appendPackage(manufacturer.M_packages[i], manufacturer)
    }
    
    for(var i  = 0; i < manufacturer.M_items.length; i++){
        mainPage.adminPage.adminManufacturerPage.appendItem(manufacturer.M_items[i], manufacturer)
    }
    
    
    mainPage.adminPage.adminManufacturerPage.panel.getChildById(manufacturer.M_id + "-addPackageButton").element.onclick = function(manufacturer){
        return function(){
            var date = new Date()
            mainPage.adminPage.adminManufacturerPage.panel.getChildById(manufacturer.M_id + "-packagesList").prependElement({"tag" : "tr", "id" : date + "-row"}).down()
            .appendElement({"tag" : "th", "scope" : "row"}).down()
            .appendElement({"tag" : "input", "class" : "form-control", "type" : "text", "placeholder" : "ID", "id" : date  + "-packageID"}).up()
            .appendElement({"tag" : "td", "scope" : "row"})
             .appendElement({"tag":"td"}).down()
            .appendElement({"tag" : "button", "class" : "btn btn-success btn-sm", "id" : date + "-confirmBtn"}).down()
            .appendElement({"tag" : "i", "class" : "fa fa-check"}).up()
            
            autocomplete("CatalogSchema.PackageType", mainPage.adminPage.adminManufacturerPage.panel.getChildById(date+"-packageID"))
            
             mainPage.adminPage.adminManufacturerPage.panel.getChildById(date + "-packageID").element.oncomplete = function(uuid){
                mainPage.adminPage.adminManufacturerPage.panel[date+"-packageID"] = uuid
            }
            
             mainPage.adminPage.adminManufacturerPage.panel.getChildById(date+ "-confirmBtn").element.onclick = function (date, manufacturer) {
            return function(){
                document.getElementById("waitingDiv").style.display = ""
                        var newPkg = mainPage.adminPage.adminManufacturerPage.panel[date+"-packageID"]
                        var newManufacturer = jQuery.extend({},manufacturer)
                        newManufacturer.M_packages.push(newPkg)
                        server.entityManager.saveEntity(newManufacturer,
                            function(success,caller){
                             
                                mainPage.adminPage.adminManufacturerPage.panel.getChildById(date + "-row").delete()
                                    document.getElementById("waitingDiv").style.display = "none"
                                    server.entityManager.getEntityByUuid(caller.package,false,
                                        function(pkg,caller){
                                            mainPage.adminPage.adminManufacturerPage.appendPackage(pkg, caller.success)
                                        },function(){},{"success" : success})
                                    
                            },function(){},{"package" : newPkg})
             
               }
            }(date, manufacturer)
             
                
        }
    }(manufacturer)
    
     mainPage.adminPage.adminManufacturerPage.panel.getChildById(manufacturer.M_id + "-addItemButton").element.onclick = function(manufacturer){
        return function(){
            var date = new Date()
            mainPage.adminPage.adminManufacturerPage.panel.getChildById(manufacturer.M_id + "-itemsList").prependElement({"tag" : "tr", "id" : date + "-row"}).down()
            .appendElement({"tag" : "th", "scope" : "row"}).down()
            .appendElement({"tag" : "input", "class" : "form-control", "type" : "text", "placeholder" : "ID", "id" : date  + "-itemID"}).up()
            .appendElement({"tag" : "td", "scope" : "row"})
             .appendElement({"tag":"td"}).down()
            .appendElement({"tag" : "button", "class" : "btn btn-success btn-sm", "id" : date + "-confirmBtn"}).down()
            .appendElement({"tag" : "i", "class" : "fa fa-check"}).up()
            
            autocomplete("CatalogSchema.ItemType", mainPage.adminPage.adminManufacturerPage.panel.getChildById(date+"-itemID"))
            
             mainPage.adminPage.adminManufacturerPage.panel.getChildById(date + "-itemID").element.oncomplete = function(uuid){
                mainPage.adminPage.adminManufacturerPage.panel[date+"-itemID"] = uuid
            }
            
             mainPage.adminPage.adminManufacturerPage.panel.getChildById(date+ "-confirmBtn").element.onclick = function (date, manufacturer) {
            return function(){
                document.getElementById("waitingDiv").style.display = ""
                        var newItem = mainPage.adminPage.adminManufacturerPage.panel[date+"-itemID"]
                        var newManufacturer = jQuery.extend({}, manufacturer)
                        newManufacturer.M_items.push(newItem)
                        server.entityManager.saveEntity(newManufacturer,
                            function(success,caller){
                             
                                mainPage.adminPage.adminManufacturerPage.panel.getChildById(date + "-row").delete()
                                    document.getElementById("waitingDiv").style.display = "none"
                                    server.entityManager.getEntityByUuid(caller.item,false,
                                        function(newItem,caller){
                                             mainPage.adminPage.adminManufacturerPage.appendItem(newItem, caller.success)
                                        },function(){}, {"success" : success})
                                   
                            },function(){},{"item" : newItem})
                
               }
            }(date, manufacturer)
             
                
        }
    }(manufacturer)
    
    mainPage.adminPage.adminManufacturerPage.panel.getChildById(manufacturer.M_id + "-name").element.onkeyup = function(manufacturer){
       return function(){
            mainPage.adminPage.adminManufacturerPage.panel.getChildById(manufacturer.M_id + "-saveBtn").element.classList.remove("disabled")
            mainPage.adminPage.adminManufacturerPage.panel.getChildById(manufacturer.M_id + "-cancelBtn").element.classList.remove("disabled")
            mainPage.adminPage.panel.getChildById("adminManufacturerSaveState").element.innerHTML = "*"
       }
    }(manufacturer)
    
    mainPage.adminPage.adminManufacturerPage.panel.getChildById(manufacturer.M_id + "-saveBtn").element.onclick = function(manufacturer){
        return function(){
            
            document.getElementById("waitingDiv").style.display = ""
            var newManufacturer = jQuery.extend({}, manufacturer)
            newManufacturer.M_name = mainPage.adminPage.adminManufacturerPage.panel.getChildById(manufacturer.M_id + "-name").element.value
        
            if(mainPage.adminPage.adminManufacturerPage.removedItems[manufacturer.M_id]["packages"].size > 0){
                    var i = 0;
                for(var object of mainPage.adminPage.adminManufacturerPage.removedItems[manufacturer.M_id]["packages"].values()){
                    i++
                    newManufacturer.M_packages.pop(object)
                    if(i== mainPage.adminPage.adminManufacturerPage.removedItems[manufacturer.M_id]["packages"].size){
                        if(mainPage.adminPage.adminManufacturerPage.removedItems[manufacturer.M_id]["items"].size > 0){
                            var j = 0;
                            for(var item of mainPage.adminPage.adminManufacturerPage.removedItems[manufacturer.M_id]["items"].values()){
                                j++
                                newManufacturer.M_items.pop(item)
                                if(j == mainPage.adminPage.adminManufacturerPage.removedItems[manufacturer.M_id]["items"].size){
                                    server.entityManager.saveEntity(newManufacturer,
                                        function(success,caller){
                                            document.getElementById("waitingDiv").style.display = "none"
                                            mainPage.adminPage.panel.getChildById("adminManufacturerSaveState").element.innerHTML = ""
                                             mainPage.adminPage.adminManufacturerPage.panel.getChildById(newManufacturer.M_id + "-selector").element.innerHTML = newManufacturer.M_name
                                                mainPage.adminPage.adminManufacturerPage.loadAdminControl(newManufacturer)
                                            mainPage.adminPage.adminManufacturerPage.panel.getChildById(newManufacturer.M_id + "-saveBtn").element.classList.add("disabled")
                                            mainPage.adminPage.adminManufacturerPage.panel.getChildById(newManufacturer.M_id + "-cancelBtn").element.classList.add("disabled")
                                            
                                            mainPage.adminPage.adminManufacturerPage.panel.getChildById(newManufacturer.M_id + "-control").element.classList.add("active")
                                             mainPage.adminPage.adminManufacturerPage.panel.getChildById(newManufacturer.M_id + "-control").element.classList.add("show")
                                        },function(){},{})
                                }
                            }
                        }else{
                            server.entityManager.saveEntity(newManufacturer,
                                function(success,caller){
                                    document.getElementById("waitingDiv").style.display = "none"
                                     mainPage.adminPage.panel.getChildById("adminManufacturerSaveState").element.innerHTML = ""
                                     mainPage.adminPage.adminManufacturerPage.panel.getChildById(newManufacturer.M_id + "-selector").element.innerHTML = newManufacturer.M_name
                                        mainPage.adminPage.adminManufacturerPage.loadAdminControl(newManufacturer)
                                    mainPage.adminPage.adminManufacturerPage.panel.getChildById(newManufacturer.M_id + "-saveBtn").element.classList.add("disabled")
                                    mainPage.adminPage.adminManufacturerPage.panel.getChildById(newManufacturer.M_id + "-cancelBtn").element.classList.add("disabled")
                                    
                                    mainPage.adminPage.adminManufacturerPage.panel.getChildById(newManufacturer.M_id + "-control").element.classList.add("active")
                                     mainPage.adminPage.adminManufacturerPage.panel.getChildById(newManufacturer.M_id + "-control").element.classList.add("show")
                                },function(){},{})
                        }
                        
                    }
                }
            }else{
               if(mainPage.adminPage.adminManufacturerPage.removedItems[manufacturer.M_id]["items"].size > 0){
                    var j = 0;
                    for(var item of mainPage.adminPage.adminManufacturerPage.removedItems[manufacturer.M_id]["items"].values()){
                        j++
                        newManufacturer.M_items.pop(item)
                        if(j == mainPage.adminPage.adminManufacturerPage.removedItems[manufacturer.M_id]["items"].size){
                            server.entityManager.saveEntity(newManufacturer,
                                function(success,caller){
                                    document.getElementById("waitingDiv").style.display = "none"
                                     mainPage.adminPage.panel.getChildById("adminManufacturerSaveState").element.innerHTML = ""
                                     mainPage.adminPage.adminManufacturerPage.panel.getChildById(newManufacturer.M_id + "-selector").element.innerHTML = newManufacturer.M_name
                                        mainPage.adminPage.adminManufacturerPage.loadAdminControl(newManufacturer)
                                    mainPage.adminPage.adminManufacturerPage.panel.getChildById(newManufacturer.M_id + "-saveBtn").element.classList.add("disabled")
                                    mainPage.adminPage.adminManufacturerPage.panel.getChildById(newManufacturer.M_id + "-cancelBtn").element.classList.add("disabled")
                                    
                                    mainPage.adminPage.adminManufacturerPage.panel.getChildById(newManufacturer.M_id + "-control").element.classList.add("active")
                                     mainPage.adminPage.adminManufacturerPage.panel.getChildById(newManufacturer.M_id + "-control").element.classList.add("show")
                                },function(){},{})
                        }
                    }
                }else{
                    server.entityManager.saveEntity(newManufacturer,
                        function(success,caller){
                            document.getElementById("waitingDiv").style.display = "none"
                             mainPage.adminPage.panel.getChildById("adminManufacturerSaveState").element.innerHTML = ""
                             mainPage.adminPage.adminManufacturerPage.panel.getChildById(newManufacturer.M_id + "-selector").element.innerHTML = newManufacturer.M_name
                                mainPage.adminPage.adminManufacturerPage.loadAdminControl(newManufacturer)
                            mainPage.adminPage.adminManufacturerPage.panel.getChildById(newManufacturer.M_id + "-saveBtn").element.classList.add("disabled")
                            mainPage.adminPage.adminManufacturerPage.panel.getChildById(newManufacturer.M_id + "-cancelBtn").element.classList.add("disabled")
                            
                            mainPage.adminPage.adminManufacturerPage.panel.getChildById(newManufacturer.M_id + "-control").element.classList.add("active")
                             mainPage.adminPage.adminManufacturerPage.panel.getChildById(newManufacturer.M_id + "-control").element.classList.add("show")
                        },function(){},{})
                }
            }
            
            
            
            
            
            
           
        
        }
    }(manufacturer)
    
    mainPage.adminPage.adminManufacturerPage.panel.getChildById(manufacturer.M_id + "-cancelBtn").element.onclick = function(manufacturer){
        return function(){
            
            mainPage.adminPage.adminManufacturerPage.panel.getChildById(manufacturer.M_id + "-saveBtn").element.classList.add("disabled")
            mainPage.adminPage.adminManufacturerPage.panel.getChildById(manufacturer.M_id + "-cancelBtn").element.classList.add("disabled")
            mainPage.adminPage.adminManufacturerPage.loadAdminControl(manufacturer)
            mainPage.adminPage.adminManufacturerPage.panel.getChildById(manufacturer.M_id + "-control").element.classList.add("active")
             mainPage.adminPage.adminManufacturerPage.panel.getChildById(manufacturer.M_id + "-control").element.classList.add("show")
        }
    }(manufacturer)
    
    mainPage.adminPage.adminManufacturerPage.panel.getChildById(manufacturer.M_id + "-filterItems").element.onkeyup = function(manufacturerID){
        return function(){
            var input, filter, table, tr, td, i;
            input = document.getElementById(manufacturerID+"-filterItems");
            filter = input.value.toUpperCase();
            table = document.getElementById(manufacturerID + "-itemsList");
            tr = table.getElementsByTagName("tr");
        
            // Loop through all table rows, and hide those who don't match the search query
            for (i = 0; i < tr.length; i++) {
                td = tr[i].getElementsByTagName("td")[0];
                if (td) {
                    if (td.innerHTML.toUpperCase().indexOf(filter) > -1) {
                        tr[i].style.display = "";
                    } else {
                        tr[i].style.display = "none";
                    }
                } 
            }
       }
    }(manufacturer.M_id)
    
    mainPage.adminPage.adminManufacturerPage.panel.getChildById(manufacturer.M_id + "-filterPackages").element.onkeyup = function(manufacturerID){
        return function(){
            var input, filter, table, tr, td, i;
            input = document.getElementById(manufacturerID+"-filterPackages");
            filter = input.value.toUpperCase();
            table = document.getElementById(manufacturerID + "-packagesList");
            tr = table.getElementsByTagName("tr");
        
            // Loop through all table rows, and hide those who don't match the search query
            for (i = 0; i < tr.length; i++) {
                td = tr[i].getElementsByTagName("td")[0];
                if (td) {
                    if (td.innerHTML.toUpperCase().indexOf(filter) > -1) {
                        tr[i].style.display = "";
                    } else {
                        tr[i].style.display = "none";
                    }
                } 
            }
       }
    }(manufacturer.M_id)
    
    
}

AdminManufacturerPage.prototype.appendPackage = function(pkg, manufacturer){
    mainPage.adminPage.adminManufacturerPage.panel.getChildById(manufacturer.M_id +"-packagesList").appendElement({"tag" : "tr", "id" : pkg.M_id + "-adminPackageRow"}).down()
    .appendElement({"tag" :"td","innerHtml" : pkg.M_id, "id" : pkg.M_id + "-packageID", "class" : "tabLink"})
    .appendElement({"tag" : "td", "innerHtml" : pkg.M_name, "id" : pkg.M_id + "-packageName"})
    .appendElement({"tag":"td"}).down()
    .appendElement({"tag" : "button", "class" : "btn btn-danger btn-sm", "id" : pkg.M_id + "-deletePkgRowAdminBtn", "style":"display: inline-flex; height: 29px;"}).down()
    .appendElement({"tag" : "i", "class" : "fa fa-trash-o"}).up()
    
    mainPage.adminPage.adminManufacturerPage.panel.getChildById(pkg.M_id + "-deletePkgRowAdminBtn").element.onclick = function(pkg, manufacturer){
        return function(){
            mainPage.adminPage.adminManufacturerPage.removedItems[manufacturer.M_id]["packages"].set(pkg.M_id + "-pkg", pkg)
            var row = mainPage.adminPage.adminManufacturerPage.panel.getChildById(pkg.M_id + "-adminPackageRow")
            row.delete()
            mainPage.adminPage.adminManufacturerPage.panel.getChildById(manufacturer.M_id + "-saveBtn").element.classList.remove("disabled")
            mainPage.adminPage.adminManufacturerPage.panel.getChildById(manufacturer.M_id + "-cancelBtn").element.classList.remove("disabled")
            mainPage.adminPage.panel.getChildById("adminManufacturerSaveState").element.innerHTML = "*"
        }
    }(pkg, manufacturer)
    
    mainPage.adminPage.adminManufacturerPage.panel.getChildById(pkg.M_id + "-packageID").element.onclick = function(item_package){
        return function(){
            mainPage.packageDisplayPage.displayTabItem(item_package)
        }
    }(pkg)
    
}

AdminManufacturerPage.prototype.appendItem = function(item, manufacturer){
    mainPage.adminPage.adminManufacturerPage.panel.getChildById(manufacturer.M_id +"-itemsList").appendElement({"tag" : "tr", "id" : item.M_id + "-adminItemRow"}).down()
    .appendElement({"tag" :"td","innerHtml" : item.M_id, "id" : item.M_id + "-itemID", "class" : "tabLink"})
    .appendElement({"tag" : "td", "innerHtml" : item.M_name, "id" : item.M_id + "-itemName"})
    .appendElement({"tag":"td"}).down()
    .appendElement({"tag" : "button", "class" : "btn btn-danger btn-sm", "id" : item.M_id + "-deleteItemRowAdminBtn", "style":"display: inline-flex; height: 29px;"}).down()
    .appendElement({"tag" : "i", "class" : "fa fa-trash-o"}).up()
    
    mainPage.adminPage.adminManufacturerPage.panel.getChildById(item.M_id + "-deleteItemRowAdminBtn").element.onclick = function(item, manufacturer){
        return function(){
            mainPage.adminPage.adminManufacturerPage.removedItems[manufacturer.M_id]["items"].set(item.M_id + "-item", item)
            var row = mainPage.adminPage.adminManufacturerPage.panel.getChildById(item.M_id + "-adminItemRow")
            row.delete()
            mainPage.adminPage.adminManufacturerPage.panel.getChildById(manufacturer.M_id + "-saveBtn").element.classList.remove("disabled")
            mainPage.adminPage.adminManufacturerPage.panel.getChildById(manufacturer.M_id + "-cancelBtn").element.classList.remove("disabled")
            mainPage.adminPage.panel.getChildById("adminManufacturerSaveState").element.innerHTML = "*"

        }
    }(item, manufacturer)
    
    mainPage.adminPage.adminManufacturerPage.panel.getChildById(item.M_id + "-itemID").element.onclick = function(item){
        return function(){
            mainPage.itemDisplayPage.displayTabItem(item)
        }
    }(item)
    
}

