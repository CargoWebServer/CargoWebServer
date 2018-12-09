/**
 * Display the item page.
 * @param {*} parent 
 */
 
var ItemDisplayPage = function (parent) {

    // The result content.
    this.panel = new Element(parent, { "tag": "div", "id": "item_display_page_panel", "class": "item_display container-fluid" , "style" :"margin-top : 15px;"})

    this.navbar = this.panel.appendElement({"tag" : "ul", "class" : "nav nav-tabs querynav printHide","id" : "itemqueries"}).down()
    // Now I will display the list of results...
    this.itemSwitchContextBtn = this.navbar.appendElement({"tag":"button","class":"btn btn-dark toggleBtn", "innerHtml" :"Voir les recherches"}).down()
    this.itemSwitchContextBtn.element.onclick = function(){
        if(document.getElementById("item_display_page_panel") !== null){
            document.getElementById("item_display_page_panel").style.display = "none"
        }
        if(document.getElementById("item_search_result_page") !== null){
            document.getElementById("item_search_result_page").style.display = ""
        }
        fireResize()
    }
    
    this.itemResultPanel = this.panel.appendElement({ "tag": "div", "class": "search_results_content tab-content" }).down()
    this.currentItem = null
    
    return this
}


function showPackage(itemID){
    packageElement = document.getElementById(itemID+"-packagesContent")
    packageElement.scrollIntoView({
        behavior : 'smooth'
    })
}
 
function parseImage(itemID){
    server.fileManager.readDir("/Catalogue_2/photo/" +itemID,
    function(results, caller){
        var photos = results[0]
        if(photos != null){
            for(var i=0; i < photos.length; i++){
                if(i == 0){
                    if(photos[i].endsWith(".jpg") || photos[i].endsWith(".png")){
                        var child = document.createElement("li")
                        child.setAttribute("data-target", "#" + itemID + "-carousel")
                        child.setAttribute("data-slide-to" , "0")
                        child.setAttribute("class", "active")
                        document.getElementById(itemID + "-carousel-indicators").appendChild(child)
                        child = document.createElement("div")
                        child.setAttribute("class", "carousel-item active")
                        var picture = document.createElement("img")
                        picture.setAttribute("src", "/Catalogue_2/photo/" +itemID+"/" + photos[i])
                        picture.setAttribute("class", "img-fluid rounded carouselImage")
                        child.appendChild(picture)
                        document.getElementById(itemID + "-carousel-inner").appendChild(child)
                    }else{
                        var child = document.createElement("li")
                        child.setAttribute("data-target", "#" + itemID + "-carousel")
                        child.setAttribute("data-slide-to" , "0")
                        child.setAttribute("class", "active")
                        document.getElementById(itemID + "-carousel-indicators").appendChild(child)
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
                        document.getElementById(itemID + "-carousel-inner").appendChild(child)
                    }
                    
                }else{
                    if(photos[i].endsWith(".jpg") || photos[i].endsWith(".png")){
                       var child = document.createElement("li")
                        child.setAttribute("data-target", "#"+ itemID + "-carousel")
                        child.setAttribute("data-slide-to", i)
                        document.getElementById(itemID + "-carousel-indicators").appendChild(child)
                        child = document.createElement("div")
                        child.setAttribute("class", "carousel-item")
                        var picture = document.createElement("img")
                        picture.setAttribute("src", "/Catalogue_2/photo/" +itemID+"/" + photos[i])
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
                        var video = document.createElement("video")
                        video.setAttribute("src", "/Catalogue_2/photo/" +itemID+"/" + photos[i])
                        video.setAttribute("type", "video/mp4")
                        video.setAttribute("class", "img-fluid rounded carouselImage")
                        video.autoplay = true
                        video.muted = true
                        video.controls = true
                        video.load()
                        child.appendChild(video)
                        document.getElementById(itemID + "-carousel-inner").appendChild(child)
                    }
                    
                }
            }
        }
    },function(errObj,caller){
    },{"itemID":itemID})
} 

