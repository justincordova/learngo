# Slice Header

## What is a Slice Header?

A slice header is a **tiny data structure that describes all or part of a backing array**. It doesn't contain the actual elements—those live in the backing array. Instead, it contains metadata about which portion of the array the slice represents.

```go
s := []int{1, 2, 3, 4, 5}
// The slice header describes this slice
// The elements live in a backing array
```

## Slice Header Fields

A slice header contains three fields:

1. **Pointer** - Memory address of the first element in the slice
2. **Length** - Number of elements currently in the slice
3. **Capacity** - Total number of elements available in the backing array from the start of the slice

```go
type SliceHeader struct {
    Pointer  uintptr
    Length   int
    Capacity int
}
```

## Understanding the Fields with an Example

Given a slice header:
- Pointer: 100th (memory address)
- Length: 5
- Capacity: 10

And a backing array:
```go
var array [10]string
```

This describes the slice `array[:5]`:
- Points to `array[0]` (the 100th memory address)
- Contains 5 elements (length)
- Can access up to 10 elements total from this position (capacity)

**Why capacity is 10:** From index 0, there are 10 elements available in the backing array (indices 0-9).

## Nil Slice Header

A nil slice has all fields set to zero:

```go
var tasks []string
// Header: Pointer: 0, Length: 0, Capacity: 0
```

Since it has no backing array, the pointer is 0 (nil), and both length and capacity are 0.

## Memory Usage

Understanding slice headers helps with memory analysis:

```go
var array [1000]int64  // 1000 × 8 bytes = 8000 bytes

array2 := array         // Copies the array: +8000 bytes
slice := array2[:]      // Creates slice header: +24 bytes
```

**Total:** 16,024 bytes
- `array`: 8,000 bytes
- `array2`: 8,000 bytes (arrays are copied on assignment)
- `slice` header: 24 bytes (a slice header is typically 24 bytes on 64-bit systems)

The slice doesn't allocate a new backing array—it references `array2`.

## What Gets Passed to Functions

When you pass a slice to a function, you're passing the **slice header** (pointer, length, capacity):

```go
nums := []int{9, 7, 5, 3, 1}
sort.Ints(nums)  // Passes the slice header
```

The function receives a copy of the slice header (24 bytes), not the entire backing array. This makes passing slices to functions very efficient—regardless of how many elements the slice contains, you're always passing just 24 bytes.

**Important:** Because the slice header contains a pointer to the backing array, the function can modify the elements in the backing array even though it receives a copy of the header.

## Slice Header Size

The slice header is a fixed size (typically 24 bytes on 64-bit systems):
- Pointer: 8 bytes
- Length: 8 bytes
- Capacity: 8 bytes

This fixed size makes slices extremely efficient to pass around, regardless of how large the backing array is.
