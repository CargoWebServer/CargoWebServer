var ItemSupplierSearchResults = function(parent, results, query, callback){
    
    // page information
    query = query.replace("*", "-");
    
    // first of all I will regroup item by suppliers.
    var resultsBySupplier = {}
    
    for(var i=0; i < results.length; i++){
        if(resultsBySupplier[results[i].M_supplier] === undefined){
            resultsBySupplier[results[i].M_supplier] = []
        }
        resultsBySupplier[results[i].M_supplier].push(results[i])
    }
    
    // The panel result panel.
    parent.element.style.width = "100%";
    this.panel = parent.appendElement({"tag":"div", "class":"container-fluid", "id" : query + "_search_result"}).down()
    
    
    for(var supplier in resultsBySupplier){
        this.table = this.panel.appendElement({"tag":"table", "class" :"table"}).down()
        
        this.table.appendElement({"tag":"thead"}).down()
        .appendElement({"tag":"tr", "class":"thead-dark"}).down()
        .appendElement({"tag":"th", "colspan":"3", "id":supplier + "-header-title"}).up()
        .appendElement({"tag":"tr"}).down()
        .appendElement({"tag":"th", "innerHtml":"id"})
        .appendElement({"tag":"th", "innerHtml":"prix"})
        .appendElement({"tag":"th", "innerHtml":"embalage"})
        
    	var q = {}
    	q.TypeName = "CatalogSchema.SupplierType"
    	q.Fields = ["M_name"]
    	q.Query = 'CatalogSchema.SupplierType.UUID == "' + supplier + '"'
    	
    	server.dataManager.read('CatalogSchema', JSON.stringify(q), q.Fields, [],
    		// success callback
    		function (results, caller) {
    			caller.element.innerHTML = results[0][0]
    			
    		},
    		// progress callback
    		function (index, total, caller) {
    
    		},
    		// error callback
    		function (errObj, caller) {
    
    		}, this.table.getChildById(supplier + "-header-title"))
        
        this.body = this.table.appendElement({"tag":"tbody"}).down()
        
        // The main panel of the supplier search.
        for(var i=0; i < resultsBySupplier[supplier].length; i++){
            var row = this.body.appendElement({"tag":"tr"}).down()
            new ItemSupplierSearchResult(row, resultsBySupplier[supplier][i])
        }
    }
    
    callback(query)
    return this;
}

var ItemSupplierSearchResult = function(row, result){
    
    // The id of the item.
    this.id = row.appendElement({"tag":"td", "innerHtml":result.M_id}).down();
    
    this.id.element.style.textDecoration = "underline"
    this.id.element.onmouseover = function () {
        this.style.color = "steelblue"
        this.style.cursor = "pointer"
    }

    this.id.element.onmouseleave = function () {
        this.style.color = ""
        this.style.cursor = "default"
    }

    // Now the onclick... 
    this.id.element.onclick = function (entity) {
        return function () {
            server.entityManager.getEntityByUuid(entity.M_package, false, 
            function(pack){
                mainPage.packageDisplayPage.displayTabItem(pack)
            },
            function(){
                
            }, {})
           
        }
    }(entities[result.UUID])
        
    // The price of the item.
    this.price = row.appendElement({"tag":"td", "innerHtml":result.M_price.M_valueOf + " " + result.M_price.M_currency.M_valueOf}).down();
    
    // Now the pacage information.
    this.packageInfo = row.appendElement({"tag":"tbody", "id": +  result.M_package + "-package-info"}).down()
    
    server.entityManager.getEntityByUuid(result.M_package, false, 
        // success callback
        function(result, body){
            body.appendElement({"tag":"tr"}).down()
            .appendElement({"tag":"th", "scope":"row", "innerHtml":"name"})
            .appendElement({"tag":"td", "innerHtml":result.M_name}).up()
            .appendElement({"tag":"tr"}).down()
            .appendElement({"tag":"th", "scope":"row", "innerHtml":"*Qte."})
            .appendElement({"tag":"td", "innerHtml":result.M_quantity.toString(), "title":"quantity per package"}).up()
            
            // Now the localisation for that package...
            .appendElement({"tag":"tr"}).down()
            .appendElement({"tag":"th", "scope":"row", "innerHtml":"Emplacement(s)"})
            .appendElement({"tag":"tbody", "id":result.UUID + "-pacakage-emplacement-info"})
            
            // Now I will get the inventory object...
            for(var i=0; i < result.M_inventoried.length; i++){
                server.entityManager.getEntityByUuid(result.M_inventoried[i], false,
                    function(inventory, body){
                        var tr = body.appendElement({"tag":"tr"}).down()
                        var locTd = tr.appendElement({"tag":"td"}).down()
                        // Now the value of the inventory itself.
                        tr.appendElement({"tag":"td"}).down().appendElement({"tag":"tbody"}).down()
                            .appendElement({"tag":"tr"}).down()
                            .appendElement({"tag":"th", "scope":"row", "innerHtml":"safety stock"})
                            .appendElement({"tag":"td", "innerHtml":inventory.M_safetyStock}).up()
                            .appendElement({"tag":"tr"}).down()
                            .appendElement({"tag":"th", "scope":"row", "innerHtml":"reorder qte"})
                            .appendElement({"tag":"td", "innerHtml":inventory.M_reorderQty}).up()
                            .appendElement({"tag":"tr"}).down()
                            .appendElement({"tag":"th", "scope":"row", "innerHtml":"qte"})
                            .appendElement({"tag":"td", "innerHtml":inventory.M_quantity}).up()
                            
                            displayLocalisation(locTd, inventory.M_located, function(location){
                                // So here I will display the result in the search page.
                                mainPage.searchContext = "localisationSearchLnk";
                                mainPage.welcomePage.element.style.display = "none"
                                mainPage.searchResultPage.displayResults({"results":[{"data":location}], "estimate":1}, location.M_id, "localisationSearchLnk") 
                            })
                    },
                    function(){
                        
                    }, body.getChildById(result.UUID + "-pacakage-emplacement-info"))
            }
        }, 
        // error callback.
        function(){
            
        }, 
        this.packageInfo)
    
    return this;
}