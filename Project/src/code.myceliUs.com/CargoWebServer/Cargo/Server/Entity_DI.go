package Server

/** Entity Prototype creation **/
func (this *EntityManager) create_DI_DiagramElementEntityPrototype() {

	var diagramElementEntityProto EntityPrototype
	diagramElementEntityProto.TypeName = "DI.DiagramElement"
	diagramElementEntityProto.IsAbstract = true
	diagramElementEntityProto.SubstitutionGroup = append(diagramElementEntityProto.SubstitutionGroup, "BPMNDI.BPMNShape")
	diagramElementEntityProto.SubstitutionGroup = append(diagramElementEntityProto.SubstitutionGroup, "BPMNDI.BPMNPlane")
	diagramElementEntityProto.SubstitutionGroup = append(diagramElementEntityProto.SubstitutionGroup, "BPMNDI.BPMNLabel")
	diagramElementEntityProto.SubstitutionGroup = append(diagramElementEntityProto.SubstitutionGroup, "BPMNDI.BPMNEdge")
	diagramElementEntityProto.Ids = append(diagramElementEntityProto.Ids, "uuid")
	diagramElementEntityProto.Fields = append(diagramElementEntityProto.Fields, "uuid")
	diagramElementEntityProto.FieldsType = append(diagramElementEntityProto.FieldsType, "xs.string")
	diagramElementEntityProto.FieldsOrder = append(diagramElementEntityProto.FieldsOrder, 0)
	diagramElementEntityProto.FieldsVisibility = append(diagramElementEntityProto.FieldsVisibility, false)
	diagramElementEntityProto.Indexs = append(diagramElementEntityProto.Indexs, "parentUuid")
	diagramElementEntityProto.Fields = append(diagramElementEntityProto.Fields, "parentUuid")
	diagramElementEntityProto.FieldsType = append(diagramElementEntityProto.FieldsType, "xs.string")
	diagramElementEntityProto.FieldsOrder = append(diagramElementEntityProto.FieldsOrder, 1)
	diagramElementEntityProto.FieldsVisibility = append(diagramElementEntityProto.FieldsVisibility, false)

	/** members of DiagramElement **/
	diagramElementEntityProto.FieldsOrder = append(diagramElementEntityProto.FieldsOrder, 2)
	diagramElementEntityProto.FieldsVisibility = append(diagramElementEntityProto.FieldsVisibility, true)
	diagramElementEntityProto.Fields = append(diagramElementEntityProto.Fields, "M_owningDiagram")
	diagramElementEntityProto.FieldsType = append(diagramElementEntityProto.FieldsType, "DI.Diagram:Ref")
	diagramElementEntityProto.FieldsOrder = append(diagramElementEntityProto.FieldsOrder, 3)
	diagramElementEntityProto.FieldsVisibility = append(diagramElementEntityProto.FieldsVisibility, true)
	diagramElementEntityProto.Fields = append(diagramElementEntityProto.Fields, "M_owningElement")
	diagramElementEntityProto.FieldsType = append(diagramElementEntityProto.FieldsType, "DI.DiagramElement:Ref")
	diagramElementEntityProto.FieldsOrder = append(diagramElementEntityProto.FieldsOrder, 4)
	diagramElementEntityProto.FieldsVisibility = append(diagramElementEntityProto.FieldsVisibility, true)
	diagramElementEntityProto.Fields = append(diagramElementEntityProto.Fields, "M_modelElement")
	diagramElementEntityProto.FieldsType = append(diagramElementEntityProto.FieldsType, "xs.interface{}:Ref")
	diagramElementEntityProto.FieldsOrder = append(diagramElementEntityProto.FieldsOrder, 5)
	diagramElementEntityProto.FieldsVisibility = append(diagramElementEntityProto.FieldsVisibility, true)
	diagramElementEntityProto.Fields = append(diagramElementEntityProto.Fields, "M_style")
	diagramElementEntityProto.FieldsType = append(diagramElementEntityProto.FieldsType, "DI.Style:Ref")
	diagramElementEntityProto.FieldsOrder = append(diagramElementEntityProto.FieldsOrder, 6)
	diagramElementEntityProto.FieldsVisibility = append(diagramElementEntityProto.FieldsVisibility, true)
	diagramElementEntityProto.Fields = append(diagramElementEntityProto.Fields, "M_ownedElement")
	diagramElementEntityProto.FieldsType = append(diagramElementEntityProto.FieldsType, "[]DI.DiagramElement")
	diagramElementEntityProto.Ids = append(diagramElementEntityProto.Ids, "M_id")
	diagramElementEntityProto.FieldsOrder = append(diagramElementEntityProto.FieldsOrder, 7)
	diagramElementEntityProto.FieldsVisibility = append(diagramElementEntityProto.FieldsVisibility, true)
	diagramElementEntityProto.Fields = append(diagramElementEntityProto.Fields, "M_id")
	diagramElementEntityProto.FieldsType = append(diagramElementEntityProto.FieldsType, "xs.ID")

	/** associations of DiagramElement **/
	diagramElementEntityProto.FieldsOrder = append(diagramElementEntityProto.FieldsOrder, 8)
	diagramElementEntityProto.FieldsVisibility = append(diagramElementEntityProto.FieldsVisibility, false)
	diagramElementEntityProto.Fields = append(diagramElementEntityProto.Fields, "M_sourceEdgePtr")
	diagramElementEntityProto.FieldsType = append(diagramElementEntityProto.FieldsType, "[]DI.Edge:Ref")
	diagramElementEntityProto.FieldsOrder = append(diagramElementEntityProto.FieldsOrder, 9)
	diagramElementEntityProto.FieldsVisibility = append(diagramElementEntityProto.FieldsVisibility, false)
	diagramElementEntityProto.Fields = append(diagramElementEntityProto.Fields, "M_targetEdgePtr")
	diagramElementEntityProto.FieldsType = append(diagramElementEntityProto.FieldsType, "[]DI.Edge:Ref")
	diagramElementEntityProto.FieldsOrder = append(diagramElementEntityProto.FieldsOrder, 10)
	diagramElementEntityProto.FieldsVisibility = append(diagramElementEntityProto.FieldsVisibility, false)
	diagramElementEntityProto.Fields = append(diagramElementEntityProto.Fields, "M_planePtr")
	diagramElementEntityProto.FieldsType = append(diagramElementEntityProto.FieldsType, "DI.Plane:Ref")
	diagramElementEntityProto.Fields = append(diagramElementEntityProto.Fields, "childsUuid")
	diagramElementEntityProto.FieldsType = append(diagramElementEntityProto.FieldsType, "[]xs.string")
	diagramElementEntityProto.FieldsOrder = append(diagramElementEntityProto.FieldsOrder, 11)
	diagramElementEntityProto.FieldsVisibility = append(diagramElementEntityProto.FieldsVisibility, false)

	diagramElementEntityProto.Fields = append(diagramElementEntityProto.Fields, "referenced")
	diagramElementEntityProto.FieldsType = append(diagramElementEntityProto.FieldsType, "[]EntityRef")
	diagramElementEntityProto.FieldsOrder = append(diagramElementEntityProto.FieldsOrder, 12)
	diagramElementEntityProto.FieldsVisibility = append(diagramElementEntityProto.FieldsVisibility, false)

	store := GetServer().GetDataManager().getDataStore(DIDB).(*KeyValueDataStore)
	store.SetEntityPrototype(&diagramElementEntityProto)

}

