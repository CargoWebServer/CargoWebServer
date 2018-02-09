package Server

import "fmt"
import "strings"
import "log"
import "code.myceliUs.com/Utility"
import "code.myceliUs.com/CargoWebServer/Cargo/JS"
import "code.myceliUs.com/XML_Schemas"
import "strconv"
import "reflect"

/**
 * Restrictions for Datatypes
 */
type RestrictionType int

const (
	// Defines a list of acceptable values
	RestrictionType_Enumeration RestrictionType = 1 + iota

	// Specifies the maximum number of decimal places allowed. Must be equal to
	// or greater than zero
	RestrictionType_FractionDigits

	// Specifies the exact number of characters or list items allowed. Must be
	// equal to or greater than zero
	RestrictionType_Length

	// Specifies the upper bounds for numeric values (the value must be less
	// than this value)
	RestrictionType_MaxExclusive

	// Specifies the upper bounds for numeric values (the value must be less than
	// or equal to this value)
	RestrictionType_MaxInclusive

	// Specifies the maximum number of characters or list items allowed. Must be
	// equal to or greater than zero
	RestrictionType_MaxLength

	// Specifies the lower bounds for numeric values (the value must be greater
	// than this value)
	RestrictionType_MinExclusive

	// Specifies the lower bounds for numeric values (the value must be greater than or equal to this value)
	RestrictionType_MinInclusive

	// Specifies the minimum number of characters or list items allowed. Must be equal to or greater than zero
	RestrictionType_MinLength

	// Defines the exact sequence of characters that are acceptable
	RestrictionType_Pattern

	// Specifies the exact number of digits allowed. Must be greater than zero
	RestrictionType_TotalDigits

	// Specifies how white space (line feeds, tabs, spaces, and carriage returns) is handled
	RestrictionType_WhiteSpace
)

// Put constaint in a field to reduce the range of
// possibles values of a given type.
// For example an email is a string with a pattern to respect.
// so the range of string restrict by a pattern became the range of
// email.
type Restriction struct {
	TYPENAME string

	// The the of the restriction (Facet)
	Type RestrictionType
	// The value...
	Value string
}

/**
 * This structure is use to make query over the key value data store.
 */
type EntityPrototype struct {

	// Uniquely identify the entity prototype.
	UUID string

	// The name of the entity
	// The type name is compose of the package name, a comma and
	// the type name itself.
	TypeName string

	// The documentation for that entity
	Documentation string

	// True if the entity prototype is an abstrac class...
	IsAbstract bool

	// In that case the prototype define a list of given item type.
	ListOf string

	// The class derived from this entity.
	SubstitutionGroup []string

	// The list of super type, equivalent to extension.
	SuperTypeNames []string

	// Restriction of the range of possible value.
	Restrictions []*Restriction

	// The ids that compose the entity...
	Ids []string

	// The indexation of this entity
	Indexs []string

	// The list of fields of the entity
	Fields []string

	// That contain the field documentation if there is so...
	FieldsDocumentation []string

	// The list of fields type of the entity
	// ex. []string:Test.Item:Ref
	// [] means the field is an array
	// string is the format of the reference in the case
	// of type other than xsd base type.
	// Type is written like PacakageName.TypeName
	// If the field is a reference to other element, (an aggregation)
	// Ref is needed at the end. Otherwise it considere at composition.
	FieldsType []string

	// Fields visibility
	FieldsVisibility []bool

	// If the field can be nil value...
	FieldsNillable []bool

	// The order of the field, use to display in tabular form...
	FieldsOrder []int

	// The fields default value.
	FieldsDefaultValue []string

	// The prototype version.
	Version string

	// Temporary container that old information about fields modifiaction.
	FieldsToDelete []int
	FieldsToUpdate []string
}

