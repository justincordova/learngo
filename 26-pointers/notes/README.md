# Pointers

## What is a Pointer?

A pointer is a value that contains a memory address of another value. While pointers can be stored in variables, they are not solely variables—they are values just like any other type in Go.

**Key distinction:** A pointer is a value type that holds a memory address, not just a variable concept.

## Declaring Pointer Types

When declaring a pointer to a type, use `*` in front of the type name:

```go
type computer struct {
    brand string
}

var c *computer  // c is a pointer to a computer
```

The `*` before the type denotes a pointer type.

## Pointer Operators

### The `&` Operator (Address-of)
Gets the memory address of a value:

```go
c := &computer{"Apple"}  // c holds the address of a computer struct
```

### The `*` Operator (Dereference)
Gets the value that is pointed to by the pointer:

```go
type computer struct {
    brand string
}

c := &computer{"Apple"}
value := *c  // Gets the computer struct that c points to
```

**Important:** The `*` has two different meanings depending on context:
- Before a type: denotes a pointer type (`*computer`)
- Before a pointer value: dereferences the pointer to get the pointed value (`*c`)

## Pointer Comparison

Pointers can be compared for equality, but understand what is being compared:

```go
type computer struct {
    brand string
}

var a, b *computer
fmt.Print(a == b)  // true - both are nil

a = &computer{"Apple"}
b = &computer{"Apple"}
fmt.Print(" ", a == b, " ", *a == *b)  // false true
```

**Explanation:**
- Initially, both `a` and `b` are `nil`, so they are equal
- After assignment, each composite literal creates a new value at a different memory address
- The addresses are different (`a == b` is false)
- But the values they point to are identical (`*a == *b` is true)

## Pointers and Function Calls

When passing pointers to functions, remember that Go passes arguments by value. This means the pointer itself is copied, but it still points to the same memory address:

```go
type computer struct {
    brand string
}

func main() {
    a := &computer{"Apple"}
    b := a
    change(b)
    change(b)
}

func change(c *computer) {
    c.brand = "Indie"  // Modifies the pointed value
    c = nil            // Only modifies the local copy of the pointer
}
```

**Variable count:** This code creates 4 pointer variables:
- `a` in main
- `b` in main (copy of `a`)
- `c` in first call to `change`
- `c` in second call to `change`

Each function call creates new variables from its parameters.

## Addressability

Not all values in Go are addressable. **Map elements are unaddressable**, which means you cannot take their address:

```go
m := map[string]int{"key": 42}
// ptr := &m["key"]  // Compile error: cannot take address of map element
```

This restriction exists because map elements can move in memory during map operations, which would invalidate any pointers to them.
