// Hello is a trivial example of a main package.
package main

import (
        "code.google.com/p/go.example/newmath"
        "fmt"
)

func main() {
        fmt.Printf("Hello, world.  Sqrt(2) = %v\n", newmath.Sqrt(2))
}