func NewEntityPrototype() *EntityPrototype {

	prototype := new(EntityPrototype)

	// Uniquely identify entity prototype.
	prototype.UUID = Utility.RandomUUID()

	prototype.Fields = make([]string, 0)
	prototype.FieldsOrder = make([]int, 0)
	prototype.FieldsType = make([]string, 0)
	prototype.FieldsDocumentation = make([]string, 0)
	prototype.FieldsNillable = make([]bool, 0)
	prototype.FieldsVisibility = make([]bool, 0)
	prototype.SuperTypeNames = make([]string, 0)
	prototype.Restrictions = make([]*Restriction, 0)
	prototype.Indexs = make([]string, 0)
	prototype.Ids = make([]string, 0)
	prototype.FieldsDefaultValue = make([]string, 0)

	// Append the default fields at begin...
	prototype.Fields = append(prototype.Fields, "UUID")
	prototype.Ids = append(prototype.Ids, "UUID")
	prototype.FieldsOrder = append(prototype.FieldsOrder, 0)
	prototype.FieldsType = append(prototype.FieldsType, "xs.string")
	prototype.FieldsVisibility = append(prototype.FieldsVisibility, false)
	prototype.FieldsDefaultValue = append(prototype.FieldsDefaultValue, "")
	prototype.FieldsNillable = append(prototype.FieldsNillable, false)

	prototype.Fields = append(prototype.Fields, "ParentUuid")
	prototype.Indexs = append(prototype.Indexs, "ParentUuid")
	prototype.FieldsOrder = append(prototype.FieldsOrder, 1)
	prototype.FieldsType = append(prototype.FieldsType, "xs.string")
	prototype.FieldsVisibility = append(prototype.FieldsVisibility, false)
	prototype.FieldsDefaultValue = append(prototype.FieldsDefaultValue, "")
	prototype.FieldsNillable = append(prototype.FieldsNillable, true)

	prototype.Fields = append(prototype.Fields, "ParentLnk")
	prototype.FieldsOrder = append(prototype.FieldsOrder, 2)
	prototype.FieldsType = append(prototype.FieldsType, "xs.string")
	prototype.FieldsVisibility = append(prototype.FieldsVisibility, false)
	prototype.FieldsDefaultValue = append(prototype.FieldsDefaultValue, "")
	prototype.FieldsNillable = append(prototype.FieldsNillable, true)

	return prototype
}

/**
 * This function is use to retreive the position in the array of a given field.
 */
func (this *EntityPrototype) getFieldIndex(fieldName string) int {

	if this.Fields != nil {
		for i := 0; i < len(this.Fields); i++ {
			if this.Fields[i] == fieldName {
				return i
			}
		}
	}
	return -1
}

/**
 * Save the new entity prototype in the data store.
 */
func (this *EntityPrototype) Create(storeId string) error {

	// Append the default fields at end...
	if len(storeId) == 0 {
		storeId = this.TypeName[:strings.Index(this.TypeName, ".")]
	}

	store := GetServer().GetDataManager().getDataStore(storeId).(*KeyValueDataStore)
	if store != nil {
		err := store.SetEntityPrototype(this)
		if err != nil {
			log.Println("Fail to save entity prototype ", this.TypeName, " in store id ", storeId)
			return err
		}
	}

	// Register it to the vm...
	JS.GetJsRuntimeManager().AppendScript("CargoWebServer", this.generateConstructor(), true)

	// Send event message...
	var eventDatas []*MessageData
	evtData := new(MessageData)
	evtData.TYPENAME = "Server.MessageData"
	evtData.Name = "prototype"

	evtData.Value = this
	eventDatas = append(eventDatas, evtData)
	evt, _ := NewEvent(NewPrototypeEvent, PrototypeEvent, eventDatas)
	GetServer().GetEventManager().BroadcastEvent(evt)

	return nil

}

/**
 * Save the new entity prototype in the data store.
 */
