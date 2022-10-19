package butter

import "reflect"

// Merge merges 2 structs of the same type & returns a map
func MergeStructsToMap(original, incoming struct{}) map[string]interface{} {
	mergedStruct := make(map[string]interface{})
	v := reflect.ValueOf(incoming)
	t := reflect.TypeOf(incoming)

	ov := reflect.ValueOf(original)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i).Interface()
		originalVal := ov.Field(i).Interface()
		if value == nil || value == "" || value == 0 {

			mergedStruct[field.Name] = originalVal
		} else {
			mergedStruct[field.Name] = value
		}
	}

	return mergedStruct
}
