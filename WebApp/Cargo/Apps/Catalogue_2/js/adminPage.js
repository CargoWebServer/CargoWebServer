
/**
 * That function use zebra printer to print item label.
 * The Zebra software must be intall before use...
 */
var defaultPrinter = null
var availablePrinters = {}

// Change the value here to fit the model.
var WORKING_PRINTER_DPI = 300;
var printerDPI = 200;
var scale = printerDPI / WORKING_PRINTER_DPI;

function scaleValue(value){
    return value * scale
}

function printerError(text){
    // show error message!
    mainPage.showNotification("danger", text, 3000)
}

function printComplete(){
    // show success message
}

function checkPrinterStatus(finishedFunction)
{
	defaultPrinter.sendThenRead("~HQES", 
				function(text){
						var that = this;
						var statuses = new Array();
						var ok = false;
						var is_error = text.charAt(70);
						var media = text.charAt(88);
						var head = text.charAt(87);
						var pause = text.charAt(84);
						// check each flag that prevents printing
						if (is_error == '0')
						{
							ok = true;
							statuses.push("Ready to Print");
						}
						if (media == '1')
							statuses.push("Paper out");
						if (media == '2')
							statuses.push("Ribbon Out");
						if (media == '4')
							statuses.push("Media Door Open");
						if (media == '8')
							statuses.push("Cutter Fault");
						if (head == '1')
							statuses.push("Printhead Overheating");
						if (head == '2')
							statuses.push("Motor Overheating");
						if (head == '4')
							statuses.push("Printhead Fault");
						if (head == '8')
							statuses.push("Incorrect Printhead");
						if (pause == '1')
							statuses.push("Printer Paused");
						if ((!ok) && (statuses.Count == 0))
							statuses.push("Error: Unknown Error");
						finishedFunction(statuses.join());
			}, 
			printerError);
};

/**
 *  Print the package label.
 */
function sendPrinterData(pck){
	checkPrinterStatus(function (text){
		if (text == "Ready to Print"){
		    // So here I will get the location from the located information.
		    function printInventoryInfo(pck, index){
		        var uuid = pck.M_inventoried[index]
		        server.entityManager.getEntityByUuid(uuid, false, 
		            function(inventory, caller){
		                function getLocalisationInfo (uuid, localisations, pck){
		                    server.entityManager.getEntityByUuid(uuid, false, 
		                    function(localisation, caller){
		                        // The string to display is simply the split value of M_id
		                        caller.localisations.push(localisation)
		                        if(isObjectReference(localisation.M_parent)){
		                            getLocalisationInfo(localisation.M_parent, caller.localisations, caller.pck)
		                        }else{
		                            var index = 0;
		                            var localisationStr = ""
		                            
		                            while(caller.localisations.length > 0){
		                                localisation = caller.localisations.pop();
		                                localisationStr += localisation.M_name
		                                if(caller.localisations.length > 0){
		                                    localisationStr += " / ";
		                                }
		                                index++
		                            }
                        		    // The localisation label.
                        		    var s1 = "^XA"
                                    + "^CF0," + scaleValue(30) + "," + scaleValue(30)
                                    + "^FT" + scaleValue(50) + "," + scaleValue(50)
                                    + "^FD" + localisationStr +  "^FS"
                                    + "^FO" + scaleValue(50) + "," + scaleValue(60)
                                    + "^BY,2.0," + scaleValue(60)
                                    + "^B3,,,N"
                                    + "^FD" + caller.pck.M_id + "^FS"// Bar code.
                                    + "^FT" + scaleValue(50) + "," + scaleValue(152)
                                    + "^FD" + caller.pck.M_id + "^FS"
                                    + "^XZ";
                        			defaultPrinter.send(s1, printComplete, printerError);
		                        }
		                    },
		                    function(err, caller){
		                        
		                    },{"localisations":localisations, "pck":pck})
		                }
		                
		                getLocalisationInfo(inventory.M_located, [], pck)
		                if(caller.index < caller.pck.M_inventoried.length - 1){
		                    printInventoryInfo(caller.pck, caller.index + 1)
		                }
		            },
		            function(errObj, caller){
		                
		            }, 
		            {"pck":pck, "index":index})
		    }
		    
		    // Print the inventory (localisation) inforamtion.
		    printInventoryInfo(pck, 0)
		    
		    // Print the Items informations.
		    function printItemsInfo(items, index){
		        
		        server.entityManager.getEntityByUuid(items[index], false, 
		            function(item, caller){
		                var alias = ""
		                for(var i=0; i < item.M_alias.length; i++){
		                    alias += item.M_alias[i]
		                    if(i < item.M_alias.length - 1){
		                        alias += " / ";
		                    }
		                }
		                
		                // Now the item description...
		                parser=new DOMParser();
                        var divs =parser.parseFromString(item.M_description, "text/html").getElementsByTagName("div");
                        var description = ""
                        for(var i=0; i < divs.length; i++){
                            if(divs[i].lang == server.languageManager.language){
                                description = divs[i].innerHTML
                                break;
                            }
                        }
		                
		                s2 = "^XA"
                        + "^CF0," + scaleValue(30) + "," + scaleValue(30)
                        + "^FO" + scaleValue(50) + "," + scaleValue(20)
                        + "^BY,2.0," + scaleValue(60)
                        + "^B3,,,N"
                        + "^FD" + item.M_id + "^FS"
                        + "^FT" + scaleValue(50) + "," + scaleValue(110)
                        + "^FDP/N : " + item.M_id + "   Ref : " + alias + "^FS"
                        + "^FT" + scaleValue(50) + "," + scaleValue(148)
                        + "^FD" + description + "^FS"
                        + "^XZ";
                        defaultPrinter.send(s2, printComplete, printerError);
		                if(caller.index < caller.items.length - 1){
		                    printItemsInfo(caller.items, caller.index + 1)
		                }
		            },
		            function(errObj, caller){
		                
		            }, {"items":items, "index":index})
		    }
		    
            // Now print the item information.
            printItemsInfo(pck.M_items, 0)
			
		}
		else
		{
			printerError(text);
		}
	});
};

