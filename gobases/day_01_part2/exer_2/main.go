package main

import "fmt"

func main() {

	client_1 := UserInfo{23, "Raphael Borba", true, 2, 100001}

	fmt.Println(able_to_get_loan(client_1))

}

type UserInfo struct {
	age           int
	name          string
	isWorking     bool
	timeInSameJob float32
	salary        int64
}

func able_to_get_loan(client_info UserInfo) string {

	if client_info.age <= 22 || !client_info.isWorking || client_info.timeInSameJob < 1 {
		return "Not Possible to get a loan"
	}

	if client_info.salary > 100000 {
		return "Congrats! You can have this loan with no tax!"
	}
	return "Congrats! You can have this loan with tax!"
}
