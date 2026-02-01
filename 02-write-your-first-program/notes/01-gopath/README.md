## Where to save your Go source code

You can save your Go source code anywhere on your computer. Modern Go uses modules, so there's no restriction on project location.

## What is a Go module?

A Go module is defined by a `go.mod` file at the project root. This file specifies the module path and manages dependencies.

## Initializing a Go module

Use `go mod init <module-path>` to create a new module:

```bash
go mod init example.com/myproject
```

## Why go.mod matters

The `go.mod` file:
- Defines your module's import path
- Tracks all dependencies and their versions
- Enables reproducible builds