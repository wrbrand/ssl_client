package ssl

import (
	"encoding/binary"
)

type Extensions struct {
	// 2 bytes of Length, followed by up to 2^16-1 bytes of data
	length uint16
	extensions []Extension
}

func NewExtensions(extensions []Extension) Extensions {
	return Extensions{uint16(len(extensions) * 2), extensions}
}

/*
	Returns the total size in bytes of this struct
*/
func (extensions Extensions) GetSize() int {
	return 2 + int(len(extensions.extensions)*2)
}

func (extensions Extensions) SerializeInto(buf []byte) {
	binary.BigEndian.PutUint16(buf[0:2], extensions.length)

	var start int = 2

	for _, suite := range extensions.extensions {
		var end int = start + 2

		binary.BigEndian.PutUint16(buf[start:end], uint16(suite))

		start = end
	}
}

func (extensions Extensions) Serialize() []byte {
	obj := make([]byte, extensions.GetSize())
	extensions.SerializeInto(obj)
	return obj
}

func DeserializeExtensions(buf []byte) (Extensions, int) {
	var extensions []Extension

	bufferPosition := 0
	extensionsLength := binary.BigEndian.Uint16(buf[0:2])

	bufferPosition += 2

	for i := extensionsLength; i > 0; i -= 2 {
		extensions = append(extensions, NewExtension(binary.BigEndian.Uint16(buf[bufferPosition:bufferPosition+2])))
		bufferPosition += 2
	}

	return NewExtensions(extensions), bufferPosition
}

var DefaultExtensions = NewExtensions([]Extension {
	SERVER_NAME,
	MAX_FRAGMENT_LENGTH,
	STATUS_REQUEST,
	SUPPORTED_GROUPS,
	SIGNATURE_ALGORITHMS,
	USE_SRTP,
	HEARTBEAT,
	APPLICATION_LAYER_PROTOCOL_NEGOTIATION,
	SIGNED_CERTIFICATE_TIMESTAMP,
	CLIENT_CERTIFICATE_TYPE,
	SERVER_CERTIFICATE_TYPE,
	PADDING,
	PRE_SHARED_KEY,
	EARLY_DATA,
	SUPPORTED_VERSIONS,
	COOKIE,
	PSK_KEY_EXCHANGE_MODES,
	CERTIFICATE_AUTHORITIES	,
	OID_FILTERS,
	POST_HANDSHAKE_AUTH,
	SIGNATURE_ALGORITHMS_CERT,
	KEY_SHARE,
})