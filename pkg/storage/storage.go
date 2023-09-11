package storage

type IStorage interface {
	Put(key string, value []byte) error
	Get(key string) ([]byte, error)
}

type CreateStorageFunc func(name string) (IStorage, error)
