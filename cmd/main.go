package main

import (
	"github.com/dgaraujo89/personal-auto-tasks/internal/tasks"
	"github.com/thatisuday/commando"
)

func main() {
	taskList := []tasks.Task{
		{
			Name:             "test",
			Description:      "Task test",
			ShortDescription: "Task test",
			Arguments: []tasks.Argument{
				{
					Name:        "start",
					Description: "Start command",
				},
			},
			Flags: []tasks.Flag{
				{
					Name:        "token",
					Description: "Token API",
					DataType:    commando.String,
				},
			},
		},
	}

	tasks.Start(taskList)
}
