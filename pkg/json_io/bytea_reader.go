package json_io

import (
	"encoding/json"
)

type ByteAReader struct {
	bytes []byte
}

func NewByteAReader(bytes []byte) *ByteAReader {
	return &ByteAReader{bytes}
}

func NewStringReader(s string) *ByteAReader {
	return NewByteAReader([]byte(s))
}

func (reader *ByteAReader) Read(document interface{}) error {
	return json.Unmarshal(reader.bytes, &document)
}
