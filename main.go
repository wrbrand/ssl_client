package main

import (
	"./clientRandom"
	"./handshake"
	"encoding/binary"
	"fmt"
	"math/rand"
	"net"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	tryHandshake(clientRandom.NewClientRandom(), handshake.NewSessionID(0x42))
}

func tryHandshake(random clientRandom.ClientRandom, session_id handshake.SessionID) {
	conn, err := net.Dial("tcp", "1.1.1.1:443")

	if err != nil {
		fmt.Print(err)
	}

	hello := handshake.NewClientHello(random, session_id)

	binary.Write(conn, binary.BigEndian, hello.Serialize())
}
