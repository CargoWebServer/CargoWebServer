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
        .appendElement({"tag":"div", "class":"list-group", "role":"tablist", "id":"items_filter_lst", "style":"overflow-y: auto;"}).up()
        .appendElement({"tag":"div", "class":"col-md-9 bg-light", "style":"overflow-y: auto;"}).down()
        .appendElement({"tag":"div", "class":"tab-content", "id":"admin_item_display"}).down()

    this.searchBar = this.panel.getChildById("items_filter_input")
    this.searchBtn =  this.panel.getChildById("item_filter_btn")
    
    var itemFilterLst = this.panel.getChildById("items_filter_lst")
    window.addEventListener('resize', function(itemFilterLst){
        return function(evt){
            var h = itemFilterLst.element.parentNode.offsetHeight - itemFilterLst.element.parentNode.firstChild.offsetHeight - 20;
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
                    mainPage.adminPage.adminItemPage.panel.getChildById("items_filter_lst").element.innerHTML = ""

                    var items = {}

                    if (results.estimate > 0) {
                        for (var i = 0; i < results.results.length; i++) {
                            var result = results.results[i];
                            if (result.data.TYPENAME == "CatalogSchema.ItemType") {
                                var uuid = results.results[i].data.UUID;
                                if (items[uuid] == null) {
                                    if(entities[uuid] != null){
                                        items[uuid] = entities[uuid];
                                    }else{
                                        items[uuid] = eval("new " + results.results[i].data.TYPENAME + "()")
                                        items[uuid].init(results.results[i].data, false, function(item){
                                            console.log(item)
                                            mainPage.adminPage.adminItemPage.setItem(item)
                                        })
                                        entities[uuid] = items[uuid];
                                    }
                                }
                            }
                        }
                    }else{
                        caller.searchBar.element.placeholder = "No results found for '" + caller.searchBar.element.value +"'"
                        caller.searchBar.element.value = ""
                        caller.searchBar.setTimeout(function(searchBar) {
                            return function(){
                                caller.searchBar.element.placeholder = "search"
                            }
                        }(caller.searchBar), 10);
                    }
                    
                    // Clear previous search...
                    mainPage.adminPage.adminItemPage.panel.getChildById("items_filter_lst").removeAllChilds()
                    mainPage.adminPage.adminItemPage.panel.getChildById("admin_item_display").removeAllChilds()
                    
                    for(var id in items){
                        var txt = items[id].M_id
                        for(var i=0; i < items[id].M_alias.length; i++){
                            if(i==0){
                                txt += " (";
                            }
                            txt += " " +  items[id].M_alias[i]
                            if(i < items[id].M_alias.length - 2){
                                txt += ","
                            }
                            if(i==items[id].M_alias.length-1){
                                txt += " )";
                            }
                        }
                        
                        mainPage.adminPage.adminItemPage.panel.getChildById("items_filter_lst").appendElement({"tag":"a", "class" : "list-group-item list-group-item-action","id" : items[id].M_id + "-selector","data-toggle" : "tab","href" : "#"+ items[id].M_id + "-control","role":"tab", "innerHtml": txt,"aria-controls":  items[id].M_id + "-control"})
                        mainPage.adminPage.adminItemPage.loadAdminControl(items[id])
                        if(items[id].M_properties.length > 0){
                            mainPage.adminPage.adminItemPage.setItem(items[id])
                        }
                    }
                    
                    fireResize()
                },
                // error callback
                function () {
        
                }, {"keyword":keyword, "searchBar":searchBar})
        }
    }(this.searchBar)
    
    this.searchBtn.element.onclick = searchFct
    
    this.searchBar.element.onkeyup = function(searchFct){
        return function(evt){
            if(evt.keyCode == 13){
                searchFct()
            }
        }
    }(searchFct)
    return this
}

AdminItemPage.prototype.loadAdminControl = function(item){
    var itemPanel = this.itemsPanel.getChildById(item.M_id + "-control");
    if( itemPanel == null){
        itemPanel = this.itemsPanel.appendElement({"tag":"div", "class":"tab-pane", "id" : item.M_id + "-control", "aria-labelledby":item.M_id +"-selector", "style":"padding:15px;"}).down()
    }
    return itemPanel;
}

