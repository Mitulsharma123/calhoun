package db

import (
	"github.com/boltdb/bolt"
	"time"
)

var taskBucket = []byte("tasks")

var db *bolt.DB

type Task struct{
	Key int
	Value string
}

func Init(dbPath string) error{
	var err error
	db, err := bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err!=nil{
		return err
	}
	return db.Update(func(tx *bolt.Tx) error{
		_, err := tx.CreateBucketIfNotExists(taskBucket)
		return err
	})
}


/*
type User struct {
	ID int // primary key
	Group string `storm:"index"` // this field will be indexed
	Email string `storm:"unique"` // this field will be indexed with a unique constraint
	Name string // this field will not be indexed
	Age int `storm:"index"`
  }*/