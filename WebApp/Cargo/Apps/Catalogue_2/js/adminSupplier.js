var AdminSupplierPage = function(panel){
    /** The tab content panel **/
    this.panel = panel
    
    /** The panel sections. **/
    this.panel.appendElement({"tag" : "div", "class" : "row", "style" : "height:85vh;margin:0;"}).down()
        .appendElement({"tag" : "div", "class" : "col-md-3 bg-dark", "id" : "suppliersAdminNavigation"}).down()
        .appendElement({"tag" : "div", "class" : "input-group mb-3", "style":"padding:15px;"}).down()
        .appendElement({"tag" : "span", "class" : "input-group-prepend"}).down()
        .appendElement({"tag":"span", "class":"input-group-text"}).down()
        .appendElement({"tag" : "i", "class" : "fa fa-search text-dark"}).up().up()
        .appendElement({"tag" : "input", "type" : "text", "class" : "form-control", "placeholder" : "Filtrer", "id" : "supplierFilterKeywords"}).up()
        .appendElement({"tag" : "div", "class" : "list-group","role" :"tablist", "id" : "suppliersFilteredList"})
        .appendElement({"tag" : "div", "style" : "position:absolute;bottom:0; right:0;margin:30px;"}).down()
        .appendElement({"tag" : "button", "class" : "btn btn-success", "id" : "addSupplierButton", "style":"display: inline-flex; height: 36px;"}).down()
        .appendElement({"tag" : "i", "class" : "fa fa-plus"}).up().up().up()
        .appendElement({"tag" : "div", "class" : "col-md-9 bg-light"}).down()
         .appendElement({"tag" : "div", "class" : "tab-content", "id" : "suppliersAdminControl"}).up().up().up()
        
    this.currencies = []
    
    this.modifiedItems = {}
  
    this.removedItems = {}
    
    this.panel.getChildById("addSupplierButton").element.onclick = function(){
      
        mainPage.adminPage.adminSupplierPage.panel.getChildById("suppliersFilteredList").appendElement({"tag":"a", "class" : "list-group-item list-group-item-action active show","id" : "newSupplier-selector","data-toggle" : "tab","href" : "#"+ "newSupplier-control","role":"tab", "innerHtml": "Nouveau fournisseur","aria-controls": "newSupplier-control"})
         mainPage.adminPage.adminSupplierPage.panel.getChildById("suppliersAdminControl").appendElement({"tag" : "div", "class" : "tab-pane active show", "id" : "newSupplier-control", "role" : "tabpanel", "aria-labelledby" : "newSupplier-selector", "style" : "padding:15px;"}).down()
        .appendElement({"tag" : "div", "class" : "input-group mb-3"}).down()
        .appendElement({"tag" : "div", "class" : "input-group-prepend"}).down()
        .appendElement({"tag" : "span", "class" : "input-group-text", "innerHtml" : "ID"}).up()
        .appendElement({"tag" : "input", "class" : "form-control",  "type" : "text", "id" : "newSupplier-supplierID"}).up()
        .appendElement({"tag" : "div", "class" : "input-group mb-3"}).down()
        .appendElement({"tag" : "div", "class" : "input-group-prepend"}).down()
        .appendElement({"tag" : "span", "class" : "input-group-text", "innerHtml" : "Nom"}).up()
        .appendElement({"tag" : "input", "class" : "form-control", "type" : "text", "id" : "newSupplier-name"}).up()
        .appendElement({"tag" : "button", "class" : "btn btn-primary mr-3", "innerHtml" : "Enregistrer", "id" : "newSupplier-saveBtn"})
        
        mainPage.adminPage.adminSupplierPage.panel.getChildById("newSupplier-saveBtn").element.onclick = function(){
            document.getElementById("waitingDiv").style.display = ""
            var id = mainPage.adminPage.adminSupplierPage.panel.getChildById("newSupplier-supplierID").element.value
            var name = mainPage.adminPage.adminSupplierPage.panel.getChildById("newSupplier-name").element.value
            var newSupplier = new CatalogSchema.SupplierType
            newSupplier.M_id = id
            newSupplier.M_name = name
            console.log(newSupplier)
            server.entityManager.createEntity(catalog.UUID, "M_suppliers", newSupplier,
                function(success,caller){
                    document.getElementById("waitingDiv").style.display = "none"
                    mainPage.adminPage.adminSupplierPage.panel.getChildById("suppliersFilteredList").removeAllChilds()
                    mainPage.adminPage.adminSupplierPage.panel.getChildById("suppliersAdminControl").removeAllChilds()
                    mainPage.adminPage.adminSupplierPage.panel.getChildById("suppliersFilteredList").appendElement({"tag":"a", "class" : "list-group-item list-group-item-action","id" : success.M_id + "-selector","data-toggle" : "tab","href" : "#"+ success.M_id + "-control","role":"tab", "innerHtml": success.M_name,"aria-controls":  success.M_id + "-control"})
                    mainPage.adminPage.adminSupplierPage.loadAdminControl(success)
                },function(){})
            
        }
    }
    
    
    
    this.panel.getChildById("supplierFilterKeywords").element.addEventListener('change', function(){
        results = mainPage.adminPage.adminSupplierPage.getSuppliersFromKeyword(this.value.toUpperCase())
    })
    return this
}


