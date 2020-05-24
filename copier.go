package copier

import (
	"bytes"
	"encoding/json"
	"reflect"

	"github.com/golang/protobuf/jsonpb"

	"github.com/golang/protobuf/proto"
	"google.golang.org/protobuf/encoding/protojson"
	protov2 "google.golang.org/protobuf/proto"
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
func CopyPB(dst interface{}, src interface{}) interface{} {
	var b []byte
	var err error
	if srcPB, ok := src.(proto.Message); ok {
		var buf bytes.Buffer
		m := &jsonpb.Marshaler{EnumsAsInts: true}
		err = m.Marshal(&buf, srcPB)
		b = buf.Bytes()
	} else if srcPB, ok := src.(protov2.Message); ok {
		mo := protojson.MarshalOptions{UseEnumNumbers: true}
		b, err = mo.Marshal(srcPB)
	} else {
		b, err = json.Marshal(src)
	}
	if err != nil {
		panic(err)
	}
	if dstPB, ok := dst.(proto.Message); ok {
		u := jsonpb.Unmarshaler{}
		u.AllowUnknownFields = true
		err = u.Unmarshal(bytes.NewReader(b), dstPB)
		dst = dstPB
	} else if dstPB, ok := dst.(protov2.Message); ok {
		uo := protojson.UnmarshalOptions{DiscardUnknown: false}
		err = uo.Unmarshal(b, dstPB)
		dst = dstPB
	} else {
		err = json.Unmarshal(b, dst)
	}

	if err != nil {
		panic(err)
	}
	return dst
}

// CopyPBAndDereference returns the deferenced value of the copy of src as stored at dst
func CopyPBAndDereference(dst interface{}, src interface{}) interface{} {
	copy := CopyPB(dst, src)
	return reflect.ValueOf(copy).Elem().Interface()
}
