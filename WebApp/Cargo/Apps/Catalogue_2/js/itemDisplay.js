/**
 * Display the item page.
 * @param {*} parent 
 */
 
var ItemDisplayPage = function (parent) {

    // The result content.
    this.panel = new Element(parent, { "tag": "div", "id": "item_display_page_panel", "class": "item_display container-fluid" , "style" :"margin-top : 15px;"})

    this.navbar = this.panel.appendElement({"tag" : "ul", "class" : "nav nav-tabs querynav","id" : "itemqueries"}).down()
    // Now I will display the list of results...
    this.navbar.appendElement({"tag":"button","class":"btn btn-dark toggleBtn", "onclick" : "itemSwitchContext()","innerHtml" :"Voir les recherches"})
    
    this.itemResultPanel = this.panel.appendElement({ "tag": "div", "class": "search_results_content tab-content" }).down()
    this.currentItem = null
    
    return this
}
  
function itemSwitchContext(){
    document.getElementById("item_display_page_panel").style.display = "none"
    document.getElementById("item_search_result_page").style.display = ""
    fireResize()
}

function showPackage(itemID){
    packageElement = document.getElementById(itemID+"-packagesContent")
    packageElement.scrollIntoView({
        behavior : 'smooth'
    })
}
 
function checkImageExists(imagePath, callBack) {
    server.fileManager.readDir(imagePath,
    function(result, caller){
        callBack(result[0].length)
    },function(errObj,caller){
        console.log("error")
    })
}   

function parseImage(itemID){
    checkImageExists("/Catalogue_2/photo/" +itemID + "/", function(existsImage) {
        for(var i =0; i< existsImage; i++) {
            // image exist
            if(i == 0){
                var child = document.createElement("li")
                child.setAttribute("data-target", "#" + itemID + "-carousel")
                child.setAttribute("data-slide-to" , "0")
                child.setAttribute("class", "active")
                document.getElementById(itemID + "-carousel-indicators").appendChild(child)
                child = document.createElement("div")
                child.setAttribute("class", "carousel-item active")
                var picture = document.createElement("img")
                picture.setAttribute("src", "../../Catalogue_2/photo/" + itemID + "/photo_" + Number(i+1) + ".jpg")
                picture.setAttribute("class", "img-fluid rounded carouselImage")
                child.appendChild(picture)
                document.getElementById(itemID + "-carousel-inner").appendChild(child)
            }else{
                var child = document.createElement("li")
                child.setAttribute("data-target", "#"+ itemID + "-carousel")
                child.setAttribute("data-slide-to", i)
                document.getElementById(itemID + "-carousel-indicators").appendChild(child)
                child = document.createElement("div")
                child.setAttribute("class", "carousel-item")
                var picture = document.createElement("img")
                picture.setAttribute("src", "../../Catalogue_2/photo/" + itemID + "/photo_" + Number(i+1) + ".jpg")
                picture.setAttribute("class", "img-fluid rounded carouselImage")
                child.appendChild(picture)
                document.getElementById(itemID + "-carousel-inner").appendChild(child)
            }
            
        }
        
    });
} 
/**
 * Display an item information.
 * @param {*} item 
 */
