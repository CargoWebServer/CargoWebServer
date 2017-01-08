
var nbTests = 0;
var nbTestsAsserted = 0;

function AllFileTestsAsserted_Test() {
    QUnit.test("AllFileTestsAsserted_Test",
        function (assert) {
            assert.ok(nbTests == nbTestsAsserted, "nbTests: " + nbTests + ", nbTestsAsserted: " + nbTestsAsserted);
        })
}

/* FAIL: Pourtant le directory est bien crée
* Une fois que ca marche: essayer multiples endroits
*/
function createDir_Test() {
    nbTests++
    server.fileManager.createDir("createDir_Test", "",
        function (result) {
            QUnit.test("createDir_Test",
                function (result) {
                    return function (assert) {
                        console.log(result)
                        assert.ok(false);
                        nbTestsAsserted++
                    }
                } (result))
        }, function (result) {
        },
        function (result) {
        },
        undefined)
}

/**
 * PAS FINI DE TESTER TOUTES LES FONCTIONS
 */
function fileTests() {
    createDir_Test()

    setTimeout(function () {
        return function () {
            AllFileTestsAsserted_Test()
        }
    } (), 3000);
}


function testDownloadFile() {
    // Test file download...
    server.fileManager.downloadFile("/BrisOutil/Upload/700", "Hydrangeas.jpg", "application/vnd.openxmlformats-officedocument.wordprocessingml.document",
        // Progress
        function (index, total, caller) {

        },
        // success
        function (result, caller) {
            // Create a new Blob object using the 
            //response data of the onload object

            //Create a link element, hide it, direct 
            //it towards the blob, and then 'click' it programatically
            let a = document.createElement("a");
            a.style = "display: none";
            document.body.appendChild(a);
            //Create a DOMString representing the blob 
            //and point the link element towards it

            a.href = result;
            //programatically click the link to trigger the download
            a.click();

            //release the reference to the file by revoking the Object URL
            window.URL.revokeObjectURL(result);
        },
        // error
        function (errMsg, caller) {

        }, undefined)
}