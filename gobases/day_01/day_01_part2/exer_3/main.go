package main

import "fmt"

func main() {

	fmt.Println(returnMonth(122))
}

func returnMonth(month_number int) string {
	months := map[int]string{
		1:  "January",
		2:  "February",
		3:  "March",
		4:  "April",
		5:  "May",
		6:  "June",
		7:  "July",
		8:  "August",
		9:  "September",
		10: "October",
		11: "November",
		12: "December",
	}

	if month_number > 12 || month_number < 1 {
		return "Invalid Number"
	}

	return months[month_number]
}
