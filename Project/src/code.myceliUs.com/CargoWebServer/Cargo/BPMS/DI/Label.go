//+build DI
package DI

import (
	"code.myceliUs.com/CargoWebServer/Cargo/BPMS/DC"
)

type Label interface {
	/** Method of Label **/

	/** UUID **/
	GetUUID() string

	/** Bounds **/
	GetBounds() *DC.Bounds
}
