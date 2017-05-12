package actor_test

import (
	"fmt"

	"github.com/frycm/go-actor-chan/actor"
)

func ExampleActorTell() {
	system, err := actor.NewDefaultLocalSystem()
	checkAndPanic(err)
	defer system.Shutdown()

	greeter, err := system.Register(actor.Dna{Name: actor.Name{"Greeter"}, ReactionFunc: func(msg actor.Message) (actor.Message, error) {
		if name, ok := msg.Intention.(string); ok {
			fmt.Printf("Hello, %s!", name)
			return nil, nil
		} else {
			return msg, nil
		}
	}})

	<-greeter.Demand
	greeter.Tell <- actor.Message{Intention: "Martin"}
}

func ExampleActorAsk() {
	system, err := actor.NewDefaultLocalSystem()
	checkAndPanic(err)
	defer system.Shutdown()

	greeter, err := system.Register(actor.Dna{Name: actor.Name{"Repeater"}, ReactionFunc: func(msg actor.Message) (actor.Message, error) {
		if text, ok := msg.Intention.(string); ok {
			if msg.ResponseChan != nil {
				msg.ResponseChan <- actor.Message{Intention: text}
			}
			return nil, nil
		} else {
			return msg, nil
		}
	}})

	<-greeter.Demand
	responseChan := greeter.NewResponseChan()
	greeter.Tell <- actor.Message{Intention: "Hello World", ResponseChan: responseChan}

	println(<-responseChan)
}

func checkAndPanic(err error) {
	if err != nil {
		panic(err)
	}
}
