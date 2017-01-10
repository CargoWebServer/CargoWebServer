
var applicationName = document.getElementsByTagName("title")[0].text
var mainPage = null

function main() {

    var bodyElement = new Element(document.getElementsByTagName("body")[0], { "tag": "div", "style": "height: 100%; width: 100%;" });

    var tutorialsBanner = bodyElement.appendElement({ "tag": "div", "class": "tutorialsBanner", "innerHtml": "TUTORIALS" }).down()

    var tutorialsFlexConainer = bodyElement.appendElement({ "tag": "div", "class": "flexboxRowContainer" }).down()

    var tutorialsSidebarContainer = tutorialsFlexConainer.appendElement({ "tag": "div", "class": "tutorialsSidebarContainer" }).down()
    var tutorialsSidebarMenu = tutorialsSidebarContainer.appendElement({ "tag": "div", "id": "tutorialsSidebarMenu", "class": "flexboxColumnContainer" }).down()

    tutorialsSidebarMenu.appendElement({ "tag": "div", "class": "sidebarMenuTitle", "innerHtml": "Getting Started" })
        .appendElement({ "tag": "div", "class": "sidebarMenuItem", "innerHtml": "Welcome", "id": "WelcomeTutorialMenuDiv" })
        .appendElement({ "tag": "div", "class": "sidebarMenuItem", "innerHtml": "Licence" })
        .appendElement({ "tag": "div", "class": "sidebarMenuItem", "innerHtml": "Installation" })
        .appendElement({ "tag": "div", "class": "sidebarMenuItem", "innerHtml": "Configuration" })
        .appendElement({ "tag": "div", "class": "sidebarMenuTitle", "innerHtml": "Basics" })
        .appendElement({ "tag": "div", "class": "sidebarMenuItem", "innerHtml": "'Hello World' Project" })
        .appendElement({ "tag": "div", "class": "sidebarMenuItem", "innerHtml": "Front-End JS Framework" })
        .appendElement({ "tag": "div", "class": "sidebarMenuItem", "innerHtml": "Back-End Golang Server" })
        .appendElement({ "tag": "div", "class": "sidebarMenuItem", "innerHtml": "Element", "id": "ElementTutorialMenuDiv" })
        .appendElement({ "tag": "div", "class": "sidebarMenuTitle", "innerHtml": "Core Modules" })
        .appendElement({ "tag": "div", "class": "sidebarMenuItem", "innerHtml": "Events" })
        .appendElement({ "tag": "div", "class": "sidebarMenuItem", "innerHtml": "Entities" })
        .appendElement({ "tag": "div", "class": "sidebarMenuItem", "innerHtml": "Data" })
        .appendElement({ "tag": "div", "class": "sidebarMenuTitle", "innerHtml": "Advanced Tutorials" })
        .appendElement({ "tag": "div", "class": "sidebarMenuItem", "innerHtml": "MVC Applications" })
        .appendElement({ "tag": "div", "class": "sidebarMenuItem", "innerHtml": "Permissions" })
        .appendElement({ "tag": "div", "class": "sidebarMenuItem", "innerHtml": "..." })

    var tutorialsContentContainer = tutorialsFlexConainer.appendElement({ "tag": "div", "class": "tutorialsContentContainer" }).down()


    tutorialsContentContainer.appendElement({ "tag": "h1", "class": "tutorialTitleHeader", "id": "elementTutorial_h1" })
        .appendElement({ "tag": "h2", "class": "sectionTitle", "id": "sectionTitle_introduction_h2" })
        .appendElement({ "tag": "div", "class": "collapsible", "id": "collapsible_introduction" }).down()
        .appendElement({ "tag": "p", "id": "introduction_p" }).up()
        .appendElement({ "tag": "h2", "class": "sectionTitle", "id": "sectionTitle_elementClassOverview_h2" })
        .appendElement({ "tag": "div", "class": "collapsible", "id": "collapsible_elementClassOverview" }).down()
        .appendElement({ "tag": "h3", "class": "sectionTitle", "id": "sectionTitle_elementConstructor_h3" })
        .appendElement({ "tag": "div", "class": "collapsible", "id": "collapsible_elementConstructor" }).down()
        .appendElement({ "tag": "p", "id": "newElement_p" })
        .appendElement({ "tag": "div", "id": "aceEditor_reference_newElement_div" })
        .appendElement({ "tag": "ul", "id": "newElementParameters_ul" }).down()
        .appendElement({ "tag": "li", "id": "newElementParent_li" })
        .appendElement({ "tag": "li", "id": "newElementParams_li" })
        .appendElement({ "tag": "li", "id": "newElementCallback_li" })
        .appendElement({ "tag": "li", "id": "newElementIsAppendFront_li" }).up().up()
        .appendElement({ "tag": "h3", "class": "sectionTitle", "id": "sectionTitle_essentialFunctions_h3" })
        .appendElement({ "tag": "div", "class": "collapsible", "id": "collapsible_essentialFunctions" }).down()
        .appendElement({ "tag": "h4", "class": "sectionTitle", "id": "sectionTitle_appendElement_h4" })
        .appendElement({ "tag": "div", "class": "collapsible", "id": "collapsible_appendElement" }).down()
        .appendElement({ "tag": "p", "id": "appendElement_p" })
        .appendElement({ "tag": "div", "id": "aceEditor_reference_appendElement_div" })
        .appendElement({ "tag": "ul", "id": "appendElementParameters_ul" }).down()
        .appendElement({ "tag": "li", "id": "appendElementChildElement_li" }).up().up()
        .appendElement({ "tag": "h4", "class": "sectionTitle", "id": "sectionTitle_removeElement_h4" })
        .appendElement({ "tag": "div", "class": "collapsible", "id": "collapsible_removeElement" }).down()
        .appendElement({ "tag": "p", "id": "removeElement_p" })
        .appendElement({ "tag": "div", "id": "aceEditor_reference_removeElement_div" })
        .appendElement({ "tag": "ul", "id": "removeElementParameters_ul" }).down()
        .appendElement({ "tag": "li", "id": "removeElementChildElement_li" }).up().up()
        .appendElement({ "tag": "h4", "class": "sectionTitle", "id": "sectionTitle_getChildById_h4" })
        .appendElement({ "tag": "div", "class": "collapsible", "id": "collapsible_getChildById" }).down()
        .appendElement({ "tag": "p", "id": "getChildById_p" })
        .appendElement({ "tag": "div", "id": "aceEditor_reference_getChildById_div" })
        .appendElement({ "tag": "ul", "id": "getChildByIdParameters_ul" }).down()
        .appendElement({ "tag": "li", "id": "getChildByIdChildElement_li" }).up().up()
        .appendElement({ "tag": "h4", "class": "sectionTitle", "id": "sectionTitle_down_h4" })
        .appendElement({ "tag": "div", "class": "collapsible", "id": "collapsible_down" }).down()
        .appendElement({ "tag": "p", "id": "down_p" })
        .appendElement({ "tag": "div", "id": "aceEditor_reference_down_div" }).up()
        .appendElement({ "tag": "h4", "class": "sectionTitle", "id": "sectionTitle_up_h4" })
        .appendElement({ "tag": "div", "class": "collapsible", "id": "collapsible_up" }).down()
        .appendElement({ "tag": "p", "id": "up_p" })
        .appendElement({ "tag": "div", "id": "aceEditor_reference_up_div" }).up().up().up()
        .appendElement({ "tag": "h2", "class": "sectionTitle", "id": "sectionTitle_example1_h2" })
        .appendElement({ "tag": "div", "class": "collapsible", "id": "collapsible_example1" }).down()
        .appendElement({ "tag": "p", "id": "example1_01_p" })
        .appendElement({ "tag": "p", "id": "example1_02_p" })
        .appendElement({ "tag": "div", "id": "aceEditor_example_example1_02_div" })
        .appendElement({ "tag": "p", "id": "example1_03_p" })
        .appendElement({ "tag": "p", "id": "example1_04_p" })
        .appendElement({ "tag": "div", "id": "aceEditor_example_example1_04_div" })
        .appendElement({ "tag": "p", "id": "example1_05_p" }).up()



    server.fileManager.getFileByPath(
        "/ElementTutorial/txt/elementTutorialText.txt",
        -1,
        function () {
        },
        function (result) {
            var splitFile = decode64(result.M_data).split("###")
            for (var i = 2; i < splitFile.length; i = i + 2) {

                var elementToGetId = splitFile[i - 1]
                var elementToGet = bodyElement.getChildById(elementToGetId)

                if (elementToGet != undefined) {
                    elementToGet.element.innerHTML = splitFile[i].trim()

                    if (elementToGetId.startsWith("sectionTitle")) {
                        var collapsibleElement = tutorialsContentContainer.getChildById("collapsible_" + elementToGetId.split("_")[1])

                        elementToGet.element.addEventListener('click', function (collapsibleElement) {
                            return function () {
                                var targetHeight = "0px"
                                // TODO Calculer initialHeight autrement, pcq fonctionne mal si on collapse des sous divs avant de cliquer sur celui-ci
                                if (collapsibleElement.initialHeight == undefined) {
                                    collapsibleElement.initialHeight = collapsibleElement.element.offsetHeight
                                }
                                else if (collapsibleElement.element.style.height == "0px") {
                                    targetHeight = collapsibleElement.initialHeight + "px"
                                }

                                var keyframe = "100% { height: " + targetHeight + "}"
                                collapsibleElement.animate(keyframe, 1,
                                    function (collapsibleElement, targetHeight) {
                                        return function () {
                                            collapsibleElement.element.style.height = targetHeight
                                        }
                                    } (collapsibleElement, targetHeight))
                            }
                        } (collapsibleElement))
                    }
                    else if (elementToGetId.startsWith("aceEditor")) {
                        var editor = ace.edit(elementToGetId);
                        editor.getSession().setMode("ace/mode/javascript");
                        editor.setOptions({
                            fontSize: "15pt",
                            maxLines: Infinity
                        });
                        editorType = elementToGetId.split("_")[1]
                        if (editorType == "reference") {
                            elementToGet.element.style.width = "900px"
                            editor.setTheme("ace/theme/chrome");
                        } else if (editorType == "example") {
                            editor.setTheme("ace/theme/monokai");
                        }
                    }
                }
            }

        },
        function () {
        },
        undefined)

}