package asyncjob

type JobState int

const (
	StateInit JobState = iota
	StatePending
	StateQueued
	StateProcessing
	StateTimeout
	StateFailed
	StateCompleted
	StateRetryFailed
)

func (j JobState) String() string {
	return []string{"Init", "Pending", "Queued", "Processing", "Timeout", "Failed", "Completed", "RetryFailed"}[j]
}
