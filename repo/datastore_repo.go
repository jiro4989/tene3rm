package repo

type DataStoreRepo interface {
	Save(filename string, body string) error
}