func (this *EntityPrototype) Save(storeId string) error {
	if len(storeId) == 0 {
		storeId = this.TypeName[:strings.Index(this.TypeName, ".")]
	}

	// Get information of the previous entity prototype.
	prototype, err := GetServer().GetEntityManager().getEntityPrototype(this.TypeName, storeId)
	if err != nil {
		return err
	}

	var store *KeyValueDataStore

	if reflect.TypeOf(GetServer().GetDataManager().getDataStore(storeId)).String() == "*Server.SqlDataStore" {
		store = GetServer().GetDataManager().getDataStore("sql_info").(*KeyValueDataStore)
	} else {
		store = GetServer().GetDataManager().getDataStore(storeId).(*KeyValueDataStore)
	}

	if store != nil {
		// Save it inside it supertype in substitution-group.
		for i := 0; i < len(this.SuperTypeNames); i++ {
			superTypeName := this.SuperTypeNames[i]
			superType, err := GetServer().GetEntityManager().getEntityPrototype(superTypeName, superTypeName[0:strings.Index(superTypeName, ".")])
			if err == nil {
				if !Utility.Contains(superType.SubstitutionGroup, this.TypeName) {
					superType.SubstitutionGroup = append(superType.SubstitutionGroup, this.TypeName)
					// Save the superType.
					err := superType.Save(superTypeName[0:strings.Index(superTypeName, ".")])
					if err != nil {
						return err
					}
				}
			} else {
				return err
			}
		}

		// I will remove it from substitution group as neeeded...
		for i := 0; i < len(prototype.SuperTypeNames); i++ {
			if !Utility.Contains(this.SuperTypeNames, prototype.SuperTypeNames[i]) {
				// Here I will remove the prototype from superType substitution group.
				superTypeName := prototype.SuperTypeNames[i]
				superType, err := GetServer().GetEntityManager().getEntityPrototype(superTypeName, superTypeName[0:strings.Index(superTypeName, ".")])
				if err != nil {
					return err
				}

				substitutionGroup := make([]string, 0)
				for j := 0; j < len(superType.SubstitutionGroup); j++ {
					if superType.SubstitutionGroup[j] != prototype.TypeName {
						substitutionGroup = append(substitutionGroup, superType.SubstitutionGroup[j])
					}
				}
				superType.SubstitutionGroup = substitutionGroup
				err = superType.Save(superTypeName[0:strings.Index(superTypeName, ".")])
				if err != nil {
					return err
				}
			}
		}

		err := store.saveEntityPrototype(this)
		if err != nil {
			log.Println("Fail to save entity prototype ", this.TypeName, " in store id ", storeId)
			return err
		} else {
			// Register it to the vm...
			JS.GetJsRuntimeManager().AppendScript("CargoWebServer/"+this.TypeName, this.generateConstructor(), true)

			var eventDatas []*MessageData
			evtData := new(MessageData)
			evtData.TYPENAME = "Server.MessageData"
			evtData.Name = "prototype"

			evtData.Value = this
			eventDatas = append(eventDatas, evtData)
			evt, _ := NewEvent(UpdatePrototypeEvent, PrototypeEvent, eventDatas)
			GetServer().GetEventManager().BroadcastEvent(evt)

			// Update local entities if the store is local.
			if store.m_ipv4 == "127.0.0.1" { // save if is local entity prototype only.
				entities, _ := GetServer().GetEntityManager().getEntities(prototype.TypeName, nil, storeId, false)

				// Remove the fields
				for i := 0; i < len(entities); i++ {
					entity := entities[i] // Must be a dynamic entity.

					// remove it...
					for j := 0; j < len(this.FieldsToDelete); j++ {
						field := prototype.Fields[this.FieldsToDelete[j]]
						if reflect.TypeOf(entity).String() == "*Server.DynamicEntity" {
							// Dynamic entity.
							entity.(*DynamicEntity).deleteValue(field)
						}

						entity.SetNeedSave(true)
						entity.SaveEntity() // Must be save before doing something else.
					}

					// update it...
					for j := 0; j < len(this.FieldsToUpdate); j++ {
						values := strings.Split(this.FieldsToUpdate[j], ":")
						if len(values) == 2 {
							indexFrom := prototype.getFieldIndex(values[0])
							indexTo := this.getFieldIndex(values[1])
							if indexFrom > -1 && indexTo > -1 {
								if values[0] != values[1] {
									if reflect.TypeOf(entity).String() == "*Server.DynamicEntity" {
										// Set the new value with the old one
										entity.(*DynamicEntity).setValue(values[1], entity.(*DynamicEntity).getValue(values[0]))
										// Delete the old one.
										entity.(*DynamicEntity).deleteValue(values[0])
									}

									entity.SetNeedSave(true)
									prototype.Fields[indexFrom] = values[1]
								}
								var fieldTypeTo = prototype.FieldsType[indexTo]
								var fieldTypeFrom = this.FieldsType[indexFrom]
								if fieldTypeFrom != fieldTypeTo {
									log.Println("------> change field type from ", fieldTypeFrom, "with", fieldTypeTo)
									// TODO set conversion rules here for each possible types.
								}
							}
						}
					}

					// Now set new fields value inside existing entities with their default
					// value.
					for j := 0; j < len(this.Fields); j++ {
						if !Utility.Contains(prototype.Fields, this.Fields[j]) {
							// I that case I will set the new field value inside the prototype.
							var value interface{}
							if strings.HasPrefix(this.FieldsType[j], "[]") {
								value = "undefined"
							} else {
								if XML_Schemas.IsXsString(this.FieldsType[j]) {
									value = this.FieldsDefaultValue[j]
								} else if XML_Schemas.IsXsInt(this.FieldsType[j]) || XML_Schemas.IsXsTime(this.FieldsType[j]) {
									value, _ = strconv.ParseInt(this.FieldsDefaultValue[j], 10, 64)
								} else if XML_Schemas.IsXsNumeric(this.FieldsType[j]) {
									value, _ = strconv.ParseFloat(this.FieldsDefaultValue[j], 64)
								} else if XML_Schemas.IsXsDate(this.FieldsType[j]) {
									value = Utility.MakeTimestamp()
								} else if XML_Schemas.IsXsBoolean(this.FieldsType[j]) {
									if this.FieldsDefaultValue[j] == "false" {
										value = false
									} else {
										value = true
									}
								} else {
									// Object here.
									value = "undefined"
								}
							}
							entity.SetNeedSave(true)
							if reflect.TypeOf(entity).String() == "*Server.DynamicEntity" {
								entity.(*DynamicEntity).setValue(this.Fields[j], value)
							}
						}
					}

					// Save the entity.
					entity.SaveEntity()
				}

				// Now the indexation key for ids and indexs...
				for i := 0; i < len(this.FieldsToUpdate); i++ {
					values := strings.Split(this.FieldsToUpdate[i], ":")
					if len(values) == 2 {
						if Utility.Contains(this.Indexs, values[1]) || Utility.Contains(this.Ids, values[1]) {
							oldKey := this.TypeName + ":" + values[0]
							newKey := this.TypeName + ":" + values[1]
							indexations, err := store.getValue(oldKey)
							if err == nil {
								// Here I will remove the oldKey...
								store.deleteValue(oldKey)
								store.setValue([]byte(newKey), indexations)
							}
						}
					}
				}
			}
		}
	}

	return nil
}

