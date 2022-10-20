package butter

import "reflect"

func isOfTypeNumber(value interface{}) bool {
	switch true {
	case value == int(0):
		return true
	case value == uint(0):
		return true
	case value == int8(0):
		return true
	case value == int8(0):
		return true
	case value == uint8(0):
		return true
	case value == int16(0):
		return true
	case value == uint16(0):
		return true
	case value == int32(0):
		return true
	case value == uint32(0):
		return true
	case value == int64(0):
		return true
	case value == uint64(0):
		return true
	case value == float32(0):
		return true
	case value == float64(64):
		return true
	default:
		return false
	}
}

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
		if value == nil || value == "" {
			mergedStruct[field.Name] = originalVal
		} else if isOfTypeNumber(value) {
			mergedStruct[field.Name] = originalVal
		} else {
			mergedStruct[field.Name] = value
		}
	}

	return mergedStruct
}
