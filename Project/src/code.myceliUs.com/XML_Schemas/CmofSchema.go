package XML_Schemas

import (
	"encoding/xml"
)

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
