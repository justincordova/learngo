# Advanced Slice Operations

## Three-Index Slicing

### Syntax: [low:high:max]

```go
lyric := []string{"show", "me", "my", "silver", "lining"}
part  := lyric[1:3:5]
```

**Formula:**
- Length: `high - low`
- Capacity: `max - low`

**Result:**
- `lyric[1:3]` returns `["me" "my"]` (length: 2)
- Capacity limited to 4 (elements from index 1 to 5: `["me" "my" "silver" "lining"]`)
- **Length: 2, Capacity: 4**

### Limiting Capacity to Force Reallocation

```go
lyric := []string{"show", "me", "my", "silver", "lining"}
part  := lyric[:2:2]
part   = append(part, "right", "place")
```

**Result:**
- `lyric` stays unchanged: `["show" "me" "my" "silver" "lining"]` (len: 5, cap: 5)
- `part` after first line: `["show" "me"]` (len: 2, cap: 2)
- `part` after append: `["show" "me" "right" "place"]` (len: 4, cap: 4)
- Append allocates new backing array because capacity was limited to 2

This technique prevents `part` from modifying `lyric`'s backing array.

## The make Function

### When to Use make

Use `make()` to preallocate a backing array when you know the size in advance. This improves performance by avoiding multiple reallocations.

### make with Length Only

```go
tasks := make([]string, 2)
tasks  = append(tasks, "hello", "world")

fmt.Printf("%q\n", tasks)
// Prints: ["" "" "hello" "world"]
```

`make([]string, 2)` creates a slice with:
- Length: 2
- Capacity: 2
- Elements initialized to zero values (`""`)

`append()` adds after the length, so new elements go after the two empty strings.

### make with Length and Capacity

```go
tasks := make([]string, 0, 2)
tasks  = append(tasks, "hello", "world")

fmt.Printf("%q\n", tasks)
// Prints: ["hello" "world"]
```

`make([]string, 0, 2)` creates a slice with:
- Length: 0
- Capacity: 2
- No initialized elements

`append()` starts at position 0, so elements appear at the beginning. This is the common pattern when using `make` with `append`.

## The copy Function

```go
lyric := []string{"le", "vent", "nous", "portera"}
n := copy(lyric, make([]string, 4))

fmt.Printf("%d %q\n", n, lyric)
// Prints: 4 ["" "" "" ""]
```

`copy(dst, src)` copies elements from `src` to `dst`:
- Returns number of elements copied (4)
- `make([]string, 4)` creates 4 empty strings
- Copies those empty strings to `lyric`, clearing it
- Copies `min(len(dst), len(src))` elements

## Multi-Dimensional Slices

### Indexing Multi-Dimensional Slices

```go
spendings := [][]int{{200, 100}, {}, {50, 25, 75}, {500}}
total := spendings[2][1] + spendings[3][0] + spendings[0][0]

fmt.Printf("%d\n", total)
// Prints: 725
```

Breakdown:
- `spendings[2][1]` = 25 (3rd row, 2nd element)
- `spendings[3][0]` = 500 (4th row, 1st element)
- `spendings[0][0]` = 200 (1st row, 1st element)
- Total: 25 + 500 + 200 = 725

### Types of Multi-Dimensional Slices

```go
spendings := [][]int{{1,2}}

fmt.Printf("%T - ", spendings)      // [][]int
fmt.Printf("%T - ", spendings[0])   // []int
fmt.Printf("%T", spendings[0][0])   // int
```

**Type hierarchy:**
- `spendings` type: `[][]int` (2D int slice)
- `spendings[0]` type: `[]int` (1D int slice)
- `spendings[0][0]` type: `int` (single int)

### Element Types

```go
[][][3]int{{{10, 5, 9}}}
```

**Element type:** `[][3]int`

**Type breakdown:**
- `[][][3]int` is a slice of `[][3]int` elements
- `[][3]int` is a slice of `[3]int` elements
- `[3]int` is an array of 3 `int` values

## Key Points

1. Three-index slicing `[low:high:max]` limits capacity
2. Use `make([]T, 0, cap)` when planning to `append()`
3. Use `make([]T, len)` when you'll set elements by index
4. `copy()` can clear a slice by copying from zero-valued slice
5. Multi-dimensional slice indexing: `slice[row][column]`
6. Element type = the type after removing one `[]` level
