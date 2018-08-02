package main

type AlertDescription uint8

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

type AlertLevel uint8

const (
	WARNING AlertLevel = 1
	FATAL AlertLevel = 2
)

type Alert struct {
	Level AlertLevel
	Description AlertDescription
}