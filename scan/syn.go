package scan

import "time"

type synScanner struct {
	timeout     time.Duration
	maxRoutines int
}
