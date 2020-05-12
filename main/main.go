package main

import (
	"fmt"
	"sync"
)

//по очереди
//WaitGroup

func sum(array []int, flag bool) (sum int) {
	if flag { // sum even numbers
		for _, value := range array {
			if value%2 == 0 {
				sum += value
			}
		}
	} else { // sum odd index number
		for _, value := range array {
			if value%2 != 0 {
				sum += value
			}
		}
	}
	fmt.Println(sum)
	return
}

func sumWG(array []int, flag bool) (sum int) {
	if flag { // sum even numbers
		for _, value := range array {
			if value%2 == 0 {
				sum += value
			}
		}
	} else { // sum odd index number
		for _, value := range array {
			if value%2 != 0 {
				sum += value
			}
		}
	}
	fmt.Println(sum)
	wg.Done()
	return
}

var wg = sync.WaitGroup{}

func main() {
	array := []int{0, 1, 2, 44, 55}

	/*go sum(array, true)
	go sum(array, false)*/
	wg.Add(2)
	go sumWG(array, false)
	/*wg.Wait()
	wg.Add(1)*/
	go sumWG(array, true)
	wg.Wait()

	//time.Sleep(1000 * time.Millisecond)
}

/*counter++
fmt.Println("main ", counter)
if !flag {
	flag = true
	f := Funcer{
		func1: main,
		func2: main,
	}
	fmt.Println(f.ReturnMethods())
}*/
/*
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
}*/
