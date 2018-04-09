
// fonction d'entré.
function main() {
    
    var path = "C:\\Users\\mm006819\\Documents\\CargoWebServer\\WebApp\\Cargo\\Apps\\interactive-product-rater\\machines"
    server.fileManager.readDir(path,
        // success
        function (results, path) {
            for (var i = 0; i < results[0].length; i++) {
                server.fileManager.readTextFile( path + "\\" +  results[0][i],
                    // Callback success
                    function (results, caller) {
                        var lines = results[0].split("\n")
                        for (var i = 0; i < lines.length; i++) {
                            var columns = lines[i].split(" ");
                            drawGraph(columns[2], /*"PROFILEUSE 0"*/  caller.fileName.replace(".SPF", ""))
                            document.getElementById("toto").innerText = path
                        }
                    },
                    // Error callback
                    function (errObj, caller) {

                    }, {"path":path, "fileName":results[0][i]})
            }
        },
        // error
        function (error, caller) {

        }, path)

}

/**
 * Cette fonciton afficher un cadran de performance pour une machine
 * @param {*} value La valeur a afficher sur le graph.
 * @param {*} machine 
 */
function drawGraph(value, machine) {
    var tempExecution = new Date()
    // Creation des element html (le div ou le graph est dessiné)
    var parentDiv = new Element(document.getElementsByTagName("body")[0], { "tag": "div", "id": randomUUID(), "style":"display: inline;" })
    var comment =   new Element(parentDiv, { "tag": "p",  "id": randomUUID() })// parentDiv.appendElement( { "tag": "p", "innerHtml":"Ceci est un test", "id": randomUUID() }).down()   
    var graphDiv = parentDiv.appendElement( { "tag": "div", "id": randomUUID() }).down()
    var leftGraph = parentDiv.appendElement( { "tag": "div", "id": randomUUID() }).down()

    comment.element.onclick =  function(){
        alert("Click!")
    }

    comment.element.innerText = machine

    function parseLine(index, line){
        var result = {
            "index":index,
            "values":[0,1,2,3]
        }

        return result
    }

    parseLine("1 2 3 4; 4 5 6 7")
    var graphParameters = {
        "type": "angulargauge",
        "renderAt": graphDiv.element.id,
        "width": "400",
        "height": "250",
        "dataFormat": "json",
        "dataSource": {
            "chart": {
                "caption": machine,
                "subcaption": tempExecution.toLocaleDateString() + " " + tempExecution.toLocaleTimeString(),
                "lowerLimit": "0",
                "upperLimit": "100",
                "theme": "fint"
            },
            "colorRange": {
                "color": [{
                    "minValue": "0",
                    "maxValue": "33",
                    "code": "E72F1D"
                },
                {
                    "minValue": "33",
                    "maxValue": "66",
                    "code": "FFE933"
                },
                {
                    "minValue": "66",
                    "maxValue": "100",
                    "code": "42FF33"
                }
                ]
            },
            "dials": {
                "dial": [{
                    "value": "0"
                }]
            }
        }
    }

    graphParameters.dataSource.dials.dial[0].value = value
    FusionCharts.ready(function () {
        var csatGauge = new FusionCharts(graphParameters);
        csatGauge.render();
    })
}