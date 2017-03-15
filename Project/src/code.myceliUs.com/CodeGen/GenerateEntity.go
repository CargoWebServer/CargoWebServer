package main

import (
	"strconv"
	"strings"

	"code.myceliUs.com/Utility"
	"code.myceliUs.com/XML_Schemas"
)

/**
 * That function generate the persistence functionnality of object...
 */
func generateEntity(packageId string) {
	packageStr := "package Server\n"
	classStr := ""

	classes := packagesMembers[packageId]
	imports := make([]string, 0)
	imports = append(imports, outputPath+packageId)

	imports = append(imports, "encoding/json")
	imports = append(imports, "code.myceliUs.com/Utility")
	imports = append(imports, "log")
	imports = append(imports, "strings")

	for i := 0; i < len(classes); i++ {
		class := classesMap[classes[i]]

		if class != nil {
			if Utility.Contains(abstractClassLst, class.Name) == false {
				className := class.Name
				// If the class is use as a base to other class
				// an interface must be generated...
				if Utility.Contains(superClassesLst, class.Name) == true {
					className += "_impl"
				}

				classStr += "\n////////////////////////////////////////////////////////////////////////////////\n"
				classStr += "//              			" + class.Name + "\n"
				classStr += "////////////////////////////////////////////////////////////////////////////////\n"

				classStr += "/** local type **/\n"
				classStr += "type " + packageId + "_" + class.Name + "Entity struct{\n"
				classStr += "	/** not the object id, except for the definition **/\n"
				classStr += "	uuid string\n"
				classStr += "	parentPtr 			Entity\n"
				classStr += "	parentUuid 			string\n"
				classStr += "	childsPtr  			[]Entity\n"
				classStr += "	childsUuid  		[]string\n"
				classStr += "	referencesUuid  	[]string\n"
				classStr += "	referencesPtr  	    []Entity\n"
				classStr += "	prototype      		*EntityPrototype\n"

				// Keep track of object that have reference to that object.
				classStr += "	referenced  		[]EntityRef\n"

				classStr += "	object *" + packageId + "." + className + "\n"
				classStr += "}\n\n"

				// The new  entity...
				classStr += generateNewEntityFunc(packageId, class)

				classStr += "/** Entity functions **/\n"

				classStr += "func(this *" + packageId + "_" + class.Name + "Entity) GetTypeName()string{\n"
				classStr += "	return \"" + packageId + "." + class.Name + "\"\n"
				classStr += "}\n"

				classStr += "func(this *" + packageId + "_" + class.Name + "Entity) GetUuid()string{\n"
				classStr += "	return this.uuid\n"
				classStr += "}\n"

				classStr += "func(this *" + packageId + "_" + class.Name + "Entity) GetParentPtr()Entity{\n"
				classStr += "	return this.parentPtr\n"
				classStr += "}\n\n"

				classStr += "func(this *" + packageId + "_" + class.Name + "Entity) SetParentPtr(parentPtr Entity){\n"
				classStr += "	this.parentPtr=parentPtr\n"
				classStr += "}\n\n"

				classStr += "func(this *" + packageId + "_" + class.Name + "Entity) AppendReferenced(name string, owner Entity){\n"
				classStr += "	if owner.GetUuid() == this.GetUuid() {\n"
				classStr += "		return\n"
				classStr += "	}\n"

				classStr += "	var ref EntityRef\n"
				classStr += "	ref.Name = name\n"
				classStr += "	ref.OwnerUuid = owner.GetUuid()\n"
				classStr += "	for i:=0; i<len(this.referenced); i++ {\n"
				classStr += "		if this.referenced[i].Name == ref.Name && this.referenced[i].OwnerUuid == ref.OwnerUuid { \n"
				classStr += "			return;\n"
				classStr += "		}\n"
				classStr += "	}\n"

				classStr += "	this.referenced = append(this.referenced, ref)\n"
				classStr += "}\n\n"

				classStr += "func(this *" + packageId + "_" + class.Name + "Entity) GetReferenced() []EntityRef{\n"
				classStr += "	return this.referenced\n"
				classStr += "}\n\n"

				classStr += "func(this *" + packageId + "_" + class.Name + "Entity) RemoveReferenced(name string, owner Entity) {\n"
				classStr += "	var referenced []EntityRef\n"
				classStr += "	referenced = make([]EntityRef,0)\n"
				classStr += "	for i := 0; i < len(this.referenced); i++ {\n"
				classStr += "		ref := this.referenced[i]\n"
				classStr += "		if !(ref.Name == name && ref.OwnerUuid == owner.GetUuid()) {\n"
				classStr += "			referenced = append(referenced, ref)\n"
				classStr += "		}\n"
				classStr += "	}\n"
				classStr += "	// Set the reference.\n"
				classStr += "	this.referenced = referenced\n"
				classStr += "}\n\n"

				classStr += "func(this *" + packageId + "_" + class.Name + "Entity) RemoveReference(name string, reference Entity){\n"
				classStr += "	refsUuid := make([]string, 0)\n"
				classStr += "	refsPtr := make([]Entity, 0)\n"

				classStr += "	for i := 0; i < len(this.referencesUuid); i++ {\n"
				classStr += "		refUuid := this.referencesUuid[i]\n"
				classStr += "		if refUuid != reference.GetUuid() {\n"
				classStr += "			refsPtr = append(refsPtr, reference)\n"
				classStr += "			refsUuid = append(refsUuid, reference.GetUuid())\n"
				classStr += "		}\n"
				classStr += "	}\n"
				classStr += "	// Set the new array...\n"
				classStr += "	this.SetReferencesUuid(refsUuid)\n"
				classStr += "	this.SetReferencesPtr(refsPtr)\n\n"

				classStr += "	var removeMethode = \"Remove\" + strings.ToUpper(name[2:3]) + name[3:]\n"
				classStr += "	params := make([]interface{}, 1)\n"
				classStr += "	params[0] = reference.GetObject()\n"
				classStr += "	Utility.CallMethod(this.GetObject(), removeMethode, params)\n"

				classStr += "}\n\n"

				classStr += "func(this *" + packageId + "_" + class.Name + "Entity) GetChildsPtr() []Entity{\n"
				classStr += "	return this.childsPtr\n"
				classStr += "}\n\n"

				classStr += "func(this *" + packageId + "_" + class.Name + "Entity) SetChildsPtr(childsPtr[]Entity){\n"
				classStr += "	this.childsPtr = childsPtr\n"
				classStr += "}\n\n"

				classStr += "func(this *" + packageId + "_" + class.Name + "Entity) GetChildsUuid() []string{\n"
				classStr += "	return this.childsUuid\n"
				classStr += "}\n\n"

				classStr += "func(this *" + packageId + "_" + class.Name + "Entity) SetChildsUuid(childsUuid[]string){\n"
				classStr += "	this.childsUuid = childsUuid\n"
				classStr += "}\n\n"

				classStr += "/**\n"
				classStr += " * Remove a chidl uuid form the list of child in an entity.\n"
				classStr += " */\n"
				classStr += "func(this *" + packageId + "_" + class.Name + "Entity) RemoveChild(name string, uuid string) {\n"
				classStr += " 	childsUuid := make([]string, 0)\n"
				classStr += " 	for i := 0; i < len(this.GetChildsUuid()); i++ {\n"
				classStr += " 		if this.GetChildsUuid()[i] != uuid {\n"
				classStr += " 			childsUuid = append(childsUuid, this.GetChildsUuid()[i])\n"
				classStr += " 		}\n"
				classStr += " 	}\n"
				classStr += " \n"
				classStr += " 	this.childsUuid = childsUuid\n"

				classStr += "	params := make([]interface{}, 1)\n"
				classStr += " 	childsPtr := make([]Entity, 0)\n"
				classStr += " 	for i := 0; i < len(this.GetChildsPtr()); i++ {\n"
				classStr += " 		if this.GetChildsPtr()[i].GetUuid() != uuid {\n"
				classStr += " 			childsPtr = append(childsPtr, this.GetChildsPtr()[i])\n"
				classStr += " 		}else{\n"
				classStr += "			params[0] = this.GetChildsPtr()[i].GetObject()\n"
				classStr += " 		}\n"
				classStr += " 	}\n"
				classStr += " 	this.childsPtr = childsPtr\n\n"
				classStr += "	var removeMethode = \"Remove\" + strings.ToUpper(name[0:1]) + name[1:]\n"
				classStr += "	Utility.CallMethod(this.GetObject(), removeMethode, params)\n"

				classStr += " }\n\n"

				classStr += "func(this *" + packageId + "_" + class.Name + "Entity) GetReferencesUuid() []string{\n"
				classStr += "	return this.referencesUuid\n"
				classStr += "}\n\n"

				classStr += "func(this *" + packageId + "_" + class.Name + "Entity) SetReferencesUuid(refsUuid[]string){\n"
				classStr += "	this.referencesUuid = refsUuid\n"
				classStr += "}\n\n"

				classStr += "func(this *" + packageId + "_" + class.Name + "Entity) GetReferencesPtr() []Entity{\n"
				classStr += "	return this.referencesPtr\n"
				classStr += "}\n\n"

				classStr += "func(this *" + packageId + "_" + class.Name + "Entity) SetReferencesPtr(refsPtr[]Entity){\n"
				classStr += "	this.referencesPtr = refsPtr\n"
				classStr += "}\n\n"

				classStr += "func(this *" + packageId + "_" + class.Name + "Entity) GetObject() interface{}{\n"
				classStr += "	return this.object\n"
				classStr += "}\n\n"

				classStr += "func(this *" + packageId + "_" + class.Name + "Entity) NeedSave() bool{\n"
				classStr += "	return this.object.NeedSave\n"
				classStr += "}\n\n"

				classStr += "func(this *" + packageId + "_" + class.Name + "Entity) SetNeedSave(needSave bool) {\n"
				classStr += "	this.object.NeedSave = needSave\n"
				classStr += "}\n\n"

				classStr += "func(this *" + packageId + "_" + class.Name + "Entity) IsInit() bool{\n"
				classStr += "	return this.object.IsInit\n"
				classStr += "}\n\n"

				classStr += "func(this *" + packageId + "_" + class.Name + "Entity) SetInit(isInit bool) {\n"
				classStr += "	this.object.IsInit = isInit\n"
				classStr += "}\n\n"

				classStr += "func(this *" + packageId + "_" + class.Name + "Entity) GetChecksum() string{\n"
				classStr += "	objectStr, _ := json.Marshal(this.object)\n"
				classStr += "	return  Utility.GetMD5Hash(string(objectStr))\n"
				classStr += "}\n\n"

				classStr += "func(this *" + packageId + "_" + class.Name + "Entity) Exist() bool{\n"
				classStr += "	var query EntityQuery\n"
				classStr += "	query.TypeName = \"" + packageId + "." + class.Name + "\"\n"
				classStr += "	query.Indexs = append(query.Indexs, \"uuid=\"+this.uuid)\n"
				classStr += "	query.Fields = append(query.Fields, \"uuid\")\n"
				classStr += "	var fieldsType []interface {} // not use...\n"
				classStr += "	var params []interface{}\n"
				classStr += "	queryStr, _ := json.Marshal(query)\n"
				classStr += "	results, err := GetServer().GetDataManager().readData(" + packageId + "DB, string(queryStr), fieldsType, params)\n"
				classStr += "	if err != nil || len(results) == 0 {\n"
				classStr += "		return false\n"
				classStr += "	}\n"
				classStr += "	return len(results[0][0].(string)) > 0\n\n"
				classStr += "}\n\n"

				classStr += "/**\n"
				classStr += "* Return the entity prototype.\n"
				classStr += "*/\n"
				classStr += "func(this *" + packageId + "_" + class.Name + "Entity) GetPrototype() *EntityPrototype {\n"
				classStr += "	return this.prototype\n"
				classStr += "}\n"

				// Entity prototype...
				classStr += generateEntityPrototypeFunc(packageId, class)

				// The CRUD...
				// Create
				classStr += generateEntitySaveFunc(packageId, class)

				// Read
				classStr += generateEntityInitFunc(packageId, class)

				// Create the entity from a given object.
				classStr += generateNewEntityFromObject(packageId, class)

				// Delete
				classStr += generateEntityDeleteFunc(packageId, class)

				// Exist function
				classStr += generateEntityExistsFunc(packageId, class)

				// Append child
				classStr += generateAppendChildFunc(packageId, class)

				// Append reference
				classStr += generateAppendRefFunc(packageId, class)

				imports_ := getClassImports(class)
				for i := 0; i < len(imports_); i++ {
					if !Utility.Contains(imports, imports_[i]) {
						//imports = append(imports, imports_[i])
					}
				}
			} else {
				// Entity prototype...
				classStr += generateEntityPrototypeFunc(packageId, class)
			}
		}
	}

	importsStr := ""

	if len(imports) > 0 {
		importsStr += "import(\n"
		for j := 0; j < len(imports); j++ {
			importsStr += "	\"" + imports[j] + "\"\n"
		}
		importsStr += ")\n\n"
	}

	classStr += generateRegisterObjects(packageId)
	classStr += generateCreatePrototypes(packageId)

	// Wrote the file...
	if len(classStr) > 0 {
		WriteClassFile("code.myceliUs.com/CargoWebServer/Cargo/Server/", "", "Entity_"+packageId, packageStr+importsStr+classStr)
	}
}

