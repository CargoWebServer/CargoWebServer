
var testsMap = {}
var testsReturnedResultMap = {}

function AllEntityTestsReturnedResult_Test() {
    QUnit.test("AllEntityTestsReturnedResult_Test",
        function (assert) {
            var noResultReturnedTests = ""
            for (var id in testsMap) {
                if (testsReturnedResultMap[id] == undefined) {
                    noResultReturnedTests += "\n" + id
                }
            }
            if (noResultReturnedTests.length > 0) {
                assert.ok(false, "-- ENTITY TESTS -- \nNot all tests have returned a result" + noResultReturnedTests);
            } else {
                assert.ok(true, "-- ENTITY TESTS -- \nAll tests have returned a result!");
            }
        })
}

// TODO Tester les caller au lieu de les mettre 'undefined' dans les parametres des fonctions
function entityTests() {

    server.entityManager.getEntityPrototypes(
        "CargoEntities",
        function () {
            /*
                        getObjectsByType_Test()
                        getObjectsByType_typeNameError_Test()
                        getObjectsByType_typeNameError_2_Test()
                        getObjectsByType_storeIdError_Test()
                        getObjectsByType_queryStrError_Test()
            
                        getEntityByUuid_Test()
                        getEntityByUuid_UuidError_Test()
            
                        getEntityById_Test()
                        getEntityById_typeNameError_Test()
                        getEntityById_idError_Test()
           */
            createEntity_Test()
            /*
           
            createEntity_parentUuidError_Test()
            createEntity_parentUuidError_2_Test()
            createEntity_attributeNameError_Test()
            createEntity_typeNameError_Test()
            createEntity_typeNameError_2_Test()

            removeEntity_Test()
            removeEntity_uuidError_Test()

            saveEntity_Test()
            saveEntity_2_Test()


            createEntityPrototype_Test()
            createEntityPrototype_storeIdError_Test()
            createEntityPrototype_prototypeError_Test()

            getEntityPrototype_Test()
            getEntityPrototype_typeNameError_Test()
            getEntityPrototype_storeIdError_Test()

            getDerivedEntityPrototypes_Test()
            getDerivedEntityPrototypes_typeNameError_Test()

            getEntityPrototypes_Test()
            getEntityPrototypes_storeIdError_Test()
             */

        }, function () {

        },
        undefined)

    setTimeout(function () {
        return function () {
            AllEntityTestsReturnedResult_Test()
        }
    } (), 3000);
}

//EntityManager.prototype.getObjectsByType = function (typeName, storeId, queryStr, lazy, level, progressCallback, successCallback, errorCallback, caller)
function getObjectsByType_Test() {
    testsMap["getObjectsByType_Test"] = getObjectsByType_Test
    server.entityManager.getObjectsByType(
        "CargoEntities.Log",
        "CargoEntities",
        undefined,
        function () {

        }, function (result) {
            QUnit.test("getObjectsByType_Test",
                function (result) {
                    return function (assert) {
                        assert.ok(result[0].TYPENAME == "CargoEntities.Log");
                        testsReturnedResultMap["getObjectsByType_Test"] = getObjectsByType_Test
                    }
                } (result))
        },
        function () {

        },
        undefined)
}

function getObjectsByType_typeNameError_Test() {
    testsMap["getObjectsByType_typeNameError_Test"] = getObjectsByType_typeNameError_Test
    server.entityManager.getObjectsByType(
        "AAA.Log",
        "CargoEntities",
        undefined,
        function () {
        }, function () {
        },
        function (result) {
            QUnit.test("getObjectsByType_typeNameError_Test",
                function (result) {
                    return function (assert) {
                        assert.ok(result.message != undefined, result.message);
                        testsReturnedResultMap["getObjectsByType_typeNameError_Test"] = getObjectsByType_typeNameError_Test
                    }
                } (result))
        },
        undefined)
}

function getObjectsByType_typeNameError_2_Test() {
    testsMap["getObjectsByType_typeNameError_2_Test"] = getObjectsByType_typeNameError_2_Test
    server.entityManager.getObjectsByType(
        "@ #$%52 3452 5235",
        "CargoEntities",
        undefined,
        function () {
        }, function () {
        },
        function (result) {
            QUnit.test("getObjectsByType_typeNameError_2_Test",
                function (result) {
                    return function (assert) {
                        assert.ok(result.message != undefined, result.message);
                        testsReturnedResultMap["getObjectsByType_typeNameError_2_Test"] = getObjectsByType_typeNameError_2_Test
                    }
                } (result))
        },
        undefined)
}

