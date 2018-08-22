/**
 * That package contain NodeJs implementation for;
 * fs The file system module.
 */
package JS

import (
	"log"

	"code.myceliUs.com/GoJerryScript"
)

func (this *JsRuntimeManager) initNodeJs() {
	/**
	 * Node.js module/exports functionality.
	 */
	this.appendFunction("require_", func(identifier string, sessionId string) GoJerryScript.Object {
		// resolve dependencie and return the exports.
		exports, err := GetJsRuntimeManager().getExports(identifier, sessionId)
		if err != nil {
			log.Println("---> error found in require function ", err)
		}
		return exports
	})

	// The file system module.
	this.initNodeJsFs()
}
