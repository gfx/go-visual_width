package main

import (
	"fmt"
	"github.com/gfx/go-visual_width"
)

func main() {
	var inEastAsianContext = true
	var maxWidth = 20
	var omission = "..."
	var s = visual_width.Truncate("国境の長いトンネルを抜けると雪国であった", inEastAsianContext, maxWidth, omission)
	fmt.Printf("%q\n", s)
	// -> "国境の長いトンネ..."
}
