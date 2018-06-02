var SearchResultPage = function (parent) {
    this.parent = parent;
    this.panel = new Element(null, {"tag":"div","id" : "item_search_result_page", "class" : "container-fluid", "style" :"margin-top : 15px; width 100%;"})
    this.navbar = this.panel.appendElement({"tag" : "ul", "class" : "nav nav-tabs d-none d-sm-flex querynav"}).down()
    this.toggle = this.navbar.appendElement({"tag" : "button","class" :"btn btn-dark toggleBtn", "onclick" : "goToItem()", "innerHtml":"Voir les items trouvés",})
    .appendElement({"tag" : "button","class" :"btn btn-dark", "onclick" : "goToPackage()", "innerHtml":"Voir les paquets trouvés",})
    
    this.mobileContent =  this.panel.appendElement({"tag" : "div", "class" : "container-fluid d-block d-sm-none"}).down()
    this.resultsContent = this.panel.appendElement({"tag" : "div", "class" : "tab-content  d-none d-sm-flex"}).down()
  
    window.addEventListener('resize', function(searchResultPage){
        return function(evt){
            // Mobile view
            if(searchResultPage.panel.element.offsetWidth < 500){
                var contents = document.getElementsByClassName("result-tab-content")
                for(var i = 0; i < contents.length; i++){
                    var content = contents[i].childNodes[0];
                    if(content != null){
                        var mobileContainer = document.getElementById(content.id.replace("_search_result", "-mobile-query"))
                        if(mobileContainer !== null){
                            content.parentNode.removeChild(content)
                            mobileContainer.appendChild(content)
                        }
                    }
                }
            }else{
                var contents = document.getElementsByClassName("mobile-content")
                for(var i=0; i < contents.length; i++){
                    var id = contents[i].id
                    var content = contents[i].childNodes[0];
                    if(content != undefined){
                        var tabContainer = document.getElementById(content.id.replace("_search_result", "-query"))
                        if(tabContainer !== null){
                            content.parentNode.removeChild(content);
                            tabContainer.appendChild(content)
                        }
                    }
                }
            }
        }
    }(this), true);
     
    return this
}

function goToItem(){
    document.getElementById("item_display_page_panel").style.display = ""
    document.getElementById("item_search_result_page").style.display = "none"
}

function goToPackage(){
    document.getElementById("package_display_page_panel").style.display = ""
    document.getElementById("item_search_result_page").style.display = "none"
}

