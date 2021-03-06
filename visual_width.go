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
var t1 = regexp.MustCompile(fmt.Sprintf(`(%s)|(.|[\n])`, r1))
var t0 = regexp.MustCompile(fmt.Sprintf(`(%s)|(.|[\n])`, r0))

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

func Wrap(str string, inEastAsian bool, maxWidth int) string {
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

		c := str[offset : offset+runeLen]
		if c == "\n" {
			buffer.WriteString(c)
			width = 0
			offset += runeLen
			continue
		}

		if (width + w) > maxWidth {
			buffer.WriteString("\n")
			width = 0
		}
		buffer.WriteString(c)

		offset += runeLen
		width += w
	}

	return buffer.String()
}
