// Get information from sql.
function getMachines(initCallback) {
	var q = "SELECT ID_CLEF, ID_DES FROM rca_one.rca_machine"
	var fields = new Array()
	var params = new Array() // No param...
	fields[0] = "int" // Le numeros de la cmm
	fields[1] = "string" // Le nom de la cmm
	var results = server.GetDataManager().Read("RCA_ONE", q, fields, params, "", "")
	initCallback(results)
}

function getProduits(machineId, initCallback) {
	var q = "SELECT DISTINCT ID_DES FROM rca_one.rca_unit WHERE ID_MACHINE = ?;"

	var params = new Array() // No param...
	params[0] = machineId

	var fields = new Array()
	fields[0] = "string" // Le nom du produit

	var results = server.GetDataManager().Read("RCA_ONE", q, fields, params, "", "")
	initCallback(results)
}

function getCotes(produitId, initCallback) {
	var q = "SELECT DISTINCT Entry_COD  FROM rca_one.rca_entry INNER JOIN rca_unit ON rca_one.rca_entry.ID_Entry = rca_unit.ID_UNIT WHERE rca_unit.ID_DES = ? ORDER BY Entry_COD"
	var fields = []
	fields[0] = "string" // Le nom de la cote

	var params = []
	params[0] = produitId

	var results = server.GetDataManager().Read("RCA_ONE", q, fields, params, "", "")
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
	initCallback(results)
}

function getTolerances(cote, produitId, initCallback) {
	var q = "SELECT DISTINCT Cote_Nominale, Tol_Min, Tol_Maxi, Entry_DES FROM rca_one.rca_entry INNER JOIN rca_one.rca_unit ON rca_one.rca_entry.ID_Entry "
	q += "= rca_one.rca_unit.ID_UNIT WHERE rca_one.rca_entry.Entry_COD=? AND rca_one.rca_unit.ID_DES=?;"
	var fieldsType = []
	fieldsType[0] = "decimal"
	fieldsType[1] = "decimal"
	fieldsType[2] = "decimal"
	fieldsType[3] = "string"

	var params = []
	params[0] = cote
	params[1] = produitId

	var results = server.GetDataManager().Read("RCA_ONE", q, fieldsType, params, "", "")
	initCallback(results)
}

/**
 * Cette fonction permet de compiler le valeur CSP pour tous les produits.
 */
function compileAnalyseCSP(echantillonSize) {
	console.log("-----------> compile Analyse CSP size ", echantillonSize)
    // New service object
    var service = new Server("localhost", "127.0.0.1", 9595)

	// Initialyse the server connection.
    service.init(
		// onOpenConnectionCallback
		function (service, caller) {
			console.log("get various informations")
			var analyser = new com.cargo.AnalyseurCSP_Interface(service)
			var echantillonSize = caller
			var toAnalyse = []

			// Les Machines
			getMachines(function (machines) {
				for (var i = 0; i < machines.length; i++) {
					// Les Produits
					var machineId = machines[i][0]
					getProduits(machineId, function (machineId, done) {
						if (done == true) {
							console.log("----------------------> machine done!")
						}
						return function (produits) {
							for (var i = 0; i < produits.length; i++) {
								var produitId = produits[i][0]
								// Les cotes
								getCotes(produitId, function (machineId, produitId, done) {
									if (done) {
										console.log("machine and produit done!")
									}
									return function (cotes) {
										for (var i = 0; i < cotes.length; i++) {
											// Les tolerances
											getTolerances(cotes[i][0], produitId, function (machineId, produitId, cote, done) {
												return function (tolerances) {
													if (tolerances.length != 0) {
														// Les valeur d'item mesurés.
														getItems(machineId, produitId, cote, echantillonSize, function (machineId, produitId, cote, tolerances, echantillonSize, done) {
															return function (items) {
																// So here I have the list of result and the tolerance I can now evaluate quality 
																// values.
																if (items.length == echantillonSize) {
																	// Create the Analyse...
																	console.log("------>get analyse ", machineId, produitId, cote)
																	//AnalyseResults(analyser, items, tolerances[0], tolerances[1], tolerances[2], tolerances[3], function(){})
																	toAnalyse.push({ "items": items, "tolzon": tolerances[0], "lowtol": tolerances[1], "uptol": tolerances[2], "toltype": tolerances[3] })
																	//console.log(toAnalyse.length)
																}
																if (done) {
																	console.log("machine produit and cotes are done.")
																}
															}
														} (machineId, produitId, cote, tolerances[0], echantillonSize, done))
													}
												}
											} (machineId, produitId, cotes[i][0], i == cotes.length - 1 && done))
										}
									}
								} (machineId, produitId, i == produits.length - 1 && done))

							}
						}
					} (machineId, i == machines.length - 1))
				}
			})
		},
		// onCloseConnectionCallback
		function () {
		}, echantillonSize)
}

/**
 * Calculate analyse results.
 */
function AnalyseResults(analyser, data, tolzon, lowtol, uptol, toltype, callback) {

	// La list des tests.
	var tests = [
		{ "state": true, "value": 0 },
		{ "state": true, "value": 9 },
		{ "state": true, "value": 6 },
		{ "state": true, "value": 14 },
		{ "state": true, "value": 2 },
		{ "state": true, "value": 4 },
		{ "state": true, "value": 15 },
		{ "state": true, "value": 8 }
	]

	var params = []
	var isPopulation = false

	// Test function call
	// Calculate the stats...
	analyser.analyse(data, tolzon, lowtol, uptol, toltype, isPopulation, tests,
		// Success Callback
		function (result, caller) {
			console.log("------>analyse success!", JSON.stringify(result))
		},
		// Error Callback
		function (errObj, caller) {

		}, {})
}
