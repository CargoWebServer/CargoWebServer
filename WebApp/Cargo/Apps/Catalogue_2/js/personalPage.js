/** 
 * That page display the history of order for a given user.
 */
 var PersonalPage = function(userId, orders){
    this.userId = userId;
    
    // The list of orders.
    mainPage.welcomePage.removeAllChilds()
    this.panel = mainPage.welcomePage;
    
    // Now I will display the personal page...
    this.panel.appendElement({"tag":"div", "class":"jumbotron bg-dark text-light border border-dark", "style":"margin: 20px; padding : 2rem 1rem;text-align:center;" }).down()
    .appendElement({"tag" : "h1", "class" : "display-5", "id" : "personal_welcome_msg", "style" : "text-align : center;"})
    .appendElement({"tag" : "hr", "class" : "my-4 dark"})
    .appendElement({"tag" : "p", "class" : "lead"}).down()
    .appendElement({"tag" : "span","id":"welcomeTipText1"})
    .appendElement({"tag" : "li" ,"class": "btn btn-outline-light disabled", "style":"opacity : 1; font-size : 0.8rem;cursor:default;"}).down()
    .appendElement({"tag" : "i", "class" :"fa fa-navicon"})
    .appendElement({"tag" : "span","id":"quickSearchText"}).up()
    .appendElement({"tag" : "span","id":"welcomeTipText2"})
    .appendElement({"tag":"li", "class": "btn btn-outline-light disabled", "style":"opacity :1;font-size : 0.8rem;cursor:default;"}).down()
    .appendElement({"tag":"i", "class":"fa fa-angle-double-down"})
    .appendElement({"tag" :"span","id":"advancedSearchText"}).up()
    .appendElement({"tag" : "span","id":"welcomeTipText3"})
    .appendElement({"tag" : "span", "innerHtml" : " Vous pouvez accéder en tout temps à votre panier de commande avec l'icône "})
    
    .appendElement({"tag" : "li", "class" : "btn btn-outline-light disabled", "id" : "orderBtn", "style" : "cursor:default;"}).down()
    .appendElement({"tag" : "i", "class" : "fa fa-shopping-cart"}).up().up()
    .appendElement({"tag" : "p", "style" : "text-align:center;"}).down()
    .appendElement({"tag" : "hr", "class" : "my-4"}).up()
    .appendElement({"tag" : "div", "class" : "d-flex justify-content-center"}).down()
    .appendElement({"tag" : "button", "class" : "btn btn-light", "onclick" : "goToSearches()", "style":"margin-right:10px;"}).down()
    .appendElement({"tag"  : "span", "innerHtml" : "Aller aux recherches "})
    .appendElement({"tag" : "i", "class" : "fa fa-search text-dark"}).up()

    .appendElement({"tag" : "button", "class" : "btn btn-light","style" : "margin-left:10px;", "onclick" : "goToBasket()"}).down()
    .appendElement({"tag"  : "span", "innerHtml" : "Voir le panier "})
    .appendElement({"tag" : "i", "class" : "fa fa-shopping-cart text-dark"}).up().up().up()
    
    .appendElement({"tag" : "hr", "class" : "my-4", "style" : "opacity:0;"})
    
    .appendElement({"tag" : "div", "class" : "row", "style" : "margin-bottom:20px;"}).down()
    .appendElement({"tag" : "div", "class" : "col", "style" : "margin-left:20px;"}).down()
    .appendElement({"tag" : "div", "class" : "card"}).down()
    .appendElement({"tag" : "div", "class" : "card-body bg-dark text-light d-flex flex-row justify-content-between", "style": "border-top-right-radius:.25rem;border-top-left-radius:.25rem;"}).down()
    .appendElement({"tag" : "h5", "class"  :"card-title", "innerHtml" : "Commandes en attente"})
    .appendElement({"tag" : "button", "class" : "btn btn-light", "id" : "showHistoryBtn"}).down()
    .appendElement({"tag" : "span", "innerHtml" : "Voir l'historique "})
    .appendElement({"tag" : "i", "class" : "fa fa-book"}).up().up()
    
    .appendElement({"tag" : "table", "class" : "table table-hover table-condensed"}).down()
    .appendElement({"tag" : "thead"}).down()
    .appendElement({"tag" : "tr"}).down()
    .appendElement({"tag" : "th", "style" : "width:30%;", "innerHtml" : "Date de création"})
    .appendElement({"tag" : "th", "style" : "width:30%;", "innerHtml" : "Date d'envoi"})
    .appendElement({"tag" : "th", "style" : "width:30%;", "innerHtml" : "Produits commandés"})
    .appendElement({"tag" : "th", "style" : "width:10%;", "innerHtml" : ""})
    .up().up()
    
    .appendElement({"tag" : "tbody", "id" : "personalOrdersBody"}).up()
    
    .up().up()
    .appendElement({"tag" : "hr", "class" : "mx-4"})
    .appendElement({"tag" : "div", "class" : "col", "style" : "margin-right:20px;"}).down()
    .appendElement({"tag" : "div", "class" : "card"}).down()
    .appendElement({"tag" : "div", "class" : "card-body bg-dark text-light", "style": "border-top-right-radius:.25rem;border-top-left-radius:.25rem;"}).down()
    .appendElement({"tag" : "h5", "class"  :"card-title", "innerHtml" : "Produits les plus commandés"}).up()
    
    .appendElement({"tag" : "table", "class" : "table table-hover table-condensed"}).down()
    .appendElement({"tag" : "thead"}).down()
    .appendElement({"tag" : "tr"}).down()
    .appendElement({"tag" : "th", "style" : "width:20%;", "innerHtml" : "ID"})
    .appendElement({"tag" : "th", "style" : "width:40%;", "innerHtml" : "Nom"})
    .appendElement({"tag" : "th", "style" : "width:30%;", "innerHtml" : "Type"})
    .appendElement({"tag" : "th", "style" : "width:10%;", "innerHtml" : ""})
    .up().up()
    
    .appendElement({"tag" : "tbody", "id" : "personalItemsBody"}).up()
    
    .up().up()
    
    mainPage.panel.getChildById("main").appendElement({"tag" : "div", "class" : "modal", "id" : "orderModal", "tabindex" : "-1", "role" : "dialog", "aria-labelledby" : "orderModal", "aria-hidden" : "true"}).down()
    .appendElement({"tag" : "div", "class" : "modal-dialog modal-dialog-centered", "role" : "document"})
    .appendElement({"tag" : "div", "class" : "modal-content"}).down()
    .appendElement({"tag" : "div", "class" : "modal-header"}).down()
    .appendElement({"tag" : "h5", "id" : "orderNameTitle", "style" : "display: flex;justify-content: space-around;width: 100%;"})
    .appendElement({"tag" : "button", "type" : "button", "class" : "close", "data-dismiss" : "modal", "aria-label" : "Close","id" : "orderModalClose"}).down()
    .appendElement({"tag" : "span", "aria-hidden" : "true", "innerHtml" : "&times"}).up().up()
    .appendElement({"tag" : "table", "class" : "table table-hover table-condensed", "id" : "orderModalContent"})
    
    server.entityManager.getEntityById("CargoEntities.User", "CargoEntities", [userId], false, 
        function(user, caller){
            caller.getChildById("personal_welcome_msg").element.innerHTML = "Bienvenue " + user.M_firstName + " " + user.M_lastName;
            
        },
        function(){
            
        }, this.panel)
        
    
    this.panel.getChildById("showHistoryBtn").element.onclick = function(orders){
        return function(){
            mainPage.orderHistory.displayHistory(orders)
        }
        
    }(orders)
    
    // display the page.
    this.displayOrders(orders)
    
    return this;
 }
 
