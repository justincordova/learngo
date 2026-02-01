# Switch Statement Notes

## Switch vs If Statements

Switch statements provide a more readable alternative to complex if-else chains, especially when comparing a single value against multiple possibilities.

## Switch Condition Types

In Go, switch statements can use any type of value as a condition:

```go
switch condition {
    // int, bool, string, or any other type
}
```

Unlike many C-based languages, Go's switch is syntactic sugar for if statements, allowing maximum flexibility.

## Case Type Matching

Case conditions must match the type of the switch condition:

```go
switch false {  // bool condition
case true:      // must use bool values
case false:
}

switch "go" {   // string condition
case "rust":    // must use string values
case "go":
}
```

## Duplicate Case Values

Each value can only appear in one case clause:

```go
switch 5 {
case 5:   // OK
case 6:   // OK
case 5:   // ERROR: 5 already used
}
```

## Multiple Conditions Per Case

Use commas to check multiple values in one case:

```go
weather := "hot"
switch weather {
case "very cold", "cold":
    fmt.Println("Bring a jacket")
case "very hot", "hot":     // This executes when weather is "hot"
    fmt.Println("Stay hydrated")
default:
    fmt.Println("Moderate weather")
}
```

## Default Clause

The `default` clause executes when no case matches. It can appear anywhere in the switch:

```go
switch weather := "too hot"; weather {
default:                    // Executes if no case matches
    fmt.Println("Unknown")
case "very cold", "cold":
    fmt.Println("Cold")
case "very hot", "hot":
    fmt.Println("Hot")
}
```

## No Match, No Default

If no case matches and there's no default clause, nothing executes:

```go
switch weather := "too hot"; weather {
case "very cold", "cold":
    // won't execute
case "very hot", "hot":
    // won't execute
}
// No output - "too hot" doesn't match any case
```

## Fallthrough

The `fallthrough` keyword forces execution to continue to the next case without checking its condition:

```go
richter := 2.5

switch r := richter; {
case r < 2:
    fallthrough
case r >= 2 && r < 3:     // This matches
    fallthrough            // Forces next case to execute
case r >= 5 && r < 6:
    fmt.Println("Not important")  // This executes due to fallthrough
case r >= 6 && r < 7:
    fallthrough
case r >= 7:
    fmt.Println("Destructive")
}
// Prints: "Not important"
```

## Fallthrough Restrictions

`fallthrough` must be the last statement in a case block:

```go
switch r := richter; {
case r >= 5 && r < 6:
    fallthrough            // Must be last
    fmt.Println("Moderate")  // ERROR: can't come after fallthrough
default:
    fmt.Println("Unknown")
}
```

## Expression-Based Switch

Switch statements can use boolean expressions without a condition value:

```go
n := 8

switch {  // No condition - evaluates case expressions
case n > 5, n < 1:      // Evaluates left-to-right, stops at first true
    fmt.Println("n is big")  // This executes (n > 5 is true)
case n == 8:
    fmt.Println("n is 8")    // Not reached - first case already matched
}
// Prints: "n is big"
```

Switch executes top-to-bottom, and case conditions evaluate left-to-right until one matches.
