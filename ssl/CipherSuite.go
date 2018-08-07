package ssl

type CipherSuite uint16

func NewCipherSuite(value uint16) CipherSuite {
	return CipherSuite(value)
}

func (suite CipherSuite) GetValue() uint16 {
	return uint16(suite)
}