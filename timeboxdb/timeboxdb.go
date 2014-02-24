package main

import (
  "fmt"
  "log"
  "net/http"

  "github.com/jeffbmartinez/timed-storage/timedstore"
)

func getHandler(response http.ResponseWriter, request *http.Request) {
  switch request.Method {
    case "GET":
      response.Write([]byte("Implementation in Progress"))
    default:
      response.WriteHeader(http.StatusMethodNotAllowed)
      return
  }
}

func putHandler(response http.ResponseWriter, request *http.Request) {
  switch request.Method {
    case "POST":
      response.Write([]byte("Implementation in Progress"))
    default:
      response.WriteHeader(http.StatusMethodNotAllowed)
      return
  }
}

func main() {
  const (
    listenAddress = "localhost:9191"
  )

  http.HandleFunc("/get", getHandler)
  http.HandleFunc("/put", putHandler)

  fmt.Printf("Listening on %s\n", listenAddress)

  log.Fatal(http.ListenAndServe(listenAddress, nil))


  value := timedstore.NewValue("hi", 1, 2)
  fmt.Printf("%v\n", value)
}
