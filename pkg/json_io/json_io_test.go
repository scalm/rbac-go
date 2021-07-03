package json_io

import (
	"errors"
	"strings"
	"testing"
)

func Test_StringReaderNull(t *testing.T) {
	var document interface{}
	err := NewStringReader("null").Read(&document)
	if err != nil {
		t.Error(err)
		return
	}

	if document != nil {
		t.Error(errors.New("document must be nil"))
		return
	}
}

func Test_StreamReaderNull(t *testing.T) {
	var document map[string]interface{}
	err := NewStreamReader(strings.NewReader("null")).Read(&document)
	if err != nil {
		t.Error(err)
		return
	}


	if document != nil {
		t.Error(errors.New("document must be nil"))
		return
	}
}

func Test_WriterNull(t *testing.T) {
	writer := NewStringWriter()
	err := writer.Write(nil)
	if err != nil {
		t.Error(err)
		return
	}

	if *writer.String != "null" {
		t.Error(errors.New("document string must be null"))
		return
	}
}
