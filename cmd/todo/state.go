package todo

type State string

const (
	Added     State = "Added"
	InProcess State = "In process"
	Done      State = "Done"
)

func (s State) ToString() string {
	switch s {
	case Added:
		return "added"
	case InProcess:
		return "in-process"
	case Done:
		return "done"
	default:
		return ""
	}
}
