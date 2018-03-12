package main

import (
	"log"
	"strconv"
	"strings"

	"code.myceliUs.com/Utility"
	"code.myceliUs.com/XML_Schemas"
)

/////////////////////////////////////////////////////////////////////////////////
// XSD -> GO
/////////////////////////////////////////////////////////////////////////////////

func generateGoXmlParser(class *XML_Schemas.CMOF_OwnedMember, packName string) string {
	xmlParserStr := ""
	// I will retreive the releated xsd element if there is one...
	element := elementsTypeMap[class.Name]
	if element != nil {
		// Create the xml parser here...
		xmlParserStr += "type Xsd" + class.Name + " struct {\n"
		if element.Name == "Point" {
			// Exception here...
			xmlParserStr += "	XMLName xml.Name	`xml:\"waypoint\"`\n"
		} else {
			xmlParserStr += "	XMLName xml.Name	`xml:\"" + element.Name + "\"`\n"
		}

		superClasses := getSuperClasses(class.Name)
		for i := 0; i < len(superClasses); i++ {
			parentElement := elementsTypeMap[superClasses[i]]
			if parentElement != nil {
				xmlParserStr += "	/** " + superClasses[i] + " **/\n"
				xmlParserContentStr := generateGoXmlParserContent(classesMap[superClasses[i]], parentElement, packName)
				if len(xmlParserContentStr) == 0 {
					xmlParserStr += "	/** Nothing **/\n"
				} else {
					xmlParserStr += generateGoXmlParserContent(classesMap[superClasses[i]], parentElement, packName) + "\n"
				}
			}
		}

		xmlParserStr += generateGoXmlParserContent(class, element, packName)
		xmlParserStr += "}\n"
	} else if simpleTypesMap[class.Name] != nil {
		// So here I will create the parser for the enum type...
		xmlParserStr += generateGoXmlEnumParser(class, packName)
	}

	// here before return i will generate the alias parser...
	var aliasParserStr string
	for alias, typeName := range aliasElementType {
		if typeName == class.Name {
			// Here the xml parser must be copied with the name...
			aliasParserStr_ := xmlParserStr
			aliasParserStr_ = strings.Replace(aliasParserStr_, `type Xsd`+typeName, `type Xsd`+strings.ToUpper(alias[0:1])+alias[1:], -1)
			aliasParserStr_ = strings.Replace(aliasParserStr_, `"`+strings.ToLower(typeName[0:1])+typeName[1:]+`"`, `"`+alias+`"`, -1)
			aliasParserStr += "\n" + aliasParserStr_
		}
	}
	if len(aliasParserStr) > 0 {
		xmlParserStr += "/** Alias Xsd parser **/\n\n "
		xmlParserStr += aliasParserStr
	}

	return xmlParserStr
}

func generateGoXmlEnumParser(member *XML_Schemas.CMOF_OwnedMember, packName string) string {

	enum := simpleTypesMap[member.Name]
	enumName := getElementType(enum.Name)
	// Create the xml parser here...
	xmlParserStr := "type Xsd" + enumName + " struct {\n"
	xmlParserStr += "	Name	xml.Name	`xml:\"" + enumName + "\"`\n"

	// Now I will parse the possible value...
	if enum.Restriction.Base == "xsd:string" {
		for i := 0; i < len(enum.Restriction.Enumeration); i++ {
			xmlParserStr += "	M_" + enum.Restriction.Enumeration[i].Value
			xmlParserStr += "	string	`xml:\"string, omitempty\"`\n"
		}
	}

	xmlParserStr += "\n}"

	return xmlParserStr
}

