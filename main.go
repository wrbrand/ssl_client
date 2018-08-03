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

	tryHandshake(ssl3.NewClientRandom(), ssl3.NewSessionID(0))
}

func tryHandshake(random ssl3.ClientRandom, session_id ssl3.SessionID) {
	conn, err := net.Dial("tcp", "wrbrand.com:443")

	if err != nil {
		fmt.Print(err)
	}

	helloBody := ssl3.NewClientHello(random, session_id)
	helloHandshake := ssl3.NewHandshake(ssl3.CLIENT_HELLO, helloBody.Serialization)
	helloMessage := ssl3.NewSSLPlainText(ssl3.HANDSHAKE, ssl3.ProtocolVersion { Major: 3, Minor: 0}, helloHandshake.Serialization)

	binary.Write(conn, binary.BigEndian, helloMessage.Serialization.Serialize())
}
