package main

import (
	"./handshake"
	sslrand "./random"
	"fmt"
	"math/rand"
	"net"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	var x Alert = Alert{FATAL, BAD_CERTIFICATE}
	fmt.Print(x)

	var timestamp int32 = int32(time.Now().Unix())

	var random sslrand.Random = sslrand.NewRandom()

	tryHandshake(random, 0)
}

func tryHandshake(random sslrand.Random) {
	conn, err := net.Dial("tcp", "golang.org:80")

	if err != nil {
		fmt.Print(err)
	}

	var clientHello = handshake.NewClientHello(random, timestamp)

}
