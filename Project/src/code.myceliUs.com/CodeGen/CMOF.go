package main

import (
	//"log"
	"strings"

	"code.myceliUs.com/Utility"
	"code.myceliUs.com/XML_Schemas"
)

// Various maps that contain processed informations.

// Those map contain reference between members.
var classesMap map[string]*XML_Schemas.CMOF_OwnedMember

// The enumeration map...
var enumMap map[string]*XML_Schemas.CMOF_OwnedMember

// The list of class in a package...
var packagesMembers map[string][]string

// The member pacakage...
var membersPackage map[string]string

// Contain the map of super class with it's members...
var superClassMembers map[string]map[string]*XML_Schemas.CMOF_OwnedMember

// The source of the association
var associationsSrcMap map[string]*XML_Schemas.CMOF_OwnedMember

// The target of the association
var associationsDstMap map[string][]string

// The primitive types...
var cmofPrimitiveTypesMap map[string]string

// Must be abstract class...
var superClassesLst []string

// Those class cannot be instanciated...
var abstractClassLst []string

// If and entity is create by a composition relation I will keep
// the entity who is in charge of the creation...
var compositionEntityCreatorMap map[string][]string

//////////////////////////////////////////////////////////////////////////////
// Initialisation function.
//////////////////////////////////////////////////////////////////////////////

/**
 * Init the maps...
 */
