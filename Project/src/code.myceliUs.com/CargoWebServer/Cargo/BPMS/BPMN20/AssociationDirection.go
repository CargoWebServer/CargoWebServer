package BPMN20

type AssociationDirection int
const(
	AssociationDirection_None AssociationDirection = 1+iota
	AssociationDirection_One
	AssociationDirection_Both
)
