package easygolang

import (
	"math"
	"runtime"
	"sync"
	"sync/atomic"
	//	"unsafe"
)

func RuntimeGoMaxProcs(num int) {
	runtime.GOMAXPROCS(num)
}

func RuntimeGosched() {
	runtime.Gosched()
}

// ===

func InfiniteLoop() {
	select {}
}

func ChanGetOrSkip(ch chan interface{}) (interface{}, bool) {
	select {
	case v := <-ch:
		return v, true
	default:
		//CONTINUE!!
		return nil, false
	}
}

// ===

type SyncMutex struct {
	m *sync.Mutex
}

func NewSyncMutex() *SyncMutex {
	return &SyncMutex{m: new(sync.Mutex)}
}

func (s *SyncMutex) Lock() {
	s.m.Lock()
}

func (s *SyncMutex) Unlock() {
	s.m.Unlock()
}

// ===

//https://github.com/sheerun/queue/blob/master/queue.go
type SyncQueue struct {
	m   *sync.Mutex
	arr []interface{}
}

func NewSyncQueue() *SyncQueue {
	return &SyncQueue{m: new(sync.Mutex), arr: []interface{}{}}
}

func (s *SyncQueue) Length() int {
	s.m.Lock()
	length := len(s.arr)
	s.m.Unlock()
	return length
}

func (s *SyncQueue) Clear() {
	s.m.Lock()
	s.arr = []interface{}{}
	s.m.Unlock()
}

func (s *SyncQueue) Append(elem interface{}) {
	s.m.Lock()
	s.arr = append(s.arr, elem)
	s.m.Unlock()
}

func (s *SyncQueue) GetBegin() interface{} {
	s.m.Lock()
	var elem interface{}
	len0 := len(s.arr)
	if len0 > 0 {
		elem = s.arr[0]
		arr2 := make([]interface{}, len0-1)
		copy(arr2, s.arr[1:len0])
		s.arr = arr2
	}
	s.m.Unlock()
	return elem
}

func (s *SyncQueue) GetEnd() interface{} {
	s.m.Lock()
	var elem interface{}
	len1 := len(s.arr) - 1
	if len1 >= 0 {
		elem = s.arr[len1]
		arr2 := make([]interface{}, len1)
		copy(arr2, s.arr[0:len1])
		s.arr = arr2
	}
	s.m.Unlock()
	return elem
}

// ===

type SyncWaitGroup struct {
	w *sync.WaitGroup
}

func NewSyncWaitGroup() *SyncWaitGroup {
	return &SyncWaitGroup{w: new(sync.WaitGroup)}
}

func (s *SyncWaitGroup) Add(delta int) {
	s.w.Add(delta)
}

func (s *SyncWaitGroup) Done() {
	s.w.Done()
}

func (s *SyncWaitGroup) Wait() {
	s.w.Wait()
}

// ===

type SyncOnce struct {
	o *sync.Once
}

func (s *SyncOnce) Do(f func()) {
	s.o.Do(f)
}

// ===

type AInt struct {
	v int32
}

func NewAtomicInt(initv int) *AInt {
	return &AInt{v: int32(initv)}
}

func (a *AInt) Add(move int) int {
	return int(atomic.AddInt32(&a.v, int32(move)))
}

func (a *AInt) Set(value int) {
	atomic.StoreInt32(&a.v, int32(value))
}

func (a *AInt) Get() int {
	return int(atomic.LoadInt32(&a.v))
}

type AInt64 struct {
	v int64
}

func NewAtomicInt64(initv int64) *AInt64 {
	return &AInt64{v: initv}
}

func (a *AInt64) Add(move int64) int64 {
	return atomic.AddInt64(&a.v, move)
}

func (a *AInt64) Set(value int64) {
	atomic.StoreInt64(&a.v, value)
}

func (a *AInt64) Get() int64 {
	return atomic.LoadInt64(&a.v)
}

type AFloat struct {
	v uint64
}

//https://github.com/uber-go/atomic/blob/master/atomic.go
func NewAtomicFloat(initv float64) *AFloat {
	f := &AFloat{v: 0}
	f.Set(initv)
	return f
}

func (f *AFloat) Add(move float64) float64 {
	for {
		oldv := f.Get()
		newv := oldv + move
		if atomic.CompareAndSwapUint64(&f.v, math.Float64bits(oldv), math.Float64bits(newv)) {
			return newv
		}
	}
}

func (a *AFloat) Set(value float64) {
	atomic.StoreUint64(&a.v, math.Float64bits(value))
}

func (a *AFloat) Get() float64 {
	return math.Float64frombits(atomic.LoadUint64(&a.v))
}

type ABool struct {
	b    uint32
	strs [2]string
}

func NewAtomicBool(initv bool, truefalse [2]string) *ABool {
	if initv {
		return &ABool{b: 1, strs: truefalse}
	} else {
		return &ABool{b: 0, strs: truefalse}
	}
}

func (a *ABool) Set(value bool) {
	var i uint32 = 0
	if value {
		i = 1
	}
	atomic.StoreUint32(&(a.b), uint32(i))
}

func (a *ABool) Get() bool {
	if atomic.LoadUint32(&(a.b)) != 0 {
		return true
	}
	return false
}

func (a *ABool) GetStr() string {
	if atomic.LoadUint32(&(a.b)) != 0 {
		return a.strs[0]
	}
	return a.strs[1]
}

// ===

/*type AString struct {
	p unsafe.Pointer
}

func NewAtomicString(initv string) *AString {
	s := &AString{}
	s.Set(initv)
	return s
}

func (a *AString) Set(value string) {
	atomic.StorePointer(&a.p, unsafe.Pointer(&value))
}

func (a *AString) Get() string {
	return "" + *(*string)(atomic.LoadPointer(&a.p))
}*/

type AString struct {
	v atomic.Value
}

func NewAtomicString(initv string) *AString {
	s := &AString{}
	s.Set(initv)
	return s
}

func (a *AString) Set(value string) {
	a.v.Store(value)
}

func (a *AString) Get() string {
	return a.v.Load().(string)
}

// ===

type FuncArr struct {
	arr map[*func()]bool
}

func NewFuncArr() *FuncArr {
	return &FuncArr{arr: make(map[*func()]bool)}
}

func (fa *FuncArr) Add(f *func()) {
	fa.arr[f] = true
}

func (fa *FuncArr) Remove(f *func()) {
	delete(fa.arr, f)
}

func (fa *FuncArr) ExecAll() {
	for fk := range fa.arr {
		(*fk)()
	}
}
