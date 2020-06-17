package main

import (
	"go.dev.pztrn.name/flagger"
)

var f *flagger.Flagger

func main() {
	f = flagger.New("testprogram", nil)
	f.Initialize()
	_ = f.AddFlag(&flagger.Flag{
		Name:         "testflag",
		Description:  "Just a test flag",
		Type:         "bool",
		DefaultValue: false,
	})

	f.Parse()
}
