/**
 * Add autocomplete functionality to a type (the type must have a propertie id)
 * The input must have a function name 'oncomplete(uuid)'.
 */
function autocomplete(typeName, input, ids){
 	var q = {};
	q.TypeName = typeName;
	
	if(ids === undefined){
	    q.Fields = ["UUID", "M_id"];
	}else{
	    if(ids.indexOf("UUID") == -1){
	        ids.unshift("UUID");
	    }
	    q.Fields = ids;
	}
	
	var fieldsType = ["xs.string", "xs.string"];
	
	server.dataManager.read("CatalogSchema", JSON.stringify(q), fieldsType, [],
		// success callback
		function (results, caller) {
			var results = results[0];
			if (results == null) {
				return
			}
			var elementLst = [];
			var idUuid = {};
			for (var i = 0; i < results.length; i++) {
				elementLst.push(results[i][1])
				idUuid[results[i][1]] = results[i][0]
				if (caller.value == results[i][0]) {
					caller.input.element.value = results[i][1];
				}
			}
			// I will attach the autocomplete box.
			attachAutoComplete(caller.input, elementLst, true,
				function (caller, idUuid) {
					return function (id) {
						caller.input.element.value = id;
						caller.input.element.oncomplete(idUuid[id])
					}
				}(caller, idUuid));
		},
		// progress callback
		function (index, total, caller) {

		},
		// error callback
		function (errObj, caller) {

		}, {"input":input})
}

/**
 * Set multilanguage input and textarea, (work with the language manager.)
 */
function initMultilanguageInput(parent, tag, entity, propertie){
    var id = entity.getEscapedUuid() + "_" + propertie
    var value = entity[propertie]
    
    // The text area must display the content of the actual language...
    // The field M_description contain innerHtml values.
    parser=new DOMParser();
    var divs =parser.parseFromString(value, "text/html").getElementsByTagName("div");
    var languages = {}
    var editors = {}
    for(var i=0; i < divs.length; i++){
        if(tag == "textarea"){
            var editor = parent.appendElement({"tag":tag, "class":"form-control", "innerHtml":divs[i].innerText, "lang":divs[i].lang, "name":id}).down()
            editors[divs[i].lang] = editor
        }else if(tag == "input"){
            var editor = parent.appendElement({"tag":tag, "class":"form-control", "type":"text", "value":divs[i].innerText, "lang":divs[i].lang, "name":id}).down()
            editors[divs[i].lang] = editor
        }
        languages[divs[i].lang] = ""
    }
    
    // Now if some language are missing I will create the text area...
    for(var languageInfo in server.languageManager.languageInfo){
        if(languages[languageInfo] == undefined){
            var editor = parent.appendElement({"tag":tag, "class":"form-control", "lang":languageInfo, "name":id}).down() 
            if(divs.length == 0){
                // In that case the propertie is a single string...
                if(tag == "textarea"){
                    editor.element.value = value
                }else if(tag == "input"){
                    editor.element.value = value
                }
                editors[languageInfo] = editor
            }
        }
    }
    
    // The onchange action now.

    // Set the listener.
    for(var lang in editors){
        editors[lang].element.onkeyup = function(editors, entity, propertie){
            return function(){
                // So here the value of the entity propertie will be one divs for each 
                // editors.
                var str = ""
                for(var lang in editors){
                    str += '<div lang="' + lang + '">'
                    if(editors[lang].element.tagName == "TEXTAREA"){
                        str += editors[lang].element.value
                    }else if(editors[lang].element.tagName == "INPUT"){
                        str += editors[lang].element.value
                    }
                    str += "</div>"
                }
                entity[propertie] = str;
            }
        }(editors, entity, propertie)   
    }
    
    return editors;
}



