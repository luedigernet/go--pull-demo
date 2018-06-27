package main

import (
	"github.com/luedigernet/go-pull-demo/pkg/lib1"
	"github.com/luedigernet/go-pull-demo/pkg/lib2"
	"fmt"
)

func main(){
	fmt.Println(lib1.SayHello())
	fmt.Println(lib2.SayHello("Reinhard"))

	fmt.Println("demo of a pull request")
}
