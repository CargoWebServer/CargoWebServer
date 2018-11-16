package main

import (
	"log"
	"strconv"
	"strings"

	"code.myceliUs.com/Utility"
	"code.myceliUs.com/XML_Schemas"
)

// TODO
// implementation must be able to contain one enum and one uri...

/////////////////////////////////////////////////////////////////////////////////
// Factory
/////////////////////////////////////////////////////////////////////////////////

/** The list of import used by the factory... **/
var factoryImports []string

// That map will contain initialisation function for a given data type
var initFunctions map[string]string
var serializationFunctions map[string]string

/** The return the factory imports **/
func getFactoryImports() string {
	factoryStr := "import (\n"
	for i := 0; i < len(factoryImports); i++ {
		factoryStr += "	\"" + factoryImports[i] + "\"\n"
	}
	factoryStr += ")\n"
	return factoryStr
}

/**
 * Generate the Go xml factory
 */
func generateGoXmlFactory(rootElementId string, packName string, outputPath string, name string) {
	// The root element of the factory...
	rootElement := elementsNameMap[rootElementId]
	elementType := getElementType(rootElement.Type)
	elementPackName := membersPackage[elementType]
	elementPackName += "."

	// The now package to import...
	factoryImports = append(factoryImports, "golang.org/x/net/html/charset")
	factoryImports = append(factoryImports, "encoding/xml")
	factoryImports = append(factoryImports, "strings")
	factoryImports = append(factoryImports, "os")
	factoryImports = append(factoryImports, "log")
	factoryImports = append(factoryImports, "path/filepath")
	factoryImports = append(factoryImports, "code.myceliUs.com/Utility")

	// The factory is created in the BPMS package...
	factoryStr := "type " + name + "XmlFactory struct {\n"
	// The map of reference by id...
	factoryStr += "	m_references map[string] string\n"
	// That contain, the map of owner id -> the map of properties, the list of ref id to set.
	factoryStr += "	m_object map[string]map[string][]string\n"
	factoryStr += "	m_getEntityByUuid func(string)(interface{}, error)\n"
	factoryStr += "	m_setEntity func(interface{})\n"
	factoryStr += "	m_generateUuid func(interface{}) string\n"
	factoryStr += "}\n\n"

	factoryStr += "	/** Get entity by uuid function **/\n"
	factoryStr += "func (this *" + name + "XmlFactory) SetEntityGetter(fct func(string)(interface{}, error)){\n"
	factoryStr += "	this.m_getEntityByUuid = fct\n"
	factoryStr += "}\n\n"

	factoryStr += "	/** Use to put the entity in the cache **/\n"
	factoryStr += "func (this *" + name + "XmlFactory) SetEntitySetter(fct func(interface{})){\n"
	factoryStr += "	this.m_setEntity = fct\n"
	factoryStr += "}\n\n"

	factoryStr += "\n/** Set the uuid genarator function. **/\n"
	factoryStr += "func (this *" + name + "XmlFactory) SetUuidGenerator(fct func(entity interface{}) string){\n"
	factoryStr += "	this.m_generateUuid = fct\n"
	factoryStr += "}\n\n"

	// Initialisation function.
	factoryStr += "\n/** Initialization function from xml file **/\n"
	factoryStr += "func (this *" + name + "XmlFactory)InitXml(inputPath string, object *" + elementPackName + strings.ToUpper(elementType[0:1]) + elementType[1:] + ") error{\n"
	factoryStr += "	xmlFilePath, err := filepath.Abs(inputPath)\n"
	factoryStr += "	if err != nil {\n"
	factoryStr += "		log.Println(err)\n"
	factoryStr += "		os.Exit(1)\n"
	factoryStr += "	}\n"

	factoryStr += "	reader, err := os.Open(xmlFilePath)\n"
	factoryStr += "	if err != nil {\n"
	factoryStr += "		log.Println(err)\n"
	factoryStr += "		os.Exit(1)\n"
	factoryStr += "	}\n"

	factoryStr += "	var xmlElement *" + elementPackName + "Xsd" + strings.ToUpper(elementType[0:1]) + elementType[1:] + "\n"
	factoryStr += "	xmlElement = new(" + elementPackName + "Xsd" + strings.ToUpper(elementType[0:1]) + elementType[1:] + ")\n"
	factoryStr += "	decoder := xml.NewDecoder(reader)\n"
	factoryStr += "	decoder.CharsetReader = charset.NewReaderLabel\n"
	factoryStr += "	if err := decoder.Decode(xmlElement); err != nil {\n"
	factoryStr += "		return err\n"
	factoryStr += "	}\n"
	factoryStr += "	this.m_references = make(map[string] string, 0)\n"
	factoryStr += "	this.m_object = make(map[string]map[string][]string, 0)\n"
	// So here I will call the generate function...
	factoryStr += "	this.Init" + strings.ToUpper(elementType[0:1]) + elementType[1:] + "(\"\", xmlElement, object)\n"
	// Now I will set the reference inside the model...
	factoryStr += "	for ref0, refMap := range this.m_object {\n"
	factoryStr += "		refOwner, _:= this.m_getEntityByUuid(this.m_references[ref0])\n"
	factoryStr += "		if refOwner != nil {\n"
	factoryStr += "			for ref1, _ := range refMap {\n"
	factoryStr += "				refs := refMap[ref1]\n"
	factoryStr += "				for i:=0; i<len(refs); i++{\n"
	factoryStr += "					ref, _:= this.m_getEntityByUuid(this.m_references[refs[i]])\n"
	factoryStr += "					if  ref != nil {\n"
	factoryStr += "						_, err := Utility.CallMethod(refOwner, ref1, []interface{}{ref})\n"
	factoryStr += "						if err != nil {\n"
	factoryStr += "							Utility.CallMethod(refOwner, ref1, []interface{}{this.m_references[refs[i]]})\n"
	factoryStr += "						}\n"
	factoryStr += "						this.m_setEntity(refOwner)\n"
	factoryStr += "					}else{\n"
	factoryStr += "						ref, _:= this.m_getEntityByUuid(refs[i])\n"
	factoryStr += "						_, err := Utility.CallMethod(refOwner, ref1, []interface{}{ref})\n"
	factoryStr += "						if err != nil {\n"
	factoryStr += "							Utility.CallMethod(refOwner, ref1, []interface{}{refs[i]})\n"
	factoryStr += "							if err != nil {\n"
	factoryStr += "								Utility.CallMethod(refOwner,\"SetFieldValue\", []interface{}{\"M_\" + strings.ToLower(ref1[3:4]) + ref1[4:], refs[i]})\n"
	factoryStr += "							}\n"
	factoryStr += "						}\n"
	factoryStr += "						this.m_setEntity(refOwner)\n"
	factoryStr += "					}\n"
	factoryStr += "				}\n"
	factoryStr += "			}\n"
	factoryStr += "		}\n"
	factoryStr += "	}\n"
	factoryStr += "	return nil\n"
	factoryStr += "}\n\n"

	// Serialization function
	factoryStr += "\n/** Serialization to xml file **/\n"
	factoryStr += "func (this *" + name + "XmlFactory)SerializeXml(outputPath string, toSerialize *" + elementPackName + strings.ToUpper(elementType[0:1]) + elementType[1:] + ") error{\n"
	factoryStr += "	xmlFilePath, err := filepath.Abs(outputPath)\n"
	factoryStr += "	if err != nil {\n"
	factoryStr += "		log.Println(err)\n"
	factoryStr += "		os.Exit(1)\n"
	factoryStr += "	}\n"

	factoryStr += "	fo, err := os.Create(xmlFilePath)\n"
	factoryStr += "	defer func() {\n"
	factoryStr += "		if err := fo.Close(); err != nil {\n"
	factoryStr += "			panic(err)\n"
	factoryStr += "		}\n"
	factoryStr += "	}()\n\n"

	factoryStr += "	var xmlElement *" + elementPackName + "Xsd" + strings.ToUpper(elementType[0:1]) + elementType[1:] + "\n"
	factoryStr += "	xmlElement = new(" + elementPackName + "Xsd" + strings.ToUpper(elementType[0:1]) + elementType[1:] + ")\n\n"
	// So here I will call the generate function...
	factoryStr += "	this.Serialyze" + strings.ToUpper(elementType[0:1]) + elementType[1:] + "(xmlElement, toSerialize)\n"

	// Finaly I will wrote the file.
	factoryStr += "	output, err := xml.MarshalIndent(xmlElement, \"  \", \"    \")\n"
	factoryStr += "	if err != nil {\n"
	factoryStr += "		log.Println(err)\n"
	factoryStr += "		os.Exit(1)\n"
	factoryStr += "	}\n"
	factoryStr += "	fileContent := []byte(\"<?xml version=\\\"1.0\\\" encoding=\\\"UTF-8\\\" standalone=\\\"yes\\\"?>\\n\")\n"
	factoryStr += "	fileContent = append(fileContent, output...)\n"
	factoryStr += "	_, err = fo.Write(fileContent)\n"
	factoryStr += "	return nil\n"
	factoryStr += "}\n"

	// Generate the factory function recursively...
	generateGoXmlFactoryElementInitFunction(rootElement, packName, name)

	// Now print the reusult...
	for _, initFunctionSrc := range initFunctions {
		factoryStr += initFunctionSrc
	}

	generateGoXmlFactoryElementSeiralizeFunction(rootElement, packName, name)

	for _, serializationFunctionSrc := range serializationFunctions {
		factoryStr += serializationFunctionSrc
	}

	// Now of the root element I will generate the parser function...
	factoryStr = "package " + packName + "\n" + getFactoryImports() + factoryStr

	// The name of the package here will be the last dir of the path
	// and the output path will be the

	// Wrote the file...
	WriteClassFile(outputPath, "", name+"XmlFactory", factoryStr)
}

