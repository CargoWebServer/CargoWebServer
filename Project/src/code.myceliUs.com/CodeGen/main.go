// CodeGen project main.go
// here is the complete line to append
// -i -tags "CargoEntities Config BPMS BPMN20 BPMNDI DI DC"
package main

import (
	"log"
	"os"
	"path/filepath"
)

func main() {
	// Get the file contain inside 'input' directory.
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	// Generate the module inside input...
	generate(dir+"/input", "Config", "configurations")
	generate(dir+"/input", "CargoEntities", "entities")

	// Now the runtime.
	/*initMaps()

	// set the output to
	loadXSD("input/BPMS.xsd")

	// Now the cmof shcmea...
	loadCMOF("input/BPMS.cmof")

	generateCode("BPMS")

	// xml roots element...
	generateGoXmlFactory("runtimes", "Entities", outputPath, "BPMS_")

	generateEntity("BPMS")

	// Bpmn stuff
	initMaps()

	// set the output to
	loadXSD("input/DI.xsd")
	loadXSD("input/DC.xsd")
	loadXSD("input/BPMN20.xsd")
	loadXSD("input/BPMNDI.xsd")
	loadXSD("input/Semantic.xsd")

	// Now the cmof shcmea...
	loadCMOF("input/DC.cmof")
	loadCMOF("input/DI.cmof")
	loadCMOF("input/BPMNDI.cmof")
	loadCMOF("input/BPMN20.cmof")

	generateCode("DC")
	generateCode("DI")
	generateCode("BPMNDI")
	generateCode("BPMN20")

	// xml roots element...
	generateGoXmlFactory("definitions", "Entities", outputPath, "BPMN_")

	// The Entity...
	generateEntity("DC")
	generateEntity("DI")
	generateEntity("BPMNDI")
	generateEntity("BPMN20")*/

}
