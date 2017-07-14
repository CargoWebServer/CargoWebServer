// Get information from sql.
function getMachines(initCallback) {
	var q = "SELECT ID_CLEF, ID_DES FROM rca_one.rca_machine"
	var fields = new Array()
	var params = new Array() // No param...
	fields[0] = "int" // Le numeros de la cmm
	fields[1] = "string" // Le nom de la cmm
	var results = server.GetDataManager().Read("RCA_ONE", q, fields, params, "", "")
	console.log("--> getMachines ", results.length, " found!")
	initCallback(results)
}

function getProduits(machineId, initCallback) {
	var q = "SELECT DISTINCT ID_DES FROM rca_one.rca_unit WHERE ID_MACHINE = ?;"
	
	var params = new Array() // No param...
	params[0] = machineId
	
	var fields = new Array()
	fields[0] = "string" // Le nom du produit
	
	var results = server.GetDataManager().Read("RCA_ONE", q, fields, params, "", "")
	console.log("--> getProduits ", results.length, " found!")
	initCallback(results)
}

function getCotes(produitId, initCallback) {
	var q = "SELECT DISTINCT ID_Entry, Entry_COD  FROM rca_one.rca_entry INNER JOIN rca_unit ON rca_one.rca_entry.ID_Entry = rca_unit.ID_UNIT WHERE rca_unit.ID_DES = ? ORDER BY Entry_COD"
	var fields = []
	fields[0] = "int" // Le numeros de la cote
	fields[1] = "string" // Le nom de la cote

	var params = []
	params[0] = produitId

	var results = server.GetDataManager().Read("RCA_ONE", q, fields, params, "", "")
	console.log("--> getCotes ", results.length, " found!")
	initCallback(results)
}

function getItems(machine, produit, cote, echantillonSize, initCallback) {
    var q = "SELECT Cote, rca_one.rca_unit.UNIT_REF, TIMESTAMP(CONCAT (rca_one.rca_unit.UNIT_DATE, \' \', rca_one.rca_unit.UNIT_END_TIME))"
    q += "AS creation_time, true as IsActive FROM rca_one.rca_entry inner join rca_one.rca_unit on rca_unit.ID_UNIT = rca_one.rca_entry.ID_Entry"
    q += " WHERE rca_one.rca_unit.ID_Machine = ? AND rca_one.rca_unit.ID_DES = ? AND rca_one.rca_entry.Entry_COD = ? "
    q += " AND TIMESTAMP(CONCAT (rca_one.rca_unit.UNIT_DATE, \' \', rca_one.rca_unit.UNIT_END_TIME)) IS NOT NULL ORDER BY creation_time DESC LIMIT 0 , " + echantillonSize + ";"

    var fieldsType = []
    fieldsType[0] = "decimal"
    fieldsType[1] = "string"
    fieldsType[2] = "date"
    fieldsType[3] = "bit"

    var params = []
    params[0] = machine
    params[1] = produit
    params[2] = cote

	var results = server.GetDataManager().Read("RCA_ONE", q, fieldsType, params, "", "")
	console.log("--> getItems ", results.length, " found!")
	initCallback(results)
}

function getTolerances(id_entry, cote, initCallback) {
	var q = "SELECT Cote_Nominale, Tol_Min, Tol_Maxi, Entry_DES FROM rca_one.rca_entry WHERE ID_Entry=? AND Entry_COD=?;"
	var fields = []
	fields[0] = "decimal"
	fields[1] = "decimal"
	fields[2] = "decimal"
	fields[3] = "string"

	var params = []
	params[0] = id_entry
	params[1] = cote

	var results = server.GetDataManager().Read("RCA_ONE", q, fields, params, "", "")
	console.log("--> getTolerances ", results.length, " found!")
	initCallback(results)
}

// Keep echantillonSize as global variable.
var echantillonSize = 30

/**
 * Cette fonction permet de compiler le valeur CSP pour tous les produits.
 */
 function compileAnalyseCSP(echantillonSize_){
	echantillonSize = echantillonSize_
    // New service object
    service = new Server("localhost", "127.0.0.1", 9595)
	
	// Initialyse the server connection.
    service.init( 
		// onOpenConnectionCallback
		function(service){
			console.log("get various informations")
			var sayHello = new com.mycelius.SayHelloInterface(service)
			
			// Les Machines
			getMachines(function(machines){
				for(var i=0; i < machines.length; i++){
					// Les Produits
					var machineId = machines[i][0]
					getProduits(machineId, function(machineId){
						return function(produits){
							for(var i=0; i < produits.length; i++){
								var produitId = produits[i][0]
									// Les cotes
									getCotes(produitId, function(machineId, produitId){
											return function(cotes){
											for(var i=0; i < cotes.length; i++) {
												// Les tolerances
												getTolerances(cotes[i][0], cotes[i][1],function(machineId, produitId, cote){
													return function(tolerances){
														if(tolerances.length != 0){
															// Les valeur d'item mesurés.
															getItems(machineId, produitId, cote, echantillonSize, function(machineId, produitId, cote, tolerances){
																return function(items){
																	// So here I have the list of result and the tolerance I can now evaluate quality 
																	// values.
																	AnalyseResults(sayHello, items, tolerances[0], tolerances[1], tolerances[2], tolerances[3], function(){})
																	
																}
															}(machineId, produitId, cote, tolerances[0]))
														}
													}
												}(machineId, produitId, cotes[i][1]))
											}
										}
									}(machineId, produitId))
								
							}
						}
					}(machineId))
				}
			})
		},
		// onCloseConnectionCallback
		function(){
		})
 }
 
 /**
  * Calculate analyse results.
  */
 function AnalyseResults(sayHello, data, tolzon, lowtol, upTol, tolType, callback) {
	
	// La list des tests.
	var tests = [
		{ "state": true, "value": 0 },
		{ "state": true, "value": 0 },
		{ "state": true, "value": 0 },
		{ "state": true, "value": 0 },
		{ "state": true, "value": 0 },
		{ "state": true, "value": 0 },
		{ "state": true, "value": 0 },
		{ "state": true, "value": 0 }
	]
	
	// Test function call
	
	// Call say hello. 
	sayHello.sayHelloTo("Cargo!!!",
		// Success Callback
		function (result, caller) {
			console.log("------> success!", result)
		},
		// Error Callback
		function (errObj, caller) {

		}, {})
	

}