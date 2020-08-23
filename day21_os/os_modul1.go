package main

import (
	"fmt"
	"os"
)

// os 模块常用方法

func ostest1() {
	// 预定义变量，保存命令行参数
	fmt.Println(os.Args)
	// 获取主机名称
	fmt.Println(os.Hostname())
	fmt.Println(os.Getpid())
	// 获取全部环境变量
	env1 := os.Environ()
	for k, v := range env1 {
		fmt.Println(k, v) // 循环打印
	}
	//os.Exit(2) // 退出程序
	fmt.Println(os.Getenv("PATH")) // 获取一条环境变量

	// 获取当前目录
	//fmt.Println("==========")
	dir, err := os.Getwd()
	fmt.Println(dir, err)

	// 创建hello 目录
	//err=os.Mkdir(dir+"/hello",0755)
	// fmt.Println(err) // 输出值为<nil>
	//fmt.Println(os.Getwd())

	// 创建递归目录
	//err =os.MkdirAll(dir+"/h1/h2",0755)
	//fmt.Println(err)

	// 删除目录,相当于 rm -rf 删除某个目录
	//err=os.RemoveAll(dir+"/h1")
	//fmt.Println(err)

	// 删除目录，目录不为空不能直接删除
	//err=os.Remove(dir+"/hi")
	//fmt.Println(err)

	//appname :=os.Mkdir(dir+"/hi",0755)
	//println(appname)

	// 环境变量检查  Getenv os.Getenv检索环境变量并返回值，如果变量是不存在的，这将是空的
	HOME := os.Getenv("HOME")
	if err != nil {
		fmt.Println(HOME)
	} else {
		fmt.Println("HOME 目录不存在")
	}
	// 设置环境变量
	fmt.Println("设置环境变量")
	err1 := os.Setenv("HOME1", "/Users/hang.ling/")
	if err1 != nil {
		fmt.Println(err1, "===")
	}
	//HOME1 :=os.Getenv("HOME1")
	//fmt.Println(HOME1)

}

func main() {
	fmt.Println("os 模块的使用")
	ostest1()

}
