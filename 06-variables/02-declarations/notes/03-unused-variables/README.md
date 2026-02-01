## Unused variables in block scope

If you declare a variable inside a function but don't use it, the program won't compile.

```go
func main() {
    var unused int  // ERROR: unused variable
}
```

## Unused variables in package scope

Variables declared at package level (outside functions) can remain unused without causing compilation errors.

```go
package main

var globalUnused int  // OK: package-scoped variables can be unused

func main() {
}
```

## Preventing unused variable errors

Use the blank identifier `_` to explicitly discard a value:

```go
func main() {
    result, _ := someFunction()  // discard second return value
}
```
