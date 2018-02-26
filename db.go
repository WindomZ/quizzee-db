package quizzee_db

import "github.com/WindomZ/gkv"

func Open(table []byte, paths ...string) error {
	return gkv.Open(table, paths...)
}

func Close() error {
	return gkv.Close()
}

func Put(table, key, value []byte) error {
	return gkv.Put(table, key, value)
}

func Get(table, key []byte) []byte {
	return gkv.Get(table, key)
}

func Count(table []byte) int {
	return gkv.Count(table)
}
