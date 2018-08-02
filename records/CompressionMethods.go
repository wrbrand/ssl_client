package records

type CompressionMethod uint8

const (
	NULL_COMPRESSION CompressionMethod = 0
)

type CompressionMethods struct { // 1 byte of Length, followed by up to 2^8-1 bytes of data
	length  uint8
	methods []CompressionMethod
}

func NewCompressionMethods(methods []CompressionMethod) CompressionMethods {
	return CompressionMethods{uint8(len(methods)), methods}
}

/*
	Returns the total size in bytes of this struct
*/
func (methods CompressionMethods) GetSize() int {
	return 1 + int(methods.length)
}

func (methods CompressionMethods) SerializeInto(buf []byte) {
	copy(buf[0:1], []byte{methods.length})

	for index, method := range methods.methods {
		var start int = index + 1
		var end int = index + 2

		copy(buf[start:end], []byte{byte(method)})
	}
}
