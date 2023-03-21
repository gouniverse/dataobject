package dataobject

// NewDataObjectFromExistingData creates a new data object
// and hydrates it with the passed data
func NewDataObjectFromExistingData(data map[string]string) *DataObject {
	o := &DataObject{}
	o.Hydrate(data)
	return o
}
