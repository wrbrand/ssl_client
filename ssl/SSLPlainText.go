package ssl

type SSLPlaintext struct {
	Content_type  ContentType     // The higher level protocol used to process the enclosed Fragment.
	Version       ProtocolVersion // The Version of the protocol being employed
	Length        ContentSize     // The Length in bytes of the following SSLPlaintext.Fragment; should not exceed 2^14
	Fragment      Serializable
}

func (plaintext SSLPlaintext) GetSerialization() NestedSerializable {
	return NewNestedSerializable([]Serializable{
		plaintext.Content_type,
		plaintext.Version,
		plaintext.Length,
		plaintext.Fragment,
	})
}

func NewSSLPlaintext(content_type ContentType, version ProtocolVersion, fragment Serializable) SSLPlaintext {
	length := NewContentSize(fragment.GetSize())

	return SSLPlaintext{
		Content_type:  content_type,
		Version:       version,
		Length:        length,
		Fragment:      fragment,
	}
}

func DeserializeSSLPlaintext(buf []byte) SSLPlaintext {
	var plaintext = SSLPlaintext{}
	var bytesRead int

	deserializers := []func([]byte) (int){
		func(x []byte) (int) { plaintext.Content_type, bytesRead = DeserializeContentType(x); return bytesRead },
		func(x []byte) (int) { plaintext.Version, bytesRead = DeserializeProtocolVersion(x); return bytesRead },
		func(x []byte) (int) { plaintext.Length, bytesRead = DeserializeContentSize(x); return bytesRead },
		func(x []byte) (int) { obj, bytesRead := DeserializeHandshake(x); plaintext.Fragment = obj.GetSerialization(); return bytesRead },
	}

	var bufferPosition = 0;
	for _, deserializer := range deserializers {
		bufferPosition += deserializer(buf[bufferPosition:])
	}

	return plaintext
}