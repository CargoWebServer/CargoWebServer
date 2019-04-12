package XML_Schemas

import (
	"encoding/xml"
	"strings"
)

/**
 * Informative data for human or electronic agents.
 * xs:annotation is a container in which additional information can be embedded,
 * either for human consumption (with xs:documentation) or for programs (xs:appinfo).
 */
type XSD_Annotation struct {
	XMLName xml.Name `xml:"annotation"`
	/** W3C XML Schema's element ID. **/
	Id string `xml:"id,attr"`

	/** Information that an application can read... **/
	AppInfos XSD_AppInfo `xml:"appinfo,omitempty"`

	/** Human readable documentation **/
	Documentation XSD_Documentation `xml:"documentation,omitempty"`
}

/**
 * xs:all is used to describe an unordered group of elements whose number of
 * occurences may be zero or one.
 */
type XSD_All struct {
	XMLName xml.Name `xml:"all"`

	/**
	 * W3C XML Schema's element ID.
	 */
	Id string `xml:"id,attr"`

	/**
	 * Maximum number of occurrences. Note that this value is fixed to one.
	 */
	MaxOccurs string `xml:"maxOccurs,attr"`

	/**
	 * Minimum number of occurrences. Note that this value can be only zero or one.
	 */
	MinOccurs string `xml:"minOccurs,attr"`

	Elements []XSD_Element `xml:"element,omitempty"`

	Annotations []XSD_Annotation `xml:"annotation,omitempty"`
}

/**
 * xs:appinfo is a container that embeds structured information that can be used
 * by applications. Its content model is open and can accept any element from any
 * namespace (with a lax validation; W3C XML Schema elements included here must
 * be valid). xs:appinfo can be used to include any kind of information, such as
 * metadata, processing directives, or even code snippets.
 */
type XSD_AppInfo struct {
	XMLName xml.Name `xml:"appinfo"`

	/**
	Can be used to provide a link to the source of the information when a
	snippet is included, or as a semantic attribute to qualify the type of
	information that is included.
	*/
	Source string `xml:",innerxml"`
}

/**
 * Human-targeted documentation.
 */
type XSD_Documentation struct {
	XMLName xml.Name `xml:"documentation"`

	/**
	Can be used to provide a link to the source of the information when a snippet
	is included, or it can be used as a semantic attribute to qualify the type of
	information included.
	*/
	Source string `xml:",innerxml"`

	/**
	Language used for the documentation.
	*/
	Lang string `xml:"name,attr"`
}

/**
 * Global definition of a complex type that can be referenced within the same
 * schema by other schemas.
 */
type XSD_ComplexType struct {
	XMLName xml.Name `xml:"complexType"`

	/**
	 * Name of the complex type.
	 */
	Name string `xml:"name,attr"`

	/**
	 * W3C XML Schema's element ID.
	 */
	Id string `xml:"id,attr"`

	/**
	When set to true, this complex type cannot be used directly in the instance
	documents and needs to be substituted using a xsi:type attribute.
	*/
	IsAbstract bool `xml:"abstract,attr"`

	/**
	 * "#all" | list of ( "extension" | "restriction"
	 */
	Final string `xml:"final,attr"`

	/**
	 * one of "#all" | list of ( "extension" | "restriction"
	 */
	Block string `xml:"block,attr"`

	/**
	 * Defines if the content model will be mixed.
	 */
	Mixed string `xml:"mixed,attr"`

	// Annotation for comment...
	Annotations []XSD_Annotation `xml:"annotation,omitempty"`

	// One of SimpleContent or ComplexContent...
	// Those are use when derivation is use...
	SimpleContent  *XSD_SimpleContent  `xml:"simpleContent,omitempty"`
	ComplexContent *XSD_ComplexContent `xml:"complexContent,omitempty"`

	// one of All, Group, Choice or Sequence...
	All      *XSD_All      `xml:"all,omitempty"`
	Group    *XSD_Group    `xml:"group,omitempty"`
	Choice   *XSD_Choice   `xml:"choice,omitempty"`
	Sequence *XSD_Sequence `xml:"sequence,omitempty"`

	// List of attribute local to the complexType...
	// Or one of Attribute or AttributeGroup
	Attributes      []XSD_Attribute      `xml:"attribute,omitempty"`
	AttributeGroups []XSD_AttributeGroup `xml:"attributeGroup,omitempty"`

	// Or... 0 or 1
	AnyAttributes []XSD_AnyAttribute `xml:"anyAttribute,omitempty"`
}

/**
 * Definition of a complex content by derivation of a complex type.
 */
type XSD_ComplexContent struct {
	XMLName xml.Name `xml:"complexContent"`
	/**
	 * W3C XML Schema's element ID.
	 */
	Id string `xml:"id,attr"`

	/**
	When set to true, the content model is mixed; when set to false, the
	content model is "element only"; when not set, the content model is
	determined by the mixed attribute of the parent xs:complexType element.
	*/
	Mixed bool `xml:"mixed,attr"`

	// Content
	// Annotation for comment...
	Annotations []XSD_Annotation `xml:"annotation,omitempty"`

	// One of...
	Extension   *XSD_Extension   `xml:"extension,omitempty"`
	Restriction *XSD_Restriction `xml:"restriction,omitempty"`
}

/**
 * Extension of a simple content model.
 */
