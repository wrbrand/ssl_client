package ssl

import (
	"encoding/binary"
)

type SessionID struct {
	length uint8
	id     []byte
}

func NewSessionID(id uint32) SessionID {
	/*
			TODO: Support SessionIDs with bit lengths other than 32
		  	e.g var length uint8 = uint8(math.Floor(math.Log2(float64(id)))) + 1
	*/

	var length uint8 = 4

	var sessionID = SessionID{
		length: length,
		id:     make([]byte, length)}

	binary.BigEndian.PutUint32(sessionID.id[0:4], id)

	return sessionID
}

/*
	Returns the total size in bytes of this struct
*/
func (session SessionID) GetSize() int {
	return 1 + int(session.length)
}

func (session SessionID) SerializeInto(buf []byte) {
	copy(buf[0:1], []uint8{session.length})
	copy(buf[1:1+session.length], session.id[0:session.length])
}
