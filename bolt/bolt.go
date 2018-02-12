package bolt

import (
	"fmt"
	"path/filepath"
	"runtime"

	quiz "github.com/WindomZ/quizzee-db"
	"github.com/boltdb/bolt"
)

var defaultPath string

type DB struct {
	bolt.DB
}

func init() {
	_, filePath, _, _ := runtime.Caller(0)
	defaultPath = filepath.Join(filepath.Dir(filepath.Dir(filePath)),
		"data", "data.db")
}

func Open(paths ...string) quiz.DB {
	path := defaultPath
	if len(paths) != 0 {
		path = paths[0]
	}
	db, err := bolt.Open(path, 0600, nil)
	if err != nil {
		panic(err)
	}
	return &DB{
		DB: *db,
	}
}

func (db *DB) Close() error {
	return db.DB.Close()
}

func (db *DB) Register(table []byte) error {
	return db.DB.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(table)
		if err != nil {
			return fmt.Errorf("CreateBucketIfNotExists: %s",
				err.Error())
		}
		return nil
	})
}

func (db *DB) Put(table, key, value []byte) error {
	return db.DB.Update(func(tx *bolt.Tx) error {
		return tx.Bucket(table).Put(key, value)
	})
}

func (db *DB) Get(table, key []byte) (value []byte) {
	db.DB.View(func(tx *bolt.Tx) error {
		value = tx.Bucket(table).Get(key)
		return nil
	})
	return
}

func (db *DB) Count(table []byte) (i int) {
	db.DB.View(func(tx *bolt.Tx) error {
		c := tx.Bucket(table).Cursor()
		for k, _ := c.First(); k != nil; k, _ = c.Next() {
			i++
		}
		return nil
	})
	return i
}

func init() {
	quiz.Register(Open)
}
