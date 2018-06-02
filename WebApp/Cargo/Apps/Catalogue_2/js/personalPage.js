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
    .appendElement({"tag" : "div", "class" : "col"}).down()
    .appendElement({"tag" : "div", "class" : "card"}).down()
    .appendElement({"tag" : "div", "class" : "card-body bg-dark text-light", "style": "border-top-right-radius:.25rem;border-top-left-radius:.25rem;"}).down()
    .appendElement({"tag" : "h5", "class"  :"card-title", "innerHtml" : "Commandes en attente"}).up()
    
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
    .appendElement({"tag" : "div", "class" : "col"}).down()
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
    
    server.entityManager.getEntityById("CargoEntities.User", "CargoEntities", [userId], false, 
        function(user, caller){
            caller.getChildById("personal_welcome_msg").element.innerHTML = "Bienvenue " + user.M_firstName + " " + user.M_lastName;
            
        },
        function(){
            
        }, this.panel)
        
    
    // display the page.
    this.displayOrders(orders)
    
    return this;
 }
 
 PersonalPage.prototype.displayOrders = function(orders){
     console.log(96, orders.length)

    for(var i = 0; i < orders.length; i++){
        console.log(99, i, orders[i].M_status.M_valueOf)
        console.log(99, i, orders[i].M_items)
        
        if(orders[i].M_status.M_valueOf === "Pending"){
             var creationDateString = orders[i].M_creationDate[0] + orders[i].M_creationDate[1] + orders[i].M_creationDate[2] + orders[i].M_creationDate[3] + "/" +orders[i].M_creationDate[5] + orders[i].M_creationDate[6] + "/" + orders[i].M_creationDate[8] + orders[i].M_creationDate[9] + "-" + orders[i].M_creationDate[11] + orders[i].M_creationDate[12]+ ":" + orders[i].M_creationDate[14] + orders[i].M_creationDate[15]
            var completionDateString = orders[i].M_completionDate[0] + orders[i].M_completionDate[1] + orders[i].M_completionDate[2] + orders[i].M_completionDate[3] + "/" +orders[i].M_completionDate[5] + orders[i].M_completionDate[6] + "/" + orders[i].M_completionDate[8] + orders[i].M_completionDate[9] + "-" + orders[i].M_completionDate[11] + orders[i].M_completionDate[12]+ ":" + orders[i].M_completionDate[14] + orders[i].M_completionDate[15]
        
            
            mainPage.welcomePage.getChildById("personalOrdersBody").appendElement({"tag" : "tr", }).down()
            .appendElement({"tag" :"td", "data-th" : "Date de création"}).down()
            .appendElement({"tag" : "span", "innerHtml" : creationDateString}).up()
            .appendElement({"tag" : "td", "data-th" : "Date d'envoi"}).down()
            .appendElement({"tag" : "span", "innerHtml": completionDateString}).up()
            .appendElement({"tag" :"td", "data-th" : "Produits commandés"}).down()
            .appendElement({"tag" : "span", "innerHtml": orders[i].M_items.length.toString()}).up()
            .appendElement({"tag" :"td", "data-th" : ""}).down()
            .appendElement({"tag" : "button", "class" : "btn btn-outline-dark", "onclick" : ""}).down()
            .appendElement({"tag" : "i", "class" : "fa fa-search-plus", "style" : "color:inherit;"}).up()

        
        }
        
        
        for(var j =0; i<orders[i].M_items.length;i++){
 
            server.entityManager.getEntityByUuid(orders[i].M_items[j].M_itemSupplierRef, false, 
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
                            .appendElement({"tag" : "button", "class" : "btn btn-outline-dark", "onclick" : ""}).down()
                            .appendElement({"tag" : "i", "class" : "fa fa-search-plus", "style" : "color:inherit;"}).up()
                        }, function() {}, {})
                },
                function(){
                    
                },{})
        }
        
        
        
    }
 }