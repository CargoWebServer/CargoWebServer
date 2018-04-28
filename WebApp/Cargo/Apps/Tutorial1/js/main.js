function main () {
    var div = new Element(document.getElementsByTagName("body")[0],{"tag": "div", "class" : "test", "style":"width: auto; height: auto;"})
    var button = new Element(null, {"tag" : "button", "innerHtml" : "BOUTON"})

    for(var i = 0; i<10; i++){
        var totaux = div.appendElement({"tag": "div", "class" : "test"}).down()
        var span = totaux.appendElement({"tag":"span", "innerHtml":i.toString()}).down()
        totaux.element.onclick = function(div, button, span){
            return function(){
                //div.element.style.borderColor = "red"
                //this.style.backgroundColor = "red"
                //div.removeElement(toto);
                //this.parentNode.removeChild(this)
                this.appendChild(button.element);
                span.element.style.display = "none"
                button.element.onclick = function(div){
                    return function(){
                        div.parentNode.removeChild(div);
                    }
                }(this)
            }
        }(div,button, span)
    }
}