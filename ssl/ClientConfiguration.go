package ssl

import (
	"math/rand"
)

type ClientConfiguration struct {
	SessionID        SessionID
	SupportedCiphers []string
}

func NewClientConfiguration() ClientConfiguration {
	return ClientConfiguration{
		NewSessionID(rand.Uint32()),
		[]string{},
	}
}

func (client ClientConfiguration) GetCipherSuites() CipherSuites {
	var suites = []CipherSuite{}

	for _, suiteName := range client.SupportedCiphers {
		suites = append(suites, NewCipherSuite(SupportedCiphers[suiteName]))
	}

	if len(suites) == 0 {
		return DefaultCipherSuites
	} else {
		return NewCipherSuites(suites)
	}
}