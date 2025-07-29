package main

import (
	"fmt"
	"net"

	"github.com/DanCpp/GoNetwork/sclient"
)

func main() {
	connection, err := net.Dial("tcp4", "127.0.0.1:80")
	if err != nil {
		fmt.Println(err)
		return
	}

	var name string
	fmt.Scanln(&name)

	client := sclient.InitClient(name, connection)
	defer client.DeleteClient()

}
