package utils

import (
	"fmt"
	"io"
	. "konn/constants"
	"konn/tel"
	"konn/tel/events"
	"net"
)

func SendEvent(e *events.Event, conn net.Conn) {
	eb := tel.EncodeEvent(e)
	fmt.Println("Committing event.")
	Commiter(eb, conn)
}

// TODO: client
func Commiter(eb *tel.EventBag, conn net.Conn) {
	_, err := conn.Write([]byte(eb.String()))
	if err != nil {
		fmt.Printf("Bag Sending failed: %v\n", err)
	}
	
}

func Receiver(conn net.Conn, out chan string) {
	var buf [512]byte
	for {
		_, err := conn.Read(buf[:])
		if err == io.EOF {
			// TODO: connection out handle
		} else if err != nil {
			LogGenerate(ERROR, "Failed to receive message from server...")
		}
		out <- string(buf[:len(buf)-1])
		LogGenerate(DEBUG, "Message received.")
	}
}

func ClientGenerate() net.Conn {
	conn, err := net.Dial("tcp", Host+":20000")
	if err != nil {
		LogGenerate(FATAL, fmt.Sprintf("Failed to connect to Server: %v", err))
		return nil
	}
	// conn, err :+ net.Dial("tcp", "fmt.Sprintf("%s:%s", HOST, "20000")
	return conn
}
