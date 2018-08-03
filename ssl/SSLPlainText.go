package ssl

type SSLPlaintext struct {
	content_type	ContentType     // The higher level protocol used to process the enclosed fragment.
	version  		ProtocolVersion // The version of the protocol being employed
	length   		ContentSize     // The Length in bytes of the following SSLPlaintext.Fragment; should not exceed 2^14
	fragment 		Serializable
	Serialization	NestedSerializable
}

func NewSSLPlainText(content_type ContentType, version ProtocolVersion, fragment Serializable) SSLPlaintext {
	length := NewContentSize(fragment.GetSize())

	return SSLPlaintext{
		content_type: content_type,
		version: version,
		length: length,
		fragment: fragment,
		Serialization: NewNestedSerializable([]Serializable{content_type, version, length, fragment})}
}