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
