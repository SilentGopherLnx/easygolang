package tests

import (
	"testing"

	. "github.com/SilentGopherLnx/easygolang"
)

// func TestMain(m *testing.M) {
// 	Prln("Start!")
// }

// func Bench_all() {
// 	Prln("?")
// }

func Test_StringPart(t *testing.T) {
	if StringPart("12ABC", 2, 3) != "2A" {
		Log_TestFailed(t)
	}
	if StringPart("12ABC", 0, 3) != "12A" {
		Log_TestFailed(t)
	}
	if StringPart("12ABC", 0, 0) != "12ABC" {
		Log_TestFailed(t)
	}
	if StringPart("12ABC", 2, 0) != "2ABC" {
		Log_TestFailed(t)
	}
	if StringPart("12ABC", 0, 9) != "12ABC" {
		Log_TestFailed(t)
	}
	if StringPart("12ABC", 2, 9) != "2ABC" {
		Log_TestFailed(t)
	}
	if StringPart("12ABC", 3, 2) != "" {
		Log_TestFailed(t)
	}
	if StringPart("", 2, 9) != "" {
		Log_TestFailed(t)
	}
	if StringPart("", 0, 0) != "" {
		Log_TestFailed(t)
	}
}

func Test_StringEnd(t *testing.T) {

}

func Test_StringFillClip(t *testing.T) {

}

func Test_StringFind(t *testing.T) {

}

func Test_StringRemoveDoubleChar(t *testing.T) {

}

func Test_StringCommonPrefix(t *testing.T) {

}
