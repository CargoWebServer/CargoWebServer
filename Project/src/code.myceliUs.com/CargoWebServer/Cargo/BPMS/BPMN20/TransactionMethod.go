// +build BPMN
package BPMN20

type TransactionMethod int

const (
	TransactionMethod_Compensate TransactionMethod = 1 + iota
	TransactionMethod_Image
	TransactionMethod_Store
)
