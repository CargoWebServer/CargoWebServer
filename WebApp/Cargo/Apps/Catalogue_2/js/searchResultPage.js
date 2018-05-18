var SearchResultPage = function (parent) {
    this.parent = parent;
    this.panel = new Element(null, {"tag":"div","id" : "item_search_result_page", "class" : "container-fluid", "style" :"margin-top : 15px; width 100%;"})
    this.navbar = this.panel.appendElement({"tag" : "ul", "class" : "nav nav-tabs d-none d-sm-flex","id" : "queries"}).down()
    this.toggle = this.navbar.appendElement({"tag" : "button","class" :"btn btn-dark toggleBtn", "onclick" : "searchSwitchContext()", "innerHtml":"Voir les items trouvés",})
    
    this.mobileContent =  this.panel.appendElement({"tag" : "div", "class" : "container-fluid d-block d-sm-none"}).down()
    this.resultsContent = this.panel.appendElement({"tag" : "div", "class" : "tab-content  d-none d-sm-flex"}).down()
  
    window.addEventListener('resize', function(searchResultPage){
        return function(evt){
            // Mobile view
            if(searchResultPage.panel.element.offsetWidth < 500){
                var contents = document.getElementsByClassName("result-tab-content")
                for(var i = 0; i < contents.length; i++){
                    var content = contents[i].childNodes[1];
                    if(content != null){
                        content.parentNode.removeChild(content)
                        var mobileContainer = document.getElementById(content.id.replace("_search_result", "-mobile-query"))
                        mobileContainer.appendChild(content)
                    }
                }
            }else{
                var contents = document.getElementsByClassName("mobile-content")
                for(var i=0; i < contents.length; i++){
                    var id = contents[i].id
                    var content = contents[i].childNodes[0];
                    if(content != undefined){
                        content.parentNode.removeChild(content);
                        var tabContainer = document.getElementById(content.id.replace("_search_result", "-query"))
                        tabContainer.appendChild(content)
                    }
                }
            }
        }
    }(this), true);
     
    return this
}

function searchSwitchContext(){
    document.getElementById("item_display_page_panel").style.display = ""
    document.getElementById("item_search_result_page").style.display = "none"
}

SearchResultPage.prototype.displayResults = function (results, query, context) {
    this.parent.element.appendChild(this.panel.element)
    spinner.panel.element.style.display = "none";
    document.getElementById("item_display_page_panel").style.display = "none"
    document.getElementById("item_search_result_page").style.display = ""
    
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
    .appendElement({"tag" : "span", "innerHtml" : query, "style" : "display:table-cell; padding-right : 8px;"})
    .appendElement({"tag" : "div", "class" : "swoop", "style" : "display : table-cell;"})
    
    this.mobileContent.appendElement({"tag" : "div", "class" : "d-flex flex-column ", "name" : query, "id" : query + "-mobilediv"}).down()
    .appendElement({"tag" : "span", "class" : "d-flex flex-row align-items-center", /*, "href" : "#" + query + "-mobile-query", "role" : "tab", "data-toggle" : "tab", "aria-controls" : query ,"aria-expanded" : "false" ,*/ "name":query, "id" : query + "-mobiletab", "style" : "padding-top :10px;" }).down()
    .appendElement({"tag" : "span","class" : "btn btn-secondary tabLink w-100", "innerHtml" : query, "style" : "display:table-cell; padding-right : 8px;"})
    .appendElement({"tag" : "span", "class" : "btn"}).down()
    .appendElement({"tag" : "div", "class" : "swoop"}).up().up()
    .appendElement({"tag" : "div", "class":"mobile-content container", "name":query, "id" : query + "-mobile-query", "style":"padding: 0px;"})
    
    if(context == "generalSearchLnk"){
        this.resultsContent.appendElement({"tag" : "div","class" : "tab-pane fade show result-tab-content", "style" : "style : position:relative; background-color : white;", "name":query, "id" : query + "-query", "role" : "tabpanel", "aria-labelledby" :query + "-tab" }).down()
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
    
                }, { "itemSearchResultPage": itemSearchResultPage, "results": results, "uuids": uuids, "query": query, "getItem": getItem, "callback": callback })
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
                        uuids.push(results.results[i].data.UUID)
                    }
                } else if (result.data.TYPENAME == "CatalogSchema.PropertyType") {
                    if (uuids.indexOf(results.results[i].data.ParentUuid) == -1) {
                        uuids.push(results.results[i].data.ParentUuid)
                    }
                }
            }
            headerText = "About " + (uuids.length) + " results " + " (" + results.elapsedTime / 1000 + "seconds)"
        }
        
        var initItemSuppliers = function(uuids, itemSuppliers, q, itemSearchResultPage, callback){
            // The query.
            var uuid = uuids.pop()
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
        
        		}, {"uuids":uuids, "itemSuppliers":itemSuppliers, "query": q, "itemSearchResultPage":itemSearchResultPage, "callback":callback})
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
        if(uuids.length > 0){
            getPackages(uuids, [], query, this, getPackages);
        }else{
            
        }
        // this.displayTabResults(results, query)
    }else if(context == "localisationSearchLnk"){
        
    }
}

