package db

import (
	"database/sql"
	"io"
	"os"
	"path/filepath"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/zerolog"
	"skatechain.org/lib/logging"
)

var (
	DbDir      = "data/operator"
	SkateAppDB *sql.DB
	logger     = logging.NewLoggerWithConsoleWriter()
)

func init() {
	if err := os.MkdirAll(DbDir, os.ModePerm); err != nil {
		logger.Error("Create data directory false, please try manually", "data dir", DbDir)
		panic(err)
	}
	db, err := sql.Open("sqlite3", filepath.Join(DbDir, "skateapp.db"))
	db.Exec(`CREATE TABLE IF NOT EXISTS tasks (
		id       INTEGER PRIMARY KEY AUTOINCREMENT,
	  taskId   TEXT,
	  message  TEXT,
	  signer   TEXT,
	  chainId  INTEGER,
	  taskHash TEXT
	)`)
	if err != nil {
		panic("Task DB initialization failed")
	}
	SkateAppDB = db
}

func NewFileLogger(fileName string) *logging.Logger {
	logFile := createLogFile(fileName)
	plainWriter := zerolog.ConsoleWriter{Out: logFile, TimeFormat: time.RFC3339, NoColor: true}
	return logging.NewLogger(plainWriter)
}

func createLogFile(fileName string) io.Writer {
	logFilePath := filepath.Join(DbDir, fileName)
	if _, err := os.Stat(logFilePath); os.IsNotExist(err) {
		if err := os.MkdirAll(DbDir, os.ModePerm); err != nil {
			logger.Error("Create data directory false, please try manually", "data dir", DbDir)
			panic(err)
		}

		// Create the log file with sufficient permissions
		file, err := os.Create(logFilePath)
		if err != nil {
			panic(err)
		}
		return file
	}

	// Open the existing log file with append mode and sufficient permissions
	file, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		panic(err)
	}
	return file
}
