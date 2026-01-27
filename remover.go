package main

import "os"

type Remover interface {
	Remove(path string) error
}

// インタフェース実装チェック
var _ Remover = &OSRemover{}
var _ Remover = &NilRemover{}

// OSRemover は実際にファイルを削除する。
type OSRemover struct{}

func (r *OSRemover) Remove(path string) error {
	return os.RemoveAll(path)
}

// NilRemover は何もしない。
type NilRemover struct{}

func (r *NilRemover) Remove(path string) error {
	// 何もしない
	return nil
}
