package ssl

type ClientHello struct {
	client_version      ProtocolVersion
	random              ClientRandom
	session_id          SessionID
	cipher_suites       CipherSuites
	compression_methods CompressionMethods
}

func NewClientHello(version ProtocolVersion, random ClientRandom, session_id SessionID) ClientHello {
	var suites CipherSuites = NewCipherSuites([]CipherSuite{
		TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
		TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
		TLS_DHE_DSS_WITH_AES_256_GCM_SHA384,
		TLS_DHE_RSA_WITH_AES_256_GCM_SHA384,
		TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256,
		TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256,
		TLS_DHE_RSA_WITH_CHACHA20_POLY1305_SHA256,
		TLS_ECDHE_ECDSA_WITH_AES_256_CCM_8,
		TLS_ECDHE_ECDSA_WITH_AES_256_CCM,
		TLS_DHE_RSA_WITH_AES_256_CCM_8,
		TLS_DHE_RSA_WITH_AES_256_CCM,
		TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA384,
		TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA384,
		TLS_ECDHE_ECDSA_WITH_CAMELLIA_256_CBC_SHA384,
		TLS_ECDHE_RSA_WITH_CAMELLIA_256_CBC_SHA384,
		TLS_DHE_RSA_WITH_CAMELLIA_256_CBC_SHA256,
		TLS_DHE_DSS_WITH_CAMELLIA_256_CBC_SHA256,
		TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA,
		TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
		TLS_DHE_RSA_WITH_CAMELLIA_256_CBC_SHA,
		TLS_DHE_DSS_WITH_CAMELLIA_256_CBC_SHA,
		TLS_RSA_WITH_AES_256_GCM_SHA384,
		TLS_RSA_WITH_AES_256_CCM_8,
		TLS_RSA_WITH_AES_256_CCM,
		TLS_RSA_WITH_CAMELLIA_256_CBC_SHA256,
		TLS_RSA_WITH_CAMELLIA_256_CBC_SHA,
		TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
		TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
		TLS_DHE_DSS_WITH_AES_128_GCM_SHA256,
		TLS_DHE_RSA_WITH_AES_128_GCM_SHA256,
		TLS_ECDHE_ECDSA_WITH_AES_128_CCM_8,
		TLS_ECDHE_ECDSA_WITH_AES_128_CCM,
		TLS_DHE_RSA_WITH_AES_128_CCM_8,
		TLS_DHE_RSA_WITH_AES_128_CCM,
		TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256,
		TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256,
		TLS_ECDHE_ECDSA_WITH_CAMELLIA_128_CBC_SHA256,
		TLS_ECDHE_RSA_WITH_CAMELLIA_128_CBC_SHA256,
		TLS_DHE_RSA_WITH_CAMELLIA_128_CBC_SHA256,
		TLS_DHE_DSS_WITH_CAMELLIA_128_CBC_SHA256,
		TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA,
		TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA,
		TLS_DHE_RSA_WITH_SEED_CBC_SHA,
		TLS_DHE_DSS_WITH_SEED_CBC_SHA,
		TLS_DHE_RSA_WITH_CAMELLIA_128_CBC_SHA,
		TLS_DHE_DSS_WITH_CAMELLIA_128_CBC_SHA,
		TLS_RSA_WITH_AES_128_GCM_SHA256,
		TLS_RSA_WITH_AES_128_CCM_8,
		TLS_RSA_WITH_AES_128_CCM,
		TLS_RSA_WITH_CAMELLIA_128_CBC_SHA256,
		TLS_RSA_WITH_SEED_CBC_SHA,
		TLS_RSA_WITH_CAMELLIA_128_CBC_SHA,
		TLS_RSA_WITH_IDEA_CBC_SHA,
		TLS_EMPTY_RENEGOTIATION_INFO_SCSV,
	})
	var compression_methods CompressionMethods = NewCompressionMethods([]CompressionMethod{NULL_COMPRESSION})

	return ClientHello{
		client_version:      version,
		random:              random,
		session_id:          session_id,
		cipher_suites:       suites,
		compression_methods: compression_methods,
	}
}

func (hello ClientHello) GetSerialization() NestedSerializable {
	return NewNestedSerializable([]Serializable{
		hello.client_version,
		hello.random,
		hello.session_id,
		hello.cipher_suites,
		hello.compression_methods,
	})
}

func DeserializeClientHello(buf []byte) (ClientHello, int) {
	var bufferPosition = 0;

	version, bytesRead := DeserializeProtocolVersion(buf[bufferPosition:])

	bufferPosition += bytesRead

	random, bytesRead := DeserializeClientRandom(buf[bufferPosition:])

	bufferPosition += bytesRead

	session_id, bytesRead := DeserializeSessionID(buf[bufferPosition:])

	bufferPosition += bytesRead

	suites, bytesRead := DeserializeCipherSuites(buf[bufferPosition:])

	bufferPosition += bytesRead

	compression_methods, bytesRead := DeserializeCompressionMethods(buf[bufferPosition:])

	bufferPosition += bytesRead

	return ClientHello{
		version,
		random,
		session_id,
		suites,
		compression_methods,
	}, bufferPosition
}
