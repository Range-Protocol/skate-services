package db

import (
	// "encoding/hex"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"skatechain.org/contracts/bindings/SkateApp"
	"skatechain.org/lib/db"
)

var TaskLogger = db.NewFileLogger(DbDir, "skateapp_tasks.log")

const TaskSchema = "Tasks"

type Task struct {
	Id      int64
	Message string
	Chain   uint32
	Signer  string
	Hash    [32]byte
}

func InitializeSkateApp() {
	SkateAppDB.Exec(`CREATE TABLE IF NOT EXISTS ` + TaskSchema + ` (
		id       INTEGER PRIMARY KEY AUTOINCREMENT,
	  taskId   TEXT,
	  message  TEXT,
	  signer   TEXT,
	  chainId  INTEGER,
	  hash     TEXT
	)`)
}

type bindingTask = bindingSkateApp.BindingSkateAppTaskCreated

func task_dbToBinding(task *Task) *bindingTask {
	return &bindingTask{
		TaskId:   big.NewInt(task.Id),
		Message:  task.Message,
		Signer:   common.HexToAddress(task.Signer),
		TaskHash: task.Hash,
		Chain:    task.Chain,
	}
}

func task_bindingToDb(task *bindingSkateApp.BindingSkateAppTaskCreated) *Task {
	return &Task{
		Id:      task.TaskId.Int64(),
		Message: task.Message,
		Signer:  task.Signer.Hex(),
		Hash:    task.TaskHash,
		Chain:   task.Chain,
	}
}

func SkateApp_InsertTask(bindingTask *bindingSkateApp.BindingSkateAppTaskCreated) error {
	task := task_bindingToDb(bindingTask)
	_, err := SkateAppDB.Exec(
		"INSERT INTO "+TaskSchema+" (taskId, message, signer, chainId, hash) VALUES (?, ?, ?, ?, ?)",
		task.Id, task.Message, task.Signer, task.Chain, task.Hash,
	)
	if err != nil {
		TaskLogger.Error("InsertTask failed", "task", task, "err", err)
		return err
	}
	return nil
}

func SkateApp_SelectTasks() ([]bindingTask, error) {
	rows, err := SkateAppDB.Query("SELECT * FROM " + TaskSchema)
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
