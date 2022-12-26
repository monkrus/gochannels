package main

import (
//"fmt"
)

type Server struct {
	users  map[string]string
	userch chan string
}

func NewServer() *Server {
	return &Server{
		users:  make(map[string]string),
		userch: make(chan string),
	}
}

func (s *Server) Start() {
	go s.loop()
}

func (s *Server) loop() {
	for {
		user := <-s.userch
		s.users[user] = user
	}
}

func (s *Server) addUser(user string) {
	s.users[user] = user
}

func main() {

}

/*func sendMessage(msgch chan<- string) {
	msgch <- "hello"
	fmt.Println(msgch)

}

func readMessage(msgch <-chan string) {
	msg := <-msgch
	fmt.Println(msg)

}
*/
