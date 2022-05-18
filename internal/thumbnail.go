package internal

type Thumbnail interface {
	Open(storage *StorageFile) (Thumbnail, error)
}
