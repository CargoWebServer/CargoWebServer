
function createRpcDataTest(){

    var intVar = 1;
    var strVar = "";
    var dblVar = 1.11;
    var boolVar = true;
    var arrayVar = [];
    var objVar = {};
    var bytesVar = new Uint8Array(1);

    QUnit.test( "createRpcData - int ", function(assert){
        assert.ok( createRpcData(intVar, "INTEGER").type == 1)})
    QUnit.test( "createRpcData - string ", function(assert){
        assert.ok( createRpcData(strVar, "STRING").type == 2)})
    QUnit.test( "createRpcData - double ", function(assert){
        assert.ok( createRpcData(dblVar, "DOUBLE").type == 0)})
    QUnit.test( "createRpcData - boolean ", function(assert){
        assert.ok( createRpcData(boolVar, "BOOLEAN").type == 5)})
    QUnit.test( "createRpcData - object ", function(assert){
        assert.ok( createRpcData(objVar, "JSON_STR").type == 4)})
    QUnit.test( "createRpcData - array ", function(assert){
        assert.ok( createRpcData(arrayVar, "JSON_STR").type == 4)})
    QUnit.test( "createRpcData - bytes ", function(assert){
        assert.ok( createRpcData(bytesVar, "BYTES").type == 3)})
    QUnit.test( "createRpcData - undefined", function(assert){
        assert.ok( createRpcData(bytesVar, "aaa") == undefined)})
}

function utilityTests(){
    createRpcDataTest()
} 