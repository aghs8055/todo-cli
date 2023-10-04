package contract

type StorageReader interface {
	Read(entityName string, entities any) error
}

type StorageWriter interface {
	Write(entityName string, entities any) error
}

type StorageReaderWriter interface {
	StorageReader
	StorageWriter
}
