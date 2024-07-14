package dataobject

type TransformerInterface interface {
	Serialize(value string) (string, error)
	Deserialize(value string) (string, error)
}
