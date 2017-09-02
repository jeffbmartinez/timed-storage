# Timedstore

A key value storage system which can store multiple values per key, along with an active time for each value.

The Get call, when supplied a key, returns all values for which the active time is now.

## Tests

The standard: `go test ./...`
To ignore any vendored files: `go test $(go list ./... | grep -v /vendor/)`