SearchResultPage.prototype.displayResults = function (results, query, context) {
    queryTxt = query
    query = query.replaceAll("*", "-").replaceAll(" ", "_").replaceAll("/", "_").replaceAll('"', "_").replaceAll("'", "_").replaceAll("\\", "_")
    
    this.parent.element.appendChild(this.panel.element)
    spinner.panel.element.style.display = "none";
    document.getElementById("item_display_page_panel").style.display = "none"
    document.getElementById("item_search_result_page").style.display = ""
     document.getElementById("order_page_panel").style.display = "none"
    
    if(results.results.length == 0){
        document.getElementById("navbarSearchText").classList.add("is-invalid");
        document.getElementById("navbarSearchText").value = ""
        document.getElementById("navbarSearchText").placeholder = "No results found for " + query
        return;
    }
    
    document.getElementById("navbarSearchText").classList.remove("is-invalid");
    document.getElementById("navbarSearchText").placeholder = server.languageManager.refresh()
    
    // In case a search already exist...
    if(document.getElementById(query + "-li") !== null){
        document.getElementById(query + "-li").childNodes[0].click()
        return;
    }
    
    if(document.getElementById(query + "-mobilediv") !== null){
        document.getElementById(query + "-mobilediv").childNodes[0].click()
        return;
    }
    
    this.navbar.appendElement({"tag" : "li", "class" : "nav-item", "name":query, "id" : query + "-li", "style" : "display:flex;flex-direction:row;"}).down()
    .appendElement({"tag" : "span", "class" : "nav-link active", "href" : "#" +query + "-query", "role" : "tab", "data-toggle" : "tab", "aria-controls" : query  + "-query", "name" : query, "id" : query + "-tab", "style" : "    display: flex;flex-direction: row;align-items: center; padding: 5px; border-radius:0px;"}).down()
    .appendElement({"tag" : "span", "innerHtml" : queryTxt, "style" : "display:table-cell; padding-right : 8px;"})
    .appendElement({"tag" : "div", "class" : "swoop", "style" : "display : table-cell;"})
    
    this.mobileContent.appendElement({"tag" : "div", "class" : "d-flex flex-column ", "name" : query, "id" : query + "-mobilediv"}).down()
    .appendElement({"tag" : "span", "class" : "d-flex flex-row align-items-center", "name":query, "id" : query + "-mobiletab", "style" : "padding-top :10px;" }).down()
    .appendElement({"tag" : "span","class" : "btn btn-secondary tabLink w-100", "innerHtml" : queryTxt, "style" : "display:table-cell; padding-right : 8px;"})
    .appendElement({"tag" : "span", "class" : "btn"}).down()
    .appendElement({"tag" : "div", "class" : "swoop"}).up().up()
    .appendElement({"tag" : "div", "class":"mobile-content container", "name":query, "id" : query + "-mobile-query", "style":"padding: 0px;"})
    
    this.resultsContent.appendElement({"tag" : "div","class" : "tab-pane fade show result-tab-content", "style" : "position:relative; background-color : white;", "name":query, "id" : query + "-query", "role" : "tabpanel", "aria-labelledby" :query + "-tab" }).down()
    
    if(context == "generalSearchLnk"){
        var packages = []
        var uuids = []
        if (results.estimate > 0) {
            for (var i = 0; i < results.results.length; i++) {
                var result = results.results[i];
                if (result.data.TYPENAME == "CatalogSchema.ItemType") {
                    if (uuids.indexOf(results.results[i].data.UUID) == -1) {
                        uuids.push(results.results[i].data.UUID)
                    }
                } else if (result.data.TYPENAME == "CatalogSchema.PropertyType") {
                    if (uuids.indexOf(results.results[i].data.ParentUuid) == -1) {
                        uuids.push(results.results[i].data.ParentUuid)
                    }
                } else if (result.data.TYPENAME == "CatalogSchema.ItemSupplierType") {
                    // In case of package type...
                    if (packages.indexOf(result.data.M_package) == -1) {
                        packages.push(result.data.M_package)
                    }
                }
            }
            headerText = "About " + (uuids.length) + " results " + " (" + results.elapsedTime / 1000 + "seconds)"
        }
    
        // So here I will regroup results by type...
        var results_ = [];
    
        // Properties are asynchrone...
        var getItem = function (uuid, itemSearchResultPage, results, properties, query, getItem, callback) {
            // In that case I will get the entity from the server...
            server.entityManager.getEntityByUuid(uuid, false,
                // success callback
                function (entity, caller) {
                    caller.results.push(entity)
                    var next = caller.uuids.pop()
                    if (caller.uuids.length == 0) {
                        caller.callback(caller.results, caller.query, caller.itemSearchResultPage)
                    } else {
                        caller.getItem(next, caller.itemSearchResultPage, caller.results, caller.uuids, caller.query, caller.getItem, caller.callback)
                    }
                },
                // error callback
                function () {
    
                }, { "itemSearchResultPage": itemSearchResultPage, "results": results, "uuids": uuids, "query": queryTxt, "getItem": getItem, "callback": callback })
        }
    
        // Get the properties items if any's
        if (uuids.length > 0) {
            getItem(uuids[0], this, results_, uuids, query, getItem,
                // The callback...
                function (packages) {
                    return function (results, query, itemSearchResultPage) {
                        if (packages.length == 0) {
                            // simply display the results.
                            itemSearchResultPage.displayTabResults(results, query)
                        } else {
                            // In that case I will also append items from the packages...
                            var getPackageItems = function (pacakageUuid, packages, itemSearchResultPage, results, query, getPackageItems, callback) {
                                server.entityManager.getEntityByUuid(pacakageUuid, false,
                                    // success callback 
                                    function (results, caller) {
                                        caller.getPacakageItem = function (items, caller) {
                                            var item = items.pop()
                                            caller.items = items;
                                            // Here I will get the item reference value.
                                            server.entityManager.getEntityByUuid(item, false,
                                                // success callback 
                                                function (results, caller) {
                                                    if (!objectPropInArray(caller.results, "UUID", results.UUID)) {
                                                        caller.results.push(results)
                                                    }
                                                    // Here I will get the item reference value.
                                                    if (caller.items.length == 0) {
                                                        if (caller.packages.length == 0) {
                                                            // end of process
                                                            caller.callback(caller.results, caller.query, caller.itemSearchResultPage)
                                                        } else {
                                                            // call another time.
                                                            caller.getPackageItems(caller.packages.pop(0), caller.packages, caller.itemSearchResultPage, caller.results, caller.query, caller.getPackageItems, caller.callback)
                                                        }
                                                    } else {
                                                        caller.getPacakageItem(caller.items, caller)
                                                    }
                                                },
                                                // error callback
                                                function (errObj, caller) {
    
                                                }, caller)
                                        }
                                        
                                        caller.getPacakageItem(JSON.parse(JSON.stringify(results.M_items)), caller)
    
                                    },
                                    // error callback
                                    function (errObj, caller) {
    
                                    }, { "itemSearchResultPage": itemSearchResultPage, "packages": packages, "results": results, "query": query, "getPackageItems": getPackageItems, "callback": callback })
                            }
                            // process the first item.
                            getPackageItems(packages.pop(0), packages, itemSearchResultPage, results, query, getPackageItems, function (results, query, itemSearchResultPage) {
                                itemSearchResultPage.displayTabResults(results, query)
                            })
                        }
                    }
                }(packages)
            )
        }
    }else if(context == "supplierSearchLnk"){
        // The seach result must be about supplier and supplied items and
        // not the items itself...
        var uuids = []
        var itemSupplier = [] // contain the supplier/itemsSupplier values, that's what we want to display.
        if (results.estimate > 0) {
            for (var i = 0; i < results.results.length; i++) {
                var result = results.results[i];
                if (result.data.TYPENAME == "CatalogSchema.ItemType") {
                    if (uuids.indexOf(results.results[i].data.UUID) == -1) {
                        if(results.results[i].data.UUID != undefined){
                            uuids.push(results.results[i].data.UUID)
                        }
                    }
                } else if (result.data.TYPENAME == "CatalogSchema.PropertyType") {
                    if (uuids.indexOf(results.results[i].data.ParentUuid) == -1) {
                        if(results.results[i].data.ParentUuid != undefined){
                            uuids.push(results.results[i].data.ParentUuid)
                        }
                    }
                }else if (result.data.TYPENAME == "CatalogSchema.ItemSupplierType") {
                    itemSupplier.push(results.results[i].data)
                }
            }
            headerText = "About " + (uuids.length) + " results " + " (" + results.elapsedTime / 1000 + "seconds)"
        }
        
        var initItemSuppliers = function(uuids, itemSuppliers, q, itemSearchResultPage, callback){
            // The query.
            var uuid = uuids.pop()
            
            // In case of undefined values.
            if(uuid == undefined){
                initItemSuppliers(uuids, itemSuppliers, q, itemSearchResultPage, callback)
                if(uuids.length == 0){
                    callback(uuids, itemSuppliers, query, itemSearchResultPage, callback)
                }
                return
            }
           
        	server.entityManager.getEntityByUuid(uuid, false,
        		// success callback
        		function (itemSupplier, caller) {
        			// return the results.
        			caller.itemSuppliers.push(itemSupplier)
        			if(caller.uuids.length > 0){
        			    caller.callback(caller.uuids, caller.itemSuppliers, caller.query, caller.itemSearchResultPage, caller.callback)
        			}else{
        			    caller.itemSearchResultPage.displayTabResults(caller.itemSuppliers, caller.query)
        			}
        		},
        		// error callback
        		function (errObj, caller) {
        
        		}, {"uuids":uuids, "itemSuppliers":itemSuppliers, "query": query, "itemSearchResultPage":itemSearchResultPage, "callback":callback})
        }
        
        var getItemSuppliers = function(uuids, itemSuppliers, q, itemSearchResultPage, callback){
            // The query.
            var uuid = uuids.pop()
        	var query = {}
        	query.TypeName = "CatalogSchema.PackageType"
        	query.Fields = ["M_supplied"]
        	query.Query = 'CatalogSchema.PackageType.UUID == "' + uuid + '"'
        	server.dataManager.read('CatalogSchema', JSON.stringify(query), query.Fields, [],
        		// success callback
        		function (results, caller) {
        			// return the results.
        			caller.itemSuppliers = caller.itemSuppliers.concat(results[0][0])
        			if(caller.uuids.length > 0){
        			    caller.callback(caller.uuids, caller.itemSuppliers, caller.query, caller.itemSearchResultPage, caller.callback)
        			}else{
        			    initItemSuppliers(caller.itemSuppliers, [], caller.query, caller.itemSearchResultPage, initItemSuppliers)
        			}
        		},
        		// progress callback
        		function (index, total, caller) {
        
        		},
        		// error callback
        		function (errObj, caller) {
        
        		}, {"uuids":uuids, "itemSuppliers":itemSuppliers, "query": q, "itemSearchResultPage":itemSearchResultPage, "callback":callback})
        }
        
        // I will use the data manager to get the basic information.
        var getPackages = function(uuids, pacakages, q, itemSearchResultPage, callback){
            // The query.
            var uuid = uuids.pop()
        	var query = {}
        	query.TypeName = "CatalogSchema.ItemType"
        	query.Fields = ["M_packaged"]
        	query.Query = 'CatalogSchema.ItemType.UUID == "' + uuid + '"'
        	server.dataManager.read('CatalogSchema', JSON.stringify(query), query.Fields, [],
        		// success callback
        		function (results, caller) {
        			// return the results.
        			caller.pacakages = results[0][0].concat(caller.pacakages)
        			if(caller.uuids.length > 0){
        			    caller.callback(caller.uuids, caller.pacakages, caller.query, caller.itemSearchResultPage, caller.callback)
        			}else{
        			    getItemSuppliers(caller.pacakages, [], caller.query, caller.itemSearchResultPage, getItemSuppliers)
        			}
        		},
        		// progress callback
        		function (index, total, caller) {
        
        		},
        		// error callback
        		function (errObj, caller) {
        
        		}, {"uuids":uuids, "pacakages":pacakages, "query": q, "itemSearchResultPage":itemSearchResultPage, "callback":callback})
        }
        if(uuids.length > 0 && itemSupplier == 0){
            getPackages(uuids, [], query, this, getPackages);
        }else{
             this.displayTabResults(itemSupplier, query)
        }
        // this.displayTabResults(results, query)
    }else if(context == "localisationSearchLnk"){
        
    }
    
    fireResize()
}

