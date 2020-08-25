package main

import "fmt"

func Getnum() (int, int) {
	return 10, 20
}

func main() {
	var (
		a1 int
		b1 string
		c1 float64
		d1 func() bool
		e1 struct {
			x1 int
			y1 string
		}
	)
	fmt.Printf("a1 is %d, b1 is %s ,c1 is %f, d1 is %t ,e1 is %t\n", a1, b1, c1, d1, e1)
	// 输出
	//a1 is 0, b1 is  ,c1 is 0.000000, d1 is %!t(func() bool=<nil>) ,e1 is {%!t(int=0) %!t(string=)}

	var a int = 10
	var b int = 20
	b, a = a, b
	fmt.Println(a, b)
	// 20 10

	//a2,_:=Getnum()
	//_,b2:=Getnum()

	var flag bool
	fmt.Println(flag)

	//var name string= "go"
	var temp string
	temp = `
      x :=10
      y :=20
      z :=30
 
	fmt.Println(x," ",y," ",z)
	x,y,z = y,z,x
	fmt.Println(x," ",y," ",z)`
	fmt.Println(temp)

	//fmt.Println(name)

}
