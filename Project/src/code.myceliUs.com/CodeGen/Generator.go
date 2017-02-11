package main

import (
	"encoding/xml"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"code.myceliUs.com/Utility"
	"code.myceliUs.com/XML_Schemas"
)

/**
 * Read the content of a CMOF file.
 */
func ReadCMOF(reader io.Reader) (*XML_Schemas.CMOF_Document, error) {
	var doc *XML_Schemas.CMOF_Document
	doc = new(XML_Schemas.CMOF_Document)
	if err := xml.NewDecoder(reader).Decode(doc); err != nil {
		return nil, err
	}

	return doc, nil
}

/**
 * Load the content into various maps (CMOF.go)
 */
func loadCMOF(inputPath string) {

	cmofFilePath, err := filepath.Abs(inputPath)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	// Open the straps.xml file
	file, err := os.Open(cmofFilePath)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	// Read the straps file
	cmofDocument, err := ReadCMOF(file)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	loadDocument(cmofDocument)
}

/**
 * Read the content of XSD file.
 */
func ReadXSD(reader io.Reader) (*XML_Schemas.XSD_Schema, error) {
	var doc *XML_Schemas.XSD_Schema
	doc = new(XML_Schemas.XSD_Schema)
	if err := xml.NewDecoder(reader).Decode(doc); err != nil {
		return nil, err
	}

	return doc, nil
}

/**
 * Load the content of the file into various maps (XSD.go)
 */
func loadXSD(inputPath string) {

	xsdFilePath, err := filepath.Abs(inputPath)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	// Open the straps.xml file
	file, err := os.Open(xsdFilePath)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	// Read the straps file
	xsdDocument, err := ReadXSD(file)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	loadSchema(xsdDocument)
}

var (
	outputPath string
)

func WriteClassFile(outputPath, packName string, className string, classCode string) {
	goPath := os.Getenv("GOPATH")

	path := goPath + "/src/" + outputPath + packName
	path = strings.Replace(path, "/", "\\", -1)
	if !strings.HasSuffix(path, "\\") {
		path += "\\"
	}
	if !Utility.Exists(path) {
		os.MkdirAll(path, 0777)
	}

	log.Println("Create go file for class: ", path+className+".go")

	if !Utility.Exists(path + className + ".go") {
		err := ioutil.WriteFile(path+className+".go", []byte(classCode), 0x777)
		if err != nil {
			panic(err)
		}
	}
}

func generateCode(packName string) {

	// The interface...
	generateGoInterfacesCode(packName)

	// The class
	generateGoClassCode(packName)

	// The enumeration
	generateGoEnumerationCode(packName)

}

func initMaps() {

	// CMOF maps
	classesMap = make(map[string]*XML_Schemas.CMOF_OwnedMember)
	enumMap = make(map[string]*XML_Schemas.CMOF_OwnedMember)
	packagesMembers = make(map[string][]string)
	membersPackage = make(map[string]string)
	superClassMembers = make(map[string]map[string]*XML_Schemas.CMOF_OwnedMember)
	associationsSrcMap = make(map[string]*XML_Schemas.CMOF_OwnedMember)
	associationsDstMap = make(map[string][]string)
	cmofPrimitiveTypesMap = make(map[string]string)
	superClassesLst = make([]string, 0)
	abstractClassLst = make([]string, 0)
	compositionEntityCreatorMap = make(map[string][]string)

	// XSD MAPS
	elementsNameMap = make(map[string]*XML_Schemas.XSD_Element)
	elementsTypeMap = make(map[string]*XML_Schemas.XSD_Element)
	complexTypesMap = make(map[string]*XML_Schemas.XSD_ComplexType)
	simpleTypesMap = make(map[string]*XML_Schemas.XSD_SimpleType)
	xsdPrimitiveTypesMap = make(map[string]string)
	substitutionGroupMap = make(map[string][]string)
	aliasElementType = make(map[string]string)

	// Factory
	factoryImports = make([]string, 0)
	// That map will contain initialisation function for a given data type
	initFunctions = make(map[string]string)
	serializationFunctions = make(map[string]string)
}

/**
 * Genrate the code for the files in the given path.
 */
func generate(inputPath string, packName string, xmlRootElementName string) {
	// Clear the maps.
	initMaps()

	// Set the output...
	outputPath = "code.myceliUs.com/CargoWebServer/Cargo/Entities/"

	path := inputPath + "\\" + packName
	path = strings.Replace(path, "/", "\\", -1)

	loadCMOF(path + ".cmof")
	loadXSD(path + ".xsd")

	// generate the code
	generateCode(packName)

	// generate the factory if there a list of entities.
	if len(xmlRootElementName) > 0 {
		generateGoXmlFactory(xmlRootElementName, "Entities", outputPath, packName)
	}
	// generate the entity.
	generateEntity(packName)

}
