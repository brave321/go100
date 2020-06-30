package main
import "fmt"

// 什么是函数  
// 函数是一块执行特定任务的代码


// 函数的分类
//1 、带有名字的函数
//2、匿名函数或者lambda 函数
// 3、方法


// 函数的声明

// func functionname(paratmetername type) returntype  {
// 	// 函数体(具体实现的功能)
// }


// 解释：声明关键词 func  函数名称 functionname  returntype 返回值类型，函数也可以如参不需要参数，并且也没有返回值

// 示例1 

func calcucateBill(price int, no int) int  {
	var sumprice = price * no 
	return sumprice
}

// 函数名称 calcucateBill  如参 price,no  返回值类型 int  // 函数主体  计算价格 

// 函数多返回值,以计算长方形的面积和周长

func multiProps(length,width float64)  (float64,float64) {
	var area = length * width
	var perimeter= (length+width) * 2
	return area,perimeter 	
}

func MultiPly3Num(a,b,c int) int {
	return a * b * c 
}


func getX2AndX3(input int) (int,int) {
	return 2 * input, 3 * input
}

//编写一个函数，接收两个整数，然后返回它们的和、积与差。编写两个版本，一个是非命名返回值，一个是命名返回值。

// 1、非命名的返回值
func calculate(num1,num2 int) (int,int,int) {
	return num1 + num2 ,num1 * num2,num1-num2
}


// 2、非命名的返回值

func calculate2(num1,num2 int) (x1,x2,x3 int) {
   x1 = num1 + num2
   x2 = num1 * num2
   x3 = num1 - num2
	// return num1 + num2 ,num1 * num2,num1-num2
	return
}

//编写一个名字为 MySqrt 的函数，计算一个 float64 类型浮点数的平方根，
//如果参数是一个负数的话将返回一个错误。编写两个版本，一个是非命名返回值，一个是命名返回值。

// 非命名返回值

func MySqrt1(x float64)  float64 {
	return x * x 
}

// 命名返回值 
func MySqrt2( y float64) (y1 float64) {
	y1 = y * y
	return
}



// 函数空白符，用来匹配一些不需要的值

func BlankTest() (int,int,float64)  {
	return 1,4,6.9
}


// 函数入参数比较大小

func MinMax(a int,b int) (min int,max int) {
	if a <b {
		min=a
		max=b
	}else{
		min=b
		max=a
	}
	return
}

// 改变外部变量 

// 传递指针给函数不但可以节省内存，而且直接赋予了直接修改外部变量的能力 
// 所以被修改的变量不再需要使用 return 返回，如下面的例子，reply 是一个指向int 变量的指针
// 通过指针，可以在函数内修改了 这个int 变量的值


func Multiply(a,b int,reply *int)  {
	*reply = a * b 
}




func main()  {
	//  fmt.Println("hello go")
	//  price,no :=100,6 
	//  sumprice :=calcucateBill(price,no)
	//  fmt.Println("total price is",sumprice)
	 // 输出 total price is 600

	 // 计算多返回值 
	//  area,perimeter :=multiProps(4.6,7.5)
	//  fmt.Printf("Area %f ,Permi is %f",area,perimeter)
	 // 输出
	//  Area 34.500000 ,Permi is 24.200000

	// fmt.Printf("multiply 3 * 7 * 5=%d\n",MultiPly3Num(3,7,5))
	// 输出 multiply 3 * 7 * 5=105
	//  var input int=10
	//  getX2AndX3(input)
	//  fmt.Println(getX2AndX3(input))
	 // 输出 20 30

	 // 非命名的返回值
	//  num1,num2 :=38,23
	//  fmt.Println(calculate(num1,num2))
	//  ret1,ret2,ret3 :=calculate(num1,num2)
	//  fmt.Println(ret1,ret2,ret3)
	  // 输出61 874 15
	  
	  //命名的返回值
	//   num1,num2 :=318,231
	//   ret1,ret2,ret3 :=calculate2(num1,num2)
	//   fmt.Println(ret1,ret2,ret3)

	  // 输出值为  549 73458 87

	//   var x float64=3.2
	//   ret1 :=MySqrt1(x)
	//   fmt.Printf("value is %f",ret1)
	  // 输出 value is 10.240000

	//   y :=9.78
	//   ret1 :=MySqrt2(y)
	//   fmt.Printf("value is %f",ret1)
	  // 输出的值为 value is 95.648400
	//    var d1 int
	//    var f1 float64
	//    d1,_,f1 = BlankTest()
	//    fmt.Printf("The int: %d, the float: %f \n", d1, f1)
	   // 输出的值为 The int: 1, the float: 6.900000 ,第二个入参的值没有输出，选择为_

	//    min,max :=11,23
	//    ret1,ret2 := MinMax(min,max)
	//    fmt.Printf(“max value is %d min value is %d”,ret1,ret2)
	// 输出 11 23 
	// fmt.Printf("Minmium is: %d, Maximum is: %d\n", min, max)
	// 输出的值为
	// Minmium is: 11, Maximum is: 23
	//  n :=0
	//  reply :=&n
	//  fmt.Println(n,reply)
	//  Multiply(10,3,reply)
	//  Multiply("Multiply:",*reply)
  
	 n := 0
	 reply := &n
	 fmt.Println(n,reply)
	 Multiply(10, 5, reply)
	 fmt.Println("Multiply:", *reply)
	 // 输出 
// 	 0 0xc000086000
// Multiply: 50

}