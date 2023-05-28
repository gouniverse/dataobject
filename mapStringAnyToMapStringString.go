package dataobject

// mapStringAnyToMapStringString converts a map[string]any to map[string]string
func mapStringAnyToMapStringString(data map[string]any) map[string]string {
	result := map[string]string{}
	for k, v := range data {
		result[k] = toString(v)
	}
	return result
}
