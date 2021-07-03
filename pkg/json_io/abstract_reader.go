package json_io

type AbstractReader interface {
	Read(document interface{}) error
}
