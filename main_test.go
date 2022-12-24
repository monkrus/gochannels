package main

import (
	"fmt"
	"testing"
)

func TestAddUser(t *testing.T) {
	server := NewServer()
	for i := 0; i < 10; i++ {
		go server.addUser(fmt.Sprintf("user_%d", i))
	}
}
