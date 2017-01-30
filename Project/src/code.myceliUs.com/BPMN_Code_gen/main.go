// BPMN_Code_gen project main.go
package main

import (
	"encoding/xml"
	"io"
	"log"
	"os"
	"path/filepath"
)

/**
 * The read function...
 */
func ReadCMOF(reader io.Reader) (*CMOF_Document, error) {
	var doc *CMOF_Document
	doc = new(CMOF_Document)
	if err := xml.NewDecoder(reader).Decode(doc); err != nil {
		return nil, err
	}

	return doc, nil
}

func load_CMOF_File(inputPath string) {

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

func ReadXSD(reader io.Reader) (*XSD_Schema, error) {
	var doc *XSD_Schema
	doc = new(XSD_Schema)
	if err := xml.NewDecoder(reader).Decode(doc); err != nil {
		return nil, err
	}

	return doc, nil
}

func load_XSD_File(inputPath string) {

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

var outputPath = ""

func clearMaps() {

	// CMOF maps
	classesMap = make(map[string]*CMOF_OwnedMember)
	enumMap = make(map[string]*CMOF_OwnedMember)
	packagesMembers = make(map[string][]string)
	membersPackage = make(map[string]string)
	superClassMembers = make(map[string]map[string]*CMOF_OwnedMember)
	associationsSrcMap = make(map[string]*CMOF_OwnedMember)
	associationsDstMap = make(map[string][]string)
	primitiveMapType = make(map[string]string)
	superClassesLst = make([]string, 0)
	abstractClassLst = make([]string, 0)
	compositionEntityCreatorMap = make(map[string][]string)

	// XSD MAPS
	elementsNameMap = make(map[string]*XSD_Element)
	elementsTypeMap = make(map[string]*XSD_Element)
	complexTypesMap = make(map[string]*XSD_ComplexType)
	simpleTypesMap = make(map[string]*XSD_SimpleType)
	xsdPrimitiveTypesMap = make(map[string]string)
	substitutionGroupMap = make(map[string][]string)
}
func main() {

	outputPath = "code.myceliUs.com/CargoWebServer/Cargo/BPMS/"

	//	//Load the xsd schema...
	//	load_XSD_File("input/DI.xsd")
	//	load_XSD_File("input/DC.xsd")
	//	load_XSD_File("input/BPMN20.xsd")
	//	load_XSD_File("input/BPMNDI.xsd")
	//	load_XSD_File("input/Semantic.xsd")

	//	// Now the cmof shcmea...
	//	load_CMOF_File("input/DC.cmof")
	//	load_CMOF_File("input/DI.cmof")
	//	load_CMOF_File("input/BPMNDI.cmof")
	//	load_CMOF_File("input/BPMN20.cmof")

	//	generateCode("DC")
	//	generateCode("DI")
	//	generateCode("BPMNDI")
	//	generateCode("BPMN20")

	//	// xml roots element...
	//	generateFactory("definitions", "BPMS", "code.myceliUs.com/CargoWebServer/Cargo/BPMS", "BPMS")

	//	// The Entity...
	//	generateEntity("DC")
	//	generateEntity("DI")
	//	generateEntity("BPMNDI")
	//	generateEntity("BPMN20")

	////////////////////////////////////////////////////////////////////////////////
	//	clearMaps()
	//	load_XSD_File("input/Runtime.xsd")
	//	load_CMOF_File("input/Runtime.cmof")
	//	generateCode("BPMS_Runtime")
	//	generateFactory("runtimes", "BPMS", "code.myceliUs.com/CargoWebServer/Cargo/BPMS", "Runtime")

	//	// The Entity...
	//	generateEntity("BPMS_Runtime")

	//////////////////////////////////////////////////////////////////////////////////
	//	clearMaps()
	//	outputPath = "code.myceliUs.com/CargoWebServer/Cargo/Persistence/"

	//	load_XSD_File("input/CargoEntities.xsd")
	//	load_CMOF_File("input/CargoEntities.cmof")

	//	generateCode("CargoEntities")
	//	generateFactory("entities", "Persistence", "code.myceliUs.com/CargoWebServer/Cargo/Persistence", "Persistence")
	//	generateEntity("CargoEntities")

	//////////////////////////////////////////////////////////////////////////////////
	outputPath = "code.myceliUs.com/CargoWebServer/Cargo/Config/"
	// Configuration entities...
	load_XSD_File("input/Config.xsd")
	load_CMOF_File("input/Config.cmof")
	generateCode("CargoConfig")
	generateFactory("configurations", "Config", "code.myceliUs.com/CargoWebServer/Cargo/Config", "Config")

	//The Entity...
	generateEntity("CargoConfig")

}
