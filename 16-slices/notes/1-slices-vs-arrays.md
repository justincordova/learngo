# Slices vs Arrays

## Why Use Slices?

Slices provide **dynamic length**, unlike arrays which have fixed length. This makes slices ideal for collections whose size may change at runtime:

```go
// Array: fixed length of 3, determined at compile-time
var arr [3]int

// Slice: dynamic length, can grow or shrink at runtime
var slice []int
```

With slices, you can build collections that grow and shrink as needed without knowing the exact size in advance.

## Slice Length at Runtime

A slice's length is not part of its type, so it can change during program execution:

```go
nums := []int{1, 2, 3}    // Length: 3
nums = append(nums, 4, 5) // Length: 5
```

The length is determined at **runtime**, not compile-time. This flexibility allows you to work with data of varying sizes.

## Passing Slices to Functions

Functions that accept slices require exact type matching:

```go
func sort(nums []int) {
    // ...
}

// Correct
sort([]int{3, 1, 6})

// Incorrect - array, not slice
// sort([...]int{3, 1, 6})

// Incorrect - wrong element type
// sort([]int32{3, 1, 6})
```

Arrays and slices are different types, even if they contain the same element type. You cannot pass an array where a slice is expected.

## Nil Slices

The zero value of a slice is `nil`, unlike arrays whose zero value is an array with zero-valued elements:

```go
var tasks []string
fmt.Println(tasks)        // []
fmt.Println(tasks == nil) // true
```

## Operations on Nil Slices

You can perform certain operations on nil slices:

### Length of Nil Slice

```go
var tasks []string
fmt.Println(len(tasks))   // 0
```

The `len()` function works on nil slices and returns 0.

### Indexing Nil Slice

```go
var tasks []string
fmt.Println(tasks[0])     // Runtime panic!
```

You **cannot** index into a nil slice because it doesn't contain any elements. This will cause a runtime panic.

## Slice Declaration Syntax

Valid slice declarations use `[]Type` without specifying length:

```go
// Correct slice declarations
[]string{"hello", "world"}
[]int{1, 2, 3}
[]uint64{}

// Arrays (not slices)
[...]int{}                      // Array with inferred length
[2]string{"hello", "world"}     // Array with explicit length
```

The absence of a size (or `...`) in the brackets indicates a slice type.

## Slice Comparability

Slices **cannot be compared** to each other using `==` or `!=`:

```go
colors := []string{"red", "blue", "green"}
tones := []string{"dark", "light"}

// if colors == tones {  // Compile error!
//     ...
// }
```

The only valid comparison for a slice is checking if it's nil:

```go
var tasks []string
if tasks == nil {
    fmt.Println("tasks is nil")
}
```

To compare slice contents, you must iterate through elements or use a function like `slices.Equal()` from the `slices` package.

## Empty vs Nil Slices

An empty slice and a nil slice are different:

```go
var nilSlice []uint64        // nil slice
emptySlice := []uint64{}     // empty slice (not nil)

fmt.Println(len(nilSlice))   // 0
fmt.Println(len(emptySlice)) // 0
fmt.Println(nilSlice == nil) // true
fmt.Println(emptySlice == nil) // false
```

Both have length 0, but only the nil slice equals `nil`. An empty slice is allocated with zero elements.

## Slice Length Examples

The length of a slice is the number of elements it contains:

```go
[]uint64{}                                           // Length: 0
[]string{"I'm", "going", "to", "stay", "\"here\""}  // Length: 5
```

Escaped quotes within string literals count as part of a single string element, not as separate elements.
