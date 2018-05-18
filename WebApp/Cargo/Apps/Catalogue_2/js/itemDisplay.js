/**
 * Display the item page.
 * @param {*} parent 
 */
 

var ItemDisplayPage = function (parent) {

    // The result content.
    this.panel = new Element(parent, { "tag": "div", "id": "item_display_page_panel", "class": "item_display container-fluid" , "style" :"margin-top : 15px;"})

    this.navbar = this.panel.appendElement({"tag" : "ul", "class" : "nav nav-tabs","id" : "queries"}).down()
    // Now I will display the list of results...
    this.navbar.appendElement({"tag":"button","class":"btn btn-dark toggleBtn", "onclick" : "itemSwitchContext()","innerHtml" :"Voir les recherches"})
    this.itemResultPanel = this.panel.appendElement({ "tag": "div", "class": "search_results_content tab-content" }).down()
    this.currentItem = null
    return this
}
  
function itemSwitchContext(){
    document.getElementById("item_display_page_panel").style.display = "none"
    document.getElementById("item_search_result_page").style.display = ""
}

function showPackage(itemID){
    packageElement = document.getElementById(itemID+"-packagesContent")
    packageElement.scrollIntoView({
        behavior : 'smooth'
    })
}
 
function checkImageExists(imageUrl, callBack) {
    var imageData = new Image();
    imageData.onload = function() {
    callBack(true);
    };
    imageData.onerror = function() {
    callBack(false);
    };
    imageData.src = imageUrl;
}   

function parseImage(itemID, number){
    checkImageExists("../../Catalogue_2/photo/" +itemID + "/photo_"+number+".jpg", function(existsImage) {
        if(existsImage == true) {
            // image exist
            console.log("image exists")
            if(number == 1){
                document.getElementById(itemID + "-carousel-indicators").appendChild({"tag" : "li", "data-target" : "#" + itemID  + "-carousel", "data-slide-to" : "0", "class" : "active"})
            }else{
                document.getElementById(itemID + "-carousel-indicators")
                .appendChild({"tag" : "li", "data-target" : "#" + itemID  + "-carousel", "data-slide-to" : number - 1})
            }
            parseImage(itemID, number+1)
        }
        else {
            console.log("reached the end")
        }
    });
} 
/**
 * Display an item information.
 * @param {*} item 
 */
