package main

import "fmt"

func main() {

    // Concatenate and print the strings "Number: " and "123"
    fmt.Println("GO " + "Number")

    // Add 2 and 2, and print the result with a description
    fmt.Println("2+2 =", 2+2)
    
    // Divide 10.0 by 2.0, and print the result with a description
    fmt.Println("10.0/2.0 =", 10.0/2.0)
    
    // Evaluate the expression 5 > 3 (greater than) and print the result
    fmt.Println(5 > 3)
    
    // Evaluate the expression 7 < 2 (less than) and print the result
    fmt.Println(7 < 2)
    
    // Evaluate the expression !(4 == 4), which negates the result of 4 == 4
    // 4 == 4 is true, so !true is false, and print the result
    fmt.Println(!(4 == 4))
}

/*This code presents a range of operations, including string concatenation, 
integer addition, floating-point division, 
and logical operations that involve comparisons and the NOT operator. 
The output is adjusted to incorporate numbers.*/