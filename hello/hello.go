// Hello is a trivial example of a main package.
package main

import (
        "github.com/jcao/Go/newmath"
        "fmt"
)

func main() {
        fmt.Printf("Hello from Github, world!!!!.  Sqrt(16) = %v\n", newmath.Sqrt(16))
        fmt.Printf("Hello from Github, world!!!!.  Sqrt(64) = %v\n", newmath.Sqrt(64))
}