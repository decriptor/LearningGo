package persistence

import (
	"encoding/json"
	"os"
	"taskmanager2/internal/app"
)

type JSONStore struct {
	FilePath string
}

func NewJSONStore(filePath string) *JSONStore {
	return &JSONStore{FilePath: filePath}
}

func (js *JSONStore) LoadTasks() ([]app.Task, int, error) {
	file, err := os.Open(js.FilePath)
	if os.IsNotExist(err) {
		return nil, 0, nil
	}
	if err != nil {
		return nil, 0, err
	}
	defer file.Close()

	var data struct {
		Tasks  []app.Task `json:"tasks"`
		NextID int        `json:"nextID"`
	}
	if err := json.NewDecoder(file).Decode(&data); err != nil {
		return nil, 0, err
	}
	return data.Tasks, data.NextID, nil
}

func (js *JSONStore) SaveTasks(tasks []app.Task, nextID int) error {
	file, err := os.Create(js.FilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	data := struct {
		Tasks  []app.Task `json:"tasks"`
		NextID int        `json:"nextID"`
	}{
		Tasks:  tasks,
		NextID: nextID,
	}
	return json.NewEncoder(file).Encode(data)
}
