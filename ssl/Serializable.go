package ssl

type Serializable interface {
	GetSize() int
	SerializeInto([]byte)
}
