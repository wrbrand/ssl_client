package ssl

type ClientHello struct {
	client_version      ProtocolVersion
	random              ClientRandom
	session_id          SessionID	// Legacy
	cipher_suites       CipherSuites
	compression_methods CompressionMethods
	extensions 			 Extensions
}

func NewClientHello(random ClientRandom, config Configuration) ClientHello {
	var compression_methods CompressionMethods = NewCompressionMethods([]CompressionMethod{NULL_COMPRESSION})

	return ClientHello{
		client_version:      config.Handshake.RecordProtocolVersion,
		random:              random,
		session_id:          config.Client.SessionID,
		cipher_suites:       config.Client.GetCipherSuites(),
		compression_methods: compression_methods,
		extensions:			  config.Handshake.GetExtensions(),
	}
}

func (hello ClientHello) GetSerialization() NestedSerializable {
	return NewNestedSerializable([]Serializable{
		hello.client_version,
		hello.random,
		hello.session_id,
		hello.cipher_suites,
		hello.compression_methods,
		hello.extensions,
	})
}

func DeserializeClientHello(buf []byte) (ClientHello, int) {
	var hello = ClientHello{}
	var bytesRead int

	deserializers := []func([]byte) (int){
		func(x []byte) (int) { hello.client_version, bytesRead = DeserializeProtocolVersion(x); return bytesRead },
		func(x []byte) (int) { hello.random, bytesRead = DeserializeClientRandom(x); return bytesRead },
		func(x []byte) (int) { hello.session_id, bytesRead = DeserializeSessionID(x); return bytesRead },
		func(x []byte) (int) { hello.cipher_suites, bytesRead = DeserializeCipherSuites(x); return bytesRead },
		func(x []byte) (int) { hello.compression_methods, bytesRead = DeserializeCompressionMethods(x); return bytesRead },
	}

	var bufferPosition = 0;
	for _, deserializer := range deserializers {
		bufferPosition += deserializer(buf[bufferPosition:])
	}

	return hello, bufferPosition
}