func generateGoXmlParserContent(class *XML_Schemas.CMOF_OwnedMember, element *XML_Schemas.XSD_Element, packName string) string {

	xmlParserContentStr := ""

	var elementType = getElementType(element.Type)

	// There the element is a complex element???
	if complexTypesMap[elementType] != nil {

		// complex type here
		complexType := complexTypesMap[elementType]

		// Now I will create it content...
		if complexType.ComplexContent != nil {
			// So here the values must be sean as the base type...
			if complexType.ComplexContent.Extension != nil {
				baseType := complexType.ComplexContent.Extension.Base
				if len(baseType) > 0 {

					// Contain a sequences...
					if complexType.ComplexContent.Extension.Sequence != nil {
						for i := 0; i < len(complexType.ComplexContent.Extension.Sequence.Elements); i++ {
							element := complexType.ComplexContent.Extension.Sequence.Elements[i]
							xmlParserContentStr += generateGoXmlElement(&element, packName)
						}
					}

					// Now it list of attributes...
					if complexType.ComplexContent.Extension.Attributes != nil {
						for i := 0; i < len(complexType.ComplexContent.Extension.Attributes); i++ {
							attribute := complexType.ComplexContent.Extension.Attributes[i]
							xmlParserContentStr += generateGoXmlAttribute(&attribute)
						}
					}

					// Now the choice...
					// The choice can contain a sequence...
					if complexType.ComplexContent.Extension.Choice != nil {
						choice := complexType.ComplexContent.Extension.Choice
						if len(choice.Sequences) > 0 {
							for i := 0; i < len(choice.Sequences); i++ {
								if len(choice.Sequences[i].Elements) > 0 {
									for j := 0; j < len(choice.Sequences[i].Elements); j++ {
										element := choice.Sequences[i].Elements[j]
										xmlParserContentStr += generateGoXmlElement(&element, packName)
									}
								}
							}
						}
						for i := 0; i < len(choice.Elements); i++ {
							xmlParserContentStr += generateGoXmlElement(&choice.Elements[i], packName)
						}
					}

				}
			}
		}
		if complexType.Sequence != nil {

			// Elements contained in a sequences...
			for i := 0; i < len(complexType.Sequence.Elements); i++ {
				element := complexType.Sequence.Elements[i]
				xmlParserContentStr += generateGoXmlElement(&element, packName)
			}

			// Attribute contain in a sequences...
			for i := 0; i < len(complexType.Sequence.Any); i++ {
				attribute := complexType.Sequence.Any[i]
				if len(attribute.Name) > 0 {
					//chardata
					xmlParserContentStr += "	M_" + attribute.Name + "	string	`xml:\",innerxml\"`" + "\n"
				}
			}
		}
		// Now it list of attributes...
		if complexType.Attributes != nil {
			for i := 0; i < len(complexType.Attributes); i++ {
				attribute := complexType.Attributes[i]
				xmlParserContentStr += generateGoXmlAttribute(&attribute)
			}
		}

		// The any Attribute...
		if complexType.AnyAttributes != nil {
			for i := 0; i < len(complexType.AnyAttributes); i++ {
				attribute := complexType.AnyAttributes[i]
				if len(attribute.Name) > 0 {

					//chardata
					if attribute.Name == "other" && (elementType != "Expression" || elementType != "FormalExpression") {
						xmlParserContentStr += "//"
					}
					xmlParserContentStr += "	M_" + attribute.Name + "	string	`xml:\",innerxml\"`" + "\n"
				}
			}
		}
	} else {
		log.Println("Not found!!!", element)
	}
	xmlParserContentStr += "\n"
	return xmlParserContentStr
}

func generateBaseClassXmlMembers(elementName string, packName string, element *XML_Schemas.XSD_Element, elementType string, classIndex *int, isArray bool, isPointer bool) string {
	elementStr := ""
	classNameLst := substitutionGroupMap[elementType]

	for i := 0; i < len(classNameLst); i++ {
		className := classNameLst[i]
		subElement := elementsTypeMap[className]
		if subElement != nil {
			subElementType := getElementType(subElement.Type)

			packName_ := ""
			if len(packName) > 0 {
				if membersPackage[className] != packName {
					packName_ = membersPackage[className] + "."
				}
			}

			// if the subElement is a super class...
			if substitutionGroupMap[className] != nil {
				elementStr += generateBaseClassXmlMembers(elementName, packName, subElement, subElementType, classIndex, isArray, isPointer)
			}

			if Utility.Contains(abstractClassLst, subElementType) == false {
				if len(subElementType) > 0 {
					subElementType_ := subElement.Name
					// Also generate the
					elementStr_ := "	M_" + elementName + "_" + strconv.Itoa(*classIndex)
					if isArray {
						if Utility.Contains(abstractClassLst, subElementType) {
							elementStr_ += "	[]" + packName_ + "Xsd" + subElementType + "	`xml:\"" + subElementType_ + ",omitempty\"`"
						} else {
							elementStr_ += "	[]*" + packName_ + "Xsd" + subElementType + "	`xml:\"" + subElementType_ + ",omitempty\"`"
						}

					} else if isPointer {
						elementStr_ += "	*" + packName_ + "Xsd" + subElementType + "	`xml:\"" + subElementType_ + ",omitempty\"`"
					} else {
						if Utility.Contains(abstractClassLst, subElementType) {
							elementStr_ += "	" + packName_ + "Xsd" + subElementType + "	`xml:\"" + subElementType_ + "\"`"
						} else {
							elementStr_ += "	*" + packName_ + "Xsd" + subElementType + "	`xml:\"" + subElementType_ + "\"`"
						}
					}
					*classIndex++
					elementStr += elementStr_ + "\n"
				}

			}
		}
	}

	return elementStr
}

