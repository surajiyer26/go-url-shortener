package storage

type Store interface {
	Save(id, url string)
	Get(id string) (string, bool)
}
