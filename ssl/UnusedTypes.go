package ssl

type HandshakeBody struct{}
type HelloRequest struct{ *HandshakeBody }
type ServerHello struct{ *HandshakeBody }
type Certificate struct{ *HandshakeBody }
type ServerKeyExchange struct{ *HandshakeBody }
type CertificateRequest struct{ *HandshakeBody }
type ServerHelloDone struct{ *HandshakeBody }
type CertificateVerify struct{ *HandshakeBody }
type ClientKeyExchange struct{ *HandshakeBody }
type Finished struct{ *HandshakeBody }
