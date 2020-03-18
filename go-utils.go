package main

import (
	"flag"
)

func main() {
	beegoPtr := flag.Bool("beego", false, "")
	pathPtr := flag.String("path", "", "")
	flag.Usage = func() {}
	flag.Parse()

	if beegoPtr != nil && *beegoPtr && pathPtr != nil {
		path := *pathPtr
		RunBeego(path)
	}
}
