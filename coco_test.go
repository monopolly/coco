package coco

//testing
import (
	"fmt"
	"os/exec"
	"strings"
	"testing"
	//testing
	//go test -bench=.
	//go test --timeout 9999999999999s
)

func TestMainA(u *testing.T) {
	__(u)

	p := New(100000)
	p.Lowercase(true)

	for range 10 {
		p.AddString("1")
		p.Flush()
		fmt.Println(p.Count())
	}

	return

	p.Add([]byte("NICE"))
	p.Add([]byte("NICE1"))
	p.Add([]byte("NICE2"))
	p.Add([]byte("NICE3"))
	p.Flush()

	if !p.Has([]byte("NICE")) {
		panic("fail")
	}
	if !p.Has([]byte("NICE3")) {
		panic("fail")
	}

}

func Benchmark1(u *testing.B) {
	u.ReportAllocs()

	for u.Loop() {

	}
}

func Benchmark2(u *testing.B) {
	u.RunParallel(func(pb *testing.PB) {
		for pb.Next() {

		}
	})
}

func __(u *testing.T) {
	fmt.Printf("\033[1;32m%s\033[0m\n", strings.ReplaceAll(u.Name(), "Test", ""))
}

func cmd(name string, v ...string) {
	c := exec.Command(name, v...)
	r, err := c.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(r))
}
