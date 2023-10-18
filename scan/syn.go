package scan

import "time"

type SynScanner struct {
	timeout     time.Duration
	maxRoutines int
}
