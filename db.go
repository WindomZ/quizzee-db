package quizzee_db

import "github.com/WindomZ/gkv"

func Open(table []byte, paths ...string) error {
	return gkv.Open(table, paths...)
}

func Close() error {
	return gkv.Close()
}

func Put(key, value []byte) error {
	return gkv.Put(key, value)
}

func Get(key []byte) []byte {
	return gkv.Get(key)
}

func Count([]byte) int {
	return gkv.Count()
}