func loadDocument(doc *XML_Schemas.CMOF_Document) {

	// primitive type mapping.
	cmofPrimitiveTypesMap["http://schema.omg.org/spec/MOF/2.0/cmof.xml#String"] = "string"
	cmofPrimitiveTypesMap["DC.cmof#String"] = "string"
	cmofPrimitiveTypesMap["String"] = "string"
	cmofPrimitiveTypesMap["Byte"] = "[]uint8"
	cmofPrimitiveTypesMap["http://schema.omg.org/spec/MOF/2.0/cmof.xml#Boolean"] = "bool"
	cmofPrimitiveTypesMap["http://www.w3.org/2001/XMLSchema#boolean"] = "bool"
	cmofPrimitiveTypesMap["DC.cmof#Boolean"] = "bool"
	cmofPrimitiveTypesMap["Boolean"] = "bool"
	cmofPrimitiveTypesMap["http://schema.omg.org/spec/MOF/2.0/cmof.xml#Integer"] = "int"
	cmofPrimitiveTypesMap["http://www.w3.org/2001/XMLSchema#integer"] = "int"
	cmofPrimitiveTypesMap["http://www.w3.org/2001/XMLSchema#date"] = "int64"
	cmofPrimitiveTypesMap["DC.cmof#Integer"] = "int"
	cmofPrimitiveTypesMap["Integer"] = "int"
	cmofPrimitiveTypesMap["http://schema.omg.org/spec/MOF/2.0/cmof.xml#Real"] = "float64"
	cmofPrimitiveTypesMap["http://www.w3.org/2001/XMLSchema#double"] = "float64"
	cmofPrimitiveTypesMap["DC.cmof#Real"] = "float64"
	cmofPrimitiveTypesMap["Real"] = "float64"
	cmofPrimitiveTypesMap["DI.cmof#DiagramElement-modelElement"] = "interface{}"
	cmofPrimitiveTypesMap["http://schema.omg.org/spec/MOF/2.0/cmof.xml#Element"] = "interface{}"

	pack := doc.Package
	packagesMembers[pack.Name] = make([]string, 0)

	// Now i will iterate over the classes...
	for i := 0; i < len(pack.OwnedMembers); i++ {
		member := pack.OwnedMembers[i]
		if member.Type == "cmof:Class" || member.Type == "cmof:DataType" {
			packagesMembers[pack.Name] = append(packagesMembers[pack.Name], member.Name)
			membersPackage[member.Name] = pack.Name
			classesMap[member.Id] = &member
			if len(member.SuperClass) > 0 {
				superClasses := strings.Split(member.SuperClass, " ")
				for j := 0; j < len(superClasses); j++ {
					if Utility.Contains(superClassesLst, superClasses[j]) == false {
						superClassesLst = append(superClassesLst, superClasses[j])
					}
					if superClassMembers[superClasses[j]] == nil {
						superClassMembers[superClasses[j]] = make(map[string]*XML_Schemas.CMOF_OwnedMember)
					}

					superClassMembers[superClasses[j]][member.Name] = &member
				}
			}
			for j := 0; j < len(member.SuperClassRefs); j++ {
				superClass := member.SuperClassRefs[j].Ref[strings.Index(member.SuperClassRefs[j].Ref, "#")+1:]
				if Utility.Contains(superClassesLst, superClass) == false {
					superClassesLst = append(superClassesLst, superClass)
				}
				if superClassMembers[superClass] == nil {
					superClassMembers[superClass] = make(map[string]*XML_Schemas.CMOF_OwnedMember)
				}
				superClassMembers[superClass][member.Name] = &member
			}
		} else if member.Type == "cmof:Association" {
			associationsSrcMap[member.Id] = &member
		} else if member.Type == "cmof:Enumeration" {
			packagesMembers[pack.Name] = append(packagesMembers[pack.Name], member.Name)
			enumMap[member.Id] = &member
			membersPackage[member.Name] = pack.Name
		}
	}

	// Here I will set the destination of association...
	for i := 0; i < len(pack.OwnedMembers); i++ {
		member := pack.OwnedMembers[i]
		if member.Type == "cmof:Class" {

			for j := 0; j < len(member.Attributes); j++ {
				attribute := member.Attributes[j]
				if attribute.Type == "cmof:Property" {
					// Here I will retreive the source association...
					associationsSrc := associationsSrcMap[attribute.AssociationId]
					if associationsSrc != nil {
						if len(attribute.AssociationType) > 0 {
							associationsSrcIds := associationsDstMap[attribute.AssociationType]
							if associationsSrcIds == nil {
								associationsSrcIds = make([]string, 0)
							}
							associationsSrcIds = append(associationsSrcIds, associationsSrc.Id)
							associationsDstMap[attribute.AssociationType] = associationsSrcIds
						}
					}
					if attribute.IsComposite == "true" {
						implementations := getImplementationClasses(attribute.AssociationType)
						for j := 0; j < len(implementations); j++ {
							if _, ok := compositionEntityCreatorMap[implementations[j]]; !ok {
								//I will create the array of parent here...
								compositionEntityCreatorMap[implementations[j]] = make([]string, 0)
							}
							// Append the parent in the list if is not already exist...
							if !Utility.Contains(compositionEntityCreatorMap[implementations[j]], member.Id+"."+attribute.Id) {
								// So here i will find the destination side attribute and get it name...
								association := associationsSrcMap[attribute.AssociationId]
								if association != nil {
									compositionEntityCreatorMap[implementations[j]] = append(compositionEntityCreatorMap[implementations[j]], association.End.Name+"::"+member.Id)
								}
							}
						}
					}
				}
			}
		}
	}

}

//////////////////////////////////////////////////////////////////////////////
// Is/Has functions
//////////////////////////////////////////////////////////////////////////////
func hasId(member *XML_Schemas.CMOF_OwnedMember) bool {

	return len(getAttributeType("id", member)) > 0
}

func hasName(member *XML_Schemas.CMOF_OwnedMember) bool {
	return len(getAttributeType("name", member)) > 0
}

func IsRef(attribute *XML_Schemas.CMOF_OwnedAttribute) bool {

	var typeName, isPrimitive /*isPointer*/, _ = getAttributeTypeName(attribute)

	/** Exception here **/
	if attribute.Name == "bpmnElement" {
		return true
	}

	if typeName == "interface{}" {
		return true
	}

	if isPrimitive {
		return false
	}

	if _, ok := simpleTypesMap[typeName]; ok {
		return false
	}

	isRef := strings.HasSuffix(attribute.Name, "Refs") || strings.HasSuffix(attribute.Name, "Ref") || typeName == "interface{}"

	if !isRef {
		// If the member is an association if is not composite, that mean
		// it must be a reference.
		isRef = attribute.IsComposite != "true"
	}

	return isRef
}

