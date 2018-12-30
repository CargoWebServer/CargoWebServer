
 
var PackageDisplayPage = function (parent) {

    // The result content.
    this.panel = new Element(parent, { "tag": "div", "id": "package_display_page_panel", "class": "item_display container-fluid" , "style" :"margin-top : 15px;"})

    this.navbar = this.panel.appendElement({"tag" : "ul", "class" : "nav nav-tabs querynav printHide","id" : "queries"}).down()
    // Now I will display the list of results...
    this.packageSwitchContextBtn = this.navbar.appendElement({"tag":"button","class":"btn btn-dark toggleBtn", "innerHtml" :"Voir les recherches"}).down()
    this.packageSwitchContextBtn.element.onclick = function packageSwitchContext(){
        if(document.getElementById("item_search_result_page") != null){
            document.getElementById("package_display_page_panel").style.display = "none"
            document.getElementById("item_search_result_page").style.display = ""
        }else{
            document.getElementById("package_display_page_panel").style.display = "none"
            document.getElementById("main-container").style.display = ""
        }
        
        fireResize()
    }

    this.packageResultPanel = this.panel.appendElement({ "tag": "div", "class": "search_results_content tab-content" }).down()
    this.currentItem = null
    
    return this
}
  



function showPackage(itemID){
    packageElement = document.getElementById(itemID+"-packagesContent")
    packageElement.scrollIntoView({
        behavior : 'smooth'
    })
}
 
  

function parsePkgImage(itemID){
    server.fileManager.readDir("/Catalogue_2/photo/" +itemID,
    function(results, caller){
        var photos = results[0]
        if(photos != null){
            console.log(photos)
            for(var i=0; i < photos.length; i++){
                if(i == 0){
                    if(photos[i].endsWith(".jpg") || photos[i].endsWith(".png")){
                        var child = document.createElement("li")
                        child.setAttribute("data-target", "#" + itemID + "-carousel")
                        child.setAttribute("data-slide-to" , "0")
                        child.setAttribute("class", "active")
                        document.getElementById(itemID + "_p-carousel-indicators").appendChild(child)
                        child = document.createElement("div")
                        child.setAttribute("class", "carousel-item active")
                        var picture = document.createElement("img")
                        picture.setAttribute("src", "/Catalogue_2/photo/" +itemID+"/" + photos[i])
                        picture.setAttribute("class", "img-fluid rounded carouselImage")
                        child.appendChild(picture)
                        document.getElementById(itemID + "_p-carousel-inner").appendChild(child)
                    }else{
                        var child = document.createElement("li")
                        child.setAttribute("data-target", "#" + itemID + "-carousel")
                        child.setAttribute("data-slide-to" , "0")
                        child.setAttribute("class", "active")
                        document.getElementById(itemID + "_p-carousel-indicators").appendChild(child)
                        child = document.createElement("div")
                        child.setAttribute("class", "carousel-item active")
                        var video = document.createElement("video")
                        video.setAttribute("src", "/Catalogue_2/photo/" +itemID+"/" + photos[i])
                        video.setAttribute("type", "video/mp4")
                        video.setAttribute("class", "img-fluid rounded carouselImage")
                        video.autoplay = true
                        video.muted = true
                        video.controls = true
                        video.load()
                        child.appendChild(video)
                        document.getElementById(itemID + "_p-carousel-inner").appendChild(child)
                    }
                    
                }else{
                    if(photos[i].endsWith(".jpg") || photos[i].endsWith(".png")){
                       var child = document.createElement("li")
                        child.setAttribute("data-target", "#"+ itemID + "-carousel")
                        child.setAttribute("data-slide-to", i)
                        document.getElementById(itemID + "_p-carousel-indicators").appendChild(child)
                        child = document.createElement("div")
                        child.setAttribute("class", "carousel-item")
                        var picture = document.createElement("img")
                        picture.setAttribute("src", "/Catalogue_2/photo/" +itemID+"/" + photos[i])
                        picture.setAttribute("class", "img-fluid rounded carouselImage")
                        child.appendChild(picture)
                        document.getElementById(itemID + "_p-carousel-inner").appendChild(child)
                    }else{
                         var child = document.createElement("li")
                        child.setAttribute("data-target", "#"+ itemID + "-carousel")
                        child.setAttribute("data-slide-to", i)
                        document.getElementById(itemID + "_p-carousel-indicators").appendChild(child)
                        child = document.createElement("div")
                        child.setAttribute("class", "carousel-item")
                        var video = document.createElement("video")
                        video.setAttribute("src", "/Catalogue_2/photo/" +itemID+"/" + photos[i])
                        video.setAttribute("type", "video/mp4")
                        video.setAttribute("class", "img-fluid rounded carouselImage")
                        video.autoplay = true
                        video.muted = true
                        video.controls = true
                        video.load()
                        child.appendChild(video)
                        document.getElementById(itemID + "_p-carousel-inner").appendChild(child)
                    }
                    
                }
            }
        }
    },function(errObj,caller){
    },{"itemID":itemID})
}

