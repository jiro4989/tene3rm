package infra

import (
	"os"
	"path/filepath"
)

type FileRepo struct {
	dir string
}

func NewFileRepo(dir string) FileRepo {
	return FileRepo{
		dir: dir,
	}
}

func (f FileRepo) Save(filename string, body string) error {
	fp, err := os.Create(f.FullPath(filename))
	if err != nil {
		return err
	}
	defer fp.Close()

	_, err = fp.Write([]byte(body))
	if err != nil {
		return err
	}

	return nil
}

func (f FileRepo) FullPath(filename string) string {
	return filepath.Join(f.dir, filename)
}
