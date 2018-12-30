/**
 * That package contain NodeJs implementation for;
 * fs The file system module.
 */
package JS

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"strings"
	utf16 "unicode/utf16"

	//utf8 "unicode/utf8"

	"code.myceliUs.com/Utility"
)

func (this *JsRuntimeManager) initNodeJsFs() {

	/**
	 * Returns true if the file exists, false otherwise.
	 */
	this.appendFunction("fs.existsSync", func(path string) bool {
		return false
	})

	/**
	 * Returns true if the file exists, false otherwise.
	 */
	this.appendFunction("fs.readFileSync", func(path string, encoding string) []byte {
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
				// text = string(b) // convert content to a 'string'
				return b
			}

		}

		return nil
	})

	/**
	 * Returns true if the file exists, false otherwise.
	 */
	this.appendFunction("fs.copyFileSync", func(src string, dst string) {
		log.Println("69 ---> copy file sync ", src, dst)
		// I will create the directory before and the file after.
		if strings.Index(dst, "\\") != -1 {
			os.MkdirAll(dst[0:strings.LastIndex(dst, "\\")], 0777)
		} else if strings.Index(dst, "/") != -1 {
			os.MkdirAll(dst[0:strings.LastIndex(dst, "/")], 0777)
		}
		err := Utility.CopyFile(src, dst)
		if err != nil {
			log.Println("---> ", err)
			return
		}
	})

	/**
	 * Return the stats of a file for a give path.
	 */
	this.appendFunction("fs.fstatSync", func(path string) map[string]interface{} {
		stats, err := os.Stat(path)
		if err != nil {
			log.Println(err)
			return nil
		}

		values := make(map[string]interface{}, 0)
		values["Mode"] = stats.Mode()
		values["Time"] = stats.ModTime().Unix()
		values["Size"] = stats.Size()
		values["Sys"] = stats.Sys()
		values["Name"] = stats.Name()

		return values
	})

}
