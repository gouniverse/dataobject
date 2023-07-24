package dataobject

import "encoding/json"

var _ DataObjectFluentInterface = (*DataObject)(nil) // verify it extends the data object interface

type DataObject struct {
	data        map[string]string
	dataChanged map[string]string
}

// ID returns the ID of the object
func (do *DataObject) ID() string {
	return do.Get("id")
}

// SetID sets the ID of the object
func (do *DataObject) SetID(id string) DataObjectFluentInterface {
	do.Set("id", id)
	return do
}

// Data returns all the data of the object
func (do *DataObject) Data() map[string]string {
	do.Init()
	return do.data
}

// DataChanged returns only the modified data
func (do *DataObject) DataChanged() map[string]string {
	do.Init()
	return do.dataChanged
}

// MarkAsNotDirty marks the object as not dirty
func (do *DataObject) MarkAsNotDirty() {
	do.dataChanged = map[string]string{}
}

// IsDirty returns if data has been modified
func (do *DataObject) IsDirty() bool {
	do.Init()
	return len(do.dataChanged) > 0
}

// SetData sets the data for the object and marks it as dirty
// see Hydrate for assignment without marking as dirty
func (do *DataObject) SetData(data map[string]string) DataObjectFluentInterface {
	for k, v := range data {
		do.Set(k, v)
	}
	return do
}

// Init initializes the data object if it is not already initialized
func (do *DataObject) Init() {
	if len(do.data) < 1 {
		do.data = map[string]string{}
	}
	if len(do.dataChanged) < 1 {
		do.dataChanged = map[string]string{}
	}
}

// Set helper setter method
func (do *DataObject) Set(key string, value string) {
	do.Init()
	do.data[key] = value
	do.dataChanged[key] = value
}

// Get helper getter method
func (do *DataObject) Get(key string) string {
	do.Init()
	return do.data[key]
}

// Hydrate sets the data for the object without marking it as dirty
func (do *DataObject) Hydrate(data map[string]string) {
	do.Init()
	do.data = data
}

// ToJSON converts the DataObject to a JSON string
//
// Returns:
// - the JSON string representation of the DataObject
// - an error if any
func (do *DataObject) ToJSON() (string, error) {
	jsonValue, jsonError := json.Marshal(do.data)
	if jsonError != nil {
		return "", jsonError
	}

	return string(jsonValue), nil
}
