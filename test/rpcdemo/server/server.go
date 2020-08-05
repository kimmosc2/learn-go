package main

import (
	"learn-go/test/rpcdemo"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	rpc.Register(rpcdemo.DemoService{})

	listen, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Printf("accept error:%v", err)
			continue
		}
		go jsonrpc.ServeConn(conn)
	}
}
