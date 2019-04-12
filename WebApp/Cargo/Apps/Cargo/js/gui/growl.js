/*
 * The panel that will contain the message to display to the user.
 */
var GrowlPanel = function(){

	// That map contain the list of message to be display...
	this.messages = {}

	// The parent of the growl is the body element.
	var body = document.getElementsByTagName("body")[0]
	this.div = new Element(body, {"tag":"div", "id" : "growl_panel"})

	return this
}

GrowlPanel.prototype.displayMessage = function(message, messageType, autoHide){

	var message = new GrowlMessagePanel(this, message, messageType, autoHide)
	this.messages[message.id] = message
	message.div.element.className += " show"

	return message
}

var GrowlMessagePanel = function(parent, message, messageType, autoHide){
	this.displayTime = 3000
	this.parent = parent

	this.div = parent.div.appendElement({"tag":"div", "class":"growl_message_panel"}).down()
	this.id = this.div.id

	// The message header...
	this.header = this.div.appendElement({"tag":"div", "class":"growl_message_panel_header"}).down()
	
	this.icon = this.header.appendElement({"tag":"div", "class":"message_icon"}).down()
        .appendElement({"tag":"i", "class":"fa  fa-info-circle"}).down()

	// The close button.
	this.closeBtn = this.header.appendElement({"tag":"div", "class":"message_delete_button"}).down()
        .appendElement({"tag":"i", "class":"fa fa-close"}).down()

	// Now the text to diplay to the user...
	this.message = this.div.appendElement({"tag":"span", "innerHtml":message}).down()

	// the function to execute to remove the message...
	var hidePanelFct = function(growlPanel, messagePanel){
			return function(){
				var endAnimationListenerName = "animationend"

				if(getNavigatorName() == "Chrome" || getNavigatorName() == "Safari"){
					endAnimationListenerName = "webkitAnimationEnd"
				}

				// End animation event...
				messagePanel.div.element.addEventListener(endAnimationListenerName, function(growlPanel, messagePanel){
					return function(){
						messagePanel.div.element.style.display = "none"
					}
				}(growlPanel, messagePanel))

				messagePanel.div.element.className = "growl_message_panel hide"

			}
		}(this.parent, this)

	// show the message for tow second...
	if(autoHide == true){
		setTimeout(hidePanelFct, this.displayTime);
	}else{
		this.closeBtn.element.onclick = hidePanelFct
	}

	return this;
}