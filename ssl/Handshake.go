package ssl

type Handshake struct {
	msg_type      HandshakeType
	length        HandshakeSize // Per spec, this should be a 24-bit uint
	body          Serializable
}

func NewHandshake(msg_type HandshakeType, body Serializable) Handshake {
	length := NewHandshakeSize(body.GetSize())

	return Handshake{
		msg_type: msg_type,
		length:   length,
		body:     body,
	}
}

func (handshake Handshake) GetSerialization() NestedSerializable {
	return NewNestedSerializable([]Serializable{
		handshake.msg_type,
		handshake.length,
		handshake.body,
	})
}

func DeserializeHandshake(buf []byte) (Handshake, int) {
	var handshake = Handshake{}
	var bytesRead int

	deserializers := []func([]byte) (int){
		func(buf []byte) (int) { handshake.msg_type, bytesRead = DeserializeHandshakeType(buf); return bytesRead },
		func(buf []byte) (int) { handshake.length, bytesRead = DeserializeHandshakeSize(buf); return bytesRead },
		func(buf []byte) (int) {
			switch (handshake.msg_type) {
				case CLIENT_HELLO:
					body, bytesRead := DeserializeClientHello(buf)

					handshake.body = body.GetSerialization()
					return bytesRead
				default:
					return 0
			}
		},
	}

	var bufferPosition = 0;
	for _, deserializer := range deserializers {
		bufferPosition += deserializer(buf[bufferPosition:])
	}

	return handshake, bufferPosition
}