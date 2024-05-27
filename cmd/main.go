package main

import (
    "fmt"
    "Homework_testing/pkg/factorial"
    "Homework_testing/pkg/reverse"
    "Homework_testing/pkg/wordcount"
)

func main() {
    fmt.Println("Factorial:", factorial.Factorial(12))
    fmt.Println("Reversed string:", reverse.Reverse("hello"))
    fmt.Println("Word count:", wordcount.Wordcount("hello world"))
}