package main

import "github.com/xifanyan/opentext/discovery/data/arm"

func main() {
	arm.NewArmDBBuilder().WithBasePath("testdata").WithID("db01").Build()
}
