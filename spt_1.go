package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// Date object to map array
type Date struct {
	Month int
	Day   int
	Year  int
}

// By is the type of a "less" function that defines the ordering of its Date arguments.
type By func(d1, d2 *Date) bool

// Sort is a method on the function type, By, that sorts the argument slice according to the function.
func (by By) Sort(dates []Date) {
	ds := &dateSorter{
		dates: dates,
		by:    by,
	}
	sort.Sort(ds)
}

// Len is part of sort.Interface.
func (d *dateSorter) Len() int {
	return len(d.dates)
}

// Swap is part of sort.Interface.
func (d *dateSorter) Swap(i, j int) {
	d.dates[i], d.dates[j] = d.dates[j], d.dates[i]
}

// Less is part of sort.Interface. It is implemented by calling the "by" closure in the sorter.
func (d *dateSorter) Less(i, j int) bool {
	return d.by(&d.dates[i], &d.dates[j])
}

// dateSorter joins a By function and a slice of dates to be sorted.
type dateSorter struct {
	dates []Date
	by    func(d1, d2 *Date) bool // Closure used in the Less method.
}

func main() {
	// input array
	dateArr := [7]string{"Oct 7, 2009", "Nov 10, 2009", "Jan 10, 2009", "Oct 22, 2009", "Dec 1, 2019", "Sep 20, 2010", "Aug 2, 1912"}

	// months array to pass down for conversion
	months := [12]string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}

	// call the function on the original array
	answer := dateSorterFunc(dateArr, months)
	fmt.Println("Answer", answer)
}

func dateSorterFunc(arr [7]string, m [12]string) []string {
	dates := []Date{}
	dates = mapArrToObj(arr, m)

	year := func(d1, d2 *Date) bool {
		return d1.Year > d2.Year
	}
	month := func(d1, d2 *Date) bool {
		return d1.Month > d2.Month
	}
	day := func(d1, d2 *Date) bool {
		var t bool
		if d1.Month == d2.Month {
			t = d1.Day > d2.Day
		}
		return t
	}

	// sort them dates!
	By(month).Sort(dates)
	By(year).Sort(dates)
	By(day).Sort(dates)

	var result []string

	// convert dates []Date back to []string as final answer
	for _, date := range dates {
		strDay := strconv.Itoa(date.Day)
		strYear := strconv.Itoa(date.Year)
		result = append(result, ""+monthIntToStr(date.Month, m)+" "+strDay+", "+strYear)
	}
	return result
}

func mapArrToObj(arr [7]string, m [12]string) []Date {
	dates := []Date{}
	for i := 0; i < 7; i++ {
		var d string
		var y string
		// convert month to int
		mo := monthStrToInt(arr[i][0:3], m)

		// trim out the comma for the day
		if strings.Contains(arr[i][3:6], ",") {
			d = strings.TrimSuffix(arr[i][3:6], ",")
		} else {
			d = arr[i][3:6]
		}

		//trim out extra space of the year
		if len(arr[i][7:]) == 4 {
			y = arr[i][7:]
		} else {
			y = strings.TrimPrefix(arr[i][7:], " ")
		}

		day, _ := strconv.Atoi(strings.TrimPrefix(d, " "))
		year, _ := strconv.Atoi(y)
		n := Date{Month: mo, Day: day, Year: year}
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
