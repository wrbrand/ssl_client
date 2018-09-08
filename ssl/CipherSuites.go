package ssl

import (
	"encoding/binary"
)

type CipherSuites struct {
	// 2 bytes of Length, followed by up to 2^16-1 bytes of data
	length uint16
	suites []CipherSuite
}

func NewCipherSuites(suites []CipherSuite) CipherSuites {
	return CipherSuites{uint16(len(suites) * 2), suites}
}

/*
	Returns the total size in bytes of this struct
*/
func (suites CipherSuites) GetSize() int {
	return 2 + int(len(suites.suites)*2)
}

func (suites CipherSuites) SerializeInto(buf []byte) {
	binary.BigEndian.PutUint16(buf[0:2], suites.length)

	var start int = 2

	for _, suite := range suites.suites {
		var end int = start + 2

		binary.BigEndian.PutUint16(buf[start:end], uint16(suite))

		start = end
	}
}

func (suites CipherSuites) Serialize() []byte {
	obj := make([]byte, suites.GetSize())
	suites.SerializeInto(obj)
	return obj
}

func DeserializeCipherSuites(buf []byte) (CipherSuites, int) {
	var suites []CipherSuite

	bufferPosition := 0
	suitesLength := binary.BigEndian.Uint16(buf[0:2])

	bufferPosition += 2

	for i := suitesLength; i > 0; i -= 2 {
		suites = append(suites, NewCipherSuite(binary.BigEndian.Uint16(buf[bufferPosition:bufferPosition+2])))
		bufferPosition += 2
	}

	return NewCipherSuites(suites), bufferPosition
}

var DefaultCipherSuites = NewCipherSuites([]CipherSuite{
	TLS_AES_128_GCM_SHA256,
	TLS_AES_256_GCM_SHA384,
	TLS_CHACHA20_POLY1305_SHA256,
	TLS_AES_128_CCM_SHA256,
	TLS_AES_128_CCM_8_SHA256,
})