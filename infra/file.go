package infra

import (
	"encoding/json"
	"errors"
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

func (f FileRepo) Exists(filename string) (bool, error) {
	_, err := os.Stat(f.FullPath(filename))
	if err == nil {
		return true, nil
	}
	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	}
	return false, err
}

func (f FileRepo) createFile(filename string) (*os.File, error) {
	return os.Create(f.FullPath(filename))
}

func (f FileRepo) Save(filename string, data string) error {
	fp, err := f.createFile(filename)
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
	fp, err := f.createFile(filename)
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

func (f FileRepo) LoadJSON(filename string, data any) error {
	if ok, err := f.Exists(filename); err != nil {
		return err
	} else if ok {
		b, err := os.ReadFile(f.FullPath(filename))
		if err != nil {
			return err
		}
		if err := json.Unmarshal(b, &data); err != nil {
			return err
		}
	}
	return nil
}

func (f FileRepo) FullPath(filename string) string {
	return filepath.Join(f.dir, filename)
}
