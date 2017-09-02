package timedstore

import (
	"testing"
)

func TestEmptyStore(t *testing.T) {
	store := &Store{}

	if values := store.GetActiveNow("doesn't exist"); len(values) > 0 {
		t.Errorf("Empty Store should empty list of values but didn't. Returned %v as value", values)
	}
}

func TestBasicStorage(t *testing.T) {
	store := &Store{}

	currentTime := CurrentTime()
	const oneHour = 60 * 60

	const string1 = "data1"
	const string2 = "data2"
	const string3 = "data3"

	value1 := NewValue(string1, currentTime, currentTime+oneHour)
	value2 := NewValue(string2, currentTime-oneHour, currentTime+2*oneHour)
	value3 := NewValue(string3, currentTime+oneHour, currentTime+2*oneHour)

	store.Put("one", *value1)
	store.Put("one", *value2)
	store.Put("two", *value3)

	values := store.GetActive("one", currentTime+1)
	if len(values) != 2 {
		t.Errorf("Expected exactly 2 values returned. Got %v", len(values))
	}

	if values[0].Data != string1 {
		t.Errorf("First element was not as expected. Received '%v', expected '%s'", values[0].Data, string1)
	}

	if values[1].Data != string2 {
		t.Errorf("Second returned value not as expected. Got '%v', expected '%s'", values[1].Data, string2)
	}

	if values := store.GetActive("two", currentTime); len(values) != 0 {
		t.Errorf("Expected to receive exactly one value. Received '%v'", values)
	}

	if values := store.GetActive("two", currentTime+oneHour+1); len(values) != 1 {
		t.Errorf("Expected to receive exactly one value. Received '%v'", values)
	}
}
