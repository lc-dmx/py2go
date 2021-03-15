package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/lc-dmx/py2go/packages/os/utils"
)

func main() {
	fmt.Println("----------os package----------")
	// os 包
	for idx, arg := range os.Args[1:] {
		fmt.Println(utils.NumToStr(idx+1), "参数", arg)
	}
	fmt.Println()

	fmt.Println("----------flag package----------")
	// flag 包
	var (
		name    = flag.String("name", "", "名称")
		age     = flag.Uint64("age", 0, "年龄")
		height  = flag.Uint64("height", 0, "身高")
		weight  = flag.Uint64("weight", 0, "体重")
		married = flag.Bool("married", false, "是否已婚")
	)
	flag.Parse()
	format := "%-10s:%v\n"
	fmt.Printf(format, "name", *name)
	fmt.Printf(format, "age", *age)
	fmt.Printf(format, "height", *height)
	fmt.Printf(format, "weight", *weight)
	fmt.Printf(format, "married", *married)
}
