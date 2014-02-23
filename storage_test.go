package timedstore

import (
  "testing"
)

func TestEmptyStore(t *testing.T) {
  store := &Store{}

  if values, ok := store.Get("doesn't exist") ; ok {
    t.Errorf("Empty Store should return ok == false but didn't. Returned %v as value", values)
  }
}

func TestBasicStorage(t *testing.T) {
  store := &Store{}

  currentTime := CurrentTime()
  const oneHour = 60 * 60

  const string1 = "data1"
  const string2 = "data2"
  const string3 = "data3"

  value1 := NewValue(string1, currentTime, currentTime + oneHour)
  value2 := NewValue(string2, currentTime - oneHour, currentTime + 2 * oneHour)
  value3 := NewValue(string3, currentTime + oneHour, currentTime + 2 * oneHour)

  store.Put("one", *value1)
  store.Put("one", *value2)
  store.Put("two", *value3)

  if values, ok := store.Get("one") ; ok {
    element := values.Front()
    if element.Value != string1 {
      t.Errorf("Received '%s', expected '%s'", element.Value, string1)
    }

    element = element.Next()
    if element == nil || element.Value != string2 {
      t.Errorf("Did not received expected second element '%s'", string2)
    }
  } else {
    t.Errorf("Expected to Get some values but didn't")
  }

  if values, ok := store.Get("two") ; ok && values.Len() != 0 {
    t.Errorf("Expected to receive no values, but received '%v'", values)
  }
}
