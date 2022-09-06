package main

import (
	"fmt"
	"time"
)

func dateFormat(now time.Time) {

	// date format
	fmt.Printf("%s \n", now.Format("2006"))
	fmt.Printf("%s \n", now.Format("01"))
	fmt.Printf("%s \n", now.Format("02"))
	fmt.Printf("%s \n", now.Format("15"))
	fmt.Printf("%s \n", now.Format("04"))
	fmt.Printf("%s \n", now.Format("05"))
	fmt.Printf("%s \n", now.Format("2006-01-02 15:04:05"))
	// get year month week day
	fmt.Printf("year %T %v \n", now.Year(), now.Year())
	fmt.Printf("month %T %v \n", now.Month(), now.Month())
	fmt.Printf("weekday %T %v \n", now.Weekday(), now.Weekday())
	fmt.Printf("day %T %v \n", now.Day(), now.Day())

	wkd := uint8(now.Weekday())
	mn := uint8(now.Month())
	// wkd := uint8(now.Weekday())
	fmt.Println(wkd)
	fmt.Println(mn)
}

func dataCalc(now time.Time) {
	d := now.Add(24 * time.Hour)
	ds := d.Format("2006-01-02 15:04:05")
	fmt.Println(ds)
	// year
	dd := now.AddDate(0, 0, -7)
	fmt.Println(now.GoString())
	fmt.Println(dd)
}

func getFirstAndLastDayOfMonth(now time.Time) (first, last time.Time) {
	year, month, _ := now.Date()
	locl := now.Location()
	first = time.Date(year, month, 1, 0, 0, 0, 0, locl)
	last = first.AddDate(0, 1, -1)
	return
}

func main() {
	now := time.Now()
	dateFormat(now)
	dataCalc(now)
	fst, lst := getFirstAndLastDayOfMonth(now)
	fmt.Printf("%s %s \n", fst, lst)
}
