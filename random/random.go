package random

import (
	"math/rand"
	"time"
)

type Random struct {
	gmt_unix_time uint32
	random_bytes  []byte // 28 bytes long
}

func NewRandom() Random {
	var random = Random{
		gmt_unix_time: uint32(time.Now().Unix()),
		random_bytes:  make([]byte, 28)}

	rand.Read(random.random_bytes)

	return random
}