function getObjectsByType_storeIdError_Test() {
    testsMap["getObjectsByType_storeIdError_Test"] = getObjectsByType_storeIdError_Test
    server.entityManager.getObjectsByType(
        "CargoEntities.Log",
        "aaa",
        undefined,
        function () {
        },
        function () {
        },
        function (result) {
            QUnit.test("getObjectsByType_storeIdError_Test",
                function (result) {
                    return function (assert) {
                        assert.ok(result.message != undefined, result.message);
                        testsReturnedResultMap["getObjectsByType_storeIdError_Test"] = getObjectsByType_storeIdError_Test
                    }
                } (result))
        },
        undefined)
}

function getObjectsByType_queryStrError_Test() {
    testsMap["getObjectsByType_queryStrError_Test"] = getObjectsByType_queryStrError_Test
    server.entityManager.getObjectsByType(
        "CargoEntities.Log",
        "CargoEntities",
        function () {

        },
        function () {

        },
        function () {

        },
        function (result) {

            QUnit.test("getObjectsByType_queryStrError_Test",
                function (result) {
                    return function (assert) {
                        assert.ok(result.message != undefined, result.message);
                        testsReturnedResultMap["getObjectsByType_queryStrError_Test"] = getObjectsByType_queryStrError_Test
                    }
                } (result))
        },
        undefined)
}

// EntityManager.prototype.getEntityByUuid = function (uuid, lazy,level, successCallback, errorCallback, caller)
function getEntityByUuid_Test() {
    testsMap["getEntityByUuid_Test"] = getEntityByUuid_Test
    server.entityManager.getEntityById(
        "CargoEntities.Entities",
        "CARGO_ENTITIES",
        true,
        -1,
        function (result) {
            server.entityManager.getEntityByUuid(
                result.UUID,
                true,
                -1,
                function (result) {
                    QUnit.test("getEntityByUuid_Test",
                        function (result) {
                            return function (assert) {
                                assert.ok(result.M_id == "CARGO_ENTITIES");
                                testsReturnedResultMap["getEntityByUuid_Test"] = getEntityByUuid_Test
                            }
                        } (result))
                },
                function () {

                },
                undefined)
        }, function () {

        },
        undefined,
        undefined)
}

function getEntityByUuid_UuidError_Test() {
    testsMap["getEntityByUuid_UuidError_Test"] = getEntityByUuid_UuidError_Test

    server.entityManager.getEntityByUuid(
        "eee",
        true,
        -1,
        function () {

        },
        function (result) {
            QUnit.test("getEntityByUuid_UuidError_Test",
                function (result) {
                    return function (assert) {
                        assert.ok(result.message != undefined, result.message);
                        testsReturnedResultMap["getEntityByUuid_UuidError_Test"] = getEntityByUuid_UuidError_Test
                    }
                } (result))

        },
        undefined)

}


//EntityManager.prototype.getEntityById = function (typeName, id, lazy, level, successCallback, errorCallback, caller, parent)
function getEntityById_Test() {
    testsMap["getEntityById_Test"] = getEntityById_Test
    server.entityManager.getEntityById(
        "CargoEntities.Entities",
        "CARGO_ENTITIES",
        true,
        -1,
        function (result) {
            QUnit.test("getEntityById_Test",
                function (result) {
                    return function (assert) {
                        assert.ok(result.M_id == "CARGO_ENTITIES");
                        testsReturnedResultMap["getEntityById_Test"] = getEntityById_Test
                    }
                } (result))
        }, function () {

        },
        undefined,
        undefined)
}

function getEntityById_typeNameError_Test() {
    testsMap["getEntityById_typeNameError_Test"] = getEntityById_typeNameError_Test
    server.entityManager.getEntityById(
        "ccc",
        "CARGO_ENTITIES",
        true,
        -1,
        function () {

        }, function (result) {
            QUnit.test("getEntityById_typeNameError_Test",
                function (result) {
                    return function (assert) {
                        assert.ok(result.message != undefined, result.message);
                        testsReturnedResultMap["getEntityById_typeNameError_Test"] = getEntityById_typeNameError_Test
                    }
                } (result))
        },
        undefined,
        undefined)
}

