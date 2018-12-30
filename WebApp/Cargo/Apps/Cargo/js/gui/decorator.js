
/*
 * This class is use to insert auto-completion functionnality
 * to an existing text box...
 * @constructor
 * @param control The associated control where the message came from...
 * @param selectValue if definied that function is call when the value is selected.
 */
var autocompleteDiv = null
function attachAutoComplete(control, elementLst, autoComplete, selectValue) {
    if (control.element.parentNode != undefined) {
        control.element.parentNode.style.position = "relative"
    }

    // The value must be in the list...
    if (autoComplete == undefined) {
        autoComplete = true
    }
    
    control.element.addEventListener("blur", 
        function(selectValue){
            return function(evt){
                if(this.style.display == "none" && autocompleteDiv != null){
                    // if the mouse is not over the autocompleteDiv...
                    var box = autocompleteDiv.element.getBoundingClientRect();
                    // If the mouse i not over the box.
                    var isOver = getMouseX() > box.left && getMouseX() < box.right && getMouseY() < box.bottom && getMouseY() > box.top;
                    if(!isOver){
                        var isSelect = false;
                        for(var id in autocompleteDiv.childs){
                            var c = autocompleteDiv.childs[id];
                            if(c.element.style.backgroundColor == "darkgrey"){
                                c.element.click()
                                isSelect = true
                                break;
                            }
                        }
                        // if  no selection is made...
                        if(!isSelect){
                            autocompleteDiv.removeAllChilds()
                            autocompleteDiv.element.style.display = "none"
                            this.value = "";
                        }
                    }
                }
        }
    }(selectValue), true)
        
    var currentIndex = -1
    
    /* Save the key down event **/
    control.element.addEventListener("keyup", function (control,elementLst, autoComplete, selectValue) {
        return function (evt) {

            /* The div that contain items **/
            var coord = getCoords(control.element)
            var minWidth = control.element.offsetWidth + "px"
            if (control.element.value.length >= 1) {
                if(autocompleteDiv == null){
                    autocompleteDiv = new Element(document.getElementsByTagName("body")[0], { "tag": "div","id":"autocompleteDiv", "class": "autoCompleteDiv", "style": "z-index: 10;" })
                }else{
                    autocompleteDiv.removeAllChilds()
                    autocompleteDiv.element.style.display = "none";
                }
                
                autocompleteDiv.element.style.minWidth = minWidth
                autocompleteDiv.element.style.top = coord.top + control.element.offsetHeight + "px";
                autocompleteDiv.element.style.left = coord.left + -1 + "px";
                autocompleteDiv.element.style.display = "";
                
                // Filter the values...
                values = _.select(elementLst, function (val) {
                    return function (inputValue) {
                        if (isString(inputValue)) {
                            val = val.replace("[]", "")
                            return inputValue.substring(0, val.length).toUpperCase() == val.toUpperCase()
                        } else if (isInt(inputValue) || isNumeric(inputValue)) {
                            inputValue.toString().substring(0, val.length).toUpperCase() == val.toUpperCase()
                        }
                    }
                }(control.element.value))

                if (values.length > 1 || (values.length >= 1 && !autoComplete)) {
                    // Append the element...
                    for (var i = 0; i < values.length; i++) {
                        var value = values[i]
                        if (!isString(value)) {
                            value = value.toString()
                        }
                        var elementDiv = autocompleteDiv.appendElement({ "tag": "div", "innerHtml": value, "contenteditable":"true", "style": "display: block;", "id": i }).down()
                        // Here i will append the click event...
                        elementDiv.element.onclick = function (control, autocompleteDiv, value, selectValue) {
                            return function (evt) {
                                evt.stopPropagation()
                                
                                if (control.element.value != value) {
                                     control.element.value = value
                                    if (selectValue != undefined) {
                                        // Here the caller want to get control...
                                        selectValue(value)
                                    }
                                }
                                if (control.element.onchange != null) {
                                    control.element.onchange()
                                }
                                currentIndex = -1
                                autocompleteDiv.removeAllChilds()
                                autocompleteDiv.element.style.display = "none"
                            }
                        }(control, autocompleteDiv, values[i], selectValue)
                    }

                    // Display the list...
                    autocompleteDiv.element.style.display = "block"
                } else if (values.length == 1) {
                    if (selectValue == undefined) {
                        control.element.value = values[0]
                        control.element.select()
                    } else {
                        selectValue(values[0])
                    }
                    autocompleteDiv.element.style.display = "none"
                } else if (values.length == 0) {
                    autocompleteDiv.element.style.display = "none"
                    currentIndex = -1
                }
            } else if (autocompleteDiv != undefined) {
                autocompleteDiv.removeAllChilds();
                autocompleteDiv.element.style.display = "none"
            }

            if (evt.keyCode == 40 || evt.keyCode == 38 || evt.keyCode == 13) {
                var max = Object.keys(autocompleteDiv.childs).length
                if (max > 0) {
                    var index = currentIndex
                    if (evt.keyCode == 40) {
                        if (index < max - 1) {
                            index++
                        }
                    } else if (evt.keyCode == 38) {
                        if (index > 0) {
                            index--
                        }
                    }

                    if (index == -1) {
                        autocompleteDiv.element.style.display = "none"
                        return
                    }

                    var c = autocompleteDiv.childs[index]
                    c.element.style.backgroundColor = "darkgrey"
                    if (evt.keyCode == 13) {
                        c.element.click()
                        currentIndex = -1
                    }
                    currentIndex = index
                }
            }else if(evt.keyCode == 27){
                for(var id in autocompleteDiv.childs){
                    var c = autocompleteDiv.childs[id];
                    c.element.style.backgroundColor = ""
                }
                // exit the selection.
                control.element.blur()
            }
        }
    }(control, elementLst, autoComplete, selectValue), true)

}

