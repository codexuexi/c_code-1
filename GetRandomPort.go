package c_code

import (
	"fmt"
	"net"
)

func GetRandomPort() (port int) {
	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		panic(err)
	}
	port = listener.Addr().(*net.TCPAddr).Port
	fmt.Println("Using port:", listener.Addr().(*net.TCPAddr).Port)
	defer listener.Close()
	return
}
