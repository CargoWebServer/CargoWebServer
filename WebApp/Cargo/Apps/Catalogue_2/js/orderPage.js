var OrderPage = function (parent) {

    // The result content.
    this.panel = new Element(parent, { "tag": "div", "id": "order_page_panel", "class": "item_display container-fluid" , "style" :"margin-top : 15px;"})
    
    this.currentOrder = null
    

    
    if(this.panel.getChildById("orderPanel") == null){
         this.panel.appendElement({"tag" : "div", "class" : "container-fluid", "id" : "orderPanel"}).down()
        .appendElement({"tag" : "div", "class" : "jumbotron"}).down()
        .appendElement({"tag" : "h1", "innerHtml" : "Panier de commande"})
        .appendElement({"tag" : "table", "class" : "table table-hover table-condensed"}).down()
        .appendElement({"tag" : "thead"}).down()
        .appendElement({"tag" : "tr"}).down()
        .appendElement({"tag" : "th", "style" : "width:50%;", "innerHtml" : "Produit"})
        .appendElement({"tag" : "th", "style" : "width:10%;", "innerHtml" : "Prix"})
        .appendElement({"tag" : "th", "style" : "width:8%;", "innerHtml" : "Quantité"})
        .appendElement({"tag" : "th", "style" : "width:22%;", "class" : "text-center", "innerHtml" : "Sous-total"})
        .appendElement({"tag" : "th", "style" : "width:10%;"}).up().up()
        
        .appendElement({"tag" : "tbody", "id" : "orderBody"})
        
        .appendElement({"tag" : "tfoot"}).down()
        .appendElement({"tag" : "tr", "class" : "d-table-row d-sm-none"}).down()
        .appendElement({"tag" : "td", "class" : "text-center"}).down()
        .appendElement({"tag" : "strong", "innerHtml" : "Total"}).up().up()
        .appendElement({"tag" : "tr"}).down()
        .appendElement({"tag" : "td"}).down()
        .appendElement({"tag" : "a", "onclick" : "goToHome()", "class" : "btn btn-outline-dark"}).down()
        .appendElement({"tag" : "i", "class" : "fa fa-angle-left"})
        .appendElement({"tag" : "span", "innerHtml" : " Retour à la page d'accueil"}).up().up()
        .appendElement({"tag" : "td", "colspan" : "2", "class" : "d-none d-sm-table-cell"})
        .appendElement({"tag" : "td", "class" : "d-none d-sm-table-cell text-center"}).down()
        .appendElement({"tag" : "strong", "innerHtml" : "Total "})
        .appendElement({"tag" : "strong", "id" : "orderTotal"}).up()
        .appendElement({"tag" : "td"}).down()
        .appendElement({"tag" :"a", "class" : "btn btn-dark text-light btn-block", "id":"send_cmd_btn"}).down()
        .appendElement({"tag" : "span" , "innerHtml" : "Envoyer la commande "})
        .appendElement({"tag" : "i", "class" : "fa fa-cart-arrow-down"})
        
        this.sendCmdBtn = this.panel.getChildById("send_cmd_btn")
        
        this.sendCmdBtn.element.onclick = function(orderPage){
            return function(){
                
                // save the current order.
                // console.log(orderPage.currentOrder)
                // When the order is save I will remove the current order
                if(orderPage.currentOrder.M_items.length == 0){
                    mainPage.showNotification("danger", "Votre panier est vide!", 4000)
                }else{
                    document.getElementById("waitingDiv").style.display = ""
                    orderPage.currentOrder.M_status.M_valueOf = "Pending";
                    server.entityManager.saveEntity(orderPage.currentOrder, 
                    function(result, orderPage){
                        localStorage.removeItem(server.account.UUID + "_currentOrder");
                        orderPage.panel.getChildById("orderBody").removeAllChilds()
                        orderPage.panel.getChildById("orderTotal").element.innerHTML = ""
                        var userId = orderPage.currentOrder.M_userId;
                        orderPage.currentOrder = null;
                        mainPage.showNotification("success", "Votre commande a été créée avec succès!", 4000)
                        mainPage.panel.getChildById("cartCount").element.innerHTML = 0
                        
                        // Recreate a new empty order.
                        var order = new CatalogSchema.OrderType()
                        order.M_userId = userId
                        order.M_creationDate = new Date()
                        order.M_status = new CatalogSchema.OrderStatus()
                        order.M_status.M_valueOf = "Open";
                        order.M_total = new CatalogSchema.PriceType()
                        order.M_total.M_valueOf = 0
                        document.getElementById("waitingDiv").style.display = "none"
                        server.entityManager.createEntity(catalog.UUID, "M_orders", order, 
                            function (order,orderPage){
                                
                                orderPage.setOrder(order)
                            },function(errObj,caller){
                                
                            },orderPage)
                        
                    },
                    function(){
                        
                    }, orderPage);
                }
                

            }
        }(this)
    }
    
    return this
}

