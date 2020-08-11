package main

// map 排序

import (
	"fmt"
	"sort"
)

// map 元素 排序

func map1() {
	var (
		barVal = map[string]int{"alpha": 34, "bravo": 56, "charlie": 23,
			"delta": 87, "echo": 56, "foxtrot": 12,
			"golf": 34, "hotel": 16, "indio": 87,
			"juliet": 65, "kili": 43, "lima": 98}
	)

	for k, v := range barVal {
		fmt.Printf("key:%v,Value:%v", k, v)
	}

	keys := make([]string, len(barVal))
	i := 0
	for k, _ := range barVal {
		keys[i] = k
		i++
	}
	sort.Strings(keys)

	fmt.Println()
	fmt.Println("sorted:")

	for _, k := range keys {
		fmt.Printf("Key: %v, Value: %v /", k, barVal[k])
	}

	// fmt.Println()

	// 输出
	// key:charlie,Value:23key:delta,Value:87key:indio,Value:87key:juliet,Value:65key:kili,Value:43key:lima,Value:98key:bravo,Value:56key:echo,Value:56key:foxtrot,Value:12key:golf,Value:34key:hotel,Value:16key:alpha,Value:34
	// sorted:
	// Key: alpha, Value: 34 /Key: bravo, Value: 56 /Key: charlie, Value: 23 /Key: delta, Value: 87 /Key: echo, Value: 56 /Key: foxtrot, Value: 12 /Key: golf, Value: 34 /Key: hotel, Value: 16 /Key: indio, Value: 87 /Key: juliet, Value: 65 /Key: kili, Value: 43 /Key: lima, Value: 98 /

	// map  中的键值对调

	inMap := make(map[int]string, len(barVal))
	for k, v := range barVal {
		inMap[v] = k
	}

	fmt.Println("inverted\n:")

	for k, v := range inMap {
		fmt.Printf("key is %v ,value:%v / ", k, v)
	}

}

func main() {
	map1()
}
