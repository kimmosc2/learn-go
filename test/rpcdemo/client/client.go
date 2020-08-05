package main

import (
	"fmt"
	"learn-go/test/rpcdemo"
	"net"
	"net/rpc/jsonrpc"
)

func main() {
	dial, err := net.Dial("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	client := jsonrpc.NewClient(dial)
	var result float64
	err = client.Call("DemoService.Div", rpcdemo.Args{10, 3}, &result)
	fmt.Println(result, err)

}
