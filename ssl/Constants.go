package ssl

const (
	WARNING AlertLevel = 1
	FATAL AlertLevel = 2
)

const (
	CLOSE_NOTIFY AlertDescription = 0
	UNEXPECTED_MESSAGE = 10
	BAD_RECORD_MAC = 20
	DECOMPRESSION_FAILURE = 30
	HANDSHAKE_FAILURE = 40
	NO_CERTIFICATE = 41
	BAD_CERTIFICATE = 42
	UNSUPPORTED_CERTIFICATE = 43
	CERTIFICATE_REVOKED = 44
	CERTIFICATE_EXPIRED = 45
	CERTIFICATE_UNKNOWN = 46
	ILLEGAL_PARAMETER = 47
)

const (
	NULL_COMPRESSION CompressionMethod = 0
)

const (
	CHANGE_CIPHER_SPEC ContentType = 20
	ALERT                          = 21
	HANDSHAKE                      = 22
	APPLICATION_DATA               = 23
)

const (
	HELLO_REQUEST       HandshakeType = 0
	CLIENT_HELLO                      = 1
	SERVER_HELLO                      = 2
	CERTIFICATE                       = 11
	SERVER_KEY_EXCHANGE               = 12
	CERTIFICATE_REQUEST               = 13
	SERVER_HELLO_DONE                 = 14
	CERTIFICATE_VERIFY                = 15
	CLIENT_KEY_EXCHANGE               = 16
	FINISHED                          = 20
)