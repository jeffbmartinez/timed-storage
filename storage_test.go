package timedstore

import (
	"testing"
)

func TestEmptyStore(t *testing.T) {
	store := NewStore()

	if values := store.GetActiveNow("doesn't exist"); len(values) > 0 {
		t.Errorf("Empty Store should empty list of values but didn't. Returned %v as value", values)
	}
}

func TestBasicStorage1(t *testing.T) {
	store := NewStore()

	currentTime := CurrentTime()
	const oneHour = 60 * 60

	const string1 = "data1"
	const string2 = "data2"
	const string3 = "data3"

	value1 := NewValue(currentTime, currentTime+oneHour, string1)
	value2 := NewValue(currentTime-oneHour, currentTime+2*oneHour, string2)
	value3 := NewValue(currentTime+oneHour, currentTime+2*oneHour, string3)

	store.Put("one", value1)
	store.Put("one", value2)
	store.Put("two", value3)

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

func TestBasicStorage2(t *testing.T) {
	store := NewStore()

	const timeBefore = 100
	const timeEventStart = 200
	const timeDuring = 250
	const timeEventEnd = 300
	const timeAfter = 400

	const key = "key"
	const value = "value"

	store.Put(key, NewValue(timeEventStart, timeEventEnd, value))

	if values := store.GetActive(key, timeBefore); len(values) != 0 {
		t.Errorf("Expected zero values returned. Got '%v'", values)
	}

	if values := store.GetActive(key, timeDuring); values[0].Data != value {
		t.Errorf("Expected '%s'. Got '%v'.", value, values[0].Data)
	}

	if values := store.GetActive(key, timeAfter); len(values) != 0 {
		t.Errorf("Expected zero values returned. Got '%v'", values)
	}
}

func TestBasicStorage3(t *testing.T) {
	store := NewStore()

	const key = "key"
	const value1 = "value1"
	const value2 = "value2"
	const value3 = "value3"

	store.Put(key, NewValue(100, 400, value1))
	store.Put(key, NewValue(200, 500, value2))
	store.Put(key, NewValue(300, 600, value3))

	if values := store.GetActive(key, 350); len(values) != 3 {
		t.Errorf("Expected exactly 3 values. Got %v: '%v'", len(values), values)
	}

	if values := store.GetActive(key, 250); len(values) != 2 {
		t.Errorf("Expected exactly 2 values. Got %v: '%v'", len(values), values)
	}

	if values := store.GetActive(key, 150); len(values) != 1 {
		t.Errorf("Expected exactly 1 value. Got %v: '%v'", len(values), values)
	}

	if values := store.GetActive(key, 50); len(values) != 0 {
		t.Errorf("Expected exactly zero values. Got %v: '%v'", len(values), values)
	}

	if expired := store.RemoveExpiredForTime(key, 550); len(expired) != 2 {
		t.Errorf("Expected exactly 2 expired. Got %v: '%v'", len(expired), expired)
	}

	if values := store.GetActive(key, 350); len(values) != 1 {
		t.Errorf("Expected exactly 1 value. Got %v: '%v'", len(values), values)
	}

	if values := store.GetActive(key, 250); len(values) != 0 {
		t.Errorf("Expected exactly zero values. Got %v: '%v'", len(values), values)
	}
}

func TestEternalValueStorage(t *testing.T) {
	store := NewStore()

	const key = "key"
	value := NewValue(100, 200, "value")
	eternal := NewEternalValue("eternal")

	store.Put(key, value)
	store.Put(key, eternal)

	if values := store.GetActive(key, 150); len(values) != 2 {
		t.Errorf("Expected exactly 2 values. Got %v: '%v'", len(values), values)
	}

	if values := store.GetActive(key, 500); len(values) != 1 {
		t.Errorf("Expected exactly 1 value. Got %v: '%v'", len(values), values)

		if values[0].Data != "eternal" {
			t.Errorf("Expected to get back the eternal value. Got '%v' instead.", values[0].Data)
		}
	}
}

func TestRemovingExpired(t *testing.T) {
	store := NewStore()

	const key = "key"
	const value1 = "value1"
	const value2 = "value2"

	store.Put(key, NewValue(200, 300, value1))
	store.Put(key, NewValue(800, 900, value2))

	if values := store.GetActive(key, 250); len(values) != 1 {
		t.Errorf("Expected exactly 1 value. Got %v: '%v'", len(values), values)
	}

	if values := store.GetActive(key, 850); len(values) != 1 {
		t.Errorf("Expected exactly 1 value. Got %v: '%v'", len(values), values)
	}

	if values := store.GetActive(key, 1000); len(values) != 0 {
		t.Errorf("Expected exactly zero values. Got %v: '%v'", len(values), values)
	}

	if expiredValues := store.RemoveExpiredForTime(key, 500); len(expiredValues) != 1 {
		t.Errorf("Expected exactly 1 expired value. Got %v: '%v'", len(expiredValues), expiredValues)

		if expiredValues[0].Data != value1 {
			t.Errorf("Expected expired value to be '%v'. Got '%v'", value2, expiredValues[0].Data)
		}
	}

	if values := store.GetActive(key, 250); len(values) != 0 {
		t.Errorf("Expected exactly zero values. Got %v: '%v'", len(values), values)
	}

	if values := store.GetActive(key, 850); len(values) != 1 {
		t.Errorf("Expected exactly 1 value. Got %v: '%v'", len(values), values)
	}
}
