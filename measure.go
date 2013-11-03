package visual_width

import (
	"regexp"
	"unicode/utf8"
)

var r1 = fullwide + "|" + wide + "|" + ambiguous
var r0 = fullwide + "|" + wide

var c1 = regexp.MustCompile(r1)
var c0 = regexp.MustCompile(r0)

func Measure(str string) int {
	var count = utf8.RuneCountInString(str)
	var nFullWidthChars = len(c1.FindAllString(str, -1))
	return nFullWidthChars*2 + (count - nFullWidthChars)
}
