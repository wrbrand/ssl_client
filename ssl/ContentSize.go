package ssl

import "encoding/binary"

type ContentSize [2]byte

func NewContentSize(num int) ContentSize {
	var ret [2]byte

	binary.BigEndian.PutUint16(ret[0:2], uint16(num))

	return ret
}

func (num ContentSize) GetSize() int {
	return 2
}

func (num ContentSize) Serialize() []byte {
	return num[0:2]
}

func (num ContentSize) SerializeInto(buf []byte) {
	copy(buf[0:num.GetSize()], num.Serialize())
}
