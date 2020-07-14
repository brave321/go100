package main

import (
	"fmt"
	"time"
	"bufio"
	"os"
)

// go 常用模块的使用

// time 模块的使用

func timetest() {
	t := time.Now()
	fmt.Println(t) //2020-07-14 08:26:33.956618 +0800 CST m=+0.000062895
	// fmt.Println(t.Date())    //2020 July 14
	// fmt.Println(t.Year())    // 2020
	// fmt.Println(t.Month())   //July
	// fmt.Println(t.Day())     //14
	// fmt.Println(t.Hour())    //8
	// fmt.Println(t.Minute())  // 26
	// fmt.Println(t.Second())  // 33
	// fmt.Println(t.Weekday()) //Tuesday
	// fmt.Println(t.YearDay()) //196
	// fmt.Println(t.Zone())    //CST 28800

	// 格式化字符串与Time类型转换
	// fmt.Println(t, reflect.TypeOf(t))

}

// 文件读写操作

fun iotest(){
	file, _:=os.Open("test.txt")
	defer file.Close()

	// / 字节切片缓存 存放每次读取的字节
	buf := make([]byte, 1024)

	for {
        // 返回本次读取的字节数
        count, err := file.Read(buf)

        // 检测是否到了文件末尾
        if err == io.EOF {
            break;
        }

        // 取出本次读取的数据
        currBytes := buf[:count]

        // 将读取到的数据 追加到字节切片中
        bytes = append(bytes, currBytes...)
	}
	// 将字节切片转为字符串 最后打印出来文件内容
    fmt.Println(string(bytes))
	
}

func main() {
	timetest()
	iotest()
}
