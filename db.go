package quizzee_db

import (
	"fmt"

	"github.com/boltdb/bolt"
)

var (
	dbPath string
	db     *bolt.DB
)

func init() {
	dbPath = ConfigString("db_path")

	var err error
	db, err = bolt.Open(dbPath, 0600, nil)
	if err != nil {
		panic(err)
	}
}

func DBRegister(name []byte) error {
	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(name)
		if err != nil {
			return fmt.Errorf("CreateBucketIfNotExists: %s",
				err.Error())
		}
		return nil
	})
}

func DBPut(name, key, value []byte) error {
	return db.Update(func(tx *bolt.Tx) error {
		return tx.Bucket(name).Put(key, value)
	})
}

func DBGet(name, key []byte) (value []byte) {
	db.View(func(tx *bolt.Tx) error {
		value = tx.Bucket(name).Get(key)
		return nil
	})
	return
}

func DBCount(name []byte) (i int) {
	db.View(func(tx *bolt.Tx) error {
		c := tx.Bucket(name).Cursor()
		for k, _ := c.First(); k != nil; k, _ = c.Next() {
			i++
		}
		return nil
	})
	return i
}
