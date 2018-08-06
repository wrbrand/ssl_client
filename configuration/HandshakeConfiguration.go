package configuration

import (
	"../ssl"
)

type HandshakeConfiguration struct {
	RecordProtocolVersion ssl.ProtocolVersion
	HandshakeProtocolVersion  ssl.ProtocolVersion
}

func NewHandshakeConfiguration() HandshakeConfiguration {
	return HandshakeConfiguration{
		ssl.ProtocolVersion{ 3, 3},
		ssl.ProtocolVersion{ 3, 3},
	}
}