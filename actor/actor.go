package actor

type Name []string

type ReactionHandler func(Message) (Message, error)

type Strategy func(Message, error) ()

type Dna struct {
	Name Name
	ReactionFunc ReactionHandler
}

type incarnation interface {

}