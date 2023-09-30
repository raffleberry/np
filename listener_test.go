package main

import (
	"io"
	"net"
	"testing"
)

func TestListener(t *testing.T) {
	listener, err := net.Listen("tcp", "127.0.0.1:64064")
	t.Logf("Hello")
	if err != nil {
		t.Fatal(err)
	}
	defer listener.Close()
	t.Logf("Listening on %q", listener.Addr())
	for {
		conn, err := listener.Accept()
		if err != nil {
			t.Errorf("error which accepting - %q", conn)
		}
		go func(c net.Conn) {
			t.Logf("Client Connected : %q", c.RemoteAddr().String())
			defer c.Close()
			buf := make([]byte, 0, 4096)
			tmp := make([]byte, 256)
			for {
				n, err := c.Read(tmp)
				if err == io.EOF {
					t.Logf("EOF")
					break
				}
				if err != nil {
					t.Logf("Error reading data from %q", c.RemoteAddr().String())
					break
				}
				buf = append(buf, tmp[:n]...)
				t.Logf("%s", buf)
			}
			t.Logf("Client Sent : %s", buf)
		}(conn)
	}
}
