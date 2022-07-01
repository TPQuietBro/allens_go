package registraty

type Registration struct {
	ServiceName ServiceName
	ServiceUrl  string
}
type ServiceName string

const (
	LogService = ServiceName("LogService")
)
