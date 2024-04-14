package main

import (
	"fmt"
	"time"
)

type Server struct {
	quitch chan struct{} //0 bytes instead of bool which takes 1 or  2bytes
	msgch  chan string
}

func newServer() *Server {
	return &Server{
		quitch: make(chan struct{}),
		msgch:  make(chan string, 128),
	}
}

func (s *Server) start() {
	fmt.Println("starting server...")
	s.loop()
}

func (s *Server) loop() {
mainloop:
	for {
		select {
		case <-s.quitch:
			fmt.Println("quitting server")
			break mainloop
		case msg := <-s.msgch:
			s.handleMessage(msg)
		}
	}
	fmt.Println("server is shutting down...")
}

func (s *Server) quit() {
	s.quitch <- struct{}{}
}

func (s *Server) sendMessage(msg string) {
	s.msgch <- msg
}

func (s *Server) handleMessage(msg string) {
	fmt.Println("we received a message: ", msg)
}

func main() {
	msgServer := newServer()

	go func() {
		msgServer.sendMessage("COULD BE THE LAST MSG ::::")
		time.Sleep(time.Second * 5)
		msgServer.quit()
	}()

	msgServer.start()

}
