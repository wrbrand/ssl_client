package main

import (
	"./ssl"
	"encoding/binary"
	"fmt"
	"math/rand"
	"net"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	tryHandshake(ssl.NewClientRandom(), ssl.NewSessionID(rand.Uint32()))
}

func tryHandshake(random ssl.ClientRandom, session_id ssl.SessionID) {
	conn, err := net.Dial("tcp", "example.com:443")

	if err != nil {
		fmt.Print(err)
	}

	helloBody := ssl.NewClientHello(ssl.ProtocolVersion{3, 3}, random, session_id)
	helloHandshake := ssl.NewHandshake(ssl.CLIENT_HELLO, helloBody.Serialization)
	helloMessage := ssl.NewSSLPlainText(ssl.HANDSHAKE, ssl.ProtocolVersion { Major: 3, Minor: 3}, helloHandshake.Serialization)

	binary.Write(conn, binary.BigEndian, helloMessage.Serialization.Serialize())
}