//////////////////////////////////////////////////////////////////////////////
// Getter functions
//////////////////////////////////////////////////////////////////////////////
func getClassAssociations(class *XML_Schemas.CMOF_OwnedMember) []*XML_Schemas.CMOF_OwnedAttribute {

	// Keep the association attribute...
	attributes := make([]*XML_Schemas.CMOF_OwnedAttribute, 0)
	associationNames := make([]string, 0)

	// The associations
	associations := associationsDstMap[class.Name]
	for index := range associations {
		srcAssociation := associationsSrcMap[associations[index]]
		if len(srcAssociation.End.Id) > 0 {
			var att XML_Schemas.CMOF_OwnedAttribute
			att.Visibility = srcAssociation.Visibility
			att.Name = srcAssociation.End.Name
			att.Lower = srcAssociation.End.Lower
			att.Upper = srcAssociation.End.Upper
			att.AssociationType = srcAssociation.End.Type
			if Utility.Contains(associationNames, att.Name) == false {
				associationNames = append(associationNames, att.Name)
				attributes = append(attributes, &att)
			}
		}
	}

	// Now the super class...
	superClasses := getSuperClasses(class.Name)
	for i := 0; i < len(superClasses); i++ {
		superClassesAssociations := getClassAssociations(classesMap[superClasses[i]])
		for j := 0; j < len(superClassesAssociations); j++ {
			if Utility.Contains(associationNames, superClassesAssociations[j].Name) == false {
				associationNames = append(associationNames, superClassesAssociations[j].Name)
				attributes = append(attributes, superClassesAssociations[j])
			}
		}
	}

	return attributes
}

func getClassPackName(className string, packName string) string {

	classPackName := membersPackage[className]
	if len(classPackName) > 0 {
		if classPackName != packName {
			if Utility.Contains(factoryImports, outputPath+classPackName) == false {
				factoryImports = append(factoryImports, outputPath+classPackName)
			}
			classPackName += "."
		} else {
			classPackName = ""
		}
	}
	return classPackName
}

func getAbstractClassNameByAttributeName(attributeName string, member *XML_Schemas.CMOF_OwnedMember, packName string) string {
	abstractClassName := getInterfaceByAttributeName(attributeName, member)
	if len(abstractClassName) > 0 {
		abstractClassName = getClassPackName(abstractClassName, packName) + abstractClassName
	}
	return abstractClassName
}

/**
	Retreive the type of an attribute via is associated cmof type.
	look recursively in the member super class if the type is not found.
**/
func getAttributeType(attributeName string, member *XML_Schemas.CMOF_OwnedMember) string {

	for i := 0; i < len(member.Attributes); i++ {
		if member.Attributes[i].Name == attributeName {

			if len(member.Attributes[i].AssociationType) > 0 {
				return member.Attributes[i].AssociationType
			}
			if len(cmofPrimitiveTypesMap[member.Attributes[i].Type]) > 0 {
				return cmofPrimitiveTypesMap[member.Attributes[i].Type]
			}
			// Here the type must be considering a primitive type.
			if len(member.Attributes[i].AssociationTypeRef.Ref) > 0 {
				if len(cmofPrimitiveTypesMap[member.Attributes[i].AssociationTypeRef.Ref]) > 0 {
					return cmofPrimitiveTypesMap[member.Attributes[i].AssociationTypeRef.Ref]
				} else {
					return getElementType(strings.Replace(member.Attributes[i].AssociationTypeRef.Ref, ".cmof#", ":", -1))
				}
			}
		}
	}

	// Here I will try to get from it parent class...
	superClasses := getSuperClasses(member.Name)
	for i := 0; i < len(superClasses); i++ {
		elementType := getAttributeType(attributeName, classesMap[superClasses[i]])
		if len(elementType) > 0 {
			return elementType
		}
	}

	return ""
}

