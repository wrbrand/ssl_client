package ssl

type ContentType uint8

const (
	CHANGE_CIPHER_SPEC ContentType = 20
	ALERT                          = 21
	HANDSHAKE                      = 22
	APPLICATION_DATA               = 23
)

func (num ContentType) GetSize() int {
	return 1
}

func (num ContentType) Serialize() []byte {
	return []byte{byte(num)}
}

func (num ContentType) SerializeInto(buf []byte) {
	copy(buf[0:num.GetSize()], num.Serialize())
}