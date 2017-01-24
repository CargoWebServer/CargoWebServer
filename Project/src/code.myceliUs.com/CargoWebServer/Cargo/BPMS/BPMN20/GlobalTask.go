// +build BPMN
package BPMN20

type GlobalTask interface {
	/** Method of GlobalTask **/

	/** UUID **/
	GetUUID() string

	/** ResourceRole **/
	GetResourceRole() []ResourceRole
}
