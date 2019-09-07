package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	fmt.Println("New client connected")

	go func() {
		for {
			_, err := io.WriteString(c, time.Now().Format("15:04:05\n\r"))
			if err != nil {
				return
			}
			time.Sleep(1 * time.Second)
		}
	}()

	for {
		fmt.Println("Waiting for command from client")
		msg := ""
		_, err := fmt.Fscanln(c, &msg)
		if err != nil {
			fmt.Printf("Error reading command: %v", err)
			return
		}

		fmt.Println("Command recevied: " + msg)
		if msg == "exit" {
			fmt.Println("Exit")
			return
		}
	}
}