/**
 * That function is use to generate the new entity...
 */
func generateNewEntityFunc(packageId string, class *XML_Schemas.CMOF_OwnedMember) string {
	entityConstructorStr := "/** Constructor function **/\n"

	className := class.Name
	if Utility.Contains(superClassesLst, class.Name) == true {
		className += "_impl"
	}

	packageId_ := packageId
	if packageId == "BPMNDI" || packageId == "DC" || packageId == "DI" {
		packageId_ = "BPMN20"
	}

	entityConstructorStr += "func (this *EntityManager) New" + packageId + class.Name + "Entity(parentUuid string, objectId string, object interface{}) *" + packageId + "_" + class.Name + "Entity{\n"
	entityConstructorStr += "	var uuidStr string\n"
	entityConstructorStr += "	if len(objectId) > 0 {\n"
	entityConstructorStr += "		if Utility.IsValidEntityReferenceName(objectId){\n"
	entityConstructorStr += "			uuidStr = objectId\n"
	entityConstructorStr += "		}else{\n"
	entityConstructorStr += "			uuidStr  = " + packageId + class.Name + "Exists(objectId)\n"
	entityConstructorStr += "		}\n"
	entityConstructorStr += "	}\n"
	entityConstructorStr += "	if object != nil{\n"
	entityConstructorStr += "		object.(*" + packageId + "." + className + ").TYPENAME = \"" + packageId + "." + class.Name + "\"\n"
	entityConstructorStr += "	}\n"
	entityConstructorStr += "	prototype, _ := GetServer().GetEntityManager().getEntityPrototype(\"" + packageId + "." + className + "\",\"" + packageId_ + "\")\n"
	entityConstructorStr += "	if len(uuidStr) > 0 {\n"
	entityConstructorStr += "		if object != nil{\n"
	entityConstructorStr += "			object.(*" + packageId + "." + className + ").UUID = uuidStr\n"
	entityConstructorStr += "		}\n"
	entityConstructorStr += "		if val, ok := this.contain(uuidStr);ok {\n"
	entityConstructorStr += "			if object != nil{\n"
	entityConstructorStr += "				this.setObjectValues(val, object)\n\n"
	entityConstructorStr += "			}\n"
	entityConstructorStr += "			return val.(*" + packageId + "_" + class.Name + "Entity)\n"
	entityConstructorStr += "		}\n"
	entityConstructorStr += "	}else{\n"
	entityConstructorStr += "		if len(prototype.Ids) == 1 {\n"
	entityConstructorStr += "			// Here there is a new entity...\n"
	entityConstructorStr += "			uuidStr = \"Config.Configurations%\" + Utility.RandomUUID()\n"
	entityConstructorStr += "		} else {\n"
	entityConstructorStr += "			var keyInfo string\n"
	entityConstructorStr += "			if len(parentUuid) > 0{\n"
	entityConstructorStr += "				keyInfo += parentUuid + \":\"\n"
	entityConstructorStr += "			}\n"
	entityConstructorStr += "			keyInfo += prototype.TypeName + \":\"\n"
	entityConstructorStr += "			for i := 1; i < len(prototype.Ids); i++ {\n"
	entityConstructorStr += "				var getter = \"Get\" + strings.ToUpper(prototype.Ids[i][2:3]) + prototype.Ids[i][3:]\n"
	entityConstructorStr += "				params := make([]interface{}, 0)\n"
	entityConstructorStr += "				value, _ := Utility.CallMethod(object, getter, params)\n"
	entityConstructorStr += "				keyInfo += Utility.ToString(value)\n"
	entityConstructorStr += "				// Append underscore for readability in case of problem...\n"
	entityConstructorStr += "				if i < len(prototype.Ids)-1 {\n"
	entityConstructorStr += "					keyInfo += \"_\"\n"
	entityConstructorStr += "				}\n"
	entityConstructorStr += "			}\n\n"
	entityConstructorStr += "			// The uuid is in that case a MD5 value.\n"
	entityConstructorStr += "			uuidStr = prototype.TypeName + \"%\" + Utility.GenerateUUID(keyInfo)\n"
	entityConstructorStr += "		}\n"
	entityConstructorStr += "	}\n"
	entityConstructorStr += "	entity := new(" + packageId + "_" + class.Name + "Entity)\n"
	entityConstructorStr += "	if object == nil{\n"
	entityConstructorStr += "		entity.object = new(" + packageId + "." + className + ")\n"
	entityConstructorStr += "		entity.SetNeedSave(true)\n"
	entityConstructorStr += "	}else{\n"
	entityConstructorStr += "		entity.object = object.(*" + packageId + "." + className + ")\n"
	entityConstructorStr += "		entity.SetNeedSave(true)\n"
	entityConstructorStr += "	}\n"

	entityConstructorStr += "	entity.object.TYPENAME = \"" + packageId + "." + class.Name + "\"\n\n"
	entityConstructorStr += "	entity.object.UUID = uuidStr\n"
	entityConstructorStr += "	entity.SetInit(false)\n"
	entityConstructorStr += "	entity.uuid = uuidStr\n"
	entityConstructorStr += "	this.insert(entity)\n"

	if strings.HasSuffix(className, "_impl") {
		className = strings.Replace(className, "_impl", "", -1)
	}

	entityConstructorStr += "	entity.prototype = prototype\n"

	entityConstructorStr += "	return entity\n"
	entityConstructorStr += "}\n\n"

	return entityConstructorStr
}

func generateRegisterObjects(packageId string) string {

	functionStr := "/** Register the entity to the dynamic typing system. **/\n"
	functionStr += "func (this *EntityManager) register" + packageId + "Objects(){\n"
	classes := packagesMembers[packageId]
	for i := 0; i < len(classes); i++ {
		class := classesMap[classes[i]]

		if class != nil {
			if Utility.Contains(abstractClassLst, class.Name) == false {
				className := class.Name
				if Utility.Contains(superClassesLst, class.Name) {
					className += "_impl"
				}
				functionStr += "	Utility.RegisterType((*" + packageId + "." + className + ")(nil))\n"
			}
		}
	}

	functionStr += "}\n\n"

	return functionStr
}

