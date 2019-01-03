package processor

import (
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
)

func getExampleValue(t descriptor.FieldDescriptorProto_Type) (interface{}, error) {
	switch t {
	// floats
	case descriptor.FieldDescriptorProto_TYPE_DOUBLE:
		return float64(1), nil
	case descriptor.FieldDescriptorProto_TYPE_FLOAT:
		return float32(1), nil

	// 64-bits int
	case descriptor.FieldDescriptorProto_TYPE_INT64:
		return int64(1), nil
	case descriptor.FieldDescriptorProto_TYPE_UINT64:
		return uint64(1), nil
	case descriptor.FieldDescriptorProto_TYPE_FIXED64:
		return uint64(1), nil
	case descriptor.FieldDescriptorProto_TYPE_SINT64:
		return int64(1), nil

	// 32 bits int
	case descriptor.FieldDescriptorProto_TYPE_INT32:
		return int32(1), nil
	case descriptor.FieldDescriptorProto_TYPE_UINT32:
		return uint32(1), nil
	case descriptor.FieldDescriptorProto_TYPE_FIXED32:
		return uint32(1), nil
	case descriptor.FieldDescriptorProto_TYPE_SINT32:
		return int32(1), nil

	case descriptor.FieldDescriptorProto_TYPE_BOOL:
		return true, nil
	case descriptor.FieldDescriptorProto_TYPE_STRING:
		return "string", nil
	case descriptor.FieldDescriptorProto_TYPE_BYTES:
		return []byte{1, 1}, nil
	default:
		return nil, ErrNotSupportedType
	}
}

func isBasicType(t descriptor.FieldDescriptorProto_Type) bool {
	if t == descriptor.FieldDescriptorProto_TYPE_ENUM ||
		t == descriptor.FieldDescriptorProto_TYPE_MESSAGE {
		return false
	}
	return true
}

func isMessageType(t descriptor.FieldDescriptorProto_Type) bool {
	return t == descriptor.FieldDescriptorProto_TYPE_MESSAGE
}

func isEnumType(t descriptor.FieldDescriptorProto_Type) bool {
	return t == descriptor.FieldDescriptorProto_TYPE_ENUM
}
