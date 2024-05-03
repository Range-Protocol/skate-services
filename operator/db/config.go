package db

import (
	"database/sql"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
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
	if err != nil {
		panic("Relayer DB initialization failed")
	}
	SkateAppDB = db
  InitializeSkateApp()
}
