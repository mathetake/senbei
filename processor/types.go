package processor

import (
	"errors"

	"github.com/golang/protobuf/protoc-gen-go/descriptor"
)

var (
	errNotSupportedType = errors.New("not supported type")
)

func getEaxmpleValue(t descriptor.FieldDescriptorProto_Type) (string, error) {
	switch t {
	case descriptor.FieldDescriptorProto_TYPE_DOUBLE:
		return "1", nil
	default:
		return "", errNotSupportedType
	}
}

// returns whether the given type is basic one (which has the example `value`)
func isBaseType(t descriptor.FieldDescriptorProto_Type) bool {
	return true
}
