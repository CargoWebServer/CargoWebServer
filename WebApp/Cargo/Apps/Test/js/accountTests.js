
var nbTests = 0;
var nbTestsAsserted = 0;

function makeRandom5characterString() {
    var text = "";
    var possible = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789";

    for (var i = 0; i < 5; i++)
        text += possible.charAt(Math.floor(Math.random() * possible.length));

    return text;
}

function AllAccountTestsAsserted_Test() {
    QUnit.test("AllAccountTestsAsserted_Test",
        function (assert) {
            assert.ok(nbTests == nbTestsAsserted, "nbTests: " + nbTests + ", nbTestsAsserted: " + nbTestsAsserted);
        })
}


function register_Test() {
    nbTests++
    server.accountManager.register(makeRandom5characterString(), "aaa", "a@hotmail.com",
        function (result) {
            QUnit.test("register_Test",
                function (result) {
                    return function (assert) {
                        assert.ok(result.M_email == "a@hotmail.com");
                        nbTestsAsserted++
                    }
                } (result))
        }, function (result) {
        },
        undefined)
}

function register_AccountAlreadyExistsError_Test() {
    nbTests++
    server.accountManager.register("admin", "aaa", "a@hotmail.com",
        function (result) {

        }, function (result) {
            QUnit.test("register_AccountAlreadyExistsError_Test",
                function (result) {
                    return function (assert) {
                        assert.ok(result.message != undefined);
                        nbTestsAsserted++
                    }
                } (result))
        },
        undefined)
}

/* FAIL: email 'aaa' devrait pas etre accepte
accountManager.go ln48 : isEmailValid := GetServer().GetSmtpManager().ValidateEmail(email)
*/
function register_InvalidEmailError_Test() {
    nbTests++
    server.accountManager.register(makeRandom5characterString(), "aaa", "aaa",
        function (result) {

        }, function (result) {
            QUnit.test("register_InvalidEmailError_Test",
                function (result) {
                    return function (assert) {
                        assert.ok(result.message != undefined);
                        nbTestsAsserted++
                    }
                } (result))
        },
        undefined)
}

function getAccountById_Test() {
    nbTests++
    server.accountManager.getAccountById("admin",
        function (result) {
            QUnit.test("getAccountById_Test",
                function (result) {
                    return function (assert) {
                        assert.ok(result.M_id == "admin");
                        nbTestsAsserted++
                    }
                } (result))
        }, function (result) {
        },
        undefined)
}

function getAccountById_AccountDoesntExistError_Test() {
    nbTests++
    server.accountManager.getAccountById("bbb",
        function (result) {
        }, function (result) {
            QUnit.test("getAccountById_AccountDoesntExistError_Test",
                function (result) {
                    return function (assert) {
                        assert.ok(result.message != undefined);
                        nbTestsAsserted++
                    }
                } (result))
        },
        undefined)
}

/* FAIL: Il n'y a pas de admin user par defaut 
    Question: Est que admin account devrait etre creer dans accountmanager initialize() avec adminUser au lieu de dans securityManager initialize()
*/ 
function getUserById_Test() {
    nbTests++
    server.accountManager.getUserById("admin",
        function (result) {
            QUnit.test("getUserById_Test",
                function (result) {
                    return function (assert) {
                        assert.ok(result.M_id == "admin");
                        nbTestsAsserted++
                    }
                } (result))
        }, function (result) {
        },
        undefined)
}


function getUserById_UserIdDoesntExistError_Test() {
    nbTests++
    server.accountManager.getUserById("bbb",
        function (result) {
        }, function (result) {
            QUnit.test("getUserById_UserIdDoesntExistError_Test",
                function (result) {
                    return function (assert) {
                        assert.ok(result.message != undefined);
                        nbTestsAsserted++
                    }
                } (result))
        },
        undefined)
}

function accountTests() {
    register_Test()
    register_AccountAlreadyExistsError_Test()
    register_InvalidEmailError_Test()
    getAccountById_Test()
    getAccountById_AccountDoesntExistError_Test()
    getUserById_Test()
    getUserById_UserIdDoesntExistError_Test()

    setTimeout(function () {
        return function () {
            AllAccountTestsAsserted_Test()
        }
    } (), 3000);
} 