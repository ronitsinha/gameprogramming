package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

var reader *bufio.Reader
var done chan bool

func ReceiveMessages(conn net.Conn) {
	serverReader := bufio.NewReader(conn)
	for {
		message, err := serverReader.ReadString('\n')
		if err != nil {
			fmt.Println("Lost connection with the server")
			done <- true
			return
		}
		fmt.Print(message)
	}
}

func SendMessages(conn net.Conn) {
	for {
		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error " + err.Error())
			done <- true
		}
		if text == "/exit\n" {
			done <- true
			return
		}
		fmt.Fprintf(conn, text)
	}
}

func main() {
	reader = bufio.NewReader(os.Stdin)

	done = make(chan bool)

	if len(os.Args) <= 1 {
		fmt.Println("Please enter an IP address")
		return
	}

	// connect to this socket
	conn, err := net.Dial("tcp", os.Args[1]+":8080")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Print("Enter your name: ")
	name, _ := reader.ReadString('\n')
	fmt.Fprintf(conn, name)

	go SendMessages(conn)
	go ReceiveMessages(conn)

	<-done
}