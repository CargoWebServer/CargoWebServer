//+build DI
package DI

import (
	"code.myceliUs.com/CargoWebServer/Cargo/BPMS/DC"
)

type Shape interface {
	/** Method of Shape **/

	/** UUID **/
	GetUUID() string

	/** Bounds **/
	GetBounds() *DC.Bounds
}
