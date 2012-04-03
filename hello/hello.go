// Hello is a trivial example of a main package.
package main

import (
        "github.com/jcao/Go/newmath"
        "fmt"
)

func main() {
        fmt.Printf("Hello from Github, world.  Sqrt(9) = %v\n", newmath.Sqrt(9))
}