SearchResultPage.prototype.displayTabResults = function (results, query, tabContent) {

    // Here I will set the tow string function for the properties type
    // that function is use by the column filter to help filtering the 
    // table values.
    var queryTxt = query
    query = query.replaceAll("*", "-").replaceAll(" ", "_").replaceAll("/", "_").replaceAll('"', "_").replaceAll("'", "_").replaceAll("\\", "_")
    
    
    document.getElementById(query + "-tab").innerHTML = ""
    document.getElementById(query + "-mobiletab").innerHTML = ""
    var tab = this.navbar.getChildById(query + "-tab").appendElement({"tag":"a","class" : "tabLink", "href" : "#" +query + "-query", "innerHtml" : queryTxt + " - " + results.length + " results", "style" : "text-decoration : none;"})
    
    var mobileTab = this.mobileContent.getChildById(query+"-mobiletab")
    mobileTab.element.classList.add("btn-group")
    mobileTab.appendElement({"tag":"a","class" : " btn btn-secondary tabLink w-100", "href" : "#" +query + "-query", "innerHtml" : queryTxt + " - " + results.length + " results"})
    
    // The close button.
    var closeBtn = new Element(null, {"tag" : "a", "name" : query, "aria-label" : "Close" , "class" : "tabLink tabLink", "id" : query+"-closebtn"})
    closeBtn.appendElement({"tag" : "span", "aria-hidden" :"true", "innerHtml" : "&times;", "style" : "padding-left: 5px;color:#007bff"}).up()
    
    closeBtn.element.onclick = function(tab, query){
        return function(){
            tabParent = tab.element.parentNode;
            var i = Array.prototype.indexOf.call(tabParent.childNodes, tab.element);
            tabParent.removeChild(tab.element)
            // The first tow button are not search item tabs...
            if(tabParent.childNodes.length > 2){
                if(i > tabParent.childNodes.length - 1){
                    tabParent.childNodes[tabParent.childNodes.length - 1].childNodes[0].click()
                }else{
                    tabParent.childNodes[i].childNodes[0].click()
                }
            }
            
            // Now I will remove all occurence of the query interface.
            var toDelete = document.getElementsByName(query)
            while(toDelete.length > 0){
                var i = toDelete.length - 1
                toDelete[i].parentNode.removeChild(toDelete[i]);
            }
        }
    }(this.navbar.getChildById(query + "-li"), query)
    
    mobileTab.element.onclick =  function(closeBtn, query, searchPage){
        return function(){
            if(closeBtn.element.parentNode !== null){
                try{
                    closeBtn.element.className = "tabLink"
                    closeBtn.element.parentNode.removeChild(closeBtn.element)
                }catch(err){
                    
                }
            }else{
                closeBtn.element.className = " btn btn-secondary tabLink"
                this.appendChild(closeBtn.element)
            }
            
            var mobileContainers = document.getElementsByClassName( "mobile-content");
            function collapse(panel){
                var keyframe = "100% {max-height: 0px;}"
                panel.animate(keyframe, .75, 
                    function(panel){
                        return function(){
                             panel.element.style.maxHeight = "0px";
                        }
                }(panel))
            }
            
            function expand(panel){
                var keyframe = "100% {max-height: 400px;}"
                panel.animate(keyframe, .75, 
                    function(panel){
                        return function(){
                             panel.element.style.maxHeight = "400px";
                        }
                }(panel))
            }
            
            for(var i = 0; i < mobileContainers.length; i++){
                var mobileContainer = searchPage.panel.getChildById(mobileContainers[i].id);
                if(mobileContainer.id !== query + "-mobile-query"){
                    collapse(mobileContainer)
                }else{
                    if(mobileContainer.element.style.maxHeight == "0px"){
                        expand(mobileContainer)
                    }else{
                        collapse(mobileContainer)
                    }
                }
            }
            fireResize()
        }
    }(closeBtn, query, this)
    
    tab.element.onmouseover =  function(closeBtn){
        return function(){
            try{
                closeBtn.element.parentNode.removeChild(closeBtn.element)
            }catch(err){
                
            }
            this.appendChild(closeBtn.element)
        }
    }(closeBtn)

    mobileTab.element.onmouseleave = tab.element.onmouseleave = function(closeBtn){
        return function(){
            closeBtn.element.className = "tabLink"
            try{
                closeBtn.element.parentNode.removeChild(closeBtn.element)
            }catch(err){
                
            }
        }
    }(closeBtn)
    
    if(mainPage.searchContext == "generalSearchLnk"){

        // This will be the next seach result table.
        new ItemSearchResultTable(this.resultsContent.getChildById(query+"-query"), results, query, function(query){
            // init callback.
            spinner.panel.element.style.display = "none";
            var tab = document.getElementById(query + "-tab")
            tab.classList.remove("active")
            tab.click()
            
            var mobiletab = document.getElementById(query + "-mobiletab")
            mobiletab.classList.remove("active")
            tab.click()
            
             // triger resize event
            fireResize()

            // refresh the language.
            server.languageManager.refresh();
            document.getElementById("advancedSearchText").innerHTML = document.getElementById(mainPage.searchContext).innerHTML
        })
        
    } else if(mainPage.searchContext == "supplierSearchLnk"){
        
        // Display the search results for that client.
        var supplierSearchResults = new ItemSupplierSearchResults(this.resultsContent.getChildById(query+"-query"), results, query, function(query){
            var tab = document.getElementById(query + "-tab")
            tab.classList.remove("active")
            tab.click()
                    
            var mobiletab = document.getElementById(query + "-mobiletab")
            mobiletab.classList.remove("active")
            tab.click()
                    
             /** triger resize event **/
            fireResize()
    
            // refresh the language.
            server.languageManager.refresh();
            document.getElementById("advancedSearchText").innerHTML = document.getElementById(mainPage.searchContext).innerHTML

        })

    }
}


