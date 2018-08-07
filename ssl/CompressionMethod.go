package ssl

type CompressionMethod uint8

func NewCompressionMethod(value uint8) CompressionMethod {
	return CompressionMethod(value)
}