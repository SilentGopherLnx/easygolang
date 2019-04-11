package easygolang

import (
	"fmt"
	"reflect" //typeof
)

type Log struct {
	Arr     [][]string
	Counter int
	Mutex   *SyncMutex
	MaxSize int
}

func NewLog(MaxSize int) *Log {
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

func TypeOf(v interface{}) string {
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

func TypeDetail(t interface{}) string {
	return fmt.Sprintf("%#v\n", t)
}

func AboutVersion(version string) {
	args := AppRunArgs()
	if len(args) == 2 && args[1] == "-v" {
		Prln(version)
		AppExit(0)
	}
}
