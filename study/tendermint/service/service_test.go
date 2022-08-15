package service

import (
	"log"
	"testing"
)

func TestService(t *testing.T) {
	server := NewSocketServer("tcp://127.0.0.1:8080")
	err := server.Start()
	if err != nil {
		log.Fatal(err)
	}
}
