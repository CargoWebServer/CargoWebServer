var events_nbTests = 0;
var events_nbTestsAsserted = 0;
var events_EventManager_attach_test_number = 300;
var events_EventManager_detach_test_number = 301;
var events_EventChannel_broadcastEvent_test_number = 302;

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
    var eventManagerEventsGroup = "TestEvent"
    var myEventManager = new EventManager(eventManagerEventsGroup)

    QUnit.test("events_EventManager_constructor_test",
        function (assert) {
            assert.ok(
                myEventManager.channelId == eventManagerEventsGroup,
                " eventsGroup: " + myEventManager.eventsGroup)
            events_nbTestsAsserted++
        })
}

function events_EventManager_attach_test() {
    events_nbTests++
    var eventManagerEventsGroup = "TestEvent"
    var myEventManager_attach = new EventManager(eventManagerEventsGroup)
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
                myElement.observable.id == myEventManager_attach.id && 
                myElement.observable.observers[300] != undefined,
                "myElement.observable.id: " + myElement.observable.id)
            events_nbTestsAsserted++
        })
}

function events_EventManager_detach_test() {
    events_nbTests++
    var eventManagerEventsGroup = "TestEvent"
    var myEventManager_detach = new EventManager(eventManagerEventsGroup)
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

function events_EventChannel_broadcastEvent_test() {
    events_nbTests++
    var eventManagerEventsGroup = "TestEvent"
    var myEventManager_broadcastEvent = new EventManager(eventManagerEventsGroup)

    server.eventHandler.addEventManager(
        myEventManager_broadcastEvent,
        function () {
            server.eventHandler.appendEventFilter(
                "\\.*",
                eventManagerEventsGroup,
                function () { 

                    var myElement = new Element(
                        document.getElementsByTagName("body")[0],
                        { "tag": "div" })
                    var message = "events_EventChannel_broadcastEvent_test_message"

                    myEventManager_broadcastEvent.attach(
                        myElement, 
                        events_EventChannel_broadcastEvent_test_number, 
                        function(evt){
                            console.log(evt)
                            myElement.element.innerHTML += evt.dataMap.message

                            QUnit.test("events_EventChannel_broadcastEvent_test",
                                function (assert) {
                                    assert.ok(
                                        evt.dataMap.message == message, 
                                        evt.dataMap.message)
                                    events_nbTestsAsserted++
                                })

                        })

                    var evt = {
                        "code": events_EventChannel_broadcastEvent_test_number, 
                        "channelId": eventManagerEventsGroup,
                        "dataMap": {
                            "message" : message
                        }
                    }

                    server.eventHandler.BroadcastEvent(evt)
                },
                function () { },
                undefined)
        }
        )




    

    
    

    
}


function eventTests() {

    events_EventManager_constructor_test()
    events_EventManager_attach_test()
    events_EventManager_detach_test()
    events_EventChannel_broadcastEvent_test()

    setTimeout(function () {
        return function () {
            events_allTestsAsserted_Test()
        }
    } (), 3000);

} 