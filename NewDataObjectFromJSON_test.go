package dataobject

import (
	"testing"
)

func TestNewDataObjectFromJSON(t *testing.T) {
	user := NewDataObject()
	user.Set("first_name", "Jon")
	user.Set("last_name", "Doe")

	json, errJson := user.ToJSON()

	if errJson != nil {
		t.Error("Error must be nil, but found:", errJson.Error())
	}

	do, err := NewDataObjectFromJSON(json)

	if err != nil {
		t.Error("Error must be nil, but found:", err.Error())
	}

	if do == nil {
		t.Error("DataObject must NOT be nil, but found:", nil)
	}

	if do.Get("last_name") != "Doe" {
		t.Error("Expected: Doe, but found:", do.Get("last_name"))
	}
}
