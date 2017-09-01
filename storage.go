package timedstore

import (
  "container/list"
)

type Key interface {}

type Store map[Key]list.List

// Returns all active values stored by the key. Also cleans expired
// Values. Values are considered expired if the current time is greater
// than or equal the value's end time.
func (s Store) Get(k Key) (activeValues *list.List, ok bool) {
  values, ok := s[k]

  if ok {
    activeValues = list.New()
    expiredValues := list.New()

    currentTime := CurrentTime()

    for element := values.Front() ; element != nil ; element = element.Next() {      
      if value, ok := element.Value.(Value) ; ok {
        if value.IsActiveForTime(currentTime) {
          activeValues.PushBack(value.Data)
        } else if value.IsExpiredForTime(currentTime) {
          // Keep track of list elements to remove later. They cannot be removed
          // during iteration.
          expiredValues.PushBack(element)
        }
      }
    }

    // Go back and remove any elements that are expired. These can never become
    // active again so they are removed from storage
    for element := expiredValues.Front() ; element != nil ; element = element.Next() {
      expiredElement, ok := element.Value.(*list.Element)

      if ok {
        expiredValues.Remove(expiredElement)
      }
    }
  }

  return
}

// Adds a new value behind a key. This does not replace an existing value,
// it adds to a list of stored values.
func (s Store) Put(k Key, v Value) {
  values := s[k]
  values.PushBack(v)
  s[k] = values
}