AdminSupplierPage.prototype.getSuppliersFromKeyword = function(keyword){
     server.entityManager.getEntityPrototype("CatalogSchema.CurrencyType", "CatalogSchema", 
        function(success,caller){
            mainPage.adminPage.adminSupplierPage.currencies = []
            for(var i = 0; i < success.Restrictions.length; i++){
                
                mainPage.adminPage.adminSupplierPage.currencies.push(success.Restrictions[i].Value)
            }

        }, function(){},{})
    
    
    var q = new EntityQuery()
    q.TypeName = "CatalogSchema.SupplierType"
    q.Fields = ["M_name","M_id"]
    q.Query = 'CatalogSchema.SupplierType.M_name=="'+ keyword +'" '
    
    server.entityManager.getEntities("CatalogSchema.SupplierType", "CatalogSchema", q, 0,-1,[],true,true, 
        function(index,total, caller){
            
        },
        function(suppliers,caller){
     
            mainPage.adminPage.adminSupplierPage.panel.getChildById("suppliersFilteredList").removeAllChilds()
            mainPage.adminPage.adminSupplierPage.panel.getChildById("suppliersAdminControl").removeAllChilds()
           
            
           
            for(var i = 0; i < suppliers.length; i++){
                mainPage.adminPage.adminSupplierPage.panel.getChildById("suppliersFilteredList").appendElement({"tag":"a", "class" : "list-group-item list-group-item-action","id" : suppliers[i].M_id + "-selector","data-toggle" : "tab","href" : "#"+ suppliers[i].M_id + "-control","role":"tab", "innerHtml": suppliers[i].M_name,"aria-controls":  suppliers[i].M_id + "-control"})
                mainPage.adminPage.adminSupplierPage.loadAdminControl(suppliers[i])
               
                
            }
            
        },
        function(){
        },{})
    
   
}


