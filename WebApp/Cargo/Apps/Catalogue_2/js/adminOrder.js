var AdminOrderPage = function(panel){
    /** The tab content panel **/
    this.panel = panel
    
    /** The panel sections. **/
    this.panel.appendElement({"tag" : "div", "class" : "row", "style" : "height:85vh;margin:0;"}).down()
        .appendElement({"tag" : "div", "class" : "col-md-3 bg-dark", "id" : "ordersAdminNavigation"}).down()
        .appendElement({"tag" : "div", "class" : "list-group","role" :"tablist", "id" : "ordersList", "style" : "margin-top : 30px;height:80vh;overflow-y:auto;"}).up()
        .appendElement({"tag" : "div", "class" : "col-md-9 bg-light","style":"overflow-y:auto;"}).down()
         .appendElement({"tag" : "div", "class" : "tab-content", "id" : "ordersAdminControl"}).up().up().up()
        
  
    
    
    this.loadOrders = function(){
        this.pendingOrders = []
        var q = new EntityQuery()
        q.TypeName = "CatalogSchema.OrderStatus"
        q.Fields = ["M_valueOf" ]
        q.Query = 'CatalogSchema.OrderStatus.M_valueOf=="Pending"'
        
        server.entityManager.getEntities("CatalogSchema.OrderStatus", "CatalogSchema", q, 0,-1,[],true,false, 
            function(index,total, caller){
                
            },
            function(orderStatus,caller){
                 mainPage.adminPage.adminOrderPage.panel.getChildById("ordersList").removeAllChilds()
                mainPage.adminPage.adminOrderPage.panel.getChildById("ordersAdminControl").removeAllChilds()
                for(var i = 0; i < orderStatus.length; i++){
                    server.entityManager.getEntityByUuid(orderStatus[i].ParentUuid,false,
                        function(order,caller){
                            mainPage.adminPage.adminOrderPage.pendingOrders.push(order)
                            
                            if(mainPage.adminPage.adminOrderPage.pendingOrders.length == caller.length){
                                mainPage.adminPage.adminOrderPage.pendingOrders.sort(function(a,b){
                                    
                                    return new Date(b.M_completionDate) - new Date(a.M_completionDate)
                                })
                                for(var i = 0; i < mainPage.adminPage.adminOrderPage.pendingOrders.length; i++){
                                    mainPage.adminPage.adminOrderPage.panel.getChildById("ordersList").appendElement({"tag":"a", "class" : "list-group-item list-group-item-action","id" : mainPage.adminPage.adminOrderPage.pendingOrders[i].getEscapedUuid() + "-selector","data-toggle" : "tab","href" : "#"+ mainPage.adminPage.adminOrderPage.pendingOrders[i].getEscapedUuid() + "-control","role":"tab","aria-controls":  mainPage.adminPage.adminOrderPage.pendingOrders[i].getEscapedUuid() + "-control"})
                                    mainPage.adminPage.adminOrderPage.loadAdminControl(mainPage.adminPage.adminOrderPage.pendingOrders[i])
                                    server.entityManager.getEntityById("CargoEntities.User", "CargoEntities", [mainPage.adminPage.adminOrderPage.pendingOrders[i].M_userId], false, 
                                        function(user, caller){
                                            caller.panel.element.innerHTML = user.M_firstName + " " + user.M_lastName  + ", " +  moment(caller.order.M_completionDate).format('DD/MM/YYYY, h:mm a')
                                            
                                        },
                                        function(){
                                            
                                        }, {"panel":mainPage.adminPage.adminOrderPage.panel.getChildById(mainPage.adminPage.adminOrderPage.pendingOrders[i].getEscapedUuid() + "-selector"), "order" :mainPage.adminPage.adminOrderPage.pendingOrders[i] })
        
                                }
                            }
                            
                           
                           
                            
                  
                       
                        },function(){},{"length" : orderStatus.length})
                }
               
                
            },
            function(){
            },{})
        }
  
    this.loadOrders()

    return this
}

