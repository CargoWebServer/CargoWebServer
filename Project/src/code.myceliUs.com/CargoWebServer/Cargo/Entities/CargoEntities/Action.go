// +build CargoEntities

package CargoEntities

import (
	"encoding/xml"
)

type Action struct {

	/** The entity UUID **/
	UUID string
	/** The entity TypeName **/
	TYPENAME string
	/** The parent uuid if there is some. **/
	ParentUuid string
	/** The relation name with the parent. **/
	ParentLnk string
	/** If the entity value has change... **/
	NeedSave bool

	/** If the entity is fully initialyse **/
	IsInit bool

	/** members of Action **/
	M_name       string
	M_doc        string
	M_parameters []*Parameter
	M_results    []*Parameter
	M_accessType AccessType

	/** Associations **/
	m_entitiesPtr *Entities
	/** If the ref is a string and not an object **/
	M_entitiesPtr string
}

/** Xml parser for Action **/
type XsdAction struct {
	XMLName xml.Name `xml:"action"`
	M_name  string   `xml:"name,attr"`
	M_doc   string   `xml:"doc,attr"`
}

/** UUID **/
func (this *Action) GetUUID() string {
	return this.UUID
}

/** Name **/
func (this *Action) GetName() string {
	return this.M_name
}

/** Init reference Name **/
func (this *Action) SetName(ref interface{}) {
	if this.M_name != ref.(string) {
		this.M_name = ref.(string)
		if this.IsInit == true {
			this.NeedSave = true
		}
	}
}

/** Remove reference Name **/

/** Doc **/
func (this *Action) GetDoc() string {
	return this.M_doc
}

/** Init reference Doc **/
func (this *Action) SetDoc(ref interface{}) {
	if this.M_doc != ref.(string) {
		this.M_doc = ref.(string)
		if this.IsInit == true {
			this.NeedSave = true
		}
	}
}

/** Remove reference Doc **/

/** Parameters **/
func (this *Action) GetParameters() []*Parameter {
	return this.M_parameters
}

/** Init reference Parameters **/
func (this *Action) SetParameters(ref interface{}) {
	isExist := false
	var parameterss []*Parameter
	for i := 0; i < len(this.M_parameters); i++ {
		if this.M_parameters[i].GetUUID() != ref.(*Parameter).GetUUID() {
			parameterss = append(parameterss, this.M_parameters[i])
		} else {
			isExist = true
			parameterss = append(parameterss, ref.(*Parameter))
		}
	}
	if !isExist {
		parameterss = append(parameterss, ref.(*Parameter))
		if this.IsInit == true {
			this.NeedSave = true
		}
		this.M_parameters = parameterss
	}
}

/** Remove reference Parameters **/
func (this *Action) RemoveParameters(ref interface{}) {
	toDelete := ref.(*Parameter)
	parameters_ := make([]*Parameter, 0)
	for i := 0; i < len(this.M_parameters); i++ {
		if toDelete.GetUUID() != this.M_parameters[i].GetUUID() {
			parameters_ = append(parameters_, this.M_parameters[i])
		} else {
			this.NeedSave = true
		}
	}
	this.M_parameters = parameters_
}

/** Results **/
func (this *Action) GetResults() []*Parameter {
	return this.M_results
}

/** Init reference Results **/
func (this *Action) SetResults(ref interface{}) {
	isExist := false
	var resultss []*Parameter
	for i := 0; i < len(this.M_results); i++ {
		if this.M_results[i].GetUUID() != ref.(*Parameter).GetUUID() {
			resultss = append(resultss, this.M_results[i])
		} else {
			isExist = true
			resultss = append(resultss, ref.(*Parameter))
		}
	}
	if !isExist {
		resultss = append(resultss, ref.(*Parameter))
		if this.IsInit == true {
			this.NeedSave = true
		}
		this.M_results = resultss
	}
}

/** Remove reference Results **/
func (this *Action) RemoveResults(ref interface{}) {
	toDelete := ref.(*Parameter)
	results_ := make([]*Parameter, 0)
	for i := 0; i < len(this.M_results); i++ {
		if toDelete.GetUUID() != this.M_results[i].GetUUID() {
			results_ = append(results_, this.M_results[i])
		} else {
			this.NeedSave = true
		}
	}
	this.M_results = results_
}

/** AccessType **/
func (this *Action) GetAccessType() AccessType {
	return this.M_accessType
}

/** Init reference AccessType **/
func (this *Action) SetAccessType(ref interface{}) {
	if this.M_accessType != ref.(AccessType) {
		this.M_accessType = ref.(AccessType)
		if this.IsInit == true {
			this.NeedSave = true
		}
	}
}

/** Remove reference AccessType **/

/** Entities **/
func (this *Action) GetEntitiesPtr() *Entities {
	return this.m_entitiesPtr
}

/** Init reference Entities **/
func (this *Action) SetEntitiesPtr(ref interface{}) {
	if _, ok := ref.(string); ok {
		if this.M_entitiesPtr != ref.(string) {
			this.M_entitiesPtr = ref.(string)
			if this.IsInit == true {
				this.NeedSave = true
			}
		}
	} else {
		if this.M_entitiesPtr != ref.(*Entities).GetUUID() {
			this.M_entitiesPtr = ref.(*Entities).GetUUID()
			if this.IsInit == true {
				this.NeedSave = true
			}
		}
		this.m_entitiesPtr = ref.(*Entities)
	}
}

/** Remove reference Entities **/
func (this *Action) RemoveEntitiesPtr(ref interface{}) {
	toDelete := ref.(*Entities)
	if this.m_entitiesPtr != nil {
		if toDelete.GetUUID() == this.m_entitiesPtr.GetUUID() {
			this.m_entitiesPtr = nil
			this.M_entitiesPtr = ""
			this.NeedSave = true
		}
	}
}
