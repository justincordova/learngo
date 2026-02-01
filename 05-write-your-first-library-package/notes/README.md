## Library packages

Library packages provide reusable code that can be imported by other packages.

**Key facts:**
- You cannot run a library package directly
- Library packages don't need a `main` function (only executable packages do)
- You can compile a library package, but you don't have to — it's automatically built when imported
- When you import a library, it gets compiled along with your program

## Exporting names

To export a name (function, variable, type, etc.), capitalize its first letter:

```go
// Exported - accessible from other packages
func Calculate() {}
var MaxValue int

// Not exported - only accessible within the same package
func helper() {}
var minValue int
```

Exported names must be at package scope, not inside functions.

## Using functions from your library

To use a function from your library in an executable program:

1. Import the library package using its module path
2. Access the exported names from that package

You must import the package using its module path as declared in `go.mod`.


## Example: Exported names

```go
package wizard

import "fmt"

func doMagic() {
    fmt.Println("enchanted!")
}

func Fireball() {
    fmt.Println("fireball!!!")
}
```

**Exported:** `Fireball` (starts with capital letter)

**Not exported:** `doMagic` (starts with lowercase letter)

Note: `fmt` is an imported package name, and `Println` is exported from the `fmt` package, but neither are exported from this `wizard` package.

## Example: Multiple exported names

```go
package wizard
import "fmt"

var one string
var Two string
var greenTrees string

func doMagic() {
    fmt.Println("enchanted!")
}

func Fireball() {
    fmt.Println("fireball!!!")
}
```

**Exported:** `Fireball` and `Two` (both start with capital letters)

**Not exported:** `doMagic`, `one`, and `greenTrees` (all start with lowercase letters)
