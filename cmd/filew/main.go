package main

import (
  "fmt"
  "github.com/karlpokus/filew"
)

func main() {
  fpath := "testdata/file1"
  events, err := filew.Watch(fpath)
  if err != nil {
    panic(err)
  }
  fmt.Printf("watching %s\n", fpath)
  for e := range events {
    fmt.Println(e)
  }
}
