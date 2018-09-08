package ssl

type HandshakeConfiguration struct {
	RecordProtocolVersion ProtocolVersion
	HandshakeProtocolVersion  ProtocolVersion
	SupportedExtensions []string
}

func NewHandshakeConfiguration() HandshakeConfiguration {
	return HandshakeConfiguration{
		TLS12,
		TLS12,
		[]string{},
	}
}

func (handshake HandshakeConfiguration) GetExtensions() Extensions {
	var extensions = []Extension{}

	for _, extensionName := range handshake.SupportedExtensions {
		extensions = append(extensions, NewExtension(SupportedExtensions[extensionName]))
	}

	if len(extensions) == 0 {
		return DefaultExtensions
	} else {
		return NewExtensions(extensions)
	}
}