package main

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

func main() {
	fmt.Println("Type Constraints in Go")
	fmt.Println("======================")
	fmt.Println()

	// Example 1: Built-in constraints
	fmt.Println("1. Built-in constraints (any, comparable):")
	PrintAny(42)
	PrintAny("hello")
	PrintAny(3.14)
	PrintAny([]int{1, 2, 3})

	fmt.Println("\nComparable values:")
	fmt.Printf("Equal(5, 5): %v\n", Equal(5, 5))
	fmt.Printf("Equal(\"hello\", \"world\"): %v\n", Equal("hello", "world"))
	fmt.Println()

	// Example 2: constraints.Ordered
	fmt.Println("2. constraints.Ordered (supports <, >, <=, >=):")
	fmt.Printf("Max(10, 20) = %d\n", Max(10, 20))
	fmt.Printf("Max(3.14, 2.71) = %.2f\n", Max(3.14, 2.71))
	fmt.Printf("Max(\"apple\", \"banana\") = %s\n", Max("apple", "banana"))
	fmt.Println()

	// Example 3: Custom numeric constraint
	fmt.Println("3. Custom Number constraint:")
	intNums := []int{1, 2, 3, 4, 5}
	floatNums := []float64{1.1, 2.2, 3.3, 4.4, 5.5}

	fmt.Printf("Sum of %v = %d\n", intNums, Sum(intNums))
	fmt.Printf("Sum of %v = %.2f\n", floatNums, Sum(floatNums))
	fmt.Printf("Average of %v = %.2f\n", intNums, Average(intNums))
	fmt.Printf("Average of %v = %.2f\n", floatNums, Average(floatNums))
	fmt.Println()

	// Example 4: Custom constraint with methods
	fmt.Println("4. Custom constraint with methods:")
	rect := Rectangle{Width: 5, Height: 3}
	circle := Circle{Radius: 4}

	fmt.Printf("Rectangle area: %.2f\n", GetArea(rect))
	fmt.Printf("Circle area: %.2f\n", GetArea(circle))
	fmt.Println()

	// Example 5: Union type constraints
	fmt.Println("5. Union type constraints:")
	fmt.Printf("IsZero(0) = %v\n", IsZero(0))
	fmt.Printf("IsZero(42) = %v\n", IsZero(42))
	fmt.Printf("IsZero(0.0) = %v\n", IsZero(0.0))
	fmt.Printf("IsZero(\"\") = %v\n", IsZero(""))
	fmt.Printf("IsZero(\"hello\") = %v\n", IsZero("hello"))
	fmt.Println()

	// Example 6: Constraint composition
	fmt.Println("6. Constraint composition:")
	nums1 := []int{5, 2, 8, 1, 9}
	nums2 := []float64{3.14, 2.71, 1.41, 1.73}

	fmt.Printf("Original: %v, Min: %d, Max: %d\n", nums1, MinSlice(nums1), MaxSlice(nums1))
	fmt.Printf("Original: %v, Min: %.2f, Max: %.2f\n", nums2, MinSlice(nums2), MaxSlice(nums2))
}

// PrintAny accepts any type (no constraints)
func PrintAny[T any](value T) {
	fmt.Printf("Value: %v (type: %T)\n", value, value)
}

// Equal checks if two comparable values are equal
// comparable allows use of == and !=
func Equal[T comparable](a, b T) bool {
	return a == b
}

// Max returns the larger of two ordered values
// constraints.Ordered includes all types that support < operator
func Max[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// Number is a custom constraint for numeric types
// The | symbol creates a union of types
type Number interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 |
		float32 | float64
}

// Sum calculates the sum of numeric values
func Sum[T Number](numbers []T) T {
	var total T
	for _, n := range numbers {
		total += n
	}
	return total
}

// Average calculates the average of numeric values
// Returns float64 regardless of input type
func Average[T Number](numbers []T) float64 {
	if len(numbers) == 0 {
		return 0
	}
	total := Sum(numbers)
	return float64(total) / float64(len(numbers))
}

// Shape is a constraint requiring an Area method
// This is a method-based constraint
type Shape interface {
	Area() float64
}

// GetArea works with any type that implements the Shape interface
func GetArea[T Shape](s T) float64 {
	return s.Area()
}

// Rectangle implements Shape
type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Circle implements Shape
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14159 * c.Radius * c.Radius
}

// SignedInteger is a constraint for signed integer types
type SignedInteger interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

// UnsignedInteger is a constraint for unsigned integer types
type UnsignedInteger interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

// Integer is a composition of signed and unsigned integers
// The ~ operator means "underlying type"
type Integer interface {
	SignedInteger | UnsignedInteger
}

// ZeroCheckable is a constraint for types that can be compared to zero
type ZeroCheckable interface {
	int | float64 | string
}

// IsZero checks if a value is the zero value for its type
func IsZero[T ZeroCheckable](value T) bool {
	var zero T
	return value == zero
}

// MinSlice finds the minimum value in a slice
// Uses constraints.Ordered to support comparison
func MinSlice[T constraints.Ordered](slice []T) T {
	if len(slice) == 0 {
		var zero T
		return zero
	}

	min := slice[0]
	for _, v := range slice[1:] {
		if v < min {
			min = v
		}
	}
	return min
}

// MaxSlice finds the maximum value in a slice
func MaxSlice[T constraints.Ordered](slice []T) T {
	if len(slice) == 0 {
		var zero T
		return zero
	}

	max := slice[0]
	for _, v := range slice[1:] {
		if v > max {
			max = v
		}
	}
	return max
}
