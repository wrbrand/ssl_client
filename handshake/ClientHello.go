package handshake

import (
	"../clientRandom"
	"../records"
	"fmt"
)

type Serializable interface {
	GetSize() int
	SerializeInto([]byte)
}

type ClientHello struct {
	client_version      records.ProtocolVersion
	random              clientRandom.ClientRandom
	session_id          SessionID
	serialization       []Serializable
	cipher_suites       records.CipherSuites
	compression_methods records.CompressionMethods
}

func NewClientHello(random clientRandom.ClientRandom, session_id SessionID) ClientHello {
	version := records.ProtocolVersion{3, 0}

	var suites records.CipherSuites = records.NewCipherSuites([]records.CipherSuite{
		records.SSL_NULL_WITH_NULL_NULL,
		records.SSL_DH_DSS_EXPORT_WITH_DES40_CBC_SDA,
		records.SSL_DH_DSS_WITH_DES_CBC_SDA})

	var compression_methods records.CompressionMethods = records.NewCompressionMethods([]records.CompressionMethod{
		records.NULL_COMPRESSION})

	return ClientHello{
		client_version:      version,
		random:              random,
		session_id:          session_id,
		cipher_suites:       suites,
		compression_methods: compression_methods,
		serialization:       []Serializable{version, random, session_id, suites, compression_methods}}
}

func (hello ClientHello) GetSize() int {
	size := 0

	for _, component := range hello.serialization {
		size += component.GetSize()
	}

	return size
}

func (hello ClientHello) Serialize() []byte {
	message := make([]byte, hello.GetSize())
	index := 0

	for _, component := range hello.serialization {
		newIndex := index + component.GetSize()

		component.SerializeInto(message[index:newIndex])

		index = newIndex
	}

	fmt.Printf("%x", message)

	return message
	/*
		fmt.Print(mes, "\n")
		var b bytes.Buffer

		binary.Write(&b, binary.BigEndian, hello.client_version)
		b = append(b, hello.client_version.Serialize()...)

		binary.Write(&b, binary.BigEndian, hello.random)
		binary.Write(&b, binary.BigEndian, hello.session_id)*/
}
