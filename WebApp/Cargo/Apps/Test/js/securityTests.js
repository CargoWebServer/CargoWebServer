
var nbTests = 0;
var nbTestsAsserted = 0;

function AllSecurityTestsAsserted_Test() {
    QUnit.test("AllSecurityTestsAsserted_Test",
        function (assert) {
            assert.ok(nbTests == nbTestsAsserted, "nbTests: " + nbTests + ", nbTestsAsserted: " + nbTestsAsserted);
        })
}

function createRole_Test(roleId) {
    nbTests++
    server.securityManager.createRole(roleId,
        function (result, caller) {
            QUnit.test("createRole_Test",
                function (result, roleId) {
                    return function (assert) {
                        assert.ok(result.M_id == roleId);
                        nbTestsAsserted++
                    }
                } (result, caller))
        }, function (result) {
        },
        roleId)
}

function createRole_Error_Test(roleId) {
    nbTests++
    server.securityManager.createRole(roleId,
        function () {

        }, function (result, caller) {
            QUnit.test("createRole_Error_Test",
                function (result, roleId) {
                    return function (assert) {
                        assert.ok(result.message != undefined);
                        nbTestsAsserted++
                    }
                } (result, caller))
        },
        roleId)
}


function getRole_Test(roleId) {
    nbTests++
    server.securityManager.getRole(roleId,
        function (result, caller) {
            QUnit.test("getRole_Test",
                function (result, roleId) {
                    return function (assert) {
                        assert.ok(result.M_id == roleId);
                        nbTestsAsserted++
                    }
                } (result, caller))
        }, function () {

        },
        roleId)
}

function getRole_Error_Test(roleId) {
    nbTests++
    server.securityManager.getRole(roleId,
        function () {
        }, function (result) {
            QUnit.test("getRole_Error_Test",
                function (result) {
                    return function (assert) {
                        assert.ok(result.message != undefined);
                        nbTestsAsserted++
                    }
                } (result))
        },
        roleId)
}

function appendAccount_Test(roleId, accountId) {
    nbTests++
    server.securityManager.appendAccount(roleId, accountId,
        function (result, caller) {
            var roleId = caller.roleId
            var accountId = caller.accountId
            server.securityManager.hasAccount(roleId, accountId,
                function (result, caller) {
                    QUnit.test("appendAccount_Test",
                        function (result, caller) {
                            return function (assert) {
                                assert.ok(true);
                                nbTestsAsserted++
                            }
                        } (result, caller))
                }, function (result) {
                },
                undefined)
        }, function (result) {
        },
        { "roleId": roleId, "accountId": accountId })
}

function appendAccount_Error_Test(roleId, accountId) {
    nbTests++
    server.securityManager.appendAccount(roleId, accountId,
        function () {

        }, function (result, caller) {
            QUnit.test("appendAccount_Error_Test",
                function (result, roleId) {
                    return function (assert) {
                        assert.ok(result.message != undefined);
                        nbTestsAsserted++
                    }
                } (result, caller))
        },
        { "roleId": roleId, "accountId": accountId })
}

function hasAccount_Test(roleId, accountId) {
    nbTests++
    server.securityManager.hasAccount(roleId, accountId,
        function (result, caller) {
            QUnit.test("hasAccount_Test",
                function (result, caller) {
                    return function (assert) {
                        assert.ok(result);
                        nbTestsAsserted++
                    }
                } (result, caller))
        }, function () {
        },
        [roleId, accountId])
}

function hasAccount_False_Test(roleId, accountId) {
    nbTests++
    server.securityManager.hasAccount(roleId, accountId,
        function (result, caller) {
            QUnit.test("hasAccount_False_Test",
                function (result, caller) {
                    return function (assert) {
                        assert.ok(!result);
                        nbTestsAsserted++
                    }
                } (result, caller))
        }, function () {
        },
        [roleId, accountId])
}


function removeAccount_Test(roleId, accountId) {
    nbTests++
    server.securityManager.removeAccount(roleId, accountId,
        function (result, caller) {
            var roleId = caller.roleId
            var accountId = caller.accountId
            server.securityManager.hasAccount(roleId, accountId,
                function (result, caller) {
                    QUnit.test("removeAccount_Test",
                        function (result, caller) {
                            return function (assert) {
                                assert.ok(!result);
                                nbTestsAsserted++
                            }
                        } (result, caller))
                }, function (result) {
                },
                undefined)
        }, function (result) {
        },
        { "roleId": roleId, "accountId": accountId })
}


function removeAccount_Error_Test(roleId, accountId) {
    nbTests++
    server.securityManager.removeAccount(roleId, accountId,
        function () {
        }, function (result, caller) {

            QUnit.test("removeAccount_Error_Test",
                function (result, roleId) {
                    return function (assert) {
                        assert.ok(result.message != undefined);
                        nbTestsAsserted++
                    }
                } (result, caller))
        },
        { "roleId": roleId, "accountId": accountId })
}


