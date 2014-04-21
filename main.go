package main

import (
  "os"
  "fmt"
  "translate"
)

func main() {
  str := os.Args[1]
  transformed := piglatin.Translate(str)
  fmt.Println(transformed)
}
