package easygolang

import (
	"os" //args
	"os/user"
	"runtime" //GOOS, GOARCH
	//"syscall"
	//	"unsafe" //WARNING
)

const OS_WINDOWS = 1
const OS_MAC = 2
const OS_LINUX = 3

func GetOS() int {
	// /usr/local/go/bin/go tool dist list -json
	os := runtime.GOOS
	switch os {
	case "windows":
		return OS_WINDOWS
	case "darwin":
		return OS_MAC
	case "linux":
		return OS_LINUX
	default:
		return 0
	}
}

func GetOS_Name() string {
	// /usr/local/go/bin/go tool dist list -json
	os := runtime.GOOS
	switch os {
	case "windows":
		return "Windows"
	case "darwin":
		return "Mac"
	case "linux":
		return "Linux"
	default:
		return os
	}
}

func GetOS_Bits() int {
	if StringFind(runtime.GOARCH, "s390x") > 0 { // s390 - 32, s390x - 64
		return 64
	}
	if StringFind(runtime.GOARCH, "32") > 0 { // amd64p32 mips64p32 mips64p32le
		return 32
	}
	if StringFind(runtime.GOARCH, "64") > 0 { // amd64 arm64arm64be ppc64 ppc64le mips64 mips64le sparc64
		return 64
	} else {
		return 32 // 386 arm armbe mips mipsle ppc s390 sparc
	}
}

func GetOS_Slash() string {
	/*os := runtime.GOOS
	if os == "windows" {
		return "\\"
	} else {
		return "/"
	}*/
	return string(os.PathSeparator)
}

// ===============

func GetPC() string {
	name, _ := os.Hostname()
	return name
}

func GetPC_UserUidLoginName() (string, string, string) {
	user, err := user.Current()
	if err != nil {
		return "", "", ""
	}
	return user.Uid, user.Username, user.Name
}

func GetPC_CountCores() int {
	return runtime.NumCPU()
}

//by app
func GetPC_MemoryUsageMb() float64 {
	// https://golang.org/pkg/runtime/#MemStats
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return float64(m.Alloc) / float64(BytesInMb)

}

/*func GetPC_MemoryMaxMb() int {
	// https://golang.org/pkg/runtime/#MemStats
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return int(m. / BytesInMb)

}*/

/*func GetPC_MemoryCountMb() int {
	os := GetOS()
	if os == OS_LINUX {
		in := &syscall.Sysinfo_t{}
		err := syscall.Sysinfo(in)
		if err != nil {
			return 0
		}
		return int(uint64(in.Totalram) * uint64(in.Unit) / BytesInMb)
	} else if os == OS_MAC {
		syscall.Sysctl
		s, err := syscall.sysctlUint64("hw.memsize")
		if err != nil {
			return 0
		}
		return int(s / BytesInMb)
	}
	return 0
}*/

/*//total, not only app
func GetPC_UsagePercentCPU() int {
	//clockSeconds := float64(C.clock()-startTicks) / float64(C.CLOCKS_PER_SEC)
	//realSeconds := time.Since(startTime).Seconds()
	//return clockSeconds / realSeconds * 100
	return -1
}*/
