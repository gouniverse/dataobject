package dataobject

import "encoding/json"

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
