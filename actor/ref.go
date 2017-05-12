package actor

type Ref struct {
	Demand <-chan uint
	Tell chan<- Message
}

func (ref *Ref) NewResponseChan() (chan Message) {
	return nil
}