/** Entity Prototype creation **/
func (this *EntityManager) create_DI_NodeEntityPrototype() {

	var nodeEntityProto EntityPrototype
	nodeEntityProto.TypeName = "DI.Node"
	nodeEntityProto.IsAbstract = true
	nodeEntityProto.SuperTypeNames = append(nodeEntityProto.SuperTypeNames, "DI.DiagramElement")
	nodeEntityProto.SubstitutionGroup = append(nodeEntityProto.SubstitutionGroup, "BPMNDI.BPMNShape")
	nodeEntityProto.SubstitutionGroup = append(nodeEntityProto.SubstitutionGroup, "BPMNDI.BPMNPlane")
	nodeEntityProto.SubstitutionGroup = append(nodeEntityProto.SubstitutionGroup, "BPMNDI.BPMNLabel")
	nodeEntityProto.Ids = append(nodeEntityProto.Ids, "uuid")
	nodeEntityProto.Fields = append(nodeEntityProto.Fields, "uuid")
	nodeEntityProto.FieldsType = append(nodeEntityProto.FieldsType, "xs.string")
	nodeEntityProto.FieldsOrder = append(nodeEntityProto.FieldsOrder, 0)
	nodeEntityProto.FieldsVisibility = append(nodeEntityProto.FieldsVisibility, false)
	nodeEntityProto.Indexs = append(nodeEntityProto.Indexs, "parentUuid")
	nodeEntityProto.Fields = append(nodeEntityProto.Fields, "parentUuid")
	nodeEntityProto.FieldsType = append(nodeEntityProto.FieldsType, "xs.string")
	nodeEntityProto.FieldsOrder = append(nodeEntityProto.FieldsOrder, 1)
	nodeEntityProto.FieldsVisibility = append(nodeEntityProto.FieldsVisibility, false)

	/** members of DiagramElement **/
	nodeEntityProto.FieldsOrder = append(nodeEntityProto.FieldsOrder, 2)
	nodeEntityProto.FieldsVisibility = append(nodeEntityProto.FieldsVisibility, true)
	nodeEntityProto.Fields = append(nodeEntityProto.Fields, "M_owningDiagram")
	nodeEntityProto.FieldsType = append(nodeEntityProto.FieldsType, "DI.Diagram:Ref")
	nodeEntityProto.FieldsOrder = append(nodeEntityProto.FieldsOrder, 3)
	nodeEntityProto.FieldsVisibility = append(nodeEntityProto.FieldsVisibility, true)
	nodeEntityProto.Fields = append(nodeEntityProto.Fields, "M_owningElement")
	nodeEntityProto.FieldsType = append(nodeEntityProto.FieldsType, "DI.DiagramElement:Ref")
	nodeEntityProto.FieldsOrder = append(nodeEntityProto.FieldsOrder, 4)
	nodeEntityProto.FieldsVisibility = append(nodeEntityProto.FieldsVisibility, true)
	nodeEntityProto.Fields = append(nodeEntityProto.Fields, "M_modelElement")
	nodeEntityProto.FieldsType = append(nodeEntityProto.FieldsType, "xs.interface{}:Ref")
	nodeEntityProto.FieldsOrder = append(nodeEntityProto.FieldsOrder, 5)
	nodeEntityProto.FieldsVisibility = append(nodeEntityProto.FieldsVisibility, true)
	nodeEntityProto.Fields = append(nodeEntityProto.Fields, "M_style")
	nodeEntityProto.FieldsType = append(nodeEntityProto.FieldsType, "DI.Style:Ref")
	nodeEntityProto.FieldsOrder = append(nodeEntityProto.FieldsOrder, 6)
	nodeEntityProto.FieldsVisibility = append(nodeEntityProto.FieldsVisibility, true)
	nodeEntityProto.Fields = append(nodeEntityProto.Fields, "M_ownedElement")
	nodeEntityProto.FieldsType = append(nodeEntityProto.FieldsType, "[]DI.DiagramElement")
	nodeEntityProto.Ids = append(nodeEntityProto.Ids, "M_id")
	nodeEntityProto.FieldsOrder = append(nodeEntityProto.FieldsOrder, 7)
	nodeEntityProto.FieldsVisibility = append(nodeEntityProto.FieldsVisibility, true)
	nodeEntityProto.Fields = append(nodeEntityProto.Fields, "M_id")
	nodeEntityProto.FieldsType = append(nodeEntityProto.FieldsType, "xs.ID")

	/** members of Node **/
	/** No members **/

	/** associations of Node **/
	nodeEntityProto.FieldsOrder = append(nodeEntityProto.FieldsOrder, 8)
	nodeEntityProto.FieldsVisibility = append(nodeEntityProto.FieldsVisibility, false)
	nodeEntityProto.Fields = append(nodeEntityProto.Fields, "M_sourceEdgePtr")
	nodeEntityProto.FieldsType = append(nodeEntityProto.FieldsType, "[]DI.Edge:Ref")
	nodeEntityProto.FieldsOrder = append(nodeEntityProto.FieldsOrder, 9)
	nodeEntityProto.FieldsVisibility = append(nodeEntityProto.FieldsVisibility, false)
	nodeEntityProto.Fields = append(nodeEntityProto.Fields, "M_targetEdgePtr")
	nodeEntityProto.FieldsType = append(nodeEntityProto.FieldsType, "[]DI.Edge:Ref")
	nodeEntityProto.FieldsOrder = append(nodeEntityProto.FieldsOrder, 10)
	nodeEntityProto.FieldsVisibility = append(nodeEntityProto.FieldsVisibility, false)
	nodeEntityProto.Fields = append(nodeEntityProto.Fields, "M_planePtr")
	nodeEntityProto.FieldsType = append(nodeEntityProto.FieldsType, "DI.Plane:Ref")
	nodeEntityProto.Fields = append(nodeEntityProto.Fields, "childsUuid")
	nodeEntityProto.FieldsType = append(nodeEntityProto.FieldsType, "[]xs.string")
	nodeEntityProto.FieldsOrder = append(nodeEntityProto.FieldsOrder, 11)
	nodeEntityProto.FieldsVisibility = append(nodeEntityProto.FieldsVisibility, false)

	nodeEntityProto.Fields = append(nodeEntityProto.Fields, "referenced")
	nodeEntityProto.FieldsType = append(nodeEntityProto.FieldsType, "[]EntityRef")
	nodeEntityProto.FieldsOrder = append(nodeEntityProto.FieldsOrder, 12)
	nodeEntityProto.FieldsVisibility = append(nodeEntityProto.FieldsVisibility, false)

	store := GetServer().GetDataManager().getDataStore(DIDB).(*KeyValueDataStore)
	store.SetEntityPrototype(&nodeEntityProto)

}

