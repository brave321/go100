package main

import "fmt"

//

// 判断value 值是否存在  Tuesday 和 Hollyday。
func map1() {
	var map1 map[string]string
	map1 = map[string]string{"1": "Monday", "2": "Tuesday", "3": "Wednesday", "4": "Thursday", "5": "Friday", "6": "Saturday", "7": "Sunday"}
	// fmt.Println(map1)

	for k, v := range map1 {
		if _, ok := map1[k]; ok {
			fmt.Println(k, v)
		}

		if map1[k] == "Tuesday" {
			fmt.Println("the value contains Tuesday")
		}

		if map1[k] == "Hollyday" {
			fmt.Println("the value contains Hollyday ")
		} else {
			fmt.Println("the value  not contains Hollyday ")
		}

		// i, ok := map1["2"]

		// i, ok = map1["2"]
		// fmt.Println(i, ok)

		// 输出
		// 		6 Saturday
		// the value  not contains Hollyday
		// 7 Sunday
		// the value  not contains Hollyday
		// 1 Monday
		// the value  not contains Hollyday
		// 2 Tuesday
		// the value contains Tuesday
		// the value  not contains Hollyday
		// 3 Wednesday
		// the value  not contains Hollyday
		// 4 Thursday
		// the value  not contains Hollyday
		// 5 Friday
		// the value  not contains Hollyday
		//

	}
}

func main() {
	fmt.Println("hello ")
	map1()
}
