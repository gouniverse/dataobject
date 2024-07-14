package dataobject

// DataObjectInterface is an interface for a data object
type DataObjectInterface interface {

	// ID returns the ID of the object
	ID() (string, error)

	// SetID sets the ID of the object
	SetID(id string) error

	// GetData returns the data for the object
	Data() map[string]string

	// GetChangedData returns the data that has been changed since the last hydration
	DataChanged() map[string]string

	// Hydrates the data object with data
	Hydrate(map[string]string)
}
