package main

import (
    "fmt"
    "math/rand"
    "os"
    "time"
)

func main() {
    rand.Seed(time.Now().UnixNano())
    eyes := rand.Intn(6) + 1
    now := time.Now().Format("2025-11-10 15:04:05")

    fmt.Fprintln(os.Stdout, eyes)
    fmt.Fprintln(os.Stderr, now)  
}