func generateCreatePrototypes(packageId string) string {

	functionStr := "/** Create entity prototypes contain in a package **/\n"
	functionStr += "func (this *EntityManager) create" + packageId + "Prototypes(){\n"
	classes := packagesMembers[packageId]
	for i := 0; i < len(classes); i++ {
		class := classesMap[classes[i]]
		if class != nil {
			functionStr += "	this.create_" + packageId + "_" + class.Name + "EntityPrototype() \n"
		}
	}

	functionStr += "}\n\n"

	return functionStr
}

func generateNewEntityFromObject(packageId string, class *XML_Schemas.CMOF_OwnedMember) string {
	objectType := "*" + packageId + "." + class.Name
	if Utility.Contains(superClassesLst, class.Name) {
		objectType += "_impl"
	}

	functionStr := "/** instantiate a new entity from an existing object. **/\n"
	functionStr += "func (this *EntityManager) New" + packageId + class.Name + "EntityFromObject(object " + objectType + ") *" + packageId + "_" + class.Name + "Entity {\n"
	functionStr += "	 return this.New" + packageId + class.Name + "Entity(\"\", object.UUID, object)\n"
	functionStr += "}\n\n"

	return functionStr
}

func generateAppendRefFunc(packageId string, class *XML_Schemas.CMOF_OwnedMember) string {
	functionStr := "/** Append reference entity into parent entity. **/\n"
	functionStr += "func (this *" + packageId + "_" + class.Name + "Entity) AppendReference(reference Entity) {\n\n"
	functionStr += "	 // Here i will append the reference uuid\n"
	functionStr += "	 index := -1\n"

	functionStr += "	 for i := 0; i < len(this.referencesUuid); i++ {\n"
	functionStr += "	 	refUuid := this.referencesUuid[i]\n"
	functionStr += "	 	if refUuid == reference.GetUuid() {\n"
	functionStr += "	 		index = i\n"
	functionStr += "	 		break\n"
	functionStr += "	 	}\n"
	functionStr += "	 }\n"

	functionStr += "	 if index == -1 {\n"
	functionStr += "	 	this.referencesUuid = append(this.referencesUuid, reference.GetUuid())\n"
	functionStr += "	 	this.referencesPtr = append(this.referencesPtr, reference)\n"
	functionStr += "	 }else{\n"
	functionStr += "	 	// The reference must be update in that case.\n"
	functionStr += "	 	this.referencesPtr[index]  = reference\n"
	functionStr += "	 }\n"

	functionStr += "}\n"

	return functionStr
}

func generateAppendChildFunc(packageId string, class *XML_Schemas.CMOF_OwnedMember) string {
	functionStr := "/** Append child entity into parent entity. **/\n"
	functionStr += "func (this *" + packageId + "_" + class.Name + "Entity) AppendChild(attributeName string, child Entity) error {\n\n"

	functionStr += "	// Append child if is not there...\n"
	functionStr += "	if !Utility.Contains(this.childsUuid, child.GetUuid()) {\n"
	functionStr += "		this.childsUuid = append(this.childsUuid, child.GetUuid())\n"
	functionStr += "		this.childsPtr = append(this.childsPtr, child)\n"
	functionStr += "	} else {\n"
	functionStr += "		childsPtr := make([]Entity, 0)\n"
	functionStr += "		for i := 0; i < len(this.childsPtr); i++ {\n"
	functionStr += "			if this.childsPtr[i].GetUuid() != child.GetUuid() {\n"
	functionStr += "				childsPtr = append(childsPtr, this.childsPtr[i])\n"
	functionStr += "			}\n"
	functionStr += "		}\n"
	functionStr += "		childsPtr = append(childsPtr, child)\n"
	functionStr += "		this.SetChildsPtr(childsPtr)\n"
	functionStr += "	}\n"

	functionStr += "	// Set this as parent in the child\n"
	functionStr += "	child.SetParentPtr(this)\n\n"
	functionStr += "	params := make([]interface{}, 1)\n"
	functionStr += "	params[0] = child.GetObject()\n"
	functionStr += "	attributeName = strings.Replace(attributeName,\"M_\", \"\", -1)\n"
	functionStr += "	methodName := \"Set\" + strings.ToUpper(attributeName[0:1]) + attributeName[1:]\n"
	functionStr += "	_, invalidMethod := Utility.CallMethod(this.object, methodName, params)\n"
	functionStr += "	if invalidMethod != nil {\n"
	functionStr += "		return invalidMethod.(error)\n"
	functionStr += "	}\n"
	functionStr += "	return nil\n"
	functionStr += "}\n"

	return functionStr
}

func generateEntityAttribute(attribute *XML_Schemas.CMOF_OwnedAttribute, packName string, prototypeName string, index *int, isAssociation bool, isId bool, isIndex bool) string {
	var memberStr string
	var typeName, isPrimitive, _ = getAttributeTypeName(attribute)

	// Here I will test if the attribute is an id...
	//log.Println("attribute!", attribute.Name)

	// Append attribute to ids
	if isId {
		memberStr += "	" + prototypeName + ".Ids = append(" + prototypeName + ".Ids,\"M_" + attribute.Name + "\")\n"
	}

	// Append attribute to indexs
	if isIndex {
		memberStr += "	" + prototypeName + ".Indexs = append(" + prototypeName + ".Indexs,\"M_" + attribute.Name + "\")\n"
	}

	memberStr += "	" + prototypeName + ".FieldsOrder = append(" + prototypeName + ".FieldsOrder," + strconv.Itoa(*index) + ")\n"
	if isAssociation {
		memberStr += "	" + prototypeName + ".FieldsVisibility = append(" + prototypeName + ".FieldsVisibility,false)\n"
	} else {
		memberStr += "	" + prototypeName + ".FieldsVisibility = append(" + prototypeName + ".FieldsVisibility,true)\n"
	}

	*index++
	attrNs := getAttributeNs(attribute)

	//isInterface := Utility.Contains(superClassesLst, typeName)

	if len(attrNs) > 0 && isPrimitive == false {
		if len(packName) > 0 {
			if attrNs != packName {
				typeName = attrNs + "." + typeName
			}
		}
	}

	var memberName = attribute.Name
	if isAssociation {
		memberName += "Ptr"
	}

	memberStr += "	" + prototypeName + ".Fields = append(" + prototypeName + ".Fields,\"M_" + memberName + "\")\n"
	memberTypeName := typeName

	if IsRef(attribute) {
		memberTypeName += ":Ref"
	}

	if isPrimitive {
		// Primitive are string, bool, int, float64 or interface{}
		if isId {
			// Set the type name as id.
			memberTypeName = "ID"
		}

		// Name that must be change...
		if memberTypeName == "bool" {
			memberTypeName = "boolean"
		}

		if memberTypeName == "[]unit8" {
			memberTypeName = "base64Binary"
		}

		if memberTypeName == "int64" {
			memberTypeName = "long"
		}

		memberTypeName = "xs." + memberTypeName

		if attribute.Upper == "*" {
			memberTypeName = "[]" + memberTypeName
		}

		memberStr += "	" + prototypeName + ".FieldsType = append(" + prototypeName + ".FieldsType,\"" + memberTypeName + "\")\n"
	} else if simpleTypesMap[typeName] != nil {
		// Here the type is an enum...
		enumStr := "enum:"
		enum := enumMap[typeName]
		for i := 0; i < len(enum.Litterals); i++ {
			litteral := enum.Litterals[i]
			enumValue := litteral.Name
			enumValue = strings.ToUpper(enumValue[0:1]) + enumValue[1:]
			enumStr += typeName + "_" + enumValue
			if i < len(enum.Litterals)-1 && len(enum.Litterals) > 1 {
				enumStr += ":"
			}
		}
		memberStr += "	" + prototypeName + ".FieldsType = append(" + prototypeName + ".FieldsType,\"" + enumStr + "\")\n"
	} else {
		// Here the type is another type...
		if attribute.Upper == "*" {
			memberTypeName = "[]" + packName + "." + memberTypeName
		} else {
			memberTypeName = packName + "." + memberTypeName
		}
		memberStr += "	" + prototypeName + ".FieldsType = append(" + prototypeName + ".FieldsType,\"" + memberTypeName + "\")\n"
	}

	return memberStr
}

func generateEntityAttributes(class *XML_Schemas.CMOF_OwnedMember, packageId string, prototypeVar string, index *int) string {
	// The super classes code...
	classStr := ""
	superClasses := getSuperClasses(class.Name)
	for i := 0; i < len(superClasses); i++ {
		superClass := classesMap[superClasses[i]]
		classStr += "\n	/** members of " + superClass.Name + " **/\n"
		if len(superClass.Attributes) == 0 {
			classStr += "	/** No members **/\n"
		}
		for j := 0; j < len(superClass.Attributes); j++ {
			attribute := superClass.Attributes[j]
			classStr += generateEntityAttribute(&attribute, packageId, prototypeVar, index, false, isId(superClass.Name, attribute.Name), isIndex(superClass.Name, attribute.Name))
		}
	}

	// Now the attributes...
	classStr += "\n	/** members of " + class.Name + " **/\n"
	if len(class.Attributes) == 0 {
		classStr += "	/** No members **/\n"
	}

	for i := 0; i < len(class.Attributes); i++ {
		attribute := class.Attributes[i]
		classStr += generateEntityAttribute(&attribute, packageId, prototypeVar, index, false, isId(class.Name, attribute.Name), isIndex(class.Name, attribute.Name))
	}

	return classStr
}

func generateEntityAssociations(class *XML_Schemas.CMOF_OwnedMember, packageId string, prototypeVar string, index *int) string {
	// The super classes code...
	classStr := ""
	associations := getClassAssociations(class)
	if len(associations) > 0 {
		classStr += "\n	/** associations of " + class.Name + " **/\n"
	}

	for _, association := range associations {
		classStr += generateEntityAttribute(association, packageId, prototypeVar, index, true, isId(class.Name, association.Name), isIndex(class.Name, association.Name))
	}

	return classStr
}

