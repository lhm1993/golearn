// Interface project Interface.go
package Interface

import (
	"fmt"
)

type ByteCount int

var i = 0

func (c *ByteCount) Write(p []byte) (int, error) {
	*c += ByteCount(len(p))
	i++
	fmt.Println(i, *c)
	return int(*c), nil
}

func InterfaceDemo1() {
	var c ByteCount
	b, err := c.Write([]byte("kello"))
	fmt.Fprintf(&c, "hello,%s\n", "1222")
	fmt.Println(b, c, err)
}
