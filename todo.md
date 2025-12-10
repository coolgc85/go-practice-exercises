# Go Study Plan

This document tracks the learning topics covered in this Go exercise repository.

## Note on Go Environment
This project uses **Go Modules** for dependency management, which is the standard in modern Go. The older `GOPATH` workspace model is a legacy concept, but understanding it can be helpful when working with older projects.

## Completed Topics
- [x] Introduction to Go
- [x] Environment Setup (Go Modules)
- [x] Basic Syntax (variables, types, constants) - (covered in most exercises)
- [x] Control Flow (`if`/`else`, `for`, `switch`) - (`controlflow`)
- [x] Functions (including multiple returns) - (`multiple_returns`)
- [x] Packages & Imports - (`package`)
- [x] Pointers - (`pointers`)
- [x] Arrays - (`array_slices`)
- [x] Slices (dynamic arrays) - (`array_slices`)
- [x] Maps - (`salary_maps`)
- [x] Structs (custom types) - (`book_structs`)
- [x] Error Handling (error returns) - (`multiple_returns`)
- [x] Basic I/O (Console I/O) - (covered in most exercises)
- [x] Resource Management(`defer` keyword) - (`defer`)

## To-Do
Here are the topics from the initial study list that are not yet covered in this project.

- [ ] **Basic I/O: File I/O**
  - Create an exercise that reads from or writes to a file using the `os` or `io/ioutil` package.

- [ ] **Concurrency**
  - **Goroutines**: Create an exercise that uses the `go` keyword to start a simple, concurrent task.
  - **sync.WaitGroup**: Extend the goroutine exercise to use a `sync.WaitGroup` to wait for the concurrent task to complete before the program exits.