package main

import (
  "fmt"
  "github.com/karlpokus/filew"
)

func main() {
  fp := "testdata/file1"
  events, err := filew.Watch(fp, nil)
  if err != nil {
    panic(err)
  }
  fmt.Printf("watching %s\n", fpath)
  for e := range events {
    fmt.Println(e)
  }
}
