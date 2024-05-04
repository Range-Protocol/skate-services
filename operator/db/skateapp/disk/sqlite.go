package disk

import (
	// "encoding/hex"
	"math/big"
	"database/sql"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"

	"github.com/ethereum/go-ethereum/common"
	"skatechain.org/contracts/bindings/SkateApp"
	"skatechain.org/lib/db"
	config "skatechain.org/operator/db"
)

var (
	SkateAppDB *sql.DB
	TaskLogger = db.NewFileLogger(config.DbDir, "skateapp_tasks.log")
)

func init() {
	db, err := sql.Open("sqlite3", filepath.Join(config.DbDir, "skateapp.db"))
	if err != nil {
		panic("Relayer DB initialization failed")
	}
	SkateAppDB = db
	InitializeSkateApp()
}


const TaskSchema = "Tasks"

type Task struct {
	Id      int64
	Message string
	Chain   uint32
	Signer  string
	Hash    []byte
}

func InitializeSkateApp() {
	SkateAppDB.Exec(`CREATE TABLE IF NOT EXISTS ` + TaskSchema + ` (
		id       INTEGER PRIMARY KEY AUTOINCREMENT,
	  taskId   TEXT,
	  message  TEXT,
	  signer   TEXT,
	  chainId  INTEGER,
	  hash     BLOB
	)`)
}

type bindingTask = bindingSkateApp.BindingSkateAppTaskCreated

func task_dbToBinding(task *Task) *bindingTask {
	return &bindingTask{
		TaskId:   big.NewInt(task.Id),
		Message:  task.Message,
		Signer:   common.HexToAddress(task.Signer),
		TaskHash: [32]byte(task.Hash),
		Chain:    task.Chain,
	}
}

func task_bindingToDb(task *bindingSkateApp.BindingSkateAppTaskCreated) *Task {
	return &Task{
		Id:      task.TaskId.Int64(),
		Message: task.Message,
		Signer:  task.Signer.Hex(),
		Hash:    task.TaskHash[:],
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