func generateGoXmlElement(element *XML_Schemas.XSD_Element, packName string) string {

	// Keep track of the min and max occurrence here...
	maxOccurs := element.MaxOccurs
	minOccurs := element.MinOccurs
	isRef := false

	// Set the ref...
	if len(element.Ref) > 0 {
		var elementRef = element.Ref
		var element_ = element
		if strings.Index(elementRef, ":") > -1 {
			elementRef = getElementType(elementRef)
		}
		element = elementsNameMap[elementRef]
		if element == nil {
			if len(element_.Type) > 0 {
				element = elementsTypeMap[getElementType(element_.Type)]
				element.Name = elementRef
				isRef = true
			}
		}
	}

	if isRef == false {
		if strings.HasSuffix(element.Name, "Ref") {
			isRef = true
			element.Ref = element.Name
		}
	}

	if element == nil {
		return ""
	}

	var elementType = getElementType(element.Type)

	elementStr := ""

	if len(elementType) == 0 {
		return ""
	}

	// Here the member is a reference...
	if isRef {
		elementType = "xsd:QName"
	}

	if len(xsdPrimitiveTypesMap[elementType]) > 0 {
		// So here the attribute can be a reference
		typeName := xsdPrimitiveTypesMap[elementType]
		elementStr += "	M_" + element.Name
		if maxOccurs == "unbounded" {
			elementStr += "	[]" + typeName + "	`xml:\"" + element.Name + "\"`"
		} else if minOccurs == "0" || (elementType == "xsd:QName" && minOccurs != "1") {
			elementStr += "	*" + typeName + "	`xml:\"" + element.Name + "\"`"
		} else {
			elementStr += "	" + typeName + "	`xml:\"" + element.Name + "\"`"
		}

	} else if complexTypesMap[elementType] != nil {
		packName_ := ""

		if membersPackage[elementType] != packName {
			packName_ = membersPackage[elementType] + "."
		}
		classIndex := 0
		if substitutionGroupMap[elementType] != nil {
			// So here I will get the list of subtype for that super class..
			isArray := false
			isPointer := false

			if maxOccurs != "" {
				isArray = maxOccurs == "unbounded"
			}

			if minOccurs != "" {
				isPointer = (minOccurs == "0" || minOccurs == "1") && !isArray
			}
			// Nothing todo here...
			if _, ok := aliasElementType[element.Name]; !ok {
				elementStr += generateBaseClassXmlMembers(element.Name, packName, element, elementType, &classIndex, isArray, isPointer)
			}
		}

		if Utility.Contains(abstractClassLst, elementType) == false {
			elementStr += "	M_" + element.Name
			elementType_ := element.Name
			if _, ok := aliasElementType[element.Name]; ok {
				elementType = strings.ToUpper(element.Name[0:1]) + element.Name[1:]
				minOccurs = "0"
			}
			if classIndex > 0 {
				elementStr += "_" + strconv.Itoa(classIndex)
			}

			if maxOccurs == "unbounded" {
				if Utility.Contains(abstractClassLst, elementType) {
					elementStr += "	[]" + packName_ + "Xsd" + elementType + "	`xml:\"" + elementType_ + ",omitempty\"`"
				} else {
					elementStr += "	[]*" + packName_ + "Xsd" + elementType + "	`xml:\"" + elementType_ + ",omitempty\"`"
				}
			} else if minOccurs == "0" || Utility.Contains(superClassesLst, elementType) {
				elementStr += "	*" + packName_ + "Xsd" + elementType + "	`xml:\"" + elementType_ + ",omitempty\"`"
			} else {
				elementStr += "	" + packName_ + "Xsd" + elementType + "	`xml:\"" + elementType_ + "\"`"
			}
		}

	} else {
		log.Println("==============> Not found: ", elementType)
		return ""
	}

	return elementStr + "\n"
}

func generateGoXmlAttribute(attribute *XML_Schemas.XSD_Attribute) string {
	if len(attribute.Name) == 0 {
		return ""
	}

	attributeStr := "	M_" + attribute.Name
	typeName := ""
	var attributeType = getElementType(attribute.Type)

	// The attribute is a refencre to a complex type...
	if complexTypesMap[attributeType] != nil {
		typeName = attribute.Type
		attributeStr += "	*Xsd" + typeName + "	`xml:\"" + attribute.Name + "\"`"
	} else if len(xsdPrimitiveTypesMap[attributeType]) > 0 {
		// So here the attribute can be a reference
		typeName = xsdPrimitiveTypesMap[attributeType]
		attributeStr += "	" + typeName + "	`xml:\"" + attribute.Name + ",attr\"`"
	} else if simpleTypesMap[attributeType] != nil {
		// So here the found value is an enumeration...
		attributeStr += "	string" + "	`xml:\"" + attribute.Name + ",attr\"`"
	} else {
		attributeStr += "	string" + "	`xml:\"" + attribute.Name + ",attr\"`"
	}

	return attributeStr + "\n"
}

/////////////////////////////////////////////////////////////////////////////////
// CMOF -> GO
/////////////////////////////////////////////////////////////////////////////////

/**
 * Generate the variables...
 */
