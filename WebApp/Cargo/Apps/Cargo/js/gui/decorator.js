
/*
 * This class is use to insert auto-completion functionnality
 * to an existing text box...
 * @constructor
 * @param control The associated control where the message came from...
 * @param initLstCallback init callback
 * @param callback the function that is call to filter items...
 * @param finishCallback this is call when the value is retreive from the list...
 */
function attachAutoComplete(control, elementLst, autoComplete) {
    control.element.parentNode.style.position = "relative"

    // The value must be in the list...
    if (autoComplete == undefined) {
        autoComplete = true
    }

    // I will always append the auto complete inside the body element.
    control.autocompleteDiv = new Element(document.getElementsByTagName("body")[0], { "tag": "div", "class": "autoCompleteDiv", "style": " display:none;" })
    var currentIndex = -1

    /* Save the key down event **/
    control.element.addEventListener("keyup", function (control, autocompleteDiv, elementLst, autoComplete) {
        return function (evt) {

            /* The div that contain items **/
            /* var viewportOffset = control.element.parentNode.getBoundingClientRect();
            var doc = document.documentElement;
            var left = (window.pageXOffset || doc.scrollLeft) - (doc.clientLeft || 0);
            var top = (window.pageYOffset || doc.scrollTop) - (doc.clientTop || 0);*/
            var coord = getCoords(control.element)
            var minWidth = control.element.offsetWidth + "px"
            autocompleteDiv.element.style.minWidth = minWidth
            autocompleteDiv.element.style.top =  coord.top + control.element.offsetHeight + "px";
            autocompleteDiv.element.style.left = coord.left + -1 + "px";
            autocompleteDiv.removeAllChilds()

            if (control.element.value.length >= 1) {
                // Filter the values...
                values = _.select(elementLst, function (val) {
                    return function (user) {
                        return user.substring(0, val.length).toUpperCase() == val.toUpperCase()
                    }
                } (control.element.value))

                if (values.length > 1 || (values.length >= 1 && !autoComplete)) {
                    // Append the element...
                    for (var i = 0; i < values.length; i++) {
                        var elementDiv = autocompleteDiv.appendElement({ "tag": "div", "innerHtml": values[i], "style": "display: block;", "id": i }).down()
                        // Here i will append the click event...
                        elementDiv.element.onclick = function (control, autocompleteDiv, value) {
                            return function () {
                                autocompleteDiv.removeAllChilds()
                                autocompleteDiv.element.style.display = "none"
                                control.element.value = value
                                currentIndex = -1
                                if (control.element.onchange != null) {
                                    control.element.onchange()
                                }
                            }
                        } (control, autocompleteDiv, values[i])
                    }

                    // Display the list...
                    autocompleteDiv.element.style.display = "block"
                } else if (values.length == 1) {
                    control.element.value = values[0]
                    control.element.select()
                    autocompleteDiv.element.style.display = "none"
                } else if (values.length == 0) {
                    autocompleteDiv.element.style.display = "none"
                    currentIndex = -1
                }
            } else {
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
            }
        }
    } (control, control.autocompleteDiv, elementLst, autoComplete), true)

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

    var validationCall = function (validator, control, errorDiv) {
        return function () {
            if (!display) {
                return
            }

            setTimeout(function (validator, errorDiv, control, delay) {
                return function () {
                    if (validator() == false) {
                        // The validor widow...
                        errorDiv.element.style.display = "block"
                        errorDiv.element.style.top = (-1 * (control.element.clientHeight + errorDiv.element.clientHeight)) + "px"
                        errorDiv.element.style.left = control.element.offsetLeft + "px"

                        // I will remove the red border if the user change the text..
                        control.element.addEventListener("change",
                            function (control) {
                                return function () {
                                    control.element.style.border = ""
                                }
                            } (control))

                        // Here i will display the error message for 2 second...
                        setTimeout(function (errorDiv, control) {
                            return function () {
                                errorDiv.element.style.display = "none"
                                errorDiv.element.style.top = "0px"
                                errorDiv.element.style.left = "0px"
                                control.element.value = ""
                                control.element.focus()
                                control.element.style.border = "1px solid red"
                                display = false
                            }
                        } (errorDiv, control), delay)
                    }
                }
            } (validator, errorDiv, control, delay), 200); // 200 ms wait before validating...
        }
    } (validator, control, errorDiv)

    control.element.addEventListener("blur", validationCall)

    // Keep a reference to the function call so it will be possible to reusse it latter at validation...
    validator.validate = validationCall

    /* Reactivate the validation here **/
    control.element.addEventListener("change", function () {
        display = true
    })
}
