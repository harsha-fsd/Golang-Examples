// a server that serves multiple requests concurrently
/* to test this program run this program using "go run concurrent-server.go"
open different terminal windows and run "nc localhost 9999" cmd
*/
package main

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"math/big"
	"net"
	"strings"
)

var connections = 0

func main() {
	l, err := net.Listen("tcp4", "127.0.0.1:9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go handleConnection(c)
		connections++
	}
}

func handleConnection(c net.Conn) {
	fmt.Println("***New client connected***")
	randomId, _ := rand.Int(rand.Reader, big.NewInt(9999))
	clientId := randomId.Int64() + 1
	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		clientMsg := strings.TrimSpace(string(netData))
		if clientMsg == "STOP" {
			connections--
			break
		}
		fmt.Println("clientId-" + fmt.Sprintf("%d", clientId) + ":" + clientMsg)
		//reply back to the client
		c.Write([]byte("Thanks for the message!"))
	}
	c.Close()
}