func generateGoMemberCode(attribute *XML_Schemas.CMOF_OwnedAttribute, isAssociation bool, packName string) string {

	//log.Println("-------------------------> ", attribute.Name, "is composite", attribute.IsComposite, "isAssociation", isAssociation)
	var memberStr string
	var typeName, isPrimitive, isPointer = getAttributeTypeName(attribute)

	attrNs := getAttributeNs(attribute)

	isInterface := Utility.Contains(superClassesLst, typeName) || Utility.Contains(abstractClassLst, typeName)

	if len(attrNs) > 0 && isPrimitive == false {
		if len(packName) > 0 {
			if attrNs != packName {
				typeName = attrNs + "." + typeName
			}
		}
	}

	var memberName = attribute.Name
	if isPointer == true && !isInterface && simpleTypesMap[typeName] == nil {
		typeName = "*" + typeName
	}

	if attribute.Upper == "*" {
		typeName = "[]" + typeName
	}

	ref := ""
	if isAssociation {
		ref = "Ptr"
	}

	// The string key to that reference...
	if isPrimitive || enumMap[typeName] != nil {
		memberStr += "	M_" + memberName + ref + " " + typeName + "\n"
	} else {
		if attribute.Upper == "*" {
			memberStr += "	M_" + memberName + ref + " []string\n"
		} else {
			memberStr += "	M_" + memberName + ref + " string\n"
		}
	}

	return memberStr
}

/**
 * Generate all member of a class
 */
func generateGoClassMembersCode(class *XML_Schemas.CMOF_OwnedMember, packageId string) string {
	var classStr string
	classStr += "	/** members of " + class.Name + " **/\n"
	if len(class.Attributes) == 0 {
		classStr += "	/** No members **/\n"
	}
	for k := 0; k < len(class.Attributes); k++ {
		attribute := class.Attributes[k]
		classStr += generateGoMemberCode(&attribute, false, packageId)
	}

	classStr += "\n"
	return classStr
}

/**
 * This is the prototype of a methode...
 * like GetResolution()float32
 */
