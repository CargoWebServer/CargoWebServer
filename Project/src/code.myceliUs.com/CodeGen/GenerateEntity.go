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
	classStr += "import (\n"
	classStr += "\"code.myceliUs.com/CargoWebServer/Cargo/Entities/" + packageId + "\"\n"
	classStr += "\"code.myceliUs.com/Utility\"\n"
	classStr += ")\n\n"

	classes := packagesMembers[packageId]

	for i := 0; i < len(classes); i++ {
		class := classesMap[classes[i]]
		if class != nil {
			// Entity prototype...
			classStr += generateEntityPrototypeFunc(packageId, class)
		}
	}

	classStr += generateRegisterObjects(packageId)
	classStr += generateCreatePrototypes(packageId)

	// Wrote the file...
	if len(classStr) > 0 {
		WriteClassFile("code.myceliUs.com/CargoWebServer/Cargo/Server/", "", "Entity_"+packageId, packageStr+"\n"+classStr)
	}
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
		memberStr += "	" + prototypeName + ".FieldsDefaultValue = append(" + prototypeName + ".FieldsDefaultValue,\"undefined\")\n"
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

		// int64 can be a time, a date or a long.
		if memberTypeName == "int64" {
			if attribute.AssociationTypeRef.Ref == "http://www.w3.org/2001/XMLSchema#date" {
				memberTypeName = "date"
			} else if attribute.AssociationTypeRef.Ref == "http://www.w3.org/2001/XMLSchema#time" {
				memberTypeName = "time"
			} else {
				memberTypeName = "long"
			}
		}

		memberTypeName = "xs." + memberTypeName

		if attribute.Upper == "*" {
			memberTypeName = "[]" + memberTypeName
		}

		memberStr += "	" + prototypeName + ".FieldsType = append(" + prototypeName + ".FieldsType,\"" + memberTypeName + "\")\n"

		// Set the default attribute value here.
		if attribute.Upper == "*" {
			memberStr += "	" + prototypeName + ".FieldsDefaultValue = append(" + prototypeName + ".FieldsDefaultValue,\"[]\")\n"
		} else if XML_Schemas.IsXsString(memberTypeName) || XML_Schemas.IsXsId(memberTypeName) || XML_Schemas.IsXsRef(memberTypeName) {
			memberStr += "	" + prototypeName + ".FieldsDefaultValue = append(" + prototypeName + ".FieldsDefaultValue,\"\")\n"
		} else if XML_Schemas.IsXsInt(memberTypeName) || XML_Schemas.IsXsTime(memberTypeName) {
			memberStr += "	" + prototypeName + ".FieldsDefaultValue = append(" + prototypeName + ".FieldsDefaultValue,\"0\")\n"
		} else if XML_Schemas.IsXsNumeric(memberTypeName) {
			memberStr += "	" + prototypeName + ".FieldsDefaultValue = append(" + prototypeName + ".FieldsDefaultValue,\"0.0\")\n"
		} else if XML_Schemas.IsXsDate(memberTypeName) {
			memberStr += "	" + prototypeName + ".FieldsDefaultValue = append(" + prototypeName + ".FieldsDefaultValue,\"new Date()\")\n"
		} else if XML_Schemas.IsXsBoolean(memberTypeName) {
			memberStr += "	" + prototypeName + ".FieldsDefaultValue = append(" + prototypeName + ".FieldsDefaultValue,\"false\")\n"
		} else {
			// Object here.
			memberStr += "	" + prototypeName + ".FieldsDefaultValue = append(" + prototypeName + ".FieldsDefaultValue,\"undefined\")\n"
		}

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
			if i == 0 {
				//memberStr += "	" + prototypeName + ".FieldsDefaultValue = append(" + prototypeName + ".FieldsDefaultValue,\"" + typeName + "_" + enumValue + "\")\n"
				memberStr += "	" + prototypeName + ".FieldsDefaultValue = append(" + prototypeName + ".FieldsDefaultValue,\"" + strconv.Itoa(i+1) + "\")\n"
			}
		}
		memberStr += "	" + prototypeName + ".FieldsType = append(" + prototypeName + ".FieldsType,\"" + enumStr + "\")\n"

	} else {
		// Here the type is another type...
		if attribute.Upper == "*" {
			memberTypeName = "[]" + packName + "." + memberTypeName
			memberStr += "	" + prototypeName + ".FieldsDefaultValue = append(" + prototypeName + ".FieldsDefaultValue,\"[]\")\n"
		} else {
			memberTypeName = packName + "." + memberTypeName
			memberStr += "	" + prototypeName + ".FieldsDefaultValue = append(" + prototypeName + ".FieldsDefaultValue,\"undefined\")\n"
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
	entityPrototypeStr += "	" + prototypeVar + ".Ids = append(" + prototypeVar + ".Ids,\"UUID\")\n"
	entityPrototypeStr += "	" + prototypeVar + ".Fields = append(" + prototypeVar + ".Fields,\"UUID\")\n"
	entityPrototypeStr += "	" + prototypeVar + ".FieldsType = append(" + prototypeVar + ".FieldsType,\"xs.string\")\n"
	entityPrototypeStr += "	" + prototypeVar + ".FieldsOrder = append(" + prototypeVar + ".FieldsOrder," + strconv.Itoa(index) + ")\n"
	entityPrototypeStr += "	" + prototypeVar + ".FieldsVisibility = append(" + prototypeVar + ".FieldsVisibility,false)\n"
	entityPrototypeStr += "	" + prototypeVar + ".FieldsDefaultValue = append(" + prototypeVar + ".FieldsDefaultValue,\"\")\n"

	index++

	// The parent uuid
	entityPrototypeStr += "	" + prototypeVar + ".Indexs = append(" + prototypeVar + ".Indexs,\"ParentUuid\")\n"
	entityPrototypeStr += "	" + prototypeVar + ".Fields = append(" + prototypeVar + ".Fields,\"ParentUuid\")\n"
	entityPrototypeStr += "	" + prototypeVar + ".FieldsType = append(" + prototypeVar + ".FieldsType,\"xs.string\")\n"
	entityPrototypeStr += "	" + prototypeVar + ".FieldsOrder = append(" + prototypeVar + ".FieldsOrder," + strconv.Itoa(index) + ")\n"
	entityPrototypeStr += "	" + prototypeVar + ".FieldsVisibility = append(" + prototypeVar + ".FieldsVisibility,false)\n"
	entityPrototypeStr += "	" + prototypeVar + ".FieldsDefaultValue = append(" + prototypeVar + ".FieldsDefaultValue,\"\")\n"

	index++

	entityPrototypeStr += "	" + prototypeVar + ".Fields = append(" + prototypeVar + ".Fields,\"ParentLnk\")\n"
	entityPrototypeStr += "	" + prototypeVar + ".FieldsType = append(" + prototypeVar + ".FieldsType,\"xs.string\")\n"
	entityPrototypeStr += "	" + prototypeVar + ".FieldsOrder = append(" + prototypeVar + ".FieldsOrder," + strconv.Itoa(index) + ")\n"
	entityPrototypeStr += "	" + prototypeVar + ".FieldsVisibility = append(" + prototypeVar + ".FieldsVisibility,false)\n"
	entityPrototypeStr += "	" + prototypeVar + ".FieldsDefaultValue = append(" + prototypeVar + ".FieldsDefaultValue,\"\")\n"

	index++

	// Generate the entity attributes.
	entityPrototypeStr += generateEntityAttributes(class, packageId, prototypeVar, &index)

	// Generate the entity associations...
	entityPrototypeStr += generateEntityAssociations(class, packageId, prototypeVar, &index)

	entityPrototypeStr += "\n	store := GetServer().GetDataManager().getDataStore(" + packageId + "DB).(*GraphStore)\n"
	entityPrototypeStr += "	store.CreateEntityPrototype(&" + prototypeVar + ")\n"

	entityPrototypeStr += "\n"

	entityPrototypeStr += "}\n\n"

	return entityPrototypeStr
}
