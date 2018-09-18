function displayLocalisation(parent, localisationId, onClick){
    var getLocalisationIds = function(id, localisations, callback, div){
                            
    server.entityManager.getEntityByUuid(id, false,
        function(localisation, caller){
            caller.localisations.unshift(localisation)
            if(localisation.M_parent != ""){
                getLocalisationIds(localisation.M_parent, caller.localisations, caller.callback, caller.div, caller.onClick)
            }else{
                var div = caller.div
                
                var breadcrumb = div.appendElement({"tag" : "nav", "aria-label" : "breadcrumb"}).down()
                .appendElement({"tag" : "ol", "class" : "breadcrumb", "style" : "background-color : transparent;"}).down()
                
                // The position
                div.element.style.position = "relative"
                
                var packages = []
                // I will also append the print label button.
                var printBtn = div.appendElement({"tag":"div", "style":"position:absolute; right: 10px; top: 5px;", "title":"imprimer les étiquettes."}).down()
                    .appendElement({"tag":"i", "class":"fa fa-print"}).down()
                
                printBtn.element.onmouseenter = function(){
                    this.style.cursor = "pointer"
                }
                 
                printBtn.element.onmouseleave = function(){
                    this.style.cursor = "default"
                }
                
                printBtn.element.onclick = function(localisation){
                    return function(){
                        // print the localisation information.
                        printLocalisationLabels(localisation)
                    }
                }(caller.localisations[caller.localisations.length-1])
                    
                for(var i=0; i < caller.localisations.length; i++){
                    var id = caller.localisations[i].getEscapedUuid() + "_localisation_display"
                    var lnk = breadcrumb.appendElement({"tag" : "li", "class" : "breadcrumb-item"}).down()
                    .appendElement({"tag" : "a", "href" : "#" + id, "innerHtml" : caller.localisations[i].M_name}).down()
                    
                    if(caller.localisations.M_subLocalisations != undefined && caller.localisations.M_parent != null){
                        lnk.element.id = id;
                    }

                    // generic onclick wiht localisation.
                    lnk.element.onclick = function(onClick, localisation, id){
                        return function(){
                            // Hide other page if there is there.
                            if(mainPage.welcomePage !== null){
                                mainPage.welcomePage.element.style.display = "none"
                            }
                            if(mainPage.personalPage !== null){
                                mainPage.personalPage.element.style.display = "none"
                            }
                            if(mainPage.itemDisplayPage !== null){
                                mainPage.itemDisplayPage.panel.element.style.display = "none"
                            }
                            if(mainPage.packageDisplayPage !== null){
                                mainPage.packageDisplayPage.panel.element.style.display = "none"
                            }
                            if(mainPage.orderPage !== null){
                                mainPage.orderPage.panel.element.style.display = "none"
                            }
                            if(mainPage.adminPage !== null){
                                mainPage.adminPage.panel.element.style.display = "none"
                            }
                            
                            // Show the search result page.
                            if(mainPage.searchResultPage !== null){
                                mainPage.searchResultPage.panel.element.style.display = ""
                            }
                            
                            if(document.getElementById(id) == undefined){
                                onClick(localisation)
                            }else{
                              // Activate the item tab and display the localisation information. 
                               function click_localisation_tab(localisation){
                                   var id = localisation.M_id + "-tab"
                                   if(document.getElementById(id) != undefined){
                                       document.getElementById(id).click()
                                       return
                                   }
                                   if(localisation.M_parent != null){
                                       server.entityManager.getEntityByUuid(localisation.M_parent, true, 
                                        function(localisation, caller){
                                            click_localisation_tab(localisation)
                                        },
                                        function(){
                                            
                                        }, {})
                                   }
                               }
                               // click on lolcalisation.
                               click_localisation_tab(localisation)
                            }
                        }
                    }(caller.onClick, caller.localisations[i], id)
                    
                }
            }
        },
        function(){
            
        },{"localisations":localisations, "callback":callback, "div":div, "onClick":onClick})
    }
    
    getLocalisationIds(localisationId, [], getLocalisationIds, parent) 
}

/**
 * That page display the localisations and their sub-localisations or inventories.
 */
var localisationPage = function(parent, localisations, query, callback){
    
    // The panel result panel.
    parent.element.style.width = "100%";
    this.panel = parent.appendElement({"tag":"div", "class":"container-fluid", "id" : query + "_search_result"}).down()
    

    // Here I will display the localisations.
    for(var i=0; i < localisations.length; i++){
        if(isObject(localisations[i])){
            new LocalisationDisplay(this.panel, localisations[i])
        }
    }
    
    // call when the initialisation is done.
    callback(query)
    
    return this;
}