func generateGoMethodCode(attribute *XML_Schemas.CMOF_OwnedAttribute, owner *XML_Schemas.CMOF_OwnedMember, isAssociation bool, packName string) string {

	// Here we are in a struct...
	typeName, isPrimitive, isPointer := getAttributeTypeName(attribute)
	var methodName = strings.ToUpper(attribute.Name[0:1]) + attribute.Name[1:]
	var attributeName = attribute.Name
	if isAssociation {
		attributeName += "Ptr"
		methodName += "Ptr"
	}

	// The getter function...
	getterStr := ""
	if typeName == "bool" || typeName == "boolean" {
		if !strings.HasPrefix(methodName, "Is") {

			getterStr += "Is"
		}
	} else {
		getterStr += "Get"
	}

	getterStr += methodName + "()"
	if attribute.Upper == "*" {
		getterStr += "[]"
	}
	isInterface := Utility.Contains(superClassesLst, typeName) || Utility.Contains(abstractClassLst, typeName)
	if isPointer == true && !isInterface && simpleTypesMap[typeName] == nil {
		getterStr += "*"
	}
	if membersPackage[typeName] != packName {
		if !isPrimitive && enumMap[typeName] == nil {
			typeName = membersPackage[typeName] + "." + typeName
		}
	}

	getterStr += typeName
	if owner != nil {
		ownerName := owner.Name
		if Utility.Contains(superClassesLst, ownerName) == true {
			ownerName += "_impl"
		}
		getterStr = "func (this *" + ownerName + ") " + getterStr
		getterStr += "{\n"

		// The getter function body here.
		if isPrimitive || enumMap[typeName] != nil {
			getterStr += "	return this.M_" + attribute.Name + "\n"
		} else {

			if attribute.Upper == "*" {
				getterStr += "	values := make([]"

				if isPointer == true && !isInterface {
					getterStr += "*"
				}

				getterStr += typeName
				getterStr += ", 0)\n"
				getterStr += "	for i := 0; i < len(this.M_" + attributeName + "); i++ {\n"
				getterStr += "		entity, err := this.getEntityByUuid(this.M_" + attributeName + "[i])\n"
				getterStr += "		if err == nil {\n"
				getterStr += "			values = append( values, entity.("
				if isPointer == true && !isInterface {
					getterStr += "*"
				}
				getterStr += typeName
				getterStr += "))\n"
				getterStr += "		}\n"
				getterStr += "	}\n"
				getterStr += "	return values\n"
			} else {
				getterStr += "	entity, err := this.getEntityByUuid(this.M_" + attributeName + ")\n"
				getterStr += "	if err == nil {\n"
				getterStr += "		return entity.("
				if isPointer == true && !isInterface {
					getterStr += "*"
				}
				getterStr += typeName
				getterStr += ")\n"
				getterStr += "	}\n"
				getterStr += "	return nil\n"
			}
		}
		getterStr += "}\n"
	}

	// The setter function...
	setterStr := "Set"
	setterStr += methodName
	setterStr += "(val "
	if attribute.Upper == "*" {
		setterStr += "[]"
	}
	if isPointer == true && !isInterface && simpleTypesMap[typeName] == nil {
		setterStr += "*"
	}
	setterStr += typeName + ")"

	if owner != nil {
		ownerName := owner.Name
		if Utility.Contains(superClassesLst, ownerName) == true {
			ownerName += "_impl"
		}
		setterStr = "func (this *" + ownerName + ") " + setterStr
		setterStr += "{\n"
		// The setter function body here.
		if isPrimitive || enumMap[typeName] != nil {
			setterStr += "	this.M_" + attributeName + "= val\n"
		} else {
			if attribute.Upper == "*" {
				setterStr += "	this.M_" + attributeName + "= make([]string,0)\n"
				setterStr += "	for i:=0; i < len(val); i++{\n"
				if !IsRef(attribute) && enumMap[attribute.Name] == nil {
					// Set the parent infos in the child.
					setterStr += "		val[i].SetParentUuid(this.UUID)\n"
					setterStr += "		val[i].SetParentLnk(\"M_" + attributeName + "\")\n"
				}
				setterStr += "		this.M_" + attributeName + "=append(this.M_" + attributeName + ", val[i].GetUuid())\n"
				setterStr += "		this.setEntity(val[i])\n"
				setterStr += "	}\n"
				setterStr += "	this.setEntity(this)\n"
			} else {
				if !IsRef(attribute) && enumMap[attribute.Name] == nil {
					// Set the parent infos in the child.
					setterStr += "	val.SetParentUuid(this.UUID)\n"
					setterStr += "	val.SetParentLnk(\"M_" + attributeName + "\")\n"
					setterStr += "  this.setEntity(val)\n"
				}
				setterStr += "	this.M_" + attributeName + "= val.GetUuid()\n"
				setterStr += "	this.setEntity(this)\n"
			}

		}
		setterStr += "}\n"
	}

	// In case of array I will also create append method...
	var appendStr string
	if attribute.Upper == "*" {
		appendStr = " Append"
		appendStr += methodName
		appendStr += "(val "
		if isPointer == true && !isInterface && simpleTypesMap[typeName] == nil {
			appendStr += "*"
		}
		appendStr += typeName + ")"
		if owner != nil {
			ownerName := owner.Name
			if Utility.Contains(superClassesLst, ownerName) == true {
				ownerName += "_impl"
			}
			appendStr = "\nfunc (this *" + ownerName + ")" + appendStr
			appendStr += "{\n"
			// The append function body here.
			if isPrimitive {
				appendStr += "	this.M_" + attributeName + "=append(this.M_" + attributeName + ", val)\n"
			} else {
				appendStr += "	for i:=0; i < len(this.M_" + attributeName + "); i++{\n"
				appendStr += "		if this.M_" + attributeName + "[i] == val.GetUuid() {\n"
				appendStr += "			return\n"
				appendStr += "		}\n"
				appendStr += "	}\n"
				if !IsRef(attribute) && enumMap[attribute.Name] == nil {
					// Set the parent infos in the child.
					appendStr += "	val.SetParentUuid(this.UUID)\n"
					appendStr += "	val.SetParentLnk(\"M_" + attributeName + "\")\n"
					appendStr += "  this.setEntity(val)\n"
				}
				appendStr += "	this.M_" + attributeName + " = append(this.M_" + attributeName + ", val.GetUuid())\n"
				appendStr += "	this.setEntity(this)\n"
			}
			appendStr += "}\n"
		}
	}

	var removerStr string
	if (isPointer || attribute.Upper == "*") && !isPrimitive {
		// The remover function...
		if attribute.Upper == "*" {
			removerStr += "Remove"
			removerStr += methodName + "("
			if !isPrimitive {
				removerStr += "val "
				if isPointer == true && !isInterface && simpleTypesMap[typeName] == nil {
					removerStr += "*"
				}
				removerStr += typeName
			} else {
				removerStr += "index int"
			}
			removerStr += ")"
		} else {
			removerStr += "Reset" + methodName + "()"
		}

		if owner != nil {
			ownerName := owner.Name
			if Utility.Contains(superClassesLst, ownerName) == true {
				ownerName += "_impl"
			}

			removerStr = "func (this *" + ownerName + ") " + removerStr
			removerStr += "{\n"
			// The delete function body here.
			if attribute.Upper == "*" {
				if isPrimitive {
					removerStr += "	this.M_" + attributeName + "=append(this.M_" + attributeName + "[0:index], [index+1:], ...)\n"
				} else {
					removerStr += "	values := make([]string,0)\n"
					removerStr += "	for i:=0; i < len(this.M_" + attributeName + "); i++{\n"
					removerStr += "		if this.M_" + attributeName + "[i] != val.GetUuid() {\n"
					removerStr += "			values = append(values, val.GetUuid())\n"
					removerStr += "		}\n"
					removerStr += "	}\n"
					removerStr += "	this.M_" + attributeName + " = values\n"
					removerStr += "	this.setEntity(this)\n"
				}
			} else {
				if enumMap[typeName] != nil {
					removerStr += "	this.M_" + attributeName + "= 0\n"
				} else {
					removerStr += "	this.M_" + attributeName + "= \"\"\n"
				}
			}

			removerStr += "}\n"
		}
	}

	return getterStr + "\n" + setterStr + "\n" + appendStr + "\n" + removerStr + "\n"

}

