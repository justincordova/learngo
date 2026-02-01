# Loop Statement Notes

## Go's Loop Statement

Go has only one loop statement: `for`. There are no `while`, `until`, or `forever` keywords.

## Basic For Loop

```go
for i := 3; i > 0; i-- {
    fmt.Println(i)
}
// Prints: 3 2 1
```

The loop has three parts:
1. Init statement: `i := 3`
2. Condition: `i > 0`
3. Post statement: `i--`

## Omitting the Post Statement

```go
for i := 3; i > 0; {
    i--
    fmt.Println(i)
}
// Prints: 2 1 0
```

When you move the decrement inside the loop body and place it before the print, the output changes because `i` is decremented before printing.

## Omitting Condition (Using break)

```go
for i := 3; ; {
    if i <= 0 {
        break
    }

    i--
    fmt.Println(i)
}
// Prints: 2 1 0
```

Without a condition in the for statement, use `break` to exit the loop explicitly.

## Continue Statement

```go
for i := 2; i <= 9; i++ {
    if i % 3 != 0 {
        continue
    }

    fmt.Println(i)
}
// Prints: 3 6 9
```

`continue` skips the rest of the current iteration and moves to the next one. This prints only numbers divisible by 3.

## Infinite Loop

```go
for ; true ; {
    // runs forever
}
```

Can be simplified to:

```go
for {
    // runs forever
}
```

An empty `for` statement creates an infinite loop.

## Range Loops

### Range with Index and Value

```bash
go run main.go go is awesome
```

```go
for i, v := range os.Args[1:] {
    fmt.Println(i+1, v)
}
// Prints:
// 1 go
// 2 is
// 3 awesome
```

`range` returns both the index and value for each element. Adding `i+1` adjusts the index to start from 1 instead of 0.

### Range with Index Only

```go
for i := range os.Args[1:] {
    fmt.Println(i+1)
}
// Prints:
// 1
// 2
// 3
```

Omitting the second variable gives you only the index.

### Range with Value Only

```go
for _, v := range os.Args[1:] {
    fmt.Println(v)
}
// Prints:
// go
// is
// awesome
```

Use `_` to discard the index and keep only the value.

### Range for Counting

```bash
go run main.go go is awesome
```

```go
var i int

for range os.Args {
    i++
}

fmt.Println(i)
// Prints: 4
```

You can use `for range` without variables to simply count elements. This counts all arguments including the program name, giving 4 total elements.