/**
 * Display the seach results.
 */
var ItemSearchResultTable = function(parent, items, query, callback){
    parent.element.style.width = "100%";
    
    // Callback informations.
    this.callback = callback
    this.query = query
    
    // page information
    this.pageIndex = 0;
    this.pageSize = 20;
    
    // Keep the items to display.
    this.items = items;
    
    // Ajout de la table de recherhe.
    this.panel = parent.appendElement({"tag":"div", "class":"container-fluid", "id": query + "_search_result"}).down()
    this.table = this.panel.appendElement({"tag":"table", "class":"table"}).down()
    
    // Now I will append the header.
    this.table.appendElement({"tag":"thead", "class":"table_header thead-dark"}).down()
        .appendElement({"tag":"tr"}).down()
        .appendElement({"tag":"th", "id":"item-id-header", "innerHtml":"OVMM"})
        .appendElement({"tag":"th", "id":"item-type-header", "innerHtml":"Type"})
        .appendElement({"tag":"th", "id":"item-description-header", "innerHtml":"Description"})
        .appendElement({"tag":"th", "id":"item-alias-header", "innerHtml":"alias"})
        .appendElement({"tag":"th", "id":"item-keyword-header", "innerHtml":"mot clés"})
        .appendElement({"tag":"th", "id":"item-comments-header", "innerHtml":"commentaire(s)"})
        .appendElement({"tag":"th", "id":"item-specs-header", "innerHtml":"spécification"})
        .appendElement({"tag":"th", "id":"item-basket-header", "class":"append_item_btn"})
        
    // Now the table body.
    this.tbody = this.table.appendElement({"tag":"tbody"}).down()
    
    // Now the pagination...
    var ul = this.panel.appendElement({"tag":"nav"}).down()
        .appendElement({"tag":"ul", "class":"pagination", "style":"display: flex; justify-content: center; flex-wrap: wrap;"}).down()
    
    // The previous link.
    /*this.previous = ul.appendElement({"tag":"li", "class":"page-item"}).down()
    this.previous.appendElement({"tag":"a", "class":"page-link", "href":"#", "aria-label":"Previous"}).down()
    .appendElement({"tag":"span", "aria-hidden":"true", "innerHtml":"&laquo;"})
    .appendElement({"tag":"sr-only", "aria-hidden":"true"}).up().up()*/
    
    this.pageLnks = []
    for(var i=0; i < items.length / this.pageSize; i++ ){
        this.pageLnks[i] = ul.appendElement({"tag":"li", "class":"page-item"}).down()
        this.pageLnks[i].appendElement({"tag":"a", "class":"page-link", "href":"#", "innerHtml":(i + 1).toString()}).down()
        this.pageLnks[i].element.onclick = function(index, itemSearchResultTable){
            return function(){
                for(var i=0; i < itemSearchResultTable.pageLnks.length; i++){
                    itemSearchResultTable.pageLnks[i].element.classList.remove("active")
                }
                this.classList.add("active")
                itemSearchResultTable.displayPage(index)
            }
        }(i, this)
        
    }
    
    this.pageLnks[0].element.click()

    // Now the page to list of page to be inserted.
    /*this.next = ul.appendElement({"tag":"li", "class":"page-item"}).down()
    this.next.appendElement({"tag":"a", "class":"page-link", "href":"#", "aria-label":"Next"}).down()
    .appendElement({"tag":"span", "aria-hidden":"true", "innerHtml":"&raquo;"})
    .appendElement({"tag":"sr-only", "aria-hidden":"true"})*/
        
    // So here i will react to the login event.
    
    // The login event.
	catalogMessageHub.attach(this, welcomeEvent, function(evt, resultPage){
	    if(server.account === undefined){
	        return
	    }
	    
        var appendItemBtns = document.getElementsByClassName("append_item_btn");
        for(var i=0; i < appendItemBtns.length; i++){
            appendItemBtns[i].style.display = ""
        }
	})

    return this
}

