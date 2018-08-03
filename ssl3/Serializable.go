package ssl3

type Serializable interface {
	GetSize() int
	SerializeInto([]byte)
}
