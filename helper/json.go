package helper

import (
	"bytes"
	"encoding/json"
)

// declare string map to handle queue message
type Message map[string]interface{}

func Deserialize(b []byte) (Message, error) {
	// declare message as map string interface
	var msg Message
	// creates and initializes a new Buffer
	buf := bytes.NewBuffer(b)
	// A Decoder reads and decodes JSON values from an input stream.
	decoder := json.NewDecoder(buf)
	// Decode reads the next JSON-encoded value from its
	err := decoder.Decode(&msg)
	// Return the value and error
	return msg, err
}
