// In this program tickers are used when you need to perform a task at consistent intervals repeatedly.

package main

import (
    "fmt"
    "time"
)

func main() {

    ticker := time.NewTicker(500 * time.Millisecond)
    done := make(chan bool)

    go func() {
        for {
            select {
            case <-done:
                return
            case t := <-ticker.C:
                fmt.Println("Tick at", t)
            }
        }
    }()

    time.Sleep(1600 * time.Millisecond)
    ticker.Stop()
    done <- true
    fmt.Println("Ticker stopped")
}

/*

# Output

Tick at 2024-06-05 18:53:57.7275694 +0800 +08 m=+0.505179401
Tick at 2024-06-05 18:53:58.2343142 +0800 +08 m=+1.011924201
Tick at 2024-06-05 18:53:58.7247729 +0800 +08 m=+1.502382901
Ticker stopped

*/
