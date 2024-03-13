package http

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/yudai2929/http/pkg/net"
)

func ListenAndServe(port int) error {
	server := &Server{
		Port: port,
	}
	return server.ListenAndServe()
}

type Server struct {
	Port int
}

func (s *Server) ListenAndServe() error {
	ln, err := net.Listen(s.Port)
	if err != nil {
		return err
	}

	return s.Serve(ln)
}

func (s *Server) Serve(ln net.Listener) error {
	defer ln.Close()

	for {
		rw, err := ln.Accept()
		if err != nil {
			return err
		}
		// _, err = readRequest(rw)
		// if err != nil {
		// 	return err
		// }

		go handle(rw)
	}
}

func readRequest(rw net.Conn) (*Request, error) {
	buf := bufio.NewReader(rw)
	line, _, err := buf.ReadLine()
	if err != nil {
		return nil, err
	}
	str := strings.Split(string(line), " ")
	fmt.Println(str)
	return &Request{
		Method: str[0],
		URL:    str[1],
	}, nil
}

func handle(conn net.Conn) {
	defer conn.Close()

	request := make([]byte, 1024)
	conn.Read(request)
	fmt.Println(string(request))

	body := []byte("Hello, World!")
	conn.Write([]byte(fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Length: %d\r\n\r\n%s", len(body), body)))
}