type XSD_Extension struct {
	XMLName xml.Name `xml:"extension"`

	/**
	 * W3C XML Schema's element ID.
	 */
	Id string `xml:"id,attr"`

	/**
	Qualified name of the base type (simple type or simple content complex type).
	*/
	Base string `xml:"base,attr"`

	// Annotation for comment...
	Annotations []XSD_Annotation `xml:"annotation,omitempty"`

	// Content
	Attributes      []XSD_Attribute      `xml:"attribute,omitempty"`
	AttributeGroups []XSD_AttributeGroup `xml:"attributeGroup,omitempty"`
	AnyAttributes   []XSD_AnyAttribute   `xml:"anyAttributes,omitempty"`

	// Inside complex element...
	Choice   *XSD_Choice   `xml:"choice,omitempty"`
	All      *XSD_All      `xml:"all,omitempty"`
	Group    *XSD_Group    `xml:"group,omitempty"`
	Sequence *XSD_Sequence `xml:"sequence,omitempty"`
}

/**
 * Definition of the field to use for a uniqueness constraint.
 */
type XSD_Field struct {
	XMLName xml.Name `xml:"field"`

	/**
	 * W3C XML Schema's element ID.
	 */
	Id string `xml:"id,attr"`

	/**
	Relative XPath expression identifying the field(s) composing the key, key
	reference, or unique constraint.
	*/
	Xpath string `xml:"xpath,attr"`

	// Annotation for comment...
	Annotations []XSD_Annotation `xml:"annotation,omitempty"`
}

/**
 * Facet to define the number of fractional digits of a numerical datatype..
 */
type XSD_FractionDigits struct {
	XMLName xml.Name `xml:"fractionDigits"`

	/**
	 * W3C XML Schema's element ID.
	 */
	Id string `xml:"id,attr"`

	/**
	When set to true , the value of the facet cannot be modified during further
	restrictions.
	*/
	Fixed bool `xml:"fixed,attr"`

	/**
	Value of the facet.
	*/
	Value string `xml:"value,attr"`

	// Annotation for comment...
	Annotations []XSD_Annotation `xml:"annotation,omitempty"`
}

/**
 * Global elements group declaration that can be referenced within the same schema
 * by other schemas.
 */
type XSD_Group struct {
	XMLName xml.Name `xml:"group"`

	Name string `xml:"name,attr"`

	/**
	 * W3C XML Schema's element ID.
	 */
	Id string `xml:"id,attr"`

	/**
	 * Maximum number of occurrences. Note that this value is fixed to one.
	 */
	MaxOccurs string `xml:"maxOccurs,attr"`

	/**
	 * Minimum number of occurrences. Note that this value can be only zero or one.
	 */
	MinOccurs string `xml:"minOccurs,attr"`

	/**
	 * Qualified name of the group to include.
	 */
	Ref string `xml:"ref,attr"`

	// Annotation for comment...
	Annotations []XSD_Annotation `xml:"annotation,omitempty"`

	// The about anything...
	Any *XSD_Any `xml:"any,omitempty"`

	// One of...
	Choice *XSD_Choice `xml:"choice,omitempty"`

	// All of
	All *XSD_All `xml:"all,omitempty"`

	// Sequence of
	Sequence *XSD_Sequence `xml:"sequence,omitempty"`
}

/**
 * Import of a W3C XML Schema for another namespace.
 */
type XSD_Import struct {
	XMLName xml.Name `xml:"import"`

	/**
	 * W3C XML Schema's element ID.
	 */
	Id string `xml:"id,attr"`

	/**
	Namespace URI of the components to import. If this attribute is missing, the
	imported components are expected to have no namespace. When present, its value
	must be different than the target namespace of the importing schema.
	*/
	namespace string `xml:"namespace,attr"`

	/**
	Location of the schema to import. If this attribute is missing, the validator might
	expect to get the information from the application, or try to find it on the Internet.
	*/
	schemaLocation string `xml:"schemaLocation,attr"`

	// Annotation for comment...
	Annotations []XSD_Annotation `xml:"annotation,omitempty"`
}

/**
 * Inclusion of a W3C XML Schema for the same target namespace.
 */
type XSD_Include struct {
	XMLName xml.Name `xml:"include"`

	/**
	 * W3C XML Schema's element ID.
	 */
	Id string `xml:"id,attr"`

	/**
	Location of the schema to include.
	*/
	SchemaLocation string `xml:"schemaLocation,attr"`

	// Annotation for comment...
	Annotations []XSD_Annotation `xml:"annotation,omitempty"`
}

/**
 * Definition of a key.
 */
type XSD_Key struct {
	XMLName xml.Name `xml:"key"`

	/**
	 * W3C XML Schema's element ID.
	 */
	Id string `xml:"id,attr"`

	/**
	 * The name of the key.
	 */
	Name string `xml:"name,attr"`

	/**
	 * The selector
	 */
	Selector XSD_Selector `xml:"selector,attr"`

	/**
	 * The fields...
	 */
	Fields []XSD_Field `xml:"field,attr"`
}

/**
 * Definition of a key reference.
 */
