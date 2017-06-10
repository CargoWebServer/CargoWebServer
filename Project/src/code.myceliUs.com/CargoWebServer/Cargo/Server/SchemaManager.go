package Server

import (
	"bytes"
	"encoding/xml"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"

	"code.myceliUs.com/CargoWebServer/Cargo/Entities/CargoEntities"
	"code.myceliUs.com/CargoWebServer/Cargo/Entities/Config"
	"code.myceliUs.com/Utility"
	"code.myceliUs.com/XML_Schemas"
	"github.com/kokardy/saxlike"
)

////////////////////////////////////////////////////////////////////////////////
// The schemas manager.
////////////////////////////////////////////////////////////////////////////////
type SchemaManager struct {

	// Map of schema
	schemas map[string]*XML_Schemas.XSD_Schema

	// To get the schema id for a given element...
	elementSchema map[string]string

	// Map of attributes...
	globalAttributes     map[string]*XML_Schemas.XSD_Attribute
	globalAnyAttributes  map[string]*XML_Schemas.XSD_AnyAttribute
	globalAttributeGroup map[string]*XML_Schemas.XSD_AttributeGroup

	// Map of compositor object...
	globalSimpleTypes  map[string]*XML_Schemas.XSD_SimpleType
	globalComplexTypes map[string]*XML_Schemas.XSD_ComplexType
	globalGroups       map[string]*XML_Schemas.XSD_Group
	globalElements     map[string]*XML_Schemas.XSD_Element

	// Map of xsd primitives...
	xsdPrimitiveTypesMap map[string]string

	// list of Prototypes...
	prototypes map[string]*EntityPrototype
}

var schemaManager *SchemaManager

func (this *Server) GetSchemaManager() *SchemaManager {
	if schemaManager == nil {
		schemaManager = newSchemaManager()
	}
	return schemaManager
}

func newSchemaManager() *SchemaManager {

	schemaManager := new(SchemaManager)

	// Creation of maps...
	schemaManager.schemas = make(map[string]*XML_Schemas.XSD_Schema)
	schemaManager.globalAttributes = make(map[string]*XML_Schemas.XSD_Attribute)
	schemaManager.globalAnyAttributes = make(map[string]*XML_Schemas.XSD_AnyAttribute)
	schemaManager.globalAttributeGroup = make(map[string]*XML_Schemas.XSD_AttributeGroup)
	schemaManager.globalSimpleTypes = make(map[string]*XML_Schemas.XSD_SimpleType)
	schemaManager.globalComplexTypes = make(map[string]*XML_Schemas.XSD_ComplexType)
	schemaManager.globalGroups = make(map[string]*XML_Schemas.XSD_Group)
	schemaManager.globalElements = make(map[string]*XML_Schemas.XSD_Element)
	schemaManager.elementSchema = make(map[string]string)

	// Set the primitive type...
	schemaManager.xsdPrimitiveTypesMap = make(map[string]string)

	// Here I will keep track of prototypes...
	schemaManager.prototypes = make(map[string]*EntityPrototype)

	// Basic primitive...
	schemaManager.xsdPrimitiveTypesMap["NCName"] = "xs.string"
	schemaManager.xsdPrimitiveTypesMap["uint"] = "xs.unsignedInt"
	schemaManager.xsdPrimitiveTypesMap["bool"] = "xs.boolean"
	schemaManager.xsdPrimitiveTypesMap["ID"] = "xs.ID"

	return schemaManager
}

////////////////////////////////////////////////////////////////////////////////
// Service functions
////////////////////////////////////////////////////////////////////////////////

func (this *SchemaManager) initialize() {

	log.Println("--> Initialize SchemaManager")
	GetServer().GetConfigurationManager().setServiceConfiguration(this.getId())

	// Here I will initialyse the schema found in the schema directory.
	// Those schema must be xml schema...
	schemasDir, _ := ioutil.ReadDir(GetServer().GetConfigurationManager().GetSchemasPath())
	for _, f := range schemasDir {
		if strings.HasSuffix(strings.ToUpper(f.Name()), ".XSD") == true {
			schemasXsdPath := GetServer().GetConfigurationManager().GetSchemasPath() + "/" + f.Name()
			var schema *XML_Schemas.XSD_Schema
			schema = new(XML_Schemas.XSD_Schema)
			file, err := os.Open(schemasXsdPath)
			defer file.Close()
			if err = xml.NewDecoder(file).Decode(schema); err != nil {
				log.Println("==----> ", err)
			} else {
				// if the schema does not exist I will get the existing schema...
				if len(schema.TargetNamespace) != 0 {
					schema.Id = schema.TargetNamespace[strings.LastIndex(schema.TargetNamespace, "/")+1:]
					//log.Println("--> Load file: ", schemasXsdPath)
					this.schemas[schema.Id] = schema

					// Read and populate maps of element, complex type etc...
					this.parseSchema(schema)
				}
			}
		}
	}

	// Force the xsd schemas first...
	this.genereatePrototype(this.schemas["xs"])

	for _, schema := range this.schemas {
		// use information from previous step to generate entity prototypes.

		/* TODO manage the schemas version...
		if len(schema.Version) > 0 {
			version := strings.Replace(schema.Version, ".", "_", -1)
			storeId += "_" + version
		}*/

		// Set the schema id with the store id values...

		store := GetServer().GetDataManager().getDataStore(schema.Id)
		if store == nil {
			// I will create the new store here...
			var errObj *CargoEntities.Error
			store, errObj = GetServer().GetDataManager().createDataStore(schema.Id, Config.DataStoreType_KEY_VALUE_STORE, Config.DataStoreVendor_MYCELIUS)
			if errObj != nil {
				log.Println(errObj.GetBody())
			}
		}

		if schema.TargetNamespace != "http://www.w3.org/2001/XMLSchema/xs" {
			this.genereatePrototype(schema)
		}
	}

	// Before i create the prototy I will expand supertypes,
	// that mean I will copy fields of super type into the
	// prototype itself...
	for _, prototype := range this.prototypes {
		// Set the super type fields...
		this.setSuperTypeField(prototype)
	}

	// Save entity prototypes...
	for _, prototype := range this.prototypes {
		// Set the field order.
		for i := 0; i < len(prototype.Fields); i++ {
			prototype.FieldsOrder = append(prototype.FieldsOrder, i)
		}
		// Here i will save the prototype...
		prototype.Create("")

		// Print the list of prototypes...
		//prototype.Print()
	}

	// Now I will import the xml file from the schema directory...
	for _, f := range schemasDir {
		if strings.HasSuffix(strings.ToUpper(f.Name()), ".XSD") == false {
			log.Println("import file: ", GetServer().GetConfigurationManager().GetSchemasPath()+"/"+f.Name())
			this.importXmlFile(GetServer().GetConfigurationManager().GetSchemasPath() + "/" + f.Name())
		}
	}
}

func (this *SchemaManager) getId() string {
	return "SchemaManager"
}

func (this *SchemaManager) start() {
	log.Println("--> Start SchemaManager")
}

func (this *SchemaManager) stop() {
	log.Println("--> Stop SchemaManager")
}

func removeNs(id string) string {
	if strings.Index(id, ":") > 0 {
		return id[strings.Index(id, ":")+1:]
	}
	return id
}

/**
 * Import a new schema whit a given file path.
 */
func (this *SchemaManager) importSchema(schemasXsdPath string) *CargoEntities.Error {

	var schema *XML_Schemas.XSD_Schema
	schema = new(XML_Schemas.XSD_Schema)
	file, err := os.Open(schemasXsdPath)
	defer file.Close()
	if err = xml.NewDecoder(file).Decode(schema); err != nil {
		cargoError := NewError(Utility.FileLine(), FILE_OPEN_ERROR, SERVER_ERROR_CODE, err)
		return cargoError
	} else {
		// if the schema does not exist I will get the existing schema...
		if len(schema.TargetNamespace) != 0 {
			schema.Id = schema.TargetNamespace[strings.LastIndex(schema.TargetNamespace, "/")+1:]
			//log.Println("--> Load file: ", schemasXsdPath)
			this.schemas[schema.Id] = schema

			// Read and populate maps of element, complex type etc...
			this.parseSchema(schema)
		}
	}

	// Set the schema id with the store id values...
	store := GetServer().GetDataManager().getDataStore(schema.Id)
	if store == nil {
		// I will create the new store here...
		var errObj *CargoEntities.Error
		store, errObj = GetServer().GetDataManager().createDataStore(schema.Id, Config.DataStoreType_KEY_VALUE_STORE, Config.DataStoreVendor_MYCELIUS)
		if errObj != nil {
			return errObj
		}
	}

	if schema.TargetNamespace != "http://www.w3.org/2001/XMLSchema/xs" {
		this.genereatePrototype(schema)
	}

	// Before i create the prototy I will expand supertypes,
	// that mean I will copy fields of super type into the
	// prototype itself...
	for _, prototype := range this.prototypes {
		// Set the super type fields...
		this.setSuperTypeField(prototype)
	}

	// Save entity prototypes...
	for _, prototype := range this.prototypes {
		// Set the field order.
		for i := 0; i < len(prototype.Fields); i++ {
			prototype.FieldsOrder = append(prototype.FieldsOrder, i)
		}
		// Here i will save the prototype...
		prototype.Create("")

		// Print the list of prototypes...
		//prototype.Print()
	}

	return nil
}

/**
 * Return the list of fields including parent fields
 */
