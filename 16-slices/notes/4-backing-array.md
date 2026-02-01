# Backing Arrays

## What is a Backing Array?

A slice doesn't directly store its elements. Instead, it **references a backing array** where the actual elements are stored. The slice is just a data structure that points to a segment of this array.

```go
s := []string{"I'm", "a", "slice"}
// The elements are stored in a backing array
// The slice 's' just references that array
```

**Key points:**
- A slice value itself is a small data structure (pointer, length, capacity)
- The actual elements live in the backing array
- Multiple slices can reference the same backing array

## Slicing Creates New Slices, Not New Arrays

When you slice a slice, the result is a **new slice value that references the same backing array**:

```go
s := []string{"I'm", "a", "slice"}
s2 := s[2:]  // New slice, same backing array
```

No new array is created—only a new slice header that points to a different position in the existing backing array.

## Why Slices Are Efficient

Because backing arrays are **contiguous in memory**, accessing elements is very fast:

```go
s := []int{1, 2, 3, 4, 5}
value := s[3]  // Fast O(1) lookup
```

Go can calculate the exact memory address of any element using simple arithmetic, making indexing and slicing extremely efficient.

## Backing Arrays from Array Slicing

When you create a slice from an array, **that array becomes the backing array**:

```go
arr := [...]int{1, 2, 3}
slice1 := arr[2:3]  // Backing array: arr
slice2 := slice1[:1]  // Backing array: still arr
```

Both `slice1` and `slice2` reference the same backing array: `arr`.

## Backing Arrays from Slice Literals

A slice literal creates a **hidden backing array**:

```go
arr := [...]int{1, 2, 3}
slice := []int{1, 2, 3}  // Creates a new hidden array
```

Even though the elements are identical, `slice` has its own backing array, distinct from `arr`.

## Each Slice Literal Creates a New Array

Every slice literal creates its own backing array:

```go
slice1 := []int{1, 2, 3}  // Creates backing array #1
slice2 := []int{1, 2, 3}  // Creates backing array #2
```

`slice1` and `slice2` have **different backing arrays**, even though their elements are the same.

## Slicing Shares Backing Arrays

When you slice an existing slice, the new slice shares the same backing array:

```go
slice1 := []int{1, 2, 3}     // Creates backing array #1
slice2 := []int{1, 2, 3}     // Creates backing array #2
slice3 := slice1[:]          // Shares backing array #1
slice4 := slice2[:]          // Shares backing array #2
```

**Relationships:**
- `slice1` and `slice3` share the same backing array (#1)
- `slice2` and `slice4` share the same backing array (#2)
- But #1 and #2 are different arrays

## Slicing Doesn't Modify the Backing Array

Slicing changes what elements are visible through the slice, but the backing array remains intact:

```go
nums := []int{9, 7, 5, 3, 1}
nums = nums[:1]

fmt.Println(nums)  // [9]
// But the backing array still contains: [9 7 5 3 1]
```

The backing array still holds all five elements. The slice `nums` now only "sees" the first element, but elements 7, 5, 3, and 1 still exist in memory.

## Shared Backing Arrays and Mutations

When multiple slices share a backing array, changes through one slice affect the others:

```go
arr   := [...]int{9, 7, 5, 3, 1}
nums  := arr[2:]   // [5 3 1]
nums2 := nums[1:]  // [3 1]

arr[2]++           // Modifies arr[2] to 6
nums[1]  -= arr[4] - 4  // nums[1] is arr[3], becomes 3 - (1-4) = 6
nums2[1] += 5      // nums2[1] is arr[4], becomes 1 + 5 = 6

fmt.Println(nums)  // [6 6 6]
```

**Explanation:**
- `nums` references `arr[2:5]` → [5 3 1]
- `nums2` references `arr[3:5]` → [3 1]
- All modifications affect the same backing array `arr`
- After modifications: `arr` = [9 7 6 6 6]
- So `nums` (arr[2:5]) = [6 6 6]

**This demonstrates why understanding backing arrays is crucial:** changes made through one slice can unexpectedly affect other slices that share the same backing array.
