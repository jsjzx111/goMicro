package main

import (
	"fmt"
	"net"
	"net/rpc"
)

func (hello Hello) Say(req string, res *string) error {
	*res = "hello， " + req
	return nil
}

type Hello struct {
}

func main() {
	err := rpc.RegisterName("hello", new(Hello))
	if err != nil {
		fmt.Println(err)
	}
	listen, err2 := net.Listen("tcp", "127.0.0.1:8888")
	if err2 != nil {
		fmt.Println(err2)
	}
	defer listen.Close()

	for {
		fmt.Println("等待建立连接...")
		conn, err3 := listen.Accept()
		if err3 != nil {
			fmt.Println(err3)
		}
		rpc.ServeConn(conn)
	}

}
