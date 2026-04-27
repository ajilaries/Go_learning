package commands

import (
    "fmt"
    "strconv"
)

func Calc(args []string) {
    if len(args) < 5 {
        fmt.Println("Usage: calc num1 op num2")
        return
    }

    a, _ := strconv.ParseFloat(args[2], 64)
    op := args[3]
    b, _ := strconv.ParseFloat(args[4], 64)

    switch op {
    case "+":
        fmt.Println(a + b)
    case "-":
        fmt.Println(a - b)
    case "*":
        fmt.Println(a * b)
    case "/":
        if b == 0 {
            fmt.Println("Cannot divide by zero")
            return
        }
        fmt.Println(a / b)
    default:
        fmt.Println("Invalid operator")
    }
}