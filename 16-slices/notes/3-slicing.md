# Slicing

## Slicing Syntax

A slice expression creates a new slice from an existing slice or array using the syntax `[start:stop]`:

```go
nums := []int{9, 7, 5, 2, 4, 6}
subset := nums[2:4]  // [5 2]
```

**Important:**
- `start` is the starting index (inclusive)
- `stop` is the stopping position (exclusive)
- The result includes elements from `nums[start]` up to but not including `nums[stop]`

## Slicing Examples

### Middle Slice

```go
nums := []int{9, 7, 5}
nums = append(nums, 2, 4, 6)  // [9 7 5 2 4 6]

fmt.Println(nums[2:4])  // [5 2]
```

- `nums[2]` is 5 (starting element)
- `nums[3]` is 2 (last element included)
- `nums[4]` is 4 (stopping position, not included)

### From Start

```go
nums := []int{9, 7, 5, 2, 4, 6}
fmt.Println(nums[:2])  // [9 7]
```

Omitting the start index defaults to 0, so `nums[:2]` is equivalent to `nums[0:2]`.

### To End

```go
nums := []int{9, 7, 5, 2, 4, 6}
fmt.Println(nums[len(nums)-2:])  // [4 6]
```

- `len(nums)` is 6
- `len(nums)-2` is 4
- `nums[4:]` gets all elements from index 4 to the end

Omitting the stop index defaults to the length of the slice.

### Full Slice

```go
names := []string{"einstein", "rosen", "newton"}
names = names[:]  // [einstein rosen newton]
fmt.Println(names[:1])  // [einstein]
```

`names[:]` is equivalent to `names[0:len(names)]`, creating a slice of the entire slice.

## Slice Expression vs Index Expression

### Slicing Returns a Slice

```go
names := []string{"einstein", "rosen", "newton"}
result := names[2:3]  // Type: []string
```

A slicing expression (`[start:stop]`) always returns a slice, even if it contains only one element.

### Indexing Returns an Element

```go
names := []string{"einstein", "rosen", "newton"}
result := names[2]  // Type: string
```

An index expression (`[index]`) returns a single element of the slice's element type.

## Relative Indexing

After slicing, indices are relative to the new slice:

```go
names := []string{"einstein", "rosen", "newton"}
names = names[1:len(names) - 1]  // ["rosen"]

fmt.Println(names[0])  // "rosen"
// names[1] would be out of bounds!
```

After `names[1:2]`, the slice only contains one element, so valid indices are only `[0]`.

## Chained Slicing

You can slice a slice multiple times:

```go
names := []string{"einstein", "rosen", "newton"}
names = names[1:]  // ["rosen", "newton"]
names = names[1:]  // ["newton"]
fmt.Println(names)  // [newton]
```

Each slicing operation returns a new slice, which can be sliced again.

## fmt.Sprintf

The `fmt.Sprintf` function works like `fmt.Printf` but returns a string instead of printing:

```go
i := 2
s := fmt.Sprintf("i = %d * %d = %d", i, i, i*i)
fmt.Print(s)  // i = 2 * 2 = 4
```

This is useful when you need to format a string for later use rather than immediately printing it.
