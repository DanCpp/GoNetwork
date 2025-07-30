package sclient

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func (cl *client) clearMsgBuffer() {
	cl.message_buffer = ""
}

type client struct {
	name           string
	message_buffer string
	read_buffer    []byte
	connection     net.Conn
	reader         *bufio.Reader
}

func InitClient(name string, connection net.Conn) *client {
	return &client{
		name:       name,
		connection: connection,
		reader:     bufio.NewReader(os.Stdin),
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
		fmt.Println(string(cl.read_buffer[0 : n-1])) //Reading to '\n'
	}
}

func (cl *client) SendMessage() {
	var err error
	cl.message_buffer, err = cl.reader.ReadString('\n')
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
