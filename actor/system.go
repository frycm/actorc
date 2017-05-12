package actor

type SystemName string

type System interface {
	Introduce(name... string) Ref
	Register(dna Dna) (Ref, error)
	Shutdown()
}
