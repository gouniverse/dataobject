package dataobject

import (
	"encoding/json"

	"github.com/gouniverse/uid"
)

// NewDataObject creates a new data object and generates an ID
func NewDataObject() *DataObject {
	o := &DataObject{}
	o.SetID(uid.HumanUid())
	return o
}

// NewDataObjectFromExistingData creates a new data object
// and hydrates it with the passed data
func NewDataObjectFromExistingData(data map[string]string) *DataObject {
	o := &DataObject{}
	o.Hydrate(data)
	return o
}

func NewDataObjectFromJSON(jsonString string) (do *DataObject, err error) {
	var e interface{}

	jsonError := json.Unmarshal([]byte(jsonString), &e)

	if jsonError != nil {
		return do, jsonError
	}

	data := mapStringAnyToMapStringString(e.(map[string]any))

	do = NewDataObjectFromExistingData(data)

	return do, nil
}
