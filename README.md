# Senbei [![CircleCI](https://circleci.com/gh/mathetake/senbei.svg?style=svg)](https://circleci.com/gh/mathetake/senbei) [![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat)](LICENSE)


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