AdminSupplierPage.prototype.loadAdminControl = function(supplier){
    mainPage.adminPage.panel.getChildById("adminSupplierSaveState").element.innerHTML = ""
    mainPage.adminPage.adminSupplierPage.modifiedItems[supplier.M_id] = new Map()
    mainPage.adminPage.adminSupplierPage.removedItems[supplier.M_id] = new Map()
    if(mainPage.adminPage.adminSupplierPage.panel.getChildById(supplier.M_id + "-control") != undefined){
        mainPage.adminPage.adminSupplierPage.panel.getChildById(supplier.M_id + "-control").delete()
    }
    mainPage.adminPage.adminSupplierPage.panel.getChildById("suppliersAdminControl").appendElement({"tag" : "div", "class" : "tab-pane", "id" : supplier.M_id + "-control", "role" : "tabpanel", "aria-labelledby" : supplier.M_id + "-selector", "style" : "padding:15px;"}).down()
    .appendElement({"tag" : "div", "class" : "input-group mb-3"}).down()
    .appendElement({"tag" : "div", "class" : "input-group-prepend"}).down()
    .appendElement({"tag" : "span", "class" : "input-group-text", "innerHtml" : "ID"}).up()
    .appendElement({"tag" : "span", "class" : "form-control",  "innerHtml" : supplier.M_id}).up()
    .appendElement({"tag" : "div", "class" : "input-group mb-3"}).down()
    .appendElement({"tag" : "div", "class" : "input-group-prepend"}).down()
    .appendElement({"tag" : "span", "class" : "input-group-text", "innerHtml" : "Nom"}).up()
    .appendElement({"tag" : "input", "class" : "form-control", "type" : "text", "value" : supplier.M_name, "id" : supplier.M_id + "-name"}).up()
    .appendElement({"tag" : "div", "class" : "row"}).down()
    .appendElement({"tag":"label", "for" : supplier.M_id + "-itemsList", "innerHtml" : "Produits","class" : "col-3 d-flex align-items-center"})
    .appendElement({"tag": "input", "class" : "form-control col-5 m-1", "id" : supplier.M_id + "-filterItemSuppliers","type" :"text", "placeholder" : "Filtrer.."}).up()
    .appendElement({"tag" : "div"}).down()
    .appendElement({"tag":"table", "class":"table table-hover", "style" : "display: block;height: 32em;overflow: auto;"}).down()
    .appendElement({"tag":"thead"}).down()
    .appendElement({"tag" : "tr"}).down()
    .appendElement({"tag" : "th", "innerHtml" : "ID transaction", "scope" : "col"})
    .appendElement({"tag" : "th", "innerHtml" : "ID Paquet", "scope" : "col"})
    .appendElement({"tag" : "th", "innerHtml" : "Quantité", "scope" : "col"})
    .appendElement({"tag" : "th", "innerHtml" : "Prix", "scope" : "col"})
    .appendElement({"tag" : "th", "innerHtml" : "Date", "scope" : "col"})
    .appendElement({"tag" : "th", "scope" : "col"}).down()
    .appendElement({"tag" : "button", "class" : "btn btn-success btn-sm", "id" : supplier.M_id +"-addItemSupplierButton", "style":"display: inline-flex; height: 29px;"}).down()
    .appendElement({"tag" : "i", "class" : "fa fa-plus"}).up().up()
    .up().up()
    .appendElement({"tag" : "tbody", "id" : supplier.M_id + "-itemsList"})
    .up().up()
    .appendElement({"tag" : "div", "class" : "d-flex justify-content-center"}).down()
    .appendElement({"tag" : "button", "class" : "btn btn-primary disabled mr-3", "innerHtml" : "Enregistrer", "id" : supplier.M_id + "-saveBtn"})
    .appendElement({"tag" : "button", "class" : "btn disabled mr-3", "innerHtml" : "Annuler les modifications", "id" : supplier.M_id + "-cancelBtn"})
    .appendElement({"tag" : "button", "class" : "btn btn-danger", "innerHtml" : "Supprimer l'entité", "id" : supplier.M_id + "-deleteBtn", "data-toggle" : "modal", "data-target" : "#" + supplier.M_id + "-modal"})
    .appendElement({"tag" : "div", "class" : "modal fade", "id"  : supplier.M_id + "-modal", "tabindex" : "-1", "role"  :"dialog", "aria-labelledby" : supplier.M_id + "-modal", "aria-hidden" : "true"}).down()
    .appendElement({"tag" : "div", "class" : "modal-dialog", "role" : "document"}).down()
    .appendElement({"tag" : "div", "class" : "modal-content"}).down()
    .appendElement({"tag" : "div", "class" : "modal-header"}).down()
    .appendElement({"tag" : "h5", "class" : "modal-title", "innerHtml" : "Supprimer le fournisseur"})
    .appendElement({"tag" : "button", "class" : "close", "data-dismiss" : "modal", "aria-label" : "Close"}).down()
    .appendElement({"tag" : "span", "aria-hidden" : "true", "innerHtml" : "&times;"}).up().up()
    .appendElement({"tag" : "div", "class" : "modal-body", "innerHtml" : "Êtes-vous certain de vouloir supprimer ce fournisseur? Cela supprimera aussi toutes les transactions associées à celui-ci."})
    .appendElement({"tag" : "div", "class" : "modal-footer"}).down()
    .appendElement({"tag" : "button", "class" : "btn btn-secondary", "data-dismiss" : "modal", "innerHtml" : "Fermer", "id" : supplier.M_id + "-closeDelete"})
    .appendElement({"tag" : "button", "type" : "button", "class" : "btn btn-danger", "id" : supplier.M_id + "-deleteConfirmBtn", "innerHtml" : "Confirmer la suppression"})
    mainPage.adminPage.adminSupplierPage.panel.getChildById(supplier.M_id + "-name").element.onkeyup = function(supplier){
       return function(){
           mainPage.adminPage.panel.getChildById("adminSupplierSaveState").element.innerHTML = "*"
            mainPage.adminPage.adminSupplierPage.panel.getChildById(supplier.M_id + "-saveBtn").element.classList.remove("disabled")
            mainPage.adminPage.panel.getChildById(supplier.M_id + "-cancelBtn").element.classList.remove("disabled")
           var newName = mainPage.adminPage.adminSupplierPage.panel.getChildById(supplier.M_id + "-name").element.value
           var newSupplier = supplier
           newSupplier.M_name = newName
           
           mainPage.adminPage.adminSupplierPage.modifiedItems[supplier.M_id].set(supplier.M_id, newSupplier)
       }
    }(supplier)
    
    mainPage.adminPage.adminSupplierPage.panel.getChildById(supplier.M_id + "-deleteConfirmBtn").element.onclick = function(supplier){
        return function(){
             document.getElementById("waitingDiv").style.display = ""
            var q = new EntityQuery()
            q.TypeName = "CatalogSchema.ItemSupplierType"
            q.Fields = ["M_supplier"]
            q.Query = 'CatalogSchema.ItemSupplierType.M_supplier=="'+ supplier.UUID +'" '
            
            server.entityManager.getEntities("CatalogSchema.ItemSupplierType", "CatalogSchema", q, 0,-1,[],true,true, 
                function(index,total,caller){},
                function(item_suppliers, caller){
                    mainPage.adminPage.adminSupplierPage.panel.getChildById(caller.supplier.M_id + "-closeDelete").element.click()
                   
                
                    for(var i = 0; i < item_suppliers.length; i++){
                        server.entityManager.removeEntity(item_suppliers[i].UUID,
                            function(success,caller){
                                console.log(success)
                                
                            },function(){},{"length" : item_suppliers.length, "supplier" : caller.supplier})
                        
                    }
                        
                    
                    
                    
                },function(){},{"supplier" : supplier})
                
            server.entityManager.removeEntity(supplier.UUID,
                function(success,caller){
                    document.getElementById("waitingDiv").style.display = "none"
                    mainPage.adminPage.adminSupplierPage.panel.getChildById(caller.supplier.M_id + "-control").delete()
                    mainPage.adminPage.adminSupplierPage.panel.getChildById(caller.supplier.M_id + "-selector").delete()
                    console.log(success)
                },function(){},{"supplier" : supplier})
        }
    }(supplier)
    
    mainPage.adminPage.adminSupplierPage.panel.getChildById(supplier.M_id + "-saveBtn").element.onclick = function(supplier){
        return function(){
            
            if(mainPage.adminPage.adminSupplierPage.modifiedItems[supplier.M_id].size > 0){
                 var i = 0
                 document.getElementById("waitingDiv").style.display = ""

                for(var object of mainPage.adminPage.adminSupplierPage.modifiedItems[supplier.M_id].values()){
                
                    server.entityManager.saveEntity(object,
                        function(result, caller){
                           
                            caller.i++
                            if(caller.i == caller.length){
                                if(mainPage.adminPage.adminSupplierPage.removedItems[caller.supplier.M_id].size > 0){
                                    var j = 0
                                    for(var removeObject of mainPage.adminPage.adminSupplierPage.removedItems[caller.supplier.M_id].values()){
                                     server.entityManager.removeEntity(removeObject.UUID,
                                        function(result,caller){
                                            caller.j++
                                            if(caller.j == caller.length){
                                                
                                                server.entityManager.getEntityByUuid(caller.supplier.UUID, true,
                                                    function(supplier,caller){
                                                        mainPage.adminPage.panel.getChildById("adminSupplierSaveState").element.innerHTML = ""
                                                        document.getElementById("waitingDiv").style.display = "none"
                                                        mainPage.adminPage.adminSupplierPage.panel.getChildById(supplier.M_id + "-selector").element.innerHTML = supplier.M_name
                                                        mainPage.adminPage.adminSupplierPage.loadAdminControl(supplier)
                                                    mainPage.adminPage.adminSupplierPage.panel.getChildById(supplier.M_id + "-saveBtn").element.classList.add("disabled")
                                                    mainPage.adminPage.adminSupplierPage.panel.getChildById(supplier.M_id + "-cancelBtn").element.classList.add("disabled")
                                                    
                                                    mainPage.adminPage.adminSupplierPage.panel.getChildById(supplier.M_id + "-control").element.classList.add("active")
                                                     mainPage.adminPage.adminSupplierPage.panel.getChildById(supplier.M_id + "-control").element.classList.add("show")
                                                    },function(){},{})
                                                
                                            }
                                         
                                        },function(){},{"length" : mainPage.adminPage.adminSupplierPage.removedItems[supplier.M_id].size, "j" : j, "supplier" : caller.supplier})
                       
                                     }
                                }else{
                                   server.entityManager.getEntityByUuid(caller.supplier.UUID, true,
                                    function(supplier,caller){
                                        mainPage.adminPage.panel.getChildById("adminSupplierSaveState").element.innerHTML = ""
                                        document.getElementById("waitingDiv").style.display = "none"
                                        mainPage.adminPage.adminSupplierPage.panel.getChildById(supplier.M_id + "-selector").element.innerHTML = supplier.M_name
                                        mainPage.adminPage.adminSupplierPage.loadAdminControl(supplier)
                                    mainPage.adminPage.adminSupplierPage.panel.getChildById(supplier.M_id + "-saveBtn").element.classList.add("disabled")
                                    mainPage.adminPage.adminSupplierPage.panel.getChildById(supplier.M_id + "-cancelBtn").element.classList.add("disabled")
                                    
                                    mainPage.adminPage.adminSupplierPage.panel.getChildById(supplier.M_id + "-control").element.classList.add("active")
                                     mainPage.adminPage.adminSupplierPage.panel.getChildById(supplier.M_id + "-control").element.classList.add("show")
                                    },function(){},{})
                                }
                               
                                
                            }
                        }, function () {}, {"i" : i, "length" : mainPage.adminPage.adminSupplierPage.modifiedItems[supplier.M_id].size, "supplier" : supplier})
                }
            }else{
                if(mainPage.adminPage.adminSupplierPage.removedItems[supplier.M_id].size > 0){
                     document.getElementById("waitingDiv").style.display = ""
                    var j = 0
                    for(var removeObject of mainPage.adminPage.adminSupplierPage.removedItems[supplier.M_id].values()){
                     server.entityManager.removeEntity(removeObject.UUID,
                        function(result,caller){
                            caller.j++
                            if(caller.j == caller.length){
                                server.entityManager.getEntityByUuid(caller.supplier.UUID, true,
                                function(supplier,caller){
                                    mainPage.adminPage.panel.getChildById("adminSupplierSaveState").element.innerHTML = ""
                                    document.getElementById("waitingDiv").style.display = "none"
                                    mainPage.adminPage.adminSupplierPage.panel.getChildById(supplier.M_id + "-selector").element.innerHTML = supplier.M_name
                                    mainPage.adminPage.adminSupplierPage.loadAdminControl(supplier)
                                mainPage.adminPage.adminSupplierPage.panel.getChildById(supplier.M_id + "-saveBtn").element.classList.add("disabled")
                                mainPage.adminPage.adminSupplierPage.panel.getChildById(supplier.M_id + "-cancelBtn").element.classList.add("disabled")
                                
                                mainPage.adminPage.adminSupplierPage.panel.getChildById(supplier.M_id + "-control").element.classList.add("active")
                                 mainPage.adminPage.adminSupplierPage.panel.getChildById(supplier.M_id + "-control").element.classList.add("show")
                                },function(){},{})
                            }
                         
                        },function(){},{"length" : mainPage.adminPage.adminSupplierPage.removedItems[supplier.M_id].size, "j" : j, "supplier" : supplier})
                    }
                }
                
            }
            
            
            
           
        
        }
    }(supplier)
    
    mainPage.adminPage.adminSupplierPage.panel.getChildById(supplier.M_id + "-cancelBtn").element.onclick = function(supplier){
        return function(){
           
            mainPage.adminPage.adminSupplierPage.panel.getChildById(supplier.M_id + "-saveBtn").element.classList.add("disabled")
            mainPage.adminPage.adminSupplierPage.panel.getChildById(supplier.M_id + "-cancelBtn").element.classList.add("disabled")
            server.entityManager.getEntityByUuid(supplier.UUID, false,
                function(success, caller){
                    mainPage.adminPage.adminSupplierPage.loadAdminControl(supplier)
                },function(){},{})
            
            mainPage.adminPage.adminSupplierPage.panel.getChildById(supplier.M_id + "-control").element.classList.add("active")
             mainPage.adminPage.adminSupplierPage.panel.getChildById(supplier.M_id + "-control").element.classList.add("show")
        }
    }(supplier)
    
    
    mainPage.adminPage.adminSupplierPage.panel.getChildById(supplier.M_id + "-filterItemSuppliers").element.onkeyup = function(supplierID){
        return function(){
            var input, filter, table, tr, td, i;
            input = document.getElementById(supplierID+"-filterItemSuppliers");
            filter = input.value.toUpperCase();
            table = document.getElementById(supplierID + "-itemsList");
            tr = table.getElementsByTagName("tr");
        
            // Loop through all table rows, and hide those who don't match the search query
            for (i = 0; i < tr.length; i++) {
                th = tr[i].getElementsByTagName("th")[0];
                td = tr[i].getElementsByTagName("td")[0];
                if (td && th) {
                    if (td.innerHTML.toUpperCase().indexOf(filter) > -1 || th.innerHTML.toUpperCase().indexOf(filter) > -1 ) {
                        tr[i].style.display = "";
                    } else {
                        tr[i].style.display = "none";
                    }
                } 
                
            }
       }
    }(supplier.M_id)
    
    mainPage.adminPage.adminSupplierPage.panel.getChildById(supplier.M_id + "-addItemSupplierButton").element.onclick = function(supplier){
        return function(){
            var date = new Date()
            mainPage.adminPage.adminSupplierPage.panel.getChildById(supplier.M_id + "-itemsList").prependElement({"tag" : "tr", "id" : date + "-row"}).down()
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
            
            for(var i = 0; i < mainPage.adminPage.adminSupplierPage.currencies.length; i++){
                mainPage.adminPage.adminSupplierPage.panel.getChildById(date + "-currencyTypes").appendElement({"tag" : "a", "class" : "dropdown-item", "innerHtml" : mainPage.adminPage.adminSupplierPage.currencies[i], "id" : date + "-currencyTypes-" + mainPage.adminPage.adminSupplierPage.currencies[i]})
                
                mainPage.adminPage.adminSupplierPage.panel.getChildById(date + "-currencyTypes-" + mainPage.adminPage.adminSupplierPage.currencies[i]).element.onclick = function (divID, currency){
                    return function(){
                        mainPage.adminPage.adminSupplierPage.panel.getChildById(divID + "-currentCurrency").element.innerHTML = currency
                    }
                }(date, mainPage.adminPage.adminSupplierPage.currencies[i])
            }
            
            autocomplete("CatalogSchema.PackageType", mainPage.adminPage.adminSupplierPage.panel.getChildById(date + "-id"))
            
            mainPage.adminPage.adminSupplierPage.panel.getChildById(date + "-id").element.oncomplete = function(uuid){
                mainPage.adminPage.adminSupplierPage.panel[date+"-id"] = uuid
            }
                    
                    
                    
            mainPage.adminPage.adminSupplierPage.panel.getChildById(date+ "-confirmBtn").element.onclick = function (date, supplier) {
           return function(){
                document.getElementById("waitingDiv").style.display = ""
                var newPackage = mainPage.adminPage.adminSupplierPage.panel[date+"-id"]
                var newItem = new CatalogSchema.ItemSupplierType()
                newItem.M_package = newPackage
                newItem.M_quantity = mainPage.adminPage.adminSupplierPage.panel.getChildById(date + "-qty").element.value
                newItem.M_id = mainPage.adminPage.adminSupplierPage.panel.getChildById(date + "-transactionID").element.value
                newItem.M_price = new CatalogSchema.PriceType()
                newItem.M_price.M_valueOf = mainPage.adminPage.adminSupplierPage.panel.getChildById(date + "-price").element.value
                newItem.M_supplier = supplier.UUID
                if(mainPage.adminPage.adminSupplierPage.panel.getChildById(date + "-currentCurrency").element.innerHTML != "Currency"){
                    newItem.M_price.M_currency.M_valueOf = mainPage.adminPage.adminSupplierPage.panel.getChildById(date + "-currentCurrency").element.innerHTML
                }
                
                newItem.M_date = mainPage.adminPage.adminSupplierPage.panel.getChildById(date + "-date").element.value
                server.entityManager.createEntity(supplier, "M_items", newItem, 
                    function(success, caller){
     
                         caller.supplier.M_items.push(success.UUID)
                  
                         server.entityManager.saveEntity(caller.supplier, 
                            function(result,caller){
                                
                                mainPage.adminPage.adminSupplierPage.panel.getChildById(date + "-row").delete()
                                
                                document.getElementById("waitingDiv").style.display = "none"
                                server.entityManager.getEntityByUuid(caller.package,false,
                                    function(pkg,caller){
                                        mainPage.adminPage.adminSupplierPage.appendItemSupplier(caller.newLine, caller.supplier.M_id, pkg)
                                    },function(){}, {"newLine" : caller.newLine, "supplier" : caller.supplier})
                                
                                console.log(result)
                            },function(){},{"supplier" : caller.supplier, "package" : caller.package, "newLine" : success})
                        
                    },function(){},{"supplier" : supplier, "package" : newPackage })
                    
               
                    
                    
                    
                    
                
                           
               
               }
            }(date, supplier)
        }
    }(supplier)
    
    for(var i =0; i < supplier.M_items.length; i++){
        if(supplier.M_items[i].M_id == undefined){
            if(supplier.M_items[i] != null){
             server.entityManager.getEntityByUuid(supplier.M_items[i],false,
            function(item_supplier ,caller){
           
    
            server.entityManager.getEntityByUuid(item_supplier.M_package, false,
                function(item_package, caller){
                    mainPage.adminPage.adminSupplierPage.appendItemSupplier(caller.itemsupplier, caller.supplierID, item_package) 
                   
                  
                    
                },function(){}, {"itemsupplier" : item_supplier, "supplierID" : caller.supplierID})
            },
            function(){
            },{"supplierID" : supplier.M_id})
            }
        }else{
            server.entityManager.getEntityByUuid(supplier.M_items[i].M_package, false,
                function(item_package, caller){
                    mainPage.adminPage.adminSupplierPage.appendItemSupplier(caller.itemsupplier, caller.supplierID, item_package) 
                   
                  
                    
                },function(){}, {"itemsupplier" : supplier.M_items[i], "supplierID" : supplier.M_id})
        }
        
       
    }
    
}

