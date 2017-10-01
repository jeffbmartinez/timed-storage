package timedstore

type Key interface{}

// type Store map[Key][]Value

type Store struct {
	unique  map[string]*Value
	storage map[Key][]*Value
}

func NewStore() Store {
	return Store{
		map[string]*Value{},
		map[Key][]*Value{},
	}
}

/* GetActive returns all active values stored by the key and the provided time. Values are
considered expired (not active) if the specified time is greater than or equal to the value's
end time.

t is in seconds since January 1, 1970 UTC (unix/epoch time)

If the key is invalid, an empty list is still returned.
*/
func (s Store) GetActive(k Key, t int64) (activeValues []Value) {
	values, ok := s.storage[k]
	if !ok {
		return
	}

	for _, value := range values {
		if value.IsActiveForTime(t) {
			activeValues = append(activeValues, *value)
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
	values, ok := s.storage[k]
	if !ok {
		return []Value{}
	}

	expiredValues := []Value{}
	nonExpiredValues := []*Value{}

	for _, value := range values {
		if value.IsExpiredForTime(t) {
			expiredValues = append(expiredValues, *value)
			delete(s.unique, value.GetUniqueID())
		} else {
			nonExpiredValues = append(nonExpiredValues, value)
		}
	}

	s.storage[k] = nonExpiredValues

	return expiredValues
}

// Put adds a new Value behind a key. This does not replace an existing value,
// it adds to a list of stored values. A unique ID is returned which can
// be used to recall the speciic value with GetUnique().
func (s Store) Put(k Key, v Value) {
	values := s.storage[k]
	values = append(values, &v)
	s.storage[k] = values

	s.unique[v.GetUniqueID()] = &v
}

// GetUnique retrieves the Value with the provided unique ID.
// It returns a value, ok pair, just like a regular map lookup in Go.
func (s Store) GetUnique(id string) (*Value, bool) {
	value, ok := s.unique[id]
	return value, ok
}
