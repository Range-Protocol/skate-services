package disk

import (
	"database/sql"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
	"skatechain.org/lib/db"
	config "skatechain.org/relayer/db"

	"skatechain.org/lib/logging"
)

var (
	TaskLogger = db.NewFileLogger(config.DbDir, "skateapp_tasks.log")
	SkateAppDB *sql.DB
)

const (
	SignedTaskSchema    = "SignedTasks"
	CompletedTaskSchema = "CompletedTasks"
)

func init() {
	logger := logging.NewLoggerWithConsoleWriter()
	db, err := sql.Open("sqlite3", filepath.Join(config.DbDir, "skateapp.db"))
	if err != nil {
		logger.Fatal("Relayer DB initialization failed")
		panic("Relayer DB initialization failed")
	}
	SkateAppDB = db
	InitializeSkateApp()
}

type SignedTask struct {
	TaskId    uint32
	Message   string
	Initiator string
	ChainType uint32
	ChainId   uint32
	Hash      []byte
	Operator  string
	Signature []byte
}

type CompletedTask struct {
	TaskId    uint32
	ChainType uint32
	ChainId   uint32
}

func InitializeSkateApp() {
	// TODO: decouple this table, [operator/signature] references [remaining]
	SkateAppDB.Exec(`CREATE TABLE IF NOT EXISTS ` + SignedTaskSchema + ` (
		id           INTEGER PRIMARY KEY AUTOINCREMENT,
	  taskId       INTEGER,
	  message      TEXT,
	  initiator    TEXT,
	  chainId      INTEGER,
	  chainType    INTEGER,
	  hash         BLOB,
	  operator     TEXT,
	  signature    BLOB
	)`)

	SkateAppDB.Exec(`CREATE TABLE IF NOT EXISTS ` + CompletedTaskSchema + ` (
		id           INTEGER PRIMARY KEY AUTOINCREMENT,
	  taskId       INTEGER,
	  chainId      INTEGER,
	  chainType    INTEGER,

    UNIQUE (taskId, chainType, chainId)
	)`)
}

func InsertSignedTask(signedTask SignedTask) error {
	_, err := SkateAppDB.Exec(
		"INSERT OR IGNORE INTO "+SignedTaskSchema+" (taskId, message, initiator, chainId, chainType, hash, operator, signature) VALUES (?,?,?,?,?,?,?,?)",
		signedTask.TaskId, signedTask.Message, signedTask.Initiator, signedTask.ChainId,
		signedTask.ChainType, signedTask.Hash, signedTask.Operator, signedTask.Signature,
	)
	if err != nil {
		TaskLogger.Error("InsertSignedTask failed", "task", signedTask, "err", err)
		return err
	}
	return nil
}

func RetrieveSignedTasks() ([]SignedTask, error) {
	rows, err := SkateAppDB.Query("SELECT * FROM " + SignedTaskSchema)
	if err != nil {
		TaskLogger.Error("RetrieveSignedTask failed", "err", err)
		return nil, err
	}
	defer rows.Close()

	var resultTasks []SignedTask

	for rows.Next() {
		var task SignedTask
		var entryid int

		err := rows.Scan(
			&entryid, &task.TaskId, &task.Message, &task.Initiator,
			&task.ChainId, &task.ChainType, &task.Hash, &task.Operator, &task.Signature,
		)
		if err != nil {
			return nil, err
		}
		TaskLogger.Info("Task", "task", task)
		resultTasks = append(resultTasks, task)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return resultTasks, nil
}

func InsertCompletedTask(completedTask CompletedTask) error {
	_, err := SkateAppDB.Exec(
		"INSERT OR IGNORE INTO "+CompletedTaskSchema+" (taskId, chainId, chainType) VALUES (?,?,?)",
		completedTask.TaskId, completedTask.ChainId, completedTask.ChainType,
	)
	if err != nil {
		TaskLogger.Error("InsertCompletedTask failed", "task", completedTask, "err", err)
		return err
	}
	return nil
}
