var languageInfo = {
    "en": {
        "LanguageManagerTest": "LanguageManagerTest - ENGLISH"
    },
    "fr": {
        "LanguageManagerTest": "LanguageManagerTest - FRANCAIS"
    },
}

function LanguageManagerTest() {
    var i = 0
    for (var languageInfoId in LanguageManager().languageInfo) {
        i = LanguageManager().languageInfo[languageInfoId]
    }
    QUnit.test("LanguageManagerTest", function (assert) {
        assert.ok(i == 0);
    })
}

function LanguageManager_prototype_appendLanguageInfoTest() {
    var languageInfo2 = {
        "en": {
            "LanguageManagerTest2": "LanguageManagerTest2 - ENGLISH"
        },
        "fr": {
            "LanguageManagerTest2": "LanguageManagerTest2 - FRANCAIS"
        },
    }
    server.languageManager.appendLanguageInfo(languageInfo2)

    QUnit.test("LanguageManager_prototype_appendLanguageInfoTest", function (assert) {
        assert.ok(server.languageManager.languageInfo["en"]["LanguageManagerTest2"]
            == "LanguageManagerTest2 - ENGLISH");
    })
}

function LanguageManager_prototype_setElementTextTest() {
    var element1 = local_createNewDivElement()
    server.languageManager.setElementText(element1, "LanguageManagerTest")

    QUnit.test("LanguageManager_prototype_setElementTextTest", function (assert) {
        assert.ok(element1.element.innerHTML == "LanguageManagerTest - ENGLISH");
    })
}

function LanguageManager_prototype_setLanguageTest() {
    var element1 = local_createNewDivElement()
    server.languageManager.setElementText(element1, "LanguageManagerTest")

    QUnit.test("LanguageManager_prototype_setLanguageTest", function (assert) {
        server.languageManager.setLanguage("en")
        assert.ok(element1.element.innerHTML == "LanguageManagerTest - ENGLISH");
        server.languageManager.setLanguage("fr")
        assert.ok(element1.element.innerHTML == "LanguageManagerTest - FRANCAIS");
    })
}

function LanguageManager_prototype_getTextTest() {
    QUnit.test("LanguageManager_prototype_getTextTest", function (assert) {
        server.languageManager.setLanguage("en")
        assert.ok(server.languageManager.getText(["LanguageManagerTest"]) == "LanguageManagerTest - ENGLISH");
    })
}

function languageManagerTests() {
    /* Setup */
    server.languageManager.appendLanguageInfo(languageInfo)
    server.languageManager.setLanguage("en")

    /* Tests */
    LanguageManagerTest()
    LanguageManager_prototype_appendLanguageInfoTest()
    LanguageManager_prototype_setElementTextTest()
    LanguageManager_prototype_setLanguageTest()
    LanguageManager_prototype_getTextTest()

    /*
    // registerElementText est une fonction locale a languagemanager, pas besoin de l'utiliser exterieurement?
    LanguageManager_prototype_registerElementTextTest()
    */

    /*
    NOTES: on doit mentionner que fonctionne avec des divs et innerHTML
        Est qu'on veut que ca fonctionne avec tout les types d'elements?
        -> Sinon mentionner la liste des types d'elements que ca marche
    */
} 