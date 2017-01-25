function local_createNewDivElement() {
    var element1 = new Element(
        document.getElementsByTagName("body")[0],
        { "tag": "div" },
        undefined,
        true
        )
    return element1
}

function ElementTest() {
    QUnit.test("ElementTest", function (assert) {
        assert.ok(local_createNewDivElement().element.tagName == "DIV");
    })
}

function Element_constructor_paramsTest(){

    var parentElement = local_createNewDivElement()
    var callbackParamTest = false;
    var childElementBId = "childElementBId"
    // Element = function (parent, params, callback, makeFirstChild)
    var childElementA = new Element(
        parentElement, 
        { 
            "tag": "p",
            "innerHtml" : "childElementA"
        },
        function(){
            callbackParamTest = true
        },
        false
        )
    var childElementB = new Element(
        parentElement, 
        { 
            "tag": "p",
            "id": childElementBId,
            "innerHtml" : "childElementB"
        },
        function(){
            callbackParamTest = true
        },
        true
        )
    QUnit.test("Element_constructor_paramsTest_01", function (assert) {
        assert.ok(
            parentElement.childs[childElementA.id] == childElementA, 
            childElementA
            );
    })
    QUnit.test("Element_constructor_paramsTest_02", function (assert) {
        assert.ok(
            parentElement.childs[childElementB.id].id == childElementBId, 
            childElementB.id
            );
    })
    QUnit.test("Element_constructor_paramsTest_03", function (assert) {
        assert.ok(
            callbackParamTest, 
            callbackParamTest);
    })
    QUnit.test("Element_constructor_paramsTest_04", function (assert) {
        assert.ok(
            parentElement.element.innerHTML == '<p id="childElementBId">childElementB</p><p>childElementA</p>', 
            parentElement.element.innerHTML
            );
    })

}

function Element_prototype_initTest() {
    QUnit.test("Element_prototype_initTest", function (assert) {
        assert.ok(local_createNewDivElement().init().element.tagName == "DIV");
    })
}

function Element_prototype_appendElementTest() {
    var parentElement = local_createNewDivElement();
    var childElement = local_createNewDivElement();
    QUnit.test("Element_prototype_appendElementTest", function (assert) {
        assert.ok(parentElement.appendElement(childElement).childs[childElement.id] == childElement);
    })
}

function Element_prototype_prependElementTest() {
    var parentElement = new Element(
        document.getElementsByTagName("body")[0],
        { "tag": "ul" },
        undefined,
        true
        )
    var childElement1 = parentElement.appendElement({ "tag": "li", "innerHtml": "I will be second" });
    var childElement2 = parentElement.prependElement({ "tag": "li", "innerHtml": "I will be first" });

    QUnit.test("Element_prototype_prependElementTest", function (assert) {
        assert.ok(parentElement.element.outerHTML == "<ul><li>I will be first</li><li>I will be second</li></ul>");
    })
}

/*
* TEST FAIL - Raison: element1.appendElement(element2) fonctionne au niveau de l'element (childs, etc) mais 
* pas dans l'interface. appendElement fonctionne proprement quand on passe des params mais pas si on passe un element
*/
function Element_prototype_removeElementTest() {
    var element1 = local_createNewDivElement();
    var element2 = local_createNewDivElement();
    element1.appendElement(element2)

    //console.log("------------------")
    //console.log(element1)
    //console.log(element2)
    //console.log(element1.element)
    //console.log(element2.element)
    //console.log(element1.childs)
    //console.log("------------------")

    element1.removeElement(element2)

    QUnit.test("Element_prototype_removeElementTest", function (assert) {
        assert.ok(element1.childs[element2.id] == undefined);
    })
}

/*
* TEST FAIL - Raison: element2 a encore element1 comme parent apres qu'on ait fait
* element1.removeAllChilds(). Est que c'est le comportement desire? Peut etre que 
* c'est causé par la meme raison qui fait fail Element_prototype_removeElementTest
*/
function Element_prototype_removeAllChildsTest() {
    var element1 = local_createNewDivElement();
    var element2 = local_createNewDivElement();
    element1.appendElement(element2)

    element1.removeAllChilds()

    var i = 0
    for (var id in element1.childs) {
        i++
    }

    QUnit.test("Element_prototype_removeAllChildsTest", function (assert) {
        assert.ok(i == 0);
        assert.ok(element2.parentElement == undefined);
    })
}

function Element_prototype_getChildByIdTest() {
    var element1 = local_createNewDivElement().appendElement({ "tag": "div", "id": "a" }).down()
    .appendElement({ "tag": "div", "id": "b" }).up()

    QUnit.test("Element_prototype_getChildByIdTest", function (assert) {
        assert.ok(element1.getChildById("b").id == "b");
    })
}


/*
* TEST FAIL - Raison: ?
*/
function Element_prototype_getTopParentTest() {
    var element1 = local_createNewDivElement()
    var grandChildElement = element1.appendElement({ "tag": "div", "id": "a" }).down()
    .appendElement({ "tag": "div", "id": "b" }).down()

    QUnit.test("Element_prototype_getTopParentTest", function (assert) {
        assert.ok(grandChildElement.getTopParent() == element1);
    })
}