var LocalisationDisplay = function(parent, localisation){
    this.id = localisation.getEscapedUuid() + "_localisation_display";
    if(parent.getChildById(this.id)!==null){
        this.panel = parent.getChildById(this.id);
    }else{
        this.panel = parent.appendElement({"tag":"table", "id":this.id, "class":"table table-hover table-condensed .table-striped"}).down()
    }
    
    // So here I will display the the component header...
    this.header = this.panel.appendElement({"tag":"thead"}).down()
    
    // Display the localisation in the header.
    displayLocalisation(this.header.appendElement({"tag":"tr", "class":"thead-dark"}).down().appendElement({"tag":"th", "colspan":"3"}).down(), localisation.UUID, 
        function(location){
            mainPage.searchContext = "localisationSearchLnk";
            mainPage.searchResultPage.displayResults({"results":[{"data":location}], "estimate":1}, location.M_id, "localisationSearchLnk") 
        })
    
    this.body = this.panel.appendElement({"tag":"tbody"}).down()
    if(localisation.M_subLocalisations.length > 0){
        // In that case I will display the sub-localisations.
        var getSublocalisation = function(index, subLocalisations, div){
            var uuid = subLocalisations[index]
            if(uuid != undefined){
                server.entityManager.getEntityByUuid(uuid, false,
                    // success callback.
                    function(result, caller){
                        if(result.M_inventories.length > 0){
                            // Here I need to initialyse the list of inventories.
                            var setInventory = function(index, inventories, localisation, div, subLocalisationIndex, subLocalisations){
                                var uuid = inventories[index]
                                if(isObject(uuid)){
                                    uuid = uuid.UUID
                                }
                                server.entityManager.getEntityByUuid(uuid, false, 
                                    function(result, caller){
                                        caller.localisation.M_inventories[caller.index] = result
                                        if(caller.index < caller.inventories.length - 1){
                                            setInventory(caller.index + 1, caller.inventories, caller.localisation, caller.div, caller.subLocalisationIndex, caller.subLocalisations)
                                        }else{
                                            // All is done...
                                            new LocalisationDisplay(caller.div, caller.localisation);
                                            if(caller.subLocalisationIndex < caller.subLocalisations.length - 1){
                                                getSublocalisation(caller.subLocalisationIndex + 1, caller.subLocalisations, caller.div)
                                            }
                                        }
                                    },
                                    function(errObj, caller){
                                        
                                    }, {"index":index, "inventories":inventories, "localisation":localisation, "div":div, "subLocalisationIndex":subLocalisationIndex, "subLocalisations":subLocalisations})
                            
                            }
                            // set the inventory...
                            setInventory(0, result.M_inventories, result, caller.div, caller.index, caller.subLocalisations)
                            
                        } else {
                            new LocalisationDisplay(caller.div, result);
                            if(caller.index < caller.subLocalisations.length - 1){
                                getSublocalisation(caller.index + 1, caller.subLocalisations, caller.div)
                            }
                        }
                    },
                    // error callback
                    function(errObj, caller){
                        
                    }, 
                    {"index":index, "subLocalisations":subLocalisations, "div":div})
            }else{
                getSublocalisation(caller.index + 1, caller.subLocalisations, caller.div)
            }
        }
        
        if(localisation.M_subLocalisations.length> 0){
            getSublocalisation(0, localisation.M_subLocalisations, this.body)
        }
        
     } else if(localisation.M_inventories !== undefined){
       if(localisation.M_inventories.length > 0){
           this.header.appendElement({"tag":"tr"}).down()
            .appendElement({"tag":"th", "innerHtml":"Package"})
            .appendElement({"tag":"th", "innerHtml":"safety stock"})
            .appendElement({"tag":"th", "innerHtml":"Qty"})
            for(var i=0; i < localisation.M_inventories.length; i++){
                if(isObject(localisation.M_inventories[i])){
                    var row = this.body.appendElement({"tag":"tr", "style":"border-bottom: 1px solid #dfe3e7;"}).down()
                    var packageDiv = row.appendElement({"tag":"tbody"}).down()
                    row.appendElement({"tag":"td", "innerHtml":localisation.M_inventories[i].M_safetyStock})
                    row.appendElement({"tag":"td", "innerHtml":localisation.M_inventories[i].M_quantity})
                    
                    // Now I will get the package information...
                    server.entityManager.getEntityByUuid(localisation.M_inventories[i].M_package, false,
                        function(result, caller){
                            // So here I will display information related to package...
                            var div = caller.div
                            div.appendElement({"tag":"tr"}).down()
                            .appendElement({"tag":"th","scope":"row", "innerHtml":"id"})
                            .appendElement({"tag":"td", "id":result.getEscapedUuid() + "_lnk", "innerHtml":result.M_id}).up()
                            .appendElement({"tag":"tr"}).down()
                            .appendElement({"tag":"th","scope":"row", "innerHtml":"Name"})
                            .appendElement({"tag":"td", "innerHtml":result.M_name}).up()
                            .appendElement({"tag":"tr"}).down()
                            .appendElement({"tag":"th","scope":"row", "innerHtml":"Unit"})
                            .appendElement({"tag":"td", "innerHtml":result.M_unitOfMeasure.M_valueOf}).up()
                            
                            var id = div.getChildById(result.getEscapedUuid() + "_lnk")
                            
                            id.element.style.textDecoration = "underline"
                            id.element.onmouseover = function () {
                                this.style.color = "steelblue"
                                this.style.cursor = "pointer"
                            }
                        
                            id.element.onmouseleave = function () {
                                this.style.color = ""
                                this.style.cursor = "default"
                            }
                        
                            // Now the onclick... 
                            id.element.onclick = function (entity) {
                                return function () {
                                    mainPage.packageDisplayPage.displayTabItem(entity)
                                }
                            }(entities[result.UUID])
                        },
                        function(){
                            
                        }, {"div":packageDiv})
                }
            }
       }
    }
    
    return this;
}
