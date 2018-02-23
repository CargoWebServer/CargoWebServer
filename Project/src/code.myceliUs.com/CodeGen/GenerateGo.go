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

	isRef := IsRef(attribute)
	// The reference pointer here

	if isRef {
		// The initialyse memory reference...
		memberStr += "	m_" + memberName + ref + " " + typeName + "\n"

		// The string key to that reference...
		memberStr += "	/** If the ref is a string and not an object **/\n"
		if attribute.Upper == "*" {
			memberStr += "	M_" + memberName + ref + " []string\n"
		} else {
			memberStr += "	M_" + memberName + ref + " string\n"
		}

	} else {
		memberStr += "	M_" + memberName + ref + " " + typeName + "\n"
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
	var methodStr string
	var getter string
	var setter string

	// Here I will print the property...
	// is abstract or not...
	var typeName, isPrimitive, isPointer = getAttributeTypeName(attribute)
	typeName_ := typeName
	if !isPrimitive {
		typeName = getClassPackName(typeName, packName) + typeName
	}

	//log.Println(attribute.Name)
	var methodName = strings.ToUpper(attribute.Name[0:1]) + attribute.Name[1:]
	isInterface := Utility.Contains(superClassesLst, typeName) || Utility.Contains(abstractClassLst, typeName)

	if isPointer == true && !isInterface && simpleTypesMap[typeName] == nil {
		typeName = "*" + typeName
	}

	ref := ""
	if isAssociation == true {
		ref = "Ptr"
	}

	// If the method is a boolean and is name beging with is...
	if strings.HasPrefix(methodName, "Is") && (typeName == "bool" || typeName == "boolean") {
		getter += methodName + "() "
		if attribute.Upper == "*" {
			getter += "[]"
		}
		getter += typeName
	} else {
		getter += "Get" + methodName + ref + "() "
		if attribute.Upper == "*" {
			getter += "[]"
		}
		getter += typeName
	}

	if owner != nil {
		ownerName := owner.Name
		if Utility.Contains(superClassesLst, ownerName) == true {
			ownerName += "_impl"
		}

		// Here we are in a struct...
		getter = "func (this *" + ownerName + ") " + getter + "{\n"
		isRef := IsRef(attribute)

		if isRef && attribute.Name != "other" {
			// Here the reference are not initialyse at first...
			if attribute.Upper == "*" {
				getter += "	if this.m_" + attribute.Name + ref + " == nil {\n"
				getter += "		this.m_" + attribute.Name + ref + " = make([]" + typeName + ", 0)\n"
				getter += "		for i := 0; i < len(this.M_" + attribute.Name + ref + "); i++ {\n"
				getter += "			entity, err := this.getEntityByUuid(this.M_" + attribute.Name + ref + "[i])\n"
				getter += "			if err == nil {\n"
				getter += "				this.m_" + attribute.Name + ref + " = append(this.m_" + attribute.Name + ref + ", entity.(" + typeName + "))\n"
				getter += "			}\n"
				getter += "		}\n"
				getter += "	}\n"
			} else {
				getter += "	if this.m_" + attribute.Name + ref + " == nil {\n"
				getter += "		entity, err := this.getEntityByUuid(this.M_" + attribute.Name + ref + ")\n"
				getter += "		if err == nil {\n"
				getter += "			this.m_" + attribute.Name + ref + " = entity.(" + typeName + ")\n"
				getter += "		}\n"
				getter += "	}\n"
			}
			getter += "	return this.m_" + attribute.Name + ref + "\n"
		} else {
			getter += "	return this.M_" + attribute.Name + ref + "\n"
		}

		getter += "}\n"
		if isRef && attribute.Name != "other" {
			if attribute.Upper == "*" {
				getter += "func (this *" + ownerName + ") Get" + methodName + ref + "Str() []string{\n"
			} else {
				getter += "func (this *" + ownerName + ") Get" + methodName + ref + "Str() string{\n"
			}
			getter += "	return this.M_" + attribute.Name + ref + "\n"
			getter += "}\n"
		}

		methodStr = "\n/** " + methodName + " **/\n" + getter

		/** Here I will create the function to initialyse a reference... **/
		methodStr += "\n/** Init reference " + methodName + " **/\n"
		setter = "Set" + methodName + ref + "(ref interface{})\n"

		refSetter := ""
		refSetter += "func (this *" + ownerName + ") Set" + methodName + ref + "(ref interface{}){\n"

		cast := typeName_
		elementPackName := ""

		if classesMap[cast] != nil {
			_, owner_ := getOwnedAttributeByName("id", classesMap[cast])
			cast = getClassPackName(owner_.Name, packName) + owner_.Name
			elementPackName = membersPackage[owner_.Name]
		}

		if attribute.Upper == "*" {
			if isRef {
				refSetter += "	if refStr, ok := ref.(string); ok {\n"
				refSetter += "		for i:=0; i < len(this.M_" + attribute.Name + ref + "); i++ {\n"
				refSetter += "			if this.M_" + attribute.Name + ref + "[i] == refStr {\n"
				refSetter += "				return\n"
				refSetter += "			}\n"
				refSetter += "		}\n"
				refSetter += "		this.M_" + attribute.Name + ref + " = append(this.M_" + attribute.Name + ref + ", ref.(string))\n"
				refSetter += "		this.NeedSave = true\n"
				refSetter += "	}else{\n"
				if (isRef || attribute.IsComposite == "true") && classesMap[typeName_] != nil {
					refSetter += "		for i:=0; i < len(this.m_" + attribute.Name + ref + "); i++ {\n"
					refSetter += "			if this.m_" + attribute.Name + ref + "[i].GetUuid() == ref.(" + typeName + ").GetUuid() {\n"
					refSetter += "				return\n"
					refSetter += "			}\n"
					refSetter += "		}\n"
					refSetter += "		isExist := false\n"
					refSetter += "		for i:=0; i < len(this.M_" + attribute.Name + ref + "); i++ {\n"
					refSetter += "			if this.M_" + attribute.Name + ref + "[i] == ref.(" + typeName + ").GetUuid() {\n"
					refSetter += "				isExist = true\n"
					refSetter += "			}\n"
					refSetter += "		}\n"
				}
				refSetter += "		this.m_" + attribute.Name + ref + " = append(this.m_" + attribute.Name + ref + ", ref.(" + typeName + "))\n"
				if classesMap[cast] != nil {
					isInterfaceCast := Utility.Contains(abstractClassLst, cast) || Utility.Contains(superClassesLst, cast)

					if elementPackName != packName {
						cast = elementPackName + "." + cast
					}
					if isPointer && !isInterfaceCast {
						cast = "*" + cast
					}
					refSetter += "	if !isExist {\n"
					refSetter += "		this.M_" + attribute.Name + ref + " = append(this.M_" + attribute.Name + ref + ", ref.(" + cast + ").GetUuid())\n"
					refSetter += "		this.NeedSave = true\n"
					refSetter += "	}\n"
				}
				refSetter += "	}\n"
			} else {
				refSetter += "	isExist := false\n"
				refSetter += "	var " + attribute.Name + "s []" + typeName + "\n"
				refSetter += "	for i:=0; i<len(this.M_" + attribute.Name + "); i++ {\n"

				if typeName == "string" {
					refSetter += "		if this.M_" + attribute.Name + "[i] != ref.(" + typeName + ") {\n"
				} else if typeName == "int" {
					refSetter += "		if this.M_" + attribute.Name + "[i] != ref.(" + typeName + ") {\n"
				} else {
					cast := cast[strings.Index(cast, ".")+1:]

					if classesMap[cast] != nil {
						isInterfaceCast := Utility.Contains(abstractClassLst, cast) || Utility.Contains(superClassesLst, cast)

						if elementPackName != packName {
							cast = elementPackName + "." + cast
						}
						if isPointer && !isInterfaceCast {
							cast = "*" + cast
						}
						if isInterface {
							refSetter += "		if this.M_" + attribute.Name + "[i].(" + cast + ").GetUuid() != ref.(" + cast + ").GetUuid() {\n"
						} else {
							refSetter += "		if this.M_" + attribute.Name + "[i].GetUuid() != ref.(" + cast + ").GetUuid() {\n"
						}

					} else {
						log.Println("------------> class not found! ", cast)
					}

				}

				refSetter += "			" + attribute.Name + "s = append(" + attribute.Name + "s, this.M_" + attribute.Name + "[i])\n"
				refSetter += "		} else {\n"
				refSetter += "			isExist = true\n"
				refSetter += "			" + attribute.Name + "s = append(" + attribute.Name + "s, ref.(" + typeName + "))\n"
				refSetter += "		}\n"
				refSetter += "	}\n"

				refSetter += "	if !isExist {\n"
				refSetter += "		" + attribute.Name + "s = append(" + attribute.Name + "s, ref.(" + typeName + "))\n"
				refSetter += "		this.NeedSave = true\n"
				refSetter += "		this.M_" + attribute.Name + " = " + attribute.Name + "s\n"
				refSetter += "	}\n"
			}
		} else {
			if isRef {
				refSetter += "	if _, ok := ref.(string); ok {\n"
				refSetter += "		if this.M_" + attribute.Name + ref + " != ref.(string) {\n"
				refSetter += "			this.M_" + attribute.Name + ref + " = ref.(string)\n"
				refSetter += "			this.NeedSave = true\n"
				refSetter += "		}\n"
				refSetter += "	}else{\n"

				if classesMap[cast] != nil {
					isInterfaceCast := Utility.Contains(abstractClassLst, cast) || Utility.Contains(superClassesLst, cast)
					if elementPackName != packName {
						cast = elementPackName + "." + cast
					}
					if isPointer && !isInterfaceCast {
						cast = "*" + cast
					}
					refSetter += "		if this.M_" + attribute.Name + ref + " != ref.(" + cast + ").GetUuid() {\n"
					refSetter += "			this.M_" + attribute.Name + ref + " = ref.(" + cast + ").GetUuid()\n"
					refSetter += "			this.NeedSave = true\n"
					refSetter += "		}\n"
				} else {
					hasUtility[packName+"."+ownerName] = true
					refSetter += "		if this.M_" + attribute.Name + ref + " != ref.(Utility.Referenceable).GetUuid() {\n"
					refSetter += "			this.M_" + attribute.Name + ref + " = ref.(Utility.Referenceable).GetUuid()\n"
					refSetter += "			this.NeedSave = true\n"
					refSetter += "		}\n"
				}

				refSetter += "		this.m_" + attribute.Name + ref + " = ref.(" + typeName + ")\n"

				refSetter += "	}\n"
			} else {
				refSetter += "	if this.M_" + attribute.Name + ref + " != ref.(" + typeName + ") {\n"
				refSetter += "		this.M_" + attribute.Name + ref + " = ref.(" + typeName + ")\n"
				refSetter += "		this.NeedSave = true\n"
				refSetter += "	}\n"
			}
		}
		refSetter += "}\n"
		methodStr += refSetter

		// Now the ref remover..
		/** Here I will create the function to initialyse a reference... **/
		methodStr += "\n/** Remove reference " + methodName + " **/\n"
		refRemover := ""
		if (isRef || attribute.IsComposite == "true") && classesMap[typeName_] != nil {

			refRemover += "func (this *" + ownerName + ") Remove" + methodName + ref + "(ref interface{}){\n"
			// Here the reference is an array...
			var isInterfaceCast bool
			if strings.Contains(cast, ".") {
				isInterfaceCast = Utility.Contains(abstractClassLst, strings.Split(cast, ".")[1]) || Utility.Contains(superClassesLst, strings.Split(cast, ".")[1])
			} else {
				isInterfaceCast = Utility.Contains(abstractClassLst, cast) || Utility.Contains(superClassesLst, cast)
			}

			if isPointer && !isInterfaceCast && !strings.HasPrefix(cast, "*") {
				cast = "*" + cast
			}

			refRemover += "	toDelete := ref.(" + cast + ")\n"

			if attribute.Upper == "*" {

				if Utility.Contains(abstractClassLst, typeName_) || Utility.Contains(superClassesLst, typeName_) {
					refRemover += "	" + attribute.Name + ref + "_ := make([]" + getClassPackName(typeName_, packName) + typeName_ + ", 0)\n"
				} else {
					refRemover += "	" + attribute.Name + ref + "_ := make([]*" + getClassPackName(typeName_, packName) + typeName_ + ", 0)\n"
				}

				if len(attribute.IsComposite) > 0 {
					refRemover += "	for i := 0; i < len(this.M_" + attribute.Name + ref + "); i++ {\n"

					if Utility.Contains(abstractClassLst, typeName_) || Utility.Contains(superClassesLst, typeName_) {
						refRemover += "		if toDelete.GetUuid() != this.M_" + attribute.Name + ref + "[i].(" + cast + ").GetUuid() {\n"
					} else {
						refRemover += "		if toDelete.GetUuid() != this.M_" + attribute.Name + ref + "[i].GetUuid() {\n"
					}
					refRemover += "			" + attribute.Name + ref + "_ = append(" + attribute.Name + ref + "_, this.M_" + attribute.Name + ref + "[i])\n"
					refRemover += "		}else{\n"
					refRemover += "			this.NeedSave = true\n"
					refRemover += "		}\n"
					refRemover += "	}\n"

					refRemover += "	this.M_" + attribute.Name + ref + " = " + attribute.Name + ref + "_\n"
				} else {
					refRemover += "	" + attribute.Name + ref + "Uuid := make([]string, 0)\n"

					refRemover += "	for i := 0; i < len(this.m_" + attribute.Name + ref + "); i++ {\n"
					if Utility.Contains(abstractClassLst, typeName_) || Utility.Contains(superClassesLst, typeName_) {
						refRemover += "		if toDelete.GetUuid() != this.m_" + attribute.Name + ref + "[i].(" + cast + ").GetUuid() {\n"
					} else {
						refRemover += "		if toDelete.GetUuid() != this.m_" + attribute.Name + ref + "[i].GetUuid() {\n"
					}

					refRemover += "			" + attribute.Name + ref + "_ = append(" + attribute.Name + ref + "_, this.m_" + attribute.Name + ref + "[i])\n"
					refRemover += "			" + attribute.Name + ref + "Uuid = append(" + attribute.Name + ref + "Uuid, this.M_" + attribute.Name + ref + "[i])\n"
					refRemover += "		}else{\n"
					refRemover += "			this.NeedSave = true\n"
					refRemover += "		}\n"
					refRemover += "	}\n"

					refRemover += "	this.m_" + attribute.Name + ref + " = " + attribute.Name + ref + "_\n"
					refRemover += "	this.M_" + attribute.Name + ref + " = " + attribute.Name + ref + "Uuid\n"
				}

			} else {

				if len(attribute.IsComposite) > 0 {
					// The attribute is not an array...
					if Utility.Contains(abstractClassLst, typeName_) || Utility.Contains(superClassesLst, typeName_) {
						refRemover += "	if toDelete.GetUuid() == this.M_" + attribute.Name + ref + ".(" + cast + ").GetUuid() {\n"
					} else {
						refRemover += "	if toDelete.GetUuid() == this.M_" + attribute.Name + ref + ".GetUuid() {\n"
					}
					refRemover += "		this.M_" + attribute.Name + ref + " = nil\n"
					refRemover += "		this.NeedSave = true\n"
				} else {
					// The attribute is not an array...
					refRemover += "	if this.m_" + attribute.Name + ref + "!= nil {\n"
					if Utility.Contains(abstractClassLst, typeName_) || Utility.Contains(superClassesLst, typeName_) {
						refRemover += "		if toDelete.GetUuid() == this.m_" + attribute.Name + ref + ".(" + cast + ").GetUuid() {\n"
					} else {
						refRemover += "		if toDelete.GetUuid() == this.m_" + attribute.Name + ref + ".GetUuid() {\n"
					}
					refRemover += "			this.m_" + attribute.Name + ref + " = nil\n"
					refRemover += "			this.M_" + attribute.Name + ref + " = \"\"\n"
					refRemover += "			this.NeedSave = true\n"
					refRemover += "		}\n"
				}
				refRemover += "	}\n"
			}
			refRemover += "}\n"
			methodStr += refRemover

		} else {
			//log.Println("-----------> no remover for ", attribute.Name)
		}
	} else {

		methodStr = "	/** " + methodName + " **/\n	" + getter + "\n"

		// Now if the property is public i will also generate it's setter.
		if attribute.Visibility == "public" {
			setter += "Set" + methodName + ref + "(interface{}) "
			methodStr += "	" + setter + "\n"
		}
	}

	return methodStr
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

	classStr += "	/** UUID **/\n"
	classStr += "	GetUuid() string\n\n"
	for j := 0; j < len(class.Attributes); j++ {
		attribute := class.Attributes[j]
		if attribute.Type == "cmof:Property" {
			classStr += generateGoMethodCode(&attribute, nil, false, membersPackage[class.Name])
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
				classStr += "	/** If the entity value has change... **/\n"
				classStr += "	NeedSave bool\n"
				classStr += "	/** Get entity by uuid function **/\n"
				classStr += "	getEntityByUuid func(string)(interface{}, error)\n\n"

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

				classStr += "/** Evaluate if an entity needs to be saved. **/\n"
				classStr += "func (this *" + className + ") IsNeedSave() bool{\n"
				classStr += "	return this.NeedSave\n"
				classStr += "}\n"
				classStr += "func (this *" + className + ") ResetNeedSave(){\n"
				classStr += "	this.NeedSave=false\n"
				classStr += "}\n\n"

				classStr += "/** Give access to entity manager GetEntityByUuid function from Entities package. **/\n"
				classStr += "func (this *" + className + ") SetEntityGetter(fct func(uuid string)(interface{}, error)){\n"
				classStr += "	this.getEntityByUuid = fct\n"
				classStr += "}\n"

				// Now the method...
				// The superclass methode...
				for j := 0; j < len(superClasses); j++ {
					superClass := classesMap[superClasses[j]]
					for k := 0; k < len(superClass.Attributes); k++ {
						classStr += generateGoMethodCode(&superClass.Attributes[k], class, false, membersPackage[class.Name])
					}
				}

				// The class methode
				for j := 0; j < len(class.Attributes); j++ {
					classStr += generateGoMethodCode(&class.Attributes[j], class, false, membersPackage[class.Name])
				}

				// The associations getter / setter
				for _, att := range associations {
					associationsStr := generateGoMethodCode(att, class, true, membersPackage[class.Name])
					if strings.Index(classStr, associationsStr) == -1 {
						classStr += associationsStr
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
