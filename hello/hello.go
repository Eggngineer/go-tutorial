// getting-started :1
// package main

// import "fmt"

// func main() {
//     fmt.Println("Hello, World!")
// }

// getting-started :2

// package main

// import "fmt"

// import "rsc.io/quote"

// func main() {
//     fmt.Println(quote.Go())
// }

// call-module-code :1
package main

import (
    "fmt"

    "example.com/greetings"
)

func main() {
    // Get a greeting message and print it.
    message := greetings.Hello("Eggngineer")
    fmt.Println(message)
}