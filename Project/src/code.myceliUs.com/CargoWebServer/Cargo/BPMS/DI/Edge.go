package DI

import(
"code.myceliUs.com/CargoWebServer/Cargo/BPMS/DC"
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