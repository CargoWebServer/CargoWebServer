var events_nbTests = 0;
var events_nbTestsAsserted = 0;
var events_EventManager_attach_test_number = 300;
var events_EventManager_detach_test_number = 301;

function events_allTestsAsserted_Test() {
    QUnit.test("events_allTestsAsserted_Test",
        function (assert) {
            assert.ok(
                events_nbTests == events_nbTestsAsserted, 
                "events_nbTests: " + events_nbTests + ", events_nbTestsAsserted: " + events_nbTestsAsserted);
        })
}


function events_EventManager_constructor_test() {
    events_nbTests++
    var eventManagerId = "myEventManager"
    var eventManagerEventsGroup = "TestEvent"
    var myEventManager = new EventManager(eventManagerId, eventManagerEventsGroup)

    QUnit.test("events_EventManager_constructor_test",
        function (assert) {
            assert.ok(
                myEventManager.id == eventManagerId && 
                myEventManager.eventsGroup == eventManagerEventsGroup,
                "id: " + myEventManager.id + " eventsGroup: " + myEventManager.eventsGroup)
            events_nbTestsAsserted++
        })
}

function events_EventManager_attach_test() {
    events_nbTests++
    var eventManagerId = "myEventManager_attach"
    var eventManagerEventsGroup = "TestEvent"
    var myEventManager_attach = new EventManager(eventManagerId, eventManagerEventsGroup)
    var myElement = new Element(
        document.getElementsByTagName("body")[0],
        { "tag": "div" })
    myEventManager_attach.attach(
        myElement, 
        events_EventManager_attach_test_number, 
        function(){
            myElement.element.innerHTML += "Event Received"
        })

    QUnit.test("events_EventManager_attach_test",
        function (assert) {
            assert.ok(
                myElement.observable.id == eventManagerId && 
                myElement.observable.observers[300] != undefined,
                "myElement.observable.id: " + myElement.observable.id)
            events_nbTestsAsserted++
        })
}

function events_EventManager_detach_test() {
    events_nbTests++
    var eventManagerId = "myEventManager_detach"
    var eventManagerEventsGroup = "TestEvent"
    var myEventManager_detach = new EventManager(eventManagerId, eventManagerEventsGroup)
    var myElement = new Element(
        document.getElementsByTagName("body")[0],
        { "tag": "div" })
    myEventManager_detach.attach(
        myElement, 
        events_EventManager_detach_test_number, 
        function(){
            myElement.element.innerHTML += "Event Received"
        })
    myEventManager_detach.detach(
        myElement, 
        events_EventManager_detach_test_number)

    QUnit.test("events_EventManager_detach_test",
        function (assert) {
            assert.ok(
                myElement.observable == undefined,
                "myElement.observable: " + myElement.observable)
            events_nbTestsAsserted++
        })
}


function eventTests() {

    events_EventManager_constructor_test()
    events_EventManager_attach_test()
    events_EventManager_detach_test()

    setTimeout(function () {
        return function () {
            events_allTestsAsserted_Test()
        }
    } (), 3000);

} 