package dataobject

import "github.com/gouniverse/uid"

// NewDataObject creates a new data object and generates an ID
func NewDataObject() *DataObject {
	o := &DataObject{}
	o.SetID(uid.HumanUid())
	return o
}
