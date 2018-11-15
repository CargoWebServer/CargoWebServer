var AdminCategoryPage = function(panel){
    /** The tab content panel **/
    this.panel = panel
    
    /** The panel sections. **/
    this.panel.appendElement({"tag" : "div", "class" : "row", "style" : "height:85vh;margin:0;"}).down()
        .appendElement({"tag" : "div", "class" : "col-md-3 bg-dark", "id" : "categoriesAdminNavigation"}).down()
        .appendElement({"tag" : "div", "class" : "input-group mb-3", "style":"padding:15px;"}).down()
        .appendElement({"tag" : "span", "class" : "input-group-prepend"}).down()
        .appendElement({"tag":"span", "class":"input-group-text"}).down()
        .appendElement({"tag" : "i", "class" : "fa fa-search text-dark"}).up().up()
        .appendElement({"tag" : "input", "type" : "text", "class" : "form-control", "placeholder" : "Filtrer", "id" : "categoryFilterKeywords"}).up()
        .appendElement({"tag" : "div", "class" : "list-group","role" :"tablist", "id" : "categoriesFilteredList"})
        .appendElement({"tag" : "div", "style" : "position:absolute;bottom:0; right:0;margin:30px;"}).down()
        .appendElement({"tag" : "button", "class" : "btn btn-success", "id" : "addCategoryButton"}).down()
        .appendElement({"tag" : "i", "class" : "fa fa-plus"}).up().up().up()
        .appendElement({"tag" : "div", "class" : "col-md-9 bg-light","style":"overflow-y:auto;"}).down()
         .appendElement({"tag" : "div", "class" : "tab-content", "id" : "categoriesAdminControl"}).up().up().up()
        
  
    this.removedItems = {}
    
    this.panel.getChildById("categoryFilterKeywords").element.addEventListener('change', function(){
        results = mainPage.adminPage.adminCategoryPage.getCategoriesFromKeyword(this.value)
    })
    
    
    this.panel.getChildById("addCategoryButton").element.onclick = function(){
      
        mainPage.adminPage.adminCategoryPage.panel.getChildById("categoriesFilteredList").appendElement({"tag":"a", "class" : "list-group-item list-group-item-action active show","id" : "newCategory-selector","data-toggle" : "tab","href" : "#"+ "newCategory-control","role":"tab", "innerHtml": "Nouvelle catégorie","aria-controls": "newCategory-control"})
         mainPage.adminPage.adminCategoryPage.panel.getChildById("categoriesAdminControl").appendElement({"tag" : "div", "class" : "tab-pane active show", "id" : "newCategory-control", "role" : "tabpanel", "aria-labelledby" : "newCategory-selector", "style" : "padding:15px;"}).down()
        .appendElement({"tag" : "div", "class" : "input-group mb-3"}).down()
        .appendElement({"tag" : "div", "class" : "input-group-prepend"}).down()
        .appendElement({"tag" : "span", "class" : "input-group-text", "innerHtml" : "ID"}).up()
        .appendElement({"tag" : "input", "class" : "form-control",  "type" : "text", "id" : "newCategory-categoryID"}).up()
        .appendElement({"tag" : "div", "class" : "input-group mb-3"}).down()
        .appendElement({"tag" : "div", "class" : "input-group-prepend"}).down()
        .appendElement({"tag" : "span", "class" : "input-group-text", "innerHtml" : "Nom"}).up()
        .appendElement({"tag" : "input", "class" : "form-control", "type" : "text", "id" : "newCategory-name"}).up()
 
        .appendElement({"tag" : "button", "class" : "btn btn-primary mr-3", "innerHtml" : "Enregistrer", "id" : "newCategory-saveBtn"})
        
        mainPage.adminPage.adminCategoryPage.panel.getChildById("newCategory-saveBtn").element.onclick = function(){
            document.getElementById("waitingDiv").style.display = ""
            var id = mainPage.adminPage.adminCategoryPage.panel.getChildById("newCategory-categoryID").element.value
            var name = mainPage.adminPage.adminCategoryPage.panel.getChildById("newCategory-name").element.value
            var newCategory = new CatalogSchema.CategoryType
            newCategory.M_id = id
            newCategory.M_name = name

           
          
            server.entityManager.createEntity(catalog.UUID, "M_categories", newCategory,
                function(success,caller){
                    
          
                            document.getElementById("waitingDiv").style.display = "none"
                            mainPage.adminPage.adminCategoryPage.panel.getChildById("categoriesFilteredList").removeAllChilds()
                            mainPage.adminPage.adminCategoryPage.panel.getChildById("categoriesAdminControl").removeAllChilds()
                            mainPage.adminPage.adminCategoryPage.panel.getChildById("categoriesFilteredList").appendElement({"tag":"a", "class" : "list-group-item list-group-item-action","id" : caller.newCategory.M_id + "-selector","data-toggle" : "tab","href" : "#"+ caller.newCategory.M_id + "-control","role":"tab", "innerHtml": caller.newCategory.M_name,"aria-controls":  caller.newCategory.M_id + "-control"})
                            mainPage.adminPage.adminCategoryPage.loadAdminControl(caller.newCategory)
                            console.log(success)
                        
                },function(){},{"newCategory" : newCategory})
           
        }
    }
    
  
    
    
    return this
}

