//interface learning
package Interface_demo3

import (
	"flag"
	"fmt"
)

type Celsius float64
type Fahrenheit float64

//摄氏华氏转换
func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

//实现flag.Value()接口的String()方法
/*
	type Value struct {
		String() string
		Set(string) error
	}
*/

func (c Celsius) String() string {
	fmt.Println(1)
	//var temp int = 21
	//fmt.Println(&c)
	return fmt.Sprintf("%g℃", c)
}
func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g℉", f)
}

type celsiusflag struct{ Celsius }

func (cf *celsiusflag) Set(s string) error {
	fmt.Println(2)
	var unit string
	var value float64
	fmt.Sscanf(s, "%g%s", &value, &unit) //fmt.Sscanf()用于将字符串s按照format解析并将值存放在后面的变量中，需要使用地址
	fmt.Println(unit, value)
	switch unit {
	case "C", "℃":
		cf.Celsius = Celsius(value)
		return nil
	case "F", "℉":
		cf.Celsius = FToC(Fahrenheit(value))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

//flag
func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	//fmt.Println(6)
	f := celsiusflag{value} //celsiusflag 实现了flag.value接口
	fmt.Println(7)
	flag.CommandLine.Var(&f, name, usage) //将自定义的flag转换成可以执行的模式，在这里会去调用String()方法，格式化输入的值
	fmt.Println(8)
	return &f.Celsius
}

//entrance
func Interface_demo3() {
	var temp = Interface_demo3.CelsiusFlag("temp", 20.0, "the temperature")
	fmt.Println(3)
	flag.Parse() //在输入的时候带上参数 -temp的时候才会解析所有传进来的参数，并调用Set()方法
	fmt.Println(4)
	fmt.Println(*temp)
}
