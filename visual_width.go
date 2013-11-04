package visual_width

import (
	"bytes"
	"fmt"
	"regexp"
	"unicode/utf8"
)

var r1 = fullwide + "|" + wide + "|" + ambiguous
var r0 = fullwide + "|" + wide

// Regexps for Measure()
var c1 = regexp.MustCompile(r1)
var c0 = regexp.MustCompile(r0)

// Regexp for Truncate()
var t1 = regexp.MustCompile(fmt.Sprintf("(%s)|(.)", r1))
var t0 = regexp.MustCompile(fmt.Sprintf("(%s)|(.)", r0))

func Measure(str string, inEastAsian bool) int {
	var re *regexp.Regexp
	if inEastAsian {
		re = c1
	} else {
		re = c0
	}

	var count = utf8.RuneCountInString(str)
	var nFullWidthChars = len(re.FindAllString(str, -1))
	return nFullWidthChars*2 + (count - nFullWidthChars)
}

func Truncate(str string, inEastAsian bool, maxWidth int, omission string) string {
	if Measure(str, inEastAsian) <= maxWidth {
		return str
	}

	var re *regexp.Regexp
	if inEastAsian {
		re = t1
	} else {
		re = t0
	}

	var buffer bytes.Buffer
	var width = 0

	maxWidth -= Measure(omission, inEastAsian)

	for offset := 0; offset < len(str); {
		loc := re.FindStringSubmatchIndex(str[offset:])
		wEnd := loc[3]
		hEnd := loc[5]
		var w int
		var runeLen int
		if wEnd != -1 {
			w = 2
			runeLen = wEnd
		} else {
			w = 1
			runeLen = hEnd
		}

		if (width + w) > maxWidth {
			break
		}
		buffer.WriteString(str[offset : offset+runeLen])

		offset += runeLen
		width += w
	}

	buffer.WriteString(omission)
	return buffer.String()
}
