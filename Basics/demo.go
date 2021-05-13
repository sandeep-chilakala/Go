package main

import (
	"fmt"
	"math"
	"time"
)

func sumOfNumbers(nums ...int) int {

	sum := 0
	for _, num := range nums {
		sum += num
	}

	return sum
}

func squareRoot(a float64) (b float64) {
	b = math.Sqrt(a)
	return b
}

func runInLoop(a int) {
	for i := 1; i <= a; i++ {
		fmt.Println("Loop ", i)
	}
}

var daysOfWeek = [...]string{
	"Sunday",
	"Monday",
	"Tuesday",
	"Wednesday",
	"Thursday",
	"Friday",
	"Saturday",
}
var daysOfWeekMap = map[string]int{
	"Sunday":    0,
	"Monday":    1,
	"Tuesday":   2,
	"Wednesday": 3,
	"Thursday":  4,
	"Friday":    5,
	"Saturday":  6,
}

func parseWeekday(v string) (time.Weekday, error) {
	for i := range daysOfWeek {
		if daysOfWeek[i] == v {
			return time.Weekday(i), nil
		}
	}

	return time.Sunday, fmt.Errorf("invalid weekday '%s'", v)
}

func findWhenIsDay(day string) (b string) {
	dayInt := daysOfWeekMap[day]
	todayStr := time.Now().Weekday().String()
	//fmt.Println("todayStr ", todayStr, dayInt)
	todayInt := daysOfWeekMap[todayStr]
	//fmt.Println("dayInt ", dayInt, todayInt)
	//diff := dayInt - todayInt
	switch dayInt - todayInt {
	case 0:
		fmt.Println("Today")
	case 1:
		fmt.Println("Tomorrow")
	case 2:
		fmt.Println("2 days from today")
	default:
		fmt.Println("Not Near")
	}
	return b
}

func printStart(num int) {
	for i := 0; i < num; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}

func isPalandromeNum(num int) bool {
	temp := 0
	for i := num; i > 0; i = i / 10 {
		modNum := i % 10
		temp = (temp * 10) + modNum
	}
	return temp == num
}

func isPalandromeStr(str string) bool {
	temp := ""
	for i := len(str) - 1; i >= 0; i-- {
		temp = temp + string(str[i])
	}
	return temp == str
}

func main() {
	fmt.Println(isPalandromeStr("tet"))
	//fmt.Println(len("test"))

	/*
		printStart(2)
		var num1 = map[string]string{"name": "sandeep"}
		//num2 := 2
		runInLoop(2)
		parseWeekday("Sunday")
		fmt.Println("Sum = ", sumOfNumbers(1, 2, 3, 4, 5))
		fmt.Println("Square Root of ", squareRoot(float64(10)))
		fmt.Println("Today date is ", time.Now().Format("2-January-2016"), " Day is ", time.Now().Weekday()+1, " is saturday ? ", time.Saturday)
		fmt.Println(findWhenIsDay("Sunday"))
		//num2 := *&num1
		fmt.Println(num1)
	*/
}
