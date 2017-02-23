// init syntax highlighter
hljs.initHighlightingOnLoad();

// top navigation
$('#bs-example-navbar-collapse-1 li').click(function (e) {
	$('#bs-example-navbar-collapse-1 li.active').removeClass('active');
	var $this = $(this);
	if (!$this.hasClass('active')) {
		$this.addClass('active');
	}
	e.preventDefault();
	hidePages(this.id)
});

$('#a-cargo').click(function (e) {
	$('#bs-example-navbar-collapse-1 li.active').removeClass('active');
	e.preventDefault();
	hidePages(this.id)
});

$('#li-gettingStarted').click(function (e) {
	hidePages("page-gettingStarted")
});

function hidePages(exceptId) {
	$('.pageRow').addClass('hidden')
	$('#page-' + exceptId.split('-')[1]).removeClass('hidden')
}

// tutorials navigation
$(".tutorials-link").click(function () {
	$('html, body').animate({
		scrollTop: $('#tutorialsSection-' + this.id.split('-')[1]).offset().top
	}, 500);
});

$("#frontEndNextTutorialLink").click(function () {
	$('html, body').animate({
		scrollTop: $('#tutorialsSection-AsynchronousProgramming').offset().top
	}, 500);
});

$("#dowloadBtn").click(function () {
	var lnk = this
	var OSName = "Unknown OS";
	if (navigator.appVersion.indexOf("Win") != -1) OSName = "Windows";
	else if (navigator.appVersion.indexOf("Mac") != -1) OSName = "MacOS";
	else if (navigator.appVersion.indexOf("X11") != -1) OSName = "UNIX";
	else if (navigator.appVersion.indexOf("Linux") != -1) OSName = "Linux";

	if (OSName == "Windows") {
		lnk.href = "http://cargowebserver.com/distro/windows/64/CargoWebServer.zip"
	} else if (OSName == "Linux" || OSName == "UNIX" ) {
		lnk.href = "http://cargowebserver.com/distro/linux/64/CargoWebServer.tar.gz"
	}

	lnk.click()
});

$('#getStarted-btn').click(function(){
	$('#li-gettingStarted').click()
})

// main
main()

function main() {

	// Element example
	var containerElement = new Element(
		document.getElementById("elementExampleContainer"),
		{ "tag": "div", "style": "height: 100%; width: 100%;" }
	);
	var headerElement = containerElement.appendElement({
		"tag": "div",
		"style": "height: 10%; width: 100%; background: black; color: white; text-align:center"
	}).down();

	headerElement.element.innerHTML = "header"

	containerElement.appendElement({
		"id": "middle",
		"tag": "div",
		"style": "height: 80%; width: 100%; color: white; text-align:center",
	}).down().appendElement({
		"tag": "div",
		"style": "height: 50px; width: 200px; background: lightgray; color: black;",
		"innerHtml": "rectangle"
	}).up().appendElement({
		"tag": "div",
		"style": "height: 10%; width: 100%; background: darkgray; color: white; text-align:center",
		"innerHtml": "footer"
	})

	containerElement.getChildById("middle").element.style.background = "gray"

	/*
	// Event manager
	var myEventManager = new EventManager("myEventManager", "myEvent")

	server.eventHandler.AddEventManager(
		myEventManager,
		function () {
			server.eventHandler.appendEventFilter(
				"\\.*",
				"myEvent",
				function () { },
				function () { },
				undefined)
		}
		)

	// broadcast local event 
	var evt = {
		"code": 100, 
		"name": "myEvent",
		"dataMap": {
			"message" : "hello"
		}
	}

	server.eventHandler.BroadcastEvent(evt)

	// broadcast network event 

	var entityInfo = {
		"TYPENAME": "Server.MessageData",
		"Name": "name",
		"Value": { 
			"message": "hello world" 
		}
	}

	server.eventHandler.broadcastEventData(
		101,
		"myEvent",
		[entityInfo],
		function () { },
		function () { },
		undefined
		)

	// attach event 
	myEventManager.attach(
		element,
		100,
		function (evt) {
			element.innerHTML = evt.dataMap["message"]
		})
	*/

}
