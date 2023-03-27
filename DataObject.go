package dataobject

var _ DataObjectFluentInterface = (*DataObject)(nil) // verify it extends the data object interface

type DataObject struct {
	data        map[string]string
	dataChanged map[string]string
}

func (do *DataObject) ID() string {
	return do.Get("id")
}

func (do *DataObject) SetID(id string) DataObjectFluentInterface {
	do.Set("id", id)
	return do
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

// IsDirty returns if data has been modified
func (do *DataObject) IsDirty() bool {
	do.Init()
	return len(do.dataChanged) > 0
}

// SetData sets the data for the object and marks it as dirty
// see Hydrate for dirtyless assignment
func (do *DataObject) SetData(data map[string]string) DataObjectFluentInterface {
	for k, v := range data {
		do.Set(k, v)
	}
	return do
}

func (do *DataObject) Init() {
	if len(do.data) < 1 {
		do.data = map[string]string{}
	}
	if len(do.dataChanged) < 1 {
		do.dataChanged = map[string]string{}
	}
}

// Set helper setter method
func (do *DataObject) Set(key string, value string) {
	do.Init()
	do.data[key] = value
	do.dataChanged[key] = value
}

// Get helper getter method
func (do *DataObject) Get(key string) string {
	do.Init()
	return do.data[key]
}

// Hybernate sets the data for the object without marking as dirty
func (do *DataObject) Hydrate(data map[string]string) {
	do.Init()
	do.data = data
}
