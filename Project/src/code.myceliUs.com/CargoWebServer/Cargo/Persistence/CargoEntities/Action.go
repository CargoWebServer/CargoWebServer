package CargoEntities

type Action struct{

	/** The entity UUID **/
	UUID string
	/** The entity TypeName **/
	TYPENAME string
	/** If the entity value has change... **/
	NeedSave bool

	/** If the entity is fully initialyse **/
	IsInit   bool

	/** members of Entity **/
	M_id string

	/** members of Action **/
	M_name string


	/** Associations **/
	m_entitiesPtr *Entities
	/** If the ref is a string and not an object **/
	M_entitiesPtr string
}/** UUID **/
func (this *Action) GetUUID() string{
	return this.UUID
}

/** Id **/
func (this *Action) GetId() string{
	return this.M_id
}

/** Init reference Id **/
func (this *Action) SetId(ref interface{}){
	this.NeedSave = true
	this.M_id = ref.(string)
}

/** Remove reference Id **/

/** Name **/
func (this *Action) GetName() string{
	return this.M_name
}

/** Init reference Name **/
func (this *Action) SetName(ref interface{}){
	this.NeedSave = true
	this.M_name = ref.(string)
}

/** Remove reference Name **/

/** Entities **/
func (this *Action) GetEntitiesPtr() *Entities{
	return this.m_entitiesPtr
}

/** Init reference Entities **/
func (this *Action) SetEntitiesPtr(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_entitiesPtr = ref.(string)
	}else{
		this.m_entitiesPtr = ref.(*Entities)
		this.M_entitiesPtr = ref.(*Entities).GetUUID()
	}
}

/** Remove reference Entities **/
func (this *Action) RemoveEntitiesPtr(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(*Entities)
	if toDelete.GetUUID() == this.m_entitiesPtr.GetUUID() {
		this.m_entitiesPtr = nil
		this.M_entitiesPtr = ""
	}
}
