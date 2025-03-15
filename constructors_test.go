package dataobject

import (
	"bytes"
	"encoding/gob"
	"testing"
)

func TestNewDataObject(t *testing.T) {
	do := NewDataObject()

	// Verify a new data object has an ID
	if do.ID() == "" {
		t.Error("Expected new DataObject to have an ID, but it was empty")
	}

	// Verify data and dataChanged are initialized
	if do.data == nil {
		t.Error("Expected data to be initialized, but it was nil")
	}

	if do.dataChanged == nil {
		t.Error("Expected dataChanged to be initialized, but it was nil")
	}

	// Verify the ID is in the data map
	if do.data[propertyId] == "" {
		t.Error("Expected data[\"id\"] to be set, but it was empty")
	}

	// Verify the ID is in the dataChanged map
	if do.dataChanged[propertyId] == "" {
		t.Error("Expected dataChanged[\"id\"] to be set, but it was empty")
	}
}

func TestNewDataObjectFromExistingData(t *testing.T) {
	// Create test data
	data := map[string]string{
		"id":         "test-id-123",
		"first_name": "Jane",
		"last_name":  "Smith",
		"age":        "30",
	}

	// Create data object from existing data
	do := NewDataObjectFromExistingData(data)

	// Verify the data object has the correct ID
	if do.ID() != "test-id-123" {
		t.Errorf("Expected ID to be \"test-id-123\", but found %s", do.ID())
	}

	// Verify all data was hydrated correctly
	for key, value := range data {
		if do.Get(key) != value {
			t.Errorf("Expected do.Get(\"%s\") to be \"%s\", but found %s", key, value, do.Get(key))
		}
	}

	// It directly creates a DataObject and hydrates it, which doesn't mark it as dirty
	if do.IsDirty() {
		t.Error("Expected object to NOT be dirty as it is existing data, but it was dirty")
	}
}

func TestNewDataObjectFromJSON(t *testing.T) {
	user := NewDataObject()
	user.Set("first_name", "Jon")
	user.Set("last_name", "Doe")

	json, errJson := user.ToJSON()

	if errJson != nil {
		t.Error("Error must be nil, but found:", errJson.Error())
		return
	}

	do, err := NewDataObjectFromJSON(json)

	if err != nil {
		t.Error("Error must be nil, but found:", err.Error())
		return
	}

	if do == nil {
		t.Error("DataObject must NOT be nil, but found:", nil)
		return
	}

	if do.Get("last_name") != "Doe" {
		t.Error("Expected: Doe, but found:", do.Get("last_name"))
	}

	// It directly creates a DataObject and hydrates it, which doesn't mark it as dirty
	if do.IsDirty() {
		t.Error("Expected object to NOT be dirty as it is existing data, but it was dirty")
	}
}

func TestNewDataObjectFromJSONInvalidJSON(t *testing.T) {
	// Test with invalid JSON
	invalidJSON := "{invalid:json}"

	do, err := NewDataObjectFromJSON(invalidJSON)

	// Verify error is returned
	if err == nil {
		t.Error("Expected error for invalid JSON, but got nil")
	}

	// Verify data object is nil
	if do != nil {
		t.Error("Expected data object to be nil for invalid JSON, but it was not nil")
	}
}

func TestNewDataObjectFromJSONEmptyJSON(t *testing.T) {
	// Test with empty JSON object
	emptyJSON := "{}"

	do, err := NewDataObjectFromJSON(emptyJSON)

	// Verify error is returned for empty JSON
	if err == nil {
		t.Error("Expected error for empty JSON, but got none")
	}

	// Verify the error message is correct
	expectedError := "invalid json: must be a valid dataobject json object"
	if err.Error() != expectedError {
		t.Errorf("Expected error message '%s', but got: '%s'", expectedError, err.Error())
	}

	// Verify data object is nil when error occurs
	if do != nil {
		t.Error("Expected data object to be nil when error occurs, but it was not nil")
	}
}

