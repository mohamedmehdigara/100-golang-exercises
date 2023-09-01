package main

import (
    "fmt"
    "exercise005"
)

func main() {
    fmt.Println("Enter a string:")
    exercise005.ReadString()

    result := exercise005.PrintString()
    fmt.Printf("Uppercase: %s\n", result)
}
