package model

type IOpenAPI3 interface {
	GetJSON() ([]byte, error)
	GetYAML() ([]byte, error)
}