func (this *SchemaManager) GetFieldsFieldsType(prototype *EntityPrototype, path *[]string) ([]string, []string) {

	var fields []string
	var fieldsType []string

	if Utility.Contains(*path, prototype.TypeName) {
		return fields, fieldsType
	}

	*path = append(*path, prototype.TypeName)

	for i := 0; i < len(prototype.Fields); i++ {
		if prototype.Fields[i] != "UUID" && prototype.Fields[i] != "ParentUuid" && prototype.Fields[i] != "childsUuid" && prototype.Fields[i] != "referenced" {
			fields = append(fields, prototype.Fields[i])
			fieldsType = append(fieldsType, prototype.FieldsType[i])
		}
	}

	for i := 0; i < len(prototype.SuperTypeNames); i++ {
		p := this.prototypes[prototype.SuperTypeNames[i]]

		fields_, fieldsType_ := this.GetFieldsFieldsType(p, path)
		for j := 0; j < len(fields_); j++ {
			if Utility.Contains(fields, fields_[j]) == false {
				fields = append(fields, fields_[j])
				fieldsType = append(fieldsType, fieldsType_[j])
			}
		}
	}

	// Return the list of fields...
	return fields, fieldsType
}

//Append the super type fields...
func (this *SchemaManager) setSuperTypeField(prototype *EntityPrototype) {
	path := make([]string, 0)
	fields, fieldsType := this.GetFieldsFieldsType(prototype, &path)
	for i := 0; i < len(fields); i++ {
		if Utility.Contains(prototype.Fields, fields[i]) == false {
			prototype.Fields = append(prototype.Fields, fields[i])
			prototype.FieldsType = append(prototype.FieldsType, fieldsType[i])
			prototype.FieldsVisibility = append(prototype.FieldsVisibility, true)
		}
	}

}

/**
 * Generate an entity prototype from given schema...
 */
func (this *SchemaManager) parseSchema(schema *XML_Schemas.XSD_Schema) {

	// The annotation...
	for i := 0; i < len(schema.Annotations); i++ {
		this.parseAnnotation(schema.Annotations[i])
	}

	// The include...

	// Now the includes files...
	for i := 0; i < len(schema.Includes); i++ {
		schemasXsdPath := GetServer().GetConfigurationManager().GetSchemasPath() + "/" + schema.Includes[i].SchemaLocation
		schema_ := new(XML_Schemas.XSD_Schema)
		file, err := os.Open(schemasXsdPath)
		defer file.Close()
		if err = xml.NewDecoder(file).Decode(schema_); err != nil {
			log.Println("==----> ", err)
		} else {
			// Set the same schema id...
			schema_.Id = schema.Id
			// Read and populate maps of element, complex type etc...
			this.parseSchema(schema_)
		}
	}

	// The import ...
	for i := 0; i < len(schema.Imports); i++ {
	}

	// The Redefines ...
	for i := 0; i < len(schema.Redefines); i++ {
	}

	// The Attributes ...
	for i := 0; i < len(schema.Attributes); i++ {
		this.parseAttribute(&schema.Attributes[i], schema, true)
	}

	// The Attributes groups...
	for i := 0; i < len(schema.AttributeGroups); i++ {
		this.parseAttributeGroups(&schema.AttributeGroups[i], schema, true)
	}

	// The Any Attributes...
	for i := 0; i < len(schema.AnyAttributes); i++ {
		this.parseAnyAttribute(&schema.AnyAttributes[i], schema, true)
	}

	// The elements...
	for i := 0; i < len(schema.Elements); i++ {
		this.parseElement(&schema.Elements[i], schema, true)
	}

	// The complex type
	for i := 0; i < len(schema.ComplexTypes); i++ {
		this.parseComplexType(&schema.ComplexTypes[i], schema, true)
	}

	// The simple type
	for i := 0; i < len(schema.SimpleTypes); i++ {
		this.parseSimpleType(&schema.SimpleTypes[i], schema, true)
	}

	// The group
	for i := 0; i < len(schema.Groups); i++ {
		this.parseGroup(&schema.Groups[i], schema, true)
	}

	// TODO Notation...

}

////////////////////////////////////////////////////////////////////////////////
// Attributes
////////////////////////////////////////////////////////////////////////////////
/**
 * Parse Attribute Group
 */
func (this *SchemaManager) parseAttributeGroups(group *XML_Schemas.XSD_AttributeGroup, schema *XML_Schemas.XSD_Schema, isGlobal bool) {
	id := schema.Id + "." + removeNs(group.Name)
	if isGlobal {
		this.globalAttributeGroup[id] = group
	}
}

/**
 * Parse Any Attribute
 */
func (this *SchemaManager) parseAnyAttribute(any *XML_Schemas.XSD_AnyAttribute, schema *XML_Schemas.XSD_Schema, isGlobal bool) {
	id := schema.Id + "." + removeNs(any.Name)
	if isGlobal {
		this.globalAnyAttributes[id] = any
	}
}

/**
 * Parse Attribute
 */
func (this *SchemaManager) parseAttribute(attribute *XML_Schemas.XSD_Attribute, schema *XML_Schemas.XSD_Schema, isGlobal bool) {
	id := schema.Id + "." + removeNs(attribute.Name)
	if isGlobal {
		this.globalAttributes[id] = attribute
	}
}

////////////////////////////////////////////////////////////////////////////////
// Compositor elements...
// Those element generate prototypes...
////////////////////////////////////////////////////////////////////////////////
func (this *SchemaManager) parseAnnotation(annotation XML_Schemas.XSD_Annotation) string {
	annotationStr := ""
	return annotationStr
}

/**
 * Parse Element
 */
func (this *SchemaManager) parseElement(element *XML_Schemas.XSD_Element, schema *XML_Schemas.XSD_Schema, isGlobal bool) {
	id := schema.Id + "." + removeNs(element.Name)
	if isGlobal {
		this.globalElements[id] = element
	}

	// Keep the schema information here...
	element.SchemaId = schema.Id
	this.elementSchema[element.Name] = schema.Id
}

/**
 * Parse SimpleType
 */
func (this *SchemaManager) parseSimpleType(simpleType *XML_Schemas.XSD_SimpleType, schema *XML_Schemas.XSD_Schema, isGlobal bool) {
	id := schema.Id + "." + removeNs(simpleType.Name)
	if isGlobal {
		this.globalSimpleTypes[id] = simpleType
	}
}

/**
 * Parse ComplexType
 */
func (this *SchemaManager) parseComplexType(complexType *XML_Schemas.XSD_ComplexType, schema *XML_Schemas.XSD_Schema, isGlobal bool) {

	id := schema.Id + "." + removeNs(complexType.Name)
	if isGlobal {
		// A globel complex type must have it one prototype
		this.globalComplexTypes[id] = complexType
	}
}

/**
 * Parse Group
 */
func (this *SchemaManager) parseGroup(group *XML_Schemas.XSD_Group, schema *XML_Schemas.XSD_Schema, isGlobal bool) {
	var groupId string

	if len(group.Id) > 0 {
		groupId = group.Id
	} else {
		groupId = group.Name
	}

	id := schema.Id + "." + removeNs(groupId)

	if isGlobal {
		this.globalGroups[id] = group
	}

}

/**
 * Now the prototypes generation
 */
func (this *SchemaManager) genereatePrototype(schema *XML_Schemas.XSD_Schema) {
	//Global simple type...
	for k, s := range this.globalSimpleTypes {
		if strings.HasPrefix(k, schema.Id) {
			prototype := this.createPrototypeSimpleType(schema, s)
			if strings.HasPrefix(prototype.TypeName, "xs.") {
				if Utility.Contains(prototype.Fields, "M_valueOf") == false {
					prototype.Fields = append(prototype.Fields, "M_valueOf")
					prototype.FieldsType = append(prototype.FieldsType, prototype.TypeName)
					prototype.FieldsOrder = append(prototype.FieldsOrder, len(prototype.FieldsOrder))
					prototype.FieldsVisibility = append(prototype.FieldsVisibility, true)
				}
			}
		}
	}

	//Global Complex type...
	for k, c := range this.globalComplexTypes {
		if strings.HasPrefix(k, schema.Id) {
			this.createPrototypeComplexType(schema, c)
		}
	}

	// Global element...
	for k, e := range this.globalElements {
		if strings.HasPrefix(k, schema.Id) {
			this.createPrototypeElement(schema, e)
		}
	}
}

func (this *SchemaManager) createPrototypeElement(schema *XML_Schemas.XSD_Schema, element *XML_Schemas.XSD_Element) *EntityPrototype {

	var simpleType *XML_Schemas.XSD_SimpleType
	var complexType *XML_Schemas.XSD_ComplexType

	if len(element.Type) > 0 {
		// Here the content is define outside of the element...
		if val, ok := this.globalComplexTypes[schema.Id+"."+removeNs(element.Type)]; ok {
			complexType = val
		} else if val, ok := this.globalSimpleTypes[schema.Id+"."+removeNs(element.Type)]; ok {
			simpleType = val
		}
	} else {
		if element.ComplexType != nil {
			complexType = element.ComplexType
			complexType.Name = element.Name
		} else if element.SimpleType != nil {
			simpleType = element.SimpleType
			simpleType.Name = element.Name
		}
	}

	var prototype *EntityPrototype

	// So here the element content can be a complex or simple type...
	if complexType != nil {
		prototype = this.createPrototypeComplexType(schema, complexType)
	} else if simpleType != nil {
		prototype = this.createPrototypeSimpleType(schema, simpleType)
	} else {
		// In that case the class has no simple or complex type...
		prototype = NewEntityPrototype()
		prototype.TypeName = schema.Id + "." + element.Name

		// Append the prototype to the map...
		this.prototypes[prototype.TypeName] = prototype
	}

	// Set the prototype as abstract...
	if element.IsAbstract == "true" {
		prototype.IsAbstract = true
	}
	if len(element.SubstitutionGroup) > 0 {
		this.appendPrototypeSuperBaseType(schema.Id+"."+element.SubstitutionGroup, prototype)
		var parent *EntityPrototype
		if p, ok := this.prototypes[schema.Id+"."+element.SubstitutionGroup]; ok {
			parent = p
		} else {
			if pe, ok := this.globalElements[schema.Id+"."+removeNs(element.SubstitutionGroup)]; ok {
				parent = this.createPrototypeElement(schema, pe)
			}
		}

		// Append the type name of this prototype to it parent prototype...
		parent.SubstitutionGroup = append(parent.SubstitutionGroup, schema.Id+"."+element.Name)
		this.prototypes[schema.Id+"."+element.SubstitutionGroup] = parent
	}
	return prototype
}

