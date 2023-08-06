package utils

import (
	"fmt"
	"io"
	"konn/constants"
	"net"
)

type GameServer net.Conn

var toServer GameServer

// TODO: client
func Commiter(str string) {
	_, err := toServer.Write([]byte(str))
	if err != nil {
		fmt.Printf("Bag Sending failed: %v\n", err)
	}
}

func Receiver(out chan string) {
	var buf [512]byte
	for {
		_, err := toServer.Read(buf[:])
		if err == io.EOF {
			// TODO: connection out handle
		} else if err != nil {
			Logger.Error("Failed to receive message from server...")
		}
		out <- string(buf[:len(buf)-1])
		Logger.Debug("Message received.")
	}
}

func ClientGenerate() *GameServer {
	conn, err := net.Dial("tcp", constants.Host+":20000")
	if err != nil {
		Logger.Fatal(fmt.Sprintf("Failed to connect to Server: %v", err))
		return nil
	}
	// conn, err :+ net.Dial("tcp", "fmt.Sprintf("%s:%s", HOST, "20000")
	toServer = conn
	return &toServer
}