func TestNewDataObjectFromJSONComplexData(t *testing.T) {
	// Test with complex JSON including numbers, booleans, and nested objects
	complexJSON := `{
		"id": "complex-123",
		"name": "Test Object",
		"active": true,
		"count": 42,
		"price": 19.99,
		"nested": {
			"key1": "value1",
			"key2": "value2"
		},
		"array": [1, 2, 3]
	}`

	do, err := NewDataObjectFromJSON(complexJSON)

	// Verify no error is returned
	if err != nil {
		t.Errorf("Expected no error for complex JSON, but got: %s", err.Error())
	}

	// Verify data object is not nil
	if do == nil {
		t.Error("Expected data object to not be nil for complex JSON, but it was nil")
	}

	// Verify primitive values are converted correctly
	if do.Get("id") != "complex-123" {
		t.Errorf("Expected id to be \"complex-123\", but found %s", do.Get("id"))
	}

	if do.Get("name") != "Test Object" {
		t.Errorf("Expected name to be \"Test Object\", but found %s", do.Get("name"))
	}

	if do.Get("active") != "true" {
		t.Errorf("Expected active to be \"true\", but found %s", do.Get("active"))
	}

	// Numbers are converted to strings with 4 decimal places for floating point values
	if do.Get("count") != "42.0000" {
		t.Errorf("Expected count to be \"42.0000\", but found %s", do.Get("count"))
	}

	// It directly creates a DataObject and hydrates it, which doesn't mark it as dirty
	if do.IsDirty() {
		t.Error("Expected object to NOT be dirty as it is existing data, but it was dirty")
	}
}

func TestNew(t *testing.T) {
	// Create a new data object using the shorter constructor
	do := New()

	// Verify data object is not nil
	if do == nil {
		t.Error("Expected data object to not be nil, but it was nil")
	}

	// Verify ID is set
	if do.ID() == "" {
		t.Error("Expected ID to be set, but it was empty")
	}

	// Verify the object is marked as dirty (since ID was set)
	if !do.IsDirty() {
		t.Error("Expected object to be dirty after setting ID, but it was not dirty")
	}
}

func TestNewFromData(t *testing.T) {
	// Test data
	data := map[string]string{
		"id":    "test-id-123",
		"name":  "Test Object",
		"value": "42",
	}

	// Create a new data object from existing data using the shorter constructor
	do := NewFromData(data)

	// Verify data object is not nil
	if do == nil {
		t.Error("Expected data object to not be nil, but it was nil")
	}

	// Verify ID is set correctly
	if do.ID() != "test-id-123" {
		t.Errorf("Expected ID to be 'test-id-123', but got: %s", do.ID())
	}

	// Verify other data is set correctly
	if do.Get("name") != "Test Object" {
		t.Errorf("Expected name to be 'Test Object', but got: %s", do.Get("name"))
	}

	if do.Get("value") != "42" {
		t.Errorf("Expected value to be '42', but got: %s", do.Get("value"))
	}

	// Verify the object is not marked as dirty (since it's existing data)
	if do.IsDirty() {
		t.Error("Expected object to NOT be dirty for existing data, but it was dirty")
	}
}

func TestNewFromJSON(t *testing.T) {
	// Test with valid JSON object
	validJSON := `{"id":"test-id-456","name":"JSON Object","active":"true"}`

	do, err := NewFromJSON(validJSON)

	// Verify no error is returned
	if err != nil {
		t.Errorf("Expected no error for valid JSON, but got: %s", err.Error())
	}

	// Verify data object is not nil
	if do == nil {
		t.Error("Expected data object to not be nil for valid JSON, but it was nil")
	}

	// Verify ID is set correctly
	if do.ID() != "test-id-456" {
		t.Errorf("Expected ID to be 'test-id-456', but got: %s", do.ID())
	}

	// Verify other data is set correctly
	if do.Get("name") != "JSON Object" {
		t.Errorf("Expected name to be 'JSON Object', but got: %s", do.Get("name"))
	}

	if do.Get("active") != "true" {
		t.Errorf("Expected active to be 'true', but got: %s", do.Get("active"))
	}

	// Verify the object is not marked as dirty (since it's existing data)
	if do.IsDirty() {
		t.Error("Expected object to NOT be dirty for existing data, but it was dirty")
	}
}

func TestNewFromJSONInvalidJSON(t *testing.T) {
	// Test with invalid JSON
	invalidJSON := `{"id":"test-id-789","name":"Invalid JSON",}`

	do, err := NewFromJSON(invalidJSON)

	// Verify error is returned
	if err == nil {
		t.Error("Expected error for invalid JSON, but got none")
	}

	// Verify data object is nil when error occurs
	if do != nil {
		t.Error("Expected data object to be nil when error occurs, but it was not nil")
	}
}

