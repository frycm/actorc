package actor

type Def struct {
	In       interface{}
	Params   interface{}
	Behavior interface{}
}

func newActorContext(system *system, def Def, ref Ref) (Context, error) {
	panic("TODO implement")
	return Context{}, nil
}