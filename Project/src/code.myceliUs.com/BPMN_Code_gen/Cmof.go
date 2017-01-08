package main

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

/** Utility function **/
func contains(slice []string, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}

	_, ok := set[item]
	return ok
}

// Exists reports whether the named file or directory exists.
func Exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func WriteClassFile(outputPath, packName string, className string, classCode string) {

	if !Exists("../" + outputPath + packName) {
		os.MkdirAll("../"+outputPath+packName, 0777)
	}
	log.Println("Create file with name: ", className)
	if !Exists("../" + outputPath + packName + "/" + className + ".go") {
		err := ioutil.WriteFile("../"+outputPath+packName+"/"+className+".go", []byte(classCode), 0x777)
		if err != nil {
			panic(err)
		}
	}
}

/**
 * Here the simple cmof xml reader.
 */
type CMOF_Document struct {
	XMI     xml.Name     `xml:"XMI"`
	Version string       `xml:"version,attr"`
	XmiNs   string       `xml:"xmi,attr"`
	CmofNs  string       `xml:"cmof,attr"`
	Package CMOF_Package `xml:"Package"`
}

type CMOF_Package struct {
	XMLName      xml.Name           `xml:"Package"`
	Id           string             `xml:"id,attr"`
	Name         string             `xml:"name,attr"`
	Uri          string             `xml:"uri,attr"`
	OwnedMembers []CMOF_OwnedMember `xml:"ownedMember"`
}

type CMOF_OwnedMember struct {
	XMLName        xml.Name              `xml:"ownedMember"`
	SuperClass     string                `xml:"superClass,attr"`
	Type           string                `xml:"type,attr"`
	Id             string                `xml:"id,attr"`
	Name           string                `xml:"name,attr"`
	Attributes     []CMOF_OwnedAttribute `xml:"ownedAttribute"`
	Visibility     string                `xml:"visibility,attr"`
	End            CMOF_OwnedEnd         `xml:"ownedEnd"`
	Litterals      []CMOF_OwnedLiteral   `xml:"ownedLiteral"`
	SuperClassRefs []CMOF_SuperClass     `xml:"superClass"`
}

type CMOF_OwnedEnd struct {
	XMLName           xml.Name `xml:"ownedEnd"`
	Id                string   `xml:"id,attr"`
	Type              string   `xml:"_type,attr"`
	Name              string   `xml:"name,attr"`
	Association       string   `xml:"association,attr"`
	OwningAssociation string   `xml:"owningAssociation,attr"`
	Lower             string   `xml:"lower,attr"`
	Upper             string   `xml:"upper,attr"`
}

type CMOF_SuperClass struct {
	XMLName xml.Name `xml:"superClass"`
	Type    string   `xml:"type,attr"`
	Ref     string   `xml:"href,attr"`
}

type CMOF_OwnedLiteral struct {
	XMLName    xml.Name `xml:"ownedLiteral"`
	Id         string   `xml:"id,attr"`
	Type       string   `xml:"type,attr"`
	Name       string   `xml:"name,attr"`
	Classifier string   `xml:"classifier,attr"`
}

type CMOF_OwnedAttribute struct {
	XMLName            xml.Name  `xml:"ownedAttribute"`
	Type               string    `xml:"type,attr"`
	Id                 string    `xml:"id,attr"`
	Name               string    `xml:"name,attr"`
	Lower              string    `xml:"lower,attr"`
	Upper              string    `xml:"upper,attr"`
	IsComposite        string    `xml:"isComposite,attr"`
	Visibility         string    `xml:"visibility,attr"`
	AssociationId      string    `xml:"association,attr"`
	AssociationType    string    `xml:"_type,attr"`
	AssociationTypeRef CMOF_Type `xml:"type"`
}

type CMOF_Type struct {
	XMLName xml.Name `xml:"type"`
	Type    string   `xml:"type,attr"`
	Ref     string   `xml:"href,attr"`
}

// Those map contain reference between members.
var classesMap = make(map[string]*CMOF_OwnedMember)

// The enumeration map...
var enumMap = make(map[string]*CMOF_OwnedMember)

// The list of class in a package...
var packagesMembers = make(map[string][]string)

// The member pacakage...
var membersPackage = make(map[string]string)

// Contain the map of super class with it's members...
var superClassMembers = make(map[string]map[string]*CMOF_OwnedMember)