PersonalPage.prototype.displayOrders = function(orders){
   
    
    orders.sort(function(a,b){
        
        return new Date(b.M_completionDate) - new Date(a.M_completionDate)
    })
    for(var i = 0; i < orders.length ; i++){

        if(orders[i].M_status.M_valueOf === "Pending"){
           
            var creationDate = moment(orders[i].M_creationDate).format('DD/MM/YYYY, h:mm a')
            var completionDate = moment(orders[i].M_completionDate).format('DD/MM/YYYY, h:mm a')
            
            mainPage.welcomePage.getChildById("personalOrdersBody").appendElement({"tag" : "tr", }).down()
            .appendElement({"tag" :"td", "data-th" : "Date de création"}).down()
            .appendElement({"tag" : "span", "innerHtml" : creationDate}).up()
            .appendElement({"tag" : "td", "data-th" : "Date d'envoi"}).down()
            .appendElement({"tag" : "span", "innerHtml": completionDate}).up()
            .appendElement({"tag" :"td", "data-th" : "Produits commandés"}).down()
            .appendElement({"tag" : "span", "innerHtml": orders[i].M_items.length.toString()}).up()
            .appendElement({"tag" :"td", "data-th" : ""}).down()
            .appendElement({"tag" : "button", "class" : "btn btn-outline-dark","data-toggle": "modal","data-target" : "#orderModal", "id" : orders[i].M_completionDate + "-searchBtn"}).down()
            .appendElement({"tag" : "i", "class" : "fa fa-search-plus", "style" : "color:inherit;"}).up()
            
            mainPage.welcomePage.getChildById(orders[i].M_completionDate+"-searchBtn").element.onclick = function(order){
                return function() {
                    displayOrderPreview(order)
                }
                
            }(orders[i])
        }
       
  
        server.entityManager.attach(this, UpdateEntityEvent, function (evt, codeEditor) {
        if (evt.dataMap.entity !== undefined) {
            var file = evt.dataMap.entity;
            if(file.TYPENAME == "CatalogSchema.OrderType"){
   
                server.entityManager.getEntityByUuid(file.UUID, false, 
                    function(order, caller){
                        caller.orders.push(order)
                        updatePopularItems(orders, 10)
                            
                }, function(){}, {"orders" : orders})
               
               var creationDate = moment(file.M_creationDate).format('DD/MM/YYYY, h:mm a')
               var completionDate = moment(file.M_completionDate).format('DD/MM/YYYY, h:mm a')
            
                mainPage.welcomePage.getChildById("personalOrdersBody").appendElement({"tag" : "tr", }).down()
                    .appendElement({"tag" :"td", "data-th" : "Date de création"}).down()
                    .appendElement({"tag" : "span", "innerHtml" : creationDate}).up()
                    .appendElement({"tag" : "td", "data-th" : "Date d'envoi"}).down()
                    .appendElement({"tag" : "span", "innerHtml": completionDate}).up()
                    .appendElement({"tag" :"td", "data-th" : "Produits commandés"}).down()
                    .appendElement({"tag" : "span", "innerHtml": file.M_items.length.toString()}).up()
                    .appendElement({"tag" :"td", "data-th" : ""}).down()
                    .appendElement({"tag" : "button", "class" : "btn btn-outline-dark", "onclick" : ""}).down()
                    .appendElement({"tag" : "i", "class" : "fa fa-search-plus", "style" : "color:inherit;"}).up()
                
                }
            }
        });
        
        
        
    }
    
    updatePopularItems(orders, 10)
    
    
 }
 
 
 
