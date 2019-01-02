# Senbei 🍘

Senbei (煎餅, 🍘) is a protocol buffers' third party plugin for generating [grpc_cli](https://github.com/grpc/grpc/blob/master/doc/command_line_tool.md) snippets


## installation

```bash
go get -u -d github.com/mathetake/senbei
cd $GOPATH/src/github.com/mathetake/senbei
go build -o $GOPATH/bin/protoc-gen-senbei
```

## usage

```$xslt
protoc -I. --senbei_out=. path/to/foo.proto
```

## Author

[@mathetake](https://twitter.com/mathetake)


## LICENSE
MIT
