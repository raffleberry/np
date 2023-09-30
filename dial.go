package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:64064")
	defer conn.Close()
	if err != nil {
		fmt.Printf("Dial error : %q", err)
		return
	}
	fmt.Fprintf(conn, "Hello")
}
