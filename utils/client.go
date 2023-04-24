package utils

import (
	"fmt"
	. "konn/constants"
	"konn/tel"
	"net"
)

func SendEvent(e tel.Event) {
	eb := tel.EncodeEvent(e)
	fmt.Println("Committing event.")
	Commiter(eb)
}

// TODO: client
func Commiter(eb *tel.EventBag, conn net.Conn) {
	if err != nil {
		fmt.Printf("Bag Sending failed: %v\n", err)
	}
	conn.Write([]byte(eb.String()))
}

func Receiver(conn net.Conn, out string chan) {
	var buf [512]byte
	for {
		_, err := conn.Read(buf[:])
		if err == io.EOF {
			// TODO: connection out handle
		} else if err != nil {
			LogGenerate(ERROR, "Failed to receive message from server...")
		}
		out <- string{buf[:len(buf)-1]}
		LogGenerate(DEBUG, "Message received.")
	}
}

func ClientGenerate() net.Conn {
	conn, err := net.Dial("tcp", Host+":20000")
	if err != nil {
		LogGenerate(FATAL, "Failed to connect to Server: %v", err)
		return nil
	}
	// conn, err :+ net.Dial("tcp", "fmt.Sprintf("%s:%s", HOST, "20000")
	return conn
}
