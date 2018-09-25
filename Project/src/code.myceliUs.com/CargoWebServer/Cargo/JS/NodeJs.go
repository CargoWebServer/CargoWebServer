/**
 * That package contain NodeJs implementation for;
 * fs The file system module.
 */
package JS

import (
	"log"
)

func (this *JsRuntimeManager) initNodeJs() {
	/**
	 * Node.js module/exports functionality.
	 */
	this.appendFunction("require_", func(identifier string, sessionId string) string {
		// resolve dependencie and return the exports.
		exports, err := GetJsRuntimeManager().getExports(identifier, sessionId)
		if err != nil {
			log.Println("---> error found in require function ", err)
		}

		// Get the export variable export uuid
		uuid, err := exports.Get("exports_id__")
		if err == nil {
			return uuid.Val.(string)
		}

		log.Println("---> error found in require function ", err)

		return ""
	})

	// The file system module.
	this.initNodeJsFs()
}
