// +build BPMN20

package BPMN20

type Gateway interface{
	/** Method of Gateway **/

	/** UUID **/
	GetUUID() string

	/** GatewayDirection **/
	GetGatewayDirection() GatewayDirection

}