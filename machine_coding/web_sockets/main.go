package main

import (
	"fmt"
	"golang.org/x/net/websocket"
	"io"
	"net/http"
)

type Server struct {
	conns map[*websocket.Conn]bool
}

func NewServer() *Server {
	return &Server{
		conns: make(map[*websocket.Conn]bool),
	}
}

func (s *Server) handleChatWS(ws *websocket.Conn) {
	fmt.Println("new incoming connection from client: ", ws.RemoteAddr())

	// todo: maps in golang are not concurrent safe, add mutex to prevent race condition
	s.conns[ws] = true

	// for each connection, we read continuously, so that we can respond back
	s.ReadLoop(ws)
}

func (s *Server) ReadLoop(ws *websocket.Conn) {
	buff := make([]byte, 1024)
	for {
		n, err := ws.Read(buff)
		if err != nil {
			if err == io.EOF {
				break // connection on other side has closed itself, so we can break the while loop
			}
			fmt.Println("read error: ", err)
			continue // clients are allowed to make a malformed message, and we will continue to read the frames from them, without dropping connection
		}

		message := buff[:n]
		fmt.Println(string(message))

		// you may optionally reply to the message
		ws.Write([]byte("thank you for the message!"))
	}
}

func main() {
	server := NewServer()
	http.Handle("/ws/chat", websocket.Handler(server.handleChatWS))
	http.ListenAndServe(":3000", nil)
}
