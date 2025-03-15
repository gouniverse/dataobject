package dataobject

import (
	"testing"
)

// TestDataObjectInterface tests the DataObjectInterface implementation
func TestDataObjectInterface(t *testing.T) {
	// Create a new data object
	do := NewDataObject()
	
	// Test that it implements the interface
	var _ DataObjectInterface = do
	
	// Test ID methods
	t.Run("ID Methods", func(t *testing.T) {
		// Test SetID and ID
		testID := "test-interface-id"
		do.SetID(testID)
		
		if do.ID() != testID {
			t.Errorf("Expected ID to be %s, but got %s", testID, do.ID())
		}
	})
	
	// Test Data method
	t.Run("Data Method", func(t *testing.T) {
		// Set up test data
		do := NewDataObject()
		do.Set("key1", "value1")
		do.Set("key2", "value2")
		
		data := do.Data()
		
		// Check data contains all keys
		if data["key1"] != "value1" || data["key2"] != "value2" {
			t.Errorf("Data() did not return expected values: %v", data)
		}
	})
	
	// Test DataChanged method
	t.Run("DataChanged Method", func(t *testing.T) {
		// Set up test data
		do := NewDataObject()
		do.MarkAsNotDirty() // Clear initial ID setting
		
		// Set some values
		do.Set("key1", "value1")
		do.Set("key2", "value2")
		
		dataChanged := do.DataChanged()
		
		// Check dataChanged contains all keys that were set
		if dataChanged["key1"] != "value1" || dataChanged["key2"] != "value2" {
			t.Errorf("DataChanged() did not return expected values: %v", dataChanged)
		}
		
		// Mark as not dirty and check dataChanged is empty
		do.MarkAsNotDirty()
		dataChanged = do.DataChanged()
		
		if len(dataChanged) != 0 {
			t.Errorf("Expected DataChanged() to be empty after MarkAsNotDirty, but got: %v", dataChanged)
		}
	})
	
	// Test Hydrate method
	t.Run("Hydrate Method", func(t *testing.T) {
		// Set up test data
		do := NewDataObject()
		do.MarkAsNotDirty() // Clear initial ID setting
		
		testData := map[string]string{
			"id":   "hydrate-test-id",
			"key1": "value1",
			"key2": "value2",
		}
		
		// Test that Hydrate sets the data
		do.Hydrate(testData)
		
		// Check data was set correctly
		data := do.Data()
		for k, v := range testData {
			if data[k] != v {
				t.Errorf("Expected data[%s] to be %s after Hydrate, but got %s", k, v, data[k])
			}
		}
		
		// Hydrate doesn't mark the object as dirty, but it also doesn't clear existing dirty flags
		// Since we're using a new object and setting an ID, the object is already dirty
		// We need to check if the dataChanged map is empty after Hydrate
		dataChanged := do.DataChanged()
		
		// The dataChanged map should be empty since we called MarkAsNotDirty before Hydrate
		if len(dataChanged) != 0 {
			t.Errorf("Expected dataChanged to be empty after Hydrate, but got: %v", dataChanged)
		}
	})
}

// TestDataObjectInterfaceEdgeCases tests edge cases for the DataObjectInterface implementation
func TestDataObjectInterfaceEdgeCases(t *testing.T) {
	t.Run("Empty Data", func(t *testing.T) {
		// Test with empty data
		do := NewDataObject()
		do.MarkAsNotDirty() // Clear initial ID setting
		
		emptyData := map[string]string{}
		do.Hydrate(emptyData)
		
		// Check data is empty except for ID
		data := do.Data()
		if len(data) != 0 {
			t.Errorf("Expected Data() to be empty after Hydrate with empty data, but got: %v", data)
		}
	})
	
	t.Run("Nil Data", func(t *testing.T) {
		// Test with nil data
		do := &DataObject{}
		
		// This should not panic
		data := do.Data()
		
		// Check data is initialized
		if data == nil {
			t.Error("Expected Data() to return initialized map for nil data, but got nil")
		}
		
		if len(data) != 0 {
			t.Errorf("Expected Data() to be empty for nil data, but got: %v", data)
		}
	})
	
	t.Run("Overwrite Values", func(t *testing.T) {
		// Test overwriting values
		do := NewDataObject()
		do.MarkAsNotDirty() // Clear initial ID setting
		
		// Set initial values
		do.Set("key", "value1")
		
		// Check initial value
		if do.Get("key") != "value1" {
			t.Errorf("Expected Get(\"key\") to be \"value1\", but got %s", do.Get("key"))
		}
		
		// Overwrite value
		do.Set("key", "value2")
		
		// Check updated value
		if do.Get("key") != "value2" {
			t.Errorf("Expected Get(\"key\") to be \"value2\" after overwrite, but got %s", do.Get("key"))
		}
		
		// Check dataChanged contains the updated value
		dataChanged := do.DataChanged()
		if dataChanged["key"] != "value2" {
			t.Errorf("Expected dataChanged[\"key\"] to be \"value2\" after overwrite, but got %s", dataChanged["key"])
		}
	})
	
	t.Run("Get Non-Existent Key", func(t *testing.T) {
		// Test getting a non-existent key
		do := NewDataObject()
		
		// Get should return empty string for non-existent keys
		if do.Get("non-existent") != "" {
			t.Errorf("Expected Get(\"non-existent\") to be empty string, but got %s", do.Get("non-existent"))
		}
	})
}

// TestDataObjectInterfaceWithLargeData tests the DataObjectInterface with large data sets
func TestDataObjectInterfaceWithLargeData(t *testing.T) {
	t.Run("Large Data Set", func(t *testing.T) {
		// Create a large data set
		largeData := make(map[string]string)
		for i := 0; i < 1000; i++ {
			key := "key" + toString(i)
			value := "value" + toString(i)
			largeData[key] = value
		}
		
		// Create data object with large data
		do := NewDataObjectFromExistingData(largeData)
		
		// Check all data was set correctly
		for key, value := range largeData {
			if do.Get(key) != value {
				t.Errorf("Expected Get(%s) to be %s, but got %s", key, value, do.Get(key))
			}
		}
		
		// Check Data() returns all values
		data := do.Data()
		if len(data) != len(largeData) {
			t.Errorf("Expected Data() to have %d entries, but got %d", len(largeData), len(data))
		}
	})
}

// TestDataObjectInterfaceChaining tests method chaining patterns
func TestDataObjectInterfaceChaining(t *testing.T) {
	t.Run("Method Chaining", func(t *testing.T) {
		// Test a typical usage pattern with method chaining
		do := NewDataObject()
		do.SetID("chaining-test")
		do.Set("name", "Test Object")
		do.Set("active", "true")
		
		// Verify all operations worked correctly
		if do.ID() != "chaining-test" {
			t.Errorf("Expected ID to be \"chaining-test\", but got %s", do.ID())
		}
		
		if do.Get("name") != "Test Object" {
			t.Errorf("Expected Get(\"name\") to be \"Test Object\", but got %s", do.Get("name"))
		}
		
		if do.Get("active") != "true" {
			t.Errorf("Expected Get(\"active\") to be \"true\", but got %s", do.Get("active"))
		}
		
		// Verify object is dirty
		if !do.IsDirty() {
			t.Error("Expected object to be dirty after chained operations, but it was not")
		}
		
		// Mark as not dirty and verify
		do.MarkAsNotDirty()
		if do.IsDirty() {
			t.Error("Expected object to not be dirty after MarkAsNotDirty, but it was")
		}
	})
}
