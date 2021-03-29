package demon

//IEvent ...
type IEvent interface {
	Run() error
}

const (
	RollBackEvent = iota
)