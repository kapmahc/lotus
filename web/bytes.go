package web

import (
	"bytes"
	"encoding/gob"
)

//FromBytes bytes to object
func FromBytes(data []byte, val interface{}) error {
	var buf bytes.Buffer
	dec := gob.NewDecoder(&buf)
	buf.Write(data)
	return dec.Decode(val)
}

//ToBytes object to bytes
func ToBytes(val interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(val)
	return buf.Bytes(), err
}
