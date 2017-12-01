/**
 * That package contain NodeJs implementation for;
 * fs The file system module.
 */
package Server

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	utf16 "unicode/utf16"
	//utf8 "unicode/utf8"

	"code.myceliUs.com/CargoWebServer/Cargo/JS"
	"github.com/robertkrimen/otto"
)

func initNodeJs() {
	/**
	 * Node.js module/exports functionality.
	 */
	JS.GetJsRuntimeManager().AppendFunction("require", func(identifier string) *otto.Object {
		// resolve dependencie et return the exports.
		exports, err := JS.GetJsRuntimeManager().GetExports(identifier)
		if err != nil {
			log.Println("---> error found in require function ", err)
		}

		return exports
	})

	// The file system module.
	initNodeJsFs()
}

func initNodeJsFs() {

	/**
	 * Returns true if the file exists, false otherwise.
	 */
	JS.GetJsRuntimeManager().AppendFunction("fs.existsSync", func(path string) bool {
		return false
	})

	/**
	 * Returns true if the file exists, false otherwise.
	 */
	JS.GetJsRuntimeManager().AppendFunction("fs.readFileSync", func(path string, encoding string) string {
		file, err := os.Open(path)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		fscanner := bufio.NewScanner(file)
		var text string
		for fscanner.Scan() {
			arrbytes := fscanner.Bytes()
			if encoding == "utf16" {
				slice := make([]uint16, len(arrbytes))
				for i := range arrbytes {
					slice[i] = uint16(arrbytes[i])
				}
				newarr := utf16.Decode(slice)
				text += string(newarr)
			} else if encoding == "utf8" {
				//text += utf8.FullRune(arrbytes)
			} else {
				b, err := ioutil.ReadFile(path) // just pass the file name
				if err != nil {
					log.Println(err)
				}
				text = string(b) // convert content to a 'string'

			}

		}

		return text
	})
}