/**
 * Generate the JavaScript class defefinition.
 */
func (this *EntityPrototype) generateConstructor() string {
	var packageName string

	values := strings.Split(this.TypeName, ".")
	packageName = values[0]
	var constructorSrc = "var " + packageName + " = " + packageName + "|| {};\n"

	// create sub-namspace if there is some.
	if len(values) > 2 {
		for i := 1; i < len(values)-1; i++ {
			packageName += "." + values[i]
			constructorSrc += packageName + " = " + packageName + "|| {};\n"
		}
	}

	constructorSrc += this.TypeName + " = function(){\n"

	// Common properties share by all entity.
	constructorSrc += " this.__class__ = \"" + this.TypeName + "\"\n"
	constructorSrc += " this.TYPENAME = \"" + this.TypeName + "\"\n"
	constructorSrc += " this.UUID = \"\"\n"
	constructorSrc += " this.ParentUuid = \"\"\n"
	constructorSrc += " this.ParentLnk = \"\"\n"
	constructorSrc += " this.childsUuid = []\n"
	constructorSrc += " this.references = []\n"
	constructorSrc += " this.NeedSave = true\n"
	constructorSrc += " this.IsInit = false\n"
	constructorSrc += " this.exist = false\n"
	constructorSrc += " this.initCallback = undefined\n"

	// Create field and set her initial values.
	for i := 3; i < len(this.Fields); i++ {
		if len(this.FieldsDefaultValue[i]) != 0 {
			constructorSrc += " this." + this.Fields[i] + " = " + this.FieldsDefaultValue[i] + "\n"
		} else {
			if strings.HasPrefix(this.FieldsType[i], "[]") {
				constructorSrc += " this." + this.Fields[i] + " = []\n"
			} else {
				constructorSrc += " this." + this.Fields[i] + " = null\n"
			}
		}
	}

	// Keep the reference on the entity prototype.
	constructorSrc += " return this\n"
	constructorSrc += "}\n"
	return constructorSrc
}

