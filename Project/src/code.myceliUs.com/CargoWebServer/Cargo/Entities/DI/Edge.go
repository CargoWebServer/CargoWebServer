// +build DI

package DI

import(
"code.myceliUs.com/CargoWebServer/Cargo/Entities/DC"
)

type Edge interface{
	/** Method of Edge **/

	/** UUID **/
	GetUUID() string

	/** Source **/
	GetSource() DiagramElement
	SetSource(interface{}) 

	/** Target **/
	GetTarget() DiagramElement
	SetTarget(interface{}) 

	/** Waypoint **/
	GetWaypoint() []*DC.Point
	SetWaypoint(interface{}) 

}