////////////////////////////////////////////////////////////////////////////////
// Prototype can be created from, simple type, complex type element and group
////////////////////////////////////////////////////////////////////////////////

/**
 * Create simple type prototype...
 */
func (this *SchemaManager) createPrototypeSimpleType(schema *XML_Schemas.XSD_Schema, simpleType *XML_Schemas.XSD_SimpleType) *EntityPrototype {
	// The new prototype to append to parent...
	prototype := NewEntityPrototype()

	if len(simpleType.Name) > 0 {
		// Here if the prototype exist on the list of prototype I will not recreate it
		if val, ok := this.prototypes[schema.Id+"."+simpleType.Name]; ok {
			return val
		}
		// Append the prototype to the map...
		prototype.TypeName = schema.Id + "." + simpleType.Name
		// Append the prototype to the map...
		this.prototypes[prototype.TypeName] = prototype
	}

	if val, ok := this.prototypes[schema.Id+"."+simpleType.Name]; ok {
		prototype = val
	} else if len(simpleType.Name) > 0 {
		// Take it from the xsd simple type...
		if val, ok := this.prototypes["xs."+simpleType.Name]; ok {
			prototype = val
		} else {
			// Append the simple type object...
			this.createPrototypeSimpleType(schema, simpleType)
		}
	}

	// Derivation of a simple datatype by restriction.
	if simpleType.Restriction != nil {
		this.appendPrototypeRestriction(schema, prototype, simpleType.Restriction, simpleType, nil, nil)
	} else if simpleType.List != nil {
		this.appendPrototypeList(schema, prototype, simpleType.List)
	} else if simpleType.Union != nil {
		this.appendPrototypeUnion(schema, prototype, simpleType.Union)
	}

	return prototype
}

/**
 * Create complex type prototype...
 */
func (this *SchemaManager) createPrototypeComplexType(schema *XML_Schemas.XSD_Schema, complexType *XML_Schemas.XSD_ComplexType) *EntityPrototype {
	prototype := NewEntityPrototype()

	if complexType.IsAbstract {
		prototype.IsAbstract = true
	}

	if len(complexType.Name) > 0 {
		// Here if the prototype exist on the list of prototype I will not recreate it
		if val, ok := this.prototypes[schema.Id+"."+complexType.Name]; ok {
			return val
		}

		prototype.TypeName = schema.Id + "." + complexType.Name
		// Append the prototype to the map...
		this.prototypes[prototype.TypeName] = prototype
	}

	if val, ok := this.prototypes[schema.Id+"."+complexType.Name]; ok {
		prototype = val
	}

	// The attributes...
	for i := 0; i < len(complexType.Attributes); i++ {
		this.appendPrototypeAttribute(schema, prototype, &complexType.Attributes[i])
	}

	// In local group only...
	for i := 0; i < len(complexType.AttributeGroups); i++ {
		this.appendPrototypeAttributeGroup(schema, prototype, &complexType.AttributeGroups[i])
	}

	for i := 0; i < len(complexType.AnyAttributes); i++ {
		this.appendPrototypeAnyAttribute(schema, prototype, &complexType.AnyAttributes[i])
	}

	// The content of the complex type can be
	if complexType.ComplexContent != nil {
		if complexType.ComplexContent.Extension != nil {
			this.appendPrototypeExtention(schema, prototype, complexType.ComplexContent.Extension, complexType.ComplexContent, nil)
		} else if complexType.ComplexContent.Restriction != nil {
			this.appendPrototypeRestriction(schema, prototype, complexType.ComplexContent.Restriction, nil, complexType.ComplexContent, nil)
		}

	} else if complexType.SimpleContent != nil {
		if complexType.SimpleContent.Extension != nil {
			this.appendPrototypeExtention(schema, prototype, complexType.SimpleContent.Extension, nil, complexType.SimpleContent)
		} else if complexType.SimpleContent.Restriction != nil {
			this.appendPrototypeRestriction(schema, prototype, complexType.SimpleContent.Restriction, nil, nil, complexType.SimpleContent)
		}
	} else {
		if complexType.Group != nil {
			this.appendPrototypeGroup(schema, prototype, complexType.Group, complexType.Group.MinOccurs, complexType.Group.MaxOccurs)
		} else if complexType.All != nil {
			this.appendPrototypeAll(schema, prototype, complexType.All, complexType.All.MinOccurs, complexType.All.MaxOccurs)
		} else if complexType.Choice != nil {
			this.appendPrototypeChoice(schema, prototype, complexType.Choice, complexType.Choice.MinOccurs, complexType.Choice.MaxOccurs)
		} else if complexType.Sequence != nil {
			this.appendPrototypeSequence(schema, prototype, complexType.Sequence, complexType.Sequence.MinOccurs, complexType.Sequence.MaxOccurs)
		}
	}

	return prototype
}

////////////////////////////////////////////////////////////////////////////////
// Append field to prototypes...
////////////////////////////////////////////////////////////////////////////////

/**
 * Append attribute to the prototype...
 * TODO connect the default value when fixed is true...
 */
func (this *SchemaManager) appendPrototypeAttribute(schema *XML_Schemas.XSD_Schema, prototype *EntityPrototype, attribute *XML_Schemas.XSD_Attribute) {

	// I will get the attribute from it reference...
	if len(attribute.Ref) > 0 {
		attribute = this.globalAttributes[schema.Id+"."+removeNs(attribute.Ref)]
	}

	if attribute == nil {
		log.Println("Attribut not found!!!")
		return
	}

	var simpleType *XML_Schemas.XSD_SimpleType
	if len(attribute.Type) > 0 {
		simpleType = this.globalSimpleTypes[schema.Id+"."+removeNs(attribute.Type)]
		if simpleType == nil {
			// if the name is nil, I will try to get it from the xs namespace...
			primitiveName := attribute.Type[strings.Index(attribute.Type, ":")+1:]
			// set xs as package name...
			primitiveType := "xs." + primitiveName
			if val, ok := this.globalSimpleTypes[primitiveType]; ok {
				simpleType = val
				schema = this.schemas["xs"]
			}
		}
	} else {
		simpleType = attribute.SimpleType
	}

	// Try to get it from xsd base type.
	if simpleType == nil {
		// Here I will try with a XSD primitive...
		if len(attribute.Type) > 0 {
			// Keep the name part only
			primitiveName := attribute.Type[strings.Index(attribute.Type, ":")+1:]
			if val, ok := this.xsdPrimitiveTypesMap[primitiveName]; ok {
				if val, ok := this.globalSimpleTypes[val]; ok {
					simpleType = val
					schema = this.schemas["xs"]
				}
			}
		}
	}

	if simpleType != nil {

		if len(simpleType.Name) == 0 {
			simpleType.Name = attribute.Name + "_Type"
		}
		p := this.createPrototypeSimpleType(schema, simpleType)
		if !Utility.Contains(prototype.Fields, "M_"+attribute.Name) {
			prototype.Fields = append(prototype.Fields, "M_"+attribute.Name)

			// TODO append other reference type as needed in id's or indexs...
			// Set the id's
			isId := strings.HasSuffix(strings.ToUpper(attribute.Name), "-ID") || strings.HasSuffix(strings.ToUpper(attribute.Name), "_ID") || strings.ToUpper(attribute.Name) == "ID" || strings.HasSuffix(p.TypeName, ".ID")
			if isId {
				prototype.Ids = append(prototype.Ids, "M_"+attribute.Name)
			}

			// Set indexation key...
			isIndex := strings.ToUpper(attribute.Name) == "NAME" || strings.HasSuffix(p.TypeName, ".anyURI") || strings.HasSuffix(p.TypeName, ".NCName") || strings.HasSuffix(p.TypeName, ".IDREF") || strings.HasSuffix(p.TypeName, ".IDREFS")
			if isIndex {
				prototype.Indexs = append(prototype.Indexs, "M_"+attribute.Name)
			}

			prototype.FieldsType = append(prototype.FieldsType, p.TypeName)
			prototype.FieldsOrder = append(prototype.FieldsOrder, len(prototype.FieldsOrder))
			prototype.FieldsVisibility = append(prototype.FieldsVisibility, true)
		}
	}
}

/**
 * Append wildcard to replace any element.
 */
func (this *SchemaManager) appendPrototypeAny(schema *XML_Schemas.XSD_Schema, parent *EntityPrototype, any *XML_Schemas.XSD_Any) {
	// TODO read anything here...
}

/**
 * Append any attribute to the prototype...
 */
func (this *SchemaManager) appendPrototypeAnyAttribute(schema *XML_Schemas.XSD_Schema, parent *EntityPrototype, anyAttribute *XML_Schemas.XSD_AnyAttribute) {
	// TODO read any attributes...
}

