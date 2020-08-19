package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/user"
)

func main() {
	fmt.Println("hello world")
	u, err := user.Current()
	// fmt.Println(u, err)
	fmt.Println("错误信息err: ", err)
	fmt.Println("用户ID:Uid: ", u.Uid)
	fmt.Println("用户名:Username: ", u.Username)
	fmt.Println("用户组ID:Gid: ", u.Gid)
	fmt.Println("用户组名称:Name: ", u.Name)
	fmt.Println("用户家目录:HomeDir: ", u.HomeDir)
	// 使用 os 模块对文件进行操作
	// errMsg := os.Mkdir("./TEST", os.ModeDir)
	// if errMsg != nil {
	// 	fmt.Println("文件存在", errMsg)
	// 	fmt.Println("删除文件", os.RemoveAll("./TEST"))
	// 	return
	// } else {
	// 	fmt.Println("创建目录成功")
	// }

	// os.MkdirAll()可以用于父目录不存在,或者要创建的目录已经存在的情况;父目录不存在会同时创建;
	// errMsg2 := os.MkdirAll("./A/B", os.ModeDir)
	// if errMsg2 != nil {
	// 	fmt.Println("创建目录失败", errMsg2)
	// 	return
	// } else {
	// 	fmt.Println("创建目录成功")
	// }

	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		fmt.Println(f.Name())
	}

}
