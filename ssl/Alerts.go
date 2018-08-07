package ssl

type AlertDescription uint8
func (description AlertDescription) GetSize() int { return 1 }
func (description AlertDescription) Serialize() []byte { return []byte { uint8(description) } }
func (description AlertDescription) SerializeInto(buf []byte) { buf[0] = uint8(description) }

type AlertLevel uint8
func (level AlertLevel) GetSize() int { return 1 }
func (level AlertLevel) Serialize() []byte { return []byte { uint8(level) } }
func (level AlertLevel) SerializeInto(buf []byte) { buf[0] = uint8(level) }

type Alert struct {
	Level AlertLevel
	Description AlertDescription
}

func (alert Alert) GetSize() int { return 2 }
func (alert Alert) SerializeInto(buf []byte) {
	copy(buf, []byte{
		uint8(alert.Level),
		uint8(alert.Description),
	})
}
func (alert Alert) Serialize() []byte {
	obj := make([]byte, alert.GetSize())
	alert.SerializeInto(obj)
	return obj
}

func DeserializeAlert(buf []byte) (Alert, int) {
	var handshake = Alert{}

	deserializers := []func([]byte) (int){
		func(buf []byte) (int) {
			handshake.Level = AlertLevel(buf[0])
			return 1
		},
		func(buf []byte) (int) {
			handshake.Description = AlertDescription(buf[0])
			return 1
		},
	}

	var bufferPosition = 0;
	for _, deserializer := range deserializers {
		bufferPosition += deserializer(buf[bufferPosition:])
	}

	return handshake, bufferPosition
}