SearchResultPage.prototype.displayTabResults = function (results, query, tabContent) {

    // Here I will set the tow string function for the properties type
    // that function is use by the column filter to help filtering the 
    // table values.
    document.getElementById(query + "-tab").innerHTML = ""
    document.getElementById(query + "-mobiletab").innerHTML = ""
    var tab = this.navbar.getChildById(query + "-tab").appendElement({"tag":"a","class" : "tabLink", "href" : "#" +query + "-query", "innerHtml" : query + " - " + results.length + " results", "style" : "text-decoration : none;"})
    
    var mobileTab = this.mobileContent.getChildById(query+"-mobiletab")
    mobileTab.element.classList.add("btn-group")
    mobileTab.appendElement({"tag":"a","class" : " btn btn-secondary tabLink w-100", "href" : "#" +query + "-query", "innerHtml" : query + " - " + results.length + " results"})
    
    // The close button.
    var closeBtn = new Element(null, {"tag" : "a", "name" : query, "aria-label" : "Close" , "class" : "tabLink tabLink", "id" : query+"-closebtn"})
    closeBtn.appendElement({"tag" : "span", "aria-hidden" :"true", "innerHtml" : "&times;", "style" : "padding-left: 5px;color:#007bff"}).up()
    
    closeBtn.element.onclick = function(tab, query){
        return function(){
            tabParent = tab.element.parentNode;
            var i = Array.prototype.indexOf.call(tabParent.childNodes, tab.element);
            tabParent.removeChild(tab.element)
            if(tabParent.childNodes.length > 1){
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
        // Now I will append the found element in a table of entities.
        var proto = getEntityPrototype("CatalogSchema.ItemType")
    
        // Set fields visibility.
        proto.FieldsVisibility[10] = false;
        proto.FieldsVisibility[11] = false;
    
        var model = new EntityTableModel(proto);
        model.editable[2] = true;
        model.editable[3] = true;
    
        // Set the list of entities in the model.
        model.entities = results;
        
        var table = new Table(query+ "_search_result", this.resultsContent.getChildById(query+"-query"))
        table.div.element.classList.add("table")
        table.div.element.classList.add("table-responsive")
        table.div.element.classList.add("table-bordered")
    
        for (var i = 0; i < results.length; i++) {
            for (var j = 0; j < results[i].M_properties.length; j++) {
                var property = results[i].M_properties[j]
                if (isObject(property.M_dimensionValue)) {
                    // In that case I will display the value...
                    property.toString = function (dimension) {
                        return function () {
                            // Return the dimension value.
                            return dimension.M_valueOf + " " + dimension.M_unitOfMeasure.M_valueOf
                        }
                    }(property.M_dimensionValue)
                } else if (property.M_stringValue != undefined) {
                    property.toString = function (str) {
                        return function () {
                            return str
                        }
                    }(property.M_stringValue)
                } else if (property.M_numericValue != undefined) {
                    property.toString = function (val) {
                        return function () {
                            if (isInt(val)) {
                                return formatValue(val, "xs.int")
                            } else {
                                return formatValue(val, "xs.decimal")
                            }
                        }
                    }(property.M_numericValue)
                } else if (property.M_booleanValue != undefined) {
                    property.toString = function (val) {
                        return function () {
                            if (val == true) {
                                return "true"
                            } else {
                                return "false"
                            }
                        }
                    }(property.M_booleanValue)
                }
            }
        }
          
        // Here I will set the renderer for the type CatalogSchema.DimensionType
        // for the table.
        table.renderFcts["CatalogSchema.DimensionType"] = function (value) {
            // So here I will return the string that contain the unit of measure.
            var str = value.M_valueOf
    
            if (value.M_unitOfMeasure.M_valueOf == "IN (inch)") {
                str += '"'
            } else if (value.M_unitOfMeasure.M_valueOf == "FT (feet)") {
                str += "'"
            } else {
                str += " " + value.M_unitOfMeasure.M_valueOf
            }
            return new Element(null, { "tag": "div", "innerHtml": str });
        }
    
        // The property display function...
        table.renderFcts["[]CatalogSchema.PropertyType"] = function (properties) {
            // So here I will return the string that contain the unit of measure.
            var panel = new Element(null, { "tag": "div", "style": "display: table;width: 100%; border-collapse: collapse;margin-top: 2px;margin-bottom: 2px;" })
            for (var i = 0; i < properties.length; i++) {
                var property = properties[i]
                var row = panel.appendElement({ "tag": "div", "style": "display:table-row; width: 100%;" }).down()
    
                // The propertie name.
                row.appendElement({ "tag": "div", "style": "display:table-cell; padding-right: 5px;", "innerHtml": property.M_name + ": " })
    
                var valueDiv = row.appendElement({ "tag": "div", "style": "display:table-cell; width: 100%;" }).down()
                // The propertie value...
                if (isObject(property.M_dimensionValue)) {
                    // Render the Dimension type with the function....
                    valueDiv.appendElement(table.renderFcts["CatalogSchema.DimensionType"](property.M_dimensionValue))
                } else if (property.M_stringValue != undefined) {
                    valueDiv.element.innerHTML = property.M_stringValue
                }
            }
            return panel;
        }
    
        // I will append information to the model
        table.setModel(model, function (table,query) {
            return function () {
                // Set hidden column here.
                // init the table.
                table.parent.element.style.position = "relative"
                table.init()
                table.header.maximizeBtn.element.click();
                table.header.div.element.style.display = "table-header-group"
                spinner.panel.element.style.display = "none";
    
                // I will set the action on the first cell...
                for (var i = 0; i < table.rows.length; i++) {
                    var cell = table.rows[i].cells[0]
                    cell.div.element.style.textDecoration = "underline"
                    cell.div.element.onmouseover = function () {
                        this.style.color = "steelblue"
                        this.style.cursor = "pointer"
                    }
    
                    cell.div.element.onmouseleave = function () {
                        this.style.color = ""
                        this.style.cursor = "default"
                    }
    
                    // Now the onclick... 
                    var entity = table.getModel().entities[i]
                    cell.div.element.onclick = function (entity) {
                        return function () {
                            //mainPage.searchResultPage.resultsContent.element.parentNode.removeChild(mainPage.searchResultPage.resultsContent.element)
                            mainPage.itemDisplayPage.displayTabItem(entity)
                        }
                    }(entity)
                }
                
                var tab = document.getElementById(query + "-tab")
                tab.classList.remove("active")
                tab.click()
                
                var mobiletab = document.getElementById(query + "-mobiletab")
                mobiletab.classList.remove("active")
                tab.click()
                
                /** triger resize event **/
                fireResize()
                
                // Set the table header dark.
                table.header.div.element.classList.add("thead-dark")
                
                // refresh the language.
                server.languageManager.refresh();
            }
        }(table,query))
    } else if(mainPage.searchContext == "supplierSearchLnk"){
        console.log(results)
    }
}


