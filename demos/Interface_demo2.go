//Interface_demo2

package Interface_demo2

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func Interface_Demo2() {
	input := "122\n2343 hkafh wuwu\n dds"
	scanner := bufio.NewScanner(strings.NewReader(input))

	//这里的bufio.ScanWords & bufio.ScanLines 并不是作为一个独立的方法被调用的，而是作为一种分割方式被调用
	scanner.Split(bufio.ScanWords)
	count := 0
	//scanner.Scan()将scanner移动到下一个可读的位置，这个取决于scanner的类型，
	for scanner.Scan() {
		count++
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
	fmt.Println(count)

}