/**
 * Generate the entity proptype creation funtion of a given class.
 */
func generateEntityPrototypeFunc(packageId string, class *XML_Schemas.CMOF_OwnedMember) string {

	entityPrototypeStr := "/** Entity Prototype creation **/\n"
	entityPrototypeStr += "func (this *EntityManager) create_" + packageId + "_" + class.Name + "EntityPrototype() {\n\n"

	// I will create the prototype...
	prototypeVar := strings.ToLower(class.Name[0:1]) + class.Name[1:] + "EntityProto"

	entityPrototypeStr += "	var " + prototypeVar + " EntityPrototype\n"

	entityPrototypeStr += "	" + prototypeVar + ".TypeName = \"" + packageId + "." + class.Name + "\"\n"

	if Utility.Contains(abstractClassLst, class.Name) {
		entityPrototypeStr += "	" + prototypeVar + ".IsAbstract=true\n"
	}

	superClasses := getSuperClasses(class.Name)
	for i := 0; i < len(superClasses); i++ {
		superClass := classesMap[superClasses[i]]
		superClassPackName := membersPackage[superClass.Name]
		entityPrototypeStr += "	" + prototypeVar + ".SuperTypeNames = append(" + prototypeVar + ".SuperTypeNames, \"" + superClassPackName + "." + superClass.Name + "\")\n"
	}

	// Now the implementation class.
	implementationClasses := getImplementationClasses(class.Name)
	for i := 0; i < len(implementationClasses); i++ {
		implementationClass := classesMap[implementationClasses[i]]
		implementationClassPackName := membersPackage[implementationClass.Name]
		if implementationClassPackName+"."+implementationClass.Name != packageId+"."+class.Name {
			entityPrototypeStr += "	" + prototypeVar + ".SubstitutionGroup = append(" + prototypeVar + ".SubstitutionGroup, \"" + implementationClassPackName + "." + implementationClass.Name + "\")\n"
		}
	}

	// Generate the entity attributes...
	index := 0

	// The entity uuid
	entityPrototypeStr += "	" + prototypeVar + ".Ids = append(" + prototypeVar + ".Ids,\"uuid\")\n"
	entityPrototypeStr += "	" + prototypeVar + ".Fields = append(" + prototypeVar + ".Fields,\"uuid\")\n"
	entityPrototypeStr += "	" + prototypeVar + ".FieldsType = append(" + prototypeVar + ".FieldsType,\"xs.string\")\n"
	entityPrototypeStr += "	" + prototypeVar + ".FieldsOrder = append(" + prototypeVar + ".FieldsOrder," + strconv.Itoa(index) + ")\n"
	entityPrototypeStr += "	" + prototypeVar + ".FieldsVisibility = append(" + prototypeVar + ".FieldsVisibility,false)\n"

	index++

	// The parent uuid
	entityPrototypeStr += "	" + prototypeVar + ".Indexs = append(" + prototypeVar + ".Indexs,\"parentUuid\")\n"
	entityPrototypeStr += "	" + prototypeVar + ".Fields = append(" + prototypeVar + ".Fields,\"parentUuid\")\n"
	entityPrototypeStr += "	" + prototypeVar + ".FieldsType = append(" + prototypeVar + ".FieldsType,\"xs.string\")\n"
	entityPrototypeStr += "	" + prototypeVar + ".FieldsOrder = append(" + prototypeVar + ".FieldsOrder," + strconv.Itoa(index) + ")\n"
	entityPrototypeStr += "	" + prototypeVar + ".FieldsVisibility = append(" + prototypeVar + ".FieldsVisibility,false)\n"

	index++

	// Generate the entity attributes.
	entityPrototypeStr += generateEntityAttributes(class, packageId, prototypeVar, &index)

	// Generate the entity associations...
	entityPrototypeStr += generateEntityAssociations(class, packageId, prototypeVar, &index)

	// The childs uuid
	entityPrototypeStr += "	" + prototypeVar + ".Fields = append(" + prototypeVar + ".Fields,\"childsUuid\")\n"
	entityPrototypeStr += "	" + prototypeVar + ".FieldsType = append(" + prototypeVar + ".FieldsType,\"[]xs.string\")\n"
	entityPrototypeStr += "	" + prototypeVar + ".FieldsOrder = append(" + prototypeVar + ".FieldsOrder," + strconv.Itoa(index) + ")\n"
	entityPrototypeStr += "	" + prototypeVar + ".FieldsVisibility = append(" + prototypeVar + ".FieldsVisibility,false)\n\n"

	index++

	// The referenced
	entityPrototypeStr += "	" + prototypeVar + ".Fields = append(" + prototypeVar + ".Fields,\"referenced\")\n"
	entityPrototypeStr += "	" + prototypeVar + ".FieldsType = append(" + prototypeVar + ".FieldsType,\"[]EntityRef\")\n"
	entityPrototypeStr += "	" + prototypeVar + ".FieldsOrder = append(" + prototypeVar + ".FieldsOrder," + strconv.Itoa(index) + ")\n"
	entityPrototypeStr += "	" + prototypeVar + ".FieldsVisibility = append(" + prototypeVar + ".FieldsVisibility,false)\n"

	entityPrototypeStr += "\n	store := GetServer().GetDataManager().getDataStore(" + packageId + "DB).(*KeyValueDataStore)\n"
	entityPrototypeStr += "	store.SetEntityPrototype(&" + prototypeVar + ")\n"

	entityPrototypeStr += "\n"

	entityPrototypeStr += "}\n\n"

	return entityPrototypeStr
}

func generateEntityQueryFields(class *XML_Schemas.CMOF_OwnedMember, packageId string) string {
	// The super classes code...
	classStr := ""

	superClasses := getSuperClasses(class.Name)
	for i := 0; i < len(superClasses); i++ {
		superClass := classesMap[superClasses[i]]
		classStr += "\n	/** members of " + superClass.Name + " **/\n"
		if len(superClass.Attributes) == 0 {
			classStr += "	/** No members **/\n"
		}
		for j := 0; j < len(superClass.Attributes); j++ {
			attribute := superClass.Attributes[j]
			classStr += "	query.Fields = append(query.Fields, \"M_" + attribute.Name + "\")\n"
		}
	}

	// Now the attributes...
	classStr += "\n	/** members of " + class.Name + " **/\n"
	if len(class.Attributes) == 0 {
		classStr += "	/** No members **/\n"
	}

	for i := 0; i < len(class.Attributes); i++ {
		attribute := class.Attributes[i]
		classStr += "	query.Fields = append(query.Fields, \"M_" + attribute.Name + "\")\n"
	}

	// Finaly the associtions
	associations := getClassAssociations(class)
	if len(associations) > 0 {
		classStr += "\n		/** associations of " + class.Name + " **/\n"
	}

	for _, association := range associations {
		//classStr += generateEntityAttribute(association, packageId, prototypeVar, index, true)
		classStr += "	query.Fields = append(query.Fields, \"M_" + association.Name + "Ptr\")\n"

	}
	classStr += "\n"

	return classStr
}

