/**
 * Contain the main function...
 */
var applicationName = document.getElementsByTagName("title")[0].text
var mainPage = null

/**
 * This function is the entry point of the application...
 */
function main() {
    
    // This will act as the main layout.
    var bodyElement = new Element(document.getElementsByTagName("body")[0], { "tag": "div", "style": "height: 100%; width: 100%;" });

    // The main grid layout...
    var mainPage = new MainPage(bodyElement)
}