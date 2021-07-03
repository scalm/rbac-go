package json_io

import (
	"encoding/json"
	"io"
)

type StreamReader struct {
	downstreamReader io.Reader
}

func NewStreamReader(reader io.Reader) *StreamReader {
	return &StreamReader{reader}
}

func (reader *StreamReader) Read(document interface{}) error {
	decoder := json.NewDecoder(reader.downstreamReader)
	return decoder.Decode(&document)
}