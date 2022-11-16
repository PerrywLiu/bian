package main

import (
	"bian/src"
	"fmt"
	"time"
)

func main() {
	fmt.Println("biAn go")
	//src.GetSystemState()
	src.GetAllCoinInfo()

	fmt.Println((time.Now().Unix()))
}
