package repo

type DataStoreRepo interface {
	Save(filename string, data string) error
	SaveJSON(filename string, data any) error
}
