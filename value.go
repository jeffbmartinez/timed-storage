package timedstore

// This is the struct used to internally represent Value objects, which are
// only active during a given time range. The start time and end time are both
// provided in seconds since January 1, 1970, also known as unix time or epoch
// time.
type Value struct {
	Data interface{}

	StartSeconds int64
	EndSeconds   int64
}

// Creates a new Value given a start time and end time. Times are provided in
// seconds since January 1, 1970 UTC (unix/epoch time)
func NewValue(data interface{}, startSeconds int64, endSeconds int64) *Value {
	return &Value{data, startSeconds, endSeconds}
}

// Creates a new Value given a start time and a duration. Start time is provided in
// seconds since January 1, 1970 UTC (unix/epoch time). Duration is also in seconds.
func NewValueFromDuration(data interface{}, startSeconds int64, duration int64) *Value {
	endSeconds := startSeconds + duration
	return NewValue(data, startSeconds, endSeconds)
}

// Returns the duration of this value in seconds. Calculated by subtracting the
// start time from the end time.
func (v Value) Duration() int64 {
	return v.EndSeconds - v.StartSeconds
}

// Determine if this value is active during the provided time. It is inclusive of the
// start time and exclusive of the end time.
func (v Value) IsActiveForTime(time int64) bool {
	return time >= v.StartSeconds && time < v.EndSeconds
}

// Determine if this value's active period is later than the provided time. Can
// be used to determine if a value's useful time period has expired.
func (v Value) IsExpiredForTime(time int64) bool {
	return time >= v.EndSeconds
}