type XSD_KeyRef struct {
	XMLName xml.Name `xml:"keyref"`

	/**
	 * W3C XML Schema's element ID.
	 */
	Id string `xml:"id,attr"`

	/**
	 * Name of the key reference.
	 */
	Name string `xml:"name,attr"`

	/**
	 * Name of the key or unique constraint referred by the key reference.
	 */
	Refer string `xml:"refer,attr"`

	/**
	 * The selector
	 */
	Selector XSD_Selector `xml:"selector,attr"`

	/**
	 * The fields...
	 */
	Fields []XSD_Field `xml:"field,attr"`
}

/**
 * Facet to define the length of a value.
 */
type XSD_Length struct {
	XMLName xml.Name `xml:"length"`

	/**
	 * W3C XML Schema's element ID.
	 */
	Id string `xml:"id,attr"`

	/**
	When set to true , the value of the facet cannot be modified during further
	restrictions.
	*/
	Fixed bool `xml:"fixed,attr"`

	/**
	Value of the facet.
	*/
	Value string `xml:"value,attr"`

	// Annotation for comment...
	Annotations []XSD_Annotation `xml:"annotation,omitempty"`
}

/**
 * Derivation by list
 */
type XSD_List struct {
	XMLName xml.Name `xml:"list"`

	/**
	 * W3C XML Schema's element ID.
	 */
	Id string `xml:"id,attr"`

	/**
	 * Reference to the item type when not defined inline by a xs:simpleType element.
	 */
	ItemType string `xml:"itemType,attr"`

	// Annotation for comment...
	Annotations []XSD_Annotation `xml:"annotation,omitempty"`

	SimpleType *XSD_SimpleType `xml:"simpleType,omitempty"`
}

/**
 * Facet to define a maximum (exclusive) value.
 */
type XSD_MaxExclusive struct {
	XMLName xml.Name `xml:"maxExclusive"`

	/**
	 * W3C XML Schema's element ID.
	 */
	Id string `xml:"id,attr"`

	/**
	When set to true , the value of the facet cannot be modified during further
	restrictions.
	*/
	Fixed bool `xml:"fixed,attr"`

	/**
	Value of the facet. Can be almost any type...
	*/
	Value string `xml:"value,attr"`

	// Annotation for comment...
	Annotations []XSD_Annotation `xml:"annotation,omitempty"`
}

/**
 * Facet to define a maximum (inclusive) value.
 */
type XSD_MaxInclusive struct {
	XMLName xml.Name `xml:"maxInclusive"`

	/**
	 * W3C XML Schema's element ID.
	 */
	Id string `xml:"id,attr"`

	/**
	When set to true , the value of the facet cannot be modified during further
	restrictions.
	*/
	Fixed bool `xml:"fixed,attr"`

	/**
	Value of the facet. Can be almost any type...
	*/
	Value string `xml:"value,attr"`

	// Annotation for comment...
	Annotations []XSD_Annotation `xml:"annotation,omitempty"`
}

/**
 * Facet to define a maximum length.
 */
type XSD_MaxLength struct {
	XMLName xml.Name `xml:"maxLength"`

	/**
	 * W3C XML Schema's element ID.
	 */
	Id string `xml:"id,attr"`

	/**
	When set to true , the value of the facet cannot be modified during further
	restrictions.
	*/
	Fixed bool `xml:"fixed,attr"`

	/**
	Value of the facet.
	*/
	Value string `xml:"value,attr"`

	// Annotation for comment...
	Annotations []XSD_Annotation `xml:"annotation,omitempty"`
}

/**
 * Facet to define a minimum (exclusive) value.
 */
type XSD_MinExclusive struct {
	XMLName xml.Name `xml:"minExclusive"`

	/**
	 * W3C XML Schema's element ID.
	 */
	Id string `xml:"id,attr"`

	/**
	When set to true , the value of the facet cannot be modified during further
	restrictions.
	*/
	Fixed bool `xml:"fixed,attr"`

	/**
	Value of the facet. Can be almost any type...
	*/
	Value string `xml:"value,attr"`

	// Annotation for comment...
	Annotations []XSD_Annotation `xml:"annotation,omitempty"`
}

/**
 * Facet to define a minimum (inclusive) value.
 */
type XSD_MinInclusive struct {
	XMLName xml.Name `xml:"minInclusive"`

	/**
	 * W3C XML Schema's element ID.
	 */
	Id string `xml:"id,attr"`

	/**
	When set to true , the value of the facet cannot be modified during further
	restrictions.
	*/
	Fixed bool `xml:"fixed,attr"`

	/**
	Value of the facet. Can be almost any type...
	*/
	Value string `xml:"value,attr"`

	// Annotation for comment...
	Annotations []XSD_Annotation `xml:"annotation,omitempty"`
}

/**
 * Facet to define a minimum length.
 */
type XSD_MinLength struct {
	XMLName xml.Name `xml:"minLength"`

	/**
	 * W3C XML Schema's element ID.
	 */
	Id string `xml:"id,attr"`

	/**
	When set to true , the value of the facet cannot be modified during further
	restrictions.
	*/
	Fixed bool `xml:"fixed,attr"`

	/**
	Value of the facet.
	*/
	Value string `xml:"value,attr"`

	// Annotation for comment...
	Annotations []XSD_Annotation `xml:"annotation,omitempty"`
}

/**
 * Declaration of a notation.
 */
