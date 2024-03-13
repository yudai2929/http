package net

import "syscall"

type Conn interface {
	Read(b []byte) (n int, err error)
	Write(b []byte) (n int, err error)
	Close() error
}

type tcpConn struct {
	fd int
	sa syscall.Sockaddr
}

func (c *tcpConn) Read(b []byte) (n int, err error) {
	return syscall.Read(c.fd, b)
}

func (c *tcpConn) Write(b []byte) (n int, err error) {
	return syscall.Write(c.fd, b)
}

func (c *tcpConn) Close() error {
	return syscall.Close(c.fd)
}
