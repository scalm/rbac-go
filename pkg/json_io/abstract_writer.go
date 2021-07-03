package json_io

type AbstractWriter interface {
	Write(document interface{}) error
}