ItemDisplayPage.prototype.displayTabItem = function (item) {
    
    this.currentItem = item
    
   
    document.getElementById("item_display_page_panel").style.display = ""
    document.getElementById("item_search_result_page").style.display = "none"

     this.navbar.appendElement({"tag" : "li", "class" : "nav-item","id" : item.M_id + "-li", "style" : "display:flex;flex-direction:row;"}).down()
    .appendElement({"tag" : "span", "class" : "nav-link active", "href" : "#" + item.M_id + "-query", "role" : "tab", "data-toggle" : "tab", "aria-controls" : item.M_id  + "-query","aria-selected" : "true", "id" : item.M_id + "-tab", "style" : "display: flex;flex-direction: row;align-items: center; padding: 5px;border-radius:0;"}).down()
    .appendElement({"tag":"a","class" : "tabLink", "href" : "#" +item.M_id + "-query", "innerHtml" : item.M_id , "style" : "text-decoration : none;"})
    .appendElement({"tag" : "a","id" : item.M_id+"-close", "aria-label" : "Close" , "class" : "tabLink", "id" : item.M_id+"-closebtn"}).down()
    .appendElement({"tag" : "span", "aria-hidden" :"true", "innerHtml" : "&times;", "style" : "padding-left: 5px;color:#007bff"}).up()
    
    this.itemResultPanel.appendElement({"tag" : "div","class" : "tab-pane fade show result-tab-content", "style" : "position:relative; background-color : white;", "id" : item.M_id + "-query", "role" : "tabpanel", "aria-labelledby" :item.M_id + "-tab" }).down()
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
    .appendElement({"tag" : "ol", "class" : "carousel-indicators", "id" : item.M_id + "-carousel-indicators"}).down()
    .appendElement({"tag" : "li", "data-target" : "#" + item.M_id  + "-carousel", "data-slide-to" : "0", "class" : "active"})
    .appendElement({"tag" : "li", "data-target" : "#" + item.M_id  + "-carousel", "data-slide-to" : "1"})
    .appendElement({"tag" : "li", "data-target" : "#" + item.M_id  + "-carousel", "data-slide-to" : "2"}).up()
    .appendElement({"tag" : "div", "class" : "carousel-inner rounded", "id" : item.M_id + "-carousel-inner"}).down()
    .appendElement({"tag" : "div", "class" : "carousel-item active"}).down()
    .appendElement({"tag" : "img" , "src" : "../../Catalogue_2/photo/" + item.M_id + "/photo_1.jpg", "class" : "img-fluid rounded"}).up()
    .appendElement({"tag" : "div", "class" : "carousel-item"}).down()
    .appendElement({"tag" : "img" ,  "src" : "../../Catalogue_2/photo/OVMMFO58079/photo_1.jpg", "alt" : "../../Catalogue_2/photo/" + item.M_id + "/photo_1.jpg", "class" : "img-fluid rounded"}).up()
    .appendElement({"tag" : "div", "class" : "carousel-item"}).down()
    .appendElement({"tag" : "img" ,   "src" : "../../Catalogue_2/photo/OVMMSS57283/photo_1.jpg", "alt" : "../../Catalogue_2/photo/" + item.M_id + "/photo_1.jpg", "class" : "img-fluid rounded"}).up().up()
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
            caller.pkgNav.appendElement({"tag" : "li", "class" : "nav-item"}).down()
            .appendElement({"tag" : "span", "class" : "nav-link active", "href" : "#" + item_package.M_id + "-pkgContent", "role" : "tab", "data-toggle" : "tab", "aria-controls" :item_package.M_id + "-pkgContent", "aria-selected" : "true", "id" : item_package.M_id + "-pkgNav", "innerHtml" : item_package.M_id}).down()
           

           
           
           
            caller.pkgContent.appendElement({"tag" : "div", "class" : "tab-pane show active", "id" : item_package.M_id + "-pkgContent", "role" : "tabpanel", "aria-labelledby" : item_package.M_id + "-pkgNav"}).down()
            .appendElement({"tag" : "div", "class" : "list-group", "id" : item_package.M_id + "-packageDetails"}).down()
            .appendElement({"tag" : "h4", "class" : "row d-flex justify-content-center list-group-item bg-dark text-light", "innerHtml" : "Inventaire"})
            .appendElement({"tag" : "div", "class" :"row list-group-item itemRow d-flex"}).down()
            .appendElement({"tag" : "div", "class" : "col largeCol", "innerHtml" : "Quantité en stock", "style" : "border-right : 1px solid #e9ecef;"})
            .appendElement({"tag" : "div", "class" : "col largeCol  list-group-item-action d-flex justify-content-center", "id": item_package.M_id + "-qtyStock"}).down()
            .appendElement({"tag" : "div", "class" : "swoop"}).up().up()
            .appendElement({"tag" : "div", "class" :"row list-group-item itemRow d-flex"}).down()
            .appendElement({"tag" : "div", "class" : "col largeCol", "innerHtml" : "Safety stock", "style" : "border-right : 1px solid #e9ecef;"})
            .appendElement({"tag" : "div", "class" : "col largeCol  list-group-item-action d-flex justify-content-center", "id": item_package.M_id + "-safetyStock"}).down()
            .appendElement({"tag" : "div", "class" : "swoop"}).up().up()
            .appendElement({"tag" : "div", "class" :"row list-group-item itemRow d-flex"}).down()
            .appendElement({"tag" : "div", "class" : "col largeCol", "innerHtml" : "Emplacement", "style" : "border-right : 1px solid #e9ecef;"})
            .appendElement({"tag" : "div", "class" : "col largeCol  list-group-item-action d-flex justify-content-center", "id": item_package.M_id + "-itemLocation"}).down()
            .appendElement({"tag" : "div", "class" : "swoop"}).up().up().up()
            var inventory = new CatalogSchema.InventoryType()
            var location1 = new CatalogSchema.LocalisationType()
            if(item.M_id.startsWith("OVMM")){
                location1.M_name = "Magasin"
            }else{
                location1.M_name = "Toolcrib"
            }
            inventory.M_quantity = 10
            inventory.M_safetyStock = 5
            var location2 = new CatalogSchema.LocalisationType()
            location2.M_name = "34D01"
            var location3 = new CatalogSchema.LocalisationType()
            location3.M_name = "P304"
            location2.M_subLocalisations.push(location3)
            location1.M_subLocalisations.push(location2)
            inventory.M_located = location1
            item_package.M_inventoried.push(inventory);
            caller.pkgContent.getChildById(item.M_id+ "-qtyStock").element.innerHTML = inventory.M_quantity
            caller.pkgContent.getChildById(item.M_id+ "-safetyStock").element.innerHTML = inventory.M_safetyStock
            caller.pkgContent.getChildById(item.M_id  +"-itemLocation").element.innerHTML = ""
            caller.pkgContent.getChildById(item.M_id + "-itemLocation").appendElement({"tag" : "nav", "aria-label" : "breadcrumb"}).down()
            .appendElement({"tag" : "ol", "class" : "breadcrumb", "style" : "background-color : white;"}).down()
            .appendElement({"tag" : "li", "class" : "breadcrumb-item"}).down()
            .appendElement({"tag" : "a", "href" : "#", "innerHtml" : location1.M_name}).up()
            .appendElement({"tag" : "li", "class" : "breadcrumb-item"}).down()
            .appendElement({"tag" : "a", "href" : "#", "innerHtml" : location2.M_name}).up()
            .appendElement({"tag" : "li", "class" : "breadcrumb-item"}).down()
            .appendElement({"tag" : "a", "href" : "#", "innerHtml" : location3.M_name}).up()
            
            caller.pkgContent.getChildById(item_package.M_id + "-pkgContent")
            .appendElement({"tag" : "hr", "class" : "my-4", "style" : "margin:1.5rem;opacity:0;"})
            .appendElement({"tag" : "div", "class" : "list-group", "id" : item_package.M_id + "-suppliersDetails"}).down()
            .appendElement({"tag" : "h4", "class" : "row d-flex justify-content-center list-group-item bg-dark text-light", "innerHtml" : "Fournisseurs"})
            .appendElement({"tag" : "hr", "class" : "my-4", "style" : "margin:1.5rem;opacity:0;"})
            
            for(var i = 0; i < item_package.M_supplied.length; i++){
                server.entityManager.getEntityByUuid(item_package.M_supplied[i], false, 
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
                        .appendElement({"tag" : "hr", "class" : "my-4", "style" : "margin:1.5rem;opacity:0;"})                   
                    }, 
                    function(){},
                    {"itemSupplier" : caller.pkgContent.getChildById(item_supplier.M_id + "-supplierDetails"), "supplierInfo" : item_supplier})
                },
                function(){}, { "pkgContent" : caller.pkgContent.getChildById(item_package.M_id+ "-suppliersDetails"), "item_package" : item_package}
                )
        }
        
            
            
        }, 
        //error callback
        function(){
            
        }, {"detailsGroup": this.itemResultPanel.getChildById(item.M_id + "-detailsGroup"), "packageDetails" : this.itemResultPanel.getChildById(item.M_id + "-packageDetails"), "packageId" : item.M_packaged,
            "pkgNav" : this.itemResultPanel.getChildById(item.M_id + "-packagesNav"), "pkgContent" : this.itemResultPanel.getChildById(item.M_id + "-packagesContent")
        })
        
        parseImage(item.M_id,1)
    
    }
   
    
    
    
    
    
    
    
    
    
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
    }(this.navbar.getChildById(item.M_id + "-li"),  this.itemResultPanel.getChildById(item.M_id+"-query") )
    
    var tab = document.getElementById(item.M_id + "-tab")
            tab.classList.remove("active")
            tab.click()
}



