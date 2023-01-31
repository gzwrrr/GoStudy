package main

import "fmt"
import "errors"

func readConf(name string) (err error) {
	if name == "config.ini" {
		return nil
	} else {
		return errors.New("读取文件错误")
	}
}

func main() {
	err := readConf("cogfig.ini")
	fmt.Println("程序继续执行...")
	
	err = readConf("config.in");
	if err != nil {
		panic(err)
	}
	fmt.Println("程序继续执行...")


}