/////////////////////////////////////////////////////////////////////////////////
// Initialisation from xml
/////////////////////////////////////////////////////////////////////////////////

// Generate the init function for a given xsd element.
func generateGoXmlFactoryElementInitFunction(element *XML_Schemas.XSD_Element, packName string, name string) {

	elementFunctionParserStr := ""

	// Set the ref...
	if len(element.Ref) > 0 {
		if len(element.Type) > 0 {
			element = elementsTypeMap[getElementType(element.Type)]
		} else {
			element = elementsNameMap[getElementType(element.Ref)]
		}
	}

	if element == nil {
		return
	}

	elementType := getElementType(element.Type)
	if _, ok := initFunctions[elementType]; ok {
		// The function already exist...
		return
	}

	// Set empty string...
	initFunctions[elementType] = elementFunctionParserStr

	if Utility.Contains(abstractClassLst, elementType) {
		// I will generate the element for it base class...
		if substitutionGroupMap[elementType] != nil {
			for i := 0; i < len(substitutionGroupMap[elementType]); i++ {
				subElement := elementsTypeMap[substitutionGroupMap[elementType][i]]
				subElementType := getElementType(subElement.Type)
				if _, ok := initFunctions[subElementType]; !ok {
					generateGoXmlFactoryElementInitFunction(subElement, packName, name)
				}
			}
		}
	} else if complexTypesMap[elementType] != nil {
		/** The package to import if there is one... **/
		elementPackName := membersPackage[elementType]
		if len(elementPackName) > 0 {
			if elementPackName != packName {
				if Utility.Contains(factoryImports, outputPath+elementPackName) == false {
					factoryImports = append(factoryImports, outputPath+elementPackName)
				}
				elementPackName += "."
			} else {
				elementPackName = ""
			}
		}

		// Create the initialisation function here...
		elementFunctionParserStr += "\n/** inititialisation of " + elementType + " **/\n"
		className := strings.ToUpper(elementType[0:1]) + elementType[1:]

		// if the class is also a superClass that's mean the class name must ha the _impl suffix.
		elementFunctionParserStr += "func (this *" + name + "XmlFactory) Init" + className + "(parentUuid string, "

		elementFunctionParserStr += "xmlElement *" + elementPackName + "Xsd" + strings.ToUpper(elementType[0:1]) + elementType[1:]

		impl := ""
		if Utility.Contains(superClassesLst, elementType) && !Utility.Contains(abstractClassLst, elementType) {
			// Here the class is an implementation
			impl = "_impl"
		}

		elementFunctionParserStr += ",object *" + elementPackName + className + impl
		elementFunctionParserStr += "){\n"
		elementFunctionParserStr += "	log.Println(\"Initialize " + className + "\")\n"
		elementFunctionParserStr += "	object.TYPENAME = \"" + elementPackName + className + "\"\n"
		elementFunctionParserStr += "	object.SetParentUuid(parentUuid)\n"
		elementFunctionParserStr += "	object.SetUuidGenerator(this.m_generateUuid)\n"
		elementFunctionParserStr += "	object.SetEntitySetter(this.m_setEntity)\n"
		elementFunctionParserStr += "	object.SetEntityGetter(this.m_getEntityByUuid)\n"

		elementFunctionParserStr += generateGoXmlFactoryElementContent(elementType, elementType, packName, true, name)

		// A little exception here for the expression class...
		if elementType == "Expression" || elementType == "FormalExpression" {
			elementFunctionParserStr += "	/** other content **/\n"
			elementFunctionParserStr += "	exprStr := xmlElement.M_other\n"
			elementFunctionParserStr += "	object.SetOther(exprStr)\n"
		}

		elementFunctionParserStr += "}\n"
		initFunctions[elementType] = elementFunctionParserStr

	}

}

