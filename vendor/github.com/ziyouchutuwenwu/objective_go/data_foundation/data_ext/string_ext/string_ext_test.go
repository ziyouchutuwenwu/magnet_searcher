package string_ext

import "testing"

func TestStringExt(t *testing.T) {
	aa := "别码代码了，码代码能娶媳妇吗？"
	bb := "01234"
	cc := "asdasdasdasd"

	t.Log(Lenth(aa))

	dd := Append(aa,bb,cc)
	t.Log(dd)

	x := SubString(dd,2,4)
	t.Log(x)
}
