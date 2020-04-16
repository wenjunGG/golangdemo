package main

import (
	"fmt"
	"net"
)

//懒汉式写法
var (
	age   int     = 10
	name  string  = "lwj"
	money float64 = 0.12
	sex   func() bool
	sum   struct {
		x int
		b string
	}
)

var i = float64(age) * money

var conn net.Conn
var err error

func changeLocal(num [5]int) {
	num[0] = 55
	fmt.Println("inside function ", num)

}

func changeLocalStr(num [5]string) {
	num[0] = "hhh"
	fmt.Println("changeLocalStr inside function ", num)

}

func main() {
	//age = 10
	//name = "lwj"
	var money float64 = 12.2

	//内部申明
	hp := 10

	//连接tcp
	conn, err = net.Dial("tcp", "127.0.0.1:8080")

	fmt.Println("hello world啦啦啦啦啦")

	fmt.Println(float64(age))
	fmt.Println(name)
	fmt.Println(money)
	fmt.Println(i)
	fmt.Println(hp)
	fmt.Println(money)
	fmt.Println(err)

	//交换变量
	var a int = 100
	var b int = 200
	a = a ^ b
	b = b ^ a
	a = a ^ b

	var c int = 100
	var d int = 200

	fmt.Println(c, d)
	//////////////////////////////////// 数组的概念 //////////////////////////////////////
	//int数组赋值，通过下标方式
	var IntSlice [1]int
	IntSlice[0] = 1

	//int数组赋值，通过下标方式  申明的时候数组必须有长度
	// var IntSlicea []int
	// IntSlicea[0] = 1
	// fmt.Println("IntSlicea", IntSlicea)

	//int 数组初始化赋值 写=
	var intArr = [2]int{1, 2}
	fmt.Println(intArr)

	//string 数组赋值
	var names [2]string
	names[0] = "a"
	names[1] = "b"

	fmt.Println(names)

	//可变长度 int 数组
	abb := [...]int{12, 78, 50}
	fmt.Println(abb)

	//动态赋值数组
	acc := [...]int{2, 5, 6, 3}
	// var add [6]int
	// add = acc
	// fmt.Println(add)
	//程序出错，不能将其赋值

	//正确的赋值，但是不能超过其数组的长度赋值
	var aee = acc
	aee[0] = 89
	//aee[4] = 89
	fmt.Println(aee)

	//int 不会改变 num,因为这是值类型
	num := [...]int{5, 6, 7, 8, 8}
	fmt.Println("before passing to function ", num)
	changeLocal(num) //num is passed by value
	fmt.Println("after passing to function ", num)

	//string 不会改变 numStr,因为这是值类型
	numStr := [...]string{"aa", "bb", "cc", "dd", "ee"}
	fmt.Println("numStr before passing to function ", numStr)
	changeLocalStr(numStr) //num is passed by value
	fmt.Println("numStr after passing to function ", numStr)

	//数组长度
	fmt.Println(len(num))

	//数组循环 for
	for i := 0; i < len(numStr); i++ {
		fmt.Println("numStr:" + numStr[i] + "ss")
	}

	//数组循环 range
	for i, v := range numStr {
		fmt.Println("numstr:", i, v)
	}

	//多维数组
	//初始化赋值 写=
	var darr = [3][2]int{
		{1, 2},
		{2, 5},
		{6, 7}, //这里必须跟一个逗号
	}
	fmt.Println(darr)

	//简化
	daee := [3][2]int{
		{1, 2},
		{2, 5},
		{6, 7},
	}
	fmt.Println(daee)

	//int下标方式赋值 没有为0
	var datt [3][2]int
	datt[0][0] = 4

	fmt.Println(datt)

	//int下标方式赋值 没有为空
	var dayy [3][2]string
	dayy[0][0] = "sss"
	fmt.Println(dayy)

	//多维数组嵌套打印
	for i, v := range daee {
		for _, y := range v {
			fmt.Println(i, y)
		}
	}

	/////////////////////////// 数组切片 ////////////////////////////////////

	//会改变原来的数据 完全赋值 gg 或者 gg[:]
	gg := []int{8, 9, 6}
	ggc := gg
	ggc[0] = 43
	gaa := gg[:]
	fmt.Println("ggc", ggc, gg, gaa)

	//截取赋值
	cc := [5]int{76, 77, 78, 79, 80}
	var ccb []int = cc[1:4] //creates a slice from a[1] to a[3]
	fmt.Println(ccb)

	//初始化的时候可以不定义长度 ，跟... 是一样的效果
	cvv := []int{6, 7, 8} //creates and array and returns a slice reference
	fmt.Println(cvv)

	//申明的时候必须要有长度，不然不能使用，初始化的时候，长度可以不用有
	// var cvvd []int
	// cvvd[0] = 2
	// fmt.Println(cvvd)

	//因此切片可以影响到原来的数组 ，切片属于引用类型操作, 之前的前面数组传递，不会引起数组的改变,因为考虑到作用域的问题
	darrt := [...]int{57, 89, 90, 82, 100, 78, 67, 69, 59}
	dslice := darrt[2:5]
	fmt.Println("darrt array before", darrt)
	for i := range dslice {
		dslice[i]++
	}
	fmt.Println("darrt array after", darrt)

	//当多个切片共享同一个数组时，对每一个切片的修改都将会反映到这个数组中
	numa := [3]int{78, 79, 80}
	nums1 := numa[:] //creates a slice which contains all elements of the array
	nums2 := numa[:]
	fmt.Println("array before change 1", numa)
	nums1[0] = 100
	fmt.Println("array after modification to slice nums1", numa)
	nums2[1] = 101
	fmt.Println("array after modification to slice nums2", numa)

}
