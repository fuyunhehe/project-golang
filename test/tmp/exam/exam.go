package exam

import (
	"fmt"
	"runtime"
	"sync"
	"math/rand"
	"time"
)

//no.1
//打印后
//打印中
//打印前
//panic: 触发异常
func DeferCall() {
	defer func() { fmt.Println("打印前") }()
	defer func() { fmt.Println("打印中") }()
	defer func() { fmt.Println("打印后") }()

	panic("触发异常")
}

//no.2
type student struct {
	Name string
	Age  int
}

//m中保留的student为同一个
func PaseStudent() {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	for _, stu := range stus {
		m[stu.Name] = &stu
	}
	fmt.Println(m)
}

//no.3
//任意值-重复
//i: 10
//...
//i: 10
//乱序-不重复
//i:0
//...
//i:10
func GoCall() {
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	wg.Add(20)
	for i := 0; i < 10; i++ {
		go func() {
			r := rand.Intn(5)
			time.Sleep(time.Second * time.Duration(r))
			fmt.Println("1i: ", i)
			wg.Done()
		}()
		time.Sleep(time.Second)
	}
	for i := 0; i < 10; i++ {
		go func(i int) {
			r := rand.Intn(5)
			time.Sleep(time.Second * time.Duration(r))
			fmt.Println("2i: ", i)
			wg.Done()
		}(i)
		time.Sleep(time.Second)
	}
	wg.Wait()
}

//no.4
type people struct{}

func (p *people) showA() {
	fmt.Println("showA")
	p.showB()
}
func (p *people) showB() {
	fmt.Println("showB")
}

type teacher struct {
	people
}

func (t *teacher) showB() {
	fmt.Println("teacher showB")
}

//showA
//showB
func TeacherShow() {
	t := teacher{}
	t.showA()
}

//no.5
//会panic，因为select是随机选取已准备就绪的channel
func SelectCall() {
	runtime.GOMAXPROCS(1)
	int_chan := make(chan int, 1)
	string_chan := make(chan string, 1)
	int_chan <- 1
	string_chan <- "hello"
	select {
	case value := <-int_chan:
		fmt.Println(value)
	case value := <-string_chan:
		panic(value)
	}
}

//no.6
//10 1 2 3
//20 0 2 2
//2 0 2 2
//1 1 3 4
func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func CalcCall() {
	a := 1
	b := 2
	defer calc("1", a, calc("10", a, b))
	a = 0
	defer calc("2", a, calc("20", a, b))
	b = 1
}

//no.7
//0 0 0 0 0 1 2 3
func SliceCall() {
	s := make([]int, 5)
	s = append(s, 1, 2, 3)
	fmt.Println(s)
}

//no.8
//读的时候也需要加读写锁
type userAges struct {
	ages map[string]int
	sync.Mutex
}

func (ua *userAges) Add(name string, age int) {
	ua.Lock()
	defer ua.Unlock()
	ua.ages[name] = age
}

func (ua *userAges) Get(name string) int {
	if age, ok := ua.ages[name]; ok {
		return age
	}
	return -1
}

//no.9
type threadSafeSet struct {
	s rune
	sync.RWMutex
}

func (set *threadSafeSet) Iter() <-chan interface{} {
	ch := make(chan interface{})
	go func() {
		set.RLock()

		for _, elem := range set.s {
			ch <- elem
		}

		close(ch)
		set.RUnlock()
	}()
	return ch
}

func ChanCall() {
	tss := threadSafeSet{
		s:       make([]interface{}, 0, 100),
		RWMutex: sync.RWMutex{},
	}
	for i := 0; i < 100; i++ {
		tss.s = append(tss.s, i*2)
	}
	c := tss.Iter()
	for s := range c {
		fmt.Println(s)
	}
}

//no.10
type iPeople interface {
	Speak(string) string
}

type stduent struct{}

func (stu *stduent) Speak(think string) (talk string) {
	if think == "bitch" {
		talk = "You are a good boy"
	} else {
		talk = "hi"
	}
	return
}

func PanicCall() {
	// todo panic,这里 student必须使用引用,或者student的Speak方法使用非指针声明
	var peo iPeople = &stduent{}
	think := "bitch"
	fmt.Println(peo.Speak(think))
}

//no.11
// BBBBBBB 指针值，指向nil
type iPeople2 interface {
	Show()
}

type student2 struct{}

func (stu *student2) Show() {

}

func live() iPeople2 {
	var stu *student2
	return stu
}

func ObjCall() {
	l := live()
	if l == nil {
		fmt.Println("AAAAAAA", l)
	} else {
		fmt.Println("BBBBBBB", l, nil)
	}
}
