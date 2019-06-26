package main

import (
	"fmt"
	"net"
	"os"
	"os/exec"
	// "time"
)

// type Connection struct {
// 	conn net.Conn
// }

func checkError(err error) {
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
		os.Exit(1)
	}
}

// func (c *Connection) sendMessage(msg string) {
// 	c.conn.Write([]byte(msg))
// }

// func (c *Connection) receiveMessage() string {
// 	buf := make([]byte, 50)
// 	// n为接收到的长度
// 	n, err := newConnection.conn.Read(buf)
// 	// checkError(err)

// 	if err != nil {
// 		fmt.Println("conn closed")
// 		return
// 	}
// 	return c.conn.Read(c.conn)
// }

type Connection struct {
	// ConnectionID    int    `json:"ConnectionID"`
	ProcessID       int      `json:"ProcessID"`
	SenderID        string   `json:"SenderID"`
	ReceiverID      string   `json:"ReceiverID"`
	Timestamp       string   `json:"Timestamp"`
	RandomarrayFrom []int    `json:"RandomarrayFrom"`
	RandomarrayTo   []int    `json:"RandomarrayTo"`
	SessionKey      string   `json:"SessionKey"`
	Sign            string   `json:"Sign"`
	Conn            net.Conn `json:"Conn"`
}

func requestConnection(address string) {
	// conn, err := net.Dial("tcp", address)
	// checkError(err)
	// newConnectionObject := Connection{
	// 	1,
	// 	"clinet",
	// 	"server",
	// 	getCurrentTime(),
	// }
	fmt.Printf("address: %s", address)
	cmd := exec.Command("peer", "chaincode invoke -n mycc -c '{\"Args\":[\"getUserInformationByUserID\",\"peer0\"]}' -C myc")
	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("Execute Shell:%s failed with error:%s\n", cmd, err.Error())
	}
	fmt.Printf("Execute Shell:%s finished with output:\n%s\n", cmd, string(output))
}

func main() {
	// conn, err := net.Dial("tcp", "127.0.0.1:10000")

	// newConnection := Connection{conn}
	// checkError(err)
	// defer newConnection.conn.Close()

	// newConnection.sendMessage("01\n")
	// newConnection.receiveMessage()

	// fmt.Println("send msg")
	address := "127.0.0.1:10000"
	requestConnection(address)

}

// func getCurrentTime() string {
// 	return time.Now().Format("20060102030405")
// }
