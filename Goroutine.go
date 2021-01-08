package main

import (
    "fmt"
    "time"
)
 
 
func main() {
    ch := make(chan int, 5)
 
    // Start a goroutine that reads a value from the channel every second and prints it
    go func(ch chan int) {
        for {
            time.Sleep(time.Second)
            fmt.Printf("Goroutine received: %d\n", <-ch)
        }
 
    }(ch)
 
    // Start a goroutine that prints a dash every second
    go func() {
        for i := 0; i < 5; i++ {
            time.Sleep(time.Second)
            fmt.Println("-")
        }
    }()
 
    // Push values to the channel as fast as possible
    for i := 0; i < 5; i++ {
        ch <- i
        fmt.Printf("main() pushed: %d\n", i)
    }
 
    // Sleep five more seconds to let all goroutines finish
    time.Sleep(5 * time.Second)
}