function getEntityById_idError_Test() {
    testsMap["getEntityById_idError_Test"] = getEntityById_idError_Test
    server.entityManager.getEntityById(
        "CargoEntities.Entities",
        "ddd",
        true,
        -1,
        function () {

        }, function (result) {
            QUnit.test("getEntityById_idError_Test",
                function (result) {
                    return function (assert) {
                        assert.ok(result.message != undefined, result.message);
                        testsReturnedResultMap["getEntityById_idError_Test"] = getEntityById_idError_Test
                    }
                } (result))
        },
        undefined,
        undefined)
}

//
//EntityManager.prototype.createEntity = function (parentUuid, attributeName, typeName, id, entity, successCallback, errorCallback, caller)
function createEntity_Test() {
    testsMap["createEntity_Test"] = createEntity_Test
    server.entityManager.getEntityById(
        "CargoEntities.Entities",
        "CARGO_ENTITIES",
        true,
        -1,
        function (result) {
            server.entityManager.createEntity(
                result.UUID,
                "M_roles",
                "CargoEntities.Role",
                //TODO mettre un id aleatoire
                "createEntity_Test",
                new CargoEntities.Role(),
                function (result) {
                    QUnit.test("createEntity_Test",
                        function (result) {
                            return function (assert) {
                                assert.ok(result.M_id == "createEntity_Test");
                                testsReturnedResultMap["createEntity_Test"] = createEntity_Test
                            }
                        } (result))
                },
                function () {

                },
                undefined)
        }, function () {

        },
        undefined,
        undefined)
}

function createEntity_parentUuidError_Test() {
    testsMap["createEntity_parentUuidError_Test"] = createEntity_parentUuidError_Test
    server.entityManager.createEntity(
        "ggg",
        "M_roles",
        "CargoEntities.Role",
        "createEntity_parentUuidError_Test",
        new CargoEntities.Role(),
        function () {

        },
        function (result) {
            QUnit.test("createEntity_parentUuidError_Test",
                function (result) {
                    return function (assert) {
                        assert.ok(result.message != undefined, result.message);
                        testsReturnedResultMap["createEntity_parentUuidError_Test"] = createEntity_parentUuidError_Test
                    }
                } (result))
        },
        undefined)
}

function createEntity_parentUuidError_2_Test() {
    testsMap["createEntity_parentUuidError_2_Test"] = createEntity_parentUuidError_2_Test
    server.entityManager.createEntity(
        "!@#%^#$^$6 2452#$%@#@#!../ 23!@$",
        "M_roles",
        "CargoEntities.Role",
        "createEntity_parentUuidError_2_Test",
        new CargoEntities.Role(),
        function () {

        },
        function (result) {
            QUnit.test("createEntity_parentUuidError_2_Test",
                function (result) {
                    return function (assert) {
                        assert.ok(result.message != undefined, result.message);
                        testsReturnedResultMap["createEntity_parentUuidError_2_Test"] = createEntity_parentUuidError_2_Test
                    }
                } (result))
        },
        undefined)
}

function createEntity_attributeNameError_Test() {
    testsMap["createEntity_attributeNameError_Test"] = createEntity_attributeNameError_Test
    server.entityManager.getEntityById(
        "CargoEntities.Entities",
        "CARGO_ENTITIES",
        true,
        -1,
        function (result) {
            server.entityManager.createEntity(
                result.UUID,
                "hhh",
                "CargoEntities.Role",
                "TestRole",
                new CargoEntities.Role(),
                function () {

                },
                function (result) {
                    QUnit.test("createEntity_attributeNameError_Test",
                        function (result) {
                            return function (assert) {
                                assert.ok(result.message != undefined, result.message);
                                testsReturnedResultMap["createEntity_attributeNameError_Test"] = createEntity_attributeNameError_Test
                            }
                        } (result))

                },
                undefined)
        }, function () {

        },
        undefined,
        undefined)
}

function createEntity_typeNameError_Test() {
    testsMap["createEntity_typeNameError_Test"] = createEntity_typeNameError_Test
    server.entityManager.getEntityById(
        "CargoEntities.Entities",
        "CARGO_ENTITIES",
        true,
        -1,
        function (result) {
            server.entityManager.createEntity(
                result.UUID,
                "M_roles",
                "iii",
                "TestRole",
                new CargoEntities.Role(),
                function () {

                },
                function (result) {
                    QUnit.test("createEntity_typeNameError_Test",
                        function (result) {
                            return function (assert) {
                                assert.ok(result.message != undefined, result.message);
                                testsReturnedResultMap["createEntity_typeNameError_Test"] = createEntity_typeNameError_Test
                            }
                        } (result))

                },
                undefined)
        }, function () {

        },
        undefined,
        undefined)
}

