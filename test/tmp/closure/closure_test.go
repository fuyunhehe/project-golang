package closure

import "testing"

func TestCreateFunc(t *testing.T) {
	fList := CreateFunc()
	//fmt.Println(fList)
	for _, f := range fList{
		f()
	}
}

func BenchmarkPerformance(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i:=0; i<b.N; i++ {
		Performance()
	}

}

func BenchmarkPerformance2(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i:=0; i<b.N; i++ {
		Performance2()
	}
}