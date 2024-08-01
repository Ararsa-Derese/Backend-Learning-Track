package data

import (
    "errors"
    "taskmanager/models"
    "sync"
)

var (
    tasks = make([]models.Task, 0)
    mu    sync.Mutex
)

func GetAllTasks() []models.Task {
    mu.Lock()
    defer mu.Unlock()
    return tasks
}

func GetTaskByID(id string) (*models.Task, error) {
    mu.Lock()
    defer mu.Unlock()
    for _, task := range tasks {
        if task.ID == id {
            return &task, nil
        }
    }
    return nil, errors.New("task not found")
}
func UpdateTask(id string, newTask *models.Task)  {
    mu.Lock()
    defer mu.Unlock()
    for i, task := range tasks {
        if task.ID == id {
            tasks[i] = *newTask
        }
    }
    
}
func AddTask(task *models.Task) {
	mu.Lock()
	defer mu.Unlock()
	tasks = append(tasks, *task)
}

func DeleteTask(id string) {
    mu.Lock()
    defer mu.Unlock()
    for i, task := range tasks {
        if task.ID == id {
            tasks = append(tasks[:i], tasks[i+1:]...)
            break
        }
    }
}