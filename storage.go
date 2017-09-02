package timedstore

type Key interface{}

type Store map[Key][]Value

/* Returns all active values stored by the key and the provided time. Values are
considered expired (not active) if the specified time is greater than or equal to the value's
end time.

t is in seconds since January 1, 1970 UTC (unix/epoch time)

If the key is invalid, an empty list is still returned.
*/
func (s Store) GetActive(k Key, t int64) (activeValues []Value) {
	values, ok := s[k]
	if !ok {
		return
	}

	for _, value := range values {
		if value.IsActiveForTime(t) {
			activeValues = append(activeValues, value)
		}
	}

	return
}

// Same as GetActive using the current time as the time.
func (s Store) GetActiveNow(k Key) (activeValues []Value) {
	currentTime := CurrentTime()
	return s.GetActive(k, currentTime)
}

/* Remove all values that are expired for a given key and time. Expired values
are ones that have an end time greater than or equal to the provided time.

Returns a slice of Values that were expired and removed, if any.
If the key does not exist nothing is done and an empty slice is returned.
*/
func (s Store) RemoveExpiredForTime(k Key, t int64) []Value {
	values, ok := s[k]
	if !ok {
		return []Value{}
	}

	activeValues := []Value{}
	expiredValues := []Value{}

	for _, value := range values {
		if value.IsActiveForTime(t) {
			activeValues = append(activeValues, value)
		} else {
			expiredValues = append(expiredValues, value)
		}
	}

	s[k] = activeValues

	return expiredValues
}

// Adds a new Value behind a key. This does not replace an existing value,
// it adds to a list of stored values.
func (s Store) Put(k Key, v Value) {
	values := s[k]
	values = append(values, v)
	s[k] = values
}