// Interface
func generateGoInterfacesCode(packageId string) {
	classes := packagesMembers[packageId]
	for i := 0; i < len(classes); i++ {
		class := classesMap[classes[i]]
		if class != nil {
			if Utility.Contains(abstractClassLst, class.Name) == true {
				generateGoInterfaceCode(packageId, class)
			}
		}
	}
}

// If the utility package must be import...
var hasUtility = make(map[string]bool, 0)

func generateGoInterfaceCode(packageId string, class *XML_Schemas.CMOF_OwnedMember) {

	// Set the pacakge here.
	var classStr string
	classStr = "package " + packageId + "\n\n"

	imports := getClassImports(class)

	if len(imports) > 0 {
		classStr += "import(\n"
		for j := 0; j < len(imports); j++ {
			classStr += "\"" + imports[j] + "\"\n"
		}

		classStr += ")\n\n"
	}

	classStr += "type " + class.Name + " interface{\n"

	// Now The method...
	if len(class.Attributes) > 0 {
		classStr += "	/** Method of " + class.Name + " **/\n\n"
	}

	classStr += "	/**\n"
	classStr += "	 * Return the type name of an entity\n"
	classStr += "	 */\n"
	classStr += "	GetTypeName() string\n\n"

	classStr += "	/**\n"
	classStr += "	 * Get an entity's uuid\n"
	classStr += "	 * Each entity must have one uuid.\n"
	classStr += "	 */\n"
	classStr += "	GetUuid() string\n"
	classStr += "	SetUuid(uuid string)\n\n"

	classStr += "	/**\n"
	classStr += "	 * Return the array of entity id's without it uuid.\n"
	classStr += "	 */\n"
	classStr += "	Ids() []interface{}\n\n"

	classStr += "	/**\n"
	classStr += "	 * Get an entity's parent UUID if it have a parent.\n"
	classStr += "	 */\n"
	classStr += "	GetParentUuid() string\n"
	classStr += "	SetParentUuid(uuid string)\n\n"

	classStr += "	/**\n"
	classStr += "	 * The name of the relation with it parent.\n"
	classStr += "	 */\n"
	classStr += "	GetParentLnk() string\n"
	classStr += "	SetParentLnk(string)\n\n"

	classStr += "	/**\n"
	classStr += "	 * Return link to entity childs.\n"
	classStr += "	 */\n"
	classStr += "	GetChilds() []interface{}\n\n"

	classStr += "	/**\n"
	classStr += "	 * Set the function GetEntityByUuid as a pointer. The entity manager can't\n"
	classStr += "	 * be access by Entities package...\n"
	classStr += "	 */\n"
	classStr += "	 SetEntityGetter(func(uuid string) (interface{}, error))\n\n"

	for j := 0; j < len(class.Attributes); j++ {
		attribute := class.Attributes[j]
		if attribute.Type == "cmof:Property" {
			classStr += "\n" + generateGoMethodCode(&attribute, nil, false, membersPackage[class.Name])
			classStr += "\n"
		}
	}

	classStr += "}"

	WriteClassFile(outputPath, packageId, class.Name, classStr)

}