function parseFirstImage(itemID, div){
    server.fileManager.readDir("Catalogue_2/photo/" + itemID,
        function(results,caller){
            for(var i = 0; i < results.length; i++){
                if(results[i].endsWith(".jpg") || results[i].endsWith(".png")){
                    caller.div.element.src  = "Catalogue_2/photo/" +  caller.itemID  + "/" + results[i]
                    break;
                }
            }
        },function(){}, {"div" : div, "itemID" : itemID})
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
    
    if(document.getElementById("admin_page_panel") != null){
        document.getElementById("admin_page_panel").style.display = "none"
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
    .appendElement({"tag" : "div", "class": "d-flex justify-content-center flex-row"}).down()
    .appendElement({"tag" : "span"}).down()
    .appendElement({"tag" : "h1", "style" : "text-align:center;", "innerHtml" : item.M_name}).up()
    .appendElement({"tag" : "span", "class" : "d-flex align-items-center"}).down()
    .appendElement({"tag" : "button", "class" : "btn btn-outline-dark ml-3", "id" : item.M_id + "-loadItemAdminEdit"}).down()
    .appendElement({"tag" : "span",  "innerHtml" : "Modifier l'item"})
    .appendElement({"tag" : "i", "class" : "fa fa-wrench ml-1"}).up()
    .appendElement({"tag" : "button", "class" : "btn btn-outline-dark ml-3", "id" : item.M_id + "-printBtn"}).down()
    .appendElement({"tag" : "i", "class" : "fa fa-print", "title" : "Imprimer l'item"}).up().up().up()
    .appendElement({"tag" : "hr", "class" : "my-4"})
    .appendElement({"tag" :"div", "class" : "container-fluid"}).down()
    .appendElement({"tag" : "span", "class" : "row"}).down()
    .appendElement({"tag" : "div", "class" : "col-md order-0 printHide"}).down()
    .appendElement({"tag" : "div", "class" : "row paddedhor"}).down()
    .appendElement({"tag" : "div", "id" : item.M_id + "-carousel", "class" : "carousel slide d-flex justify-content-center ", "data-ride": "carousel"}).down()
    .appendElement({"tag" : "ol", "class" : "carousel-indicators printHide", "id" : item.M_id + "-carousel-indicators"})
    .appendElement({"tag" : "div", "class" : "carousel-inner rounded", "id" : item.M_id + "-carousel-inner"})
    .appendElement({"tag" : "a", "class" : "carousel-control-prev printHide", "href" : "#" + item.M_id + "-carousel", "role" : "button","data-slide" :"prev" , "id" : item.M_id + "-carousel-control-prev"}).down()
    .appendElement({"tag" : "span", "class" : "carousel-control-prev-icon", "aria-hidden" : "true"})
    .appendElement({"tag" : "span", "class" : "sr-only", "innerHtml" : "Previous"}).up()
    .appendElement({"tag" : "a", "class" : "carousel-control-next printHide", "href" : "#" + item.M_id + "-carousel", "role" : "button","data-slide" :"next" , "id" : item.M_id + "-carousel-control-next"}).down()
    .appendElement({"tag" : "span", "class" : "carousel-control-next-icon", "aria-hidden" : "true"})
    .appendElement({"tag" : "span", "class" : "sr-only", "innerHtml" : "Previous"}).up().up().up().up()
    
    .appendElement({"tag" : "div", "class" : "col-md order-1"}).down()
    .appendElement({"tag" : "div", "class" : "list-group ", "id" : item.M_id + "-detailsGroup"}).down()
    .appendElement({"tag" : "h4", "class" : "row d-flex justify-content-center list-group-item bg-dark text-light", "innerHtml" : "Détails"})
    .appendElement({"tag" : "div", "class" :"row list-group-item itemRow d-flex"}).down()
    .appendElement({"tag" : "div", "class" : "col largeCol", "innerHtml" : "ID", "style" : "border-right :1px solid #e9ecef;"})
    .appendElement({"tag" : "div", "class" : "col largeCol list-group-item-action ", "innerHtml" : item.M_id}).up()
    .appendElement({"tag" : "div", "class" :"row list-group-item itemRow d-flex"}).down()
    .appendElement({"tag" : "div", "class" : "col largeCol", "innerHtml" : "Description", "style" : "border-right :1px solid #e9ecef;"})
    .appendElement({"tag" : "div", "class" : "col largeCol list-group-item-action ", "innerHtml" : item.M_description}).up()
    .appendElement({"tag" : "div", "class" :"row list-group-item  itemRow d-flex"}).down()
    .appendElement({"tag" : "div", "class" : "col  largeCol", "innerHtml" : "Alias", "style" : "border-right : 1px solid #e9ecef;"})
    .appendElement({"tag" : "div", "class" : "col  largeCol list-group-item-action", "innerHtml" : item.M_alias}).up()
    .appendElement({"tag" : "div", "class" :"row list-group-item  itemRow d-flex"}).down()
    .appendElement({"tag" : "div", "class" : "col largeCol", "innerHtml" : "Équivalents", "style" : "border-right : 1px solid #e9ecef;"})
    .appendElement({"tag" : "div", "class" : "col  largeCol list-group-item-action"}).down()
    .appendElement({"tag" : "ul", "class" : "list-group borderless ", "id": item.M_id + "-equivalent"}).up().up()
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
    .appendElement({"tag" : "div", "class" : "text-light d-flex justify-content-center align-items-center flex-direction-row bg-dark text-light list-group-item printNoBorder printNoPadding", "style"  : "font-size : 1.5rem;"}).down()
    .appendElement({"tag" : "span", "innerHtml" : " Paquets"})
    .appendElement({"tag" : "i", "class" : "fa fa-object-group ml-3"}).up()
    
    .appendElement({"tag": "div","class" : "tab-content", "id" : item.M_id  + "-packageSection"})
    
    parseImage(item.M_id)
    
    
    this.itemResultPanel.getChildById(item.M_id  + "-printBtn").element.onclick = function(){
        window.print()
    }
    
    this.itemResultPanel.getChildById(item.M_id + "-loadItemAdminEdit").element.onclick = function(item){
        return function(){
             mainPage.adminPage.adminItemPage.searchBar.element.value = item.M_id 
            mainPage.adminPage.adminItemPage.searchFct()
             document.getElementById("item_display_page_panel").style.display = "none"
              document.getElementById("admin_page_panel").style.display = ""
              document.getElementById("item-adminTab").click()
        }
    }(item)

    // Display the equivalent list.
    var equivalents = this.itemResultPanel.getChildById(item.M_id + "-equivalent")
    for(var i=0; i < item.M_equivalents.length; i++){
        var uuid = item.M_equivalents[i];
        // So here I will display the equivalency...
        getEntityIdsFromUuid(uuid, function(lst, uuid){
            return function(ids){
                // So here I will set the lnk to the item...
                var li = lst.appendElement({"tag":"li", "class":"list-group-item printNoBorder"}).down()
                var lnk = li.appendElement({"tag":"a", "href":"#", "innerHtml":ids[1]}).down()
                // Now the lnk action...
                lnk.element.onclick = function(uuid){
                    return function(){
                        spinner.panel.element.style.display = "";
                        server.entityManager.getEntityByUuid(uuid, false, 
                            function(entity, caller){
                                mainPage.itemDisplayPage.displayTabItem(entity)
                                spinner.panel.element.style.display = "none";
                            },
                            function(){
                                
                            }, {})
                    }
                }(uuid)
                
            }
        }(equivalents, uuid))   
    }
    
    for(var i = 0; i<item.M_properties.length;i++){
        var div = this.itemResultPanel.getChildById(item.M_id + "-specsGroup").appendElement({"tag" : "div", "class" :"row list-group-item  itemRow d-flex"}).down()
        .appendElement({"tag" : "div", "class" : "col largeCol", "innerHtml" : item.M_properties[i].M_name, "style" : "border-right :1px solid #e9ecef;"})
        var property = item.M_properties[i]
        if(property.M_kind.M_valueOf == "boolean"){
            div.appendElement({"tag" : "div", "class" : "col  largeCol list-group-item-action", "innerHtml" : property.M_booleanValue.toString()})
        }else if(property.M_kind.M_valueOf == "numeric"){
            div.appendElement({"tag" : "div", "class" : "col  largeCol list-group-item-action", "innerHtml" : property.M_numericValue.toString()})
        }else if(property.M_kind.M_valueOf == "dimension"){
            // So in that particular case I will display the dimension type information.
            var cell = div.appendElement({"tag" : "div", "class" : "col  largeCol list-group-item-action"}).down()
            cell.appendElement({"tag":"span", "innerHtml":property.M_dimensionValue.M_valueOf.toString()})
            cell.appendElement({"tag":"span", "innerHtml":property.M_dimensionValue.M_unitOfMeasure.M_valueOf, "style":"padding-left: 4px;"})
        }else{
            div.appendElement({"tag" : "div", "class" : "col  largeCol list-group-item-action", "innerHtml" : item.M_properties[i].M_stringValue})
        }
    }
    
    
    for(var i  = 0; i < item.M_packaged.length; i++){
       server.entityManager.getEntityByUuid(item.M_packaged[i], false, 
       function(pkg, caller){
           caller.packageSection.appendElement({"tag" : "div", "class" :"list-group-item row itemRow d-flex innerShadow printNoBorder", "id" : pkg.M_id + "-itemRow", "style" : "margin-left : 0px; margin-right:0px;"}).down()
           .appendElement({"tag" : "div", "class" : "col-md-3 d-flex justify-content-center align-items-center"}).down()
           .appendElement({"tag": "img","class" : "img-fluid rounded printHide", "style" : "margin:10px;", "id" : pkg.M_id + "-image"}).up()
           .appendElement({"tag" : "div", "class" : "col-md largeCol"}).down()
           .appendElement({"tag" : "div", "class" : "list-group"}).down()
           .appendElement({"tag" : "div", "class" : "row list-group-item itemRow d-flex"}).down()
           .appendElement({"tag" : "div", "class" : "col largeCol", "innerHtml" : "ID"})
           .appendElement({"tag" : "div", "class" : "col largeCol list-group-item-action", "innerHtml" : pkg.M_id}).up()
           .appendElement({"tag" : "div", "class" : "row list-group-item itemRow d-flex"}).down()
           .appendElement({"tag" : "div", "class" : "col largeCol", "innerHtml" : "Nom"})
           .appendElement({"tag" : "div", "class" : "col largeCol list-group-item-action", "innerHtml" : pkg.M_name}).up()
           .appendElement({"tag" : "div", "class" : "row list-group-item itemRow d-flex"}).down()
           .appendElement({"tag" : "div", "class" : "col largeCol", "innerHtml" : "Description"})
           .appendElement({"tag" : "div", "class" : "col largeCol  list-group-item-action", "innerHtml" : pkg.M_description}).up().up().up()
           .appendElement({"tag" : "div", "class" : "col-md-2 largeCol d-flex justify-content-center align-items-center"}).down()
           .appendElement({"tag" : "button", "class" : "btn btn-outline-dark", "id" : pkg.M_id + "-showPackageButton"}).down()
           .appendElement({"tag" : "span", "innerHtml" : "Voir le paquet"})
           .appendElement({"tag" : "i", "class" : "fa fa-object-ungroup", "style" : "margin:10px;"}).up()
           
           parseFirstImage(pkg.M_id, caller.packageSection.getChildById(pkg.M_id + "-image"))
           
        caller.packageSection.getChildById(pkg.M_id + "-showPackageButton").element.addEventListener('click', function(){
            mainPage.packageDisplayPage.displayTabItem(pkg)
        })
          
           
       }, function (){
           
       }, {"packageSection" : this.itemResultPanel.getChildById(item.M_id + "-packageSection")})
   }
    
    /*for(var i = 0; i< item.M_packaged.length; i++){
         server.entityManager.getEntityByUuid(item.M_packaged[i], false,
        //success callback
        function(item_package,caller){
            
            caller.pkgNav.appendElement({"tag" : "li", "class" : "nav-item", "id" : item_package.M_id + "-itemli", "style" : "display:flex; flex-direction:row;"}).down()
            .appendElement({"tag" : "a", "class" : "nav-link", "href" : "#" + item_package.M_id + "-pkgContent", "role" : "tab", "data-toggle" : "tab", "aria-controls" :item_package.M_id + "-pkgContent", "aria-selected" : "true", "id" : item_package.M_id + "-pkgNav", "style" : "display:flex;flex-direction:row;align-items:center;padding:5px;border-radius:0;", "innerHtml" : item_package.M_id})
           

            caller.pkgContent.appendElement({"tag" : "div", "class" : "tab-pane show", "id" : item_package.M_id + "-pkgContent", "role" : "tabpanel", "aria-labelledby" : item_package.M_id + "-pkgNav"}).down()
            .appendElement({"tag" : "div", "class" : "row d-flex justify-content-center"}).down()
            .appendElement({"tag" : "button","id" : item_package.M_id+"-showbtn", "aria-label" : "Show" , "class" : "btn btn-outline-dark", "style" : "margin:10px;font-size:30px;"}).down()
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
    
    }*/
   
    
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