/**
 * Append attribute group to the prototype...
 */
func (this *SchemaManager) appendPrototypeAttributeGroup(schema *XML_Schemas.XSD_Schema, parent *EntityPrototype, groupAttribute *XML_Schemas.XSD_AttributeGroup) {
	// If is global...
	if _, ok := this.globalAttributeGroup[schema.Id+"."+removeNs(groupAttribute.Name)]; ok {
		for i := 0; i < len(groupAttribute.AnyAttributes); i++ {
			this.appendPrototypeAnyAttribute(schema, parent, groupAttribute.AnyAttributes[i])
		}
	}
}

/**
 * set the last field cardinality
 */
func (this *SchemaManager) setLastPrototypeFieldCardinality(prototype *EntityPrototype, minOccurs string, maxOccurs string) {
	cardinality := ""

	if maxOccurs == "unbounded" {
		cardinality = "[]"
		prototype.FieldsNillable = append(prototype.FieldsNillable, false)
	} else if maxOccurs == "1" && minOccurs == "1" {
		prototype.FieldsNillable = append(prototype.FieldsNillable, false)
	} else if maxOccurs == "1" && minOccurs == "0" {
		prototype.FieldsNillable = append(prototype.FieldsNillable, true)
	}

	// Set brakets to the fieldType...
	if strings.HasPrefix(prototype.FieldsType[len(prototype.FieldsType)-1], "[]") == false {
		prototype.FieldsType[len(prototype.FieldsType)-1] = cardinality + prototype.FieldsType[len(prototype.FieldsType)-1]
	}
}

// The rest of stuff...
/**
 *  append element to prototype...
 */
func (this *SchemaManager) appendPrototypeElement(schema *XML_Schemas.XSD_Schema, prototype *EntityPrototype, element *XML_Schemas.XSD_Element, minOccurs string, maxOccurs string) {

	// Here the element to append is not by reference, so
	// it can be a subtype...
	// The super class of an element...
	if len(element.SubstitutionGroup) > 0 {
		this.appendPrototypeSuperBaseType(schema.Id+"."+element.SubstitutionGroup, prototype)
		var parent *EntityPrototype
		if p, ok := this.prototypes[schema.Id+"."+element.SubstitutionGroup]; ok {
			parent = p
		} else {
			if pe, ok := this.globalElements[schema.Id+"."+removeNs(element.SubstitutionGroup)]; ok {
				parent = this.createPrototypeElement(schema, pe)
			}
		}

		// Append the type name of this prototype to it parent prototype...
		parent.SubstitutionGroup = append(parent.SubstitutionGroup, schema.Id+"."+element.Name)
		this.prototypes[schema.Id+"."+element.SubstitutionGroup] = parent
	}

	// Resolve the reference, and append an element of this existing
	// type.
	if len(element.Ref) > 0 {
		isRef := strings.HasSuffix(element.Type, "anyURI")
		if val, ok := this.globalElements[schema.Id+"."+removeNs(element.Ref)]; ok {
			element = val

			if p, ok := this.prototypes[schema.Id+"."+element.Name]; ok {
				// Here the prototype is found...
				if !Utility.Contains(prototype.Fields, "M_"+element.Name) {
					prototype.Fields = append(prototype.Fields, "M_"+element.Name)

					if p.TypeName == "xs.anyURI" || getRootTypeName(p.TypeName) == "xs.anyURI" || isRef {
						prototype.FieldsType = append(prototype.FieldsType, p.TypeName+":Ref")
					} else {
						prototype.FieldsType = append(prototype.FieldsType, p.TypeName)
					}
					prototype.FieldsOrder = append(prototype.FieldsOrder, len(prototype.FieldsOrder))
					prototype.FieldsVisibility = append(prototype.FieldsVisibility, true)
					this.setLastPrototypeFieldCardinality(prototype, minOccurs, maxOccurs)
				}
				return
			} else {
				p := this.createPrototypeElement(schema, element)
				if p != nil {
					if !Utility.Contains(prototype.Fields, "M_"+element.Name) {
						prototype.Fields = append(prototype.Fields, "M_"+element.Name)

						if p.TypeName == "xs.anyURI" || getRootTypeName(p.TypeName) == "xs.anyURI" || isRef {
							prototype.FieldsType = append(prototype.FieldsType, p.TypeName+":Ref")
						} else {
							prototype.FieldsType = append(prototype.FieldsType, p.TypeName)
						}

						prototype.FieldsOrder = append(prototype.FieldsOrder, len(prototype.FieldsOrder))
						prototype.FieldsVisibility = append(prototype.FieldsVisibility, true)
						this.setLastPrototypeFieldCardinality(prototype, minOccurs, maxOccurs)
					}
				} else {
					log.Println("-------> prototype: ", element.Name, " Not found!!!")
				}
				return
			}
		} else {
			log.Println("-------> element: ", element.Ref, " Not found!!!")
			return
		}
	}

	// Here it's an element with local scope...
	var simpleType *XML_Schemas.XSD_SimpleType
	var complexType *XML_Schemas.XSD_ComplexType
	if len(element.Type) > 0 {
		// Here the content is define outside of the element...
		if val, ok := this.globalComplexTypes[schema.Id+"."+removeNs(element.Type)]; ok {
			complexType = val
		} else if val, ok := this.globalSimpleTypes[schema.Id+"."+removeNs(element.Type)]; ok {
			simpleType = val
		} else {
			elementType := element.Type[strings.Index(element.Type, ":")+1:]
			if val, ok := this.globalSimpleTypes["xs."+elementType]; ok {
				simpleType = val
				schema = this.schemas["xs"]
			}
		}
	} else {
		if element.ComplexType != nil {
			complexType = element.ComplexType
		} else if element.SimpleType == nil {
			simpleType = element.SimpleType
		}
	}

	// So here the element content can be a complex or simple type...
	if complexType != nil {
		if len(element.Type) > 0 {
			if !Utility.Contains(prototype.Fields, "M_"+element.Name) {
				prototype.Fields = append(prototype.Fields, "M_"+element.Name)
				prototype.FieldsType = append(prototype.FieldsType, schema.Id+"."+element.Type)
				prototype.FieldsOrder = append(prototype.FieldsOrder, len(prototype.FieldsOrder))
				prototype.FieldsVisibility = append(prototype.FieldsVisibility, true)
				this.setLastPrototypeFieldCardinality(prototype, minOccurs, maxOccurs)
			}
		} else {
			// For complex type local to an element I will generate complex type
			// with name separated by "." so the pattern is
			// pacakegeName.Type.InnerType.InnerType.InnerType... etc...
			if len(complexType.Name) == 0 {
				values := strings.Split(prototype.TypeName, ".")
				prototypeTypeName := ""
				for i := 1; i < len(values); i++ {
					if i > 1 {
						prototypeTypeName += "."
					}
					prototypeTypeName += values[i]
				}

				if len(prototypeTypeName) == 0 {
					complexType.Name = element.Name + "_Type"
				} else {
					complexType.Name = prototypeTypeName + "." + element.Name + "_Type"
				}

			}
			p := this.createPrototypeComplexType(schema, complexType)
			if !Utility.Contains(prototype.Fields, "M_"+element.Name) {
				prototype.Fields = append(prototype.Fields, "M_"+element.Name)
				prototype.FieldsType = append(prototype.FieldsType, p.TypeName)
				prototype.FieldsOrder = append(prototype.FieldsOrder, len(prototype.FieldsOrder))
				prototype.FieldsVisibility = append(prototype.FieldsVisibility, true)
				this.setLastPrototypeFieldCardinality(prototype, minOccurs, maxOccurs)
			}
		}
	} else if simpleType != nil {
		p := this.createPrototypeSimpleType(schema, simpleType)
		if !Utility.Contains(prototype.Fields, "M_"+element.Name) {
			prototype.Fields = append(prototype.Fields, "M_"+element.Name)
			prototype.FieldsType = append(prototype.FieldsType, p.TypeName)
			prototype.FieldsOrder = append(prototype.FieldsOrder, len(prototype.FieldsOrder))
			prototype.FieldsVisibility = append(prototype.FieldsVisibility, true)
			this.setLastPrototypeFieldCardinality(prototype, minOccurs, maxOccurs)
		}
	}

	for i := 0; i < len(element.Key); i++ {
		this.appendPrototypeKey(schema, prototype, &element.Key[i])
	}

	for i := 0; i < len(element.KeyRef); i++ {
		this.appendPrototypeKeyRef(schema, prototype, &element.KeyRef[i])
	}

	for i := 0; i < len(element.Unique); i++ {
		this.appendPrototypeUnique(schema, prototype, &element.Unique[i])
	}
}

/**
 *  Append group...
 */
func (this *SchemaManager) appendPrototypeGroup(schema *XML_Schemas.XSD_Schema, prototype *EntityPrototype, group *XML_Schemas.XSD_Group, minOccurs string, maxOccurs string) *EntityPrototype {

	if len(group.Ref) > 0 {
		minOccurs := group.MinOccurs
		maxOccurs := group.MaxOccurs
		group = this.globalGroups[schema.Id+"."+removeNs(group.Ref)]

		group.MaxOccurs = maxOccurs
		group.MinOccurs = minOccurs
	}

	if group.All != nil {
		this.appendPrototypeAll(schema, prototype, group.All, group.MinOccurs, group.MaxOccurs)
	}

	if group.Any != nil {
		this.appendPrototypeAny(schema, prototype, group.Any)
	}

	if group.Choice != nil {
		this.appendPrototypeChoice(schema, prototype, group.Choice, group.Choice.MinOccurs, group.Choice.MaxOccurs)
	}

	if group.Sequence != nil {
		this.appendPrototypeSequence(schema, prototype, group.Sequence, group.Sequence.MinOccurs, group.Sequence.MaxOccurs)
	}

	this.setLastPrototypeFieldCardinality(prototype, minOccurs, maxOccurs)

	return prototype
}

