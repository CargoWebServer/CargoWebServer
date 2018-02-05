
var TabPanel = function(parent){
    this.parent = parent;
    if(parent == undefined){
        this.parent = null;
    }

    this.panel = new Element(this.parent, {"tag":"div", "class":"tab-panel"});

    // The header...
    this.header = this.panel.appendElement({"tag":"div", "class":"tab-panel-header"}).down();
    this.previousBtn = this.header.appendElement({"tag":"div", "class":"tab-panel-btn"}).down();
    this.previousBtn.appendElement({"tag":"i", "class":"fa fa-caret-left"})

    this.tabsDiv = this.header.appendElement({"tag":"div", "class":"tab-panel-tabs-div"}).down();
    this.nextBtn = this.header.appendElement({"tag":"div", "class":"tab-panel-btn"}).down();
    this.nextBtn.appendElement({"tag":"i", "class":"fa fa-caret-right"})

    // The content...
    this.content = this.panel.appendElement({"tag":"div", "class":"tab-panel-content"}).down();

    this.tabs = {};

    return this;
}

/**
 * Append a new Tab panel
 */
TabPanel.prototype.appendTab = function(id){
    var header =  this.tabsDiv.appendElement({"tag":"div", "id":id + "_tab", "class":"tab-div"}).down();
    var title = header.appendElement({"tag":"div"}).down();
    var closeBtn = header.appendElement({"tag":"div", "class":"tab-panel-close-btn"}).down()
    closeBtn.appendElement({"tag":"i", "class":"fa fa-times"})
    var content = this.content.appendElement({"tag":"div", "id":id+"_content"}).down();

    if(Object.keys(this.tabs).length == 0){
        header.element.className = "tab-div active"
    }

    var tab = {"title":title, "content":content}
    this.tabs[id] = tab;
    header.element.onmouseover = function(closeBtn){
        return function(){
            closeBtn.element.style.visibility = "visible";
        }
    }(closeBtn)

    header.element.onmouseleave = function(closeBtn){
        return function(){
            closeBtn.element.style.visibility = "hidden";
        }
    }(closeBtn)

    /** The tab content. */
    header.element.onclick = function(content, tabs){
        return function(){
            for(var id in tabs){
                tabs[id].title.parentElement.element.className = "tab-div";
                tabs[id].content.element.className = "";
            }
            this.className = "tab-div active";
            content.element.className = "active";
        }
    }(content,  this.tabs)

    return tab;
}