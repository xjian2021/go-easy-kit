package maps

import "reflect"

func MapValToSlice(data interface{}) []interface{} {
	return mapToSlice(data, "value")
}

func MapKeyToSlice(data interface{}) []interface{} {
	return mapToSlice(data, "key")
}

func MapToSlice(data interface{}) []interface{} {
	return mapToSlice(data, "")
}

func mapToSlice(data interface{}, needKV string) []interface{} {
	var (
		reflectValue = reflect.ValueOf(data)
		reflectKind  = reflectValue.Kind()
	)
	// 如果传入的是指针 就一层层地获取它的值
	for reflectKind == reflect.Ptr {
		reflectValue = reflectValue.Elem()
		reflectKind = reflectValue.Kind()
	}
	switch reflectKind {
	case reflect.Map:
		array := make([]interface{}, 0)
		for _, key := range reflectValue.MapKeys() {
			switch needKV {
			case "key":
				array = append(array, key.Interface())
			case "value":
				array = append(array, reflectValue.MapIndex(key).Interface())
			default:
				array = append(array, key.Interface())
				array = append(array, reflectValue.MapIndex(key).Interface())
			}
		}
		return array
	}
	return nil
}
