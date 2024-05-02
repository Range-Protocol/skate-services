package main

import (
	"skatechain.org/lib/logging"
	diskDb "skatechain.org/relayer/db/skateapp/disk"
)

func main() {
	logger := logging.NewLoggerWithConsoleWriter()
	tasks, _ := diskDb.SkateApp_RetriveSignedTasks("SELECT * FROM skateapp_signed_tasks")
	for _, task := range tasks {
    logger.Info("Task", "task", task)
	}
}
