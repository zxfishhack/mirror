package simple

import (
	"github.com/zxfishhack/mirror/pkg/storage"
	"os"
	"path/filepath"
)

type simpleStorage struct {
	basePath string
	name     string
}

func (s *simpleStorage) Put(key string, value []byte) error {
	realPath := filepath.Join(s.basePath, key)
	return os.WriteFile(realPath, value, 0644)
}

func (s *simpleStorage) Get(key string) ([]byte, error) {
	realPath := filepath.Join(s.basePath, key)
	_ = os.MkdirAll(filepath.Dir(realPath), 0755)
	return os.ReadFile(realPath)
}

func CreateFunc(storagePath string) storage.CreateStorageFunc {
	return func(name string) (storage.IStorage, error) {
		return &simpleStorage{basePath: storagePath, name: name}, nil
	}
}