/**
 *  Append the content of a sequence to a prototype...
 */
func (this *SchemaManager) appendPrototypeSequence(schema *XML_Schemas.XSD_Schema, prototype *EntityPrototype, sequence *XML_Schemas.XSD_Sequence, minOccurs string, maxOccurs string) {

	// The sequence inner element
	for i := 0; i < len(sequence.Elements); i++ {
		element := sequence.Elements[i]
		this.appendPrototypeElement(schema, prototype, &element, element.MinOccurs, element.MaxOccurs)
		this.setLastPrototypeFieldCardinality(prototype, minOccurs, maxOccurs)
	}

	// The sequence inner items...

	// The groups
	for i := 0; i < len(sequence.Groups); i++ {
		group := sequence.Groups[i]
		this.appendPrototypeGroup(schema, prototype, &group, group.MinOccurs, group.MaxOccurs)
		this.setLastPrototypeFieldCardinality(prototype, minOccurs, maxOccurs)
	}

	// The Choices
	for i := 0; i < len(sequence.Choices); i++ {
		choice := sequence.Choices[i]
		this.appendPrototypeChoice(schema, prototype, &choice, choice.MinOccurs, choice.MaxOccurs)
		this.setLastPrototypeFieldCardinality(prototype, minOccurs, maxOccurs)
	}

	// Recursive sequence.
	for i := 0; i < len(sequence.Sequences); i++ {
		sequence_ := sequence.Sequences[i]
		this.appendPrototypeSequence(schema, prototype, &sequence_, sequence_.MinOccurs, sequence_.MaxOccurs)
		this.setLastPrototypeFieldCardinality(prototype, minOccurs, maxOccurs)
	}

	// Any
	for i := 0; i < len(sequence.Any); i++ {
		any := sequence.Any[i]
		this.appendPrototypeAny(schema, prototype, &any)
	}
}

/**
 * Append Field.
 */
func (this *SchemaManager) appendPrototypeField(schema *XML_Schemas.XSD_Schema, prototype *EntityPrototype, field *XML_Schemas.XSD_Field) {

	// This is kind of attribute that contain a xpath string...
	if !Utility.Contains(prototype.Fields, "M_"+field.Id) {
		prototype.Fields = append(prototype.Fields, "M_"+field.Id)
		prototype.FieldsType = append(prototype.FieldsType, "xs.ID")
		prototype.FieldsOrder = append(prototype.FieldsOrder, len(prototype.FieldsOrder))
		prototype.FieldsVisibility = append(prototype.FieldsVisibility, true)
		prototype.Indexs = append(prototype.Indexs, "M_"+field.Id)
	}

}

/**
 * Append selector.
 */
func (this *SchemaManager) appendPrototypeSelector(schema *XML_Schemas.XSD_Schema, parent *EntityPrototype, selector *XML_Schemas.XSD_Selector) {

}

/**
 * Append unique element.
 */
func (this *SchemaManager) appendPrototypeUnique(schema *XML_Schemas.XSD_Schema, prototype *EntityPrototype, unique *XML_Schemas.XSD_Unique) {
	this.appendPrototypeSelector(schema, prototype, &unique.Selector)

	for i := 0; i < len(unique.Fields); i++ {
		this.appendPrototypeField(schema, prototype, &unique.Fields[i])
	}
}

/**
 * Append unique element.
 */
func (this *SchemaManager) appendPrototypeList(schema *XML_Schemas.XSD_Schema, prototype *EntityPrototype, lst *XML_Schemas.XSD_List) {
	if lst.SimpleType != nil {
		// Create a list of simple type element.
		p := this.createPrototypeSimpleType(schema, lst.SimpleType)
		if !Utility.Contains(prototype.Fields, "M_"+lst.SimpleType.Name) {
			prototype.Fields = append(prototype.Fields, "M_"+lst.SimpleType.Name)
			prototype.FieldsType = append(prototype.FieldsType, "[]"+p.TypeName)
			prototype.FieldsOrder = append(prototype.FieldsOrder, len(prototype.FieldsOrder))
			prototype.FieldsVisibility = append(prototype.FieldsVisibility, true)
		}
	} else if len(lst.ItemType) > 0 {
		var itemType string
		if strings.HasPrefix(lst.ItemType, "xs:") || strings.HasPrefix(lst.ItemType, "xsd:") {
			itemType = "xs." + lst.ItemType[strings.Index(lst.ItemType, ":")+1:]
		} else {
			itemType = schema.Id + "." + lst.ItemType
		}

		// Append a field to hold the values...
		prototype.ListOf = itemType
		if !Utility.Contains(prototype.Fields, "M_listOf") {
			prototype.Fields = append(prototype.Fields, "M_listOf")
			prototype.FieldsType = append(prototype.FieldsType, "[]"+itemType)
			prototype.FieldsOrder = append(prototype.FieldsOrder, len(prototype.FieldsOrder))
			prototype.FieldsVisibility = append(prototype.FieldsVisibility, true)
		}

	}
}

/**
 * Append key ref element
 */
func (this *SchemaManager) appendPrototypeKeyRef(schema *XML_Schemas.XSD_Schema, parent *EntityPrototype, keyRef *XML_Schemas.XSD_KeyRef) {
	this.appendPrototypeSelector(schema, parent, &keyRef.Selector)

	for i := 0; i < len(keyRef.Fields); i++ {
		this.appendPrototypeField(schema, parent, &keyRef.Fields[i])
	}
}

/**
 * Append key element
 */
func (this *SchemaManager) appendPrototypeKey(schema *XML_Schemas.XSD_Schema, parent *EntityPrototype, key *XML_Schemas.XSD_Key) {

	this.appendPrototypeSelector(schema, parent, &key.Selector)

	for i := 0; i < len(key.Fields); i++ {
		this.appendPrototypeField(schema, parent, &key.Fields[i])
	}
}

/**
 *  Append the content of a choice to a prototype...
 */
func (this *SchemaManager) appendPrototypeChoice(schema *XML_Schemas.XSD_Schema, prototype *EntityPrototype, choice *XML_Schemas.XSD_Choice, minOccurs string, maxOccurs string) {
	// Append Any...
	for i := 0; i < len(choice.Anys); i++ {
		this.appendPrototypeAny(schema, prototype, &choice.Anys[i])
	}

	// Append sequences...
	for i := 0; i < len(choice.Sequences); i++ {
		this.appendPrototypeSequence(schema, prototype, &choice.Sequences[i], choice.Sequences[i].MinOccurs, choice.Sequences[i].MaxOccurs)
		this.setLastPrototypeFieldCardinality(prototype, minOccurs, maxOccurs)
	}

	// Append elements...
	for i := 0; i < len(choice.Elements); i++ {
		element := choice.Elements[i]
		this.appendPrototypeElement(schema, prototype, &element, element.MinOccurs, element.MaxOccurs)
		this.setLastPrototypeFieldCardinality(prototype, minOccurs, maxOccurs)
	}

	// Append groups...
	for i := 0; i < len(choice.Groups); i++ {
		group := choice.Groups[i]
		this.appendPrototypeGroup(schema, prototype, &group, group.MinOccurs, group.MaxOccurs)
		this.setLastPrototypeFieldCardinality(prototype, minOccurs, maxOccurs)
	}

	// Append choices...
	for i := 0; i < len(choice.Choices); i++ {
		this.appendPrototypeChoice(schema, prototype, &choice.Choices[i], choice.Choices[i].MinOccurs, choice.Choices[i].MaxOccurs)
		this.setLastPrototypeFieldCardinality(prototype, minOccurs, maxOccurs)
	}
}

/**
 *  Append the content of a all to a prototype...
 */
func (this *SchemaManager) appendPrototypeAll(schema *XML_Schemas.XSD_Schema, prototype *EntityPrototype, all *XML_Schemas.XSD_All, minOccurs string, maxOccurs string) {
	// here I will append the list of all element...
	for i := 0; i < len(all.Elements); i++ {
		// Append Element from all to the prototype...
		element := all.Elements[i]
		this.appendPrototypeElement(schema, prototype, &element, element.MinOccurs, element.MaxOccurs)
		this.setLastPrototypeFieldCardinality(prototype, minOccurs, maxOccurs)
	}
}

/**
 * Append union... The content can be one of simple type...
 */
func (this *SchemaManager) appendPrototypeUnion(schema *XML_Schemas.XSD_Schema, prototype *EntityPrototype, union *XML_Schemas.XSD_Union) {
	// here I will append the union of element...
	for i := 0; i < len(union.SimpleType); i++ {
		this.createPrototypeSimpleType(schema, &union.SimpleType[i])
	}
}

func (this *SchemaManager) appendPrototypeSuperBaseType(superTypeName string, prototype *EntityPrototype) {
	superTypePrototype := this.prototypes[superTypeName]

	if prototype.TypeName != superTypeName && superTypePrototype != nil {
		// Apppend the supertype into the supertype list of the input prototype.
		if Utility.Contains(prototype.SuperTypeNames, superTypeName) == false {
			prototype.SuperTypeNames = append(prototype.SuperTypeNames, superTypeName)
		}
		// Append the prototype into the substitution group of the supertype...
		if Utility.Contains(superTypePrototype.SubstitutionGroup, prototype.TypeName) == false {
			superTypePrototype.SubstitutionGroup = append(superTypePrototype.SubstitutionGroup, prototype.TypeName)
		}

	}
}

