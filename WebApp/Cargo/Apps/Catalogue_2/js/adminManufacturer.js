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
        .appendElement({"tag" : "div", "class" : "col-md-9 bg-light"}).down()
         .appendElement({"tag" : "div", "class" : "tab-content", "id" : "manufacturersAdminControl"}).up().up().up()
        

    this.modifiedItems = {}
  
    this.removedItems = {}
    
    this.panel.getChildById("manufacturerFilterKeywords").element.addEventListener('change', function(){
        results = mainPage.adminPage.adminManufacturerPage.getManufacturersFromKeyword(this.value.toUpperCase())
    })
    
    
    this.panel.getChildById("addManufacturerButton").element.onclick = function(){
      
        mainPage.adminPage.adminManufacturerPage.panel.getChildById("manufacturersFilteredList").appendElement({"tag":"a", "class" : "list-group-item list-group-item-action","id" : "newManufacturer-selector","data-toggle" : "tab","href" : "#"+ "newManufacturer-control","role":"tab", "innerHtml": "Nouveau manufacturier","aria-controls": "newManufacturer-control"})
         mainPage.adminPage.adminManufacturerPage.panel.getChildById("manufacturersAdminControl").appendElement({"tag" : "div", "class" : "tab-pane", "id" : "newManufacturer-control", "role" : "tabpanel", "aria-labelledby" : "newManufacturer-selector", "style" : "padding:15px;"}).down()
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
    q.TypeName = "CatalogSchema.Manufacturer"
    q.Fields = ["M_name","M_id"]
    q.Query = 'CatalogSchema.ManufacturerType.M_name~="'+ keyword +'" '
    
    server.entityManager.getEntities("CatalogSchema.ManufacturerType", "CatalogSchema", q, 0,-1,[],true,true, 
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
    mainPage.adminPage.adminManufacturerPage.modifiedItems[manufacturer.M_id] = new Map()
    mainPage.adminPage.adminManufacturerPage.removedItems[manufacturer.M_id] = new Map()
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
    .appendElement({"tag":"label", "for" : manufacturer.M_id + "-packagesList", "innerHtml" : "Produits","class" : "col-3 d-flex align-items-center"})
    .appendElement({"tag": "input", "class" : "form-control col-5 m-1", "id" : manufacturer.M_id + "-filterPackages","type" :"text", "placeholder" : "Filtrer.."}).up()
    .appendElement({"tag" : "div"}).down()
    .appendElement({"tag":"table", "class":"table table-hover", "style" : "display: block;height: 32em;overflow: auto;"}).down()
    .appendElement({"tag":"thead"}).down()
    .appendElement({"tag" : "tr"}).down()
    .appendElement({"tag" : "th", "innerHtml" : "ID Paquet", "scope" : "col"})
    .appendElement({"tag" : "th", "innerHtml" : "Nom", "scope" : "col"})

    .appendElement({"tag" : "th", "scope" : "col"}).down()
    .appendElement({"tag" : "button", "class" : "btn btn-success btn-sm", "id" : manufacturer.M_id +"-addPackageButton", "style":"display: inline-flex; height: 29px;"}).down()
    .appendElement({"tag" : "i", "class" : "fa fa-plus"}).up().up()
    .up().up()
    .appendElement({"tag" : "tbody", "id" : manufacturer.M_id + "-packagesList"})
    .up().up()
    .appendElement({"tag" : "div", "class" : "d-flex justify-content-center"}).down()
    .appendElement({"tag" : "button", "class" : "btn btn-primary disabled mr-3", "innerHtml" : "Enregistrer", "id" : manufacturer.M_id + "-saveBtn"})
    .appendElement({"tag" : "button", "class" : "btn disabled", "innerHtml" : "Annuler les modifications", "id" : manufacturer.M_id + "-cancelBtn"})
    
    
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
                    
             
             mainPage.adminPage.adminManufacturerPage.panel.getChildById(date+ "-confirmBtn").element.onclick = function (date, manufacturer) {
           return function(){
                document.getElementById("waitingDiv").style.display = ""
                var pkgName = mainPage.adminPage.adminManufacturerPage.panel.getChildById(date+"-packageID").element.value
                var q = new EntityQuery()
                q.TypeName = "CatalogSchema.PackageType"
                q.Fields = ["M_id"]
                q.Query = 'CatalogSchema.PackageType.M_id=="'+ pkgName +'" '
                server.entityManager.getEntities("CatalogSchema.PackageType", "CatalogSchema", q, 0, -1, [], true, false,
                    function(index,total,caller){},
                    function(packages, caller){
                        var newManufacturer = caller.manufacturer
                        newManufacturer.M_packages.push(packages[0])
                        server.entityManager.saveEntity(newManufacturer,
                            function(success,caller){
                             
                                mainPage.adminPage.adminManufacturerPage.panel.getChildById(date + "-row").delete()
                                    document.getElementById("waitingDiv").style.display = "none"
                                    mainPage.adminPage.adminManufacturerPage.appendPackage(caller.package, success)
                            },function(){},{"package" : packages[0]})
                           
                                
                                
                                
                                
                            
                           
                    },function(){},{"date" : date, "manufacturer" : manufacturer})
               }
            }(date, manufacturer)
             
                
        }
    }(manufacturer)
    
    
    
    
}

AdminManufacturerPage.prototype.appendPackage = function(pkg, manufacturer){
    console.log(pkg,manufacturer)
}