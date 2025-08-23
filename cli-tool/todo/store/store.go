package store

import (
	"encoding/json"
	"os"
)

type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

type TaskStore interface {
	LoadTask() ([]Task, error)
	SaveTask([]Task) error
}

type FileTaskStore struct {
	Path string
}

func (f *FileTaskStore) LoadTask() ([]Task, error) {
	data, err := os.ReadFile(f.Path)

	if err != nil {
		return nil, err
	}

	var tasks []Task
	err = json.Unmarshal(data, &tasks)

	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (f *FileTaskStore) SaveTask(tasks []Task) error {

	data, err := json.MarshalIndent(tasks, "", " ")

	if err != nil {
		return err
	}

	return os.WriteFile(f.Path, data, 0644)
}
