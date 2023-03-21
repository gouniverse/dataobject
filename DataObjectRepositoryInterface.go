package dataobject

type DataObjectRepositoryInterface interface {
	Create(dataObject DataObjectInterface) error

	Find(id string) (DataObjectInterface, error)

	List() ([]DataObjectInterface, error)

	Update(dataObject DataObjectInterface) error
}
