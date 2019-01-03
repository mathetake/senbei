package processor

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/Jeffail/gabs"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
	"github.com/pkg/errors"
)

type snippet struct {
	service string
	method  string
	data    *gabs.Container
}

func (s *snippet) compile() string {
	return fmt.Sprintf("grpc_cli call %s%s %s.%s --json_input '%s'\n",
		Address, Port, s.service, s.method, s.data.StringIndent("", "\t"),
	)
}

type V3Processor struct {
	files    map[string]*descriptor.FileDescriptorProto
	messages map[string]*descriptor.DescriptorProto
	enums    map[string]*descriptor.EnumDescriptorProto
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
		out += fmt.Sprintf("[%s.%s]\n%s\n", s.service, s.method, s.compile())
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
	p.enums = make(map[string]*descriptor.EnumDescriptorProto)

	// store files and message definitions to which we refer later
	for _, f := range p.req.ProtoFile {
		// store file
		p.files[f.GetName()] = f

		prefix := "." + f.GetPackage()
		// store message definitions
		for _, m := range f.MessageType {
			p.storeMessage(prefix, m)
		}

		for _, e := range f.EnumType {
			p.enums[prefix+"."+e.GetName()] = e
		}
	}

	for _, fn := range p.req.FileToGenerate {
		f := p.files[fn]
		for _, srv := range f.Service {
			for _, m := range srv.Method {

				// process method
				inMsg, ok := p.messages[m.GetInputType()]
				if !ok {
					return ErrInvalidMethod
				}

				obj, err := p.getMessageJson(inMsg)
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

func (p *V3Processor) storeMessage(prefix string, m *descriptor.DescriptorProto) {
	key := prefix + "." + m.GetName()
	p.messages[key] = m

	for _, nm := range m.NestedType {
		p.storeMessage(key, nm)
	}

	for _, ne := range m.EnumType {
		p.enums[key+"."+ne.GetName()] = ne
	}
}

func (p *V3Processor) getMessageJson(in *descriptor.DescriptorProto) (*gabs.Container, error) {
	var ret = gabs.New()

	// fields
	for _, f := range in.Field {
		if isBasicType(f.GetType()) {
			v, err := getExampleValue(f.GetType())
			if err != nil {
				return nil, errors.Wrap(err, "getExampleValue failed")
			}

			if f.GetLabel() == descriptor.FieldDescriptorProto_LABEL_REPEATED {
				v = []interface{}{v, v, v}
			}
			ret.Set(v, f.GetJsonName())
			continue
		}

		if isMessageType(f.GetType()) {
			// process message type
			msg, ok := p.messages[f.GetTypeName()]
			if !ok {
				return nil, ErrInvalidTypeName
			}

			// recursive
			obj, err := p.getMessageJson(msg)
			if err != nil {
				return nil, errors.Wrap(err, "getMessageJson failed")
			}

			jsonParsed, err := gabs.ParseJSON([]byte(fmt.Sprintf(
				`{"%s": %s}`, f.GetJsonName(), obj.String())))

			if err != nil {
				return nil, errors.Wrap(err, "gabs.ParseJSON failed")
			}

			err = ret.Merge(jsonParsed)
			if err != nil {
				return nil, errors.Wrap(err, "container.Merge failed")
			}
			continue
		}

		if isEnumType(f.GetType()) {
			// process enum type
			e, ok := p.enums[f.GetTypeName()]
			if !ok {
				return nil, ErrInvalidTypeName
			}

			n := e.Value[0].GetName()
			if f.GetLabel() == descriptor.FieldDescriptorProto_LABEL_REPEATED {
				ret.Set([]string{n}, f.GetJsonName())
			} else {
				ret.Set(n, f.GetJsonName())
			}

			continue
		}

		return nil, ErrInvalidTypeName
	}
	return ret, nil
}
