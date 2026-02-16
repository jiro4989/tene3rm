package infra

import (
	"encoding/json"

	"github.com/jiro4989/tene3rm/repo"
)

var _ repo.DataStoreRepo = InMemoryRepo{}

type InMemoryRepo struct {
	data map[string][]byte
	err  error
}

func NewInMemoryRepo(data map[string][]byte, err error) InMemoryRepo {
	return InMemoryRepo{
		data: data,
		err:  err,
	}
}

func NewDefaultInMemoryRepo() InMemoryRepo {
	return NewInMemoryRepo(make(map[string][]byte), nil)
}

func (i InMemoryRepo) Exists(key string) (bool, error) {
	if i.err != nil {
		return false, i.err
	}
	_, ok := i.data[key]
	return ok, nil
}

func (i InMemoryRepo) Save(key string, data string) error {
	if i.err != nil {
		return i.err
	}
	i.data[key] = []byte(data)
	return nil
}

func (i InMemoryRepo) SaveJSON(key string, data any) error {
	if i.err != nil {
		return i.err
	}

	b, err := json.Marshal(data)
	if err != nil {
		return err
	}
	i.data[key] = b

	return nil
}

func (i InMemoryRepo) LoadJSON(key string, data any) error {
	if i.err != nil {
		return i.err
	}

	if err := json.Unmarshal(i.data[key], &data); err != nil {
		return err
	}

	return nil
}