func generateGoXmlFactoryElementContent(elementType string, baseElementType string, packName string, isInitialisation bool, name string) string {

	elementParserFunctionContentStr := ""

	// So here I will get the complex element that define that element...
	complexElement := complexTypesMap[elementType]
	className := elementType
	member := classesMap[className]

	if complexElement != nil && member != nil {

		// Now I will create it content...
		if complexElement.ComplexContent != nil {
			if complexElement.ComplexContent.Extension != nil {
				// So here the values must be sean as the base type...
				baseType_ := getElementType(complexElement.ComplexContent.Extension.Base)

				// Reset the member to keep the original member...
				member := classesMap[baseElementType]
				if len(baseType_) > 0 {
					elementParserFunctionContentStr += generateGoXmlFactoryElementContent(baseType_, baseElementType, packName, isInitialisation, name)

					// Now it list of attributes...
					for i := 0; i < len(complexElement.ComplexContent.Extension.Attributes); i++ {
						attribute := complexElement.ComplexContent.Extension.Attributes[i]
						if len(attribute.Name) > 0 && member != nil {
							elementParserFunctionContentStr += generateGoXmlFactoryElementAttribute(&attribute, member, baseElementType, packName, baseElementType, isInitialisation)
						}
					}

					// Contain a sequences...
					if complexElement.ComplexContent.Extension.Sequence != nil {
						for i := 0; i < len(complexElement.ComplexContent.Extension.Sequence.Elements); i++ {
							subElement := complexElement.ComplexContent.Extension.Sequence.Elements[i]
							if isInitialisation == true {
								generateGoXmlFactoryElementInitFunction(&subElement, packName, name)
							} else {
								generateGoXmlFactoryElementSeiralizeFunction(&subElement, packName, name)
							}
							elementParserFunctionContentStr += generateGoXmlFactoryElementReference(elementType, member, &subElement, packName, isInitialisation, name)
						}
					}

					// Now the choise...
					if complexElement.ComplexContent.Extension.Choice != nil {
						choice := complexElement.ComplexContent.Extension.Choice
						// The choice can contain a sequence...
						for i := 0; i < len(choice.Sequences); i++ {
							for j := 0; j < len(choice.Sequences[i].Elements); j++ {
								subElement := choice.Sequences[i].Elements[j]
								if isInitialisation == true {
									generateGoXmlFactoryElementInitFunction(&subElement, packName, name)
								} else {
									generateGoXmlFactoryElementSeiralizeFunction(&subElement, packName, name)
								}
								elementParserFunctionContentStr += generateGoXmlFactoryElementReference(elementType, member, &subElement, packName, isInitialisation, name)
							}
						}
						for i := 0; i < len(choice.Elements); i++ {
							if isInitialisation == true {
								generateGoXmlFactoryElementInitFunction(&choice.Elements[i], packName, name)
							} else {
								generateGoXmlFactoryElementSeiralizeFunction(&choice.Elements[i], packName, name)
							}
							elementParserFunctionContentStr += generateGoXmlFactoryElementReference(elementType, member, &choice.Elements[i], packName, isInitialisation, name)
						}
					}
				}
			}
		}

		// The attributes
		if complexElement.Attributes != nil {
			// Now it list of complexElement...
			for i := 0; i < len(complexElement.Attributes); i++ {
				attribute := complexElement.Attributes[i]
				if len(attribute.Name) > 0 && member != nil {
					elementParserFunctionContentStr += generateGoXmlFactoryElementAttribute(&attribute, member, elementType, packName, baseElementType, isInitialisation)
				}
			}
		}

		// The any attributes
		if complexElement.AnyAttributes != nil {
			// The any Attribute...
			for i := 0; i < len(complexElement.AnyAttributes); i++ {
				attribute := complexElement.AnyAttributes[i]
				if len(attribute.Name) > 0 {
					elementParserFunctionContentStr += generateGoXmlFactoryElementAnyAttribute(&attribute, member, baseElementType, packName, isInitialisation)
				}
			}
		}

		if !Utility.Contains(abstractClassLst, elementType) {
			// If the attribute Name is id I will keep the objectect in the
			// the map...
			if hasId(member) {
				// Set the element ref by id...
				elementParserFunctionContentStr += "	if len(object.M_id) > 0 {\n"
				elementParserFunctionContentStr += "		this.m_references[object.M_id] = object.GetUuid()\n"
				elementParserFunctionContentStr += "	}\n"
			} else if hasName(member) {
				// Set the element ref by id...
				elementParserFunctionContentStr += "	if len(object.M_name) > 0 {\n"
				elementParserFunctionContentStr += "		this.m_references[object.M_name] = object.GetUuid()\n"
				elementParserFunctionContentStr += "	}\n"
			} else {
				// Set element ref by type
				elementParserFunctionContentStr += "	this.m_references[object.GetUuid()] = object.GetUuid()\n"
			}
		}

		// The sequence
		if complexElement.Sequence != nil {
			// Attribute contain in a sequences...
			for i := 0; i < len(complexElement.Sequence.Any); i++ {
				attribute := complexElement.Sequence.Any[i]
				if len(attribute.Name) > 0 {
					elementParserFunctionContentStr += generateGoXmlFactoryElementAny(&attribute, member, baseElementType, packName, isInitialisation)
				}
			}

			// Elements contained in a sequences...
			for i := 0; i < len(complexElement.Sequence.Elements); i++ {
				subElement := complexElement.Sequence.Elements[i]
				if isInitialisation == true {
					generateGoXmlFactoryElementInitFunction(&subElement, packName, name)
				} else {
					generateGoXmlFactoryElementSeiralizeFunction(&subElement, packName, name)
				}
				elementParserFunctionContentStr += generateGoXmlFactoryElementReference(elementType, member, &subElement, packName, isInitialisation, name)
			}
		}

	} else {
		log.Println("================> type not found...", className)
	}

	return elementParserFunctionContentStr
}

/////////////////////////////////////////////////////////////////////////////////
// Xml element
/////////////////////////////////////////////////////////////////////////////////

/**
 * Member intialisation.
 */
