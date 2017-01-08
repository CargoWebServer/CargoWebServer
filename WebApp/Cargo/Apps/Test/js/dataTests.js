var testsMap = {}
var testsReturnedResultMap = {}

function AllDataTestsReturnedResult_Test() {
    QUnit.test("AllDataTestsReturnedResult_Test",
        function (assert) {
            var noResultReturnedTests = ""
            for (var id in testsMap) {
                if (testsReturnedResultMap[id] == undefined) {
                    noResultReturnedTests += "\n" + id
                }
            }
            if (noResultReturnedTests.length > 0) {
                assert.ok(false, "-- DATA TESTS -- \nNot all tests have returned a result" + noResultReturnedTests);
            } else {
                assert.ok(true, "-- DATA TESTS -- \nAll tests have returned a result!");
            }
        })
}

//DataManager.prototype.createDataStore = function (storeId, storeType, storeVendor, successCallback, errorCallback, caller)
function createDataStore_Test() {
    testsMap["createDataStore_Test"] = createDataStore_Test
    server.dataManager.createDataStore(
        "createDataStore_Test",
        3,
        1,
        function (result) {
            QUnit.test("createDataStore_Test",
                function (result) {
                    return function (assert) {
                        assert.ok(true);
                        testsReturnedResultMap["createDataStore_Test"] = createDataStore_Test
                    }
                } (result))
        },
        function () {

        },
        undefined)
}

//** PAS FINI **//
function dataTests() {

    createDataStore_Test()

    setTimeout(function () {
        return function () {
            AllDataTestsReturnedResult_Test()
        }
    } (), 3000);
} 