package copier

import (
	"bytes"
	"encoding/json"
	"reflect"

	"github.com/gogo/protobuf/jsonpb"

	"github.com/golang/protobuf/proto"
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
	return reflect.ValueOf(copy).Elem().Interface()
}

// CopyPB returns a deep copy of src as stored at dst
func CopyPB(dst proto.Message, src interface{}) interface{} {
	b, err := json.Marshal(src)
	if err != nil {
		panic(err)
	}
	err = jsonpb.Unmarshal(bytes.NewReader(b), dst)
	if err != nil {
		panic(err)
	}
	return dst
}

// CopyPBAndDereference returns the deferenced value of the copy of src as stored at dst
func CopyPBAndDereference(dst proto.Message, src interface{}) interface{} {
	copy := CopyPB(dst, src)
	return reflect.ValueOf(copy).Elem().Interface()
}
