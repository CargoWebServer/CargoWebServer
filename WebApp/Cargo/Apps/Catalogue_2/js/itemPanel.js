/**
 * That function show items results.
 */
var ItemPanel = function (item, prototype) {
    // Here I got an item to display...
    if (prototype.SuperTypeNames.indexOf("CatalogSchema.ItemType") != -1) {
        var panel = this.parent.appendElement({ "tag": "div", "class": "item_result_panel" }).down()

        var col0 = panel.appendElement({ "tag": "div", "id": item.UUID + "_col_0", "style": "flex:1;flex-basis: 20%;" }).down() // first column 0
        var col1 = panel.appendElement({ "tag": "div", "id": item.UUID + "_col_1", "style": "flex:1;flex-basis: 60%; flex-grow: 1" }).down() // first column 1
        var col2 = panel.appendElement({ "tag": "div", "id": item.UUID + "_col_2", "style": "flex:1;flex-basis: 20%" }).down() // first column 2

        // The first column will contain the item image and localisation...
        var imgDiv = col0.appendElement({ "tag": "div", "style": "display: flex;" }).down()
        imgDiv_ = imgDiv.appendElement({ "tag": "div", "style": "width: 128px; height: 128px; border: 1px solid grey;" }).down()

        // here I will try to get the image...
        /*server.fileManager.downloadFile("photos/" + item.M_id, "photo_1.jpg", "image/jpeg",
            // Progress callback
            function (index, total, caller) {

            },
            // Success callback
            function (result, caller) {
                var reader = new FileReader();
                reader.onloadend = function (caller) {
                    return function (e) {
                        // Read as array buffer...
                        var imgUrl = resizeImage(e.target.result)
                        caller.appendElement({"tag":"img", "src":imgUrl, "style":"width: 128px; height: 128px;"})
                    }
                }(caller)
                // Read the file
                reader.readAsDataURL(result);
            },
            // Error callback.
            function () {
                console.log("inage not found!")
                imgDiv_.appendElement({ "tag": "i", "class": "fa fa-file-picture-o", "style": "padding-left: 5px;" })                
            }, imgDiv_)*/



        // Here I will use the get file by path to get the image.
        // base64toBlob(filePanel.fileInfo.M_data, filePanel.fileInfo.M_mime)
        // var filePath = "/Catalogue/photos/" + item.M_id + "/photo_1.jpg"
        //"photos/" + item.M_id.toUpperCase() + "/photo_1.jpg"
        /*server.fileManager.getFileByPath(filePath,
            // Success callback.
            function (result, caller) {
                // console.log("picture found!")
                result.M_
            },
            // Error callback.
            function (errObj, caller) {
                console.log(errObj)
            },
            {}
        )*/

        /* var q = {}
        var fieldsType = ["xs.string"]
        q.TypeName = "CargoEntities.File"
        q.Fields = ["M_thumbnail"]
        q.Query = 'CargoEntities.File.M_path == "/Catalogue/photos/' + item.M_id + '" && CargoEntities.File.M_name == "photo_1.jpg"'

        server.dataManager.read("CargoEntities", JSON.stringify(q), ["xs.string"], [],
            // success callback
            function (results, caller) {
                if (results.length == 1) {
                    if (results[0].length == 1) {
                        var thumbnail = results[0][0][0]
                        caller.appendElement({ "tag": "img", "src": thumbnail, "style": "width: 128px; height: 128px;" })
                        // TODO connect the image viewer here.
                    } else {
                        caller.appendElement({ "tag": "i", "class": "fa fa-file-picture-o", "style": "padding-left: 5px;" })
                    }
                }
            },
            // progress callback
            function (index, total, caller) {

            },
            // error callback
            function (errObj, caller) {
                console.log(errObj)
            }, imgDiv_)
            */

        // The localisation... div.
        var locDiv = col0.appendElement({ "tag": "div", "style": "display: flex;" }).down()
        server.entityManager.getEntityById("CatalogSchema.PackageType", "CatalogSchema", [item.M_id], false,
            // Success callback
            function(result, caller){
                console.log(result)
            },
            // Error callback
            function(){

            },
            {})

        // The second column will contain the description div and the specifications.
        var descriptionDiv = col1.appendElement({ "tag": "div", "style": "display: flex;" }).down()
        var specsDiv = col1.appendElement({ "tag": "div", "style": "display: table; border-spacing:10px 2px;" }).down()

        // That function will display the propertie.
        function displayPropertie(div, UUID, value, propertie, propertieType) {
            var isArray = propertieType.indexOf("[]") == 0;
            if (isArray) {
                propertieType.replace("[]", "")
            }

            if (isXsString(propertieType)) {
                var str = value
                if (value.M_valueOf != undefined) {
                    str = value.M_valueOf
                }
                if (!isArray) {
                    div.appendElement({ "tag": "label", "for": UUID + "_" + propertie + "_div", "style": "display: table-cell;", "innerHtml": propertie.substr(2).replaceAll("_", " ") + ":" })
                }

                div.appendElement({ "tag": "div", "id": UUID + "_" + propertie + "_div", "style": "display: table-cell;", "innerHtml": str })

            } else if (propertieType == "CatalogSchema.DimensionType") {
                // In that case I will display the dimension type.
                var dimension = value.M_valueOf + " " + value.M_unitOfMeasure.M_valueOf
                if (!isArray) {
                    div.appendElement({ "tag": "label", "for": UUID + "_" + propertie + "_div", "style": "display: table-cell;", "innerHtml": propertie.substr(2).replaceAll("_", " ") + ":" })
                }

                div.appendElement({ "tag": "div", "id": UUID  + "_" + propertie +  "_div", "style": "display: table-cell;", "innerHtml": dimension })

            } else if (propertieType.endsWith(":Ref")) {
                // So Here i will display a link to the other entity...
                item["set_" + propertie + "_" + value + "_ref"](
                    function (div) {
                        return function (subItem) {
                            // TODO display the sub item link...
                            var lnk = div.appendElement({ "tag": "a", "href":"#", "id": UUID  + "_" + propertie +  "_lnk", "style": "display: table-cell;", "innerHtml": subItem.M_id })
                            lnk.element.onclick = function(){
                                alert("--> display me!")
                            }

                        }
                    }(div))

            } else {
                var baseType = getBaseTypeExtension(propertieType)
                if (isXsBaseType(baseType)) {
                    displayPropertie(div, item.UUID, item[propertie], propertie, baseType)
                }
            }
        }

        // Now I wil display specs...
        for (var i = 0; i < prototype.Fields.length; i++) {
            var propertie = prototype.Fields[i];
            if (propertie.startsWith("M_")) {
                var propertieType = prototype.FieldsType[i]
                if (propertieType.startsWith("[]")) {
                    // First the array header.
                    if (propertie != "M_keywords" && propertie != "M_images" && propertie != "M_references") {
                        var specsDiv_ = specsDiv.appendElement({ "tag": "div", "style": "display: table-row;" })
                        specsDiv_.appendElement({ "tag": "label", "for": item.UUID + "_" + propertie + "_div", "style": "display: table-cell;", "innerHtml": propertie.substr(2).replaceAll("_", " ") + ":" })
                        var specsDiv__ = specsDiv_.appendElement({ "tag": "div", "style": "display: table; padding: 5px; width: 100%;" }).down();
                        for (var j = 0; j < item[propertie].length; j++) {
                            displayPropertie(specsDiv__.appendElement({ "tag": "div", "style": "display: table-row; width: 100%;" }).down(), item.UUID, item[propertie][j], propertie, propertieType)
                        }
                    }
                } else {
                    if (propertie != "M_id") {
                        displayPropertie(specsDiv.appendElement({ "tag": "div", "style": "display: table-row;" }).down(), item.UUID, item[propertie], propertie, propertieType)
                    }
                }
            }
        }

        // The third column div will contain the ovmm and the keyword div.
        var ovmmDiv = col2.appendElement({ "tag": "div", "style": "display: flex;" }).down()
        var lnk = ovmmDiv.appendElement({ "tag": "label", "for": item.UUID + "_id_span", "innerHtml": "ID: " })
            .appendElement({ "tag": "a", "href":"#", "id": item.UUID + "_id_lnk", "innerHtml": item.M_id, "style":"padding-left: 5px;"}).down()
        
        lnk.element.onclick = function(){
                alert("--> display me!")
            }

        var keywordsDiv = col2.appendElement({ "tag": "div", "style": "display: flex;" }).down()
        for (var i = 0; i < item.M_keywords.length; i++) {
            keywordsDiv.appendElement({ "tag": "span", "innerHtml": item.M_keywords[i] })
        }

    } else {
        // Here I get other property... I must found the item that has it property in order to display it.
        console.log(item)
    }

    // Return the item panel.
    return this;
}