// Class
func generateGoClassCode(packageId string) {
	classes := packagesMembers[packageId]
	for i := 0; i < len(classes); i++ {
		class := classesMap[classes[i]]
		// Array of super class name...
		var superClasses []string
		if class != nil {
			if Utility.Contains(abstractClassLst, class.Name) == false {
				className := class.Name
				// If the class is use as a base to other class
				// an interface must be generated...
				if Utility.Contains(superClassesLst, class.Name) == true {
					generateGoInterfaceCode(packageId, class)
					className += "_impl"
				}

				var classStr = "type " + className + " struct{\n"

				classStr += "\n	/** The entity UUID **/\n"
				classStr += "	UUID string\n"
				classStr += "	/** The entity TypeName **/\n"
				classStr += "	TYPENAME string\n"
				classStr += "	/** The parent uuid if there is some. **/\n"
				classStr += "	ParentUuid string\n"
				classStr += "	/** The relation name with the parent. **/\n"
				classStr += "	ParentLnk string\n"
				classStr += "	/** Get entity by uuid function **/\n"
				classStr += "	getEntityByUuid func(string)(interface{}, error)\n"
				classStr += "	/** Use to put the entity in the cache **/\n"
				classStr += "	setEntity func(interface{})\n"
				classStr += "	/** Generate the entity uuid **/\n"
				classStr += "	generateUuid func(interface{}) string\n\n"

				superClasses = getSuperClasses(class.Name)
				for j := 0; j < len(superClasses); j++ {
					superClass := classesMap[superClasses[j]]
					classStr += generateGoClassMembersCode(superClass, packageId)
				}
				// Now the member of the class itself...
				classStr += generateGoClassMembersCode(class, packageId)

				associations := getClassAssociations(class)
				if len(associations) > 0 {
					classStr += "\n	/** Associations **/\n"
				}

				for _, att := range associations {
					associationsStr := generateGoMemberCode(att, true, packageId)
					if strings.Index(classStr, associationsStr) == -1 {
						classStr += associationsStr
					}
				}

				classStr += "}"
				imports := getClassImports(class)

				// The xml code
				xmlParserStr := generateGoXmlParser(class, packageId)
				if len(xmlParserStr) > 0 {
					imports = append(imports, "encoding/xml")
				}

				if len(xmlParserStr) > 0 {
					classStr += "\n\n/** Xml parser for " + class.Name + " **/\n"
					classStr += xmlParserStr
				}

				classStr += "/***************** Entity **************************/\n\n"

				classStr += "/** UUID **/\n"
				classStr += "func (this *" + className + ") GetUuid() string{\n"
				classStr += "	if len(this.UUID) == 0 {\n"
				classStr += "		this.SetUuid(this.generateUuid(this))\n"
				classStr += "	}\n"
				classStr += "	return this.UUID\n"
				classStr += "}\n"

				classStr += "func (this *" + className + ") SetUuid(uuid string){\n"
				classStr += "	this.UUID = uuid\n"
				classStr += "}\n\n"

				classStr += "/** Return the array of entity id's without it uuid **/\n"
				classStr += "func (this *" + className + ") Ids() []interface{} {\n"
				classStr += "	ids := make([]interface{}, 0)\n"

				superClasses := getSuperClasses(class.Name)
				for j := 0; j < len(superClasses); j++ {
					superClass := classesMap[superClasses[j]]
					for k := 0; k < len(superClass.Attributes); k++ {
						attribute := superClass.Attributes[k]
						if isId(superClass.Name, attribute.Name) {
							classStr += "	ids = append(ids, this.M_" + attribute.Name + ")\n"
						}
					}
				}

				// So here I will
				for j := 0; j < len(class.Attributes); j++ {
					attribute := class.Attributes[j]
					if isId(class.Name, attribute.Name) {
						classStr += "	ids = append(ids, this.M_" + attribute.Name + ")\n"
					}
				}

				classStr += "	return ids\n"
				classStr += "}\n\n"

				classStr += "/** The type name **/\n"
				classStr += "func (this *" + className + ") GetTypeName() string{\n"
				classStr += "	this.TYPENAME = \"" + packageId + "." + className + "\"\n"
				classStr += "	return this.TYPENAME\n"
				classStr += "}\n\n"

				classStr += "/** Return the entity parent UUID **/\n"
				classStr += "func (this *" + className + ") GetParentUuid() string{\n"
				classStr += "	return this.ParentUuid\n"
				classStr += "}\n\n"

				classStr += "/** Set it parent UUID **/\n"
				classStr += "func (this *" + className + ") SetParentUuid(parentUuid string){\n"
				classStr += "	this.ParentUuid = parentUuid\n"
				classStr += "}\n\n"

				classStr += "/** Return it relation with it parent, only one parent is possible by entity. **/\n"
				classStr += "func (this *" + className + ") GetParentLnk() string{\n"
				classStr += "	return this.ParentLnk\n"
				classStr += "}\n"
				classStr += "func (this *" + className + ") SetParentLnk(parentLnk string){\n"
				classStr += "	this.ParentLnk = parentLnk\n"
				classStr += "}\n\n"

				classStr += "func (this *" + className + ") GetParent() interface{}{\n"
				classStr += "	parent, err := this.getEntityByUuid(this.ParentUuid)\n"
				classStr += "	if err != nil {\n"
				classStr += "		return nil\n"
				classStr += "	}\n"
				classStr += "	return parent\n"
				classStr += "}\n\n"

				classStr += "/** Return it relation with it parent, only one parent is possible by entity. **/\n"
				classStr += "func (this *" + className + ") GetChilds() []interface{}{\n"
				classStr += "	var childs []interface{}\n"

				hasChild := false
				for j := 0; j < len(superClasses); j++ {
					superClass := classesMap[superClasses[j]]
					for k := 0; k < len(superClass.Attributes); k++ {
						attribute := superClass.Attributes[k]
						typeName, isPrimitive, _ := getAttributeTypeName(&attribute)
						if !isPrimitive && enumMap[typeName] == nil && !IsRef(&attribute) {
							if !hasChild {
								classStr += "	var child interface{}\n"
								classStr += "	var err error\n"
								hasChild = true
							}
							if attribute.Upper == "*" {
								// Here The attribute is an array
								classStr += "	for i:=0; i < len(this.M_" + attribute.Name + "); i++ {\n"
								classStr += "		child, err = this.getEntityByUuid( this.M_" + attribute.Name + "[i])\n"
								classStr += "		if err == nil {\n"
								classStr += "			childs = append( childs, child)\n"
								classStr += "		}\n"
								classStr += "	}\n"
							} else {
								// Here the attribute is not an array....
								classStr += "	child, err = this.getEntityByUuid( this.M_" + attribute.Name + ")\n"
								classStr += "	if err == nil {\n"
								classStr += "		childs = append( childs, child)\n"
								classStr += "	}\n"
							}
						}
					}
				}

				for j := 0; j < len(class.Attributes); j++ {
					attribute := class.Attributes[j]
					typeName, isPrimitive, _ := getAttributeTypeName(&attribute)
					if !isPrimitive && enumMap[typeName] == nil && !IsRef(&attribute) {
						if !hasChild {
							classStr += "	var child interface{}\n"
							classStr += "	var err error\n"
							hasChild = true
						}
						if attribute.Upper == "*" {
							// Here The attribute is an array
							classStr += "	for i:=0; i < len(this.M_" + attribute.Name + "); i++ {\n"
							classStr += "		child, err = this.getEntityByUuid( this.M_" + attribute.Name + "[i])\n"
							classStr += "		if err == nil {\n"
							classStr += "			childs = append( childs, child)\n"
							classStr += "		}\n"
							classStr += "	}\n"
						} else {
							// Here the attribute is not an array....
							classStr += "	child, err = this.getEntityByUuid( this.M_" + attribute.Name + ")\n"
							classStr += "	if err == nil {\n"
							classStr += "		childs = append( childs, child)\n"
							classStr += "	}\n"
						}
					}
				}

				classStr += "	return childs\n"
				classStr += "}\n"

				classStr += "/** Give access to entity manager GetEntityByUuid function from Entities package. **/\n"
				classStr += "func (this *" + className + ") SetEntityGetter(fct func(uuid string)(interface{}, error)){\n"
				classStr += "	this.getEntityByUuid = fct\n"
				classStr += "}\n"

				classStr += "/** Use it the set the entity on the cache. **/\n"
				classStr += "func (this *" + className + ") SetEntitySetter(fct func(entity interface{})){\n"
				classStr += "	this.setEntity = fct\n"
				classStr += "}\n"

				classStr += "/** Set the uuid generator function **/\n"
				classStr += "func (this *" + className + ") SetUuidGenerator(fct func(entity interface{}) string){\n"
				classStr += "	this.generateUuid = fct\n"
				classStr += "}\n"

				// Now the method...
				// The superclass methode...
				for j := 0; j < len(superClasses); j++ {
					superClass := classesMap[superClasses[j]]
					for k := 0; k < len(superClass.Attributes); k++ {
						classStr += "\n" + generateGoMethodCode(&superClass.Attributes[k], class, false, membersPackage[class.Name])
					}
				}

				// The class methode
				for j := 0; j < len(class.Attributes); j++ {
					classStr += "\n" + generateGoMethodCode(&class.Attributes[j], class, false, membersPackage[class.Name])
				}

				// The associations getter / setter
				for _, att := range associations {
					associationsStr := generateGoMethodCode(att, class, true, membersPackage[class.Name])
					if strings.Index(classStr, associationsStr) == -1 {
						classStr += "\n" + associationsStr
					}
				}

				var classStr_ = "package " + packageId + "\n\n"
				if hasUtility[packageId+"."+class.Name] == true || hasUtility[packageId+"."+class.Name+"_impl"] == true {
					imports = append(imports, "code.myceliUs.com/Utility")
				}

				if len(imports) > 0 {
					classStr_ += "import(\n"
					for j := 0; j < len(imports); j++ {
						classStr_ += "	\"" + imports[j] + "\"\n"
					}
					classStr_ += ")\n\n"
				}

				classStr = classStr_ + classStr

				// Here I will write the string in a file...
				WriteClassFile(outputPath, packageId, className, classStr)
			}
		}
	}
}

// Enumeration
func generateGoEnumerationCode(packageId string) {
	for _, enum := range enumMap {
		// Generate in the good folder...
		if membersPackage[enum.Name] == packageId {
			var enumType = getElementType(enum.Name)
			enumStr := "package " + packageId + "\n\n"
			enumStr += "type " + enumType + " int\n"
			enumStr += "const(\n"
			for i := 0; i < len(enum.Litterals); i++ {
				litteral := enum.Litterals[i]
				enumValue := litteral.Name
				enumValue = strings.ToUpper(enumValue[0:1]) + enumValue[1:]
				enumStr += "	" + enumType + "_" + enumValue
				if i == 0 {
					enumStr += " " + enumType + " = 1+iota"
				}
				enumStr += "\n"

			}
			enumStr += ")\n"
			WriteClassFile(outputPath, packageId, enumType, enumStr)
		}
	}
}
