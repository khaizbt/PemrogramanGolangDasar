package main

import (
	"fmt"
	time2 "time"
)

func main() {
	var time = time2.Now()
	fmt.Printf("time 1 %v \n", time)

	var time1 = time2.Date(1997, 12, 17, 9, 9, 0, 0, time2.Local)
	fmt.Printf("time2 %v \n", time1)

	//Karena time2 time merupakan struct, maka memiliki beberapa method sbb:
	fmt.Println("year: ", time.Year(), "month : ", time.Month(), "jam : ", time.Hour(), ":", time.Minute())
	//semua method bisa di cek disini https://dasarpemrogramangolang.novalagung.com/38-time.html

	//Parsing dari string ke time
	var layoutFormat, value string
	var date time2.Time
	//layoutFormat = "2006-01-01 15:04:05"
	layoutFormat = "2006-January-2 15:00:05"
	value = "2004-May-04 06:00:00"
	date, _ = time2.Parse(layoutFormat, value)
	fmt.Println(value, "\t->", date.String())

	//Predefined Layout Format Untuk Keperluan Parsing Time
	var date2, _ = time2.Parse(time2.RFC822, "02 Sep 15 08:00 WIB")
	fmt.Println(date2.String())

	//Format dato time ke string
	var date3, _ = time2.Parse(time2.RFC822, "02 Sep 15 08:00 WIB")

	var dateS1 = date3.Format("Monday 02, January 2006 15:04 MST")
	fmt.Println("dateS1", dateS1)

	var dateS2 = date3.Format(time2.RFC3339)
	fmt.Println("dateS2", dateS2)

	//Handle Error parsing time

	var dateErr, err = time2.Parse("06 Jan 15", "02 Sep 15 08:00 WIB")

	if err != nil {
		fmt.Println("error", err.Error())
		return
	}

	fmt.Println(dateErr)
}
