package main

import (
    "fmt"
    "time"
)

func main() {
    pet := "dog" 

    // Switch statement to print the type of pet
    switch pet {
    case "dog":
        fmt.Println("I have a dog.")
    case "cat":
        fmt.Println("I have a cat.")
    case "fish":
        fmt.Println("I have a fish.")
    default:
        fmt.Println("I have another type of pet.")
    }

    // Switch statement based on the day of the week
    switch time.Now().Weekday() {
    case time.Saturday, time.Sunday:
        fmt.Println("It's the weekend.")
    default:
        fmt.Println("It's a weekday.")
    }

    // Switch statement based on the time of day
    t := time.Now()
    switch {
    case t.Hour() < 12:
        fmt.Println("It's morning.")
    default:
        fmt.Println("It's afternoon.")
    }

    // Function to identify the type of pet
    whatPetAmI := func(pet interface{}) {
        switch p := pet.(type) {
        case string:
            fmt.Println("I have a", p)
        default:
            fmt.Println("I have another type of pet.")
        }
    }

    // Call the function to identify the type of pet
    whatPetAmI("dog")
    whatPetAmI(1) 
    whatPetAmI(true) 
}

/*

# Output

I have a dog.
It's a weekday.
It's afternoon.
I have a dog
I have another type of pet.
I have another type of pet.


*/