type XSD_Notation struct {
	XMLName xml.Name `xml:"notation"`

	/**
	 * W3C XML Schema's element ID.
	 */
	Id string `xml:"id,attr"`

	/**
	 * The name of the key.
	 */
	Name string `xml:"name,attr"`

	/**
	Public identifier (usually its content type).
	*/
	Public string

	/**
	System identifier (typically the location of a resource that might be used
	to process the content type associated with the notation).
	*/
	System string `xml:"system,attr"`

	// Annotation for comment...
	Annotations []XSD_Annotation `xml:"annotation,omitempty"`
}

/**
 * Facet to define a regular expression pattern constraint.
 */
type XSD_Pattern struct {
	XMLName xml.Name `xml:"pattern"`

	/**
	 * W3C XML Schema's element ID.
	 */
	Id string `xml:"id,attr"`

	/**
	Value of the facet. Can be almost any type...
	*/
	Value string `xml:"value,attr"`

	// Annotation for comment...
	Annotations []XSD_Annotation `xml:"annotation,omitempty"`
}

/**
 * Inclusion of a W3C XML Schema for the same namespace with possible override.
 */
type XSD_Redefine struct {
	XMLName xml.Name `xml:"redefine"`

	/**
	 * W3C XML Schema's element ID.
	 */
	Id string `xml:"id,attr"`

	/**
	Location of the schema to redefine.
	*/
	SchemaLocation string `xml:"id,attr"`

	// Annotation for comment...
	Annotations []XSD_Annotation `xml:"annotation,omitempty"`
}

/**
 * Derivation of a simple datatype by restriction
 */
type XSD_Restriction struct {
	XMLName xml.Name `xml:"restriction"`

	/**
	 * W3C XML Schema's element ID.
	 */
	Id string `xml:"id,attr"`

	/**
	Qualified name of the base datatype when defined by reference.
	*/
	Base string `xml:"base,attr"`

	// Content

	// Annotation for comment...
	Annotations []XSD_Annotation `xml:"annotation,omitempty"`

	// Simple type
	SimpleType *XSD_SimpleType `xml:"simpleType,omitempty"`

	// inside Complex type context...
	Group    *XSD_Group    `xml:"group,omitempty"`
	All      *XSD_All      `xml:"all,omitempty"`
	Choice   *XSD_Choice   `xml:"choice,omitempty"`
	Sequence *XSD_Sequence `xml:"sequence,omitempty"`

	// Facets
	MaxExclusive   []XSD_MaxExclusive   `xml:"maxExclusive,omitempty"`
	MinExclusive   []XSD_MinExclusive   `xml:"minExclusive,omitempty"`
	MaxInclusive   []XSD_MaxInclusive   `xml:"maxInclusive,omitempty"`
	MinInclusive   []XSD_MinInclusive   `xml:"minInclusive,omitempty"`
	MaxLength      []XSD_MaxLength      `xml:"maxLength,omitempty"`
	MinLength      []XSD_MinLength      `xml:"minLength,omitempty"`
	Length         []XSD_Length         `xml:"length,omitempty"`
	FractionDigits []XSD_FractionDigits `xml:"fractionDigits,omitempty"`
	TotalDigits    []XSD_TotalDigits    `xml:"totalDigits,omitempty"`
	WhiteSpace     []XSD_WhiteSpace     `xml:"whiteSpace,omitempty"`
	Pattern        []XSD_Pattern        `xml:"pattern,omitempty"`
	Enumeration    []XSD_Enumeration    `xml:"enumeration"`

	Attributes      []XSD_Attribute      `xml:"attribute,omitempty"`
	AttributeGroups []XSD_AttributeGroup `xml:"attributeGroup,omitempty"`
	AnyAttributes   []XSD_AnyAttribute   `xml:"anyAttributes,omitempty"`
}

/**
 * Document element of a W3C XML Schema.
 */
type XSD_Schema struct {
	Schema xml.Name `xml:"schema"`

	/**
	 * W3C XML Schema's element ID.
	 */
	Id string `xml:"id,attr"`

	/**
	Default value for the form attributes of xs:attribute, determining whether
	attributes will be namespace-qualified by default.
	Possible value...
	"qualified" | "unqualified" ) : "unqualified"
	*/
	AttributeFormDefault string `xml:"attributeFormDefault,attr"`

	/**
	Default value of the block attribute of xs:element and xs:complexType.
	Possible value...
	"#all" | list of ( "extension" | "restriction" | "substitution"
	*/
	BlockDefault string `xml:"blockDefault,attr"`

	/**
	Default value for the form attributes of xs:element, determining whether
	attributes will be namespace-qualified by default
	Posible value
	"qualified" | "unqualified" ) : "unqualified"
	*/
	ElementFormDefault string `xml:"elementFormDefault,attr"`

	/**
	Default value of the final attribute of xs:element and xs:complexType.
	Possible value
	"#all" | list of ( "extension" | "restriction") : ""
	*/
	FinalDefault string `xml:"finalDefault,attr"`

	/**
	Namespace attached to this schema. All the qualified elements and attributes
	defined in this schema will belong to this namespace. This namespace will also
	be attached to all the global components.
	*/
	TargetNamespace string `xml:"targetNamespace,attr"`

	/**
	Version of the schema (for user convenience).
	*/
	Version string `xml:"version,attr"`

	/**
	Language of the schema.
	*/
	Language string `xml:"lang,attr"`

	// Content

	// Annotation for comment...
	Annotations []XSD_Annotation `xml:"annotation,omitempty"`

	Includes  []XSD_Include  `xml:"include,omitempty"`
	Imports   []XSD_Import   `xml:"import,omitempty"`
	Redefines []XSD_Redefine `xml:"redefine,omitempty"`

	Attributes      []XSD_Attribute      `xml:"attribute,omitempty"`
	AttributeGroups []XSD_AttributeGroup `xml:"attributeGroup,omitempty"`
	AnyAttributes   []XSD_AnyAttribute   `xml:"anyAttributes,omitempty"`

	Elements     []XSD_Element     `xml:"element,omitempty"`
	ComplexTypes []XSD_ComplexType `xml:"complexType,omitempty"`
	SimpleTypes  []XSD_SimpleType  `xml:"simpleType,omitempty"`
	Groups       []XSD_Group       `xml:"group,omitempty"`
}

