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

	/** Target **/
	GetTarget() DiagramElement

	/** Waypoint **/
	GetWaypoint() []*DC.Point

}