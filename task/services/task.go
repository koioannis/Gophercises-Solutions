package services

import (
	"encoding/binary"
	"fmt"

	bolt "go.etcd.io/bbolt"
)

type Task struct {
	ID      int
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

		return b.Put(itob(t.ID), []byte(t.Content))
	})
}

func (ts *TaskService) GetAll() ([]Task, error) {
	var tasks []Task

	err := ts.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Tasks"))

		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			tasks = append(tasks, Task{
				ID:      btoi(k),
				Content: string(v),
			})
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return tasks, nil
}


func (ts *TaskService) Close() {
	ts.db.Close()
}

// integer to bytes
func itob(v int) []byte {
	b := make([]byte, 8)

	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

// bytes to integer
func btoi(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}
