package utils

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

// TODO: client
func clientGenerate() {
	conn, err := net.Dial("tcp", "106.53.70.120:20000")
	// conn, err :+ net.Dial("tcp", "fmt.Sprintf("%s:%s", HOST, "20000")
	if err != nil {
		fmt.Printf("Error connecting to server: %v\n", err)
		return
	}
	fmt.Println("Connection established.")
	defer conn.Close()

	inputReader := bufio.NewReader(os.Stdin)

	for {
		input, _ := inputReader.ReadString('\n')
		inputInfo := strings.Trim(input, "\r\n")

		_, err := conn.Write([]byte(inputInfo))
		if err != nil {
			fmt.Printf("error.")
		}

		var buf [256]byte

		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Println("recv failed: ", err)
		}
		fmt.Println(string(buf[:n]))
	}
}
