package easylinux

import (
	. "github.com/SilentGopherLnx/easygolang"
)

func LinuxFolderFreeSpace(folder string) int64 {
	out, _, _ := ExecCommand("df", "-T", "-h", "-B", "1", folder)
	//Prln(out)
	l := StringSplitLines(out)
	if len(l) >= 2 {
		str := l[1]
		str = StringRemoveDoubleChar(str, " ")
		arr := StringSplit(str, " ")
		if len(arr) >= 7 {
			return S2I64(arr[4])
		}
	}
	return 0
}

func LinuxFileGetParent(path string) string {
	tpath := path
	mv := 1
	if StringEnd(tpath, 1) == "/" {
		mv = 2
	}
	s := StringSplit(tpath, "/")
	s = s[:len(s)-mv]
	tpath = StringJoin(s, "/")
	if len(tpath) == 0 {
		tpath = "/"
	}
	return tpath
}

func LinuxFileNameFromPath(path string) string {
	tpath := path
	if StringEnd(tpath, 1) == "/" {
		tpath = StringPart(tpath, 1, StringLength(tpath)-1)
	}
	s := StringSplit(tpath, "/")
	if len(s) > 0 {
		return s[len(s)-1]
	}
	return ""
}