AdminOrderPage.prototype.loadAdminControl = function(order){
    
    var creationDate = moment(order.M_creationDate).format('DD/MM/YYYY, h:mm a')
    var completionDate = moment(order.M_completionDate).format('DD/MM/YYYY, h:mm a')
    this.panel.getChildById("ordersAdminControl").appendElement({"tag" : "div", "class" : "tab-pane", "id" : order.getEscapedUuid() + "-control", "role" : "tabpanel", "aria-labelledby" : order.getEscapedUuid() + "-selector", "style" : "padding:15px;"}).down()
    .appendElement({"tag" : "div", "class" : "input-group mb-3"}).down()
    .appendElement({"tag" : "div", "class" : "input-group-prepend"}).down()
    .appendElement({"tag" : "span", "class" : "input-group-text", "innerHtml" : "Utilisateur"}).up()
    .appendElement({"tag" : "span", "class" : "form-control",  "id" : order.getEscapedUuid() + "-user"}).up()
    .appendElement({"tag" : "div", "class" : "input-group mb-3"}).down()
    .appendElement({"tag" : "div", "class" : "input-group-prepend"}).down()
    .appendElement({"tag" : "span", "class" : "input-group-text", "innerHtml" : "Date de création"}).up()
    .appendElement({"tag" : "span", "class" : "form-control",  "innerHtml" : creationDate}).up()
    .appendElement({"tag" : "div", "class" : "input-group mb-3"}).down()
    .appendElement({"tag" : "div", "class" : "input-group-prepend"}).down()
    .appendElement({"tag" : "span", "class" : "input-group-text", "innerHtml" : "Date de complétion"}).up()
    .appendElement({"tag" : "span", "class" : "form-control",  "innerHtml" : completionDate}).up()
    
    .appendElement({"tag" : "div"}).down()
    .appendElement({"tag":"table", "class":"table table-hover table-condensed"}).down()
    .appendElement({"tag" : "thead"}).down()
        .appendElement({"tag" : "tr"}).down()
        .appendElement({"tag" : "th", "style" : "width:40%;", "innerHtml" : "Produit"})
        .appendElement({"tag" : "th", "style" : "width:20%;", "innerHtml" : "Prix"})
        .appendElement({"tag" : "th", "style" : "width:8%;", "innerHtml" : "Quantité"})
        .appendElement({"tag" : "th", "style" : "width:22%;", "class" : "text-center", "innerHtml" : "Sous-total"})
        .appendElement({"tag" : "th", "style" : "width:10%;"}).up().up()
        .appendElement({"tag" : "tbody", "id" : order.getEscapedUuid() + "-adminOrderBody"}).up()
     .appendElement({"tag" : "div", "class" : "d-flex justify-content-center"}).down()
    .appendElement({"tag" : "button", "class" : "btn btn-success  mr-3", "innerHtml" : "Approuver la commande", "id" : order.getEscapedUuid() + "-saveBtn"})
    .appendElement({"tag" : "button", "class" : "btn btn-danger mr-3", "innerHtml" : "Annuler la commande", "id" : order.getEscapedUuid() + "-cancelBtn"})
        .appendElement({"tag" : "button", "class" : "btn mr-3", "id" : order.getEscapedUuid() + "-printBtn"}).down()
        .appendElement({"tag" : "span", "innerHtml" : "Imprimer"})
        .appendElement({"tag" : "i", "class" : "fa fa-print ml-1"})
        
    this.panel.getChildById(order.getEscapedUuid()+ "-printBtn").element.onclick = function(){
        window.print()
    }
        
    server.entityManager.getEntityById("CargoEntities.User", "CargoEntities", [order.M_userId], false, 
        function(user, caller){
            caller.element.innerHTML =user.M_firstName + " " + user.M_lastName;
            
        },
        function(){
            
        }, this.panel.getChildById(order.getEscapedUuid() + "-user"))
        
        
    
    this.panel.getChildById(order.getEscapedUuid() + "-saveBtn").element.onclick = function(order){
        return function(){
            order.M_status.M_valueOf = "Completed"
            server.entityManager.saveEntity(order, 
                function(success,caller){
                    mainPage.showNotification("success", "La commande a été approuvée!", 4000)
                    mainPage.adminPage.adminOrderPage.loadOrders()
                },function(){},{})
        }
    }(order)
    
    this.panel.getChildById(order.getEscapedUuid() + "-cancelBtn").element.onclick = function(order){
        return function(){
            order.M_status.M_valueOf = "Cancel"
            server.entityManager.saveEntity(order, 
                function(success,caller){
                    mainPage.showNotification("danger", "La commande a été annulée!", 4000)
                    mainPage.adminPage.adminOrderPage.loadOrders()
                },function(){},{})
        }
    }(order)
    
     
    for(var i = 0; i < order.M_items.length; i++){
        
        server.entityManager.getEntityByUuid(order.M_items[i].M_itemSupplierRef, false, 
        function(item_supplier, caller){
            server.entityManager.getEntityByUuid(item_supplier.M_package, false, 
                function(item_package,caller){
   
                         mainPage.adminPage.adminOrderPage.panel.getChildById(caller.order.getEscapedUuid() +"-adminOrderBody").appendElement({"tag" : "tr", "id" : caller.item_supplier.M_id  + "-productRow"}).down()
                        .appendElement({"tag" :"td", "data-th" : "Produit"}).down()
                        .appendElement({"tag" : "div", "class" : "row"}).down()
                        .appendElement({"tag" : "div", "class" : "col-sm-4 d-none d-sm-block"}).down()
                        .appendElement({"tag" : "img", "src" : "../../Catalogue_2/photo/" + item_package.M_id + "/photo_1.jpg", "class" : "img-fluid rounded"}).up()
                        .appendElement({"tag" : "div", "class" : "col-sm-8"}).down()
                        .appendElement({"tag" :"h4", "class" : "nomargin", "innerHtml" : item_package.M_id})
                        .appendElement({"tag" : "p", "innerHtml" : item_package.M_name})
                        .appendElement({"tag" : "span", "innerHtml" : item_package.M_id, "style" : "font-family: 'Libre Barcode 128 Text', cursive;font-size:50px;"}).up().up().up()
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
                        
                        mainPage.adminPage.adminOrderPage.panel.getChildById(item_package.M_id + "-modalSearchBtn").element.onclick = function(){
                                mainPage.packageDisplayPage.displayTabItem(item_package)
                        }
                        
                        
                        
                        
                        
                       
                    
                }, function () {}
                ,{"item_supplier" : item_supplier, "newOrderLine" : caller.newOrderLine, "order" : caller.order})
        },
        function (){
            
        }, {"newOrderLine" : order.M_items[i], "order" : order})
    }
}