/** Entity Prototype creation **/
func (this *EntityManager) create_DI_EdgeEntityPrototype() {

	var edgeEntityProto EntityPrototype
	edgeEntityProto.TypeName = "DI.Edge"
	edgeEntityProto.IsAbstract = true
	edgeEntityProto.SuperTypeNames = append(edgeEntityProto.SuperTypeNames, "DI.DiagramElement")
	edgeEntityProto.SubstitutionGroup = append(edgeEntityProto.SubstitutionGroup, "BPMNDI.BPMNEdge")
	edgeEntityProto.Ids = append(edgeEntityProto.Ids, "uuid")
	edgeEntityProto.Fields = append(edgeEntityProto.Fields, "uuid")
	edgeEntityProto.FieldsType = append(edgeEntityProto.FieldsType, "xs.string")
	edgeEntityProto.FieldsOrder = append(edgeEntityProto.FieldsOrder, 0)
	edgeEntityProto.FieldsVisibility = append(edgeEntityProto.FieldsVisibility, false)
	edgeEntityProto.Indexs = append(edgeEntityProto.Indexs, "parentUuid")
	edgeEntityProto.Fields = append(edgeEntityProto.Fields, "parentUuid")
	edgeEntityProto.FieldsType = append(edgeEntityProto.FieldsType, "xs.string")
	edgeEntityProto.FieldsOrder = append(edgeEntityProto.FieldsOrder, 1)
	edgeEntityProto.FieldsVisibility = append(edgeEntityProto.FieldsVisibility, false)

	/** members of DiagramElement **/
	edgeEntityProto.FieldsOrder = append(edgeEntityProto.FieldsOrder, 2)
	edgeEntityProto.FieldsVisibility = append(edgeEntityProto.FieldsVisibility, true)
	edgeEntityProto.Fields = append(edgeEntityProto.Fields, "M_owningDiagram")
	edgeEntityProto.FieldsType = append(edgeEntityProto.FieldsType, "DI.Diagram:Ref")
	edgeEntityProto.FieldsOrder = append(edgeEntityProto.FieldsOrder, 3)
	edgeEntityProto.FieldsVisibility = append(edgeEntityProto.FieldsVisibility, true)
	edgeEntityProto.Fields = append(edgeEntityProto.Fields, "M_owningElement")
	edgeEntityProto.FieldsType = append(edgeEntityProto.FieldsType, "DI.DiagramElement:Ref")
	edgeEntityProto.FieldsOrder = append(edgeEntityProto.FieldsOrder, 4)
	edgeEntityProto.FieldsVisibility = append(edgeEntityProto.FieldsVisibility, true)
	edgeEntityProto.Fields = append(edgeEntityProto.Fields, "M_modelElement")
	edgeEntityProto.FieldsType = append(edgeEntityProto.FieldsType, "xs.interface{}:Ref")
	edgeEntityProto.FieldsOrder = append(edgeEntityProto.FieldsOrder, 5)
	edgeEntityProto.FieldsVisibility = append(edgeEntityProto.FieldsVisibility, true)
	edgeEntityProto.Fields = append(edgeEntityProto.Fields, "M_style")
	edgeEntityProto.FieldsType = append(edgeEntityProto.FieldsType, "DI.Style:Ref")
	edgeEntityProto.FieldsOrder = append(edgeEntityProto.FieldsOrder, 6)
	edgeEntityProto.FieldsVisibility = append(edgeEntityProto.FieldsVisibility, true)
	edgeEntityProto.Fields = append(edgeEntityProto.Fields, "M_ownedElement")
	edgeEntityProto.FieldsType = append(edgeEntityProto.FieldsType, "[]DI.DiagramElement")
	edgeEntityProto.Ids = append(edgeEntityProto.Ids, "M_id")
	edgeEntityProto.FieldsOrder = append(edgeEntityProto.FieldsOrder, 7)
	edgeEntityProto.FieldsVisibility = append(edgeEntityProto.FieldsVisibility, true)
	edgeEntityProto.Fields = append(edgeEntityProto.Fields, "M_id")
	edgeEntityProto.FieldsType = append(edgeEntityProto.FieldsType, "xs.ID")

	/** members of Edge **/
	edgeEntityProto.FieldsOrder = append(edgeEntityProto.FieldsOrder, 8)
	edgeEntityProto.FieldsVisibility = append(edgeEntityProto.FieldsVisibility, true)
	edgeEntityProto.Fields = append(edgeEntityProto.Fields, "M_source")
	edgeEntityProto.FieldsType = append(edgeEntityProto.FieldsType, "DI.DiagramElement")
	edgeEntityProto.FieldsOrder = append(edgeEntityProto.FieldsOrder, 9)
	edgeEntityProto.FieldsVisibility = append(edgeEntityProto.FieldsVisibility, true)
	edgeEntityProto.Fields = append(edgeEntityProto.Fields, "M_target")
	edgeEntityProto.FieldsType = append(edgeEntityProto.FieldsType, "DI.DiagramElement")
	edgeEntityProto.FieldsOrder = append(edgeEntityProto.FieldsOrder, 10)
	edgeEntityProto.FieldsVisibility = append(edgeEntityProto.FieldsVisibility, true)
	edgeEntityProto.Fields = append(edgeEntityProto.Fields, "M_waypoint")
	edgeEntityProto.FieldsType = append(edgeEntityProto.FieldsType, "[]DI.DC.Point")

	/** associations of Edge **/
	edgeEntityProto.FieldsOrder = append(edgeEntityProto.FieldsOrder, 11)
	edgeEntityProto.FieldsVisibility = append(edgeEntityProto.FieldsVisibility, false)
	edgeEntityProto.Fields = append(edgeEntityProto.Fields, "M_sourceEdgePtr")
	edgeEntityProto.FieldsType = append(edgeEntityProto.FieldsType, "[]DI.Edge:Ref")
	edgeEntityProto.FieldsOrder = append(edgeEntityProto.FieldsOrder, 12)
	edgeEntityProto.FieldsVisibility = append(edgeEntityProto.FieldsVisibility, false)
	edgeEntityProto.Fields = append(edgeEntityProto.Fields, "M_targetEdgePtr")
	edgeEntityProto.FieldsType = append(edgeEntityProto.FieldsType, "[]DI.Edge:Ref")
	edgeEntityProto.FieldsOrder = append(edgeEntityProto.FieldsOrder, 13)
	edgeEntityProto.FieldsVisibility = append(edgeEntityProto.FieldsVisibility, false)
	edgeEntityProto.Fields = append(edgeEntityProto.Fields, "M_planePtr")
	edgeEntityProto.FieldsType = append(edgeEntityProto.FieldsType, "DI.Plane:Ref")
	edgeEntityProto.Fields = append(edgeEntityProto.Fields, "childsUuid")
	edgeEntityProto.FieldsType = append(edgeEntityProto.FieldsType, "[]xs.string")
	edgeEntityProto.FieldsOrder = append(edgeEntityProto.FieldsOrder, 14)
	edgeEntityProto.FieldsVisibility = append(edgeEntityProto.FieldsVisibility, false)

	edgeEntityProto.Fields = append(edgeEntityProto.Fields, "referenced")
	edgeEntityProto.FieldsType = append(edgeEntityProto.FieldsType, "[]EntityRef")
	edgeEntityProto.FieldsOrder = append(edgeEntityProto.FieldsOrder, 15)
	edgeEntityProto.FieldsVisibility = append(edgeEntityProto.FieldsVisibility, false)

	store := GetServer().GetDataManager().getDataStore(DIDB).(*KeyValueDataStore)
	store.SetEntityPrototype(&edgeEntityProto)

}

/** Entity Prototype creation **/
func (this *EntityManager) create_DI_DiagramEntityPrototype() {

	var diagramEntityProto EntityPrototype
	diagramEntityProto.TypeName = "DI.Diagram"
	diagramEntityProto.IsAbstract = true
	diagramEntityProto.SubstitutionGroup = append(diagramEntityProto.SubstitutionGroup, "BPMNDI.BPMNDiagram")
	diagramEntityProto.Ids = append(diagramEntityProto.Ids, "uuid")
	diagramEntityProto.Fields = append(diagramEntityProto.Fields, "uuid")
	diagramEntityProto.FieldsType = append(diagramEntityProto.FieldsType, "xs.string")
	diagramEntityProto.FieldsOrder = append(diagramEntityProto.FieldsOrder, 0)
	diagramEntityProto.FieldsVisibility = append(diagramEntityProto.FieldsVisibility, false)
	diagramEntityProto.Indexs = append(diagramEntityProto.Indexs, "parentUuid")
	diagramEntityProto.Fields = append(diagramEntityProto.Fields, "parentUuid")
	diagramEntityProto.FieldsType = append(diagramEntityProto.FieldsType, "xs.string")
	diagramEntityProto.FieldsOrder = append(diagramEntityProto.FieldsOrder, 1)
	diagramEntityProto.FieldsVisibility = append(diagramEntityProto.FieldsVisibility, false)

	/** members of Diagram **/
	diagramEntityProto.FieldsOrder = append(diagramEntityProto.FieldsOrder, 2)
	diagramEntityProto.FieldsVisibility = append(diagramEntityProto.FieldsVisibility, true)
	diagramEntityProto.Fields = append(diagramEntityProto.Fields, "M_rootElement")
	diagramEntityProto.FieldsType = append(diagramEntityProto.FieldsType, "DI.DiagramElement")
	diagramEntityProto.FieldsOrder = append(diagramEntityProto.FieldsOrder, 3)
	diagramEntityProto.FieldsVisibility = append(diagramEntityProto.FieldsVisibility, true)
	diagramEntityProto.Fields = append(diagramEntityProto.Fields, "M_name")
	diagramEntityProto.FieldsType = append(diagramEntityProto.FieldsType, "xs.string")
	diagramEntityProto.Ids = append(diagramEntityProto.Ids, "M_id")
	diagramEntityProto.FieldsOrder = append(diagramEntityProto.FieldsOrder, 4)
	diagramEntityProto.FieldsVisibility = append(diagramEntityProto.FieldsVisibility, true)
	diagramEntityProto.Fields = append(diagramEntityProto.Fields, "M_id")
	diagramEntityProto.FieldsType = append(diagramEntityProto.FieldsType, "xs.ID")
	diagramEntityProto.FieldsOrder = append(diagramEntityProto.FieldsOrder, 5)
	diagramEntityProto.FieldsVisibility = append(diagramEntityProto.FieldsVisibility, true)
	diagramEntityProto.Fields = append(diagramEntityProto.Fields, "M_documentation")
	diagramEntityProto.FieldsType = append(diagramEntityProto.FieldsType, "xs.string")
	diagramEntityProto.FieldsOrder = append(diagramEntityProto.FieldsOrder, 6)
	diagramEntityProto.FieldsVisibility = append(diagramEntityProto.FieldsVisibility, true)
	diagramEntityProto.Fields = append(diagramEntityProto.Fields, "M_resolution")
	diagramEntityProto.FieldsType = append(diagramEntityProto.FieldsType, "xs.float64")
	diagramEntityProto.FieldsOrder = append(diagramEntityProto.FieldsOrder, 7)
	diagramEntityProto.FieldsVisibility = append(diagramEntityProto.FieldsVisibility, true)
	diagramEntityProto.Fields = append(diagramEntityProto.Fields, "M_ownedStyle")
	diagramEntityProto.FieldsType = append(diagramEntityProto.FieldsType, "[]DI.Style")
	diagramEntityProto.Fields = append(diagramEntityProto.Fields, "childsUuid")
	diagramEntityProto.FieldsType = append(diagramEntityProto.FieldsType, "[]xs.string")
	diagramEntityProto.FieldsOrder = append(diagramEntityProto.FieldsOrder, 8)
	diagramEntityProto.FieldsVisibility = append(diagramEntityProto.FieldsVisibility, false)

	diagramEntityProto.Fields = append(diagramEntityProto.Fields, "referenced")
	diagramEntityProto.FieldsType = append(diagramEntityProto.FieldsType, "[]EntityRef")
	diagramEntityProto.FieldsOrder = append(diagramEntityProto.FieldsOrder, 9)
	diagramEntityProto.FieldsVisibility = append(diagramEntityProto.FieldsVisibility, false)

	store := GetServer().GetDataManager().getDataStore(DIDB).(*KeyValueDataStore)
	store.SetEntityPrototype(&diagramEntityProto)

}