// Extention and restriction...
/**
 * Append new prototype Extention.
 */
func (this *SchemaManager) appendPrototypeExtention(schema *XML_Schemas.XSD_Schema, prototype *EntityPrototype, extension *XML_Schemas.XSD_Extension, complexContent *XML_Schemas.XSD_ComplexContent, simpleContent *XML_Schemas.XSD_SimpleContent) {
	if complexContent != nil {
		if len(extension.Base) > 0 {
			// So here I will find the base type content...
			if val, ok := this.globalComplexTypes[schema.Id+"."+removeNs(extension.Base)]; ok {
				this.createPrototypeComplexType(schema, val)
				this.appendPrototypeSuperBaseType(schema.Id+"."+val.Name, prototype)

			} else if val, ok := this.globalSimpleTypes[schema.Id+"."+removeNs(extension.Base)]; ok {
				this.createPrototypeSimpleType(schema, val)
				this.appendPrototypeSuperBaseType(schema.Id+"."+val.Name, prototype)
			}
		}

		if extension.All != nil {
			this.appendPrototypeAll(schema, prototype, extension.All, extension.All.MinOccurs, extension.All.MaxOccurs)
		} else if extension.Choice != nil {
			this.appendPrototypeChoice(schema, prototype, extension.Choice, extension.Choice.MinOccurs, extension.Choice.MaxOccurs)
		} else if extension.Sequence != nil {
			this.appendPrototypeSequence(schema, prototype, extension.Sequence, extension.Sequence.MinOccurs, extension.Sequence.MaxOccurs)
		} else if extension.Group != nil {
			this.appendPrototypeGroup(schema, prototype, extension.Group, extension.Group.MinOccurs, extension.Group.MaxOccurs)
		}

	} else if simpleContent != nil {
		if len(extension.Base) > 0 {
			// The base extention...

			if val, ok := this.globalComplexTypes[schema.Id+"."+removeNs(extension.Base)]; ok {
				this.createPrototypeComplexType(schema, val)
				this.appendPrototypeSuperBaseType(schema.Id+"."+extension.Base, prototype)
			} else if val, ok := this.globalSimpleTypes[schema.Id+"."+removeNs(extension.Base)]; ok {
				this.createPrototypeSimpleType(schema, val)
				this.appendPrototypeSuperBaseType(schema.Id+"."+extension.Base, prototype)
			} else {
				baseName := extension.Base[strings.Index(extension.Base, ":")+1:]
				if _, ok := this.globalComplexTypes["xs."+baseName]; ok {
					this.appendPrototypeSuperBaseType("xs."+baseName, prototype)
				} else if _, ok := this.globalSimpleTypes["xs."+baseName]; ok {
					this.appendPrototypeSuperBaseType("xs."+baseName, prototype)
				}
			}

		}
	}

	for i := 0; i < len(extension.Attributes); i++ {
		this.appendPrototypeAttribute(schema, prototype, &extension.Attributes[i])
	}

	for i := 0; i < len(extension.AttributeGroups); i++ {
		this.appendPrototypeAttributeGroup(schema, prototype, &extension.AttributeGroups[i])
	}

	for i := 0; i < len(extension.AnyAttributes); i++ {
		this.appendPrototypeAnyAttribute(schema, prototype, &extension.AnyAttributes[i])
	}
}

/**
 * Append new restriction to the prototype...
 */
func (this *SchemaManager) appendPrototypeRestriction(schema *XML_Schemas.XSD_Schema, prototype *EntityPrototype, restriction *XML_Schemas.XSD_Restriction, simpleType *XML_Schemas.XSD_SimpleType, complexContent *XML_Schemas.XSD_ComplexContent, simpleContent *XML_Schemas.XSD_SimpleContent) {

	if len(restriction.Base) > 0 {
		// The base type name...
		baseName := restriction.Base[strings.Index(restriction.Base, ":")+1:]
		// Now the facet that will limit the value that can be set as values...
		if complexContent != nil || simpleContent != nil {
			// Set the prototype SuperTypeName...
			if simpleType != nil {
				this.appendPrototypeSuperBaseType(schema.Id+"."+simpleType.Name, prototype)
			} else {
				// Here the base type will be take from the restriction base.
				this.appendPrototypeSuperBaseType(schema.Id+"."+restriction.Base, prototype)
			}

			// In the context of simple content
			for i := 0; i < len(restriction.Attributes); i++ {
				this.appendPrototypeAttribute(schema, prototype, &restriction.Attributes[i])
			}

			for i := 0; i < len(restriction.AttributeGroups); i++ {
				this.appendPrototypeAttributeGroup(schema, prototype, &restriction.AttributeGroups[i])
			}

			for i := 0; i < len(restriction.AnyAttributes); i++ {
				this.appendPrototypeAnyAttribute(schema, prototype, &restriction.AnyAttributes[i])
			}

			if complexContent != nil {

				// In the contex of the complex content.
				// Append All element if there ones...
				if restriction.All != nil {
					this.appendPrototypeAll(schema, prototype, restriction.All, restriction.All.MinOccurs, restriction.All.MaxOccurs)
				} else if restriction.Choice != nil {
					this.appendPrototypeChoice(schema, prototype, restriction.Choice, restriction.Choice.MinOccurs, restriction.Choice.MaxOccurs)
				} else if restriction.Group != nil {
					log.Println("-----------> restriction", restriction)
					this.appendPrototypeGroup(schema, prototype, restriction.Group, restriction.Group.MinOccurs, restriction.Group.MaxOccurs)
				} else if restriction.Sequence != nil {
					this.appendPrototypeSequence(schema, prototype, restriction.Sequence, restriction.Sequence.MinOccurs, restriction.Sequence.MaxOccurs)
				}
			}

		} else if simpleType != nil {

			if baseName == "anySimpleType" {
				this.appendPrototypeSuperBaseType(schema.Id+"."+simpleType.Name, prototype)
			} else {
				// So here I will find the base type content...
				if val, ok := this.globalSimpleTypes["xs."+baseName]; ok {
					this.createPrototypeSimpleType(this.schemas["xs"], val)
					this.appendPrototypeSuperBaseType("xs."+val.Name, prototype)
				} else if val, ok := this.globalSimpleTypes[schema.Id+"."+baseName]; ok {
					this.createPrototypeSimpleType(schema, val)
					this.appendPrototypeSuperBaseType(schema.Id+"."+val.Name, prototype)
				} else if restriction.SimpleType != nil {
					this.createPrototypeSimpleType(schema, restriction.SimpleType)
					this.appendPrototypeSuperBaseType(schema.Id+"."+val.Name, prototype)
				}
			}
		}
	}

	// Facets are not applicable to complex content...
	if complexContent == nil || simpleType != nil {
		facets := make([]*Restriction, 0)

		// Here I will set the prototype restriction...
		for i := 0; i < len(restriction.MaxExclusive); i++ {
			facet := new(Restriction)
			facet.Type = RestrictionType_MaxExclusive
			facet.Value = restriction.MaxExclusive[i].Value
			facets = append(facets, facet)
		}

		for i := 0; i < len(restriction.MinExclusive); i++ {
			facet := new(Restriction)
			facet.Type = RestrictionType_MinExclusive
			facet.Value = restriction.MaxExclusive[i].Value
			facets = append(facets, facet)
		}

		for i := 0; i < len(restriction.MaxInclusive); i++ {
			facet := new(Restriction)
			facet.Type = RestrictionType_MaxInclusive
			facet.Value = restriction.MaxInclusive[i].Value
			facets = append(facets, facet)
		}

		for i := 0; i < len(restriction.MinInclusive); i++ {
			facet := new(Restriction)
			facet.Type = RestrictionType_MinInclusive
			facet.Value = restriction.MinInclusive[i].Value
			facets = append(facets, facet)
		}

		for i := 0; i < len(restriction.MaxLength); i++ {
			facet := new(Restriction)
			facet.Type = RestrictionType_MaxLength
			facet.Value = restriction.MaxLength[i].Value
			facets = append(facets, facet)
		}

		for i := 0; i < len(restriction.MinLength); i++ {
			facet := new(Restriction)
			facet.Type = RestrictionType_MinLength
			facet.Value = restriction.MinLength[i].Value
			facets = append(facets, facet)
		}

		for i := 0; i < len(restriction.Length); i++ {
			facet := new(Restriction)
			facet.Type = RestrictionType_Length
			facet.Value = restriction.Length[i].Value
			facets = append(facets, facet)
		}

		for i := 0; i < len(restriction.FractionDigits); i++ {
			facet := new(Restriction)
			facet.Type = RestrictionType_FractionDigits
			facet.Value = restriction.FractionDigits[i].Value
			facets = append(facets, facet)
		}

		for i := 0; i < len(restriction.TotalDigits); i++ {
			facet := new(Restriction)
			facet.Type = RestrictionType_TotalDigits
			facet.Value = restriction.TotalDigits[i].Value
			facets = append(facets, facet)
		}

		for i := 0; i < len(restriction.WhiteSpace); i++ {
			facet := new(Restriction)
			facet.Type = RestrictionType_WhiteSpace
			facet.Value = restriction.WhiteSpace[i].Value
			facets = append(facets, facet)
		}

		for i := 0; i < len(restriction.Pattern); i++ {
			facet := new(Restriction)
			facet.Type = RestrictionType_Pattern
			facet.Value = restriction.Pattern[i].Value
			facets = append(facets, facet)
		}

		for i := 0; i < len(restriction.Enumeration); i++ {
			facet := new(Restriction)
			facet.Type = RestrictionType_Enumeration
			facet.Value = restriction.Enumeration[i].Value
			facets = append(facets, facet)

		}

		prototype.Restrictions = facets
	}
	// Now the facets...
	for i := 0; i < len(restriction.AttributeGroups); i++ {
		this.appendPrototypeAttributeGroup(schema, prototype, &restriction.AttributeGroups[i])
	}
}

