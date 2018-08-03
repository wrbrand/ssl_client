package ssl

type SSLCompressed struct {
	Type     ContentType
	Version  ProtocolVersion
	Length   uint16
	Fragment []byte
}

type Cipher struct {
	Content []byte // of Length SSLCompressed.Length
	MAC     []byte // of Length CipherSpec.hash_size
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
