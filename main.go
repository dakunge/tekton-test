package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0
	tk := time.NewTicker(time.Second)
	for {
		select {
		case <-tk.C:
			fmt.Println("receive ticker signal", i)
			i++
			//i++
		}
	}
}