ItemSearchResultTable.prototype.displayPage = function(index){
    // set the index.
    this.pageIndex = index;
    this.tbody.removeAllChilds();
    
    for(var i=0; i < this.pageSize && (this.pageSize*index) + i < this.items.length; i++){
        var item = this.items[(this.pageSize*index) + i];
        var tr = this.tbody.appendElement({"tag":"tr"}).down()
        
        var lnk = tr.appendElement({"tag":"td", "innerHtml":item.M_id}).down()
        lnk.element.style.textDecoration = "underline"
        lnk.element.onmouseover = function () {
            this.style.color = "steelblue"
            this.style.cursor = "pointer"
        }

        lnk.element.onmouseleave = function () {
            this.style.color = ""
            this.style.cursor = "default"
        }

        // Now the onclick... 
        var entity = entities[item.UUID]
        lnk.element.onclick = function (entity) {
            return function () {
                mainPage.itemDisplayPage.displayTabItem(entity)
            }
        }(entity)
        
        tr.appendElement({"tag":"td", "innerHtml":item.M_name})
        tr.appendElement({"tag":"td", "innerHtml":item.M_description})
        
        // The alias or supplier 
        var alias = tr.appendElement({"tag":"td"}).down()
            .appendElement({"tag":"tbody"}).down()
            
        for(var j=0; j < item.M_alias.length; j++){
             alias.appendElement({"tag":"tr"}).down()
                .appendElement({"tag":"td", "innerHtml":item.M_alias[j]})
        }
        
        // Now the keyword...
        var keywords =  tr.appendElement({"tag":"td"}).down()
            .appendElement({"tag":"tbody"}).down()
            
        for(var j=0; j < item.M_keywords.length; j++){
             keywords.appendElement({"tag":"tr"}).down()
                .appendElement({"tag":"td", "innerHtml":item.M_keywords[j]})
        }
        
        var comments =  tr.appendElement({"tag":"td"}).down()
            .appendElement({"tag":"tbody"}).down()
            
        for(var j=0; j < item.M_comments.length; j++){
             comments.appendElement({"tag":"tr"}).down()
                .appendElement({"tag":"td", "innerHtml":item.M_comments[j]})
        }

        // Now the properties...
        var properties =  tr.appendElement({"tag":"td"}).down()
            .appendElement({"tag":"tbody"}).down()
            
        var appendItemBtn = properties.up().up().appendElement({"tag":"td"}).down()
            .appendElement({"tag":"div", "class":"append_item_btn", "style":"position: relative; display: none;"}).down()
        
        appendItemBtn.appendElement({"tag":"span", "class":"fa fa-cart-plus"})
        
        // display if the user is connected.
        if(server.account !== undefined){
            appendItemBtn.element.style.display = "";
        }
        
        appendItemBtn.element.onclick = function(item){
            return function(){
                mainPage.itemDisplayPage.displayTabItem(item, function(item){
                    return function(){
                    packageElement = document.getElementById(item.M_id + "-suppliersDetails")
                    packageElement.scrollIntoView({
                        behavior : 'smooth'
                    })
                }}(item))
              
            }
        }(item)
	    
        // TODO display other type of values.
        for(var j=0; j < item.M_properties.length; j++){
            var property = item.M_properties[j]
            // console.log(property.M_name)
            var row = properties.appendElement({"tag":"tr"}).down();
            row.appendElement({"tag":"td", "innerHtml":property.M_name})
            if(property.M_stringValue.length > 0){
                row.appendElement({"tag":"td", "innerHtml":property.M_stringValue})
            }
        }
    }
    
    // call the callback
    this.callback(this.query)
}

