package pkg

import "reflect"

func KindOfData(data interface{}) reflect.Kind {
	var (
		value = reflect.ValueOf(data)
		kind = value.Kind()
	)
	for kind == reflect.Ptr {
		value = value.Elem()
		kind = value.Kind()
	}
	return kind
}
