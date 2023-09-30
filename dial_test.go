package main

import (
	"bufio"
	"io"
	"net"
	"testing"
	"time"
)

func TestDial(t *testing.T) {
	listener, err := net.Listen("tcp", "127.0.0.1:64064")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Listening on %q", listener.Addr())
	fin := make(chan struct{})
	go func() { // listener routine
		for {
			conn, err := listener.Accept()
			if err != nil {
				return
			}
			go func(conn net.Conn) { // con routine
				defer conn.Close()
				conn.SetReadDeadline(time.Now().Add(time.Second * 3))
				c := bufio.NewReader(conn)
				t.Logf("someone dialed in : %q", conn.RemoteAddr().String())
				buf := make([]byte, 1024)
				for {
					select {
					case <-fin:
						break
					default:
						n, err := c.Read(buf)
						if err != nil {
							if err != io.EOF {
								t.Error(err)
							}
							return
						}
						t.Logf("received: %s", buf[:n])
					}
				}
			}(conn)
		}
	}()
	conn, err := net.Dial("tcp", "127.0.0.1:64064")
	conn.Write([]byte("hello"))
	if err != nil {
		t.Fatal(err)
	}
	conn.Close()
	listener.Close()
}