ItemDisplayPage.prototype.displayTabItem = function (item, callback) {
    
    this.currentItem = item
    
    document.getElementById("item_display_page_panel").style.display = ""
    if(document.getElementById("item_search_result_page") != null){
        document.getElementById("item_search_result_page").style.display = "none"
    }
    
    document.getElementById("package_display_page_panel").style.display = "none"
    if(document.getElementById(item.M_id + "-itemli") !== null){
        document.getElementById(item.M_id + "-itemli").childNodes[0].click()
        return;
    }
    
    this.navbar.appendElement({"tag" : "li", "class" : "nav-item","id" : item.M_id + "-itemli", "style" : "display:flex;flex-direction:row;"}).down()
    .appendElement({"tag" : "span", "class" : "nav-link active", "href" : "#" + item.M_id + "-itemquery", "role" : "tab", "data-toggle" : "tab", "aria-controls" : item.M_id  + "-itemquery","aria-selected" : "true", "id" : item.M_id + "-tab", "style" : "display: flex;flex-direction: row;align-items: center; padding: 5px;border-radius:0;"}).down()
    .appendElement({"tag":"a","class" : "tabLink", "href" : "#" +item.M_id + "-itemquery", "innerHtml" : item.M_id , "style" : "text-decoration : none;"})
    .appendElement({"tag" : "a","id" : item.M_id+"-close", "aria-label" : "Close" , "class" : "tabLink", "id" : item.M_id+"-closebtn"}).down()
    .appendElement({"tag" : "span", "aria-hidden" :"true", "innerHtml" : "&times;", "style" : "padding-left: 5px;color:#007bff"}).up()
    
    this.itemResultPanel.appendElement({"tag" : "div","class" : "tab-pane show result-tab-content", "style" : "position:relative; background-color : white;", "id" : item.M_id + "-itemquery", "role" : "tabpanel", "aria-labelledby" :item.M_id + "-tab" }).down()
    .appendElement({"tag" : "div", "class" : "jumbotron"}).down()
    .appendElement({"tag" : "div", "class": "d-flex justify-content-around"}).down()
    .appendElement({"tag" : "h1", "style" : "text-align:center;", "innerHtml" : item.M_name})
    .appendElement({"tag" : "button", "class" : "btn btn-outline-dark", "id" : item.M_id + "-packageButton"}).down()
    .appendElement({"tag" : "span","style" : "font-size : 1.75rem;","innerHtml" : "Aller aux paquets "})
    .appendElement({"tag" : "i", "class" : "fa fa-object-group","style" : "font-size:1.75rem;"}).up().up()
    .appendElement({"tag" : "hr", "class" : "my-4"})
    .appendElement({"tag" :"div", "class" : "container-fluid"}).down()
    .appendElement({"tag" : "span", "class" : "row"}).down()
    .appendElement({"tag" : "div", "class" : "col-md order-0"}).down()
    .appendElement({"tag" : "div", "class" : "row paddedhor"}).down()
    .appendElement({"tag" : "div", "id" : item.M_id + "-carousel", "class" : "carousel slide d-flex justify-content-center ", "data-ride": "carousel"}).down()
    .appendElement({"tag" : "ol", "class" : "carousel-indicators", "id" : item.M_id + "-carousel-indicators"})
    .appendElement({"tag" : "div", "class" : "carousel-inner rounded", "id" : item.M_id + "-carousel-inner"})
    .appendElement({"tag" : "a", "class" : "carousel-control-prev", "href" : "#" + item.M_id + "-carousel", "role" : "button","data-slide" :"prev" , "id" : item.M_id + "-carousel-control-prev"}).down()
    .appendElement({"tag" : "span", "class" : "carousel-control-prev-icon", "aria-hidden" : "true"})
    .appendElement({"tag" : "span", "class" : "sr-only", "innerHtml" : "Previous"}).up()
    .appendElement({"tag" : "a", "class" : "carousel-control-next", "href" : "#" + item.M_id + "-carousel", "role" : "button","data-slide" :"next" , "id" : item.M_id + "-carousel-control-next"}).down()
    .appendElement({"tag" : "span", "class" : "carousel-control-next-icon", "aria-hidden" : "true"})
    .appendElement({"tag" : "span", "class" : "sr-only", "innerHtml" : "Previous"}).up().up().up().up()
    
    .appendElement({"tag" : "div", "class" : "col-md order-1"}).down()
    .appendElement({"tag" : "div", "class" : "list-group ", "id" : item.M_id + "-detailsGroup"}).down()
    .appendElement({"tag" : "h4", "class" : "row d-flex justify-content-center list-group-item bg-dark text-light", "innerHtml" : "Détails"})
    .appendElement({"tag" : "div", "class" :"row list-group-item itemRow d-flex"}).down()
    .appendElement({"tag" : "div", "class" : "col largeCol", "innerHtml" : "Description", "style" : "border-right :1px solid #e9ecef;"})
    .appendElement({"tag" : "div", "class" : "col largeCol list-group-item-action ", "value" : item.M_description}).up()
    .appendElement({"tag" : "div", "class" :"row list-group-item  itemRow d-flex"}).down()
    .appendElement({"tag" : "div", "class" : "col  largeCol", "innerHtml" : "Alias", "style" : "border-right : 1px solid #e9ecef;"})
    .appendElement({"tag" : "div", "class" : "col  largeCol list-group-item-action", "innerHtml" : item.M_alias}).up()
    .appendElement({"tag" : "div", "class" :"row list-group-item  itemRow d-flex"}).down()
    .appendElement({"tag" : "div", "class" : "col largeCol", "innerHtml" : "Équivalents", "style" : "border-right : 1px solid #e9ecef;"})
    .appendElement({"tag" : "div", "class" : "col  largeCol list-group-item-action", "innerHtml" : item.M_equivalents}).up()
    .appendElement({"tag" : "div", "class" :"row list-group-item  itemRow d-flex"}).down()
    .appendElement({"tag" : "div", "class" : "col largeCol", "innerHtml" : "Commentaires", "style" : "border-right :1px solid #e9ecef;"})
    .appendElement({"tag" : "div", "class" : "col  largeCol list-group-item-action", "innerHtml" : item.M_comments}).up()
    .appendElement({"tag" : "div", "class" :"row list-group-item  itemRow d-flex"}).down()
    .appendElement({"tag" : "div", "class" : "col largeCol", "innerHtml" : "Mots-clés", "style" : "border-right : 1px solid #e9ecef;"})
    .appendElement({"tag" : "div", "class" : "col largeCol  list-group-item-action", "innerHtml" : item.M_keywords}).up().up().up().up()
    
    .appendElement({"tag" : "hr", "class" : "my-4", "style" : "margin:1.5rem;opacity:0;"})
    
    .appendElement({"tag" : "span" , "class" : "row order-2"}).down()
    .appendElement({"tag" : "div", "class" : "col-md"}).down()
    .appendElement({"tag" : "div", "class" : "list-group ", "id" : item.M_id + "-specsGroup"}).down()
    .appendElement({"tag" : "h4", "class" : "row d-flex justify-content-center list-group-item bg-dark text-light", "innerHtml" : "Spécifications"}).up().up().up()
    
    .appendElement({"tag" : "hr", "class" : "my-4", "style" : "margin:1.5rem;opacity:0;"})
    
    .appendElement({"tag" : "div", "id" : item.M_id + "-packages", "class" : "border rounded border-dark"}).down()
    .appendElement({"tag" : "div", "class" : "text-light d-flex justify-content-center align-items-center flex-direction-row", "style"  : "font-size : 2rem; background-color:#868e96;"}).down()
    .appendElement({"tag" : "i", "class" : "fa fa-object-group", "style" : "margin:10px;"})
    .appendElement({"tag" : "span", "innerHtml" : " Paquets"}).up()
    .appendElement({"tag" : "ul", "class" : "nav nav-tabs", "id" : item.M_id + "-packagesNav"})
    
    .appendElement({"tag": "div","class" : "tab-content", "id" : item.M_id  + "-packagesContent"})
    
    this.itemResultPanel.getChildById(item.M_id + "-packageButton").element.addEventListener('click', function(){
        showPackage(item.M_id)
    })
     
    for(var i = 0; i<item.M_properties.length;i++){
        this.itemResultPanel.getChildById(item.M_id + "-specsGroup").appendElement({"tag" : "div", "class" :"row list-group-item  itemRow d-flex"}).down()
        .appendElement({"tag" : "div", "class" : "col largeCol", "innerHtml" : item.M_properties[i].M_name, "style" : "border-right :1px solid #e9ecef;"})
        .appendElement({"tag" : "div", "class" : "col  largeCol list-group-item-action", "innerHtml" : item.M_properties[i].M_stringValue}).up()
    }
    
    for(var i = 0; i< item.M_packaged.length; i++){
         server.entityManager.getEntityByUuid(item.M_packaged[i], false,
        //success callback
        function(item_package,caller){
            caller.pkgNav.appendElement({"tag" : "li", "class" : "nav-item", "id" : item_package.M_id + "-itemli", "style" : "display:flex; flex-direction:row;"}).down()
            .appendElement({"tag" : "span", "class" : "nav-link active", "href" : "#" + item_package.M_id + "-pkgContent", "role" : "tab", "data-toggle" : "tab", "aria-controls" :item_package.M_id + "-pkgContent", "aria-selected" : "true", "id" : item_package.M_id + "-pkgNav", "style" : "display:flex;flex-direction:row;align-items:center;padding:5px;border-radius:0;"}).down()
            .appendElement({"tag":"a","class" : "tabLink", "href" : "#" +item_package.M_id + "-pkgContent", "innerHtml" : item_package.M_id , "style" : "text-decoration : none;"})

            caller.pkgContent.appendElement({"tag" : "div", "class" : "tab-pane show active", "id" : item_package.M_id + "-pkgContent", "role" : "tabpanel", "aria-labelledby" : item_package.M_id + "-pkgNav"}).down()
            .appendElement({"tag" : "div", "class" : "row d-flex justify-content-center"}).down()
            .appendElement({"tag" : "button","id" : item_package.M_id+"-showbtn", "aria-label" : "Show" , "class" : "btn btn-outline-dark", "style" : "margin:10px;"}).down()
             .appendElement({"tag" : "span", "innerHtml" : "Voir le paquet "})
             .appendElement({"tag" : "i", "class" : "fa fa-object-group"}).up().up()
            .appendElement({"tag" : "div", "class" : "list-group", "id" : item_package.M_id + "-packageDetails"}).down()
            .appendElement({"tag" : "h4", "class" : "row d-flex justify-content-center list-group-item bg-dark text-light", "innerHtml" : "Inventaire"})
  
             for(var i=0; i < item_package.M_inventoried.length; i++){
                server.entityManager.getEntityByUuid(item_package.M_inventoried[i], false,
                    function(inventory, caller){
                        var id = (caller.packageID+i+"-place")
                        var temp = document.createElement("div")
                        temp.setAttribute("class","swoop")
                        var child = document.createElement("div")
                        child.setAttribute("class", "row list-group-item itemRow d-flex")
                        var col1 = document.createElement("div")
                        col1.setAttribute("class", "col largeCol")
                        col1.innerHTML = "Emplacement"
                        col1.setAttribute("style", "border-right :1px solid #e9ecef;")
                        var col2 = document.createElement("div")
                        col2.setAttribute("class","col largeCol list-group-item-action d-flex justify-content-center")
                        col2.setAttribute("id",id)
                        col2.appendChild(temp)
                        child.appendChild(col1)
                        child.appendChild(col2)
                        document.getElementById(caller.packageID+"-packageDetails").appendChild(child)
                        var list = new Array()
                       parseLocalisation(inventory.M_located,list, caller.packageID, id )
                       caller.inventoryContainer.appendElement({"tag" : "div", "class" :"row list-group-item itemRow d-flex"}).down()
                        .appendElement({"tag" : "div", "class" : "col largeCol", "innerHtml" : "Quantité en stock", "style" : "border-right : 1px solid #e9ecef;"})
                        .appendElement({"tag" : "div", "class" : "col largeCol  list-group-item-action d-flex justify-content-center", "innerHtml" : inventory.M_quantity}).up()
                        .appendElement({"tag" : "div", "class" :"row list-group-item itemRow d-flex"}).down()
                        .appendElement({"tag" : "div", "class" : "col largeCol", "innerHtml" : "Stock minimal", "style" : "border-right : 1px solid #e9ecef;"})
                        .appendElement({"tag" : "div", "class" : "col largeCol  list-group-item-action d-flex justify-content-center", "innerHtml" : inventory.M_safetyStock}).up()
                        .appendElement({"tag" : "div", "class" :"row list-group-item itemRow d-flex"}).down()
                        .appendElement({"tag" : "div", "class" : "col largeCol", "innerHtml" : "Quantité en commande", "style" : "border-right : 1px solid #e9ecef;"})
                        .appendElement({"tag" : "div", "class" : "col largeCol  list-group-item-action d-flex justify-content-center", "innerHtml" : inventory.M_reorderQty}).up()
                    },
                    function(){
                        
                    }, {"packageID" : item_package.M_id, "i" : i, "inventoryContainer" : caller.pkgContent.getChildById(item_package.M_id + "-packageDetails")})
            }
            
            caller.pkgContent.getChildById(item_package.M_id + "-pkgContent")
            
            .appendElement({"tag" : "hr", "class" : "my-4", "style" : "margin:1.5rem;opacity:0;"})
            .appendElement({"tag" : "div", "class" : "list-group innerShadow", "id" : item_package.M_id + "-suppliersDetails"}).down()
            .appendElement({"tag" : "h4", "class" : "row d-flex justify-content-center list-group-item bg-dark text-light", "innerHtml" : "Fournisseurs"})
            .appendElement({"tag" : "hr", "class" : "my-4", "style" : "margin:1.5rem;opacity:0;"})
            
            
            caller.pkgContent.getChildById(item_package.M_id + "-showbtn").element.addEventListener('click', function(){
                mainPage.packageDisplayPage.displayTabItem(item_package)
            })
    
            
            for(var i = 0; i < item_package.M_supplied.length; i++){
                server.entityManager.getEntityByUuid(item_package.M_supplied[i], false, 
                function(item_supplier, caller) {
                    caller.pkgContent.appendElement({"tag" : "div", "class" : "list-group", "id" : item_supplier.M_id + "-supplierDetails"}).down()
                    .appendElement({"tag" : "div", "class" : "swoop", "style" : "margin: 10px;"})
                    
                    server.entityManager.getEntityByUuid(item_supplier.M_supplier, false, 
                    function(supplier, caller){
                        caller.itemSupplier.element.innerHTML = ""
                        caller.itemSupplier.appendElement({"tag" : "div", "class" : "list-group", "id" : supplier.M_id + "-list"}).down()
                        .appendElement({"tag" : "h4", "class" : "row d-flex justify-content-center list-group-item bg-dark text-light", "innerHtml" : supplier.M_name})
                        .appendElement({"tag" : "div", "class" :"row list-group-item itemRow d-flex"}).down()
                        .appendElement({"tag" : "div", "class" : "col largeCol", "innerHtml" : "ID", "style" : "border-right :1px solid #e9ecef;"})
                        .appendElement({"tag" : "div", "class" : "col largeCol list-group-item-action ", "innerHtml" : caller.supplierInfo.M_id}).up()
                        .appendElement({"tag" : "div", "class" :"row list-group-item itemRow d-flex"}).down()
                        .appendElement({"tag" : "div", "class" : "col largeCol", "innerHtml" : "Prix", "style" : "border-right :1px solid #e9ecef;"})
                        .appendElement({"tag" : "div", "class" : "col largeCol list-group-item-action ", "innerHtml" : caller.supplierInfo.M_price.M_valueOf +" " + caller.supplierInfo.M_price.M_currency.M_valueOf}).up()
                        
                        .appendElement({"tag" : "div", "class" : "row list-group-item itemRow d-flex", "id":supplier.M_id + "-order-line", "style":"visibility: hidden;"}).down()
                        .appendElement({"tag" : "div", "class" : "col largeCol", "innerHtml" : "Quantité à ajouter", "style" : "border-right:1px solid #e9ecef"})
                        .appendElement({"tag" : "div", "class" : "col largeCol list-group-item-action "}).down()
                        .appendElement({"tag" : "div", "class" : "input-group"}).down()
                        
                        .appendElement({"tag" :"input", "type" : "number", "class" : "form-control text-center", "value" : "1", "id" : caller.supplierInfo.M_id + "-qtyToOrder"})
                        .appendElement({"tag" : "div", "class" : "input-group-append"}).down()

                        .appendElement({"tag" : "button", "class" : "btn btn-dark", "type" : "button", "id" : caller.supplierInfo.M_id + "-addButton"}).down()
                        .appendElement({"tag" : "span","innerHtml" : "Ajouter "})
                        .appendElement({"tag" : "i", "class" : "fa fa-cart-plus", "style" : "margin-left:5px;"}).up().up().up().up().up()
                        .appendElement({"tag" : "hr", "class" : "my-4", "style" : "margin:1.5rem;opacity:0;"})             
                        
                        // Here I will the line to order...
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
                    {"itemSupplier" : caller.pkgContent.getChildById(item_supplier.M_id + "-supplierDetails"), "supplierInfo" : item_supplier, "callback":caller.callback})
                    
                },
                function(){}, { "pkgContent" : caller.pkgContent.getChildById(item_package.M_id+ "-suppliersDetails"), "item_package" : item_package, "callback":caller.callback}
                )
        }
        
            
            
        }, 
        //error callback
        function(){
            
        }, {"detailsGroup": this.itemResultPanel.getChildById(item.M_id + "-detailsGroup"), "packageDetails" : this.itemResultPanel.getChildById(item.M_id + "-packageDetails"), "packageId" : item.M_packaged,
            "pkgNav" : this.itemResultPanel.getChildById(item.M_id + "-packagesNav"), "pkgContent" : this.itemResultPanel.getChildById(item.M_id + "-packagesContent"), "callback":callback
        })
        
        parseImage(item.M_id)
    
    }
   
    
    document.getElementById("item_display_page_panel").scrollIntoView()
    
    this.navbar.getChildById(item.M_id+"-closebtn").element.onclick = function(tab,content){
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
    }(this.navbar.getChildById(item.M_id + "-itemli"),  this.itemResultPanel.getChildById(item.M_id+"-itemquery") )
    
    var tab = document.getElementById(item.M_id + "-tab")
            tab.classList.remove("active")
            tab.click()
}



/**
 * Display the item page.
 * @param {*} parent 
 */
 
var ItemDisplayPage = function (parent) {

    // The result content.
    this.panel = new Element(parent, { "tag": "div", "id": "item_display_page_panel", "class": "item_display container-fluid" , "style" :"margin-top : 15px;"})

    this.navbar = this.panel.appendElement({"tag" : "ul", "class" : "nav nav-tabs querynav","id" : "itemqueries"}).down()
    // Now I will display the list of results...
    this.navbar.appendElement({"tag":"button","class":"btn btn-dark toggleBtn", "onclick" : "itemSwitchContext()","innerHtml" :"Voir les recherches"})
    
    this.itemResultPanel = this.panel.appendElement({ "tag": "div", "class": "search_results_content tab-content" }).down()
    this.currentItem = null
    
    return this
}
  
function itemSwitchContext(){
    document.getElementById("item_display_page_panel").style.display = "none"
    document.getElementById("item_search_result_page").style.display = ""
    fireResize()
}

function showPackage(itemID){
    packageElement = document.getElementById(itemID+"-packagesContent")
    packageElement.scrollIntoView({
        behavior : 'smooth'
    })
}
 
function checkImageExists(imagePath, callBack) {
    server.fileManager.readDir(imagePath,
    function(result, caller){
        callBack(result[0].length)
    },function(errObj,caller){
        console.log("error")
    })
}   

function parseImage(itemID){
    checkImageExists("/Catalogue_2/photo/" +itemID + "/", function(existsImage) {
        for(var i =0; i< existsImage; i++) {
            // image exist
            if(i == 0){
                var child = document.createElement("li")
                child.setAttribute("data-target", "#" + itemID + "-carousel")
                child.setAttribute("data-slide-to" , "0")
                child.setAttribute("class", "active")
                document.getElementById(itemID + "-carousel-indicators").appendChild(child)
                child = document.createElement("div")
                child.setAttribute("class", "carousel-item active")
                var picture = document.createElement("img")
                picture.setAttribute("src", "../../Catalogue_2/photo/" + itemID + "/photo_" + Number(i+1) + ".jpg")
                picture.setAttribute("class", "img-fluid rounded carouselImage")
                child.appendChild(picture)
                document.getElementById(itemID + "-carousel-inner").appendChild(child)
            }else{
                var child = document.createElement("li")
                child.setAttribute("data-target", "#"+ itemID + "-carousel")
                child.setAttribute("data-slide-to", i)
                document.getElementById(itemID + "-carousel-indicators").appendChild(child)
                child = document.createElement("div")
                child.setAttribute("class", "carousel-item")
                var picture = document.createElement("img")
                picture.setAttribute("src", "../../Catalogue_2/photo/" + itemID + "/photo_" + Number(i+1) + ".jpg")
                picture.setAttribute("class", "img-fluid rounded carouselImage")
                child.appendChild(picture)
                document.getElementById(itemID + "-carousel-inner").appendChild(child)
            }
            
        }
        
    });
} 
/**
 * Display an item information.
 * @param {*} item 
 */
ItemDisplayPage.prototype.displayTabItem = function (item, callback) {
    
    this.currentItem = item
    
    document.getElementById("item_display_page_panel").style.display = ""
    if(document.getElementById("item_search_result_page") != null){
        document.getElementById("item_search_result_page").style.display = "none"
    }
    
    document.getElementById("package_display_page_panel").style.display = "none"
    if(document.getElementById(item.M_id + "-itemli") !== null){
        document.getElementById(item.M_id + "-itemli").childNodes[0].click()
        return;
    }
    
    this.navbar.appendElement({"tag" : "li", "class" : "nav-item","id" : item.M_id + "-itemli", "style" : "display:flex;flex-direction:row;"}).down()
    .appendElement({"tag" : "span", "class" : "nav-link active", "href" : "#" + item.M_id + "-itemquery", "role" : "tab", "data-toggle" : "tab", "aria-controls" : item.M_id  + "-itemquery","aria-selected" : "true", "id" : item.M_id + "-tab", "style" : "display: flex;flex-direction: row;align-items: center; padding: 5px;border-radius:0;"}).down()
    .appendElement({"tag":"a","class" : "tabLink", "href" : "#" +item.M_id + "-itemquery", "innerHtml" : item.M_id , "style" : "text-decoration : none;"})
    .appendElement({"tag" : "a","id" : item.M_id+"-close", "aria-label" : "Close" , "class" : "tabLink", "id" : item.M_id+"-closebtn"}).down()
    .appendElement({"tag" : "span", "aria-hidden" :"true", "innerHtml" : "&times;", "style" : "padding-left: 5px;color:#007bff"}).up()
    
    this.itemResultPanel.appendElement({"tag" : "div","class" : "tab-pane show result-tab-content", "style" : "position:relative; background-color : white;", "id" : item.M_id + "-itemquery", "role" : "tabpanel", "aria-labelledby" :item.M_id + "-tab" }).down()
    .appendElement({"tag" : "div", "class" : "jumbotron"}).down()
    .appendElement({"tag" : "div", "class": "d-flex justify-content-around"}).down()
    .appendElement({"tag" : "h1", "style" : "text-align:center;", "innerHtml" : item.M_name})
    .appendElement({"tag" : "button", "class" : "btn btn-outline-dark", "id" : item.M_id + "-packageButton"}).down()
    .appendElement({"tag" : "span","style" : "font-size : 1.75rem;","innerHtml" : "Aller aux paquets "})
    .appendElement({"tag" : "i", "class" : "fa fa-object-group","style" : "font-size:1.75rem;"}).up().up()
    .appendElement({"tag" : "hr", "class" : "my-4"})
    .appendElement({"tag" :"div", "class" : "container-fluid"}).down()
    .appendElement({"tag" : "span", "class" : "row"}).down()
    .appendElement({"tag" : "div", "class" : "col-md order-0"}).down()
    .appendElement({"tag" : "div", "class" : "row paddedhor"}).down()
    .appendElement({"tag" : "div", "id" : item.M_id + "-carousel", "class" : "carousel slide d-flex justify-content-center ", "data-ride": "carousel"}).down()
    .appendElement({"tag" : "ol", "class" : "carousel-indicators", "id" : item.M_id + "-carousel-indicators"})
    .appendElement({"tag" : "div", "class" : "carousel-inner rounded", "id" : item.M_id + "-carousel-inner"})
    .appendElement({"tag" : "a", "class" : "carousel-control-prev", "href" : "#" + item.M_id + "-carousel", "role" : "button","data-slide" :"prev" , "id" : item.M_id + "-carousel-control-prev"}).down()
    .appendElement({"tag" : "span", "class" : "carousel-control-prev-icon", "aria-hidden" : "true"})
    .appendElement({"tag" : "span", "class" : "sr-only", "innerHtml" : "Previous"}).up()
    .appendElement({"tag" : "a", "class" : "carousel-control-next", "href" : "#" + item.M_id + "-carousel", "role" : "button","data-slide" :"next" , "id" : item.M_id + "-carousel-control-next"}).down()
    .appendElement({"tag" : "span", "class" : "carousel-control-next-icon", "aria-hidden" : "true"})
    .appendElement({"tag" : "span", "class" : "sr-only", "innerHtml" : "Previous"}).up().up().up().up()
    
    .appendElement({"tag" : "div", "class" : "col-md order-1"}).down()
    .appendElement({"tag" : "div", "class" : "list-group ", "id" : item.M_id + "-detailsGroup"}).down()
    .appendElement({"tag" : "h4", "class" : "row d-flex justify-content-center list-group-item bg-dark text-light", "innerHtml" : "Détails"})
    .appendElement({"tag" : "div", "class" :"row list-group-item itemRow d-flex"}).down()
    .appendElement({"tag" : "div", "class" : "col largeCol", "innerHtml" : "Description", "style" : "border-right :1px solid #e9ecef;"})
    .appendElement({"tag" : "div", "class" : "col largeCol list-group-item-action ", "value" : item.M_description}).up()
    .appendElement({"tag" : "div", "class" :"row list-group-item  itemRow d-flex"}).down()
    .appendElement({"tag" : "div", "class" : "col  largeCol", "innerHtml" : "Alias", "style" : "border-right : 1px solid #e9ecef;"})
    .appendElement({"tag" : "div", "class" : "col  largeCol list-group-item-action", "innerHtml" : item.M_alias}).up()
    .appendElement({"tag" : "div", "class" :"row list-group-item  itemRow d-flex"}).down()
    .appendElement({"tag" : "div", "class" : "col largeCol", "innerHtml" : "Équivalents", "style" : "border-right : 1px solid #e9ecef;"})
    .appendElement({"tag" : "div", "class" : "col  largeCol list-group-item-action", "innerHtml" : item.M_equivalents}).up()
    .appendElement({"tag" : "div", "class" :"row list-group-item  itemRow d-flex"}).down()
    .appendElement({"tag" : "div", "class" : "col largeCol", "innerHtml" : "Commentaires", "style" : "border-right :1px solid #e9ecef;"})
    .appendElement({"tag" : "div", "class" : "col  largeCol list-group-item-action", "innerHtml" : item.M_comments}).up()
    .appendElement({"tag" : "div", "class" :"row list-group-item  itemRow d-flex"}).down()
    .appendElement({"tag" : "div", "class" : "col largeCol", "innerHtml" : "Mots-clés", "style" : "border-right : 1px solid #e9ecef;"})
    .appendElement({"tag" : "div", "class" : "col largeCol  list-group-item-action", "innerHtml" : item.M_keywords}).up().up().up().up()
    
    .appendElement({"tag" : "hr", "class" : "my-4", "style" : "margin:1.5rem;opacity:0;"})
    
    .appendElement({"tag" : "span" , "class" : "row order-2"}).down()
    .appendElement({"tag" : "div", "class" : "col-md"}).down()
    .appendElement({"tag" : "div", "class" : "list-group ", "id" : item.M_id + "-specsGroup"}).down()
    .appendElement({"tag" : "h4", "class" : "row d-flex justify-content-center list-group-item bg-dark text-light", "innerHtml" : "Spécifications"}).up().up().up()
    
    .appendElement({"tag" : "hr", "class" : "my-4", "style" : "margin:1.5rem;opacity:0;"})
    
    .appendElement({"tag" : "div", "id" : item.M_id + "-packages", "class" : "border rounded border-dark"}).down()
    .appendElement({"tag" : "div", "class" : "text-light d-flex justify-content-center align-items-center flex-direction-row", "style"  : "font-size : 2rem; background-color:#868e96;"}).down()
    .appendElement({"tag" : "i", "class" : "fa fa-object-group", "style" : "margin:10px;"})
    .appendElement({"tag" : "span", "innerHtml" : " Paquets"}).up()
    .appendElement({"tag" : "ul", "class" : "nav nav-tabs", "id" : item.M_id + "-packagesNav"})
    
    .appendElement({"tag": "div","class" : "tab-content", "id" : item.M_id  + "-packagesContent"})
    
    this.itemResultPanel.getChildById(item.M_id + "-packageButton").element.addEventListener('click', function(){
        showPackage(item.M_id)
    })
     
    for(var i = 0; i<item.M_properties.length;i++){
        this.itemResultPanel.getChildById(item.M_id + "-specsGroup").appendElement({"tag" : "div", "class" :"row list-group-item  itemRow d-flex"}).down()
        .appendElement({"tag" : "div", "class" : "col largeCol", "innerHtml" : item.M_properties[i].M_name, "style" : "border-right :1px solid #e9ecef;"})
        .appendElement({"tag" : "div", "class" : "col  largeCol list-group-item-action", "innerHtml" : item.M_properties[i].M_stringValue}).up()
    }
    
    for(var i = 0; i< item.M_packaged.length; i++){
         server.entityManager.getEntityByUuid(item.M_packaged[i], false,
        //success callback
        function(item_package,caller){
            caller.pkgNav.appendElement({"tag" : "li", "class" : "nav-item", "id" : item_package.M_id + "-itemli", "style" : "display:flex; flex-direction:row;"}).down()
            .appendElement({"tag" : "span", "class" : "nav-link active", "href" : "#" + item_package.M_id + "-pkgContent", "role" : "tab", "data-toggle" : "tab", "aria-controls" :item_package.M_id + "-pkgContent", "aria-selected" : "true", "id" : item_package.M_id + "-pkgNav", "style" : "display:flex;flex-direction:row;align-items:center;padding:5px;border-radius:0;"}).down()
            .appendElement({"tag":"a","class" : "tabLink", "href" : "#" +item_package.M_id + "-pkgContent", "innerHtml" : item_package.M_id , "style" : "text-decoration : none;"})

            caller.pkgContent.appendElement({"tag" : "div", "class" : "tab-pane show active", "id" : item_package.M_id + "-pkgContent", "role" : "tabpanel", "aria-labelledby" : item_package.M_id + "-pkgNav"}).down()
            .appendElement({"tag" : "div", "class" : "row d-flex justify-content-center"}).down()
            .appendElement({"tag" : "button","id" : item_package.M_id+"-showbtn", "aria-label" : "Show" , "class" : "btn btn-outline-dark", "style" : "margin:10px;"}).down()
             .appendElement({"tag" : "span", "innerHtml" : "Voir le paquet "})
             .appendElement({"tag" : "i", "class" : "fa fa-object-group"}).up().up()
            .appendElement({"tag" : "div", "class" : "list-group", "id" : item_package.M_id + "-packageDetails"}).down()
            .appendElement({"tag" : "h4", "class" : "row d-flex justify-content-center list-group-item bg-dark text-light", "innerHtml" : "Inventaire"})
  
             for(var i=0; i < item_package.M_inventoried.length; i++){
                server.entityManager.getEntityByUuid(item_package.M_inventoried[i], false,
                    function(inventory, caller){
                        var id = (caller.packageID+i+"-place")
                        var child = new Element(document.getElementById(caller.packageID+"-packageDetails"), {"tag":"div", "class":"row list-group-item itemRow d-flex"})
                        var col1 = new Element(child, {"tag":"div", "class":"col largeCol","style":"border-right :1px solid #e9ecef;", "innerHtml":"Emplacement"})
                        var col2 = new Element(child, {"tag":"div", "class":"col largeCol list-group-item-action d-flex justify-content-center", "id":id})
                        //var temp = new Element(col2, {"tag":"div", "class":"swoop"})
                
                        displayLocalisation(col2, inventory.M_located, function(location){
                            mainPage.searchContext = "localisationSearchLnk";
                            mainPage.welcomePage.element.style.display = "none"
                            mainPage.searchResultPage.displayResults({"results":[{"data":location}], "estimate":1}, location.M_id, "localisationSearchLnk") 
                        })
                       
                       caller.inventoryContainer.appendElement({"tag" : "div", "class" :"row list-group-item itemRow d-flex"}).down()
                        .appendElement({"tag" : "div", "class" : "col largeCol", "innerHtml" : "Quantité en stock", "style" : "border-right : 1px solid #e9ecef;"})
                        .appendElement({"tag" : "div", "class" : "col largeCol  list-group-item-action d-flex justify-content-center", "innerHtml" : inventory.M_quantity}).up()
                        .appendElement({"tag" : "div", "class" :"row list-group-item itemRow d-flex"}).down()
                        .appendElement({"tag" : "div", "class" : "col largeCol", "innerHtml" : "Stock minimal", "style" : "border-right : 1px solid #e9ecef;"})
                        .appendElement({"tag" : "div", "class" : "col largeCol  list-group-item-action d-flex justify-content-center", "innerHtml" : inventory.M_safetyStock}).up()
                        .appendElement({"tag" : "div", "class" :"row list-group-item itemRow d-flex"}).down()
                        .appendElement({"tag" : "div", "class" : "col largeCol", "innerHtml" : "Quantité en commande", "style" : "border-right : 1px solid #e9ecef;"})
                        .appendElement({"tag" : "div", "class" : "col largeCol  list-group-item-action d-flex justify-content-center", "innerHtml" : inventory.M_reorderQty}).up()
                    },
                    function(){
                        
                    }, {"packageID" : item_package.M_id, "i" : i, "inventoryContainer" : caller.pkgContent.getChildById(item_package.M_id + "-packageDetails")})
            }
            
            caller.pkgContent.getChildById(item_package.M_id + "-pkgContent")
            
            .appendElement({"tag" : "hr", "class" : "my-4", "style" : "margin:1.5rem;opacity:0;"})
            .appendElement({"tag" : "div", "class" : "list-group innerShadow", "id" : item_package.M_id + "-suppliersDetails"}).down()
            .appendElement({"tag" : "h4", "class" : "row d-flex justify-content-center list-group-item bg-dark text-light", "innerHtml" : "Fournisseurs"})
            .appendElement({"tag" : "hr", "class" : "my-4", "style" : "margin:1.5rem;opacity:0;"})
            
            
            caller.pkgContent.getChildById(item_package.M_id + "-showbtn").element.addEventListener('click', function(){
                mainPage.packageDisplayPage.displayTabItem(item_package)
            })
    
            
            for(var i = 0; i < item_package.M_supplied.length; i++){
                server.entityManager.getEntityByUuid(item_package.M_supplied[i], false, 
                function(item_supplier, caller) {
                    caller.pkgContent.appendElement({"tag" : "div", "class" : "list-group", "id" : item_supplier.M_id + "-supplierDetails"}).down()
                    .appendElement({"tag" : "div", "class" : "swoop", "style" : "margin: 10px;"})
                    
                    server.entityManager.getEntityByUuid(item_supplier.M_supplier, false, 
                    function(supplier, caller){
                        caller.itemSupplier.element.innerHTML = ""
                        caller.itemSupplier.appendElement({"tag" : "div", "class" : "list-group", "id" : supplier.M_id + "-list"}).down()
                        .appendElement({"tag" : "h4", "class" : "row d-flex justify-content-center list-group-item bg-dark text-light", "innerHtml" : supplier.M_name})
                        .appendElement({"tag" : "div", "class" :"row list-group-item itemRow d-flex"}).down()
                        .appendElement({"tag" : "div", "class" : "col largeCol", "innerHtml" : "ID", "style" : "border-right :1px solid #e9ecef;"})
                        .appendElement({"tag" : "div", "class" : "col largeCol list-group-item-action ", "innerHtml" : caller.supplierInfo.M_id}).up()
                        .appendElement({"tag" : "div", "class" :"row list-group-item itemRow d-flex"}).down()
                        .appendElement({"tag" : "div", "class" : "col largeCol", "innerHtml" : "Prix", "style" : "border-right :1px solid #e9ecef;"})
                        .appendElement({"tag" : "div", "class" : "col largeCol list-group-item-action ", "innerHtml" : caller.supplierInfo.M_price.M_valueOf +" " + caller.supplierInfo.M_price.M_currency.M_valueOf}).up()
                        
                        .appendElement({"tag" : "div", "class" : "row list-group-item itemRow d-flex", "id":supplier.M_id + "-order-line", "style":"visibility: hidden;"}).down()
                        .appendElement({"tag" : "div", "class" : "col largeCol", "innerHtml" : "Quantité à ajouter", "style" : "border-right:1px solid #e9ecef"})
                        .appendElement({"tag" : "div", "class" : "col largeCol list-group-item-action "}).down()
                        .appendElement({"tag" : "div", "class" : "input-group"}).down()
                        
                        .appendElement({"tag" :"input", "type" : "number", "class" : "form-control text-center", "value" : "1", "id" : caller.supplierInfo.M_id + "-qtyToOrder"})
                        .appendElement({"tag" : "div", "class" : "input-group-append"}).down()

                        .appendElement({"tag" : "button", "class" : "btn btn-dark", "type" : "button", "id" : caller.supplierInfo.M_id + "-addButton"}).down()
                        .appendElement({"tag" : "span","innerHtml" : "Ajouter "})
                        .appendElement({"tag" : "i", "class" : "fa fa-cart-plus", "style" : "margin-left:5px;"}).up().up().up().up().up()
                        .appendElement({"tag" : "hr", "class" : "my-4", "style" : "margin:1.5rem;opacity:0;"})             
                        
                        // Here I will the line to order...
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
                    {"itemSupplier" : caller.pkgContent.getChildById(item_supplier.M_id + "-supplierDetails"), "supplierInfo" : item_supplier, "callback":caller.callback})
                    
                },
                function(){}, { "pkgContent" : caller.pkgContent.getChildById(item_package.M_id+ "-suppliersDetails"), "item_package" : item_package, "callback":caller.callback}
                )
        }
        
            
            
        }, 
        //error callback
        function(){
            
        }, {"detailsGroup": this.itemResultPanel.getChildById(item.M_id + "-detailsGroup"), "packageDetails" : this.itemResultPanel.getChildById(item.M_id + "-packageDetails"), "packageId" : item.M_packaged,
            "pkgNav" : this.itemResultPanel.getChildById(item.M_id + "-packagesNav"), "pkgContent" : this.itemResultPanel.getChildById(item.M_id + "-packagesContent"), "callback":callback
        })
        
        parseImage(item.M_id)
    
    }
   
    
    document.getElementById("item_display_page_panel").scrollIntoView()
    
    this.navbar.getChildById(item.M_id+"-closebtn").element.onclick = function(tab,content){
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
    }(this.navbar.getChildById(item.M_id + "-itemli"),  this.itemResultPanel.getChildById(item.M_id+"-itemquery") )
    
    var tab = document.getElementById(item.M_id + "-tab")
            tab.classList.remove("active")
            tab.click()
}



