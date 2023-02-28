package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	name := flag.String("name", "world", "specify the name you want to say hi")
	version := flag.String("version", "1.1", "specify the name you want to say hi")
	// 第 0 个参数是自身执行命令,第二个是参数
	//❯ ./main --name gqa
	//os args is: [./main --name gqa]
	//input parameter is: gqa
	//Hello gqa from Go

	flag.Parse()
	fmt.Println("os args is:", os.Args)
	fmt.Println("input parameter is:", *name)
	fullString := fmt.Sprintf("Hello %s from Go\n version is %s", *name, *version)
	fmt.Println(fullString)
	err, value := DuplicateString(*name)
	if err == nil {
		fmt.Printf(value)
	} else {
		fmt.Println(err)
	}

}

// 判断参数
func DuplicateString(input string) (error, string) {
	if input == "aaa" {
		return fmt.Errorf("aaa is not allowed"), ""
	}
	return nil, input + input
}
