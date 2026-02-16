package repo

type DataStoreRepo interface {
	Exists(filename string) (bool, error)
	Save(filename string, data string) error
	SaveJSON(filename string, data any) error
	LoadJSON(filename string, data any) error
}
