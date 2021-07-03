package json_io

import (
	"encoding/json"
)

type ByteAWriter struct {
	Bytes []byte
}


func NewByteAWriter() *ByteAWriter {
	return &ByteAWriter{}
}

func (writer *ByteAWriter) Write(document interface{}) error {
	bytes, err := json.Marshal(document)
	if err != nil {
		return err
	}

	writer.Bytes = bytes
	return nil
}

type StringWriter struct {
	byteaWriter *ByteAWriter
	String *string
}

func NewStringWriter() *StringWriter {
	return &StringWriter{NewByteAWriter(), nil}
}

func (writer *StringWriter) Write(document interface{}) error {
	err := writer.byteaWriter.Write(document)
	if err != nil {
		return err
	}

	s := string(writer.byteaWriter.Bytes)
	writer.String = &s
	return nil
}