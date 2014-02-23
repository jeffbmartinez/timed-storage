package timedstore

import (
  "time"
  "container/list"
)

type Key interface {}

type Store map[Key]list.List

// Returns all active values stored by the key
func (s Store) Get(k Key) (activeValues *list.List, ok bool) {
  values, ok := s[k]

  if ok {
    activeValues = list.New()
    expiredValues := list.New()

    currentTime := time.Now().UTC().Unix()

    for element := values.Front() ; element != nil ; element = element.Next() {
      value, ok := element.Value.(Value)
      if ok {
        if value.IsActiveForTime(currentTime) {
          activeValues.PushBack(value.Data)
        } else if value.IsExpiredForTime(currentTime) {
          expiredValues.PushBack(element) // Keep track of values to be removed
        }
      }
    }

    // Go back and remove any elements that are expired. These can never become
    // active again so they are removed from storage
    for element := values.Front() ; element != nil ; element = element.Next() {
      expiredElement, ok := element.Value.(*list.Element)

      if ok {
        values.Remove(expiredElement)
      }
    }
  }

  return
}

// Adds a new value behind a key
func (s Store) Put(k Key, v Value) {
  
}