/** Entity Prototype creation **/
func (this *EntityManager) create_DI_ShapeEntityPrototype() {

	var shapeEntityProto EntityPrototype
	shapeEntityProto.TypeName = "DI.Shape"
	shapeEntityProto.IsAbstract = true
	shapeEntityProto.SuperTypeNames = append(shapeEntityProto.SuperTypeNames, "DI.DiagramElement")
	shapeEntityProto.SuperTypeNames = append(shapeEntityProto.SuperTypeNames, "DI.Node")
	shapeEntityProto.SubstitutionGroup = append(shapeEntityProto.SubstitutionGroup, "BPMNDI.BPMNShape")
	shapeEntityProto.Ids = append(shapeEntityProto.Ids, "uuid")
	shapeEntityProto.Fields = append(shapeEntityProto.Fields, "uuid")
	shapeEntityProto.FieldsType = append(shapeEntityProto.FieldsType, "xs.string")
	shapeEntityProto.FieldsOrder = append(shapeEntityProto.FieldsOrder, 0)
	shapeEntityProto.FieldsVisibility = append(shapeEntityProto.FieldsVisibility, false)
	shapeEntityProto.Indexs = append(shapeEntityProto.Indexs, "parentUuid")
	shapeEntityProto.Fields = append(shapeEntityProto.Fields, "parentUuid")
	shapeEntityProto.FieldsType = append(shapeEntityProto.FieldsType, "xs.string")
	shapeEntityProto.FieldsOrder = append(shapeEntityProto.FieldsOrder, 1)
	shapeEntityProto.FieldsVisibility = append(shapeEntityProto.FieldsVisibility, false)

	/** members of DiagramElement **/
	shapeEntityProto.FieldsOrder = append(shapeEntityProto.FieldsOrder, 2)
	shapeEntityProto.FieldsVisibility = append(shapeEntityProto.FieldsVisibility, true)
	shapeEntityProto.Fields = append(shapeEntityProto.Fields, "M_owningDiagram")
	shapeEntityProto.FieldsType = append(shapeEntityProto.FieldsType, "DI.Diagram:Ref")
	shapeEntityProto.FieldsOrder = append(shapeEntityProto.FieldsOrder, 3)
	shapeEntityProto.FieldsVisibility = append(shapeEntityProto.FieldsVisibility, true)
	shapeEntityProto.Fields = append(shapeEntityProto.Fields, "M_owningElement")
	shapeEntityProto.FieldsType = append(shapeEntityProto.FieldsType, "DI.DiagramElement:Ref")
	shapeEntityProto.FieldsOrder = append(shapeEntityProto.FieldsOrder, 4)
	shapeEntityProto.FieldsVisibility = append(shapeEntityProto.FieldsVisibility, true)
	shapeEntityProto.Fields = append(shapeEntityProto.Fields, "M_modelElement")
	shapeEntityProto.FieldsType = append(shapeEntityProto.FieldsType, "xs.interface{}:Ref")
	shapeEntityProto.FieldsOrder = append(shapeEntityProto.FieldsOrder, 5)
	shapeEntityProto.FieldsVisibility = append(shapeEntityProto.FieldsVisibility, true)
	shapeEntityProto.Fields = append(shapeEntityProto.Fields, "M_style")
	shapeEntityProto.FieldsType = append(shapeEntityProto.FieldsType, "DI.Style:Ref")
	shapeEntityProto.FieldsOrder = append(shapeEntityProto.FieldsOrder, 6)
	shapeEntityProto.FieldsVisibility = append(shapeEntityProto.FieldsVisibility, true)
	shapeEntityProto.Fields = append(shapeEntityProto.Fields, "M_ownedElement")
	shapeEntityProto.FieldsType = append(shapeEntityProto.FieldsType, "[]DI.DiagramElement")
	shapeEntityProto.Ids = append(shapeEntityProto.Ids, "M_id")
	shapeEntityProto.FieldsOrder = append(shapeEntityProto.FieldsOrder, 7)
	shapeEntityProto.FieldsVisibility = append(shapeEntityProto.FieldsVisibility, true)
	shapeEntityProto.Fields = append(shapeEntityProto.Fields, "M_id")
	shapeEntityProto.FieldsType = append(shapeEntityProto.FieldsType, "xs.ID")

	/** members of Node **/
	/** No members **/

	/** members of Shape **/
	shapeEntityProto.FieldsOrder = append(shapeEntityProto.FieldsOrder, 8)
	shapeEntityProto.FieldsVisibility = append(shapeEntityProto.FieldsVisibility, true)
	shapeEntityProto.Fields = append(shapeEntityProto.Fields, "M_Bounds")
	shapeEntityProto.FieldsType = append(shapeEntityProto.FieldsType, "DI.DC.Bounds")

	/** associations of Shape **/
	shapeEntityProto.FieldsOrder = append(shapeEntityProto.FieldsOrder, 9)
	shapeEntityProto.FieldsVisibility = append(shapeEntityProto.FieldsVisibility, false)
	shapeEntityProto.Fields = append(shapeEntityProto.Fields, "M_sourceEdgePtr")
	shapeEntityProto.FieldsType = append(shapeEntityProto.FieldsType, "[]DI.Edge:Ref")
	shapeEntityProto.FieldsOrder = append(shapeEntityProto.FieldsOrder, 10)
	shapeEntityProto.FieldsVisibility = append(shapeEntityProto.FieldsVisibility, false)
	shapeEntityProto.Fields = append(shapeEntityProto.Fields, "M_targetEdgePtr")
	shapeEntityProto.FieldsType = append(shapeEntityProto.FieldsType, "[]DI.Edge:Ref")
	shapeEntityProto.FieldsOrder = append(shapeEntityProto.FieldsOrder, 11)
	shapeEntityProto.FieldsVisibility = append(shapeEntityProto.FieldsVisibility, false)
	shapeEntityProto.Fields = append(shapeEntityProto.Fields, "M_planePtr")
	shapeEntityProto.FieldsType = append(shapeEntityProto.FieldsType, "DI.Plane:Ref")
	shapeEntityProto.Fields = append(shapeEntityProto.Fields, "childsUuid")
	shapeEntityProto.FieldsType = append(shapeEntityProto.FieldsType, "[]xs.string")
	shapeEntityProto.FieldsOrder = append(shapeEntityProto.FieldsOrder, 12)
	shapeEntityProto.FieldsVisibility = append(shapeEntityProto.FieldsVisibility, false)

	shapeEntityProto.Fields = append(shapeEntityProto.Fields, "referenced")
	shapeEntityProto.FieldsType = append(shapeEntityProto.FieldsType, "[]EntityRef")
	shapeEntityProto.FieldsOrder = append(shapeEntityProto.FieldsOrder, 13)
	shapeEntityProto.FieldsVisibility = append(shapeEntityProto.FieldsVisibility, false)

	store := GetServer().GetDataManager().getDataStore(DIDB).(*KeyValueDataStore)
	store.SetEntityPrototype(&shapeEntityProto)

}

