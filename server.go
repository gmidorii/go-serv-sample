package main

import (
	"fmt"
	"log"
	"net"
)

func server() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()
	fmt.Println("Listen")

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
		// 1 request to 1 goroutine
		go func() {
			defer conn.Close()
			var req []byte
			_, err := conn.Read(req)
			if err != nil {
				log.Print(err)
			}
			conn.Write([]byte("Accepted"))
		}()
	}
}
