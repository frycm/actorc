package actor_test

import (
	"github.com/frycm/go-actor-chan/actor"
	"sync"
)

func ExampleActorTell() {
	system, shutdown, err := createExampleSystem()
	defer shutdown()
	orPanic(err)

	var wg sync.WaitGroup

	{
		behaviorFunc := func(ctx actor.Context, in <-chan string) error {
			defer func (){
				wg.Done()
			}()

			loop: for {
				select {
				case msg := <-in:
					println("Hello, ", msg)
					break loop
				case s := <-ctx.Signal:
					if s.Type() == actor.StopSignal {
						break loop
					} else {
						ctx.Pass(s)
					}
				}
			}
			return nil
		}
		wg.Add(1)
		err := system.NewActor(actor.Def{In: make(<-chan string), Behavior: behaviorFunc})
		orPanic(err)
	}

	{
		greeterC := make(chan string)
		_, err := system.Introduce(greeterC, actor.UserRef("greeter"))
		orPanic(err)
		greeterC <- "Martin"
	}

	wg.Wait()
}

func ExampleActorAsk() {
	_, shutdown, err := createExampleSystem()
	defer shutdown()
	orPanic(err)
}

func createExampleSystem() (actor.System, func(), error) {
	system, err := actor.NewDefaultLocalSystem()
	if err != nil {
		return nil, func() {}, err
	}
	shutdownFunc := func() {
		system.Shutdown()
	}

	return system, shutdownFunc, nil
}

func orPanic(err error) {
	if err != nil {
		panic(err)
	}
}
