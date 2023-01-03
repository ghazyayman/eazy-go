package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port", os.Args[0])
		os.Exit(1)
	}

	service := os.Args[1]

	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s", err.Error())
		os.Exit(1)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s", err.Error())
		os.Exit(1)
	}

	fmt.Println("Connected to", service)
	conn.Close()
}

// To use this port scanner, save the code to a file (e.g. go-eazy.go) and build it using go build.
// Then, run the resulting executable and pass it the host and port to scan as a command-line argument

// ex. ./go-eazy localhost:80

//This will attempt to connect to the specified host and port and report whether the connection was successful or not.
