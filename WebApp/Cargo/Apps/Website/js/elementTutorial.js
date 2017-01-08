var elementTutorial = function (parent) {

    var elementClassJs = "http://" + server.ipv4 + ":" + server.port + "/Cargo/js/element.js"

    var constructors = getContructorsInFile(elementClassJs)

    console.log(constructors)

    this.parent = parent

    this.parent.appendElement({
        "tag": "h1",
        "class": "tutorialTitleHeader",
        "innerHtml": "Element Tutorial"
    }).appendElement({
        "tag": "h2",
        "innerHtml": "Introduction"
    }).appendElement({
        "tag": "p",
        "innerHtml": "The 'Element' class included in the Cargo javascript framework offers a basic Html element " +
        "wrapper class to help you create client side graphical interfaces. All the functionnalities related to this class " +
        "can be found within the file <a>element.js</a>. You do <b>not</b> need to use the 'Element' class or any of it's related functionnalities " +
        "to develop a web application using the Cargo framework. If you are an experienced front-end web application developper and " +
        "already have techniques which you are familiar with, feel free to skip to the next tutorial. "
    }).appendElement({
        "tag": "h2",
        "innerHtml": "Using the Element class"
    }).appendElement({
        "tag": "p",
        "innerHtml": "To create a new 'Element', simply call it's constructor function. "
    }).appendElement({
        "tag": "div",
        "id": "tutorialEditor",
        "style": "height: 210px; text-align: left;",
        "innerHtml": "/** \n *@param parent The parent element of this element. Can be an Element or a DOM element." +
        "\n *@param params The list of parameters. \n *@param callback This function called after the initialization is completed. " +
        "\n *@param isAppendAtFront True: the element is put in front of the other elements. \n   False: it will be at the end. " +
        "\n *@returns {HTMLElement} \n */ \n" +
        "var myElement = new Element(parent, params, callback, isAppendAtFront)"
    }).appendElement({
        "tag": "div",
        "id": "test01",
        "style": "height: 210px; text-align: left;",
        "innerHtml": '<object style="width: 100%" data="' + elementClassJs + '" type="text/html"><embed src="' + elementClassJs + '" type="text/html" /></object>'
    })

    console.log(this.parent.getChildById("test01").element.data)
    console.log(this.parent.getChildById("test01").element.children[0].contentDocument.body.innerHTML)

    var editor = ace.edit("tutorialEditor");
    editor.setTheme("ace/theme/monokai");
    editor.getSession().setMode("ace/mode/javascript");
    editor.setOptions({ fontSize: "15pt" });


}

function getContructorsInFile(file) {
    var commentStart = file.split("/**");

    return commentStart
}