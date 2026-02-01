# Array Basics

## What is an Array?

An array is a collection of values with a **fixed length** and **fixed type**. All elements in an array must be of the same type, and the array's size is determined at declaration time and cannot change.

## Memory Layout: Variables vs Arrays

### Individual Variables

Regular variables can be stored anywhere in memory with no guaranteed relationship to each other:

```go
var (
    first  int32 = 100
    second int32 = 150
)
```

If `first` is stored at memory location 20, `second` could be stored at any location—there's no guarantee they'll be adjacent.

### Array Elements

Unlike individual variables, **array elements are stored in contiguous memory locations**. This is one of the key characteristics that makes arrays efficient for certain operations.

```go
// Let's say nums array starts at memory location 500
var nums [5]int64
```

Memory layout for this array:
- 1st element (index 0): location 500
- 2nd element (index 1): location 508 (500 + 8 bytes)
- 3rd element (index 2): location 516 (500 + 16 bytes)
- 4th element (index 3): location 524 (500 + 24 bytes)
- 5th element (index 4): location 532 (500 + 32 bytes)

**Formula:** `Memory Location = Starting Position + Element Size × Index`

For the 3rd element: `516 = 500 + 8 × 2`

## Array Variables

An array variable stores **one value**—the entire array:

```go
var gophers [10]string
```

The variable `gophers` stores a single array value (which happens to contain 10 string elements). Through this variable, you can access individual string values inside the array.

## Array Length

The length of an array is the number of elements it can hold:

```go
var gophers [5]int  // Length is 5
```

You can use constant expressions when declaring array length:

```go
const length = 5 * 2
var gophers [length - 1]int  // Length is 9 (5 * 2 - 1)
```

## Array Type

An array's type consists of both its **length** and its **element type**:

```go
var luminosity [100]float32
```

- **Element type:** `float32` (the type of each individual element)
- **Array type:** `[100]float32` (the complete type including length)

Arrays with different lengths are **different types**, even if they have the same element type:

```go
var a [5]int
var b [10]int
// a and b have different types: [5]int vs [10]int
```

## Working with Array Elements

You can access and modify array elements using index notation:

```go
package main
import "fmt"

func main() {
    var names [3]string

    names[len(names)-1] = "!"        // names[2] = "!"
    names[1] = "think" + names[2]    // names[1] = "think!"
    names[0] = "Don't"
    names[0] += " "                  // names[0] = "Don't "

    fmt.Println(names[0] + names[1] + names[2])
    // Prints: "Don't think!!"
}
```

Array indexing is zero-based, so valid indices for a `[3]string` array are 0, 1, and 2.

## Range Loop Behavior

The `for range` loop creates a **copy** of the array being iterated over. Modifications to the original array during iteration won't affect the values seen by the loop:

```go
package main
import "fmt"

func main() {
    var sum [5]int

    for i, v := range sum {
        if i == len(sum) - 1 {
            break
        }

        sum[i+1] = 10      // Modifies original array
        fmt.Print(v, " ")  // But prints from the copy
    }
    // Prints: 0 0 0 0
}
```

**Explanation:**
- The `range` clause makes a copy of the `sum` array at the start of the loop
- Loop variable `v` receives values from this copy (all zeros)
- Modifications to `sum[i+1]` affect the original array, not the copy
- Therefore, `v` always prints `0` even though we're setting `sum[i+1]` to `10`

If you print the `sum` array after the loop completes, you'll see the modifications: `[0 10 10 10 10]`.
