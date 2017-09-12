// Goroutine project Goroutine.go
package servers

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func Server1() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err) //fmt.Print & os.Exit(1)
	}
	for {
		connect, err := listener.Accept()
		if err != nil {
			log.Panic(err)
			continue
		}
		go handleConnect(connect)
	}
}

func handleConnect(conn net.Conn) {
	defer conn.Close()
	for {
		_, err := io.WriteString(conn, fmt.Sprintf("%v\n", time.Now()))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}

}
