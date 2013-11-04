package visual_width

import (
	a "github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

var str = strings.Trim(`
Googleは2009年11月10日(米国時間)、オープンソースのプログラミング言語「Go」を発表しました。
Go言語は、Linux、Mac、Native Clientで動作する開発言語で、Android携帯上でも動作します。
まだ発表されたばかりなのでこれからの動向が注目されています。

特徴はGoogleによると・・

・シンプルな言語である。
・コンパイル・実行速度が早い。
・安全性が高い。
・同期処理が容易に行える。
・なにより楽しい。
・オープンソースである。

これらについてもう少し掘り下げてみてみましょう。
`, " \r\n")

func TestMeasureAscii(t *testing.T) {
	a.Equal(t, Measure("", true), 0)

	a.Equal(t, Measure("foo", true), 3)
	a.Equal(t, Measure("...", true), 3)

	a.Equal(t, Measure("", false), 0)

	a.Equal(t, Measure("foo", false), 3)
	a.Equal(t, Measure("...", false), 3)
}

func TestMeasureHiragana(t *testing.T) {
	a.Equal(t, Measure("こんにちは", true), 10)
	a.Equal(t, Measure("こ ん に ち は", true), 14)

	a.Equal(t, Measure(str, true), 459)
}

func TestMeasureHalfKatakana(t *testing.T) {
	a.Equal(t, Measure("ｺﾝﾆﾁﾊ", true), 5)
	a.Equal(t, Measure("ｺ ﾝ ﾆ ﾁ ﾊ", true), 9)

	a.Equal(t, Measure("ｺﾝﾆﾁﾊ", false), 5)
	a.Equal(t, Measure("ｺ ﾝ ﾆ ﾁ ﾊ", false), 9)
}

func TestMeasureAmbiguous(t *testing.T) {
	a.Equal(t, Measure("αβγδε", true), 10)
	a.Equal(t, Measure("α β γ δ ε", true), 14)

	a.Equal(t, Measure("αβγδε", false), 5)
	a.Equal(t, Measure("α β γ δ ε", false), 9)
}

func TestMeasureMixed(t *testing.T) {
	a.Equal(t, Measure("こんにちは ｺﾝﾆﾁﾊ αβγδε", true), 27)
	a.Equal(t, Measure("こんにちは ｺﾝﾆﾁﾊ αβγδε", false), 22)
}

func TestTruncateAscii(t *testing.T) {
	a.Equal(t, Truncate("abcdefghij", true, 7, "..."), "abcd...")
	a.Equal(t, Truncate("abcdefghij", true, 8, "..."), "abcde...")
	a.Equal(t, Truncate("abcdefghij", true, 9, "..."), "abcdef...")
	a.Equal(t, Truncate("abcdefghij", true, 10, "..."), "abcdefghij")
}

func TestTruncateHiragana(t *testing.T) {
	a.Equal(t, Truncate("こんにちは", true, 7, "..."), "こん...")
	a.Equal(t, Truncate("こんにちは", true, 8, "..."), "こん...")
	a.Equal(t, Truncate("こんにちは", true, 9, "..."), "こんに...")
	a.Equal(t, Truncate("こんにちは", true, 10, "..."), "こんにちは")
}

func TestWrap(t *testing.T) {
	a.Equal(t, Wrap("こんにちは", true, 10), "こんにちは")
	a.Equal(t, Wrap("こんにちはこんにちは", true, 10), "こんにちは\nこんにちは")
	a.Equal(t, Wrap("こんにちは\nこんにちは", true, 10), "こんにちは\nこんにちは")
}

// benchmarks

func BenchmarkMeasure(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var _ = Measure(str, true)
	}
}

func BenchmarkTruncate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var _ = Truncate(str, true, 100, "...")
	}
}