func getAttribute(attributeName string, member *XML_Schemas.CMOF_OwnedMember) *XML_Schemas.CMOF_OwnedAttribute {

	for i := 0; i < len(member.Attributes); i++ {
		if member.Attributes[i].Name == attributeName {
			return &member.Attributes[i]
		}
	}

	// Here I will try to get from it parent class...
	superClasses := getSuperClasses(member.Name)
	for i := 0; i < len(superClasses); i++ {
		attribute := getAttribute(attributeName, classesMap[superClasses[i]])
		if attribute != nil {
			return attribute
		}
	}

	return nil
}

/**
 * Return the name of the interface that own a property of a given name.
 */
func getInterfaceByAttributeName(name string, member *XML_Schemas.CMOF_OwnedMember) string {
	for i := 0; i < len(member.Attributes); i++ {
		if member.Attributes[i].Name == name {
			return member.Name
		}
	}

	// Here I will try to get from it parent class...
	superClasses := getSuperClasses(member.Name)
	for i := 0; i < len(superClasses); i++ {
		attribute, _ := getOwnedAttributeByName(name, classesMap[superClasses[i]])
		if attribute != nil {
			return superClasses[i]
		}
	}

	return ""
}

func getOwnedAttributeByName(name string, member *XML_Schemas.CMOF_OwnedMember) (*XML_Schemas.CMOF_OwnedAttribute, *XML_Schemas.CMOF_OwnedMember) {

	for i := 0; i < len(member.Attributes); i++ {
		if member.Attributes[i].Name == name {
			return &member.Attributes[i], member
		}
	}

	// Here I will try to get from it parent class...
	superClasses := getSuperClasses(member.Name)
	for i := 0; i < len(superClasses); i++ {
		attribute, owner := getOwnedAttributeByName(name, classesMap[superClasses[i]])
		if attribute != nil {
			// also set the parent class that owned that attribute.
			return attribute, owner
		}
	}
	return nil, member
}

func getAttributeTypeName(attribute *XML_Schemas.CMOF_OwnedAttribute) (string, bool, bool) {

	var typeName = attribute.AssociationType
	var isPrimitive = false
	var isPointer = false

	if len(typeName) == 0 {
		if attribute.AssociationTypeRef.Type == "cmof:PrimitiveType" {
			if len(attribute.AssociationTypeRef.Ref) > 0 {
				typeName = cmofPrimitiveTypesMap[attribute.AssociationTypeRef.Ref]
			}
		} else if attribute.AssociationTypeRef.Type == "cmof:DataType" || attribute.AssociationTypeRef.Type == "cmof:Class" {
			if attribute.AssociationTypeRef.Ref == "http://schema.omg.org/spec/MOF/2.0/cmof.xml#Element" {
				typeName = "interface{}"
			} else {
				// Remove the  part...
				tmpStr := strings.Replace(attribute.AssociationTypeRef.Ref, "cmof#", "", -1)
				if strings.Index(tmpStr, ".") > -1 {
					typeName = tmpStr[strings.Index(tmpStr, ".")+1:]
				}

				if Utility.Contains(abstractClassLst, typeName) == false {
					isPointer = true
				}
			}
		}

	} else {
		if len(cmofPrimitiveTypesMap[typeName]) > 0 {
			typeName = cmofPrimitiveTypesMap[typeName]
			isPrimitive = true
		} else if Utility.Contains(abstractClassLst, typeName) == false {
			isPointer = true
		}
	}

	// Set primitive is so...
	if typeName == "[]uint8" || typeName == "string" || typeName == "interface{}" || typeName == "int" || typeName == "bool" || typeName == "float64" || typeName == "int64" {
		isPrimitive = true
	}

	return typeName, isPrimitive, isPointer
}

