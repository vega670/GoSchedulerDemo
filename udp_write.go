package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

const (
	maxport  = 65535
	numConns = 100
)

func main() {
	for conns, port := 0, 1024; conns < numConns && port < maxport; port++ {
		raddr := &net.UDPAddr{
			Port: port,
			IP:   net.ParseIP("127.0.0.1"),
		}

		conn, err := net.DialUDP("udp", nil, raddr)
		if err != nil {
			fmt.Printf("Error dialing port %d: %s\n", port, err)
			continue
		}
		defer conn.Close()

		go writeConnection(conn)
		conns++

		if conns%100 == 0 {
			fmt.Printf("Opened %d connections...\n", conns)
		}
	}

	fmt.Println("Press enter to continue...")
	bufio.NewReader(os.Stdin).ReadString('\n')
}

func writeConnection(conn *net.UDPConn) {
	bytes := []byte("derping!")
	for {
		conn.Write(bytes)
	}
}