/** Entity Prototype creation **/
func (this *EntityManager) create_DI_PlaneEntityPrototype() {

	var planeEntityProto EntityPrototype
	planeEntityProto.TypeName = "DI.Plane"
	planeEntityProto.IsAbstract = true
	planeEntityProto.SuperTypeNames = append(planeEntityProto.SuperTypeNames, "DI.DiagramElement")
	planeEntityProto.SuperTypeNames = append(planeEntityProto.SuperTypeNames, "DI.Node")
	planeEntityProto.SubstitutionGroup = append(planeEntityProto.SubstitutionGroup, "BPMNDI.BPMNPlane")
	planeEntityProto.Ids = append(planeEntityProto.Ids, "uuid")
	planeEntityProto.Fields = append(planeEntityProto.Fields, "uuid")
	planeEntityProto.FieldsType = append(planeEntityProto.FieldsType, "xs.string")
	planeEntityProto.FieldsOrder = append(planeEntityProto.FieldsOrder, 0)
	planeEntityProto.FieldsVisibility = append(planeEntityProto.FieldsVisibility, false)
	planeEntityProto.Indexs = append(planeEntityProto.Indexs, "parentUuid")
	planeEntityProto.Fields = append(planeEntityProto.Fields, "parentUuid")
	planeEntityProto.FieldsType = append(planeEntityProto.FieldsType, "xs.string")
	planeEntityProto.FieldsOrder = append(planeEntityProto.FieldsOrder, 1)
	planeEntityProto.FieldsVisibility = append(planeEntityProto.FieldsVisibility, false)

	/** members of DiagramElement **/
	planeEntityProto.FieldsOrder = append(planeEntityProto.FieldsOrder, 2)
	planeEntityProto.FieldsVisibility = append(planeEntityProto.FieldsVisibility, true)
	planeEntityProto.Fields = append(planeEntityProto.Fields, "M_owningDiagram")
	planeEntityProto.FieldsType = append(planeEntityProto.FieldsType, "DI.Diagram:Ref")
	planeEntityProto.FieldsOrder = append(planeEntityProto.FieldsOrder, 3)
	planeEntityProto.FieldsVisibility = append(planeEntityProto.FieldsVisibility, true)
	planeEntityProto.Fields = append(planeEntityProto.Fields, "M_owningElement")
	planeEntityProto.FieldsType = append(planeEntityProto.FieldsType, "DI.DiagramElement:Ref")
	planeEntityProto.FieldsOrder = append(planeEntityProto.FieldsOrder, 4)
	planeEntityProto.FieldsVisibility = append(planeEntityProto.FieldsVisibility, true)
	planeEntityProto.Fields = append(planeEntityProto.Fields, "M_modelElement")
	planeEntityProto.FieldsType = append(planeEntityProto.FieldsType, "xs.interface{}:Ref")
	planeEntityProto.FieldsOrder = append(planeEntityProto.FieldsOrder, 5)
	planeEntityProto.FieldsVisibility = append(planeEntityProto.FieldsVisibility, true)
	planeEntityProto.Fields = append(planeEntityProto.Fields, "M_style")
	planeEntityProto.FieldsType = append(planeEntityProto.FieldsType, "DI.Style:Ref")
	planeEntityProto.FieldsOrder = append(planeEntityProto.FieldsOrder, 6)
	planeEntityProto.FieldsVisibility = append(planeEntityProto.FieldsVisibility, true)
	planeEntityProto.Fields = append(planeEntityProto.Fields, "M_ownedElement")
	planeEntityProto.FieldsType = append(planeEntityProto.FieldsType, "[]DI.DiagramElement")
	planeEntityProto.Ids = append(planeEntityProto.Ids, "M_id")
	planeEntityProto.FieldsOrder = append(planeEntityProto.FieldsOrder, 7)
	planeEntityProto.FieldsVisibility = append(planeEntityProto.FieldsVisibility, true)
	planeEntityProto.Fields = append(planeEntityProto.Fields, "M_id")
	planeEntityProto.FieldsType = append(planeEntityProto.FieldsType, "xs.ID")

	/** members of Node **/
	/** No members **/

	/** members of Plane **/
	planeEntityProto.FieldsOrder = append(planeEntityProto.FieldsOrder, 8)
	planeEntityProto.FieldsVisibility = append(planeEntityProto.FieldsVisibility, true)
	planeEntityProto.Fields = append(planeEntityProto.Fields, "M_DiagramElement")
	planeEntityProto.FieldsType = append(planeEntityProto.FieldsType, "[]DI.DiagramElement")

	/** associations of Plane **/
	planeEntityProto.FieldsOrder = append(planeEntityProto.FieldsOrder, 9)
	planeEntityProto.FieldsVisibility = append(planeEntityProto.FieldsVisibility, false)
	planeEntityProto.Fields = append(planeEntityProto.Fields, "M_sourceEdgePtr")
	planeEntityProto.FieldsType = append(planeEntityProto.FieldsType, "[]DI.Edge:Ref")
	planeEntityProto.FieldsOrder = append(planeEntityProto.FieldsOrder, 10)
	planeEntityProto.FieldsVisibility = append(planeEntityProto.FieldsVisibility, false)
	planeEntityProto.Fields = append(planeEntityProto.Fields, "M_targetEdgePtr")
	planeEntityProto.FieldsType = append(planeEntityProto.FieldsType, "[]DI.Edge:Ref")
	planeEntityProto.FieldsOrder = append(planeEntityProto.FieldsOrder, 11)
	planeEntityProto.FieldsVisibility = append(planeEntityProto.FieldsVisibility, false)
	planeEntityProto.Fields = append(planeEntityProto.Fields, "M_planePtr")
	planeEntityProto.FieldsType = append(planeEntityProto.FieldsType, "DI.Plane:Ref")
	planeEntityProto.Fields = append(planeEntityProto.Fields, "childsUuid")
	planeEntityProto.FieldsType = append(planeEntityProto.FieldsType, "[]xs.string")
	planeEntityProto.FieldsOrder = append(planeEntityProto.FieldsOrder, 12)
	planeEntityProto.FieldsVisibility = append(planeEntityProto.FieldsVisibility, false)

	planeEntityProto.Fields = append(planeEntityProto.Fields, "referenced")
	planeEntityProto.FieldsType = append(planeEntityProto.FieldsType, "[]EntityRef")
	planeEntityProto.FieldsOrder = append(planeEntityProto.FieldsOrder, 13)
	planeEntityProto.FieldsVisibility = append(planeEntityProto.FieldsVisibility, false)

	store := GetServer().GetDataManager().getDataStore(DIDB).(*KeyValueDataStore)
	store.SetEntityPrototype(&planeEntityProto)

}

