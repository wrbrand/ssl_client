package handshake

import (
	"../records"
)

/*
These processes are performed
   in the handshake protocol, which can be summarized as follows: the
   client sends a client hello message to which the server must respond
   with a server hello message, or else a fatal error will occur and the
   connection will fail.  The client hello and server hello are used to
   establish security enhancement capabilities between client and
   server.  The client hello and server hello establish the following
   attributes: Protocol Version, Session ID, Cipher Suite, and
   Compression Method.  Additionally, two random values are generated
   and exchanged: ClientHello.random and ServerHello.random.
*/

/*

   Client                                                Server

   ClientHello                   -------->
                                                    ServerHello
                                                   Certificate*
                                             ServerKeyExchange*
                                            CertificateRequest*
                                 <--------      ServerHelloDone
   Certificate*
   ClientKeyExchange
   CertificateVerify*
   [ChangeCipherSpec]
   Finished                      -------->
                                             [ChangeCipherSpec]
                                 <--------             Finished
   Application Data              <------->     Application Data

   * Indicates optional or situation-dependent messages that are not
     always sent.
*/

type HandshakeType uint8

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

type Handshake struct {
	Msg_type HandshakeType
	length   []byte // Per spec, this should be a 24-bit uint
	body     HandshakeBody
}

type HandshakeBody struct {
}

type HelloRequest struct {
	*HandshakeBody
}

type SessionID struct {
	length uint8
	id     []byte
}

func NewSessionID(id uint) {
	// TODO: Finish generating session IDs
	/*var length = math.Floor(math.Log2(float64(id))) + 1
	var sessionID = SessionID {
		length: length,
		id: make([]byte, length)}

	binary.Write(sessionID.id, binary.LittleEndian, id)

	return sessionID*/
}

type ClientHello struct {
	*HandshakeBody
	client_version      records.ProtocolVersion
	random              Random
	session_id          SessionID
	cipher_suites       []byte // 2 bytes of length, followed by up to 2^16-1 bytes of data
	compression_methods []byte // 1 byte of length, followed by up to 2^8-1 bytes of data
}

func NewClientHello(random Random, session_id SessionID) ClientHello {
	var hello = ClientHello{
		client_version:      records.ProtocolVersion{3, 0},
		random:              random,
		session_id:          session_id,
		cipher_suites:       nil,
		compression_methods: nil}

	return hello
}

type ServerHello struct {
	*HandshakeBody
}

type Certificate struct {
	*HandshakeBody
}

type ServerKeyExchange struct {
	*HandshakeBody
}

type CertificateRequest struct {
	*HandshakeBody
}

type ServerHelloDone struct {
	*HandshakeBody
}

type CertificateVerify struct {
	*HandshakeBody
}

type ClientKeyExchange struct {
	*HandshakeBody
}

type Finished struct {
	*HandshakeBody
}
