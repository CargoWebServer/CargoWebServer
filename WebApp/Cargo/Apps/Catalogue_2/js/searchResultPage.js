var SearchResultPage = function (parent) {
    this.parent = parent;
    this.panel = new Element(null, {"tag":"div","id" : "item_search_result_page", "class" : "container-fluid", "style" :"margin-top : 15px; width 100%;"})
    this.navbar = this.panel.appendElement({"tag" : "ul", "class" : "nav nav-tabs d-none d-sm-flex querynav"}).down()
    this.getItemBtn = this.navbar.appendElement({"tag" : "button","class" :"btn btn-dark toggleBtn", "innerHtml":"Voir les items trouvés",}).down()
    this.getItemBtn.element.onclick = function(){
        document.getElementById("item_display_page_panel").style.display = ""
        document.getElementById("item_search_result_page").style.display = "none"
        fireResize()
    }
    this.getPackageBtn = this.navbar.appendElement({"tag" : "button","class" :"btn btn-dark", "innerHtml":"Voir les paquets trouvés",}).down()
    this.getPackageBtn.element.onclick = function(){
        document.getElementById("package_display_page_panel").style.display = ""
        document.getElementById("item_search_result_page").style.display = "none"
        fireResize()
    }

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
                            try{
                                content.parentNode.removeChild(content)
                                mobileContainer.appendChild(content)
                            }catch(err){
                                
                            }
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


SearchResultPage.prototype.displayResults = function (results, query, context) {
    mainPage.itemDisplayPage.panel.element.style.display = "none"
    mainPage.packageDisplayPage.panel.element.style.display = "none"
    mainPage.welcomePage.element.style.display = "none"
    
    queryTxt = query
    query = query.replaceAll("*", "-").replaceAll(" ", "_").replaceAll("/", "_").replaceAll('"', "_").replaceAll("'", "_").replaceAll("\\", "_")
    
    this.parent.element.appendChild(this.panel.element)
    spinner.panel.element.style.display = "none";
    document.getElementById("admin_page_panel").style.display = "none"
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
    .appendElement({"tag" : "span", "class" : "nav-link active", "href" : "#" +query + "-query", "role" : "tab", "data-toggle" : "tab", "aria-controls" : query  + "-query", "name" : query, "id" : query + "-tab", "style" : "    display: flex;flex-direction: row;align-items: center; padding: 5px; border-radius:0px;border-left : 1px solid white;"}).down()
    .appendElement({"tag" : "span", "innerHtml" : queryTxt, "style" : "display:table-cell; padding-right : 8px;"})
    .appendElement({"tag" : "div", "class" : "swoop", "style" : "display : table-cell;"})
    
    this.mobileContent.appendElement({"tag" : "div", "class" : "d-flex flex-column ", "name" : query, "id" : query + "-mobilediv"}).down()
    .appendElement({"tag" : "span", "class" : "d-flex flex-row align-items-center", "name":query, "id" : query + "-mobiletab", "style" : "padding-top :10px;" }).down()
    .appendElement({"tag" : "span","class" : "btn btn-secondary tabLink w-100", "innerHtml" : queryTxt, "style" : "display:table-cell; padding-right : 8px;"})
    .appendElement({"tag" : "span", "class" : "btn"}).down()
    .appendElement({"tag" : "div", "class" : "swoop"}).up().up()
    .appendElement({"tag" : "div", "class":"mobile-content container", "name":query, "id" : query + "-mobile-query", "style":"padding: 0px;"})
    
    this.resultsContent.appendElement({"tag" : "div","class" : "tab-pane show result-tab-content", "style" : "position:relative; background-color : white;", "name":query, "id" : query + "-query", "role" : "tabpanel", "aria-labelledby" :query + "-tab" }).down()
    
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
        server.entityManager.getEntitiesByUuid(uuids, 
          function(index, total, caller){
              console.log("---> transfert ", index, "/", total)
          },
          function(results, caller){
              var items = []
              for(i=0; i < results.length; i++){
                  var item = new CatalogSchema.ItemType()
                  item.init(results[i], false)
                  items.push(item)
              }
              caller.itemSearchResultPage.displayTabResults(items, caller.query)
              
          },
          function(){
              
          }, {"query":query, "itemSearchResultPage":this})
          
    }else if(context == "supplierSearchLnk"){
        // The seach result must be about supplier and supplied items and
        // not the items itself...
        var uuids = []
        var itemSuppliers = [] // contain the supplier/itemsSupplier values, that's what we want to display.
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
                    itemSuppliers.push(results.results[i].data)
                }
            }
            headerText = "About " + (uuids.length) + " results " + " (" + results.elapsedTime / 1000 + "seconds)"
        }

        var getItemSuppliers = function(packages, itemSuppliers, q, itemSearchResultPage){
            
            server.entityManager.getEntitiesByUuid(packages, 
              function(index, total, caller){
                  console.log("---> transfert ", index, "/", total)
              },
              function(results, caller){
                  var itemSuppliers = caller.itemSuppliers
                  for(var i=0; i < results.length; i++){
                      for(var j=0; j < results[i].M_supplied.length; j++){
                          if(results[i].M_supplied[j] != null){
                              if(itemSuppliers.indexOf(results[i].M_supplied[j]) == -1){
                                itemSuppliers.push(results[i].M_supplied[j])
                              }
                          }
                      }
                  }
                  
                  server.entityManager.getEntitiesByUuid(itemSuppliers, 
                    function(index, total, caller){
                        console.log("---> transfert ", index, "/", total)
                    },
                    function(results, caller){
                        var itemSuppliers = []
                        for(var i=0; i < results.length; i++){
                            var itemSupplier = new CatalogSchema.ItemSupplierType();
                            itemSupplier.init(results[i], false);
                            itemSuppliers.push(itemSupplier);
                        }
                        caller.itemSearchResultPage.displayTabResults(itemSuppliers, caller.query)
                    },
                    function(err, caller){},
                    caller)
      
              },
              function(){
                  
              },  {"query":q, "itemSuppliers":itemSuppliers, "itemSearchResultPage":itemSearchResultPage})
        }
        
        // I will use the data manager to get the basic information.
        var getPackages = function(uuids, pacakages, itemSuppliers, q, itemSearchResultPage, callback){
            // The query.
            var uuid = uuids.pop()
        	var query = {}
        	query.TypeName = "CatalogSchema.ItemType"
        	query.Fields = ["M_packaged"]
        	query.Query = 'CatalogSchema.ItemType.UUID=="' + uuid + '"'
        	server.dataManager.read('CatalogSchema', JSON.stringify(query), query.Fields, [],
        		// success callback
        		function (results, caller) {
        			// return the results.
        			if(results[0][0][0] != undefined){
        			    caller.pacakages = results[0][0][0].concat(caller.pacakages)
        			}
        			if(caller.uuids.length > 0){
        			    caller.callback(caller.uuids, caller.pacakages, caller.itemSuppliers, caller.query, caller.itemSearchResultPage, caller.callback)
        			}else{
        			    getItemSuppliers(caller.pacakages, caller.itemSuppliers, caller.query, caller.itemSearchResultPage)
        			}
        		},
        		// progress callback
        		function (index, total, caller) {
        
        		},
        		// error callback
        		function (errObj, caller) {
                
        		}, {"uuids":uuids, "pacakages":pacakages, "itemSuppliers":itemSuppliers, "query": q, "itemSearchResultPage":itemSearchResultPage, "callback":callback})
        }
        if(uuids.length > 0){
            var uuids_ = []
            for(var i=0; i < itemSuppliers.length; i++){
                uuids_.push(itemSuppliers[i].UUID)
            }
            getPackages(uuids, [], uuids_, query, this, getPackages);
        }else if(itemSuppliers.length > 0){
            this.displayTabResults(itemSuppliers, query)
        }
        
    }else if(context == "localisationSearchLnk"){
        
        // In that case localisation is the the main think to look for...
        var items = []
        var localisations = []
        
        for (var i = 0; i < results.results.length; i++) {
            var result = results.results[i];
            if (result.data.TYPENAME == "CatalogSchema.LocalisationType") {
                var localisation = new CatalogSchema.LocalisationType()
                localisation.init(result.data)
                localisations.push(localisation)
            } else if (result.data.TYPENAME == "CatalogSchema.ItemType") {
                if(items.indexOf(result.data.UUID)==-1){
                    items.push(result.data.UUID)
                }
            } else if (result.data.TYPENAME == "CatalogSchema.PropertyType") {
                if(items.indexOf(result.data.ParentUuid)==-1){
                    items.push(result.data.ParentUuid)
                }
            }
        }
        
        // I will get the list of localisations for the packages related to those items.
        var packages = []
 
        var getItem =  function(items, packages, localisations, callback){
            var uuid = items.pop()
            server.entityManager.getEntityByUuid(uuid, false,
                // success callback
                function(item, caller){
                    if(item.M_packaged != undefined){
                        for(var i=0; i < item.M_packaged.length; i++){
                            if(caller.packages.indexOf(item.M_packaged[i]) == -1){
                                caller.packages.push(item.M_packaged[i]);
                            }
                        }
                    }
                    if(caller.items.length == 0){
                        caller.callback(caller.packages, caller.localisations)
                    }else{
                        getItem(caller.items, caller.packages, caller.localisations, caller.callback)
                    }
                },
                // error callback
                function(errObj, caller){
                    if(caller.items.length == 0){
                        caller.callback(caller.packages, caller.localisations)
                    }else{
                        getItem(caller.items, caller.packages, caller.localisations, caller.callback)
                    }
                },{"items":items, "packages":packages, "localisations":localisations, "callback":callback})
        }

        
        var inventories = [];
        
        var getPackage = function(packages, inventories, localisations, callback){
            var uuid = packages.pop()
            server.entityManager.getEntityByUuid(uuid, false,
                // success callback
                function(package_, caller){
                    if(package_.M_inventoried != undefined){
                        for(var i=0; i < package_.M_inventoried.length; i++){
                            if(caller.inventories.indexOf(package_.M_inventoried[i]) == -1){
                                caller.inventories.push(package_.M_inventoried[i]);
                            }
                        }
                    }
                    if(caller.packages.length == 0){
                        caller.callback(caller.inventories, caller.localisations)
                    }else{
                        getPackage(caller.packages, caller.inventories, caller.localisations, caller.callback)
                    }
                },
                // error callback
                function(errObj, caller){
                    if(caller.packages.length == 0){
                        caller.callback(caller.inventories, caller.localisations)
                    }else{
                        getPackage(caller.packages, caller.inventories, caller.localisations, caller.callback)
                    }
                },{"packages":packages, "inventories":inventories, "localisations":localisations, "callback":callback})
        }
        
        var getInventory = function(index, inventories, localisations, callback){
            if(inventories == undefined){
                callback(localisations, inventories)
                return;
            }
            if(inventories.length == 0){
                // No inventories here.
                callback(localisations, inventories)
                return;
            }
            var uuid = inventories[index]
            if(uuid == null || uuid == undefined){
                if(index < inventories.length){
                    getInventory(index + 1, inventories, localisations, callback)
                    return
                }else{
                    callback(localisations, inventories)
                    return
                }
            }
            server.entityManager.getEntityByUuid(uuid, false,
                function(inventory, caller){
                    if(inventory.M_located != undefined){
                        var exist = false
                        for(var i=0; i < caller.localisations.length; i++){
                            if(isObject(caller.localisations[i])){
                                if(caller.localisations[i].UUID == inventory.M_located){
                                    exist = true;
                                    break
                                }
                            }else if(caller.localisations[i] == inventory.M_located){
                                    exist = true;
                                    break                               
                            }
                        }
                        if(!exist){
                            caller.localisations.push(inventory.M_located)
                        }
                    }
                    caller.inventories[caller.index] = inventory
                    if(caller.index == caller.inventories.length - 1){
                        caller.callback(caller.localisations, caller.inventories)
                    }else{
                        getInventory(caller.index + 1, caller.inventories, caller.localisations, caller.callback)
                    }
                },
                function(errObj, caller){
                    if(caller.index == caller.inventories.length - 1){
                        caller.callback(caller.localisations, caller.inventories)
                    }else{
                        getInventory(caller.index + 1, caller.inventories, caller.localisations, caller.callback)
                    }
                }, {"inventories":inventories, "index":index, "localisations":localisations, "callback":callback})
        }
        
        var setInventory = function(index, localisations, callback){
            if(localisations[index].M_inventories == null){
                localisations[index].M_inventories = []
            }
            getInventory(0, localisations[index].M_inventories, localisations, 
                function(index,callback){
                    return function(localisations, inventories){
                        callback(localisations, inventories, index)
                    }
                }(index, callback))
        }
        
        var getLocalisation = function(index, localisations, inventories, callback){
            if(index == localisations.length){
                callback(localisations, inventories)
            }else if(!isObject(localisations[index])){
               // In that case I will get the localisation from the 
               // the server.
               var uuid = localisations[index]
               if(isObjectReference(uuid)){
                   server.entityManager.getEntityByUuid(localisations[index], false, 
                    function(localisation, caller){
                        caller.localisations[caller.index] = localisation
                        getLocalisation(caller.index + 1, caller.localisations, caller.inventories, caller.callback)
                    },
                    function(){
                        
                    },
                    {"index":index, "localisations":localisations, "inventories":inventories, "callback":callback})
               }else{
                   getLocalisation(index + 1, localisations, inventories, callback)
               }
            }else{
                setInventory(index, localisations, function(index, localisations, inventories, callback){
                    return function(){
                         getLocalisation(index, localisations, inventories, callback)
                    }
                }(index + 1, localisations, inventories, callback))
               
            }
        }
        
        if(items.length > 0){
            getItem(items, packages, localisations, function(searchResultPage){
                return function(packages, localisations){
                    getPackage(packages, inventories, localisations, function(inventories){
                         getInventory(0, inventories, localisations, function(localisations, inventories){
                             getLocalisation(0, localisations, inventories, function(localisations, inventories){
                                // here I will link the localisation and inventories...
                                for(var i=0; i < localisations.length; i++){
                                    if(localisations[i].M_inventories != undefined){
                                        for(var j=0; j < localisations[i].M_inventories.length; j++){
                                            for(var k=0; k < inventories.length; k++){
                                                if(inventories[k].UUID == localisations[i].M_inventories[j]){
                                                    localisations[i].M_inventories[j] = inventories[k];
                                                    break;
                                                }
                                            }
                                        }
                                    }
                                }
                                mainPage.searchResultPage.displayTabResults(localisations, query)
                             })
                         })
                    })
                }
            }(this))
            
        }else{
            // Here All inventory and sublocalisation must be initialyse.
           var index = 0;
            setInventory_ = function(index, localisations){
                if(index < localisations.length){
                    setInventory(index, localisations, 
                    function(localisations, inventories, index){
                        if(localisations[index].M_inventories != undefined){
                            for(var j=0; j < localisations[index].M_inventories.length; j++){
                                for(var k=0; k < inventories.length; k++){
                                    if(inventories[k].UUID == localisations[index].M_inventories[j]){
                                        localisations[index].M_inventories[j] = inventories[k];
                                        break;
                                    }
                                }
                            }
                        }
                        // set the next localisation inventory
                        setInventory_(index + 1, localisations)
                    })
                }else{
                    // display the result in the page.
                    mainPage.searchResultPage.displayTabResults(localisations, query) 
                }
            }
            setInventory_(0, localisations)
            
        }
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

    }else if(mainPage.searchContext == "localisationSearchLnk"){
        var localisationSearchResult = new localisationPage(this.resultsContent.getChildById(query+"-query"), results, query, function(query){
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
                fireResize()
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
	    
        for(var j=0; j < item.M_properties.length; j++){
            var property = entities[item.M_properties[j].UUID]
            // console.log(property.M_name)
            var row = properties.appendElement({"tag":"tr"}).down();
            row.appendElement({"tag":"td", "innerHtml":property.M_name})

            // Here I will display the value of the property.
            if(property.M_kind.M_valueOf == "boolean"){
                row.appendElement({"tag":"td", "innerHtml":property.M_booleanValue.toString()})
            }else if(property.M_kind.M_valueOf == "numeric"){
                row.appendElement({"tag":"td", "innerHtml":property.M_numericValue.toString()})
            }else if(property.M_kind.M_valueOf == "dimension"){
                // So in that particular case I will display the dimension type information.
                var cell = row.appendElement({"tag":"td"}).down()
                cell.appendElement({"tag":"span", "innerHtml":property.M_dimensionValue.M_valueOf.toString()})
                cell.appendElement({"tag":"span", "innerHtml":property.M_dimensionValue.M_unitOfMeasure.M_valueOf, "style":"padding-left: 4px;"})
            }else{
                row.appendElement({"tag":"td", "innerHtml":property.M_stringValue})
            }
            
        }
    }
    
    // call the callback
    this.callback(this.query)
}

