# What's a variable?

## Where variables live

Variables are stored in computer memory (RAM), not on the hard disk or in the CPU.

## Why use variable names

You use a variable's name to access and modify its value later in your program.

## Changing variable values

Change a variable's value by using its name:

```go
var count int = 5
count = 10  // change the value using the name
```

## Static typing

After declaration, you cannot change a variable's type. Go is **statically typed** — types are fixed at compile time.

```go
var age int = 25
age = "twenty-five"  // ERROR: cannot change type from int to string
```
