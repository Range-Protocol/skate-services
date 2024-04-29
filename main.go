package main

import (
	"skatechain.org/lib/logging"
	operatorDb "skatechain.org/operator/db"
)

func main() {
	logger := logging.NewLoggerWithConsoleWriter()
	task := &operatorDb.Task{
		Id:      1,
		Message: "I'm not fine",
		Signer:  "0xe254a253Cd380d67c57EA20B97ef60678E379eB0",
		Chain:   1,
		Hash:    "0xe254a253Cd380d67c57EA20B97ef60678E379eB0",
	}
	defer operatorDb.SkateAppDB.Close()
	logger.Info("Insert", "task", task)
	operatorDb.InsertTask(task)
	tasks, _ := operatorDb.SelectAllTasks()
	for _, task := range tasks {
		logger.Info("Got tasks", "task", task)
	}
}
