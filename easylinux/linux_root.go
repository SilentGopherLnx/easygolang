package easylinux

import (
	"fmt"
	"os"
	"os/exec"

	//"path/filepath"
	"strconv"
)

func CheckRootLinux() int {
	cmd := exec.Command("id", "-u")
	output, err := cmd.Output()

	if err != nil {
		fmt.Println(err)
		return -1
	}

	// output has trailing \n
	// need to remove the \n
	// otherwise it will cause error for strconv.Atoi
	// log.Println(output[:len(output)-1])

	// 0 = root, 501 = non-root user
	i, err := strconv.Atoi(string(output[:len(output)-1]))

	if err != nil {
		fmt.Println(err)
		return -1
	}

	if i == 0 {
		return 1
	} else {
		return 0
	}
}

func StartRootLinux(app string, cmd string) {
	//    pkexec /mnt/dm-1/golang/code/ssh/ssh
	/*dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Println(err)
	}*/
	fmt.Println(os.Args[0])
	var obj *exec.Cmd
	if len(cmd) > 0 {
		obj = exec.Command("pkexec", app, cmd)
	} else {
		obj = exec.Command("pkexec", app)
	}
	output, err := obj.Output()

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(output))
	}
}
