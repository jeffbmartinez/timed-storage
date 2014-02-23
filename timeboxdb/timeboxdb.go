package main

import (
  "fmt"
  "github.com/jeffbmartinez/timed-storage/timedstore"
)

func main() {
  value := timedstore.NewValue("hi", 1, 2)
  fmt.Printf("%v\n", value)
}
