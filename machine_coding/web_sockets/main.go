package main

import (
	"fmt"
	"golang.org/x/net/websocket"
	"io"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type Details struct {
	Ip       string
	IsActive bool
	Name     string
}

type Server struct {
	conns map[*websocket.Conn]*Details
}

func NewServer() *Server {
	return &Server{
		conns: make(map[*websocket.Conn]*Details),
	}
}

func (s *Server) handleRelayDocumentUploadPercentage(ws *websocket.Conn) {
	fmt.Println("new incoming connection from client: ", ws.RemoteAddr())
	perc := 0
	for {
		payload := fmt.Sprintf("percentage uploaded: %d", perc)
		ws.Write([]byte(payload)) // payload write starts happening as soon as websocket connection is opened, if event data is not published on web, data is lost
		time.Sleep(time.Millisecond * 100)
		if perc == 100 {
			break
		}
		perc++
	}
	fmt.Println("closing connection from client: ", ws.RemoteAddr())
}

func (s *Server) handleChatWS(ws *websocket.Conn) {
	fmt.Println("new incoming connection from client: ", ws.RemoteAddr())

	// todo: maps in golang are not concurrent safe, add mutex to prevent race condition
	s.conns[ws] = &Details{
		Ip:       ws.Request().Host,
		IsActive: true, // open connection
		Name:     "Player" + strconv.Itoa(rand.Int()),
	}
	fmt.Println(fmt.Sprintf("%s joined the chat!!", s.conns[ws].Name))

	// for each connection, we read continuously, so that we can respond back
	s.ReadLoop(ws)
}

func (s *Server) ReadLoop(ws *websocket.Conn) {
	buff := make([]byte, 1024)
	for {
		n, err := ws.Read(buff)
		if err != nil {
			if err == io.EOF {
				s.conns[ws].IsActive = false
				fmt.Println(fmt.Sprintf("%s left the chat!!", s.conns[ws].Name))
				break // connection on other side has closed itself, so we can break the while loop
			}
			fmt.Println("read error: ", err)
			continue // clients are allowed to make a malformed message, and we will continue to read the frames from them, without dropping connection
		}

		message := buff[:n]
		draftMessage := fmt.Sprintf("%s: %s", s.conns[ws].Name, string(message))
		fmt.Println(draftMessage)

		// you may optionally reply to the message
		// ws.Write([]byte("thank you for the message!"))

		// broadcast message to all the websocket connections connect to this server
		s.BroadCast([]byte(draftMessage))
	}
}

func (s *Server) BroadCast(b []byte) {
	for ws := range s.conns {
		if s.conns[ws].IsActive == true {
			go func(ws *websocket.Conn) {
				if _, err := ws.Write(b); err != nil {
					fmt.Println("write error: ", err)
				}
			}(ws)
		}
	}
}

func main() {
	server := NewServer()
	http.Handle("/ws/chat", websocket.Handler(server.handleChatWS))
	http.Handle("/ws/document-upload-perc", websocket.Handler(server.handleRelayDocumentUploadPercentage))
	http.ListenAndServe(":3000", nil)
}
