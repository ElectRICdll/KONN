package utils

import (
	"fmt"
	"konn/tel"
	"net"
)

// TODO: client
func Commiter(eb *tel.EventBag) {
	conn, err := clientGenerate()
	if err != nil {
		fmt.Printf("Bag Sending failed: %v\n", err)
	}
	defer 
	conn.Write([]byte(eb.String()))
}

func clientGenerate() (net.Conn, error) {
	conn, err := net.Dial("tcp", "106.53.70.120:20000")
	// conn, err :+ net.Dial("tcp", "fmt.Sprintf("%s:%s", HOST, "20000")
	return conn, err
}
