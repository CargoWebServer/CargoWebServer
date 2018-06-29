var AdminPackagePage = function(panel){
    /** The tab content panel **/
    this.panel = panel
    
    /** The panel sections. **/
    this.panel.appendElement({"tag" : "div", "class" : "row", "style" : "height:85vh;margin:0;"}).down()
        .appendElement({"tag" : "div", "class" : "col-md-3 bg-dark", "id" : "packagesAdminNavigation"}).down()
        .appendElement({"tag" : "div", "class" : "input-group mb-3", "style":"padding:15px;"}).down()
        .appendElement({"tag" : "span", "class" : "input-group-prepend"}).down()
        .appendElement({"tag":"span", "class":"input-group-text"}).down()
        .appendElement({"tag" : "i", "class" : "fa fa-search text-dark"}).up().up()
        .appendElement({"tag" : "input", "type" : "text", "class" : "form-control", "placeholder" : "Filtrer", "id" : "packageFilterKeywords"}).up()
        .appendElement({"tag" : "div", "class" : "list-group","role" :"tablist", "id" : "packagesFilteredList", "style" : "overflow-y: auto;height: 68vh;"})
        .appendElement({"tag" : "div", "style" : "position:absolute;bottom:0; right:0;margin:30px;"}).down()
        .appendElement({"tag" : "button", "class" : "btn btn-success", "id" : "addPackageButton"}).down()
        .appendElement({"tag" : "i", "class" : "fa fa-plus"}).up().up().up()
        .appendElement({"tag" : "div", "class" : "col-md-9 bg-light","style":"overflow-y:auto;"}).down()
         .appendElement({"tag" : "div", "class" : "tab-content", "id" : "packagesAdminControl"}).up().up().up()
        
  
    this.removedItems = {}
    
    this.currencies = []
    
    this.modifiedItems = {}
    
    this.panel.getChildById("packageFilterKeywords").element.addEventListener('change', function(){
        results = mainPage.adminPage.adminPackagePage.getPackagesFromKeyword(this.value)
    })
    
    this.unitOfMeasure = null
    
    
    
    
    this.panel.getChildById("addPackageButton").element.onclick = function(){
      
        mainPage.adminPage.adminPackagePage.panel.getChildById("packagesFilteredList").appendElement({"tag":"a", "class" : "list-group-item list-group-item-action show active","id" : "newPackage-selector","data-toggle" : "tab","href" : "#"+ "newPackage-control","role":"tab", "innerHtml": "Nouvelle catégorie","aria-controls": "newPackage-control"})
         mainPage.adminPage.adminPackagePage.panel.getChildById("packagesAdminControl").appendElement({"tag" : "div", "class" : "tab-pane show active", "id" : "newPackage-control", "role" : "tabpanel", "aria-labelledby" : "newPackage-selector", "style" : "padding:15px;"}).down()
        .appendElement({"tag" : "div", "class" : "input-group mb-3"}).down()
        .appendElement({"tag" : "div", "class" : "input-group-prepend"}).down()
        .appendElement({"tag" : "span", "class" : "input-group-text", "innerHtml" : "ID"}).up()
        .appendElement({"tag" : "input", "class" : "form-control",  "type" : "text", "id" : "newPackage-packageID"}).up()
        .appendElement({"tag" : "div", "class" : "input-group mb-3"}).down()
        .appendElement({"tag" : "div", "class" : "input-group-prepend"}).down()
        .appendElement({"tag" : "span", "class" : "input-group-text", "innerHtml" : "Nom"}).up()
        .appendElement({"tag" : "input", "class" : "form-control", "type" : "text", "id" : "newPackage-name"}).up()
 
        .appendElement({"tag" : "button", "class" : "btn btn-primary mr-3", "innerHtml" : "Enregistrer", "id" : "newPackage-saveBtn"})
        
        mainPage.adminPage.adminPackagePage.panel.getChildById("newPackage-saveBtn").element.onclick = function(){
            document.getElementById("waitingDiv").style.display = ""
            var id = mainPage.adminPage.adminPackagePage.panel.getChildById("newPackage-packageID").element.value
            var name = mainPage.adminPage.adminPackagePage.panel.getChildById("newPackage-name").element.value
            var newPackage = new CatalogSchema.PackageType
            newPackage.M_id = id
            newPackage.M_name = name

           
          
            server.entityManager.createEntity(catalog.UUID, "M_packages", newPackage,
                function(success,caller){
                            document.getElementById("waitingDiv").style.display = "none"
                            mainPage.adminPage.adminPackagePage.panel.getChildById("packagesFilteredList").removeAllChilds()
                            mainPage.adminPage.adminPackagePage.panel.getChildById("packagesAdminControl").removeAllChilds()
                            mainPage.adminPage.adminPackagePage.panel.getChildById("packagesFilteredList").appendElement({"tag":"a", "class" : "list-group-item list-group-item-action","id" : caller.newPackage.M_id + "-selector","data-toggle" : "tab","href" : "#"+ caller.newPackage.M_id + "-control","role":"tab", "innerHtml": caller.newPackage.M_name,"aria-controls":  caller.newPackage.M_id + "-control"})
                            mainPage.adminPage.adminPackagePage.loadAdminControl(caller.newPackage)
                        
                },function(){},{"newPackage" : newPackage})
           
        }
    }
    
    return this
}

AdminPackagePage.prototype.getPackagesFromKeyword = function(keyword){
    
    server.entityManager.getEntityPrototype("CatalogSchema.CurrencyType", "CatalogSchema", 
        function(success,caller){
            mainPage.adminPage.adminPackagePage.currencies = []
            for(var i = 0; i < success.Restrictions.length; i++){
                
                mainPage.adminPage.adminPackagePage.currencies.push(success.Restrictions[i].Value)
            }

        }, function(){},{})
    xapian.search(
        dbpaths,
        keyword.toUpperCase(),
        [],
        "en",
        0,
        1000,
        // success callback
        function (results, caller) {
            mainPage.adminPage.adminPackagePage.panel.getChildById("packagesFilteredList").removeAllChilds()
            mainPage.adminPage.adminPackagePage.panel.getChildById("packagesAdminControl").removeAllChilds()
            var uuids = []
            for (var i = 0; i < results.results.length; i++) {
                    var result = results.results[i];
                    if (result.data.TYPENAME == "CatalogSchema.PackageType") {
                        if(uuids.indexOf(result.data.UUID) == -1){
                            uuids.push(result.data.UUID)
                            server.entityManager.getEntityByUuid(result.data.UUID,false,
                                function(pkg, caller){
                                    mainPage.adminPage.adminPackagePage.panel.getChildById("packagesFilteredList").appendElement({"tag":"a", "class" : "list-group-item list-group-item-action","id" : pkg.M_id + "-selector","data-toggle" : "tab","href" : "#"+ pkg.M_id + "-control","role":"tab", "innerHtml": pkg.M_name,"aria-controls":  pkg.M_id + "-control"})
                                    
                                    mainPage.adminPage.adminPackagePage.panel.getChildById(pkg.M_id + "-selector").element.onclick = function(pkg){
                                        return function(){
                                            mainPage.adminPage.adminPackagePage.loadAdminControl(pkg)
                                        }
                                    }(pkg)
                                },function(){},{})
                        }
                    }
            }
       
        },
        // error callback
        function () {

        }, {"keyword":keyword})
}