/**
 * Definition of the the path selecting an element for a
 * uniqueness constraint.
 */
type XSD_Selector struct {
	XMLName xml.Name `xml:"selector"`

	/**
	 * W3C XML Schema's element ID.
	 */
	Id string `xml:"id,attr"`

	/**
	Relative XPath expression identifying the element on which the constraint
	applies.
	*/
	Xpath string `xml:"id,attr"`

	// Content

	// Annotation for comment...
	Annotations []XSD_Annotation `xml:"annotation,omitempty"`
}

/**
 * Compositor to define an ordered group of elements.
 */
type XSD_Sequence struct {
	XMLName xml.Name `xml:"sequence"`

	/**
	 * W3C XML Schema's element ID.
	 */
	Id string `xml:"id,attr"`

	/**
	 * Maximum number of occurrences. Note that this value is fixed to one.
	 */
	MaxOccurs string `xml:"maxOccurs,attr"`

	/**
	 * Minimum number of occurrences. Note that this value can be only zero or one.
	 */
	MinOccurs string `xml:"minOccurs,attr"`

	// Content

	// Annotation for comment...
	Annotations []XSD_Annotation `xml:"annotation,omitempty"`

	Elements  []XSD_Element  `xml:"element,omitempty"`
	Groups    []XSD_Group    `xml:"group,omitempty"`
	Choices   []XSD_Choice   `xml:"choice,omitempty"`
	Sequences []XSD_Sequence `xml:"sequence,omitempty"`
	Any       []XSD_Any      `xml:"any,omitempty"`
}

/**
 * Simple content model declaration.
 */
type XSD_SimpleContent struct {
	XMLName xml.Name `xml:"simpleContent"`

	/**
	 * W3C XML Schema's element ID.
	 */
	Id string `xml:"id,attr"`

	// Content
	// Annotation for comment...
	Annotations []XSD_Annotation `xml:"annotation,omitempty"`

	Restriction *XSD_Restriction `xml:"restriction,omitempty"`
	Extension   *XSD_Extension   `xml:"extension,omitempty"`
}

/**
 * Global simple type declaration that can be referenced within the same
 * schema by other schemas.
 */
type XSD_SimpleType struct {
	XMLName xml.Name `xml:"simpleType"`

	/**
	 * W3C XML Schema's element ID.
	 */
	Id string `xml:"id,attr"`

	/**
	When set, this attribute blocks any further derivations of this datatype
	(by list, union, derivation, or all).
	Possible value.
	"#all" | ( "list" | "union" | "restriction" )
	*/
	Final string `xml:"final,attr"`

	/**
	Unqualified name of this datatype.
	*/
	Name string `xml:"name,attr"`

	// Content
	// Annotation for comment...
	Annotations []XSD_Annotation `xml:"annotation,omitempty"`
	Restriction *XSD_Restriction `xml:"restriction,omitempty"`
	List        *XSD_List        `xml:"list,omitempty"`
	Union       *XSD_Union       `xml:"union,omitempty"`
}

/**
 * Facet to define the total number of digits of a numeric datatype.
 */
type XSD_TotalDigits struct {
	XMLName xml.Name `xml:"totalDigits"`

	/**
	 * W3C XML Schema's element ID.
	 */
	Id string `xml:"id,attr"`

	/**
	When set to true , the value of the facet cannot be modified during further
	restrictions.
	*/
	Fixed bool `xml:"fixed,attr"`

	/**
	Value of the facet.
	*/
	Value string `xml:"value,attr"`

	// Annotation for comment...
	Annotations []XSD_Annotation `xml:"annotation,omitempty"`
}

/**
 * Derivation of simple datatypes by union.
 */
type XSD_Union struct {
	XMLName xml.Name `xml:"union"`

	/**
	 * W3C XML Schema's element ID.
	 */
	Id string `xml:"id,attr"`

	/**
	List of member types (member types can also be embedded as xs:simpleType in
	the xs:union element).
	*/
	MemberTypes string `xml:"xml:anyURI"`

	Annotations []XSD_Annotation `xml:"annotation,omitempty"`
	SimpleType  []XSD_SimpleType `xml:"simpleType"`
}

/**
 * Definition of a uniqueness constraint.
 */
type XSD_Unique struct {
	XMLName xml.Name `xml:"unique"`

	/**
	 * W3C XML Schema's element ID.
	 */
	Id string `xml:"id,attr"`

	/**
	Name of the unique constraint.
	*/
	Name string `xml:"name,attr"`

	Annotations []XSD_Annotation `xml:"annotation,omitempty"`

	Selector XSD_Selector `xml:"selector,omitempty"`
	Fields   []XSD_Field  `xml:"field,omitempty"`
}

