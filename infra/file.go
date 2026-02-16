package infra

import (
	"encoding/json"
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

func (f FileRepo) Save(filename string, data string) error {
	fp, err := os.Create(f.FullPath(filename))
	if err != nil {
		return err
	}
	defer fp.Close()

	_, err = fp.Write([]byte(data))
	if err != nil {
		return err
	}

	return nil
}

func (f FileRepo) SaveJSON(filename string, data any) error {
	fp, err := os.Create(f.FullPath(filename))
	if err != nil {
		return err
	}
	defer fp.Close()

	b, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = fp.Write(b)
	if err != nil {
		return err
	}

	return nil
}

func (f FileRepo) FullPath(filename string) string {
	return filepath.Join(f.dir, filename)
}
