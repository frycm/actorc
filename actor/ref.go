package actor

type RefPath []string

type RefProtocol string

const (
	Local = RefProtocol("local")
)

type Ref struct {
	protocol RefProtocol
	system SystemName
	path RefPath
}

func UserRef(components... string) Ref {
	return Ref{protocol:Local, path: append([]string{"user"}, components...)}
}