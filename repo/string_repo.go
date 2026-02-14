package repo

type StringRepo interface {
	Save(filename string, body string) error
}