/**
 * Recursively get the field type name.
 */
func (this *SchemaManager) getFieldType(typeName string, fieldName string) string {
	prototype := this.prototypes[typeName]
	index := prototype.getFieldIndex(fieldName)
	// First of all i will
	if index > -1 {
		return prototype.FieldsType[index]
	}

	// Now I will try to get it from it super types.
	for i := 0; i < len(prototype.SuperTypeNames); i++ {
		superType := prototype.SuperTypeNames[i]
		if _, ok := this.prototypes[superType]; !ok {
			// Here will try to find if a global element exist
			if element, ok := this.globalElements[removeNs(superType)]; ok {
				if len(element.Type) > 0 {
					superType = superType[0:strings.Index(superType, ".")+1] + element.Type
				}
			}
		}

		if typeName != superType {
			fieldType := this.getFieldType(superType, fieldName)
			if len(fieldType) > 0 {
				return fieldType
			}
		} else {
			return typeName
		}
	}

	// nothing was found...
	return ""
}

////////////////////////////////////////////////////////////////////////////////
// XML Stuff here...
////////////////////////////////////////////////////////////////////////////////

type XmlDocumentHandler struct {

	/** The schema to use to interpret the data. **/
	SchemaId string

	/** Keep entities reference here... **/
	references map[string]*DynamicEntity

	/** Contain the stack of element currently process. **/
	objects Utility.Stack

	/** Contain the list of object to be saved **/
	globalObjects []map[string]interface{}

	// The last property to set
	lastProperty string
}

/**
 * Init the content of the entity from the content of the element.
 */
func (this *XmlDocumentHandler) InitElement(entity *DynamicEntity, element *XML_Schemas.XSD_Element) {

}

/**
 * Create an new handler to process the xml document...
 */
func NewXmlDocumentHandler() *XmlDocumentHandler {

	// The xml handler...
	handler := new(XmlDocumentHandler)

	// Set the global object array.
	handler.globalObjects = make([]map[string]interface{}, 0)

	return handler
}

func (this *XmlDocumentHandler) getLastObject() map[string]interface{} {
	lastObject := this.objects.Pop()
	if lastObject != nil {
		this.objects.Push(lastObject) // keep in the stack...
		return lastObject
	}
	return nil
}

func (this *XmlDocumentHandler) StartDocument() {

}

/**
 * Remove xs object and keep only there values...
 */
func compactObject(object map[string]interface{}) {
	for key, value := range object {
		switch v := value.(type) {
		case map[string]interface{}:
			if v["TYPENAME"] != nil {
				// if the type is a base type I will take it value...
				if strings.HasPrefix(v["TYPENAME"].(string), "xs.") {
					// Keep the value only...
					object[key] = v["M_valueOf"]
				} else {
					compactObject(v)
				}
			}
		case []interface{}:
			for i := 0; i < len(v); i++ {
				if reflect.TypeOf(v[i]).String() == "map[string]interface {}" {
					v_ := v[i].(map[string]interface{})
					if v_["TYPENAME"] != nil {
						// if the type is a base type I will take it value...
						if strings.HasPrefix(v_["TYPENAME"].(string), "xs.") {
							// Keep the value only...
							v[i] = v_["M_valueOf"]
						} else {
							compactObject(v_)
						}
					}
				}
			}
		}
	}
}

func (this *XmlDocumentHandler) EndDocument() {
	// Here I will compact the object

	// Here I will create dynamic entities for those objects.
	for i := 0; i < len(this.globalObjects); i++ {
		object := this.globalObjects[i]
		compactObject(object)
		object["NeedSave"] = true
		entity, errObj := GetServer().GetEntityManager().newDynamicEntity("", object)
		if errObj == nil {
			entity.SaveEntity()

		}
	}
}

func (this *XmlDocumentHandler) setObjectValue(object map[string]interface{}, typeName string, name string, value string, isArray bool) interface{} {
	// Set the object value...
	if typeName == "xs.boolean" {
		//////////////////////////// Boolean types ////////////////////////////
		if value == "true" {
			if isArray == true {
				object[name] = append(object[name].([]interface{}), true)
			} else {
				object[name] = true
			}
		} else {
			if isArray == true {
				object[name] = append(object[name].([]interface{}), false)
			} else {
				object[name] = false
			}
		}
	} else if typeName == "xs.byte" || typeName == "xs.short" || typeName == "xs.int" || typeName == "xs.integer" || typeName == "xs.long" || typeName == "xs.positiveInteger" || typeName == "xs.nonPositiveInteger" || typeName == "xs.nonNegativeInteger" || typeName == "xs.negativeInteger" {
		//////////////////////////// Integer types ////////////////////////////
		if typeName == "xs.byte" {
			i, err := strconv.ParseInt(value, 10, 8)
			if err != nil {
				panic(err)
			}
			if isArray == true {
				object[name] = append(object[name].([]interface{}), int8(i))
			} else {
				object[name] = int8(i)
			}

		} else if typeName == "xs.short" {
			i, err := strconv.ParseInt(value, 10, 16)
			if err != nil {
				panic(err)
			}
			if isArray == true {
				object[name] = append(object[name].([]interface{}), int32(i))
			} else {
				object[name] = int32(i)
			}
		} else if typeName == "xs.int" || typeName == "xs.integer" || typeName == "xs.positiveInteger" || typeName == "xs.nonPositiveInteger" || typeName == "xs.nonNegativeInteger" || typeName == "xs.negativeInteger" || typeName == "xs.timestampNumeric" {
			i, err := strconv.ParseInt(value, 10, 32)
			if err != nil {
				panic(err)
			}
			if isArray == true {
				object[name] = append(object[name].([]interface{}), int32(i))
			} else {
				object[name] = int32(i)
			}
		} else {
			i, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				panic(err)
			}
			if isArray == true {
				object[name] = append(object[name].([]interface{}), i)
			} else {
				object[name] = i
			}
		}

	} else if typeName == "xs.unsignedByte" || typeName == "xs.unsignedShort" || typeName == "xs.unsignedInt" || typeName == "xs.unsignedLong" {
		//////////////////////////// Unsigned types ////////////////////////////
		if typeName == "xs.unsignedByte" {
			i, err := strconv.ParseUint(value, 10, 8)
			if err != nil {
				panic(err)
			}
			if isArray == true {
				object[name] = append(object[name].([]interface{}), uint8(i))
			} else {
				object[name] = uint8(i)
			}
		} else if typeName == "xs.unsignedShort" {
			i, err := strconv.ParseUint(value, 10, 16)
			if err != nil {
				panic(err)
			}
			if isArray == true {
				object[name] = append(object[name].([]interface{}), uint32(i))
			} else {
				object[name] = uint32(i)
			}
		} else if typeName == "xs.unsignedInt" {

			i, err := strconv.ParseUint(value, 10, 32)
			if err != nil {
				panic(err)
			}
			if isArray == true {
				object[name] = append(object[name].([]interface{}), uint32(i))
			} else {
				object[name] = uint32(i)
			}
		} else {
			i, err := strconv.ParseUint(value, 10, 64)
			if err != nil {
				panic(err)
			}
			if isArray == true {
				object[name] = append(object[name].([]interface{}), uint64(i))
			} else {
				object[name] = uint64(i)
			}
		}
	} else if typeName == "xs.decimal" || typeName == "xs.double" || typeName == "xs.float" || typeName == "xs.money" || typeName == "xs.smallmoney" {
		//////////////////////////// Numeric types ////////////////////////////
		if typeName == "xs.float" {
			f, err := strconv.ParseFloat(value, 64)
			if err != nil {
				panic(err)
			}
			if isArray == true {
				object[name] = append(object[name].([]interface{}), float64(f))
			} else {
				object[name] = float32(f)
			}
		} else {
			f, err := strconv.ParseFloat(value, 32)
			if err != nil {
				panic(err)
			}
			if isArray == true {
				object[name] = append(object[name].([]interface{}), f)
			} else {
				object[name] = f
			}
		}
	} else if typeName == "xs.date" || typeName == "xs.datetime" || typeName == "xs.time" {
		//////////////////////////// Time and date ////////////////////////////
		if typeName == "xs.dateTime" {
			dateTime, _ := Utility.MatchISO8601_DateTime(value)

			if isArray == true {
				object[name] = append(object[name].([]interface{}), dateTime.UnixNano()/int64(time.Millisecond))
			} else {
				object[name] = dateTime
			}
		} else if typeName == "xs.time" {
			time_, _ := Utility.MatchISO8601_Time(value)

			if isArray == true {
				object[name] = append(object[name].([]interface{}), time_.UnixNano()/int64(time.Millisecond))
			} else {
				object[name] = time_
			}
		} else if typeName == "xs.date" {
			date, _ := Utility.MatchISO8601_Date(value)
			if isArray == true {
				object[name] = append(object[name].([]interface{}), date.UnixNano()/int64(time.Millisecond))
			} else {
				object[name] = date
			}
		}
	} else {
		//////////////////////////// String types ////////////////////////////
		if isArray == true {
			object[name] = append(object[name].([]interface{}), value)
		} else {
			object[name] = value
		}
	}

	if typeName == "xs.ID" {
		// The list of object with their respective id's
		//this.objectsById[value] = object
	}

	return object[name]
}

