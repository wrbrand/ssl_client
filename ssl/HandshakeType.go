package ssl

type HandshakeType uint8

const (
	HELLO_REQUEST       HandshakeType = 0
	CLIENT_HELLO                      = 1
	SERVER_HELLO                      = 2
	CERTIFICATE                       = 11
	SERVER_KEY_EXCHANGE               = 12
	CERTIFICATE_REQUEST               = 13
	SERVER_HELLO_DONE                 = 14
	CERTIFICATE_VERIFY                = 15
	CLIENT_KEY_EXCHANGE               = 16
	FINISHED                          = 20
)

func (num HandshakeType) GetSize() int {
	return 1
}

func (num HandshakeType) Serialize() []byte {
	return []byte{byte(num)}
}

func (num HandshakeType) SerializeInto(buf []byte) {
	copy(buf[0:num.GetSize()], num.Serialize())
}
