package main

import (
    "fmt"
    "os"
    "devtool/commands"
)

func main() {
    args := os.Args

    if len(args) < 2 {
        fmt.Println("Usage: command")
        return
    }

    switch args[1] {

    case "greet":
        commands.Greet(args)

    case "calc":
        commands.Calc(args)

    default:
        fmt.Println("Unknown command")
    }
}