// Save the order temporaly in the cache.
OrderPage.prototype.saveOrder = function (order) {
    if(this.currentOrder !== null){
        localStorage.setItem(server.account.UUID + "_currentOrder", order.stringify());
    }
}

// Get the order from the cache.
OrderPage.prototype.getOrder = function () {
    if(localStorage.getItem(server.account.UUID + "_currentOrder") !== null){
        var obj = JSON.parse(localStorage.getItem(server.account.UUID + "_currentOrder"))
        var order = new CatalogSchema.OrderType()
        setObjectValues(order, obj, true)
        return order
    }
    
    return null
}

OrderPage.prototype.displayOrder = function () {
    document.getElementById("admin_page_panel").style.display = "none"
    document.getElementById("main-container").style.display = "none"
    document.getElementById("package_display_page_panel").style.display = "none"
    document.getElementById("item_display_page_panel").style.display = "none"
    
    if(document.getElementById("item_search_result_page") != null){
         document.getElementById("item_search_result_page").style.display = "none"
    }
   
    document.getElementById("order_page_panel").style.display = ""
    if(document.getElementById("order_history_page_panel") != null){
      document.getElementById("order_history_page_panel").style.display = "none"  
    }
    
    
   
    
    
    
    catalogMessageHub.attach(this, createOrderEvent, function (evt, orderPage) {
        console.log(evt)
        
    })
    
    catalogMessageHub.attach(this, modifiedOrderEvent, function (evt, orderPage) {
        console.log(evt)
        
    })
    
    
}