/**
 * Facet to define whitespace behavior..
 */
type XSD_WhiteSpace struct {
	XMLName xml.Name `xml:"whiteSpace"`

	/**
	When set to true , the value of the facet cannot be modified during further
	restrictions.
	*/
	Fixed bool `xml:"fixed,attr"`

	/**
	Value of the facet .
	Possible values
	"preserve" | "replace" | "collapse"
	*/
	Value string `xml:"value,attr"`

	// Annotation for comment...
	Annotations []XSD_Annotation `xml:"annotation,omitempty"`

	/**
	 * W3C XML Schema's element ID.
	 */
	Id string `xml:"id,attr"`
}

/**
 * Global element definition that can be referenced within the same schema by
 * other schemas.
 */
type XSD_Element struct {
	/**
	 * Place to keep the information about the schema that host that element.
	 */
	SchemaId string

	XMLName xml.Name `xml:"element"`

	/**
	 * W3C XML Schema's element ID.
	 */
	Id string `xml:"id,attr"`

	/**
	Controls whether the element may be used directly in instance documents.
	When set to true, the element may still be used to define content models,
	but it must be substituted through a substitution group in the instance
	document.
	*/
	IsAbstract string `xml:"abstract,attr"`

	/**
	Controls whether the element can be subject to a type or substitution group
	substitution. #all blocks any substitution, substitution blocks any substitution
	through substitution groups, and extension and restriction block any substitution
	(both through xsi:type and substitution groups) by elements or types, derived
	respectively by extension or restriction from the type of the element.
	Its default value is defined by the blockDefault attribute of the parent
	xs:schema.
	*/
	Block string `xml:"block,attr"`

	/**
	Local name of the element (without namespace prefix).
	*/
	Name string `xml:"name,attr"`

	/**
	A simple content element may be fixed to a specific value using this attribute.
	This value is also used as a default value, and if the element is empty, it is
	supplied to the application. The fixed and default attributes are mutually
	exclusive.
	*/
	Fixed string `xml:"fixed, attr"`

	/**
	Controls whether the element can be used as the head of a substitution group
	for elements whose types are derived by extension or restriction from the type
	of the element. Its default value is defined by the finalDefault attribute of
	the parent
	*/
	Final string `xml:"final, attr"`

	/**
	Default value of the element. Defined in an attribute, element default values
	must be simple contents. Also note that default values apply only to elements
	that are present in the document and empty. The fixed and default attributes
	are mutually exclusive.
	*/
	Default string `xml:"default, attr"`

	/**
	Qualified name of a simple or complex type (must be omitted when a simple
	or complex type definition is embedded).
	*/
	Type string `xml:"type,attr"`

	/**
	Reference to a global element definition (mutually exclusive with the name,
	block, and type attributes and any embedded type definition.
	*/
	Ref string `xml:"ref,attr"`

	/**
	Qualified name of the head of the substitution group to which this element
	belongs.
	*/
	SubstitutionGroup string `xml:"substitutionGroup,attr"`

	/**
	When this attribute is set to true, the element can be declared as nil using
	an xsi:nil attribute in the instance documents.
	*/
	Nillable bool `xml:"nillable,attr"`

	/**
	Maximum number of occurrences of the element. Can take only the values 0 or
	1 within a xs:all compositor.
	*/
	MinOccurs string `xml:"minOccurs,attr"`

	/**
	Minimum number of occurrences of the element. Can take only the values 0 or
	1 within a xs:all compositor.
	*/
	MaxOccurs string `xml:"maxOccurs,attr"`

	// Content

	// Comment...
	Annotations []XSD_Annotation `xml:"annotation,omitempty"`

	// One of complexType or Simple type...
	SimpleType  *XSD_SimpleType  `xml:"simpleType,omitempty"`
	ComplexType *XSD_ComplexType `xml:"complexType,omitempty"`

	// One of unique, key or keyref
	Unique []XSD_Unique `xml:"unique,omitempty"`
	Key    []XSD_Key    `xml:"key,omitempty"`
	KeyRef []XSD_KeyRef `xml:"keyref,omitempty"`
}

/**
 * xs:anyAttribute is a wildcard that allows the insertion of any attribute
 * belonging to a list of namespaces. This particle must be used wherever an
 * attribute local declaration of reference can be used (i.e., within complexType
 * or attributeGroup definitions).
 */
type XSD_AnyAttribute struct {
	XMLName xml.Name `xml:"anyAttribute"`
	/**
	 * W3C XML Schema's element ID.
	 */
	Id string `xml:"id,attr"`

	Name string `xml:"name,attr"`

	/**
	 * Permitted namespaces.
	 */
	Namespace string `xml:"namespace,attr"`

	/**
	 * Type of validation required on the elements permitted for this wildcard.
	 */
	ProcessContents string `xml:"processContents,attr"`

	// Comment...
	Annotations []XSD_Annotation `xml:"annotation,omitempty"`
}

/**
 * Global attribute definition that can be referenced within the same schema by
 * other schemas.
 */
