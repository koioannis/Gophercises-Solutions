package services

import (
	"encoding/binary"
	"encoding/json"
	"fmt"

	bolt "go.etcd.io/bbolt"
)

type Task struct {
	ID int
	Content string 
}

func NewTaskService() (*TaskService, error) {
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		return nil, err
	}

	tx, err := db.Begin(true)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	_, err = tx.CreateBucketIfNotExists([]byte("Tasks"))
	if err != nil {
		return nil, fmt.Errorf("create bucket: %s", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return &TaskService{
		db: db,
	}, nil
}

type TaskService struct {
	db *bolt.DB
}

func (ts *TaskService) Create(t *Task) error {
	return ts.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Tasks"))

		id, err := b.NextSequence()
		if err != nil {
			return err
		}

		t.ID = int(id)

		buf, err := json.Marshal(t)
		if err != nil {
			return err
		}
		
		return b.Put(itob(t.ID), buf)
	})
}

func (ts *TaskService) Close() {
	ts.db.Close()
}

func itob(v int) []byte {
	b := make([]byte, 8)

	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}
