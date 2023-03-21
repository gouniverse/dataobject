package dataobject

var _ DataObjectInterface = (*DataObject)(nil) // verify it extends the task interface

type DataObject struct {
	data        map[string]string
	dataChanged map[string]string
}

func (do *DataObject) ID() string {
	return do.Get("id")
}

func (do *DataObject) SetID(id string) DataObjectInterface {
	do.Set("id", id)
	return do
}

// Data returns all the data of the object
func (do *DataObject) Data() map[string]string {
	return do.data
}

// DataChanged returns only the modified data
func (do *DataObject) DataChanged() map[string]string {
	return do.dataChanged
}

// IsDirty returns if data has been modified
func (do *DataObject) IsDirty() bool {
	return len(do.dataChanged) > 0
}

// SetData sets the data for the object and marks it as dirty
// see Hydrate for dirtyless assignment
func (do *DataObject) SetData(data map[string]string) DataObjectInterface {
	for k, v := range data {
		do.Set(k, v)
	}
	return do
}

// Set helper setter method
func (do *DataObject) Set(key string, value string) {
	do.data[key] = value
	do.dataChanged[key] = value
}

// Get helper getter method
func (do *DataObject) Get(key string) string {
	return do.data[key]
}

// Hybernate sets the data for the object without marking as dirty
func (do *DataObject) Hydrate(data map[string]string) {
	do.data = data
}
