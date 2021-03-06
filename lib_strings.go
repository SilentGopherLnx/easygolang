package easygolang

import (
	"fmt"
	"regexp"
	"strings"
	"unicode/utf8"
)

func StringFormat(format string, args ...interface{}) string {
	return fmt.Sprintf(format, args...)
}

func StringLength(str string) int {
	runes := []rune(str)
	return len(runes)
}

//result from 1, end can be 0
func StringPart(str string, start int, end int) string {
	runes := []rune(str)
	rsize := len(runes)
	if rsize > 0 {
		start0 := MAXI(start, 1)
		if start0 <= rsize {
			end0 := MINI(MAXI(end, 1), rsize)
			if end > 0 {
				if end0 >= start0 {
					return string(runes[start0-1 : end0])
				} else {
					return ""
				}
			} else {
				return string(runes[start0-1:])
			}
		} else {
			return ""
		}
	} else {
		return ""
	}
}

func StringEnd(str string, count int) string {
	runes := []rune(str)
	rsize := len(runes)
	esize := MINI(MAXI(count, 0), rsize)
	if rsize == 0 || esize == 0 {
		return ""
	}
	if rsize == esize {
		return str
	}
	return string(runes[rsize-esize:])
}

func StringFill(str string, length int) string {
	return fmt.Sprintf("%"+I2S(length)+"s", str)
}

func StringClip(str string, maxlen int) string {
	if StringLength(str) > maxlen {
		return StringPart(str, 1, maxlen)
	} else {
		return str
	}
}

func StringJoin(target []string, separator string) string {
	return strings.Join(target, separator)
}

func StringSplit(target string, separator string) []string {
	return strings.Split(target, separator)
}

func StringSplitLines(target string) []string {
	return strings.Split(target, "\n")
}

func StringReplace(str string, from string, to string) string {
	return strings.Replace(str, from, to, -1)
}

//not found 0, start from 1
func StringFind(where string, what string) int {
	ind := strings.Index(where, what)
	if ind > -1 {
		return utf8.RuneCountInString(where[:ind]) + 1
	} else {
		return 0
	}
}

//not found 0, start from 1
func StringFindEnd(where string, what string) int {
	ind := strings.LastIndex(where, what)
	if ind > -1 {
		return utf8.RuneCountInString(where[:ind]) + 1
	} else {
		return 0
	}
}

func StringUp(str string) string {
	return strings.ToUpper(str)
}

func StringDown(str string) string {
	return strings.ToLower(str)
}

func StringTitle(str string) string {
	return strings.Title(str)
}

func StringTrim(str string) string {
	return strings.TrimSpace(str)
}

func StringRemoveDoubleChar(str string, one_char string) string {
	chars := one_char + one_char
	str2 := ""
	str3 := str
	for str3 != str2 {
		str2 = str3
		str3 = StringReplace(str2, chars, one_char)
	}
	return str3
}

func StringCommonPrefix(a string, b string) string {
	r1 := []rune(a)
	r2 := []rune(b)
	min_len := MINI(len(r1), len(r2))
	res := ""
	j := 0
	for j < min_len && r1[j] == r2[j] {
		res += string(r1[j])
		j++
	}
	return res
}

func StringCommonPrefixArr(a []string) string {
	lena := len(a)
	if lena == 0 {
		return ""
	}
	if lena == 1 {
		return a[0]
	}
	min_len := len(a[0])
	r := make([][]rune, lena)
	for j := 0; j < lena; j++ {
		r[j] = []rune(a[j])
		min_len = MINI(min_len, len(r[j]))
	}
	res := ""
	j := 0
	ok := true
	for j < min_len && ok {
		for k := 1; k < lena; k++ {
			if r[k-1][j] != r[k][j] {
				ok = false
			}
		}
		if ok {
			res += string(r[0][j])
			j++
		}
	}
	return res
}

func StringInArray(val string, array []string) int {
	for i := range array {
		if ok := array[i] == val; ok {
			return i
		}
	}
	return -1
}

var utf8_bom = string([]byte{239, 187, 191}) // UTF-8 specific

func StringFixUTF8(line string) string {
	return strings.TrimPrefix(line, utf8_bom)
}

// ======

type RegExp struct {
	r *regexp.Regexp
}

func StringFilterCompile(pattern string) RegExp {
	return RegExp{regexp.MustCompile(pattern)}
}

func StringFilter(str string, reg RegExp) string {
	return reg.r.ReplaceAllString(str, "")
}
