package main

import (
	"fmt"
	"log"
)

func main() {
	server := NewServer("0.0.0.0:5000")
	fmt.Println("server launched on port 5000")
	log.Fatal(server.Run())
}
