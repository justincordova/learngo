# Functions

## Function Declaration Syntax

The correct syntax for declaring a function in Go:

```go
func run(a, b int) {}
```

**Key elements:**
- Starts with `func` keyword
- Function name follows
- Parameters in parentheses with types
- When consecutive parameters share the same type, you can declare the type once: `(a, b int)` instead of `(a int, b int)`

## Function Signatures

A function signature consists of its input parameters and result values:

```go
func run(p Process, id1, id2 int) (pid int, err error) {}
```

**Breaking down the signature:**
- **Inputs:** `p`, `id1`, and `id2`
- **Input types:** `Process`, `int`, `int`
- **Results:** `pid`, `err`
- **Result types:** `int`, `error`

Note how `id1` and `id2` share the type `int`, similar to the parameter grouping syntax.

## Return Statements

A return statement terminates a function by returning zero or more values to the calling function. It does not terminate the entire program—only the current function.

### Return with Results

Functions that declare result values must return those values:

```go
func add(a, b int) int {
    return a + b
}
```

**Common mistake:**
```go
func add(a, b int) {
    return a + b  // Error: function has no result type
}
```

To fix: declare an `int` result value.

## Pass-by-Value Semantics

Go is a 100% pass-by-value language. When you pass arguments to a function, the function receives copies of those values. Changes to parameters inside the function do not affect the original values.

### Example Problem

```go
func incr(a int) {
    a++
    return
}

num := 10
incr(num)
fmt.Println(num)  // Still prints 10, not 11
```

**Solution:** Return the modified value:

```go
func incr(a int) int {
    a++
    return a
}

num := 10
num = incr(num)
fmt.Println(num)  // Now prints 11
```

## Package-Level Variables

While package-level variables are allowed, they can increase code coupling and lead to fragile code because anyone can access and modify them. This makes it harder to reason about your code's behavior.

**Prefer:** Passing values as function parameters and returning results.

## Error Handling

Always return errors when operations can fail, rather than silently ignoring them:

```go
// Good: Returns error to caller
func incr(n string) (int, error) {
    m, err := strconv.Atoi(n)
    if err != nil {
        return 0, err
    }
    return m + 1, nil
}

// Bad: Silently ignores errors
func incr(n string) int {
    m, _ := strconv.Atoi(n)  // Ignoring error!
    return m + 1
}
```

When `strconv.Atoi` encounters an error, it returns `0`, but the caller has no way to know if `0` is the intended result or an error condition. Always let the caller decide how to handle errors.

## Named Result Values and Naked Returns

You can name your result values in the function signature. Named result values are automatically initialized to their zero values and can be returned without explicitly specifying them:

```go
func spread(samples int, P int) (estimated float64) {
    for i := 0; i < P; i++ {
        estimated += estimate(i, P)
    }
    return  // Naked return - automatically returns 'estimated'
}
```

A naked `return` statement automatically returns all named result values.

**Note:** While naked returns are possible, they can reduce code clarity in longer functions. Use them judiciously.

## Reference Types and Mutability

### Maps Are Reference Types

Map values are pointers, so functions can modify the underlying map:

```go
func main() {
    stats := map[int]int{1: 10, 10: 2}
    incrAll(stats)
    fmt.Print(stats)  // Prints: map[1:11 10:3]
}

func incrAll(stats map[int]int) {
    for k := range stats {
        stats[k]++
    }
}
```

This works because the map value is a pointer to the underlying map structure. Even though the parameter is passed by value, the copied pointer still references the same map.

### Slices Require Care

Slice headers are passed by value. While you can modify existing elements through the slice, operations that change the slice header (like `append`) only affect the local copy:

```go
func main() {
    stats := []int{10, 5}
    add(stats, 2)
    fmt.Print(stats)  // Prints: [10 5] - unchanged!
}

func add(stats []int, n int) {
    stats = append(stats, n)  // Only modifies local copy of slice header
}
```

**Solution:** Return the updated slice:

```go
func add(stats []int, n int) []int {
    return append(stats, n)
}

func main() {
    stats := []int{10, 5}
    stats = add(stats, 2)
    fmt.Print(stats)  // Prints: [10 5 2]
}
```

The append operation may create a new underlying array, so you must return and reassign the slice to capture the updated header.
