/*
In this code using go, I model swimming pools with an interface and two structs, 
implementing area and perimeter methods. 
The measure function calculates and displays pool details.
*/

package main

import (
    "fmt"
    "math"
)

// Define an interface named pool with area() and perimeter() methods
type pool interface {
    area() float64
    perimeter() float64
}

// Define a struct for rectangular pools
type rectangle struct {
    length, width float64
}

// Define a struct for circular pools
type circle struct {
    radius float64
}

// Method to calculate the area of a rectangle
func (r rectangle) area() float64 {
    return r.length * r.width
}

// Method to calculate the perimeter of a rectangle
func (r rectangle) perimeter() float64 {
    return 2*r.length + 2*r.width
}

// Method to calculate the area of a circle
func (c circle) area() float64 {
    return math.Pi * c.radius * c.radius
}

// Method to calculate the perimeter of a circle
func (c circle) perimeter() float64 {
    return 2 * math.Pi * c.radius
}

// Function to measure and print details of a pool (implements pool interface)
func measure(p pool) {
    fmt.Println("Pool Details:")
    fmt.Println("Area:", p.area())
    fmt.Println("Perimeter:", p.perimeter())
}

func main() {
    // Create a rectangular pool instance
    rectanglePool := rectangle{length: 10, width: 5}

    // Create a circular pool instance
    circularPool := circle{radius: 7}

    // Measure and print details of the rectangular pool
    measure(rectanglePool)

    // Measure and print details of the circular pool
    measure(circularPool)
}


/*

# Output

Pool Details:
Area: 50
Perimeter: 30
Pool Details:
Area: 153.93804002589985
Perimeter: 43.982297150257104

*/