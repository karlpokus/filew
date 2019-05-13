package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/karlpokus/filew"
	"github.com/karlpokus/filew/internal/mem"
)

var (
	m     = flag.Bool("m", false, "toggle mem usage output")
	fpath = flag.String("p", ".", "path to dir to watch")
	version = flag.Bool("version", false, "print cli version and exit")
)

func main() {
	flag.Parse()
	if *version {
		fmt.Println(filew.Version)
		os.Exit(0)
	}
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
