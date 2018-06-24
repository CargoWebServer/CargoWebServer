/**
 *  That class contain the code to admin item.
 */
var AdminItemPage = function(panel){
    /** The tab content panel **/
    this.panel = panel
    
    /** The panel sections. **/
    this.itemsPanel = this.panel.appendElement({"tag":"div", "class":"row", "style":"height:85vh;margin:0;"}).down()
        .appendElement({"tag":"div","class":"col-md-3 bg-dark", "id":"itemAdminNavigation"}).down()
        .appendElement({"tag":"div","class":"input-group mb-3", "style":"padding:15px;"}).down()
        .appendElement({"tag":"div", "class":"input-group-prepend", "id":"item_filter_btn"}).down()
        .appendElement({"tag":"div", "class":"input-group-text"}).down()
        .appendElement({"tag":"i", "class":"fa fa-search text-dark"}).up().up()
        .appendElement({"tag":"input", "class":"form-control", "id":"items_filter_input", "style":"text", "placeholder":"Filtrer"}).up()
        .appendElement({"tag":"div", "class":"list-group", "role":"tablist", "id":"items_filter_lst", "style":"overflow-y: auto;"})
        .appendElement({"tag":"button","class":"btn btn-success", "id":"new_item_btn", "style":"display: flex; position: absolute; height: 36px; bottom: 30px; right: 30px;"}).down()
        .appendElement({"tag":"i", "class":"fa fa-plus"}).up().up()
        .appendElement({"tag":"div", "class":"col-md-9 bg-light", "style":"overflow-y: auto;"}).down()
        .appendElement({"tag":"div", "class":"tab-content", "id":"admin_item_display"}).down()

    this.searchBar = this.panel.getChildById("items_filter_input")
    this.searchBtn =  this.panel.getChildById("item_filter_btn")
    
    this.newItemBtn = this.panel.getChildById("new_item_btn")
    
    this.newItemBtn.element.onclick = function(adminItemPage){
        return function(){
            // So here I will create a new item and load it page...
            adminItemPage.panel.getChildById("admin_item_display").removeAllChilds()
            adminItemPage.panel.getChildById("items_filter_lst").removeAllChilds()
            
            // So here i will use the item display div to display the new 
            // item to create.
            var newItemCreationPanel = adminItemPage.panel.getChildById("admin_item_display")
                .appendElement({"tag":"div", "class":"tab-pane active show", "style":"padding: 15px;" })
                
            var idInput = newItemCreationPanel.appendElement({"tag":"div", "class":"input-group mb-3"}).down()
                .appendElement({"tag":"div","class":"input-group-prepend"}).down()
                .appendElement({"tag":"span", "class":"input-group-text", "innerHtml":"id"}).up()
                .appendElement({"tag":"input", "class":"form-control"}).down()
                
            var manufacturerIdInput = newItemCreationPanel.appendElement({"tag":"div", "class":"input-group mb-3"}).down()
                .appendElement({"tag":"div","class":"input-group-prepend"}).down()
                .appendElement({"tag":"span", "class":"input-group-text", "innerHtml":"Manufacturer"}).up()
                .appendElement({"tag":"input", "class":"form-control"}).down()
                
            // The save button.
            var createBtn = newItemCreationPanel.appendElement({"tag":"button", "class":"btn btn-primary mr-3", "innerHtml":"Enregistrer"}).down()
            
            // set the focus to the id...
            idInput.element.focus()
            
            createBtn.element.onclick = function(idInput, manufacturerIdInput){
                return function(){
                    var id = idInput.element.value;
                    var manufacturerId = manufacturerIdInput.element.value
                    if(manufacturerId.length === 0){
                        manufacturerId = "NA";
                    }
                    
                    if(id.length === 0){
                        // TODO display error message here.
                    }else{
                        spinner.panel.element.style.display = "";
                        server.entityManager.getEntityById("CatalogSchema.ManufacturerType", "CatalogSchema", [manufacturerId], true,
                                // Success callback.
                                function(supplier, caller){
                                    var item = new CatalogSchema.ItemType()
                                    item.M_id = caller.id;
                                    server.entityManager.createEntity(supplier.UUID, "M_items", item, 
                                    function(item, caller){
                                        // Set the item.
                                        mainPage.adminPage.adminItemPage.panel.getChildById("admin_item_display").removeAllChilds()
                                        mainPage.adminPage.adminItemPage.panel.getChildById("items_filter_lst").removeAllChilds()
                                        
                                        // Set the active item.
                                        var itemPanel = mainPage.adminPage.adminItemPage.setItem(item)
                                                                                
                                        // Reset the display attribute.
                                        var itemsPanel = document.getElementsByClassName("item_panel");
                                        for(var i=0; i < itemsPanel.length; i++){
                                            itemsPanel[i].style.display = ""
                                        }
                                        
                                        itemPanel.element.style.display = "block"
                                        
                                        // Set the focus to the name input.
                                        var nameInputs = document.getElementsByName(item.getEscapedUuid() + "_M_name")
                                        for(var i=0; i < nameInputs.length; i++){
                                            if(server.languageManager.language == nameInputs[i].lang){
                                                nameInputs[i].focus()
                                                break;
                                            }
                                        }
                                        spinner.panel.element.style.display = "none";
                                    },
                                    function(){
                                        spinner.panel.element.style.display = "none";
                                    }, {})
                                    
                                },
                                // Error callback
                                function(errObj){
                                    spinner.panel.element.style.display = "none";
                                },{"id":id}
                            )
                        
                    }
                }
            }(idInput, manufacturerIdInput)
        }
    }(this)
    
    var itemFilterLst = this.panel.getChildById("items_filter_lst")
    window.addEventListener('resize', function(itemFilterLst){
        return function(evt){
            var h = itemFilterLst.element.parentNode.offsetHeight - itemFilterLst.element.parentNode.firstChild.offsetHeight - 96;
            itemFilterLst.element.style.height = h + "px";
        }
    }(itemFilterLst));

    var searchFct = function(searchBar){
        return function(){
            // So here I will populate the search values...
            var keyword = searchBar.element.value;
            var fields = []
            
            // First of a ll I will clear the search panel.
            xapian.search(
                dbpaths,
                keyword.toUpperCase(),
                fields,
                "en",
                0,
                1000,
                // success callback
                function (results, caller) {
                    var uuids = []
                    for (var i = 0; i < results.results.length; i++) {
                            var result = results.results[i];
                            if (result.data.TYPENAME == "CatalogSchema.ItemType") {
                                if(uuids.indexOf(result.data.UUID) == -1){
                                    uuids.push(result.data.UUID)
                                }
                            }else if(result.data.TYPENAME == "CatalogSchema.PropertyType"){
                                if(uuids.indexOf(result.data.ParentUuid) == -1){
                                    uuids.push(result.data.ParentUuid)
                                }
                            }
                    }
                    function getItems(uuids){
                        var uuid = uuids.pop()
                        server.entityManager.getEntityByUuid(uuid, false, 
                            function(item, caller){
                                var txt = item.M_id
                                for(var i=0; i < item.M_alias.length; i++){
                                    if(i==0){
                                        txt += " (";
                                    }
                                    txt += " " +  item.M_alias[i]
                                    if(i < item.M_alias.length - 2){
                                        txt += ","
                                    }
                                    if(i==item.M_alias.length-1){
                                        txt += " )";
                                    }
                                }
                                var lnk = mainPage.adminPage.adminItemPage.panel.getChildById("items_filter_lst").appendElement({"tag":"a", "class" : "list-group-item list-group-item-action","id" : item.getEscapedUuid() + "-selector","data-toggle" : "tab","href" : "#"+ item.getEscapedUuid() + "-control","role":"tab", "innerHtml": txt,"aria-controls":  item.getEscapedUuid() + "-control"}).down()
                                lnk.element.onclick = function(item){
                                    return function(){
                                        // reset item_panel properties.
                                        var itemsPanel = document.getElementsByClassName("item_panel");
                                        for(var i=0; i < itemsPanel.length; i++){
                                            itemsPanel[i].style.display = ""
                                        }
                                        
                                        // create the display once.
                                        if(document.getElementById(item.getEscapedUuid() + "-control") == undefined) {
                                            mainPage.adminPage.adminItemPage.setItem(item)
                                        }
                                    }
                                }(item)
                                
                                if(caller.uuids.length > 0){
                                    getItems(caller.uuids)
                                }
                            },
                            function(errObj, caller){
                                if(caller.uuids.length > 0){
                                    getItems(caller.uuids)
                                }
                            }, {"uuids":uuids})
                    }
                    getItems(uuids)
                    fireResize()
                },
                // error callback
                function () {
        
                }, {"keyword":keyword, "searchBar":searchBar})
        }
    }(this.searchBar)
    
    this.searchBtn.element.onclick = searchFct
    
    this.searchBar.element.onkeyup = function(searchFct, itemFilterLst, itemsPanel){
        return function(evt){
            if(evt.keyCode == 13){
                itemFilterLst.removeAllChilds()
                itemsPanel.removeAllChilds()
                searchFct()
            }else if(this.value.length === 0){
                // In that case the result was clear...
                itemFilterLst.removeAllChilds()
                itemsPanel.removeAllChilds()
            }
        }
    }(searchFct, itemFilterLst, this.itemsPanel)
    
    return this
}

