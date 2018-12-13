package main

import (
	"syscall"
	"fmt"
	"net"
)

type udpClient struct {

}

func (t *udpClient) run() {
	var (
		//buf = make([]byte, 100)
	)
	//socket
	fd, err := syscall.Socket(syscall.AF_UNIX, syscall.SOCK_DGRAM, 0)
	if err != nil {
		fmt.Println(err)
		syscall.Exit(1)
	}
	//connect
	rsa := &syscall.SockaddrInet4{
		Port: 7760,
	}
	copy(rsa.Addr[:], net.ParseIP("127.0.0.1").To4())
	//if err := syscall.Connect(fd, rsa); err != nil {
	//	fmt.Println(err)
	//	syscall.Exit(1)
	//}
	syscall.Sendto(fd, []byte("hello world!"), 0, rsa)
	//syscall.Write(fd, []byte("hello world!"))
	//size, err := syscall.Read(fd, buf)
	//fmt.Println(string(buf), size)
}