OrderPage.prototype.addOrderItem = function (newOrderLine, firstLoad) {
    if(firstLoad == undefined){
        firstLoad = false
    }
    server.entityManager.getEntityByUuid(newOrderLine.M_itemSupplierRef, false, 
        function(item_supplier, caller){
            server.entityManager.getEntityByUuid(item_supplier.M_package, false, 
                function(item_package,caller){
                    if(mainPage.orderPage.panel.getChildById(caller.item_supplier.M_id + "-productRow") != null){
                        for(var i = 0; i < mainPage.orderPage.currentOrder.M_items.length; i++){
                            if(caller.newOrderLine.M_itemSupplierRef == mainPage.orderPage.currentOrder.M_items[i].M_itemSupplierRef){
                                mainPage.orderPage.currentOrder.M_items[i].M_quantity = Number(mainPage.orderPage.currentOrder.M_items[i].M_quantity) + Number(caller.newOrderLine.M_quantity)  
                                console.log(mainPage.orderPage.currentOrder.M_items[i].M_quantity)
                                mainPage.orderPage.panel.getChildById(caller.item_supplier.M_id  + "-qtyToOrder").element.value =  mainPage.orderPage.currentOrder.M_items[i].M_quantity
                                if(caller.newOrderLine.M_quantity == 1){
                                    mainPage.showNotification("primary", "Le paquet " + item_package.M_id + " est déjà dans le panier, " + caller.newOrderLine.M_quantity + " instance a été rajoutée.", 4000) 
                                }else{
                                    mainPage.showNotification("primary", "Le paquet " + item_package.M_id + " est déjà dans le panier, " + caller.newOrderLine.M_quantity + " instances ont été rajoutées.", 4000) 
                                }   
                            }
                        }
                    }else{
                        if(!firstLoad){
                            if(caller.newOrderLine.M_quantity == 1){
                                mainPage.showNotification("primary","Le paquet " + item_package.M_id + " a été rajouté au panier" , 4000) 
                            }else{
                                mainPage.showNotification("primary", caller.newOrderLine.M_quantity + " paquets "+ item_package.M_id + " ont été rajoutés au panier.", 4000) 
                            }   
                        }
                      
                        var isExist = false
                        for(var i = 0; i < mainPage.orderPage.currentOrder.M_items.length; i++){
                            if(caller.newOrderLine.M_itemSupplierRef == mainPage.orderPage.currentOrder.M_items[i].M_itemSupplierRef){
                                isExist = true
                            }
                        }
                        if(!isExist){
                            mainPage.orderPage.currentOrder.M_items.push(caller.newOrderLine)
                        }
                        
                        mainPage.panel.getChildById("cartCount").element.innerHTML = mainPage.orderPage.currentOrder.M_items.length
                       
                         mainPage.orderPage.panel.getChildById("orderBody").appendElement({"tag" : "tr", "id" : caller.item_supplier.M_id  + "-productRow"}).down()
                        .appendElement({"tag" :"td", "data-th" : "Produit"}).down()
                        .appendElement({"tag" : "div", "class" : "row"}).down()
                        .appendElement({"tag" : "div", "class" : "col-sm-2 d-none d-sm-block"}).down()
                        .appendElement({"tag" : "img", "src" : "../../Catalogue_2/photo/" + item_package.M_id + "/photo_1.jpg", "class" : "img-fluid rounded"}).up()
                        .appendElement({"tag" : "div", "class" : "col-sm-10"}).down()
                        .appendElement({"tag" :"h4", "class" : "nomargin", "innerHtml" : item_package.M_id})
                        .appendElement({"tag" : "a", "innerHtml" : item_package.M_name}).up().up().up()
                        .appendElement({"tag" : "td", "data-th" : "Prix"}).down()
                        .appendElement({"tag" : "span", "innerHtml":caller.item_supplier.M_price.M_valueOf})
                        .appendElement({"tag" : "span","innerHtml" : caller.item_supplier.M_price.M_currency.M_valueOf, "style" : "margin-left:5px;"}).up()
                        .appendElement({"tag" :"td", "data-th" : "Quantité"}).down()
                        .appendElement({"tag" :"input", "type" : "number", "class" : "form-control text-center", "value" : caller.newOrderLine.M_quantity, "id" : caller.item_supplier.M_id + "-qtyToOrder"}).up()
                        .appendElement({"tag" : "td", "data-th" : "Sous-total", "class" : "text-center"}).down()
                        .appendElement({"tag" : "span","innerHtml" : (caller.item_supplier.M_price.M_valueOf * caller.newOrderLine.M_quantity).toFixed(2), "id" : caller.item_supplier.M_id + "-subtotal"})
                        .appendElement({"tag" : "span", "innerHtml" : caller.item_supplier.M_price.M_currency.M_valueOf, "style" : "margin-left: 5px;"}).up()
                        .appendElement({"tag":"td", "class" : "actions", "data-th" : ""}).down()
                        .appendElement({"tag" : "button", "class" : "btn btn-danger", "id" : caller.item_supplier.M_id + "-deleteRowBtn"}).down()
                        .appendElement({"tag" : "i", "class" : "fa fa-trash-o"}).up()
                        .appendElement({"tag" : "hr", "class" : "my-2", "style" : "opacity:0;"})
                        .appendElement({"tag" : "button", "class" : "btn btn-outline-dark", "id" : item_package.M_id + "-searchBtn"}).down()
                                .appendElement({"tag" : "i", "class" : "fa fa-search-plus", "style" : "color:inherit;"}).up()
                        
                        mainPage.orderPage.recalculateTotal()
                        
                        mainPage.orderPage.panel.getChildById(item_package.M_id + "-searchBtn").element.onclick = function(item_package){
                            return function () {
                                mainPage.packageDisplayPage.displayTabItem(item_package)
                            }
                        }(item_package)
                        
                        mainPage.orderPage.panel.getChildById(caller.item_supplier.M_id  + "-deleteRowBtn").element.onclick = function (orderLine, rowID){
                            return function () {
                                mainPage.orderPage.removeOrderItem(orderLine, rowID)
                                
                               
                            }
                        }(caller.newOrderLine, caller.item_supplier.M_id + "-productRow")
                        
                        mainPage.orderPage.panel.getChildById(caller.item_supplier.M_id  + "-qtyToOrder").element.addEventListener('change', function () {
                            mainPage.orderPage.panel.getChildById(caller.item_supplier.M_id  + "-subtotal").element.innerHTML = (caller.item_supplier.M_price.M_valueOf * mainPage.orderPage.panel.getChildById(caller.item_supplier.M_id  + "-qtyToOrder").element.value).toFixed(2)
                            mainPage.orderPage.currentOrder.M_items.find(function (newOrderLine, newQty, price){ 
                               
                               return function (obj) {
                                    if(obj.M_itemSupplierRef == newOrderLine.M_itemSupplierRef) {
                                        obj.M_quantity = newQty
                                        mainPage.orderPage.recalculateTotal()
                                        
                                    }
                                }
                                }(caller.newOrderLine, mainPage.orderPage.panel.getChildById(caller.item_supplier.M_id  + "-qtyToOrder").element.value, item_supplier.M_price))
                            })
                    }
                }, function () {}
                ,{"item_supplier" : item_supplier, "newOrderLine" : caller.newOrderLine})
        },
        function (errObj, caller){
            console.log(errObj)
        }, 
        {"newOrderLine" : newOrderLine})
    
}

