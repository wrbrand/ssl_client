package ssl3

import (
	"encoding/binary"
)

type CipherSuite uint16

func NewCipherSuite(value uint16) CipherSuite {
	return CipherSuite(value)
}

func (suite CipherSuite) GetValue() uint16 {
	return uint16(suite)
}

const (
	SSL_NULL_WITH_NULL_NULL = 0x0000

	SSL_RSA_WITH_NULL_MD5              = 0x0001
	SSL_RSA_WITH_NULL_SHA              = 0x0002
	SSL_RSA_EXPORT_WITH_RC4_40_MD5     = 0x0003
	SSL_RSA_WITH_RC4_128_MD5           = 0x0004
	SSL_RSA_WITH_RC4_128_SHA           = 0x0005
	SSL_RSA_EXPORT_WITH_RC2_CBC_40_MD5 = 0x0006
	SSL_RSA_WITH_IDEA_CBC_SHA          = 0x0007
	SSL_RSA_EXPORT_WITH_DES40_CBC_SHA  = 0x0008
	SSL_RSA_WITH_DES_CBC_SHA           = 0x0009
	SSL_RSA_WITH_3DES_EDE_CBC_SHA      = 0x000A

	SSL_DH_DSS_EXPORT_WITH_DES40_CBC_SDA = 0x000B
	SSL_DH_DSS_WITH_DES_CBC_SDA          = 0x000C
	SSL_DH_DSS_WITH_3DES_EDE_CBC_SDA     = 0x000D
	SSL_DH_RSA_EXPORT_WITH_DES40_CBC_SHA = 0x000E
	SSL_DH_RSA_WITH_DES_CBC_SHA          = 0x000F
	SSL_DH_RSA_WITH_3DES_EDE_CBC_SHA     = 0x0010

	SSL_DHE_DSS_EXPORT_WITH_DES40_CBC_SHA = 0x0011
	SSL_DHE_DSS_WITH_DES_CBC_SHA          = 0x0012
	SSL_DHE_DSS_WITH_3DES_EDE_CBC_SHA     = 0x0013
	SSL_DHE_RSA_EXPORT_WITH_DES40_CBC_SHA = 0x0014
	SSL_DHE_RSA_WITH_DES_CBC_SHA          = 0x0015
	SSL_DHE_RSA_WITH_3DES_EDE_CBC_SHA     = 0x0016

	SSL_DH_anon_EXPORT_WITH_RC4_40_MD5    = 0x0017
	SSL_DH_anon_WITH_RC4_128_MD5          = 0x0018
	SSL_DH_anon_EXPORT_WITH_DES40_CBC_SHA = 0x0019
	SSL_DH_anon_WITH_DES_CBC_SHA          = 0x001A
	SSL_DH_anon_WITH_3DES_EDE_CBC_SHA     = 0x001B

	SSL_FORTEZZA_KEA_WITH_NULL_SHA         = 0x001C
	SSL_FORTEZZA_KEA_WITH_FORTEZZA_CBC_SHA = 0x001D
	SSL_FORTEZZA_KEA_WITH_RC4_128_SHA      = 0x001E
)

type CipherSuites struct { // 2 bytes of Length, followed by up to 2^16-1 bytes of data
	length uint16
	suites []CipherSuite
}

func NewCipherSuites(suites []CipherSuite) CipherSuites {
	return CipherSuites{uint16(len(suites)), suites}
}

/*
	Returns the total size in bytes of this struct
*/
func (suites CipherSuites) GetSize() int {
	return 2 + int(suites.length*2)
}

func (suites CipherSuites) SerializeInto(buf []byte) {
	binary.BigEndian.PutUint16(buf[0:2], suites.length)

	for index, suite := range suites.suites {
		var start int = (index + 1) * 2
		var end int = (index + 2) * 2

		binary.BigEndian.PutUint16(buf[start:end], uint16(suite))
	}
}
