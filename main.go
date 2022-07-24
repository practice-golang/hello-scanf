package main // import "hello-scanf"

import (
	"fmt"
	"time"
)

func main() {
	var a int
	fmt.Println("Please enter any strings")
	fmt.Scanf("%d", &a)
	fmt.Println(a)

	fmt.Println("Done")
	time.Sleep(time.Second * 2)
}
