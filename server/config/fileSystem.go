package config

import (
	"io/ioutil"
)

var iouUtil fileSystem = osFS{}

type fileSystem interface {
	ReadFile(name string) ([]byte, error)
}

// osFS implements fileSystem using the local disk.
type osFS struct{}

func (osFS) ReadFile(name string) ([]byte, error) { return ioutil.ReadFile(name) }
