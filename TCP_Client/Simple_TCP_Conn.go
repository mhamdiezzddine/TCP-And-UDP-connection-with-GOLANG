package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
)

var (
	host, port string
)

func init() {
	flag.StringVar(&host, "host", "exemple.com", "target host")
	flag.StringVar(&port, "port", "80", "target port")
}

func main() {

	flag.Parse()

	//convert the format of the address to host:port
	address := net.JoinHostPort(host, port)

	connection, err := net.Dial("tcp", address)
	if err != nil {
		panic(err)
	}

	//create a GET request to the server
	req := "GET / HTTP/1.1\r\nHost: example.com\r\nConnection: close\r\n\r\n"
	fmt.Fprintf(connection, req)

	connReader := bufio.NewReader(connection)

	// Create a scanner
	scanner := bufio.NewScanner(connReader)

	// Combined into one line
	// scanner := bufio.NewScanner(bufio.NewReader(conn))

	// Read from the scanner and print
	// Scanner reads until an I/O error
	for scanner.Scan() {
		fmt.Printf("%s\n", scanner.Text())
	}

	// Check if scanner has quit with an error
	if err := scanner.Err(); err != nil {
		fmt.Println("Scanner error", err)

	}
}