// The source of the association
var associationsSrcMap = make(map[string]*CMOF_OwnedMember)

// The target of the association
var associationsDstMap = make(map[string][]string)

// The primitive types...
var primitiveMapType = make(map[string]string)

// Must be abstract class...
var superClassesLst = make([]string, 0)

// Those class cannot be instanciated...
var abstractClassLst = make([]string, 0)

// If and entity is create by a composition relation I will keep
// the entity who is in charge of the creation...
var compositionEntityCreatorMap = make(map[string][]string)

func getClassAssociations(class *CMOF_OwnedMember) []*CMOF_OwnedAttribute {

	// Keep the association attribute...
	attributes := make([]*CMOF_OwnedAttribute, 0)
	associationNames := make([]string, 0)

	// The associations
	associations := associationsDstMap[class.Name]
	for index := range associations {
		srcAssociation := associationsSrcMap[associations[index]]
		if len(srcAssociation.End.Id) > 0 {
			var att CMOF_OwnedAttribute
			att.Visibility = srcAssociation.Visibility
			att.Name = srcAssociation.End.Name
			att.Lower = srcAssociation.End.Lower
			att.Upper = srcAssociation.End.Upper
			att.AssociationType = srcAssociation.End.Type
			if contains(associationNames, att.Name) == false {
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
			if contains(associationNames, superClassesAssociations[j].Name) == false {
				associationNames = append(associationNames, superClassesAssociations[j].Name)
				attributes = append(attributes, superClassesAssociations[j])
			}
		}
	}

	return attributes
}

func hasId(member *CMOF_OwnedMember) bool {

	return len(getAttributeType("id", member)) > 0
}

func hasName(member *CMOF_OwnedMember) bool {
	return len(getAttributeType("name", member)) > 0
}

/**
	Retreive the type of an attribute via is associated cmof type.
	look recursively in the member super class if the type is not found.
**/
func getAttributeType(attributeName string, member *CMOF_OwnedMember) string {

	for i := 0; i < len(member.Attributes); i++ {
		if member.Attributes[i].Name == attributeName {

			if len(member.Attributes[i].AssociationType) > 0 {
				return member.Attributes[i].AssociationType
			}
			if len(primitiveMapType[member.Attributes[i].Type]) > 0 {
				return primitiveMapType[member.Attributes[i].Type]
			}
			// Here the type must be considering a primitive type.
			if len(member.Attributes[i].AssociationTypeRef.Ref) > 0 {
				if len(primitiveMapType[member.Attributes[i].AssociationTypeRef.Ref]) > 0 {
					return primitiveMapType[member.Attributes[i].AssociationTypeRef.Ref]
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

func getAttribute(attributeName string, member *CMOF_OwnedMember) *CMOF_OwnedAttribute {

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
func getInterfaceByAttributeName(name string, member *CMOF_OwnedMember) string {
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

func getOwnedAttributeByName(name string, member *CMOF_OwnedMember) (*CMOF_OwnedAttribute, *CMOF_OwnedMember) {

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

func getAttributeTypeName(attribute *CMOF_OwnedAttribute) (string, bool, bool) {

	var typeName = attribute.AssociationType
	var isPrimitive = false
	var isPointer = false

	if len(typeName) == 0 {
		if attribute.AssociationTypeRef.Type == "cmof:PrimitiveType" {
			if len(attribute.AssociationTypeRef.Ref) > 0 {
				typeName = primitiveMapType[attribute.AssociationTypeRef.Ref]
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

				if contains(abstractClassLst, typeName) == false {
					isPointer = true
				}
			}
		}

	} else {
		if len(primitiveMapType[typeName]) > 0 {
			typeName = primitiveMapType[typeName]
			isPrimitive = true
		} else if contains(abstractClassLst, typeName) == false {
			isPointer = true
		}
	}

	// Set primitive is so...
	if typeName == "[]uint8" || typeName == "string" || typeName == "interface{}" || typeName == "int" || typeName == "bool" || typeName == "float64" || typeName == "int64" {
		isPrimitive = true
	}

	return typeName, isPrimitive, isPointer
}

func getAttributeNs(attribute *CMOF_OwnedAttribute) string {
	typeName, _, _ := getAttributeTypeName(attribute)
	return membersPackage[typeName]
}

func getClassImports(class *CMOF_OwnedMember) []string {
	imports := make([]string, 0)
	if contains(abstractClassLst, class.Name) == false {
		superClasses := getSuperClasses(class.Name)
		for i := 0; i < len(superClasses); i++ {
			superClassesImports := getClassImports(classesMap[superClasses[i]])
			for j := 0; j < len(superClassesImports); j++ {
				if !contains(imports, superClassesImports[j]) {
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
				if !contains(imports, ns) {
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
				if !contains(imports, ns) {
					imports = append(imports, ns)
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
				if contains(results, superSuperClasses[j]) == false {
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
			if contains(results, superSuperClasses[j]) == false {
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
	if !contains(abstractClassLst, className) {
		if contains(implementations, className) == false {
			implementations = append(implementations, className)
		}
	}
	childs := make(map[string]*CMOF_OwnedMember)
	if _, ok := superClassMembers[className]; ok {
		childs = superClassMembers[className]
	}
	if len(childs) > 0 {
		for child, _ := range childs {
			if contains(implementations, child) == false {
				implementations_ := getImplementationClasses(child)
				for i := 0; i < len(implementations_); i++ {
					if contains(implementations, implementations_[i]) == false {
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

/**
 * Init the maps...
 */
func loadDocument(doc *CMOF_Document) {

	// Set the primitive types here...
	primitiveMapType["http://schema.omg.org/spec/MOF/2.0/cmof.xml#String"] = "string"
	primitiveMapType["DC.cmof#String"] = "string"
	primitiveMapType["String"] = "string"
	primitiveMapType["Byte"] = "[]uint8"
	primitiveMapType["http://schema.omg.org/spec/MOF/2.0/cmof.xml#Boolean"] = "bool"
	primitiveMapType["http://www.w3.org/2001/XMLSchema#boolean"] = "bool"
	primitiveMapType["DC.cmof#Boolean"] = "bool"
	primitiveMapType["Boolean"] = "bool"
	primitiveMapType["http://schema.omg.org/spec/MOF/2.0/cmof.xml#Integer"] = "int"
	primitiveMapType["http://www.w3.org/2001/XMLSchema#integer"] = "int"
	primitiveMapType["http://www.w3.org/2001/XMLSchema#date"] = "int64"
	primitiveMapType["DC.cmof#Integer"] = "int"
	primitiveMapType["Integer"] = "int"
	primitiveMapType["http://schema.omg.org/spec/MOF/2.0/cmof.xml#Real"] = "float64"
	primitiveMapType["http://www.w3.org/2001/XMLSchema#double"] = "float64"
	primitiveMapType["DC.cmof#Real"] = "float64"
	primitiveMapType["Real"] = "float64"
	primitiveMapType["DI.cmof#DiagramElement-modelElement"] = "interface{}"
	primitiveMapType["http://schema.omg.org/spec/MOF/2.0/cmof.xml#Element"] = "interface{}"

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
					if contains(superClassesLst, superClasses[j]) == false {
						superClassesLst = append(superClassesLst, superClasses[j])
					}
					if superClassMembers[superClasses[j]] == nil {
						superClassMembers[superClasses[j]] = make(map[string]*CMOF_OwnedMember)
					}

					superClassMembers[superClasses[j]][member.Name] = &member
				}

			}
			for j := 0; j < len(member.SuperClassRefs); j++ {
				superClass := member.SuperClassRefs[j].Ref[strings.Index(member.SuperClassRefs[j].Ref, "#")+1:]
				if contains(superClassesLst, superClass) == false {
					superClassesLst = append(superClassesLst, superClass)
				}
				if superClassMembers[superClass] == nil {
					superClassMembers[superClass] = make(map[string]*CMOF_OwnedMember)
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
							if !contains(compositionEntityCreatorMap[implementations[j]], member.Id+"."+attribute.Id) {
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

func generateCode(packName string) {

	///////////////////////////////////////////////////////////////////////////
	// Go Go Go!!!
	// The interface...
	generateGoInterfacesCode(packName)

	// The class
	generateGoClassCode(packName)

	// The enumeration
	generateGoEnumerationCode(packName)

	///////////////////////////////////////////////////////////////////////////
	// Js Js Js!!!
	//generateJsClassCode(pack.Name)
}

func generateFactory(rootElementId string, packName string, outpoutPath string, name string) {

	// The go xml factory
	generateGoXmlFactory(rootElementId, packName, outpoutPath, name)

	// The JS xml factory
	//generateJsXmlFactory()

}
