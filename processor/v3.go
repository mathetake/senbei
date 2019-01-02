package processor

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/Jeffail/gabs"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
)

type snippet struct {
	service string
	method  string
	data    *gabs.Container
}

func (s *snippet) compile() string {
	return fmt.Sprintf("grpc_cli call %s%s %s.%s --json_input '%s'\n",
		Address, Port, s.service, s.method, s.data.String(),
	)
}

type V3Processor struct {
	files    map[string]*descriptor.FileDescriptorProto
	messages map[string]*descriptor.DescriptorProto
	output   []snippet
	req      *plugin.CodeGeneratorRequest
}

var _ Processor = &V3Processor{}

func (p *V3Processor) ParseReq() error {
	buf, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return err
	}

	p.req = &plugin.CodeGeneratorRequest{}
	if err = proto.Unmarshal(buf, p.req); err != nil {
		return err
	}
	return nil
}

func (p *V3Processor) EmitResp() error {
	var out string
	for _, s := range p.output {
		out += s.compile()
	}

	resp := &plugin.CodeGeneratorResponse{
		File: []*plugin.CodeGeneratorResponse_File{
			{
				Name:    proto.String(OutFileName),
				Content: proto.String(out),
			},
		},
	}

	buf, err := proto.Marshal(resp)
	if err != nil {
		return err
	}
	_, err = os.Stdout.Write(buf)
	return err
}

func (p *V3Processor) ProcessReq() error {
	p.files = make(map[string]*descriptor.FileDescriptorProto)
	p.messages = make(map[string]*descriptor.DescriptorProto)

	// store files and message definitions to which we refer later
	for _, f := range p.req.ProtoFile {

		// store files
		p.files[f.GetName()] = f

		// store message definitions
		for _, m := range f.MessageType {
			p.messages["."+f.GetPackage()+"."+m.GetName()] = m
		}
	}

	for _, fn := range p.req.FileToGenerate {
		f := p.files[fn]
		for _, srv := range f.Service {
			for _, m := range srv.Method {

				// process method
				obj, err := p.processMethod(m)
				if err != nil {
					return err
				}

				p.output = append(p.output, snippet{
					service: srv.GetName(),
					method:  m.GetName(),
					data:    obj,
				})
			}
		}
	}
	return nil
}

func (p *V3Processor) processMethod(m *descriptor.MethodDescriptorProto) (*gabs.Container, error) {
	jsonObj := gabs.New()
	jsonObj.Array("senbei_types")
	jsonObj.ArrayAppend("wei", "senbei_types")
	// "senbei_types": ["wei", "soiya"], "max_price": 10}
	jsonObj.Set(10, "max_price")
	return jsonObj, nil
}