type XSD_Attribute struct {
	XMLName xml.Name `xml:"attribute"`
	/**
	 * W3C XML Schema's element ID.
	 */
	Id string `xml:"id,attr"`

	/**
	 * Local name of the attribute (without namespace prefix).
	 */
	Name string `xml:"name,attr"`

	/**
	 * Qualified name of a simple type of the attribute (must be omitted when a
	 * simple type definition is embedded).
	 */
	Type string `xml:"type,attr"`

	/**
		 * Default value. When specified, an attribute is added by the schema processor
		 * (if it is missing from the instance document) and it is given this value. The
	     * default and fixed attributes are mutually exclusive.
	*/
	Default string `xml:"default,attr"`

	/**
	 * When specified, the value of the attribute is fixed and must be equal to
	 * the value specified here. The default and fixed attributes are mutually exclusive.
	 */
	Fixed string `xml:"fixed,attr"`

	/**
	Possible usage of the attribute. Marking an attribute "prohibited" is useful
	to exclude attributes during derivations by restriction.
	possible value are: "prohibited" | "optional" | "required"
	by default is "optional"
	*/
	Use string `xml:"use,attr"`

	/**
	Specifies if the attribute is qualified (i.e., must have a namespace prefix
	in the instance document) or not. The default value for this attribute is
	specified by the attributeFormDefault attribute of the xs:schema document
	element—local definition only.
	possible value are: "qualified" | "unqualified"
	*/
	Form string `xml:"form,attr"`

	/**
	 * Qualified name of a globally defined attribute—reference only.
	 */
	Ref string `xml:"ref,attr"`

	/**
	 * The attribute content, only simple type allow.
	 */
	SimpleType *XSD_SimpleType `xml:"simpleType"`

	/**
	 * Comment about the attribute.
	 */
	Annotations []XSD_Annotation `xml:"annotation,omitempty"`
}

/**
 * Global attributes group declaration that can be referenced within the same
 * schema by other schemas.
 */
type XSD_AttributeGroup struct {
	XMLName xml.Name `xml:"attributeGroup"`
	/**
	 * W3C XML Schema's element ID.
	 */
	Id string `xml:"id,attr"`

	/**
	 * Local name of the attribute (without namespace prefix).
	 */
	Name string `xml:"name,attr"`

	/**
	 * Qualified name of the attribute group to reference.
	 */
	Ref string `xml:"ref,attr"`

	// Content...
	Annotations     []XSD_Annotation      `xml:"annotation,omitempty"`
	Attributes      []*XSD_Attribute      `xml:"attribute"`
	AnyAttributes   []*XSD_AnyAttribute   `xml:"anyAttribute"`
	AttributeGroups []*XSD_AttributeGroup `xml:"attributeGroup"`
}

/**
 * Compositor to define group of mutually exclusive elements or compositors.
 */
type XSD_Choice struct {
	XMLName xml.Name `xml:"choice"`

	/**
	 * W3C XML Schema's element ID.
	 */
	Id string `xml:"id,attr"`

	/**
	 * Maximum number of occurrences. Note that this value is fixed to one.
	 */
	MaxOccurs string `xml:"maxOccurs,attr"`

	/**
	 * Minimum number of occurrences. Note that this value can be only zero or one.
	 */
	MinOccurs string `xml:"minOccurs,attr"`

	// Content
	Elements  []XSD_Element  `xml:"element,omitempty"`
	Sequences []XSD_Sequence `xml:"sequence,omitempty"`
	Groups    []XSD_Group    `xml:"group,omitempty"`
	Choices   []XSD_Choice   `xml:"choice,omitempty"`
	Anys      []XSD_Any      `xml:"any,omitempty"`
}

/**
 * Facet to restrict a datatype to a finite set of values.
 */
type XSD_Enumeration struct {
	XMLName xml.Name `xml:"enumeration"`

	/**
	 * W3C XML Schema's element ID.
	 */
	Id string `xml:"id,attr"`

	/**
	 * Value to be added to the list of possible values for this datatype.
	 */
	Value string `xml:"value,attr"`

	// Content
	Annotations []XSD_Annotation `xml:"annotation,omitempty"`
}

/**
 * xs:any is a wildcard that allows the insertion of any element belonging to a
 * list of namespaces. This particle can be used like xs:element within choices
 * (xs:choice) and sequences (xs:sequence), and the number of occurrences of the
 * elements that are allowed can be controlled by its minOccurs and maxOccurs
 * attributes.
 */
type XSD_Any struct {
	XMLName xml.Name `xml:"any"`

	/**
	 * W3C XML Schema's element ID.
	 */
	Id string `xml:"id,attr"`

	Name string `xml:"name,attr"`

	/**
	 * Maximum number of elements permitted for this wildcard.
	 */
	MaxOccurs string `xml:"maxOccurs,attr"`

	/**
	 * Minimum number of elements permitted for this wildcard.
	 */
	MinOccurs string `xml:"minOccurs,attr"`

	/**
	 * Permitted namespaces.
	 */
	Namespace string `xml:"namespace,attr"`

	/**
	 * Type of validation required on the elements permitted for this wildcard.
	 */
	ProcessContents string `xml:"processContents,attr"`

	// Comment...
	Annotations []XSD_Annotation `xml:"annotation,omitempty"`
}

// Helper functions.
/**
 * Dertermine if the value is a base type.
 */
