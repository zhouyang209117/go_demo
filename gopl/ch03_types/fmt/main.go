package main

import "fmt"

type Student struct {
	StuNum int64
	StuName string
}
func main() {
	//%v	the value in a default format
	//when printing structs, the plus flag (%+v) adds field names 没有明白怎么回事
	//%#v	a Go-syntax representation of the value
	//%T	a Go-syntax representation of the type of the value
	//%%	a literal percent sign; consumes no value
	stu := Student{
		StuNum: int64(1),
		StuName: "name1",
	}
	fmt.Printf("stu=%v\n", stu)
	fmt.Printf("stu=%#v\n", stu)
	fmt.Printf("stu=%T\n", stu)
	//fmt.Printf("stu=%%\n", stu) //没有明白怎么回事

	fmt.Printf("result=%t\n", false)
	fmt.Printf("%8b\n", 10)
	fmt.Printf("%08b\n", 10)
	fmt.Printf("%c\n", '中') //显示unicode字符
	fmt.Printf("%o\n", 10)

	fmt.Printf("%q\n", 0x4E2D) //没看懂是什么意思

	fmt.Printf("%x\n", 10)
	fmt.Printf("%X\n", 11)
	fmt.Printf("%U", '中')
}
