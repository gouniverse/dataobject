package dataobject

import (
	"encoding/json"
	"errors"

	"github.com/gouniverse/uid"
)

// NewDataObject creates a new data object with a unique ID
//
// Business logic:
// - instantiates a new data object
// - generates an ID, using uid.HumanUid()
//
// Note! The object is marked as dirty, as ID is set
//
// Returns:
// - a new data object
func NewDataObject() *DataObject {
	o := &DataObject{}
	o.SetID(uid.HumanUid())
	return o
}

// NewDataObjectFromExistingData creates a new data object
// and hydrates it with the passed data
//
// Note: the object is marked as not dirty, as it is existing data
//
// Business logic:
// - instantiates a new data object
// - hydrates it with the passed data
//
// Returns:
// - a new data object
func NewDataObjectFromExistingData(data map[string]string) *DataObject {
	o := &DataObject{}
	o.Hydrate(data)
	return o
}

// NewDataObjectFromJSON creates a new data object
// and hydrates it with the passed JSON string.
//
// # The JSON string is expected to be a valid DataObject JSON object
//
// Note: the object is marked as not dirty, as it is existing data
//
// Business logic:
// - instantiates a new data object
// - hydrates it with the passed JSON string
//
// Returns:
// - a new data object
// - an error if any
func NewDataObjectFromJSON(jsonString string) (do *DataObject, err error) {
	if !isDataObjectJSON(jsonString) {
		return nil, errors.New("invalid json: missing id")
	}

	var e any

	jsonError := json.Unmarshal([]byte(jsonString), &e)

	if jsonError != nil {
		return nil, jsonError
	}

	data := mapStringAnyToMapStringString(e.(map[string]any))

	if data == nil {
		return nil, errors.New("invalid data from json")
	}

	do = NewDataObjectFromExistingData(data)

	return do, nil
}
