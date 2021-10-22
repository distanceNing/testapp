package filestore

import "io"

const (
	CosStore = 1
)

type FileStore interface {
	Get(path string) error
	Put(path string, r io.Reader) (string, error)
}

func MakeFileStore(storeType int) FileStore {
	if storeType == CosStore {
		return NewCosFileStore(&DefaultCosConf)
	} else {
		return nil
	}
}
