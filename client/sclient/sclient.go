package sclient

import (
	"fmt"
	"net"
)

func (cl *client) clearMsgBuffer() {
	cl.message_buffer = ""
}

type client struct {
	name           string
	message_buffer string
	read_buffer    []byte
	connection     net.Conn
}

func InitClient(name string, connection net.Conn) *client {
	return &client{
		name:       name,
		connection: connection,
	}
}

func (cl *client) DeleteClient() {
	cl.connection.Close()
}

func (cl *client) HandleMessage() {
	for {
		cl.read_buffer = make([]byte, 1024)
		n, err := cl.connection.Read(cl.read_buffer)
		if n == 0 || err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(string(cl.read_buffer[0:n]))
	}
}

func (cl *client) SendMessage() {
	_, err := fmt.Scanln(&cl.message_buffer)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer cl.clearMsgBuffer()
	if n, err := cl.connection.Write([]byte(cl.message_buffer)); n == 0 || err != nil {
		fmt.Println(err)
		return
	}
}
