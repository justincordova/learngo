# Arrays - Advanced Topics

## Array Literals with Ellipsis

You can use `...` in an array literal to let Go automatically determine the length based on the number of elements:

```go
gadgets := [...]string{"Mighty Mouse", "Amazing Keyboard", "Shiny Monitor"}
// Type: [3]string, Length: 3
```

Go counts the elements in the literal and sets the array length accordingly.

### Empty Array Literals

```go
gadgets := [...]string{}
// Type: [0]string, Length: 0
```

With no elements in the list, Go creates an array of length 0.

## Zero Values in Arrays

When you initialize an array with fewer elements than its declared length, Go fills the remaining positions with zero values:

```go
package main
import "fmt"

func main() {
    gadgets := [3]string{"Confused Drone"}
    fmt.Printf("%q\n", gadgets)
    // Prints: ["Confused Drone" "" ""]
}
```

The first element is set to `"Confused Drone"`, and the remaining two elements are initialized to the zero value for strings (empty string `""`).

## Array Comparability

Arrays can be compared using `==` and `!=`, but **only if they have the same type**:

```go
gadgets := [3]string{"Confused Drone"}
gears   := [...]string{"Confused Drone"}

fmt.Println(gadgets == gears)  // Compile error!
```

**Why this fails:**
- `gadgets` has type `[3]string`
- `gears` has type `[1]string` (Go inferred length from the single element)
- Different lengths mean different types, making them incomparable

## Arrays Are Value Types

When you assign an array to another variable, **Go creates a complete copy** of the array:

```go
gadgets := [3]string{"Confused Drone", "Broken Phone"}
gears   := gadgets  // Creates a copy

gears[2] = "Shiny Mouse"

fmt.Printf("%q\n", gadgets)
// Prints: ["Confused Drone" "Broken Phone" ""]
```

The arrays are independent. Modifying `gears` doesn't affect `gadgets` because they're separate copies.

## Multidimensional Arrays

You can create arrays of arrays (multidimensional arrays):

```go
digits := [...][5]string{
    {
        "## ",
        " # ",
        " # ",
        " # ",
        "###",
    },
    [5]string{
        "###",
        "  #",
        "###",
        "  #",
        "###",
    },
}
// Type: [2][5]string
```

**Type breakdown:**
- Outer array: length 2 (two inner arrays)
- Inner arrays: length 5 each (five strings)
- Complete type: `[2][5]string`

Note: The `[5]string` type annotation on the second element is optional—Go can infer it.

## Keyed Elements

You can use indices as keys when initializing array elements:

```go
rates := [...]float64{
    5: 1.5,   // Set index 5 to 1.5
    2.5,      // Set index 6 to 2.5 (next position)
    0: 0.5,   // Set index 0 to 0.5
}

fmt.Printf("%#v\n", rates)
// Prints: [7]float64{0.5, 0, 0, 0, 0, 1.5, 2.5}
```

**How this works:**
1. `5: 1.5` sets index 5 to 1.5 (length becomes at least 6)
2. `2.5` goes to the next index (6), making length 7
3. `0: 0.5` sets index 0 to 0.5
4. All unspecified indices get zero values (0.0)

## Named Types and Comparability

### Underlying Types

Arrays with named types can be compared if they have the same underlying type:

```go
type three [3]int

nums  := [3]int{1, 2, 3}
nums2 := three{1, 2, 3}

fmt.Println(nums == nums2)  // true
```

This works because both `[3]int` and `three` have the same underlying type: `[3]int`.

### Type Conversion

You can compare arrays of different named types if you convert one to the other:

```go
type (
    threeA [3]int
    threeB [3]int
)

nums  := threeA{1, 2, 3}
nums2 := threeA(threeB{1, 2, 3})  // Convert threeB to threeA

fmt.Println(nums == nums2)  // true
```

Without the conversion, `threeA` and `threeB` would be incomparable even though they have the same underlying type, because they are distinct defined types. The conversion `threeA(...)` makes them comparable by ensuring both values have type `threeA`.
