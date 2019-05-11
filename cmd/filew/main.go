package main

import (
  "fmt"
  "github.com/karlpokus/filew"
)

func main() {
  fpath := "testdata/small"
  events, err := filew.Watch(fpath, nil)
  if err != nil {
    panic(err)
  }
  fmt.Printf("watching %s\n", fpath)
  for ev := range events {
    fmt.Println(ev)
  }
}
