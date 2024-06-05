package main

import "fmt"

func main() {

    // Check if 7 is odd or even
    if 7%2 == 1 {
        fmt.Println("The number 7 is an odd number.")
    } else {
        fmt.Println("The number 7 is an even number.")
    }

    // Check if 8 is divisible by 4
    if 8%4 == 0 {
        fmt.Println("The number 8 is divisible by 4.")
    }

    // Check if either 8 or 7 is even
    if 8%2 == 0 || 7%2 == 0 {
        fmt.Println("At least one of the numbers 8 or 7 is even.")
    }

    num := 9
    // Check the number's properties (negative, single-digit, or multiple-digit)
    if num < 0 {
        fmt.Println("The number", num, "is negative.")
    } else if num < 10 {
        fmt.Println("The number", num, "is a single-digit number.")
    } else {
        fmt.Println("The number", num, "has more than one digit.")
    }
}

/*

# Output

The number 7 is an odd number.
The number 8 is divisible by 4.
At least one of the numbers 8 or 7 is even.
The number 9 is a single-digit number.

*/