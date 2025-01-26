package enums

type TaskStatus int

const (
	InProgress TaskStatus = iota
	ToDo
	Completed
)

func (t TaskStatus) String() string {
	switch t {
	case InProgress:
		return "in_progress"
	case ToDo:
		return "to_do"
	case Completed:
		return "completed"
	default:
		return "unknown"
	}
}
