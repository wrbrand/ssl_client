package records

type ProtocolVersion struct {
	Major uint8
	Minor uint8
}

type ContentType uint8

const (
	CHANGE_CIPHER_SPEC ContentType = 20
	ALERT                          = 21
	HANDSHAKE                      = 22
	APPLICATION_DATA               = 23
)

type SSLPlaintext struct {
	Type     ContentType     // The higher level protocol used to process the enclosed fragment.
	Version  ProtocolVersion // The versionof the protocol being employed
	Length   uint16          // The length in bytes of the following SSLPlaintext.Fragment; should not exceed 2^14
	Fragment []byte
}

type SSLCompressed struct {
	Type     ContentType
	Version  ProtocolVersion
	Length   uint16
	Fragment []byte
}

type Cipher struct {
	Content []byte // of length SSLCompressed.length
	MAC     []byte // of length CipherSpec.hash_size
}

type StreamCipher struct {
	*Cipher
}

type BlockCipher struct {
	*Cipher
	Padding       uint8
	PaddingLength uint8
}

type SSLCiphertext struct {
	Type     ContentType
	Version  ProtocolVersion
	Length   uint16
	Fragment Cipher
}

type ChangeCipherSpec struct {
	Type uint8 // Should always be 1
}
