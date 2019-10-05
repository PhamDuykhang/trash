package main

import (
	"fmt"
	"github.com/PhamDuyKhang/trash/aa/myfun"
	"time"
)

func main() {
	a := []string{"a", "f", "fail", "s"}
	for _, d := range a {
		fmt.Println("sending ", d)
		myfun.Send(d)
		fmt.Println("Seen")
	}
	fmt.Println("end of for")
	time.Sleep(2 * time.Minute)
	fmt.Println("end of program")
	return
}