/** Entity Prototype creation **/
func (this *EntityManager) create_DI_LabeledEdgeEntityPrototype() {

	var labeledEdgeEntityProto EntityPrototype
	labeledEdgeEntityProto.TypeName = "DI.LabeledEdge"
	labeledEdgeEntityProto.IsAbstract = true
	labeledEdgeEntityProto.SuperTypeNames = append(labeledEdgeEntityProto.SuperTypeNames, "DI.DiagramElement")
	labeledEdgeEntityProto.SuperTypeNames = append(labeledEdgeEntityProto.SuperTypeNames, "DI.Edge")
	labeledEdgeEntityProto.SubstitutionGroup = append(labeledEdgeEntityProto.SubstitutionGroup, "BPMNDI.BPMNEdge")
	labeledEdgeEntityProto.Ids = append(labeledEdgeEntityProto.Ids, "uuid")
	labeledEdgeEntityProto.Fields = append(labeledEdgeEntityProto.Fields, "uuid")
	labeledEdgeEntityProto.FieldsType = append(labeledEdgeEntityProto.FieldsType, "xs.string")
	labeledEdgeEntityProto.FieldsOrder = append(labeledEdgeEntityProto.FieldsOrder, 0)
	labeledEdgeEntityProto.FieldsVisibility = append(labeledEdgeEntityProto.FieldsVisibility, false)
	labeledEdgeEntityProto.Indexs = append(labeledEdgeEntityProto.Indexs, "parentUuid")
	labeledEdgeEntityProto.Fields = append(labeledEdgeEntityProto.Fields, "parentUuid")
	labeledEdgeEntityProto.FieldsType = append(labeledEdgeEntityProto.FieldsType, "xs.string")
	labeledEdgeEntityProto.FieldsOrder = append(labeledEdgeEntityProto.FieldsOrder, 1)
	labeledEdgeEntityProto.FieldsVisibility = append(labeledEdgeEntityProto.FieldsVisibility, false)

	/** members of DiagramElement **/
	labeledEdgeEntityProto.FieldsOrder = append(labeledEdgeEntityProto.FieldsOrder, 2)
	labeledEdgeEntityProto.FieldsVisibility = append(labeledEdgeEntityProto.FieldsVisibility, true)
	labeledEdgeEntityProto.Fields = append(labeledEdgeEntityProto.Fields, "M_owningDiagram")
	labeledEdgeEntityProto.FieldsType = append(labeledEdgeEntityProto.FieldsType, "DI.Diagram:Ref")
	labeledEdgeEntityProto.FieldsOrder = append(labeledEdgeEntityProto.FieldsOrder, 3)
	labeledEdgeEntityProto.FieldsVisibility = append(labeledEdgeEntityProto.FieldsVisibility, true)
	labeledEdgeEntityProto.Fields = append(labeledEdgeEntityProto.Fields, "M_owningElement")
	labeledEdgeEntityProto.FieldsType = append(labeledEdgeEntityProto.FieldsType, "DI.DiagramElement:Ref")
	labeledEdgeEntityProto.FieldsOrder = append(labeledEdgeEntityProto.FieldsOrder, 4)
	labeledEdgeEntityProto.FieldsVisibility = append(labeledEdgeEntityProto.FieldsVisibility, true)
	labeledEdgeEntityProto.Fields = append(labeledEdgeEntityProto.Fields, "M_modelElement")
	labeledEdgeEntityProto.FieldsType = append(labeledEdgeEntityProto.FieldsType, "xs.interface{}:Ref")
	labeledEdgeEntityProto.FieldsOrder = append(labeledEdgeEntityProto.FieldsOrder, 5)
	labeledEdgeEntityProto.FieldsVisibility = append(labeledEdgeEntityProto.FieldsVisibility, true)
	labeledEdgeEntityProto.Fields = append(labeledEdgeEntityProto.Fields, "M_style")
	labeledEdgeEntityProto.FieldsType = append(labeledEdgeEntityProto.FieldsType, "DI.Style:Ref")
	labeledEdgeEntityProto.FieldsOrder = append(labeledEdgeEntityProto.FieldsOrder, 6)
	labeledEdgeEntityProto.FieldsVisibility = append(labeledEdgeEntityProto.FieldsVisibility, true)
	labeledEdgeEntityProto.Fields = append(labeledEdgeEntityProto.Fields, "M_ownedElement")
	labeledEdgeEntityProto.FieldsType = append(labeledEdgeEntityProto.FieldsType, "[]DI.DiagramElement")
	labeledEdgeEntityProto.Ids = append(labeledEdgeEntityProto.Ids, "M_id")
	labeledEdgeEntityProto.FieldsOrder = append(labeledEdgeEntityProto.FieldsOrder, 7)
	labeledEdgeEntityProto.FieldsVisibility = append(labeledEdgeEntityProto.FieldsVisibility, true)
	labeledEdgeEntityProto.Fields = append(labeledEdgeEntityProto.Fields, "M_id")
	labeledEdgeEntityProto.FieldsType = append(labeledEdgeEntityProto.FieldsType, "xs.ID")

	/** members of Edge **/
	labeledEdgeEntityProto.FieldsOrder = append(labeledEdgeEntityProto.FieldsOrder, 8)
	labeledEdgeEntityProto.FieldsVisibility = append(labeledEdgeEntityProto.FieldsVisibility, true)
	labeledEdgeEntityProto.Fields = append(labeledEdgeEntityProto.Fields, "M_source")
	labeledEdgeEntityProto.FieldsType = append(labeledEdgeEntityProto.FieldsType, "DI.DiagramElement")
	labeledEdgeEntityProto.FieldsOrder = append(labeledEdgeEntityProto.FieldsOrder, 9)
	labeledEdgeEntityProto.FieldsVisibility = append(labeledEdgeEntityProto.FieldsVisibility, true)
	labeledEdgeEntityProto.Fields = append(labeledEdgeEntityProto.Fields, "M_target")
	labeledEdgeEntityProto.FieldsType = append(labeledEdgeEntityProto.FieldsType, "DI.DiagramElement")
	labeledEdgeEntityProto.FieldsOrder = append(labeledEdgeEntityProto.FieldsOrder, 10)
	labeledEdgeEntityProto.FieldsVisibility = append(labeledEdgeEntityProto.FieldsVisibility, true)
	labeledEdgeEntityProto.Fields = append(labeledEdgeEntityProto.Fields, "M_waypoint")
	labeledEdgeEntityProto.FieldsType = append(labeledEdgeEntityProto.FieldsType, "[]DI.DC.Point")

	/** members of LabeledEdge **/
	labeledEdgeEntityProto.FieldsOrder = append(labeledEdgeEntityProto.FieldsOrder, 11)
	labeledEdgeEntityProto.FieldsVisibility = append(labeledEdgeEntityProto.FieldsVisibility, true)
	labeledEdgeEntityProto.Fields = append(labeledEdgeEntityProto.Fields, "M_ownedLabel")
	labeledEdgeEntityProto.FieldsType = append(labeledEdgeEntityProto.FieldsType, "[]DI.Label")

	/** associations of LabeledEdge **/
	labeledEdgeEntityProto.FieldsOrder = append(labeledEdgeEntityProto.FieldsOrder, 12)
	labeledEdgeEntityProto.FieldsVisibility = append(labeledEdgeEntityProto.FieldsVisibility, false)
	labeledEdgeEntityProto.Fields = append(labeledEdgeEntityProto.Fields, "M_sourceEdgePtr")
	labeledEdgeEntityProto.FieldsType = append(labeledEdgeEntityProto.FieldsType, "[]DI.Edge:Ref")
	labeledEdgeEntityProto.FieldsOrder = append(labeledEdgeEntityProto.FieldsOrder, 13)
	labeledEdgeEntityProto.FieldsVisibility = append(labeledEdgeEntityProto.FieldsVisibility, false)
	labeledEdgeEntityProto.Fields = append(labeledEdgeEntityProto.Fields, "M_targetEdgePtr")
	labeledEdgeEntityProto.FieldsType = append(labeledEdgeEntityProto.FieldsType, "[]DI.Edge:Ref")
	labeledEdgeEntityProto.FieldsOrder = append(labeledEdgeEntityProto.FieldsOrder, 14)
	labeledEdgeEntityProto.FieldsVisibility = append(labeledEdgeEntityProto.FieldsVisibility, false)
	labeledEdgeEntityProto.Fields = append(labeledEdgeEntityProto.Fields, "M_planePtr")
	labeledEdgeEntityProto.FieldsType = append(labeledEdgeEntityProto.FieldsType, "DI.Plane:Ref")
	labeledEdgeEntityProto.Fields = append(labeledEdgeEntityProto.Fields, "childsUuid")
	labeledEdgeEntityProto.FieldsType = append(labeledEdgeEntityProto.FieldsType, "[]xs.string")
	labeledEdgeEntityProto.FieldsOrder = append(labeledEdgeEntityProto.FieldsOrder, 15)
	labeledEdgeEntityProto.FieldsVisibility = append(labeledEdgeEntityProto.FieldsVisibility, false)

	labeledEdgeEntityProto.Fields = append(labeledEdgeEntityProto.Fields, "referenced")
	labeledEdgeEntityProto.FieldsType = append(labeledEdgeEntityProto.FieldsType, "[]EntityRef")
	labeledEdgeEntityProto.FieldsOrder = append(labeledEdgeEntityProto.FieldsOrder, 16)
	labeledEdgeEntityProto.FieldsVisibility = append(labeledEdgeEntityProto.FieldsVisibility, false)

	store := GetServer().GetDataManager().getDataStore(DIDB).(*KeyValueDataStore)
	store.SetEntityPrototype(&labeledEdgeEntityProto)

}