func generateGoXmlFactoryElementMember(ownerElementType string, elementBaseType string, relationName string, member *XML_Schemas.CMOF_OwnedMember, element *XML_Schemas.XSD_Element, packName string, memberIndex *int, minOccurs string, maxOccurs string, isInitialisation bool, implementationTypes *[]string, name string) string {

	elementType := getElementType(element.Type)

	// Get the cmof reference...
	attribute, _ := getOwnedAttributeByName(relationName, member)

	var isRef bool
	if attribute != nil {
		isRef = IsRef(attribute)
	}

	elementMemberInitStr := ""
	//if Utility.Contains(aliasElementName, relationName) {
	if val, ok := aliasElementType[relationName]; ok {
		log.Println("---------------> element whit type", val, "was found!!!")
		//elementType = val
	}

	if !isRef {
		classNameLst := substitutionGroupMap[elementType]
		for i := 0; i < len(classNameLst); i++ {
			if memberIndex == nil {
				memberIndex = new(int)
			}
			elementMemberInitStr += generateGoXmlFactoryElementMember(ownerElementType, elementBaseType, relationName, member, elementsTypeMap[classNameLst[i]], packName, memberIndex, minOccurs, maxOccurs, isInitialisation, implementationTypes, name)
		}
	}

	if !Utility.Contains(abstractClassLst, elementType) {
		if isInitialisation {
			elementMemberInitStr += "\n	/** Init " + element.Name + " **/\n"
		}

		// Here the xml contain the information about the objectect.
		if attribute != nil {
			attributeTypeName, _, _ := getAttributeTypeName(attribute)

			elementPackName0 := membersPackage[attributeTypeName]
			if len(elementPackName0) > 0 {
				if elementPackName0 != packName {
					if Utility.Contains(factoryImports, outputPath+elementPackName0) == false {
						factoryImports = append(factoryImports, outputPath+elementPackName0)
					}
					elementPackName0 += "."
				} else {
					elementPackName0 = ""
				}
			}

			elementPackName1 := membersPackage[elementType]
			if len(elementPackName1) > 0 {
				if elementPackName1 != packName {
					if Utility.Contains(factoryImports, outputPath+elementPackName1) == false {
						factoryImports = append(factoryImports, outputPath+elementPackName1)
					}
					elementPackName1 += "."
				} else {
					elementPackName1 = ""
				}
			}

			// Create the init function if none exist...
			if isInitialisation == true {
				generateGoXmlFactoryElementInitFunction(element, packName, name)
			} else {
				generateGoXmlFactoryElementSeiralizeFunction(element, packName, name)
			}

			relationName_ := relationName
			if memberIndex != nil {
				relationName_ += "_" + strconv.Itoa(*memberIndex)
			}

			impl := ""
			implCast := ""

			if Utility.Contains(superClassesLst, elementType) && !Utility.Contains(abstractClassLst, elementType) {
				impl = "_impl"
			}

			if Utility.Contains(superClassesLst, attributeTypeName) {
				implCast = ".(*" + elementPackName1 + elementType + impl + ")"
			}

			*implementationTypes = append(*implementationTypes, elementType)

			// Here if the alias type exist...
			aliasType := aliasElementType[relationName]
			if len(aliasType) > 0 {
				aliasType = "*" + elementPackName1 + "Xsd" + aliasType
			}

			if maxOccurs == "unbounded" || attribute.Upper == "*" {
				// The element is an array
				// So here I will create an array...
				// Initialyse the values here...
				if isInitialisation {
					// So now I will create the initialisation loop...
					if !isRef {
						elementMemberInitStr += "	for i:=0;i<len(xmlElement.M_" + relationName_ + "); i++{\n"
						elementMemberInitStr += "		if object.M_" + relationName + "==nil{\n"
						elementMemberInitStr += "			object.M_" + relationName + "= make([]string,0)\n"
						elementMemberInitStr += "		}\n"
						elementMemberInitStr += "		val:=new(" + elementPackName1 + elementType + impl + ")\n"
						elementMemberInitStr += "		this.Init" + elementType + "(object.GetUuid(), xmlElement.M_" + relationName_ + "[i],val)\n"
						elementMemberInitStr += "		object.Append" + strings.ToUpper(relationName[0:1]) + relationName[1:] + "(val)\n"

						elementMemberInitStr += generateGoXmlFactoryElementCompositeAssociation(relationName, ownerElementType, elementType, implCast, attribute, true)
						elementMemberInitStr += "	}\n"
					}
				}
			} else if minOccurs == "0" || (attribute.Upper == "1" && attribute.Lower == "0") {
				if isInitialisation {
					if !isRef {
						elementMemberInitStr += "	if xmlElement.M_" + relationName_ + "!= nil{\n"
						elementMemberInitStr += "		val := new(" + elementPackName1 + elementType + ")\n"
						elementMemberInitStr += "		this.Init" + elementType + "(object.GetUuid(), xmlElement.M_" + relationName_ + ", val)\n"
						elementMemberInitStr += "		object.Set" + strings.ToUpper(relationName[0:1]) + relationName[1:] + "(val)\n"
						elementMemberInitStr += generateGoXmlFactoryElementCompositeAssociation(relationName, ownerElementType, elementType, implCast, attribute, false)
						elementMemberInitStr += "	}\n"
					}
				}
			} else {
				// Single xsd parser not a pointer...
				relationName_ := relationName
				if memberIndex != nil {
					relationName_ += "_" + strconv.Itoa(*memberIndex)
				}
				if isInitialisation {
					if Utility.Contains(abstractClassLst, attributeTypeName) {
						if !isRef {
							elementMemberInitStr += "	if xmlElement.M_" + relationName_ + "!= nil{\n"
							elementMemberInitStr += "		val := new(" + elementPackName1 + elementType + impl + ")\n"
							elementMemberInitStr += "		this.Init" + elementType + "(object.GetUuid(), xmlElement.M_" + relationName_ + ", val)\n"
							elementMemberInitStr += "		object.Set" + strings.ToUpper(relationName[0:1]) + relationName[1:] + "(val)\n"
							elementMemberInitStr += generateGoXmlFactoryElementCompositeAssociation(relationName, ownerElementType, elementType, implCast, attribute, false)
							elementMemberInitStr += "	}\n"
						}
					} else if minOccurs == "0" || attribute.Lower == "0" {
						if !isRef {
							elementMemberInitStr += "	if xmlElement.M_" + relationName_ + "!= nil{\n"
							elementMemberInitStr += "		val := new(" + elementPackName1 + elementType + impl + ")\n"
							if len(aliasType) > 0 {
								elementMemberInitStr += "	this.Init" + elementType + "(object.GetUuid(), (" + aliasType + ")(unsafe.Pointer(xmlElement.M_" + relationName + implCast + ")), val)\n"
							} else {
								elementMemberInitStr += "	this.Init" + elementType + "(object.GetUuid(), xmlElement.M_" + relationName_ + ", val)\n"
							}
							elementMemberInitStr += "		object.Set" + strings.ToUpper(relationName[0:1]) + relationName[1:] + "(val)\n"
							elementMemberInitStr += generateGoXmlFactoryElementCompositeAssociation(relationName, ownerElementType, elementType, implCast, attribute, false)
							elementMemberInitStr += "	}\n"
						}
					} else {
						if !isRef {
							elementMemberInitStr += "		val := new(" + elementPackName1 + elementType + impl + ")\n"
							if len(aliasType) > 0 {
								elementMemberInitStr += "	this.Init" + elementType + "(object.GetUuid(), (" + aliasType + ")(unsafe.Pointer(&xmlElement.M_" + relationName_ + ")), val)\n"
							} else {
								elementMemberInitStr += "	this.Init" + elementType + "(object.GetUuid(), &xmlElement.M_" + relationName_ + ", val)\n"
							}
							elementMemberInitStr += "		object.Set" + strings.ToUpper(relationName[0:1]) + relationName[1:] + "(val)\n"
							elementMemberInitStr += generateGoXmlFactoryElementCompositeAssociation(relationName, ownerElementType, elementType, implCast, attribute, false)
						}
					}

				}
			}

			if memberIndex != nil {
				*memberIndex++
			}
		} else {
			log.Println("================> Relation  ", relationName, " not found in ", member.Name)
		}
	}

	return elementMemberInitStr
}