function deleteRole_Test(roleId) {
    nbTests++
    server.securityManager.deleteRole(roleId,
        function (result, caller) {
            server.securityManager.getRole(roleId,
                function (result, caller) {

                }, function () {
                    QUnit.test("deleteRole_Test",
                        function (result, roleId) {
                            return function (assert) {
                                assert.ok(true);
                                nbTestsAsserted++
                            }
                        } (result, caller))
                },
                roleId)
        }, function () {
        },
        roleId)
}

function deleteRole_Error_Test(roleId) {
    nbTests++
    server.securityManager.deleteRole(roleId,
        function () {
        }, function (result) {
            QUnit.test("deleteRole_Error_Test",
                function (result) {
                    return function (assert) {
                        assert.ok(result.message != undefined);
                        nbTestsAsserted++
                    }
                } (result))
        },
        roleId)
}


function setPermissionRole_Test(roleId) {
    nbTests++
    server.securityManager.createRole(roleId,
        function (result, caller) {
            server.securityManager.setPermissionRole(roleId, result.UUID, 1,
                function () {
                    QUnit.test("setPermissionRole_Test",
                        function (result) {
                            return function (assert) {
                                assert.ok(true);
                                nbTestsAsserted++
                            }
                        } (result))
                }, function () {

                }, undefined)
        }, function (result) {
        },
        roleId)
}

function setPermissionRole_roleIdError_Test(roleId) {
    nbTests++
    server.securityManager.createRole(roleId,
        function (result, caller) {

            server.securityManager.setPermissionRole("aaa", result.UUID, 1,
                function () {

                }, function (result) {
                    QUnit.test("setPermissionRole_roleIdError_Test",
                        function (result) {
                            return function (assert) {
                                assert.ok(result.message != undefined);
                                nbTestsAsserted++
                            }
                        } (result))
                }, undefined)

        }, function () {
        },
        roleId)
}


function setPermissionRole_entityUuidError_Test(roleId) {
    nbTests++
    server.securityManager.createRole(roleId,
        function (result, caller) {

            server.securityManager.setPermissionRole(roleId, "aaa", 1,
                function () {

                }, function (result) {
                    QUnit.test("setPermissionRole_entityUuidError_Test",
                        function (result) {
                            return function (assert) {
                                assert.ok(result.message != undefined);
                                nbTestsAsserted++
                            }
                        } (result))
                }, undefined)

        }, function () {
        },
        roleId)
}


function removePermissionRole_Test(roleId) {
    nbTests++
    server.securityManager.createRole(roleId,
        function (result, caller) {
            server.securityManager.setPermissionRole(roleId, result.UUID, 1,
                function (result, resultUUID) {
                    server.securityManager.removePermissionRole(roleId, resultUUID, 1,
                        function () {
                            QUnit.test("removePermissionRole_Test",
                                function (assert) {
                                    assert.ok(true);
                                    nbTestsAsserted++
                                })

                        }, function () {

                        }, undefined)
                }, function () {
                }, result.UUID)
        }, function (result) {
        },
        roleId)

}

function removePermissionRole_roleIdError_Test(roleId) {
    nbTests++
    server.securityManager.createRole(roleId,
        function (result, caller) {
            server.securityManager.setPermissionRole(roleId, result.UUID, 1,
                function (result, resultUUID) {
                    server.securityManager.removePermissionRole("aaa", resultUUID, 1,
                        function () {

                        }, function () {
                            QUnit.test("removePermissionRole_roleIdError_Test",
                                function (result) {
                                    return function (assert) {
                                        assert.ok(result.message != undefined);
                                        nbTestsAsserted++
                                    }
                                } (result))

                        }, undefined)
                }, function () {
                }, result.UUID)
        }, function (result) {
        },
        roleId)

}

function removePermissionRole_entityUuidError_Test(roleId) {
    nbTests++
    server.securityManager.createRole(roleId,
        function (result, caller) {
            server.securityManager.setPermissionRole(roleId, result.UUID, 1,
                function (result, resultUUID) {
                    server.securityManager.removePermissionRole(roleId, "aaa", 1,
                        function () {

                        }, function () {
                            QUnit.test("removePermissionRole_entityUuidError_Test",
                                function (result) {
                                    return function (assert) {
                                        assert.ok(result.message != undefined);
                                        nbTestsAsserted++
                                    }
                                } (result))

                        }, undefined)
                }, function () {
                }, result.UUID)
        }, function (result) {
        },
        roleId)
}

function getPermissionsByRole_Test(roleId) {
    nbTests++
    server.securityManager.createRole(roleId,
        function (result, caller) {
            server.securityManager.setPermissionRole(roleId, result.UUID, 1,
                function (result, resultUUID) {
                    server.securityManager.getPermissionsByRole(roleId, resultUUID,
                        function () {

                            QUnit.test("getPermissionsByRole_Test",
                                function (assert) {
                                    assert.ok(true);
                                    nbTestsAsserted++
                                })

                        }, function () {
                        }, undefined)
                }, function () {
                }, result.UUID)
        }, function (result) {
        },
        roleId)
}

