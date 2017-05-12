package string_ext

import (
	"bytes"
	"unicode/utf8"
	"strings"
)

func Lenth(srcStr string) int {
	return utf8.RuneCountInString(srcStr)
}

func SubString(srcStr string, beginPos, length int) string {

	srcRuneStr := []rune(srcStr)
	runeStrLen := len(srcRuneStr)

	// 简单的越界判断
	if beginPos < 0 {
		beginPos = 0
	}
	if beginPos >= runeStrLen {
		beginPos = runeStrLen
	}
	endPos := beginPos + length
	if endPos > runeStrLen {
		endPos = runeStrLen
	}

	// 返回子串
	return string(srcRuneStr[beginPos:endPos])
}

func Append(srcStr string, strToAppends ...string) string{
	array := append([]string{srcStr},strToAppends...)

	var buffer bytes.Buffer
	for _, element := range array {
		buffer.WriteString(element)
	}

	return buffer.String()
}

func IsNonSensitiveEqual(str1 string, str2 string) bool{
	return strings.EqualFold(str1, str2)
}