# go-actor-chan

Actor based channels for Go.

Differences in comparison with plain Go channels:
* Input channel is not processed by Actor directly but by Dispatcher
* actor.Ref owning input channel is abstraction over Actor location and type
* Various Dispatcher and Inbox types could be used
* actor.System manages all Actors and allow cooperation with other Actors in cluster

## Current status

In development (prototype phase).

## Remaining tasks for minimal version

- [ ] System, Ref, Dispatcher, Actor and Message public API draft (in progress)
- [ ] Default local API implementation
- [ ] Hello world working
- [ ] Cluster support
- [ ] Common actors implementation (routing, circuit breaking, bulk heading, ...)