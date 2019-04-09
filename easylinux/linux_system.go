package easylinux

import (
	. "github.com/SilentGopherLnx/easygolang"
)

func ExecCommandBash(arg string) (string, string, string) {
	return ExecCommand("bash", "-c", arg)
}
