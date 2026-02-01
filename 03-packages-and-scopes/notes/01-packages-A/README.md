## Package file organization

All source code files belonging to a package should be in a single directory.

## Package clause

The `package` clause tells Go which package a file belongs to. It must be the first code in a Go source file and can only appear once per file.

**Correct syntax:**
```go
package main
```

## Files in the same package

All files in the same package can call each other's functions directly without importing.

## Running multiple Go files

To run multiple Go files together:

```bash
go run *.go
```

Or explicitly list them:

```bash
go run file1.go file2.go file3.go
```
