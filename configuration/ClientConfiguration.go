package configuration

import (
	"math/rand"
	"../ssl"
)

type ClientConfiguration struct {
	SessionID ssl.SessionID
}

func NewClientConfiguration() ClientConfiguration {
	return ClientConfiguration{
		ssl.NewSessionID(rand.Uint32()),
	}
}