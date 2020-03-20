package main

import (
	"flag"
	"fmt"
)

func main() {
	beegoPtr := flag.Bool("beego", false, "")
	pathPtr := flag.String("path", "", "")
	utilsPathPtr := flag.String("utilsPath", "", "")
	oldModelsPathPtr := flag.String("oldModelsPath", "", "")
	newModelsPathPtr := flag.String("newModelsPath", "", "")
	flag.Usage = func() {}
	flag.Parse()

	fmt.Println("go-utils v1.2")
	if beegoPtr != nil && *beegoPtr && pathPtr != nil && utilsPathPtr != nil && oldModelsPathPtr != nil && newModelsPathPtr != nil {
		path := *pathPtr
		utilsPath := *utilsPathPtr
		oldModelsPath := *oldModelsPathPtr
		newModelsPath := *newModelsPathPtr
		RunBeego(path, utilsPath, oldModelsPath, newModelsPath)
	}
}
