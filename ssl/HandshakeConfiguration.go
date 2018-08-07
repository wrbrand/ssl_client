package ssl

type HandshakeConfiguration struct {
	RecordProtocolVersion ProtocolVersion
	HandshakeProtocolVersion  ProtocolVersion
}

func NewHandshakeConfiguration() HandshakeConfiguration {
	return HandshakeConfiguration{
		TLS12,
		TLS12,
	}
}