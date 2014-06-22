# `visual_width` package

This is a go-lang package to deal with visual width, or Unicode East Asian Width.

## Usage

```go
package main

import (
    "fmt"
    "github.com/gfx/go-visual_width"
)

func main() {
    var str = "国境の長いトンネルを抜けると雪国であった"
    var inEastAsian = true
    var maxWidth = 20
    var omission = "..."
    var s = visual_width.Truncate(str, inEastAsian, maxWidth, omission)
    fmt.Printf("%q\n", s)
    // -> "国境の長いトンネ..."
}
```

## SPECIFICATION

* http://www.unicode.org/reports/tr11/
* http://www.unicode.org/Public/UNIDATA/EastAsianWidth.txt

## LICENSE

This library can be distributed under the MIT License.

Copyright (c) 2013 Fuji, Goro (gfx) <gfuji@cpan.org>.

Permission is hereby granted, free of charge, to any person obtaining a
copy of this software and associated documentation files (the "Software"),
to deal in the Software without restriction, including without limitation
the rights to use, copy, modify, merge, publish, distribute, sublicense,
and/or sell copies of the Software, and to permit persons to whom the
Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER
DEALINGS IN THE SOFTWARE.

