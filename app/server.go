package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

// Ensures gofmt doesn't remove the "net" and "os" imports above (feel free to remove this!)
var _ = net.Listen
var _ = os.Exit

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}

	conn, err := l.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}

	defer conn.Close()
	input := bufio.NewScanner(conn)
	for input.Scan() {
		var startLine string
		var headers string
		var data string

		request := strings.Split(input.Text(), "\r\n")
		numRequestParts := len(request)
		if numRequestParts > 0 {
			startLine = request[0]
		}
		if numRequestParts > 1 {
			headers = request[1]
		}
		if numRequestParts > 2 {
			data = request[2]
		}

		fmt.Printf("> %s\n", startLine)
		fmt.Printf("> %s\n", headers)
		fmt.Printf("> %s\n", data)
		startLineParts := strings.Fields(startLine)

		method, requestTarget, protocol := startLineParts[0], startLineParts[1], startLineParts[2]
		fmt.Printf("method: %s\n", method)
		fmt.Printf("requestTarget: %s\n", requestTarget)
		fmt.Printf("protocol: %s\n", protocol)
		if requestTarget == "/" {
			_, err := conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\n"))
			if err != nil {
				fmt.Println("Error writing to connection", err)
			}
		} else {
			_, err := conn.Write([]byte("HTTP/1.1 404 Not Found\r\n\r\n"))
			if err != nil {
				fmt.Println("Error writing to connection", err)
			}
		}
	}
}
