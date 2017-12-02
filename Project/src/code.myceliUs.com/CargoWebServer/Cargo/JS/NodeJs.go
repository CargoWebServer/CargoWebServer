/**
 * That package contain NodeJs implementation for;
 * fs The file system module.
 */
package JS

import (
	"log"

	"github.com/robertkrimen/otto"
)

func (this *JsRuntimeManager) initNodeJs() {
	/**
	 * Node.js module/exports functionality.
	 */
	this.appendFunction("require", func(identifier string) *otto.Object {
		// resolve dependencie and return the exports.
		exports, err := GetJsRuntimeManager().getExports(identifier)
		if err != nil {
			log.Println("---> error found in require function ", err)
		}

		return exports
	})

	// The file system module.
	this.initNodeJsFs()
}
