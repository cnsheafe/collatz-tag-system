package main

import (
  "fmt"
  "strings"
)

func main()  {
  input := "aaa"
  prodRules := map[byte]string {
    'a': "bc",
    'b': "a",
    'c': "aaa",
  }

  for len(input) > 1 {
    fmt.Println(input)
    head := input[0]
    input = input[2:]
    input = strings.Join([]string{input, prodRules[head]}, "")
  }
  fmt.Printf("%v\n(halt)\n", input)
}
