# The Mechanics of Append

## When Append Allocates a New Backing Array

Given a slice:
```go
words := []string{"lucy", "in", "the", "sky", "with", "diamonds"}
```

### Operations that reuse the backing array:

```go
words = append(words[:3], "crystals")
// Overwrites the 4th element ("sky")
// No reallocation - enough capacity

words = append(words[:4], "crystals")
// Overwrites the 5th element ("with")
// No reallocation - enough capacity

words = append(words[:5], "crystals")
// Overwrites the last element ("diamonds")
// No reallocation - enough capacity
```

### Operation that allocates new backing array:

```go
words = append(words[:5], "crystals", "and", "diamonds")
// Tries to append 3 elements after position 5
// Not enough capacity - allocates new backing array
```

When append needs more capacity than available, it allocates a new backing array.

## Complex Append Operations

```go
words := []string{"lucy", "in", "the", "sky", "with", "diamonds"}
words = append(words[:1], "is", "everywhere")
words = append(words, words[len(words)+1:cap(words)]...)
```

**Result:** `["lucy" "is" "everywhere" "with" "diamonds"]`

**Explanation:**
- Line 2: Keeps `["lucy"]`, then appends `"is"` and `"everywhere"`, overwriting positions 2-3
- Line 3: Appends remaining elements from the backing array (`["with" "diamonds"]`)

## Capacity Growth Patterns

### Growth by Doubling (< 1024 elements)

Starting with 1023 elements:
```go
words := []string{1022: ""}  // Creates slice with 1023 elements
words = append(words, "boom!")
```

**Result:** Length: 1024, Capacity: 2048

Append doubles the capacity when growing slices with < 1024 elements.

### Growth by ~25% (≥ 1024 elements)

Starting with 1024 elements:
```go
words := []string{1023: ""}  // Creates slice with 1024 elements
words = append(words, "boom!")
```

**Result:** Length: 1025, Capacity: 1280

After 1024 elements, append grows capacity by approximately 25% instead of doubling to conserve memory.

## Key Points

1. Append reuses the backing array if there's enough capacity
2. Append allocates a new backing array when capacity is exceeded
3. Growth pattern:
   - **< 1024 elements:** Double the capacity
   - **≥ 1024 elements:** Grow by ~25%
4. Use three-index slicing to limit capacity and force reallocation when needed