func generateEntityInitAttribute(class *XML_Schemas.CMOF_OwnedMember, attribute *XML_Schemas.CMOF_OwnedAttribute, packName string, prototypeName string, index *int, isAssociation bool) string {
	var memberStr string
	var typeName, isPrimitive, isPointer = getAttributeTypeName(attribute)

	var memberName = attribute.Name
	if isAssociation {
		memberName += "Ptr"
	}

	memberStr += "\n		/** " + memberName + " **/\n"
	memberTypeName := typeName

	if isPrimitive && memberTypeName != "interface{}" {
		memberStr += " 		if results[0][" + strconv.Itoa(*index) + "] != nil{\n"
		// Primitive are string, bool, int, float64 or interface{}
		if isPointer == true {
			memberStr += " 			this.object.M_" + memberName + "=results[0][" + strconv.Itoa(*index) + "].(*" + memberTypeName + ")\n"
		} else if attribute.Upper == "*" {
			memberStr += " 			this.object.M_" + memberName + "= append(this.object.M_" + memberName + ", results[0][" + strconv.Itoa(*index) + "].([]" + memberTypeName + ") ...)\n"
		} else {
			memberStr += " 			this.object.M_" + memberName + "=results[0][" + strconv.Itoa(*index) + "].(" + memberTypeName + ")\n"
		}

		memberStr += " 		}\n"

	} else if simpleTypesMap[typeName] != nil {
		// Here the type is an enum...
		enum := enumMap[typeName]
		memberStr += " 		if results[0][" + strconv.Itoa(*index) + "] != nil{\n"
		memberStr += " 			enumIndex := results[0][" + strconv.Itoa(*index) + "].(int)\n"
		for i := 0; i < len(enum.Litterals); i++ {

			if i == 0 {
				memberStr += "			if"
			} else {
				memberStr += "			} else if"
			}
			memberStr += " enumIndex == " + strconv.Itoa(i) + "{\n"

			// Now the code...
			litteral := enum.Litterals[i]
			enumValue := litteral.Name
			enumValue = strings.ToUpper(enumValue[0:1]) + enumValue[1:]

			// The litteral value...
			memberStr += " 				this.object.M_" + memberName + "=" + membersPackage[typeName] + "." + typeName + "_" + enumValue + "\n"

			if i == len(enum.Litterals)-1 {
				memberStr += " 			}\n"
			}
		}
		memberStr += " 		}\n"
	} else {
		// Here the type is another type...
		isRef := IsRef(attribute)
		memberStr += " 		if results[0][" + strconv.Itoa(*index) + "] != nil{\n"
		packName_ := packName
		if len(membersPackage[memberName]) > 0 {
			packName_ = membersPackage[memberName]
		}
		if attribute.Upper == "*" {
			if isRef || isAssociation {
				// Here the value is a reference...
				memberStr += "			idsStr :=results[0][" + strconv.Itoa(*index) + "].(string)\n"
				memberStr += "			ids :=make([]string,0)\n"
				memberStr += "			err := json.Unmarshal([]byte(idsStr), &ids)\n"
				memberStr += "			if err != nil {\n"
				memberStr += "				return err\n"
				memberStr += "			}\n"
				memberStr += "			for i:=0; i<len(ids); i++{\n"
				memberStr += "				if len(ids[i]) > 0 {\n"
				memberStr += "					refTypeName:=\"" + packName_ + "." + memberTypeName + "\"\n"
				memberStr += "					id_:= refTypeName + \"$$\" + ids[i]\n"
				memberStr += "					this.object.M_" + memberName + " = append(this.object.M_" + memberName + ",ids[i])\n"
				memberStr += "					GetServer().GetEntityManager().appendReference(\"" + memberName + "\",this.object.UUID, id_)\n"
				memberStr += "				}\n"
				memberStr += "			}\n"
			} else {
				// Here the type is not a reference and the value
				// represent a uuid...
				memberStr += "			uuidsStr :=results[0][" + strconv.Itoa(*index) + "].(string)\n"
				memberStr += "			uuids :=make([]string,0)\n"
				memberStr += "			err := json.Unmarshal([]byte(uuidsStr), &uuids)\n"
				memberStr += "			if err != nil {\n"
				memberStr += "				return err\n"
				memberStr += "			}\n"
				memberStr += "			for i:=0; i<len(uuids); i++{\n"
				// So here I will
				// Now I will initalyse the value...
				if Utility.Contains(abstractClassLst, typeName) {
					// Here I need to cast the type to it implementation type...
					memberStr += "			typeName := uuids[i][0:strings.Index(uuids[i], \"%\")]\n"
					memberStr += "			if err!=nil{\n"
					memberStr += "				log.Println(\"type \", typeName, \" not found!\")\n"
					memberStr += "				return err\n"
					memberStr += "			}\n"
					// Now I will iterate over implementation...
					implementations := getImplementationClasses(typeName)

					for i := 0; i < len(implementations); i++ {
						impl := ""
						if Utility.Contains(superClassesLst, implementations[i]) {
							impl = "_impl"
						}
						if i == 0 {
							memberStr += "				if typeName == \"" + packName + "." + implementations[i] + impl + "\"{\n"

						} else {
							memberStr += "				} else if typeName == \"" + packName + "." + implementations[i] + impl + "\"{\n"
						}
						elementPackName := membersPackage[implementations[i]]
						if Utility.Contains(superClassesLst, implementations[i]) {
							memberStr += "						/** abstract class **/\n"
							implementations_ := getImplementationClasses(implementations[i])
							for j := 0; j < len(implementations_); j++ {
								implementation := implementations_[j]
								if j == 0 {
									memberStr += "						if strings.HasPrefix(uuids[i], \"" + elementPackName + "." + implementation + "\"){\n"
								} else {
									memberStr += "						} else if strings.HasPrefix(uuids[i], \"" + elementPackName + "." + implementation + "\"){\n"
								}
								memberStr += "							if len(uuids[i]) > 0 {\n"
								memberStr += "								var " + memberName + "Entity *" + elementPackName + "_" + implementations[i] + "Entity\n"
								memberStr += "								if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {\n"
								memberStr += "									" + memberName + "Entity = instance.(*" + elementPackName + "_" + implementations[i] + "Entity)\n"
								memberStr += "								}else{\n"
								memberStr += "									" + memberName + "Entity = GetServer().GetEntityManager().New" + elementPackName + implementations[i] + "Entity(this.GetUuid(), uuids[i], nil)\n"
								memberStr += "									" + memberName + "Entity.InitEntity(uuids[i])\n"
								memberStr += "									GetServer().GetEntityManager().insert(" + memberName + "Entity)\n"
								memberStr += "								}\n"
								memberStr += "								" + memberName + "Entity.AppendReferenced(\"" + memberName + "\", this)\n"
								memberStr += "								this.AppendChild(\"" + memberName + "\"," + memberName + "Entity)\n"

								memberStr += "							}\n"
								if j == len(implementations_)-1 {
									memberStr += "				}\n"
								}
							}
						} else {
							memberStr += "						if len(uuids[i]) > 0 {\n"
							memberStr += "							var " + memberName + "Entity *" + elementPackName + "_" + implementations[i] + "Entity\n"
							memberStr += "							if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {\n"
							memberStr += "								" + memberName + "Entity = instance.(*" + elementPackName + "_" + implementations[i] + "Entity)\n"
							memberStr += "							}else{\n"
							memberStr += "								" + memberName + "Entity = GetServer().GetEntityManager().New" + elementPackName + implementations[i] + "Entity(this.GetUuid(), uuids[i], nil)\n"
							memberStr += "								" + memberName + "Entity.InitEntity(uuids[i])\n"
							memberStr += "								GetServer().GetEntityManager().insert(" + memberName + "Entity)\n"
							memberStr += "							}\n"

							memberStr += "							" + memberName + "Entity.AppendReferenced(\"" + memberName + "\", this)\n"
							memberStr += "							this.AppendChild(\"" + memberName + "\"," + memberName + "Entity)\n"

							memberStr += "						}\n"
						}

						if i == len(implementations)-1 {
							memberStr += "				}\n"
						}
					}
				} else {
					elementPackName := membersPackage[typeName]
					if Utility.Contains(superClassesLst, typeName) {
						memberStr += "				/** abstract class **/\n"
						implementations_ := getImplementationClasses(typeName)
						for j := 0; j < len(implementations_); j++ {
							implementation := implementations_[j]
							if j == 0 {
								memberStr += "				if strings.HasPrefix(uuids[i], \"" + elementPackName + "." + implementation + "\"){\n"
							} else {
								memberStr += "				} else if strings.HasPrefix(uuids[i], \"" + elementPackName + "." + implementation + "\"){\n"
							}
							memberStr += "					if len(uuids[i]) > 0 {\n"
							memberStr += "						var " + memberName + "Entity *" + elementPackName + "_" + implementation + "Entity\n"
							memberStr += "						if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {\n"
							memberStr += "						" + memberName + "Entity = instance.(*" + elementPackName + "_" + implementation + "Entity)\n"
							memberStr += "						}else{\n"
							memberStr += "							" + memberName + "Entity = GetServer().GetEntityManager().New" + elementPackName + implementation + "Entity(this.GetUuid(), uuids[i], nil)\n"
							memberStr += "							" + memberName + "Entity.InitEntity(uuids[i])\n"
							memberStr += "							GetServer().GetEntityManager().insert(" + memberName + "Entity)\n"
							memberStr += "						}\n"

							memberStr += "						" + memberName + "Entity.AppendReferenced(\"" + memberName + "\", this)\n"
							memberStr += "						this.AppendChild(\"" + memberName + "\"," + memberName + "Entity)\n"

							memberStr += "					}\n"
							if j == len(implementations_)-1 {
								memberStr += "				}\n"
							}
						}
					} else {
						memberStr += "				if len(uuids[i]) > 0 {\n"
						memberStr += "					var " + memberName + "Entity *" + elementPackName + "_" + typeName + "Entity\n"
						memberStr += "					if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {\n"
						memberStr += "						" + memberName + "Entity = instance.(*" + elementPackName + "_" + typeName + "Entity)\n"
						memberStr += "					}else{\n"
						memberStr += "						" + memberName + "Entity = GetServer().GetEntityManager().New" + elementPackName + typeName + "Entity(this.GetUuid(), uuids[i], nil)\n"
						memberStr += "						" + memberName + "Entity.InitEntity(uuids[i])\n"
						memberStr += "						GetServer().GetEntityManager().insert(" + memberName + "Entity)\n"
						memberStr += "					}\n"

						memberStr += "					" + memberName + "Entity.AppendReferenced(\"" + memberName + "\", this)\n"
						memberStr += "					this.AppendChild(\"" + memberName + "\"," + memberName + "Entity)\n"

						memberStr += "				}\n"
					}

				}
				memberStr += " 			}\n"
			}

		} else {
			if isRef || isAssociation {
				packName_ := packName
				if len(membersPackage[memberName]) > 0 {
					packName_ = membersPackage[memberName]
				}
				// Here the value is a refrerence...
				memberStr += "			id :=results[0][" + strconv.Itoa(*index) + "].(string)\n"
				memberStr += "			if len(id) > 0 {\n"
				memberStr += "				refTypeName:=\"" + packName_ + "." + memberTypeName + "\"\n"
				memberStr += "				id_:= refTypeName + \"$$\" + id\n"
				memberStr += "				this.object.M_" + memberName + "= id\n"
				memberStr += "				GetServer().GetEntityManager().appendReference(\"" + memberName + "\",this.object.UUID, id_)\n"
				if memberName == "bpmnElement" {
					memberStr += "				this.object.M_bpmnElement = id\n"
				}
				memberStr += "			}\n"
			} else {
				// Here the type is not a reference and the value
				// represent a uuid...
				memberStr += "			uuid :=results[0][" + strconv.Itoa(*index) + "].(string)\n"
				// Now I will initalyse the value...
				if Utility.Contains(abstractClassLst, typeName) {
					// Here I need to cast the type to it implementation type...
					memberStr += "			if len(uuid) > 0 {\n"
					memberStr += "				typeName := uuid[0:strings.Index(uuid, \"%\")]\n"
					memberStr += "				if err!=nil{\n"
					memberStr += "					log.Println(\"type \", typeName, \" not found!\")\n"
					memberStr += "					return err\n"
					memberStr += "				}\n"

					// Now I will iterate over implementation...
					implementations := getImplementationClasses(typeName)
					for i := 0; i < len(implementations); i++ {
						impl := ""
						if Utility.Contains(superClassesLst, implementations[i]) {
							impl = "_impl"
						}
						if i == 0 {
							memberStr += "			if typeName == \"" + packName + "." + implementations[i] + impl + "\"{\n"

						} else {
							memberStr += "			} else if typeName == \"" + packName + "." + implementations[i] + impl + "\"{\n"
						}
						elementPackName := membersPackage[implementations[i]]
						if Utility.Contains(superClassesLst, implementations[i]) {
							memberStr += "				/** abstract class **/\n"
							implementations_ := getImplementationClasses(implementations[i])
							for j := 0; j < len(implementations_); j++ {
								implementation := implementations_[j]
								if j == 0 {
									memberStr += "				if strings.HasPrefix(uuid, \"" + elementPackName + "." + implementation + "\"){\n"
								} else {
									memberStr += "				} else if strings.HasPrefix(uuid, \"" + elementPackName + "." + implementation + "\"){\n"
								}

								// So here I will create the good entity type...
								memberStr += "					if len(uuid) > 0 {\n"
								memberStr += "						var " + memberName + "Entity *" + elementPackName + "_" + implementation + "Entity\n"
								memberStr += "						if instance, ok := GetServer().GetEntityManager().contain(uuid); ok {\n"
								memberStr += "							" + memberName + "Entity = instance.(*" + elementPackName + "_" + implementation + "Entity)\n"
								memberStr += "						}else{\n"
								memberStr += "							" + memberName + "Entity = GetServer().GetEntityManager().New" + elementPackName + implementation + "Entity(this.GetUuid(), uuid, nil)\n"
								memberStr += "							" + memberName + "Entity.InitEntity(uuid)\n"
								memberStr += "							GetServer().GetEntityManager().insert(" + memberName + "Entity)\n"
								memberStr += "						}\n"

								memberStr += "						" + memberName + "Entity.AppendReferenced(\"" + memberName + "\", this)\n"
								memberStr += "						this.AppendChild(\"" + memberName + "\"," + memberName + "Entity)\n"

								memberStr += "					}\n"

								if j == len(implementations_)-1 {
									memberStr += "				}\n"
								}
							}
						} else {

							memberStr += "				if len(uuid) > 0 {\n"
							memberStr += "					var " + memberName + "Entity *" + elementPackName + "_" + implementations[i] + "Entity\n"
							memberStr += "					if instance, ok := GetServer().GetEntityManager().contain(uuid); ok {\n"
							memberStr += "						" + memberName + "Entity = instance.(*" + elementPackName + "_" + implementations[i] + "Entity)\n"
							memberStr += "					}else{\n"
							memberStr += "						" + memberName + "Entity = GetServer().GetEntityManager().New" + elementPackName + implementations[i] + "Entity(this.GetUuid(), uuid, nil)\n"
							memberStr += "						" + memberName + "Entity.InitEntity(uuid)\n"
							memberStr += "						GetServer().GetEntityManager().insert(" + memberName + "Entity)\n"
							memberStr += "					}\n"

							memberStr += "					" + memberName + "Entity.AppendReferenced(\"" + memberName + "\", this)\n"
							memberStr += "					this.AppendChild(\"" + memberName + "\"," + memberName + "Entity)\n"

							memberStr += "				}\n"

						}
						if i == len(implementations)-1 {
							memberStr += "			}\n"
						}
					}
					memberStr += "			}\n"

				} else {
					elementPackName := membersPackage[typeName]
					if Utility.Contains(superClassesLst, typeName) {
						memberStr += "				/** abstract class **/\n"
						implementations_ := getImplementationClasses(typeName)
						for j := 0; j < len(implementations_); j++ {
							implementation := implementations_[j]
							if j == 0 {
								memberStr += "			if strings.HasPrefix(uuid, \"" + elementPackName + "." + implementation + "\"){\n"
							} else {
								memberStr += "			} else if strings.HasPrefix(uuid, \"" + elementPackName + "." + implementation + "\"){\n"
							}
							memberStr += "			if len(uuid) > 0 {\n"
							memberStr += "				var " + memberName + "Entity *" + elementPackName + "_" + implementation + "Entity\n"
							memberStr += "				if instance, ok := GetServer().GetEntityManager().contain(uuid); ok {\n"
							memberStr += "					" + memberName + "Entity = instance.(*" + elementPackName + "_" + implementation + "Entity)\n"
							memberStr += "				}else{\n"
							memberStr += "					" + memberName + "Entity = GetServer().GetEntityManager().New" + elementPackName + implementation + "Entity(this.GetUuid(), uuid, nil)\n"
							memberStr += "					" + memberName + "Entity.InitEntity(uuid)\n"
							memberStr += "					GetServer().GetEntityManager().insert(" + memberName + "Entity)\n"
							memberStr += "				}\n"

							memberStr += "				" + memberName + "Entity.AppendReferenced(\"" + memberName + "\", this)\n"
							memberStr += "				this.AppendChild(\"" + memberName + "\"," + memberName + "Entity)\n"

							memberStr += "			}\n"
							if j == len(implementations_)-1 {
								memberStr += "				}\n"
							}
						}
					} else {
						memberStr += "			if len(uuid) > 0 {\n"
						memberStr += "				var " + memberName + "Entity *" + elementPackName + "_" + typeName + "Entity\n"
						memberStr += "				if instance, ok := GetServer().GetEntityManager().contain(uuid); ok {\n"
						memberStr += "					" + memberName + "Entity = instance.(*" + elementPackName + "_" + typeName + "Entity)\n"
						memberStr += "				}else{\n"
						memberStr += "					" + memberName + "Entity = GetServer().GetEntityManager().New" + elementPackName + typeName + "Entity(this.GetUuid(), uuid, nil)\n"
						memberStr += "					" + memberName + "Entity.InitEntity(uuid)\n"
						memberStr += "					GetServer().GetEntityManager().insert( " + memberName + "Entity)\n"
						memberStr += "				}\n"

						memberStr += "				" + memberName + "Entity.AppendReferenced(\"" + memberName + "\", this)\n"
						memberStr += "				this.AppendChild(\"" + memberName + "\"," + memberName + "Entity)\n"
						memberStr += "			}\n"
					}
				}
			}
		}
		memberStr += " 		}\n"
	}

	// Increment the index...
	*index++

	return memberStr
}