/**
 * Display an item information.
 * @param {*} item 
 */
PackageDisplayPage.prototype.displayTabItem = function (pkg) {
    
    var packageID = pkg.M_id  + "_p"
    if(document.getElementById("admin_page_panel") != null){
        document.getElementById("admin_page_panel").style.display = "none"
    }
    document.getElementById("main-container").style.display = "none"
    document.getElementById("package_display_page_panel").style.display = ""
    document.getElementById("item_display_page_panel").style.display = "none"
    document.getElementById("order_page_panel").style.display = "none"
    if(document.getElementById("order_history_page_panel") != null){
        document.getElementById("order_history_page_panel").style.display = "none"
    }
    if(document.getElementById("item_search_result_page")!=null){
        document.getElementById("item_search_result_page").style.display = "none"
    }
    
    
    if(document.getElementById(packageID + "-li") !== null){
        document.getElementById(packageID + "-li").childNodes[0].click()
        return;
    }
     this.navbar.appendElement({"tag" : "li", "class" : "nav-item","id" : packageID + "-li", "style" : "display:flex;flex-direction:row;"}).down()
    .appendElement({"tag" : "span", "class" : "nav-link active", "href" : "#" + packageID + "-query", "role" : "tab", "data-toggle" : "tab", "aria-controls" : packageID  + "-query","aria-selected" : "true", "id" : packageID + "-tab", "style" : "display: flex;flex-direction: row;align-items: center; padding: 5px;border-radius:0;"}).down()
    .appendElement({"tag":"a","class" : "tabLink", "href" : "#" +packageID + "-query", "innerHtml" : packageID , "style" : "text-decoration : none;"})
    .appendElement({"tag" : "a","id" : packageID+"-close", "aria-label" : "Close" , "class" : "tabLink", "id" : packageID+"-closebtn"}).down()
    .appendElement({"tag" : "span", "aria-hidden" :"true", "innerHtml" : "&times;", "style" : "padding-left: 5px;color:#007bff"}).up()
    
    this.packageResultPanel.appendElement({"tag" : "div","class" : "tab-pane show result-tab-content", "style" : "position:relative; background-color : white;", "id" : packageID + "-query", "role" : "tabpanel", "aria-labelledby" :packageID + "-tab" }).down()
    .appendElement({"tag" : "div", "class" : "jumbotron"}).down()
    .appendElement({"tag" : "div", "class": "d-flex justify-content-center flex-row"}).down()
   .appendElement({"tag" : "span"}).down()
    .appendElement({"tag" : "h1", "style" : "text-align:center;", "innerHtml" : pkg.M_name}).up()
    .appendElement({"tag" : "span", "class" : "d-flex align-items-center"}).down()
    .appendElement({"tag" : "button", "class" : "btn btn-outline-dark ml-3", "id" : packageID + "-loadPkgAdminEdit"}).down()
    .appendElement({"tag" : "span",  "innerHtml" : "Modifier le paquet"})
    .appendElement({"tag" : "i", "class" : "fa fa-wrench ml-1"}).up()
    .appendElement({"tag" : "button", "class" : "btn btn-outline-dark ml-3", "id" : packageID + "-printBtn"}).down()
    .appendElement({"tag" : "i", "class" : "fa fa-print", "title" : "Imprimer le paquet"}).up().up().up()
    .appendElement({"tag" : "hr", "class" : "my-4"})
    .appendElement({"tag" :"div", "class" : "container-fluid"}).down()
    .appendElement({"tag" : "span", "class" : "row"}).down()
    .appendElement({"tag" : "div", "class" : "col-md order-0 printHide"}).down()
    .appendElement({"tag" : "div", "class" : "row paddedhor"}).down()
    .appendElement({"tag" : "div", "id" : packageID + "-carousel", "class" : "carousel slide d-flex justify-content-center ", "data-ride": "carousel", "data-interval" : "false"}).down()
    .appendElement({"tag" : "ol", "class" : "carousel-indicators printHide", "id" : packageID + "-carousel-indicators"})
    .appendElement({"tag" : "div", "class" : "carousel-inner rounded", "id" : packageID + "-carousel-inner"})
    .appendElement({"tag" : "a", "class" : "carousel-control-prev printHide", "href" : "#" + packageID + "-carousel", "role" : "button","data-slide" :"prev" , "id" : packageID + "-carousel-control-prev"}).down()
    .appendElement({"tag" : "span", "class" : "carousel-control-prev-icon", "aria-hidden" : "true"})
    .appendElement({"tag" : "span", "class" : "sr-only", "innerHtml" : "Previous"}).up()
    .appendElement({"tag" : "a", "class" : "carousel-control-next printHide", "href" : "#" + packageID + "-carousel", "role" : "button","data-slide" :"next" , "id" : packageID + "-carousel-control-next"}).down()
    .appendElement({"tag" : "span", "class" : "carousel-control-next-icon", "aria-hidden" : "true"})
    .appendElement({"tag" : "span", "class" : "sr-only", "innerHtml" : "Previous"}).up().up().up().up()
    
    .appendElement({"tag" : "div", "class" : "col-md order-1"}).down()
    .appendElement({"tag" : "div", "class" : "list-group ", "id" : packageID + "-detailsGroup"}).down()
    .appendElement({"tag" : "h4", "class" : "row d-flex justify-content-center list-group-item bg-dark text-light", "innerHtml" : "Détails"})
    .appendElement({"tag" : "div", "class" :"row list-group-item itemRow d-flex"}).down()
    .appendElement({"tag" : "div", "class" : "col largeCol", "innerHtml" : "ID", "style" : "border-right :1px solid #e9ecef;"})
    .appendElement({"tag" : "div", "class" : "col largeCol list-group-item-action ", "innerHtml" : pkg.M_id}).up()
    .appendElement({"tag" : "div", "class" :"row list-group-item itemRow d-flex"}).down()
    .appendElement({"tag" : "div", "class" : "col largeCol", "innerHtml" : "Description", "style" : "border-right :1px solid #e9ecef;"})
    .appendElement({"tag" : "div", "class" : "col largeCol list-group-item-action ", "innerHtml" : pkg.M_description}).up()
    .appendElement({"tag" : "div", "class" :"row list-group-item  itemRow d-flex"}).down()
    .appendElement({"tag" : "div", "class" : "col  largeCol", "innerHtml" : "Quantité", "style" : "border-right : 1px solid #e9ecef;"})
    .appendElement({"tag" : "div", "class" : "col  largeCol list-group-item-action", "innerHtml" : pkg.M_quantity.toString()}).up()
    .appendElement({"tag" : "div", "class" :"row list-group-item  itemRow d-flex"}).down()
    .appendElement({"tag" : "div", "class" : "col largeCol", "innerHtml" : "Paquets", "style" : "border-right : 1px solid #e9ecef;"})
    .appendElement({"tag" : "div", "class" : "col  largeCol list-group-item-action","id" : packageID + "-packaged"}).up().up()
    
    .appendElement({"tag" : "hr", "class" : "my-4", "style" : "margin:1.5rem;opacity:0;"})
    
    
    .appendElement({"tag" : "div", "class" : "list-group ", "id" : packageID + "-inventoryGroup"}).down()
    .appendElement({"tag" : "h4", "class" : "row d-flex justify-content-center list-group-item bg-dark text-light", "innerHtml" : "Inventaire"}).up().up().up().up()
    .appendElement({"tag" : "span" , "class" : "row order-2"}).down()
    .appendElement({"tag" : "div", "class" : "col-md"}).down()
    .appendElement({"tag" : "hr", "class" : "my-4", "style" : "margin:1.5rem;opacity:0;"})
  
    
    .appendElement({"tag" : "div", "id" : packageID + "-items", "class" : "list-group"}).down()
    .appendElement({"tag" : "div", "class" : "text-light bg-dark d-flex justify-content-center align-items-center flex-direction-row list-group-item printNoBorder printNoPadding", "style"  : "font-size : 1.5rem;"}).down()
    .appendElement({"tag" : "span", "innerHtml" : "Items"})
     .appendElement({"tag" : "i", "class" : "fa fa-object-ungroup", "style" : "margin:10px;"}).up().up()
     
     .appendElement({"tag" : "hr", "class" : "my-4", "style" : "margin:1.5rem;opacity:0;"})
    .appendElement({"tag" : "div", "class" : "list-group innerShadow printNoBorder", "id" : packageID + "-suppliersDetails"}).down()
    .appendElement({"tag" : "h4", "class" : "row d-flex justify-content-center list-group-item bg-dark text-light printNoBorder", "innerHtml" : "Fournisseurs"})
    .appendElement({"tag" : "hr", "class" : "my-4 printHide", "style" : "margin:1.5rem;opacity:0;"})
    
    this.packageResultPanel.getChildById(packageID + "-printBtn").element.onclick = function(){
        window.print()
    }
    
    this.packageResultPanel.getChildById(packageID+ "-loadPkgAdminEdit").element.onclick = function(pkg){
        return function(){
            mainPage.adminPage.adminPackagePage.getPackagesFromKeyword(pkg.M_id)
             document.getElementById("package_display_page_panel").style.display = "none"
              document.getElementById("admin_page_panel").style.display = ""
              document.getElementById("package-adminTab").click()
        }
    }(pkg)
    
   
   if(pkg.M_packages.length == 0){
       this.packageResultPanel.getChildById(packageID+"-packaged").element.innerHTML = "Aucun"
   }else{
       for(var i  = 0; i < pkg.M_packages.length; i++){
           server.entityManager.getEntityByUuid(pkg.M_packages[i], false, 
           function(item_package, caller){
                caller.packagesDiv.appendElement({"tag" : "span", "class" : "d-flex"}).down()
                .appendElement({"tag" : "a", "aria-label" : "Show" , "class" : "m-1"}).down()
                .appendElement({"tag" : "span","class" : "tabLink","innerHtml" : item_package.M_id,"id" : item_package.M_id+"-showbtn"})
                
                caller.packagesDiv.getChildById(item_package.M_id + "-showbtn").element.onclick = function(item_package){
                    return function(){
                        mainPage.packageDisplayPage.displayTabItem(item_package)
                    }
                }(item_package)
           }, function () {},
           {"packagesDiv" : this.packageResultPanel.getChildById(packageID + "-packaged")})
       }
   }
   
   
   for(var i  = 0; i < pkg.M_items.length; i++){
       server.entityManager.getEntityByUuid(pkg.M_items[i], false, 
       function(item, caller){
           caller.itemSection.appendElement({"tag" : "div", "class" :"list-group-item row itemRow d-flex innerShadow printNoBorder", "id" : item.M_id + "-itemRow", "style" : "margin-left : 0px; margin-right:0px;"}).down()
           .appendElement({"tag" : "div", "class" : "col-md-3 d-flex justify-content-center align-items-center"}).down()
           .appendElement({"tag": "img","class" : "img-fluid rounded", "style" : "margin:10px;", "id" : item.M_id + "-image"}).up()
           .appendElement({"tag" : "div", "class" : "col-md largeCol"}).down()
           .appendElement({"tag" : "div", "class" : "list-group"}).down()
           .appendElement({"tag" : "div", "class" : "row list-group-item itemRow d-flex"}).down()
           .appendElement({"tag" : "div", "class" : "col largeCol", "innerHtml" : "ID"})
           .appendElement({"tag" : "div", "class" : "col largeCol list-group-item-action", "innerHtml" : item.M_id}).up()
           .appendElement({"tag" : "div", "class" : "row list-group-item itemRow d-flex"}).down()
           .appendElement({"tag" : "div", "class" : "col largeCol", "innerHtml" : "Nom"})
           .appendElement({"tag" : "div", "class" : "col largeCol list-group-item-action", "innerHtml" : item.M_name}).up()
           .appendElement({"tag" : "div", "class" : "row list-group-item itemRow d-flex"}).down()
           .appendElement({"tag" : "div", "class" : "col largeCol", "innerHtml" : "Alias"})
           .appendElement({"tag" : "div", "class" : "col largeCol  list-group-item-action", "innerHtml" : item.M_alias}).up().up().up()
           .appendElement({"tag" : "div", "class" : "col-md-2 largeCol d-flex justify-content-center align-items-center"}).down()
           .appendElement({"tag" : "button", "class" : "btn btn-outline-dark", "id" : item.M_id + "-showItemButton"}).down()
           .appendElement({"tag" : "span", "innerHtml" : "Voir l'item"})
           .appendElement({"tag" : "i", "class" : "fa fa-object-ungroup", "style" : "margin:10px;"}).up()
           
            parseFirstImage(item.M_id, caller.itemSection.getChildById(item.M_id + "-image"))
           
        caller.itemSection.getChildById(item.M_id + "-showItemButton").element.addEventListener('click', function(){
            mainPage.itemDisplayPage.displayTabItem(item)
        })
          
           
       }, function (){
           
       }, {"itemSection" : this.packageResultPanel.getChildById(packageID + "-items")})
   }
   
   for(var i = 0; i < pkg.M_supplied.length; i++){
        server.entityManager.getEntityByUuid(pkg.M_supplied[i], false, 
        function(item_supplier, caller) {
            caller.pkgContent.appendElement({"tag" : "div", "class" : "list-group", "id" : item_supplier.M_id + "-supplierDetails"}).down()
            .appendElement({"tag" : "div", "class" : "swoop", "style" : "margin: 10px;"})
            
           
            server.entityManager.getEntityByUuid(item_supplier.M_supplier, false, 
            function(supplier, caller){
                caller.itemSupplier.element.innerHTML = ""
                caller.itemSupplier.appendElement({"tag" : "div", "class" : "list-group", "id" : supplier + "-list"}).down()
                .appendElement({"tag" : "h4", "class" : "row d-flex justify-content-center list-group-item bg-dark text-light", "innerHtml" : supplier.M_name})
                .appendElement({"tag" : "div", "class" :"row list-group-item itemRow d-flex"}).down()
                .appendElement({"tag" : "div", "class" : "col largeCol", "innerHtml" : "ID", "style" : "border-right :1px solid #e9ecef;"})
                .appendElement({"tag" : "div", "class" : "col largeCol list-group-item-action ", "innerHtml" : caller.supplierInfo.M_id}).up()
                .appendElement({"tag" : "div", "class" :"row list-group-item itemRow d-flex"}).down()
                .appendElement({"tag" : "div", "class" : "col largeCol", "innerHtml" : "Prix", "style" : "border-right :1px solid #e9ecef;"})
                .appendElement({"tag" : "div", "class" : "col largeCol list-group-item-action ", "innerHtml" : caller.supplierInfo.M_price.M_valueOf +" " + caller.supplierInfo.M_price.M_currency.M_valueOf}).up()
                .appendElement({"tag" : "div", "class" :"row list-group-item itemRow d-flex"}).down()
                .appendElement({"tag" : "div", "class" : "col largeCol", "innerHtml" : "Date", "style" : "border-right :1px solid #e9ecef;"})
                .appendElement({"tag" : "div", "class" : "col largeCol list-group-item-action ", "innerHtml" : caller.supplierInfo.M_date}).up()
                
                .appendElement({"tag" : "div", "class" : "row list-group-item itemRow d-flex printHide", "id":supplier.M_id + "-order-line", "style":"visibility: hidden;"}).down()
                .appendElement({"tag" : "div", "class" : "col largeCol", "innerHtml" : "Quantité à ajouter", "style" : "border-right:1px solid #e9ecef"})
                .appendElement({"tag" : "div", "class" : "col largeCol list-group-item-action "}).down()
                .appendElement({"tag" : "div", "class" : "input-group"}).down()
                
                .appendElement({"tag" :"input", "type" : "number", "class" : "form-control text-center", "value" : "1", "id" : caller.supplierInfo.M_id + "-qtyToOrder"})
                .appendElement({"tag" : "div", "class" : "input-group-append"}).down()

                .appendElement({"tag" : "button", "class" : "btn btn-dark", "type" : "button", "id" : caller.supplierInfo.M_id + "-addButton"}).down()
                .appendElement({"tag" : "span","innerHtml" : "Ajouter "})
                .appendElement({"tag" : "i", "class" : "fa fa-cart-plus", "style" : "margin-left:5px;"}).up().up().up().up().up()
                        
                .appendElement({"tag" : "hr", "class" : "my-4", "style" : "margin:1.5rem;opacity:0;"})
                
                // Here I will attach the line to order...
                catalogMessageHub.attach(caller.itemSupplier.getChildById(supplier.M_id + "-order-line"), welcomeEvent, function(evt, orderLine){
            	    if(server.account === undefined){
            	        return
            	    }
            	    orderLine.element.style.visibility = "visible";
                })
                
                if(server.account !== undefined){
            	   caller.itemSupplier.getChildById(supplier.M_id + "-order-line").element.style.visibility = "visible";
            	}
            	   
                if(caller.callback != undefined){
                    caller.callback()
                }
                
                caller.itemSupplier.getChildById(caller.supplierInfo.M_id + "-addButton").element.onclick = function(supplierRef, qty, ovmm){
                    return function () {
                        
                       var newOrderLine = new CatalogSchema.ItemSupplierOrderType()
                        newOrderLine.M_itemSupplierRef = supplierRef
                        newOrderLine.M_quantity = qty.element.value
                        mainPage.orderPage.addOrderItem(newOrderLine)
                    }
                    
                }(caller.supplierInfo.UUID, caller.itemSupplier.getChildById(caller.supplierInfo.M_id+"-qtyToOrder"), caller.supplierInfo.M_id)
                
            }, 
            function(){},
            {"itemSupplier" : caller.pkgContent.getChildById(item_supplier.M_id + "-supplierDetails"), "supplierInfo" : item_supplier})
        },
        function(){}, { "pkgContent" : this.packageResultPanel.getChildById(packageID+ "-suppliersDetails"), "item_package" : pkg}
        )
    }
   
   
   
   for(var i=0; i < pkg.M_inventoried.length; i++){
        server.entityManager.getEntityByUuid(pkg.M_inventoried[i], false,
            function(inventory, caller){
                var id = (caller.packageID+i+"-place")
                var child = new Element(document.getElementById(packageID+"-inventoryGroup"), {"tag":"div", "class":"row list-group-item itemRow d-flex"})
                var col1 = new Element(child, {"tag":"div", "class":"col largeCol","style":"border-right :1px solid #e9ecef;", "innerHtml":"Emplacement"})
                var col2 = new Element(child, {"tag":"div", "class":"col largeCol list-group-item-action d-flex justify-content-center", "id":id})
                //var temp = new Element(col2, {"tag":"div", "class":"swoop"})
                displayLocalisation(col2, inventory.M_located, function(location){
                    
                    mainPage.searchContext = "localisationSearchLnk";
                    mainPage.searchResultPage.displayResults({"results":[{"data":location}], "estimate":1}, location.M_id, "localisationSearchLnk") 
                })
                       
               caller.inventoryContainer.appendElement({"tag" : "div", "class" :"row list-group-item itemRow d-flex"}).down()
                .appendElement({"tag" : "div", "class" : "col largeCol", "innerHtml" : "Quantité en stock", "style" : "border-right : 1px solid #e9ecef;"})
                .appendElement({"tag" : "div", "class" : "col largeCol  list-group-item-action d-flex justify-content-center", "innerHtml" : inventory.M_quantity}).up()
                .appendElement({"tag" : "div", "class" :"row list-group-item itemRow d-flex"}).down()
                .appendElement({"tag" : "div", "class" : "col largeCol", "innerHtml" : "Stock minimal", "style" : "border-right : 1px solid #e9ecef;"})
                .appendElement({"tag" : "div", "class" : "col largeCol  list-group-item-action d-flex justify-content-center", "innerHtml" : inventory.M_safetyStock}).up()
                .appendElement({"tag" : "div", "class" :"row list-group-item itemRow d-flex"}).down()
                .appendElement({"tag" : "div", "class" : "col largeCol", "innerHtml" : "Quantité pour re-order", "style" : "border-right : 1px solid #e9ecef;"})
                .appendElement({"tag" : "div", "class" : "col largeCol  list-group-item-action d-flex justify-content-center", "innerHtml" : inventory.M_reorderQty}).up()
            },
            function(){
                
            }, {"packageID" : packageID, "i" : i, "inventoryContainer" : this.packageResultPanel.getChildById(packageID + "-inventoryGroup")})
    }
   
   
   
    parsePkgImage(pkg.M_id)
    
    this.navbar.getChildById(packageID+"-closebtn").element.onclick = function(tab,content){
        return function(){
            tabParent = tab.element.parentNode;
            var i = Array.prototype.indexOf.call(tabParent.childNodes, tab.element);
            tabParent.removeChild(tab.element)
            content.element.parentNode.removeChild(content.element)
            if(tabParent.childNodes.length > 1){
                if(i > tabParent.childNodes.length - 1){
                    tabParent.childNodes[tabParent.childNodes.length - 1].childNodes[0].click()
                }else{
                    tabParent.childNodes[i].childNodes[0].click()
                }
            }
        }
    }(this.navbar.getChildById(packageID + "-li"),  this.packageResultPanel.getChildById(packageID+"-query") )
    
    this.navbar.getChildById(packageID+"-closebtn").element.onmouseover = function(){
        this.style.cursor = "pointer"
    }

    this.navbar.getChildById(packageID+"-closebtn").element.onmouseleave = function(){
        this.style.cursor = "default"
    }

    var tab = document.getElementById(packageID + "-tab")
            tab.classList.remove("active")
            tab.click()
            
    document.getElementById("package_display_page_panel").scrollIntoView()
}


