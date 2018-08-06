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

func DeserializeHandshake(buf []byte) (Handshake, int) {
	var bufferPosition = 0;

	msgType, bytesRead := DeserializeHandshakeType(buf[bufferPosition:])

	bufferPosition += bytesRead

	length, bytesRead := DeserializeHandshakeSize(buf[bufferPosition:])

	bufferPosition += bytesRead

	switch (msgType) {
		case CLIENT_HELLO:
			body, bytesRead := DeserializeClientHello(buf[4:4 + length.GetValue()])

			bufferPosition += bytesRead

			return Handshake{
				msgType,
				length,
				body.Serialization,
				NewNestedSerializable([]Serializable{msgType, length, body.Serialization}),
			}, bufferPosition
	}

	return Handshake{}, bufferPosition
}