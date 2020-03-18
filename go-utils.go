package main

import (
	"flag"
)

func main() {
	beegoPtr := flag.Bool("beego", false, "")
	pathPtr := flag.String("path", "", "")
	utilsPathPtr := flag.String("utilSpath", "", "")
	oldModelsPathPtr := flag.String("oldModelsPath", "", "")
	newModelsPathPtr := flag.String("newModelsPath", "", "")
	flag.Usage = func() {}
	flag.Parse()

	if beegoPtr != nil && *beegoPtr && pathPtr != nil && utilsPathPtr != nil && oldModelsPathPtr != nil && newModelsPathPtr != nil {
		path := *pathPtr
		utilsPath := *utilsPathPtr
		oldModelsPath := *oldModelsPathPtr
		newModelsPath := *newModelsPathPtr
		RunBeego(path, utilsPath, oldModelsPath, newModelsPath)
	}
}
