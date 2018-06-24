/**
 * That little compnent display the list of image inside a given
 * directory.
 */
var ImagePicker = function(parent, path){
    
    this.panel = parent.appendElement({"tag":"div", "class":"image_picker"}).down()
    
    // Display the local image selector.
    this.selectImageBtn = this.panel.appendElement({"tag":"span", "class":"append_picture_btn"}).down()
    
    // set the icon.
    this.selectImageBtn.appendElement({"tag":"i", "class":"fa fa-plus"})
    
    // The item images path.
    this.path = path;
    
    // Now I will create the image...
    server.fileManager.readDir(path,
        function(results, caller){
            if(results[0] == null){
                var index = caller.path.lastIndexOf("/");
                server.fileManager.createDir(caller.path.substr(index+1), caller.path.substr(0, index), 
                    // success callback.
                    function(dir, caller){
                        
                    },
                    // error callback.
                    function(errObj, caller){
                        console.log(errObj)
                    },{})   
            }else{
                for(var i=0; i < results[0].length; i++){
                    caller.imagePicker.appendImage(caller.path + "/" + results[0][i])
                }
            }
        },function(errObj,caller){
            console.log("error")
            // If the directory dosent exist i will create it.
            var index = caller.path.lastIndexOf("/");
            server.fileManager.createDir(caller.path.substr(index+1), caller.path.substr(0, index), 
                // success callback.
                function(dir, caller){
                    
                },
                // error callback.
                function(errObj, caller){
                    console.log(errObj)
                },{})            
        }, {"imagePicker":this, "path":path})
    
    this.imagePanels = {}
    
    // Ajout d'un image a l'aide du selecteur de fichier.
    this.fileInput = this.panel.appendElement({ "tag": "input", "type": "file", "multiple": "", "style": "display: none;" }).down()

    this.selectImageBtn.element.onclick = function (fileInput) {
        return function () {
            fileInput.element.click()
        }
    } (this.fileInput)

    this.fileInput.element.onchange = function (imagePicker) {
        return function (evt) {
            /* Here I will get the list of files... **/
            var files = evt.target.files // FileList object.
            for (var i = 0, f; f = files[i]; i++) {
                imagePicker.uploadFile(f, imagePicker.path)
            }
        }
    } (this)

    return this;
}

ImagePicker.prototype.uploadFile = function(file, path){

   // Upload the file here.
   var formData = new FormData()
   
   formData.append("multiplefiles", file, file.name)
   formData.append("path", path)
   // Use the post function to upload the file to the server.
   var xhr = new XMLHttpRequest()
   xhr.open('POST', '/uploads', true)
   // In case of error or success...
   xhr.onload = function (xhr, filePath, imagePicker) {
       return function (e) {
           if (xhr.readyState === 4) {
               if (xhr.status === 200) {
                   // Here I will create the file...
                   imagePicker.appendImage(filePath)
               } else {
                   console.log(xhr.statusText);
               }
           }
       }
   } (xhr, path + "/" + file.name, this)
   
   // now the progress event...
   xhr.upload.onprogress = function (e){
           
    }
    
    xhr.send(formData);
}

ImagePicker.prototype.appendImage = function(path){
    this.imagePanels[path] = new ImagePanel(this.panel, path)
}

/**
 * Display a single image.
 */
var ImagePanel = function(parent, path){
    this.path = path
    this.panel = parent.appendElement({"tag":"div", "class":"image_panel"}).down()
    
    this.image = this.panel.appendElement({"tag":"img", "src":"http://" + server.hostName + ":"+ server.port + path}).down()

    // Now the delete button...
    this.deleteBtn = this.panel.appendElement({"tag":"button", "class":"btn btn-danger btn-sm", "style":"position: absolute; top:2px; right:2px; display: none;"}).down()
    this.deleteBtn.appendElement({"tag":"i", "class":"fa fa-trash"})
    
    this.deleteBtn.element.onmouseenter = function(){
        this.style.cursor = "pointer";
    }
    
    this.deleteBtn.element.onmouseleave = function(){
        this.style.cursor = "default";
    }
    
    // Now the delete warning...
    this.deleteBtn.element.onclick = function(imagePanel){
        return function(e){
            // this.panel
            e.stopPropagation()
            var confirmDialog = new Dialog(randomUUID(), undefined, true)
            confirmDialog.title.element.innerHTML = "Supprimer l'image"
            confirmDialog.content.appendElement({ "tag": "span", "innerHtml": "Voulez-vous supprimer " + imagePanel.path + "?" })
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
            
            confirmDialog.ok.element.onclick = function (dialog, imagePanel) {
                return function () {
                    // I will call delete file
                    imagePanel.panel.element.parentNode.removeChild(imagePanel.panel.element)
                    
                    server.fileManager.removeFile(imagePanel.path,
                        function(result, caller){
                            /** nothing to do here */
                        },
                        function(errObj, caller){
                            /** display error here **/
                        }, {})
                    
                    dialog.close()
                }
            } (confirmDialog, imagePanel)
        }
    }(this)
    
    this.panel.element.onmouseenter = function(deleteBtn){
        return function(){
            deleteBtn.element.style.display = "block";
        }
    }(this.deleteBtn)
    
    this.panel.element.onmouseleave = function(deleteBtn){
        return function(){
            deleteBtn.element.style.display = "none";
        }
    }(this.deleteBtn)
}