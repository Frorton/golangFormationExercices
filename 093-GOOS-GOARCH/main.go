package main

import (
	"fmt"
	"runtime"
)

//package runtime has all the stuff about gopath,arch,os,etc...

func main() {
	fmt.Println("OS:", runtime.GOOS)
	fmt.Println("----------")
	fmt.Println("ARCH:", runtime.GOARCH)
}
