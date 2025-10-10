package status

type Status int

const (
	TODO Status = iota
	IN_PROGRESS
	DONE
)

func (s Status) String() string {
	return [...]string{"todo", "inProgress", "done"}[s]
}
