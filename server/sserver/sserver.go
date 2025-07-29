package sserver

import (
	"fmt"
	"net"
)

type server struct {
	connections   map[net.Conn]int
	serv_listener net.Listener
}

func NewServer(srv net.Listener) *server {
	return &server{
		serv_listener: srv,
		connections:   make(map[net.Conn]int),
	}
}

func (srv *server) AddConnection(new_connection net.Conn) {
	srv.connections[new_connection] = 1
}

func (srv *server) DeleteServer() {
	srv.serv_listener.Close()
}

func (srv *server) AcceptNewClient() (net.Conn, error) {
	connection, err := srv.serv_listener.Accept()
	if err != nil {
		return nil, err
	}
	srv.connections[connection] = 1
	go srv.handleConnection(connection)
	return connection, nil
}

func (srv *server) handleConnection(connection net.Conn) {
	defer connection.Close()
	defer delete(srv.connections, connection)
	for {
		input_buffer := make([]byte, 1024)
		n, err := connection.Read(input_buffer)
		if n == 0 || err != nil {
			fmt.Println(err)
			break
		}
		input_string := string(input_buffer[0:n])
		for key, value := range srv.connections {
			if key == connection || value == 0 {
				continue
			}
			key.Write([]byte(input_string))
		}
	}
}
