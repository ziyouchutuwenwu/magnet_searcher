package data_convert

import (
	"bytes"
	"encoding/gob"
	"strconv"
)

func StrToInt(srcStr string) int {
	integer, _ := strconv.Atoi(srcStr)
	return integer
}

func IntToStr(integer int) string {
	return strconv.Itoa(integer)
}

func BytesToString(b []byte) string {
	return string(b)
}

func StringToBytes(s string) []byte {
	return []byte(s)
}

func FloatToStr(number float64) string {
	return strconv.FormatFloat(number, 'f', -1, 64)
}

func FloatToStrWithFloatPartLenth(number float64, floatPartLenth int) string {
	return strconv.FormatFloat(number, 'f', floatPartLenth, 64)
}

func StrToFloat(str string) float64{
	value ,_ := strconv.ParseFloat(str,64)
	return value
}

func StringToBool(isTrueStr string) bool{
	isTrue, _ := strconv.ParseBool(isTrueStr)
	return isTrue
}

func BoolToString(isTrue bool) string{
	return strconv.FormatBool(isTrue)
}

func IntToBool(val int) bool{
	var result bool = val != 0
	return result
}

func BoolToInt(val bool) int {
	if val {
		return 1
	}
	return 0
}

func InterfaceToBytes(val interface{}) []byte{
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)
	err := encoder.Encode(val)
	if err != nil {
		return nil
	}
	return buf.Bytes()
}