func getAttributeNs(attribute *XML_Schemas.CMOF_OwnedAttribute) string {
	typeName, _, _ := getAttributeTypeName(attribute)
	return membersPackage[typeName]
}

func getClassImports(class *XML_Schemas.CMOF_OwnedMember) []string {
	imports := make([]string, 0)
	if Utility.Contains(abstractClassLst, class.Name) == false {
		superClasses := getSuperClasses(class.Name)
		for i := 0; i < len(superClasses); i++ {
			superClassesImports := getClassImports(classesMap[superClasses[i]])
			for j := 0; j < len(superClassesImports); j++ {
				if !Utility.Contains(imports, superClassesImports[j]) {
					imports = append(imports, superClassesImports[j])
				}
			}
		}

		// The super class refs...
		for i := 0; i < len(class.SuperClassRefs); i++ {
			ns := class.SuperClassRefs[i].Ref[0:strings.Index(class.SuperClassRefs[i].Ref, ".")]
			// Set the correct folder latter...
			if len(ns) > 0 {
				ns = outputPath + ns
				if !Utility.Contains(imports, ns) {
					imports = append(imports, ns)
				}
			}
		}
	}

	for i := 0; i < len(class.Attributes); i++ {
		attribute := class.Attributes[i]
		ns := getAttributeNs(&attribute)
		// Set the correct folder latter...
		if len(ns) > 0 {
			if ns != membersPackage[class.Name] {
				ns = outputPath + ns
				if !Utility.Contains(imports, ns) {
					imports = append(imports, ns)
				}
				// recursive call...
				if classesMap[attribute.Name] != nil {
					imports_ := getClassImports(classesMap[attribute.Name])
					for j := 0; j < len(imports_); j++ {
						if !Utility.Contains(imports, imports_[j]) {
							imports = append(imports, imports_[j])
						}
					}
				}
			}
		}
	}
	return imports
}

/**
 * Recursive function to get the list of all super classes.
 */
func getSuperClasses(className string) []string {
	class := classesMap[className]
	results := make([]string, 0)
	if len(class.SuperClass) > 0 {
		superClasses := strings.Split(class.SuperClass, " ")
		for i := 0; i < len(superClasses); i++ {
			superSuperClasses := getSuperClasses(superClasses[i])
			for j := 0; j < len(superSuperClasses); j++ {
				if Utility.Contains(results, superSuperClasses[j]) == false {
					results = append(results, superSuperClasses[j])
				}
			}
			results = append(results, superClasses[i])
		}
	}

	for i := 0; i < len(class.SuperClassRefs); i++ {
		superClass := class.SuperClassRefs[i].Ref[strings.Index(class.SuperClassRefs[i].Ref, "#")+1:]
		superSuperClasses := getSuperClasses(superClass)
		for j := 0; j < len(superSuperClasses); j++ {
			if Utility.Contains(results, superSuperClasses[j]) == false {
				results = append(results, superSuperClasses[j])
			}
		}
		results = append(results, superClass)
	}

	return results
}

func getImplementationClasses(className string) []string {

	implementations := make([]string, 0)
	// If the class name is not abstract...
	if !Utility.Contains(abstractClassLst, className) {
		if Utility.Contains(implementations, className) == false {
			implementations = append(implementations, className)
		}
	}
	childs := make(map[string]*XML_Schemas.CMOF_OwnedMember)
	if _, ok := superClassMembers[className]; ok {
		childs = superClassMembers[className]
	}
	if len(childs) > 0 {
		for child, _ := range childs {
			if Utility.Contains(implementations, child) == false {
				implementations_ := getImplementationClasses(child)
				for i := 0; i < len(implementations_); i++ {
					if Utility.Contains(implementations, implementations_[i]) == false {
						if len(implementations_[i]) > 0 {
							implementations = append(implementations, implementations_[i])
						}
					}
				}
			}
		}
	}
	return implementations
}
