package maps

// MapDelete deletes all `keys` from map `data`.
func MapDelete(data map[string]interface{}, keys ...string) {
	if len(data) == 0 {
		return
	}
	for _, key := range keys {
		delete(data, key)
	}
}

// MapMerge 合并多个map，返回合并后的map
// 注：若多个map中含有相同的key，则后面map的value会覆盖前面的
func MapMerge(src ...map[string]interface{}) (newMap map[string]interface{}) {
	newMap = make(map[string]interface{})
	for _, m := range src {
		for k, v := range m {
			newMap[k] = v
		}
	}
	return
}

// 从map提取返回float64
func GetMapToFloat64(_data map[string]interface{}, _key string, _defaultVal float64) float64 {
	if _data[_key] == nil {
		return _defaultVal
	}

	result, err := _data[_key].(float64)
	if err == false {
		result = _defaultVal
	}
	return result
}

// 从map提取返回int
func GetMapToInt(_data map[string]interface{}, _key string, _defaultVal int) int {
	result := _defaultVal
	switch _data[_key].(type) {
	case float64:
		resultFloat := GetMapToFloat64(_data, _key, float64(_defaultVal))
		result = int(resultFloat)
	case int:
		result = _data[_key].(int)
	case int32:
		int32Val := _data[_key].(int32)
		result = int(int32Val)
	}

	return result
}

// 从map提取返回string
func GetMapToString(_data map[string]interface{}, _key string, _defaultVal string) string {
	if _data[_key] == nil {
		return _defaultVal
	}
	result, err := _data[_key].(string)
	if err == false {
		result = _defaultVal
	}
	return result
}

// 从map提取返回bool
func GetMapToBool(_data map[string]interface{}, _key string, _defaultVal bool) bool {
	if _data[_key] == nil {
		return _defaultVal
	}
	result, err := _data[_key].(bool)
	if err == false {
		result = _defaultVal
	}
	return result
}

func GetMapToArr(data map[string]interface{}, key string, defaultVal []map[string]interface{}) []map[string]interface{} {
	obj, ok := data[key]
	if !ok {
		return defaultVal
	}

	// TODO 待改进
	var objArr []interface{}
	objArr, ok = obj.([]interface{}) // 断言可能会失败 元素类型并非interface
	if !ok {
		return defaultVal
	}

	var resultArr []map[string]interface{}
	for _, v := range objArr {
		if result, ok := v.(map[string]interface{}); ok {
			resultArr = append(resultArr, result)
		}
	}
	if resultArr == nil {
		return defaultVal
	}
	return resultArr
}
