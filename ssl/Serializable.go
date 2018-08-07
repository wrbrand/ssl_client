package ssl

type Serializable interface {
	GetSize() int
	Serialize() []byte
	SerializeInto([]byte)
}