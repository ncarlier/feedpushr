package aggregator

import "time"

const h24 = time.Duration(24) * time.Hour

// Action is used to manage a process
type Action uint8

const (
	// StartAction is the action to start a process
	StartAction Action = iota
	// StopAction is the action to stop a process
	StopAction
)

// String converts an action to a readable string
func (a Action) String() string {
	switch a {
	case StartAction:
		return "start"
	case StopAction:
		return "stop"
	}
	return ""
}

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

// FeedStatus status of a feed
type FeedStatus struct {
	CheckedAt          time.Time `json:"checked_at"`
	EtagHeader         string    `json:"etag_header"`
	LastModifiedHeader string    `json:"last_modified_header"`
	ExpiresHeader      string    `json:"expires_header"`
	ErrorMsg           string    `json:"error_message"`
	ErrorCount         int       `json:"error_count"`
}

// Err set error status message and counter
// Reset if err is null
func (fs *FeedStatus) Err(err error) {
	if err != nil {
		fs.ErrorCount++
		fs.ErrorMsg = err.Error()
	} else {
		fs.ErrorCount = 0
		fs.ErrorMsg = ""
	}
}

// ComputeNextCheckDate computes next dat to check regarding  some rules
// The duration is multiply the number of error.
// The limit is h24
func (fs *FeedStatus) ComputeNextCheckDate(base time.Duration) time.Time {
	if fs.ErrorCount > 0 {
		base = base * time.Duration(fs.ErrorCount)
	}
	if base > h24 {
		base = h24
	}
	return fs.CheckedAt.Add(base)
}