/////////////////////////////////////////////////////////////////////////////////
// Xml Attributes
/////////////////////////////////////////////////////////////////////////////////
func generateGoXmlFactorySimpleType(simpleType *XML_Schemas.XSD_SimpleType, attribute *XML_Schemas.XSD_Attribute, simpleTypeName string, elementPackName string, isInitialisation bool) string {
	simpleTypesStr := "\n"
	for i := 0; i < len(simpleType.Restriction.Enumeration); i++ {
		if isInitialisation {
			if i == 0 {
				simpleTypesStr += "	if "
			} else if i < len(simpleType.Restriction.Enumeration) {
				simpleTypesStr += " else if "
			}
			enumValue := strings.Replace(simpleType.Restriction.Enumeration[i].Value, "##", "", -1)
			enumValue = strings.ToUpper(enumValue[0:1]) + enumValue[1:]
			if i < len(simpleType.Restriction.Enumeration) {
				simpleTypesStr += "xmlElement." + "M_" + attribute.Name + "==\"##" + strings.Replace(simpleType.Restriction.Enumeration[i].Value, "##", "", -1) + "\"{\n"
			}

			simpleTypesStr += "		object.M_" + attribute.Name + "=" + elementPackName + simpleTypeName + "_" + enumValue + "\n" //ItemKind_Physical
			simpleTypesStr += "	}"
			if i == len(simpleType.Restriction.Enumeration)-1 {
				// The the default value...
				if len(attribute.Default) > 0 {
					enumValue = strings.Replace(attribute.Default, "##", "", -1)
					enumValue = strings.ToUpper(enumValue[0:1]) + enumValue[1:]
					simpleTypesStr += " else{\n"
					simpleTypesStr += "		object.M_" + attribute.Name + "=" + elementPackName + simpleTypeName + "_" + enumValue + "\n" //ItemKind_Physical
					simpleTypesStr += " 	}\n"
				} else {
					simpleTypesStr += "\n"
				}
			}
		} else {
			if i == 0 {
				simpleTypesStr += "	if "
			} else if i < len(simpleType.Restriction.Enumeration) {
				simpleTypesStr += " else if "
			}
			enumValue := strings.Replace(simpleType.Restriction.Enumeration[i].Value, "##", "", -1)
			enumValue = strings.ToUpper(enumValue[0:1]) + enumValue[1:]
			if i < len(simpleType.Restriction.Enumeration) {
				simpleTypesStr += "object." + "M_" + attribute.Name + "==" + elementPackName + simpleTypeName + "_" + enumValue + "{\n"
			}

			simpleTypesStr += "		xmlElement." + "M_" + attribute.Name + "=\"##" + strings.Replace(simpleType.Restriction.Enumeration[i].Value, "##", "", -1) + "\"\n"
			simpleTypesStr += "	}"

			// Set the default value here.
			if i == len(simpleType.Restriction.Enumeration)-1 {
				// The the default value...
				if len(attribute.Default) > 0 {
					simpleTypesStr += " else{\n"
					simpleTypesStr += "		xmlElement." + "M_" + attribute.Name + "=\"##" + strings.Replace(attribute.Default, "##", "", -1) + "\"\n"
					simpleTypesStr += " 	}\n"
				} else {
					simpleTypesStr += "\n"
				}
				//simpleTypesStr += "\n"
			}
		}
	}
	return simpleTypesStr
}

// Generate the attribute intialisation string.
func generateGoXmlFactoryElementAttribute(attribute *XML_Schemas.XSD_Attribute, member *XML_Schemas.CMOF_OwnedMember, elementType string, packName string, elementType_ string, isInitialisation bool) string {
	attributeInitialisationStr := ""
	elementPackName := membersPackage[elementType_]
	if len(elementPackName) > 0 {
		if elementPackName != packName {
			if Utility.Contains(factoryImports, outputPath+elementPackName) == false {
				factoryImports = append(factoryImports, outputPath+elementPackName)
			}
			elementPackName += "."
		} else {
			elementPackName = ""
		}
	}

	if len(attribute.Name) > 0 {
		attributeType := getElementType(attribute.Type)
		if attributeType == "xsd:QName" || attributeType == "xsd:IDREF" {
			// log.Println("Reference initialisation here with name: ", attribute.Name)
			// So here the attribute is a reference
			attributeInitialisationStr += generateGoXmlFactoryElementRef(attribute.Name, attributeType, false, false, member, isInitialisation, packName)
		} else {
			//log.Println("Object initialisation here with name: ", attribute.Name)
			if simpleTypesMap[attributeType] != nil {
				// Here the enumerations...
				simpleType := simpleTypesMap[attributeType]
				simpleTypeName := getElementType(simpleType.Name)

				if simpleType.Union != nil {
					if !isInitialisation {
						attributeInitialisationStr += "	if len(object.M_" + attribute.Name + "Str)>0{\n"
						attributeInitialisationStr += "		xmlElement." + "M_" + attribute.Name + "=object.M_" + attribute.Name + "Str\n"
						attributeInitialisationStr += " }else{\n"
					} else {
						attributeInitialisationStr += "	if !strings.HasPrefix(xmlElement." + "M_" + attribute.Name + ",\"##\"){\n"
						attributeInitialisationStr += "		object.M_" + attribute.Name + "Str=xmlElement.M_" + attribute.Name + "\n"
						attributeInitialisationStr += "	}else{\n"
					}
					for i := 0; i < len(simpleType.Union.SimpleType); i++ {
						attributeInitialisationStr += generateGoXmlFactorySimpleType(&simpleType.Union.SimpleType[i], attribute, simpleTypeName, elementPackName, isInitialisation)
					}
					attributeInitialisationStr += "	}\n"

				} else {
					attributeInitialisationStr += generateGoXmlFactorySimpleType(simpleType, attribute, simpleTypeName, elementPackName, isInitialisation)
				}

			} else {
				// here this is simple attribute.
				// I will made the conversion if the type need one...
				primitiveType := getAttributeType(attribute.Name, member)
				if len(primitiveType) > 0 {
					attributeInitialisationStr += "\n	/** " + elementType + " **/\n"
					if isInitialisation {
						attributeInitialisationStr += "	object.M_" + attribute.Name + "= xmlElement." + "M_" + attribute.Name + "\n"
						if attribute.Name == "id" {
							attributeInitialisationStr += "	if len(object.M_" + attribute.Name + ")==0 {\n"
							attributeInitialisationStr += "		object.Set" + strings.ToUpper(attribute.Name[0:1]) + attribute.Name[1:] + "(Utility.RandomUUID())\n"
							attributeInitialisationStr += "	}\n"
						}
					} else {
						attributeInitialisationStr += "	xmlElement.M_" + attribute.Name + "= object." + "M_" + attribute.Name + "\n"
					}
				}
			}
		}
	}
	return attributeInitialisationStr
}

func generateGoXmlFactoryElementAnyAttribute(attribute *XML_Schemas.XSD_AnyAttribute, member *XML_Schemas.CMOF_OwnedMember, elementType string, packName string, isInitialisation bool) string {
	attributeInitialisationStr := ""
	primitiveType := getAttributeType(attribute.Name, member)
	if len(attribute.Name) > 0 {
		if attribute.Name == "other" {
			attributeInitialisationStr += "//"

		}
		if isInitialisation {
			attributeInitialisationStr += "	object.M_" + attribute.Name + "= xmlElement." + "M_" + attribute.Name + "\n"
		} else {
			if primitiveType == "interface{}" {
				attributeInitialisationStr += "	xmlElement.M_" + attribute.Name + "= object." + "M_" + attribute.Name + ".(string)\n"
			} else {
				attributeInitialisationStr += "	xmlElement.M_" + attribute.Name + "= object." + "M_" + attribute.Name + "\n"
			}
		}
	}
	return attributeInitialisationStr
}

