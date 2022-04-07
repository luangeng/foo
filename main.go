package foo

import "fmt"

func init() {
    fmt.Println("foo init invoke")
}

func Bar() string {
    return "bar"
}