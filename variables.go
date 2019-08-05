// 変数名: 1文字目に注意

package main

import "fmt"

func main() {
	msg := "hello world"
	fmt.Println(msg)

    // var a, b int
    // a, b = 10, 15
    // a, b := 10, 15

   /* 
    var (
        c int
        d string
    )
    c = 20
    d = "hoge"
	*/
	
	var a, b, c = 3, 4, "foo"

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("a is of type %T\n", a)
	fmt.Printf("b is of type %T\n", b)
	fmt.Printf("c is of type %T\n", c)
}