AdminCategoryPage.prototype.getCategoriesFromKeyword = function(keyword){
    var q = new EntityQuery()
    q.TypeName = "CatalogSchema.CategoryType"
    q.Fields = ["M_name"]
    q.Query = 'CatalogSchema.CategoryType.M_name=="'+ keyword +'"'
    
    server.entityManager.getEntities("CatalogSchema.CategoryType", "CatalogSchema", q, 0,-1,[],true,false, 
        function(index,total, caller){
            
        },
        function(categories,caller){
     
            mainPage.adminPage.adminCategoryPage.panel.getChildById("categoriesFilteredList").removeAllChilds()
            mainPage.adminPage.adminCategoryPage.panel.getChildById("categoriesAdminControl").removeAllChilds()
           
            
           
            for(var i = 0; i < categories.length; i++){
                mainPage.adminPage.adminCategoryPage.panel.getChildById("categoriesFilteredList").appendElement({"tag":"a", "class" : "list-group-item list-group-item-action","id" : categories[i].M_id + "-selector","data-toggle" : "tab","href" : "#"+ categories[i].M_id + "-control","role":"tab", "innerHtml": categories[i].M_name,"aria-controls":  categories[i].M_id + "-control"})
                mainPage.adminPage.adminCategoryPage.loadAdminControl(categories[i])
               console.log(categories[i])
                
            }
            
        },
        function(){
        },{})
}

