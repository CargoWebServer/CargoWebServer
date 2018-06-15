var OrderHistory = function(orders, parent){
    this.panel = new Element(parent, { "tag": "div", "id": "order_history_page_panel", "class": "item_display container-fluid" , "style" :"margin-top : 15px;"})
    return this
}

OrderHistory.prototype.displayHistory = function (orders) {
    document.getElementById("main-container").style.display = "none"
    document.getElementById("order_history_page_panel").style.display = ""
    
    if(this.panel.getChildById("orderHistoryContainer") != null){
        return
    }
    this.panel.appendElement({"tag" : "div", "class" : "jumbotron", "id" : "orderHistoryContainer"}).down()
    .appendElement({"tag":"h4", "innerHtml" : "Historique des commandes"})
    .appendElement({"tag" : "table", "class" : "table table-hover table-condensed"}).down()
    .appendElement({"tag" : "thead"}).down()
    .appendElement({"tag" : "tr"}).down()
    .appendElement({"tag" : "th", "style" : "width:20%;", "innerHtml" : "Date de création"})
    .appendElement({"tag" : "th", "style" : "width:20%;", "innerHtml" : "Date d'envoi"})
    .appendElement({"tag" : "th", "style" : "width:20%;", "innerHtml" : "Produits commandés"})
    .appendElement({"tag" : "th", "style" : "width:10%;", "innerHtml" : "Quantité"})
    .appendElement({"tag" : "th", "style" : "width:20%;", "innerHtml" : "Total"})
    .appendElement({"tag" : "th", "style" : "width:10%;", "innerHtml" : "Status"})
    .up().up()
    .appendElement({"tag" : "tbody", "id" : "ordersHistoryBody"}).up()
    
    console.log(orders)
    displayOrders(orders)
}

function displayOrders(orders){
    for(var i = 0; i < orders.length; i++){
        if(orders[i].M_status.M_valueOf != "Open"){
            var creationDate = moment(orders[i].M_creationDate).format('DD/MM/YYYY, h:mm a')
            var completionDate = moment(orders[i].M_completionDate).format('DD/MM/YYYY, h:mm a')
                
            mainPage.orderHistory.panel.getChildById("ordersHistoryBody").appendElement({"tag" : "tr", }).down()
            .appendElement({"tag" :"td", "data-th" : "Date de création"}).down()
            .appendElement({"tag" : "span", "innerHtml" : creationDate}).up()
            .appendElement({"tag" : "td", "data-th" : "Date d'envoi"}).down()
            .appendElement({"tag" : "span", "innerHtml": completionDate}).up()
            .appendElement({"tag" :"td", "data-th" : "Produits commandés"}).down()
            .appendElement({"tag" : "div", "id" : orders[i].UUID+"-items"}).up()
            .appendElement({"tag" :"td", "data-th" : "Quantité"}).down()
            .appendElement({"tag" : "div", "id" : orders[i].UUID+"-quantity"}).up()
            .appendElement({"tag" :"td", "data-th" : "Total"}).down()
            .appendElement({"tag" : "div", "id" : orders[i].UUID+"-total"}).up()
            .appendElement({"tag" :"td", "data-th" : "Status"}).down()
            .appendElement({"tag" : "span", "innerHtml": orders[i].M_status.M_valueOf, "class" : orders[i].M_status.M_valueOf}).up()
            mainPage.orderHistory.panel.getChildById(orders[i].UUID +  "-total").appendElement({"tag" : "div", "innerHtml" : orders[i].M_total.M_valueOf.toFixed(2)})
            
            for(var j = 0; j < orders[i].M_items.length; j++){
                mainPage.orderHistory.panel.getChildById(orders[i].UUID +  "-quantity").appendElement({"tag" : "div", "innerHtml" : orders[i].M_items[j].M_quantity})
                
                server.entityManager.getEntityByUuid(orders[i].M_items[j].M_itemSupplierRef, false, 
                    function(item_supplier, caller){
                         console.log(item_supplier)
                        server.entityManager.getEntityByUuid(item_supplier.M_package, false,
                            function(item_package, caller){
                               
                                mainPage.orderHistory.panel.getChildById(caller.contentID + "-items")
                                .appendElement({"tag" : "div","class" : "simpleLink", "innerHtml" : item_package.M_id, "id" : item_package.M_id + "-searchBtn"})
                            
                                mainPage.orderHistory.panel.getChildById(item_package.M_id + "-searchBtn").element.onclick = function(){
                                    mainPage.packageDisplayPage.displayTabItem(item_package)
                                }
                                
                            }, function () {}, {"contentID" : caller.contentID})
                    }, function() {}, {"contentID" : orders[i].UUID})
            }
        }
    }
}