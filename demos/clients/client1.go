package clients

import (
	"Util"
	"log"
	"net"
	"os"
)

func Client1() {
	req, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer req.Close()
	Util.MustCopy(os.Stdout, req)
}
