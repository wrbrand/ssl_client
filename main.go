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
	initialize()

	config := ssl.LoadConfiguration("config.json")

	//tryHandshake(ssl.NewClientRandom(), config)
	//tryDecodeHandshake(ssl.NewClientRandom(), config)
	tryHandshake(ssl.NewClientRandom(), config)
}

func initialize() {
	rand.Seed(time.Now().UnixNano())
}

func tryHandshake(random ssl.ClientRandom, config ssl.Configuration) {
	helloBody := ssl.NewClientHello(random, config)
	helloHandshake := ssl.NewHandshake(ssl.CLIENT_HELLO, helloBody.GetSerialization())
	helloMessage := ssl.NewSSLPlaintext(ssl.HANDSHAKE, config.Handshake.HandshakeProtocolVersion, helloHandshake.GetSerialization())

	response := getResponse("tcp", "example.com:443", helloMessage.GetSerialization().Serialize())
	fmt.Printf("Response: %x \n", response)

	var message = ssl.DeserializeSSLPlaintext(response)
	fmt.Print(message.Content_type)
	if message.Content_type == ssl.ALERT {
		fmt.Print(ssl.DeserializeAlert(message.GetSerialization().Serialize()))
	}
}

func tryDecodeHandshake(random ssl.ClientRandom, config ssl.Configuration) {
	helloBody := ssl.NewClientHello(random, config)
	helloHandshake := ssl.NewHandshake(ssl.CLIENT_HELLO, helloBody.GetSerialization())
	helloMessage := ssl.NewSSLPlaintext(ssl.HANDSHAKE, config.Handshake.HandshakeProtocolVersion, helloHandshake.GetSerialization())

	tmp := make([]byte, 65536)
	helloMessage.GetSerialization().SerializeInto(tmp)
	var message = ssl.DeserializeSSLPlaintext(tmp)

	fmt.Print("Message deserialized: ", message, "\n")
	fmt.Print("Content type: ", message.Content_type, "\n")
	fmt.Print("Version: ", message.Version, "\n")
	fmt.Print("Length: ", message.Length, "\n")
	fmt.Print("Fragment: ", message.Fragment, "\n")
}

func getResponse(network string, address string, message []byte) []byte {
	timeout := 30 * time.Second

	conn, err := net.Dial(network, address)
	if err != nil {
		panic(err)
	}

	binary.Write(conn, binary.BigEndian, message)

	defer conn.Close()

	response := make([]byte, 0, 65536)
	tmp := make([]byte, 65536)
	conn.SetReadDeadline((time.Now().Add(timeout)))

	n, err := conn.Read(tmp)
	if err != nil {
		panic(err)
	}
	response = append(response, tmp[:n]...)

	defer conn.Close();

	return response
}