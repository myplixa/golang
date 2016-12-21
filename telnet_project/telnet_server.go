package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func main() {
	fmt.Println("Launching server...")
	ln, err := net.Listen("tcp", "127.0.0.1:8081")
	if err != nil {
		panic(err)
	}
	defer ln.Close()

	for {
		conn, _ := ln.Accept()

	}
}
