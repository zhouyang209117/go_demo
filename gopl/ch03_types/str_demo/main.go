package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s1 := "hel"
	fmt.Println(len(s1))
	s2 := "中国人"
	fmt.Println(len(s2))
	fmt.Println(utf8.RuneCountInString(s2))

	for i := 0; i < len(s2);{
		r, size := utf8.DecodeRuneInString(s2[i:])
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}

	s3 := []rune("中华人民共和国")
	fmt.Printf("%c", s3[0])
	fmt.Printf("%c", s3[1])
	fmt.Println(string(s3))
}
