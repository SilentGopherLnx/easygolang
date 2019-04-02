package easylinux

import (
	. "github.com/SilentGopherLnx/easygolang"
)

func FolderLinuxFreeSpace(folder string) int64 {
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