function getPermissionsByRole_roleIdError_Test(roleId) {
    nbTests++
    server.securityManager.createRole(roleId,
        function (result, caller) {
            server.securityManager.setPermissionRole(roleId, result.UUID, 1,
                function (result, resultUUID) {
                    server.securityManager.getPermissionsByRole("aaa", resultUUID, 1,
                        function () {

                        }, function () {
                            QUnit.test("getPermissionsByRole_roleIdError_Test",
                                function (result) {
                                    return function (assert) {
                                        assert.ok(result.message != undefined);
                                        nbTestsAsserted++
                                    }
                                } (result))

                        }, undefined)
                }, function () {
                }, result.UUID)
        }, function (result) {
        },
        roleId)
}

function getPermissionsByRole_entityUuidError_Test(roleId) {
    nbTests++
    server.securityManager.createRole(roleId,
        function (result, caller) {
            server.securityManager.setPermissionRole(roleId, result.UUID, 1,
                function (result, resultUUID) {
                    server.securityManager.getPermissionsByRole(roleId, "aaa", 1,
                        function () {

                        }, function () {
                            QUnit.test("getPermissionsByRole_entityUuidError_Test",
                                function (result) {
                                    return function (assert) {
                                        assert.ok(result.message != undefined);
                                        nbTestsAsserted++
                                    }
                                } (result))

                        }, undefined)
                }, function () {
                }, result.UUID)
        }, function (result) {
        },
        roleId)
}

function getPermissions_Test(roleId) {
    nbTests++
    server.securityManager.createRole(roleId,
        function (result, caller) {
            server.securityManager.setPermissionRole(roleId, result.UUID, 1,
                function (result, resultUUID) {
                    server.securityManager.getPermissions(resultUUID,
                        function () {

                            QUnit.test("getPermissions_Test",
                                function (assert) {
                                    assert.ok(true);
                                    nbTestsAsserted++
                                })

                        }, function () {
                        }, undefined)
                }, function () {
                }, result.UUID)
        }, function (result) {
        },
        roleId)

}

function getPermissions_entityUuidError_Test(roleId) {
    nbTests++
    server.securityManager.createRole(roleId,
        function (result, caller) {
            server.securityManager.setPermissionRole(roleId, result.UUID, 1,
                function (result, resultUUID) {
                    server.securityManager.getPermissions("aaa",
                        function () {

                        }, function () {
                            QUnit.test("getPermissions_entityUuidError_Test",
                                function (result) {
                                    return function (assert) {
                                        assert.ok(result.message != undefined);
                                        nbTestsAsserted++
                                    }
                                } (result))

                        }, undefined)
                }, function () {
                }, result.UUID)
        }, function (result) {
        },
        roleId)
}


function securityTests() {

    // Roles related tests

    var roleId = randomUUID()
    createRole_Test(roleId)
    createRole_Error_Test("admin")
    setTimeout(function (roleId) {
        return function () {
            getRole_Test(roleId)
            getRole_Error_Test("aaa")
            setTimeout(function (roleId) {
                return function () {
                    appendAccount_Test(roleId, "admin")
                    appendAccount_Error_Test(roleId, "aaa")
                    appendAccount_Error_Test("aaa", "admin")
                    setTimeout(function (roleId) {
                        return function () {
                            hasAccount_Test(roleId, "admin")
                            hasAccount_False_Test(roleId, "aaa")
                            setTimeout(function (roleId) {
                                return function () {
                                    removeAccount_Test(roleId, "admin")
                                    removeAccount_Error_Test(roleId, "aaa")
                                    removeAccount_Error_Test("aaa", "admin")
                                    setTimeout(function (roleId) {
                                        return function () {
                                            deleteRole_Test(roleId)
                                            setTimeout(function (roleId) {
                                                return function () {
                                                    deleteRole_Error_Test(roleId)
                                                }
                                            } (roleId), 100);
                                        }
                                    } (roleId), 100);
                                }
                            } (roleId), 100);
                        }
                    } (roleId), 100);
                }
            } (roleId), 100);
        }
    } (roleId), 100);


    //Permissions related tests

    setPermissionRole_Test(randomUUID())
    setPermissionRole_roleIdError_Test(randomUUID())
    setPermissionRole_entityUuidError_Test(randomUUID())

    removePermissionRole_Test(randomUUID())
    removePermissionRole_roleIdError_Test(randomUUID())
    removePermissionRole_entityUuidError_Test(randomUUID())

    getPermissionsByRole_Test(randomUUID())
    getPermissionsByRole_roleIdError_Test(randomUUID())
    getPermissionsByRole_entityUuidError_Test(randomUUID())

    getPermissions_Test(randomUUID())
    getPermissions_entityUuidError_Test(randomUUID())

    setTimeout(function () {
        return function () {
            AllSecurityTestsAsserted_Test()
        }
    } (), 3000);


} 
