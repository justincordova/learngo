# Structs

## When to Use Struct Types

Use a struct type to **combine different types in a single type to represent a concept**:

```go
type Person struct {
    name string
    age  int
    city string
}
```

**When NOT to use structs:**
- For storing the same type of values → use arrays, slices, or maps
- For dynamic fields that can be added/removed at runtime → structs have fixed fields at compile-time

Structs are ideal for modeling real-world concepts with multiple properties of different types.

## Struct Field Properties

Each struct field must have:
- **A unique name** within the struct
- **A type** (can be different from other fields)

```go
type weather struct {
    temperature float64
    humidity    float64
    windSpeed   float64
}
```

You can define multiple fields of the same type in parallel:

```go
type weather struct {
    temperature, humidity float64  // Parallel definition
    windSpeed             float64
}
```

### Field Name Uniqueness

Field names must be unique within a struct:

```go
type weather struct {
    temperature, humidity float64
    windSpeed             float64
    temperature           float64  // Error: duplicate field name
}
```

This will cause a compile error because `temperature` is defined twice.

## Zero Values

When you declare a struct variable without initialization, its fields are set to their respective zero values:

```go
var movie struct {
    title, genre string
    rating       float64
    released     bool
}

// Zero value: {title: "", genre: "", rating: 0, released: false}
```

Each field gets the zero value for its type:
- `string` → `""`
- `float64` → `0`
- `bool` → `false`

## Struct Types

A struct's type includes **both field names and field types**:

```go
avengers := struct {
    title, genre string
    rating       float64
    released     bool
}{
    "avengers: end game", "sci-fi", 8.9, true,
}

fmt.Printf("%T\n", avengers)
// Prints: struct{ title string; genre string; rating float64; released bool }
```

Two structs are the same type only if they have the exact same field names and types in the same order.

## Creating Struct Values

You can create struct values with or without field names:

```go
type movie struct {
    title, genre string
    rating       float64
    released     bool
}

// Without field names (positional)
avengers := movie{"avengers: end game", "sci-fi", 8.9, true}

// With field names (named)
clone := movie{
    title: "avengers: end game", genre: "sci-fi",
    rating: 8.9, released: true,
}

fmt.Println(avengers == clone)  // true
```

Both forms create equivalent values. Named fields are more explicit and allow you to specify fields in any order or omit some fields (they'll get zero values).

## Comparing Struct Values

Structs are comparable if all their fields are comparable:

```go
type movie struct {
    title, genre string
    rating       float64
    released     bool
}

avengers := movie{
    title: "avengers: end game", genre: "sci-fi",
    rating: 8.9, released: true,
}

clone := movie{title: "avengers: end game", genre: "sci-fi"}
// clone.rating = 0, clone.released = false (zero values)

fmt.Println(avengers == clone)  // false
```

When you omit fields, they get zero values. Here, `clone` has `rating: 0` and `released: false`, so it's not equal to `avengers`.

## Named Types and Comparison

Struct types with different names are **not comparable**, even if they have identical fields:

```go
type item        struct { title string }
type movie       struct { item }
type performance struct { item }

// movie and performance are different types
```

However, you can convert one to the other because they have the same underlying structure:

```go
m := movie{}
p := performance{}

// fmt.Println(m == p)  // Error: different types
fmt.Println(m == movie(p))  // OK: convert performance to movie
```

## Embedded Fields

Embedded fields allow you to compose structs:

```go
type item struct{ title string }

type movie struct {
    item
    title string
}

m := movie{
    title: "avengers: end game",
    item:  item{"midnight in paris"},
}

fmt.Println(m.title, "&", m.item.title)
// Prints: avengers: end game & midnight in paris
```

**Field shadowing:** When the outer and inner types have fields with the same name:
- `m.title` accesses the outer field (avengers: end game)
- `m.item.title` explicitly accesses the inner field (midnight in paris)

The outer type always takes priority for unqualified access.

## Field Tags

Field tags are **metadata** associated with struct fields, written as string literals:

```go
type movie struct {
    title string `json:"title"`
}
```

**Key points about field tags:**
- They are just string values with no inherent meaning
- Other packages (like `json`, `xml`, `yaml`) read and interpret them
- They cannot be changed at runtime (they're part of the type definition)
- While they can have any value, packages expect specific formats (e.g., `key:"value"`)

## JSON Encoding and Exported Fields

The `json` package can only encode/decode **exported** (capitalized) fields:

```go
type movie struct {
    title string `json:"title"`  // Unexported - won't be encoded
}

m := movie{"black panthers"}
encoded, _ := json.Marshal(m)

fmt.Println(string(encoded))  // Prints: {}
```

**Fix:** Export the field:

```go
type movie struct {
    Title string `json:"title"`  // Exported - will be encoded
}

m := movie{"black panthers"}
encoded, _ := json.Marshal(m)

fmt.Println(string(encoded))  // Prints: {"title":"black panthers"}
```

## Unmarshaling with Pointers

You must pass a **pointer** to `json.Unmarshal` so it can update the original value:

```go
var m movie
err := json.Unmarshal(data, &m)  // Pass pointer with &
```

**Why?** Go is pass-by-value. Without the pointer, `Unmarshal` would receive a copy and could only modify the copy, leaving your original value unchanged. By passing a pointer, you give `Unmarshal` the memory address so it can modify the original value directly.
