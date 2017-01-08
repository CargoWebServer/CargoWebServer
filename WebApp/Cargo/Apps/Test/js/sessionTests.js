var nbTests = 0;
var nbTestsAsserted = 0;

function AllSessionTestsAsserted_Test() {
    QUnit.test("AllSessionTestsAsserted_Test",
        function (assert) {
            assert.ok(nbTests == nbTestsAsserted, "nbTests: " + nbTests + ", nbTestsAsserted: " + nbTestsAsserted);
        })
}


function login_Test() {
    nbTests++
    server.sessionManager.login("admin", "adminadmin",
        function (result) {
            QUnit.test("login_Test",
                function (result) {
                    return function (assert) {
                        assert.ok(result.M_accountPtr.M_name == "admin");
                        nbTestsAsserted++
                    }
                } (result))
        }, function (result) {

        },
        undefined)
}

function login_accountNameError_Test() {
    nbTests++
    server.sessionManager.login("aaa", "adminadmin",
        function (result) {

        }, function (result) {
            QUnit.test("login_accountNameError_Test",
                function (result) {
                    return function (assert) {
                        assert.ok(result.message != undefined);
                        nbTestsAsserted++
                    }
                } (result))
        },
        undefined)
}


function login_passwordError_Test() {
    nbTests++
    server.sessionManager.login("admin", "aaa",
        function (result) {

        }, function (result) {
            QUnit.test("login_passwordError_Test",
                function (result) {
                    return function (assert) {
                        assert.ok(result.message != undefined);
                        nbTestsAsserted++
                    }
                } (result))
        },
        undefined)
}


function getActiveSessions_Test() {
    nbTests++
    server.sessionManager.login("admin", "adminadmin",
        function () {
            server.sessionManager.getActiveSessions(
                function (result) {
                    QUnit.test("getActiveSessions_Test",
                        function (result) {
                            return function (assert) {
                                assert.ok(result[0].M_accountPtr == "admin");
                                nbTestsAsserted++
                            }
                        } (result))
                }, function (result) {

                },
                undefined)
        }, function () {

        },
        undefined)


}

function getActiveSessionByAccountId_Test() {
    nbTests++
    server.sessionManager.login("admin", "adminadmin",
        function () {
            server.sessionManager.getActiveSessions(
                function (result) {
                    QUnit.test("getActiveSessionByAccountId_Test",
                        function (result) {
                            return function (assert) {
                                assert.ok(result[0].M_accountPtr == "admin");
                                nbTestsAsserted++
                            }
                        } (result))
                }, function (result) {

                },
                undefined)
        }, function () {

        },
        undefined)
}

/* FAIL: Fonction updateSessionState retourne rien
    cause par GetActiveSessions
*/
function updateSessionState_1_Test() {
    nbTests++
    server.sessionManager.updateSessionState(1,
        function (result) {
            QUnit.test("updateSessionState_1_Test",
                function (result) {
                    return function (assert) {
                        console.log("updateSessionState_1_Test " + result)
                        assert.ok(false);
                        nbTestsAsserted++
                    }
                } (result))
        }, function (result) {

        },
        undefined)
}
/* FAIL: Fonction updateSessionState retourne rien
    cause par GetActiveSessions
*/
function updateSessionState_2_Test() {
    nbTests++
    server.sessionManager.updateSessionState(2,
        function (result) {
            QUnit.test("updateSessionState_2_Test",
                function (result) {
                    return function (assert) {
                        console.log("updateSessionState_2_Test " + result)
                        assert.ok(false);
                        nbTestsAsserted++
                    }
                } (result))
        }, function (result) {

        },
        undefined)
}
/* FAIL: Fonction updateSessionState retourne rien
    cause par GetActiveSessions
*/
function updateSessionState_3_Test() {
    nbTests++
    server.sessionManager.updateSessionState(3,
        function (result) {
            QUnit.test("updateSessionState_3_Test",
                function (result) {
                    return function (assert) {
                        console.log("updateSessionState_3_Test " + result)
                        assert.ok(false);
                        nbTestsAsserted++
                    }
                } (result))
        }, function (result) {

        },
        undefined)
}

/* PAS FINI
*/
function updateSessionState_Error_Test() {
    nbTests++
    server.sessionManager.updateSessionState(1,
        function (result) {

        }, function (result) {
            QUnit.test("updateSessionState_Error_Test",
                function (result) {
                    return function (assert) {
                        console.log("updateSessionState_Error_Test " + result)
                        assert.ok(false);
                        nbTestsAsserted++
                    }
                } (result))
        },
        undefined)
}

/* FAIL: passe dans le errorCallback
    cause par GetActiveSessions
*/
function logout_Test(state) {
    nbTests++
    server.sessionManager.logout(
        function (result) {
            QUnit.test("logout_Test",
                function (result) {
                    return function (assert) {
                        console.log("logout_Test" + result)
                        assert.ok(false);
                        nbTestsAsserted++
                    }
                } (result))
        }, function (result) {
        },
        undefined)
}


/* PAS FINI */
function logout_Error_Test(state) {
    nbTests++
    server.sessionManager.logout(
        function (result) {

        }, function (result) {
            QUnit.test("logout_Error_Test",
                function (result) {
                    return function (assert) {
                        assert.ok(result.message != undefined);
                        nbTestsAsserted++
                    }
                } (result))
        },
        undefined)
}



function sessionTests() {

    login_Test()
    login_accountNameError_Test()
    login_passwordError_Test()

    getActiveSessions_Test()
    getActiveSessionByAccountId_Test()
    /*
    updateSessionState_1_Test()
    updateSessionState_2_Test()
    updateSessionState_3_Test()
    updateSessionState_Error_Test()
    logout_Test()
    logout_Error_Test
    */




    setTimeout(function () {
        return function () {
            AllSessionTestsAsserted_Test()
        }
    } (), 3000);

    //getActiveSessionsTest();
    //getActiveSessionByAccountIdTest(/*accountId*/);
    //updateSessionStateTest(/*state*/);
    //logoutTest();
} 