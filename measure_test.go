package visual_width

import (
	a "github.com/stretchr/testify/assert"
	"testing"
)

func TestMeasureAscii(t *testing.T) {
	a.Equal(t, Measure(""), 0)

	a.Equal(t, Measure("foo"), 3)
	a.Equal(t, Measure("..."), 3)
}

func TestMeasureHiragana(t *testing.T) {
	a.Equal(t, Measure("こんにちは"), 10)
	a.Equal(t, Measure("こ ん に ち は"), 14)
}

func TestMeasureHalfKatakana(t *testing.T) {
	a.Equal(t, Measure("ｺﾝﾆﾁﾊ"), 5)
	a.Equal(t, Measure("ｺ ﾝ ﾆ ﾁ ﾊ"), 9)
}

func TestMeasureAmbiguous(t *testing.T) {
	a.Equal(t, Measure("αβγδε"), 10)
	a.Equal(t, Measure("α β γ δ ε"), 14)
}

func TestMeasureMixed(t *testing.T) {
	a.Equal(t, Measure("こんにちは ｺﾝﾆﾁﾊ αβγδε"), 27)
}
