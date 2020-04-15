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

	fmt.Println(a, b)
}
