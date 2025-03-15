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

func TestDataObjectInit(t *testing.T) {
	// Test initialization of a new data object
	do := &DataObject{}
	
	// Verify data and dataChanged are nil before initialization
	if do.data != nil {
		t.Error("Expected data to be nil before Init, but it was not nil")
	}
	
	if do.dataChanged != nil {
		t.Error("Expected dataChanged to be nil before Init, but it was not nil")
	}
	
	// Call Init
	do.Init()
	
	// Verify data and dataChanged are initialized
	if do.data == nil {
		t.Error("Expected data to be initialized after Init, but it was nil")
	}
	
	if do.dataChanged == nil {
		t.Error("Expected dataChanged to be initialized after Init, but it was nil")
	}
	
	// Call Init again to ensure it doesn't overwrite existing maps
	do.data["test"] = "value"
	do.dataChanged["test"] = "value"
	
	do.Init()
	
	// Verify data and dataChanged were not reset
	if do.data["test"] != "value" {
		t.Error("Expected data to retain values after second Init, but it did not")
	}
	
	if do.dataChanged["test"] != "value" {
		t.Error("Expected dataChanged to retain values after second Init, but it did not")
	}
}

func TestDataObjectData(t *testing.T) {
	do := NewDataObject()
	do.Set("key1", "value1")
	do.Set("key2", "value2")
	
	data := do.Data()
	
	if len(data) != 3 { // 2 keys + id
		t.Errorf("Expected data to have 3 entries, but found %d", len(data))
	}
	
	if data["key1"] != "value1" {
		t.Errorf("Expected data[\"key1\"] to be \"value1\", but found %s", data["key1"])
	}
	
	if data["key2"] != "value2" {
		t.Errorf("Expected data[\"key2\"] to be \"value2\", but found %s", data["key2"])
	}
}

func TestDataObjectDataChanged(t *testing.T) {
	do := NewDataObject()
	do.Set("key1", "value1")
	do.Set("key2", "value2")
	
	dataChanged := do.DataChanged()
	
	if len(dataChanged) != 3 { // 2 keys + id
		t.Errorf("Expected dataChanged to have 3 entries, but found %d", len(dataChanged))
	}
	
	if dataChanged["key1"] != "value1" {
		t.Errorf("Expected dataChanged[\"key1\"] to be \"value1\", but found %s", dataChanged["key1"])
	}
	
	if dataChanged["key2"] != "value2" {
		t.Errorf("Expected dataChanged[\"key2\"] to be \"value2\", but found %s", dataChanged["key2"])
	}
	
	// Test after MarkAsNotDirty
	do.MarkAsNotDirty()
	dataChanged = do.DataChanged()
	
	if len(dataChanged) != 0 {
		t.Errorf("Expected dataChanged to be empty after MarkAsNotDirty, but found %d entries", len(dataChanged))
	}
}

func TestDataObjectIsDirty(t *testing.T) {
	do := NewDataObject()
	
	// New object with ID should be dirty
	if !do.IsDirty() {
		t.Error("Expected new object to be dirty, but it was not")
	}
	
	// After marking as not dirty, should not be dirty
	do.MarkAsNotDirty()
	if do.IsDirty() {
		t.Error("Expected object to not be dirty after MarkAsNotDirty, but it was")
	}
	
	// After setting a value, should be dirty again
	do.Set("key", "value")
	if !do.IsDirty() {
		t.Error("Expected object to be dirty after Set, but it was not")
	}
}

func TestDataObjectSetData(t *testing.T) {
	do := NewDataObject()
	do.MarkAsNotDirty() // Clear initial ID setting
	
	data := map[string]string{
		"key1": "value1",
		"key2": "value2",
	}
	
	do.SetData(data)
	
	// Check data was set correctly
	if do.Get("key1") != "value1" {
		t.Errorf("Expected do.Get(\"key1\") to be \"value1\", but found %s", do.Get("key1"))
	}
	
	if do.Get("key2") != "value2" {
		t.Errorf("Expected do.Get(\"key2\") to be \"value2\", but found %s", do.Get("key2"))
	}
	
	// Check object is marked as dirty
	if !do.IsDirty() {
		t.Error("Expected object to be dirty after SetData, but it was not")
	}
	
	// Check dataChanged contains the new values
	dataChanged := do.DataChanged()
	if dataChanged["key1"] != "value1" || dataChanged["key2"] != "value2" {
		t.Error("Expected dataChanged to contain the new values, but it did not")
	}
}

func TestDataObjectHydrate(t *testing.T) {
	do := NewDataObject()
	do.MarkAsNotDirty() // Clear initial ID setting
	
	data := map[string]string{
		"key1": "value1",
		"key2": "value2",
	}
	
	do.Hydrate(data)
	
	// Check data was set correctly
	if do.Get("key1") != "value1" {
		t.Errorf("Expected do.Get(\"key1\") to be \"value1\", but found %s", do.Get("key1"))
	}
	
	if do.Get("key2") != "value2" {
		t.Errorf("Expected do.Get(\"key2\") to be \"value2\", but found %s", do.Get("key2"))
	}
	
	// Check object is NOT marked as dirty after Hydrate
	if do.IsDirty() {
		t.Error("Expected object to NOT be dirty after Hydrate, but it was")
	}
}

func TestDataObjectIDMethods(t *testing.T) {
	do := NewDataObject()
	originalID := do.ID()
	
	// Verify ID is not empty
	if originalID == "" {
		t.Error("Expected ID to not be empty, but it was")
	}
	
	// Set a new ID
	newID := "custom-id-123"
	do.SetID(newID)
	
	// Verify ID was updated
	if do.ID() != newID {
		t.Errorf("Expected ID to be %s after SetID, but found %s", newID, do.ID())
	}
}