func TestNewFromGob(t *testing.T) {
	// Create test data
	data := map[string]string{
		"id":         "test-gob-123",
		"first_name": "John",
		"last_name":  "Doe",
		"age":        "25",
	}

	// Encode the data to gob
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)
	err := encoder.Encode(data)
	if err != nil {
		t.Errorf("Failed to encode data to gob: %v", err)
		return
	}

	// Create a new data object from the gob data
	do, err := NewFromGob(buf.Bytes())

	// Verify no error is returned
	if err != nil {
		t.Errorf("Expected no error, but got: %s", err.Error())
		return
	}

	// Verify data object is not nil
	if do == nil {
		t.Error("Expected data object to not be nil, but it was nil")
		return
	}

	// Verify the data object has the correct ID
	if do.ID() != "test-gob-123" {
		t.Errorf("Expected ID to be \"test-gob-123\", but found %s", do.ID())
	}

	// Verify all data was hydrated correctly
	for key, value := range data {
		if do.Get(key) != value {
			t.Errorf("Expected do.Get(\"%s\") to be \"%s\", but found %s", key, value, do.Get(key))
		}
	}

	// It directly creates a DataObject and hydrates it, which doesn't mark it as dirty
	if do.IsDirty() {
		t.Error("Expected object to NOT be dirty as it is existing data, but it was dirty")
	}
}

func TestNewFromGobInvalidData(t *testing.T) {
	// Create invalid gob data (not a proper gob encoding of map[string]string)
	invalidGobData := []byte("invalid gob data")

	// Attempt to create a data object from invalid gob data
	do, err := NewFromGob(invalidGobData)

	// Verify error is returned
	if err == nil {
		t.Error("Expected error for invalid gob data, but got nil")
	}

	// Verify data object is nil
	if do != nil {
		t.Error("Expected data object to be nil for invalid gob data, but it was not nil")
	}
}

func TestNewFromGobMissingID(t *testing.T) {
	// Create test data without an ID
	data := map[string]string{
		"first_name": "John",
		"last_name":  "Doe",
		"age":        "25",
	}

	// Encode the data to gob
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)
	err := encoder.Encode(data)
	if err != nil {
		t.Errorf("Failed to encode data to gob: %v", err)
		return
	}

	// Attempt to create a data object from gob data without an ID
	do, err := NewFromGob(buf.Bytes())

	// Verify error is returned
	if err == nil {
		t.Error("Expected error for gob data without ID, but got nil")
	}

	// Verify the error message is correct
	expectedError := "invalid gob data: missing id"
	if err.Error() != expectedError {
		t.Errorf("Expected error message '%s', but got: '%s'", expectedError, err.Error())
	}

	// Verify data object is nil when error occurs
	if do != nil {
		t.Error("Expected data object to be nil when error occurs, but it was not nil")
	}
}

func TestNewFromJSONComplexData(t *testing.T) {
	// Test with complex JSON including numbers, booleans, and nested objects
	complexJSON := `{
		"id": "complex-123",
		"name": "Test Object",
		"active": true,
		"count": 42,
		"price": 19.99,
		"nested": {
			"key1": "value1",
			"key2": "value2"
		},
		"array": [1, 2, 3]
	}`

	do, err := NewDataObjectFromJSON(complexJSON)

	// Verify no error is returned
	if err != nil {
		t.Errorf("Expected no error for complex JSON, but got: %s", err.Error())
	}

	// Verify data object is not nil
	if do == nil {
		t.Error("Expected data object to not be nil for complex JSON, but it was nil")
	}

	// Verify primitive values are converted correctly
	if do.Get("id") != "complex-123" {
		t.Errorf("Expected id to be \"complex-123\", but found %s", do.Get("id"))
	}

	if do.Get("name") != "Test Object" {
		t.Errorf("Expected name to be \"Test Object\", but found %s", do.Get("name"))
	}

	if do.Get("active") != "true" {
		t.Errorf("Expected active to be \"true\", but found %s", do.Get("active"))
	}

	// Numbers are converted to strings with 4 decimal places for floating point values
	if do.Get("count") != "42.0000" {
		t.Errorf("Expected count to be \"42.0000\", but found %s", do.Get("count"))
	}

	// It directly creates a DataObject and hydrates it, which doesn't mark it as dirty
	if do.IsDirty() {
		t.Error("Expected object to NOT be dirty as it is existing data, but it was dirty")
	}
}
