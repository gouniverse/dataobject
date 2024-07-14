package dataobject

import (
	"encoding/json"
	"errors"
)

var _ DataObjectInterface = (*DataObject)(nil) // verify it extends the data object interface

type DataObject struct {
	data         map[string]string
	dataChanged  map[string]string
	transformers map[string]TransformerInterface
}

func (do *DataObject) SetTransformer(key string, transformer TransformerInterface) error {
	_, exists := do.transformers[key]
	if exists {
		return errors.New(`transformer for key "` + key + `" already set`)
	}

	do.transformers[key] = transformer
	return nil
}

func (do *DataObject) GetTransformer(key string) TransformerInterface {
	transformer, exists := do.transformers[key]

	if exists {
		return transformer
	}

	return nil
}

// ID returns the ID of the object
func (do *DataObject) ID() (string, error) {
	return do.Get("id")
}

// SetID sets the ID of the object
func (do *DataObject) SetID(id string) error {
	return do.Set("id", id)
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
func (do *DataObject) SetData(data map[string]string) {
	for k, v := range data {
		do.Set(k, v)
	}
}

// Init initializes the data object if it is not already initialized
func (do *DataObject) Init() {
	if len(do.data) < 1 {
		do.data = map[string]string{}
	}
	if len(do.dataChanged) < 1 {
		do.dataChanged = map[string]string{}
	}
	if len(do.transformers) < 1 {
		do.transformers = map[string]TransformerInterface{}
	}
}

// Set helper setter method
func (do *DataObject) Set(key string, value string) error {
	do.Init()

	transformer := do.GetTransformer(key)
	if transformer != nil {
		valueSerialized, err := transformer.Serialize(value)
		if err != nil {
			return err
		}
		value = valueSerialized
	}

	do.data[key] = value
	do.dataChanged[key] = value

	return nil
}

// Get helper getter method
func (do *DataObject) Get(key string) (string, error) {
	do.Init()

	value := do.data[key]

	transformer := do.GetTransformer(key)
	if transformer != nil {
		valueDeserialized, err := transformer.Deserialize(value)
		if err != nil {
			return "", err
		}
		value = valueDeserialized
	}

	return value, nil
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