func generateGoXmlFactoryElementAny(attribute *XML_Schemas.XSD_Any, member *XML_Schemas.CMOF_OwnedMember, elementType string, packName string, isInitialisation bool) string {
	attributeInitialisationStr := ""
	primitiveType := getAttributeType(attribute.Name, member)
	if len(attribute.Name) > 0 {
		if attribute.Name == "other" {
			attributeInitialisationStr += "//"

		}

		if isInitialisation {
			attributeInitialisationStr += "	object.M_" + attribute.Name + "= xmlElement." + "M_" + attribute.Name + "\n"
		} else {
			if primitiveType == "interface{}" {
				attributeInitialisationStr += "	xmlElement.M_" + attribute.Name + "= object." + "M_" + attribute.Name + ".(string)\n"
			} else {
				attributeInitialisationStr += "	xmlElement.M_" + attribute.Name + "= object." + "M_" + attribute.Name + "\n"
			}
		}

	}
	return attributeInitialisationStr
}

/////////////////////////////////////////////////////////////////////////////////
// Xml References
/////////////////////////////////////////////////////////////////////////////////

// Generate the initialisation code for reference.
func generateGoXmlFactoryElementRef(attributeName string, attributeType string, isRefPointer bool, isRefArray bool, member *XML_Schemas.CMOF_OwnedMember, isInitialisation bool, packName string) string {

	// I will make an exception for bmpnElement here...
	if attributeName == "bpmnElement" {
		bpmnElementStr := ""
		if isInitialisation {
			bpmnElementStr = "	/** Init bpmnElement **/\n"
			bpmnElementStr += "	if len(xmlElement.M_bpmnElement)>0{\n"
			bpmnElementStr += "		if _, ok:= this.m_object[object.M_id]; !ok {\n"
			bpmnElementStr += "			this.m_object[object.M_id]=make(map[string][]string)\n"
			bpmnElementStr += "		}\n"
			bpmnElementStr += "		if _, ok:= this.m_object[object.M_id][\"SetBpmnElement\"]; !ok {\n"
			bpmnElementStr += "			this.m_object[object.M_id][\"SetBpmnElement\"]=make([]string,0)\n"
			bpmnElementStr += "		}\n"
			bpmnElementStr += "		this.m_object[object.M_id][\"SetBpmnElement\"] = append(this.m_object[object.M_id][\"SetBpmnElement\"], xmlElement.M_bpmnElement)\n"
			bpmnElementStr += "	}\n"
		} else {
			bpmnElementStr = "	/** bpmnElement **/\n"
			bpmnElementStr += "	xmlElement.M_bpmnElement = object.M_bpmnElement\n"
		}
		return bpmnElementStr
	}
	mapMemberId := ""
	if hasId(member) {
		// Set with it id...
		mapMemberId = "object.M_id"
	} else if hasName(member) {
		// Set with it name
		mapMemberId = "object.M_name"
	} else {
		// If there is no id i will set with the
		mapMemberId = "\"" + attributeType + "\""
	}
	attributeInitialisationStr := ""

	if isInitialisation {
		// Now I will set the value...
		// attributeInitialisationStr += "	object.M_" + attributeName + " = xmlElement.M_" + attributeName + "\n"
		attributeInitialisationStr += "\n	/** Init ref " + attributeName + " **/\n"
		attributeInitialisationStr += "	if len(" + mapMemberId + ") == 0 {\n"
		attributeInitialisationStr += "		this.m_references[object.GetUuid()] = object.GetUuid()\n"

		attributeInitialisationStr += "	}\n"
		attributeInitialisationStr += "	if _, ok:= this.m_object[" + mapMemberId + "]; !ok {\n"
		attributeInitialisationStr += "		this.m_object[" + mapMemberId + "]=make(map[string][]string)\n"
		attributeInitialisationStr += "	}\n"
		// Now I will set the value...
		if isRefArray {
			attributeInitialisationStr += "	for i:=0; i < len(xmlElement.M_" + attributeName + "); i++ {\n"
		} else {
			if isRefPointer == true {
				attributeInitialisationStr += "	if xmlElement.M_" + attributeName + " !=nil {\n"
			} else {
				attributeInitialisationStr += "	if len(xmlElement.M_" + attributeName + ") > 0 {\n"
			}
		}

		// Set the value
		if isRefArray {
			attributeInitialisationStr += "		if _, ok:= this.m_object[" + mapMemberId + "][\"Append" + strings.ToUpper(attributeName[0:1]) + attributeName[1:] + "\"]; !ok {\n"
			attributeInitialisationStr += "			this.m_object[" + mapMemberId + "][\"Append" + strings.ToUpper(attributeName[0:1]) + attributeName[1:] + "\"]=make([]string,0)\n"
			attributeInitialisationStr += "		}\n"
			attributeInitialisationStr += "		this.m_object[" + mapMemberId + "][\"Append" + strings.ToUpper(attributeName[0:1]) + attributeName[1:] + "\"] = append(this.m_object[" + mapMemberId + "][\"Append" + strings.ToUpper(attributeName[0:1]) + attributeName[1:] + "\"], xmlElement.M_" + attributeName + "[i])\n"
		} else {
			attributeInitialisationStr += "		if _, ok:= this.m_object[" + mapMemberId + "][\"Set" + strings.ToUpper(attributeName[0:1]) + attributeName[1:] + "\"]; !ok {\n"
			attributeInitialisationStr += "			this.m_object[" + mapMemberId + "][\"Set" + strings.ToUpper(attributeName[0:1]) + attributeName[1:] + "\"]=make([]string,0)\n"
			attributeInitialisationStr += "		}\n"
			if isRefPointer == true {
				attributeInitialisationStr += "		this.m_object[" + mapMemberId + "][\"Set" + strings.ToUpper(attributeName[0:1]) + attributeName[1:] + "\"] = append(this.m_object[" + mapMemberId + "][\"Set" + strings.ToUpper(attributeName[0:1]) + attributeName[1:] + "\"], *xmlElement.M_" + attributeName + ")\n"
			} else {
				attributeInitialisationStr += "		this.m_object[" + mapMemberId + "][\"Set" + strings.ToUpper(attributeName[0:1]) + attributeName[1:] + "\"] = append(this.m_object[" + mapMemberId + "][\"Set" + strings.ToUpper(attributeName[0:1]) + attributeName[1:] + "\"], xmlElement.M_" + attributeName + ")\n"
			}
		}
		attributeInitialisationStr += "	}\n"
	} else {
		// Serialization...
		attributeInitialisationStr += "\n	/** Serialyze ref " + attributeName + " **/\n"
		if isRefPointer {
			attributeInitialisationStr += "	xmlElement.M_" + attributeName + "=&object.M_" + attributeName + "\n"
		} else {
			attributeInitialisationStr += "	xmlElement.M_" + attributeName + "=object.M_" + attributeName + "\n"
		}
	}
	return attributeInitialisationStr
}

/**
 * Generate the reference type.
 */
