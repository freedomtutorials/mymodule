package mymodule

import "fmt"

func MyFunc(name string) string {
    message := fmt.Sprintf("Hello, %v. Welcome!", name)
    return message
}
