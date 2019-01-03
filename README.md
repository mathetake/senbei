# Senbei [![CircleCI](https://circleci.com/gh/mathetake/senbei.svg?style=svg)](https://circleci.com/gh/mathetake/senbei) [![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat)](LICENSE)


Senbei (煎餅, 🍘) is a protocol buffers' third party plugin for generating [grpc_cli](https://github.com/grpc/grpc/blob/master/doc/command_line_tool.md) snippets


[![Watch the video](https://uploader.xzy.pw/upload/20190103152936_b8ffcb31_6547484662.png)](https://twitter.com/i/videos/1080709959271165952)

## usage

To install the package, run the following:

```bash
git clone git@github.com:mathetake/senbei.git
cd senbei
go build -o $GOPATH/bin/protoc-gen-senbei
```

To obtain snippets, run `protoc` with `protoc-gen-senbei` as an protocol buffers' plugin: 

```$xslt
protoc -I. --senbei_out=. path/to/foo.proto
```

Then you can find `grpc_snippets.txt` file in your current directory which contains generated snippets of grpc_cli for your service specification.

For example, given the following `.proto`, 
```proto
syntax = "proto3";

package senbei.example;
option go_package = "main";


service SenbeiService {
    rpc GetSenbeis (SenbeiRuest) returns (Senbei) {}
    rpc echo(Messages) returns(Messages) {}
}

message Message {
    string message = 1;
}

message Messages {
    repeated Message messages = 1;
}

message SenbeiRuest {
    repeated string senbei_types = 1;
    uint32 max_price = 2;
    enum NestedEnum {
        nestedEnum0 = 0;
        nestedEnum1 = 1;
    }

    repeated NestedEnum repeatedNestedEnum = 3;
    NestedEnum nestedEnum = 4;

    message NestedMessage {
        uint64 nestedMessage1 = 1;
        uint64 nestedMessage2 = 2;
        message NestedNestedMessage {
            uint64 nestedNestedMessage1 = 1;
            uint64 nestedNestedMessage2 = 2;
        }

        NestedNestedMessage nestedNestedMessage = 3;
    }

    NestedMessage nestedMessage = 5;

    double float1 = 6;
    float float2 = 7;

    int32 int32_1 = 8;
    uint32 int32_2 = 9;
    fixed32 int32_3 = 10;
    sint32 int32_4 = 11;


    int64 int64_1 = 12;
    uint64 int64_2 = 13;
    fixed64 int64_3 = 14;
    sint64 int64_4 = 15;

    bool boo1 = 16;
    string str = 17;
    bytes bs = 18;
}

message Senbei {}
```

the generated snippets look like

```
❯❯❯ cat grpc_snippets.txt
[SenbeiService.GetSenbeis]
grpc_cli call localhost:50051 SenbeiService.GetSenbeis --json_input '{
	"boo1": true,
	"bs": "AQE=",
	"float1": 1,
	"float2": 1,
	"int321": 1,
	"int322": 1,
	"int323": 1,
	"int324": 1,
	"int641": 1,
	"int642": 1,
	"int643": 1,
	"int644": 1,
	"maxPrice": 1,
	"nestedEnum": "nestedEnum0",
	"nestedMessage": {
		"nestedMessage1": 1,
		"nestedMessage2": 1,
		"nestedNestedMessage": {
			"nestedNestedMessage1": 1,
			"nestedNestedMessage2": 1
		}
	},
	"repeatedNestedEnum": [
		"nestedEnum0"
	],
	"senbeiTypes": [
		"string",
		"string",
		"string"
	],
	"str": "string"
}'

[SenbeiService.echo]
grpc_cli call localhost:50051 SenbeiService.echo --json_input '{
	"messages": [
		{
			"message": "string"
		},
		{
			"message": "string"
		}
	]
}'
```

## Author

[@mathetake](https://twitter.com/mathetake)


## LICENSE
MIT
