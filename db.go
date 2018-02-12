package quizzee_db

import (
	"errors"
	"fmt"
	"path"
	"runtime"

	"github.com/boltdb/bolt"
)

var db *bolt.DB

var defaultPath string

func init() {
	_, filePath, _, _ := runtime.Caller(0)
	defaultPath = path.Join(path.Dir(filePath), "data", "data.db")
}

func Open(dataPath ...string) (err error) {
	dbPath := defaultPath
	if len(dataPath) != 0 {
		dbPath = dataPath[0]
	}
	db, err = bolt.Open(dbPath, 0600, nil)
	return
}

func Close() error {
	if db != nil {
		return db.Close()
	}
	return nil
}

func Register(name []byte) error {
	if db == nil {
		if err := Open(); err != nil {
			return err
		}
	}
	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(name)
		if err != nil {
			return fmt.Errorf("CreateBucketIfNotExists: %s",
				err.Error())
		}
		return nil
	})
}

func Put(name, key, value []byte) error {
	if db == nil {
		return errors.New("the db service is not started")
	}
	return db.Update(func(tx *bolt.Tx) error {
		return tx.Bucket(name).Put(key, value)
	})
}

func Get(name, key []byte) (value []byte) {
	if db == nil {
		return nil
	}
	db.View(func(tx *bolt.Tx) error {
		value = tx.Bucket(name).Get(key)
		return nil
	})
	return
}

func Count(name []byte) (i int) {
	if db == nil {
		return 0
	}
	db.View(func(tx *bolt.Tx) error {
		c := tx.Bucket(name).Cursor()
		for k, _ := c.First(); k != nil; k, _ = c.Next() {
			i++
		}
		return nil
	})
	return i
}
