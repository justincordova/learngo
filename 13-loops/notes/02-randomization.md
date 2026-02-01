# Randomization in Go

## Pseudorandom Number Generation

**Pseudorandom numbers** appear to be randomly generated, but they are actually produced by a deterministic algorithm. Computers are deterministic machines and cannot generate truly random numbers without external input from physical processes (like atmospheric noise or radioactive decay).

Instead, computers use mathematical algorithms to generate sequences of numbers that have statistical properties similar to random sequences, but are entirely predictable if you know the starting state.

## Seed Numbers

A **seed number** is used to initialize a pseudorandom number generator. The seed determines the sequence of numbers that will be generated—the same seed will always produce the same sequence of pseudorandom numbers.

```go
rand.Seed(42)  // Initialize with seed 42
```

This deterministic behavior is useful for:
- Debugging (reproducible results)
- Testing (consistent test outcomes)
- Simulations that need to be repeatable

## The `rand` Package

Go's `math/rand` package provides pseudorandom number generation:

```go
import "math/rand"

// Generate a random number in the range [0, 10)
n := rand.Intn(10)
```

## Mathematical Interval Notation

The notation `[0, 5)` represents a range of numbers:
- `[` (square bracket) means **inclusive** (includes the number)
- `)` (parenthesis) means **exclusive** (excludes the number)

**Examples:**
- `[0, 5)` = 0, 1, 2, 3, 4 (includes 0, excludes 5)
- `[0, 5]` = 0, 1, 2, 3, 4, 5 (includes both 0 and 5)
- `(0, 5)` = 1, 2, 3, 4 (excludes both 0 and 5)

## Using `rand.Intn()`

The `rand.Intn(n)` function generates a pseudorandom number in the range `[0, n)`.

### Invalid Usage

```go
rand.Intn(0)  // Error!
```

This doesn't work because `Intn(0)` would need to return a number in the range `[0, 0)`, which means including 0 but excluding 0—an impossible constraint.

## How Seeds Affect Output

Seeding the random number generator with the same value always produces the same sequence:

```go
package main

import (
    "fmt"
    "math/rand"
)

func main() {
    for i := 0; i < 3; i++ {
        rand.Seed(int64(i))
        fmt.Print(rand.Intn(11), " ")
        fmt.Print(rand.Intn(11), " ")
    }
}
// Output: 3 3 1 1 10 1
```

**Explanation:**
- Loop iteration 0: `Seed(0)` produces 3, 3
- Loop iteration 1: `Seed(1)` produces 1, 1
- Loop iteration 2: `Seed(2)` produces 10, 1

Each time you reseed, the generator resets to the beginning of that seed's sequence.

## Generating Different Numbers Each Run

To get different random numbers every time your program runs, seed the generator with the current time:

```go
import (
    "math/rand"
    "time"
)

func main() {
    rand.Seed(time.Now().UnixNano())
    // Now each run will produce different numbers
}
```

`time.Now().UnixNano()` returns the current time as a Unix timestamp in nanoseconds, which is different each time you run the program. This ensures a different seed and thus a different sequence of pseudorandom numbers.

**Note:** In Go 1.20+, the global random number generator is automatically seeded, so explicit seeding is no longer necessary for most use cases. However, understanding seeding is still important for reproducible results in testing and simulations.
