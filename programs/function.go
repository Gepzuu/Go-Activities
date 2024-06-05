//The add function accepts two integer values, x and y, and produces their sum as an integer result.

package main

import "fmt"

func add(x int, y int) int {
    return x + y
}

func main() {
    result := add(4, 83)
    fmt.Println("The sum is:", result)
}

/*

# Output

The sum is: 87

*/