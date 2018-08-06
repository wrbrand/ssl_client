package ssl

type SSLPlaintext struct {
	Content_type  ContentType     // The higher level protocol used to process the enclosed Fragment.
	Version       ProtocolVersion // The Version of the protocol being employed
	Length        ContentSize     // The Length in bytes of the following SSLPlaintext.Fragment; should not exceed 2^14
	Fragment      Serializable
	Serialization NestedSerializable
}

func NewSSLPlaintext(content_type ContentType, version ProtocolVersion, fragment Serializable) SSLPlaintext {
	length := NewContentSize(fragment.GetSize())

	return SSLPlaintext{
		Content_type:  content_type,
		Version:       version,
		Length:        length,
		Fragment:      fragment,
		Serialization: NewNestedSerializable([]Serializable{content_type, version, length, fragment}),
	}
}

func DeserializeSSLPlaintext(buf []byte) SSLPlaintext {
	var bufferPosition = 0;

	contentType, bytesRead := DeserializeContentType(buf[bufferPosition:])

	bufferPosition += bytesRead

	version, bytesRead := DeserializeProtocolVersion(buf[bufferPosition:])

	bufferPosition += bytesRead

	_, bytesRead = DeserializeContentSize(buf[bufferPosition:])

	bufferPosition += bytesRead

	fragment, bytesRead := DeserializeHandshake(buf[bufferPosition:])

	return NewSSLPlaintext(contentType, version, fragment.Serialization)
}