/**
 * Process the content of the element
 */
func (this *XmlDocumentHandler) StartElement(e xml.StartElement) {
	attributes := e.Attr
	elementNameLocal := e.Name.Local
	var object map[string]interface{}

	// Get the element schema...
	if SchemaId, ok := GetServer().GetSchemaManager().elementSchema[elementNameLocal]; ok {
		this.SchemaId = SchemaId
	}

	// If the SchemaId is not already set I will try to get it
	// from the list of attribute of the element
	if len(this.SchemaId) == 0 {
		for i := 0; i < len(attributes); i++ {
			attr := attributes[i]
			if attr.Name.Local == "xmlns" {
				this.SchemaId = attr.Value[strings.LastIndex(attr.Value, "/")+1:]
			}
		}
	}

	if len(this.SchemaId) > 0 {
		// Get the parent of this object...
		lastObject := this.getLastObject()
		if lastObject != nil {
			// The element is a member of the last object...
			prototype, _ := GetServer().GetEntityManager().getEntityPrototype(lastObject["TYPENAME"].(string), lastObject["TYPENAME"].(string)[0:strings.Index(lastObject["TYPENAME"].(string), ".")])
			fieldType := GetServer().GetSchemaManager().getFieldType(prototype.TypeName, "M_"+elementNameLocal)
			// Here I will try to
			if len(fieldType) > 0 {
				// The field type...
				fieldType := strings.Replace(fieldType, "[]", "", -1)
				fieldType = strings.Replace(fieldType, ":Ref", "", -1)
				object = make(map[string]interface{})
				object["TYPENAME"] = fieldType
				object["NeedSave"] = true
				this.objects.Push(object)
			} else {
				// The element is not found with it name in the prototype...
				// so I will try to get it from the global scope...
				var element *XML_Schemas.XSD_Element
				if val, ok := GetServer().GetSchemaManager().globalElements[this.SchemaId+"."+elementNameLocal]; ok {
					element = val
					// The element can be name with it superType name in case of abstract type...
					if len(element.SubstitutionGroup) > 0 {
						index := prototype.getFieldIndex("M_" + element.SubstitutionGroup)
						if index > 0 {
							elementNameLocal = element.SubstitutionGroup
							fieldType := strings.Replace(prototype.FieldsType[index], "[]", "", -1)
							fieldType = strings.Replace(fieldType, ":Ref", "", -1)
							object = make(map[string]interface{})
							if len(element.Type) > 0 {
								object["TYPENAME"] = this.SchemaId + "." + element.Type
							} else {
								object["TYPENAME"] = this.SchemaId + "." + element.Name
							}
							object["NeedSave"] = true
							this.objects.Push(object)
						} else {
							log.Println(" 1615 -------------> Element not found ", e.Name.Local)
						}
					} else {
						log.Println(" 1617 -------------> Element not found ", e.Name.Local)
					}
				} else {
					log.Println(" 1621 -------------> Element not found ", e.Name.Local)
				}
			}
		} else {
			// Try to get element from global element here.
			var element *XML_Schemas.XSD_Element
			if val, ok := GetServer().GetSchemaManager().globalElements[this.SchemaId+"."+elementNameLocal]; ok {
				element = val
			}
			object = make(map[string]interface{})
			// Here the object is in the global scope so the element must no]
			// be nil
			if len(element.Type) > 0 {
				object["TYPENAME"] = this.SchemaId + "." + element.Type
			} else {
				object["TYPENAME"] = this.SchemaId + "." + elementNameLocal
			}
			object["NeedSave"] = true
			this.objects.Push(object)
			this.globalObjects = append(this.globalObjects, object)
		}

		if object != nil {
			// Here I will set the attributes...
			storeId := object["TYPENAME"].(string)[0:strings.Index(object["TYPENAME"].(string), ".")]
			prototype, err := GetServer().GetEntityManager().getEntityPrototype(object["TYPENAME"].(string), storeId)
			if err != nil {
				log.Panicln("No prototype found for class ", object["TYPENAME"].(string))
			}

			// Set the attributes...
			for i := 0; i < len(attributes); i++ {
				attr := attributes[i]
				attrName := "M_" + attr.Name.Local
				// I will get the index of that field to test if it exist in the
				// prototype...
				fieldType := GetServer().GetSchemaManager().getFieldType(prototype.TypeName, attrName)
				if len(fieldType) > 0 {
					// In that case the field is not a base type so I need to create a object of the good type and set the
					// value of field...
					attrObject := make(map[string]interface{})
					attrObject["TYPENAME"] = fieldType
					attrObject["M_valueOf"] = attr.Value
					attrObject["NeedSave"] = true
					object[attrName] = attrObject
				}
			}

			// Here I will set the object into the parent...
			if lastObject != nil {
				this.lastProperty = "M_" + elementNameLocal
				prototype, _ := GetServer().GetEntityManager().getEntityPrototype(lastObject["TYPENAME"].(string), lastObject["TYPENAME"].(string)[0:strings.Index(lastObject["TYPENAME"].(string), ".")])
				fieldType := GetServer().GetSchemaManager().getFieldType(prototype.TypeName, "M_"+elementNameLocal)
				//log.Println(lastObject["TYPENAME"].(string), "-> M_"+elementNameLocal, ":", fieldType)
				isArray := strings.Index(fieldType, "[]") > -1
				// If the field type is tag as reference...
				if isArray {
					if lastObject["M_"+elementNameLocal] == nil {
						lastObject[this.lastProperty] = make([]interface{}, 0)
					}
					lastObject[this.lastProperty] = append(lastObject[this.lastProperty].([]interface{}), object)
				} else {
					lastObject[this.lastProperty] = object
				}
			}
		}

	} else {

		log.Println("Schema not found!")
	}
}

func (this *XmlDocumentHandler) EndElement(e xml.EndElement) {
	this.objects.Pop()
	this.lastProperty = ""
}

func getRootTypeName(typeName string) string {
	typeName = strings.Replace(typeName, "[]", "", -1)

	prototype, err := GetServer().GetEntityManager().getEntityPrototype(typeName, typeName[0:strings.Index(typeName, ".")])
	if err == nil {
		if prototype.SuperTypeNames != nil {
			if len(prototype.SuperTypeNames) > 0 {
				if strings.HasPrefix(typeName, "xs.") {
					return prototype.TypeName
				} else {
					return getRootTypeName(prototype.SuperTypeNames[0])
				}
			}
		}

		if len(prototype.ListOf) > 0 {
			return getRootTypeName(prototype.ListOf)
		}

		return prototype.TypeName
	}

	return ""

}

/**
 * Recursive function to determine if the data type must be considerate as an
 * array.
 */
func isArray(typeName string) bool {

	if strings.HasPrefix(typeName, "[]") {
		return true
	}

	if strings.HasPrefix(typeName, "xs.") {
		return false
	}

	// retrieve the type name...
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype(typeName, typeName[0:strings.Index(typeName, ".")])
	// Must be considerate as an array...
	if len(prototype.ListOf) > 0 {
		return true
	}

	// Find in super type...
	for i := 0; i < len(prototype.SuperTypeNames); i++ {
		if isArray(prototype.SuperTypeNames[i]) == true {
			return true
		}
	}

	// Now restriction...
	for i := 0; i < len(prototype.Restrictions); i++ {
		restriction := prototype.Restrictions[i]
		if restriction.Type == RestrictionType_Length || restriction.Type == RestrictionType_MinExclusive || restriction.Type == RestrictionType_MinInclusive || restriction.Type == RestrictionType_MaxExclusive || restriction.Type == RestrictionType_MaxInclusive || restriction.Type == RestrictionType_MaxLength || restriction.Type == RestrictionType_MinLength {
			return true
		}
	}

	return false
}

func (this *XmlDocumentHandler) CharData(e xml.CharData) {
	// Set the data...
	if len(string([]byte(e))) > 0 {
		if strings.TrimSpace(string([]byte(e))) != "" {
			// So here I will get the last object
			object := this.getLastObject()
			if object != nil {
				if isArray(object["TYPENAME"].(string)) {
					// Here I will split the values...
					values := strings.Split(string([]byte(e)), " ")
					object["M_listOf"] = make([]interface{}, 0)
					for i := 0; i < len(values); i++ {
						this.setObjectValue(object, object["TYPENAME"].(string), "M_listOf", values[i], true)
					}
				} else {
					this.setObjectValue(object, object["TYPENAME"].(string), "M_valueOf", string([]byte(e)), false)
				}
			}
		}
	}
}

func (this *XmlDocumentHandler) Comment(e xml.Comment) {

}

func (this *XmlDocumentHandler) ProcInst(xml.ProcInst) {

}

func (this *XmlDocumentHandler) Directive(xml.Directive) {

}

/**
 * The xml file to read must be place in the schema directory. Everything else
 * than xsd extension will be read as xml file.
 */
func (this *SchemaManager) importXmlFile(filePath string) error {
	source, err := ioutil.ReadFile(filePath)
	if err == nil {
		r := bytes.NewReader([]byte(source))
		handler := NewXmlDocumentHandler()
		parser := saxlike.NewParser(r, handler)
		parser.Parse()
	} else {
		return err
	}
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// Api
////////////////////////////////////////////////////////////////////////////////
