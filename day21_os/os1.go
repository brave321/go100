package main

import (
	"fmt"
	"os"
)

var EnvCfg=Cfg{
	AppName: "pod-event-router",
	Port: "8080",
}

type Cfg struct{
	AppName string
	Port  string
}

func init()  {
	if appName,ok :=os.LookupEnv(key: "APP_NAME");ok{
		EnvCfg.appName = appName
		// fmt.Println("这是 EnvCfg.appName 的值",EnvCfg.appName)
	}
}




func main() {
	// 获取当前路径
	dir, _ := os.Getwd()
	fmt.Println("当前路径", dir)

	// 获取系统环境变量的值
	path1 := os.Getenv("GOPATH")
	fmt.Println("环境变量GOPATH的值是", path1)

	// os.chidir() 改变路径

	beforedir, _ := os.Getwd()
	fmt.Println("起始路径", beforedir)
	//Getenv()、LookupEnv()传入参数与环境变量名严格一致（包括大小写）
	//LookupEnv 可以区分空环境变量或者缺失环境变量，更加清新的显示环境变量的真实情况

	//如果有环境变量v，内容为空
	fmt.Println(os.Getenv("v"), "hello=========") //返回`空`
	fmt.Println(os.LookupEnv("v"))                //俩返回值，会返回`空`以及`true`
	//如果不存在环境变量n
	// fmt.Println(os.Getenv("n"))    //返回`空`
	// fmt.Println(os.LookupEnv("n")) //俩返回值，会返回`空`以及`false`

	fmt.Println("#######################")
	init()


}
