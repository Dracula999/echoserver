package main

import (
	"fmt"
	"net"
	"os"
)

// As I'm new to this lang I looked up an example for simple tcp servers. This is the link https://medium.com/go-to-golang/%D1%81%D0%BE%D0%B7%D0%B4%D0%B0%D0%BD%D0%B8%D0%B5-%D0%BF%D1%80%D0%BE%D1%81%D1%82%D0%BE%D0%B3%D0%BE-tcp-%D1%81%D0%B5%D1%80%D0%B2%D0%B5%D1%80%D0%B0-%D0%BD%D0%B0-go-dfbd21957e94
func main() {
	cmdArgs := os.Args
	// check if additional arg is not provided
	if len(cmdArgs) < 3 {
		fmt.Println("Provide port as arg")
		os.Exit(1)
	}
	// expect a second arg to be a free port
	connPort := ":" + cmdArgs[1]
	// Listen for incoming connections.
	l, err := net.Listen("tcp", connPort)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	var conns []net.Conn
	// Close the listener when the application closes.
	defer l.Close()
	fmt.Println("Listening on localhost:" + connPort)
	for {
		// Listen for an incoming connection.
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		// Handle connections in a new goroutine.
		go handleRequest(conn)
	}
}

// Handles incoming requests.
func handleRequest(conn net.Conn) {
	// Make a buffer to hold incoming data.
	buf := make([]byte, 1024)
	// Read the incoming connection into the buffer.
	reqLen, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	fmt.Printf("Received msg: %v, len: %v\n", string(buf), reqLen)
	// Send a response back to person contacting us.
	conn.Write(buf)
	// Close the connection when you're done with it.
	conn.Close()
}
