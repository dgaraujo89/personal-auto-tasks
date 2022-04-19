package tasks

type Argument struct {
	Name         string
	Description  string
	DefaultValue string
}

type Flag struct {
	Name         string
	Description  string
	DataType     int
	DefaultValue interface{}
}

type Task struct {
	Name             string
	Description      string
	ShortDescription string
	Arguments        []Argument
	Flags            []Flag
	Action           TaskAction
}

type TaskAction interface {
	Run(args map[string]string, flags map[string]string)
}
