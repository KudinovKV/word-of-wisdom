package messages

//go:generate go run github.com/dmarkham/enumer -type=ClientType -json
type ClientType int32

const (
	ClientTypeRequest ClientType = iota
	ClientTypeResponse
)

//go:generate go run github.com/dmarkham/enumer -type=ServerType -json
type ServerType int32

const (
	ServerTypeRequest ServerType = iota
	ServerTypeResponse
)
