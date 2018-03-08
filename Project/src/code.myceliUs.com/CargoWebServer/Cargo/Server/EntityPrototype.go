package Server

import "fmt"
import "strings"
import "log"
import "code.myceliUs.com/Utility"

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
 * That function is use to create an extension of a given prototype.
 */
func (this *EntityPrototype) setSuperTypeFields() {
	var index = 3 // The start index is after the uuid and parentUuid and the parent lnk.
	if len(this.ListOf) > 0 {
		// It can not have superpetype and list of a the same time.
		this.SuperTypeNames = make([]string, 0)
		// In that particular case the entity is a list of the given type.
		if !Utility.Contains(this.Fields, "M_listOf") {
			Utility.InsertStringAt(index, "M_listOf", &this.Fields)
			Utility.InsertStringAt(index, "[]"+this.ListOf, &this.FieldsType)
			Utility.InsertBoolAt(index, true, &this.FieldsVisibility)
			Utility.InsertStringAt(index, "[]", &this.FieldsDefaultValue)
			Utility.InsertStringAt(index, "", &this.FieldsDocumentation)
		}
	} else {
		for i := 0; i < len(this.SuperTypeNames); i++ {
			superTypeName := this.SuperTypeNames[i]
			superPrototype, err := GetServer().GetEntityManager().getEntityPrototype(superTypeName, superTypeName[0:strings.Index(superTypeName, ".")])
			if err == nil {
				// I will merge the fields
				// The first to fields are always the uuid, parentUuid, parentLnk and the last is the childUuids and referenced
				for j := 3; j < len(superPrototype.Fields)-2; j++ {
					if !Utility.Contains(this.Fields, superPrototype.Fields[j]) && strings.HasPrefix(superPrototype.Fields[j], "M_") {

						Utility.InsertStringAt(index, superPrototype.Fields[j], &this.Fields)
						Utility.InsertStringAt(index, superPrototype.FieldsType[j], &this.FieldsType)
						Utility.InsertBoolAt(index, superPrototype.FieldsVisibility[j], &this.FieldsVisibility)
						Utility.InsertStringAt(index, superPrototype.FieldsDefaultValue[j], &this.FieldsDefaultValue)

						// create a new index at the end...
						if superPrototype.FieldsNillable != nil {
							isNillable := false
							if j < len(superPrototype.FieldsNillable) {
								isNillable = superPrototype.FieldsNillable[j]
							}
							Utility.InsertBoolAt(index, isNillable, &this.FieldsNillable)
						} else {
							this.FieldsNillable = append(this.FieldsNillable, true)
						}

						if superPrototype.FieldsDocumentation != nil {
							documentation := ""
							if j < len(superPrototype.FieldsDocumentation) {
								documentation = superPrototype.FieldsDocumentation[j]
							}
							if index < len(this.FieldsDocumentation) {
								Utility.InsertStringAt(index, documentation, &this.FieldsDocumentation)
							} else {
								this.FieldsDocumentation = append(this.FieldsDocumentation, documentation)
							}
						} else {
							this.FieldsDocumentation = append(this.FieldsDocumentation, "")
						}

						index++
					}
				}
				// Now the index...
				for j := 0; j < len(superPrototype.Indexs); j++ {
					if !Utility.Contains(this.Indexs, superPrototype.Indexs[j]) {
						this.Indexs = append(this.Indexs, superPrototype.Indexs[j])
					}
				}
				// Now the ids
				for j := 0; j < len(superPrototype.Ids); j++ {
					if !Utility.Contains(this.Ids, superPrototype.Ids[j]) {
						this.Ids = append(this.Ids, superPrototype.Ids[j])
					}
				}

				// Now I will append the new prototype to the list of substitution group of the super type.
				if !Utility.Contains(superPrototype.SubstitutionGroup, this.TypeName) {
					if !Utility.Contains(superPrototype.SubstitutionGroup, this.TypeName) {
						superPrototype.SubstitutionGroup = append(superPrototype.SubstitutionGroup, this.TypeName)
						// save it to it store...
						store := GetServer().GetDataManager().getDataStore(superTypeName[0:strings.Index(superTypeName, ".")])
						store.SaveEntityPrototype(superPrototype)
					}
				}
			} else {
				log.Println("error ", err)
			}
		}
	}

	// reset the field orders.
	this.FieldsOrder = make([]int, len(this.Fields))
	for i := 0; i < len(this.Fields); i++ {
		this.FieldsOrder[i] = i
	}

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
