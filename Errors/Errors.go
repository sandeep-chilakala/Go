package main

import (
	"errors"
	"fmt"
	"log"
)

type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f1(arg int) (int, error) {
	if arg == 42 {

		return -1, errors.New("can't work with 42")

	}

	return arg + 3, nil
}

func f2(arg int) (int, error) {
	if arg == 42 {

		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

//my code start
type customError struct {
	start  int
	end    int
	errStr string
}

func (e *customError) Error() string {
	return fmt.Sprintf("Your start number %d should be less than end number %d - %s", e.start, e.end, e.errStr)
}

func sumOfNumbers(start, end int) (int, error) {
	if start > end {
		return -1, errors.New("Start should be less than end")
	} else if end == 0 {
		return -1, &customError{start, end, "custom Error"}
	}
	sum := 0
	for i := start; i < end; i++ {
		sum += i
	}
	return sum, nil
}

//my code end

func main() {

	//my code start
	log.Println(sumOfNumbers(10, 8))
	//get each value from the custom error struct
	_, er := sumOfNumbers(-1, 0)
	if ce, ok := er.(*customError); ok { // this to instantate the error object to get the values in it
		log.Println(ce.start)
		log.Println(ce.end)
		log.Println(ce.errStr)
	}
	//my code end
	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}
	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}

	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}
