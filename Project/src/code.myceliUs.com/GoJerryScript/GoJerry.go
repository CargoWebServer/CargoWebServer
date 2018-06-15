package GoJerryScript

import (
	"log"
)

// Various type conversion functions.

// The Uint8 Type represent a 8 bit char.
type Uint8 struct {
	val uintptr
}

func (self Uint8) Swigcptr() uintptr {
	log.Println("----> Swigcptr call!", self.val)
	return self.val
}

/**
 * The JerryScript JS engine.
 */
type GoJerry struct {
}