function createEntity_typeNameError_2_Test() {
    testsMap["createEntity_typeNameError_2_Test"] = createEntity_typeNameError_2_Test
    server.entityManager.getEntityById(
        "CargoEntities.Entities",
        "CARGO_ENTITIES",
        true,
        -1,
        function (result) {
            server.entityManager.createEntity(
                result.UUID,
                "M_roles",
                "iii.aaa",
                "TestRole",
                new CargoEntities.Role(),
                function () {

                },
                function (result) {
                    QUnit.test("createEntity_typeNameError_2_Test",
                        function (result) {
                            return function (assert) {
                                assert.ok(result.message != undefined, result.message);
                                testsReturnedResultMap["createEntity_typeNameError_2_Test"] = createEntity_typeNameError_2_Test
                            }
                        } (result))

                },
                undefined)
        }, function () {

        },
        undefined,
        undefined)
}


// EntityManager.prototype.removeEntity = function (uuid, successCallback, errorCallback, caller)

function removeEntity_Test() {
    testsMap["removeEntity_Test"] = removeEntity_Test
    server.entityManager.getEntityById(
        "CargoEntities.Entities",
        "CARGO_ENTITIES",
        true,
        -1,
        function (result) {
            server.entityManager.createEntity(
                result.UUID,
                "M_roles",
                "CargoEntities.Role",
                "removeEntity_Test",
                new CargoEntities.Role(),
                function (result) {
                    server.entityManager.removeEntity(
                        result.UUID,
                        function (result) {
                            QUnit.test("removeEntity_Test",
                                function (result) {
                                    return function (assert) {
                                        assert.ok(result == true);
                                        testsReturnedResultMap["removeEntity_Test"] = removeEntity_Test
                                    }
                                } (result))
                        },
                        function () {

                        },
                        undefined)
                },
                function () {

                },
                undefined)
        }, function () {

        },
        undefined,
        undefined)
}

function removeEntity_uuidError_Test() {
    testsMap["removeEntity_uuidError_Test"] = removeEntity_uuidError_Test
    server.entityManager.getEntityById(
        "CargoEntities.Entities",
        "CARGO_ENTITIES",
        true,
        -1,
        function (result) {
            server.entityManager.createEntity(
                result.UUID,
                "M_roles",
                "CargoEntities.Role",
                "removeEntity_uuidError_Test",
                new CargoEntities.Role(),
                function (result) {
                    server.entityManager.removeEntity(
                        "lll",
                        function () {

                        },
                        function (result) {
                            QUnit.test("removeEntity_uuidError_Test",
                                function (result) {
                                    return function (assert) {
                                        assert.ok(result.message != undefined, result.message);
                                        testsReturnedResultMap["removeEntity_uuidError_Test"] = removeEntity_uuidError_Test
                                    }
                                } (result))
                        },
                        undefined)
                },
                function () {

                },
                undefined)
        }, function () {

        },
        undefined,
        undefined)
}

// EntityManager.prototype.saveEntity = function (entity, successCallback, errorCallback, caller) 

function saveEntity_Test() {
    testsMap["saveEntity_Test"] = saveEntity_Test
    server.entityManager.getEntityById(
        "CargoEntities.Entities",
        "CARGO_ENTITIES",
        true,
        -1,
        function (result) {
            server.entityManager.createEntity(
                result.UUID,
                "M_roles",
                "CargoEntities.Role",
                "saveEntity_Test",
                new CargoEntities.Role(),
                function (result) {
                    server.entityManager.saveEntity(
                        result,
                        function (result) {
                            QUnit.test("saveEntity_Test",
                                function (result) {
                                    return function (assert) {
                                        assert.ok(result.M_id == "saveEntity_Test");
                                        testsReturnedResultMap["saveEntity_Test"] = saveEntity_Test
                                    }
                                } (result))
                        },
                        function () {

                        },
                        undefined)
                },
                function () {

                },
                undefined)
        }, function () {

        },
        undefined,
        undefined)
}


function saveEntity_2_Test() {
    testsMap["saveEntity_2_Test"] = saveEntity_2_Test

    server.entityManager.saveEntity(
        new CargoEntities.Role(),
        function (result) {
            QUnit.test("saveEntity_2_Test",
                function (result) {
                    return function (assert) {
                        assert.ok(result.TYPENAME == "CargoEntities.Role");
                        testsReturnedResultMap["saveEntity_2_Test"] = saveEntity_2_Test
                    }
                } (result))
        },
        function () {

        },
        undefined)
}