func generateGoXmlFactoryElementReference(ownerElementType string, member *XML_Schemas.CMOF_OwnedMember, element *XML_Schemas.XSD_Element, packName string, isInitialisation bool, name string) string {

	elementReferenceStr := ""
	maxOccurs := ""
	minOccurs := ""
	elementRefName := ""

	isRef := false

	if len(element.Ref) > 0 {
		elementRefName = element.Ref
		if strings.Index(elementRefName, ":") > -1 {
			elementRefName = getElementType(elementRefName)
		}
		maxOccurs = element.MaxOccurs
		minOccurs = element.MinOccurs
		if len(element.Type) > 0 {
			isRef = true
		}
	} else {
		elementRefName = element.Name
	}

	elementType := getElementType(element.Type)
	attribute, _ := getOwnedAttributeByName(elementRefName, member)
	if attribute != nil {
		isRef = IsRef(attribute)
	} else {
		attribute = getAttribute(elementRefName, member)
	}
	// The reference is outside the memeber...
	if elementType == "xsd:IDREF" || elementType == "xsd:QName" || isRef {
		attributeType := getAttributeType(elementRefName, member)

		isArray := attribute.Upper == "*"

		isPointer := !isArray && (attribute.Lower == "0")

		if maxOccurs != "" {
			isArray = maxOccurs == "unbounded"
		}

		if minOccurs != "" {
			isPointer = (minOccurs == "0") && !isArray
		}

		// Here the reference is a string containing the id of the objectect...
		// Array...
		elementReferenceStr += generateGoXmlFactoryElementRef(elementRefName, attributeType, isPointer, isArray, member, isInitialisation, packName) + "\n"

	} else {
		elementType = getAttributeType(elementRefName, member)

		element = elementsTypeMap[elementType]
		if element == nil {
			element = elementsNameMap[elementRefName]
		}

		if element != nil {
			implementationTypes := new([]string)
			elementReferenceStr += generateGoXmlFactoryElementMember(ownerElementType, elementType, elementRefName, member, element, packName, nil, minOccurs, maxOccurs, isInitialisation, implementationTypes, name)
			if !isInitialisation {
				attribute, _ := getOwnedAttributeByName(elementRefName, member)

				if attribute != nil {
					attributeTypeName, _, _ := getAttributeTypeName(attribute)
					isArray := attribute.Upper == "*"
					isPointer := !isArray && (attribute.Lower == "0")

					if maxOccurs != "" {
						isArray = maxOccurs == "unbounded"
					}

					if minOccurs != "" {
						isPointer = (minOccurs == "0") && !isArray
					}

					// Generate the serialysation
					elementReferenceStr += generateGoXmlElementSerialyze(elementRefName, attributeTypeName, packName, isArray, isPointer, implementationTypes)

					elementReferenceStr += generateGoXmlElementSerialyzeInit(ownerElementType, attributeTypeName, elementRefName, isArray, isPointer, implementationTypes, packName)
				} else {
					log.Println("===============> Attribute ", elementRefName, " ", elementType, " not found!", " ", member)
				}
			}

		} else {

			log.Println("===============> Element ", elementType, " not found!", " ", member)
		}
	}
	return elementReferenceStr
}

/////////////////////////////////////////////////////////////////////////////////
// Xml Associations
/////////////////////////////////////////////////////////////////////////////////
/**
 * Set the association in memory.
 */
func generateGoXmlFactoryElementCompositeAssociation(relationName string, ownerElementType string, elementType string, implCast string, attribute *XML_Schemas.CMOF_OwnedAttribute, isArray bool) string {
	elementMemberInitStr := ""

	if attribute.IsComposite == "true" {
		elementMemberInitStr += "\n		/** association initialisation **/\n"
		associationId := strings.ToLower(ownerElementType[0:1]) + ownerElementType[1:]
		associations := getClassAssociations(classesMap[elementType])
		var association *XML_Schemas.CMOF_OwnedAttribute
		for i := 0; i < len(associations); i++ {
			if associations[i].Name == associationId {
				association = associations[i]
				break
			}
		}
		if association != nil {
			associationTypeName, _, _ := getAttributeTypeName(association)
			isArray = association.Upper == "*"
			if associationTypeName == ownerElementType {
				// In that case I will set association...
				if isArray == true {
					elementMemberInitStr += "		val.Append" + strings.ToUpper(associationId[0:1]) + associationId[1:] + "Ptr(object)\n"
				} else {
					elementMemberInitStr += "		val.Set" + strings.ToUpper(associationId[0:1]) + associationId[1:] + "Ptr(object)\n"
				}
			} else {
				log.Println("==============> bad relation type, expected ", ownerElementType, " found ", associationTypeName)
			}
		}
	}
	return elementMemberInitStr
}

/////////////////////////////////////////////////////////////////////////////////
// Serialization to xml
/////////////////////////////////////////////////////////////////////////////////

/**
 * Generate the serialization code to xml.
 */
func generateGoXmlFactoryElementSeiralizeFunction(element *XML_Schemas.XSD_Element, packName string, name string) {

	elementFunctionStr := ""

	// Set the ref...
	if len(element.Ref) > 0 {
		var elementRef = getElementType(element.Ref)
		element = elementsNameMap[elementRef]
	}

	if element == nil {
		return
	}

	elementType := getElementType(element.Type)
	if _, ok := serializationFunctions[elementType]; ok {
		// The function already exist...
		return
	}

	// Set empty string...
	serializationFunctions[elementType] = elementFunctionStr

	if Utility.Contains(abstractClassLst, elementType) {
		// I will generate the element for it base class...
		if substitutionGroupMap[elementType] != nil {
			for i := 0; i < len(substitutionGroupMap[elementType]); i++ {
				subElement := elementsTypeMap[substitutionGroupMap[elementType][i]]
				subElementType := getElementType(subElement.Type)
				if _, ok := serializationFunctions[subElementType]; !ok {
					generateGoXmlFactoryElementSeiralizeFunction(subElement, packName, name)
				}
			}
		}
	} else if complexTypesMap[elementType] != nil {
		/** The package to import if there is one... **/
		elementPackName := membersPackage[elementType]
		if len(elementPackName) > 0 {
			if elementPackName != packName {
				if Utility.Contains(factoryImports, outputPath+elementPackName) == false {
					factoryImports = append(factoryImports, outputPath+elementPackName)
				}
				elementPackName += "."
			} else {
				elementPackName = ""
			}
		}

		// Create the initialisation function here...
		elementFunctionStr += "\n/** serialysation of " + elementType + " **/\n"
		className := strings.ToUpper(elementType[0:1]) + elementType[1:]

		// if the class is also a superClass that's mean the class name must ha the _impl suffix.
		elementFunctionStr += "func (this *" + name + "XmlFactory) Serialyze" + className + "("
		elementFunctionStr += "xmlElement *" + elementPackName + "Xsd" + strings.ToUpper(elementType[0:1]) + elementType[1:]

		impl := ""
		if Utility.Contains(superClassesLst, elementType) && !Utility.Contains(abstractClassLst, elementType) {
			// Here the class is an implementation
			impl = "_impl"
		}

		elementFunctionStr += ",object *" + elementPackName + className + impl
		elementFunctionStr += "){\n"

		// If the xml element is nil I will return empty string...
		elementFunctionStr += "	if xmlElement == nil{\n"
		elementFunctionStr += "		return\n"
		elementFunctionStr += "	}\n"

		elementFunctionStr += generateGoXmlFactoryElementContent(elementType, elementType, packName, false, name)

		// A little exception here for the expression class...
		if elementType == "Expression" || elementType == "FormalExpression" {
			elementFunctionStr += "	/** other content **/\n"
			elementFunctionStr += "	exprStr := object.GetOther()\n"
			elementFunctionStr += "	xmlElement.M_other = exprStr\n"
		}

		elementFunctionStr += "}\n"
		serializationFunctions[elementType] = elementFunctionStr
	}
}

