package timedstore

import (
	"math"
	"testing"
)

func TestNewValue(t *testing.T) {
	const expectedDuration = 10

	start := CurrentTime()
	end := start + expectedDuration

	value := NewValue(start, end, "my value")

	if value.Duration() != expectedDuration {
		t.Errorf("Value duration is %d, expected %d", value.Duration(), expectedDuration)
	}
}

func TestNewEternalValue(t *testing.T) {
	eternal := NewEternalValue("always-active")
	if eternal.Duration() != math.MaxUint64 {
		t.Errorf("Eternal value duration is wrong. Expected '%d'. Got '%d'", uint64(math.MaxUint64), eternal.Duration())
	}
}

func TestNewValueFromDuration(t *testing.T) {
	const duration = 10

	now := CurrentTime()
	expectedEndSeconds := now + duration

	value := NewValueFromDuration(now, duration, "my value")
	if value.EndSeconds != expectedEndSeconds {
		t.Errorf("Value end time is %d, expected %d", value.EndSeconds, expectedEndSeconds)
	}
}

func TestIsActiveForTime(t *testing.T) {
	const startTime = 1000
	const endTime = 9000
	const duration = 10000

	const timeToCheckActive = 5000
	const timeInactiveTooLow = 10
	const timeInactiveTooHigh = 10000000

	value := NewValue(startTime, endTime, "my value")
	if value.IsActiveForTime(timeToCheckActive) != true {
		t.Errorf("1: value.IsActiveForTime did not return expected truth value")
	}

	if value.IsActiveForTime(timeInactiveTooLow) != false {
		t.Errorf("1: value.IsActiveForTime did not return expected truth value")
	}

	if value.IsActiveForTime(timeInactiveTooHigh) != false {
		t.Errorf("1: value.IsActiveForTime did not return expected truth value")
	}

	value2 := NewValueFromDuration(startTime, duration, "my value")
	if value2.IsActiveForTime(timeToCheckActive) != true {
		t.Errorf("value2.IsActiveForTime did not return expected truth value")
	}

	if value2.IsActiveForTime(timeInactiveTooLow) != false {
		t.Errorf("value2.IsActiveForTime did not return expected truth value")
	}

	if value2.IsActiveForTime(timeInactiveTooHigh) != false {
		t.Errorf("value2.IsActiveForTime did not return expected truth value")
	}

	eternal := NewEternalValue("eternal")

	if !eternal.IsActiveForTime(100) {
		t.Errorf("Eternal values should always be active")
	}
}

func TestExpiredForTime(t *testing.T) {
	const startTime = 1000
	const endTime = 10000

	const timeToCheckExpired = endTime + 100
	const timeToCheckNotExpired = startTime - 100

	value := NewValue(startTime, endTime, "my value")
	if !value.IsExpiredForTime(timeToCheckExpired) {
		t.Errorf("Expected value to be expired, but was not")
	}

	if value.IsExpiredForTime(startTime) {
		t.Errorf("Did not expect value to be expired")
	}

	if value.IsExpiredForTime(timeToCheckNotExpired) {
		t.Errorf("Did not expect value to be expired")
	}

	eternal := NewEternalValue("eternal")

	if eternal.IsExpiredForTime(100) {
		t.Errorf("Eternal values should never be expired")
	}
}
