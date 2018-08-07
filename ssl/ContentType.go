package ssl

type ContentType uint8

func (num ContentType) GetSize() int {
	return 1
}

func (num ContentType) Serialize() []byte {
	return []byte{byte(num)}
}

func (num ContentType) SerializeInto(buf []byte) {
	copy(buf[0:num.GetSize()], num.Serialize())
}

func DeserializeContentType(buf []byte) (ContentType, int) {
	return ContentType(buf[0]), 1
}