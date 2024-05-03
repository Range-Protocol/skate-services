package db

import (
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"skatechain.org/contracts/bindings/SkateApp"
	"skatechain.org/lib/db"
)

var (
	TaskLogger = db.NewFileLogger(DbDir, "skateapp_tasks.log")
)

const TaskSchema = "Tasks"

type Task struct {
	Id      int64
	Message string
	Chain   uint32
	Signer  string
	Hash    string
}

func InitializeSkateApp() {
	SkateAppDB.Exec(`CREATE TABLE IF NOT EXISTS ? (
		id       INTEGER PRIMARY KEY AUTOINCREMENT,
	  taskId   TEXT,
	  message  TEXT,
	  signer   TEXT,
	  chainId  INTEGER,
	  hash     TEXT
	)`, TaskSchema)
}

type bindingTask = bindingSkateApp.BindingSkateAppTaskCreated

func task_dbToBinding(task *Task) *bindingTask {
	hashBytes, _ := hex.DecodeString(task.Hash)
	if len(hashBytes) != 32 {
		panic(fmt.Sprintf("Malformed hash, must be 32 bytes, got %d", len(hashBytes)))
	}
	return &bindingTask{
		TaskId:   big.NewInt(task.Id),
		Message:  task.Message,
		Signer:   common.HexToAddress(task.Signer),
		TaskHash: [32]byte(hashBytes),
		Chain:    task.Chain,
	}
}

func task_bindingToDb(task *bindingSkateApp.BindingSkateAppTaskCreated) *Task {
	return &Task{
		Id:      task.TaskId.Int64(),
		Message: task.Message,
		Signer:  task.Signer.Hex(),
		Hash:    hex.EncodeToString(task.TaskHash[:]),
		Chain:   task.Chain,
	}
}

func SkateApp_InsertTask(bindingTask *bindingSkateApp.BindingSkateAppTaskCreated) error {
	task := task_bindingToDb(bindingTask)
	_, err := SkateAppDB.Exec(
		"INSERT INTO ? (taskId, message, signer, chainId, hash) VALUES (?, ?, ?, ?, ?)",
    TaskSchema,
		task.Id, task.Message, task.Signer, task.Chain, task.Hash,
	)
	if err != nil {
		TaskLogger.Error("InsertTask failed", "task", task, "err", err)
		return err
	}
	return nil
}

func SkateApp_SelectTasks(query string, args ...any) ([]bindingTask, error) {
	rows, err := SkateAppDB.Query(query, args)
	if err != nil {
		TaskLogger.Error("SelectAllTasks failed", "err", err)
		return nil, err
	}
	defer rows.Close()

	var bindingTasks []bindingTask

	for rows.Next() {
		var task Task
		var entryid int

		err := rows.Scan(&entryid, &task.Id, &task.Message, &task.Signer, &task.Chain, &task.Hash)
		if err != nil {
			return nil, err
		}
		TaskLogger.Info("Task", "task", task)
		bindingTask := task_dbToBinding(&task)
		bindingTasks = append(bindingTasks, *bindingTask)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return bindingTasks, nil
}
