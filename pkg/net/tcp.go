package net

import (
	"syscall"
)

func Listen(port int) (Listener, error) {
	// TCP ソケットを作成 (IPv4でストリーム型)
	socket, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, syscall.IPPROTO_IP)
	if err != nil {
		return nil, err
	}

	// サーバのアドレスを設定
	sa := syscall.SockaddrInet4{
		Port: port,
	}
	// ソケットにアドレスをバインド
	if err := syscall.Bind(socket, &sa); err != nil {
		return nil, err
	}

	// サーバにリッスン
	if err := syscall.Listen(socket, syscall.SOMAXCONN); err != nil {
		return nil, err
	}

	return &tcpListener{fd: socket}, nil

}