AdminSupplierPage.prototype.appendItemSupplier = function(item_supplier, supplierID, item_package){
    mainPage.adminPage.adminSupplierPage.panel.getChildById(supplierID+"-itemsList").appendElement({"tag" : "tr", "id" : item_supplier.M_id + "-adminItemRow"}).down()
    .appendElement({"tag" :"td","innerHtml" : item_supplier.M_id, "id" : item_supplier.M_id + "-transactionID"})
    .appendElement({"tag" : "th", "scope" : "row" ,"innerHtml" : item_package.M_id, "id" : item_supplier.M_id + "-id", "class" : "tabLink"})
    .appendElement({"tag" :"td"}).down()
    .appendElement({"tag" : "input", "class" : "form-control", "type" : "number", "value" : item_supplier.M_quantity, "id" : item_supplier.M_id + "-qty"}).up()
    .appendElement({"tag" : "td"}).down()
    .appendElement({"tag" :"div", "class" : "input-group"}).down()
    .appendElement({"tag":"input", "class" : "form-control", "type" : "number", "value" : item_supplier.M_price.M_valueOf,  "id" : item_supplier.M_id + "-price" })
    .appendElement({"tag":"div", "class" : "input-group-append"}).down()
    .appendElement({"tag" :"button", "class": "btn btn-outline-dark dropdown-toggle", "type" : "button", "data-toggle" : "dropdown", "aria-haspopup" : "true", "aria-expanded":"false","innerHtml" : item_supplier.M_price.M_currency.M_valueOf, "id" : item_supplier.M_id + "-currentCurrency"})
    .appendElement({"tag":"div", "class":"dropdown-menu dropDown", "id" : item_supplier.M_id + "-currencyTypes", "style": "transform: translate3d(137px, 0px, 0px)!important"}).down()
    .up().up().up().up()
    .appendElement({"tag":"td", "id" : item_supplier.M_id + "-dateSelector"})
    .appendElement({"tag":"td"}).down()
    .appendElement({"tag" : "button", "class" : "btn btn-danger btn-sm", "id" : item_supplier.M_id + "-deleteRowAdminBtn", "style":"display: inline-flex; height: 29px;"}).down()
    .appendElement({"tag" : "i", "class" : "fa fa-trash-o"}).up()
    
     if(moment(item_supplier.M_date).format('MM/DD/YYYY') != "Invalid date"){
        mainPage.adminPage.adminSupplierPage.panel.getChildById(item_supplier.M_id + "-dateSelector").appendElement({"tag" : "span", "innerHtml":moment(item_supplier.M_date).format('MM/DD/YYYY'), "id" : item_supplier.M_id + "-date"})
    }else{
        mainPage.adminPage.adminSupplierPage.panel.getChildById(item_supplier.M_id + "-dateSelector") .appendElement({"tag" : "input", "type" : "date", "class"  :"form-control", "id" : item_supplier.M_id + "-date"})
    }


    mainPage.adminPage.adminSupplierPage.panel.getChildById(item_supplier.M_id + "-id").element.onclick = function(item_package){
        return function(){
            mainPage.packageDisplayPage.displayTabItem(item_package)
        }
    }(item_package)
    
    for(var i = 0; i < mainPage.adminPage.adminSupplierPage.currencies.length; i++){
        mainPage.adminPage.adminSupplierPage.panel.getChildById(item_supplier.M_id + "-currencyTypes").appendElement({"tag" : "a", "class" : "dropdown-item", "innerHtml" : mainPage.adminPage.adminSupplierPage.currencies[i], "id" : item_supplier.M_id + "-currencyTypes-" + mainPage.adminPage.adminSupplierPage.currencies[i]})
        
        //Save changes to the currency type
        mainPage.adminPage.adminSupplierPage.panel.getChildById(item_supplier.M_id + "-currencyTypes-" + mainPage.adminPage.adminSupplierPage.currencies[i]).element.onclick = function (item_supplier, currency, supplierID){
            return function(){
                mainPage.adminPage.panel.getChildById("adminSupplierSaveState").element.innerHTML = "*"
                mainPage.adminPage.adminSupplierPage.panel.getChildById(item_supplier.M_id + "-currentCurrency").element.innerHTML = currency
                
                var newitem_supplier = jQuery.extend({},item_supplier)
                newitem_supplier.M_price.M_currency.M_valueOf = currency
                mainPage.adminPage.adminSupplierPage.modifiedItems[supplierID].set(newitem_supplier.M_id, newitem_supplier)
                
                mainPage.adminPage.adminSupplierPage.panel.getChildById(supplierID + "-saveBtn").element.classList.remove("disabled")
                mainPage.adminPage.adminSupplierPage.panel.getChildById(supplierID + "-cancelBtn").element.classList.remove("disabled")

            }
        }(item_supplier, mainPage.adminPage.adminSupplierPage.currencies[i], supplierID)
    }
    
    //Save changes to the quantity, price
    mainPage.adminPage.adminSupplierPage.panel.getChildById(item_supplier.M_id  + "-adminItemRow").element.onchange = function(item_supplier, supplierID){
            return function(){
                mainPage.adminPage.panel.getChildById("adminSupplierSaveState").element.innerHTML = "*"
                var qty = mainPage.adminPage.adminSupplierPage.panel.getChildById(item_supplier.M_id+"-qty").element.value
                var price = mainPage.adminPage.adminSupplierPage.panel.getChildById(item_supplier.M_id + "-price").element.value
                if(mainPage.adminPage.adminSupplierPage.panel.getChildById(item_supplier.M_id + "-date").element.value != null){
                    var date = mainPage.adminPage.adminSupplierPage.panel.getChildById(item_supplier.M_id + "-date").element.value
                }else{
                    var date = mainPage.adminPage.adminSupplierPage.panel.getChildById(item_supplier.M_id + "-date").element.innerHTML
                }
                
                var newitem_supplier = jQuery.extend({}, item_supplier)
                newitem_supplier.M_date = date
                newitem_supplier.M_price.M_valueOf = price
                newitem_supplier.M_quantity= qty
                mainPage.adminPage.adminSupplierPage.modifiedItems[supplierID].set(newitem_supplier.M_id, newitem_supplier)

                mainPage.adminPage.adminSupplierPage.panel.getChildById(supplierID + "-saveBtn").element.classList.remove("disabled")
                mainPage.adminPage.adminSupplierPage.panel.getChildById(supplierID + "-cancelBtn").element.classList.remove("disabled")
                
            }
            
            
        }(item_supplier, supplierID)
        
    
    mainPage.adminPage.adminSupplierPage.panel.getChildById(item_supplier.M_id + "-deleteRowAdminBtn").element.onclick = function(item_supplier, supplierID){
        return function(){
            mainPage.adminPage.panel.getChildById("adminSupplierSaveState").element.innerHTML = "*"
            mainPage.adminPage.adminSupplierPage.removedItems[supplierID].set(item_supplier.M_id, item_supplier)
            var row = mainPage.adminPage.adminSupplierPage.panel.getChildById(item_supplier.M_id + "-adminItemRow")
            row.delete()
            mainPage.adminPage.adminSupplierPage.panel.getChildById(supplierID + "-saveBtn").element.classList.remove("disabled")
            mainPage.adminPage.panel.getChildById(supplierID + "-cancelBtn").element.classList.remove("disabled")
        }
    }(item_supplier, supplierID)
}
