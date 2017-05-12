package actor

func NewLocalSystem(systemName SystemName) (System, error) {
	return nil, nil // TODO implement
}

func NewDefaultLocalSystem() (System, error) {
	return NewLocalSystem("DefaultSystem")
}