/**
 *  Recursively print the content of the localistion package.
 */
function printLocalisationLabels(localisation, callback){
    var sublocalisations = []
    var packages = []
    
    function getSublocalisations(sublocalisations, localisation, index, callback){
        var uuid = localisation.M_subLocalisations[index]
        server.entityManager.getEntityByUuid(uuid, false, 
            function(sublocalisation, caller){
                caller.sublocalisations.push(sublocalisation)
                if(caller.index < caller.localisation.M_subLocalisations.length-1){
                    // Here the callback look for further sub-localisation.
                    if(sublocalisation.M_subLocalisations.length > 0){
                        getSublocalisations(caller.sublocalisations, sublocalisation, 0, function(caller){
                            return function(){
                                getSublocalisations(caller.sublocalisations, caller.localisation, caller.index + 1, caller.callback)
                            }
                        }(caller))
                    }else{
                        getSublocalisations(caller.sublocalisations, caller.localisation, caller.index + 1, caller.callback)
                    }
                }else{
                    caller.callback(caller.sublocalisations)
                }
            },
            function(errObj, caller){
                
            }, {"sublocalisations":sublocalisations, "localisation":localisation, "index":index, "callback":callback})
    }
    
    
    function getPacackages(inventories, packages, index, callback){
        var inventory = inventories[index]
        if(isObject(inventory)){
            server.entityManager.getEntityByUuid(inventory.M_package, false,
                // success callback
                function(result, caller){
                    caller.packages.push(result)
                    if(caller.index < caller.inventories.length - 1){
                        getPacackages(caller.inventories, caller.packages, caller.index + 1, caller.callback)
                    }else{
                        caller.callback(caller.packages)
                    }
                },
                // error callback
                function(){
                    
                }, 
                // caller
                {"inventories":inventories, "packages":packages, "index":index, "callback":callback})
        }else{
            // Here I will get the inventory object and then set it packages.
            server.entityManager.getEntityByUuid(inventory, false,
                function(inventory, caller){
                    caller.inventories[caller.index] = inventory
                    // get packages...
                    getPacackages(caller.inventories, caller.packages, caller.index, caller.callback)
                },
                function(){
                    
                }, {"inventories":inventories, "packages":packages, "index":index, "callback":callback})
        }
    }
    
    if(localisation.M_subLocalisations.length > 0){
        getSublocalisations(sublocalisations, localisation, 0, function(sublocalisations){
            var inventories = []
            for(var i=0; i < sublocalisations.length; i++){
                inventories = inventories.concat(sublocalisations[i].M_inventories)
            }
            getPacackages(inventories, packages, 0, function(packages){
                printPackagesLabels(packages)
            })
        })
    }else{
        getPacackages(localisation.M_inventories, packages, 0, function(packages){
            printPackagesLabels(packages)
        })
    }
}

/**
 *  Print labels for one package.
 */
