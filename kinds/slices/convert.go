package slices

import (
	"github.com/gogf/gf/util/gutil"
)

// SliceToMapWithColumnAsKey converts slice type variable `slice` to `map[interface{}]interface{}`
// The value of specified column use as the key for returned map.
// Eg:
// SliceToMapWithColumnAsKey([{"K1": "v1", "K2": 1}, {"K1": "v2", "K2": 2}], "K1") => {"v1": {"K1": "v1", "K2": 1}, "v2": {"K1": "v2", "K2": 2}}
// SliceToMapWithColumnAsKey([{"K1": "v1", "K2": 1}, {"K1": "v2", "K2": 2}], "K2") => {1: {"K1": "v1", "K2": 1}, 2: {"K1": "v2", "K2": 2}}
func SliceToMapWithColumnAsKey(slice interface{}, key string) map[interface{}]interface{} {
	return gutil.SliceToMapWithColumnAsKey(slice, key)
}

// ArrayColumn
// Eg:
// input:[struct{Name:"a",Age:10},{Name:"b",Age:15}] key:"Name" => ["a","b"]
// input:[map{"k1":"v1","k2":"v2"},map{"k1":"vv1","k2":"vv2"}] key:"k1" => ["v1","vv1"]
func ArrayColumn(input []interface{}, key string) []interface{} {
	keyArr := make([]interface{}, len(input))
	for i, i2 := range input {
		value, ok := gutil.ItemValue(i2, key)
		if ok {
			keyArr[i] = value
		}
	}
	return keyArr
}
