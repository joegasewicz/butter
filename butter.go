package butter

import "reflect"

// Merge merges two structs of the same type & returns a map
//
//	 		user1 := user{
//				Name: "joe",
//			}
//			user2 := user{
//				Msg: "hello!",
//			}
//			result := MergeStructsToMap(user1, user2)
func MergeStructsToMap(original, incoming interface{}) map[string]interface{} {
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
