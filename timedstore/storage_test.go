package timedstore

import (
  "fmt"
  "testing"
  "container/list"
)

// func TestEmptyStore(t *testing.T) {
//   store := &Store{}

//   if values, ok := store.Get("doesn't exist") ; ok {
//     t.Errorf("Empty Store should return ok == false but didn't. Returned %v as value", values)
//   }
// }

// func TestBasicStorage(t *testing.T) {
//   store := &Store{}

//   currentTime := CurrentTime()
//   const oneHour = 60 * 60

//   const string1 = "data1"
//   const string2 = "data2"
//   const string3 = "data3"

//   value1 := NewValue(string1, currentTime, currentTime + oneHour)
//   value2 := NewValue(string2, currentTime - oneHour, currentTime + 2 * oneHour)
//   value3 := NewValue(string3, currentTime + oneHour, currentTime + 2 * oneHour)

//   store.Put("one", *value1)
//   store.Put("one", *value2)
//   store.Put("two", *value3)

//   if values, ok := store.Get("one") ; ok {
//     element := values.Front()
//     if element.Value != string1 {
//       t.Errorf("Received '%s', expected '%s'", element.Value, string1)
//     }

//     element = element.Next()
//     if element == nil || element.Value != string2 {
//       t.Errorf("Did not received expected second element '%s'", string2)
//     }
//   } else {
//     t.Errorf("Expected to Get some values but didn't")
//   }

//   if values, ok := store.Get("two") ; ok && values.Len() != 0 {
//     t.Errorf("Expected to receive no values, but received '%v'", values)
//   }
// }

// Make sure that values that have expired are removed from the storage
// so as to not have a memory leak. They should be cleaned up after each
// Get operation.
func TestRemoveExpiredValues(t *testing.T) {
  fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~TestRemoveExpiredValues~~~~~~~~~~~")
  store := &Store{}

  duration := int64(1000)
  expiredTime := CurrentTime() - (duration * 2)

  value1 := NewValue("expired data", expiredTime, expiredTime + duration)
  //value2 := NewValue("expired data 2", expiredTime, expiredTime + duration)
  key1 := "one"

  store.Put(key1, *value1)
  //store.Put(key1, *value2)

  if values, ok := store.GetAllValues(key1) ; !ok || values.Len() != 1 {
    t.Errorf("Stored value does not appear in list of values")
  }

  _, _ = store.Get(key1)
  
  if values, ok := store.GetAllValues(key1) ; !ok || values.Len() != 0 {
    t.Errorf("Expected all values to have expired, but still have values: '%v'", values)
  }

  fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~TestRemoveExpiredValues~~~~~~~~~~~")
}

func (s Store) GetAllValues(k Key) (values list.List, ok bool) {
  values, ok = s[k]
  fmt.Println("------------\nValues:")
  for e := values.Front(); e != nil; e = e.Next() {
    fmt.Printf("\t%v\n", e.Value)
  }
  fmt.Println("------------")
  return
}