function printPackagesLabels(pcks){
    // Set the default printer...
    BrowserPrint.getDefaultDevice('printer', function(printer){
        if(printer == undefined){
           mainPage.showNotification("danger", "Veuillez configurer votre imprimante, <a href='http://mon176:9393/Catalogue_2/zebra/download/software-browser-print-user-guide-en-us.pdf'>aide</a>", 4000)
        }
        
        defaultPrinter = printer;
        BrowserPrint.getLocalDevices(function(printers){
            availablePrinters = printers;
            // Display the list of available printers.
            if(printers.printers == undefined){
                mainPage.showNotification("danger", "Veuillez configurer votre imprimante, <a href='http://mon176:9393/Catalogue_2/zebra/download/software-browser-print-user-guide-en-us.pdf'>aide</a>", 4000)
                return
            }
            // so here I will set the print dialog...
            var printDialog = new Dialog(randomUUID(), undefined, true)
            printDialog.title.element.innerHTML = "Imprimer les étiquettes pour l'item " + pck.M_id
            var select = printDialog.content.appendElement({ "tag": "span", "innerHtml": "Veuillez Selectionner l'imprimante "})
                .appendElement({"tag":"select", "class":"form-control"}).down()
           

            for(var i=0; i < printers.printer.length; i++){
                availablePrinters[printers.printer[i].name] = printers.printer[i]
                select.appendElement({"tag":"option", "value":printers.printer[i].name, "innerHtml":printers.printer[i].name})
            }
            
            // Set the default printer.
            select.element.value = defaultPrinter.name
            
            select.element.onchange = function(){
                // set the default printer.
                defaultPrinter = availablePrinters[this.value]
            }
            
            printDialog.ok.element.focus();
            printDialog.div.element.style.maxWidth = "650px"
            printDialog.setCentered()
            printDialog.header.element.style.padding = "5px";
            printDialog.header.element.style.color = "white";
            
            printDialog.header.element.classList.add("bg-dark")
            printDialog.ok.element.classList.remove("btn-default")
            printDialog.ok.element.classList.add("btn-primary")
            printDialog.ok.element.classList.add("mr-3")
            printDialog.deleteBtn.element.style.color = "white"
            printDialog.deleteBtn.element.style.fontSize="16px"
            printDialog.ok.element.innerHTML = "imprimer"
            
            printDialog.ok.element.onclick = function (pcks, dialog) {
                return function(){
                    dialog.close()
                    for(var i=0; i < pcks.length; i++){
                        sendPrinterData(pcks[i])
                    }
                }
            }(pcks, printDialog)
        })
        
    })
}

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
     .appendElement({"tag": "li","class" : "nav-item"}).down()
    .appendElement({"tag" : "a", "class" : "nav-link ", "href" : "#order-adminContent", "role" :"tab", "data-toggle" : "tab", "aria-controls" : "order-adminContent", "id":"order-adminTab"}).down()
    .appendElement({"tag" : "span", "innerHtml" : "Commandes"})
    .appendElement({"tag" : "span", "style" : "margin-left:5px;", "id" : "adminOrderSaveState"}).up().up()
    .appendElement({"tag" : "li", "style" : "padding:10px;display:flex;"}).down()
    .appendElement({"tag" : "button", "class" : "btn btn-primary disabled ml-3", "innerHtml" : "Enregistrer tout", "style" : "padding:1px 6px;"}).up().up()
    .appendElement({"tag" : "div", "class" : "tab-content"}).down()
    .appendElement({"tag" : "div", "class" : "tab-pane active show result-tab-content","id" : "supplier-adminContent", "role" : "tabpanel", "aria-labelledby" : "supplier-adminTab" })
    .appendElement({"tag" : "div", "class" : "tab-pane result-tab-content","id" : "localisation-adminContent", "role" : "tabpanel", "aria-labelledby" : "localisation-adminTab" })
    .appendElement({"tag" : "div", "class" : "tab-pane show result-tab-content","id" : "manufacturer-adminContent", "role" : "tabpanel", "aria-labelledby" : "manufacturer-adminTab" })
    .appendElement({"tag" : "div", "class" : "tab-pane show result-tab-content","id" : "category-adminContent", "role" : "tabpanel", "aria-labelledby" : "category-adminTab" })
    .appendElement({"tag" : "div", "class" : "tab-pane show result-tab-content","id" : "item-adminContent", "role" : "tabpanel", "aria-labelledby" : "item-adminTab" })
    .appendElement({"tag" : "div", "class" : "tab-pane show result-tab-content","id" : "package-adminContent", "role" : "tabpanel", "aria-labelledby" : "package-adminTab" })
    .appendElement({"tag" : "div", "class" : "tab-pane show result-tab-content","id" : "order-adminContent", "role" : "tabpanel", "aria-labelledby" : "order-adminTab" })
  
    this.adminItemPage = new AdminItemPage(this.panel.getChildById("item-adminContent"));
    
    this.adminSupplierPage = new AdminSupplierPage(this.panel.getChildById("supplier-adminContent"))
    
    this.adminLocalisationPage = new AdminLocalisationPage(this.panel.getChildById("localisation-adminContent"))
  
    this.adminManufacturerPage = new AdminManufacturerPage(this.panel.getChildById("manufacturer-adminContent"))
    
    this.adminCategoryPage = new AdminCategoryPage(this.panel.getChildById("category-adminContent"))
    
    this.adminPackagePage = new AdminPackagePage(this.panel.getChildById("package-adminContent"))
    
    this.adminOrderPage = new AdminOrderPage(this.panel.getChildById("order-adminContent"))
  
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


