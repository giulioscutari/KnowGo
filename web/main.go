package main

import (
	"fmt"
	"log"
)

func main() {
	store, err := NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	server := NewServer(":5000", store)
	fmt.Println("server launched on port 5000")
	log.Fatal(server.Run())
}
