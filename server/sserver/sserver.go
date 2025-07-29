package sserver

import (
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
