package timedstore

import "time"

// Returns current UTC time in seconds since January 1, 1970.
// This is also known as Unix time or epoch time.
func CurrentTime() int64 {
  return time.Now().UTC().Unix()
}