var AdminPage = function (parent){
    this.panel = new Element(parent, { "tag": "div", "id": "admin_page_panel", "class": "item_display container-fluid" , "style" :"margin-top : 15px; display: none;"})
    this.panel.appendElement({"tag" : "ul", "class" : "nav nav-tabs querynav"}).down()
    .appendElement({"tag": "li","class" : "nav-item"}).down()
    .appendElement({"tag" : "a", "class" : "nav-link active show", "href" : "#supplier-adminContent", "role" :"tab", "data-toggle" : "tab", "aria-controls" : "supplier-adminContent", "id":"supplier-adminTab"}).down()
    .appendElement({"tag" : "span", "innerHtml" : "Fournisseur"})
    .appendElement({"tag" : "span", "style" : "margin-left:5px;", "id" : "adminSupplierSaveState"}).up().up()
    .appendElement({"tag": "li","class" : "nav-item"}).down()
    .appendElement({"tag" : "a", "class" : "nav-link", "href" : "#localisation-adminContent", "role" :"tab", "data-toggle" : "tab", "aria-controls" : "localisation-adminContent", "id":"localisation-adminTab"}).down()
    .appendElement({"tag" : "span", "innerHtml" : "Localisation"})
    .appendElement({"tag" : "span", "style" : "margin-left:5px;", "id" : "adminLocalisationSaveState"}).up().up()
    .appendElement({"tag": "li","class" : "nav-item"}).down()
    .appendElement({"tag" : "a", "class" : "nav-link ", "href" : "#manufacturer-adminContent", "role" :"tab", "data-toggle" : "tab", "aria-controls" : "manufacturer-adminContent", "id":"manufacturer-adminTab"}).down()
    .appendElement({"tag" : "span", "innerHtml" : "Fabricant"})
    .appendElement({"tag" : "span", "style" : "margin-left:5px;", "id" : "adminManufacturerSaveState"}).up().up()
    .appendElement({"tag": "li","class" : "nav-item"}).down()
    .appendElement({"tag" : "a", "class" : "nav-link ", "href" : "#category-adminContent", "role" :"tab", "data-toggle" : "tab", "aria-controls" : "category-adminContent", "id":"category-adminTab"}).down()
    .appendElement({"tag" : "span", "innerHtml" : "Catégorie"})
    .appendElement({"tag" : "span", "style" : "margin-left:5px;", "id" : "adminCategorySaveState"}).up().up()
    .appendElement({"tag": "li","class" : "nav-item"}).down()
    .appendElement({"tag" : "a", "class" : "nav-link ", "href" : "#item-adminContent", "role" :"tab", "data-toggle" : "tab", "aria-controls" : "item-adminContent", "id":"item-adminTab"}).down()
    .appendElement({"tag" : "span", "innerHtml" : "Item"})
    .appendElement({"tag" : "span", "style" : "margin-left:5px;", "id" : "adminItemSaveState"}).up().up()
    .appendElement({"tag": "li","class" : "nav-item"}).down()
    .appendElement({"tag" : "a", "class" : "nav-link ", "href" : "#package-adminContent", "role" :"tab", "data-toggle" : "tab", "aria-controls" : "package-adminContent", "id":"package-adminTab"}).down()
    .appendElement({"tag" : "span", "innerHtml" : "Paquet"})
    .appendElement({"tag" : "span", "style" : "margin-left:5px;", "id" : "adminPackageSaveState"}).up().up()
    .appendElement({"tag" : "li", "style" : "padding:10px;display:flex;"}).down()
    .appendElement({"tag" : "button", "class" : "btn btn-primary disabled ml-3", "innerHtml" : "Enregistrer tout", "style" : "padding:1px 6px;"}).up().up()
    .appendElement({"tag" : "div", "class" : "tab-content"}).down()
    .appendElement({"tag" : "div", "class" : "tab-pane active show result-tab-content","id" : "supplier-adminContent", "role" : "tabpanel", "aria-labelledby" : "supplier-adminTab" })
    .appendElement({"tag" : "div", "class" : "tab-pane result-tab-content","id" : "localisation-adminContent", "role" : "tabpanel", "aria-labelledby" : "localisation-adminTab" })
    .appendElement({"tag" : "div", "class" : "tab-pane show result-tab-content","id" : "manufacturer-adminContent", "role" : "tabpanel", "aria-labelledby" : "manufacturer-adminTab" })
    .appendElement({"tag" : "div", "class" : "tab-pane show result-tab-content","id" : "category-adminContent", "role" : "tabpanel", "aria-labelledby" : "category-adminTab" })
    .appendElement({"tag" : "div", "class" : "tab-pane show result-tab-content","id" : "item-adminContent", "role" : "tabpanel", "aria-labelledby" : "item-adminTab" })
    .appendElement({"tag" : "div", "class" : "tab-pane show result-tab-content","id" : "package-adminContent", "role" : "tabpanel", "aria-labelledby" : "package-adminTab" })
  
    this.adminItemPage = new AdminItemPage(this.panel.getChildById("item-adminContent"));
    
    this.adminSupplierPage = new AdminSupplierPage(this.panel.getChildById("supplier-adminContent"))
    
    this.adminLocalisationPage = new AdminLocalisationPage(this.panel.getChildById("localisation-adminContent"))
  
    this.adminManufacturerPage = new AdminManufacturerPage(this.panel.getChildById("manufacturer-adminContent"))
    
    this.adminCategoryPage = new AdminCategoryPage(this.panel.getChildById("category-adminContent"))
    
    this.adminPackagePage = new AdminPackagePage(this.panel.getChildById("package-adminContent"))
  
    return this
}

AdminPage.prototype.displayAdminPage = function(){
    document.getElementById("main-container").style.display = "none";
    document.getElementById("admin_page_panel").style.display = "";
    if(document.getElementById("item_search_result_page") !== null){
        document.getElementById("item_search_result_page").style.display = "none";
    }
    if(document.getElementById("item_display_page_panel") !== null){
        document.getElementById("item_display_page_panel").style.display = "none";
    }
    document.getElementById("order_page_panel").style.display = "none";
    
    if(document.getElementById("order_history_page_panel") !== null){
        document.getElementById("order_history_page_panel").style.display = "none";
    }
    document.getElementById("package_display_page_panel").style.display = "none";
}


