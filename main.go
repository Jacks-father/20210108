package main

import (
	_ "20210208/routers"
	"fmt"
	"github.com/astaxie/beego"
)

func main() {
	fmt.Println("hello world")
	fmt.Println("写入内容")
	beego.Run()
}

