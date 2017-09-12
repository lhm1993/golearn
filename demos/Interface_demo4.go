/*
	type assertion
*/
package Interface_demo4

import (
	//"bufio"
	"fmt"
	"io"
	"os"
)

var w io.Writer

func Interface_Demo4() {
	w = os.Stdout
	if w, ok := w.(*os.File); ok {
		fmt.Println("ok")
		fmt.Printf("%#v\n", w)
	} else {
		fmt.Println("not ok")
		fmt.Printf("%#v\n", w)
	}
}

func test(x interface{}) string {
	switch x.(type) {
	case nil:
		return "NULL"
	case int:
		return fmt.Sprintf("%d", x)
	case string:
		return x.(string)
	case bool:
		if x.(bool) {
			return "TRUE"
		} else {
			return "FLASE"
		}
	default:
		panic(fmt.Sprintf("worng type %T:%v", x, x))
	}
}
