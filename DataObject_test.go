package dataobject

import (
	"strings"
	"testing"
)

func TestDataObjectSetAndGet(t *testing.T) {
	user := NewDataObject()
	user.Set("first_name", "Jon")
	user.Set("last_name", "Doe")

	if user.Get("first_name") != "Jon" {
		t.Error("Expected: Jon, but found:", user.Get("first_name"))
	}

	if user.Get("last_name") != "Doe" {
		t.Error("Expected: Doe, but found:", user.Get("last_name"))
	}
}

func TestDataObjectToJSON(t *testing.T) {
	user := NewDataObject()
	user.Set("first_name", "Jon")
	user.Set("last_name", "Doe")

	json, err := user.ToJSON()

	if err != nil {
		t.Error("Error must be nil, but found:", err.Error())
	}

	if !strings.Contains(json, `"first_name":"Jon"`) {
		t.Error(`Expected to contain: "first_name":"Jon", but found:`, json)
	}

	if !strings.Contains(json, `"last_name":"Doe"`) {
		t.Error(`Expected to contain: "last_name":"Doe", but found:`, json)
	}
}