AdminCategoryPage.prototype.loadAdminControl = function(category){
    mainPage.adminPage.adminCategoryPage.removedItems[category.M_id] = {}
    mainPage.adminPage.adminCategoryPage.removedItems[category.M_id]["items"] = new Map()
    mainPage.adminPage.adminCategoryPage.removedItems[category.M_id]["subCategories"] = new Map()
    if(mainPage.adminPage.adminCategoryPage.panel.getChildById(category.M_id + "-control") != undefined){
        mainPage.adminPage.adminCategoryPage.panel.getChildById(category.M_id + "-control").delete()
    }
    mainPage.adminPage.adminCategoryPage.panel.getChildById("categoriesAdminControl").appendElement({"tag" : "div", "class" : "tab-pane", "id" : category.M_id + "-control", "role" : "tabpanel", "aria-labelledby" : category.M_id + "-selector", "style" : "padding:15px;"}).down()
    .appendElement({"tag" : "div", "class" : "input-group mb-3"}).down()
    .appendElement({"tag" : "div", "class" : "input-group-prepend"}).down()
    .appendElement({"tag" : "span", "class" : "input-group-text", "innerHtml" : "ID"}).up()
    .appendElement({"tag" : "span", "class" : "form-control",  "innerHtml" : category.M_id}).up()
    .appendElement({"tag" : "div", "class" : "input-group mb-3"}).down()
    .appendElement({"tag" : "div", "class" : "input-group-prepend"}).down()
    .appendElement({"tag" : "span", "class" : "input-group-text", "innerHtml" : "Nom"}).up()
    .appendElement({"tag" : "input", "class" : "form-control", "type" : "text", "value" : category.M_name, "id" : category.M_id + "-name"}).up()
    .appendElement({"tag" : "div", "class" : "row"}).down()
    .appendElement({"tag":"label", "for" : category.M_id + "-subCategoriesList", "innerHtml" : "Sous-catégories","class" : "col-3 d-flex align-items-center"})
    .appendElement({"tag": "input", "class" : "form-control col-5 m-1", "id" : category.M_id + "-filterSubCategories","type" :"text", "placeholder" : "Filtrer.."}).up()
    .appendElement({"tag" : "div"}).down()
    .appendElement({"tag":"table", "class":"table table-hover table-condensed"}).down()
    .appendElement({"tag":"thead"}).down()
    .appendElement({"tag" : "tr"}).down()
    .appendElement({"tag" : "th", "innerHtml" : "ID Catégorie", "style" : "width:30%;"})
    .appendElement({"tag" : "th", "innerHtml" : "Nom", "style":"width:50%;"})

    .appendElement({"tag" : "th", "style" : "width:20%;"}).down()
    .appendElement({"tag" : "button", "class" : "btn btn-success btn-sm", "id" : category.M_id +"-addSubCategoryButton", "style":"display: inline-flex; height: 29px;"}).down()
    .appendElement({"tag" : "i", "class" : "fa fa-plus"}).up().up()
    .up().up()
    .appendElement({"tag" : "tbody", "id" : category.M_id + "-subCategoriesList"})
    .up().up()
    .appendElement({"tag" : "div", "class" : "row"}).down()
    .appendElement({"tag":"label", "for" : category.M_id + "-itemsList", "innerHtml" : "Items","class" : "col-3 d-flex align-items-center"})
    .appendElement({"tag": "input", "class" : "form-control col-5 m-1", "id" : category.M_id + "-filterItems","type" :"text", "placeholder" : "Filtrer.."}).up()
    .appendElement({"tag" : "div"}).down()
    .appendElement({"tag":"table", "class":"table table-hover table-condensed"}).down()
    .appendElement({"tag":"thead"}).down()
    .appendElement({"tag" : "tr"}).down()
    .appendElement({"tag" : "th", "innerHtml" : "ID Item", "style" : "width:30%;"})
    .appendElement({"tag" : "th", "innerHtml" : "Nom", "style":"width:50%;"})

    .appendElement({"tag" : "th", "style" : "width:20%;"}).down()
    .appendElement({"tag" : "button", "class" : "btn btn-success btn-sm", "id" : category.M_id +"-addItemButton", "style":"display: inline-flex; height: 29px;"}).down()
    .appendElement({"tag" : "i", "class" : "fa fa-plus"}).up().up()
    .up().up()
    .appendElement({"tag" : "tbody", "id" : category.M_id + "-itemsList"})
    .up().up()
    .appendElement({"tag" : "div", "class" : "d-flex justify-content-center"}).down()
    .appendElement({"tag" : "button", "class" : "btn btn-primary disabled mr-3", "innerHtml" : "Enregistrer", "id" : category.M_id + "-saveBtn"})
    .appendElement({"tag" : "button", "class" : "btn disabled mr-3", "innerHtml" : "Annuler les modifications", "id" : category.M_id + "-cancelBtn"})
    .appendElement({"tag" : "button", "class" : "btn btn-danger", "innerHtml" : "Supprimer l'entité", "id" : category.M_id + "-deleteBtn", "data-toggle" : "modal", "data-target" : "#" + category.M_id + "-modal"})
    .appendElement({"tag" : "div", "class" : "modal fade", "id"  : category.M_id + "-modal", "tabindex" : "-1", "role"  :"dialog", "aria-labelledby" : category.M_id + "-modal", "aria-hidden" : "true"}).down()
    .appendElement({"tag" : "div", "class" : "modal-dialog", "role" : "document"}).down()
    .appendElement({"tag" : "div", "class" : "modal-content"}).down()
    .appendElement({"tag" : "div", "class" : "modal-header"}).down()
    .appendElement({"tag" : "h5", "class" : "modal-title", "innerHtml" : "Supprimer la catégorie"})
    .appendElement({"tag" : "button", "class" : "close", "data-dismiss" : "modal", "aria-label" : "Close"}).down()
    .appendElement({"tag" : "span", "aria-hidden" : "true", "innerHtml" : "&times;"}).up().up()
    .appendElement({"tag" : "div", "class" : "modal-body", "innerHtml" : "Êtes-vous certain de vouloir supprimer cette catégorie?"})
    .appendElement({"tag" : "div", "class" : "modal-footer"}).down()
    .appendElement({"tag" : "button", "class" : "btn btn-secondary", "data-dismiss" : "modal", "innerHtml" : "Fermer", "id" : category.M_id + "-closeDelete"})
    .appendElement({"tag" : "button", "type" : "button", "class" : "btn btn-danger", "id" : category.M_id + "-deleteConfirmBtn", "innerHtml" : "Confirmer la suppression"})
    
     mainPage.adminPage.adminCategoryPage.panel.getChildById(category.M_id + "-deleteConfirmBtn").element.onclick = function(category){
        return function(){
            document.getElementById("waitingDiv").style.display = ""
            mainPage.adminPage.adminCategoryPage.panel.getChildById(category.M_id + "-closeDelete").element.click()
            mainPage.adminPage.adminCategoryPage.deleteSubCategories(category.M_categories)
            
                
            
                
            server.entityManager.removeEntity(category.UUID,
            function(success,caller){
                mainPage.adminPage.adminCategoryPage.panel.getChildById(caller.category.M_id + "-control").delete()
                mainPage.adminPage.adminCategoryPage.panel.getChildById(caller.category.M_id + "-selector").delete()
                document.getElementById("waitingDiv").style.display = "none"
            }, function(){},{"category" : category})
        }
    }(category)
    
    for(var i = 0; i < category.M_categories.length; i++){
        server.entityManager.getEntityByUuid(category.M_categories[i], false,
            function(subCategory,caller){
                mainPage.adminPage.adminCategoryPage.appendSubCategory(subCategory, caller.category)
            },function(){},{"category" : category})
    }
    
    for(var i = 0; i < category.M_items.length; i++){
        server.entityManager.getEntityByUuid(category.M_items[i], false,
            function(item, caller){
                mainPage.adminPage.adminCategoryPage.appendItem(item, caller.category)
            },function(){}, {"category" : category})
    }
    
    mainPage.adminPage.adminCategoryPage.panel.getChildById(category.M_id + "-filterItems").element.onkeyup = function(categoryID){
        return function(){
            var input, filter, table, tr, td, i;
            input = document.getElementById(categoryID+"-filterItems");
            filter = input.value.toUpperCase();
            table = document.getElementById(categoryID + "-itemsList");
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
    }(category.M_id)
    
    mainPage.adminPage.adminCategoryPage.panel.getChildById(category.M_id + "-filterSubCategories").element.onkeyup = function(categoryID){
        return function(){
            var input, filter, table, tr, td, i;
            input = document.getElementById(categoryID+"-filterSubCategories");
            filter = input.value.toUpperCase();
            table = document.getElementById(categoryID + "-subCategoriesList");
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
    }(category.M_id)
    
    mainPage.adminPage.adminCategoryPage.panel.getChildById(category.M_id + "-name").element.onkeyup = function(category){
       return function(){
           mainPage.adminPage.panel.getChildById("adminCategorySaveState").element.innerHTML = "*"
            mainPage.adminPage.adminCategoryPage.panel.getChildById(category.M_id + "-saveBtn").element.classList.remove("disabled")
            mainPage.adminPage.adminCategoryPage.panel.getChildById(category.M_id + "-cancelBtn").element.classList.remove("disabled")
       }
    }(category)
    
    
    mainPage.adminPage.adminCategoryPage.panel.getChildById(category.M_id + "-addSubCategoryButton").element.onclick = function(category){
        return function(){
            var date = new Date()
            mainPage.adminPage.adminCategoryPage.panel.getChildById(category.M_id + "-subCategoriesList").prependElement({"tag" : "tr", "id" : date + "-row"}).down()
            .appendElement({"tag" : "th", "scope" : "row"}).down()
            .appendElement({"tag" : "input", "class" : "form-control", "type" : "text", "placeholder" : "ID", "id" : date  + "-subCategoryID"}).up()
            .appendElement({"tag" : "td", "scope" : "row"})
             .appendElement({"tag":"td"}).down()
            .appendElement({"tag" : "button", "class" : "btn btn-success btn-sm", "id" : date + "-confirmBtn"}).down()
            .appendElement({"tag" : "i", "class" : "fa fa-check"}).up()
            
            autocomplete("CatalogSchema.CategoryType", mainPage.adminPage.adminCategoryPage.panel.getChildById(date+"-subCategoryID"))
            
            mainPage.adminPage.adminCategoryPage.panel.getChildById(date + "-subCategoryID").element.oncomplete = function(uuid){
                mainPage.adminPage.adminCategoryPage.panel[date+"-subCategoryID"] = uuid
            }
            
             mainPage.adminPage.adminCategoryPage.panel.getChildById(date+ "-confirmBtn").element.onclick = function (date, category) {
            return function(){
                document.getElementById("waitingDiv").style.display = ""
                var newSubCategory = mainPage.adminPage.adminCategoryPage.panel[date+"-subCategoryID"]
                var newCategory = jQuery.extend({},category)
                newCategory.M_categories.push(newSubCategory)
                server.entityManager.saveEntity(newCategory,
                    function(success,caller){
                     
                        mainPage.adminPage.adminCategoryPage.panel.getChildById(date + "-row").delete()
                        document.getElementById("waitingDiv").style.display = "none"
                        server.entityManager.getEntityByUuid(caller.category,false,
                            function(newSubCategory,caller){
                                mainPage.adminPage.adminCategoryPage.appendSubCategory(newSubCategory, caller.success)
                            },function(){}, {"success" : success})
                            
                    },function(){},{"category" : newSubCategory})
               }
            }(date, category)
             
                
        }
    }(category)
    
     mainPage.adminPage.adminCategoryPage.panel.getChildById(category.M_id + "-addItemButton").element.onclick = function(category){
        return function(){
            var date = new Date()
            mainPage.adminPage.adminCategoryPage.panel.getChildById(category.M_id + "-itemsList").prependElement({"tag" : "tr", "id" : date + "-row"}).down()
            .appendElement({"tag" : "th", "scope" : "row"}).down()
            .appendElement({"tag" : "input", "class" : "form-control", "type" : "text", "placeholder" : "ID", "id" : date  + "-itemID"}).up()
            .appendElement({"tag" : "td", "scope" : "row"})
             .appendElement({"tag":"td"}).down()
            .appendElement({"tag" : "button", "class" : "btn btn-success btn-sm", "id" : date + "-confirmBtn"}).down()
            .appendElement({"tag" : "i", "class" : "fa fa-check"}).up()
            
            autocomplete("CatalogSchema.ItemType", mainPage.adminPage.adminCategoryPage.panel.getChildById(date+"-itemID"))
            
             mainPage.adminPage.adminCategoryPage.panel.getChildById(date+ "-confirmBtn").element.onclick = function (date, category) {
            return function(){
                document.getElementById("waitingDiv").style.display = ""
                var pkgName = mainPage.adminPage.adminCategoryPage.panel.getChildById(date+"-itemID").element.value
                var q = new EntityQuery()
                q.TypeName = "CatalogSchema.ItemType"
                q.Fields = ["M_id"]
                q.Query = 'CatalogSchema.ItemType.M_id=="'+ pkgName +'" '
                server.entityManager.getEntities("CatalogSchema.ItemType", "CatalogSchema", q, 0, -1, [], true, false,
                    function(index,total,caller){},
                    function(items, caller){
                        var newCategory = caller.category
                        newCategory.M_items.push(items[0])
                        server.entityManager.saveEntity(newCategory,
                            function(success,caller){
                             
                                mainPage.adminPage.adminCategoryPage.panel.getChildById(date + "-row").delete()
                                    document.getElementById("waitingDiv").style.display = "none"
                                    mainPage.adminPage.adminCategoryPage.appendItem(caller.item, success)
                            },function(){},{"item" : items[0]})
                    },function(){},{"date" : date, "category" : category})
               }
            }(date, category)
             
                
        }
    }(category)
    
      mainPage.adminPage.adminCategoryPage.panel.getChildById(category.M_id + "-saveBtn").element.onclick = function(category){
        return function(){
            document.getElementById("waitingDiv").style.display = ""
            var newCategory = jQuery.extend({}, category)
            newCategory.M_name = mainPage.adminPage.adminCategoryPage.panel.getChildById(category.M_id + "-name").element.value
        
            if(mainPage.adminPage.adminCategoryPage.removedItems[category.M_id]["subCategories"].size > 0){
                    var i = 0;
                for(var object of mainPage.adminPage.adminCategoryPage.removedItems[category.M_id]["subCategories"].values()){
                    i++
                    newCategory.M_categories.pop(object)
                    if(i== mainPage.adminPage.adminCategoryPage.removedItems[category.M_id]["subCategories"].size){
                        if(mainPage.adminPage.adminCategoryPage.removedItems[category.M_id]["items"].size > 0){
                            var j = 0;
                            for(var item of mainPage.adminPage.adminCategoryPage.removedItems[category.M_id]["items"].values()){
                                j++
                                newCategory.M_items.pop(item)
                                if(j == mainPage.adminPage.adminCategoryPage.removedItems[category.M_id]["items"].size){
                                    server.entityManager.saveEntity(newCategory,
                                        function(success,caller){
                                            document.getElementById("waitingDiv").style.display = "none"
                                            mainPage.adminPage.panel.getChildById("adminCategorySaveState").element.innerHTML = ""
                                             mainPage.adminPage.adminCategoryPage.panel.getChildById(newCategory.M_id + "-selector").element.innerHTML = newCategory.M_name
                                                mainPage.adminPage.adminCategoryPage.loadAdminControl(newCategory)
                                            mainPage.adminPage.adminCategoryPage.panel.getChildById(newCategory.M_id + "-saveBtn").element.classList.add("disabled")
                                            mainPage.adminPage.adminCategoryPage.panel.getChildById(newCategory.M_id + "-cancelBtn").element.classList.add("disabled")
                                            
                                            mainPage.adminPage.adminCategoryPage.panel.getChildById(newCategory.M_id + "-control").element.classList.add("active")
                                             mainPage.adminPage.adminCategoryPage.panel.getChildById(newCategory.M_id + "-control").element.classList.add("show")
                                        },function(){},{})
                                }
                            }
                        }else{
                            server.entityManager.saveEntity(newCategory,
                                function(success,caller){
                                    document.getElementById("waitingDiv").style.display = "none"
                                    mainPage.adminPage.panel.getChildById("adminCategorySaveState").element.innerHTML = ""
                                     mainPage.adminPage.adminCategoryPage.panel.getChildById(newCategory.M_id + "-selector").element.innerHTML = newCategory.M_name
                                        mainPage.adminPage.adminCategoryPage.loadAdminControl(newCategory)
                                    mainPage.adminPage.adminCategoryPage.panel.getChildById(newCategory.M_id + "-saveBtn").element.classList.add("disabled")
                                    mainPage.adminPage.adminCategoryPage.panel.getChildById(newCategory.M_id + "-cancelBtn").element.classList.add("disabled")
                                    
                                    mainPage.adminPage.adminCategoryPage.panel.getChildById(newCategory.M_id + "-control").element.classList.add("active")
                                     mainPage.adminPage.adminCategoryPage.panel.getChildById(newCategory.M_id + "-control").element.classList.add("show")
                                },function(){},{})
                        }
                        
                    }
                }
            }else{
               if(mainPage.adminPage.adminCategoryPage.removedItems[category.M_id]["items"].size > 0){
                    var j = 0;
                    for(var item of mainPage.adminPage.adminCategoryPage.removedItems[category.M_id]["items"].values()){
                        j++
                        newCategory.M_items.pop(item)
                        if(j == mainPage.adminPage.adminCategoryPage.removedItems[category.M_id]["items"].size){
                            server.entityManager.saveEntity(newCategory,
                                function(success,caller){
                                    document.getElementById("waitingDiv").style.display = "none"
                                    mainPage.adminPage.panel.getChildById("adminCategorySaveState").element.innerHTML = ""
                                     mainPage.adminPage.adminCategoryPage.panel.getChildById(newCategory.M_id + "-selector").element.innerHTML = newCategory.M_name
                                        mainPage.adminPage.adminCategoryPage.loadAdminControl(newCategory)
                                    mainPage.adminPage.adminCategoryPage.panel.getChildById(newCategory.M_id + "-saveBtn").element.classList.add("disabled")
                                    mainPage.adminPage.adminCategoryPage.panel.getChildById(newCategory.M_id + "-cancelBtn").element.classList.add("disabled")
                                    
                                    mainPage.adminPage.adminCategoryPage.panel.getChildById(newCategory.M_id + "-control").element.classList.add("active")
                                     mainPage.adminPage.adminCategoryPage.panel.getChildById(newCategory.M_id + "-control").element.classList.add("show")
                                },function(){},{})
                        }
                    }
                }else{
                    server.entityManager.saveEntity(newCategory,
                        function(success,caller){
                            document.getElementById("waitingDiv").style.display = "none"
                            mainPage.adminPage.panel.getChildById("adminCategorySaveState").element.innerHTML = ""
                             mainPage.adminPage.adminCategoryPage.panel.getChildById(newCategory.M_id + "-selector").element.innerHTML = newCategory.M_name
                                mainPage.adminPage.adminCategoryPage.loadAdminControl(newCategory)
                            mainPage.adminPage.adminCategoryPage.panel.getChildById(newCategory.M_id + "-saveBtn").element.classList.add("disabled")
                            mainPage.adminPage.adminCategoryPage.panel.getChildById(newCategory.M_id + "-cancelBtn").element.classList.add("disabled")
                            
                            mainPage.adminPage.adminCategoryPage.panel.getChildById(newCategory.M_id + "-control").element.classList.add("active")
                             mainPage.adminPage.adminCategoryPage.panel.getChildById(newCategory.M_id + "-control").element.classList.add("show")
                        },function(){},{})
                }
            }
        
        
        }
    }(category)
    
    mainPage.adminPage.adminCategoryPage.panel.getChildById(category.M_id + "-cancelBtn").element.onclick = function(category){
        return function(){
           mainPage.adminPage.panel.getChildById("adminCategorySaveState").element.innerHTML = ""
            mainPage.adminPage.adminCategoryPage.panel.getChildById(category.M_id + "-saveBtn").element.classList.add("disabled")
            mainPage.adminPage.adminCategoryPage.panel.getChildById(category.M_id + "-cancelBtn").element.classList.add("disabled")
            mainPage.adminPage.adminCategoryPage.loadAdminControl(category)
            mainPage.adminPage.adminCategoryPage.panel.getChildById(category.M_id + "-control").element.classList.add("active")
             mainPage.adminPage.adminCategoryPage.panel.getChildById(category.M_id + "-control").element.classList.add("show")
        }
    }(category)
}

AdminCategoryPage.prototype.appendSubCategory = function(subCategory, category){
    if(subCategory != undefined && category != undefined){
        mainPage.adminPage.adminCategoryPage.panel.getChildById(category.M_id +"-subCategoriesList").appendElement({"tag" : "tr", "id" : subCategory.M_id + "-adminSubCategoryRow"}).down()
        .appendElement({"tag" :"td","innerHtml" : subCategory.M_id, "id" : subCategory.M_id + "-subCategoryID"})
        .appendElement({"tag" : "td", "innerHtml" : subCategory.M_name, "id" : subCategory.M_id + "-subCategoryName"})
        .appendElement({"tag":"td"}).down()
        .appendElement({"tag" : "button", "class" : "btn btn-danger btn-sm", "id" : subCategory.M_id + "-deleteSubCategoryRowAdminBtn", "style":"display: inline-flex; height: 29px;"}).down()
        .appendElement({"tag" : "i", "class" : "fa fa-trash-o"}).up()
        
         mainPage.adminPage.adminCategoryPage.panel.getChildById(subCategory.M_id + "-deleteSubCategoryRowAdminBtn").element.onclick = function(subCategory, category){
            return function(){
                mainPage.adminPage.panel.getChildById("adminCategorySaveState").element.innerHTML = "*"
                mainPage.adminPage.adminCategoryPage.removedItems[category.M_id]["subCategories"].set(subCategory.M_id + "-subCategory", subCategory)
                var row = mainPage.adminPage.adminCategoryPage.panel.getChildById(subCategory.M_id + "-adminSubCategoryRow")
                row.delete()
                mainPage.adminPage.adminCategoryPage.panel.getChildById(category.M_id + "-saveBtn").element.classList.remove("disabled")
                mainPage.adminPage.adminCategoryPage.panel.getChildById(category.M_id + "-cancelBtn").element.classList.remove("disabled")
                console.log( mainPage.adminPage.adminCategoryPage.removedItems[category.M_id])
            }
        }(subCategory, category)
    }
    
    
}

AdminCategoryPage.prototype.appendItem = function(item, category){
    mainPage.adminPage.adminCategoryPage.panel.getChildById(category.M_id +"-itemsList").appendElement({"tag" : "tr", "id" : item.M_id + "-adminItemRow"}).down()
    .appendElement({"tag" :"td","innerHtml" : item.M_id, "id" : item.M_id + "-itemID", "class" : "tabLink"})
    .appendElement({"tag" : "td", "innerHtml" : item.M_name, "id" : item.M_id + "-itemName"})
    .appendElement({"tag":"td"}).down()
    .appendElement({"tag" : "button", "class" : "btn btn-danger btn-sm", "id" : item.M_id + "-deleteItemRowAdminBtn", "style":"display: inline-flex; height: 29px;"}).down()
    .appendElement({"tag" : "i", "class" : "fa fa-trash-o"}).up()
    
    mainPage.adminPage.adminCategoryPage.panel.getChildById(item.M_id + "-deleteItemRowAdminBtn").element.onclick = function(item, category){
        return function(){
            mainPage.adminPage.panel.getChildById("adminCategorySaveState").element.innerHTML = "*"
            mainPage.adminPage.adminCategoryPage.removedItems[category.M_id]["items"].set(item.M_id + "-item", item)
            var row = mainPage.adminPage.adminCategoryPage.panel.getChildById(item.M_id + "-adminItemRow")
            row.delete()
            mainPage.adminPage.adminCategoryPage.panel.getChildById(category.M_id + "-saveBtn").element.classList.remove("disabled")
            mainPage.adminPage.adminCategoryPage.panel.getChildById(category.M_id + "-cancelBtn").element.classList.remove("disabled")
            console.log( mainPage.adminPage.adminCategoryPage.removedItems[category.M_id])
        }
    }(item, category)
    
     mainPage.adminPage.adminCategoryPage.panel.getChildById(item.M_id + "-itemID").element.onclick = function(item){
        return function(){
            mainPage.itemDisplayPage.displayTabItem(item)
        }
    }(item)
    
}

AdminCategoryPage.prototype.deleteSubCategories = function(subCategories){
    for(var i = 0; i< subCategories.length; i++){
        server.entityManager.getEntityByUuid(subCategories[i],false,
            function(subCategory, caller){
                
                if(subLocalisation.M_categories.length > 0){
                    mainPage.adminPage.adminCategoryPage.deleteSubCategories(subCategory.M_categories)
                }
                server.entityManager.removeEntity(subCategory.UUID,
                    function(success,caller){
                        console.log(success)
                    },function(){},{})
            },function(){},{})
        
    }
}