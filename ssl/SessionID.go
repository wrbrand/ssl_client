package ssl

import (
	"encoding/binary"
)

type SessionID struct {
	Length uint8
	ID     []byte
}

func NewSessionID(id uint32) SessionID {
	/*
			TODO: Support SessionIDs with bit lengths other than 32
		  	e.g var Length uint8 = uint8(math.Floor(math.Log2(float64(ID)))) + 1
	*/

	var length uint8 = 4

	var sessionID = SessionID{
		Length: length,
		ID:     make([]byte, length)}

	binary.BigEndian.PutUint32(sessionID.ID[0:4], id)

	return sessionID
}

/*
	Returns the total size in bytes of this struct
*/
func (session SessionID) GetSize() int {
	return 1 + int(session.Length)
}

func (session SessionID) SerializeInto(buf []byte) {
	copy(buf[0:1], []uint8{session.Length})
	copy(buf[1:1+session.Length], session.ID[0:session.Length])
}
