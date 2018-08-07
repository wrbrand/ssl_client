package ssl

type HandshakeType uint8

func (num HandshakeType) GetSize() int {
	return 1
}

func (num HandshakeType) Serialize() []byte {
	return []byte{byte(num)}
}

func (num HandshakeType) SerializeInto(buf []byte) {
	copy(buf[0:num.GetSize()], num.Serialize())
}

func DeserializeHandshakeType(buf []byte) (HandshakeType, int) {
	return HandshakeType(buf[0]), 1
}