package actor

type SystemName string

type System interface {
	NewActor(def Def) error
	Introduce(channel interface{}, ref Ref) (Ref, error)
	Shutdown()
}

type system struct {
	name SystemName
}

func (sys *system) NewActor(def Def) error {
	def.Behavior(nil, nil)
}

func (sys *system) Introduce(channel interface{}, ref Ref) (Ref, error) {
	panic("implement me")
}

func (sys *system) Shutdown() {
	panic("implement me")
}