/* This function is use to informe the user if he make a mistake...
 * @param msg The message to display
 * @param control The control where the error appear...
 * @param validator A function that validate
 */
function setValidator(msg, control, validator, delay) {
    if (control == null) {
        return
    }
    // Set the control retalive...
    control.element.parentNode.style.position = "relative"
    var errorDiv = new Element(control.element.parentNode, { "tag": "div", "class": "errorDiv", "style": "display:none;" })
    var msgDiv = errorDiv.appendElement({ "tag": "div", "class": "errorMsgDiv", "style": "font-size: 10pt", "innerHtml": msg }).down()
    var triangle = errorDiv.appendElement({ "tag": "div", "class": "triangle-down" }).down()
    var display = true

    /* Needed for final validation **/
    if (control.validators == undefined) {
        control.validators = []
    }
    control.validators.push(validator)

    /* Remove the red border if there is one here... **/
    control.element.addEventListener("keyup", function () {
        this.style.border = ""
    });

    var validationCall = function (validator, control, errorDiv, msgDiv) {
        return function () {
            if (!display) {
                return
            }

            setTimeout(function (validator, msgDiv, errorDiv, control, delay) {
                return function () {
                    if (validator(msgDiv) == false) {
                        // The validor widow...
                        errorDiv.element.style.display = "block"
                        errorDiv.element.style.top = (-1 * (control.element.clientHeight + errorDiv.element.clientHeight - 4)) + "px"
                        errorDiv.element.style.left = control.element.offsetLeft + "px"

                        // I will remove the red border if the user change the text..
                        control.element.addEventListener("change",
                            function (control) {
                                return function () {
                                    control.element.style.border = ""
                                }
                            }(control))

                        // Here i will display the error message for 2 second...
                        setTimeout(function (errorDiv, control) {
                            return function () {
                                errorDiv.element.style.display = "none"
                                errorDiv.element.style.top = "0px"
                                errorDiv.element.style.left = "0px"
                                control.element.value = ""
                                control.element.style.border = "1px solid red"
                                display = false
                            }
                        }(errorDiv, control), delay)
                    }
                }
            }(validator, msgDiv, errorDiv, control, delay), 200); // 200 ms wait before validating...
        }
    }(validator, control, errorDiv, msgDiv)

    control.element.addEventListener("blur", validationCall)

    // Keep a reference to the function call so it will be possible to reusse it latter at validation...
    validator.validate = validationCall

    /* Reactivate the validation here **/
    control.element.addEventListener("change", function () {
        display = true
    })

    return this
}
