package copier

import (
	"encoding/json"
	"reflect"
)

// Copy returns a deep copy of src as stored at dst
func Copy(dst, src interface{}) interface{} {
	b, err := json.Marshal(src)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(b, dst)
	if err != nil {
		panic(err)
	}
	return dst
}

// CopyAndDereference returns the deferenced value of the copy of src as stored at dst
func CopyAndDereference(dst, src interface{}) interface{} {
	copy := Copy(dst, src)
	return reflect.ValueOf(copy).Elem()
}