OrderPage.prototype.recalculateTotal = function () {
    var total = 0
    var itemSupplierRefs = []
    for(var i = 0; i < mainPage.orderPage.currentOrder.M_items.length; i++){
        if(mainPage.orderPage.currentOrder.M_items[i] != null){
            var uuid = mainPage.orderPage.currentOrder.M_items[i].M_itemSupplierRef
            if(isObjectReference(uuid)){
                itemSupplierRefs.push({"uuid":uuid, "qty":Number(mainPage.orderPage.currentOrder.M_items[i].M_quantity)})
            }
        }
    }
    
    // Async call of callculate.
    var recalculateTotal = function(uuids, total, callback){
        var item = uuids.pop()
        var uuid = item.uuid
        var qty = item.qty
        server.entityManager.getEntityByUuid(uuid, false,
            function(item_supplier,caller){
                caller.total = Number(total) + (caller.qty * Number(item_supplier.M_price.M_valueOf))
                if(caller.uuids.length == 0){
                    mainPage.orderPage.panel.getChildById("orderTotal").element.innerHTML = caller.total.toFixed(2);
                    mainPage.orderPage.currentOrder.M_total.M_valueOf = caller.total
                }else{
                    caller.callback(caller.uuids, caller.total, caller.callback);
                }
            },
            function(){
                
            },{"uuids":uuids, "total":total, "qty":qty, "callback":callback})
    }
    if(itemSupplierRefs.length > 0){
        recalculateTotal(itemSupplierRefs, total, recalculateTotal);
    }else{
         mainPage.orderPage.panel.getChildById("orderTotal").element.innerHTML = "";
    }
    
    
    this.saveOrder(this.currentOrder)
}


OrderPage.prototype.removeOrderItem = function (orderLine,rowID) {
    var row = mainPage.orderPage.panel.getChildById(rowID)
    this.currentOrder.M_items.pop(orderLine)
    mainPage.panel.getChildById("cartCount").element.innerHTML = mainPage.orderPage.currentOrder.M_items.length
    row.delete()
    //this.panel.getChildById(rowID).element.parentNode.removeChild(mainPage.orderPage.panel.getChildById(rowID).element)
    this.saveOrder(this.currentOrder)
    this.recalculateTotal()
    
}

OrderPage.prototype.setOrder = function (order) {
    if(this.getOrder() !== null){
        order = this.getOrder()
    }
    this.firstLoad = true
    mainPage.orderPage.currentOrder = order
    
    mainPage.panel.getChildById("cartCount").element.innerHTML = order.M_items.length
    for(var i = 0; i < order.M_items.length; i++){
        mainPage.orderPage.addOrderItem(order.M_items[i], true)
    }

}