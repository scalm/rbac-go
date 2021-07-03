package json_io

import (
	"encoding/json"
	"io"
)

type StreamWriter struct {
	upstreamWriter io.Writer
}

func NewStreamWriter(writer io.Writer) *StreamWriter {
	return &StreamWriter{writer}
}

func (writer *StreamWriter) Write(document interface{}) error {
	encoder := json.NewEncoder(writer.upstreamWriter)
	return encoder.Encode(&document)
}
