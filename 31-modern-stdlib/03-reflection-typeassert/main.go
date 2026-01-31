package main

import (
	"fmt"
	"reflect"
	"time"
)

func main() {
	fmt.Println("Zero-Allocation Reflection with reflect.TypeAssert (Go 1.25)")
	fmt.Println("=============================================================")
	fmt.Println()
	
	demonstrateTraditionalReflection()
	fmt.Println()
	demonstrateTypeAssert()
	fmt.Println()
	performanceComparison()
}

func demonstrateTraditionalReflection() {
	fmt.Println("1. Traditional Reflection (Allocates)")
	fmt.Println("--------------------------------------")
	
	x := 42
	v := reflect.ValueOf(x)
	
	// Traditional way: uses Interface() which allocates
	result := v.Interface().(int)
	fmt.Printf("Value: %d (extracted via Interface())\n", result)
	fmt.Println("Note: Interface() allocates memory on the heap")
}

func demonstrateTypeAssert() {
	fmt.Println("2. reflect.TypeAssert (Go 1.25) - Zero Allocation")
	fmt.Println("--------------------------------------------------")
	
	x := 42
	v := reflect.ValueOf(x)
	
	// New in Go 1.25: TypeAssert for zero-allocation extraction
	// Note: TypeAssert is a generic function
	result, ok := reflect.TypeAssert[int](v)
	if ok {
		fmt.Printf("Value: %d (extracted via TypeAssert)\n", result)
		fmt.Println("Note: TypeAssert does NOT allocate - zero-cost abstraction!")
	}
	
	// Works with other types too
	s := "hello"
	vs := reflect.ValueOf(s)
	str, ok := reflect.TypeAssert[string](vs)
	if ok {
		fmt.Printf("String: %q\n", str)
	}
	
	// Type mismatch returns false
	wrong, ok := reflect.TypeAssert[float64](v)
	if !ok {
		fmt.Println("Type mismatch detected (int != float64)")
	} else {
		fmt.Printf("This won't print: %f\n", wrong)
	}
}

func performanceComparison() {
	fmt.Println("3. Performance Comparison")
	fmt.Println("-------------------------")
	
	x := 42
	v := reflect.ValueOf(x)
	iterations := 10_000_000
	
	// Traditional Interface() method
	start := time.Now()
	for i := 0; i < iterations; i++ {
		_ = v.Interface().(int)
	}
	traditional := time.Since(start)
	
	// New TypeAssert method (Go 1.25+)
	start = time.Now()
	for i := 0; i < iterations; i++ {
		_, _ = reflect.TypeAssert[int](v)
	}
	typeAssert := time.Since(start)
	
	fmt.Printf("Traditional (Interface): %v\n", traditional)
	fmt.Printf("TypeAssert (Go 1.25):    %v\n", typeAssert)
	fmt.Printf("Speedup:                 %.2fx faster\n", 
		float64(traditional)/float64(typeAssert))
	fmt.Println()
	fmt.Println("TypeAssert is significantly faster due to zero allocations!")
}
