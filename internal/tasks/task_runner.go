package tasks

import "github.com/thatisuday/commando"

const VERSION = "0.0.1"

func Start(tasks []Task) {
	commando.SetExecutableName("personal-auto-tasks").
		SetVersion(VERSION).
		SetDescription("Tools for automate tasks")

	registerTasks(tasks)

	commando.Parse(nil)
}

func registerTasks(tasks []Task) {
	for indexTasks := range tasks {
		command := commando.Register(tasks[indexTasks].Name).
			SetDescription(tasks[indexTasks].Description).
			SetShortDescription(tasks[indexTasks].ShortDescription)

		for indexArguments := range tasks[indexTasks].Arguments {
			command.AddArgument(tasks[indexTasks].Arguments[indexArguments].Name,
				tasks[indexTasks].Arguments[indexArguments].Description,
				tasks[indexTasks].Arguments[indexArguments].DefaultValue)
		}

		for indexFlags := range tasks[indexTasks].Flags {
			command.AddFlag(tasks[indexTasks].Flags[indexFlags].Name,
				tasks[indexTasks].Flags[indexFlags].Description,
				tasks[indexTasks].Flags[indexFlags].DataType,
				tasks[indexTasks].Flags[indexFlags].DefaultValue)
		}
	}
}