func generateGoXmlElementSerialyzeInit(ownerElementType string, relationType string, relationName string, isArray bool, isPointer bool, implementationTypes *[]string, packName string) string {

	// So here I will create the initialisation code...
	xmlElementSerialyzeStr := "\n	/** Now I will save the value of " + relationName + " **/\n"

	// Here I will create array to keep value before the serialyse function is call
	if isArray {
		xmlElementSerialyzeStr += "	for i:=0; i<len(object.M_" + relationName + ");i++{\n"
	} else {
		xmlElementSerialyzeStr += "	if object.Get" + strings.ToUpper(relationName[0:1]) + relationName[1:] + "()!=nil{\n"
	}

	// Here I will make a type switch to get the good value to init...
	if len(*implementationTypes) == 1 && !Utility.Contains(superClassesLst, (*implementationTypes)[0]) {
		elementType := (*implementationTypes)[0]
		impl := ""
		if Utility.Contains(superClassesLst, elementType) {
			impl = "_impl"
		}
		cast := ""
		suffix := ""
		if relationType != elementType {
			cast = ".(*" + getClassPackName(elementType, packName) + elementType + ")"
			suffix = "_0"
		}
		serializeFunctionName := "this.Serialyze" + elementType

		aliasType := aliasElementType[relationName]
		if len(aliasType) > 0 {
			aliasType = "*" + getClassPackName(elementType, packName) + "Xsd" + aliasType
		}

		if isArray {
			xmlElementSerialyzeStr += "		xmlElement.M_" + relationName + "=append(xmlElement.M_" + relationName + ",new(" + getClassPackName(elementType, packName) + "Xsd" + elementType + "))\n"
			xmlElementSerialyzeStr += "		" + serializeFunctionName + "(xmlElement.M_" + relationName + "[i],object.Get" + strings.ToUpper(relationName[0:1]) + relationName[1:] + impl + "()[i]" + cast + ")\n"
		} else if isPointer {
			if len(aliasType) > 0 {
				xmlElementSerialyzeStr += "		" + serializeFunctionName + "((" + aliasType + ")(unsafe.Pointer(xmlElement.M_" + relationName + suffix + ")),object.Get" + strings.ToUpper(relationName[0:1]) + relationName[1:] + impl + cast + "())\n"
			} else {
				xmlElementSerialyzeStr += "		" + serializeFunctionName + "(xmlElement.M_" + relationName + suffix + ",object.Get" + strings.ToUpper(relationName[0:1]) + relationName[1:] + impl + cast + "())\n"
			}
		} else {
			if len(aliasType) > 0 {
				xmlElementSerialyzeStr += "		" + serializeFunctionName + "((" + aliasType + ")(unsafe.Pointer(&xmlElement.M_" + relationName + suffix + ")),object.Get" + strings.ToUpper(relationName[0:1]) + relationName[1:] + impl + cast + "())\n"
			} else {
				xmlElementSerialyzeStr += "		" + serializeFunctionName + "(&xmlElement.M_" + relationName + suffix + ",object.Get" + strings.ToUpper(relationName[0:1]) + relationName[1:] + impl + cast + "())\n"
			}
		}
	} else {
		if isArray {
			xmlElementSerialyzeStr += "		switch v:= object.Get" + strings.ToUpper(relationName[0:1]) + relationName[1:] + "()[i].(type){\n"
		} else {
			xmlElementSerialyzeStr += "		switch v:= object.Get" + strings.ToUpper(relationName[0:1]) + relationName[1:] + "().(type){\n"
		}
		for i := 0; i < len(*implementationTypes); i++ {
			elementType := (*implementationTypes)[i]
			impl := ""
			if Utility.Contains(superClassesLst, elementType) {
				impl = "_impl"
			}
			implementationType := "*" + getClassPackName(elementType, packName) + elementType + impl
			xmlElementSerialyzeStr += "			case " + implementationType + ":\n"

			serializeFunctionName := "this.Serialyze" + elementType

			if isArray {
				xmlElementSerialyzeStr += "				xmlElement.M_" + relationName + "_" + strconv.Itoa(i) + "=append(xmlElement.M_" + relationName + "_" + strconv.Itoa(i) + ",new(" + getClassPackName(elementType, packName) + "Xsd" + elementType + "))\n"
				xmlElementSerialyzeStr += "				" + serializeFunctionName + "(xmlElement.M_" + relationName + "_" + strconv.Itoa(i) + "[len(xmlElement.M_" + relationName + "_" + strconv.Itoa(i) + ")-1],v)\n"
			} else {
				xmlElementSerialyzeStr += "				xmlElement.M_" + relationName + "_" + strconv.Itoa(i) + "=new(" + getClassPackName(elementType, packName) + "Xsd" + elementType + ")\n"
				xmlElementSerialyzeStr += "				" + serializeFunctionName + "(xmlElement.M_" + relationName + "_" + strconv.Itoa(i) + ",v)\n"
			}
			//xmlElementSerialyzeStr += "				break\n"
			//xmlElementSerialyzeStr += "				log.Println(\"Serialyze " + ownerElementType + ":" + relationName + ":" + elementType + "\")\n"
		}
		// Reset the list of types.
		xmlElementSerialyzeStr += "		}\n"
	}

	xmlElementSerialyzeStr += "	}\n"
	return xmlElementSerialyzeStr
}

func generateGoXmlElementSerialyze(relationName string, elementType string, packName string, isArray bool, isPointer bool, implementationTypes *[]string) string {
	xmlElementSerialyzeStr := ""
	xmlElementSerialyzeStr += "\n	/** Serialyze " + elementType + " **/\n"

	aliasType := aliasElementType[relationName]
	if len(aliasType) > 0 {
		aliasType = getClassPackName(elementType, packName) + "Xsd" + strings.ToUpper(relationName[0:1]) + relationName[1:]
	}

	relationName_ := relationName
	for i := 0; i < len(*implementationTypes); i++ {

		if len(*implementationTypes) > 0 {
			if Utility.Contains(superClassesLst, elementType) {
				relationName = relationName_ + "_" + strconv.Itoa(i)
			}
		}

		if isArray {
			xmlElementSerialyzeStr += "	if len(object.M_" + relationName_ + ") > 0 {\n"
			xmlElementSerialyzeStr += "		xmlElement.M_" + relationName + "= make([]*" + getClassPackName((*implementationTypes)[i], packName) + "Xsd" + (*implementationTypes)[i] + ",0)\n"
			xmlElementSerialyzeStr += "	}\n"
		} else if isPointer {
			xmlElementSerialyzeStr += "	if object.Get" + strings.ToUpper(relationName_[0:1]) + relationName_[1:] + "()!= nil {\n"
			if len(aliasType) > 0 {
				xmlElementSerialyzeStr += "		xmlElement.M_" + relationName + "= new(" + aliasType + ")\n"
			} else {
				xmlElementSerialyzeStr += "		xmlElement.M_" + relationName + "= new(" + getClassPackName((*implementationTypes)[i], packName) + "Xsd" + (*implementationTypes)[i] + ")\n"
			}

			xmlElementSerialyzeStr += "	}\n"
		} // Else nothing todo...
	}

	return xmlElementSerialyzeStr
}
