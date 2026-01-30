# Concurrency in Go

Learn how to write concurrent programs using goroutines and channels.

## Topics Covered

1. **Goroutines Basics** - Launching concurrent tasks
2. **Channels** - Communication between goroutines
3. **Select Statement** - Multiplexing channel operations
4. **Mutexes** - Protecting shared state
5. **WaitGroups** - Coordinating goroutine completion
6. **Worker Pool Pattern** - Practical concurrent design

## Prerequisites

- Section 25: Functions
- Section 27: Error Handling

## Important Notes

- Goroutines are lightweight (not OS threads)
- Channels provide safe communication
- "Don't communicate by sharing memory; share memory by communicating"
- Always use `-race` flag during development: `go run -race program.go`

## Race Detection

Run any concurrent program with race detector:
```bash
go run -race main.go
```

This helps catch data races during development.
