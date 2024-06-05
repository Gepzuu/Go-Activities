package main

import (
    "fmt"
    "errors"
)

func main() {
    // Array of strings to represent airports
    airportCodes := [5]string{"MNL", "CEB", "DVO", "BCD", "CRK"}

    // Check if the array length is less than 5
    if len(airportCodes) < 5 {
        err := errors.New("airportCodes array length is less than 5")
        fmt.Println(err)
        return
    }

    // Print out the airports array
    fmt.Println("Airport codes in the Philippines:", airportCodes)
}

/*

# Output

Airport codes in the Philippines: [MNL CEB DVO BCD CRK]

*/