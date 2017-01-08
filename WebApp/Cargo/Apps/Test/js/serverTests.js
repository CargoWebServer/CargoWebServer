
function SayHelloTo(to){
	var helloTo = "Hello " + to + "!"
	return helloTo
}

function executeJsFunctionTest(){
	var to = "champion"
	var params = []
	params.push(createRpcData(to, "STRING"))

    // Call it on the server.
    server.executeJsFunction(
        SayHelloTo.toString(), // The function to execute remotely on server
        params, // The parameters to pass to that function
        function (index, total, caller) { // The progress callback
        },
        function (result, caller) {
        	QUnit.test( "executeJsFunction - success callback", function(expected){
        		return function(assert){
        			assert.ok( expected == "Hello " + to + "!");
        		}}(result[0]))
        },
        function (errMsg, caller) {
        	QUnit.test( "executeJsFunction - error callback", function(assert){
        		assert.ok( false);
        	})
        }, // Error callback
        undefined
        )
}

function serverTests(){
	executeJsFunctionTest()
} 