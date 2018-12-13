package main

import (
	"syscall"
	"fmt"
	"net"
)

type udpServer struct {

}

func (u *udpServer) run() {
	var (
		err     error
		lsa     *syscall.SockaddrInet4
		buf = make([]byte, 100)
		size int
	)
	//socket
	fd, err := syscall.Socket(syscall.AF_UNIX, syscall.SOCK_DGRAM, 0)
	if err != nil {
		fmt.Println(err)
		syscall.Exit(1)
	}
	defer syscall.Close(fd)
	//bind
	lsa = &syscall.SockaddrInet4{
		Port: 7760,
	}
	copy(lsa.Addr[:], net.ParseIP("127.0.0.1").To4())
	syscall.Bind(fd, lsa)
	////listen
	//if err = syscall.Listen(fd, syscall.SOMAXCONN); err != nil {
	//	fmt.Println(err)
	//	//if err = syscall.Closesocket(fd); err != nil {
	//	//	fmt.Println(err)
	//	//}
	//	syscall.Exit(1)
	//}
	//accept
	//nfd, _, err := syscall.Accept(fd)
	//if err != nil {
	//	fmt.Println(err)
	//	syscall.Exit(1)
	//}
	//read net fd
	size, _, err = syscall.Recvfrom(fd, buf, 0)

	//if size, err = syscall.Read(fd, buf); err != nil {
	//	fmt.Println(err)
	//	syscall.Exit(1)
	//}
	fmt.Println(string(buf), size)
	//syscall.Write(fd, []byte("i received"))
}