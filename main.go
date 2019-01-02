package main

import (
	"fmt"

	"github.com/mathetake/senbei/processor"
)

// TODO: to support other proto versions
const defaultVersion processor.ProtoVer = "v3"

func main() {
	prc, err := processor.GetProcessor(defaultVersion)
	if err != nil {
		panic(fmt.Sprintf("GetProcessor failed %v", err))
	}

	if err = prc.ParseReq(); err != nil {
		panic(fmt.Sprintf("ParseReq failed: %v", err))
	}

	if err = prc.ProcessReq(); err != nil {
		panic(fmt.Sprintf("ProcessReq failed: %v", err))
	}

	if err = prc.EmitResp(); err != nil {
		panic(fmt.Sprintf("EmitResp failed: %v", err))
	}
}
