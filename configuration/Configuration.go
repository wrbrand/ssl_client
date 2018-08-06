package configuration

import (
	"os"
	"encoding/json"
	"fmt"
)

type Configuration struct {
	Client ClientConfiguration
	Handshake HandshakeConfiguration
}

func NewConfiguration() Configuration {
	return Configuration{
		Client: NewClientConfiguration(),
		Handshake: NewHandshakeConfiguration(),
	}
}

func Load(filename string) Configuration {
	file, _ := os.Open(filename)
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := NewConfiguration()

	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("Error attempting to decode configuration file:", err)
	}

	fmt.Printf("Session ID: %x\n", configuration.Client.SessionID)
	fmt.Printf("Record Protocol Version: %x\n", configuration.Handshake.RecordProtocolVersion)
	fmt.Printf("Handshake Protocol Version: %x\n", configuration.Handshake.HandshakeProtocolVersion)

	return configuration
}