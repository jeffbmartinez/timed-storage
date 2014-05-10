package timedstore

import (
  "container/list"
  "fmt"
)

type Key interface {}

type Store map[Key]list.List

// Returns all active values stored by the key. Also cleans expired
// Values. Values are considered expired if the current time is greater
// than or equal the value's end time.
func (s Store) Get(k Key) (activeValues *list.List, ok bool) {
  values, ok := s[k]

  fmt.Printf("s[k]:\n\t%v\n", s[k])
  fmt.Printf("List:\n\t%v\n", values)

  if ok {
    activeValues = list.New()
    expiredElements := make([]*list.Element, 0, 10)

    currentTime := CurrentTime()

    for element := values.Front() ; element != nil ; element = element.Next() {      
      if value, ok := element.Value.(Value) ; ok {
        if value.IsActiveForTime(currentTime) {
          activeValues.PushBack(value.Data)
        } else if value.IsExpiredForTime(currentTime) {
          // Keep track of list elements to remove later. They cannot be removed
          // during iteration.
          //expiredElements = append(expiredElements, element)
          //values.MoveToFront(element)
          fmt.Printf("Added to expiredElements:\n\t%v\n", element)
        }
      }
    }

    fmt.Printf("expiredElements length: %v\n", len(expiredElements))

    // Go back and remove any elements that are expired. These can never become
    // active again so they are removed from storage
    for i := 0 ; i < len(expiredElements) ; i++ {
      value := values.Front()
      fmt.Printf("Expired element removed:\n\t%v\n", value)
      //values.MoveToFront(element)
      values.Remove(value)
    }

    s[k] = values
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
