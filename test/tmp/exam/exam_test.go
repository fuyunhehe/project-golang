package exam

import (
	"testing"
	"fmt"
)

func recoverFunc(f func()) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover", err)
		}
	}()
	f()
}

func TestDeferCall(t *testing.T) {
	recoverFunc(DeferCall)
}

func TestPaseStudent(t *testing.T) {
	PaseStudent()
}

func TestGoCall(t *testing.T) {
	GoCall()
}

func TestTeacherShow(t *testing.T) {
	TeacherShow()
}

func TestSelectCall(t *testing.T) {
	for i := 0; i < 10; i++ {
		recoverFunc(SelectCall)
	}
}

func TestCalcCall(t *testing.T) {
	CalcCall()
}

func TestSliceCall(t *testing.T) {
	SliceCall()
}

func TestChanCall(t *testing.T) {
	ChanCall()
}

func TestPanicCall(t *testing.T) {
	PanicCall()
}

func TestObjCall(t *testing.T) {
	ObjCall()
}