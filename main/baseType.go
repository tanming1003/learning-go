package main

import "fmt"

const x, y int = 1, 2

const (
	SUNDAY = iota
	MONDAY
	TUESDAY
	WEDNESDAY
	THIRSDAY
	FRIDAY
	SATURDAY
)

func main() {
	base()
	strings()
	maps()
	testPanic()
}

func base() {
	var a int
	a = 3
	fmt.Println(a)
	b := 3.5
	fmt.Println(b)
	s := "你好"
	fmt.Println(s)
	var c byte
	c = 255;
	fmt.Println(c)
	n, t := 1, "sss"
	fmt.Println(n, t)
	fmt.Println(x, y)
	fmt.Println(SUNDAY, MONDAY, TUESDAY, WEDNESDAY, THIRSDAY, FRIDAY, SATURDAY)
}

func strings() {
	s := "你好啊，世界，hello,world！"
	fmt.Print("\n")
	for i, c := range s {
		fmt.Printf("%d,%c\n", i, c)
	}
}

func maps() {
	m := map[string]int{"1":1, "2":2}
	for k, v := range m {
		fmt.Println(k, v)
	}
}

func testPanic() {
	defer func() {
		if err := recover(); err != nil {
			println(err.(string))
		}
	}()
	panic("test panic")
}
