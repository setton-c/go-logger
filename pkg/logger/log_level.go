package logger

type Log_level int

const (
	INFO Log_level = iota
	WARNING
	ERROR
)

func (l Log_level) String() string {
	switch l {
	case INFO:
		return "info"
	case WARNING:
		return "warning"
	case ERROR:
		return "error"
	default:
		return "unknown"
	}
}
