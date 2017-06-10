/**
 * That file contain informations structure use by
 * DT 3.0 application.
 */

/**
 * Program, the part family. A350, A320, A321 etc...
 */
var programPrototype = new EntityPrototype()
programPrototype.init(
    {
        "TypeName": "DT3_informations.Program",
        "Ids": ["M_id"],
        "Indexs": ["M_id", "M_number", "M_description"], // put here field that you want to use as search key...
        "Fields": ["M_id", "M_number", "M_description", "M_actif", "M_parts"],
        "FieldsType": ["xs.int", "xs.string", "xs.string", "xs.boolean", "[]DT3_informations.Part"],
        "FieldsVisibility": [true, true],
        "FieldsOrder": [0, 1, 2, 3, 4]
    }
)

/**
 * The Part number,  ex. 201522351, 20152269...
 */
var partPrototype = new EntityPrototype()
partPrototype.init(
    {
        "TypeName": "DT3_informations.Part",
        "Ids": ["M_id"],
        "Indexs": ["M_id", "M_number", "M_description"], // put here field that you want to use as search key...
        "Fields": ["M_id", "M_number", "M_description", "M_actif", "M_workorders"],
        "FieldsType": ["xs.int", "xs.string", "xs.string", "xs.boolean", "[]DT3_informations.Workorder"],
        "FieldsVisibility": [true, true, true],
        "FieldsOrder": [0, 1, 2, 3, 4]
    }
)

var departmentPrototype = new EntityPrototype()
departmentPrototype.init(
    {
        "TypeName": "DT3_informations.Department",
        "Ids": ["M_id"],
        "Indexs": ["M_id", "M_description"], // put here field that you want to use as search key...
        "Fields": ["M_id", "M_description"],
        "FieldsType": ["xs.string", "xs.string"],
        "FieldsVisibility": [true, true],
        "FieldsOrder": [0, 1]
    }
)

/**
 * The workpoint information.
 */
var workpointPrototype = new EntityPrototype()
workpointPrototype.init(
    {
        "TypeName": "DT3_informations.Workpoint",
        "Ids": ["M_id"],
        "Indexs": ["M_id", "M_department", "M_description"], // put here field that you want to use as search key...
        "Fields": ["M_id", "M_department", "M_description"],
        "FieldsType": ["xs.string", "DT3_informations.Department:Ref", "xs.string"],
        "FieldsVisibility": [true, true, true],
        "FieldsOrder": [0, 1, 2]
    }
)

/**
 * The work path regroup workpoint information operation number.
 */
var workpathPrototype = new EntityPrototype()
workpathPrototype.init(
    {
        "TypeName": "DT3_informations.Workpath",
        "Ids": ["M_id"], // Number in the system.
        "Indexs": ["M_id"], // put here field that you want to use as search key...
        "Fields": ["M_id", "M_workpoint", "M_part", "M_stepNum"],
        "FieldsType": ["xs.string", "DT3_informations.Workpoint:Ref", "DT3_informations.Part:Ref", "xs.int"],
        "FieldsVisibility": [true, true, true, true],
        "FieldsOrder": [0, 1, 2, 3]
    }
)

/**
 * The work process keep information of the workorder state of production and 
 * his history.
 */
var workprocessPrototype = new EntityPrototype()
workprocessPrototype.init(
    {
        "TypeName": "DT3_informations.Workprocess",
        "Ids": [], // Number in the system.
        "Indexs": [], // put here field that you want to use as search key...
        "Fields": ["M_workorder", "M_workpoint", "M_workpath", "M_scheduledDate", "M_dueDate", "M_lastActivityDate"],
        "FieldsType": ["DT3_informations.Workorder:Ref", "DT3_informations.Workpoint:Ref", "DT3_informations.Workpath:Ref", "xs.dateTime", "xs.dateTime", "xs.dateTime"],
        "FieldsVisibility": [true, true, true, true, true, true],
        "FieldsOrder": [0, 1, 2, 3, 4, 5]
    }
)

/**
 * The workorder information. This represent unit of work. CDE...
 */
var workorderPrototype = new EntityPrototype()
workorderPrototype.init(
    {
        "TypeName": "DT3_informations.Workorder",
        "Ids": ["M_id"], // Number in the system.
        "Indexs": ["M_id", "M_serial", "M_order"], // put here field that you want to use as search key...
        "Fields": ["M_id", "M_order", "M_serial", "M_creation_date", "M_part"],
        "FieldsType": ["xs.int", "xs.string", "xs.string", "xs.date", "DT3_informations.Part:Ref"],
        "FieldsVisibility": [true, true],
        "FieldsOrder": [0, 1]
    }
)

/////////////////////////////////////////////////////////////////////////
// Initialisations
/////////////////////////////////////////////////////////////////////////
var departements = {}

// Get department infomation from shopvue.
function getDepartmentsFromShopVue() {
	var query = "SELECT [DepartmentNum] "
	query += ",[DepartmentDescription] "
	query += ",[IsInactive] "
	query += "FROM [ShopVue].[dbo].[fDepartment]"

	var params = new Array() // No param...
	var fields = new Array()
	fields[0] = "string"
	fields[1] = "string"
	fields[2] = "int"
	
	var results = server.GetDataManager().Read("ShopVue", query, fields, params, "", "")
	// Keep the first object.
	//results = results[0]
	var departments = []
	for (var i = 0; i < results.length; i++) {
		var val = results[i]
		var department = new DT3_informations.Department()
		department.M_id = val[0]
		department.M_description = val[1]
		departments.push(department)
	}
	return departments
}

function updateDepartmentData(){
	// Set department
	var departements_ = getDepartmentsFromShopVue()
	for(var i=0; i < departements_.length; i++){
		
		var departement = server.GetEntityManager().GetObjectById("DT3_informations", departements_[i].TYPENAME, [departements_[i].M_id], "", "")
		if(departement == null){
			console.log("---->" + departements_[i].TYPENAME)
			departement = server.GetEntityManager().SaveEntity(departements_[i], departements_[i].TYPENAME, "", "") 
		}
		departements[departement.M_id] = departement
	}
	// console.log("---->" + departements)
}

/**
 * That function is use to initialise informations.
 */
function InitInformations() {
	console.log("Initialise Dt3 informations")
	var store = server.GetDataManager().GetDataStore("DT3_informations", "", "")
	// Creation of entity prototypes.
	if(store == null){
		console.log("Create DT3_informations data store")
		
		// Create the data store
		server.GetDataManager().CreateDataStore("DT3_informations", 2, 1, "", "")
		
		console.log("Create DT3_informations entity prototypes...")
		
		// Create the data type
		server.GetEntityManager().CreateEntityPrototype("DT3_informations", departmentPrototype, "", "")
		server.GetEntityManager().CreateEntityPrototype("DT3_informations", programPrototype, "", "")
		server.GetEntityManager().CreateEntityPrototype("DT3_informations", partPrototype, "", "")
		server.GetEntityManager().CreateEntityPrototype("DT3_informations", workpointPrototype, "", "")
		server.GetEntityManager().CreateEntityPrototype("DT3_informations", workorderPrototype, "", "")
		server.GetEntityManager().CreateEntityPrototype("DT3_informations", workpathPrototype, "", "")
		server.GetEntityManager().CreateEntityPrototype("DT3_informations", workprocessPrototype, "", "")
		
		
	}
	
	// Connect to the store.
	store.Connect()
	
	// Now I will initialize the data...
	updateDepartmentData()
	
}

// TODO put in the task manager...
InitInformations()