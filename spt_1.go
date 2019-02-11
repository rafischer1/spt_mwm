// not quitew done yet...almost there with the mapping to a struct and migrating back to the array of strings...

package main

import (
	"fmt"
	"strconv"
)

// Date object to map array
type Date struct {
	Month int
	Day   int
	Year  int
}

func main() {
	// input array
	dateArr := [7]string{"Oct 7, 2009", "Nov 10, 2009", "Jan 10, 2009", "Oct 22, 2009", "Dec 1, 2019", "Sep 20, 2010", "Aug 2, 1912"}

	// months array to pass down for conversion
	months := [12]string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}

	// call the function on the original array
	answer := dateSorter(dateArr, months)
	fmt.Println("Answer", answer)
}

func dateSorter(arr [7]string, m [12]string) [7]string {
	var result [7]string
	d := []Date{}
	d = mapArrToObj(arr, m)
	fmt.Println(d)
	// for _, arr := range d {
	// 	result = append(result, monthIntToStr(d[0].Month, m)+"2,"+"2007")
	// }
	return result
}

func mapArrToObj(arr [7]string, m [12]string) []Date {
	dates := []Date{}
	for i := 0; i < 7; i++ {
		fmt.Println(arr[i][0:3], arr[i][3:6], arr[i][7:])
		x, _ := strconv.Atoi(arr[i][3:6])
		y, _ := strconv.Atoi(arr[i][7:len(arr[i])])
		n := Date{Month: monthStrToInt(arr[i][0:3], m), Day: x, Year: y}
		dates = append(dates, n)
	}
	return dates
}

func monthStrToInt(s string, m [12]string) int {
	var num int
	for i := 0; i < len(m); i++ {
		if s == m[i] {
			num = i
		}
	}
	return num
}

func monthIntToStr(n int, m [12]string) string {
	var str string
	for i := 0; i < len(m); i++ {
		if n == i {
			str = m[i]
		}
	}
	return str
}
