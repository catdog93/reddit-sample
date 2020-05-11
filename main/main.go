package main

import "fmt"

/*
funcs & closures
*/

var (
	flag    bool
	counter int
)

func main() {
	counter++
	fmt.Println("main ", counter)
	if !flag {
		flag = true
		f := Funcer{
			func1: main,
			func2: main,
		}
		fmt.Println(f.ReturnMethods())
	}
}

type funcer interface {
	ReturnMethods() []*func()
}

type Funcer struct {
	func1 func()
	func2 func()
}

//ReturnMethods
func (f *Funcer) ReturnMethods() (funcSlice []*func()) {
	if f != nil {
		funcSlice = append(funcSlice, &f.func1)
		funcSlice = append(funcSlice, &f.func2)
	}
	return
}
