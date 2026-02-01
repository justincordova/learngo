# Appending to Slices

## How append() Works

The `append()` function adds new elements to a slice and **returns a new slice**. It does not modify the original slice unless you reassign the result:

```go
nums := []int{1, 2, 3}
result := append(nums, 4, 5)

fmt.Println(nums)    // [1 2 3] - original unchanged
fmt.Println(result)  // [1 2 3 4 5] - new slice
```

**Key point:** Always capture the return value of `append()`, as the new slice may have a different underlying array.

## Where Elements Are Appended

The `append()` function adds new elements **after the length** of the given slice:

```go
nums := []int{9, 7, 5}  // Length: 3
nums = append(nums, 2, 4, 6)
// New elements go at indices 3, 4, 5
```

Elements are always added to the end, extending the slice's length.

## Common Mistake: Not Capturing the Result

```go
nums := []int{9, 7, 5}
append(nums, []int{2, 4, 6}...)  // Result discarded!

fmt.Println(nums[3])  // Runtime panic: index out of range
```

**Problem:** The `append()` call returns a new slice with 6 elements, but that slice is discarded. The original `nums` still has only 3 elements, so `nums[3]` is invalid.

**Fix:** Capture the result:

```go
nums := []int{9, 7, 5}
nums = append(nums, []int{2, 4, 6}...)
fmt.Println(nums[3])  // 2
```

## Appending Without Reassignment

When you don't reassign the result, the original slice remains unchanged:

```go
nums := []int{9, 7, 5}
evens := append(nums, []int{2, 4, 6}...)

fmt.Println(nums, evens)
// Prints: [9 7 5] [9 7 5 2 4 6]
```

- `nums` still contains [9 7 5]
- `evens` contains the new slice [9 7 5 2 4 6]

The original slice is unaffected because `append()` creates a new slice.

## Appending With Reassignment

To update the original slice, reassign the result:

```go
nums := []int{9, 7, 5}
nums = append(nums, 2, 4, 6)

fmt.Println(nums)
// Prints: [9 7 5 2 4 6]
```

This overwrites the `nums` variable with the new slice returned by `append()`, effectively "updating" the slice.

## Appending Multiple Elements

You can append multiple elements in one call:

```go
// Append individual elements
nums = append(nums, 1, 2, 3)

// Append elements from another slice (use ... to unpack)
more := []int{4, 5, 6}
nums = append(nums, more...)
```

The `...` operator unpacks the slice into individual elements for appending.
