package main

import (
    "fmt"
    "time"
)

func main() {
    // Get the current time
    currentTime := time.Now()
    fmt.Println("Current Time:", currentTime)

    // Create a specific time using time.Date
    specificTime := time.Date(2024, time.June, 5, 12, 0, 0, 0, time.UTC)
    fmt.Println("Specific Time:", specificTime)

    // Extract and print individual components of the specific time
    year := specificTime.Year()
    month := specificTime.Month()
    day := specificTime.Day()
    hour := specificTime.Hour()
    minute := specificTime.Minute()
    second := specificTime.Second()
    fmt.Printf("Year: %d, Month: %s, Day: %d, Hour: %d, Minute: %d, Second: %d\n", year, month, day, hour, minute, second)

    // Calculate and print the difference between two times
    timeDiff := currentTime.Sub(specificTime)
    fmt.Println("Time Difference:", timeDiff)

    // Format time in a specific layout
    formattedTime := specificTime.Format("January 2, 2006 15:04:05")
    fmt.Println("Formatted Time:", formattedTime)
}


/*

# Output

Current Time: 2024-06-05 19:00:38.8133335 +0800 +08 m=+0.002022501
Specific Time: 2024-06-05 12:00:00 +0000 UTC
Year: 2024, Month: June, Day: 5, Hour: 12, Minute: 0, Second: 0
Time Difference: -59m21.1866665s
Formatted Time: June 5, 2024 12:00:00

*/