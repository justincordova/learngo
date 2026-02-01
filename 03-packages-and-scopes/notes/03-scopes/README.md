## What is scope?

Scope determines the visibility of declared names in your code — where you can access variables, functions, and other identifiers.

## Example code

```go
package awesome

import "fmt"

var enabled bool

func block() {
    var counter int
    fmt.Println(counter)
}
```

## Package scope

Names declared outside any function are **package scoped**. All code in the same package can access them.

**Package scoped in example:** `enabled`, `block()`

## File scope

Imported package names are **file scoped**. Only code in the same file can use them.

**File scoped in example:** `fmt`

## Function/Block scope

Names declared inside a function are **function scoped**. Only code within that function can access them.

**Function scoped in example:** `counter` (only accessible inside `block()`)

## Visibility rules

**Can `block()` see `enabled`?**
Yes — `enabled` is package scoped, so all code in the package can see it.

**Can other files in `awesome` package see `counter`?**
No — `counter` is function scoped to `block()`. Only code inside `block()` can see it.

**Can other files in `awesome` package see `fmt`?**
No — imported packages are file scoped. Each file must import packages separately.

## Redeclaring names in the same scope

```go
var enabled bool
var enabled bool  // ERROR: already declared
```

You cannot declare the same name twice in the same scope.

## Shadowing (declaring in inner scope)

```go
package awesome

import "fmt"

var enabled bool  // package scope

func block() {
    var enabled bool  // function scope - shadows package scope

    var counter int
    fmt.Println(counter)
}
```

You **can** declare the same name in an inner scope. The inner declaration shadows (overrides) the outer one within that scope. Both can exist, but the inner scope sees only its own version.
