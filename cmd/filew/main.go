package main

import (
  "fmt"
  "github.com/karlpokus/filew"
)

func main() {
  fpath := "testdata"
  events, err := filew.Watch(fpath)
  if err != nil {
    panic(err)
  }
  fmt.Printf("watching %s\n", fpath)
  for ev := range events {
    fmt.Println(ev)
  }
}
