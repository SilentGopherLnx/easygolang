package easygolang

import (
	"crypto/md5"
	"crypto/sha1"
	"math/rand"
	"strconv"
	"time"
)

//var BOOL_STYLE_10 = [2]string{"1", "0"}
//var BOOL_STYLE_TRUEFALSE = [2]string{"true", "false"}
//var BOOL_STYLE_YESNO = [2]string{"yes", "no"}

const INT_SEPARATOR = ' '

func init() {
	RandNew()
	//Prln("random genetator inited")
}

func XOR(X bool, Y bool) bool {
	return (X || Y) && !(X && Y)
}

func B2S(v bool, str_true string, str_false string) string {
	if v {
		return str_true
	} else {
		return str_false
	}
}

func B2S_10(v bool) string {
	return B2S(v, "1", "0")
}

func B2S_TF(v bool) string {
	return B2S(v, "true", "false")
}

func B2S_YN(v bool) string {
	return B2S(v, "yes", "no")
}

//for sort
func CompareBoolLess(a bool, b bool) bool {
	ai := 0
	bi := 0
	if a {
		ai = 1
	}
	if b {
		bi = 1
	}
	return ai < bi
}

// =================

func RandNew() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func RandInt(min int, max int) int {
	return rand.Intn(max-min) + min
}

// =================

func I2S(v int) string {
	return strconv.Itoa(v)
}

func I2S64(v int64) string {
	return strconv.FormatInt(v, 10)
}

func S2I(v string) int {
	num, _ := strconv.Atoi(v)
	return num
}

func S2I64(v string) int64 {
	num, _ := strconv.ParseInt(v, 10, 64)
	return num
}

func IsInt(v string) bool {
	_, err := strconv.Atoi(v)
	return err == nil
}

func I2Ss(v int64) string {
	in := strconv.FormatInt(v, 10)
	out := make([]byte, len(in)+(len(in)-2+int(in[0]/'0'))/3)
	if in[0] == '-' {
		in, out[0] = in[1:], '-'
	}
	for i, j, k := len(in)-1, len(out)-1, 0; ; i, j = i-1, j-1 {
		out[j] = in[i]
		if i == 0 {
			return string(out)
		}
		if k++; k == 3 {
			j, k = j-1, 0
			out[j] = INT_SEPARATOR
		}
	}
	return I2S64(v)
}

func F2S(v float64, prec int) string {
	return strconv.FormatFloat(v, 'f', prec, 64)
}

func S2F(v string) float64 {
	num, _ := strconv.ParseFloat(v, 64)
	return num
}

func MAXI(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func MINI(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func MAXI64(a int64, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func MINI64(a int64, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func MAXF(a float64, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func MINF(a float64, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

func AbcI(a int) int {
	if a >= 0 {
		return a
	} else {
		return -a
	}
}

func AbcF(a float64) float64 {
	if a >= 0 {
		return a
	} else {
		return -a
	}
}

func RoundF(a float64) int {
	return int(a + 0.499999)
}

func IntInArray(val int, array []int) int {
	for i := range array {
		if ok := array[i] == val; ok {
			return i
		}
	}
	return -1
}

// =================

func Crypto_MD5(data []byte) string {
	// var s
	// fmt.SPrintf(&s, "%x", md5.Sum(data))
	// return s
	return StringFormat("%x", md5.Sum(data))
}

func Crypto_SHA1(data []byte) string {
	// var s string = ""
	// fmt.SPrintf(&s, "%x", sha1.Sum(data))
	// return s
	return StringFormat("%x", sha1.Sum(data))
}
