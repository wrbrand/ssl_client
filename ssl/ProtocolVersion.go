package ssl

type ProtocolVersion struct {
	Major uint8
	Minor uint8
}

var SSL3 ProtocolVersion = ProtocolVersion { 3, 0 }
var TLS10 ProtocolVersion = ProtocolVersion { 3, 1 }
var TLS11 ProtocolVersion = ProtocolVersion { 3, 2 }
var TLS12 ProtocolVersion = ProtocolVersion { 3, 3 }

func (version ProtocolVersion) GetSize() int {
	return 2
}

/*
	Serializes this struct into a given buffer, which is assumed to be 2 bytes.
*/
func (version ProtocolVersion) SerializeInto(buf []byte) {
	buf[0] = version.Major
	buf[1] = version.Minor
}

func (version ProtocolVersion) Serialize() []byte {
	obj := make([]byte, version.GetSize())
	version.SerializeInto(obj)
	return obj
}

func DeserializeProtocolVersion(buf []byte) (ProtocolVersion, int) {
	return ProtocolVersion{buf[0], buf[1]}, 2
}