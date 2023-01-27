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

// create a *TCPAddr
func createTCPadd(target, port string) (*net.TCPAddr, error) {
	return net.ResolveTCPAddr("tcp", net.JoinHostPort(host, port))
}

func main() {

	//Converting host and port to *TCPAddr
	addr, err := createTCPadd(host, port)
	if err != nil {
		panic(err)
	}

	//Connect to a server(TCPconn)
	//local addr is nil 
	tcpConn, err := net.DialTCP("tcp",nil,addr)
	if err != nil {
		panic(err)
	}

	// Write the GET request to connection
	// Note we are closing the HTTP connection with the Connection: close header
	// Fprintf writes to an io.writer
	req := "GET / HTTP/1.1\r\nHost: example.com\r\nConnection: close\r\n\r\n"
	fmt.Fprintf(tcpConn, req)

	// Reading the response

	// Create a scanner
	scanner := bufio.NewScanner(bufio.NewReader(tcpConn))

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
