package net

import (
	"syscall"
)

type Listener interface {
	Accept() (Conn, error)
	Close() error
}

type tcpListener struct {
	fd int
}

func (l *tcpListener) Accept() (Conn, error) {
	fd, sa, err := syscall.Accept(l.fd)
	if err != nil {
		return nil, err
	}
	return &tcpConn{fd: fd, sa: sa}, nil
}

func (l *tcpListener) Close() error {
	return syscall.Close(l.fd)
}