/** Entity Prototype creation **/
func (this *EntityManager) create_DI_LabeledShapeEntityPrototype() {

	var labeledShapeEntityProto EntityPrototype
	labeledShapeEntityProto.TypeName = "DI.LabeledShape"
	labeledShapeEntityProto.IsAbstract = true
	labeledShapeEntityProto.SuperTypeNames = append(labeledShapeEntityProto.SuperTypeNames, "DI.DiagramElement")
	labeledShapeEntityProto.SuperTypeNames = append(labeledShapeEntityProto.SuperTypeNames, "DI.Node")
	labeledShapeEntityProto.SuperTypeNames = append(labeledShapeEntityProto.SuperTypeNames, "DI.Shape")
	labeledShapeEntityProto.SubstitutionGroup = append(labeledShapeEntityProto.SubstitutionGroup, "BPMNDI.BPMNShape")
	labeledShapeEntityProto.Ids = append(labeledShapeEntityProto.Ids, "uuid")
	labeledShapeEntityProto.Fields = append(labeledShapeEntityProto.Fields, "uuid")
	labeledShapeEntityProto.FieldsType = append(labeledShapeEntityProto.FieldsType, "xs.string")
	labeledShapeEntityProto.FieldsOrder = append(labeledShapeEntityProto.FieldsOrder, 0)
	labeledShapeEntityProto.FieldsVisibility = append(labeledShapeEntityProto.FieldsVisibility, false)
	labeledShapeEntityProto.Indexs = append(labeledShapeEntityProto.Indexs, "parentUuid")
	labeledShapeEntityProto.Fields = append(labeledShapeEntityProto.Fields, "parentUuid")
	labeledShapeEntityProto.FieldsType = append(labeledShapeEntityProto.FieldsType, "xs.string")
	labeledShapeEntityProto.FieldsOrder = append(labeledShapeEntityProto.FieldsOrder, 1)
	labeledShapeEntityProto.FieldsVisibility = append(labeledShapeEntityProto.FieldsVisibility, false)

	/** members of DiagramElement **/
	labeledShapeEntityProto.FieldsOrder = append(labeledShapeEntityProto.FieldsOrder, 2)
	labeledShapeEntityProto.FieldsVisibility = append(labeledShapeEntityProto.FieldsVisibility, true)
	labeledShapeEntityProto.Fields = append(labeledShapeEntityProto.Fields, "M_owningDiagram")
	labeledShapeEntityProto.FieldsType = append(labeledShapeEntityProto.FieldsType, "DI.Diagram:Ref")
	labeledShapeEntityProto.FieldsOrder = append(labeledShapeEntityProto.FieldsOrder, 3)
	labeledShapeEntityProto.FieldsVisibility = append(labeledShapeEntityProto.FieldsVisibility, true)
	labeledShapeEntityProto.Fields = append(labeledShapeEntityProto.Fields, "M_owningElement")
	labeledShapeEntityProto.FieldsType = append(labeledShapeEntityProto.FieldsType, "DI.DiagramElement:Ref")
	labeledShapeEntityProto.FieldsOrder = append(labeledShapeEntityProto.FieldsOrder, 4)
	labeledShapeEntityProto.FieldsVisibility = append(labeledShapeEntityProto.FieldsVisibility, true)
	labeledShapeEntityProto.Fields = append(labeledShapeEntityProto.Fields, "M_modelElement")
	labeledShapeEntityProto.FieldsType = append(labeledShapeEntityProto.FieldsType, "xs.interface{}:Ref")
	labeledShapeEntityProto.FieldsOrder = append(labeledShapeEntityProto.FieldsOrder, 5)
	labeledShapeEntityProto.FieldsVisibility = append(labeledShapeEntityProto.FieldsVisibility, true)
	labeledShapeEntityProto.Fields = append(labeledShapeEntityProto.Fields, "M_style")
	labeledShapeEntityProto.FieldsType = append(labeledShapeEntityProto.FieldsType, "DI.Style:Ref")
	labeledShapeEntityProto.FieldsOrder = append(labeledShapeEntityProto.FieldsOrder, 6)
	labeledShapeEntityProto.FieldsVisibility = append(labeledShapeEntityProto.FieldsVisibility, true)
	labeledShapeEntityProto.Fields = append(labeledShapeEntityProto.Fields, "M_ownedElement")
	labeledShapeEntityProto.FieldsType = append(labeledShapeEntityProto.FieldsType, "[]DI.DiagramElement")
	labeledShapeEntityProto.Ids = append(labeledShapeEntityProto.Ids, "M_id")
	labeledShapeEntityProto.FieldsOrder = append(labeledShapeEntityProto.FieldsOrder, 7)
	labeledShapeEntityProto.FieldsVisibility = append(labeledShapeEntityProto.FieldsVisibility, true)
	labeledShapeEntityProto.Fields = append(labeledShapeEntityProto.Fields, "M_id")
	labeledShapeEntityProto.FieldsType = append(labeledShapeEntityProto.FieldsType, "xs.ID")

	/** members of Node **/
	/** No members **/

	/** members of Shape **/
	labeledShapeEntityProto.FieldsOrder = append(labeledShapeEntityProto.FieldsOrder, 8)
	labeledShapeEntityProto.FieldsVisibility = append(labeledShapeEntityProto.FieldsVisibility, true)
	labeledShapeEntityProto.Fields = append(labeledShapeEntityProto.Fields, "M_Bounds")
	labeledShapeEntityProto.FieldsType = append(labeledShapeEntityProto.FieldsType, "DI.DC.Bounds")

	/** members of LabeledShape **/
	labeledShapeEntityProto.FieldsOrder = append(labeledShapeEntityProto.FieldsOrder, 9)
	labeledShapeEntityProto.FieldsVisibility = append(labeledShapeEntityProto.FieldsVisibility, true)
	labeledShapeEntityProto.Fields = append(labeledShapeEntityProto.Fields, "M_ownedLabel")
	labeledShapeEntityProto.FieldsType = append(labeledShapeEntityProto.FieldsType, "[]DI.Label")

	/** associations of LabeledShape **/
	labeledShapeEntityProto.FieldsOrder = append(labeledShapeEntityProto.FieldsOrder, 10)
	labeledShapeEntityProto.FieldsVisibility = append(labeledShapeEntityProto.FieldsVisibility, false)
	labeledShapeEntityProto.Fields = append(labeledShapeEntityProto.Fields, "M_sourceEdgePtr")
	labeledShapeEntityProto.FieldsType = append(labeledShapeEntityProto.FieldsType, "[]DI.Edge:Ref")
	labeledShapeEntityProto.FieldsOrder = append(labeledShapeEntityProto.FieldsOrder, 11)
	labeledShapeEntityProto.FieldsVisibility = append(labeledShapeEntityProto.FieldsVisibility, false)
	labeledShapeEntityProto.Fields = append(labeledShapeEntityProto.Fields, "M_targetEdgePtr")
	labeledShapeEntityProto.FieldsType = append(labeledShapeEntityProto.FieldsType, "[]DI.Edge:Ref")
	labeledShapeEntityProto.FieldsOrder = append(labeledShapeEntityProto.FieldsOrder, 12)
	labeledShapeEntityProto.FieldsVisibility = append(labeledShapeEntityProto.FieldsVisibility, false)
	labeledShapeEntityProto.Fields = append(labeledShapeEntityProto.Fields, "M_planePtr")
	labeledShapeEntityProto.FieldsType = append(labeledShapeEntityProto.FieldsType, "DI.Plane:Ref")
	labeledShapeEntityProto.Fields = append(labeledShapeEntityProto.Fields, "childsUuid")
	labeledShapeEntityProto.FieldsType = append(labeledShapeEntityProto.FieldsType, "[]xs.string")
	labeledShapeEntityProto.FieldsOrder = append(labeledShapeEntityProto.FieldsOrder, 13)
	labeledShapeEntityProto.FieldsVisibility = append(labeledShapeEntityProto.FieldsVisibility, false)

	labeledShapeEntityProto.Fields = append(labeledShapeEntityProto.Fields, "referenced")
	labeledShapeEntityProto.FieldsType = append(labeledShapeEntityProto.FieldsType, "[]EntityRef")
	labeledShapeEntityProto.FieldsOrder = append(labeledShapeEntityProto.FieldsOrder, 14)
	labeledShapeEntityProto.FieldsVisibility = append(labeledShapeEntityProto.FieldsVisibility, false)

	store := GetServer().GetDataManager().getDataStore(DIDB).(*KeyValueDataStore)
	store.SetEntityPrototype(&labeledShapeEntityProto)

}

