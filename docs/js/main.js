
// The application tite.
var applicationName = document.getElementsByTagName("title")[0].text

// Connection with the cargo server.
//var server = new Server("CargoWebServer.com", "54.218.110.52", 9393)
var server = new Server("localhost", "127.0.0.1", 9393)

// Set the language from the html lang attribute.
server.languageManager.setLanguage($('html').attr('lang'))

// The entry point of the application.
function main() {
	// init syntax highlighter
	hljs.initHighlightingOnLoad();

	// Create the home page.
	var homePage = new HomePage(document.getElementsByTagName("body")[0])

	// Create the get started page.
	var gettingSatartedPage = new GettingStartedPage(homePage.pages)

	// The tutorial page.
	var tutorialsPage = new TutorialsPage(homePage.pages)

	// The documentation page 
	var documentationPage = new DocumentationPage(homePage.pages)
}
