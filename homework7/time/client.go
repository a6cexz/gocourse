package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	go func() {
		io.Copy(os.Stdout, conn)
	}()

	for {
		fmt.Println("Enter command:")
		fmt.Print("$>")
		msg := ""
		_, err := fmt.Scanln(&msg)
		if err != nil {
			fmt.Printf("Error scanning command: %v", err)
			return
		}

		fmt.Println("Sending to server: " + msg)
		_, err = fmt.Fprintln(conn, msg)
		if err != nil {
			fmt.Printf("Error sending command: %v", err)
			return
		}
	}
}
