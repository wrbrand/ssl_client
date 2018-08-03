package ssl

type Handshake struct {
	msg_type      HandshakeType
	length        HandshakeSize // Per spec, this should be a 24-bit uint
	body          Serializable
	Serialization NestedSerializable
}

func NewHandshake(msg_type HandshakeType, body Serializable) Handshake {
	length := NewHandshakeSize(body.GetSize())

	return Handshake{
		msg_type:      msg_type,
		length:        length,
		body:          body,
		Serialization: NewNestedSerializable([]Serializable{msg_type, length, body})}
}

type HandshakeBody struct{}
type HelloRequest struct{ *HandshakeBody }
type ServerHello struct{ *HandshakeBody }
type Certificate struct{ *HandshakeBody }
type ServerKeyExchange struct{ *HandshakeBody }
type CertificateRequest struct{ *HandshakeBody }
type ServerHelloDone struct{ *HandshakeBody }
type CertificateVerify struct{ *HandshakeBody }
type ClientKeyExchange struct{ *HandshakeBody }
type Finished struct{ *HandshakeBody }
