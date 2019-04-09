package easylinux

// #include <unistd.h>
import "C"

import (
	. "github.com/SilentGopherLnx/easygolang"
)

func LinuxMemoryTotalMB() int {
	return int(uint64(C.sysconf(C._SC_PHYS_PAGES)*C.sysconf(C._SC_PAGE_SIZE)) / BytesInMb)
}

func LinixMemoryUsedMB(pid int) float64 {
	res, _, _ := ExecCommandBash("cat /proc/" + I2S(pid) + "/status | grep RssAnon")
	res = StringReplace(StringReplace(StringDown(res), "rssanon:", ""), "kb", "")
	v := S2I(StringTrim(res))
	return float64(v) / 1024.0
}
