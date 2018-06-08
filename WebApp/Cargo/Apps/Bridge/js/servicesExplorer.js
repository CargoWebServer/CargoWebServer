var ServicesExplorer = function (parent) {

    this.parent = parent
    this.panel = parent.appendElement({ "tag": "div", "class": "data_explorer" }).down()

    // set the scrolling shadow...
    this.panel.element.onscroll = function (header) {
        return function () {
            var position = this.scrollTop;
            if (this.scrollTop > 0) {
                if (header.className.indexOf(" scrolling") == -1) {
                    header.className += " scrolling"
                }
            } else {
                header.className = header.className.replaceAll(" scrolling", "")
            }
        }
    }(this.parent.element.firstChild)

    this.parent.element.firstChild.style.borderBottom = "1px solid gray"

    // Set the resize event.
    window.addEventListener('resize',
        function (taskInstancesExplorer) {
            return function () {
                taskInstancesExplorer.resize()
            }
        }(this), true);

    // The list of services.
    this.servicesView = {}
    
    return this
}

ServicesExplorer.prototype.resize = function () {
    var height = this.parent.element.offsetHeight - this.parent.element.firstChild.offsetHeight;
    this.panel.element.style.height = height - 10 + "px"
}

/**
 * Display the data schema of a given data store.
 */
ServicesExplorer.prototype.initService = function (serviceConfig, initCallback) {
    // init one time.
    if (this.servicesView[serviceConfig.M_id] != undefined) {
        return
    }

    this.servicesView[serviceConfig.M_id] = new Element(this.panel, { "tag": "div", "class": "shemas_view", "style":"display: none;" })
    
    if(Object.keys(this.servicesView).length === 1){
        this.servicesView[serviceConfig.M_id].element.style.display = "";
    }
    
    // Now I will get the list of action for a given services.
    server.serviceManager.getServiceActions(serviceConfig.M_id,
        // success callback
        function (results, parent) {
            // Now I will display the list of action in panel.
            for (var i = 0; i < results.length; i++) {
                var result = results[i]
                new EntityPanel(parent, result.TYPENAME, function (entity) {
                    return function (panel) {
                        panel.setEntity(entity);
                        panel.panel.element.style.padding = "10px";
                        panel.panel.element.style.borderTop = "1px solid";
                        panel.panel.element.style.marginBottom =  "30px";
                        
                        // Now I will display the action documentation correctly...
                        var documentationInput = CargoEntities.Action_M_doc
                        //panel.fields["M_doc"].panel.element.style.display = "none"
                        var doc = panel.fields["M_doc"].value.element.innerText
                        if (doc.indexOf("@src") != -1) {
                            doc = doc.split("@src")[0]
                        }

                        var values = doc.split("@")
                        doc = ""
                        for (var i = 0; i < values.length; i++) {
                            doc += "<div>"
                            if (values[i].startsWith("api")) {
                                doc += values[i].replaceAll("api 1.0", "<span style='color: green;'>api 1.0</span>")
                            } else if (values[i].startsWith("param") && values[i].indexOf("{callback}") == -1) {
                                var values_ = values[i].split("param")[1].split(" ")
                                doc += "<span class='doc_tag' style='vertical-align: top;'>param</span><span>"
                                var description = ""
                                for (var j = 1; j < values_.length; j++) {
                                    if (j == 1) {
                                        // The type:
                                        doc += "<span style='color: darkgreen; vertical-align: text-top'>" + values_[j] + "</span>"
                                    } else if (j == 2) {
                                        // The name
                                        doc += "<span style='color: color: #657383; font-weight:bold; vertical-align: text-top'>" + values_[j] + "</span>"
                                    } else {
                                        description = description + " " + values_[j]
                                    }
                                }
                                doc += "<span style='vertical-align: text-top'>" + description + "</span>"
                                doc += "</span>"
                            }
                            doc += "</div>"
                        }
                        panel.fields["M_doc"].value.element.innerHTML = doc
                    }
                }(result), undefined, false, result, "")
            }
        },
        // error callback
        function (errObj, caller) {

        }, this.servicesView[serviceConfig.M_id])  
}


/**
 * Display the data schema of a given data store.
 */
ServicesExplorer.prototype.setService = function (serviceId) {

    // here I will calculate the height...
    this.resize()

    for (var id in this.servicesView) {
        this.servicesView[id].element.style.display = "none"
    }
    
    if (this.servicesView[serviceId] != undefined) {
        this.servicesView[serviceId].element.style.display = ""
    }
}
