/*
I use two functions: zeroCalories sets a calorie value to zero within the function, and 
zeroPtrCalories uses a pointer to directly modify the original calorie variable. This demonstrates 
the difference between modifying values by value and by pointer in Go.
*/

package main

import "fmt"

func zeroCalories(calories int) {
    calories = 0
}

func zeroPtrCalories(calPtr *int) {
    *calPtr = 0
}

func main() {
    burgerCal := 500
    fmt.Println("Initial calories in burger:", burgerCal)

    zeroCalories(burgerCal)
    fmt.Println("Calories after zeroval:", burgerCal)

    zeroPtrCalories(&burgerCal)
    fmt.Println("Calories after zeroptr:", burgerCal)

    fmt.Println("Memory address of burgerCal:", &burgerCal)
}

/*

# Output

Initial calories in burger: 500
Calories after zeroval: 500
Calories after zeroptr: 0
Memory address of burgerCal: 0xc00000a0f8

*/