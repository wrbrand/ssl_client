package main

import (
	"./handshake"
	sslrand "./random"
	"encoding/gob"
	"fmt"
	"math/rand"
	"net"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	var random sslrand.Random = sslrand.NewRandom()

	tryHandshake(random, handshake.NewSessionID(0))
}

func tryHandshake(random sslrand.Random, session_id handshake.SessionID) {
	conn, err := net.Dial("tcp", "1.1.1.1:443")

	if err != nil {
		fmt.Print(err)
	}

	enc := gob.NewEncoder(conn)
	err = enc.Encode(handshake.NewClientHello(random, session_id))

	if err != nil {
		fmt.Print(err)
	}
}
