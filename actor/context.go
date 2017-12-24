package actor

type Context struct {
	system *system
	self Ref
	Signal <-chan Signal
}

func (c * Context) Pass(msg interface{}) {
	// TODO implement, send it to DLQ Actor
}

type SignalType int

const (
	InitSignal SignalType = iota
	StopSignal
	RebornSignal
)

type Signal interface {
	Type() SignalType
}
