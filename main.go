package main

import (
	"./ssl3"
	"encoding/binary"
	"fmt"
	"math/rand"
	"net"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	tryHandshake(ssl3.NewClientRandom(), ssl3.NewSessionID(0x42))
}

func tryHandshake(random ssl3.ClientRandom, session_id ssl3.SessionID) {
	conn, err := net.Dial("tcp", "1.1.1.1:443")

	if err != nil {
		fmt.Print(err)
	}

	helloBody := ssl3.NewClientHello(random, session_id)

	hello := ssl3.NewHandshake(ssl3.CLIENT_HELLO, helloBody.Serialization)

	binary.Write(conn, binary.BigEndian, hello.Serialization.Serialize())
}
