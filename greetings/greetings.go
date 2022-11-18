// create-module-codes
// package greetings

// import "fmt"

// // Hello returns a greeting for the named person.
// func Hello(name string) string {
//     // Return a greeting that embeds the name in a message.
//     message := fmt.Sprintf("Hi, %v. Welcome!", name)
//     return message
// }

// return and handle an error
package greetings

import (
    "errors"
    "fmt"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
    // If no name was given, return an error with a message.
    if name == "" {
        return "", errors.New("empty name")
    }

    // If a name was received, return a value that embeds the name
    // in a greeting message.
    message := fmt.Sprintf("Hi, %v. Welcome!", name)
    return message, nil
}