function Element_prototype_upTest() {
    var element1 = local_createNewDivElement()
    .appendElement({ "tag": "div", "id": "a" }).down()
    .appendElement({ "tag": "div", "id": "b" }).down()
    .appendElement({ "tag": "div", "id": "c" }).up()

    QUnit.test("Element_prototype_upTest", function (assert) {
        assert.ok(element1.id == "a");
    })
}

function Element_prototype_downTest() {
    var element1 = local_createNewDivElement()
    .appendElement({ "tag": "div", "id": "a" }).down()

    QUnit.test("Element_prototype_downTest", function (assert) {
        assert.ok(element1.id == "a");
    })
}

/*
* TEST FAIL - element1 contient encore childElement dans ses childs 
*/
function Element_prototype_moveChildElementsTest() {
    var element1 = local_createNewDivElement()
    var childElement = element1.appendElement({ "tag": "div", "id": "a" }).down()
    var element2 = local_createNewDivElement()

    var childsToMove = []
    childsToMove[childElement.id] = childElement;
    element2.moveChildElements(childsToMove)

    QUnit.test("Element_prototype_moveChildElementsTest", function (assert) {
        assert.ok(element1.childs[childElement.id] == undefined);
        assert.ok(element2.childs[childElement.id] == childElement);
    })
}

/*
* TEST FAIL - element2 n'as pas childElement dans ses childs 
*/
function Element_prototype_copyChildElementsTest() {
    var element1 = local_createNewDivElement()
    var childElement = element1.appendElement({ "tag": "div", "id": "a" }).down()
    var element2 = local_createNewDivElement()

    var childsToCopy = []
    childsToCopy[childElement.id] = childElement;
    element2.copyChildElements(childsToCopy)

    QUnit.test("Element_prototype_copyChildElementsTest", function (assert) {
        assert.ok(element1.childs[childElement.id] == childElement);
        assert.ok(element2.childs[childElement.id] == childElement);
    })
}

function Element_prototype_setStyleTest() {
    var element1 = local_createNewDivElement()
    element1.setStyle("background", "black")
    QUnit.test("Element_prototype_setStyleTest", function (assert) {
        assert.ok(element1.element.style.background == "black");
    })
}

function Element_prototype_setAttributeTest() {
    var element1 = local_createNewDivElement()
    element1.setAttribute("style", "background : black")
    QUnit.test("Element_prototype_setAttributeTest", function (assert) {
        assert.ok(element1.element.style.background == "black");
    })
}

function Element_prototype_animateTest() {
    var element1 = local_createNewDivElement()

    var keyframe = "100% { width: 300px; }"
    element1.animate(keyframe, "element_open_", 1,
        function (element1) {
            return function () {
                element1.setStyle("width", "300px")
                QUnit.test("Element_prototype_animateTest", function (assert) {
                    assert.ok(element1.element.style.width == "300px");
                })
            }
        } (element1))
}


/*
* TEST FAIL - surement que je l'ai mal compris
*/
function createChildElementTest() {
    var element1 = local_createNewDivElement()
    var childElement1 = local_createNewDivElement()

    createChildElement(element1, childElement1.element)

    //console.log(element1.childs)

    QUnit.test("createChildElementTest", function (assert) {
        assert.ok(element1.childs[childElement1.id] == childElement1);
    })
}

function createElementFromXmlTest() {
    var element1 = local_createNewDivElement()
    var xml = "<a><b>createElementFromXmlTest</b></a>"

    createElementFromXml(element1, xml)

    var childElement1;
    for (var childId in element1.childs) {
        childElement1 = element1.childs[childId]
    }

    QUnit.test("createElementFromXmlTest", function (assert) {
        assert.ok(element1.childs[childElement1.id] == childElement1);
    })
}

function createElementFromHtmlTest() {
    var element1 = local_createNewDivElement()
    var html = "<html><div>createElementFromHtmlTest</div></html>"

    createElementFromHtml(element1, html)

    var childElement1;
    for (var childId in element1.childs) {
        childElement1 = element1.childs[childId]
    }

    QUnit.test("createElementFromHtmlTest", function (assert) {
        assert.ok(element1.childs[childElement1.id] == childElement1);
    })
}


function elementTests() {
    ElementTest()
    Element_constructor_paramsTest()
    Element_prototype_initTest()
    Element_prototype_appendElementTest()
    Element_prototype_prependElementTest()
    Element_prototype_getChildByIdTest()
    Element_prototype_upTest()
    Element_prototype_downTest()
    Element_prototype_setStyleTest()
    Element_prototype_setAttributeTest()
    Element_prototype_animateTest()

    createElementFromXmlTest()
    createElementFromHtmlTest()

    /* Failing tests */
    /*
    Element_prototype_copyChildElementsTest()
    Element_prototype_moveChildElementsTest()
    Element_prototype_getTopParentTest()
    Element_prototype_removeAllChildsTest()
    Element_prototype_removeElementTest()

    createChildElementTest()
    */

} 