AdminItemPage.prototype.loadAdminControl = function(item){
    var itemPanel = this.itemsPanel.getChildById(item.getEscapedUuid() + "-control");
    if( itemPanel == null){
        itemPanel = this.itemsPanel.appendElement({"tag":"div", "class":"tab-pane item_panel", "id" : item.getEscapedUuid() + "-control", "aria-labelledby":item.getEscapedUuid() +"-selector", "style":"padding:15px;"}).down()
    }
    
    return itemPanel;
}

function setSaveCancelBtn(item){
    var saveBtn = document.getElementById(item.getEscapedUuid() + "_save_btn")
    var cancelBtn = document.getElementById(item.getEscapedUuid() + "_cancel_btn")
    var itemAdminTabTitle = document.getElementById("item-adminTab").firstChild
    if(!itemAdminTabTitle.innerHTML.endsWith("*")){
        itemAdminTabTitle.innerHTML = itemAdminTabTitle.innerHTML + "*"
    }
    saveBtn.classList.remove("disabled")
    cancelBtn.classList.remove("disabled")
}  

AdminItemPage.prototype.setItem = function(item){
       
    var itemPanel = this.loadAdminControl(item)
    itemPanel.removeAllChilds()
    
    // Now I will display the properties of the items.
    // The id, unmutable.
    itemPanel.appendElement({"tag":"div", "class":"row mb-3"}).down()
        .appendElement({"tag":"div", "class":"col-sm-6", "id":"general_infos_div"})
        .appendElement({"tag":"div", "class":"col-sm-6", "id":"picture_div"})
        
    var generalInfoDiv = itemPanel.getChildById("general_infos_div")
    var pictureDiv = itemPanel.getChildById("picture_div")
    
    // Here I will set the item picture selector.
    new ImagePicker(pictureDiv,"/Catalogue_2/photo/" + item.M_id)
     
    generalInfoDiv.appendElement({"tag":"div", "class":"input-group mb-3"}).down()
    .appendElement({"tag":"div","class":"input-group-prepend"}).down()
    .appendElement({"tag":"a","href":"#", "class":"input-group-text", "innerHtml":"id", "id":"item_lnk"}).up()
    .appendElement({"tag":"span", "class":"form-control", "innerHtml":item.M_id})
    
    // Link to the item.
    generalInfoDiv.getChildById("item_lnk").element.onclick = function(item){
        return function(){
            spinner.panel.element.style.display = "";
            mainPage.itemDisplayPage.displayTabItem(item)
            spinner.panel.element.style.display = "none";
        }
    }(item)
    
    // The Type name.
    var typeName = generalInfoDiv.appendElement({"tag":"div", "class":"input-group mb-3"}).down()
    .appendElement({"tag":"div","class":"input-group-prepend"}).down()
    .appendElement({"tag":"span", "class":"input-group-text", "innerHtml":"Nom"}).up()
 
    // Set the input
    var editors = initMultilanguageInput(typeName, "input", item, "M_name")
    for(var id in editors){
        editors[id].element.onchange = function(item){
            return function(){
                setSaveCancelBtn(item)            
            }
        }(item)
    }
    
    // The description
    var desciptions = generalInfoDiv.appendElement({"tag":"div", "class":"input-group"}).down()
    .appendElement({"tag":"div","class":"input-group-prepend"}).down()
    .appendElement({"tag":"span", "class":"input-group-text", "innerHtml":"Description"}).up()
    
    // Set the decription.
    editors_ = initMultilanguageInput(desciptions, "textarea", item, "M_description")
    for(var id in editors_){
        editors_[id].element.onchange = function(item){
            return function(){
                setSaveCancelBtn(item)
            }
        }(item)
    }
    
    // Now the alias...
    itemPanel.appendElement({"tag":"div", "class":"row mb-3"}).down()
    .appendElement({"tag":"div", "class":"col-sm-2", "style":"padding-right: 0px; padding-left: 0px;"}).down()
    .appendElement({"tag":"ul", "class":"list-group borderless", "id":"alias"}).up()
    .appendElement({"tag":"div", "class":"col-sm-3", "style":"padding-right: 0px; padding-left: 0px;"}).down()
    .appendElement({"tag":"ul", "class":"list-group borderless", "id":"keywords"}).up()
    .appendElement({"tag":"div", "class":"col-sm-5", "style":"padding-right: 0px; padding-left: 0px;"}).down()
    .appendElement({"tag":"ul", "class":"list-group borderless", "id":"comments"}).up()
    .appendElement({"tag":"div", "class":"col-sm-2", "style":"padding-right: 0px; padding-left: 0px;"}).down()
    .appendElement({"tag":"ul", "class":"list-group borderless", "id":"equivalents"})
    
    // List of alia(s)
    var alias = itemPanel.getChildById("alias")
    // List of comment's
    var comments = itemPanel.getChildById("comments")
    // List of keyword's
    var keywords = itemPanel.getChildById("keywords")
    // List of equivalent product.
    var equivalents = itemPanel.getChildById("equivalents")
    
    // The equivalentcy...
    var equivalentInput = equivalents.appendElement({"tag":"li", "class":"list-group-item", "style" : "background-color: #e9ecef; border: 1px solid #ced4da; color: #495057; position: relative;"}).down()
        .appendElement({"tag":"input", "class":"form-control", "type":"text", "placeholder":"Equivalent(s)"}).down()
    
    // Append an equivalent item to a list.
    function appendEquivalent(item, uuid, lst){
        // So here I will display the equivalency...
        getEntityIdsFromUuid(uuid, function(lst, item, uuid){
            return function(ids){
                // So here I will set the lnk to the item...
                var li = lst.appendElement({"tag":"li", "class":"list-group-item"}).down()
                var lnk = li.appendElement({"tag":"a", "href":"#", "innerHtml":ids[1]}).down()
                var deleteBtn = li.appendElement({"tag":"span"}).down()
                deleteBtn.appendElement({"tag":"i", "class":"fa fa-trash", "style":"color: #495057; padding-left: 7px"})
                deleteBtn.element.onmouseenter = function(){
                    this.style.cursor = "pointer"
                }
                deleteBtn.element.onmouseleave = function(){
                    this.style.cursor = "default"
                }
                // The delete action...
                deleteBtn.element.onclick = function(uuid, item, li){
                    return function(){
                        // remove the li from the ul.
                        li.element.parentNode.removeChild(li.element)
                        // Now I will remove it from the item equivalents...
                        item.M_equivalents.splice(item.M_equivalents.indexOf(uuid), 1)
                        // display the save and cancel button...
                        setSaveCancelBtn(item)
                    }
                }(uuid, item, li)
                
                // Now the lnk action...
                lnk.element.onclick = function(uuid){
                    return function(){
                        spinner.panel.element.style.display = "";
                        server.entityManager.getEntityByUuid(uuid, false, 
                            function(entity, caller){
                                mainPage.itemDisplayPage.displayTabItem(entity)
                                spinner.panel.element.style.display = "none";
                            },
                            function(){
                                
                            }, {})
                    }
                }(uuid)
                
            }
        }(lst, item, uuid))            
    }
    
    // Append new equivalency..
    equivalentInput.element.oncomplete = function(item, equivalents){
        return function(uuid){
            this.value = "";
            if(item.M_equivalents.indexOf(uuid) == -1){
                item.M_equivalents.push(uuid) // set in the item.
                appendEquivalent(item, uuid, equivalents)
                // display the save and cancel button...
                setSaveCancelBtn(item)
            }
        }
    }(item, equivalents)
    
    // I will append the list of existing equivalent items.
    for(var i=0; i<item.M_equivalents.length; i++){
        appendEquivalent(item,  item.M_equivalents[i], equivalents)
    }
    
    autocomplete("CatalogSchema.ItemType", equivalentInput)
    
    // The alias
    var appendAliasBtn = alias.appendElement({"tag":"li", "class":"list-group-item", "style" : "background-color: #e9ecef; border: 1px solid #ced4da; color: #495057; position: relative;"}).down()
        .appendElement({"tag":"span", "innerHtml":"Alias"})
        .appendElement({"tag":"i", "class":"fa fa-plus", "style":"position: absolute;right: 12px;top: 17px;"}).down()
        
    for(var i=0; i < item.M_alias.length; i++){
        var row = alias.appendElement({"tag":"li", "class":"list-group-item"}).down()
        var deletBtn = row.appendElement({"tag":"div", "class":"input-group"}).down()
            .appendElement({"tag":"input", "class":"form-control","type":"text", "id":"row_content", "value":item.M_alias[i]})
            .appendElement({"tag":"div","class":"input-group-append"}).down()
            .appendElement({"tag":"span","class":"input-group-text"}).down()
            
        deletBtn.appendElement({"tag":"i","class":"fa fa-trash"}).down()
        
        deletBtn.element.onmouseenter = function(){
            this.style.cursor = "pointer"
        }
        
        deletBtn.element.onmouseleave = function(){
            this.style.cursor = "default"
        }
        
        // Remove the row.
        deletBtn.element.onclick = function(row, index, item){
            return function(){
                row.element.parentNode.removeChild(row.element)
                item.M_alias.splice(index, 1);
                setSaveCancelBtn(item)
            }
        }(row, i, item)
        
        // Edit the content.
        row.getChildById("row_content").element.onchange = function(index, item){
            return function(){
                // Set the new value
                item.M_alias[index] = this.value;
                setSaveCancelBtn(item)
            }
        }(i, item)
    }
    
    // The keywords
    var appendKeywordBtn = keywords.appendElement({"tag":"li", "class":"list-group-item", "style" : "background-color: #e9ecef; border: 1px solid #ced4da; color: #495057; position: relative;"}).down()
        .appendElement({"tag":"span", "innerHtml":"Mot(s) clés"})
        .appendElement({"tag":"i", "class":"fa fa-plus", "style":"position: absolute;right: 12px;top: 17px;"}).down()
        
    for(var i=0; i < item.M_keywords.length; i++){
        var row =  keywords.appendElement({"tag":"li", "class":"list-group-item"}).down()
        var deletBtn = row.appendElement({"tag":"div", "class":"input-group"}).down()
            .appendElement({"tag":"input", "class":"form-control", "id":"row_content", "type":"text", "value":item.M_keywords[i]})
            .appendElement({"tag":"div","class":"input-group-append"}).down()
            .appendElement({"tag":"span","class":"input-group-text"}).down()
            
        deletBtn.appendElement({"tag":"i","class":"fa fa-trash"}).down()
        
        deletBtn.element.onmouseenter = function(){
            this.style.cursor = "pointer"
        }
        
        deletBtn.element.onmouseleave = function(){
            this.style.cursor = "default"
        }
        
        // Remove the row.
        deletBtn.element.onclick = function(row, index, item){
            return function(){
                row.element.parentNode.removeChild(row.element)
                item.M_keywords.splice(index, 1);
                setSaveCancelBtn(item)
            }
        }(row, i, item)
        
        row.getChildById("row_content").element.onchange = function(index, item){
            return function(){
                // Set the new value
                item.M_keywords[index] = this.value;
                setSaveCancelBtn(item)
            }
        }(i, item)
    }
    
    // The comments.
    var appendCommentBtn = comments.appendElement({"tag":"li", "class":"list-group-item", "style" : "background-color: #e9ecef; border: 1px solid #ced4da; color: #495057; position: relative;"}).down()
     .appendElement({"tag":"span", "innerHtml":"Commentaire(s)"})
     .appendElement({"tag":"i", "class":"fa fa-plus", "style":"position: absolute;right: 12px;top: 17px;"}).down()
     
    for(var i=0; i < item.M_comments.length; i++){
        var row = comments.appendElement({"tag":"li", "class":"list-group-item"}).down()
        var deletBtn = row.appendElement({"tag":"div", "class":"input-group"}).down()
            .appendElement({"tag":"textarea", "class":"form-control", "id":"row_content", "innerHtml":item.M_comments[i]})
            .appendElement({"tag":"div","class":"input-group-append"}).down()
            .appendElement({"tag":"span","class":"input-group-text"}).down()
        
        deletBtn.appendElement({"tag":"i","class":"fa fa-trash"}).down()
        
        deletBtn.element.onmouseenter = function(){
            this.style.cursor = "pointer"
        }
        
        deletBtn.element.onmouseleave = function(){
            this.style.cursor = "default"
        }
        
        // Remove the row.
        deletBtn.element.onclick = function(row, index, item){
            return function(){
                row.element.parentNode.removeChild(row.element)
                item.M_comments.splice(index, 1);
                setSaveCancelBtn(item)
            }
        }(row, i, item)
        
        // Edit the content.
        row.getChildById("row_content").element.onchange = function(index, item){
            return function(){
                // Set the new value
                item.M_comments[index] = this.value;
                setSaveCancelBtn(item)
            }
        }(i, item)
    }
    
    appendCommentBtn.element.onmouseenter = appendAliasBtn.element.onmouseenter = appendKeywordBtn.element.onmouseenter = function(){
        this.style.cursor = "pointer"
    }
    
    appendCommentBtn.element.onmouseleave = appendAliasBtn.element.onmouseleave = appendKeywordBtn.element.onmouseleave = function(){
        this.style.cursor = "default"
    }
    
    appendKeywordBtn.element.onclick = function(keywords, item){
        return function(){
            var row =  keywords.appendElement({"tag":"li", "class":"list-group-item"}).down()
            var deletBtn = row.appendElement({"tag":"div", "class":"input-group"}).down()
                .appendElement({"tag":"input", "class":"form-control", "id":"row_content", "type":"text"})
                .appendElement({"tag":"div","class":"input-group-append"}).down()
                .appendElement({"tag":"span","class":"input-group-text"}).down()
                
            deletBtn.appendElement({"tag":"i","class":"fa fa-trash"}).down()
            
            deletBtn.element.onmouseenter = function(){
                this.style.cursor = "pointer"
            }
            
            deletBtn.element.onmouseleave = function(){
                this.style.cursor = "default"
            }
            
            var i = item.M_keywords.length
            item.M_keywords.push("")
            
            // Remove the row.
            deletBtn.element.onclick = function(row, index, item){
                return function(){
                    row.element.parentNode.removeChild(row.element)
                    item.M_keywords.splice(index, 1);
                    setSaveCancelBtn(item)
                }
            }(row, i, item)
            
            row.getChildById("row_content").element.focus()
            
            row.getChildById("row_content").element.onchange = function(index, keywords, item){
                return function(){
                    // Set the new value
                    keywords[index] = this.value;
                    setSaveCancelBtn(item)
                }
            }(i, item.M_keywords, item)
            
            setSaveCancelBtn(item)
        }
    }(keywords, item)
    
    appendAliasBtn.element.onclick = function(alias, item){
        return function(){
            var row = alias.appendElement({"tag":"li", "class":"list-group-item"}).down()
            var deletBtn = row.appendElement({"tag":"div", "class":"input-group"}).down()
                .appendElement({"tag":"input", "class":"form-control","type":"text", "id":"row_content", "value":""})
                .appendElement({"tag":"div","class":"input-group-append"}).down()
                .appendElement({"tag":"span","class":"input-group-text"}).down()
                
            deletBtn.appendElement({"tag":"i","class":"fa fa-trash"}).down()
            
            deletBtn.element.onmouseenter = function(){
                this.style.cursor = "pointer"
            }
            
            deletBtn.element.onmouseleave = function(){
                this.style.cursor = "default"
            }
            
            var i = item.M_alias.length
            item.M_alias.push("")
            
            // Remove the row.
            deletBtn.element.onclick = function(row, index, item){
                return function(){
                    row.element.parentNode.removeChild(row.element)
                    item.M_alias.splice(index, 1);
                    setSaveCancelBtn(item)
                }
            }(row, i, item)
            
            row.getChildById("row_content").element.focus()
            
            // Edit the content
            row.getChildById("row_content").element.onchange = function(index, item){
                return function(){
                    // Set the new value
                    item.M_alias[index] = this.value;
                    setSaveCancelBtn(item)
                }
            }(i, item)
            
            setSaveCancelBtn(item)
        }
    }(alias, item)
    
    appendCommentBtn.element.onclick = function(comments, item){
        return function(){
            var row = comments.appendElement({"tag":"li", "class":"list-group-item"}).down()
            var deletBtn = row.appendElement({"tag":"div", "class":"input-group"}).down()
                .appendElement({"tag":"textarea", "class":"form-control", "id":"row_content"})
                .appendElement({"tag":"div","class":"input-group-append"}).down()
                .appendElement({"tag":"span","class":"input-group-text"}).down()
                
            row.getChildById("row_content").element.focus()
            
            deletBtn.appendElement({"tag":"i","class":"fa fa-trash"}).down()
            var i = item.M_comments.length
            item.M_comments.push("")
            
            deletBtn.element.onmouseenter = function(){
                this.style.cursor = "pointer"
            }
            
            deletBtn.element.onmouseleave = function(){
                this.style.cursor = "default"
            }
            
            // Remove the row.
            deletBtn.element.onclick = function(row, index, item){
                return function(){
                    row.element.parentNode.removeChild(row.element)
                    item.M_comments.splice(index, 1);
                    setSaveCancelBtn(item)
                }
            }(row, i, item)
            
            // Edit the content.
            row.getChildById("row_content").element.onchange = function(index, item){
                return function(){
                    // Set the new value
                    item.M_comments[index] = this.value;
                    setSaveCancelBtn(item)
                }
            }(i, item)
            
            setSaveCancelBtn(item)
        }
    }(comments, item)  
    
    // Now The properties table.
    var properties = itemPanel.appendElement({"tag":"div", "class":"row mb-3"}).down()
        .appendElement({"tag":"div", "class":"col-sm-12 mb-3"}).down()
        .appendElement({"tag":"table", "class":"table table-hover"}).down()
    
    // The headers section.
    var appendPropertieBtn = properties.appendElement({"tag":"thead"}).down()
        .appendElement({"tag":"tr", "class":"row"}).down()
        .appendElement({"tag":"th", "scope":"col", "class":"col-sm-1", "innerHtml":"Id"})
        .appendElement({"tag":"th", "scope":"col", "class":"col-sm-2", "innerHtml":"Name"})
        .appendElement({"tag":"th", "scope":"col", "class":"col-sm-3", "innerHtml":"Description"})
        .appendElement({"tag":"th", "scope":"col", "class":"col-sm-5", "innerHtml":"Value"})
        .appendElement({"tag":"th", "scope":"col", "class":"col-sm-1"}).down()
        .appendElement({"tag":"button", "class":"btn btn-success btn-sm", "style":"display: inline-flex; height: 29px;"}).down()
        // the icon.
        appendPropertieBtn.appendElement({"tag":"i", "class":"fa fa-plus"})
        

    var body = properties.appendElement({"tag":"tbody"}).down()
    
      
    appendPropertieBtn.element.onclick = function(item, body){
        return function(){
            // Append a new propertie in the item.
            var lastIndex = 0;
            for(var i=0; i < item.M_properties.length; i++){
                if(item.M_properties[i].M_id > lastIndex){
                    lastIndex = item.M_properties[i].M_id;
                }
            }
            
            // Create the new propertie.
            var propertie = new CatalogSchema.PropertyType()
            propertie.M_kind = new CatalogSchema.PropertyKind()
            propertie.M_kind.M_valueOf = "string"
            propertie.M_kind.UUID = "CatalogSchema.PropertyKind%" + randomUUID()
            propertie.M_stringValue = ""
            propertie.M_dimensionValue = ""
            propertie.M_numericValue = 0
            propertie.M_booleanValue = false
            propertie.M_id = lastIndex + 1
            propertie.UUID = "CatalogSchema.PropertyType%" + randomUUID()
            propertie.M_kind.ParentUuid =  propertie.UUID;
            propertie.M_kind.ParentLnk = "M_kind"
            propertie.ParentUuid = item.UUID
            propertie.ParentLnk = "M_properties"
            entities[propertie.UUID] = propertie
            entities[propertie.M_kind.UUID] = propertie.M_kind
  
            item.M_properties.push(propertie)
            appendProperties(body, propertie, item, propertie.M_id - 1)
            
            setSaveCancelBtn(item)
        }
    }(item, body)
    
    // Display and set the value into the editor.
    function createPropertieValueEditor(valueDiv, propertie, item){

        var div = valueDiv.appendElement({"tag":"div", "class":"row"}).down()
        
        var kindSelect = div.appendElement({"tag":"select", "class":"form-control col-sm-4"}).down()

        kindSelect.appendElement({"tag":"option", "value":"string", "innerHtml":"text"})
            .appendElement({"tag":"option", "value":"boolean", "innerHtml":"bool"})
            .appendElement({"tag":"option", "value":"numeric", "innerHtml":"numeric"})
            .appendElement({"tag":"option", "value":"dimension", "innerHtml":"dimension"})
            
        // set it value.
        kindSelect.element.value = propertie.M_kind.M_valueOf
        
        var editorDiv = div.appendElement({"tag":"div", "class":"col-sm-8"}).down()
        
        // So here I will set the different property editor.
        editors = initMultilanguageInput(editorDiv, "textarea", propertie, "M_stringValue")
        for(var id in editors){
            editors[id].element.onchange = function(item){
                return function(){
                    setSaveCancelBtn(item)
                }
            }(item)
        }
        
        var textEditors = editorDiv.getChildsByName(propertie.getEscapedUuid() + "_M_stringValue")

        // The boolean editor.
        var booleanEditor = editorDiv.appendElement({"tag":"label", "class":"form-control switch"}).down()
        booleanEditor.appendElement({"tag":"input", "type":"checkbox", "id" : "checkbox"})
            .appendElement({"tag":"span", "class":"slider round"})
            
        booleanEditor.getChildById("checkbox").element.onchange = function(propertie, item){
            return function(){
                // Set the propertie value.
                propertie.M_booleanValue = this.checked;
                setSaveCancelBtn(item)
            }
        }(propertie, item)
        
        // The numeric value editor.
        var numericEditor = editorDiv.appendElement({"tag":"input", "type":"number", "class":"form-control"}).down()
        numericEditor.element.onchange = function(propertie, item){
            return function(){
                propertie.M_numericValue = this.value;
                setSaveCancelBtn(item)
            }
        }(propertie, item)
        
        // Now the property editor.
        var unitOfMesure = getEntityPrototype("CatalogSchema.UnitOfMeasureEnum")
        var dimensionEditor = editorDiv.appendElement({"tag":"div", "class":"list-group"}).down()
        
        var numberValueInput = dimensionEditor.appendElement({"tag":"div", "class":"row"}).down()
            .appendElement({"tag":"input", "type":"number", "class":"form-control"}).down()
        
        var unitSelector = dimensionEditor.appendElement({"tag":"div", "class":"row"}).down()
            .appendElement({"tag":"select", "class":"form-control"}).down()
    
        for(var i=0; i < unitOfMesure.Restrictions.length; i++){
            unitSelector.appendElement({"tag":"option", "value":unitOfMesure.Restrictions[i].Value, "innerHtml":unitOfMesure.Restrictions[i].Value})
        }
        
        unitSelector.element.onchange = function(propertie, item){
            return function(){
                if(propertie.M_dimensionValue == ""){
                    propertie.M_dimensionValue = new CatalogSchema.DimensionType()
                    propertie.M_dimensionValue.M_unitOfMeasure = new CatalogSchema.UnitOfMeasureEnum()
                }
                propertie.M_dimensionValue.M_unitOfMeasure.M_valueOf = this.value
                setSaveCancelBtn(item)
            }
        }(propertie, item)
        
        // Set the propertie value.
        numberValueInput.element.onchange = function(propertie, item){
            return function(){
                propertie.M_dimensionValue.M_valueOf = this.value;
                setSaveCancelBtn(item)
            }
        }(propertie, item)
        
        // Set the selector value.
        function setEditorValue(propertie, booleanEditor, numericEditor, numberValueInput, unitSelector){
            
            booleanEditor.getChildById("checkbox").element.checked = propertie.M_booleanValue
            
            numericEditor.element.value =  propertie.M_numericValue;
            if(propertie.M_dimensionValue != ""){
                numberValueInput.element.value = propertie.M_dimensionValue.M_valueOf;
                unitSelector.element.value = propertie.M_dimensionValue.M_unitOfMeasure.M_valueOf;
            }
        }
        
        function displayEditor(kind, textEditors, booleanEditor, numericEditor, dimensionEditor, unitSelector){
            // so here I will set the property type selector.
            propertie.M_kind.M_valueOf = kind // Set the type.
            
            if(kind == "string"){
                for(var i=0; i < textEditors.length; i++){
                    textEditors[i].element.style.display = "";
                    booleanEditor.element.style.display = "none";
                    numericEditor.element.style.display = "none";
                    dimensionEditor.element.style.display = "none";
                }
                
            }else if(kind == "boolean"){
                booleanEditor.element.style.display = "";
                for(var i=0; i < textEditors.length; i++){
                    textEditors[i].element.style.display = "none";
                }   
                numericEditor.element.style.display = "none";
                dimensionEditor.element.style.display = "none";
                booleanEditor.element.checked = propertie.M_booleanValue
            }else if(kind == "numeric"){
                numericEditor.element.style.display = "";
                for(var i=0; i < textEditors.length; i++){
                    textEditors[i].element.style.display = "none";
                }  
                booleanEditor.element.style.display = "none";
                dimensionEditor.element.style.display = "none";
                numericEditor.element.value = propertie.M_numericValue
            }else if(kind == "dimension"){
                dimensionEditor.element.style.display = "";
                for(var i=0; i < textEditors.length; i++){
                    textEditors[i].element.style.display = "none";
                }
                booleanEditor.element.style.display = "none";
                numericEditor.element.style.display = "none";
                if(propertie.M_dimensionValue == ""){
                    propertie.M_dimensionValue = new CatalogSchema.DimensionType()
                    propertie.M_dimensionValue.M_unitOfMeasure = new CatalogSchema.UnitOfMeasureEnum()
                    propertie.M_dimensionValue.M_unitOfMeasure.M_valueOf = unitSelector.element.value
                }
                unitSelector.element.value =  propertie.M_dimensionValue.M_unitOfMeasure.M_valueOf 
                dimensionEditor.element.value = propertie.M_dimensionValue.M_valueOf
            }
        }
        
        // display the current propertie editor.
        displayEditor(propertie.M_kind.M_valueOf, textEditors, booleanEditor, numericEditor, dimensionEditor, unitSelector)
        setEditorValue(propertie, booleanEditor, numericEditor, numberValueInput, unitSelector)
        
        // Now the control actions...
        kindSelect.element.onchange = function(textEditors, booleanEditor, numericEditor, dimensionEditor, unitSelector, item){
            return function(){
                setSaveCancelBtn(item)
                displayEditor(this.value, textEditors, booleanEditor, numericEditor, dimensionEditor, unitSelector)
            }
        }(textEditors, booleanEditor, numericEditor, dimensionEditor, unitSelector, item)
        
    }
    
    
    // Now the append properties
    function appendProperties(body, propertie, item, index){
        var row = body.appendElement({"tag":"tr", "class":"row"}).down()
        
        // The propertie id, immutable.
        var idSpan = row.appendElement({"tag":"td", "class":"col-sm-1"}).down()
            .appendElement({"tag":"span", "innerHtml":propertie.M_id.toString()}).down()
        
        // The propertie name.
        var nameInput = row.appendElement({"tag":"td", "class":"col-sm-2"}).down()
        var editors = initMultilanguageInput(nameInput, "input", propertie, "M_name")  
        for(var id in editors){
            editors[id].element.onchange = function(item){
                return function(){
                    setSaveCancelBtn(item)
                }
            }(item)
        }
        
        var descriptionTextarea = row.appendElement({"tag":"td", "class":"col-sm-3"}).down()
        var editors_ = initMultilanguageInput(descriptionTextarea, "textarea", propertie, "M_description")  
        for(var id in editors_){
            editors_[id].element.onchange = function(item){
                return function(){
                    setSaveCancelBtn(item)
                }
            }(item)
        }
        var valuediv = row.appendElement({"tag":"td", "class":"col-sm-5"}).down()
        createPropertieValueEditor(valuediv, propertie, item)

        var deleteBtn = row.appendElement({"tag":"td", "class":"col-sm-1"}).down()
        .appendElement({"tag":"button", "class":"btn btn-danger btn-sm", "style":"display: inline-flex; height: 29px;"}).down()
        
        // append the icon.
        deleteBtn.appendElement({"tag":"i", "class":"fa fa-trash"})
        
        // So here I will delete the propertie from the local object.
        deleteBtn.element.onclick = function(row, item, index){
            return function(){
                // remove the items at a given index.
                item.M_properties.splice(index, 1) 
                row.element.parentNode.removeChild(row.element)
                setSaveCancelBtn(item)
            }
        }(row, item, index)
        
    }
    
    // I will append the existing properties to the editor.
    for(var i=0; i < item.M_properties.length; i++){
        appendProperties(body, item.M_properties[i], item, i)
    }

    // Now the save and cancel button.
    var btnRow = itemPanel.appendElement({"tag":"div", "class":"d-flex justify-content-center"}).down()
      

    var saveBtn = btnRow.appendElement({"tag":"button", "class":"btn btn-primary disabled mr-3", "innerHtml":"Enregistrer", "id":item.getEscapedUuid() + "_save_btn"}).down()
    var saveAsBtn = btnRow.appendElement({"tag":"button", "class":"btn btn-primary mr-3", "innerHtml":"Enregistrer Sous", "id":item.getEscapedUuid() + "_save_btn"}).down()
    var cancelBtn = btnRow.appendElement({"tag":"button", "class":"btn disabled", "innerHtml":"Annuler les modifications", "id":item.getEscapedUuid() + "_cancel_btn"}).down()
    var deleteBtn = btnRow.appendElement({"tag":"button", "class":"btn btn-danger ml-3", "innerHtml":"Supprimer", "id":item.getEscapedUuid() + "_delete_btn"}).down()
    
    // Save as is use to create a copy of an existing item so it can be a starting point for similar items...
    saveAsBtn.element.onclick = function(item){
        return function(){
            var copy = jQuery.extend(true, {},item)
            copy.UUID = ""
            copy.M_properties = []
            for(var i=0; i < item.M_properties.length; i++){
                var property = new CatalogSchema.PropertyType()
                property.M_name = item.M_properties[i].M_name
                property.M_id = item.M_properties[i].M_id
                property.M_description = item.M_properties[i].M_description
                property.M_stringValue = item.M_properties[i].M_stringValue
                property.M_numericValue = item.M_properties[i].M_numericValue
                property.M_booleanValue = item.M_properties[i].M_booleanValue
                if(item.M_properties[i].M_dimensionValue == ""){
                     property.M_dimensionValue = ""
                }else{
                    property.M_dimensionValue = new CatalogSchema.DimensionType()
                    property.M_dimensionValue.M_valueOf = item.M_properties[i].M_dimensionValue.M_valueOf
                    property.M_dimensionValue.M_unitOfMeasure = new CatalogSchema.UnitOfMeasureEnum()
                    property.M_dimensionValue.M_unitOfMeasure.M_valueOf = item.M_properties[i].M_dimensionValue.M_unitOfMeasure.M_valueOf
                }
                property.M_kind = new CatalogSchema.PropertyKind()
                property.M_kind.M_valueOf = item.M_properties[i].M_kind.M_valueOf
                copy.M_properties.push(property)
            }
            
            // Now I will display the save as dialog.
            var confirmDialog = new Dialog(randomUUID(), undefined, true)
            confirmDialog.title.element.innerHTML = "Enregister sous"
            var idInput = confirmDialog.content.appendElement({ "tag": "span", "innerHtml": "Veuillez entrer le ID pour le nouvel item à crée "})
                .appendElement({"tag":"input", "type":"text", "class":"form-control"}).down()
                
            idInput.element.focus();
            
            confirmDialog.div.element.style.maxWidth = "650px"
            confirmDialog.setCentered()
            confirmDialog.header.element.style.padding = "5px";
            confirmDialog.header.element.style.color = "white";
            
            confirmDialog.header.element.classList.add("bg-dark")
            confirmDialog.ok.element.classList.remove("btn-default")
            confirmDialog.ok.element.classList.add("btn-primary")
            confirmDialog.ok.element.classList.add("mr-3")
            confirmDialog.deleteBtn.element.style.color = "white"
            confirmDialog.deleteBtn.element.style.fontSize="16px"
            
            confirmDialog.ok.element.onclick = function (dialog, item, idInput) {
                return function () {
                    // I will call delete file
                    spinner.panel.element.style.display = "";
                    // So here I will create the new item...
                    copy.M_id = idInput.element.value;
                    server.entityManager.createEntity(copy.ParentUuid, copy.ParentLnk, copy, 
                         // The success callback.
                        function(item, caller){
                            // So now I will set the new Item in the display
                            mainPage.adminPage.adminItemPage.panel.getChildById("admin_item_display").removeAllChilds()
                            mainPage.adminPage.adminItemPage.panel.getChildById("items_filter_lst").removeAllChilds()
                            
                            // Set the active item.
                            var itemPanel = mainPage.adminPage.adminItemPage.setItem(item)
                                                                    
                            // Reset the display attribute.
                            var itemsPanel = document.getElementsByClassName("item_panel");
                            for(var i=0; i < itemsPanel.length; i++){
                                itemsPanel[i].style.display = ""
                            }
                            
                            itemPanel.element.style.display = "block"
                            
                            // Set the focus to the name input.
                            var nameInputs = document.getElementsByName(item.getEscapedUuid() + "_M_name")
                            for(var i=0; i < nameInputs.length; i++){
                                if(server.languageManager.language == nameInputs[i].lang){
                                    nameInputs[i].focus()
                                    break;
                                }
                            }
                            spinner.panel.element.style.display = "none";
                            
                        },
                        // The error callback.
                        function(errObj, caller){
                            spinner.panel.element.style.display = "none";
                        },{})
                        
                    
                    dialog.close()
                }
            } (confirmDialog, item, idInput)
        }
    }(item)
    
    // Now I will save the entity...
    saveBtn.element.onclick = function(item, saveBtn, cancelBtn){
        return function(){
            if(saveBtn.element.classList.contains("disabled")){
                return
            }
            saveBtn.element.classList.add("disabled")
            cancelBtn.element.classList.add("disabled")
            spinner.panel.element.style.display = "";
            server.entityManager.saveEntity(item,
                // success callback
                function(item, caller){
                    /** Nothing here **/
                    spinner.panel.element.style.display = "none";
                    var itemAdminTabTitle = document.getElementById("item-adminTab").firstChild
                    if(itemAdminTabTitle.innerHTML.endsWith("*")){
                        itemAdminTabTitle.innerHTML = itemAdminTabTitle.innerHTML.replace("*", "")
                    }
                }, 
                // error callback
                function(){
                    spinner.panel.element.style.display = "none";
                }, 
                {})
        }
    }(item, saveBtn, cancelBtn)
    
    // Now the cancel button action.
    // create a backup object.
    var backup = jQuery.extend(true, {},item)
    for(var i=0; i < backup.M_properties.length; i++){
        backup.M_properties[i] =  jQuery.extend(true, {},backup.M_properties[i])
        backup.M_properties[i].M_kind = jQuery.extend(true, {},backup.M_properties[i].M_kind)
        backup.M_properties[i].M_dimensionValue = jQuery.extend(true, {},backup.M_properties[i].M_dimensionValue)
        backup.M_properties[i].M_dimensionValue.M_unitOfMeasure = jQuery.extend(true, {},backup.M_properties[i].M_dimensionValue.M_unitOfMeasure )
    }
    
    cancelBtn.element.onclick = function(item, saveBtn, cancelBtn){
        return function(){
            if(saveBtn.element.classList.contains("disabled")){
                return
            }
            
            saveBtn.element.classList.add("disabled")
            cancelBtn.element.classList.add("disabled")
            
            entities[item.UUID] = item

            // Set the item.
            var itemPanel = mainPage.adminPage.adminItemPage.setItem(item)
            itemPanel.element.style.display = "block";
            
            var itemAdminTabTitle = document.getElementById("item-adminTab").firstChild
            if(itemAdminTabTitle.innerHTML.endsWith("*")){
                itemAdminTabTitle.innerHTML = itemAdminTabTitle.innerHTML.replace("*", "")
            }
        }
    }(backup, saveBtn, cancelBtn)
    
    deleteBtn.element.onclick = function(item){
        return function(){
            // Suppression d'un item de la base de donnée.
            
            var confirmDialog = new Dialog(randomUUID(), undefined, true)
            confirmDialog.title.element.innerHTML = "Supprimer l'item"
            confirmDialog.content.appendElement({ "tag": "span", "innerHtml": "Voulez-vous vraiment supprimer l'item " + item.M_id + "?" })
            confirmDialog.div.element.style.maxWidth = "650px"
            confirmDialog.setCentered()
            confirmDialog.header.element.style.padding = "5px";
            confirmDialog.header.element.style.color = "white";
            
            confirmDialog.header.element.classList.add("bg-dark")
            confirmDialog.ok.element.classList.remove("btn-default")
            confirmDialog.ok.element.classList.add("btn-primary")
            confirmDialog.ok.element.classList.add("mr-3")
            confirmDialog.deleteBtn.element.style.color = "white"
            confirmDialog.deleteBtn.element.style.fontSize="16px"
            
            confirmDialog.ok.element.onclick = function (dialog, item) {
                return function () {
                    // I will call delete file
                    spinner.panel.element.style.display = "";
                    // remove the item form the editor...
                    server.entityManager.removeEntity(item.UUID, 
                    // Success callback
                    function(result, caller){
                        // remove the div.
                        spinner.panel.element.style.display = "none";
                        var selector = document.getElementById(caller + "-selector")
                        selector.parentNode.removeChild(selector)
                        var control = document.getElementById(caller + "-control")
                        control.parentNode.removeChild(control)
                    },
                    // Error callback
                    function(errObj, caller){
                        spinner.panel.element.style.display = "none";
                    },item.getEscapedUuid())
                    dialog.close()
                }
            } (confirmDialog, item)
        }
    }(item)
    
    server.languageManager.refresh()
    
    return itemPanel;
}
