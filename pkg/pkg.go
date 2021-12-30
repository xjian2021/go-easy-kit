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

//ContainsUpper 字符串中是否有大写字母
func ContainsUpper(str string) bool {
	for _, r := range str {
		//大写字母ascii码值
		if r > 64 && r < 91 {
			return true
		}
	}
	return false
}

//ContainsLower 字符串中是否有小写字母
func ContainsLower(str string) bool {
	for _, r := range str {
		//小写字母ascii码值
		if r > 96 && r < 123 {
			return true
		}
	}
	return false
}