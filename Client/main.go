package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	conn, err := rpc.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer conn.Close()

	var reply string

	err2 := conn.Call("hello.Say", "张三", &reply)
	if err2 != nil {
		fmt.Println(err2)
		return
	}

	fmt.Printf("reply: %v\n", reply)
}
