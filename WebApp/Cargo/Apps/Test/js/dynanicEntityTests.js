/**
 * Create a new data schema name Test and class name Item and
 * also create instance of Item.
 */
function testDynamicEntity() {
    // Create the dataStore...
    server.dataManager.createDataStore("Test", 2, 1,
        function (result, caller) {
            // The prototype define a JSON object schema that hold 
            // information on the server side.
            // The information is indexed and can be queried  one the sever side 
            // via other libraries like  underscore.js or link.js as example.
            var itemPrototype = new EntityPrototype()
            itemPrototype.init(
                {
                    "TypeName": "Test.Item",
                    "Ids": ["M_id"],
                    "Indexs": ["M_name", "M_stringLst", "M_description", "M_price", "M_inStock", "M_qte", "M_date"], // put here field that you want to use as search key...
                    "Fields": ["M_id", "M_name", "M_description", "M_equivalent", "M_subitem", "M_stringLst", "M_creator", "M_price", "M_inStock", "M_qte", "M_date"],
                    "FieldsType": ["xs.string", "xs.string", "xs.string", "[]Test.Item:Ref", "[]Test.Item", "[]xs.string", "CargoEntities.User:Ref", "xs.double", "xs.boolean", "xs.int", "xs.dateTime"], // A reference to another item here..
                    "FieldsVisibility": [true, true, true, true, true, true, true, true, true, true, true],
                    "FieldsOrder": [0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10]
                }
            )
            // Create the prototype on the server...
            server.entityManager.createEntityPrototype(
                "Test",
                itemPrototype,
                function (result, caller) {
                    server.entityManager.getEntityPrototype("Test.Item", "Test",
                        function (result, caller) {
                            // Now I will create a new subItem type derived from the Test.Item type.
                            var derivedUserPrototype = new EntityPrototype()
                            derivedUserPrototype.init(
                                {
                                    "TypeName": "Test.User",
                                    "Ids": [],
                                    "Indexs": ["M_derivedName"], // put here field that you want to use as search key...
                                    "Fields": ["M_derivedName"],
                                    "FieldsType": ["xs.string"], // A reference to another item here..
                                    "FieldsVisibility": [true],
                                    "FieldsOrder": [0],
                                    "SuperTypeNames": ["CargoEntities.User"]
                                }
                            )

                            server.entityManager.createEntityPrototype(
                                "Test",
                                derivedUserPrototype,
                                function (result, caller) {
                                    server.entityManager.getEntityPrototype("Test.User", "Test",
                                        function (result, caller) {
                                            // i will create the derived item prototype now...
                                            var derivedItemPrototype = new EntityPrototype()
                                            derivedItemPrototype.init(
                                                {
                                                    "TypeName": "Test.DerivedItem",
                                                    "Ids": [],
                                                    "Indexs": ["M_derivedName"], // put here field that you want to use as search key...
                                                    "Fields": ["M_derivedName", "M_derivedUser"],
                                                    "FieldsType": ["xs.string", "Test.User:Ref"], // A reference to another item here..
                                                    "FieldsVisibility": [true, true],
                                                    "FieldsOrder": [0, 1],
                                                    "SuperTypeNames": ["Test.Item"]
                                                }
                                            )

                                            server.entityManager.createEntityPrototype(
                                                "Test",
                                                derivedItemPrototype,
                                                function (result, caller) {
                                                    server.entityManager.getEntityPrototype("Test.DerivedItem", "Test",
                                                        function (result, caller) {
                                                            // i will create the derived item prototype now...
                                                            testCreateDynamicEntity()
                                                        },
                                                        null)
                                                }, null)
                                        },
                                        null)
                                }, null)
                        },
                        null)
                }, null)
        },
        function (errMsg, caller) {

        }, undefined)


}

function testCreateDynamicEntity() {
    server.entityManager.getEntityPrototypes("Test",
        function (result, caller) {
            // After the prototype was created I can create instance of the new class...
            var instance_1 = new Test.Item()
            instance_1.M_id = "item_1"
            instance_1.M_name = "item 1"
            instance_1.M_stringLst = ["toto", "tata", "titi"]
            instance_1.M_description = "Ceci est une description de " + instance_1.M_name
            instance_1.M_price = 2.50
            instance_1.M_inStock = false
            instance_1.M_qte = 0
            instance_1.M_date = new Date("2016-12-12T15:42:22.720Z")

            var instance_2 = new Test.Item()
            instance_2.M_id = "item_2"
            instance_2.M_name = "item 2"
            instance_2.M_description = "Ceci est une description de " + instance_2.M_name
            instance_2.M_price = 2.25
            instance_2.M_inStock = true
            instance_2.M_qte = 10
            instance_2.M_date = new Date("2016-11-12T15:42:22.720Z")

            // Will be save as the same time as item 1
            var instance_3 = new Test.Item()
            instance_3.M_id = "item_3"
            instance_3.M_name = "item 3"
            instance_3.M_description = "Ceci est une description de " + instance_3.M_name
            instance_3.M_price = 2.50
            instance_3.M_inStock = true
            instance_3.M_qte = 12
            instance_3.M_date = new Date("2016-09-12T15:42:22.720Z")

            var instance_4 = new Test.Item()
            instance_4.M_id = "item_4"
            instance_4.M_name = "item 4"
            instance_4.M_stringLst = ["toto", "tata", "titi"]
            instance_4.M_description = "Ceci est une description de " + instance_4.M_name
            instance_4.M_price = 3.75
            instance_4.M_inStock = false
            instance_4.M_qte = 0
            instance_4.M_date = new Date("2016-08-12T15:42:22.720Z")

            var instance_5 = new Test.DerivedItem()
            instance_5.M_id = "item_5"
            instance_5.M_name = "item 5"
            instance_5.M_description = "Ceci est une description de " + instance_5.M_name
            instance_5.M_price = 4.75
            instance_5.M_inStock = true
            instance_5.M_qte = 12
            instance_5.M_date = new Date("2016-07-12T15:42:22.720Z")
            // only in DerivedItem
            instance_5.M_derivedName = "derived item 5 !-)"

            // Set subitem
            instance_3.setSubitem(instance_4)
            instance_3.setSubitem(instance_5)
            instance_1.setSubitem(instance_3)

            // simply save the first entity...
            server.entityManager.saveEntity(instance_1,
                function (result, caller) {
                })

            // Test save entity
            server.entityManager.saveEntity(instance_2,
                function (result, caller) {
                    // Set the equivalent... reference...
                    server.entityManager.getEntityByUuid(caller.instance_1_uuid,
                        function (instance_1, instance_2) {
                            instance_1.setEquivalent(instance_2)
                            instance_2.setEquivalent(instance_1)
                            // save it again with it references ...
                            server.entityManager.saveEntity(instance_2,
                                function (result, caller) {
                                    console.log("---> instance 2", result)
                                    server.entityManager.saveEntity(instance_1,
                                        function (instance1, caller) {
                                            console.log("---> instance 1", result)
                                            /*server.entityManager.getEntityById("CargoEntities.User", "mm006819@ud6.uf6", 
                                            function(result, instance1){
                                                instance1.setCreator(result)
                                               server.entityManager.saveEntity(instance_1)
                                            }, function(){
        
                                            }, instance1)*/
                                        },
                                        function () {

                                        })
                                },
                                function () { }, instance_1)

                        },
                        function () {

                        }, result)
                }, function () {

                }, { "instance_1_uuid": instance_1.UUID })

        },
        function () { },
        undefined)
}