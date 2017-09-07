// Test project Test.go
package Test

import (
	"fmt"
)

//"fmt"

func Test() (x int) {
	//	time.Date(time.Now().Year())
	//fmt.Println(time.Now().Date(), time.Now().Year())
	defer func() { fmt.Printf("%d", x+x) }() //这里的执行顺序需要注意
	//defer 在其所在函数执行完之后才会执行，但是其执行是在返回到主调函数之前执行的，也就是，函数整体空间还存在的时候，
	//执行的defer，如果是匿名函数，需要在其末尾加上圆括号()，否则会报错
	return 22
}