// Take an array of objects, and return two arrays: one with the different objects, and one with the indexes at which they can be found
function arraySort(arr){
    arr.sort()
    var a = [], b = [], prev;
    for ( var i = 0; i < arr.length; i++ ) {
        if ( arr[i] !== prev ) {
            a.push(arr[i]);
            b.push(1);
        } else {
            b[b.length-1]++;
        }
        prev = arr[i];
        
    }
    return [a,b]
}
 
 
//Sorts an array of objects by their occurence, and adds them to a map 
function mapObjectOccurences(arr){
    var sortedItems = arraySort(arr)

    var map = {}
    
    for(var i = 0; i<sortedItems[0].length; i++){
        if( map[sortedItems[1][i]] == undefined){
             map[sortedItems[1][i]] = [sortedItems[0][i]]
        }else{
             map[sortedItems[1][i]].push(sortedItems[0][i])
        }
       
    }
    return map
 }
 
 
//With an object occurence map and their keys array, draws the most popular items
function drawPopularItems(keysArray, map, amountToShow){
    mainPage.welcomePage.getChildById("personalItemsBody").removeAllChilds()
    for(var i = keysArray.length - 1; i >= 0 && i>=keysArray.length - amountToShow ; i--){
        for(var j = 0; j < map[keysArray[i]].length; j++){
            var itemSupplier = map[keysArray[i]][j]
             server.entityManager.getEntityByUuid(itemSupplier, false, 
                function(item_supplier,caller){
                    server.entityManager.getEntityByUuid(item_supplier.M_package,false,
                        function(item_package, caller){
                            mainPage.welcomePage.getChildById("personalItemsBody").appendElement({"tag" : "tr", }).down()
                            .appendElement({"tag" :"td", "data-th" : "ID"}).down()
                            .appendElement({"tag" : "span", "innerHtml" : item_package.M_id}).up()
                            .appendElement({"tag" : "td", "data-th" : "Nom"}).down()
                            .appendElement({"tag" : "span", "innerHtml": item_package.M_name}).up()
                            .appendElement({"tag" :"td", "data-th" : "Type"}).down()
                            .appendElement({"tag" : "span", "innerHtml": ""}).up()
                            .appendElement({"tag" :"td", "data-th" : ""}).down()
                            .appendElement({"tag" : "button", "class" : "btn btn-outline-dark", "id" : item_package.M_id + "-searchBtn"}).down()
                            .appendElement({"tag" : "i", "class" : "fa fa-search-plus", "style" : "color:inherit;"}).up()
                            
                            mainPage.welcomePage.getChildById(item_package.M_id + "-searchBtn").element.onclick = function(){
                                mainPage.packageDisplayPage.displayTabItem(item_package)
                            }
                        }, function() {}, {})
                },
                function(){
                    
                },{})
                
                    
                
        }
        
           
    }
}
 
 
//Updates the popular items map, and draws them inside the container 
function updatePopularItems(orders, amountToShow){
    var arr = []
    for(var i = 0; i < orders.length; i++){
        for(var j =0; j<orders[i].M_items.length;j++){
            arr.push(orders[i].M_items[j].M_itemSupplierRef)
        }
    }
    
    var map = mapObjectOccurences(arr)
    var keysArray = Object.keys(map)
    keysArray.sort(function(a,b){
        return a-b
    })
    drawPopularItems(keysArray, map, amountToShow)
 }
 
 
 function displayOrderPreview(order){
    mainPage.panel.getChildById("orderModalContent").element.innerHTML = ""
    mainPage.panel.getChildById("orderNameTitle").element.innerHTML = ""
    mainPage.panel.getChildById("orderNameTitle").appendElement({"tag":"span", "innerHtml" : moment(order.M_completionDate).format('DD/MM/YYYY, h:mm a')})
    .appendElement({"tag" : "span","class" :  order.M_status.M_valueOf, "innerHtml" : "Status : " + order.M_status.M_valueOf})
    mainPage.panel.getChildById("orderModalContent").appendElement({"tag" : "thead"}).down()
        .appendElement({"tag" : "tr"}).down()
        .appendElement({"tag" : "th", "style" : "width:50%;", "innerHtml" : "Produit"})
        .appendElement({"tag" : "th", "style" : "width:10%;", "innerHtml" : "Prix"})
        .appendElement({"tag" : "th", "style" : "width:8%;", "innerHtml" : "Quantité"})
        .appendElement({"tag" : "th", "style" : "width:22%;", "class" : "text-center", "innerHtml" : "Sous-total"})
        .appendElement({"tag" : "th", "style" : "width:10%;"}).up().up()
        .appendElement({"tag" : "tbody", "id" : "orderBodyModal"})
        
        
    for(var i = 0; i < order.M_items.length; i++){
        
        server.entityManager.getEntityByUuid(order.M_items[i].M_itemSupplierRef, false, 
        function(item_supplier, caller){
            server.entityManager.getEntityByUuid(item_supplier.M_package, false, 
                function(item_package,caller){
   
                         mainPage.panel.getChildById("orderBodyModal").appendElement({"tag" : "tr", "id" : caller.item_supplier.M_id  + "-productRow"}).down()
                        .appendElement({"tag" :"td", "data-th" : "Produit"}).down()
                        .appendElement({"tag" : "div", "class" : "row"}).down()
                        .appendElement({"tag" : "div", "class" : "col-sm-2 d-none d-sm-block"}).down()
                        .appendElement({"tag" : "img", "src" : "../../Catalogue_2/photo/" + item_package.M_id + "/photo_1.jpg", "class" : "img-fluid rounded"}).up()
                        .appendElement({"tag" : "div", "class" : "col-sm-10"}).down()
                        .appendElement({"tag" :"h4", "class" : "nomargin", "innerHtml" : item_package.M_id})
                        .appendElement({"tag" : "p", "innerHtml" : item_package.M_name}).up().up().up()
                        .appendElement({"tag" : "td", "data-th" : "Prix"}).down()
                        .appendElement({"tag" : "span", "innerHtml":caller.item_supplier.M_price.M_valueOf})
                        .appendElement({"tag" : "span","innerHtml" : caller.item_supplier.M_price.M_currency.M_valueOf, "style" : "margin-left:5px;"}).up()
                        .appendElement({"tag" :"td", "data-th" : "Quantité"}).down()
                        .appendElement({"tag" :"div", "type" : "number", "class" : "form-control text-center", "innerHtml" : caller.newOrderLine.M_quantity, "id" : caller.item_supplier.M_id + "-qtyToOrder"}).up()
                        .appendElement({"tag" : "td", "data-th" : "Sous-total", "class" : "text-center"}).down()
                        .appendElement({"tag" : "span","innerHtml" : (caller.item_supplier.M_price.M_valueOf * caller.newOrderLine.M_quantity).toFixed(2), "id" : caller.item_supplier.M_id + "-subtotal"})
                        .appendElement({"tag" : "span", "innerHtml" : caller.item_supplier.M_price.M_currency.M_valueOf, "style" : "margin-left: 5px;"}).up()
                        .appendElement({"tag":"td", "class" : "actions", "data-th" : ""}).down()
                        .appendElement({"tag" : "button", "class" : "btn btn-outline-dark", "id" : item_package.M_id + "-modalSearchBtn"}).down()
                        .appendElement({"tag" : "i", "class" : "fa fa-search-plus", "style" : "color:inherit;"}).up()
                        
                        mainPage.panel.getChildById(item_package.M_id + "-modalSearchBtn").element.onclick = function(){
                                mainPage.panel.getChildById("orderModalClose").element.click()
                                mainPage.packageDisplayPage.displayTabItem(item_package)
                        }
                        
                        
                        
                       
                    
                }, function () {}
                ,{"item_supplier" : item_supplier, "newOrderLine" : caller.newOrderLine})
        },
        function (){
            
        }, {"newOrderLine" : order.M_items[i]})
    }
 }