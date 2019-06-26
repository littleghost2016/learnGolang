package main

import (
	"fmt"
	"net"
	"os"
)

type Connection struct {
	conn net.Conn
}

func checkError(err error) {
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		os.Exit(1)
	}
}

func recvConnMsg(newConnection Connection) {

	//  var buf [50]byte
	buf := make([]byte, 50)

	defer newConnection.conn.Close()

	for {
		// n为接收到的长度
		n, err := newConnection.conn.Read(buf)
		// checkError(err)

		if err != nil {
			fmt.Println("conn closed")
			return
		}
		newConnection.conn.Write([]byte("answer"))
		// fmt.Printf("buf[0:2]: %s\n", buf[0:2])
		fmt.Printf("n: %d\n", n)
		//fmt.Println("recv msg:", buf[0:n])

	}
}

func main() {
	listen_sock, err := net.Listen("tcp", "localhost:10000")
	checkError(err)
	defer listen_sock.Close()

	for {
		conn, err := listen_sock.Accept()
		newConnection := Connection{conn}
		if err != nil {
			continue
		}

		go recvConnMsg(newConnection)
	}

}
