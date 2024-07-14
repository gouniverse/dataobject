package dataobject

import (
	"strings"
	"testing"
)

func TestDataObjectSetAndGet(t *testing.T) {
	user := NewDataObject()
	user.Set("first_name", "Jon")
	user.Set("last_name", "Doe")

	firstName, err := user.Get("first_name")

	if err != nil {
		t.Error(err)
	}

	if firstName != "Jon" {
		t.Error("Expected: Jon, but found:", firstName)
	}

	lastName, err := user.Get("last_name")

	if err != nil {
		t.Error(err)
	}

	if lastName != "Doe" {
		t.Error("Expected: Doe, but found:", lastName)
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

func TestDataObjectSetAndGetTransformers(t *testing.T) {
	user := NewDataObject()
	user.SetTransformer("first_name", &testTransformer{})
	user.Set("first_name", "Jon")
	user.Set("last_name", "Doe")

	firstName, err := user.Get("first_name")

	if err != nil {
		t.Error(err)
	}

	if firstName != "Jon" {
		t.Error("Expected: Jon, but found:", firstName)
	}

	lastName, err := user.Get("last_name")

	if err != nil {
		t.Error(err)
	}

	if lastName != "Doe" {
		t.Error("Expected: Doe, but found:", lastName)
	}
}

type testTransformer struct{}

func (t *testTransformer) Serialize(key string) (string, error) {
	return "serialized_" + key, nil
}

func (t *testTransformer) Deserialize(value string) (string, error) {
	return strings.TrimPrefix(value, "serialized_"), nil
}