AdminPackagePage.prototype.loadAdminControl = function(pkg){
    
    if(mainPage.adminPage.adminPackagePage.unitOfMeasure == null){
        server.entityManager.getEntityPrototype("CatalogSchema.UnitOfMeasureEnum", "CatalogSchema",
        function(success, caller){
            mainPage.adminPage.adminPackagePage.unitOfMeasure = success.Restrictions
        })
    }
    mainPage.adminPage.adminPackagePage.modifiedItems[pkg.M_id] = new Map()
    mainPage.adminPage.adminPackagePage.removedItems[pkg.M_id] = new Map()
    if(mainPage.adminPage.adminPackagePage.panel.getChildById(pkg.M_id + "-control") != undefined){
        mainPage.adminPage.adminPackagePage.panel.getChildById(pkg.M_id + "-control").delete()
    }
    mainPage.adminPage.adminPackagePage.panel.getChildById("packagesAdminControl").appendElement({"tag" : "div", "class" : "tab-pane", "id" : pkg.M_id + "-control", "role" : "tabpanel", "aria-labelledby" : pkg.M_id + "-selector", "style" : "padding:15px;"}).down()
    .appendElement({"tag" : "div", "class" : "row mb-3"}).down()
    .appendElement({"tag":  "div" , "class" : "col-7"}).down()
    .appendElement({"tag" : "div", "class" : "input-group mb-3"}).down()
    .appendElement({"tag" : "div", "class" : "input-group-prepend"}).down()
    .appendElement({"tag" : "span", "class" : "input-group-text", "innerHtml" : "ID"}).up()
    .appendElement({"tag" : "span", "class" : "form-control",  "innerHtml" : pkg.M_id})
    .appendElement({"tag":"div","class":"input-group-append"}).down()
    .appendElement({"tag":"span", "class":"input-group-text", "id":"print_label_btn", "title":"imprimer les étiquettes."}).down()
    .appendElement({"tag":"i", "class":"fa fa-print", "style":"color: #495057;"}).up().up().up()
    .appendElement({"tag" : "div", "class" : "input-group mb-3"}).down()
    .appendElement({"tag" : "div", "class" : "input-group-prepend"}).down()
    .appendElement({"tag" : "span", "class" : "input-group-text", "innerHtml" : "Nom"}).up()
    .appendElement({"tag" : "input", "class" : "form-control", "type" : "text", "value" : pkg.M_name, "id" : pkg.M_id + "-name"}).up()
    .appendElement({"tag" : "div", "class" : "input-group mb-3", "id" : pkg.M_id  + "-descriptionInput"}).down()
    .appendElement({"tag" : "div", "class" : "input-group-prepend"}).down()
    .appendElement({"tag" : "span", "class" : "input-group-text", "innerHtml" : "Description"}).up().up()
    .appendElement({"tag" : "div", "class" : "input-group mb-3"}).down()
    .appendElement({"tag" : "div", "class" : "input-group-prepend"}).down()
    .appendElement({"tag" : "span", "class" : "input-group-text", "innerHtml" : "Quantité"}).up()
    .appendElement({"tag" : "input", "class" : "form-control", "type" : "number", "value" : pkg.M_quantity, "id" : pkg.M_id + "-quantity"})
    .appendElement({"tag" : "div", "class" : "input-group-append"}).down()
    .appendElement({"tag" : "button", "class" : "btn btn-outline-dark dropdown-toggle", "type" : "button", "data-toggle" : "dropdown", "aria-haspopup" : "true", "aria-expanded" : "false", "innerHtml" : pkg.M_unitOfMeasure.M_valueOf, "id" : pkg.M_id + "-currentUnitOfMeasure"})
    .appendElement({"tag" : "div", "class" : "dropdown-menu dropDownEnum", "id" : pkg.M_id + "-unitOfMeasureEnum"}).up().up().up()
    
    .appendElement({"tag" : "div", "class" : "col-5"}).down()
    .appendElement({"tag" : "div", "id" : pkg.M_id + "-picture_div", "style" : "height:100%;"}).up().up()
    
    .appendElement({"tag" : "div", "class" : "row"}).down()
    .appendElement({"tag":"label", "for" : pkg.M_id + "-packagesList", "innerHtml" : "Paquets","class" : "col-3 d-flex align-items-center"})
    .appendElement({"tag": "input", "class" : "form-control col-5 m-1", "id" : pkg.M_id + "-filterPackages","type" :"text", "placeholder" : "Filtrer.."}).up()
    
    .appendElement({"tag" : "div"}).down()
    .appendElement({"tag":"table", "class":"table table-hover table-condensed"}).down()
    .appendElement({"tag":"thead"}).down()
    .appendElement({"tag" : "tr"}).down()
    .appendElement({"tag" : "th", "innerHtml" : "ID Paquet", "style" : "width:30%;"})
    .appendElement({"tag" : "th", "innerHtml" : "Nom", "style":"width:60%;"})
    .appendElement({"tag" : "th", "style" : "width:10%;"}).down()
    .appendElement({"tag" : "button", "class" : "btn btn-success btn-sm", "id" : pkg.M_id +"-addPackageButton", "style":"display: inline-flex; height: 29px;"}).down()
    .appendElement({"tag" : "i", "class" : "fa fa-plus"}).up().up()
    .up().up()
    .appendElement({"tag" : "tbody", "id" : pkg.M_id + "-packagesList"})
    .up().up()
    
    .appendElement({"tag" : "div", "class" : "row"}).down()
    .appendElement({"tag":"label", "for" : pkg.M_id + "-itemsList", "innerHtml" : "Items","class" : "col-3 d-flex align-items-center"})
    .appendElement({"tag": "input", "class" : "form-control col-5 m-1", "id" : pkg.M_id + "-filterItems","type" :"text", "placeholder" : "Filtrer.."}).up()
    
    .appendElement({"tag" : "div"}).down()
    .appendElement({"tag":"table", "class":"table table-hover table-condensed"}).down()
    .appendElement({"tag":"thead"}).down()
    .appendElement({"tag" : "tr"}).down()
    .appendElement({"tag" : "th", "innerHtml" : "ID Item", "style" : "width:30%;"})
    .appendElement({"tag" : "th", "innerHtml" : "Nom", "style":"width:60%;"})
    .appendElement({"tag" : "th", "style" : "width:10%;"}).down()
    .appendElement({"tag" : "button", "class" : "btn btn-success btn-sm", "id" : pkg.M_id +"-addItemButton", "style":"display: inline-flex; height: 29px;"}).down()
    .appendElement({"tag" : "i", "class" : "fa fa-plus"}).up().up()
    .up().up()
    .appendElement({"tag" : "tbody", "id" : pkg.M_id + "-itemsList"})
    .up().up()
    
    .appendElement({"tag" : "div", "class" : "row"}).down()
    .appendElement({"tag":"label", "for" : pkg.M_id + "-itemsList", "innerHtml" : "Inventaires","class" : "col-3 d-flex align-items-center"})
    .appendElement({"tag": "input", "class" : "form-control col-5 m-1", "id" : pkg.M_id + "-filterInventories","type" :"text", "placeholder" : "Filtrer.."}).up()
    
    .appendElement({"tag" : "div"}).down()
    .appendElement({"tag":"table", "class":"table table-hover table-condensed"}).down()
    .appendElement({"tag":"thead"}).down()
    .appendElement({"tag" : "tr"}).down()
    .appendElement({"tag" : "th", "innerHtml" : "Localisation", "style" : "width:20%;"})
    .appendElement({"tag" : "th", "innerHtml" : "Stock", "style":"width:20%;"})
    .appendElement({"tag" : "th", "innerHtml" : "Stock minimal", "style":"width:20%;"})
    .appendElement({"tag" : "th", "innerHtml" : "Quantité pour re-order", "style":"width:30%;"})
    .appendElement({"tag" : "th", "style" : "width:10%;"}).down()
    .appendElement({"tag" : "button", "class" : "btn btn-success btn-sm", "id" : pkg.M_id +"-addInventoryButton", "style":"display: inline-flex; height: 29px;"}).down()
    .appendElement({"tag" : "i", "class" : "fa fa-plus"}).up().up()
    .up().up()
    .appendElement({"tag" : "tbody", "id" : pkg.M_id + "-inventoriesList"})
    .up().up()
    
    .appendElement({"tag" : "div", "class" : "row"}).down()
    .appendElement({"tag":"label", "for" : pkg.M_id + "-suppliersList", "innerHtml" : "Fournisseurs","class" : "col-3 d-flex align-items-center"})
    .appendElement({"tag": "input", "class" : "form-control col-5 m-1", "id" : pkg.M_id + "-filterSuppliers","type" :"text", "placeholder" : "Filtrer.."}).up()
    
    .appendElement({"tag" : "div"}).down()
    .appendElement({"tag":"table", "class":"table table-hover table-condensed"}).down()
    .appendElement({"tag":"thead"}).down()
    .appendElement({"tag" : "tr"}).down()
    .appendElement({"tag" : "th", "innerHtml" : "ID Transation", "style" : "width:15%;"})
    .appendElement({"tag" : "th", "innerHtml" : "Fournisseur", "style" : "width:15%;"})
    .appendElement({"tag" : "th", "innerHtml" : "Quantité", "style":"width:15%;"})
    .appendElement({"tag" : "th", "innerHtml" : "Prix", "style":"width:30%;"})
    .appendElement({"tag" : "th", "innerHtml" : "Date", "style":"width:20%;"})
    .appendElement({"tag" : "th", "style" : "width:5%;"}).down()
    .appendElement({"tag" : "button", "class" : "btn btn-success btn-sm", "id" : pkg.M_id +"-addSupplierButton", "style":"display: inline-flex; height: 29px;"}).down()
    .appendElement({"tag" : "i", "class" : "fa fa-plus"}).up().up()
    .up().up()
    .appendElement({"tag" : "tbody", "id" : pkg.M_id + "-suppliersList"})
    .up().up()
    
    
    .appendElement({"tag" : "div", "class" : "d-flex justify-content-center"}).down()
    .appendElement({"tag" : "button", "class" : "btn btn-primary disabled mr-3", "innerHtml" : "Enregistrer", "id" : pkg.M_id + "-saveBtn"})
    .appendElement({"tag" : "button", "class" : "btn disabled mr-3", "innerHtml" : "Annuler les modifications", "id" : pkg.M_id + "-cancelBtn"})
    .appendElement({"tag" : "button", "class" : "btn btn-danger", "innerHtml" : "Supprimer l'entité", "id" : pkg.M_id + "-deleteBtn", "data-toggle" : "modal", "data-target" : "#" + pkg.M_id + "-modal"})
    .appendElement({"tag" : "div", "class" : "modal fade", "id"  : pkg.M_id + "-modal", "tabindex" : "-1", "role"  :"dialog", "aria-labelledby" : pkg.M_id + "-modal", "aria-hidden" : "true"}).down()
    .appendElement({"tag" : "div", "class" : "modal-dialog", "role" : "document"}).down()
    .appendElement({"tag" : "div", "class" : "modal-content"}).down()
    .appendElement({"tag" : "div", "class" : "modal-header"}).down()
    .appendElement({"tag" : "h5", "class" : "modal-title", "innerHtml" : "Supprimer le paquet"})
    .appendElement({"tag" : "button", "class" : "close", "data-dismiss" : "modal", "aria-label" : "Close"}).down()
    .appendElement({"tag" : "span", "aria-hidden" : "true", "innerHtml" : "&times;"}).up().up()
    .appendElement({"tag" : "div", "class" : "modal-body", "innerHtml" : "Êtes-vous certain de vouloir supprimer ce paquet? Tous les inventaires et transactions y étant associés seront supprimés."})
    .appendElement({"tag" : "div", "class" : "modal-footer"}).down()
    .appendElement({"tag" : "button", "class" : "btn btn-secondary", "data-dismiss" : "modal", "innerHtml" : "Fermer", "id" : pkg.M_id + "-closeDelete"})
    .appendElement({"tag" : "button", "type" : "button", "class" : "btn btn-danger", "id" : pkg.M_id + "-deleteConfirmBtn", "innerHtml" : "Confirmer la suppression"})
    
    var printBtn = this.panel.getChildById("print_label_btn")
    printBtn.element.onmouseenter = function(){
        this.style.cursor = "pointer"
    }
     
    printBtn.element.onmouseleave = function(){
        this.style.cursor = "default"
    }
    
    // Here the code to print the items label on zebra printer.
    printBtn.element.onclick = function(pkg){
        return function(){
            printPackagesLabels([pkg])
        }
    }(pkg)
    
    server.languageManager.refresh()
        
    var pictureDiv = this.panel.getChildById(pkg.M_id + "-picture_div")
    new ImagePicker(pictureDiv,"/Catalogue_2/photo/" + pkg.M_id)
    var editors = initMultilanguageInput(this.panel.getChildById(pkg.M_id + "-descriptionInput"), "textarea", pkg, "M_description")
    for(var id in editors){

        editors[id].element.onchange = function(pkg){
            return function(){
                mainPage.adminPage.panel.getChildById("adminPackageSaveState").element.innerHTML = "*"
                mainPage.adminPage.adminPackagePage.panel.getChildById(pkg.M_id + "-saveBtn").element.classList.remove("disabled")
                mainPage.adminPage.adminPackagePage.panel.getChildById(pkg.M_id + "-cancelBtn").element.classList.remove("disabled")
     
                if(mainPage.adminPage.adminPackagePage.modifiedItems[pkg.M_id].get(pkg.M_id) != undefined){
                    var newPkg = jQuery.extend({},mainPage.adminPage.adminPackagePage.modifiedItems[pkg.M_id].get(pkg.M_id))
                    newPkg.M_description = pkg.M_description
                    mainPage.adminPage.adminPackagePage.modifiedItems[newPkg.M_id].set(newPkg.M_id, newPkg)
                }else{
                    var newPkg = jQuery.extend({},pkg)
                    newPkg.M_description = pkg.M_description
                 mainPage.adminPage.adminPackagePage.modifiedItems[newPkg.M_id].set(newPkg.M_id, newPkg)
                }
            }
        }(pkg)
    }

    
    mainPage.adminPage.adminPackagePage.panel.getChildById(pkg.M_id + "-deleteConfirmBtn").element.onclick = function(pkg){
        return function(){
            document.getElementById("waitingDiv").style.display = ""
            mainPage.adminPage.adminPackagePage.panel.getChildById(pkg.M_id + "-closeDelete").element.click()
         
            for(var i = 0; i < pkg.M_items.length; i++){
                server.entityManager.getEntityByUuid(pkg.M_items[i], true,
                    function(item, caller){
                        item.M_packaged.pop(caller.pkg)
                        server.entityManager.saveEntity(item,
                            function(success, caller){
                            },function(){},{})
                    },function(){},{"pkg" : pkg})
            }
            
            for(var i = 0; i < pkg.M_inventoried.length; i++){
                server.entityManager.removeEntity(pkg.M_inventoried[i],
                    function(success,caller){
                    },function(){},{})
            }
            
            for(var i = 0; i < pkg.M_supplied.length; i++){
                server.entityManager.getEntityByUuid(pkg.M_supplied[i],false,
                    function(transaction, caller){
                        server.entityManager.getEntityByUuid(transaction.M_supplier, false,
                            function(supplier,caller){
                                supplier.M_items.pop(caller.transaction)
                                server.entityManager.saveEntity(supplier,
                                    function(success,caller){
                                    },function(){},{})
                            },function(){}, {"transaction" : transaction})
                    },function(){}, {})
                
                server.entityManager.removeEntity(pkg.M_supplied[i], 
                    function(success,caller){
                    },function(){},{})
                
            }
            
            server.entityManager.removeEntity(pkg.UUID,
                function(success,caller){
                     mainPage.adminPage.adminPackagePage.panel.getChildById(caller.pkg.M_id + "-control").delete()
                    mainPage.adminPage.adminPackagePage.panel.getChildById(caller.pkg.M_id + "-selector").delete()
                    document.getElementById("waitingDiv").style.display = "none"
                },function(){}, {"pkg" : pkg})
            
                
            
                
            /*server.entityManager.removeEntity(pkg.UUID,
            function(success,caller){
                mainPage.adminPage.adminPackagePage.panel.getChildById(caller.pkg.M_id + "-control").delete()
                mainPage.adminPage.adminPackagePage.panel.getChildById(caller.pkg.M_id + "-selector").delete()
                document.getElementById("waitingDiv").style.display = "none"
            }, function(){},{"pkg" : pkg})*/
        }
    }(pkg)
    
    mainPage.adminPage.adminPackagePage.panel.getChildById(pkg.M_id + "-addPackageButton").element.onclick = function(pkg){
        return function(){
            var date = new Date()
            mainPage.adminPage.adminPackagePage.panel.getChildById(pkg.M_id + "-packagesList").prependElement({"tag" : "tr", "id" : date + "-row"}).down()
            .appendElement({"tag" : "th", "scope" : "row"}).down()
            .appendElement({"tag" : "input", "class" : "form-control", "type" : "text", "placeholder" : "ID", "id" : date  + "-packageID"}).up()
            .appendElement({"tag" : "td", "scope" : "row"})
             .appendElement({"tag":"td"}).down()
            .appendElement({"tag" : "button", "class" : "btn btn-success btn-sm", "id" : date + "-confirmBtn"}).down()
            .appendElement({"tag" : "i", "class" : "fa fa-check"}).up()
            
            autocomplete("CatalogSchema.PackageType", mainPage.adminPage.adminPackagePage.panel.getChildById(date + "-packageID"))
            
            mainPage.adminPage.adminPackagePage.panel.getChildById(date + "-packageID").element.oncomplete = function(uuid){
                mainPage.adminPage.adminPackagePage.panel[date+"-packageID"] = uuid
            }
            
             mainPage.adminPage.adminPackagePage.panel.getChildById(date+ "-confirmBtn").element.onclick = function (date, pkg) {
            return function(){
                document.getElementById("waitingDiv").style.display = ""
                    var pkgToAdd = mainPage.adminPage.adminPackagePage.panel[date+"-packageID"]

                    pkg.M_packages.push(pkgToAdd)
                   
                    server.entityManager.saveEntity(pkg,
                        function(success,caller){
                            mainPage.adminPage.adminPackagePage.panel.getChildById(date + "-row").delete()
                                document.getElementById("waitingDiv").style.display = "none"
                                server.entityManager.getEntityByUuid(caller.pkgToAdd,false,
                                    function(pkgToAdd,caller){
                                        mainPage.adminPage.adminPackagePage.appendPackage(pkgToAdd, caller.success)
                                    },function(){}, { "success" : success})
                                
                        },function(){},{"pkgToAdd" : pkgToAdd})

                            

               }
            }(date, pkg)
        }
    }(pkg)
    
    mainPage.adminPage.adminPackagePage.panel.getChildById(pkg.M_id + "-addItemButton").element.onclick = function(pkg){
        return function(){
            var date = new Date()
            mainPage.adminPage.adminPackagePage.panel.getChildById(pkg.M_id + "-itemsList").prependElement({"tag" : "tr", "id" : date + "-row"}).down()
            .appendElement({"tag" : "th", "scope" : "row"}).down()
            .appendElement({"tag" : "input", "class" : "form-control", "type" : "text", "placeholder" : "ID", "id" : date  + "-itemID"}).up()
            .appendElement({"tag" : "td", "scope" : "row"})
             .appendElement({"tag":"td"}).down()
            .appendElement({"tag" : "button", "class" : "btn btn-success btn-sm", "id" : date + "-confirmBtn"}).down()
            .appendElement({"tag" : "i", "class" : "fa fa-check"}).up()
            
            autocomplete("CatalogSchema.ItemType", mainPage.adminPage.adminPackagePage.panel.getChildById(date + "-itemID"))
            
            mainPage.adminPage.adminPackagePage.panel.getChildById(date + "-itemID").element.oncomplete = function(uuid){
                mainPage.adminPage.adminPackagePage.panel[date+"-itemID"] = uuid
            }
            
             mainPage.adminPage.adminPackagePage.panel.getChildById(date+ "-confirmBtn").element.onclick = function (date, pkg) {
            return function(){
                document.getElementById("waitingDiv").style.display = ""
                    var newItem = mainPage.adminPage.adminPackagePage.panel[date+"-itemID"]
                    server.entityManager.getEntityByUuid(newItem,false,
                        function(newItem,caller){
                             var newPkg = caller.pkg
                            newPkg.M_items.push(newItem)
                            newItem.M_packaged.push(newPkg)
                            server.entityManager.saveEntity(newItem,
                                function(newItem,caller){
                                    server.entityManager.saveEntity(caller.newPkg,
                                    function(success,caller){
                                        mainPage.adminPage.adminPackagePage.panel.getChildById(date + "-row").delete()
                                            document.getElementById("waitingDiv").style.display = "none"
                                            mainPage.adminPage.adminPackagePage.appendItem(caller.item, success)
                                    },function(){},{"item" : newItem})
                                },function(){}, {"newPkg" : newPkg})
                            
                        },function(){}, {"pkg" : pkg})
               }
            }(date, pkg)
        }
    }(pkg)
    
    mainPage.adminPage.adminPackagePage.panel.getChildById(pkg.M_id + "-addInventoryButton").element.onclick = function(pkg){
        return function(){
            var date = new Date()
            mainPage.adminPage.adminPackagePage.panel.getChildById(pkg.M_id + "-inventoriesList").prependElement({"tag" : "tr", "id" : date + "-row"}).down()
            .appendElement({"tag" : "th", "scope" : "row"}).down()
            .appendElement({"tag" : "input", "class" : "form-control", "type" : "text", "placeholder" : "ID Localisation", "id" : date  + "-localisationID"}).up()
            .appendElement({"tag" : "td", "scope" : "row"}).down()
            .appendElement({"tag" : "input", "class" : "form-control", "type" : "number", "value":0,"id" : date  + "-quantity"}).up()
            .appendElement({"tag" : "td", "scope" : "row"}).down()
            .appendElement({"tag" : "input", "class" : "form-control", "type" : "number", "value" : 0, "id" : date  + "-safetyStock"}).up()
            .appendElement({"tag" : "td", "scope" : "row"}).down()
            .appendElement({"tag" : "input", "class" : "form-control", "type" : "number","value" :0, "id" : date  + "-reorderQty"}).up()
             .appendElement({"tag":"td"}).down()
            .appendElement({"tag" : "button", "class" : "btn btn-success btn-sm", "id" : date + "-confirmBtn"}).down()
            .appendElement({"tag" : "i", "class" : "fa fa-check"}).up()
            
            autocomplete("CatalogSchema.LocalisationType", mainPage.adminPage.adminPackagePage.panel.getChildById(date+"-localisationID"))
            
            mainPage.adminPage.adminPackagePage.panel.getChildById(date + "-localisationID").element.oncomplete = function(uuid){
                mainPage.adminPage.adminPackagePage.panel[date+"-tempLocalisation"] = uuid
            }
            
             mainPage.adminPage.adminPackagePage.panel.getChildById(date+ "-confirmBtn").element.onclick = function (date, pkg) {
            return function(){
                document.getElementById("waitingDiv").style.display = ""
                
                
                        var newLocalisation = mainPage.adminPage.adminPackagePage.panel[date+"-tempLocalisation"]
                         var newInventory = new CatalogSchema.InventoryType
                         
                         server.entityManager.createEntity(catalog.UUID, "M_inventories",newInventory,
                            function(newInventory, caller){
                                newInventory.M_quantity = mainPage.adminPage.adminPackagePage.panel.getChildById(caller.date+"-quantity").element.value
                                newInventory.M_safetyStock = mainPage.adminPage.adminPackagePage.panel.getChildById(caller.date+"-safetyStock").element.value
                                newInventory.M_reorderQty = mainPage.adminPage.adminPackagePage.panel.getChildById(caller.date+"-reorderQty").element.value
                                 newInventory.M_located = caller.localisation
                                 newInventory.M_package = caller.pkg.UUID
                               
                            
                                server.entityManager.saveEntity(newInventory,
                                    function(success,caller){
                                        caller.pkg.M_inventoried.push(success)
                                        server.entityManager.saveEntity(caller.pkg,
                                            function(pkg, caller){
                                                mainPage.adminPage.adminPackagePage.panel.getChildById(caller.date + "-row").delete()
                                                document.getElementById("waitingDiv").style.display = "none"
                                                mainPage.adminPage.adminPackagePage.appendInventory(caller.inventory, pkg)
                                            },function(){},{"inventory" : success, "date" : caller.date})
                                            
                                    },function(){},{"pkg" : caller.pkg, "date" : caller.date})
                            },function(){}, {"date" : date, "pkg" : pkg, "localisation" : newLocalisation})
                       
              
               }
            }(date, pkg)
             
                
        }
    }(pkg)
    
    mainPage.adminPage.adminPackagePage.panel.getChildById(pkg.M_id + "-addSupplierButton").element.onclick = function(pkg){
        return function(){
            var date = new Date()
            mainPage.adminPage.adminPackagePage.panel.getChildById(pkg.M_id + "-suppliersList").prependElement({"tag" : "tr", "id" : date + "-row"}).down()
            .appendElement({"tag" : "th", "scope" : "row"}).down()
            .appendElement({"tag" : "input", "class" : "form-control", "type" : "text", "placeholder" : "ID", "id" : date  + "-transactionID"}).up()
            .appendElement({"tag" : "th", "scope" : "row"}).down()
            .appendElement({"tag" : "input", "class" : "form-control", "type" : "text", "placeholder" : "ID", "id" : date  + "-id"}).up()
            .appendElement({"tag" : "td"}).down()
            .appendElement({"tag" : "input", "class" : "form-control", "type" : "number", "placeholder" : "0", "id" : date + "-qty"}).up()
            .appendElement({"tag" : "td"}).down()
            .appendElement({"tag" :"div", "class" : "input-group"}).down()
            .appendElement({"tag":"input", "class" : "form-control", "type" : "number", "placeholder" : "0", "id" : date + "-price"})
            .appendElement({"tag":"div", "class" : "input-group-append"}).down()
            .appendElement({"tag" :"button", "class": "btn btn-outline-dark dropdown-toggle", "type" : "button", "data-toggle" : "dropdown", "aria-haspopup" : "true", "aria-expanded":"false", "innerHtml" : "Currency", "id" : date + "-currentCurrency"})
            .appendElement({"tag":"div", "class":"dropdown-menu dropDown", "id" : date + "-currencyTypes"}).down()
            .up().up().up().up()
            .appendElement({"tag" : "td"}).down()
            .appendElement({"tag" : "input", "type" : "date", "class"  :"form-control", "id" : date + "-date"}).up()
             .appendElement({"tag":"td"}).down()
            .appendElement({"tag" : "button", "class" : "btn btn-success btn-sm", "id" : date + "-confirmBtn"}).down()
            .appendElement({"tag" : "i", "class" : "fa fa-check"}).up()
            
            var fields = ["UUID", "M_name"]
            
           
            
            autocomplete("CatalogSchema.SupplierType", mainPage.adminPage.adminPackagePage.panel.getChildById(date + "-id"), fields)
            
           mainPage.adminPage.adminPackagePage.panel.getChildById(date + "-id").element.oncomplete = function(uuid){
                mainPage.adminPage.adminPackagePage.panel[date+"-tempSupplier"] = uuid
            }
            
            
            
            
            for(var i = 0; i < mainPage.adminPage.adminPackagePage.currencies.length; i++){
                mainPage.adminPage.adminPackagePage.panel.getChildById(date + "-currencyTypes").appendElement({"tag" : "a", "class" : "dropdown-item", "innerHtml" : mainPage.adminPage.adminPackagePage.currencies[i], "id" : date + "-currencyTypes-" + mainPage.adminPage.adminPackagePage.currencies[i]})
                
                mainPage.adminPage.adminPackagePage.panel.getChildById(date + "-currencyTypes-" + mainPage.adminPage.adminPackagePage.currencies[i]).element.onclick = function (divID, currency){
                    return function(){
                        mainPage.adminPage.adminPackagePage.panel.getChildById(divID + "-currentCurrency").element.innerHTML = currency
                    }
                }(date, mainPage.adminPage.adminPackagePage.currencies[i])
            }
                    
            mainPage.adminPage.adminPackagePage.panel.getChildById(date+ "-confirmBtn").element.onclick = function (date, pkg) {
               return function(){
                   
                    
                    if(mainPage.adminPage.adminPackagePage.panel[date+"-tempSupplier"] != ""){
                       document.getElementById("waitingDiv").style.display = ""
                
                        var newItem = new CatalogSchema.ItemSupplierType()
                        newItem.M_supplier = mainPage.adminPage.adminPackagePage.panel[date+"-tempSupplier"]
                        newItem.M_quantity = mainPage.adminPage.adminPackagePage.panel.getChildById(date + "-qty").element.value
                        newItem.M_id = mainPage.adminPage.adminPackagePage.panel.getChildById(date + "-transactionID").element.value
                        newItem.M_price = new CatalogSchema.PriceType()
                        newItem.M_price.M_valueOf = mainPage.adminPage.adminPackagePage.panel.getChildById(date + "-price").element.value
                        newItem.M_package = pkg.UUID
                        if(mainPage.adminPage.adminPackagePage.panel.getChildById(date + "-currentCurrency").element.innerHTML != "Currency"){
                            newItem.M_price.M_currency.M_valueOf = mainPage.adminPage.adminPackagePage.panel.getChildById(date + "-currentCurrency").element.innerHTML
                        }
                        
                        newItem.M_date = mainPage.adminPage.adminPackagePage.panel.getChildById(date + "-date").element.value
                        server.entityManager.createEntity(pkg, "M_supplied", newItem, 
                            function(success, caller){
             
                                 caller.pkg.M_supplied.push(success.UUID)
                          
                                 server.entityManager.saveEntity(caller.pkg, 
                                    function(result,caller){
                                        
                                        mainPage.adminPage.adminPackagePage.panel.getChildById(date + "-row").delete()
                                        
                                        document.getElementById("waitingDiv").style.display = "none"
                                        mainPage.adminPage.adminPackagePage.appendItemSupplier(caller.newLine, caller.pkg)
                                    },function(){},{"pkg" : caller.pkg, "newLine" : success})
                                
                            },function(){},{"pkg" : pkg}) 
                    }
                    
                                    
                               
                                    
                                    
                                    
                                    
                              
               }
            }(date, pkg)
        }
    }(pkg)
    
    
    for(var i  =0 ; i< this.unitOfMeasure.length; i++){
        this.panel.getChildById(pkg.M_id + "-unitOfMeasureEnum").appendElement({"tag" : "a", "class" : "dropdown-item","id" : pkg.M_id + "-unitOfMeasureEnum-" + this.unitOfMeasure[i].Value, "innerHtml" : this.unitOfMeasure[i].Value})
    
        this.panel.getChildById(pkg.M_id + "-unitOfMeasureEnum-" + this.unitOfMeasure[i].Value).element.onclick = function (divID, unitOfMeasure, pkg){
            return function(){
                mainPage.adminPage.adminPackagePage.panel.getChildById(divID ).element.innerHTML = unitOfMeasure
                mainPage.adminPage.panel.getChildById("adminPackageSaveState").element.innerHTML = "*"
                mainPage.adminPage.adminPackagePage.panel.getChildById(pkg.M_id + "-saveBtn").element.classList.remove("disabled")
                mainPage.adminPage.adminPackagePage.panel.getChildById(pkg.M_id + "-cancelBtn").element.classList.remove("disabled")
                var newPkg = jQuery.extend({}, pkg)
                newPkg.M_unitOfMeasure.M_valueOf = unitOfMeasure
                mainPage.adminPage.adminPackagePage.modifiedItems[pkg.M_id].set(newPkg.M_id, newPkg)
            }
        }(pkg.M_id + "-currentUnitOfMeasure", this.unitOfMeasure[i].Value, pkg)
    }
    
    for(var i = 0 ; i < pkg.M_packages.length; i++){
        server.entityManager.getEntityByUuid(pkg.M_packages[i],false,
            function(pkg, caller){
                mainPage.adminPage.adminPackagePage.appendPackage(pkg, caller.package)
            }, function(){}, {"package" : pkg})
    }
    
    for(var i = 0 ; i < pkg.M_items.length; i++){
        server.entityManager.getEntityByUuid(pkg.M_items[i],false,
            function(item, caller){
                mainPage.adminPage.adminPackagePage.appendItem(item, caller.package)
            }, function(){}, {"package" : pkg})
    }
    
    for(var i = 0; i < pkg.M_inventoried.length; i++){
        server.entityManager.getEntityByUuid(pkg.M_inventoried[i], false,
            function(inventory, caller){
                mainPage.adminPage.adminPackagePage.appendInventory(inventory, caller.package)
            },function(){},{"package" : pkg})
    }
    
    for(var i = 0; i < pkg.M_supplied.length; i++){
        server.entityManager.getEntityByUuid(pkg.M_supplied[i],false,
            function(supplier, caller){
                mainPage.adminPage.adminPackagePage.appendItemSupplier(supplier, caller.package)
            },function(){},{"package" : pkg})
    }
    
    mainPage.adminPage.adminPackagePage.panel.getChildById(pkg.M_id + "-filterPackages").element.onkeyup = function(packageID){
        return function(){
            var input, filter, table, tr, td, i;
            input = document.getElementById(packageID+"-filterPackages");
            filter = input.value.toUpperCase();
            table = document.getElementById(packageID + "-packagesList");
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
    }(pkg.M_id)
    
     mainPage.adminPage.adminPackagePage.panel.getChildById(pkg.M_id + "-filterItems").element.onkeyup = function(packageID){
        return function(){
            var input, filter, table, tr, td, i;
            input = document.getElementById(packageID+"-filterItems");
            filter = input.value.toUpperCase();
            table = document.getElementById(packageID + "-itemsList");
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
    }(pkg.M_id)
    
    mainPage.adminPage.adminPackagePage.panel.getChildById(pkg.M_id + "-filterInventories").element.onkeyup = function(packageID){
        return function(){
            var input, filter, table, tr, td, i;
            input = document.getElementById(packageID+"-filterInventories");
            filter = input.value.toUpperCase();
            table = document.getElementById(packageID + "-inventoriesList");
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
    }(pkg.M_id)
    
    mainPage.adminPage.adminPackagePage.panel.getChildById(pkg.M_id + "-filterSuppliers").element.onkeyup = function(packageID){
        return function(){
            var input, filter, table, tr, td, i;
            input = document.getElementById(packageID+"-filterSuppliers");
            filter = input.value.toUpperCase();
            table = document.getElementById(packageID + "-suppliersList");
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
    }(pkg.M_id)
    
    mainPage.adminPage.adminPackagePage.panel.getChildById(pkg.M_id + "-name").element.onchange = function(pkg){
       return function(){
           mainPage.adminPage.panel.getChildById("adminPackageSaveState").element.innerHTML = "*"
            mainPage.adminPage.adminPackagePage.panel.getChildById(pkg.M_id + "-saveBtn").element.classList.remove("disabled")
            mainPage.adminPage.adminPackagePage.panel.getChildById(pkg.M_id + "-cancelBtn").element.classList.remove("disabled")
            if(mainPage.adminPage.adminPackagePage.modifiedItems[pkg.M_id].get(pkg.M_id) != undefined){
                var newPkg = jQuery.extend({},mainPage.adminPage.adminPackagePage.modifiedItems[pkg.M_id].get(pkg.M_id))
                newPkg.M_name = mainPage.adminPage.adminPackagePage.panel.getChildById(newPkg.M_id + "-name").element.value
                mainPage.adminPage.adminPackagePage.modifiedItems[newPkg.M_id].set(newPkg.M_id, newPkg)
            }else{
                var newPkg = jQuery.extend({},pkg)
                newPkg.M_name = mainPage.adminPage.adminPackagePage.panel.getChildById(newPkg.M_id + "-name").element.value
                mainPage.adminPage.adminPackagePage.modifiedItems[newPkg.M_id].set(newPkg.M_id, newPkg)
            }

       }
    }(pkg)
    
 
    mainPage.adminPage.adminPackagePage.panel.getChildById(pkg.M_id + "-quantity").element.onchange = function(pkg){
       return function(){
           mainPage.adminPage.panel.getChildById("adminPackageSaveState").element.innerHTML = "*"
            mainPage.adminPage.adminPackagePage.panel.getChildById(pkg.M_id + "-saveBtn").element.classList.remove("disabled")
            mainPage.adminPage.adminPackagePage.panel.getChildById(pkg.M_id + "-cancelBtn").element.classList.remove("disabled")
            if(mainPage.adminPage.adminPackagePage.modifiedItems[pkg.M_id].get(pkg.M_id) != undefined){
                var newPkg = jQuery.extend({},mainPage.adminPage.adminPackagePage.modifiedItems[pkg.M_id].get(pkg.M_id))
                newPkg.M_quantity = mainPage.adminPage.adminPackagePage.panel.getChildById(newPkg.M_id + "-quantity").element.value
                mainPage.adminPage.adminPackagePage.modifiedItems[newPkg.M_id].set(newPkg.M_id, newPkg)
            }else{
                var newPkg = jQuery.extend({},pkg)
                 newPkg.M_quantity = mainPage.adminPage.adminPackagePage.panel.getChildById(newPkg.M_id + "-quantity").element.value
                mainPage.adminPage.adminPackagePage.modifiedItems[newPkg.M_id].set(newPkg.M_id, newPkg)
            }

       }
    }(pkg)
    
    mainPage.adminPage.adminPackagePage.panel.getChildById(pkg.M_id + "-cancelBtn").element.onclick = function(pkg){
        return function(){
           mainPage.adminPage.panel.getChildById("adminPackageSaveState").element.innerHTML = ""
            mainPage.adminPage.adminPackagePage.panel.getChildById(pkg.M_id + "-saveBtn").element.classList.add("disabled")
            mainPage.adminPage.adminPackagePage.panel.getChildById(pkg.M_id + "-cancelBtn").element.classList.add("disabled")
            mainPage.adminPage.adminPackagePage.loadAdminControl(pkg)
            mainPage.adminPage.adminPackagePage.panel.getChildById(pkg.M_id + "-control").element.classList.add("active")
             mainPage.adminPage.adminPackagePage.panel.getChildById(pkg.M_id + "-control").element.classList.add("show")
        }
    }(pkg)
    
    mainPage.adminPage.adminPackagePage.panel.getChildById(pkg.M_id + "-saveBtn").element.onclick = function(pkg){
        return function(){
            
            if(mainPage.adminPage.adminPackagePage.modifiedItems[pkg.M_id].size > 0){
                 var i = 0
                 document.getElementById("waitingDiv").style.display = ""

                for(var object of mainPage.adminPage.adminPackagePage.modifiedItems[pkg.M_id].values()){
                
                    server.entityManager.saveEntity(object,
                        function(result, caller){
                           
                            caller.i++
                             document.getElementById("waitingDiv").style.display = "none"
                            if(caller.i == caller.length){
                                
                                if(mainPage.adminPage.adminPackagePage.removedItems[caller.pkg.M_id].size > 0){
                                    var j = 0
                                    for(var removeObject of mainPage.adminPage.adminPackagePage.removedItems[caller.pkg.M_id].values()){
                                     server.entityManager.removeEntity(removeObject.UUID,
                                        function(result,caller){
                                            caller.j++
                                            if(caller.j == caller.length){
                                                
                                                server.entityManager.getEntityByUuid(caller.pkg.UUID, true,
                                                    function(pkg,caller){
                                                        mainPage.adminPage.panel.getChildById("adminPackageSaveState").element.innerHTML = ""
                                                   
                                                        mainPage.adminPage.adminPackagePage.panel.getChildById(pkg.M_id + "-selector").element.innerHTML = pkg.M_name
                                                        mainPage.adminPage.adminPackagePage.loadAdminControl(pkg)
                                                    mainPage.adminPage.adminPackagePage.panel.getChildById(pkg.M_id + "-saveBtn").element.classList.add("disabled")
                                                    mainPage.adminPage.adminPackagePage.panel.getChildById(pkg.M_id + "-cancelBtn").element.classList.add("disabled")
                                                    
                                                    mainPage.adminPage.adminPackagePage.panel.getChildById(pkg.M_id + "-control").element.classList.add("active")
                                                     mainPage.adminPage.adminPackagePage.panel.getChildById(pkg.M_id + "-control").element.classList.add("show")
                                                    },function(){},{})
                                                
                                            }
                                         
                                        },function(){},{"length" : mainPage.adminPage.adminPackagePage.removedItems[pkg.M_id].size, "j" : j, "pkg" : caller.pkg})
                       
                                     }
                                }else{
                                   server.entityManager.getEntityByUuid(caller.pkg.UUID, true,
                                    function(pkg,caller){
                                        mainPage.adminPage.panel.getChildById("adminPackageSaveState").element.innerHTML = ""
                                   
                                        mainPage.adminPage.adminPackagePage.panel.getChildById(pkg.M_id + "-selector").element.innerHTML = pkg.M_name
                                        mainPage.adminPage.adminPackagePage.loadAdminControl(pkg)
                                    mainPage.adminPage.adminPackagePage.panel.getChildById(pkg.M_id + "-saveBtn").element.classList.add("disabled")
                                    mainPage.adminPage.adminPackagePage.panel.getChildById(pkg.M_id + "-cancelBtn").element.classList.add("disabled")
                                    
                                    mainPage.adminPage.adminPackagePage.panel.getChildById(pkg.M_id + "-control").element.classList.add("active")
                                     mainPage.adminPage.adminPackagePage.panel.getChildById(pkg.M_id + "-control").element.classList.add("show")
                                    },function(){},{})
                                }
                               
                                
                            }
                        }, function () {}, {"i" : i, "length" : mainPage.adminPage.adminPackagePage.modifiedItems[pkg.M_id].size, "pkg" : pkg})
                }
            }else{
                 document.getElementById("waitingDiv").style.display = "none"
                if(mainPage.adminPage.adminPackagePage.removedItems[pkg.M_id].size > 0){
                    
                    var j = 0
                    for(var removeObject of mainPage.adminPage.adminPackagePage.removedItems[pkg.M_id].values()){
                     server.entityManager.removeEntity(removeObject.UUID,
                        function(result,caller){
                            caller.j++
                            if(caller.j == caller.length){
                                server.entityManager.getEntityByUuid(caller.pkg.UUID, true,
                                function(pkg,caller){
                                    mainPage.adminPage.panel.getChildById("adminPackageSaveState").element.innerHTML = ""
                                   
                                    mainPage.adminPage.adminPackagePage.panel.getChildById(pkg.M_id + "-selector").element.innerHTML = pkg.M_name
                                    mainPage.adminPage.adminPackagePage.loadAdminControl(pkg)
                                mainPage.adminPage.adminPackagePage.panel.getChildById(pkg.M_id + "-saveBtn").element.classList.add("disabled")
                                mainPage.adminPage.adminPackagePage.panel.getChildById(pkg.M_id + "-cancelBtn").element.classList.add("disabled")
                                
                                mainPage.adminPage.adminPackagePage.panel.getChildById(pkg.M_id + "-control").element.classList.add("active")
                                 mainPage.adminPage.adminPackagePage.panel.getChildById(pkg.M_id + "-control").element.classList.add("show")
                                },function(){},{})
                            }
                         
                        },function(){},{"length" : mainPage.adminPage.adminPackagePage.removedItems[pkg.M_id].size, "j" : j, "pkg" : pkg})
                    }
                }
                
            }
        }
    }(pkg)
    
    
}

AdminPackagePage.prototype.appendItem = function(item, pkg){
    mainPage.adminPage.adminPackagePage.panel.getChildById(pkg.M_id +"-itemsList").appendElement({"tag" : "tr", "id" : item.M_id + "-adminItemRow"}).down()
    .appendElement({"tag" :"td","innerHtml" : item.M_id, "id" : item.M_id + "-itemID", "class" : "tabLink"})
    .appendElement({"tag" : "td", "innerHtml" : item.M_name, "id" : item.M_id + "-itemName"})
    .appendElement({"tag":"td"}).down()
    .appendElement({"tag" : "button", "class" : "btn btn-danger btn-sm", "id" : item.M_id + "-deleteItemRowAdminBtn", "style":"display: inline-flex; height: 29px;"}).down()
    .appendElement({"tag" : "i", "class" : "fa fa-trash-o"}).up()
    
    mainPage.adminPage.adminPackagePage.panel.getChildById(item.M_id + "-deleteItemRowAdminBtn").element.onclick = function(item, pkg){
        return function(){
            mainPage.adminPage.panel.getChildById("adminPackageSaveState").element.innerHTML = "*"
            mainPage.adminPage.adminPackagePage.removedItems[pkg.M_id].set(item.M_id + "-item", item)
            var row = mainPage.adminPage.adminPackagePage.panel.getChildById(item.M_id + "-adminItemRow")
            row.delete()
            mainPage.adminPage.adminPackagePage.panel.getChildById(pkg.M_id + "-saveBtn").element.classList.remove("disabled")
            mainPage.adminPage.adminPackagePage.panel.getChildById(pkg.M_id + "-cancelBtn").element.classList.remove("disabled")
        }
    }(item, pkg)
    
  mainPage.adminPage.adminPackagePage.panel.getChildById(item.M_id + "-itemID").element.onclick = function(item){
        return function(){
            mainPage.itemDisplayPage.displayTabItem(item)
        }
    }(item)
    
}

AdminPackagePage.prototype.appendPackage = function(pkgToAdd, pkg){
    mainPage.adminPage.adminPackagePage.panel.getChildById(pkg.M_id +"-packagesList").appendElement({"tag" : "tr", "id" : pkgToAdd.M_id + "-adminPackageRow"}).down()
    .appendElement({"tag" :"td","innerHtml" : pkgToAdd.M_id, "id" : pkgToAdd.M_id + "-packageID", "class" : "tabLink"})
    .appendElement({"tag" : "td", "innerHtml" : pkgToAdd.M_name, "id" : pkgToAdd.M_id + "-packageName"})
    .appendElement({"tag":"td"}).down()
    .appendElement({"tag" : "button", "class" : "btn btn-danger btn-sm", "id" : pkgToAdd.M_id + "-deletePackageRowBtn", "style":"display: inline-flex; height: 29px;"}).down()
    .appendElement({"tag" : "i", "class" : "fa fa-trash-o"}).up()
    
    mainPage.adminPage.adminPackagePage.panel.getChildById(pkgToAdd.M_id + "-deletePackageRowBtn").element.onclick = function(pkgToAdd, pkg){
        return function(){
            mainPage.adminPage.panel.getChildById("adminPackageSaveState").element.innerHTML = "*"
            mainPage.adminPage.adminPackagePage.removedItems[pkg.M_id].set(pkgToAdd.M_id + "-package", pkgToAdd)
            var row = mainPage.adminPage.adminPackagePage.panel.getChildById(pkgToAdd.M_id + "-adminPackageRow")
            row.delete()
            mainPage.adminPage.adminPackagePage.panel.getChildById(pkg.M_id + "-saveBtn").element.classList.remove("disabled")
            mainPage.adminPage.adminPackagePage.panel.getChildById(pkg.M_id + "-cancelBtn").element.classList.remove("disabled")
        }
    }(pkgToAdd, pkg)
    
    mainPage.adminPage.adminPackagePage.panel.getChildById(pkgToAdd.M_id + "-packageID").element.onclick = function(item_package){
        return function(){
            mainPage.packageDisplayPage.displayTabItem(item_package)
        }
    }(pkgToAdd)
}

AdminPackagePage.prototype.appendInventory = function(inventory, pkg){
    mainPage.adminPage.adminPackagePage.panel.getChildById(pkg.M_id +"-inventoriesList").appendElement({"tag" : "tr", "id" : inventory.UUID + "-adminItemRow"}).down()
    .appendElement({"tag" :"td", "id" : inventory.UUID + "-inventoryLocalisation"})
    .appendElement({"tag" : "td"}).down()
    .appendElement({"tag" : "input", "type" : "number", "class" : "form-control", "value" : inventory.M_quantity, "id" : inventory.UUID + "-qty"}).up()
    .appendElement({"tag":"td"}).down()
    .appendElement({"tag" : "input", "type" : "number", "class" : "form-control", "value" : inventory.M_safetyStock, "id" : inventory.UUID + "-safetyStock"}).up()
    .appendElement({"tag":"td"}).down()
    .appendElement({"tag" : "input", "type" : "number", "class" : "form-control", "value" : inventory.M_reorderQty, "id" : inventory.UUID + "-reorderQty"}).up()
    .appendElement({"tag":"td"}).down()
    .appendElement({"tag" : "button", "class" : "btn btn-danger btn-sm", "id" : inventory.UUID + "-deleteItemRowAdminBtn", "style":"display: inline-flex; height: 29px;"}).down()
    .appendElement({"tag" : "i", "class" : "fa fa-trash-o"}).up()
    
    mainPage.adminPage.adminPackagePage.panel.getChildById(inventory.UUID  + "-adminItemRow").element.onchange = function(inventory, packageID){
        return function(){
            mainPage.adminPage.panel.getChildById("adminPackageSaveState").element.innerHTML = "*"
            var stock = mainPage.adminPage.adminPackagePage.panel.getChildById(inventory.UUID+"-qty").element.value
            var safetyStock = mainPage.adminPage.adminPackagePage.panel.getChildById(inventory.UUID + "-safetyStock").element.value
            var reorderQty = mainPage.adminPage.adminPackagePage.panel.getChildById(inventory.UUID + "-reorderQty").element.value
            var newInventory = jQuery.extend({}, inventory)
            newInventory.M_quantity = stock
            newInventory.M_safetyStock = safetyStock
            newInventory.M_reorderQty= reorderQty
            mainPage.adminPage.adminPackagePage.modifiedItems[packageID].set(newInventory.UUID, newInventory)

            mainPage.adminPage.adminPackagePage.panel.getChildById(packageID + "-saveBtn").element.classList.remove("disabled")
            mainPage.adminPage.adminPackagePage.panel.getChildById(packageID + "-cancelBtn").element.classList.remove("disabled")
            
        }
        
        
    }(inventory, pkg.M_id)
    
    server.entityManager.getEntityByUuid(inventory.M_located, false,
        function(localisation, caller){
            mainPage.adminPage.adminPackagePage.panel.getChildById(caller.inventoryID).appendElement({"tag" : "div", "innerHtml" : localisation.M_id})
        },function(){}, {"inventoryID" : inventory.UUID + "-inventoryLocalisation"})
    
    mainPage.adminPage.adminPackagePage.panel.getChildById(inventory.UUID + "-deleteItemRowAdminBtn").element.onclick = function(inventory, pkg){
        return function(){
            mainPage.adminPage.panel.getChildById("adminPackageSaveState").element.innerHTML = "*"
            mainPage.adminPage.adminPackagePage.removedItems[pkg.M_id].set(inventory.UUID + "-inventory", inventory)
            var row = mainPage.adminPage.adminPackagePage.panel.getChildById(inventory.UUID + "-adminItemRow")
            row.delete()
            mainPage.adminPage.adminPackagePage.panel.getChildById(pkg.M_id + "-saveBtn").element.classList.remove("disabled")
            mainPage.adminPage.adminPackagePage.panel.getChildById(pkg.M_id + "-cancelBtn").element.classList.remove("disabled")
        }
    }(inventory, pkg)
    
    
}

AdminPackagePage.prototype.appendItemSupplier = function(supplier, pkg){
    if(supplier.M_price == ""){
        supplier.M_price = new CatalogSchema.PriceType
        supplier.M_price.M_valueOf = 0
    }
    if(supplier.M_price.M_currency == null){
        supplier.M_price.M_currency = new CatalogSchema.CurrencyType
        supplier.M_price.M_currency.M_valueOf = ""
    }
    mainPage.adminPage.adminPackagePage.panel.getChildById(pkg.M_id+"-suppliersList").appendElement({"tag" : "tr", "id" : supplier.M_id + "-adminItemRow"}).down()
    .appendElement({"tag" :"td"}).down()
    .appendElement({"tag" : "span","innerHtml" : supplier.M_id, "id" : supplier.M_id + "-transactionID"}).up()
    .appendElement({"tag" :"td", "id" : supplier.M_id + "-supplierID"})
    .appendElement({"tag" :"td"}).down()
    .appendElement({"tag" : "input", "class" : "form-control", "type" : "number", "value" : supplier.M_quantity, "id" : supplier.M_id + "-qty"}).up()
    .appendElement({"tag" : "td"}).down()
    .appendElement({"tag" :"div", "class" : "input-group"}).down()
    .appendElement({"tag":"input", "class" : "form-control", "type" : "number", "value" : supplier.M_price.M_valueOf,  "id" : supplier.M_id + "-price" })
    .appendElement({"tag":"div", "class" : "input-group-append"}).down()
    .appendElement({"tag" :"button", "class": "btn btn-outline-dark dropdown-toggle", "type" : "button", "data-toggle" : "dropdown", "aria-haspopup" : "true", "aria-expanded":"false","innerHtml" : supplier.M_price.M_currency.M_valueOf, "id" : supplier.M_id + "-currentCurrency"})
    .appendElement({"tag":"div", "class":"dropdown-menu dropDown", "id" : supplier.M_id + "-currencyTypes", "style": "transform: translate3d(137px, 0px, 0px)!important"}).down()
    .up().up().up().up()
    .appendElement({"tag":"td", "id" : supplier.M_id + "-dateSelector"})
    .appendElement({"tag":"td"}).down()
    .appendElement({"tag" : "button", "class" : "btn btn-danger btn-sm", "id" : supplier.M_id + "-deleteRowAdminBtn", "style":"display: inline-flex; height: 29px;"}).down()
    .appendElement({"tag" : "i", "class" : "fa fa-trash-o"}).up()
    
     if(moment(supplier.M_date).format('MM/DD/YYYY') != "Invalid date"){
        mainPage.adminPage.adminPackagePage.panel.getChildById(supplier.M_id + "-dateSelector").appendElement({"tag" : "span", "innerHtml":moment(supplier.M_date).format('MM/DD/YYYY'), "id" : supplier.M_id + "-date"})
    }else{
        mainPage.adminPage.adminPackagePage.panel.getChildById(supplier.M_id + "-dateSelector") .appendElement({"tag" : "input", "type" : "date", "class"  :"form-control", "id" : supplier.M_id + "-date"})
    }

    server.entityManager.getEntityByUuid(supplier.M_supplier,false,
        function(supplier,caller){
            mainPage.adminPage.adminPackagePage.panel.getChildById(caller.item_supplier.M_id + "-supplierID").appendElement({"tag" : "span", "innerHtml" : supplier.M_id + " (" + supplier.M_name + ")"})
        },function(){},{"item_supplier" : supplier})
    
    for(var i = 0; i < mainPage.adminPage.adminPackagePage.currencies.length; i++){
        mainPage.adminPage.adminPackagePage.panel.getChildById(supplier.M_id + "-currencyTypes").appendElement({"tag" : "a", "class" : "dropdown-item", "innerHtml" : mainPage.adminPage.adminPackagePage.currencies[i], "id" : supplier.M_id + "-currencyTypes-" + mainPage.adminPage.adminPackagePage.currencies[i]})
        
        //Save changes to the currency type
        mainPage.adminPage.adminPackagePage.panel.getChildById(supplier.M_id + "-currencyTypes-" + mainPage.adminPage.adminPackagePage.currencies[i]).element.onclick = function (supplier, currency, packageID){
            return function(){
                mainPage.adminPage.panel.getChildById("adminPackageSaveState").element.innerHTML = "*"
                mainPage.adminPage.adminPackagePage.panel.getChildById(supplier.M_id + "-currentCurrency").element.innerHTML = currency
                
                var newsupplier = jQuery.extend({},supplier)
                newsupplier.M_price.M_currency.M_valueOf = currency
                mainPage.adminPage.adminPackagePage.modifiedItems[packageID].set(newsupplier.M_id, newsupplier)
                
                mainPage.adminPage.adminPackagePage.panel.getChildById(packageID + "-saveBtn").element.classList.remove("disabled")
                mainPage.adminPage.adminPackagePage.panel.getChildById(packageID + "-cancelBtn").element.classList.remove("disabled")

            }
        }(supplier, mainPage.adminPage.adminPackagePage.currencies[i], pkg.M_id)
    }
    
    //Save changes to the quantity, price
    mainPage.adminPage.adminPackagePage.panel.getChildById(supplier.M_id  + "-adminItemRow").element.onchange = function(supplier, packageID){
        return function(){
            mainPage.adminPage.panel.getChildById("adminPackageSaveState").element.innerHTML = "*"
            var qty = mainPage.adminPage.adminPackagePage.panel.getChildById(supplier.M_id+"-qty").element.value
            var price = mainPage.adminPage.adminPackagePage.panel.getChildById(supplier.M_id + "-price").element.value
            if(mainPage.adminPage.adminPackagePage.panel.getChildById(supplier.M_id + "-date").element.value != undefined){
                var date = mainPage.adminPage.adminPackagePage.panel.getChildById(supplier.M_id + "-date").element.value
            }else{
                var date = mainPage.adminPage.adminPackagePage.panel.getChildById(supplier.M_id + "-date").element.innerHTML
            }
            
            var newsupplier = jQuery.extend({}, supplier)
            newsupplier.M_date = date
            newsupplier.M_price.M_valueOf = price
            newsupplier.M_quantity= qty
            mainPage.adminPage.adminPackagePage.modifiedItems[packageID].set(newsupplier.M_id, newsupplier)

            mainPage.adminPage.adminPackagePage.panel.getChildById(packageID + "-saveBtn").element.classList.remove("disabled")
            mainPage.adminPage.adminPackagePage.panel.getChildById(packageID + "-cancelBtn").element.classList.remove("disabled")
            
        }
        
        
    }(supplier, pkg.M_id)
        
    
    mainPage.adminPage.adminPackagePage.panel.getChildById(supplier.M_id + "-deleteRowAdminBtn").element.onclick = function(supplier, packageID){
        return function(){
            mainPage.adminPage.panel.getChildById("adminPackageSaveState").element.innerHTML = "*"
            mainPage.adminPage.adminPackagePage.removedItems[packageID].set(supplier.M_id, supplier)
            var row = mainPage.adminPage.adminPackagePage.panel.getChildById(supplier.M_id + "-adminItemRow")
            row.delete()
            mainPage.adminPage.adminPackagePage.panel.getChildById(packageID + "-saveBtn").element.classList.remove("disabled")
            mainPage.adminPage.panel.getChildById(packageID + "-cancelBtn").element.classList.remove("disabled")
        }
    }(supplier, pkg.M_id)
    
    
}