/** Entity Prototype creation **/
func (this *EntityManager) create_DI_LabelEntityPrototype() {

	var labelEntityProto EntityPrototype
	labelEntityProto.TypeName = "DI.Label"
	labelEntityProto.IsAbstract = true
	labelEntityProto.SuperTypeNames = append(labelEntityProto.SuperTypeNames, "DI.DiagramElement")
	labelEntityProto.SuperTypeNames = append(labelEntityProto.SuperTypeNames, "DI.Node")
	labelEntityProto.SubstitutionGroup = append(labelEntityProto.SubstitutionGroup, "BPMNDI.BPMNLabel")
	labelEntityProto.Ids = append(labelEntityProto.Ids, "uuid")
	labelEntityProto.Fields = append(labelEntityProto.Fields, "uuid")
	labelEntityProto.FieldsType = append(labelEntityProto.FieldsType, "xs.string")
	labelEntityProto.FieldsOrder = append(labelEntityProto.FieldsOrder, 0)
	labelEntityProto.FieldsVisibility = append(labelEntityProto.FieldsVisibility, false)
	labelEntityProto.Indexs = append(labelEntityProto.Indexs, "parentUuid")
	labelEntityProto.Fields = append(labelEntityProto.Fields, "parentUuid")
	labelEntityProto.FieldsType = append(labelEntityProto.FieldsType, "xs.string")
	labelEntityProto.FieldsOrder = append(labelEntityProto.FieldsOrder, 1)
	labelEntityProto.FieldsVisibility = append(labelEntityProto.FieldsVisibility, false)

	/** members of DiagramElement **/
	labelEntityProto.FieldsOrder = append(labelEntityProto.FieldsOrder, 2)
	labelEntityProto.FieldsVisibility = append(labelEntityProto.FieldsVisibility, true)
	labelEntityProto.Fields = append(labelEntityProto.Fields, "M_owningDiagram")
	labelEntityProto.FieldsType = append(labelEntityProto.FieldsType, "DI.Diagram:Ref")
	labelEntityProto.FieldsOrder = append(labelEntityProto.FieldsOrder, 3)
	labelEntityProto.FieldsVisibility = append(labelEntityProto.FieldsVisibility, true)
	labelEntityProto.Fields = append(labelEntityProto.Fields, "M_owningElement")
	labelEntityProto.FieldsType = append(labelEntityProto.FieldsType, "DI.DiagramElement:Ref")
	labelEntityProto.FieldsOrder = append(labelEntityProto.FieldsOrder, 4)
	labelEntityProto.FieldsVisibility = append(labelEntityProto.FieldsVisibility, true)
	labelEntityProto.Fields = append(labelEntityProto.Fields, "M_modelElement")
	labelEntityProto.FieldsType = append(labelEntityProto.FieldsType, "xs.interface{}:Ref")
	labelEntityProto.FieldsOrder = append(labelEntityProto.FieldsOrder, 5)
	labelEntityProto.FieldsVisibility = append(labelEntityProto.FieldsVisibility, true)
	labelEntityProto.Fields = append(labelEntityProto.Fields, "M_style")
	labelEntityProto.FieldsType = append(labelEntityProto.FieldsType, "DI.Style:Ref")
	labelEntityProto.FieldsOrder = append(labelEntityProto.FieldsOrder, 6)
	labelEntityProto.FieldsVisibility = append(labelEntityProto.FieldsVisibility, true)
	labelEntityProto.Fields = append(labelEntityProto.Fields, "M_ownedElement")
	labelEntityProto.FieldsType = append(labelEntityProto.FieldsType, "[]DI.DiagramElement")
	labelEntityProto.Ids = append(labelEntityProto.Ids, "M_id")
	labelEntityProto.FieldsOrder = append(labelEntityProto.FieldsOrder, 7)
	labelEntityProto.FieldsVisibility = append(labelEntityProto.FieldsVisibility, true)
	labelEntityProto.Fields = append(labelEntityProto.Fields, "M_id")
	labelEntityProto.FieldsType = append(labelEntityProto.FieldsType, "xs.ID")

	/** members of Node **/
	/** No members **/

	/** members of Label **/
	labelEntityProto.FieldsOrder = append(labelEntityProto.FieldsOrder, 8)
	labelEntityProto.FieldsVisibility = append(labelEntityProto.FieldsVisibility, true)
	labelEntityProto.Fields = append(labelEntityProto.Fields, "M_Bounds")
	labelEntityProto.FieldsType = append(labelEntityProto.FieldsType, "DI.DC.Bounds")

	/** associations of Label **/
	labelEntityProto.FieldsOrder = append(labelEntityProto.FieldsOrder, 9)
	labelEntityProto.FieldsVisibility = append(labelEntityProto.FieldsVisibility, false)
	labelEntityProto.Fields = append(labelEntityProto.Fields, "M_owningEdgePtr")
	labelEntityProto.FieldsType = append(labelEntityProto.FieldsType, "DI.LabeledEdge:Ref")
	labelEntityProto.FieldsOrder = append(labelEntityProto.FieldsOrder, 10)
	labelEntityProto.FieldsVisibility = append(labelEntityProto.FieldsVisibility, false)
	labelEntityProto.Fields = append(labelEntityProto.Fields, "M_owningShapePtr")
	labelEntityProto.FieldsType = append(labelEntityProto.FieldsType, "DI.LabeledShape:Ref")
	labelEntityProto.FieldsOrder = append(labelEntityProto.FieldsOrder, 11)
	labelEntityProto.FieldsVisibility = append(labelEntityProto.FieldsVisibility, false)
	labelEntityProto.Fields = append(labelEntityProto.Fields, "M_sourceEdgePtr")
	labelEntityProto.FieldsType = append(labelEntityProto.FieldsType, "[]DI.Edge:Ref")
	labelEntityProto.FieldsOrder = append(labelEntityProto.FieldsOrder, 12)
	labelEntityProto.FieldsVisibility = append(labelEntityProto.FieldsVisibility, false)
	labelEntityProto.Fields = append(labelEntityProto.Fields, "M_targetEdgePtr")
	labelEntityProto.FieldsType = append(labelEntityProto.FieldsType, "[]DI.Edge:Ref")
	labelEntityProto.FieldsOrder = append(labelEntityProto.FieldsOrder, 13)
	labelEntityProto.FieldsVisibility = append(labelEntityProto.FieldsVisibility, false)
	labelEntityProto.Fields = append(labelEntityProto.Fields, "M_planePtr")
	labelEntityProto.FieldsType = append(labelEntityProto.FieldsType, "DI.Plane:Ref")
	labelEntityProto.Fields = append(labelEntityProto.Fields, "childsUuid")
	labelEntityProto.FieldsType = append(labelEntityProto.FieldsType, "[]xs.string")
	labelEntityProto.FieldsOrder = append(labelEntityProto.FieldsOrder, 14)
	labelEntityProto.FieldsVisibility = append(labelEntityProto.FieldsVisibility, false)

	labelEntityProto.Fields = append(labelEntityProto.Fields, "referenced")
	labelEntityProto.FieldsType = append(labelEntityProto.FieldsType, "[]EntityRef")
	labelEntityProto.FieldsOrder = append(labelEntityProto.FieldsOrder, 15)
	labelEntityProto.FieldsVisibility = append(labelEntityProto.FieldsVisibility, false)

	store := GetServer().GetDataManager().getDataStore(DIDB).(*KeyValueDataStore)
	store.SetEntityPrototype(&labelEntityProto)

}

/** Entity Prototype creation **/
func (this *EntityManager) create_DI_StyleEntityPrototype() {

	var styleEntityProto EntityPrototype
	styleEntityProto.TypeName = "DI.Style"
	styleEntityProto.IsAbstract = true
	styleEntityProto.SubstitutionGroup = append(styleEntityProto.SubstitutionGroup, "BPMNDI.BPMNLabelStyle")
	styleEntityProto.Ids = append(styleEntityProto.Ids, "uuid")
	styleEntityProto.Fields = append(styleEntityProto.Fields, "uuid")
	styleEntityProto.FieldsType = append(styleEntityProto.FieldsType, "xs.string")
	styleEntityProto.FieldsOrder = append(styleEntityProto.FieldsOrder, 0)
	styleEntityProto.FieldsVisibility = append(styleEntityProto.FieldsVisibility, false)
	styleEntityProto.Indexs = append(styleEntityProto.Indexs, "parentUuid")
	styleEntityProto.Fields = append(styleEntityProto.Fields, "parentUuid")
	styleEntityProto.FieldsType = append(styleEntityProto.FieldsType, "xs.string")
	styleEntityProto.FieldsOrder = append(styleEntityProto.FieldsOrder, 1)
	styleEntityProto.FieldsVisibility = append(styleEntityProto.FieldsVisibility, false)

	/** members of Style **/
	styleEntityProto.Ids = append(styleEntityProto.Ids, "M_id")
	styleEntityProto.FieldsOrder = append(styleEntityProto.FieldsOrder, 2)
	styleEntityProto.FieldsVisibility = append(styleEntityProto.FieldsVisibility, true)
	styleEntityProto.Fields = append(styleEntityProto.Fields, "M_id")
	styleEntityProto.FieldsType = append(styleEntityProto.FieldsType, "xs.ID")

	/** associations of Style **/
	styleEntityProto.FieldsOrder = append(styleEntityProto.FieldsOrder, 3)
	styleEntityProto.FieldsVisibility = append(styleEntityProto.FieldsVisibility, false)
	styleEntityProto.Fields = append(styleEntityProto.Fields, "M_diagramElementPtr")
	styleEntityProto.FieldsType = append(styleEntityProto.FieldsType, "[]DI.DiagramElement:Ref")
	styleEntityProto.FieldsOrder = append(styleEntityProto.FieldsOrder, 4)
	styleEntityProto.FieldsVisibility = append(styleEntityProto.FieldsVisibility, false)
	styleEntityProto.Fields = append(styleEntityProto.Fields, "M_owningDiagramPtr")
	styleEntityProto.FieldsType = append(styleEntityProto.FieldsType, "DI.Diagram:Ref")
	styleEntityProto.Fields = append(styleEntityProto.Fields, "childsUuid")
	styleEntityProto.FieldsType = append(styleEntityProto.FieldsType, "[]xs.string")
	styleEntityProto.FieldsOrder = append(styleEntityProto.FieldsOrder, 5)
	styleEntityProto.FieldsVisibility = append(styleEntityProto.FieldsVisibility, false)

	styleEntityProto.Fields = append(styleEntityProto.Fields, "referenced")
	styleEntityProto.FieldsType = append(styleEntityProto.FieldsType, "[]EntityRef")
	styleEntityProto.FieldsOrder = append(styleEntityProto.FieldsOrder, 6)
	styleEntityProto.FieldsVisibility = append(styleEntityProto.FieldsVisibility, false)

	store := GetServer().GetDataManager().getDataStore(DIDB).(*KeyValueDataStore)
	store.SetEntityPrototype(&styleEntityProto)

}

/** Register the entity to the dynamic typing system. **/
func (this *EntityManager) registerDIObjects() {
}

/** Create entity prototypes contain in a package **/
func (this *EntityManager) createDIPrototypes() {
	this.create_DI_DiagramElementEntityPrototype()
	this.create_DI_NodeEntityPrototype()
	this.create_DI_EdgeEntityPrototype()
	this.create_DI_DiagramEntityPrototype()
	this.create_DI_ShapeEntityPrototype()
	this.create_DI_PlaneEntityPrototype()
	this.create_DI_LabeledEdgeEntityPrototype()
	this.create_DI_LabeledShapeEntityPrototype()
	this.create_DI_LabelEntityPrototype()
	this.create_DI_StyleEntityPrototype()
}
