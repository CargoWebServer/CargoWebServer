package main

import (
	//	"log"
	"strings"

	"code.myceliUs.com/Utility"
	"code.myceliUs.com/XML_Schemas"
)

// Various maps that contain processed informations.
var elementsNameMap map[string]*XML_Schemas.XSD_Element
var elementsTypeMap map[string]*XML_Schemas.XSD_Element
var complexTypesMap map[string]*XML_Schemas.XSD_ComplexType
var simpleTypesMap map[string]*XML_Schemas.XSD_SimpleType
var xsdPrimitiveTypesMap map[string]string
var substitutionGroupMap map[string][]string
var aliasElementType map[string]string

//////////////////////////////////////////////////////////////////////////////
// Initialisation function.
//////////////////////////////////////////////////////////////////////////////

/**
 * Init the maps...
 */
func loadSchema(schema *XML_Schemas.XSD_Schema) {

	// The xsd primitives...
	// The forced element type is use when the
	// xml attribute name is not equivalent to it's type...

	// Those tow types are reference...
	xsdPrimitiveTypesMap["xsd:QName"] = "string"
	xsdPrimitiveTypesMap["xsd:IDREF"] = "string"
	xsdPrimitiveTypesMap["xsd:anyURI"] = "string"
	xsdPrimitiveTypesMap["xsd:ID"] = "string"

	// Basic primitive...
	xsdPrimitiveTypesMap["xsd:string"] = "string"
	xsdPrimitiveTypesMap["xsd:boolean"] = "bool"
	xsdPrimitiveTypesMap["xsd:integer"] = "int"
	xsdPrimitiveTypesMap["xsd:double"] = "float64"
	xsdPrimitiveTypesMap["xsd:date"] = "int64"
	xsdPrimitiveTypesMap["xsd:byte"] = "[]uint8"

	/////////////////////////////////////////////////////////////
	// Alias type...
	/////////////////////////////////////////////////////////////
	aliasElementType["inputDataItem"] = "DataInput"
	aliasElementType["outputDataItem"] = "DataOutput"
	aliasElementType["childLaneSet"] = "LaneSet"

	// Originaly Expression...
	aliasElementType["activationCondition"] = "FormalExpression"
	aliasElementType["from"] = "FormalExpression"
	aliasElementType["to"] = "FormalExpression"
	aliasElementType["loopCardinality"] = "FormalExpression"
	aliasElementType["completionCondition"] = "FormalExpression"
	aliasElementType["loopCondition"] = "FormalExpression"
	aliasElementType["conditionExpression"] = "FormalExpression"

	// Formal Expression...
	aliasElementType["transformation"] = "FormalExpression"
	aliasElementType["messagePath"] = "FormalExpression"
	aliasElementType["dataPath"] = "FormalExpression"
	aliasElementType["condition"] = "FormalExpression"
	aliasElementType["timeDate"] = "FormalExpression"
	aliasElementType["timeDuration"] = "FormalExpression"
	aliasElementType["timeCycle"] = "FormalExpression"

	// ImplicitThrowEvent...
	aliasElementType["event"] = "ImplicitThrowEvent"

	// The element
	for i := 0; i < len(schema.Elements); i++ {

		var elementType = getElementType(schema.Elements[i].Type)
		elementsTypeMap[elementType] = &schema.Elements[i]
		elementsNameMap[schema.Elements[i].Name] = &schema.Elements[i]
		if schema.Elements[i].IsAbstract == "true" {
			if Utility.Contains(superClassesLst, elementType) == false {
				superClassesLst = append(superClassesLst, elementType)
			}
			if Utility.Contains(abstractClassLst, elementType) == false {
				abstractClassLst = append(abstractClassLst, elementType)
			}
		}
		if len(schema.Elements[i].SubstitutionGroup) > 0 {
			// So here the values must be sean as the base type...
			substitutionGroupType := getElementType(schema.Elements[i].SubstitutionGroup)

			substitutionGroupType = strings.ToUpper(substitutionGroupType[0:1]) + substitutionGroupType[1:]

			/** Substitution group are asbstract **/
			if Utility.Contains(superClassesLst, substitutionGroupType) == false {
				superClassesLst = append(superClassesLst, substitutionGroupType)
			}

			/** The substitution group map **/
			if substitutionGroupMap[substitutionGroupType] == nil {
				substitutionGroupMap[substitutionGroupType] = make([]string, 0)
			}
			if Utility.Contains(substitutionGroupMap[substitutionGroupType], elementType) == false {
				substitutionGroupMap[substitutionGroupType] = append(substitutionGroupMap[substitutionGroupType], elementType)
			}
		}
	}

	// The complex type
	for i := 0; i < len(schema.ComplexTypes); i++ {
		var elementType = getElementType(schema.ComplexTypes[i].Name)
		complexTypesMap[elementType] = &schema.ComplexTypes[i]
		if schema.ComplexTypes[i].IsAbstract == true {
			if Utility.Contains(superClassesLst, elementType) == false {
				superClassesLst = append(superClassesLst, elementType)
			}

			if Utility.Contains(abstractClassLst, elementType) == false {
				abstractClassLst = append(abstractClassLst, elementType)
			}
		}
	}

	for _, complexType := range complexTypesMap {
		if complexType.ComplexContent != nil {
			if complexType.ComplexContent.Extension != nil {
				// if the complex type contain a base class...
				if len(complexType.ComplexContent.Extension.Base) > 0 {
					// So here the values must be sean as the base type...
					baseType := getElementType(complexType.ComplexContent.Extension.Base)
					if Utility.Contains(superClassesLst, baseType) == false {
						superClassesLst = append(superClassesLst, baseType)
					}

				}
			}
		}
	}

	// The simple type
	for i := 0; i < len(schema.SimpleTypes); i++ {
		var elementType = getElementType(schema.SimpleTypes[i].Name)
		simpleTypesMap[elementType] = &schema.SimpleTypes[i]
	}
}

func isIndex(className string, attributeName string) bool {
	xsdElement := complexTypesMap[className]

	if xsdElement != nil {
		attributes := xsdElement.Attributes
		for i := 0; i < len(attributes); i++ {
			if attributes[i].Name == attributeName {
				// Here I will try to see if the type is one of...
				// TODO append other type that can represent an index value...
				return strings.HasSuffix(attributes[i].Type, ":IDREF")
			}
		}
	}

	if strings.ToUpper(attributeName) == "NAME" {
		return true
	}

	return false
}

func isId(className string, attributeName string) bool {
	xsdElement := complexTypesMap[className]

	if xsdElement != nil {
		attributes := xsdElement.Attributes
		for i := 0; i < len(attributes); i++ {
			if attributes[i].Name == attributeName {
				// Here I will try to see if the type is one of...
				// TODO append other type that can represent an index value...
				return strings.HasSuffix(attributes[i].Type, ":ID")
			}
		}
	}

	return false
}

// Filter the element type name from unwanted prefix and namesapace...
func getElementType(elementType string) string {
	// Clean the element type...
	if strings.Index(elementType, ":") > -1 && len(xsdPrimitiveTypesMap[elementType]) == 0 {
		elementType = elementType[strings.Index(elementType, ":")+1:]
	}
	if strings.HasPrefix(elementType, "t") == true {
		elementType = elementType[1:]
	}
	return elementType
}