/**
 * Generate the entity initilyzation function of a given class.
 */
func generateEntityInitFunc(packageId string, class *XML_Schemas.CMOF_OwnedMember) string {
	impl := ""
	if Utility.Contains(superClassesLst, class.Name) {
		impl = "_impl"
	}

	entityInitStr := "/** Read **/\n"
	entityInitStr += "func (this *" + packageId + "_" + class.Name + "Entity) InitEntity(id string) error{\n"

	// If the value is already in the cache I have nothing todo...
	entityInitStr += "	if this.object.IsInit == true {\n"
	entityInitStr += "		entity, err := GetServer().GetEntityManager().getEntityByUuid(id)\n"
	entityInitStr += "		if err == nil {\n"
	entityInitStr += "			// Return the already initialyse entity.\n"
	entityInitStr += "			this = entity.(*" + packageId + "_" + class.Name + "Entity)\n"
	entityInitStr += "			return nil\n"
	entityInitStr += "		}\n"
	entityInitStr += "		// I must reinit the entity if the entity manager dosent have it.\n"
	entityInitStr += "		this.object.IsInit = false\n"
	entityInitStr += "	}\n"

	// Here If the entity already exist I will it...
	entityInitStr += "	this.uuid = id\n"

	entityInitStr += "\n	// Set the reference on the map\n"

	// First of all I will generate the query...
	entityInitStr += "	var query EntityQuery\n"
	entityInitStr += "	query.TypeName = \"" + packageId + "." + class.Name + "\"\n\n"
	entityInitStr += "	query.Fields = append(query.Fields, \"uuid\")\n"
	entityInitStr += "	query.Fields = append(query.Fields, \"parentUuid\")\n"

	// Now the fields...
	entityInitStr += generateEntityQueryFields(class, packageId)

	// The child uuids...
	entityInitStr += "	query.Fields = append(query.Fields, \"childsUuid\")\n"

	// The entities that make use of this entity...
	entityInitStr += "	query.Fields = append(query.Fields, \"referenced\")\n"

	// The indexation...
	entityInitStr += "	query.Indexs = append(query.Indexs, \"uuid=\"+this.uuid)\n"

	// In the proto...
	entityInitStr += "\n	var fieldsType []interface{} // not use...\n"

	// The filters...
	entityInitStr += "	var params []interface{}\n"

	entityInitStr += "	var results [][]interface{}\n"
	entityInitStr += "	var err error\n"

	// Now I will create the query string...
	entityInitStr += "	queryStr, _ := json.Marshal(query)\n\n"

	entityInitStr += "	results, err = GetServer().GetDataManager().readData(" + packageId + "DB, string(queryStr), fieldsType, params)\n"
	entityInitStr += "	if err != nil {\n"
	entityInitStr += "		return err\n"
	entityInitStr += "	}\n"

	entityInitStr += "	// Initialisation of information of " + class.Name + "...\n"
	entityInitStr += "	if len(results) > 0 {\n"

	entityInitStr += "\n	/** initialyzation of the entity object **/\n"
	entityInitStr += "		this.object = new(" + packageId + "." + class.Name + impl + ")\n"
	entityInitStr += "		this.object.UUID = this.uuid\n"
	entityInitStr += "		this.object.TYPENAME = \"" + packageId + "." + class.Name + "\"\n\n"

	// Initialisation here...
	//entityInitStr += "		log.Println(\"initialyzation of " + class.Name + " uuid: \", results[0][0])\n"

	entityInitStr += "		this.parentUuid = results[0][1].(string)\n"

	// 0:uuid 1:parentUuid
	index := 2

	prototypeVar := strings.ToLower(class.Name[0:1]) + class.Name[1:] + "EntityProto"

	// Now the attributes...
	superClasses := getSuperClasses(class.Name)
	for i := 0; i < len(superClasses); i++ {
		superClass := classesMap[superClasses[i]]
		entityInitStr += "\n		/** members of " + superClass.Name + " **/\n"
		if len(superClass.Attributes) == 0 {
			entityInitStr += "		/** No members **/\n"
		}
		for j := 0; j < len(superClass.Attributes); j++ {
			attribute := superClass.Attributes[j]
			entityInitStr += generateEntityInitAttribute(class, &attribute, packageId, prototypeVar, &index, false)
		}
	}

	entityInitStr += "\n		/** members of " + class.Name + " **/\n"
	if len(class.Attributes) == 0 {
		entityInitStr += "		/** No members **/\n"
	}

	for i := 0; i < len(class.Attributes); i++ {
		attribute := class.Attributes[i]
		entityInitStr += generateEntityInitAttribute(class, &attribute, packageId, prototypeVar, &index, false)
	}

	// Now the assiciation initialisation...
	associations := getClassAssociations(class)
	if len(associations) > 0 {
		entityInitStr += "\n		/** associations of " + class.Name + " **/\n"
	}

	for _, association := range associations {
		entityInitStr += generateEntityInitAttribute(class, association, packageId, prototypeVar, &index, true)
	}

	// The childs uuid's
	entityInitStr += "		childsUuidStr := results[0][" + strconv.Itoa(index) + "].(string)\n"
	entityInitStr += "		this.childsUuid = make([]string, 0)\n"
	entityInitStr += "		err := json.Unmarshal([]byte(childsUuidStr), &this.childsUuid)\n"
	entityInitStr += "		if err != nil {\n"
	entityInitStr += "			return err\n"
	entityInitStr += "		}\n\n"

	index++

	// The referenced...
	entityInitStr += "		referencedStr := results[0][" + strconv.Itoa(index) + "].(string)\n"
	entityInitStr += "		this.referenced = make([]EntityRef, 0)\n"
	entityInitStr += "		err = json.Unmarshal([]byte(referencedStr), &this.referenced)\n"
	entityInitStr += "		if err != nil {\n"
	entityInitStr += "			return err\n"
	entityInitStr += "		}\n"

	entityInitStr += "	}\n\n"

	entityInitStr += "	// set need save to false.\n"
	entityInitStr += "	this.SetNeedSave(false)\n"
	entityInitStr += "	// set init done.\n"
	entityInitStr += "	this.SetInit(true)\n"

	entityInitStr += "	// Init the references...\n"
	entityInitStr += "	GetServer().GetEntityManager().InitEntity(this)\n"

	entityInitStr += "	return nil\n"

	entityInitStr += "}\n\n"
	return entityInitStr
}

