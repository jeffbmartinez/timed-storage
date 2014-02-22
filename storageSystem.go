package main

type Value struct {
  string Data

  StartDate uint64
  EndDate uint64
}

func NewValue(data string, startDate uint64, endDate uint64) *Value {
  return &Value{data, startDate, endDate}
}

func NewValueFromDuration(data string, startDate uint64, duration uint64) *Value {
  endDate := startDate + duration
  return NewValue(data, startDate, endDate)
}

func (v Value) Duration() uint64 {
  return v.EndDate - v.StartDate
}
