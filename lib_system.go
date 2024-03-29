package easygolang

import (
	"bytes"
	"errors"
	"fmt" //print scan
	"os"
	"os/exec"
	"reflect"
	"regexp"
	"runtime"
	"runtime/debug"
	"sort"
	"strconv"
	"strings"
	"time"
	"unsafe"
)

var pattern *regexp.Regexp

func init() {
	pattern = regexp.MustCompile(`[^\w@%+=:,./-]`)
}

func AppProcessID() int {
	return os.Getpid()
}

func AppRunArgs() []string {
	return os.Args
}

func AppExit(code int) {
	os.Exit(code)
}

func SleepMS(ms int) {
	time.Sleep(time.Millisecond * time.Duration(ms))
}

func GarbageCollection() {
	runtime.GC()
}

func FreeOSMemory() {
	debug.FreeOSMemory()
}

func RuntimeLockOSThread() {
	runtime.LockOSThread()
}

func Prln(v string) {
	fmt.Println(v)
	//println(v)
}

func Scln() string {
	var input string
	fmt.Scanln(&input)
	return input
}

func ErrorWithText(txt string) error {
	return errors.New(txt)
}

func InterfaceNil(a interface{}) bool {
	defer func() { recover() }()
	return a == nil || reflect.ValueOf(a).IsNil()
}

func RuntimeSetFinalizer(obj interface{}, finalFunc interface{}) {
	runtime.SetFinalizer(obj, finalFunc)
}

//goroutine id
func GoId() int {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	return id
}

func GetGolangVersion() string {
	return runtime.Version()
}

// funcAddr returns function value fn executable code address.
func funcAddr(fn interface{}) uintptr {
	// emptyInterface is the header for an interface{} value.
	type emptyInterface struct {
		typ   uintptr
		value *uintptr
	}
	e := (*emptyInterface)(unsafe.Pointer(&fn))
	return *e.value
}

// if a, b, c = func() and you need only "b" or only "c"
func A(a ...interface{}) []interface{} {
	return a
}

// if a, b, c = func() and you need only "a"
func A0(a ...interface{}) interface{} {
	return A(a...)[0]
}

// func Ptr(p interface{}) *interface{} {
// 	return &p
// }

func CloneBytesArray(arr []byte) []byte {
	data2 := make([]byte, len(arr))
	copy(data2, arr)
	return data2
}

func SortArray(slice interface{}, less func(i, j int) bool) {
	sort.SliceStable(slice, less)
}

//B2S clone
// func Select_String(condition bool, istrue string, isfalse string) string {
// 	return map[bool]string{true: istrue, false: isfalse}[condition]
// }

func Select_Int(condition bool, istrue int, isfalse int) int {
	return map[bool]int{true: istrue, false: isfalse}[condition]
}

// ==========

func ExecQuote(s string) string {
	if len(s) == 0 {
		return "''"
	}
	if pattern.MatchString(s) {
		return "'" + strings.Replace(s, "'", "'\"'\"'", -1) + "'"
	}
	return s
}

func ExecCommandBytes(input []byte, timeout_ms int, forkill chan *exec.Cmd, exe_name string, args ...string) ([]byte, []byte, string) {
	cmd := exec.Command(exe_name, args...)
	var buffer_out bytes.Buffer
	var buffer_err bytes.Buffer
	cmd.Stdout = &buffer_out
	cmd.Stderr = &buffer_err

	var err error
	if len(input) > 0 {
		p, in_err := cmd.StdinPipe()
		err = cmd.Start()
		go func() {
			if forkill != nil {
				forkill <- cmd
			}
		}()
		// go func() {
		// 	cmd.Process.Kill()
		// }()
		if in_err == nil {
			p.Write(input)
			//p.Write([]byte("\000"))
			p.Close()
			Prln("cmd.StdinPipe() - Close()")
		} else {
			Prln(in_err.Error())
		}
		//cmd.Wait()
		Prln("cmd.Wait()?")
	} else {
		if timeout_ms > 0 {
			go func() {
				SleepMS(timeout_ms)
				cmd.Process.Kill()
			}()
		}
		err = cmd.Run()
	}

	if err != nil {
		return []byte{}, []byte{}, err.Error()
	}
	return buffer_out.Bytes(), buffer_err.Bytes(), ""

	//cmd := exec.Command("ls", "-lah")
	//out, err := cmd.CombinedOutput() // !!!!!!!!!!!!!!!! FIX
	//if err != nil {
	//    log.Fatalf("cmd.Run() failed with %s\n", err)
	//}
	//fmt.Printf("combined out:\n%s\n", string(out))
}

func ExecCommand(exe_name string, args ...string) (string, string, string) {
	r1, r2, err := ExecCommandBytes([]byte{}, 0, nil, exe_name, args...)
	return string(r1), string(r2), err
}
