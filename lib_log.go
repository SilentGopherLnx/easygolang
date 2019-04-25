package easygolang

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"testing"
)

type Log struct {
	Arr     [][]string
	Counter int
	Mutex   *SyncMutex
	MaxSize int
}

func LogNew(MaxSize int) *Log {
	return &Log{Arr: [][]string{}, Counter: 1, Mutex: NewSyncMutex(), MaxSize: MaxSize}
}

func (log *Log) Push(new_msg string) {
	if len(new_msg) > 0 {
		t := TimeNowStr()
		Prln(new_msg)
		log.Mutex.Lock()
		new_rec := []string{I2S(log.Counter), new_msg, t}
		log.Counter++

		if len(log.Arr) == 0 {
			log.Arr = [][]string{new_rec}
		} else {
			log.Arr = append(log.Arr, new_rec)
		}
		delta := len(log.Arr) - log.MaxSize
		if delta > 0 {
			//tmp := make([][]string, len(log.Arr)-delta)
			//copy(tmp, log.Arr[delta:])
			//log.Arr = tmp
			log.Arr = log.Arr[delta:]
		}
		log.Mutex.Unlock()
	}
}

func (log *Log) Print(separator string, reverse bool) string {
	res := ""
	log.Mutex.Lock()
	if !reverse {
		for j := 0; j < len(log.Arr); j++ {
			res += "[" + log.Arr[j][2] + "][" + log.Arr[j][0] + "]: " + log.Arr[j][1] + separator
		}
	} else {
		for j := len(log.Arr) - 1; j >= 0; j-- {
			res += "[" + log.Arr[j][2] + "][" + log.Arr[j][0] + "]: " + log.Arr[j][1] + separator
		}
	}
	log.Mutex.Unlock()
	return res
}

// ====================

func Log_TypeOf(v interface{}) string {
	if v == nil {
		return "nil-i"
	}
	t := reflect.TypeOf(v)
	if t == nil {
		return "nil-rt"
	}
	s := t.String()
	if s == "" {
		return "?"
	}
	return s
}

func Log_TypeDetail(t interface{}) string {
	return fmt.Sprintf("%#v\n", t)
}

func AboutVersion(version string) {
	if AppHasArg("-v") {
		Prln(version)
		Prln("app  folder:[" + FolderLocation_App() + "]")
		Prln("work folder:[" + FolderLocation_WorkDir() + "]")
		Prln("home folder:[" + FolderLocation_UserHome() + "]")
		AppExit(0)
	}
}

func AppHasArg(v string) bool {
	args := AppRunArgs()
	for j := 1; j < len(args); j++ {
		if args[j] == v {
			return true
		}
	}
	return false
}

// ==================

func Log_WhereAmI() string {
	return whereAmI()
}

func Log_TestFailed(t *testing.T) {
	Prln(whereAmI())
	t.Fail()
}

func whereAmI(depthList ...int) string {
	var depth int
	if depthList == nil {
		depth = 1
	} else {
		depth = depthList[0]
	}
	depth++
	function, file, line, _ := runtime.Caller(depth)
	ind := strings.LastIndex(file, "/")
	if ind != -1 {
		file = file[ind+1:]
	}
	return fmt.Sprintf("File: %s  Function: %s Line: %d", file, runtime.FuncForPC(function).Name(), line)
}