// EntityManager.prototype.createEntityPrototype = function (storeId, prototype, successCallback, errorCallback, caller)

function createEntityPrototype_Test() {
    testsMap["createEntityPrototype_Test"] = createEntityPrototype_Test

    server.dataManager.deleteDataStore(
        "createEntityPrototype_Test",
        function () {
            createEntityPrototype_Test_Part_2()
        },
        function () {
            createEntityPrototype_Test_Part_2()
        },
        undefined
    )




}

function createEntityPrototype_Test_Part_2() {
    server.dataManager.createDataStore(
        "createEntityPrototype_Test",
        3,
        1,
        function () {

            var itemPrototype = new EntityPrototype()
            itemPrototype.init(
                {
                    "TypeName": "createEntityPrototype_Test.Item",
                    "Ids": ["M_id"],
                    "Indexs": ["M_name"], // put here field that you want to use as search key...
                    "Fields": ["M_id", "M_name", "M_description", "M_equivalent", "M_subitem"],
                    "FieldsType": ["xs.string", "xs.string", "xs.string", "[]createEntityPrototype_Test.Item:Ref", "[]createEntityPrototype_Test.Item"], // A reference to another item here..
                    "FieldsVisibility": [true, true, true, true, true],
                    "FieldsOrder": [0, 1, 2, 3, 4]
                }
            )


            server.entityManager.createEntityPrototype(
                "createEntityPrototype_Test",
                itemPrototype,
                function (result) {


                    QUnit.test("createEntityPrototype_Test",
                        function (result) {
                            return function (assert) {
                                assert.ok(result.TypeName == "createEntityPrototype_Test.Item");
                                testsReturnedResultMap["createEntityPrototype_Test"] = createEntityPrototype_Test
                            }
                        } (result))
                },
                function () {

                },
                undefined)
        },
        function () {

        },
        undefined)
}

function createEntityPrototype_storeIdError_Test() {
    server.dataManager.createDataStore(
        "createEntityPrototype_storeIdError_Test",
        3,
        1,
        function () {

            var itemPrototype = new EntityPrototype()
            itemPrototype.init(
                {
                    "TypeName": "createEntityPrototype_storeIdError_Test.Item",
                    "Ids": ["M_id"],
                    "Indexs": ["M_name"], // put here field that you want to use as search key...
                    "Fields": ["M_id", "M_name", "M_description", "M_equivalent", "M_subitem"],
                    "FieldsType": ["xs.string", "xs.string", "xs.string", "[]createEntityPrototype_storeIdError_Test.Item:Ref", "[]createEntityPrototype_storeIdError_Test.Item"], // A reference to another item here..
                    "FieldsVisibility": [true, true, true, true, true],
                    "FieldsOrder": [0, 1, 2, 3, 4]
                }
            )


            server.entityManager.createEntityPrototype(
                "rrr",
                itemPrototype,
                function () {

                },
                function (result) {

                    QUnit.test("createEntityPrototype_storeIdError_Test",
                        function (result) {
                            return function (assert) {
                                assert.ok(result.message != undefined);
                                testsReturnedResultMap["createEntityPrototype_storeIdError_Test"] = createEntityPrototype_storeIdError_Test
                            }
                        } (result))

                },
                undefined)
        },
        function () {

        },
        undefined)
}



function createEntityPrototype_prototypeError_Test() {
    testsMap["createEntityPrototype_prototypeError_Test"] = createEntityPrototype_prototypeError_Test

    server.entityManager.createEntityPrototype(
        CargoEntities,
        new CargoEntities.Role(),
        function () {

        },
        function (result) {
            QUnit.test("createEntityPrototype_prototypeError_Test",
                function (result) {
                    return function (assert) {
                        console.log(result)

                        assert.ok(result.message != undefined);
                        testsReturnedResultMap["createEntityPrototype_prototypeError_Test"] = createEntityPrototype_prototypeError_Test
                    }
                } (result))
        },
        undefined)
}

// EntityManager.prototype.getEntityPrototype = function (typeName, storeId, successCallback, errorCallback, caller)
function getEntityPrototype_Test() {
    testsMap["getEntityPrototype_Test"] = getEntityPrototype_Test

    server.entityManager.getEntityPrototype(
        "CargoEntities.Log",
        "CargoEntities",
        function (result) {
            QUnit.test("getEntityPrototype_Test",
                function (result) {
                    return function (assert) {
                        assert.ok(result.TypeName == "CargoEntities.Log");
                        testsReturnedResultMap["getEntityPrototype_Test"] = getEntityPrototype_Test
                    }
                } (result))
        },
        function () {

        },
        undefined)
}