func IsXsBaseType(fieldType string) bool {
	return IsXsId(fieldType) || IsXsRef(fieldType) || IsXsInt(fieldType) || IsXsString(fieldType) || IsXsBinary(fieldType) || IsXsNumeric(fieldType) || IsXsBoolean(fieldType) || IsXsDate(fieldType) || IsXsTime(fieldType) || IsXsMoney(fieldType)
}

/**
 * Helper function use to dertermine if a XS type must be considere integer.
 */
func IsXsInt(fieldType string) bool {
	if strings.HasSuffix(fieldType, "int") || strings.HasSuffix(fieldType, "byte") || strings.HasSuffix(fieldType, "int") || strings.HasSuffix(fieldType, "integer") || strings.HasSuffix(fieldType, "short") || strings.HasSuffix(fieldType, "unsignedInt") || strings.HasSuffix(fieldType, "unsignedBtype") || strings.HasSuffix(fieldType, "unsignedShort") || strings.HasSuffix(fieldType, "unsignedLong") || strings.HasSuffix(fieldType, "negativeInteger") || strings.HasSuffix(fieldType, "nonNegativeInteger") || strings.HasSuffix(fieldType, "nonPositiveInteger") || strings.HasSuffix(fieldType, "positiveInteger") || strings.HasSuffix(fieldType, "tinyint") || strings.HasSuffix(fieldType, "smallint") || strings.HasSuffix(fieldType, "bigint") {
		return true
	}
	return false
}

/**
 * Helper function use to dertermine if a XS type must be considere String.
 */
func IsXsString(fieldType string) bool {
	if strings.HasSuffix(fieldType, "string") || strings.HasSuffix(fieldType, "Name") || strings.HasSuffix(fieldType, "QName") || strings.HasSuffix(fieldType, "NMTOKEN") || strings.HasSuffix(fieldType, "gDay") || strings.HasSuffix(fieldType, "gMonth") || strings.HasSuffix(fieldType, "gMonthDay") || strings.HasSuffix(fieldType, "gYear") || strings.HasSuffix(fieldType, "gYearMonth") || strings.HasSuffix(fieldType, "token") || strings.HasSuffix(fieldType, "normalizedString") || strings.HasSuffix(fieldType, "hexBinary") || strings.HasSuffix(fieldType, "language") || strings.HasSuffix(fieldType, "NMTOKENS") || strings.HasSuffix(fieldType, "NOTATION") || strings.HasSuffix(fieldType, "char") || strings.HasSuffix(fieldType, "nchar") || strings.HasSuffix(fieldType, "varchar") || strings.HasSuffix(fieldType, "nvarchar") || strings.HasSuffix(fieldType, "text") || strings.HasSuffix(fieldType, "ntext") {
		return true
	}
	return false
}

/**
 * Helper function use to dertermine if a XS type must be considere binary value.
 */
func IsXsBinary(fieldType string) bool {
	if strings.HasSuffix(fieldType, "base64Binary") || strings.HasSuffix(fieldType, "varbinary") || strings.HasSuffix(fieldType, "binary") || strings.HasSuffix(fieldType, "image") || strings.HasSuffix(fieldType, "timestamp") {
		return true
	}
	return false
}

/**
 * Helper function use to dertermine if a XS type must be considere numeric value.
 */
func IsXsNumeric(fieldType string) bool {
	if strings.HasSuffix(fieldType, "double") || strings.HasSuffix(fieldType, "decimal") || strings.HasSuffix(fieldType, "float") || strings.HasSuffix(fieldType, "numeric") || strings.HasSuffix(fieldType, "real") {
		return true
	}
	return false
}

/**
 * Helper function use to dertermine if a XS type must be considere boolean value.
 */
func IsXsBoolean(fieldType string) bool {
	if strings.HasSuffix(fieldType, "boolean") || strings.HasSuffix(fieldType, "bit") {
		return true
	}
	return false
}

/**
 * Helper function use to dertermine if a XS type must be considere date value.
 */
func IsXsDate(fieldType string) bool {
	if strings.HasSuffix(fieldType, "date") || strings.HasSuffix(fieldType, "datetime") || strings.HasSuffix(fieldType, "datetime2") || strings.HasSuffix(fieldType, "smalldatetime") || strings.HasSuffix(fieldType, "datetimeoffset") {
		return true
	}
	return false
}

/**
 * Helper function use to dertermine if a XS type must be considere time value.
 */
func IsXsTime(fieldType string) bool {
	if strings.HasSuffix(fieldType, "time") || strings.HasSuffix(fieldType, "timestampNumeric") || strings.HasSuffix(fieldType, "timestamp") {
		return true
	}
	return false
}

/**
 * Helper function use to dertermine if a XS type must be considere money value.
 */
func IsXsMoney(fieldType string) bool {
	if strings.HasSuffix(fieldType, "money") || strings.HasSuffix(fieldType, "smallmoney") {
		return true
	}
	return false
}

/**
 * Helper function use to dertermine if a XS type must be considere id value.
 */
func IsXsId(fieldType string) bool {
	if strings.HasSuffix(fieldType, "ID") || strings.HasSuffix(fieldType, "NCName") || strings.HasSuffix(fieldType, "uniqueidentifier") {
		return true
	}
	return false
}

/**
 * Helper function use to dertermine if a XS type must be considere id value.
 */
func IsXsRef(fieldType string) bool {
	if strings.HasSuffix(fieldType, "anyURI") || strings.HasSuffix(fieldType, "IDREF") {
		return true
	}
	return false
}
