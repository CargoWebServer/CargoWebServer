package jsoniq

import (
	"log"
	"testing"
)

/**
 * Test getting the version.
 */
func TestQuery(t *testing.T) {
	log.Println("----> Test query 0")
	interpreter := NewInterpreter()

	q0 := `for $p in collection("persons")
		   where $p.age gt 20
		   let $home := $p.phoneNumber[][$$.type eq "home"].number
		   group by $area := substring-before($home, " ")
		   return 
		   {
		     "area code" : $area,
		     "count" : fct($p)
		   }
		`
	callback0 := func(values ...interface{}) error {
		log.Println("---> ", values)
		return nil
	}
	interpreter.Run(q0, callback0)
}