function getEntityPrototype_typeNameError_Test() {
    testsMap["getEntityPrototype_typeNameError_Test"] = getEntityPrototype_typeNameError_Test

    server.entityManager.getEntityPrototype(
        "mmm",
        "CargoEntities",
        function () {

        },
        function (result) {
            QUnit.test("getEntityPrototype_typeNameError_Test",
                function (result) {
                    return function (assert) {
                        assert.ok(result.message != undefined, result.message);
                        testsReturnedResultMap["getEntityPrototype_typeNameError_Test"] = getEntityPrototype_typeNameError_Test
                    }
                } (result))
        },
        undefined)
}

function getEntityPrototype_storeIdError_Test() {
    testsMap["getEntityPrototype_storeIdError_Test"] = getEntityPrototype_storeIdError_Test

    server.entityManager.getEntityPrototype(
        "nnn.Log",
        "nnn",
        function (result) {
            console.log(result)
        },
        function (result) {
            QUnit.test("getEntityPrototype_storeIdError_Test",
                function (result) {
                    return function (assert) {
                        assert.ok(result.message != undefined, result.message);
                        testsReturnedResultMap["getEntityPrototype_storeIdError_Test"] = getEntityPrototype_storeIdError_Test
                    }
                } (result))
        },
        undefined)
}

// EntityManager.prototype.getDerivedEntityPrototypes = function (typeName, successCallback, errorCallback, caller)
function getDerivedEntityPrototypes_Test() {
    testsMap["getDerivedEntityPrototypes_Test"] = getDerivedEntityPrototypes_Test

    server.entityManager.getDerivedEntityPrototypes(
        "CargoEntities.Entity",
        function (result) {
            QUnit.test("getDerivedEntityPrototypes_Test",
                function (result) {
                    return function (assert) {
                        var ok = result.length > 0
                        for (var i = 0; i < result.length; i++) {
                            if (!contains(result[i].SuperTypeNames, "CargoEntities.Entity")) {
                                ok = false
                            }
                        }
                        assert.ok(ok);
                        testsReturnedResultMap["getDerivedEntityPrototypes_Test"] = getDerivedEntityPrototypes_Test
                    }
                } (result))
        },
        function () {

        },
        undefined)
}

function getDerivedEntityPrototypes_typeNameError_Test() {
    testsMap["getDerivedEntityPrototypes_typeNameError_Test"] = getDerivedEntityPrototypes_typeNameError_Test

    server.entityManager.getDerivedEntityPrototypes(
        "ooo",
        function () {

        },
        function (result) {
            QUnit.test("getDerivedEntityPrototypes_typeNameError_Test",
                function (result) {
                    return function (assert) {
                        assert.ok(result.message != undefined, result.message);
                        testsReturnedResultMap["getDerivedEntityPrototypes_typeNameError_Test"] = getDerivedEntityPrototypes_typeNameError_Test
                    }
                } (result))
        },
        undefined)
}

// EntityManager.prototype.getEntityPrototypes = function (storeId, successCallback, errorCallback, caller) 
function getEntityPrototypes_Test() {
    testsMap["getEntityPrototypes_Test"] = getEntityPrototypes_Test

    server.entityManager.getEntityPrototypes(
        "CargoEntities",
        function (result) {
            QUnit.test("getEntityPrototypes_Test",
                function (result) {
                    return function (assert) {
                        assert.ok(result.length > 0);
                        testsReturnedResultMap["getEntityPrototypes_Test"] = getEntityPrototypes_Test
                    }
                } (result))
        },
        function () {

        },
        undefined)
}

function getEntityPrototypes_storeIdError_Test() {
    testsMap["getEntityPrototypes_storeIdError_Test"] = getEntityPrototypes_storeIdError_Test

    server.entityManager.getEntityPrototypes(
        "ppp",
        function () {

        },
        function (result) {
            QUnit.test("getEntityPrototypes_storeIdError_Test",
                function (result) {
                    return function (assert) {
                        assert.ok(result.message != undefined, result.message);
                        testsReturnedResultMap["getEntityPrototypes_storeIdError_Test"] = getEntityPrototypes_storeIdError_Test
                    }
                } (result))
        },
        undefined)
}

