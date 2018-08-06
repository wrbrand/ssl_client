package ssl

type ProtocolVersion struct {
	Major uint8
	Minor uint8
}

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

func DeserializeProtocolVersion(buf []byte) (ProtocolVersion, int) {
	return ProtocolVersion{buf[0], buf[1]}, 2
}