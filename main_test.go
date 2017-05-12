package main

import (
	"fmt"
	"testing"
)

var mds = []string{
	`
	asfasf
	## hell
	hello`,
	`## hell
	#afdsf`,
	`## hell`,
	`# adf
	## hell`,
	`
	`,
}

func TestTitle(t *testing.T) {
	for _, md := range mds {
		fmt.Println(Title(md))
	}
}
