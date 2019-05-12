package main

import (
	"flag"
	"fmt"

	"github.com/karlpokus/filew"
	"github.com/karlpokus/filew/internal/mem"
)

var (
	m     = flag.Bool("m", false, "toggle mem usage output")
	fpath = flag.String("p", ".", "path to dir to watch")
)

func main() {
	flag.Parse()
	if *m {
		go mem.Usage()
	}
	events, err := filew.Watch(*fpath, nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("watching %s\n", *fpath)
	for ev := range events {
		fmt.Println(ev)
	}
}
