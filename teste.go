package main

import (
	"fmt"
	"log"
	"os"
)


func main() {
    // Simple print
    fmt.Println("This is a standard log line")

    // Formatted print
    name := "Go Container"
    fmt.Printf("Logging from %s\n", name)

    // Using log package for more control
    log.SetOutput(os.Stdout)
    log.Println("This is a log message using log package")

    // Error logging
    err := fmt.Errorf("this is a sample error")
    if err != nil {
        log.Printf("Error occurred: %v\n", err)
    }
}