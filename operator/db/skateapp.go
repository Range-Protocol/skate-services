package db

import (
	"encoding/hex"

	"skatechain.org/contracts/bindings/SkateApp"
)

var TaskLogger = NewFileLogger("tasks.log")

type Task struct {
	Id      uint64
	Message string
	Chain   uint32
	Signer  string
	Hash    string
}

func SkateAppTaskToDbTask(task *bindingSkateApp.BindingSkateAppTaskCreated) *Task {
	return &Task{
		Id:      task.TaskId.Uint64(),
		Message: task.Message,
		Signer:  task.Signer.Hex(),
		Hash:    hex.EncodeToString(task.TaskHash[:]),
		Chain:   task.Chain,
	}
}

func InsertTask(task *Task) error {
	_, err := SkateAppDB.Exec(
		"INSERT INTO tasks (taskId, message, signer, chainId, taskHash) VALUES (?, ?, ?, ?, ?)",
		task.Id, task.Message, task.Signer, task.Chain, task.Hash,
	)
	if err != nil {
    TaskLogger.Error("InsertTask failed", "task", task, "err", err)
		return err
	}
	return nil
}

func SelectAllTasks() ([]Task, error) {
	rows, err := SkateAppDB.Query("SELECT * FROM tasks")
	if err != nil {
		TaskLogger.Error("SelectAllTasks failed", "err", err)
		return nil, err
	}
	defer rows.Close()

	var tasks []Task

	for rows.Next() {
		var task Task
    var entryid int

		err := rows.Scan(&entryid, &task.Id, &task.Message, &task.Signer, &task.Chain, &task.Hash)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}
