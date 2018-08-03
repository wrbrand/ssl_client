package ssl3

type ClientHello struct {
	client_version      ProtocolVersion
	random              ClientRandom
	session_id          SessionID
	cipher_suites       CipherSuites
	compression_methods CompressionMethods
	Serialization       NestedSerializable
}

func NewClientHello(random ClientRandom, session_id SessionID) ClientHello {
	version := ProtocolVersion{3, 0}

	var suites CipherSuites = NewCipherSuites([]CipherSuite{
		SSL_DH_DSS_WITH_DES_CBC_SDA,
		SSL_DH_DSS_EXPORT_WITH_DES40_CBC_SDA,
		SSL_RSA_WITH_RC4_128_MD5,
		SSL_NULL_WITH_NULL_NULL})

	var compression_methods CompressionMethods = NewCompressionMethods([]CompressionMethod{NULL_COMPRESSION})

	return ClientHello{
		client_version:      version,
		random:              random,
		session_id:          session_id,
		cipher_suites:       suites,
		compression_methods: compression_methods,
		Serialization:       NewNestedSerializable([]Serializable{version, random, session_id, suites, compression_methods})}
}