AdminItemPage.prototype.setItem = function(item){

    var itemPanel = this.loadAdminControl(item)

    // Now I will display the properties of the items.
    // The id, unmutable.
    itemPanel.appendElement({"tag":"div", "class":"input-group mb-3"}).down()
    .appendElement({"tag":"div","class":"input-group-prepend"}).down()
    .appendElement({"tag":"span", "class":"input-group-text", "innerHtml":"id"}).up()
    .appendElement({"tag":"span", "class":"form-control", "innerHtml":item.M_id})
    
    // The Type name.
    itemPanel.appendElement({"tag":"div", "class":"input-group mb-3"}).down()
    .appendElement({"tag":"div","class":"input-group-prepend"}).down()
    .appendElement({"tag":"span", "class":"input-group-text", "innerHtml":"Nom"}).up()
    .appendElement({"tag":"input", "class":"form-control","type":"text", "value":item.M_name})
    
    // The description
    var desciptions = itemPanel.appendElement({"tag":"div", "class":"input-group mb-3"}).down()
    .appendElement({"tag":"div","class":"input-group-prepend"}).down()
    .appendElement({"tag":"span", "class":"input-group-text", "innerHtml":"Description"}).up()

    // The text area must display the content of the actual language...
    // The field M_description contain innerHtml values.
    parser=new DOMParser();
    var divs =parser.parseFromString(item.M_description, "text/html").getElementsByTagName("div");
    var languages = {}
    
    for(var i=0; i < divs.length; i++){
        desciptions.appendElement({"tag":"textarea", "class":"form-control", "innerHtml":divs[i].innerText, "lang":divs[i].lang})
        languages[divs[i].lang] = ""
    }
    
    // Now if some language are missing I will create the text area...
    for(var languageInfo in server.languageManager.languageInfo){
        if(languages[languageInfo] == undefined){
           desciptions.appendElement({"tag":"textarea", "class":"form-control", "lang":languageInfo}) 
        }
    }
    
    // Now the alias...
    itemPanel.appendElement({"tag":"div", "class":"row mb-3"}).down()
    .appendElement({"tag":"div", "class":"col-sm-3", "style":"padding-right: 0px; padding-left: 0px;"}).down()
    .appendElement({"tag":"ul", "class":"list-group borderless", "id":"alias"}).up()
    .appendElement({"tag":"div", "class":"col-sm-3", "style":"padding-right: 0px; padding-left: 0px;"}).down()
    .appendElement({"tag":"ul", "class":"list-group borderless", "id":"keywords"}).up()
    .appendElement({"tag":"div", "class":"col-sm-6", "style":"padding-right: 0px; padding-left: 0px;"}).down()
    .appendElement({"tag":"ul", "class":"list-group borderless", "id":"comments"})
    
    var alias = itemPanel.getChildById("alias")
    var comments = itemPanel.getChildById("comments")
    var keywords = itemPanel.getChildById("keywords")
    
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
        deletBtn.element.onclick = function(row, index, alias){
            return function(){
                row.element.parentNode.removeChild(row.element)
                alias.splice(index, 1);
            }
        }(row, i, item.M_alias)
        
        // Edit the content.
        row.getChildById("row_content").element.onchange = function(index, alias){
            return function(){
                // Set the new value
                alias[index] = this.value;
            }
        }(i, item.M_alias)
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
        deletBtn.element.onclick = function(row, index, keywords){
            return function(){
                row.element.parentNode.removeChild(row.element)
                keywords.splice(index, 1);
            }
        }(row, i, item.M_keywords)
        
        row.getChildById("row_content").element.onchange = function(index, keywords){
            return function(){
                // Set the new value
                keywords[index] = this.value;
            }
        }(i, item.M_keywords)
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
        deletBtn.element.onclick = function(row, index, comments){
            return function(){
                row.element.parentNode.removeChild(row.element)
                comments.splice(index, 1);
            }
        }(row, i, item.M_comments)
        
        // Edit the content.
        row.getChildById("row_content").element.onchange = function(index, comments){
            return function(){
                // Set the new value
                comments[index] = this.innerHTML;
            }
        }(i, item.M_comments)
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
            
            var i = item.M_keywords.length
            item.M_keywords.push("")
            
            // Remove the row.
            deletBtn.element.onclick = function(row, index, keywords){
                return function(){
                    row.element.parentNode.removeChild(row.element)
                    keywords.splice(index, 1);
                }
            }(row, i, item.M_keywords)
            
            row.getChildById("row_content").element.onchange = function(index, keywords){
                return function(){
                    // Set the new value
                    keywords[index] = this.value;
                }
            }(i, item.M_keywords)
        }
    }(keywords, item)
    
    appendAliasBtn.element.onclick = function(alias, item){
        return function(){
            var row = alias.appendElement({"tag":"li", "class":"list-group-item"}).down()
            var deletBtn = row.appendElement({"tag":"div", "class":"input-group"}).down()
                .appendElement({"tag":"input", "class":"form-control","type":"text", "id":"row_content"})
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
            deletBtn.element.onclick = function(row, index, alias){
                return function(){
                    row.element.parentNode.removeChild(row.element)
                    alias.splice(index, 1);
                }
            }(row, i, item.M_alias)
            
            // Edit the content
            row.getChildById("row_content").element.onchange = function(index, alias){
                return function(){
                    // Set the new value
                    alias[index] = this.value;
                }
            }(i, item.M_alias)
        }
    }(alias, item)
    
    appendCommentBtn.element.onclick = function(comments, item){
        return function(){
            var row = comments.appendElement({"tag":"li", "class":"list-group-item"}).down()
            var deletBtn = row.appendElement({"tag":"div", "class":"input-group"}).down()
                .appendElement({"tag":"textarea", "class":"form-control", "id":"row_content"})
                .appendElement({"tag":"div","class":"input-group-append"}).down()
                .appendElement({"tag":"span","class":"input-group-text"}).down()
            
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
            deletBtn.element.onclick = function(row, index, comments){
                return function(){
                    row.element.parentNode.removeChild(row.element)
                    comments.splice(index, 1);
                }
            }(row, i, item.M_comments)
            
            // Edit the content.
            row.getChildById("row_content").element.onchange = function(index, comments){
                return function(){
                    // Set the new value
                    comments[index] = this.innerHTML;
                }
            }(i, item.M_comments)
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
        .appendElement({"tag":"i", "class":"fa fa-plus"}).down()
        
    var body = properties.appendElement({"tag":"tbody"}).down()
    
    // Display and set the value into the editor.
    function createPropertieValueEditor(valueDiv, propertie){
        
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
        var textEditor = editorDiv.appendElement({"tag":"textarea", "class":"form-control"}).down()
        textEditor.element.onchange = function(propertie){
            return function(){
                // Set the propertie value.
                propertie.M_stringValue = this.innerHTML;
            }
        }(propertie)
        
        // The boolean editor.
        var booleanEditor = editorDiv.appendElement({"tag":"label", "class":"form-control switch"}).down()
        booleanEditor.appendElement({"tag":"input", "type":"checkbox", "id" : "checkbox"})
            .appendElement({"tag":"span", "class":"slider round"})
            
        booleanEditor.getChildById("checkbox").element.onchange = function(propertie){
            return function(){
                // Set the propertie value.
                propertie.M_booleanValue = this.cheched;
            }
        }(propertie)
        
        // The numeric value editor.
        var numericEditor = editorDiv.appendElement({"tag":"input", "type":"number", "class":"form-control"}).down()
        numericEditor.element.onchange = function(propertie){
            return function(){
                propertie.M_numericValue = this.value;
            }
        }(propertie)
        
        // Create a new dimension type.
        if(propertie.M_dimensionValue == null){
            propertie.M_dimensionValue = new CatalogSchema.DimensionType()
        }
        
        // Now the property editor.
        var unitOfMesure = getEntityPrototype("CatalogSchema.UnitOfMeasureEnum")
        var dimensionEditor = editorDiv.appendElement({"tag":"div", "class":"list-group"}).down()
        
        var numberValueInput = dimensionEditor.appendElement({"tag":"div", "class":"row"}).down()
            .appendElement({"tag":"input", "type":"number", "class":"form-control", "number_input"}).down()
        
        var unitSelector = dimensionEditor.appendElement({"tag":"div", "class":"row"}).down()
            .appendElement({"tag":"select", "class":"form-control", "id":"unit_selector"}).down()
    
        for(var i=0; i < unitOfMesure.Restrictions.length; i++){
            unitSelector.appendElement({"tag":"option", "value":unitOfMesure.Restrictions[i].Value, "innerHtml":unitOfMesure.Restrictions[i].Value})
        }
        
        // Set the propertie value.
        numberValueInput.element.onchange = function(propertie){
            return function(){
                propertie.M_dimensionValue.M_valueOf = this.value;
            }
        }(propertie)
        
        // Set the selector value.
        
        
        function setEditorValue(propertie, textEditor, booleanEditor, numericEditor, dimensionEditor){
            textEditor.element.innerHTML = propertie.M_stringValue
            booleanEditor.getChildById("checkbox").element.cheched = propertie.M_booleanValue;
            numericEditor.element.value =  propertie.M_numericValue;
            // dimensionEditor.element.style.display = "";
        }
        
        function displayEditor(kind, textEditor, booleanEditor, numericEditor, dimensionEditor){
            // so here I will set the property type selector.
            propertie.M_kind.M_valueOf = kind // Set the type.
            
            if(kind == "string"){
                textEditor.element.style.display = "";
                booleanEditor.element.style.display = "none";
                numericEditor.element.style.display = "none";
                dimensionEditor.element.style.display = "none";
            }else if(kind == "boolean"){
                booleanEditor.element.style.display = "";
                textEditor.element.style.display = "none";
                numericEditor.element.style.display = "none";
                dimensionEditor.element.style.display = "none";
            }else if(kind == "numeric"){
                numericEditor.element.style.display = "";
                textEditor.element.style.display = "none";
                booleanEditor.element.style.display = "none";
                dimensionEditor.element.style.display = "none";
            }else if(kind == "dimension"){
                dimensionEditor.element.style.display = "";
                textEditor.element.style.display = "none";
                booleanEditor.element.style.display = "none";
                numericEditor.element.style.display = "none";
            }
        }
        
        // display the current propertie editor.
        displayEditor(propertie.M_kind.M_valueOf, textEditor, booleanEditor, numericEditor, dimensionEditor)
        setEditorValue(propertie, textEditor, booleanEditor, numericEditor, dimensionEditor)
        
        // Now the control actions...
        kindSelect.element.onchange = function(textEditor, booleanEditor, numericEditor, dimensionEditor){
            return function(){
                displayEditor(this.value, textEditor, booleanEditor, numericEditor, dimensionEditor)
            }
        }(textEditor, booleanEditor, numericEditor, dimensionEditor)
        
        
    }
    
    // Now the append properties
    function appendProperties(body, propertie){
        var row = body.appendElement({"tag":"tr", "class":"row"}).down()
        
        // The propertie id, immutable.
        var idSpan = row.appendElement({"tag":"td", "class":"col-sm-1"}).down()
            .appendElement({"tag":"span", "innerHtml":propertie.M_id.toString()}).down()
        
        // The propertie name.
        var nameInput = row.appendElement({"tag":"td", "class":"col-sm-2"}).down()
            .appendElement({"tag":"input", "class":"form-control", "type":"text", "value":propertie.M_name}).down()
            
        var descriptionTextarea = row.appendElement({"tag":"td", "class":"col-sm-3"}).down()
            .appendElement({"tag":"textarea", "class":"form-control", "innerHtml":propertie.M_description}).down()
        
        var valuediv = row.appendElement({"tag":"td", "class":"col-sm-5"}).down()
        
        createPropertieValueEditor(valuediv, propertie)
        
        var deleteBtn = row.appendElement({"tag":"td", "class":"col-sm-1"}).down()
        .appendElement({"tag":"button", "class":"btn btn-danger btn-sm", "style":"display: inline-flex; height: 29px;"}).down()
        .appendElement({"tag":"i", "class":"fa fa-trash"}).down()
        
    }
    
    // I will append the existing properties to the editor.
    for(var i=0; i < item.M_properties.length; i++){
        appendProperties(body, item.M_properties[i])
    }
    
    server.languageManager.refresh()
}
