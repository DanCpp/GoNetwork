package main

import (
	"fmt"
	"net"

	"github.com/DanCpp/GoNetwork/sserver"
)

func main() {
	srv, err := net.Listen("tcp4", ":80")
	if err != nil {
		fmt.Println(err)
		return
	}
	server := sserver.NewServer(srv)
	defer server.DeleteServer()

	for {
		_, err := server.AcceptNewClient()
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
}
