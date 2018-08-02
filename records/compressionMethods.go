package records

type CompressionMethod uint8

const (
	NULL_COMPRESSION CompressionMethod = 0
)

type CompressionMethods struct { // 1 byte of length, followed by up to 2^8-1 bytes of data
	length  uint8
	methods []CompressionMethod
}

func NewCompressionMethods(length uint8, methods []CompressionMethod) CompressionMethods {
	return CompressionMethods{length, methods}
}
