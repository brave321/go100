package main

import "fmt"

// if 条件语句

// if 是用于测试某个条件（布尔型或逻辑型）的语句，如果该条件成立，
//则会执行 if 后由大括号括起来的代码块，否则就忽略该代码块继续执行后续的代码。

// 语法句式1

// if condition {
//     // do something
// }

// 语法句式2

// if condition {
//     // do something
// } else {
//     // do something
// }

// 语法句式3

// if condition1 {
//     // do something
// } else if condition2 {
//     // do something else
// }else {
//     // catch-all or default
// }

func contrtest1() {
	bool1 := false
	if bool1 {
		fmt.Println("the value is true")
	} else {
		fmt.Println("the value is false")
	}
	// output  the value is false
}

//
func contrtest2() {
	var score = 93
	if score <= 0 {
		fmt.Println("illegele score not fu shuo")
	} else if score > 60 && score < 75 {
		fmt.Println("C")
	} else if score >= 75 && score < 90 {
		fmt.Println("B")
	} else if score >= 90 && score <= 100 {
		fmt.Println("A")
	} else {
		fmt.Println("illegele not >100")
	}
	// 输出 A

}

// switch 结构

// switch var1 {
// case val1:
// 	...
// case val2:
// 	...
// default:
// 	...
// }

// 每一个case 分支都是唯一的，从上至下，逐一测试，直到匹配为止
// 一旦成功的匹配的某个分之 ，在执行完相应的代码就会退出整个swictch 块
//可选的 default 分支可以出现在任何顺序，但最好将它放在最后。它的作用类似与 if-else 语句中的 else，表示不符合任何已给出条件时，执行相关语句。

// 如果在执行完每个分支的代码后，还希望继续执行后续分支的代码，可以使用 fallthrough 关键字来达到目的
func switchtest1() {
	var num1 int = 91
	switch num1 {
	case 98, 99:
		fmt.Println("It's equal to 98")
	case 100:
		fmt.Println("It's equal to 100")
	default:
		fmt.Println("It's not equal to 91")
		// 输出
		//It's not equal to 91

	}
}

func switchtest2() {
	k := 6
	switch k {
	case 4:
		fmt.Println("was <= 4")
		fallthrough
	case 5:
		fmt.Println("was <= 5")
		fallthrough
	case 6:
		fmt.Println("was <= 6") // 使用 fallthrough，执行完 case 6,后续的代码也会执行
		fallthrough
	case 7:
		fmt.Println("was <= 7")
		fallthrough
	case 8:
		fmt.Println("was <= 8")
		fallthrough
	default:
		fmt.Println("default case")
	}
	// 输出
	// 	was <= 6
	// was <= 7
	// was <= 8
	// default case
}

func switchtest3(num int) {
	switch num {
	case 1:
		fmt.Println("Jan,Feb,Mar")
	case 2:
		fmt.Println("Apr,May,Jun")
	case 3:
		fmt.Println("Jul,Aug,Sept")
	case 4:
		fmt.Println("Oct,Nov,Dec")
	default:
		fmt.Println("run error")
	}
}

func main() {
	// fmt.Println("hello")
	// contrtest1()
	// contrtest2()
	// switchtest1()
	// switchtest2()
	// num2 := "JUL"
	// fmt.Println(switchtest3(6))
	// num := 3
	// switchtest3(3)
	// 输出 Jul,Aug,Sept

}
