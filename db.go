package quizzee_db

import "errors"

type DB interface {
	Close() error
	Register([]byte) error
	Put([]byte, []byte, []byte) error
	Get([]byte, []byte) []byte
	Count([]byte) int
}

var db DB

type Instance func(paths ...string) DB

var inst Instance

func Register(i Instance) {
	inst = i
}

func Open(table []byte, paths ...string) error {
	if inst == nil {
		return errors.New("forgot to import the driver")
	}
	db = inst(paths...)
	return db.Register(table)
}

func Close() error {
	if db != nil {
		return db.Close()
	}
	return nil
}

func Put(table, key, value []byte) error {
	if db == nil {
		return errors.New("the db service is not started")
	}
	return db.Put(table, key, value)
}

func Get(table, key []byte) []byte {
	if db == nil {
		return nil
	}
	return db.Get(table, key)
}

func Count(table []byte) int {
	if db == nil {
		return 0
	}
	return db.Count(table)
}