/**
 * For debug purpose only...
 */
func (this *EntityPrototype) Print() {
	// The prototype Type Name...
	fmt.Println("\nTypeName:", this.TypeName)
	if len(this.SuperTypeNames) > 0 {
		fmt.Println("	Super Types:", this.SuperTypeNames)
	}

	if this.SubstitutionGroup != nil {
		fmt.Println("	Substitution Groups:", this.SubstitutionGroup)
	}

	if len(this.ListOf) > 0 {
		fmt.Println("	List of:", this.ListOf)
	}

	// Now the restrictions...
	if this.Restrictions != nil {
		for j := 0; j < len(this.Restrictions); j++ {
			if this.Restrictions[j].Type == RestrictionType_Enumeration {
				fmt.Println("	----> Enumration Restriction:", this.Restrictions[j].Value)
			} else if this.Restrictions[j].Type == RestrictionType_FractionDigits {
				fmt.Println("	----> Fraction Digits Restriction:", this.Restrictions[j].Value)
			} else if this.Restrictions[j].Type == RestrictionType_Length {
				fmt.Println("	----> Length Restriction:", this.Restrictions[j].Value)
			} else if this.Restrictions[j].Type == RestrictionType_MaxExclusive {
				fmt.Println("	----> Max Exclusive Restriction:", this.Restrictions[j].Value)
			} else if this.Restrictions[j].Type == RestrictionType_MaxInclusive {
				fmt.Println("	----> Max Inclusive Restriction:", this.Restrictions[j].Value)
			} else if this.Restrictions[j].Type == RestrictionType_MaxLength {
				fmt.Println("	----> Max Length Restriction:", this.Restrictions[j].Value)
			} else if this.Restrictions[j].Type == RestrictionType_MinExclusive {
				fmt.Println("	----> Min Exclusive Restriction:", this.Restrictions[j].Value)
			} else if this.Restrictions[j].Type == RestrictionType_MinInclusive {
				fmt.Println("	----> Min Inclusive Restriction:", this.Restrictions[j].Value)
			} else if this.Restrictions[j].Type == RestrictionType_MinLength {
				fmt.Println("	----> Min Length Restriction:", this.Restrictions[j].Value)
			} else if this.Restrictions[j].Type == RestrictionType_TotalDigits {
				fmt.Println("	----> Total Digits Restriction:", this.Restrictions[j].Value)
			} else if this.Restrictions[j].Type == RestrictionType_WhiteSpace {
				fmt.Println("	----> White Space Restriction:", this.Restrictions[j].Value)
			}
		}
	}

	// Now the fields...
	fmt.Println("	Fields:")
	for i := 0; i < len(this.Fields); i++ {
		fmt.Println("	-->", this.Fields[i], ":", this.FieldsType[i])
	}
}