func generateEntitySaveEntityInfo(class *XML_Schemas.CMOF_OwnedMember, attribute *XML_Schemas.CMOF_OwnedAttribute, packageId string, isAssociation bool) string {
	entityInfoStr := ""

	var typeName, isPrimitive, _ = getAttributeTypeName(attribute)

	ref := ""
	if isAssociation {
		ref = "Ptr"
	}

	isRef := IsRef(attribute)

	if isPrimitive {
		entityInfoStr += "	" + class.Name + "Info = append(" + class.Name + "Info, this.object.M_" + attribute.Name + ")\n"
	} else {
		// Here the field is a reference to other entity so only the id must be save...
		entityInfoStr += "\n	/** Save " + attribute.Name + " type " + typeName + " **/\n"
		if attribute.Upper == "*" {
			// Here the reference is an array...
			if !isRef {
				entityInfoStr += "	" + attribute.Name + "Ids := make([]string,0)\n"
				entityInfoStr += "	for i := 0; i < len(this.object.M_" + attribute.Name + "); i++ {\n"
				if Utility.Contains(abstractClassLst, typeName) || Utility.Contains(superClassesLst, typeName) {
					// So here I will test the type and save appropriately...
					entityInfoStr += "		switch v := this.object.M_" + attribute.Name + "[i].(type) {\n"
					implementations := getImplementationClasses(typeName)
					for i := 0; i < len(implementations); i++ {
						elementPackName := membersPackage[implementations[i]]
						impl := ""
						if Utility.Contains(superClassesLst, implementations[i]) {
							impl = "_impl"
						}
						entityInfoStr += "		case *" + elementPackName + "." + implementations[i] + impl + ":\n"
						entityInfoStr += "		" + attribute.Name + "Entity:= GetServer().GetEntityManager().New" + elementPackName + implementations[i] + "Entity(this.GetUuid(), v.UUID, v)\n"
						entityInfoStr += "		" + attribute.Name + "Ids=append(" + attribute.Name + "Ids," + attribute.Name + "Entity.uuid)\n"
						entityInfoStr += "		" + attribute.Name + "Entity.AppendReferenced(\"" + attribute.Name + "\", this)\n"
						entityInfoStr += "		this.AppendChild(\"" + attribute.Name + "\"," + attribute.Name + "Entity)\n"
						entityInfoStr += "		if " + attribute.Name + "Entity.NeedSave() {\n"
						entityInfoStr += "			" + attribute.Name + "Entity.SaveEntity()\n"
						entityInfoStr += "		}\n"

					}
					entityInfoStr += "		}\n"
				} else {
					elementPackName := membersPackage[typeName]
					entityInfoStr += "		" + attribute.Name + "Entity:= GetServer().GetEntityManager().New" + elementPackName + typeName + "Entity(this.GetUuid(), this.object.M_" + attribute.Name + "[i].UUID,this.object.M_" + attribute.Name + "[i])\n"
					entityInfoStr += "		" + attribute.Name + "Ids=append(" + attribute.Name + "Ids," + attribute.Name + "Entity.uuid)\n"
					entityInfoStr += "		" + attribute.Name + "Entity.AppendReferenced(\"" + attribute.Name + "\", this)\n"
					entityInfoStr += "		this.AppendChild(\"" + attribute.Name + "\"," + attribute.Name + "Entity)\n"
					entityInfoStr += "		if " + attribute.Name + "Entity.NeedSave() {\n"
					entityInfoStr += "			" + attribute.Name + "Entity.SaveEntity()\n"
					entityInfoStr += "		}\n"
				}
				entityInfoStr += "	}\n"
				// Now I will save the array of id's...
				entityInfoStr += "	" + attribute.Name + "Str, _ := json.Marshal(" + attribute.Name + "Ids)\n"
				entityInfoStr += "	" + class.Name + "Info = append(" + class.Name + "Info, string(" + attribute.Name + "Str))\n"
			} else {
				// The attribute is a reference.
				entityInfoStr += "	" + attribute.Name + ref + "Str, _ := json.Marshal(this.object.M_" + attribute.Name + ref + ")\n"
				entityInfoStr += "	" + class.Name + "Info = append(" + class.Name + "Info, string(" + attribute.Name + ref + "Str))\n"
			}

		} else {

			if simpleTypesMap[typeName] != nil {
				// This is an enum...
				enum := enumMap[typeName]
				elementPackName := membersPackage[typeName]
				for i := 0; i < len(enum.Litterals); i++ {
					litteral := enum.Litterals[i]
					enumValue := litteral.Name
					enumValue = strings.ToUpper(enumValue[0:1]) + enumValue[1:]
					if i == 0 {
						entityInfoStr += "	if "
					} else {
						entityInfoStr += "	} else if "
					}
					entityInfoStr += "this.object.M_" + attribute.Name + "==" + elementPackName + "." + typeName + "_" + enumValue + "{\n"
					// Now here I will set the numerical value of the enum...
					entityInfoStr += "		" + class.Name + "Info = append(" + class.Name + "Info, " + strconv.Itoa(i) + ")\n"
				}
				// Now the default value...
				entityInfoStr += "	}else{\n"
				entityInfoStr += "		" + class.Name + "Info = append(" + class.Name + "Info, 0)\n"
				entityInfoStr += "	}\n"

			} else {
				if !isRef {
					entityInfoStr += "	if this.object.M_" + attribute.Name + " != nil {\n"
					if Utility.Contains(abstractClassLst, typeName) || Utility.Contains(superClassesLst, typeName) {
						// So here I will test the type and save appropriately...
						entityInfoStr += "		switch v := this.object.M_" + attribute.Name + ".(type) {\n"
						implementations := getImplementationClasses(typeName)
						for i := 0; i < len(implementations); i++ {
							elementPackName := membersPackage[implementations[i]]
							impl := ""
							if Utility.Contains(superClassesLst, implementations[i]) {
								impl = "_impl"
							}

							entityInfoStr += "		case *" + elementPackName + "." + implementations[i] + impl + ":\n"
							entityInfoStr += "			" + attribute.Name + "Entity:= GetServer().GetEntityManager().New" + elementPackName + implementations[i] + "Entity(this.GetUuid(), v.UUID, v)\n"
							entityInfoStr += "			" + class.Name + "Info = append(" + class.Name + "Info, " + attribute.Name + "Entity.uuid)\n"
							entityInfoStr += "		    " + attribute.Name + "Entity.AppendReferenced(\"" + attribute.Name + "\", this)\n"
							entityInfoStr += "			this.AppendChild(\"" + attribute.Name + "\"," + attribute.Name + "Entity)\n"
							entityInfoStr += "			if " + attribute.Name + "Entity.NeedSave() {\n"
							entityInfoStr += "				" + attribute.Name + "Entity.SaveEntity()\n"
							entityInfoStr += "			}\n"
						}
						entityInfoStr += "			}\n"

					} else {
						elementPackName := membersPackage[typeName]
						entityInfoStr += "		" + attribute.Name + "Entity:= GetServer().GetEntityManager().New" + elementPackName + typeName + "Entity(this.GetUuid(), this.object.M_" + attribute.Name + ".UUID, this.object.M_" + attribute.Name + ")\n"
						entityInfoStr += "		" + class.Name + "Info = append(" + class.Name + "Info, " + attribute.Name + "Entity.uuid)\n"
						entityInfoStr += "		" + attribute.Name + "Entity.AppendReferenced(\"" + attribute.Name + "\", this)\n"
						entityInfoStr += "		this.AppendChild(\"" + attribute.Name + "\"," + attribute.Name + "Entity)\n"
						entityInfoStr += "		if " + attribute.Name + "Entity.NeedSave() {\n"
						entityInfoStr += "			" + attribute.Name + "Entity.SaveEntity()\n"
						entityInfoStr += "		}\n"

					}
					entityInfoStr += "	}else{\n"
					entityInfoStr += "		" + class.Name + "Info = append(" + class.Name + "Info, \"\")\n"
					entityInfoStr += "	}\n"
				} else {
					// The attribute is a reference.
					if hasId(classesMap[typeName]) {
						entityInfoStr += "		" + class.Name + "Info = append(" + class.Name + "Info,this.object.M_" + attribute.Name + ref + ")\n"
					} else {
						//log.Println("------------> attribute ", typeName, " has no method GetId, must be an error here...")
						entityInfoStr += "	/** attribute " + typeName + " has no method GetId, must be an error here...*/\n"
						entityInfoStr += "	" + class.Name + "Info = append(" + class.Name + "Info,\"\")\n"
					}
				}
			}
		}
	}
	return entityInfoStr
}

