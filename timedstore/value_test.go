package timedstore

import (
  "testing"
)

func TestNewValue(t *testing.T) {
  const expectedDuration = 10

  start := CurrentTime()
  end := start + expectedDuration

  value := NewValue("my value", start, end)

  if value.Duration() != expectedDuration {
    t.Errorf("Value duration is %d, expected %d", value.Duration(), expectedDuration)
  }
}

func TestNewValueFromDuration(t *testing.T) {
  const duration = 10

  now := CurrentTime()
  expectedEndSeconds := now + duration

  value := NewValueFromDuration("my value", now, duration)
  if value.EndSeconds != expectedEndSeconds {
    t.Errorf("Value end time is %d, expected %d", value.EndSeconds, expectedEndSeconds)
  }
}

func TestIsActiveForTime(t *testing.T) {
  const startTime = 1000
  const endTime = 9000
  const duration = 10000

  const timeToCheck = 5000
  const timeInactiveTooLow = 10
  const timeInactiveTooHigh = 10000000

  value := NewValue("my value", startTime, endTime)
  if value.IsActiveForTime(timeToCheck) != true {
    t.Errorf("1: value.IsActiveForTime did not return expected truth value")
  }

  if value.IsActiveForTime(timeInactiveTooLow) != false {
    t.Errorf("1: value.IsActiveForTime did not return expected truth value")
  }

  if value.IsActiveForTime(timeInactiveTooHigh) != false {
    t.Errorf("1: value.IsActiveForTime did not return expected truth value")
  }

  value2 := NewValueFromDuration("my value", startTime, duration)
  if value2.IsActiveForTime(timeToCheck) != true {
    t.Errorf("value2.IsActiveForTime did not return expected truth value")
  }

  if value2.IsActiveForTime(timeInactiveTooLow) != false {
    t.Errorf("value2.IsActiveForTime did not return expected truth value")
  }

  if value2.IsActiveForTime(timeInactiveTooHigh) != false {
    t.Errorf("value2.IsActiveForTime did not return expected truth value")
  }
}

func TestExpiredForTime(t *testing.T) {
  const startTime = 1000
  const endTime = 10000

  const timeToCheck = 500000

  value := NewValue("my value", startTime, endTime)
  if !value.IsExpiredForTime(timeToCheck) {
    t.Errorf("Expected value to be expired, but was not")
  }

  if value.IsExpiredForTime(startTime) {
    t.Errorf("Did not expect value to be expired")
  }
}
