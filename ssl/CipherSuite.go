package ssl

type CipherSuite uint16

func NewCipherSuite(value uint16) CipherSuite {
	return CipherSuite(value)
}

func (suite CipherSuite) GetValue() uint16 {
	return uint16(suite)
}

const (
	TLS_AES_128_GCM_SHA256 CipherSuite = 0x1301
	TLS_AES_256_GCM_SHA384		  		= 0x1302
	TLS_CHACHA20_POLY1305_SHA256 		= 0x1303
	TLS_AES_128_CCM_SHA256       		= 0x1304
	TLS_AES_128_CCM_8_SHA256     		= 0x1305
)

var SupportedCiphers = map[string]uint16 {
	"TLS_AES_128_GCM_SHA256": 0x1301,
	"TLS_AES_128_GCM_SHA384": 0x1302,
	"TLS_CHACHA20_POLY1305_SHA256": 0x1303,
	"TLS_AES_128_CCM_SHA256": 0x1304,
	"TLS_AES_128_CCM_8_SHA256": 0x1305,
}
