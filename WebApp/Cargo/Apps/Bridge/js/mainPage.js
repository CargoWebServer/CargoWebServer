/**
 * The main page is the page displayed when no user are logged in.
 */

var MainPage = function(parent, loginPage){

	// The main layout...
	this.parent = parent

	this.panel = this.parent.appendElement({"tag":"div", "class":"main_page"}).down()

	// The login page...
	this.loginPage = loginPage

    // Set the login page into the layout
	this.panel.moveChildElements( [this.loginPage.panel])
 
    // Set the focus to the element...
    this.loginPage.usernameInput.element.value = ""
    this.loginPage.passwordInput.element.value = ""
    this.loginPage.usernameInput.element.focus()
  
	return this
}