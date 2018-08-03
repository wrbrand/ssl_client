package main

import (
	"./ssl"
		"fmt"
	"math/rand"
	"net"
	"time"
	"encoding/binary"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	tryHandshake(ssl.NewClientRandom(), ssl.NewSessionID(rand.Uint32()))
}

func tryHandshake(random ssl.ClientRandom, session_id ssl.SessionID) {
	timeout := 30 * time.Second

	conn, err := net.Dial("tcp", "example.com:443")
	if err != nil {
		panic(err)
	}

	helloBody := ssl.NewClientHello(ssl.ProtocolVersion{3, 3}, random, session_id)
	helloHandshake := ssl.NewHandshake(ssl.CLIENT_HELLO, helloBody.Serialization)
	helloMessage := ssl.NewSSLPlainText(ssl.HANDSHAKE, ssl.ProtocolVersion { Major: 3, Minor: 3}, helloHandshake.Serialization)

	binary.Write(conn, binary.BigEndian, helloMessage.Serialization.Serialize())

	defer conn.Close()

	response := make([]byte, 0, 65536)
	tmp := make([]byte, 65536)
	conn.SetReadDeadline((time.Now().Add(timeout)))

	n, err := conn.Read(tmp)
	if err != nil {
		panic(err)
	}
	response = append(response, tmp[:n]...)

	conn.Close();

	fmt.Printf("%x", response)
}
