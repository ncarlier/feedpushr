package output

// Status is the process status
type Status uint8

const (
	// RunningStatus is the status of a running process
	RunningStatus Status = iota
	// StoppedStatus is the status of a stoppedprocess
	StoppedStatus
)

// String converts a status to a readable string
func (s Status) String() string {
	switch s {
	case RunningStatus:
		return "running"
	case StoppedStatus:
		return "stopped"
	}
	return ""
}