/**
 * Generate the entity creation function of a given class.
 */
func generateEntitySaveFunc(packageId string, class *XML_Schemas.CMOF_OwnedMember) string {
	entitySaveStr := "/** Create **/\n"
	entitySaveStr += "func (this *" + packageId + "_" + class.Name + "Entity) SaveEntity() {\n"
	entitySaveStr += "	if this.object.NeedSave == false {\n"
	entitySaveStr += "		return\n"
	entitySaveStr += "	}\n\n"

	entitySaveStr += "	this.SetNeedSave(false)\n"
	entitySaveStr += "	this.SetInit(true)\n"

	// So here I will save the entity data...
	entitySaveStr += "	this.object.UUID = this.uuid\n"
	entitySaveStr += "	this.object.TYPENAME = \"" + packageId + "." + class.Name + "\"\n\n"

	// Here i will create the user information in the database...
	entitySaveStr += "	var query EntityQuery\n"
	entitySaveStr += "	query.TypeName = \"" + packageId + "." + class.Name + "\"\n\n"

	entitySaveStr += "	query.Fields = append(query.Fields, \"uuid\")\n"
	entitySaveStr += "	query.Fields = append(query.Fields, \"parentUuid\")\n"

	// Now the fields...
	entitySaveStr += generateEntityQueryFields(class, packageId)

	// Now the child uuid...
	entitySaveStr += "	query.Fields = append(query.Fields, \"childsUuid\")\n"

	// referenced
	entitySaveStr += "	query.Fields = append(query.Fields, \"referenced\")\n"

	// I will create an array and put the value inside it...
	entitySaveStr += "	var " + class.Name + "Info []interface{}\n\n"

	// Save the entity uuid
	entitySaveStr += "	" + class.Name + "Info = append(" + class.Name + "Info, this.GetUuid())\n"

	// Save it parent uuid
	entitySaveStr += "	if this.parentPtr != nil {\n"
	entitySaveStr += "		" + class.Name + "Info = append(" + class.Name + "Info, this.parentPtr.GetUuid())\n"
	entitySaveStr += "	}else{\n"
	entitySaveStr += "		" + class.Name + "Info = append(" + class.Name + "Info, \"\")\n"
	entitySaveStr += "	}\n"

	// Here I will map the value to save into the database.
	superClasses := getSuperClasses(class.Name)
	for i := 0; i < len(superClasses); i++ {
		superClass := classesMap[superClasses[i]]
		entitySaveStr += "\n	/** members of " + superClass.Name + " **/\n"
		if len(superClass.Attributes) == 0 {
			entitySaveStr += "	/** No members **/\n"
		}
		for j := 0; j < len(superClass.Attributes); j++ {
			attribute := superClass.Attributes[j]
			entitySaveStr += generateEntitySaveEntityInfo(class, &attribute, packageId, false)
		}
	}

	// Now the attributes...
	entitySaveStr += "\n	/** members of " + class.Name + " **/\n"
	if len(class.Attributes) == 0 {
		entitySaveStr += "	/** No members **/\n"
	}

	for i := 0; i < len(class.Attributes); i++ {
		attribute := class.Attributes[i]
		entitySaveStr += generateEntitySaveEntityInfo(class, &attribute, packageId, false)
	}

	// Now the assiciation initialisation...
	associations := getClassAssociations(class)
	if len(associations) > 0 {
		entitySaveStr += "\n	/** associations of " + class.Name + " **/\n"
	}

	for _, association := range associations {
		entitySaveStr += generateEntitySaveEntityInfo(class, association, packageId, true)
	}

	// Save the childs uuid...
	entitySaveStr += "	childsUuidStr, _ := json.Marshal(this.childsUuid)\n"
	entitySaveStr += "	" + class.Name + "Info = append(" + class.Name + "Info, string(childsUuidStr))\n"

	// Save the referenced
	entitySaveStr += "	referencedStr, _ := json.Marshal(this.referenced)\n"
	entitySaveStr += "	" + class.Name + "Info = append(" + class.Name + "Info, string(referencedStr))\n"

	// The event data...
	entitySaveStr += "	eventData := make([]*MessageData, 1)\n"
	entitySaveStr += "	msgData := new(MessageData)\n"
	entitySaveStr += "	msgData.Name = \"entity\"\n"
	entitySaveStr += "	msgData.Value = this.GetObject()\n"
	entitySaveStr += "	eventData[0] = msgData\n"
	entitySaveStr += "	var err error\n"

	entitySaveStr += "	var evt *Event\n"
	entitySaveStr += "	if this.Exist() == true {\n"
	entitySaveStr += "		evt, _ = NewEvent(UpdateEntityEvent, EntityEvent, eventData)\n"
	entitySaveStr += "		var params []interface{}\n"
	entitySaveStr += "		query.Indexs = append(query.Indexs, \"uuid=\"+this.uuid)\n"
	entitySaveStr += "		queryStr, _ := json.Marshal(query)\n"
	entitySaveStr += "		err = GetServer().GetDataManager().updateData(" + packageId + "DB, string(queryStr), " + class.Name + "Info, params)\n"
	entitySaveStr += "	} else {\n"
	entitySaveStr += "		evt, _ = NewEvent(NewEntityEvent, EntityEvent, eventData)\n"
	entitySaveStr += "		queryStr, _ := json.Marshal(query)\n"
	entitySaveStr += "		_, err =  GetServer().GetDataManager().createData(" + packageId + "DB, string(queryStr), " + class.Name + "Info)\n"
	entitySaveStr += "	}\n"

	entitySaveStr += "	if err == nil {\n"
	entitySaveStr += "		GetServer().GetEntityManager().insert(this)\n"
	entitySaveStr += "		GetServer().GetEntityManager().setReferences(this)\n"
	entitySaveStr += "		GetServer().GetEventManager().BroadcastEvent(evt)\n"
	entitySaveStr += "	}\n"

	entitySaveStr += "}\n\n"
	return entitySaveStr
}

/**
 * Generate the entity delete function of a given class.
 */
func generateEntityDeleteFunc(packageId string, class *XML_Schemas.CMOF_OwnedMember) string {
	entityDeleteStr := "/** Delete **/\n"
	entityDeleteStr += "func (this *" + packageId + "_" + class.Name + "Entity) DeleteEntity() {\n"
	entityDeleteStr += "	GetServer().GetEntityManager().deleteEntity(this)\n"
	entityDeleteStr += "}\n\n"
	return entityDeleteStr
}

/**
 * Generate the entity delete function of a given class.
 */
func generateEntityExistsFunc(packageId string, class *XML_Schemas.CMOF_OwnedMember) string {
	entityExistStr := "/** Exists **/\n"
	entityExistStr += "func " + packageId + strings.ToUpper(class.Name[0:1]) + class.Name[1:] + "Exists(val string) string {\n"

	entityExistStr += "	var query EntityQuery\n"
	entityExistStr += "	query.TypeName = \"" + packageId + "." + class.Name + "\"\n"

	// If the class has id, or has name
	if hasId(class) {
		entityExistStr += "	query.Indexs = append(query.Indexs, \"M_id=\"+val)\n"
	} else if hasName(class) {
		entityExistStr += "	query.Indexs = append(query.Indexs, \"M_name=\"+val)\n"
	}
	entityExistStr += "	query.Fields = append(query.Fields, \"uuid\")\n"
	entityExistStr += "	var fieldsType []interface {} // not use...\n"
	entityExistStr += "	var params []interface{}\n"

	entityExistStr += "	queryStr, _ := json.Marshal(query)\n"

	entityExistStr += "	results, err := GetServer().GetDataManager().readData(" + packageId + "DB, string(queryStr), fieldsType, params)\n"

	entityExistStr += "	if err != nil || len(results) == 0 {\n"
	entityExistStr += "		return \"\"\n"
	entityExistStr += "	}\n"

	entityExistStr += "	return results[0][0].(string)\n"

	entityExistStr += "}\n\n"

	return entityExistStr
}
