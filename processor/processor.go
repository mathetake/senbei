package processor

import (
	"github.com/pkg/errors"
)

type ProtoVer string

var (
	ErrNotSupported     = errors.New("not supported proto version")
	ErrNotSupportedType = errors.New("not supported type")
)

var (
	OutFileName = "grpc_snippets.txt"
	Address     = "localhost"
	Port        = ":50051"
)

type Processor interface {
	ParseReq() error
	ProcessReq() error
	EmitResp() error
}

var verToProcessor = map[ProtoVer]Processor{
	"v3": &V3Processor{},
}

func GetProcessor(ver ProtoVer) (Processor, error) {
	prc, ok := verToProcessor[ver]
	if !ok {
		return nil, ErrNotSupported
	}

	return prc, nil
}
