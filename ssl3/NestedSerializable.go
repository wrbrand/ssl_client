package ssl3

type NestedSerializable struct {
	serialization []Serializable
}

func NewNestedSerializable(serialization []Serializable) NestedSerializable {
	return NestedSerializable{serialization}
}

func (object NestedSerializable) GetSize() int {
	size := 0

	for _, component := range object.serialization {
		size += component.GetSize()
	}

	return size
}

func (object NestedSerializable) Serialize() []byte {
	message := make([]byte, object.GetSize())
	index := 0

	for _, component := range object.serialization {
		newIndex := index + component.GetSize()

		component.SerializeInto(message[index:newIndex])

		index = newIndex
	}

	return message
}

func (object NestedSerializable) SerializeInto(buf []byte) {
	copy(buf[0:object.GetSize()], object.Serialize())
}
