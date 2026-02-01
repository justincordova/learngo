# Slice Capacity

## Length vs Capacity

**Length** and **capacity** are two distinct properties of a slice:

- **Length**: The number of elements currently in the slice
- **Capacity**: The total number of elements available in the backing array, starting from the first element of the slice

```go
s := []int{1, 2, 3, 4, 5}
s = s[:3]
// Length: 3 (elements 1, 2, 3)
// Capacity: 5 (elements 1, 2, 3, 4, 5 available in backing array)
```

**Key relationship:** Length ≤ Capacity (length can never exceed capacity)

## Capacity of a Nil Slice

A nil slice has capacity 0:

```go
var tasks []string
fmt.Println(cap(tasks))  // 0
```

Since a nil slice has no backing array, its capacity is 0. The capacity type is `int`, not a pointer, so it cannot be `nil`.

## Slice Literals and Capacity

A slice literal creates a slice with **equal length and capacity**:

```go
s := []string{"I", "have", "a", "great", "capacity"}
// Length: 5
// Capacity: 5
```

The backing array created by the literal has exactly the number of elements specified, so length equals capacity.

## Slicing and Capacity

### Slicing to Zero Length

```go
words := []string{"lucy", "in", "the", "sky", "with", "diamonds"}
words = words[:0]
// Length: 0
// Capacity: 6
```

`words[:0]` creates a slice with no visible elements (length 0), but the backing array still has 6 elements available (capacity 6).

### Slicing the Full Range

```go
words := []string{"lucy", "in", "the", "sky", "with", "diamonds"}
words = words[0:]
// Length: 6
// Capacity: 6
```

`words[0:]` includes all elements, so both length and capacity remain 6.

### Slicing from Middle

```go
words := []string{"lucy", "in", "the", "sky", "with", "diamonds"}
words = words[2:cap(words)-2]
// Equivalent to: words[2:4]
// Length: 2
// Capacity: 4
```

**Breakdown:**
- `words[2:4]` gives ["the", "sky"] → length is 2
- Starting from index 2, there are 4 elements in the backing array (indices 2, 3, 4, 5)
- Capacity is 4

**Important:** Capacity is always measured from the **start of the slice** (the first element the slice points to) to the end of the backing array, not from the beginning of the backing array.

## Visualizing Capacity

```go
backing := [6]string{"lucy", "in", "the", "sky", "with", "diamonds"}
//                      0      1     2      3       4        5

s := backing[2:4]
// s points to indices 2-3: ["the", "sky"]
// Length: 2 (from index 2 to index 3)
// Capacity: 4 (from index 2 to index 5 - end of backing array)
```

The capacity counts how many elements are available from the slice's starting position to the end of the backing array.

## Using cap() Function

The `cap()` built-in function returns a slice's capacity:

```go
words := []string{"lucy", "in", "the", "sky", "with", "diamonds"}
fmt.Println(len(words))  // 6
fmt.Println(cap(words))  // 6

words = words[2:4]
fmt.Println(len(words))  // 2
fmt.Println(cap(words))  // 4
```
