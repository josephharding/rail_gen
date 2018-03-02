package main

import (
  "fmt"
  "io/ioutil"
)

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func main() {
	fmt.Printf("writing output...\n")
  d1 := []byte("hello output")
  err := ioutil.WriteFile("data/phase0", d1, 0644)
  check(err)
}
